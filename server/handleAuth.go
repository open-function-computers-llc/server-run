package server

import (
	"net/http"
	"time"

	"github.com/dchest/uniuri"
)

func (s *Server) handleAuth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		u := r.FormValue("user")
		p := r.FormValue("pass")

		newSessionKey := uniuri.NewLen(64)

		if u == "admin" && p == "password" {
			s.sessions[newSessionKey] = session{
				expires: time.Now().Add(time.Minute * 30),
			}

			sendJSON(w, map[string]string{
				"auth-token": newSessionKey,
			})
			return
		}

		sendJSONError(w, http.StatusForbidden, map[string]string{
			"error": "Invalid username or password",
		})
	}
}
