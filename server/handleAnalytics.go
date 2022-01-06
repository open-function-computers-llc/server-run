package server

import (
	"net/http"
	"os"
)

func (s *Server) handleAnalytics() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		r.ParseForm()

		d := r.FormValue("domain")
		tp := r.FormValue("period")

		if d == "" || tp == "" {
			sendJSONError(w, http.StatusBadRequest, map[string]string{
				"error": "the query params `domain` and `period` are both required",
			})
			return
		}

		switch tp {
		case "1":
			tp = ".day"
		case "30":
			tp = ".month"
		default:
			tp = ""
		}

		filename := os.Getenv("REPORTS_ROOT") + d + tp + ".html"
		content, err := os.ReadFile(filename)
		if err != nil {
			sendJSONError(w, http.StatusInternalServerError, map[string]string{
				"error": "There was an error reading the file: " + err.Error(),
			})
			return
		}

		w.Write(content)
	}
}
