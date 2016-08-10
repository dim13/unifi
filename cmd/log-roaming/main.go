// Copyright (c) 2014 Dimitri Sokolyuk. All rights reserved.
// Use of this source code is governed by ISC-style license
// that can be found in the LICENSE file.

// log stations as they roams
package main

import (
	"flag"
	"log"
	"log/syslog"
	"os"
	"time"

	"github.com/dim13/unifi"
)

type roaming struct {
	Name    string
	IP      string
	Ap      string
	Channel int
	Essid   string
}

type roamMap map[string]roaming

var stamap roamMap

var (
	host    = flag.String("host", "unifi", "Controller hostname")
	user    = flag.String("user", "admin", "Controller username")
	pass    = flag.String("pass", "unifi", "Controller password")
	version = flag.Int("version", 2, "Controller base version")
	siteid  = flag.String("siteid", "default", "Site ID, UniFi v3 only")
	delay   = flag.Duration("delay", 5*time.Second, "delay")
)

func main() {
	flag.Parse()
	u, err := unifi.Login(*user, *pass, *host, *siteid, *version)
	if err != nil {
		log.Fatal(err)
	}

	defer u.Logout()
	apsmap, err := u.ApsMap()
	if err != nil {
		log.Fatal(err)
	}

	elog := log.New(os.Stderr, "", log.Ltime)
	slog, err := syslog.NewLogger(syslog.LOG_NOTICE|syslog.LOG_DAEMON, 0)
	if err != nil {
		log.Fatal(err)
	}
	logger := []*log.Logger{elog, slog}

	ticker := time.NewTicker(*delay)
	defer ticker.Stop()
	for range ticker.C {
		newmap := make(roamMap)
		sta, err := u.Sta()
		if err != nil {
			continue
		}
		for _, s := range sta {
			newmap[s.Mac] = roaming{
				Name:    s.Name(),
				IP:      s.IP,
				Ap:      apsmap[s.ApMac].Name,
				Channel: s.Channel,
				Essid:   s.EssID,
			}
		}
		for k, v := range newmap {
			if z, ok := stamap[k]; !ok {
				elog.SetPrefix(" → ")
				for _, l := range logger {
					l.Printf("%s appears on %s/%d %s/%s\n",
						v.Name, v.Ap, v.Channel, v.Essid, v.IP)
				}
			} else if z != v {
				elog.SetPrefix(" ↔ ")
				for _, l := range logger {
					l.Printf("%s roams from %s/%d %s/%s to %s/%d %s/%s\n",
						v.Name,
						z.Ap, z.Channel, z.Essid, z.IP,
						v.Ap, v.Channel, v.Essid, v.IP)
				}
			}
			delete(stamap, k)
		}
		for _, v := range stamap {
			elog.SetPrefix(" ← ")
			for _, l := range logger {
				l.Printf("%s vanishes from %s/%d %s/%s\n",
					v.Name, v.Ap, v.Channel, v.Essid, v.IP)
			}
		}
		stamap = newmap
	}
}
