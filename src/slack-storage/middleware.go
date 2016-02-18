/* vim: set autoindent noexpandtab tabstop=4 shiftwidth=4: */
package main

import (
	"github.com/gin-gonic/gin"
	"slack"
)

func StorageMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("storage", Storage)
		c.Next()
	}
}

func ParseSlackMessage() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("slack", &slack.Slack{
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
		})
		c.Next()
	}
}

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		if *token == "" {
			c.Next()
		} else {
			idata, _ := c.Get("slack")
			data := idata.(*slack.Slack)
			if data.Token == *token {
				c.Next()
			} else {
				c.JSON(401, gin.H{
					"text": "Unauthorized.",
				})
				c.Abort()
			}
		}
	}
}
