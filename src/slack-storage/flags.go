/* vim: set autoindent noexpandtab tabstop=4 shiftwidth=4: */
package main

import (
	"flag"
)

var (
	driver     = flag.String("driver", "amnesiadb", "Database driver")
	dsn        = flag.String("dsn", "", "Data source name")
	host       = flag.String("host", "[::]", "Host to bind on")
	port       = flag.Int("port", 7043, "Port to listen on")
	production = flag.Bool("production", false, "Production mode")
	token      = flag.String("token", "", "Slack token")
)

func flags() {
	flag.Parse()
}
