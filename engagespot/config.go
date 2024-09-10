package engagespot

// Config holds the configuration for the Engagespot client
type Config struct {
	APIKey    string
	APISecret string
	BaseURL   string
}

// DefaultConfig returns the default configuration for the Engagespot client
func DefaultConfig() *Config {
	return &Config{
		BaseURL: "https://api.engagespot.co/v3", // Adjust this URL if needed
	}
}