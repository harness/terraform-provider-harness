package convert

import (
	"fmt"

	"github.com/harness/harness-go-sdk/harness/svcdiscovery"
)

// ExpandDataCollectionConfig expands the data collection configuration from Terraform schema
func ExpandDataCollectionConfig(input []interface{}) (*svcdiscovery.DatabaseDataCollectionConfiguration, error) {
	if len(input) == 0 || input[0] == nil {
		return nil, nil
	}

	cfg := input[0].(map[string]interface{})
	data := &svcdiscovery.DatabaseDataCollectionConfiguration{}

	var err error

	data.EnableNodeAgent, err = getBool(cfg, "enable_node_agent", false)
	if err != nil {
		return nil, fmt.Errorf("data.enable_node_agent: %w", err)
	}

	data.NodeAgentSelector, err = getString(cfg, "node_agent_selector", false)
	if err != nil {
		return nil, fmt.Errorf("data.node_agent_selector: %w", err)
	}

	data.NamespaceSelector, err = getString(cfg, "namespace_selector", false)
	if err != nil {
		return nil, fmt.Errorf("data.namespace_selector: %w", err)
	}

	data.EnableOrphanedPod, err = getBool(cfg, "enable_orphaned_pod", false)
	if err != nil {
		return nil, fmt.Errorf("data.enable_orphaned_pod: %w", err)
	}

	data.EnableBatchResources, err = getBool(cfg, "enable_batch_resources", false)
	if err != nil {
		return nil, fmt.Errorf("data.enable_batch_resources: %w", err)
	}

	collectionWindow, err := getInt(cfg, "collection_window_in_min", 0)
	if err != nil {
		return nil, fmt.Errorf("data.collection_window_in_min: %w", err)
	}
	data.CollectionWindowInMin = int32(collectionWindow)

	// Handle string slices
	if blacklisted, ok := cfg["blacklisted_namespaces"].([]interface{}); ok {
		data.BlacklistedNamespaces = make([]string, 0, len(blacklisted))
		for _, ns := range blacklisted {
			if str, ok := ns.(string); ok {
				data.BlacklistedNamespaces = append(data.BlacklistedNamespaces, str)
			}
		}
	}

	if observed, ok := cfg["observed_namespaces"].([]interface{}); ok {
		data.ObservedNamespaces = make([]string, 0, len(observed))
		for _, ns := range observed {
			if str, ok := ns.(string); ok {
				data.ObservedNamespaces = append(data.ObservedNamespaces, str)
			}
		}
	}

	// Handle cron
	if cron, ok := cfg["cron"].([]interface{}); ok && len(cron) > 0 {
		if cronMap, ok := cron[0].(map[string]interface{}); ok {
			if expr, ok := cronMap["expression"].(string); ok {
				data.Cron = &svcdiscovery.DatabaseCronConfig{
					Expression: expr,
				}
			}
		}
	}

	return data, nil
}

// FlattenDataCollectionConfig flattens the data collection configuration to Terraform schema
func FlattenDataCollectionConfig(data *svcdiscovery.DatabaseDataCollectionConfiguration) (map[string]interface{}, error) {
	if data == nil {
		return nil, nil
	}

	result := map[string]interface{}{
		"enable_node_agent":        data.EnableNodeAgent,
		"node_agent_selector":      data.NodeAgentSelector,
		"namespace_selector":       data.NamespaceSelector,
		"enable_orphaned_pod":      data.EnableOrphanedPod,
		"enable_batch_resources":   data.EnableBatchResources,
		"collection_window_in_min": data.CollectionWindowInMin,
	}

	// Add slices
	if len(data.BlacklistedNamespaces) > 0 {
		result["blacklisted_namespaces"] = data.BlacklistedNamespaces
	}

	if len(data.ObservedNamespaces) > 0 {
		result["observed_namespaces"] = data.ObservedNamespaces
	}

	// Add cron
	if data.Cron != nil && data.Cron.Expression != "" {
		result["cron"] = []interface{}{
			map[string]interface{}{
				"expression": data.Cron.Expression,
			},
		}
	}

	return result, nil
}
