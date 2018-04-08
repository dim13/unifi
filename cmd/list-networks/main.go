// Copyright (c) 2014 Dimitri Sokolyuk. All rights reserved.
// Use of this source code is governed by ISC-style license
// that can be found in the LICENSE file.

// Example program list-Networks
// Prints information of all networks of a given Controller
// If no site is specified, the networks of all sites are printed

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/dim13/unifi"
)

var (
	host    = flag.String("host", "unifi", "Controller hostname")
	user    = flag.String("user", "admin", "Controller username")
	pass    = flag.String("pass", "unifi", "Controller password")
	port    = flag.String("port", "8443", "Controller port")
	version = flag.Int("version", 5, "Controller base version")
	siteid  = flag.String("siteid", "", "Site name or description")
)

func main() {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 3, ' ', 0)
	defer w.Flush()

	flag.Parse()

	u, err := unifi.Login(*user, *pass, *host, *port, *siteid, *version)
	if err != nil {
		log.Fatalln("Login returned error: ")
	}
	defer u.Logout()

	var sites []unifi.Site

	// If site specified, print networks of site
	// Else print networks of all sites
	if *siteid != "" {
		site, err := u.Site(*siteid)
		if err != nil {
			log.Fatal(err)
		}
		sites = append(sites, *site)
	} else {

		sites, err = u.Sites()
	}

	for _, s := range sites {

		if len(sites) > 1 {
			fmt.Fprintf(w, "Site: %s\n", s.Desc)
			w.Flush()
		}

		// Returns a slice of devices
		networks, err := u.Networks(&s)

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
}
