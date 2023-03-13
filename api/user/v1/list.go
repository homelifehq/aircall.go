package user

import (
	"encoding/json"
	"github.com/homelifehq/aircall.go/api/filter/v1"
	"github.com/homelifehq/aircall.go/api/meta"
	"github.com/homelifehq/aircall.go/internal/transport"
	"io"
	"net/http"
)

type ListUserParams struct {
	Filter *filter.Params
}

type ListUserResponse struct {
	Users []Resource `json:"users"`
	Meta  meta.Meta  `json:"meta"`
}

func (api API) ListUsers(params ListUserParams) (response ListUserResponse, err error) {
	version := api.version
	client := api.client
	isVerbose := client.Verbose()
	logger := client.Logger()
	auth := client.AuthStrategy()

	url := transport.BuildRequestURL(client.APIURL(), version, "users")
	if params.Filter != nil {
		if err = transport.SetQueryParams(&url, params.Filter); err != nil {
			return
		}
	}

	// Request (GET https://api.aircall.io/v1/users)
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
