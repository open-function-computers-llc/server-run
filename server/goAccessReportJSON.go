package server

type anaylticJSON struct {
	General struct {
		StartDate         string   `json:"start_date"`
		EndDate           string   `json:"end_date"`
		DateTime          string   `json:"date_time"`
		TotalRequests     int      `json:"total_requests"`
		ValidRequests     int      `json:"valid_requests"`
		FailedRequests    int      `json:"failed_requests"`
		GenerationTime    int      `json:"generation_time"`
		UniqueVisitors    int      `json:"unique_visitors"`
		UniqueFiles       int      `json:"unique_files"`
		ExcludedHits      int      `json:"excluded_hits"`
		UniqueReferrers   int      `json:"unique_referrers"`
		UniqueNotFound    int      `json:"unique_not_found"`
		UniqueStaticFiles int      `json:"unique_static_files"`
		LogSize           int      `json:"log_size"`
		Bandwidth         int      `json:"bandwidth"`
		LogPath           []string `json:"log_path"`
	} `json:"general"`
	Visitors struct {
		Metadata struct {
			Bytes struct {
				Count int `json:"count"`
				Max   int `json:"max"`
				Min   int `json:"min"`
			} `json:"bytes"`
			Visitors struct {
				Count int `json:"count"`
				Max   int `json:"max"`
				Min   int `json:"min"`
			} `json:"visitors"`
			Hits struct {
				Count int `json:"count"`
				Max   int `json:"max"`
				Min   int `json:"min"`
			} `json:"hits"`
			Data struct {
				Unique int `json:"unique"`
			} `json:"data"`
		} `json:"metadata"`
		Data []struct {
			Hits struct {
				Count   int    `json:"count"`
				Percent string `json:"percent"`
			} `json:"hits"`
			Visitors struct {
				Count   int    `json:"count"`
				Percent string `json:"percent"`
			} `json:"visitors"`
			Bytes struct {
				Count   int    `json:"count"`
				Percent string `json:"percent"`
			} `json:"bytes"`
			Data string `json:"data"`
		} `json:"data"`
	} `json:"visitors"`
	Requests struct {
		Metadata struct {
			Bytes struct {
				Count int `json:"count"`
				Max   int `json:"max"`
				Min   int `json:"min"`
			} `json:"bytes"`
			Visitors struct {
				Count int `json:"count"`
				Max   int `json:"max"`
				Min   int `json:"min"`
			} `json:"visitors"`
			Hits struct {
				Count int `json:"count"`
				Max   int `json:"max"`
				Min   int `json:"min"`
			} `json:"hits"`
			Data struct {
				Unique int `json:"unique"`
			} `json:"data"`
		} `json:"metadata"`
		Data []struct {
			Hits struct {
				Count   int    `json:"count"`
				Percent string `json:"percent"`
			} `json:"hits"`
			Visitors struct {
				Count   int    `json:"count"`
				Percent string `json:"percent"`
			} `json:"visitors"`
			Bytes struct {
				Count   int    `json:"count"`
				Percent string `json:"percent"`
			} `json:"bytes"`
			Method   string `json:"method"`
			Protocol string `json:"protocol"`
			Data     string `json:"data"`
		} `json:"data"`
	} `json:"requests"`
	StaticRequests struct {
		Metadata struct {
			Bytes struct {
				Count int `json:"count"`
				Max   int `json:"max"`
				Min   int `json:"min"`
			} `json:"bytes"`
			Visitors struct {
				Count int `json:"count"`
				Max   int `json:"max"`
				Min   int `json:"min"`
			} `json:"visitors"`
			Hits struct {
				Count int `json:"count"`
				Max   int `json:"max"`
				Min   int `json:"min"`
			} `json:"hits"`
			Data struct {
				Unique int `json:"unique"`
			} `json:"data"`
		} `json:"metadata"`
		Data []struct {
			Hits struct {
				Count   int    `json:"count"`
				Percent string `json:"percent"`
			} `json:"hits"`
			Visitors struct {
				Count   int    `json:"count"`
				Percent string `json:"percent"`
			} `json:"visitors"`
			Bytes struct {
				Count   int    `json:"count"`
				Percent string `json:"percent"`
			} `json:"bytes"`
			Method   string `json:"method"`
			Protocol string `json:"protocol"`
			Data     string `json:"data"`
		} `json:"data"`
	} `json:"static_requests"`
	NotFound struct {
		Metadata struct {
			Bytes struct {
				Count int `json:"count"`
				Max   int `json:"max"`
				Min   int `json:"min"`
			} `json:"bytes"`
			Visitors struct {
				Count int `json:"count"`
				Max   int `json:"max"`
				Min   int `json:"min"`
			} `json:"visitors"`
			Hits struct {
				Count int `json:"count"`
				Max   int `json:"max"`
				Min   int `json:"min"`
			} `json:"hits"`
			Data struct {
				Unique int `json:"unique"`
			} `json:"data"`
		} `json:"metadata"`
		Data []struct {
			Hits struct {
				Count   int    `json:"count"`
				Percent string `json:"percent"`
			} `json:"hits"`
			Visitors struct {
				Count   int    `json:"count"`
				Percent string `json:"percent"`
			} `json:"visitors"`
			Bytes struct {
				Count   int    `json:"count"`
				Percent string `json:"percent"`
			} `json:"bytes"`
			Method   string `json:"method"`
			Protocol string `json:"protocol"`
			Data     string `json:"data"`
		} `json:"data"`
	} `json:"not_found"`
	Hosts struct {
		Metadata struct {
			Bytes struct {
				Count int `json:"count"`
				Max   int `json:"max"`
				Min   int `json:"min"`
			} `json:"bytes"`
			Visitors struct {
				Count int `json:"count"`
				Max   int `json:"max"`
				Min   int `json:"min"`
			} `json:"visitors"`
			Hits struct {
				Count int `json:"count"`
				Max   int `json:"max"`
				Min   int `json:"min"`
			} `json:"hits"`
			Data struct {
				Unique int `json:"unique"`
			} `json:"data"`
		} `json:"metadata"`
		Data []struct {
			Hits struct {
				Count   int    `json:"count"`
				Percent string `json:"percent"`
			} `json:"hits"`
			Visitors struct {
				Count   int    `json:"count"`
				Percent string `json:"percent"`
			} `json:"visitors"`
			Bytes struct {
				Count   int    `json:"count"`
				Percent string `json:"percent"`
			} `json:"bytes"`
			Data    string   `json:"data"`
			Country string   `json:"country"`
			Items   []string `json:"items"`
		} `json:"data"`
	} `json:"hosts"`
	Os struct {
		Metadata struct {
			Bytes struct {
				Count int `json:"count"`
				Max   int `json:"max"`
				Min   int `json:"min"`
			} `json:"bytes"`
			Visitors struct {
				Count int `json:"count"`
				Max   int `json:"max"`
				Min   int `json:"min"`
			} `json:"visitors"`
			Hits struct {
				Count int `json:"count"`
				Max   int `json:"max"`
				Min   int `json:"min"`
			} `json:"hits"`
			Data struct {
				Unique int `json:"unique"`
			} `json:"data"`
		} `json:"metadata"`
		Data []struct {
			Hits struct {
				Count   int    `json:"count"`
				Percent string `json:"percent"`
			} `json:"hits"`
			Visitors struct {
				Count   int    `json:"count"`
				Percent string `json:"percent"`
			} `json:"visitors"`
			Bytes struct {
				Count   int    `json:"count"`
				Percent string `json:"percent"`
			} `json:"bytes"`
			Data  string `json:"data"`
			Items []struct {
				Hits struct {
					Count   int    `json:"count"`
					Percent string `json:"percent"`
				} `json:"hits"`
				Visitors struct {
					Count   int    `json:"count"`
					Percent string `json:"percent"`
				} `json:"visitors"`
				Bytes struct {
					Count   int    `json:"count"`
					Percent string `json:"percent"`
				} `json:"bytes"`
				Data string `json:"data"`
			} `json:"items"`
		} `json:"data"`
	} `json:"os"`
	Browsers struct {
		Metadata struct {
			Bytes struct {
				Count int `json:"count"`
				Max   int `json:"max"`
				Min   int `json:"min"`
			} `json:"bytes"`
			Visitors struct {
				Count int `json:"count"`
				Max   int `json:"max"`
				Min   int `json:"min"`
			} `json:"visitors"`
			Hits struct {
				Count int `json:"count"`
				Max   int `json:"max"`
				Min   int `json:"min"`
			} `json:"hits"`
			Data struct {
				Unique int `json:"unique"`
			} `json:"data"`
		} `json:"metadata"`
		Data []struct {
			Hits struct {
				Count   int    `json:"count"`
				Percent string `json:"percent"`
			} `json:"hits"`
			Visitors struct {
				Count   int    `json:"count"`
				Percent string `json:"percent"`
			} `json:"visitors"`
			Bytes struct {
				Count   int    `json:"count"`
				Percent string `json:"percent"`
			} `json:"bytes"`
			Data  string `json:"data"`
			Items []struct {
				Hits struct {
					Count   int    `json:"count"`
					Percent string `json:"percent"`
				} `json:"hits"`
				Visitors struct {
					Count   int    `json:"count"`
					Percent string `json:"percent"`
				} `json:"visitors"`
				Bytes struct {
					Count   int    `json:"count"`
					Percent string `json:"percent"`
				} `json:"bytes"`
				Data string `json:"data"`
			} `json:"items"`
		} `json:"data"`
	} `json:"browsers"`
	VisitTime struct {
		Metadata struct {
			Bytes struct {
				Count int `json:"count"`
				Max   int `json:"max"`
				Min   int `json:"min"`
			} `json:"bytes"`
			Visitors struct {
				Count int `json:"count"`
				Max   int `json:"max"`
				Min   int `json:"min"`
			} `json:"visitors"`
			Hits struct {
				Count int `json:"count"`
				Max   int `json:"max"`
				Min   int `json:"min"`
			} `json:"hits"`
			Data struct {
				Unique int `json:"unique"`
			} `json:"data"`
		} `json:"metadata"`
		Data []struct {
			Hits struct {
				Count   int    `json:"count"`
				Percent string `json:"percent"`
			} `json:"hits"`
			Visitors struct {
				Count   int    `json:"count"`
				Percent string `json:"percent"`
			} `json:"visitors"`
			Bytes struct {
				Count   int    `json:"count"`
				Percent string `json:"percent"`
			} `json:"bytes"`
			Data string `json:"data"`
		} `json:"data"`
	} `json:"visit_time"`
	ReferringSites struct {
		Metadata struct {
			Bytes struct {
				Count int `json:"count"`
				Max   int `json:"max"`
				Min   int `json:"min"`
			} `json:"bytes"`
			Visitors struct {
				Count int `json:"count"`
				Max   int `json:"max"`
				Min   int `json:"min"`
			} `json:"visitors"`
			Hits struct {
				Count int `json:"count"`
				Max   int `json:"max"`
				Min   int `json:"min"`
			} `json:"hits"`
			Data struct {
				Unique int `json:"unique"`
			} `json:"data"`
		} `json:"metadata"`
		Data []struct {
			Hits struct {
				Count   int    `json:"count"`
				Percent string `json:"percent"`
			} `json:"hits"`
			Visitors struct {
				Count   int    `json:"count"`
				Percent string `json:"percent"`
			} `json:"visitors"`
			Bytes struct {
				Count   int    `json:"count"`
				Percent string `json:"percent"`
			} `json:"bytes"`
			Data string `json:"data"`
		} `json:"data"`
	} `json:"referring_sites"`
	StatusCodes struct {
		Metadata struct {
			Bytes struct {
				Count int `json:"count"`
				Max   int `json:"max"`
				Min   int `json:"min"`
			} `json:"bytes"`
			Visitors struct {
				Count int `json:"count"`
				Max   int `json:"max"`
				Min   int `json:"min"`
			} `json:"visitors"`
			Hits struct {
				Count int `json:"count"`
				Max   int `json:"max"`
				Min   int `json:"min"`
			} `json:"hits"`
			Data struct {
				Unique int `json:"unique"`
			} `json:"data"`
		} `json:"metadata"`
		Data []struct {
			Hits struct {
				Count   int    `json:"count"`
				Percent string `json:"percent"`
			} `json:"hits"`
			Visitors struct {
				Count   int    `json:"count"`
				Percent string `json:"percent"`
			} `json:"visitors"`
			Bytes struct {
				Count   int    `json:"count"`
				Percent string `json:"percent"`
			} `json:"bytes"`
			Data  string `json:"data"`
			Items []struct {
				Hits struct {
					Count   int    `json:"count"`
					Percent string `json:"percent"`
				} `json:"hits"`
				Visitors struct {
					Count   int    `json:"count"`
					Percent string `json:"percent"`
				} `json:"visitors"`
				Bytes struct {
					Count   int    `json:"count"`
					Percent string `json:"percent"`
				} `json:"bytes"`
				Data string `json:"data"`
			} `json:"items"`
		} `json:"data"`
	} `json:"status_codes"`
	Geolocation struct {
		Metadata struct {
			Bytes struct {
				Count int `json:"count"`
				Max   int `json:"max"`
				Min   int `json:"min"`
			} `json:"bytes"`
			Visitors struct {
				Count int `json:"count"`
				Max   int `json:"max"`
				Min   int `json:"min"`
			} `json:"visitors"`
			Hits struct {
				Count int `json:"count"`
				Max   int `json:"max"`
				Min   int `json:"min"`
			} `json:"hits"`
			Data struct {
				Unique int `json:"unique"`
			} `json:"data"`
		} `json:"metadata"`
		Data []struct {
			Hits struct {
				Count   int    `json:"count"`
				Percent string `json:"percent"`
			} `json:"hits"`
			Visitors struct {
				Count   int    `json:"count"`
				Percent string `json:"percent"`
			} `json:"visitors"`
			Bytes struct {
				Count   int    `json:"count"`
				Percent string `json:"percent"`
			} `json:"bytes"`
			Data  string `json:"data"`
			Items []struct {
				Hits struct {
					Count   int    `json:"count"`
					Percent string `json:"percent"`
				} `json:"hits"`
				Visitors struct {
					Count   int    `json:"count"`
					Percent string `json:"percent"`
				} `json:"visitors"`
				Bytes struct {
					Count   int    `json:"count"`
					Percent string `json:"percent"`
				} `json:"bytes"`
				Data string `json:"data"`
			} `json:"items"`
		} `json:"data"`
	} `json:"geolocation"`
}
