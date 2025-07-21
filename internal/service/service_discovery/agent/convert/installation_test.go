package convert

import (
	"testing"

	"github.com/harness/harness-go-sdk/harness/svcdiscovery"
	"github.com/stretchr/testify/assert"
)

func TestFlattenInstallationDetails(t *testing.T) {
	tests := []struct {
		name    string
		details *svcdiscovery.DatabaseInstallationCollection
		want    map[string]interface{}
		wantErr bool
	}{
		{
			name:    "Nil input",
			details: nil,
			want:    nil,
			wantErr: false,
		},
		{
			name:    "Empty details",
			details: &svcdiscovery.DatabaseInstallationCollection{},
			want: map[string]interface{}{
				"account_identifier":      "",
				"created_at":              "",
				"created_by":              "",
				"delegate_id":             "",
				"delegate_task_id":        "",
				"delegate_task_status":    "",
				"environment_identifier":  "",
				"id":                      "",
				"is_cron_triggered":       false,
				"log_stream_created_at":   "",
				"log_stream_id":           "",
				"organization_identifier": "",
				"project_identifier":      "",
				"stopped":                 false,
				"updated_at":              "",
				"updated_by":              "",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FlattenInstallationDetails(tt.details)
			if (err != nil) != tt.wantErr {
				t.Errorf("FlattenInstallationDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
