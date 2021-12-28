package server

import (
	"net/http"

	"github.com/open-function-computers-llc/server-run/website"
)

func (s *Server) handleDetails() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		d := r.FormValue("domain")
		site := website.Site{}

		for _, s := range s.sites {
			if s.Domain == d {
				site = s
			}
		}

		if site.Domain == "" {
			// handle error
			return
		}

		sendJSON(w, site)
	}
}
