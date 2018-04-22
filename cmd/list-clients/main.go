// Copyright (c) 2014 Dimitri Sokolyuk. All rights reserved.
// Use of this source code is governed by ISC-style license
// that can be found in the LICENSE file.

// Example command list-clients
// list associated clients (stations) of a given site
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/dim13/unifi"
)

var (
	host    = flag.String("host", "unifi", "Controller hostname")
	user    = flag.String("user", "admin", "Controller username")
	pass    = flag.String("pass", "unifi", "Controller password")
	version = flag.Int("version", 5, "Controller base version")
	port    = flag.String("port", "8443", "Controller port")
	siteid  = flag.String("siteid", "default", "Sitename or description")
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

	site, err := u.Site(*siteid)
	if err != nil {
		log.Fatal(err)
	}

	devices, err := u.DeviceMap(site)
	if err != nil {
		log.Fatal(err)
	}
	sta, err := u.Sta(site)
	if err != nil {
		log.Fatal(err)
	}

	// Output headline
	fmt.Fprintln(w, "Name\tIsWired\tRadio\tChannel\tESSID\tRoamCount\tSignal\tNoise\tRSSI\tDevicename\tIP\tFirstSeen\tLastSeen\tUptime")

	for _, s := range sta {

		deviceMac := ""

		if s.ApMac != "" {
			deviceMac = s.ApMac
		} else if s.SwMac != "" {
			deviceMac = s.SwMac
		}

		deviceName := devices[deviceMac].DeviceName()

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
			time.Unix(s.FirstSeen, 0).Format("2006-01-02 15:04:05"),
			time.Unix(s.LastSeen, 0).Format("2006-01-02 15:04:05"),
			(time.Duration(s.Uptime) * time.Second).String(),
		}
		fmt.Fprintln(w, strings.Join(p, "\t"))

	}

}
