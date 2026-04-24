package security_governance

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/harness/harness-go-sdk/harness/chaos"
	"github.com/harness/harness-go-sdk/harness/chaos/graphql/model"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Helper function to convert string to string pointer
func stringPtr(s string) *string {
	return &s
}

// Helper function to safely dereference a string pointer
func safeDeref(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func resourceChaosSecurityGovernanceConditionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*internal.Session).ChaosClient
	client := chaos.NewSecurityGovernanceConditionClient(c)

	// Get the account ID from the provider config
	accountID := c.AccountId
	if accountID == "" {
		return helpers.HandleChaosGraphQLError(fmt.Errorf("account ID must be configured in the provider"), d, "create_chaos_security_governance_condition")
	}

	// Get the identifiers from the resource data
	identifiers := getIdentifiers(d, accountID)
	identifiersReq := model.IdentifiersRequest{
		AccountIdentifier: accountID,
	}

	// Set org and project identifiers if they exist
	if identifiers.OrgIdentifier != nil {
		identifiersReq.OrgIdentifier = *identifiers.OrgIdentifier
	}
	if identifiers.ProjectIdentifier != nil {
		identifiersReq.ProjectIdentifier = *identifiers.ProjectIdentifier
	}

	// In resourceChaosSecurityGovernanceConditionCreate
	infraTypeStr := d.Get("infra_type").(string)
	infraType := model.InfrastructureType(infraTypeStr)

	// Generate a unique condition ID
	conditionID := fmt.Sprintf("tf-condition-%d", time.Now().UnixNano())

	// Initialize the request with required fields
	req := model.ConditionRequest{
		ConditionID: conditionID,
		Name:        d.Get("name").(string),
		Description: stringPtr(d.Get("description").(string)),
		InfraType:   infraType,
	}

	// Set tags if provided
	if v, ok := d.GetOk("tags"); ok {
		tags := v.([]interface{})
		req.Tags = make([]*string, len(tags))
		for i, tag := range tags {
			tagStr := tag.(string)
			req.Tags[i] = &tagStr
		}
	}

	// Get and validate fault spec - this is a required field
	faultSpecs := d.Get("fault_spec").([]interface{})
	if len(faultSpecs) == 0 {
		return helpers.HandleChaosGraphQLError(fmt.Errorf("fault_spec is required"), d, "create_chaos_security_governance_condition")
	}

	faultSpec := faultSpecs[0].(map[string]interface{})
	faults := expandFaults(faultSpec["faults"].([]interface{}))
	if len(faults) == 0 {
		return helpers.HandleChaosGraphQLError(fmt.Errorf("at least one fault must be specified in fault_spec"), d, "create_chaos_security_governance_condition")
	}

	// Get the operator from the input
	operator, ok := faultSpec["operator"].(string)
	if !ok || operator == "" {
		return helpers.HandleChaosGraphQLError(fmt.Errorf("operator is required in fault_spec"), d, "create_chaos_security_governance_condition")
	}

	// Initialize FaultSpec with the operator and faults from the input
	faultSpecInput := &model.FaultSpecInput{
		Operator: model.Operator(operator),
		Faults:   faults,
	}
	req.FaultSpec = faultSpecInput

	// Log the fault spec for debugging
	log.Printf("[DEBUG] FaultSpecInput: %+v", faultSpecInput)
	log.Printf("[DEBUG] FaultSpecInput.Operator: %v", faultSpecInput.Operator)
	log.Printf("[DEBUG] Number of Faults: %d", len(faultSpecInput.Faults))
	for i, f := range faultSpecInput.Faults {
		log.Printf("[DEBUG] Fault %d: Type=%s, Name=%s", i, f.FaultType, f.Name)
	}

	// Set K8s spec if provided
	if k8sSpecs, ok := d.GetOk("k8s_spec"); ok && (req.InfraType == model.InfrastructureTypeKubernetes || req.InfraType == model.InfrastructureTypeKubernetesV2) {
		k8sSpecList, ok := k8sSpecs.([]interface{})
		if !ok || len(k8sSpecList) == 0 {
			return helpers.HandleChaosGraphQLError(fmt.Errorf("k8s_spec is required for Kubernetes infrastructure type"), d, "create_chaos_security_governance_condition")
		}

		k8sSpec, ok := k8sSpecList[0].(map[string]interface{})
		if !ok {
			return helpers.HandleChaosGraphQLError(fmt.Errorf("invalid k8s_spec format"), d, "create_chaos_security_governance_condition")
		}

		req.K8sSpec = expandK8sSpec(k8sSpec)
	}

	// Set Machine spec if provided
	if machineSpecs, ok := d.GetOk("machine_spec"); ok &&
		(req.InfraType == model.InfrastructureTypeLinux || req.InfraType == model.InfrastructureTypeWindows) {
		machineSpecList, ok := machineSpecs.([]interface{})
		if !ok || len(machineSpecList) == 0 {
			return helpers.HandleChaosGraphQLError(fmt.Errorf("machine_spec is required for machine infrastructure type"), d, "create_chaos_security_governance_condition")
		}

		machineSpec, ok := machineSpecList[0].(map[string]interface{})
		if !ok {
			return helpers.HandleChaosGraphQLError(fmt.Errorf("invalid machine_spec format"), d, "create_chaos_security_governance_condition")
		}

		req.MachineSpec = expandMachineSpec(machineSpec)
	}

	// Log the request for debugging
	log.Printf("[DEBUG] Creating security governance condition with request: %+v", req)
	log.Printf("[DEBUG] FaultSpec: %+v", req.FaultSpec)
	if req.FaultSpec != nil {
		log.Printf("[DEBUG] FaultSpec.Operator: %v", req.FaultSpec.Operator)
		log.Printf("[DEBUG] Number of Faults: %d", len(req.FaultSpec.Faults))
		for i, f := range req.FaultSpec.Faults {
			log.Printf("[DEBUG] Fault %d: Type=%s, Name=%s", i, f.FaultType, f.Name)
		}
	}

	// Log the entire request for debugging
	reqJSON, err := json.MarshalIndent(req, "", "  ")
	if err != nil {
		log.Printf("[WARN] Failed to marshal request to JSON: %v", err)
	} else {
		log.Printf("[DEBUG] Create request: %s", string(reqJSON))
	}

	// Log the identifiers
	log.Printf("[DEBUG] Identifiers: %+v", identifiersReq)

	// Create the condition
	_, err = client.Create(ctx, identifiersReq, req)
	if err != nil {
		return helpers.HandleChaosGraphQLError(fmt.Errorf("failed to create security governance condition: %v", err), d, "create_chaos_security_governance_condition")
	}

	// Set the ID in the state
	d.SetId(conditionID)

	// Read the condition back to ensure it was created successfully
	return resourceChaosSecurityGovernanceConditionRead(ctx, d, meta)
}

func resourceChaosSecurityGovernanceConditionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*internal.Session).ChaosClient
	client := chaos.NewSecurityGovernanceConditionClient(c)

	// Parse the ID to get the account, org, and project information
	accountID := c.AccountId
	if accountID == "" {
		return helpers.HandleChaosGraphQLError(fmt.Errorf("account ID must be configured in the provider"), d, "read_chaos_security_governance_condition")
	}

	log.Printf("[DEBUG] Reading security governance condition with ID: %s", d.Id())

	// Parse the ID to get the condition ID and scope information
	conditionID := d.Id()
	// _, conditionID, err := parseID(d.Id())
	// if err != nil {
	// 	log.Printf("[ERROR] Failed to parse ID %s: %v", d.Id(), err)
	// 	return diag.Errorf("failed to parse resource ID: %v", err)
	// }

	// Get the identifiers from the resource data
	identifiers := getIdentifiers(d, accountID)
	identifiersReq := model.IdentifiersRequest{
		AccountIdentifier: accountID,
	}

	// Set org and project identifiers if they exist
	if identifiers.OrgIdentifier != nil {
		identifiersReq.OrgIdentifier = *identifiers.OrgIdentifier
	}
	if identifiers.ProjectIdentifier != nil {
		identifiersReq.ProjectIdentifier = *identifiers.ProjectIdentifier
	}

	log.Printf("[DEBUG] Getting condition with ID: %s, Account: %s, Org: %v, Project: %v",
		conditionID,
		identifiersReq.AccountIdentifier,
		identifiersReq.OrgIdentifier,
		identifiersReq.ProjectIdentifier)

	// Get the condition
	resp, err := client.Get(ctx, identifiersReq, conditionID)
	if err != nil {
		if strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "no documents in result") {
			log.Printf("[WARN] Condition %s not found, removing from state: %v", conditionID, err)
			d.SetId("")
			return nil
		}
		log.Printf("[ERROR] Failed to get condition %s: %v", conditionID, err)
		return helpers.HandleChaosGraphQLError(fmt.Errorf("failed to read security governance condition: %v", err), d, "read_chaos_security_governance_condition")
	}

	// The Get response should be a ConditionResponse
	if resp == nil || resp.Condition == nil {
		log.Printf("[WARN] Condition %s not found in response, removing from state", conditionID)
		d.SetId("")
		return nil
	}

	log.Printf("[DEBUG] Successfully retrieved condition: %+v", resp.Condition)

	// The actual condition data is in the response
	condition := resp.Condition

	// Set the resource data
	d.Set("name", condition.Name)
	d.Set("description", condition.Description)
	d.Set("infra_type", condition.InfraType)

	// Set tags
	if condition.Tags != nil {
		tags := make([]string, 0, len(condition.Tags))
		for _, tag := range condition.Tags {
			tags = append(tags, tag)
		}
		d.Set("tags", tags)
	}

	// Set fault spec
	if condition.FaultSpec != nil {
		faultSpec := map[string]interface{}{
			"operator": string(condition.FaultSpec.Operator),
		}

		// Convert FaultResponse to Fault for flattening
		if condition.FaultSpec.Faults != nil {
			faults := make([]map[string]string, len(condition.FaultSpec.Faults))
			for i, f := range condition.FaultSpec.Faults {
				faults[i] = map[string]string{
					"fault_type": string(f.FaultType),
					"name":       f.Name,
				}
			}
			faultSpec["faults"] = faults
		}

		d.Set("fault_spec", []interface{}{faultSpec})
	}

	// Set K8s spec if present
	if condition.K8sSpec != nil {
		k8sSpec := make(map[string]interface{})

		// Handle infra_spec
		if condition.K8sSpec.InfraSpec != nil {
			k8sSpec["infra_spec"] = []map[string]interface{}{
				{
					"operator":  string(condition.K8sSpec.InfraSpec.Operator),
					"infra_ids": condition.K8sSpec.InfraSpec.InfraIds,
				},
			}
		}

		// Handle application_spec
		if condition.K8sSpec.ApplicationSpec != nil {
			appSpec := map[string]interface{}{
				"operator": string(condition.K8sSpec.ApplicationSpec.Operator),
			}

			// Convert workloads if present
			if condition.K8sSpec.ApplicationSpec.Workloads != nil {
				workloads := make([]map[string]interface{}, len(condition.K8sSpec.ApplicationSpec.Workloads))
				for i, w := range condition.K8sSpec.ApplicationSpec.Workloads {
					workload := map[string]interface{}{
						"namespace":          w.Namespace,
						"kind":               safeDeref(w.Kind),
						"label":              safeDeref(w.Label),
						"services":           w.Services,
						"application_map_id": safeDeref(w.ApplicationMapID),
					}
					workloads[i] = workload
				}
				appSpec["workloads"] = workloads
			}

			k8sSpec["application_spec"] = []map[string]interface{}{appSpec}
		}

		// Handle chaos_service_account_spec
		if condition.K8sSpec.ChaosServiceAccountSpec != nil {
			k8sSpec["chaos_service_account_spec"] = []map[string]interface{}{
				{
					"operator":         string(condition.K8sSpec.ChaosServiceAccountSpec.Operator),
					"service_accounts": condition.K8sSpec.ChaosServiceAccountSpec.ServiceAccounts,
				},
			}
		}

		// Only set k8s_spec if we have at least one spec defined
		if len(k8sSpec) > 0 {
			d.Set("k8s_spec", []interface{}{k8sSpec})
		}
	}

	// Set Machine spec if present
	if condition.MachineSpec != nil && condition.MachineSpec.InfraSpec != nil {
		machineSpec := map[string]interface{}{
			"infra_spec": []map[string]interface{}{
				{
					"operator":  string(condition.MachineSpec.InfraSpec.Operator),
					"infra_ids": condition.MachineSpec.InfraSpec.InfraIds,
				},
			},
		}
		d.Set("machine_spec", []interface{}{machineSpec})
	}

	return nil
}

func resourceChaosSecurityGovernanceConditionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*internal.Session).ChaosClient
	client := chaos.NewSecurityGovernanceConditionClient(c)

	// Parse the ID to get the account, org, and project information
	accountID := c.AccountId
	if accountID == "" {
		return helpers.HandleChaosGraphQLError(fmt.Errorf("account ID must be configured in the provider"), d, "update_chaos_security_governance_condition")
	}

	conditionID := d.Id()

	// Get the identifiers from the resource data
	identifiers := getIdentifiers(d, accountID)
	identifiersReq := model.IdentifiersRequest{
		AccountIdentifier: accountID,
	}

	// Set org and project identifiers if they exist
	if identifiers.OrgIdentifier != nil {
		identifiersReq.OrgIdentifier = *identifiers.OrgIdentifier
	}
	if identifiers.ProjectIdentifier != nil {
		identifiersReq.ProjectIdentifier = *identifiers.ProjectIdentifier
	}

	// First, get the existing condition
	existing, err := client.Get(ctx, identifiersReq, conditionID)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			d.SetId("")
			return helpers.HandleChaosGraphQLError(fmt.Errorf("security governance condition not found, removing from state: %v", err), d, "update_chaos_security_governance_condition")
		}
		return helpers.HandleChaosGraphQLError(fmt.Errorf("failed to get existing condition for update: %v", err), d, "update_chaos_security_governance_condition")
	}

	// Convert the existing condition to an update request
	req := model.ConditionRequest{
		ConditionID: conditionID,
		Name:        d.Get("name").(string),
		Description: stringPtr(d.Get("description").(string)),
		InfraType:   existing.Condition.InfraType, // Keep the existing infra type
		Tags:        make([]*string, 0),           // Initialize empty tags slice
	}

	// Handle fault spec updates
	if d.HasChange("fault_spec") {
		faultSpecs := d.Get("fault_spec").([]interface{})
		if len(faultSpecs) == 0 {
			return helpers.HandleChaosGraphQLError(fmt.Errorf("fault_spec is required"), d, "update_chaos_security_governance_condition")
		}

		faultSpec := faultSpecs[0].(map[string]interface{})
		faults := expandFaults(faultSpec["faults"].([]interface{}))
		if len(faults) == 0 {
			return helpers.HandleChaosGraphQLError(fmt.Errorf("at least one fault must be specified in fault_spec"), d, "update_chaos_security_governance_condition")
		}

		// Initialize FaultSpec with the operator and faults from the input
		faultSpecInput := &model.FaultSpecInput{
			Operator: model.Operator(faultSpec["operator"].(string)),
			Faults:   faults,
		}
		req.FaultSpec = faultSpecInput

		// Log the fault spec for debugging
		log.Printf("[DEBUG] Update - FaultSpecInput: %+v", faultSpecInput)
		log.Printf("[DEBUG] Update - FaultSpecInput.Operator: %v", faultSpecInput.Operator)
		log.Printf("[DEBUG] Update - Number of Faults: %d", len(faultSpecInput.Faults))
		for i, f := range faultSpecInput.Faults {
			log.Printf("[DEBUG] Update - Fault %d: Type=%s, Name=%s", i, f.FaultType, f.Name)
		}
	} else if existing.Condition.FaultSpec != nil {
		// If fault_spec hasn't changed, use the existing one
		faults := make([]*model.Fault, len(existing.Condition.FaultSpec.Faults))
		for i, f := range existing.Condition.FaultSpec.Faults {
			faults[i] = &model.Fault{
				FaultType: f.FaultType,
				Name:      f.Name,
			}
		}
		req.FaultSpec = &model.FaultSpecInput{
			Operator: existing.Condition.FaultSpec.Operator,
			Faults:   faults,
		}
	} else {
		// This should not happen as fault_spec is required
		return helpers.HandleChaosGraphQLError(fmt.Errorf("fault_spec is required"), d, "update_chaos_security_governance_condition")
	}

	// Convert K8sSpec if it exists
	if existing.Condition.K8sSpec != nil {
		k8sSpec := &model.K8sSpecInput{}

		// Handle InfraSpec
		if existing.Condition.K8sSpec.InfraSpec != nil {
			k8sSpec.InfraSpec = &model.InfraSpecInput{
				Operator: existing.Condition.K8sSpec.InfraSpec.Operator,
				InfraIds: existing.Condition.K8sSpec.InfraSpec.InfraIds,
			}
		}

		// Handle ApplicationSpec
		if existing.Condition.K8sSpec.ApplicationSpec != nil {
			appSpec := &model.ApplicationSpecInput{
				Operator: existing.Condition.K8sSpec.ApplicationSpec.Operator,
			}

			// Handle Workloads
			if existing.Condition.K8sSpec.ApplicationSpec.Workloads != nil {
				workloads := make([]*model.WorkloadInput, len(existing.Condition.K8sSpec.ApplicationSpec.Workloads))
				for i, w := range existing.Condition.K8sSpec.ApplicationSpec.Workloads {
					workload := &model.WorkloadInput{
						Namespace: w.Namespace,
						Kind:      w.Kind,
						Label:     w.Label,
						Services:  w.Services,
					}
					if w.ApplicationMapID != nil {
						workload.ApplicationMapID = w.ApplicationMapID
					}
					workloads[i] = workload
				}
				appSpec.Workloads = workloads
			}

			k8sSpec.ApplicationSpec = appSpec
		}

		// Handle ChaosServiceAccountSpec
		if existing.Condition.K8sSpec.ChaosServiceAccountSpec != nil {
			k8sSpec.ChaosServiceAccountSpec = &model.ChaosServiceAccountSpecInput{
				Operator:        existing.Condition.K8sSpec.ChaosServiceAccountSpec.Operator,
				ServiceAccounts: existing.Condition.K8sSpec.ChaosServiceAccountSpec.ServiceAccounts,
			}
		}

		req.K8sSpec = k8sSpec
	}

	// Convert MachineSpec if it exists
	if existing.Condition.MachineSpec != nil {
		machineSpec := &model.MachineSpecInput{}

		// Handle InfraSpec if it exists
		if existing.Condition.MachineSpec.InfraSpec != nil {
			machineSpec.InfraSpec = &model.InfraSpecInput{
				Operator: existing.Condition.MachineSpec.InfraSpec.Operator,
				InfraIds: existing.Condition.MachineSpec.InfraSpec.InfraIds,
			}
		}

		req.MachineSpec = machineSpec
	}

	// Set tags if provided
	if v, ok := d.GetOk("tags"); ok {
		tags := v.([]interface{})
		req.Tags = make([]*string, len(tags))
		for i, tag := range tags {
			tagStr := tag.(string)
			req.Tags[i] = &tagStr
		}
	}

	// Set fault spec if provided
	if d.HasChange("fault_spec") {
		faultSpecs := d.Get("fault_spec").([]interface{})
		if len(faultSpecs) > 0 {
			faultSpec := faultSpecs[0].(map[string]interface{})
			req.FaultSpec = &model.FaultSpecInput{
				Operator: model.Operator(faultSpec["operator"].(string)),
				Faults:   expandFaults(faultSpec["faults"].([]interface{})),
			}
		}
	}

	// Set K8s spec if provided
	if d.HasChange("k8s_spec") {
		if k8sSpecs, ok := d.GetOk("k8s_spec"); ok {
			k8sSpec := k8sSpecs.([]interface{})[0].(map[string]interface{})
			req.K8sSpec = expandK8sSpec(k8sSpec)
		}
	}

	// Set Machine spec if provided
	if d.HasChange("machine_spec") {
		if machineSpecs, ok := d.GetOk("machine_spec"); ok {
			machineSpec := machineSpecs.([]interface{})[0].(map[string]interface{})
			req.MachineSpec = expandMachineSpec(machineSpec)
		}
	}

	// Update the condition
	_, err = client.Update(ctx, identifiersReq, req)
	if err != nil {
		return helpers.HandleChaosGraphQLError(fmt.Errorf("failed to update security governance condition: %v", err), d, "update_chaos_security_governance_condition")
	}

	return resourceChaosSecurityGovernanceConditionRead(ctx, d, meta)
}

func resourceChaosSecurityGovernanceConditionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Deleting security governance condition with ID: %s", d.Id())

	c := meta.(*internal.Session).ChaosClient
	client := chaos.NewSecurityGovernanceConditionClient(c)

	// Parse the ID to get the account, org, and project information
	accountID := c.AccountId
	if accountID == "" {
		err := "account ID must be configured in the provider"
		log.Printf("[ERROR] %s", err)
		return helpers.HandleChaosGraphQLError(fmt.Errorf("%s", err), d, "delete_chaos_security_governance_condition")
	}
	conditionID := d.Id()

	// Get the identifiers from the resource data
	identifiers := getIdentifiers(d, accountID)
	identifiersReq := model.IdentifiersRequest{
		AccountIdentifier: accountID,
	}

	// Set org and project identifiers if they exist
	if identifiers.OrgIdentifier != nil {
		identifiersReq.OrgIdentifier = *identifiers.OrgIdentifier
	}
	if identifiers.ProjectIdentifier != nil {
		identifiersReq.ProjectIdentifier = *identifiers.ProjectIdentifier
	}

	log.Printf("[DEBUG] Deleting condition with ID: %s, Account: %s, Org: %v, Project: %v",
		conditionID,
		identifiersReq.AccountIdentifier,
		identifiersReq.OrgIdentifier,
		identifiersReq.ProjectIdentifier)

	// Delete the condition
	_, err := client.Delete(ctx, identifiersReq, conditionID)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			log.Printf("[DEBUG] Condition %s not found, assuming it's already deleted", conditionID)
			d.SetId("")
			return nil
		}
		errMsg := fmt.Sprintf("failed to delete security governance condition: %v", err)
		log.Printf("[ERROR] %s", errMsg)
		return helpers.HandleChaosGraphQLError(fmt.Errorf("%s", errMsg), d, "delete_chaos_security_governance_condition")
	}

	log.Printf("[DEBUG] Successfully deleted condition with ID: %s", conditionID)
	d.SetId("")
	return nil
}
