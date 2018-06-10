// Copyright (c) 2014 Dimitri Sokolyuk. All rights reserved.
// Use of this source code is governed by ISC-style license
// that can be found in the LICENSE file.

// Example command list-events
// List Events of a given site
package main

import (
	"encoding/json"
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

	// Events in raw format
	fmt.Fprintln(w, "Timestamp\tId\tKey\tMessage")

	for _, re := range events {

		if *key == "" || *key == re.Key {
			timestamp := time.Unix(0, re.Timestamp*int64(time.Millisecond))
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", timestamp.String(), re.ID, re.Key, re.Message)
		}
	}

	// Parsed Events

	// EVT_AP_Lost_Contact
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "--- EVT_AP_Lost_Contact ---")
	fmt.Fprintln(w, "Timestamp\tId\tSubsystem\tSiteID\tAp\tApName\tPort\tMessage")

	for _, re := range events {

		switch re.Key {
		case "EVT_AP_Lost_Contact":

			var e unifi.EVT_AP_Lost_Contact
			err := json.Unmarshal(re.RawEvent, &e)
			if err == nil {
				timestamp := time.Unix(0, e.Time*int64(time.Millisecond))
				fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\t%s\n", timestamp.String(), e.ID, e.Subsystem, e.SiteID, e.Ap, e.ApName, e.Msg)
			}
		}
	}

	// EVT_SW_PoeDisconnect
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "--- EVT_SW_PoeDisconnect ---")
	fmt.Fprintln(w, "Timestamp\tId\tSubsystem\tSiteID\tSwitch\tSwitchName\tPort\tMessage")

	for _, re := range events {

		switch re.Key {
		case "EVT_SW_PoeDisconnect":

			var e unifi.EVT_SW_PoeDisconnect
			err := json.Unmarshal(re.RawEvent, &e)
			if err == nil {
				timestamp := time.Unix(0, e.Time*int64(time.Millisecond))
				fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\t%d\t%s\n", timestamp.String(), e.ID, e.Subsystem, e.SiteID, e.Sw, e.SwName, e.Port, e.Msg)
			}
		}
	}

}
