// list associated stations
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

	for _, s := range u.Aps() {
		fmt.Println(s.Mac, s.ModelName(), s.Name)
	}
}
