// logs stations as they roams
package main

import (
	"flag"
	"github.com/dim13/unifi"
	"log"
	"time"
)

type roaming struct {
	Name    string
	Ip      string
	Ap      string
	Channel int
	Essid   string
}

type roamMap map[string]roaming

var stamap roamMap

var (
	user  = flag.String("user", "admin", "User")
	pass  = flag.String("pass", "unifi", "Password")
	url   = flag.String("url", "unifi", "URL")
	delay = flag.Int("delay", 5, "delay")
)

func main() {
	flag.Parse()
	u := unifi.Login(*user, *pass, *url)
	defer u.Logout()

	apsmap := u.ApsMap()

	for {
		newmap := make(roamMap)
		for _, s := range u.Sta() {
			newmap[s.Mac] = roaming{
				Name:    s.Name(),
				Ip:      s.Ip,
				Ap:      apsmap[s.Ap_mac].Name,
				Channel: s.Channel,
				Essid:   s.Essid,
			}
		}
		for k, v := range newmap {
			if z, ok := stamap[k]; !ok {
				log.Printf("%s appears on %s/%d (%s/%s)\n",
					v.Name, v.Ap, v.Channel, v.Essid, v.Ip)
			} else if z != v {
				log.Printf("%s roams %s/%d (%s/%s) -> %s/%d (%s/%s)\n",
					v.Name,
					z.Ap, z.Channel, z.Essid, z.Ip,
					v.Ap, v.Channel, v.Essid, v.Ip)
			}
			delete(stamap, k)
		}
		for _, v := range stamap {
			log.Printf("%s vanishes from %s/%d (%s/%s)\n",
				v.Name, v.Ap, v.Channel, v.Essid, v.Ip)
		}
		stamap = newmap
		time.Sleep(time.Duration(*delay) * time.Second)
	}
}
