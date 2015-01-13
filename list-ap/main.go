// list associated stations
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
	host    = flag.String("host", "unifi", "Controller hostname")
	user    = flag.String("user", "admin", "Controller username")
	pass    = flag.String("pass", "unifi", "Controller password")
	version = flag.Int("version", 2, "Controller base version")
	siteid  = flag.String("siteid", "default", "Site ID, UniFi v3 only")
)

func main() {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 3, ' ', 0)
	defer w.Flush()

	flag.Parse()

	u, err := unifi.Login(*user, *pass, *host, *siteid, *version)
	if err != nil {
		log.Fatalln("Login returned error: ", err)
	}
	defer u.Logout()

	aps, err := u.Aps()
	if err != nil {
		log.Fatalln(err)
	}

	for _, s := range aps {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n",
			s.Mac, s.ModelName(), s.Version, s.Name, s.Status())
	}
}
