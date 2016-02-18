/* vim: set autoindent noexpandtab tabstop=4 shiftwidth=4: */
package main

import (
	"slack-storage/api/v1"
)

func route() {
	V1 := router.Group("/api/v1")

	V1.Use(Authenticate(), StorageMiddleware())
	{
		V1.POST("/parse", v1.Parse)
	}
}
