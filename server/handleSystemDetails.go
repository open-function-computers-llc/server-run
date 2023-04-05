package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
)

func (s *Server) handleSystemDetails() http.HandlerFunc {
	type SystemService struct {
		Name   string `json:"name"`
		Status string `json:"status"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		service := r.URL.Query().Get("service")
		var out bytes.Buffer

		sJSON := SystemService{
			Name: service,
		}

		cmd := exec.Command("systemctl", "status", service, "--no-pager")
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			fmt.Println(err.Error())
			w.Write([]byte("nope!"))
			return
		}
		sJSON.Status = out.String()

		outputBytes, _ := json.Marshal(sJSON)
		w.Write(outputBytes)
	}
}
