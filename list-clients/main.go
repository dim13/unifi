package main

import (
	"flag"
	"fmt"
	"github.com/dim13/unifi"
)

var (
	user = flag.String("user", "admin", "User")
	pass = flag.String("pass", "unifi", "Password")
	url  = flag.String("url", "unifi", "URL")
)

func main() {
	flag.Parse()
	u := unifi.Login(*user, *pass, *url)
	defer u.Logout()

	aps := u.ApsMap()
	for _, s := range u.Sta() {
		fmt.Printf("%24s%3s%12s%3d%5d%5d%5d%8s/%-3d%16s%4s\n",
			s.Name(), s.Radio, s.Essid, s.Roam_count, s.Signal, s.Noise, s.Rssi,
			aps[s.Ap_mac].Name, s.Channel, s.Ip, aps[s.Ap_mac].Model)
	}
}
