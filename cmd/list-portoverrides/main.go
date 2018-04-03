// Copyright (c) 2014 Dimitri Sokolyuk. All rights reserved.
// Use of this source code is governed by ISC-style license
// that can be found in the LICENSE file.

// list devices
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/BFLB/unifi"
)

var (
	host    = flag.String("host", "unifi", "Controller hostname")
	user    = flag.String("user", "admin", "Controller username")
	pass    = flag.String("pass", "unifi", "Controller password")
	port    = flag.String("port", "8443", "Controller port")
	version = flag.Int("version", 5, "Controller base version")
	siteid  = flag.String("siteid", "default", "Site ID, UniFi v3 only")

	devicename = flag.String("devicename", "", "Name of the device")
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

	// Returns a slice of devices
	usw, err := u.USW(*devicename)

	if err != nil {
		log.Fatalln(err)
	}

	pm, err := u.PortProfileMap()

	if err != nil {
		log.Fatalln(err)
	}

	overrides := usw.PortOverrides

	// Output headline
	fmt.Fprintln(w, "PortIndex\tName\tPortprofileName\tPoeMode\tOpMode\tAutoneg\tSpeed\tFullDuplex\tIsolation\tLLDPmedEnabled")

	for _, o := range overrides {
		portindex := o.PortIdx             // int
		name := o.Name                     // string
		profileID := o.PortconfID          // string
		poemode := o.POEMode               // string
		opmode := o.OpMode                 // string
		pautoneg := o.Autoneg              // bool
		pspeed := o.Speed                  // int
		pfullduplex := o.FullDuplex        // bool
		pisolation := o.Isolation          // bool
		lldpmedEnabled := o.LLDPMedEnabled // bool

		// Get port profile name
		profilename := pm[profileID].Name

		// Handle nil pointer in case of absent field
		fullduplex := ""
		if pfullduplex != nil {
			fullduplex = strconv.FormatBool(*pfullduplex)
		}

		// Handle nil pointer in case of absent field
		speed := ""
		if pspeed != nil {
			speed = strconv.FormatInt(int64(*pspeed), 10)
		}

		// Handle nil pointer in case of absent field
		autoneg := ""
		if pautoneg != nil {
			autoneg = strconv.FormatBool(*pautoneg)
		}

		// Handle nil pointer in case of absent field
		isolation := ""
		if pisolation != nil {
			isolation = strconv.FormatBool(*pisolation)
		}

		// Handle nil pointer in case of absent field
		lldpmed := ""
		if lldpmedEnabled != nil {
			lldpmed = strconv.FormatBool(*lldpmedEnabled)
		}

		fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n", portindex, name, profilename, poemode, opmode, autoneg, speed, fullduplex, isolation, lldpmed)
	}
}
