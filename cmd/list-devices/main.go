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

	"github.com/dim13/unifi"
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
	flag.Parse()

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 3, ' ', 0)
	defer w.Flush()

	u, err := unifi.Login(*user, *pass, *host, *port, *siteid, *version)
	if err != nil {
		log.Fatalln("Login returned error: ", err)
	}
	defer u.Logout()

	devices, err := u.Devices()
	if err != nil {
		log.Fatalln(err)
	}

	var numberOfAps, numberOfSwitches, numberOfDevices int

	// Output headline
	fmt.Fprintln(w, "Type\tName\tIP\tMac\tModelName\tVersion\tStatus\tNumberOfClients\tTxBytes\tRxBytes")

	for _, d := range devices {
		switch dev := d.(type) {
		case unifi.UAP:
			p := []string{
				"AP    ",
				dev.DeviceName(),
				dev.IP,
				dev.Mac,
				dev.ModelName(),
				dev.Version,
				dev.Status(),
				strconv.Itoa(dev.NumSta),
				unifi.Bytes(dev.TxBytes).String(),
				unifi.Bytes(dev.RxBytes).String(),
			}
			fmt.Fprintln(w, strings.Join(p, "\t"))
			numberOfAps++

		case unifi.USW:
			p := []string{
				"Switch",
				dev.DeviceName(),
				dev.IP,
				dev.Mac,
				dev.ModelName(),
				dev.Version,
				dev.Status(),
				strconv.Itoa(dev.NumSta),
				unifi.Bytes(dev.TxBytes).String(),
				unifi.Bytes(dev.RxBytes).String(),
			}
			fmt.Fprintln(w, strings.Join(p, "\t"))
			numberOfSwitches++

		}
		numberOfDevices++

	}

	w.Flush() // Write
	fmt.Printf("\n")
	fmt.Printf("Number of APs     : %d\n", numberOfAps)
	fmt.Printf("Number of Switches: %d\n", numberOfSwitches)
	fmt.Printf("Total             : %d\n", numberOfDevices)

}
