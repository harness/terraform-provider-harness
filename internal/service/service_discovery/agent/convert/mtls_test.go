package convert

import (
	"testing"

	"github.com/harness/harness-go-sdk/harness/svcdiscovery"
	"github.com/stretchr/testify/assert"
)

func TestExpandMtlsConfig(t *testing.T) {
	tests := []struct {
		name    string
		input   []interface{}
		want    *svcdiscovery.DatabaseMtlsConfiguration
		wantErr bool
	}{
		{
			name:  "Empty config",
			input: []interface{}{},
			want:  nil,
		},
		{
			name: "Full config",
			input: []interface{}{
				map[string]interface{}{
					"cert_path":   "/path/to/cert",
					"key_path":    "/path/to/key",
					"secret_name": "mtls-secret",
					"url":         "https://example.com",
				},
			},
			want: &svcdiscovery.DatabaseMtlsConfiguration{
				CertPath:   "/path/to/cert",
				KeyPath:    "/path/to/key",
				SecretName: "mtls-secret",
				Url:        "https://example.com",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExpandMtlsConfig(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExpandMtlsConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
