//Example of creating a new voucher, the return value - create time of the voucher

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/dim13/unifi"
)

var (
	num    = flag.String("num", "1", "The number of the new vouchers")
	multi  = flag.String("multi", "0", "If 0 is the multi-voucher")
	minute = flag.String("minute", "1440", "Duration of the voucher in minutes")
	note   = flag.String("note", "", "Note of the voucher")

	host    = flag.String("host", "unifi", "Controller hostname")
	user    = flag.String("user", "admin", "Controller username")
	pass    = flag.String("pass", "unifi", "Controller password")
	version = flag.Int("version", 4, "Controller base version")
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

	jsonData := unifi.NewVoucher{
		Cmd:          "create-voucher",
		Expire:       "custom",
		ExpireNumber: *minute,
		ExpireUnit:   "1",
		N:            *num,
		Note:         *note,
		Quota:        *multi,
	}

	res, err := u.NewVoucher(jsonData)
	if err != nil {
		log.Fatalln(err)
	}

	ct := res[0].CreateTime
	fmt.Println(ct)

}
