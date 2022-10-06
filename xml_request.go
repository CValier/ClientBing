package bing

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
)

var reporting = "https://bingads.microsoft.com/Reporting/v13"

func (b *Session) reportRequest(body interface{}, endpoint string, soapAction string) ([]byte, error) {
	return b.sendRequest(body, endpoint, soapAction, reporting)
}

// sendRequest envelopes the xml payload and sent a POST request.
func (b *Session) sendRequest(body interface{}, endpoint string, soapAction string, ns string) ([]byte, error) {
	// Set the header
	header := RequestHeader{
		BingNS:            ns,         // Setting the xml namespace of the request xmlns:
		Action:            soapAction, // Specifying the action on the request
		CustomerAccountId: b.AccountId,
		CustomerId:        b.CustomerId,
		DeveloperToken:    b.DeveloperToken,
	}
	// Set the authentication token or the
	if b.TokenSource != nil {
		token, err := b.TokenSource.Token()
		if err != nil {
			return nil, errors.New("Error getting access token from current session. " + err.Error() + "Action: " + soapAction)
		}
		header.AuthenticationToken = token.AccessToken
	} else {
		header.Username = b.Username
		header.Password = b.Password
	}

	// Envelope the entire request
	envelope := RequestEnvelope{
		EnvNS:  "http://www.w3.org/2001/XMLSchema-instance", // xmlns:i
		EnvSS:  "http://schemas.xmlsoap.org/soap/envelope/", // xmlns:s
		Header: header,                                      // Set the generated header
		Body: RequestBody{ // Set the body
			Body: body,
		},
	}

	// Indent the body
	req, err := xml.MarshalIndent(envelope, "", "  ")
	if err != nil {
		return nil, errors.New("Error indenting request. " + err.Error() + "Action: " + soapAction)
	}

	// create the http request to the specified endpoint
	httpRequest, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(req))
	if err != nil {
		return nil, errors.New("Error generating the http request. " + err.Error() + "Action: " + soapAction)
	}

	// Adding Content-type and SOAP action to the request header
	httpRequest.Header.Add("Content-Type", "text/xml; charset=utf-8")
	httpRequest.Header.Add("SOAPAction", soapAction)

	// Make the http request
	response, err := b.HTTPClient.Do(httpRequest)
	if err != nil {
		return nil, errors.New("Error making the Http Post request. " + err.Error() + "Action: " + soapAction)
	}
	// Read the response body
	raw, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.New("Error reading the request response. " + err.Error() + "Action: " + soapAction)
	}

	// Formating the response to a xml format.
	res := SoapResponseEnvelope{}
	err = xml.Unmarshal(raw, &res)
	if err != nil {
		return nil, errors.New("Error unmarshaling request response. " + err.Error() + "Action: " + soapAction)
	}

	// Handling response errors.
	switch response.StatusCode {
	case 400, 401, 403, 405, 500:
		// Unmarshaling error details.
		fault := Fault{}
		err = xml.Unmarshal(res.Body.OperationResponse, &fault)
		if err != nil {
			return res.Body.OperationResponse, err
		}
		for _, e := range fault.Detail.Errors.AdApiErrors {
			switch e.ErrorCode {
			case "AuthenticationTokenExpired", "InvalidCredentials", "InternalError", "CallRateExceeded":
				return res.Body.OperationResponse, &baseError{
					code:    e.ErrorCode,
					origErr: &fault.Detail.Errors,
				}
			}
		}
		return res.Body.OperationResponse, &fault.Detail.Errors
	}

	return res.Body.OperationResponse, err
}
