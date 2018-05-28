// Copyright (c) 2014 Dimitri Sokolyuk. All rights reserved.
// Use of this source code is governed by ISC-style license
// that can be found in the LICENSE file.

// Example command list-events
// List Events of a given site
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"text/tabwriter"
	"time"

	"github.com/dim13/unifi"
)

var (
	host    = flag.String("host", "unifi", "Controller hostname")
	user    = flag.String("user", "admin", "Controller username")
	pass    = flag.String("pass", "unifi", "Controller password")
	version = flag.Int("version", 5, "Controller base version")
	port    = flag.String("port", "8443", "Controller port")
	siteID  = flag.String("siteid", "default", "Sitename or description")
	key     = flag.String("key", "", "key (e.g. EVT_WU_Roam)")
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

	events, err := u.Events(site)
	if err != nil {
		log.Fatalln(err)
		return
	}

	// Output headline
	fmt.Fprintln(w, "Timestamp\tId\tKey\tMessage")

	for _, re := range events {

		if *key == "" || *key == re.Key {
			timestamp := time.Unix(0, re.Timestamp*int64(time.Millisecond))
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", timestamp.String(), re.ID, re.Key, re.Message)
		}
	}

}
