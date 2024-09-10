package engagespot

import "encoding/json"

// Notification represents the structure of a notification to be sent
type Notification struct {
	Workflow Workflow               `json:"workflow"`
	Data     map[string]interface{} `json:"data"`
	SendTo   SendTo                 `json:"sendTo"`
	Options  map[string]interface{} `json:"options,omitempty"`
}

// Workflow represents the workflow information for a notification
type Workflow struct {
	Identifier string `json:"identifier"`
}

// SendTo represents the recipients of a notification
type SendTo struct {
	Recipients []string `json:"recipients"`
}

// MarshalJSON implements custom JSON marshaling for Notification
func (n Notification) MarshalJSON() ([]byte, error) {
	type Alias Notification
	aux := struct {
		Alias
		*Options
	}{
		Alias: Alias(n),
	}

	if n.Options != nil {
		aux.Options = (*Options)(&n.Options)
	}

	return json.Marshal(aux)
}

// Options is a helper type for custom JSON marshaling
type Options map[string]interface{}

// SendResponse represents the response from the Send API call
type SendResponse struct {
	RequestID string `json:"requestId"`
}

// User represents an Engagespot user
type User struct {
	// Add user fields here
}

// UserResponse represents the response from user-related API calls
type UserResponse struct {
	// Add response fields here
}