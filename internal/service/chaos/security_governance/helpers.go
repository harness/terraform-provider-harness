package security_governance

import (
	"log"

	"github.com/harness/harness-go-sdk/harness/chaos/graphql/model"
)

// Helper function to convert string slice to pointer slice
func stringSliceToPtrSlice(slice []string) []*string {
	result := make([]*string, len(slice))
	for i, v := range slice {
		vCopy := v
		result[i] = &vCopy
	}
	return result
}

// Helper function to convert interface slice to string slice
func interfaceSliceToStringSlice(slice []interface{}) []string {
	result := make([]string, len(slice))
	for i, v := range slice {
		result[i] = v.(string)
	}
	return result
}

// Helper function to expand faults from Terraform schema
func expandFaults(faults []interface{}) []*model.Fault {
	log.Printf("[DEBUG] Expanding %d faults", len(faults))
	result := make([]*model.Fault, len(faults))
	for i, f := range faults {
		fault := f.(map[string]interface{})
		faultType, _ := fault["fault_type"].(string)
		faultName, _ := fault["name"].(string)
		
		// Convert to model.FaultType
		var faultTypeModel model.FaultType
		switch faultType {
		case string(model.FaultTypeFaultGroup):
			faultTypeModel = model.FaultTypeFaultGroup
		default:
			// Default to FAULT if not specified or invalid
			faultTypeModel = model.FaultTypeFault
		}
		
		log.Printf("[DEBUG] Expanded fault %d: Type=%s, Name=%s", i, faultTypeModel, faultName)
		
		result[i] = &model.Fault{
			FaultType: faultTypeModel,
			Name:      faultName,
		}
	}
	log.Printf("[DEBUG] Expanded %d faults successfully", len(result))
	return result
}



// Helper function to expand K8s spec from Terraform schema
func expandK8sSpec(spec map[string]interface{}) *model.K8sSpecInput {
	k8sSpec := &model.K8sSpecInput{}

	// Handle InfraSpec
	if infraSpecs, ok := spec["infra_spec"].([]interface{}); ok && len(infraSpecs) > 0 {
		infraSpec := infraSpecs[0].(map[string]interface{})
		infraIDs := interfaceSliceToStringSlice(infraSpec["infra_ids"].([]interface{}))
		
		k8sSpec.InfraSpec = &model.InfraSpecInput{
			Operator: model.Operator(infraSpec["operator"].(string)),
			InfraIds: infraIDs,
		}
	}

	// Handle ApplicationSpec if provided (using snake_case for Terraform conventions)
	if appSpecs, ok := spec["application_spec"].([]interface{}); ok && len(appSpecs) > 0 {
		appSpec := appSpecs[0].(map[string]interface{})
		
		var workloads []*model.WorkloadInput
		if workloadList, ok := appSpec["workloads"].([]interface{}); ok {
			workloads = make([]*model.WorkloadInput, len(workloadList))
			for i, w := range workloadList {
				workload := w.(map[string]interface{})
				workloadInput := &model.WorkloadInput{
					Namespace: workload["namespace"].(string),
				}

				// Optional fields
				if kind, ok := workload["kind"]; ok && kind != "" {
					workloadInput.Kind = stringPtr(kind.(string))
				}
				if label, ok := workload["label"]; ok && label != "" {
					workloadInput.Label = stringPtr(label.(string))
				}
				if services, ok := workload["services"].([]interface{}); ok && len(services) > 0 {
					workloadInput.Services = interfaceSliceToStringSlice(services)
				}
				if appMapID, ok := workload["application_map_id"]; ok && appMapID != "" {
					workloadInput.ApplicationMapID = stringPtr(appMapID.(string))
				}

				// Handle environment variables if provided
				if envVars, ok := workload["env"].([]interface{}); ok && len(envVars) > 0 {
					envInputs := make([]*model.EnvInput, len(envVars))
					for j, e := range envVars {
						envVar := e.(map[string]interface{})
						envInputs[j] = &model.EnvInput{
							Name:  envVar["name"].(string),
							Value: envVar["value"].(string),
						}
					}
					workloadInput.Env = envInputs
				}

				workloads[i] = workloadInput
			}
		}

		k8sSpec.ApplicationSpec = &model.ApplicationSpecInput{
			Operator:  model.Operator(appSpec["operator"].(string)),
			Workloads: workloads,
		}
	}

	// Handle ChaosServiceAccountSpec if provided (using snake_case for Terraform conventions)
	if svcAcctSpecs, ok := spec["chaos_service_account_spec"].([]interface{}); ok && len(svcAcctSpecs) > 0 {
		svcAcctSpec := svcAcctSpecs[0].(map[string]interface{})
		
		k8sSpec.ChaosServiceAccountSpec = &model.ChaosServiceAccountSpecInput{
			Operator:        model.Operator(svcAcctSpec["operator"].(string)),
			ServiceAccounts: interfaceSliceToStringSlice(svcAcctSpec["service_accounts"].([]interface{})),
		}
	}

	return k8sSpec
}

// Helper function to expand Machine spec from Terraform schema
func expandMachineSpec(spec map[string]interface{}) *model.MachineSpecInput {
	if infraSpecs, ok := spec["infra_spec"].([]interface{}); ok && len(infraSpecs) > 0 {
		infraSpec := infraSpecs[0].(map[string]interface{})
		infraIDs := interfaceSliceToStringSlice(infraSpec["infra_ids"].([]interface{}))
		
		return &model.MachineSpecInput{
			InfraSpec: &model.InfraSpecInput{
				Operator: model.Operator(infraSpec["operator"].(string)),
				InfraIds: infraIDs,
			},
		}
	}
	return nil
}

// Helper function to flatten faults to Terraform schema
func flattenFaults(faults []*model.Fault) []map[string]interface{} {
	if faults == nil {
		return nil
	}

	result := make([]map[string]interface{}, len(faults))
	for i, f := range faults {
		if f == nil {
			continue
		}

		// Ensure we only use valid fault types
		var faultType string
		switch f.FaultType {
		case model.FaultTypeFaultGroup:
			faultType = string(model.FaultTypeFaultGroup)
		default:
			faultType = string(model.FaultTypeFault)
		}

		result[i] = map[string]interface{}{
			"fault_type": faultType,
			"name":       f.Name,
		}
	}
	return result
}

// Helper function to flatten infra spec to Terraform schema
func flattenInfraSpec(infraSpec *model.InfraSpecInput) []map[string]interface{} {
	if infraSpec == nil {
		return nil
	}
	return []map[string]interface{}{
		{
			"operator":  string(infraSpec.Operator),
			"infra_ids": infraSpec.InfraIds,
		},
	}
}
