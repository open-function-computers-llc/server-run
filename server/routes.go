package server

import (
	"net/http"
)

func (s *Server) bindRoutes() {
	// API routes
	protectedRoutes := map[string]http.HandlerFunc{
		"system-load": s.handleSystemLoad(),
		"sites":       s.handleSites(),
		"details":     s.handleDetails(),
	}
	for path, handler := range protectedRoutes {
		http.Handle("/api/"+path, s.LogRequest(s.ProtectRequest(handler)))
	}
	openRoutes := map[string]http.HandlerFunc{
		"auth": s.handleAuth(),
	}
	for path, handler := range openRoutes {
		http.Handle("/api/"+path, s.LogRequest(handler))
	}

	// filesystem server for angular app
	frontendFS := http.FileServer(http.FS(s.filesystem))
	http.Handle("/", frontendFS)
}
