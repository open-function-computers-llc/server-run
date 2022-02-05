package commands

import (
	"errors"
	"strings"
)

func runScript(script string, args []string) (string, error) {
	availableScripts := []string{
		"unix-username-format",
	}
	validScript := false
	for _, s := range availableScripts {
		if s == script {
			validScript = true
			break
		}
	}
	if !validScript {
		return "", errors.New("Invalid script passed to the script runner. Valid scripts are: " + strings.Join(availableScripts, ", "))
	}
	return "", nil
}
