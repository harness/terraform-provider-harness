package convert

import (
	"fmt"

	"github.com/harness/harness-go-sdk/harness/svcdiscovery"
)

// ExpandKubernetesConfig expands the Kubernetes configuration from Terraform schema
func ExpandKubernetesConfig(input []interface{}) (*svcdiscovery.DatabaseKubernetesAgentConfiguration, error) {
	if len(input) == 0 || input[0] == nil {
		return nil, nil
	}

	cfg := input[0].(map[string]interface{})
	k8s := &svcdiscovery.DatabaseKubernetesAgentConfiguration{}

	var err error

	// Required fields
	k8s.Namespace, err = getString(cfg, "namespace", true)
	if err != nil {
		return nil, fmt.Errorf("kubernetes.namespace: %w", err)
	}

	k8s.ServiceAccount, err = getString(cfg, "service_account", false)
	if err != nil {
		return nil, fmt.Errorf("kubernetes.service_account: %w", err)
	}

	// Optional fields
	k8s.DisableNamespaceCreation, err = getBool(cfg, "disable_namespace_creation", false)
	if err != nil {
		return nil, fmt.Errorf("kubernetes.disable_namespace_creation: %w", err)
	}

	k8s.Namespaced, err = getBool(cfg, "namespaced", false)
	if err != nil {
		return nil, fmt.Errorf("kubernetes.namespaced: %w", err)
	}

	k8s.ImagePullPolicy, err = getString(cfg, "image_pull_policy", false)
	if err != nil {
		return nil, fmt.Errorf("kubernetes.image_pull_policy: %w", err)
	}

	// Handle run_as_user
	if v, ok := cfg["run_as_user"]; ok {
		if v != nil {
			// Convert to int64 first, then to int32
			if runAsUser, ok := v.(int); ok {
				k8s.RunAsUser = int32(runAsUser)
			}
		}
	}

	// Handle run_as_group
	if v, ok := cfg["run_as_group"]; ok {
		if v != nil {
			// Convert to int64 first, then to int32
			if runAsGroup, ok := v.(int); ok {
				k8s.RunAsGroup = int32(runAsGroup)
			}
		}
	}

	// Handle resources
	if resources, ok := cfg["resources"].([]interface{}); ok && len(resources) > 0 {
		res, err := expandResources(resources[0].(map[string]interface{}))
		if err != nil {
			return nil, fmt.Errorf("kubernetes.resources: %w", err)
		}
		k8s.Resources = res
	}

	// Handle tolerations
	if tols, ok := cfg["tolerations"].([]interface{}); ok {
		tolerations := make([]svcdiscovery.V1Toleration, len(tols))
		for i, tol := range tols {
			t, err := expandToleration(tol.(map[string]interface{}))
			if err != nil {
				return nil, fmt.Errorf("kubernetes.tolerations[%d]: %w", i, err)
			}
			tolerations[i] = *t
		}
		k8s.Tolerations = tolerations
	}

	// Handle node selector
	if selector, ok := cfg["node_selector"].(map[string]interface{}); ok {
		k8s.NodeSelector = make(map[string]string)
		for k, v := range selector {
			if str, ok := v.(string); ok {
				k8s.NodeSelector[k] = str
			}
		}
	}

	// Handle labels
	if labels, ok := cfg["labels"].(map[string]interface{}); ok {
		k8s.Labels = make(map[string]string)
		for k, v := range labels {
			if str, ok := v.(string); ok {
				k8s.Labels[k] = str
			}
		}
	}

	// Handle annotations
	if annotations, ok := cfg["annotations"].(map[string]interface{}); ok {
		k8s.Annotations = make(map[string]string)
		for k, v := range annotations {
			if str, ok := v.(string); ok {
				k8s.Annotations[k] = str
			}
		}
	}

	return k8s, nil
}

// FlattenKubernetesConfig flattens the Kubernetes configuration to Terraform schema
func FlattenKubernetesConfig(k8s *svcdiscovery.DatabaseKubernetesAgentConfiguration) (map[string]interface{}, error) {
	if k8s == nil {
		return nil, nil
	}

	result := map[string]interface{}{
		"disable_namespace_creation": k8s.DisableNamespaceCreation,
		"namespaced":                 k8s.Namespaced,
		"namespace":                  k8s.Namespace,
		"service_account":            k8s.ServiceAccount,
		"image_pull_policy":          k8s.ImagePullPolicy,
		"run_as_user":                int(k8s.RunAsUser),
		"run_as_group":               int(k8s.RunAsGroup),
	}

	// Add resources
	if k8s.Resources != nil {
		resources, err := flattenResources(k8s.Resources)
		if err != nil {
			return nil, fmt.Errorf("flattening resources: %w", err)
		}
		result["resources"] = []interface{}{resources}
	}

	// Add tolerations
	if len(k8s.Tolerations) > 0 {
		tols := make([]interface{}, len(k8s.Tolerations))
		for i, t := range k8s.Tolerations {
			tols[i] = flattenToleration(t)
		}
		result["tolerations"] = tols
	}

	// Add maps
	if len(k8s.NodeSelector) > 0 {
		result["node_selector"] = k8s.NodeSelector
	}
	if len(k8s.Labels) > 0 {
		result["labels"] = k8s.Labels
	}
	if len(k8s.Annotations) > 0 {
		result["annotations"] = k8s.Annotations
	}

	return result, nil
}

func expandResources(m map[string]interface{}) (*svcdiscovery.DatabaseResourceRequirements, error) {
	res := &svcdiscovery.DatabaseResourceRequirements{}

	if limits, ok := m["limits"].([]interface{}); ok && len(limits) > 0 {
		if limitMap, ok := limits[0].(map[string]interface{}); ok {
			res.Limits = &svcdiscovery.DatabaseResourceList{
				Cpu:    limitMap["cpu"].(string),
				Memory: limitMap["memory"].(string),
			}
		}
	}

	if requests, ok := m["requests"].([]interface{}); ok && len(requests) > 0 {
		if reqMap, ok := requests[0].(map[string]interface{}); ok {
			res.Requests = &svcdiscovery.DatabaseResourceList{
				Cpu:    reqMap["cpu"].(string),
				Memory: reqMap["memory"].(string),
			}
		}
	}

	return res, nil
}

func flattenResources(res *svcdiscovery.DatabaseResourceRequirements) (map[string]interface{}, error) {
	result := make(map[string]interface{})

	if res.Limits != nil {
		result["limits"] = []interface{}{
			map[string]interface{}{
				"cpu":    res.Limits.Cpu,
				"memory": res.Limits.Memory,
			},
		}
	}

	if res.Requests != nil {
		result["requests"] = []interface{}{
			map[string]interface{}{
				"cpu":    res.Requests.Cpu,
				"memory": res.Requests.Memory,
			},
		}
	}

	return result, nil
}

func expandToleration(m map[string]interface{}) (*svcdiscovery.V1Toleration, error) {
	tol := &svcdiscovery.V1Toleration{}

	var err error

	// Required fields
	tol.Key, err = getString(m, "key", true)
	if err != nil {
		return nil, fmt.Errorf("key: %w", err)
	}

	tol.Operator, err = getString(m, "operator", true)
	if err != nil {
		return nil, fmt.Errorf("operator: %w", err)
	}

	// Optional fields
	if v, ok := m["value"]; ok {
		tol.Value = v.(string)
	}

	if v, ok := m["effect"]; ok {
		tol.Effect = v.(string)
	}

	if v, ok := m["toleration_seconds"]; ok {
		tol.TolerationSeconds = int32(v.(int))
	}

	return tol, nil
}

func flattenToleration(t svcdiscovery.V1Toleration) map[string]interface{} {
	result := map[string]interface{}{
		"key":                t.Key,
		"operator":           t.Operator,
		"value":              t.Value,
		"effect":             t.Effect,
		"toleration_seconds": int(t.TolerationSeconds),
	}

	return result
}
