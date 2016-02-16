/* vim: set autoindent noexpandtab tabstop=4 shiftwidth=4: */
package main

import (
	"github.com/gin-gonic/gin"
)

func StorageMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("storage", Storage)
		c.Next()
	}
}
