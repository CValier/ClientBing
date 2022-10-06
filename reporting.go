package bing

import (
	"encoding/xml"
	"strconv"
	"strings"
)

// SubmitGenerateReportRequest contains the required fields to generate the report request.
type SubmitGenerateReportRequest struct {
	XMLName       xml.Name `xml:"SubmitGenerateReportRequest"`
	NS            string   `xml:"xmlns,attr"`
	ReportRequest interface{}
}

// SubmitGenerateReportResponse handles the ID of the generated report.
type SubmitGenerateReportResponse struct {
	ReportRequestId string
}

// PollGenerateReportRequest
type PollGenerateReportRequest struct {
	XMLName         xml.Name `xml:"PollGenerateReportRequest"`
	NS              string   `xml:"xmlns,attr"`
	ReportRequestId string
}

type PollGenerateReportResponse struct {
	ReportRequestStatus ReportRequestStatus
}

// Status :: Error | Success | Pending
// ReportRequestStatus contains the report response with the url and the status.
type ReportRequestStatus struct {
	ReportDownloadUrl string
	Status            string
}

// Report Requests
type AdGroupPerformanceReportRequest PerformanceReportRequest
type CampaignPerformanceReportRequest PerformanceReportRequest
type AccountPerformanceReportRequest PerformanceReportRequest

// PerformanceReportRequest contains the structure of the report request.
type PerformanceReportRequest struct {
	XMLName     xml.Name `xml:"ReportRequest"`
	Type        string   `xml:"i:type,attr"`
	Aggregation string
	Columns     []string
	Scope       ReportScope
	Time        ReportTime
}

// ReportScope defines the scope that the report wants to cover; accounts, campaigns and adgroups
type ReportScope struct {
	XMLName    xml.Name              `xml:"Scope"`
	AccountIds Longs                 `xml:"AccountIds>long,omitempty"`
	AdGroups   []AdGroupReportScope  `xml:"AdGroups>AdGroupReportScope,omitempty"`
	Campaigns  []CampaignReportScope `xml:"Campaigns>CampaignReportScope,omitempty"`
}

type Longs []int64

// AdGroupReportScope is used in case of wanting to specify an adGroup.
type AdGroupReportScope struct {
	AccountId  int64
	CampaignId int64
	AdGroupId  int64
}

// CampaignReportScope is used in case of wanting to specify campaigns.
type CampaignReportScope struct {
	AccountId  int64
	CampaignId int64
}

// ReportTime is used to send all the date fields in xml request.
type ReportTime struct {
	XMLName              xml.Name `xml:"Time"`
	CustomDateRangeEnd   Date     `xml:",omitempty"`
	CustomDateRangeStart Date     `xml:",omitempty"`
	PredefinedTime       string   `xml:",omitempty"`
}

// Date is used to send the daterange in xml request.
type Date struct {
	Day   int
	Month int
	Year  int
}

// RequestTime is used as a parameter to specify either a date range or a predefinedTime.
type RequestTime struct {
	StartDate      string
	EndDate        string
	PredefinedTime string
}

// parseDate transforms a date from a string to a Date struct to use it in xml requests.
func ParseDate(date string) *Date {
	dateValues := strings.Split(date, "-")
	year, _ := strconv.Atoi(dateValues[0])
	month, _ := strconv.Atoi(dateValues[1])
	day, _ := strconv.Atoi(dateValues[2])
	parsedDate := &Date{
		Year:  year,
		Month: month,
		Day:   day,
	}
	return parsedDate
}

// This section allows to return a response to every report with the proper field in a struct and a proper format.
// In case of wanting to add new metrics or dimensions, it is necessary to add the switch case with the metric name
// into de corresponding report method.

// GetAdGroupReport transform raw report to adGroupPerformanceReport
func GetAdGroupReport(report []*ReportRecord) []*AdGroupPerformanceReportColumns {
	adGroupPerformanceReport := []*AdGroupPerformanceReportColumns{}
	for _, record := range report {
		adGroupPerformanceRecord := AdGroupPerformanceReportColumns{}
		for k, v := range record.Record {
			switch string(k) {
			case "AccountId":
				adGroupPerformanceRecord.AccountId = v
			case "AccountName":
				adGroupPerformanceRecord.AccountName = v
			case "AccountNumber":
				adGroupPerformanceRecord.AccountNumber = v
			case "AdGroupId":
				adGroupPerformanceRecord.AdGroupId = v
			case "AdGroupName":
				adGroupPerformanceRecord.AdGroupName = v
			case "CampaignId":
				adGroupPerformanceRecord.CampaignId = v
			case "CampaignName":
				adGroupPerformanceRecord.CampaignName = v
			case "DeviceType":
				adGroupPerformanceRecord.DeviceType = v
			case "Impressions":
				intV, _ := strconv.ParseInt(v, 0, 64)
				adGroupPerformanceRecord.Impressions = intV
			case "Clicks":
				intV, _ := strconv.ParseInt(v, 0, 64)
				adGroupPerformanceRecord.Clicks = intV
			case "Ctr":
				floatV, _ := strconv.ParseFloat(v, 64)
				adGroupPerformanceRecord.Ctr = floatV
			case "AverageCpc":
				adGroupPerformanceRecord.AverageCpc = v
			case "Spend":
				floatV, _ := strconv.ParseFloat(v, 64)
				adGroupPerformanceRecord.Spend = floatV
			case "Conversions":
				intV, _ := strconv.ParseInt(v, 0, 64)
				adGroupPerformanceRecord.Conversions = intV
			case "Revenue":
				floatV, _ := strconv.ParseFloat(v, 64)
				adGroupPerformanceRecord.Revenue = floatV
			case "TimePeriod":
				adGroupPerformanceRecord.TimePeriod = v
			default:
				if strings.Contains(k, "TimePeriod") {
					adGroupPerformanceRecord.TimePeriod = v
				}
			}
		}
		adGroupPerformanceReport = append(adGroupPerformanceReport, &adGroupPerformanceRecord)
	}
	return adGroupPerformanceReport
}

// GetCampaignReport transform raw report to CampaignPerformanceReport
func GetCampaignReport(report []*ReportRecord) []*CampaignPerformanceReportColumns {
	campaignPerformanceReport := []*CampaignPerformanceReportColumns{}
	for _, record := range report {
		campaignPerformanceRecord := CampaignPerformanceReportColumns{}
		for k, v := range record.Record {
			switch string(k) {
			case "AccountId":
				campaignPerformanceRecord.AccountId = v
			case "AccountName":
				campaignPerformanceRecord.AccountName = v
			case "AccountNumber":
				campaignPerformanceRecord.AccountNumber = v
			case "CampaignId":
				campaignPerformanceRecord.CampaignId = v
			case "CampaignName":
				campaignPerformanceRecord.CampaignName = v
			case "DeviceType":
				campaignPerformanceRecord.DeviceType = v
			case "Impressions":
				intV, _ := strconv.ParseInt(v, 0, 64)
				campaignPerformanceRecord.Impressions = intV
			case "Clicks":
				intV, _ := strconv.ParseInt(v, 0, 64)
				campaignPerformanceRecord.Clicks = intV
			case "Ctr":
				campaignPerformanceRecord.Ctr = v
			case "AverageCpc":
				campaignPerformanceRecord.AverageCpc = v
			case "Spend":
				floatV, _ := strconv.ParseFloat(v, 64)
				campaignPerformanceRecord.Spend = floatV
			case "Conversions":
				intV, _ := strconv.ParseInt(v, 0, 64)
				campaignPerformanceRecord.Conversions = intV
			case "Revenue":
				campaignPerformanceRecord.Revenue = v
			case "TimePeriod":
				campaignPerformanceRecord.TimePeriod = v
			case "AbsoluteTopImpressionRatePercent":
				campaignPerformanceRecord.AbsoluteTopImpressionRatePercent = v
			case "TopImpressionRatePercent":
				campaignPerformanceRecord.TopImpressionRatePercent = v
			default:
				if strings.Contains(k, "TimePeriod") {
					campaignPerformanceRecord.TimePeriod = v
				}
			}
		}
		campaignPerformanceReport = append(campaignPerformanceReport, &campaignPerformanceRecord)
	}
	return campaignPerformanceReport
}

// GetAccountReport transform raw report to AccountPerformanceReport
func GetAccountReport(report []*ReportRecord) []*AccountPerformanceReportColumns {
	accountPerformanceReport := []*AccountPerformanceReportColumns{}
	for _, record := range report {
		accountPerformanceRecord := AccountPerformanceReportColumns{}
		for k, v := range record.Record {
			switch string(k) {
			case "AccountId":
				accountPerformanceRecord.AccountId = v
			case "AccountName":
				accountPerformanceRecord.AccountName = v
			case "AccountNumber":
				accountPerformanceRecord.AccountNumber = v
			case "DeviceType":
				accountPerformanceRecord.DeviceType = v
			case "Impressions":
				intV, _ := strconv.ParseInt(v, 0, 64)
				accountPerformanceRecord.Impressions = intV
			case "Clicks":
				intV, _ := strconv.ParseInt(v, 0, 64)
				accountPerformanceRecord.Clicks = intV
			case "Ctr":
				floatV, _ := strconv.ParseFloat(v, 64)
				accountPerformanceRecord.Ctr = floatV
			case "AverageCpc":
				accountPerformanceRecord.AverageCpc = v
			case "Spend":
				floatV, _ := strconv.ParseFloat(v, 64)
				accountPerformanceRecord.Spend = floatV
			case "Conversions":
				intV, _ := strconv.ParseInt(v, 0, 64)
				accountPerformanceRecord.Conversions = intV
			case "Revenue":
				floatV, _ := strconv.ParseFloat(v, 64)
				accountPerformanceRecord.Revenue = floatV
			case "TimePeriod":
				accountPerformanceRecord.TimePeriod = v
			default:
				if strings.Contains(k, "TimePeriod") {
					accountPerformanceRecord.TimePeriod = v
				}
			}
		}
		accountPerformanceReport = append(accountPerformanceReport, &accountPerformanceRecord)
	}
	return accountPerformanceReport
}
