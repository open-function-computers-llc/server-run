package server

import (
	"encoding/json"
	"net/http"
)

func (s *Server) handleAnalyticsJSON() http.HandlerFunc {
	type chartSeries struct {
		Data []int  `json:"data"`
		Type string `json:"type"`
	}
	type chartJson struct {
		Series []chartSeries `json:"series"`
	}

	return func(w http.ResponseWriter, r *http.Request) {

		r.ParseForm()

		a := r.FormValue("account")
		t := r.FormValue("type")

		if a == "" || t == "" {
			sendJSONError(w, http.StatusBadRequest, map[string]string{
				"error": "the query params `account` and `type` are both required",
			})
			return
		}

		// TODO: look in reports folder for all json files that start with the account name
		// filename := os.Getenv("REPORTS_ROOT") + a + tp + ".html"

		data := chartJson{
			Series: []chartSeries{
				{
					Data: []int{33, 44, 87, 92},
					Type: "bar",
				},
			},
		}

		output, _ := json.Marshal(data)
		w.Write(output)
	}
}
