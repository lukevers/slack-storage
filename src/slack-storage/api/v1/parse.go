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

	data := &slack.Slack{
		Token:       c.PostForm("token"),
		TeamId:      c.PostForm("team_id"),
		TeamDomain:  c.PostForm("team_domain"),
		ChannelId:   c.PostForm("channel_id"),
		ChannelName: c.PostForm("channel_name"),
		UserId:      c.PostForm("user_id"),
		UserName:    c.PostForm("user_name"),
		Command:     c.PostForm("command"),
		Text:        c.PostForm("text"),
		ResponseUrl: c.PostForm("response_url"),
	}

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
