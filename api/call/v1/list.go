package v1

import (
	"encoding/json"
	"github.com/homelifehq/aircall.go/api/filter/v1"
	"github.com/homelifehq/aircall.go/api/meta"
	"github.com/homelifehq/aircall.go/internal/transport"
	"io"
	"net/http"
)

type ListCallsParams struct {
	Filter *filter.Params
}

type ListCallsResponse struct {
	Calls []Resource `json:"calls"`
	Meta  meta.Meta  `json:"meta"`
}

// ListCalls Fetch all Calls associated to a company and their information.
func (api API) ListCalls(params ListCallsParams) (response ListCallsResponse, err error) {
	version := api.version
	client := api.client
	isVerbose := client.Verbose()
	logger := client.Logger()
	auth := client.AuthStrategy()

	url := transport.BuildRequestURL(client.APIURL(), version, "calls")
	if params.Filter != nil {
		if err = transport.SetQueryParams(&url, params.Filter); err != nil {
			return
		}
	}

	// Request (GET https://api.aircall.io/v1/calls)
	// Create client
	httpClient := &http.Client{}

	// Create request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	// Headers
	auth.SetAuth(req)
	// Fetch Request
	resp, err := httpClient.Do(req) //nolint:body close
	if err != nil {
		return
	}
	// Read Response Body
	body, _ := io.ReadAll(resp.Body)

	if err = json.Unmarshal(body, &response); err != nil {
		return
	}

	if isVerbose {
		// Display Results
		logger.Print("response Status : ", resp.Status)
		logger.Print("response Headers : ", resp.Header)
	}

	return
}
