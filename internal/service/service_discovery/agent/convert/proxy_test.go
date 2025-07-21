package convert

import (
	"testing"

	"github.com/harness/harness-go-sdk/harness/svcdiscovery"
	"github.com/stretchr/testify/assert"
)

func TestExpandProxyConfig(t *testing.T) {
	tests := []struct {
		name    string
		input   []interface{}
		want    *svcdiscovery.DatabaseProxyConfiguration
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
					"http_proxy":  "http://proxy.example.com",
					"https_proxy": "https://proxy.example.com",
					"no_proxy":    "localhost,127.0.0.1",
					"url":         "http://proxy.example.com",
				},
			},
			want: &svcdiscovery.DatabaseProxyConfiguration{
				HttpProxy:  "http://proxy.example.com",
				HttpsProxy: "https://proxy.example.com",
				NoProxy:    "localhost,127.0.0.1",
				Url:        "http://proxy.example.com",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExpandProxyConfig(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExpandProxyConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
