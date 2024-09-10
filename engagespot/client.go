package engagespot

import (
	"net/http"
	"time"
)

// Client represents the Engagespot API client
type Client struct {
	apiKey     string
	apiSecret  string
	baseURL    string
	httpClient *http.Client

	Notifications *NotificationService
	Users         *UserService
}

// NewClient creates a new Engagespot API client
func NewClient(apiKey, apiSecret string) *Client {
	c := &Client{
		apiKey:    apiKey,
		apiSecret: apiSecret,
		baseURL:   "https://api.engagespot.co/v3",
		httpClient: &http.Client{
			Timeout: time.Second * 30,
		},
	}

	c.Notifications = &NotificationService{client: c}
	c.Users = &UserService{client: c}

	return c
}

// doRequest performs an HTTP request and returns the response
func (c *Client) doRequest(req *http.Request) (*http.Response, error) {
	req.Header.Set("X-ENGAGESPOT-API-KEY", c.apiKey)
	req.Header.Set("X-ENGAGESPOT-API-SECRET", c.apiSecret)
	req.Header.Set("Content-Type", "application/json")

	return c.httpClient.Do(req)
}