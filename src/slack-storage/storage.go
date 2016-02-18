/* vim: set autoindent noexpandtab tabstop=4 shiftwidth=4: */
package main

import (
	"log"
	"storage"
	"storage/amnesiadb"
	"storage/boltdb"
)

var Storage storage.Storage

func setStorage() {
	var err error
	switch *driver {
	case "boltdb", "bolt":
		Storage, err = boltdb.New(*dsn)
	default:
		fallthrough // Default to amnesiadb
	case "amnesiadb", "memory":
		Storage, err = amnesiadb.New()
	}

	if err != nil {
		log.Fatal(err)
	}
}
