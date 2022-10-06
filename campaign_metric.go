package bing

type CampaignMetric struct {
	AccountID                        string
	AccountName                      string
	AccountNumber                    string
	CampaignID                       string
	CampaignName                     string
	Date                             string
	Clicks                           int64
	Impressions                      int64
	Ctr                              string
	AverageCpc                       string
	Spend                            float64
	Conversions                      float64
	TopImpressionRatePercent         string
	AbsoluteTopImpressionRatePercent string
	Revenue                          string
}

// GetComposedID generates an unique id of the record, in this case accountID:campaignID:date, e.x. 1234:1111:2022-01-01
func (m *CampaignMetric) GetComposedID() string {
	return m.AccountID + ":" + m.CampaignID + ":" + m.Date
}
