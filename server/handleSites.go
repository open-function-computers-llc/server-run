package server

import (
	"net/http"
)

func (s *Server) handleSites() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sendJSON(w, s.sites)
	}
}
