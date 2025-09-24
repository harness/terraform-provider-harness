package api

// Option represents a configuration option function.
type Option func(*Config) error

// APIKey sets the API key.
func APIKey(key string) Option {
	return func(c *Config) error {
		c.APIKey = key
		return nil
	}
}

// APIBaseURL sets the base URL.
func APIBaseURL(url string) Option {
	return func(c *Config) error {
		c.APIBaseURL = url
		return nil
	}
}

// UserAgent sets the user agent.
func UserAgent(userAgent string) Option {
	return func(c *Config) error {
		c.UserAgent = userAgent
		return nil
	}
}

// CustomHTTPHeaders sets custom HTTP headers.
func CustomHTTPHeaders(headers map[string]string) Option {
	return func(c *Config) error {
		c.CustomHTTPHeaders = headers
		return nil
	}
}

// ClientTimeout sets the client timeout.
func ClientTimeout(timeout int) Option {
	return func(c *Config) error {
		c.ClientTimeout = timeout
		return nil
	}
}