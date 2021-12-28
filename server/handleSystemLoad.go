package server

import (
	"net/http"
	"os/exec"
	"strings"
)

func (s *Server) handleSystemLoad() http.HandlerFunc {
	type output struct {
		OneMinute      string `json:"oneMinute"`
		FiveMinutes    string `json:"fiveMinutes"`
		FifteenMinutes string `json:"fifteenMinutes"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		output := output{}
		stdout, err := exec.Command("uptime").Output()
		if err != nil {
			// ...
		}
		stdOutString := string(stdout)
		parts := strings.Split(stdOutString, "load average: ")
		if len(parts) != 2 {
			// ...
		}

		loadParts := strings.Split(parts[1], ", ")
		output.OneMinute = loadParts[0]
		output.FiveMinutes = loadParts[1]
		output.FifteenMinutes = strings.TrimSpace(loadParts[2])

		sendJSON(w, output)
	}
}
