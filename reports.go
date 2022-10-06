package bing

type ReportRecord struct {
	Record map[string]string
}

type AccountPerformanceReportColumns struct {
	AccountName   string
	AccountNumber string
	AccountId     string
	TimePeriod    string
	Impressions   int64
	Clicks        int64
	Ctr           float64
	AverageCpc    string
	Spend         float64
	Conversions   int64
	DeviceType    string
	Revenue       float64
	// CurrencyCode    string
	// AdDistribution  string
	// AveragePosition string
	// ConversionRate                                string
	// CostPerConversion                             string
	// LowQualityClicks                              string
	// LowQualityClicksPercent                       string
	// LowQualityImpressions                         string
	// LowQualityImpressionsPercent                  string
	// LowQualityConversions                         string
	// LowQualityConversionRate                      string
	// DeviceOS                                      string
	// ImpressionSharePercent                        string
	// ImpressionLostToBudgetPercent                 string
	// ImpressionLostToRankAggPercent                string
	// PhoneImpressions                              string
	// PhoneCalls                                    string
	// Ptr                                           string
	// Network                                       string
	// TopVsOther                                    string
	// BidMatchType                                  string
	// DeliveredMatchType                            string
	// Assists                                       string
	// ReturnOnAdSpend                               string
	// CostPerAssist                                 string
	// RevenuePerConversion                          string
	// RevenuePerAssist                              string
	// AccountStatus                                 string
	// LowQualityGeneralClicks                       string
	// LowQualitySophisticatedClicks                 string
	// ExactMatchImpressionSharePercent              string
	// CustomerId                                    string
	// CustomerName                                  string
	// ClickSharePercent                             string
	// AbsoluteTopImpressionSharePercent             string
	// TopImpressionShareLostToRankPercent           string
	// TopImpressionShareLostToBudgetPercent         string
	// AbsoluteTopImpressionShareLostToRankPercent   string
	// AbsoluteTopImpressionShareLostToBudgetPercent string
	// TopImpressionSharePercent                     string
	// AbsoluteTopImpressionRatePercent              string
	// TopImpressionRatePercent                      string
	// AllConversions                                string
	// AllRevenue                                    string
	// AllConversionRate                             string
	// AllCostPerConversion                          string
	// AllReturnOnAdSpend                            string
	// AllRevenuePerConversion                       string
	// ViewThroughConversions                        string
	// Goal                                          string
	// GoalType                                      string
	// AudienceImpressionSharePercent                string
	// AudienceImpressionLostToRankPercent           string
	// AudienceImpressionLostToBudgetPercent         string
	// AverageCpm                                    string
	// ConversionsQualified                          string
	// LowQualityConversionsQualified                string
	// AllConversionsQualified                       string
	// ViewThroughConversionsQualified               string
}

type CampaignPerformanceReportColumns struct {
	AccountName                      string
	AccountNumber                    string
	AccountId                        string
	TimePeriod                       string
	CampaignName                     string
	CampaignId                       string
	Impressions                      int64
	Clicks                           int64
	Ctr                              string
	AverageCpc                       string
	Spend                            float64
	Conversions                      int64
	DeviceType                       string
	Revenue                          string
	AbsoluteTopImpressionRatePercent string
	TopImpressionRatePercent         string
	// CampaignStatus                                string
	// CurrencyCode                                  string
	// AdDistribution                                string
	// AveragePosition                               string
	// ConversionRate                                string
	// CostPerConversion                             string
	// LowQualityClicks                              string
	// LowQualityClicksPercent                       string
	// LowQualityImpressions                         string
	// LowQualityImpressionsPercent                  string
	// LowQualityConversions                         string
	// LowQualityConversionRate                      string
	// DeviceOS                                      string
	// ImpressionSharePercent                        string
	// ImpressionLostToBudgetPercent                 string
	// ImpressionLostToRankAggPercent                string
	// QualityScore                                  string
	// ExpectedCtr                                   string
	// AdRelevance                                   string
	// LandingPageExperience                         string
	// HistoricalQualityScore                        string
	// HistoricalExpectedCtr                         string
	// HistoricalAdRelevance                         string
	// HistoricalLandingPageExperience               string
	// PhoneImpressions                              string
	// PhoneCalls                                    string
	// Ptr                                           string
	// Network                                       string
	// TopVsOther                                    string
	// BidMatchType                                  string
	// DeliveredMatchType                            string
	// Assists                                       string
	// ReturnOnAdSpend                               string
	// CostPerAssist                                 string
	// RevenuePerConversion                          string
	// RevenuePerAssist                              string
	// TrackingTemplate                              string
	// CustomParameters                              string
	// AccountStatus                                 string
	// BudgetName                                    string
	// BudgetStatus                                  string
	// BudgetAssociationStatus                       string
	// LowQualityGeneralClicks                       string
	// LowQualitySophisticatedClicks                 string
	// CampaignLabels                                string
	// ExactMatchImpressionSharePercent              string
	// CustomerId                                    string
	// CustomerName                                  string
	// ClickSharePercent                             string
	// AbsoluteTopImpressionSharePercent             string
	// FinalUrlSuffix                                string
	// CampaignType                                  string
	// TopImpressionShareLostToRankPercent           string
	// TopImpressionShareLostToBudgetPercent         string
	// AbsoluteTopImpressionShareLostToRankPercent   string
	// AbsoluteTopImpressionShareLostToBudgetPercent string
	// TopImpressionSharePercent                     string
	// TopImpressionRatePercent                      string
	// BaseCampaignId                                string
	// AllConversions                                string
	// AllRevenue                                    string
	// AllConversionRate                             string
	// AllCostPerConversion                          string
	// AllReturnOnAdSpend                            string
	// AllRevenuePerConversion                       string
	// ViewThroughConversions                        string
	// Goal                                          string
	// GoalType                                      string
	// AudienceImpressionSharePercent                string
	// AudienceImpressionLostToRankPercent           string
	// AudienceImpressionLostToBudgetPercent         string
	// RelativeCtr                                   string
	// AverageCpm                                    string
	// ConversionsQualified                          string
	// LowQualityConversionsQualified                string
	// AllConversionsQualified                       string
	// ViewThroughConversionsQualified               string
}

type AdGroupPerformanceReportColumns struct {
	AccountName   string
	AccountNumber string
	AccountId     string
	TimePeriod    string
	CampaignName  string
	CampaignId    string
	AdGroupName   string
	AdGroupId     string
	Impressions   int64
	Clicks        int64
	Ctr           float64
	AverageCpc    string
	Spend         float64
	Conversions   int64
	DeviceType    string
	Revenue       float64
	// Status                                        string
	// CurrencyCode                                  string
	// AdDistribution                                string
	// AveragePosition                               string
	// ConversionRate                                string
	// CostPerConversion                             string
	// Language                                      string
	// DeviceOS                                      string
	// ImpressionSharePercent                        string
	// ImpressionLostToBudgetPercent                 string
	// ImpressionLostToRankAggPercent                string
	// QualityScore                                  string
	// ExpectedCtr                                   string
	// AdRelevance                                   string
	// LandingPageExperience                         string
	// HistoricalQualityScore                        string
	// HistoricalExpectedCtr                         string
	// HistoricalAdRelevance                         string
	// HistoricalLandingPageExperience               string
	// PhoneImpressions                              string
	// PhoneCalls                                    string
	// Ptr                                           string
	// Network                                       string
	// TopVsOther                                    string
	// BidMatchType                                  string
	// DeliveredMatchType                            string
	// Assists                                       string
	// ReturnOnAdSpend                               string
	// CostPerAssist                                 string
	// RevenuePerConversion                          string
	// RevenuePerAssist                              string
	// TrackingTemplate                              string
	// CustomParameters                              string
	// AccountStatus                                 string
	// CampaignStatus                                string
	// AdGroupLabels                                 string
	// ExactMatchImpressionSharePercent              string
	// CustomerId                                    string
	// CustomerName                                  string
	// ClickSharePercent                             string
	// AbsoluteTopImpressionSharePercent             string
	// FinalUrlSuffix                                string
	// CampaignType                                  string
	// TopImpressionShareLostToRankPercent           string
	// TopImpressionShareLostToBudgetPercent         string
	// AbsoluteTopImpressionShareLostToRankPercent   string
	// AbsoluteTopImpressionShareLostToBudgetPercent string
	// TopImpressionSharePercent                     string
	// AbsoluteTopImpressionRatePercent              string
	// TopImpressionRatePercent                      string
	// BaseCampaignId                                string
	// AllConversions                                string
	// AllRevenue                                    string
	// AllConversionRate                             string
	// AllCostPerConversion                          string
	// AllReturnOnAdSpend                            string
	// AllRevenuePerConversion                       string
	// ViewThroughConversions                        string
	// Goal                                          string
	// GoalType                                      string
	// AudienceImpressionSharePercent                string
	// AudienceImpressionLostToRankPercent           string
	// AudienceImpressionLostToBudgetPercent         string
	// RelativeCtr                                   string
	// AdGroupType                                   string
	// AverageCpm                                    string
	// ConversionsQualified                          string
	// AllConversionsQualified                       string
	// ViewThroughConversionsQualified               string
}
