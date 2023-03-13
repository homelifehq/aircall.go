package user

import (
	"encoding/json"
	"fmt"
	"github.com/homelifehq/aircall.go/internal/transport"
	"io"
	"net/http"
)

type RetrieveResponse struct {
	User Resource `json:"user"`
}

func (api API) Retrieve(userID int) (response RetrieveResponse, err error) {
	version := api.version
	client := api.client
	isVerbose := client.Verbose()
	logger := client.Logger()
	auth := client.AuthStrategy()

	url := transport.BuildRequestURL(client.APIURL(), version, fmt.Sprintf("users/%d", userID))
	fmt.Println(url)
	// Request (GET https://api.aircall.io/v1/user/:id)
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
