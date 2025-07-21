package convert

import (
	"testing"

	"github.com/harness/harness-go-sdk/harness/svcdiscovery"
	"github.com/stretchr/testify/assert"
)

func TestExpandKubernetesConfig(t *testing.T) {
	tests := []struct {
		name    string
		input   []interface{}
		want    *svcdiscovery.DatabaseKubernetesAgentConfiguration
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
					"namespace":       "test-ns",
					"service_account": "test-sa",
					"run_as_user":     1000,
					"run_as_group":    2000,
					"namespaced":      true,
					"node_selector":   map[string]interface{}{"key": "value"},
					"labels":          map[string]interface{}{"app": "test"},
					"annotations":     map[string]interface{}{"test.com/key": "value"},
				},
			},
			want: &svcdiscovery.DatabaseKubernetesAgentConfiguration{
				Namespace:      "test-ns",
				ServiceAccount: "test-sa",
				RunAsUser:      1000,
				RunAsGroup:     2000,
				Namespaced:     true,
				NodeSelector:   map[string]string{"key": "value"},
				Labels:         map[string]string{"app": "test"},
				Annotations:    map[string]string{"test.com/key": "value"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExpandKubernetesConfig(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExpandKubernetesConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestFlattenKubernetesConfig(t *testing.T) {
	tests := []struct {
		name    string
		input   *svcdiscovery.DatabaseKubernetesAgentConfiguration
		want    map[string]interface{}
		wantErr bool
	}{
		{
			name:  "Nil input",
			input: nil,
			want:  nil,
		},
		{
			name: "Full config",
			input: &svcdiscovery.DatabaseKubernetesAgentConfiguration{
				Namespace:      "test-ns",
				ServiceAccount: "test-sa",
				RunAsUser:      1000,
				RunAsGroup:     2000,
				Namespaced:     true,
				NodeSelector:   map[string]string{"key": "value"},
			},
			want: map[string]interface{}{
				"namespace":       "test-ns",
				"service_account": "test-sa",
				"run_as_user":     1000,
				"run_as_group":    2000,
				"namespaced":      true,
				"node_selector":   map[string]string{"key": "value"},
				// Add default fields that are always included
				"disable_namespace_creation": false,
				"image_pull_policy":          "",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FlattenKubernetesConfig(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("FlattenKubernetesConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
