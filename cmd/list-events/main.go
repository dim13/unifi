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
	limit   = flag.Int("limit", 3000, "Max number of returned events")
	start   = flag.Int("start", 0, "Index of first event (offset)")
	within  = flag.Int("within", 4320, "Number of hours back")
	unknown = flag.Bool("listUnknown", false, "Print unknown events in raw format (helps for adding them)")
)

//{"_limit":500,"_start":0,"_withcount":true,"archived":false} // Alarms

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

	var eventFilter unifi.EventFilter
	eventFilter.Limit = *limit
	eventFilter.Start = *start

	eventFilter.Within = *within

	// Print basic events

	if !*unknown {

		be, err := u.BasicEvents(site, eventFilter)
		if err != nil {
			log.Fatalln(err)
			return
		}

		fmt.Fprintln(w, "Timestamp\tId\tKey\tMessage")

		for _, e := range be {

			timestamp := time.Unix(0, e.Time*int64(time.Millisecond))
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", timestamp.String(), e.ID, e.Key, e.Message)

		}
	} else {

		// Print unknown events in json format

		fmt.Fprintf(w, "\n\nUnknown Events\n")

		re, err := u.RawEvents(site, eventFilter)
		if err != nil {
			log.Fatalln(err)
			return
		}

		totalEvents := 0
		unknownEvents := 0

		for _, e := range re {
			switch e.Key {
			case "EVT_AD_Login":
			case "EVT_AP_Adopted":
			case "EVT_AP_Connected":
			case "EVT_AP_DetectRogueAP":
			case "EVT_AP_DiscoveredPending":
			case "EVT_AP_Isolated":
			case "EVT_AP_Lost_Contact":
			case "EVT_AP_PossibleInterference":
			case "EVT_AP_Restarted":
			case "EVT_AP_RestartedUnknown":
			case "EVT_AP_Upgraded":
			case "EVT_AP_UpgradeScheduled":
			case "EVT_LU_Connected":
			case "EVT_LU_Disconnected":
			case "EVT_SW_Adopted":
			case "EVT_SW_Connected":
			case "EVT_SW_DiscoveredPending":
			case "EVT_SW_Lost_Contact":
			case "EVT_SW_PoeDisconnect":
			case "EVT_SW_RestartedUnknown":
			case "EVT_SW_Upgraded":
			case "EVT_SW_UpgradeScheduled":
			case "EVT_WU_Connected":
			case "EVT_WU_Disconnected":
			case "EVT_WU_Roam":
			case "EVT_WU_RoamRadio":

			default:
				j, err := json.Marshal(&e.Data)
				if err != nil {
					panic(err)
				}
				fmt.Fprintln(w, string(j))
				unknownEvents++
			}
			totalEvents++
		}

		fmt.Fprintf(w, "Total Events:   %d\n", totalEvents)
		fmt.Fprintf(w, "Unknown Events: %d\n", unknownEvents)

	}
	/*	// Events in raw format
		fmt.Fprintln(w, "Timestamp\tId\tKey\tMessage")

		for _, re := range events {

			if *key == "" || *key == re.Key {
				timestamp := time.Unix(0, re.Timestamp*int64(time.Millisecond))
				fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", timestamp.String(), re.ID, re.Key, re.Message)
			}
		}
	*/
	// Parsed Events

	/*

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

		// EVT_SW_PoeDisconnect controller version 5.8.21+ / USW firmware 3.9.27+
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
	*/
}
