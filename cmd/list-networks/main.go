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
	networks, err := u.Networks()

	if err != nil {
		log.Fatalln(err)
	}

	// Output headline
	fmt.Fprintln(w, "ID\tName\tSiteID\tPurpose\tSubnet\tNetworkgroup\tDomainname\tAttrHiddenId\tAttrNoDelete\tVlanEnabled")

	for _, n := range networks {
		id := n.ID           // string
		name := n.Name       // string
		siteid := n.SiteID   // string
		purpose := n.Purpose // string
		subnet := n.IPSubnet
		networkgroup := n.Networkgroup // string
		domaninname := n.DomainName    // string
		attrhiddenid := n.AttrHiddenID //string
		attrnodelete := n.AttrNoDelete // bool
		vlanenabled := n.VlanEnabled   // bool

		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n", id, name, siteid, purpose, subnet, networkgroup, domaninname, attrhiddenid, strconv.FormatBool(attrnodelete), strconv.FormatBool(vlanenabled))
	}
}
