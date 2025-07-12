package convert

import "fmt"

// ErrInvalidType is returned when a type assertion fails
type ErrInvalidType struct {
	Expected string
	Actual   interface{}
	Msg      string
}

func (e *ErrInvalidType) Error() string {
	return fmt.Sprintf("invalid type: expected %s, got %T", e.Expected, e.Actual)
}

// ErrRequiredField is returned when a required field is missing
type ErrRequiredField struct {
	Field string
}

func (e *ErrRequiredField) Error() string {
	return fmt.Sprintf("required field missing: %s", e.Field)
}

// ErrInvalidValue is returned when a field has an invalid value
type ErrInvalidValue struct {
	Field string
	Value interface{}
	Msg   string
}

func (e *ErrInvalidValue) Error() string {
	return fmt.Sprintf("invalid value for field %s: %v - %s", e.Field, e.Value, e.Msg)
}
