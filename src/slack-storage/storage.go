/* vim: set autoindent noexpandtab tabstop=4 shiftwidth=4: */
package main

import (
	"storage"
	"storage/memory"
)

var Storage storage.Storage

func setStorage() {
	Storage = memory.New()
}
