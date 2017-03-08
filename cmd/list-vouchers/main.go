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
	host = flag.String("host", "unifi", "Controller hostname")
	user = flag.String("user", "admin", "Controller username")
	pass = flag.String("pass", "unifi", "Controller password")
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

	u, err := unifi.Login(*user, *pass, *host, *siteid, *version)
	if err != nil {
		log.Fatal("Login returned error: ", err)
	}
	defer u.Logout()

	vouchers, err := u.VoucherMap()

	if err != nil {
		log.Fatalln(err)
	}

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
