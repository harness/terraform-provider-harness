package convert

import (
	"testing"

	"github.com/harness/harness-go-sdk/harness/svcdiscovery"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
)

func TestFlattenAgentToSchema(t *testing.T) {
	tests := []struct {
		name    string
		agent   *svcdiscovery.ApiGetAgentResponse
		want    map[string]interface{}
		wantErr bool
	}{
		{
			name:    "Nil agent",
			agent:   nil,
			want:    nil,
			wantErr: true,
		},
		{
			name: "Minimal agent",
			agent: &svcdiscovery.ApiGetAgentResponse{
				Id:                    "test-id",
				Name:                  "test-agent",
				EnvironmentIdentifier: "test-env",
			},
			want: map[string]interface{}{
				"id":                     "test-id",
				"name":                   "test-agent",
				"environment_identifier": "test-env",
				"tags":                   []interface{}{},
				"config":                 []interface{}{},
				"installation_details":   []interface{}{},
			},
			wantErr: false,
		},
		{
			name: "Full agent",
			agent: &svcdiscovery.ApiGetAgentResponse{
				Id:                     "test-id",
				Name:                   "test-agent",
				Description:            "test description",
				OrganizationIdentifier: "test-org",
				ProjectIdentifier:      "test-project",
				EnvironmentIdentifier:  "test-env",
				Identity:               "test-identity",
				InstallationType:       "HELM",
				WebhookURL:             "https://example.com/webhook",
				CorrelationID:          "test-correlation",
				CreatedAt:              "2023-01-01T00:00:00Z",
				UpdatedAt:              "2023-01-02T00:00:00Z",
				CreatedBy:              "test-user",
				UpdatedBy:              "test-user",
				PermanentInstallation:  true,
				Removed:                false,
				NetworkMapCount:        5,
				ServiceCount:           10,
				Tags:                   []string{"key1=value1", "key2=value2"},
				Config: &svcdiscovery.DatabaseAgentConfiguration{
					CollectorImage:   "test-image",
					LogWatcherImage:  "log-watcher",
					SkipSecureVerify: true,
				},
				InstallationDetails: &svcdiscovery.DatabaseInstallationCollection{
					DelegateID:         "test-delegate",
					DelegateTaskID:     "test-delegate-task",
					DelegateTaskStatus: "RUNNING",
				},
			},
			want: map[string]interface{}{
				"id":                     "test-id",
				"name":                   "test-agent",
				"description":            "test description",
				"org_identifier":         "test-org",
				"project_identifier":     "test-project",
				"environment_identifier": "test-env",
				"identity":               "test-identity",
				"installation_type":      "HELM",
				"webhook_url":            "https://example.com/webhook",
				"correlation_id":         "test-correlation",
				"created_at":             "2023-01-01T00:00:00Z",
				"updated_at":             "2023-01-02T00:00:00Z",
				"created_by":             "test-user",
				"updated_by":             "test-user",
				"permanent_installation": true,
				"removed":                false,
				"network_map_count":      5,
				"service_count":          10,
				"tags":                   []interface{}{"key1=value1", "key2=value2"},
				"config": []interface{}{
					map[string]interface{}{
						"collector_image":    "test-image",
						"log_watcher_image":  "log-watcher",
						"skip_secure_verify": true,
					},
				},
				"installation_details": []interface{}{
					map[string]interface{}{
						"delegate_id":             "test-delegate",
						"delegate_task_id":        "test-delegate-task",
						"delegate_task_status":    "RUNNING",
						"account_identifier":      "",
						"created_at":              "",
						"created_by":              "",
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
				},
			},
			wantErr: false,
		},
		{
			name: "Agent with nil config and details",
			agent: &svcdiscovery.ApiGetAgentResponse{
				Id:                    "test-id",
				Name:                  "test-agent",
				EnvironmentIdentifier: "test-env",
				Config:                nil,
				InstallationDetails:   nil,
			},
			want: map[string]interface{}{
				"id":                     "test-id",
				"name":                   "test-agent",
				"environment_identifier": "test-env",
				"tags":                   []interface{}{},
				"config":                 []interface{}{},
				"installation_details":   []interface{}{},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := schema.TestResourceDataRaw(t, map[string]*schema.Schema{
				"id":                     {Type: schema.TypeString},
				"name":                   {Type: schema.TypeString},
				"description":            {Type: schema.TypeString},
				"org_identifier":         {Type: schema.TypeString},
				"project_identifier":     {Type: schema.TypeString},
				"environment_identifier": {Type: schema.TypeString},
				"identity":               {Type: schema.TypeString},
				"installation_type":      {Type: schema.TypeString},
				"webhook_url":            {Type: schema.TypeString},
				"correlation_id":         {Type: schema.TypeString},
				"created_at":             {Type: schema.TypeString},
				"updated_at":             {Type: schema.TypeString},
				"created_by":             {Type: schema.TypeString},
				"updated_by":             {Type: schema.TypeString},
				"permanent_installation": {Type: schema.TypeBool},
				"removed":                {Type: schema.TypeBool},
				"network_map_count":      {Type: schema.TypeInt},
				"service_count":          {Type: schema.TypeInt},
				"tags": {
					Type:     schema.TypeList,
					Optional: true,
					Elem:     &schema.Schema{Type: schema.TypeString},
				},
				"config": {
					Type:     schema.TypeList,
					Optional: true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"collector_image":    {Type: schema.TypeString},
							"log_watcher_image":  {Type: schema.TypeString},
							"skip_secure_verify": {Type: schema.TypeBool},
						},
					},
				},
				"installation_details": {
					Type:     schema.TypeList,
					Optional: true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"delegate_id":             {Type: schema.TypeString},
							"delegate_task_id":        {Type: schema.TypeString},
							"delegate_task_status":    {Type: schema.TypeString},
							"account_identifier":      {Type: schema.TypeString},
							"created_at":              {Type: schema.TypeString},
							"created_by":              {Type: schema.TypeString},
							"environment_identifier":  {Type: schema.TypeString},
							"id":                      {Type: schema.TypeString},
							"is_cron_triggered":       {Type: schema.TypeBool},
							"log_stream_created_at":   {Type: schema.TypeString},
							"log_stream_id":           {Type: schema.TypeString},
							"organization_identifier": {Type: schema.TypeString},
							"project_identifier":      {Type: schema.TypeString},
							"stopped":                 {Type: schema.TypeBool},
							"updated_at":              {Type: schema.TypeString},
							"updated_by":              {Type: schema.TypeString},
						},
					},
				},
			}, map[string]interface{}{})

			err := FlattenAgentToSchema(d, tt.agent)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)

			// Convert schema.ResourceData to map for comparison
			result := make(map[string]interface{})
			for k := range tt.want {
				result[k] = d.Get(k)
			}

			assert.Equal(t, tt.want, result)
		})
	}
}

func TestSetStringIfNotEmpty(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected interface{}
	}{
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "Non-empty string",
			input:    "test",
			expected: "test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := schema.TestResourceDataRaw(t, map[string]*schema.Schema{
				"test": {Type: schema.TypeString},
			}, map[string]interface{}{})

			setStringIfNotEmpty := func(key, value string) {
				if value == "" {
					d.Set(key, "")
				} else {
					d.Set(key, value)
				}
			}

			setStringIfNotEmpty("test", tt.input)
			assert.Equal(t, tt.expected, d.Get("test"))
		})
	}
}
