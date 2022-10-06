package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/epa-datos/bing"
	"golang.org/x/oauth2"
)

func main() {
	config := bing.AuthConfig{

		Oauth2Config: &oauth2.Config{
			ClientID:     os.Getenv("CLIENT_ID"),
			ClientSecret: os.Getenv("CLIENT_SECRET"),
		},
		Oauth2Token: &oauth2.Token{
			//Use the credentials in order to connect via API with the given customer
			AccessToken:  os.Getenv("ACCESS_TOKEN"),
			RefreshToken: os.Getenv("REFRESH_TOKEN"),
		},
		// INNOVA SPORT
		// AccountId:      "141152230",
		// CustomerID:     "159920785",
		// VIVA AEROBUS
		AccountId:      "163134174",
		CustomerID:     "159920785",
		DeveloperToken: os.Getenv("DEVELOPER_TOKEN"),
	}

	auth := bing.NewAuth(config)

	reportService := bing.NewReportingService(auth)

	columns := []string{
		"TimePeriod",
		"AccountId",
		"AccountName",
		"AccountNumber",
		"CampaignId",
		"CampaignName",
		"Impressions",
		"Clicks",
		"Spend",
		"Conversions",
		"AverageCpc",
		"AbsoluteTopImpressionRatePercent",
		"Ctr",
		"TopImpressionRatePercent",
		"Revenue",
	}

	aggregation := "Daily"

	time := &bing.RequestTime{
		StartDate: "2022-09-01",
		EndDate:   "2022-10-04",
	}

	responseId, err := reportService.GetCampaignPerformanceReport(aggregation, columns, time)
	fmt.Println("RESPUESTA: ")
	parseCampaignResponse(responseId)
	fmt.Println("ERROR: ", err)

}

func parseCampaignResponse(response []*bing.CampaignPerformanceReportColumns) {

	for _, bingCampaignMetric := range response {
		metric := &bing.CampaignMetric{
			AccountID:                        bingCampaignMetric.AccountId,
			AccountName:                      bingCampaignMetric.AccountName,
			AccountNumber:                    bingCampaignMetric.AccountNumber,
			CampaignID:                       bingCampaignMetric.CampaignId,
			CampaignName:                     bingCampaignMetric.CampaignName,
			Date:                             bingCampaignMetric.TimePeriod,
			Clicks:                           bingCampaignMetric.Clicks,
			Impressions:                      bingCampaignMetric.Impressions,
			Ctr:                              bingCampaignMetric.Ctr,
			AverageCpc:                       bingCampaignMetric.AverageCpc,
			Spend:                            bingCampaignMetric.Spend,
			Conversions:                      float64(bingCampaignMetric.Conversions),
			TopImpressionRatePercent:         bingCampaignMetric.TopImpressionRatePercent,
			AbsoluteTopImpressionRatePercent: bingCampaignMetric.AbsoluteTopImpressionRatePercent,
			Revenue:                          bingCampaignMetric.Revenue,
		}

		jsonE, _ := json.Marshal(metric)
		fmt.Println(string(jsonE))
	}

}
