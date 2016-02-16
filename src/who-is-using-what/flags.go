/* vim: set autoindent noexpandtab tabstop=4 shiftwidth=4: */
package main

import (
	"flag"
)

var (
	host       = flag.String("host", "[::]", "Host to bind on")
	port       = flag.Int("port", 7043, "Port to listen on")
	production = flag.Bool("production", false, "Production mode")
)

func flags() {
	flag.Parse()
}
