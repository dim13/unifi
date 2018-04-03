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
	"reflect"
	"strconv"
	"strings"
	"text/tabwriter"

	"github.com/BFLB/unifi"
)

var (
	host    = flag.String("host", "unifi", "Controller hostname")
	user    = flag.String("user", "admin", "Controller username")
	pass    = flag.String("pass", "unifi", "Controller password")
	version = flag.Int("version", 5, "Controller base version")
	port    = flag.String("port", "8443", "Controller port")
	siteid  = flag.String("siteid", "default", "Site ID, UniFi v3 only")
)

func main() {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 3, ' ', 0)
	defer w.Flush()

	flag.Parse()
	u, err := unifi.Login(*user, *pass, *host, *port, *siteid, *version)
	if err != nil {
		log.Fatal("Login returned error: ", err)
	}
	defer u.Logout()

	aps, err := u.DeviceMap()
	if err != nil {
		log.Fatal(err)
	}
	sta, err := u.Sta()
	if err != nil {
		log.Fatal(err)
	}

	// Output headline
	fmt.Fprintln(w, "Name\tIsWired\tRadio\tChannel\tESSID\tRoamCount\tSignal\tNoise\tRSSI\tDevicename\tIP\tModelName")

	for _, s := range sta {

		deviceMac := ""
		deviceName := ""
		modelName := ""

		if s.ApMac != "" {
			deviceMac = s.ApMac
		} else if s.SwMac != "" {
			deviceMac = s.SwMac
		}

		if deviceMac != "" {

		}
		d := aps[deviceMac]
		var devicetype string
		devicetype = reflect.ValueOf(d).Type().String()

		switch devicetype {
		case "unifi.UAP":
			// Type assertion from interface to unifi.Uap
			d := d.(unifi.UAP)
			deviceName = d.DeviceName()
			modelName = d.ModelName()

		case "unifi.USW":
			// Type assertion from interface to unifi.Uap
			d := d.(unifi.USW)
			deviceName = d.DeviceName()
			modelName = d.ModelName()
		}

		p := []string{
			s.Name(),
			strconv.FormatBool(s.IsWired),
			s.Radio,
			strconv.Itoa(s.Channel),
			s.ESSID,
			strconv.Itoa(s.RoamCount),
			strconv.Itoa(s.Signal),
			strconv.Itoa(s.Noise),
			strconv.Itoa(s.Rssi),
			deviceName,
			s.IP,
			modelName,
		}
		fmt.Fprintln(w, strings.Join(p, "\t"))

	}

}
