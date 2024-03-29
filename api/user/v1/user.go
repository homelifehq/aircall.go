package user

import (
	"github.com/homelifehq/aircall.go/aircall"
	"github.com/homelifehq/aircall.go/api/number"
)

type API struct {
	version string
	client  *aircall.Client
}

type Resource struct {
	ID                 int               `json:"id"`
	DirectLink         string            `json:"direct_link"`
	Name               string            `json:"name"`
	Email              string            `json:"email"`
	Available          bool              `json:"available"`
	AvailabilityStatus string            `json:"availability_status"`
	CreatedAt          string            `json:"created_at"`
	Timezone           string            `json:"time_zone"`
	Language           string            `json:"language"`
	WrapUpTime         int               `json:"wrap_up_time"`
	Numbers            []number.Resource `json:"numbers"`
}

func NewAPI(client *aircall.Client) API {
	return API{version: "v1", client: client}
}
