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
	case "":
		// If no info is passed, run list
		code, text, err = list(data, args)
	case "list":
		code, text, err = list(data, args)
	case "add":
		code, text, err = add(data, args[1:])
	case "remove":
		code, text, err = remove(data, args)
	default:
		code, text, err = add(data, args)
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
