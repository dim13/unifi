// in memory of http://memegenerator.net/instance/37313316 and https://xkcd.com/927/

package unifi

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"reflect"
)

type Unifi struct {
	client  *http.Client
	baseURL string
	apiURL  string
}

type meta struct {
	Rc string
}

// Initializes a session.  If site != "", it's to a V3 controller.
func Login(user, pass, host, site string, version int) (*Unifi, error) {
	u := new(Unifi)
	err := u.login(user, pass, host, site, version)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (u *Unifi) login(user, pass, host, site string, version int) error {
	val := url.Values{
		"login":    {"login"},
		"username": {user},
		"password": {pass},
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	cj, _ := cookiejar.New(nil)
	u.client = &http.Client{
		Transport: tr,
		Jar:       cj,
	}
	u.baseURL = "https://" + host + ":8443/"
	if _, err := u.client.PostForm(u.baseURL+"login", val); err != nil {
		return err
	}
	if version >= 3 {
		u.apiURL = u.baseURL + "api/s/" + site + "/"
	} else {
		u.apiURL = u.baseURL + "api/"
	}
	return nil
}

// Terminates a session
func (u *Unifi) Logout() {
	u.client.Get(u.baseURL + "logout")
}

func (u *Unifi) apicmd(cmd string) ([]byte, error) {
	resp, err := u.client.Get(u.apiURL + cmd)
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

func (u *Unifi) devcmd(mac, cmd string) error {
	return u.maccmd(mac, cmd, "devmgr")
}

func (u *Unifi) stacmd(mac, cmd string) error {
	return u.maccmd(mac, cmd, "stamgr")
}

func (u *Unifi) maccmd(mac, cmd, mgr string) error {
	type command struct {
		Mac string `json:"mac"`
		Cmd string `json:"cmd"`
	}
	param, err := json.Marshal(command{Mac: mac, Cmd: cmd})
	if err != nil {
		return err
	}
	val := url.Values{"json": {string(param)}}
	if _, err := u.client.PostForm(u.apiURL+"cmd/"+mgr, val); err != nil {
		return err
	}
	return nil
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
func (u *Unifi) Aps() ([]Aps, error) {
	var response struct {
		Data []Aps
		Meta meta
	}
	err := u.parse("stat/device", &response)
	for i := range response.Data {
		response.Data[i].u = u
	}
	return response.Data, err
}

// Returns a map of access points with mac as a key
func (u *Unifi) ApsMap() (ApsMap, error) {
	aps, err := u.Aps()
	if err != nil {
		return nil, err
	}
	m := make(ApsMap)
	for _, a := range aps {
		m[a.Mac] = a
	}
	return m, nil
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

// Returns user groups
func (u *Unifi) UserGroups() ([]UserGroup, error) {
	var response struct {
		Data []UserGroup
		Meta meta
	}
	err := u.parse("list/usergroup", &response)
	return response.Data, err
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
