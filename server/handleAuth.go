package server

import (
	"net/http"
	"time"

	"github.com/dchest/uniuri"
)

func (s *Server) handleAuth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(1024)

		u := r.PostFormValue("user")
		p := r.PostFormValue("pass")

		newSessionKey := uniuri.NewLen(64)
		expires := time.Now().Add(time.Minute * 30)

		if u == s.adminUser && p == s.adminPass {
			s.sessions[newSessionKey] = session{
				expires: expires,
			}

			sendJSON(w, map[string]string{
				"authToken": newSessionKey,
				"expiresAt": expires.UTC().Format("2006-01-02T15:04:05.999Z07:00"),
			})
			return
		}

		sendJSONError(w, http.StatusForbidden, map[string]string{
			"error": "Invalid username or password",
		})
	}
}
