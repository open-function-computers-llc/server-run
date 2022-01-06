package server

import (
	"net/http"
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

		// w.Write(content)
	}
}
