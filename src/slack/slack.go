/* vim: set autoindent noexpandtab tabstop=4 shiftwidth=4: */
package slack

import (
	"strings"
)

type Slack struct {
	Token       string
	TeamId      string
	TeamDomain  string
	ChannelId   string
	ChannelName string
	UserId      string
	UserName    string
	Command     string
	Text        string
	ResponseUrl string
}

func (s *Slack) Args() []string {
	return strings.Split(s.Text, " ")
}
