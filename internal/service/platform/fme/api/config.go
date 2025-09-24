package api

// Config represents all configuration options available to user to customize the API v2.
type Config struct {
	APIBaseURL         string
	UserAgent          string
	CustomHTTPHeaders  map[string]string
	ContentTypeHeader  string
	AcceptHeader       string
	APIKey             string
	ClientTimeout      int
}

// ParseOptions parses the supplied options functions.
func (c *Config) ParseOptions(opts ...Option) error {
	// Range over each options function and apply it to our API type to
	// configure it. Options functions are applied in order, with any
	// conflicting options overriding earlier calls.
	for _, option := range opts {
		err := option(c)
		if err != nil {
			return err
		}
	}

	return nil
}