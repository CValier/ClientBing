package bing

import (
	"archive/zip"
	"bufio"
	"bytes"
	"encoding/csv"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type ReportingService struct {
	Endpoint string
	Session  *Session
}

func NewReportingService(session *Session) *ReportingService {
	return &ReportingService{
		Endpoint: "https://api.bingads.microsoft.com/Api/Advertiser/Reporting/v13/ReportingService.svc",
		Session:  session,
	}
}

// SubmitReportRequest generates the bing report and returns the reportrRequestId.
func (c *ReportingService) SubmitReportRequest(rr interface{}) (string, error) {
	req := SubmitGenerateReportRequest{
		ReportRequest: rr,
		NS:            "https://bingads.microsoft.com/Reporting/v13",
	}

	resp, err := c.Session.reportRequest(req, c.Endpoint, "SubmitGenerateReport")
	if err != nil {
		return "", errors.New("Error SubmitReportRequest: " + err.Error())
	}

	// Validate if the response failed because the token has expired
	if strings.Contains(string(resp), "token expired") {
		err := c.Session.refreshToken()
		if err != nil {
			return "", errors.New("Error SubmitReportRequest: " + err.Error())
		}
		// Repeat the request with the renewed token.
		resp, err = c.Session.reportRequest(req, c.Endpoint, "SubmitGenerateReport")
		if err != nil {
			return "", errors.New("Error SubmitReportRequest: " + err.Error())
		}
	}

	// Unmarshal response to obtain the generated id.
	reportR := SubmitGenerateReportResponse{}
	err = xml.Unmarshal(resp, &reportR)

	if err != nil {
		return "", errors.New("Error unmarshaling submitReportRequest response: " + err.Error())
	}

	return reportR.ReportRequestId, nil

}

// PollGenerateReport polls for the report to check if the report is ready and returns the url and status.
func (c *ReportingService) PollGenerateReport(id string) (*ReportRequestStatus, error) {
	req := PollGenerateReportRequest{
		ReportRequestId: id,
		NS:              "https://bingads.microsoft.com/Reporting/v13",
	}
	resp, err := c.Session.reportRequest(req, c.Endpoint, "PollGenerateReport")
	if err != nil {
		return nil, errors.New("Error PollGenerateReport request: " + err.Error())
	}

	// Unmarshal the reponse to get the report url and the status
	reportR := PollGenerateReportResponse{}
	err = xml.Unmarshal(resp, &reportR)
	if err != nil {
		return nil, errors.New("Error unmarshaling SubmitReportRequest response: " + err.Error())
	}
	return &reportR.ReportRequestStatus, nil
}

// PollReport polls for the requested report until the report status is success.
func (c *ReportingService) PollReport(id string) (string, error) {
	// Validate that the id is not empty to avoid polling report loop
	if id == "" {
		return "", errors.New("PollReport: ReportID was not generated successfully ")
	}
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)

		// Polling the report
		report, err := c.PollGenerateReport(id)
		if err != nil {
			return "", errors.New("Error polling the report: " + err.Error())
		}
		// Verify the status.
		switch report.Status {
		case "Success":
			return report.ReportDownloadUrl, nil
		case "Error":
			return "", errors.New("Error polling report, reportStatus = 'Error': " + err.Error())
		default:
		}
	}
	return "", errors.New("timeout. Poll report process took more than 10 seconds. ")
}

// GetAdGroupPerformanceReport returns the AdGroup report with the specified columns and daterange.
func (c *ReportingService) GetAdGroupPerformanceReport(aggregation string, columns []string, requestTime *RequestTime) ([]*AdGroupPerformanceReportColumns, error, []string, [][]string) {

	// Parsing data to use it in report request
	accountId, _ := strconv.ParseInt(c.Session.AccountId, 10, 64)
	time := ReportTime{
		PredefinedTime:       requestTime.PredefinedTime,
		CustomDateRangeStart: *ParseDate(requestTime.StartDate),
		CustomDateRangeEnd:   *ParseDate(requestTime.EndDate),
	}

	// Building report request
	rr := &AdGroupPerformanceReportRequest{
		Scope: ReportScope{
			AccountIds: Longs{
				accountId,
			},
		},
		Aggregation: aggregation,
		Columns:     columns,
		Time:        time,
	}

	// Generate performance report
	report, err, cols, rows := c.GetPerformanceReport(rr)
	if err != nil {
		return nil, err, nil, nil
	}

	// convert map to adGroup report.
	return GetAdGroupReport(report), nil, cols, rows
}

// GetCampaignPerformanceReport returns a campaign report with the specified columns and daterange.
func (c *ReportingService) GetCampaignPerformanceReport(aggregation string, columns []string, requestTime *RequestTime) ([]*CampaignPerformanceReportColumns, error, []string, [][]string) {

	// Parsing data to use it in report request
	accountId, _ := strconv.ParseInt(c.Session.AccountId, 10, 64)
	time := ReportTime{
		PredefinedTime:       requestTime.PredefinedTime,
		CustomDateRangeStart: *ParseDate(requestTime.StartDate),
		CustomDateRangeEnd:   *ParseDate(requestTime.EndDate),
	}

	// Building report request
	rr := &CampaignPerformanceReportRequest{
		Scope: ReportScope{
			AccountIds: Longs{
				accountId,
			},
		},
		Aggregation: aggregation,
		Columns:     columns,
		Time:        time,
	}

	// Generate performance report
	report, err, cols, rows := c.GetPerformanceReport(rr)
	if err != nil {
		return nil, err, nil, nil
	}

	// convert map to Campaign report.
	return GetCampaignReport(report), nil, cols, rows
}

// GetAccountPerformanceReport returns the account report with the specified columns and daterange.
func (c *ReportingService) GetAccountPerformanceReport(aggregation string, columns []string, requestTime *RequestTime) ([]*AccountPerformanceReportColumns, error, []string, [][]string) {

	// Parsing data to use it in report request
	accountId, _ := strconv.ParseInt(c.Session.AccountId, 10, 64)
	time := ReportTime{
		PredefinedTime:       requestTime.PredefinedTime,
		CustomDateRangeStart: *ParseDate(requestTime.StartDate),
		CustomDateRangeEnd:   *ParseDate(requestTime.EndDate),
	}

	// Building Report Request
	rr := &AccountPerformanceReportRequest{
		Scope: ReportScope{
			AccountIds: Longs{
				accountId,
			},
		},
		Aggregation: aggregation,
		Columns:     columns,
		Time:        time,
	}

	// Generate performance report
	report, err, cols, rows := c.GetPerformanceReport(rr)
	if err != nil {
		return nil, err, nil, nil
	}

	// convert map to Account report.
	return GetAccountReport(report), nil, cols, rows
}

// getPerformanceReport submits the report request, polls it and returns the report as a map
func (c *ReportingService) GetPerformanceReport(rr interface{}) ([]*ReportRecord, error, []string, [][]string) {
	// Submit the report and generate a report id.
	id, err := c.SubmitReportRequest(rr)
	if err != nil {
		return nil, errors.New("Error submiting the report: " + err.Error()), nil, nil
	}

	// Poll the report and get the url.
	url, err := c.PollReport(id)
	if err != nil {
		return nil, errors.New("Error polling the report: " + err.Error()), nil, nil
	}

	// Use the url to generate a map with the report.
	report, err, cols, rows := c.GetReportMap(url)
	if err != nil {
		return nil, errors.New("Error getting report map: " + err.Error()), nil, nil
	}
	return report, nil, cols, rows
}

// GetReportMap downloads the report, reads the csv file and returns a map with the report metrics.
func (c *ReportingService) GetReportMap(url string) ([]*ReportRecord, error, []string, [][]string) {
	//Verifying url is not empty
	if url == "" {
		return nil, errors.New("Empty Url Report, unable to find data for that period."), nil, nil
	}

	// Download ZIP file from the url
	res, err := http.DefaultClient.Get(url)
	if err != nil {
		return nil, errors.New("Error downloading report zip file: " + err.Error()), nil, nil
	}

	// Opening body of the response
	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("Error opening the zip file response: " + err.Error()), nil, nil
	}

	// Read the zip file and extract the csv file
	breader := bytes.NewReader(b)
	reader, err := zip.NewReader(breader, breader.Size())
	if err != nil {
		return nil, errors.New("Error reading the zip file: " + err.Error()), nil, nil
	}

	// Verify that exists 1 csv file
	if len(reader.File) != 1 {
		return nil, fmt.Errorf("Error reading csv file, expected 1 file, got: %d", len(reader.File)), nil, nil
	}

	// Open the csv file.
	f, err := reader.File[0].Open()
	if err != nil {
		return nil, errors.New("Error reading the csv file: " + err.Error()), nil, nil
	}

	// Reading the csv file
	defer f.Close()
	br := bufio.NewReader(f)

	// Reading the first line to extract columns.
	line, _, err := br.ReadLine()
	if err != nil {
		return nil, errors.New("Error reading the csv report columns: " + err.Error()), nil, nil
	}

	// Getting list of columns
	cols := strings.Split(strings.Replace(string(line), "\"", "", -1), ",")
	rows := [][]string{}
	// Reading other csv lines and storing metrics on a map
	report := []*ReportRecord{}
	csvLines, _ := csv.NewReader(br).ReadAll()
	for _, line := range csvLines {
		row := []string{}
		newRecord := &ReportRecord{
			Record: map[string]string{},
		}
		for j, metric := range line {
			newRecord.Record[cols[j]] = metric
			row = append(row, metric)
		}
		report = append(report, newRecord)
		rows = append(rows, row)
	}

	return report, nil, cols, rows
}
