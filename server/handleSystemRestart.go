package server

import (
	"encoding/json"
	"net/http"
	"os/exec"
)

func (s *Server) handleSystemRestart() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		service := r.URL.Query().Get("service")

		cmd := exec.Command("systemctl", "restart", service)
		cmd.Run()

		output := map[string]string{
			"message": "success",
		}
		bytes, _ := json.Marshal(output)
		w.Write(bytes)
	}
}
