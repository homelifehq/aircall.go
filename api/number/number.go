package number

import "github.com/homelifehq/aircall.go/api/message"

type Resource struct {
	ID                       int    `json:"id"`
	DirectLink               string `json:"direct_link"`
	Name                     string `json:"name"`
	Digits                   string `json:"digits"`
	Country                  string `json:"country"`
	TimeZone                 string `json:"time_zone"`
	IsOpen                   bool   `json:"open"`
	AvailabilityStatus       string `json:"availability_status"`
	IsIVR                    bool   `json:"is_ivr"`
	IsLiveRecordingActivated bool   `json:"live_recording_activated"`
	messages                 []message.Resource
}
