package cd

import "fmt"

type InvalidConfigError struct {
	Config *Config
	Field  string
	Err    error
}

func (c *Config) NewInvalidConfigError(field string, err error) InvalidConfigError {
	return InvalidConfigError{
		Config: c,
		Field:  field,
		Err:    err,
	}
}

func (e InvalidConfigError) Error() string {
	return fmt.Sprintf("invalid config: %s must be set %s", e.Field, e.Err)
}

func (e InvalidConfigError) Unwrap() error {
	return e.Err
}
