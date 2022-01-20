package server

import (
	"net/http"
	"os"
)

func (s *Server) handleUptimeDetails() http.HandlerFunc {
	type output struct {
		UptimeAvailable bool `json:"uptimeAvailable"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		uptimeDomain := os.Getenv("UPTIME_DOMAIN")
		uptimeUser := os.Getenv("UPTIME_USER")
		uptimePass := os.Getenv("UPTIME_PASS")
		o := output{
			UptimeAvailable: uptimeDomain != "" && uptimePass != "" && uptimeUser != "",
		}
		sendJSON(w, o)
	}
}
