package server

import (
	"bufio"
	"encoding/json"
	"net/http"
	"os/exec"

	"nhooyr.io/websocket"
)

func (s *Server) streamScriptRunner() http.HandlerFunc {
	availableScripts := map[string]string{
		"unlock": "ofco-unlock-site.sh",
		"lock":   "ofco-lock-site-production.sh",
	}
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		d := r.FormValue("domain")
		script := r.FormValue("script")

		if d == "" || script == "" {
			sendJSONError(w, http.StatusBadRequest, map[string]string{
				"error": "the query params `domain` and `script` are both required",
			})
			return
		}

		if _, ok := availableScripts[script]; !ok {
			sendJSONError(w, http.StatusBadRequest, map[string]string{
				"error": "Invalid script requested",
			})
			return
		}

		c, err := websocket.Accept(w, r, nil)
		if err != nil {
			sendJSONError(w, http.StatusInternalServerError, map[string]string{
				"error": "Error setting up web socket connection: " + err.Error(),
			})
			return
		}
		defer c.Close(websocket.StatusInternalError, "uh oh...")

		// this is the channel that we will use to communicate between the exec stdOut and the web socket
		commChannel := make(chan string)

		// set up the command runner
		cmd := exec.Command(s.scriptsRoot+availableScripts[script], d)
		outPipe, _ := cmd.StdoutPipe()
		cmd.Stderr = cmd.Stdout
		scanner := bufio.NewScanner(outPipe)

		go func() {
			for scanner.Scan() {
				line := scanner.Text()
				commChannel <- line
			}
			close(commChannel)
		}()

		err = cmd.Start()

		for incomingCommMessage := range commChannel {
			message := map[string]interface{}{
				"output": incomingCommMessage,
			}
			b, _ := json.Marshal(message)
			c.Write(r.Context(), websocket.MessageText, b)
		}

		c.Close(websocket.StatusNormalClosure, "")
	}
}
