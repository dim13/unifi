// Copyright (c) 2014 The unifi Authors. All rights reserved.
// Use of this source code is governed by ISC-style license
// that can be found in the LICENSE file.

// Example command list-vouchers
// List vouchers of a given site

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

	vouchers, err := u.VoucherMap(site)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Fprintln(w, "Code\tCreateTime\tDuration\tNote\tQuota\tUsed")

	for _, v := range vouchers {
		p := []string{
			v.Code,
			strconv.Itoa(v.CreateTime),
			strconv.Itoa(v.Duration),
			v.Note,
			strconv.Itoa(v.Quota),
			strconv.Itoa(v.Used),
		}
		fmt.Fprintln(w, strings.Join(p, "\t"))
	}

}
