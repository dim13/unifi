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
	"strings"
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
	profiles, err := u.PortProfiles()

	if err != nil {
		log.Fatalln(err)
	}

	// Returns a map of networks with ID as key
	nm, err := u.NetworkMap()

	if err != nil {
		log.Fatalln(err)
	}

	// Output headline
	fmt.Fprintln(w, "ID\tName\tSiteID\tOpMode\tAutoneg\tNativeVLAN\tAttrNodelete\tAttrNoEdit\tForward\tTaggedNetworks")

	for _, p := range profiles {
		id := p.ID                                 // string
		name := p.Name                             // string
		siteid := p.SiteID                         // string
		opmode := p.OpMode                         // string
		autoneg := p.Autoneg                       // bool
		nativeNetID := p.NativeNetworkconfID       // string
		attrNoDelete := p.AttrNoDelete             // bool
		attrNoEdit := p.AttrNoEdit                 // bool
		forward := p.Forward                       // string
		taggedNetworkIDs := p.TaggedNetworkconfIds //

		nativeVlan := ""
		nativeVlan = nm[nativeNetID].Name

		taggedNetworks := ""
		for _, t := range taggedNetworkIDs {
			taggedNetworks += fmt.Sprintf("%s ", nm[t].Name)
		}
		taggedNetworks = strings.TrimSpace(taggedNetworks)

		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n", id, name, siteid, opmode, strconv.FormatBool(autoneg), nativeVlan, strconv.FormatBool(attrNoDelete), strconv.FormatBool(attrNoEdit), forward, taggedNetworks)
	}
}
