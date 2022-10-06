package server

import (
	"net/http"

	"github.com/open-function-computers-llc/server-run/website"
)

func (s *Server) handleDetails() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		d := r.FormValue("domain") // TODO: switch this in ANGULAR to "account"
		var site website.Site

		for _, s := range s.sites {
			if s.Account == d {
				site = s
			}
		}

		if site.Account == "" {
			// handle error
			return
		}

		// reload the status each time the page is requested
		site.LoadStatus()

		sendJSON(w, site)
	}
}
