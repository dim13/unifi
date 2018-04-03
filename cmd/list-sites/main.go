// Copyright (c) 2014 Dimitri Sokolyuk. All rights reserved.
// Use of this source code is governed by ISC-style license
// that can be found in the LICENSE file.

// list associated stations
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/BFLB/unifi"
)

var (
	host    = flag.String("host", "unifi", "Controller hostname")
	user    = flag.String("user", "ubnt", "Controller username")
	pass    = flag.String("pass", "ubnt", "Controller password")
	port    = flag.String("port", "8443", "Controller port")
	version = flag.Int("version", 5, "Controller base version")
	site    = flag.String("siteid", "defaulid", "Site name or description, UniFi v3 only")
)

func main() {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 3, ' ', 0)
	defer w.Flush()

	flag.Parse()
	u, err := unifi.Login(*user, *pass, *host, *port, *site, *version)
	if err != nil {
		log.Fatal("Login returned error: ", err)
	}
	defer u.Logout()

	sites, err := u.Sites()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(w, "Description\tName\tRole")
	for _, s := range sites {
		p := []string{
			s.Desc,
			s.Name,
			s.Role,
		}
		fmt.Fprintln(w, strings.Join(p, "\t"))
	}
}
