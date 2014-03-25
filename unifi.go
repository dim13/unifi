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
	client *http.Client
	host   string
}

type meta struct {
	Rc string
}

// Initializes a session
func Login(user, pass, host string) *Unifi {
	u := new(Unifi)
	u.login(user, pass, host)
	return u
}

func (u *Unifi) login(user, pass, host string) {
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
	u.host = "https://" + host + ":8443/"
	if _, err := u.client.PostForm(u.host+"login", val); err != nil {
		log.Fatal(err)
	}
}

// Terminates a session
func (u *Unifi) Logout() {
	u.client.Get(u.host + "logout")
}

func (u *Unifi) apicmd(cmd string) []byte {
	resp, err := u.client.Get(u.host + "api/" + cmd)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body
}

func (u *Unifi) maccmd(mac, cmd string, mgr ...string) {
	type Command struct {
		Mac string
		Cmd string
	}
	if mgr == nil {
		mgr = append(mgr, "stamgr")
	}
	res, err := json.Marshal(Command{Mac: mac, Cmd: cmd})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s: %s %s\n", u.host, res, mgr[0])
	/* FIXME */
}

func (u *Unifi) parse(cmd string, v interface{}) {
	body := u.apicmd(cmd)
	if err := json.Unmarshal(body, &v); err != nil {
		log.Fatal(err)
	}
	m := reflect.ValueOf(v).Elem().FieldByName("Meta").Interface().(meta)
	if m.Rc != "ok" {
		log.Fatal("not ok")
	}
}

// Returns a slice of access points
func (u *Unifi) Aps() []Aps {
	var response struct {
		Data []Aps
		Meta meta
	}
	u.parse("stat/device", &response)
	for i, _ := range response.Data {
		response.Data[i].u = u
	}
	return response.Data
}

// Returns a map of access points with mac as a key
func (u *Unifi) ApsMap() ApsMap {
	m := make(ApsMap)
	for _, a := range u.Aps() {
		m[a.Mac] = a
	}
	return m
}

// Returns a slice of stations
func (u *Unifi) Sta() []Sta {
	var response struct {
		Data []Sta
		Meta meta
	}
	u.parse("stat/sta", &response)
	for i, _ := range response.Data {
		response.Data[i].u = u
	}
	return response.Data
}

// Returns a map of stations with MAC as a key
func (u *Unifi) StaMap() StaMap {
	m := make(StaMap)
	for _, s := range u.Sta() {
		m[s.Mac] = s
	}
	return m
}

// Returns a slice of known users
func (u *Unifi) Users() []User {
	var response struct {
		Data []User
		Meta meta
	}
	u.parse("list/user", &response)
	return response.Data
}

func (u *Unifi) UserGroups() {
	body := u.apicmd("list/usergroup")
	fmt.Printf("%s\n", body)
}

// Returns a Wlan config
func (u *Unifi) WlanConf() []WlanConf {
	var response struct {
		Data []WlanConf
		Meta meta
	}
	u.parse("list/wlanconf", &response)
	return response.Data
}

func (u *Unifi) CreateBackup() {
}

func (u *Unifi) GetBackup() {
}
