package v1

import (
	"encoding/json"
	"fmt"
	"github.com/homelifehq/aircall.go/internal/transport"
	"io"
	"net/http"
)

type RetrieveResponse struct {
	Call Resource `json:"call"`
}

// Retrieve asynchronously retrieve a Call data like duration,
// direction, status, timestamps, comments or tags…
func (api API) Retrieve(callID int) (response RetrieveResponse, err error) {
	version := api.version
	client := api.client
	isVerbose := client.Verbose()
	logger := client.Logger()
	auth := client.AuthStrategy()

	url := transport.BuildRequestURL(client.APIURL(), version, fmt.Sprintf("calls/%d", callID))
	fmt.Println(url)
	// Request (GET https://api.aircall.io/v1/calls/:id)
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
