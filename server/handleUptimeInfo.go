package server

import (
	"io"
	"net/http"
	"os"
)

func (s *Server) handleUptimeInfo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uptimeUser := os.Getenv("UPTIME_USER")
		uptimePass := os.Getenv("UPTIME_PASS")
		uptimeDomain := os.Getenv("UPTIME_DOMAIN")
		if uptimeUser == "" || uptimePass == "" || uptimeDomain == "" {
			sendJSONError(w, http.StatusInternalServerError, map[string]string{
				"error": "Invalid ENV settings to check uptime",
			})
			return
		}

		// get URI to check
		r.ParseForm()
		uri := r.FormValue("uri")
		if uri == "" {
			sendJSONError(w, http.StatusBadRequest, map[string]string{
				"error": "Invalid `uri` requested",
			})
			return
		}

		resp, err := http.Get("https://" + uptimeUser + ":" + uptimePass + "@" + uptimeDomain + "/details?url=" + uri)
		if err != nil {
			sendJSONError(w, http.StatusInternalServerError, map[string]string{
				"error": "Error from upstream uptime monitoring service: " + err.Error(),
			})
			return
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			sendJSONError(w, http.StatusInternalServerError, map[string]string{
				"error": "Error parsing response from upstream uptime monitoring service: " + err.Error(),
			})
			return
		}

		w.WriteHeader(resp.StatusCode)
		w.Write(body)
	}
}
