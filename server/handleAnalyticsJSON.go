package server

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
)

func (s *Server) handleAnalyticsJSON() http.HandlerFunc {
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

		chartData, err := processAnalyticDataForAccount(a, t)
		if err != nil {
			sendJSONError(w, http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
			return
		}

		// start making our output
		data := chartJson{
			Legend: chartLegend{
				Enabled: false,
			},
			XAxis: Axis{
				Categories: chartData.outputLabels,
			},
			YAxis: Axis{
				Title: chartTitle{
					Text: chartData.title,
				},
				Labels: axisLabels{
					Format: chartData.yAxisLabelFormat,
				},
			},
			Series: []chartSeries{
				{
					Data: chartData.outputValues,
					Type: chartData.chartType,
					Name: chartData.seriesName,
					Tooltip: chartTooltip{
						Suffix:   chartData.tooltipSuffix,
						Prefix:   chartData.tooltipPrefix,
						Decimals: chartData.tooltipDecimals,
					},
				},
			},
			Title: chartTitle{
				Text: chartData.title,
			},
		}

		bytes, _ := json.Marshal(data)
		w.Write(bytes)
	}
}

func processAnalyticDataForAccount(account string, dataToFetch string) (processedChartData, error) {
	chartData := processedChartData{
		outputValues:     []float64{},
		outputLabels:     []string{},
		title:            "",
		chartType:        "column",
		yAxisLabelFormat: "{value:,.0f}",
		seriesName:       "Data",
		tooltipSuffix:    "",
		tooltipPrefix:    "",
		tooltipDecimals:  0,
	}

	// TODO: look in reports folder for all json files that start with the account name
	allReports, err := os.ReadDir(os.Getenv("REPORTS_ROOT"))
	if err != nil {
		return chartData, err
	}

	for _, report := range allReports {
		_, err := report.Info()
		if err != nil {
			return chartData, err
		}

		if strings.HasPrefix(report.Name(), account) {
			reportBytes, _ := os.ReadFile(os.Getenv("REPORTS_ROOT") + report.Name())
			var reportJSON anaylticJSON
			err := json.Unmarshal(reportBytes, &reportJSON)
			if err != nil {
				continue
			}

			if dataToFetch == "total-requests" {
				chartData.outputValues = append(chartData.outputValues, float64(reportJSON.General.TotalRequests))
				chartData.title = "Total Requests"
				chartData.seriesName = "Requests"
			}

			if dataToFetch == "unique-visitors" {
				chartData.outputValues = append(chartData.outputValues, float64(reportJSON.General.UniqueVisitors))
				chartData.title = "Unique Visitors"
				chartData.seriesName = "Visitors"
			}

			if dataToFetch == "bandwidth" {
				chartData.outputValues = append(chartData.outputValues, float64(reportJSON.General.Bandwidth)/1000000)
				chartData.title = "Bandwidth"
				chartData.chartType = "line"
				chartData.yAxisLabelFormat = "{value:,.0f} MB"
				chartData.seriesName = "Bandwidth"
				chartData.tooltipSuffix = " MB"
				chartData.tooltipDecimals = 2
			}
			chartData.outputLabels = append(chartData.outputLabels, reportJSON.General.EndDate)
		}
	}

	return chartData, nil
}
