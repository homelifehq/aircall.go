package v1

import (
	"encoding/json"
	webhook "github.com/homelifehq/aircall.go/api/webhook/v1"
)

// UnMarshal decodes user related webhook call and
// converts it to a Call.Resource
func UnMarshal(data []byte) (webhook.Payload[Resource], error) {
	var payload webhook.Payload[Resource]

	if err := json.Unmarshal(data, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}
