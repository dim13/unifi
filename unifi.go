// Copyright (c) 2014 The unifi Authors. All rights reserved.
// Use of this source code is governed by ISC-style license
// that can be found in the LICENSE file.

package unifi

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"reflect"
)

var (
	ErrLoginFirst = errors.New("login first")
)

type Unifi struct {
	client  *http.Client
	baseURL string
	apiURL  string
	version int
}

type meta struct {
	Rc string
}

type login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Initializes a session.  If site != "", it's to a V3 controller.
func Login(user, pass, host, port, site string, version int) (*Unifi, error) {
	u := new(Unifi)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	cj, _ := cookiejar.New(nil)
	u.client = &http.Client{
		Transport: tr,
		Jar:       cj,
	}
	u.baseURL = "https://" + host + ":" + port + "/"
	u.version = version
	if u.version >= 4 {
		l := new(login)
		l.Username = user
		l.Password = pass
		j, err := json.Marshal(l)
		if err != nil {
			return nil, err
		}
		r := bytes.NewReader(j)
		if _, err := u.client.Post(u.baseURL+"api/login", "application/json", r); err != nil {
			fmt.Println(err)
			return nil, err
		}
	} else {
		val := url.Values{
			"login":    {"login"},
			"username": {user},
			"password": {pass},
		}
		if _, err := u.client.PostForm(u.baseURL+"login", val); err != nil {
			fmt.Println(err)
			return nil, err
		}
	}
	if u.version >= 3 {
		// Try to resolve the site by description (i.e. user friendly name)
		sitename, err := u.siteNameByDesc(site)
		if err == nil {
			site = sitename
		}

		u.apiURL = u.baseURL + "api/s/" + site + "/"
	} else {
		u.apiURL = u.baseURL + "api/"
	}
	return u, nil
}

// Terminates a session
func (u *Unifi) Logout() {
	u.client.Get(u.baseURL + "logout")
}

func (u *Unifi) apicmd(cmd string) ([]byte, error) {
	// The sites command is global, the others are site specific
	var url string
	if cmd == "api/self/sites" {
		url = u.baseURL + cmd
	} else {
		url = u.apiURL + cmd
	}

	resp, err := u.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (u *Unifi) apicmdPut(cmd string, data interface{}) error {

	url := u.apiURL + cmd

	j, err := json.Marshal(data)
	if err != nil {
		return err
	}

	r := bytes.NewReader(j)

	req, err := http.NewRequest(http.MethodPut, url, r)

	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")

	// Send request
	resp, err := u.client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if err != nil {
		return err
	}
	fmt.Println(resp.Status)
	if resp.StatusCode != 200 {
		return fmt.Errorf(resp.Status)
	}

	return nil
}

type command struct {
	Mac     string `json:"mac"`
	Cmd     string `json:"cmd"`
	Minutes int    `json:"minutes,omitempty"`
}

func (u *Unifi) devcmd(mac, cmd string) error {
	return u.maccmd("devmgr", command{Mac: mac, Cmd: cmd})
}

func (u *Unifi) stacmd(mac, cmd string, min ...int) error {
	minutes := 0
	if len(min) > 0 {
		minutes = min[0]
	}
	return u.maccmd("stamgr", command{Mac: mac, Cmd: cmd, Minutes: minutes})
}

func (u *Unifi) maccmd(mgr string, args interface{}) error {
	param, err := json.Marshal(args)
	if err != nil {
		return err
	}
	val := url.Values{"json": {string(param)}}
	_, err = u.client.PostForm(u.apiURL+"cmd/"+mgr, val)
	return err
}

func (u *Unifi) parse(cmd string, v interface{}) error {
	body, err := u.apicmd(cmd)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, &v); err != nil {
		log.Println(cmd)
		log.Println(string(body))
		return err
	}
	m := reflect.ValueOf(v).Elem().FieldByName("Meta").Interface().(meta)
	if m.Rc != "ok" {
		return fmt.Errorf("%s returned result code: %s", cmd, m.Rc)
	}
	return nil
}

// Returns a slice of access points
func (u *Unifi) Aps() ([]UAP, error) {

	var uaps []UAP

	devices, err := u.Devices()
	if err != nil {
		return uaps, err
	}

	for _, d := range devices {
		switch dev := d.(type) {
		case UAP:
			uaps = append(uaps, dev)
		}
	}
	return uaps, err
}

// Returns a map of access points with mac as a key
func (u *Unifi) ApsMap() (UAPmap, error) {
	aps, err := u.Aps()
	if err != nil {
		return nil, err
	}
	m := make(UAPmap)
	for _, a := range aps {
		m[a.Mac] = a
	}
	return m, nil
}

// Returns a slice of devices
func (u *Unifi) Devices() ([]interface{}, error) {
	// Delay parsing until we know the type
	//var rawDevices []json.RawMessage

	// Devices
	var genericDevices []interface{}
	var response struct {
		Data []json.RawMessage
		Meta meta
	}
	err := u.parse("stat/device", &response)

	// Get the device list
	//err := parse(&raw)

	// Now do the magic
	for _, device := range response.Data {

		// unmarshal into a map to check the "type" field
		var obj map[string]interface{}
		err := json.Unmarshal(device, &obj)
		if err != nil {
			fmt.Println("Raw JSON Unmarshaling failed") // TODO Remove and handle correctly
		}

		devicetype := ""
		if t, ok := obj["type"].(string); ok {
			devicetype = t
		}
		// unmarshal again into the correct type
		switch devicetype {
		case "uap":
			var uap UAP
			err = json.Unmarshal(device, &uap)

			if err == nil {
				uap.u = u // Set API pointer
				genericDevices = append(genericDevices, uap)
			} else {
				fmt.Println(err) // TODO Handle correctly
			}

		case "usw":
			var usw USW
			err = json.Unmarshal(device, &usw)
			if err == nil {
				usw.u = u // Set API pointer
				genericDevices = append(genericDevices, usw)
			}

		default:
			fmt.Println("Unknown device")
		}

	}
	return genericDevices, err
}

// Returns a map of access points with mac as a key
func (u *Unifi) DeviceMap() (DeviceMap, error) {
	devices, err := u.Devices()
	if err != nil {
		return nil, err
	}
	m := make(DeviceMap)
	for _, d := range devices {
		switch dev := d.(type) {
		case UAP:
			m[dev.Mac] = dev

		case USW:
			m[dev.Mac] = dev
		}

	}
	return m, nil
}

// Returns a USW pointer for USW with a given name
func (u *Unifi) USW(name string) (*USW, error) {
	devices, err := u.Devices()
	if err != nil {
		return nil, err
	}
	for _, d := range devices {
		switch dev := d.(type) {
		case USW:
			if name == dev.DeviceName() {
				return &dev, nil
			}
		}
	}
	return nil, fmt.Errorf("No device with name: %s", name)
}

// Returns a slice of stations
func (u *Unifi) Sta() ([]Sta, error) {
	var response struct {
		Data []Sta
		Meta meta
	}
	err := u.parse("stat/sta", &response)
	for i := range response.Data {
		response.Data[i].u = u
	}
	return response.Data, err
}

// Returns a map of stations with MAC as a key
func (u *Unifi) StaMap() (StaMap, error) {
	sta, err := u.Sta()
	if err != nil {
		return nil, err
	}
	m := make(StaMap)
	for _, s := range sta {
		m[s.Mac] = s
	}
	return m, nil
}

// Returns a slice of known users
func (u *Unifi) Users() ([]User, error) {
	var response struct {
		Data []User
		Meta meta
	}
	err := u.parse("list/user", &response)
	return response.Data, err
}

// Returns a slice of known networks
func (u *Unifi) Networks() ([]Network, error) {
	var response struct {
		Data []Network
		Meta meta
	}
	err := u.parse("rest/networkconf", &response)
	return response.Data, err
}

// Returns a map of networkconfigs with ID as key
func (u *Unifi) NetworkMap() (NetworkMap, error) {
	networks, err := u.Networks()
	if err != nil {
		return nil, err
	}
	m := make(NetworkMap)
	for _, n := range networks {
		m[n.ID] = n
	}
	return m, nil
}

// Returns a slice of known portconfigs
func (u *Unifi) PortProfiles() ([]PortProfile, error) {
	var response struct {
		Data []PortProfile
		Meta meta
	}
	err := u.parse("list/portconf", &response)
	return response.Data, err
}

// Returns a map of networkconfigs with ID as key
func (u *Unifi) PortProfileMap() (PortprofileMap, error) {
	profiles, err := u.PortProfiles()
	if err != nil {
		return nil, err
	}
	m := make(PortprofileMap)
	for _, p := range profiles {
		m[p.ID] = p
	}
	return m, nil
}

// Returns a map of networkconfigs with ID as key
func (u *Unifi) PortProfile(name string) (*PortProfile, error) {

	profiles, err := u.PortProfiles()
	if err != nil {
		return nil, err
	}

	for _, p := range profiles {
		if p.Name == name {
			return &p, err
		}
	}
	return nil, fmt.Errorf("No Profile with name: %s", name)
}

// Sets the portoverrides of a given device
func (u *Unifi) SetPortoverrides(deviceid string, o []PortOverride) error {

	cmd := fmt.Sprintf("rest/device/%s", deviceid)

	// Create a map with port_overrides as key and a slice of overrides as value
	m := make(map[string][]PortOverride)
	m["port_overrides"] = o
	err := u.apicmdPut(cmd, m)

	return err
}

// Returns user groups
func (u *Unifi) UserGroups() ([]UserGroup, error) {
	var response struct {
		Data []UserGroup
		Meta meta
	}
	err := u.parse("list/usergroup", &response)
	return response.Data, err
}

// Returns a slice of all sites
func (u *Unifi) Sites() ([]Site, error) {
	var response struct {
		Data []Site
		Meta meta
	}
	err := u.parse("api/self/sites", &response)
	return response.Data, err
}

// Returns the name (id) of a site, searched by its description (user friendly name)
func (u *Unifi) siteNameByDesc(desc string) (string, error) {
	sites, err := u.Sites()
	if err != nil {
		return "", err
	}

	for _, s := range sites {
		if s.Desc == desc {
			return s.Name, nil
		}
	}

	return "", errors.New("No site with desc: " + desc)
}

// Returns a Wlan config
func (u *Unifi) WlanConf() ([]WlanConf, error) {
	var response struct {
		Data []WlanConf
		Meta meta
	}
	err := u.parse("list/wlanconf", &response)
	return response.Data, err
}

func (u *Unifi) CreateBackup() error {
	return nil
}

func (u *Unifi) GetBackup() error {
	return nil
}
