//For Unifi Controller ver.4

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
	siteid  = flag.String("siteid", "default", "Site ID, UniFi v3 only")
)

func main() {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 3, ' ', 0)
	defer w.Flush()

	flag.Parse()
	version := new(int)
	siteid := new(string)
	*version = 4
	*siteid = "default"

	u, err := unifi.Login(*user, *pass, *host, *port, *siteid, *version)
	if err != nil {
		log.Fatal("Login returned error: ", err)
	}
	defer u.Logout()

	vouchers, err := u.VoucherMap()

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
