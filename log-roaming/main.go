package main

import (
	"flag"
	"github.com/dim13/unifi"
	"log"
	"time"
)

type Roaming struct {
	Name    string
	Ip      string
	Ap      string
	Channel int
	Essid   string
}

type RoamMap map[string]Roaming

func main() {
	var stamap RoamMap

	user := flag.String("user", "admin", "User")
	pass := flag.String("pass", "unifi", "Password")
	url := flag.String("url", "unifi", "URL")
	flag.Parse()

	u := unifi.Login(*user, *pass, *url)
	defer u.Logout()

	apsmap := u.GetApsMap()

	for {
		newmap := make(RoamMap)
		for _, s := range u.GetSta() {
			newmap[s.Mac] = Roaming{
				Name:    s.GetName(),
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
		time.Sleep(5 * time.Second)
	}
}
