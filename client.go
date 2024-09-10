package dify

import (
	"fmt"
	"net/http"
)

// ClientConfig - Configuration for the Dify client.
type ClientConfig struct {
	BaseURL string // The base URL of the Dify API.
	APIKey  string // The API key for authentication.
}

// Client - Dify client for interacting with the API.
type Client struct {
	config ClientConfig // Configuration for the client.
	client *http.Client // HTTP client to send requests.
}

// NewClient - Creates and returns a new Dify client.
func NewClient(config ClientConfig) (*Client, error) {
	if config.BaseURL == "" || config.APIKey == "" {
		return nil, fmt.Errorf("BaseURL and APIKey must be provided")
	}
	return &Client{
		config: config,
		client: &http.Client{},
	}, nil
}
