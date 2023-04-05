package server

import (
	"net/http"
)

func (s *Server) bindRoutes() {
	// API routes
	protectedRoutes := map[string]http.HandlerFunc{
		"sites":           s.handleSites(),
		"details":         s.handleDetails(),
		"update":          s.handleUpdateSite(),
		"analytics":       s.handleAnalytics(),
		"analytics-json":  s.handleAnalyticsJSON(),
		"uptime":          s.handleUptimeInfo(),
		"uptime-provider": s.handleUptimeDetails(),
		"process-list":    s.handleSystemProcesses(),
		"process-details": s.handleSystemDetails(),
		"process-restart": s.handleSystemRestart(),
	}
	for path, handler := range protectedRoutes {
		http.Handle("/api/"+path, s.LogRequest(s.ProtectRequest(handler)))
	}

	// web socket routes
	protectedStreamRoutes := map[string]http.HandlerFunc{
		"system-load": s.streamSystemLoad(),
		"script":      s.streamScriptRunner(),
	}
	for path, handler := range protectedStreamRoutes {
		http.Handle("/stream/"+path, s.LogRequest(s.ProtectRequest(handler)))
	}

	// no auth needed
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
