/* vim: set autoindent noexpandtab tabstop=4 shiftwidth=4: */
package v1

import (
	"errors"
	"slack"
	"strings"
)

func add(data *slack.Slack, args []string) (int, string, error) {
	if len(args) < 2 {
		return 400, "", errors.New("Not enough arguments to add a new item.")
	}

	key := args[0]
	val := strings.Join(args[1:], " ")

	err := Storage.Put(key, val)
	if err != nil {
		return 400, "", err
	}

	return 200, "Item added.", nil
}
