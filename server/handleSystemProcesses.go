package server

import (
	"encoding/json"
	"net/http"
	"os/exec"
)

func (s *Server) handleSystemProcesses() http.HandlerFunc {
	type SystemService struct {
		Name     string `json:"name"`
		IsActive bool   `json:"isActive"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		services := []string{
			"php-fpm",
			"mariadb",
			"fail2ban",
			// "NetworkManager",
			"httpd",
		}

		servicesJSON := []SystemService{}

		for _, service := range services {
			sJSON := SystemService{
				Name:     service,
				IsActive: false,
			}
			cmd := exec.Command("systemctl", "is-active", service)
			err := cmd.Run()
			if err == nil {
				sJSON.IsActive = true
			}
			servicesJSON = append(servicesJSON, sJSON)
		}

		outputBytes, _ := json.Marshal(servicesJSON)
		w.Write(outputBytes)
	}
}
