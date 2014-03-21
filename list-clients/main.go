package main

import (
	"flag"
	"fmt"
	"github.com/dim13/unifi"
)

func main() {
	user := flag.String("user", "admin", "User")
	pass := flag.String("pass", "unifi", "Password")
	url := flag.String("url", "unifi", "URL")
	flag.Parse()

	u := unifi.Login(*user, *pass, *url)
	defer u.Logout()

	aps := u.GetApsMap()
	for _, s := range u.GetSta() {
		fmt.Printf("%s at %s/%d\n", s.GetName(), aps[s.Ap_mac].Name, s.Channel)
	}
}
