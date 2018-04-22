// Copyright (c) 2014 Dimitri Sokolyuk. All rights reserved.
// Use of this source code is governed by ISC-style license
// that can be found in the LICENSE file.

// Example command list-aps
// List APs of a given site
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
	version = flag.Int("version", 5, "Controller base version")
	port    = flag.String("port", "8443", "Controller port")
	siteID  = flag.String("siteid", "default", "Sitename or description")
)

func main() {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 3, ' ', 0)
	defer w.Flush()

	flag.Parse()

	u, err := unifi.Login(*user, *pass, *host, *port, *siteID, *version)
	if err != nil {
		log.Fatalln("Login returned error: ", err)
		return
	}
	defer u.Logout()

	site, err := u.Site(*siteID)
	if err != nil {
		log.Fatal(err)
	}

	aps, err := u.UAPs(site)
	if err != nil {
		log.Fatalln(err)
		return
	}

	// Output headline
	fmt.Fprintln(w, "DeviceName\tIP\tMac\tModelName\tVersion\tStatus\tNumberOfUsers\tNumerOfGuests\tTxBytes\tRxBytes")

	for _, a := range aps {
		p := []string{
			a.DeviceName(),
			a.IP,
			a.Mac,
			a.ModelName(),
			a.Version,
			a.State.String(),
			strconv.Itoa(a.NumSta),
			strconv.Itoa(a.GuestNumSta),
			unifi.Bytes(a.TxBytes).String(),
			unifi.Bytes(a.RxBytes).String(),
		}
		fmt.Fprintln(w, strings.Join(p, "\t"))
	}
}
