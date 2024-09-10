package engagespot

// Notification represents the structure of a notification to be sent
type Notification struct {
	Workflow Workflow               `json:"workflow"`
	Data     map[string]interface{} `json:"data"`
	SendTo   SendTo                 `json:"sendTo"`
}

// Workflow represents the workflow information for a notification
type Workflow struct {
	Identifier string `json:"identifier"`
}

// SendTo represents the recipients of a notification
type SendTo struct {
	Recipients []string `json:"recipients"`
}

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