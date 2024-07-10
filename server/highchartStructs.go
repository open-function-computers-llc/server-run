package server

type chartSeries struct {
	Data    []float64    `json:"data"`
	Type    string       `json:"type"`
	Name    string       `json:"name"`
	Tooltip chartTooltip `json:"tooltip"`
}
type chartTooltip struct {
	Suffix   string `json:"valueSuffix"`
	Prefix   string `json:"valuePrefix"`
	Decimals int    `json:"valueDecimals"`
}
type chartTitle struct {
	Text string `json:"text"`
}
type axisLabels struct {
	Format string `json:"format"`
}
type Axis struct {
	Categories []string   `json:"categories"`
	Title      chartTitle `json:"title"`
	Labels     axisLabels `json:"labels"`
}
type chartLegend struct {
	Enabled bool `json:"enabled"`
}
type chartJson struct {
	Series []chartSeries `json:"series"`
	XAxis  Axis          `json:"xAxis"`
	YAxis  Axis          `json:"yAxis"`
	Title  chartTitle    `json:"title"`
	Legend chartLegend   `json:"legend"`
}
type thirdPartyAnalyticsJSON struct {
	Values   []float64 `json:"values"`
	Min      float64   `json:"min"`
	Max      float64   `json:"max"`
	Start    string    `json:"start"`
	End      string    `json:"end"`
	Average  float64   `json:"average"`
	Total    float64   `json:"total"`
	Duration int       `json:"duration"`
}

type processedChartData struct {
	outputValues     []float64
	outputLabels     []string
	title            string
	chartType        string
	yAxisLabelFormat string
	seriesName       string
	tooltipSuffix    string
	tooltipPrefix    string
	tooltipDecimals  int
}
