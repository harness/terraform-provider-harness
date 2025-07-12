package convert

import (
	"github.com/harness/harness-go-sdk/harness/svcdiscovery"
)

func FlattenInstallationDetails(details *svcdiscovery.DatabaseInstallationCollection) (map[string]interface{}, error) {
	if details == nil {
		return nil, nil
	}

	result := map[string]interface{}{
		"id":                      details.Id,
		"account_identifier":      details.AccountIdentifier,
		"organization_identifier": details.OrganizationIdentifier,
		"project_identifier":      details.ProjectIdentifier,
		"environment_identifier":  details.EnvironmentIdentifier,
		"delegate_task_id":        details.DelegateTaskID,
		"delegate_id":             details.DelegateID,
		"delegate_task_status":    details.DelegateTaskStatus,
		"is_cron_triggered":       details.IsCronTriggered,
		"log_stream_id":           details.LogStreamID,
		"log_stream_created_at":   details.LogStreamCreatedAt,
		"stopped":                 details.Stopped,
		"created_at":              details.CreatedAt,
		"updated_at":              details.UpdatedAt,
		"created_by":              details.CreatedBy,
		"updated_by":              details.UpdatedBy,
	}

	// Handle agent details
	if details.AgentDetails != nil {
		agentDetails := map[string]interface{}{
			"status": details.AgentDetails.Status,
		}

		// Handle cluster details
		if details.AgentDetails.Cluster != nil {
			cluster := map[string]interface{}{
				"name":      details.AgentDetails.Cluster.Name,
				"namespace": details.AgentDetails.Cluster.Namespace,
				"uid":       details.AgentDetails.Cluster.Uid,
				"status":    details.AgentDetails.Cluster.Status,
			}
			agentDetails["cluster"] = []interface{}{cluster}
		}

		result["agent_details"] = []interface{}{agentDetails}
	}

	return result, nil
}

// FlattenAgentDetails converts agent details to a Terraform-friendly format
func FlattenAgentDetails(details *svcdiscovery.DatabaseAgentDetails) (map[string]interface{}, error) {
	if details == nil {
		return nil, nil
	}

	result := map[string]interface{}{
		"status": details.Status,
	}

	// Handle cluster
	if details.Cluster != nil {
		cluster := map[string]interface{}{
			"name":      details.Cluster.Name,
			"namespace": details.Cluster.Namespace,
			"uid":       details.Cluster.Uid,
			"status":    details.Cluster.Status,
		}
		result["cluster"] = []interface{}{cluster}
	}

	// Handle lifecycle manager
	if len(details.LifecycleManager) > 0 {
		lifecycleManager := make([]interface{}, len(details.LifecycleManager))
		for i, lm := range details.LifecycleManager {
			lifecycleManager[i] = map[string]interface{}{
				"name":      lm.Name,
				"namespace": lm.Namespace,
				"uid":       lm.Uid,
				"status":    lm.Status,
			}
		}
		result["lifecycle_manager"] = lifecycleManager
	}

	// Handle nodes
	if len(details.Node) > 0 {
		nodes := make([]interface{}, len(details.Node))
		for i, node := range details.Node {
			nodes[i] = map[string]interface{}{
				"name":      node.Name,
				"namespace": node.Namespace,
				"uid":       node.Uid,
				"status":    node.Status,
			}
		}
		result["node"] = nodes
	}

	return result, nil
}
