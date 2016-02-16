/* vim: set autoindent noexpandtab tabstop=4 shiftwidth=4: */
package v1

import (
	"slack"
)

func add(data *slack.Slack, args []string) (int, string, error) {
	x, err := Storage.Get("y")
	return 200, x, err
}
