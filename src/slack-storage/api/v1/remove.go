/* vim: set autoindent noexpandtab tabstop=4 shiftwidth=4: */
package v1

import (
	"errors"
	"slack"
)

func remove(data *slack.Slack, args []string) (int, string, error) {
	if len(args) < 1 {
		return 400, "", errors.New("Not enough arguments to remove an item.")
	}

	err := Storage.Remove(args[1])
	if err != nil {
		return 400, "", err
	}

	return 200, "Item removed.", nil
}
