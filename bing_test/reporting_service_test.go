package bing_test

import (
	"context"
	"testing"

	"github.com/epa-datos/bing"
	"github.com/go-playground/assert/v2"
	"github.com/jarcoal/httpmock"
	"golang.org/x/oauth2"
)

var Oauth2Config = &oauth2.Config{
	ClientID:     "1234567",
	ClientSecret: "123534645",
}

var Oauth2Token = &oauth2.Token{
	AccessToken:  "1234567",
	RefreshToken: "1234567",
}
var ctx = context.TODO()
var session = &bing.Session{
	AccountId:      "12345",
	CustomerId:     "123456",
	DeveloperToken: "nmqiwdupbqwoensdn",
	ClientID:       "1234567",
	ClientSecret:   "123534645",
	HTTPClient:     Oauth2Config.Client(ctx, Oauth2Token),
	TokenSource: Oauth2Config.TokenSource(context.TODO(), &oauth2.Token{
		AccessToken:  Oauth2Token.AccessToken,
		RefreshToken: Oauth2Token.RefreshToken,
	}),
}

var testReportingSvc = bing.NewReportingService(session)

func TestSubmitReportRequest_Positive(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// expected response
	expectedId := "testingReportID"
	// Mock http response
	httpmock.RegisterResponder(
		"POST",
		"https://api.bingads.microsoft.com/Api/Advertiser/Reporting/v13/ReportingService.svc",
		httpmock.NewStringResponder(
			200, // Expected http code status
			`<s:Envelope xmlns:s="http://schemas.xmlsoap.org/soap/envelope/">
				<s:Header xmlns="https://bingads.microsoft.com/Reporting/v13">
			  		<TrackingId d3p1:nil="false" xmlns:d3p1="http://www.w3.org/2001/XMLSchema-instance">12345</TrackingId>
				</s:Header>
				<s:Body>
			  		<SubmitGenerateReportResponse xmlns="https://bingads.microsoft.com/Reporting/v13">
						<ReportRequestId d4p1:nil="false" xmlns:d4p1="http://www.w3.org/2001/XMLSchema-instance">testingReportID</ReportRequestId>
			  		</SubmitGenerateReportResponse>
				</s:Body>
		  	</s:Envelope>`,
		))

	rr := &bing.AdGroupPerformanceReportRequest{
		Scope: bing.ReportScope{
			AccountIds: bing.Longs{
				12345678,
			},
		},
		Aggregation: "Daily",
		Columns:     []string{"Clicks", "Investment"},
		Time: bing.ReportTime{
			PredefinedTime: "LastMonth",
		},
	}

	responseId, err := testReportingSvc.SubmitReportRequest(rr)
	assert.Equal(t, responseId, expectedId)
	assert.Equal(t, err, nil)
}

func TestPollGenerateReport_Positive(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	expectedReport := &bing.ReportRequestStatus{
		ReportDownloadUrl: "testingUrl",
		Status:            "Success",
	}
	httpmock.RegisterResponder(
		"POST",
		"https://api.bingads.microsoft.com/Api/Advertiser/Reporting/v13/ReportingService.svc",
		httpmock.NewStringResponder(
			200, // Expected http code status
			`<s:Envelope xmlns:s="http://schemas.xmlsoap.org/soap/envelope/">
				<s:Header xmlns="https://bingads.microsoft.com/Reporting/v13">
			  <TrackingId d3p1:nil="false" xmlns:d3p1="http://www.w3.org/2001/XMLSchema-instance">12345</TrackingId>
			</s:Header>
			<s:Body>
			  <PollGenerateReportResponse xmlns="https://bingads.microsoft.com/Reporting/v13">
				<ReportRequestStatus d4p1:nil="false" xmlns:d4p1="http://www.w3.org/2001/XMLSchema-instance">
				  <ReportDownloadUrl d4p1:nil="false">testingUrl</ReportDownloadUrl>
				  <Status>Success</Status>
				</ReportRequestStatus>
			  </PollGenerateReportResponse>
			</s:Body>
		  </s:Envelope>`,
		))
	responseReport, err := testReportingSvc.PollGenerateReport("TestingID")
	assert.Equal(t, expectedReport, responseReport)
	assert.Equal(t, err, nil)
}

func TestPollReport_Positive(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	expectedUrl := "testingUrl"
	httpmock.RegisterResponder(
		"POST",
		"https://api.bingads.microsoft.com/Api/Advertiser/Reporting/v13/ReportingService.svc",
		httpmock.NewStringResponder(
			200, // Expected http code status
			`<s:Envelope xmlns:s="http://schemas.xmlsoap.org/soap/envelope/">
				<s:Header xmlns="https://bingads.microsoft.com/Reporting/v13">
			  <TrackingId d3p1:nil="false" xmlns:d3p1="http://www.w3.org/2001/XMLSchema-instance">12345</TrackingId>
			</s:Header>
			<s:Body>
			  <PollGenerateReportResponse xmlns="https://bingads.microsoft.com/Reporting/v13">
				<ReportRequestStatus d4p1:nil="false" xmlns:d4p1="http://www.w3.org/2001/XMLSchema-instance">
				  <ReportDownloadUrl d4p1:nil="false">testingUrl</ReportDownloadUrl>
				  <Status>Success</Status>
				</ReportRequestStatus>
			  </PollGenerateReportResponse>
			</s:Body>
		  </s:Envelope>`,
		))
	responseUrl, err := testReportingSvc.PollReport("TestingID")
	assert.Equal(t, expectedUrl, responseUrl)
	assert.Equal(t, err, nil)
}
