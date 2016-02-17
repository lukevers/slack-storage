/* vim: set autoindent noexpandtab tabstop=4 shiftwidth=4: */
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func serveHttp() {
	if *production {
		gin.SetMode(gin.ReleaseMode)
	}

	router = gin.Default()
	route()
	router.Run(fmt.Sprintf("%s:%d", *host, *port))
}
