package bing

import "encoding/xml"

// Request structs

// RequestEnvelope contains all the structure of the enveloped request in xml format.
type RequestEnvelope struct {
	XMLName xml.Name `xml:"s:Envelope"`
	EnvNS   string   `xml:"xmlns:i,attr"`
	EnvSS   string   `xml:"xmlns:s,attr"`
	Header  RequestHeader
	Body    RequestBody
}

// RequestHeader contains the header fields of the xml format.
type RequestHeader struct {
	XMLName             xml.Name `xml:"s:Header"`
	BingNS              string   `xml:"xmlns,attr"`
	Action              string
	AuthenticationToken string `xml:"AuthenticationToken,omitempty"`
	CustomerAccountId   string `xml:"CustomerAccountId"`
	CustomerId          string `xml:"CustomerId,omitempty"`
	DeveloperToken      string `xml:"DeveloperToken"`
	Password            string `xml:"Password,omitempty"`
	Username            string `xml:"UserName,omitempty"`
}

// RequestBody allows us to specify the body of a request in xml format.
type RequestBody struct {
	XMLName xml.Name `xml:"s:Body"`
	Body    interface{}
}

// Response structs

// SoapResponseEnvelope is the main structure of a Soap response in xml format.
type SoapResponseEnvelope struct {
	XMLName xml.Name         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Header  TrackingId       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header"`
	Body    SoapResponseBody `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
}

// SoapResponseBody
type SoapResponseBody struct {
	OperationResponse []byte `xml:",innerxml"`
}

// TrackingID contains the identifier of the log entry that contains the details of the API call.
type TrackingId struct {
	Nil        bool   `xml:"http://www.w3.org/2001/XMLSchema-instance nil,attr"`
	TrackingId string `xml:"https://adcenter.microsoft.com/v8 TrackingId"`
}
