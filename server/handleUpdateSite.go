package server

import (
	"net/http"
)

func (s *Server) handleUpdateSite() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(1024)

		d := r.PostFormValue("domain")
		uri := r.PostFormValue("uri")

		site, err := s.findSiteByDomain(d)
		if err != nil {
			sendJSONError(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}

		switch r.PostFormValue("action") {
		case "update-uptime-uri":
			err = site.SetUptimeURI(uri)
		}
		if err != nil {
			sendJSONError(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}

		sendJSON(w, map[string]string{
			"message": "Updated URI to " + uri,
		})
	}
}
