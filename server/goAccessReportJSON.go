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
}
