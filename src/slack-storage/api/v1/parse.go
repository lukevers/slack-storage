/* vim: set autoindent noexpandtab tabstop=4 shiftwidth=4: */
package v1

import (
	"github.com/gin-gonic/gin"
	"slack"
	"storage"
)

var Storage storage.Storage

func Parse(c *gin.Context) {
	s, _ := c.Get("storage")
	Storage = s.(storage.Storage)

	idata, _ := c.Get("slack")
	data := idata.(*slack.Slack)

	var code int
	var text string
	var err error

	args := data.Args()
	switch args[0] {
	case "", "list":
		code, text, err = list(data, args)
	case "remove", "rm":
		code, text, err = remove(data, args)
	default:
		fallthrough
	case "add", "create", "update":
		code, text, err = add(data, args[1:])
	}

	if err != nil {
		c.JSON(code, gin.H{
			"text": err.Error(),
		})
	} else {
		c.JSON(code, gin.H{
			"text": text,
		})
	}
}
