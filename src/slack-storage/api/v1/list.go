/* vim: set autoindent noexpandtab tabstop=4 shiftwidth=4: */
package v1

import (
	"slack"
)

func list(data *slack.Slack, args []string) (int, string, error) {
	all, err := Storage.GetAll()

	if err != nil {
		return 400, "", err
	}

	var text string
	for key, val := range all {
		text += key + ": " + val + "\n"
	}

	return 200, text, nil
}
