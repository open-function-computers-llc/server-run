package server

import (
	"regexp"
	"strings"
)

func sanitize(s string) string {
	// Make a Regex to say we only want letters and numbers
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	return strings.ToLower(reg.ReplaceAllString(s, "-"))
}

func (s *Server) sanitizeEnv(envPair string) string {
	parts := strings.Split(envPair, "=")
	if len(parts) != 2 {
		s.logger.Error("Can't sanitize the env pair: " + envPair)
		return envPair
	}
	parts[1] = sanitize(parts[1])
	return strings.Join(parts, "=")
}
