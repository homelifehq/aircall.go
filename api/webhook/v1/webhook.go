package v1

type Payload[C any] struct {
	Token     string   `json:"token"`
	Resource  Resource `json:"resource"`
	Event     Event    `json:"event"`
	Timestamp int      `json:"timestamp"`
	Data      C        `"json:data"` //nolint:govet
}
