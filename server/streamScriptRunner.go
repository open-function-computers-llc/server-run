package server

import (
	"bufio"
	"encoding/json"
	"net/http"
	"os/exec"
	"strings"

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
	availableScriptsWithENVRequirements := map[string]string{
		"addAccount":   "create-new-account.sh",
		"cloneAccount": "clone-wordpress-account.sh",
	}

	// combine the above
	var availableScripts []string
	for name := range availableScriptsWithArguments {
		availableScripts = append(availableScripts, name)
	}
	for name := range availableScriptsWithoutArguments {
		availableScripts = append(availableScripts, name)
	}
	for name := range availableScriptsWithENVRequirements {
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

		// 3 ways to run a script!

		// option 1 - the script has a single argument
		if _, ok := availableScriptsWithArguments[script]; ok {
			d := r.FormValue("arg")
			cmd = exec.Command(s.scriptsRoot+availableScriptsWithArguments[script], d)
		}

		// option 2 - the script has zero arguments and nothing special in the ENV
		if _, ok := availableScriptsWithoutArguments[script]; ok {
			cmd = exec.Command(s.scriptsRoot + availableScriptsWithoutArguments[script])
		}

		// option 3 - this is the weird one
		// the only takes the script name, no arguments, but all the ENV variables
		// need to be passed into this script in a single form input valiable
		// example: /stream/script?script=cloneAccount&env=DESTINATION_ACCOUNT=newbie|SOURCE_ACCOUNT=asdfasdf
		// will be translated to
		// script = cloneAccount
		// env =
		//     DESTINATION_ACCOUNT=newbie
		//     SOURCE_ACCOUNT=asdfasdf
		if _, ok := availableScriptsWithENVRequirements[script]; ok {
			cmd = exec.Command(s.scriptsRoot + availableScriptsWithENVRequirements[script])
			e := r.FormValue("env")
			envPairs := strings.Split(e, "|")
			for _, env := range envPairs {
				cmd.Env = append(cmd.Env, s.sanitizeEnv(env))
			}
			s.logger.Info("Running: ", s.scriptsRoot+availableScriptsWithENVRequirements[script])
			s.logger.Info("Script ENV: ", cmd.Env)
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
