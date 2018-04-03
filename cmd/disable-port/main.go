// Copyright (c) 2014 Dimitri Sokolyuk. All rights reserved.
// Use of this source code is governed by ISC-style license
// that can be found in the LICENSE file.

// list devices
package main

import (
	"flag"
	"log"
	"os"
	"text/tabwriter"

	"github.com/dim13/unifi"
)

var (
	host    = flag.String("host", "unifi", "Controller hostname")
	user    = flag.String("user", "admin", "Controller username")
	pass    = flag.String("pass", "unifi", "Controller password")
	port    = flag.String("port", "8443", "Controller port")
	version = flag.Int("version", 5, "Controller base version")
	siteid  = flag.String("siteid", "default", "Site ID, UniFi v3 only")

	devicename = flag.String("devicename", "", "Name of the device")
	index      = flag.Int("index", 0, "Port Index")
)

func main() {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 3, ' ', 0)
	defer w.Flush()

	flag.Parse()

	u, err := unifi.Login(*user, *pass, *host, *port, *siteid, *version)
	if err != nil {
		log.Fatalln("Login returned error: ", err)
	}
	defer u.Logout()

	// Returns the USW
	usw, err := u.USW(*devicename)

	if err != nil {
		log.Fatalln(err)
	}

	if err != nil {
		log.Fatalln(err)
	}

	// Get port profile name
	profilename := "Disabled"
	profile, err := u.PortProfile(profilename)

	if err != nil {
		log.Fatalln(err)
	}

	overrides := usw.PortOverrides

	var found bool
	for idx, _ := range usw.PortOverrides {
		if overrides[idx].PortIdx == *index {
			overrides[idx].PortconfID = profile.ID
			overrides[idx].POEMode = unifi.POEMODE_OFF
			overrides[idx].Name = "Disable-Test"
			found = true
			break
		}
	}
	// If not found, create a new override for the given port and add it to the slice
	if !found {
		var o unifi.PortOverride
		o.PortIdx = *index
		o.PortconfID = profile.ID
		o.POEMode = unifi.POEMODE_OFF
		o.Name = "Disable-Test"
		overrides = append(overrides, o)
	}
	u.SetPortoverrides(usw.DeviceID, overrides)
}
