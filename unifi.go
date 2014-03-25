/* in memory of http://memegenerator.net/instance/37313316 and https://xkcd.com/927/ */

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

type Meta struct {
	Rc string
}

func Login(user, pass, host string) *Unifi {
	u := new(Unifi)
	u.Login(user, pass, host)
	return u
}

func (u *Unifi) Login(user, pass, host string) {
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
	fmt.Printf("%s %s\n", res, mgr[0])
	/* FIXME */
}

func (u *Unifi) parse(cmd string, v interface{}) {
	body := u.apicmd(cmd)
	if err := json.Unmarshal(body, &v); err != nil {
		log.Fatal(err)
	}
	m := reflect.ValueOf(v).Elem().FieldByName("Meta").Interface().(Meta)
	if m.Rc != "ok" {
		log.Fatal("not ok")
	}
}

func (u *Unifi) GetAps() []Aps {
	var response struct {
		Data []Aps
		Meta Meta
	}
	u.parse("stat/device", &response)
	return response.Data
}

func (u *Unifi) GetApsMap() map[string]Aps {
	m := make(map[string]Aps)
	for _, a := range u.GetAps() {
		m[a.Mac] = a
	}
	return m
}

func (u *Unifi) GetSta() []Sta {
	var response struct {
		Data []Sta
		Meta Meta
	}
	u.parse("stat/sta", &response)
	return response.Data
}

func (u *Unifi) GetStaMap() map[string]Sta {
	m := make(map[string]Sta)
	for _, s := range u.GetSta() {
		m[s.Mac] = s
	}
	return m
}

func (u *Unifi) GetUsers() []User {
	var response struct {
		Data []User
		Meta Meta
	}
	u.parse("list/user", &response)
	return response.Data
}

func (u *Unifi) GetUserGroups() {
	body := u.apicmd("list/usergroup")
	fmt.Printf("%s\n", body)
}

func (u *Unifi) GetWlanConf() []WlanConf {
	var response struct {
		Data []WlanConf
		Meta Meta
	}
	u.parse("list/wlanconf", &response)
	return response.Data
}

func (u *Unifi) BlockSta(mac string) {
	u.maccmd(mac, "block-sta")
}

func (u *Unifi) UnBlockSta(mac string) {
	u.maccmd(mac, "unblock-sta")
}

func (u *Unifi) DisconnectSta(mac string) {
	u.maccmd(mac, "kick-sta")
}

func (u *Unifi) RestartAP(mac string) {
	u.maccmd(mac, "restart", "devmgr")
}

func (u *Unifi) RestartAPbyName(name string) {
}

func (u *Unifi) CreateBackup() {
}

func (u *Unifi) GetBackup() {
}

func (u *Unifi) Restart(mac string) {
	type cmd struct {
		Mac string
		Cmd string
	}

	c := cmd{Mac: mac, Cmd: "restart"}
	fmt.Println(c)
}
