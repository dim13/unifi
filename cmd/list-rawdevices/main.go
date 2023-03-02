// Copyright (c) 2014 Dimitri Sokolyuk. All rights reserved.
// Use of this source code is governed by ISC-style license
// that can be found in the LICENSE file.

// Example command list-rawdevices
// Generates a json file with a list of devices as received by the controller api.
// Optionally devices can be filtered by device type
package main

import (
	"encoding/json"
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

	filter = flag.String("filter", "", "Filter by device type [uap|usw|ugw]")
	path   = flag.String("path", "./rawDevices.json", "The path and filename of the output")
)

func main() {
	flag.Parse()

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 3, ' ', 0)
	defer w.Flush()

	u, err := unifi.Login(*user, *pass, *host, *port, *siteid, *version)
	if err != nil {
		log.Fatalln("Login returned error: ", err)
	}
	defer u.Logout()

	site, err := u.Site(*siteid)
	if err != nil {
		log.Fatal(err)
	}

	rawDevices, err := u.RawDevices(site, *filter)
	if err != nil {
		log.Fatalln(err)
	}

	var devices []json.RawMessage

	for _, v := range rawDevices {
		devices = append(devices, v.Data)
	}

	json, _ := json.MarshalIndent(devices, "", "  ")
	err = os.WriteFile(*path, json, 0644)
	if err != nil {
		log.Fatalln(err)

	}
}
