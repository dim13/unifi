// Copyright (c) 2014 Dimitri Sokolyuk. All rights reserved.
// Use of this source code is governed by ISC-style license
// that can be found in the LICENSE file.

// Example command list-health
// list controller health information
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

	health, err := u.Health(site)
	if err != nil {
		log.Fatal(err)
	}

	var p []string

	// Print LAN health
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subsystem\tStatus\tNumAdopted\tNumSw\tNumDisconnected\tNumPending\tNumUser\tNumGuest\tRxBytesR\tTxBytesR\tLanIP\t")

	p = []string{
		health.LAN.Subsystem,
		health.LAN.Status,
		strconv.Itoa(health.LAN.NumAdopted),
		strconv.Itoa(health.LAN.NumSw),
		strconv.Itoa(health.LAN.NumDisconnected),
		strconv.Itoa(health.LAN.NumPending),
		strconv.Itoa(health.LAN.NumUser),
		strconv.Itoa(health.LAN.NumGuest),
		strconv.FormatInt(health.LAN.RxBytesR, 10),
		strconv.FormatInt(health.LAN.TxBytesR, 10),
		health.LAN.LanIP,
	}
	fmt.Fprintln(w, strings.Join(p, "\t"))

	// Print WLAN health
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subsystem\tStatus\tNumAdopted\tNumAp\tNumDisconnected\tNumPending\tNumUser\tNumGuest\tRxBytesR\tTxBytesR")

	p = []string{
		health.WLAN.Subsystem,
		health.WLAN.Status,
		strconv.Itoa(health.WLAN.NumAdopted),
		strconv.Itoa(health.WLAN.NumAp),
		strconv.Itoa(health.WLAN.NumDisconnected),
		strconv.Itoa(health.WLAN.NumUser),
		strconv.Itoa(health.WLAN.NumGuest),
		strconv.FormatInt(health.WLAN.RxBytesR, 10),
		strconv.FormatInt(health.WLAN.TxBytesR, 10),
		strconv.Itoa(health.WLAN.NumPending),
	}
	fmt.Fprintln(w, strings.Join(p, "\t"))

	// Print WAN health
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subsystem\tStatus\tNumAdopted\tNumGw\tNumDisconnected\tNumPending\tNumSta\tNumGuest\tRxBytesR\tTxBytesR\tWanIP\tNameServers")

	p = []string{
		health.WAN.Subsystem,
		health.WAN.Status,
		strconv.Itoa(health.WAN.NumAdopted),
		strconv.Itoa(health.WAN.NumGw),
		strconv.Itoa(health.WAN.NumDisconnected),
		strconv.Itoa(health.WAN.NumPending),
		strconv.Itoa(health.WAN.NumSta),
		strconv.FormatInt(health.WAN.RxBytesR, 10),
		strconv.FormatInt(health.WAN.TxBytesR, 10),
		health.WAN.WanIP,
		strings.Join(health.WAN.Nameservers, " "),
	}
	fmt.Fprintln(w, strings.Join(p, "\t"))

	// Print WWW health
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subsystem\tStatus\tRxBytesR\tTxBytesR\tDrops\tGwMac\tLatency\tSpeedtestLastrun\tSpeedtestPing\tSpeedtestStatus\tUptime\tXputDown\tXputUp")

	p = []string{
		health.WWW.Subsystem,
		health.WWW.Status,
		strconv.FormatInt(health.WWW.RxBytesR, 10),
		strconv.FormatInt(health.WWW.TxBytesR, 10),
		strconv.Itoa(health.WWW.Drops),
		health.WWW.GwMac,
		strconv.Itoa(health.WWW.Latency),
		time.Unix(int64(health.WWW.SpeedtestLastrun), 0).String(),
		strconv.Itoa(health.WWW.SpeedtestPing),
		health.WWW.SpeedtestStatus,
		(time.Duration(health.WWW.Uptime) * time.Second).String(),
		strconv.FormatFloat(health.WWW.XputDown, 'f', 6, 64),
		strconv.FormatFloat(health.WWW.XputUp, 'f', 6, 64),
	}
	fmt.Fprintln(w, strings.Join(p, "\t"))

	// Print VPN health
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subsystem\tStatus")

	p = []string{
		health.WAN.Subsystem,
		health.WAN.Status,
	}
	fmt.Fprintln(w, strings.Join(p, "\t"))

}
