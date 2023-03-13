package aircall

import (
	"github.com/charmbracelet/log"
	"github.com/homelifehq/aircall.go/internal"
)

type Client struct {
	apiURL       string
	authStrategy internal.AuthStrategy
	verbose      bool
	logger       log.Logger
}

type ClientOptions struct {
	AuthStrategy internal.AuthStrategy
	Verbose      bool
}

// NewClient ...
func NewClient(options ClientOptions) *Client {
	return &Client{
		apiURL:       "https://api.aircall.io",
		authStrategy: options.AuthStrategy,
		verbose:      options.Verbose,
		logger:       log.New(),
	}
}

// APIURL ...
func (client Client) APIURL() string {
	return client.apiURL
}

// AuthStrategy ...
func (client Client) AuthStrategy() internal.AuthStrategy {
	return client.authStrategy
}

// Verbose ...
func (client Client) Verbose() bool {
	return client.verbose
}

// Logger ...
func (client Client) Logger() log.Logger {
	return client.logger
}
