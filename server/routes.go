package server

import (
	"net/http"
)

func (s *Server) bindRoutes() {
	// API routes
	var apiRoutes = map[string]http.HandlerFunc{
		"system-load": s.handleSystemLoad(),
	}
	for path, handler := range apiRoutes {
		http.Handle("/api/"+path, handler)
	}

	// filesystem server for angular app
	frontendFS := http.FileServer(http.FS(s.filesystem))
	http.Handle("/", frontendFS)
}
