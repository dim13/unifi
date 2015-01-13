// list associated stations
package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/dim13/unifi"
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
		log.Fatal("Login returned error: ", err)
	}
	defer u.Logout()

	aps, err := u.ApsMap()
	if err != nil {
		log.Fatal(err)
	}
	sta, err := u.Sta()
	if err != nil {
		log.Fatal(err)
	}

	for _, s := range sta {
		fmt.Printf("%24s%3s%12s%3d%5d%5d%5d%8s/%-3d%16s%4s\n",
			s.Name(), s.Radio, s.EssID, s.RoamCount, s.Signal, s.Noise, s.Rssi,
			aps[s.ApMac].Name, s.Channel, s.IP, aps[s.ApMac].Model)
	}
}
