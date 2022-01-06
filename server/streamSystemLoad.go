package server

import (
	"encoding/json"
	"errors"
	"net/http"
	"os/exec"
	"strings"
	"time"

	"nhooyr.io/websocket"
)

func (s *Server) streamSystemLoad() http.HandlerFunc {
	type output struct {
		OneMinute      string `json:"oneMinute"`
		FiveMinutes    string `json:"fiveMinutes"`
		FifteenMinutes string `json:"fifteenMinutes"`
	}

	parse := func(stdout []byte) ([]byte, error) {
		output := output{}
		stdOutString := string(stdout)
		parts := strings.Split(stdOutString, "load average: ")
		if len(parts) != 2 {
			return nil, errors.New("There was an error gathering the command output")
		}

		loadParts := strings.Split(parts[1], ", ")
		output.OneMinute = loadParts[0]
		output.FiveMinutes = loadParts[1]
		output.FifteenMinutes = strings.TrimSpace(loadParts[2])

		b, _ := json.Marshal(output)

		return b, nil
	}

	return func(w http.ResponseWriter, r *http.Request) {
		c, err := websocket.Accept(w, r, nil)
		if err != nil {
			sendJSONError(w, http.StatusInternalServerError, map[string]string{
				"error": "Error setting up web socket connection: " + err.Error(),
			})
			return
		}
		defer c.Close(websocket.StatusInternalError, "uh oh...")

		for {
			stdout, err := exec.Command("uptime").Output()
			if err != nil {
				sendJSONError(w, http.StatusInternalServerError, map[string]string{
					"error": "There was an error gathering the `uptime` command output",
				})
				break
			}

			b, err := parse(stdout)
			if err != nil {
				sendJSONError(w, http.StatusInternalServerError, map[string]string{
					"error": err.Error(),
				})
				break
			}

			c.Write(r.Context(), websocket.MessageText, b)
			time.Sleep(5 * time.Second)
		}

		c.Close(websocket.StatusNormalClosure, "")
	}
}
