package server

import (
	"bufio"
	"encoding/json"
	"net/http"
	"os/exec"

	"nhooyr.io/websocket"
)

func (s *Server) streamScriptRunner() http.HandlerFunc {
	availableScriptsWithArguments := map[string]string{
		"unlock": "ofco-unlock-site.sh",
		"lock":   "ofco-lock-site-production.sh",
	}
	availableScriptsWithoutArguments := map[string]string{
		"f2banstatus": "f2bstatus",
	}

	// combine the above
	var availableScripts []string
	for name, _ := range availableScriptsWithArguments {
		availableScripts = append(availableScripts, name)
	}
	for name, _ := range availableScriptsWithoutArguments {
		availableScripts = append(availableScripts, name)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		script := r.FormValue("script")

		if script == "" {
			sendJSONError(w, http.StatusBadRequest, map[string]string{
				"error": "the query param `script` is required",
			})
			return
		}

		// verify that the script called is in the g2g list
		validScript := false
		for _, checkedName := range availableScripts {
			if script == checkedName {
				validScript = true
				break
			}
		}
		if !validScript {
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
		var cmd *exec.Cmd
		if _, ok := availableScriptsWithArguments[script]; ok {
			d := r.FormValue("domain")
			cmd = exec.Command(s.scriptsRoot+availableScriptsWithArguments[script], d)
		}
		if _, ok := availableScriptsWithoutArguments[script]; ok {
			cmd = exec.Command(s.scriptsRoot + availableScriptsWithoutArguments[script])
		}

		outPipe, _ := cmd.StdoutPipe()
		cmd.Stderr = cmd.Stdout
		cmd.Env = append(cmd.Env, "NOCONFIRM=yes")
		scanner := bufio.NewScanner(outPipe)

		go func() {
			for scanner.Scan() {
				line := scanner.Text()
				commChannel <- line
			}
			close(commChannel)
		}()

		err = cmd.Start()
		if err != nil {
			s.logger.Error(err)
		}

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
