package v1

import (
	"github.com/homelifehq/aircall.go/aircall"
	"github.com/homelifehq/aircall.go/api/number"
	"github.com/homelifehq/aircall.go/api/user/v1"
)

type API struct {
	version string
	client  *aircall.Client
}

type Resource struct {
	ID         int             `json:"id"`
	DirectLink string          `json:"direct_link"`
	Direction  string          `json:"direction"`
	Status     string          `json:"status"`
	StartedAt  int             `json:"started_at"`
	Duration   int             `json:"duration"`
	Cost       string          `json:"cost"`
	RawDigits  string          `json:"raw_digits"`
	User       user.Resource   `json:"user"`
	Number     number.Resource `json:"number"`
	Archived   bool            `json:"archived"`
	Teams      []interface{}   `json:"teams"`
	Comments   []interface{}   `json:"comments"`
	Tags       []interface{}   `json:"tags"`
}

func NewAPI(client *aircall.Client) API {
	return API{version: "v1", client: client}
}
