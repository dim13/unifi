// list associated stations
package main

import (
	"flag"
	"fmt"
	"github.com/dim13/unifi"
	"log"
)

var (
	host    = flag.String("host", "unifi", "Controller hostname")
	user    = flag.String("user", "admin", "Controller username")
	pass    = flag.String("pass", "unifi", "Controller password")
	version = flag.Int("version", 2, "Controller base version")
	siteid  = flag.String("siteid", "default", "Site ID, UniFi v3 only")
)

func main() {
	flag.Parse()
	u, err := unifi.Login(*user, *pass, *host, *siteid, *version)
	if err != nil {
		log.Fatalln("Login returned error: ", err)
	}

	defer u.Logout()

	aps, err := u.Aps()
	if err != nil {
		log.Fatalln(err)
	}

	for _, s := range aps {
		fmt.Println(s.Mac, s.ModelName(), s.Name, s.Status())
	}
}
