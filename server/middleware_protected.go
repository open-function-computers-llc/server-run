package server

import (
	"net/http"
	"strconv"
	"time"
)

func (s *Server) ProtectRequest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sessionKey := r.Header.Get("Authorization")
		session, ok := s.sessions[sessionKey]
		if !ok {
			sendJSONError(w, http.StatusUnauthorized, map[string]string{
				"error": "Invalid Session",
			})
			return
		}
		if time.Now().After(session.expires) {
			sendJSONError(w, http.StatusUnauthorized, map[string]string{
				"error": "Session Expired",
			})
			return
		}

		w.Header().Add("expires-in", strconv.Itoa(int(session.expires.Sub(time.Now()).Milliseconds())))
		next(w, r)
	}
}
