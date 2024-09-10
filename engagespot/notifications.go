package engagespot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// NotificationService handles communication with the notification related
// methods of the Engagespot API.
type NotificationService struct {
	client *Client
}

// Send sends a notification using the Engagespot API
func (s *NotificationService) Send(notification *Notification) (*SendResponse, error) {
	url := fmt.Sprintf("%s/notifications", s.client.baseURL)

	reqBody, err := json.Marshal(map[string]interface{}{
		"notification": notification,
	})
	if err != nil {
		return nil, fmt.Errorf("error marshaling request body: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	resp, err := s.client.doRequest(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status code %d: %s", resp.StatusCode, string(body))
	}

	var sendResponse SendResponse
	err = json.Unmarshal(body, &sendResponse)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %v", err)
	}

	return &sendResponse, nil
}

// Add other notification-related methods here