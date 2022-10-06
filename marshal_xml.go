package bing

import (
	"encoding/xml"
	"fmt"
)

// MarshallXML logic to encode PerformanceReportRequest.
func (s PerformanceReportRequest) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if s.Type == "" {
		return fmt.Errorf("missing report type")
	}
	return marshallPerformanceReportRequest(e, s, s.Type)
}

// MarshallXML logic to encode AdGroupPerformanceReportRequest.
func (s *AdGroupPerformanceReportRequest) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	req := PerformanceReportRequest(*s)
	return marshallPerformanceReportRequest(e, req, "AdGroupPerformanceReport")
}

// MarshallXML logic to encode CampaignPerformanceReportRequest.
func (s *CampaignPerformanceReportRequest) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	req := PerformanceReportRequest(*s)
	return marshallPerformanceReportRequest(e, req, "CampaignPerformanceReport")
}

// MarshallXML logic to encode AccountPerformanceReportRequest.
func (s *AccountPerformanceReportRequest) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	req := PerformanceReportRequest(*s)
	return marshallPerformanceReportRequest(e, req, "AccountPerformanceReport")
}

// MarshallXML logic to encode Time for the performance report.
func (s ReportTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "Time"
	e.EncodeToken(start)
	if s.PredefinedTime != "" {
		e.EncodeElement(s.PredefinedTime, st("PredefinedTime"))
	} else {
		e.EncodeElement(s.CustomDateRangeEnd, st("CustomDateRangeEnd"))
		e.EncodeElement(s.CustomDateRangeStart, st("CustomDateRangeStart"))
	}
	e.EncodeToken(start.End())
	return nil
}

// MarshallXML logic to encode Scope for the performance report.
func (s ReportScope) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "Scope"
	e.EncodeToken(start)
	if len(s.AccountIds) > 0 {
		acts := st("AccountIds", "xmlns:a1", "http://schemas.microsoft.com/2003/10/Serialization/Arrays")
		e.EncodeToken(acts)
		for i := 0; i < len(s.AccountIds); i++ {
			e.EncodeElement(s.AccountIds[i], st("a1:long"))
		}
		e.EncodeToken(acts.End())
	}

	if len(s.AdGroups) > 0 {
		acts := st("AdGroups")
		e.EncodeToken(acts)
		e.Encode(s.AdGroups)
		e.EncodeToken(acts.End())
	}

	if len(s.Campaigns) > 0 {
		acts := st("Campaigns")
		e.EncodeToken(acts)
		e.Encode(s.Campaigns)
		e.EncodeToken(acts.End())
	}
	e.EncodeToken(start.End())
	return nil
}

// marshallPerformanceReportRequest encodes the entire request in xml format for a performance report.
func marshallPerformanceReportRequest(e *xml.Encoder, s PerformanceReportRequest, t string) error {
	start := st("ReportRequest", "i:type", t+"Request")
	e.EncodeToken(start)

	excludes := []string{"ExcludeReportFooter", "ExcludeReportHeader"}
	for i := 0; i < len(excludes); i++ {
		e.EncodeElement(true, st(excludes[i]))
	}
	if s.Aggregation != "" {
		e.EncodeElement(s.Aggregation, st("Aggregation"))
	}
	if s.Columns == nil || len(s.Columns) == 0 {
		return fmt.Errorf("no columns selected")
	}
	cols := st("Columns", "i:nil", "false")
	e.EncodeToken(cols)
	for i := 0; i < len(s.Columns); i++ {
		e.EncodeElement(s.Columns[i], st(t+"Column"))
	}
	e.EncodeToken(cols.End())
	e.Encode(s.Scope)
	e.Encode(s.Time)
	e.EncodeToken(start.End())
	return nil
}

// st adds attributes to the xml start elements.
func st(name string, attrs ...string) xml.StartElement {
	ret := xml.StartElement{
		Name: xml.Name{Local: name},
	}

	for i := 0; i < len(attrs); i += 2 {
		ret.Attr = append(ret.Attr, xml.Attr{Name: xml.Name{Local: attrs[i]}, Value: attrs[i+1]})
	}

	return ret
}
