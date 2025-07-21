package convert

import (
	"testing"

	"github.com/harness/harness-go-sdk/harness/svcdiscovery"
	"github.com/stretchr/testify/assert"
)

func TestExpandDataCollectionConfig(t *testing.T) {
	tests := []struct {
		name    string
		input   []interface{}
		want    *svcdiscovery.DatabaseDataCollectionConfiguration
		wantErr bool
	}{
		{
			name:    "Nil input",
			input:   nil,
			want:    nil,
			wantErr: false,
		},
		{
			name:    "Empty input",
			input:   []interface{}{},
			want:    nil,
			wantErr: false,
		},
		{
			name: "Full config",
			input: []interface{}{
				map[string]interface{}{
					"enable_node_agent":        true,
					"node_agent_selector":      "app=test",
					"enable_batch_resources":   true,
					"enable_orphaned_pod":      true,
					"namespace_selector":       "environment=dev",
					"collection_window_in_min": 15,
					"blacklisted_namespaces":   []interface{}{"kube-system", "kube-public"},
					"observed_namespaces":      []interface{}{"default", "harness"},
				},
			},
			want: &svcdiscovery.DatabaseDataCollectionConfiguration{
				EnableNodeAgent:       true,
				NodeAgentSelector:     "app=test",
				EnableBatchResources:  true,
				EnableOrphanedPod:     true,
				NamespaceSelector:     "environment=dev",
				CollectionWindowInMin: 15,
				BlacklistedNamespaces: []string{"kube-system", "kube-public"},
				ObservedNamespaces:    []string{"default", "harness"},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExpandDataCollectionConfig(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExpandDataCollectionConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestFlattenDataCollectionConfig(t *testing.T) {
	tests := []struct {
		name    string
		input   *svcdiscovery.DatabaseDataCollectionConfiguration
		want    map[string]interface{}
		wantErr bool
	}{
		{
			name:    "Nil input",
			input:   nil,
			want:    nil,
			wantErr: false,
		},
		{
			name: "Full config",
			input: &svcdiscovery.DatabaseDataCollectionConfiguration{
				EnableNodeAgent:       true,
				NodeAgentSelector:     "app=test",
				EnableBatchResources:  true,
				EnableOrphanedPod:     true,
				NamespaceSelector:     "environment=dev",
				CollectionWindowInMin: int32(15),
				BlacklistedNamespaces: []string{"kube-system", "kube-public"},
				ObservedNamespaces:    []string{"default", "harness"},
			},
			want: map[string]interface{}{
				"enable_node_agent":        true,
				"node_agent_selector":      "app=test",
				"enable_batch_resources":   true,
				"enable_orphaned_pod":      true,
				"namespace_selector":       "environment=dev",
				"collection_window_in_min": int32(15),
				"blacklisted_namespaces":   []string{"kube-system", "kube-public"},
				"observed_namespaces":      []string{"default", "harness"},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FlattenDataCollectionConfig(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("FlattenDataCollectionConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
