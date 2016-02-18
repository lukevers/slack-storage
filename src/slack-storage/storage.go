/* vim: set autoindent noexpandtab tabstop=4 shiftwidth=4: */
package main

import (
	"storage"
	"storage/amnesiadb"
)

var Storage storage.Storage

func setStorage() {
	switch *driver {
	case "amnesiadb", "memory":
		Storage = amnesiadb.New()
	}
}
