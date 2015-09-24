// Copyright (c) 2014 Dimitri Sokolyuk. All rights reserved.
// Use of this source code is governed by ISC-style license
// that can be found in the LICENSE file.

// list associated stations
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"

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
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 3, ' ', 0)
	defer w.Flush()

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
		a := aps[s.ApMac]
		p := []string{
			s.Name(),
			s.Radio,
			strconv.Itoa(s.Channel),
			s.EssID,
			strconv.Itoa(s.RoamCount),
			strconv.Itoa(s.Signal),
			strconv.Itoa(s.Noise),
			strconv.Itoa(s.Rssi),
			a.Name,
			s.IP,
			a.ModelName(),
		}
		fmt.Fprintln(w, strings.Join(p, "\t"))
	}
}
