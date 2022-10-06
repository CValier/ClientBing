package bing

import (
	"testing"

	"github.com/epa-datos/bing"
	"github.com/go-playground/assert/v2"
)

func TestParseDate_Positive(t *testing.T) {
	date := "2022-01-01"
	expectedResponse := &bing.Date{
		Year:  2022,
		Month: 1,
		Day:   1,
	}
	response := bing.ParseDate(date)
	assert.Equal(t, response, expectedResponse)
}

func TestGetAdGroupReport_Positive(t *testing.T) {
	report := []*bing.ReportRecord{
		{
			Record: map[string]string{
				"TimePeriod": "2022-01-01",
				"Clicks":     "100",
				"Revenue":    "1000",
			},
		},
	}

	expectedResponse := []*bing.AdGroupPerformanceReportColumns{
		{
			TimePeriod: "2022-01-01",
			Clicks:     100,
			Revenue:    1000,
		},
	}

	response := bing.GetAdGroupReport(report)
	assert.Equal(t, response, expectedResponse)
}

func TestGetCampaignReport_Positive(t *testing.T) {
	report := []*bing.ReportRecord{
		{
			Record: map[string]string{
				"TimePeriod": "2022-01-01",
				"Clicks":     "100",
				"Revenue":    "1000",
			},
		},
	}

	expectedResponse := []*bing.CampaignPerformanceReportColumns{
		{
			TimePeriod: "2022-01-01",
			Clicks:     100,
			Revenue:    "1000",
		},
	}

	response := bing.GetCampaignReport(report)
	assert.Equal(t, response, expectedResponse)
}

func TestGetAccountReport_Positive(t *testing.T) {
	report := []*bing.ReportRecord{
		{
			Record: map[string]string{
				"TimePeriod": "2022-01-01",
				"Clicks":     "100",
				"Revenue":    "1000",
			},
		},
	}

	expectedResponse := []*bing.AccountPerformanceReportColumns{
		{
			TimePeriod: "2022-01-01",
			Clicks:     100,
			Revenue:    1000,
		},
	}

	response := bing.GetAccountReport(report)
	assert.Equal(t, response, expectedResponse)
}
