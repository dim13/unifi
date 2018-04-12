// Copyright (c) 2014 The unifi Authors. All rights reserved.
// Use of this source code is governed by ISC-style license
// that can be found in the LICENSE file.

// Example command list-usws
// List USWs of a given site
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
	siteid  = flag.String("siteid", "default", "Sitename or description")
)

func main() {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 3, ' ', 0)
	defer w.Flush()

	flag.Parse()

	u, err := unifi.Login(*user, *pass, *host, *port, *siteid, *version)
	if err != nil {
		log.Fatalln("Login returned error: ", err)
		return
	}
	defer u.Logout()

	site, err := u.Site(*siteid)
	if err != nil {
		log.Fatal(err)
	}

	aps, err := u.USWs(site)
	if err != nil {
		log.Fatalln(err)
		return
	}

	// Output headline
	fmt.Fprintln(w, "DeviceName\tIP\tMac\tModelName\tVersion\tStatus\tNumberOfClients\tTxBytes\tRxBytes")

	for _, s := range aps {
		p := []string{
			s.DeviceName(), // Serial if not specified
			s.IP,
			s.Mac,
			s.ModelName(),
			s.Version,
			s.Status(),
			strconv.Itoa(s.NumSta),
			strings.TrimSpace(unifi.Bytes(s.TxBytes).String()),
			strings.TrimSpace(unifi.Bytes(s.RxBytes).String()),
		}
		fmt.Fprintln(w, strings.Join(p, "\t"))
	}
}
