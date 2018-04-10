// Copyright (c) 2014 The unifi Authors. All rights reserved.
// Use of this source code is governed by ISC-style license
// that can be found in the LICENSE file.

// Example command disable-port
// Disables a port of a USW
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
	siteid  = flag.String("siteid", "default", "Sitename or description")

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

	site, err := u.Site(*siteid)
	if err != nil {
		log.Fatal(err)
	}

	// Returns the USW
	usw, err := u.USW(site, *devicename)

	if err != nil {
		log.Fatalln(err)
	}

	if err != nil {
		log.Fatalln(err)
	}

	// Get port profile name
	profilename := "Disabled"
	profile, err := u.PortProfile(site, profilename)

	if err != nil {
		log.Fatalln(err)
	}

	overrides := usw.PortOverrides

	var found bool
	for i := range usw.PortOverrides {
		if overrides[i].PortIdx == *index {
			overrides[i].PortconfID = profile.ID
			overrides[i].POEMode = unifi.POEModeOff
			overrides[i].Name = "Disabled by script"
			found = true
			break
		}
	}
	// If not found, create a new override for the given port and add it to the slice
	if !found {
		var o unifi.PortOverride
		o.PortIdx = *index
		o.PortconfID = profile.ID
		o.POEMode = unifi.POEModeOff
		o.Name = "Disabled by script"
		overrides = append(overrides, o)
	}
	u.SetPortoverrides(site, usw.DeviceID, overrides)
}
