package security_governance_v3

import (
	"fmt"
	"strings"

	"github.com/harness/harness-go-sdk/harness/chaos"
)

// parseScopedImportIDV3 parses a project-scoped import ID of the form
// "<org_id>/<project_id>/<resource_id>". idLabel names the trailing segment
// (e.g. "rule-id" or "condition-id") for use in the error message. All three
// segments are required and must be non-empty.
func parseScopedImportIDV3(importID, idLabel string) (orgID, projectID, resourceID string, err error) {
	parts := strings.Split(importID, "/")
	if len(parts) != 3 {
		return "", "", "", fmt.Errorf("invalid import ID format. Expected \"<org-id>/<project-id>/<%s>\", got: %s", idLabel, importID)
	}
	orgID, projectID, resourceID = parts[0], parts[1], parts[2]
	if orgID == "" || projectID == "" || resourceID == "" {
		return "", "", "", fmt.Errorf("org_id, project_id, and %s cannot be empty", strings.ReplaceAll(idLabel, "-", "_"))
	}
	return orgID, projectID, resourceID, nil
}

// operatorPtr returns a pointer to a SecurityGovernanceOperator from a string.
func operatorPtr(s string) *chaos.SecurityGovernanceOperator {
	op := chaos.SecurityGovernanceOperator(s)
	return &op
}

// infraTypePtr returns a pointer to a SecurityGovernanceInfraType from a string.
func infraTypePtr(s string) *chaos.SecurityGovernanceInfraType {
	it := chaos.SecurityGovernanceInfraType(s)
	return &it
}

// operatorString safely dereferences a SecurityGovernanceOperator.
func operatorString(op *chaos.SecurityGovernanceOperator) string {
	if op == nil {
		return ""
	}
	return string(*op)
}

// faultTypeToREST maps the Terraform fault_type value to the value the
// chaos-guard REST handler actually accepts and stores.
//
// The OpenAPI contract documents "FAULT_NAME", but the chaos-guard REST
// handler, the Harness UI, the GraphQL enum, and all existing conditions in
// MongoDB use the legacy value "FAULT". Sending "FAULT_NAME" makes
// Terraform-managed conditions inconsistent with UI-managed ones and can break
// ChaosGuard fault matching. We therefore send "FAULT"/"FAULT_GROUP" to stay
// consistent with the platform.
func faultTypeToREST(s string) *chaos.SecurityGovernanceFaultType {
	var ft chaos.SecurityGovernanceFaultType
	switch s {
	case "FAULT_GROUP":
		ft = chaos.GROUP_SecurityGovernanceFaultType
	default: // "FAULT" (also accept legacy "FAULT_NAME" input)
		ft = chaos.SecurityGovernanceFaultType("FAULT")
	}
	return &ft
}

// faultTypeFromREST maps the REST SDK enum back to the Terraform value so that
// state matches the user's configuration (which uses "FAULT").
func faultTypeFromREST(ft *chaos.SecurityGovernanceFaultType) string {
	if ft != nil && *ft == chaos.GROUP_SecurityGovernanceFaultType {
		return "FAULT_GROUP"
	}
	return "FAULT"
}

// interfaceSliceToStringSlice converts a Terraform []interface{} to []string.
func interfaceSliceToStringSlice(slice []interface{}) []string {
	result := make([]string, 0, len(slice))
	for _, v := range slice {
		if s, ok := v.(string); ok {
			result = append(result, s)
		}
	}
	return result
}

// expandStringMap converts a Terraform map attribute to map[string]string.
func expandStringMap(m map[string]interface{}) map[string]string {
	if len(m) == 0 {
		return nil
	}
	out := make(map[string]string, len(m))
	for k, v := range m {
		if s, ok := v.(string); ok {
			out[k] = s
		}
	}
	return out
}

// -----------------------------------------------------------------------------
// expand: Terraform schema -> REST request models
// -----------------------------------------------------------------------------

func expandFaultsV3(faults []interface{}) []chaos.SecurityGovernanceFault {
	result := make([]chaos.SecurityGovernanceFault, 0, len(faults))
	for _, f := range faults {
		fault, ok := f.(map[string]interface{})
		if !ok {
			continue
		}
		result = append(result, chaos.SecurityGovernanceFault{
			FaultType: faultTypeToREST(fault["fault_type"].(string)),
			Name:      fault["name"].(string),
		})
	}
	return result
}

func expandFaultSpecV3(faultSpecs []interface{}) *chaos.SecurityGovernanceFaultSpec {
	if len(faultSpecs) == 0 {
		return nil
	}
	faultSpec := faultSpecs[0].(map[string]interface{})
	return &chaos.SecurityGovernanceFaultSpec{
		Operator: operatorPtr(faultSpec["operator"].(string)),
		Faults:   expandFaultsV3(faultSpec["faults"].([]interface{})),
	}
}

func expandK8sSpecV3(spec map[string]interface{}) *chaos.SecurityGovernanceK8sSpec {
	k8sSpec := &chaos.SecurityGovernanceK8sSpec{}

	if infraSpecs, ok := spec["infra_spec"].([]interface{}); ok && len(infraSpecs) > 0 {
		infraSpec := infraSpecs[0].(map[string]interface{})
		k8sSpec.InfraSpec = &chaos.SecurityGovernanceInfraSpec{
			Operator: operatorPtr(infraSpec["operator"].(string)),
			InfraIds: interfaceSliceToStringSlice(infraSpec["infra_ids"].([]interface{})),
		}
	}

	if appSpecs, ok := spec["application_spec"].([]interface{}); ok && len(appSpecs) > 0 {
		appSpec := appSpecs[0].(map[string]interface{})
		applicationSpec := &chaos.SecurityGovernanceApplicationSpec{
			Operator: operatorPtr(appSpec["operator"].(string)),
		}
		if workloadList, ok := appSpec["workloads"].([]interface{}); ok {
			workloads := make([]chaos.SecurityGovernanceWorkload, 0, len(workloadList))
			for _, w := range workloadList {
				workload := w.(map[string]interface{})
				wl := chaos.SecurityGovernanceWorkload{
					Namespace: workload["namespace"].(string),
				}
				if v, ok := workload["kind"].(string); ok {
					wl.Kind = v
				}
				if v, ok := workload["label"].(string); ok {
					wl.Label = v
				}
				if v, ok := workload["application_map_id"].(string); ok {
					wl.ApplicationMapId = v
				}
				if services, ok := workload["services"].([]interface{}); ok && len(services) > 0 {
					wl.Services = interfaceSliceToStringSlice(services)
				}
				if nl, ok := workload["namespace_labels"].(map[string]interface{}); ok {
					wl.NamespaceLabels = expandStringMap(nl)
				}
				workloads = append(workloads, wl)
			}
			applicationSpec.Workloads = workloads
		}
		k8sSpec.ApplicationSpec = applicationSpec
	}

	if svcAcctSpecs, ok := spec["chaos_service_account_spec"].([]interface{}); ok && len(svcAcctSpecs) > 0 {
		svcAcctSpec := svcAcctSpecs[0].(map[string]interface{})
		k8sSpec.ChaosServiceAccountSpec = &chaos.SecurityGovernanceChaosServiceAccountSpec{
			Operator:        operatorPtr(svcAcctSpec["operator"].(string)),
			ServiceAccounts: interfaceSliceToStringSlice(svcAcctSpec["service_accounts"].([]interface{})),
		}
	}

	return k8sSpec
}

func expandMachineSpecV3(spec map[string]interface{}) *chaos.SecurityGovernanceMachineSpec {
	if infraSpecs, ok := spec["infra_spec"].([]interface{}); ok && len(infraSpecs) > 0 {
		infraSpec := infraSpecs[0].(map[string]interface{})
		return &chaos.SecurityGovernanceMachineSpec{
			InfraSpec: &chaos.SecurityGovernanceInfraSpec{
				Operator: operatorPtr(infraSpec["operator"].(string)),
				InfraIds: interfaceSliceToStringSlice(infraSpec["infra_ids"].([]interface{})),
			},
		}
	}
	return nil
}

// -----------------------------------------------------------------------------
// flatten: REST response models -> Terraform schema
// -----------------------------------------------------------------------------

func flattenFaultSpecV3(fs *chaos.SecurityGovernanceFaultSpec) []interface{} {
	if fs == nil {
		return nil
	}
	faults := make([]map[string]interface{}, 0, len(fs.Faults))
	for _, f := range fs.Faults {
		faults = append(faults, map[string]interface{}{
			"fault_type": faultTypeFromREST(f.FaultType),
			"name":       f.Name,
		})
	}
	return []interface{}{map[string]interface{}{
		"operator": operatorString(fs.Operator),
		"faults":   faults,
	}}
}

func flattenK8sSpecV3(k8s *chaos.SecurityGovernanceK8sSpec) []interface{} {
	if k8s == nil {
		return nil
	}
	m := map[string]interface{}{}

	if k8s.InfraSpec != nil {
		m["infra_spec"] = []interface{}{map[string]interface{}{
			"operator":  operatorString(k8s.InfraSpec.Operator),
			"infra_ids": k8s.InfraSpec.InfraIds,
		}}
	}

	if k8s.ApplicationSpec != nil {
		appSpec := map[string]interface{}{
			"operator": operatorString(k8s.ApplicationSpec.Operator),
		}
		if k8s.ApplicationSpec.Workloads != nil {
			workloads := make([]map[string]interface{}, 0, len(k8s.ApplicationSpec.Workloads))
			for _, w := range k8s.ApplicationSpec.Workloads {
				wl := map[string]interface{}{
					"namespace":          w.Namespace,
					"kind":               w.Kind,
					"label":              w.Label,
					"services":           w.Services,
					"application_map_id": w.ApplicationMapId,
				}
				if len(w.NamespaceLabels) > 0 {
					nl := make(map[string]interface{}, len(w.NamespaceLabels))
					for k, v := range w.NamespaceLabels {
						nl[k] = v
					}
					wl["namespace_labels"] = nl
				}
				workloads = append(workloads, wl)
			}
			appSpec["workloads"] = workloads
		}
		m["application_spec"] = []interface{}{appSpec}
	}

	if k8s.ChaosServiceAccountSpec != nil {
		m["chaos_service_account_spec"] = []interface{}{map[string]interface{}{
			"operator":         operatorString(k8s.ChaosServiceAccountSpec.Operator),
			"service_accounts": k8s.ChaosServiceAccountSpec.ServiceAccounts,
		}}
	}

	return []interface{}{m}
}

func flattenMachineSpecV3(ms *chaos.SecurityGovernanceMachineSpec) []interface{} {
	if ms == nil || ms.InfraSpec == nil {
		return nil
	}
	return []interface{}{map[string]interface{}{
		"infra_spec": []interface{}{map[string]interface{}{
			"operator":  operatorString(ms.InfraSpec.Operator),
			"infra_ids": ms.InfraSpec.InfraIds,
		}},
	}}
}
