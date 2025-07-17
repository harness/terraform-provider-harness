package convert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrInvalidType_Error(t *testing.T) {
	tests := []struct {
		name     string
		err      *ErrInvalidType
		expected string
	}{
		{
			name: "With message",
			err: &ErrInvalidType{
				Expected: "string",
				Actual:   123,
				Msg:      "test message",
			},
			// The Msg field is not included in the error message
			expected: "invalid type: expected string, got int",
		},
		{
			name: "Without message",
			err: &ErrInvalidType{
				Expected: "string",
				Actual:   true,
			},
			expected: "invalid type: expected string, got bool",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.err.Error())
		})
	}
}

func TestErrRequiredField_Error(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		expected string
	}{
		{
			name:     "Simple field",
			field:    "test_field",
			expected: "required field missing: test_field",
		},
		{
			name:     "Empty field",
			field:    "",
			expected: "required field missing: ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := &ErrRequiredField{Field: tt.field}
			assert.Equal(t, tt.expected, err.Error())
		})
	}
}

func TestErrorWrapping(t *testing.T) {
	t.Run("ErrInvalidType contains expected substring", func(t *testing.T) {
		err := &ErrInvalidType{Expected: "string", Actual: 123}
		assert.Contains(t, err.Error(), "expected string, got int")
	})

	t.Run("ErrRequiredField contains expected substring", func(t *testing.T) {
		err := &ErrRequiredField{Field: "test"}
		assert.Contains(t, err.Error(), "required field missing: test")
	})
}
