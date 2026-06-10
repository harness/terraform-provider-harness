package security_governance_v3

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/chaos"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// isKubernetesInfra reports whether the infra type uses the k8s_spec block.
func isKubernetesInfra(infraType string) bool {
	return infraType == "Kubernetes" || infraType == "KubernetesV2"
}

// buildConditionV3 assembles the shared condition payload from the resource data.
// It is used both for create (ChaosguardconditionsCreateConditionRequest) and
// update (SecurityGovernanceCondition), which share an identical field set.
func buildConditionV3(d *schema.ResourceData, conditionID string) chaos.SecurityGovernanceCondition {
	infraType := d.Get("infra_type").(string)

	condition := chaos.SecurityGovernanceCondition{
		ConditionId: conditionID,
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		InfraType:   infraTypePtr(infraType),
		FaultSpec:   expandFaultSpecV3(d.Get("fault_spec").([]interface{})),
		Tags:        interfaceSliceToStringSlice(d.Get("tags").([]interface{})),
	}

	if isKubernetesInfra(infraType) {
		if specs := d.Get("k8s_spec").([]interface{}); len(specs) > 0 && specs[0] != nil {
			condition.K8sSpec = expandK8sSpecV3(specs[0].(map[string]interface{}))
		}
	} else {
		if specs := d.Get("machine_spec").([]interface{}); len(specs) > 0 && specs[0] != nil {
			condition.MachineSpec = expandMachineSpecV3(specs[0].(map[string]interface{}))
		}
	}

	return condition
}

func resourceChaosSecurityGovernanceConditionV3Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)

	orgID := d.Get("org_id").(string)
	projectID := d.Get("project_id").(string)

	// Generate a unique, stable condition identifier (consistent with the
	// GraphQL resource so users observe the same behavior).
	conditionID := fmt.Sprintf("tf-condition-%d", time.Now().UnixNano())
	condition := buildConditionV3(d, conditionID)

	req := chaos.ChaosguardconditionsCreateConditionRequest{
		ConditionId: condition.ConditionId,
		Name:        condition.Name,
		Description: condition.Description,
		InfraType:   condition.InfraType,
		FaultSpec:   condition.FaultSpec,
		K8sSpec:     condition.K8sSpec,
		MachineSpec: condition.MachineSpec,
		Tags:        condition.Tags,
	}

	log.Printf("[DEBUG] Creating chaos security governance condition: %s (org: %s, project: %s)", conditionID, orgID, projectID)

	_, httpResp, err := c.CreateconditionApi.CreateCondition(ctx, req, c.AccountId, &chaos.CreateconditionApiCreateConditionOpts{
		OrganizationIdentifier: optional.NewString(orgID),
		ProjectIdentifier:      optional.NewString(projectID),
	})
	if err != nil {
		return helpers.HandleChaosApiError(err, d, httpResp)
	}

	d.SetId(conditionID)

	return resourceChaosSecurityGovernanceConditionV3Read(ctx, d, meta)
}

func resourceChaosSecurityGovernanceConditionV3Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)

	orgID := d.Get("org_id").(string)
	projectID := d.Get("project_id").(string)
	conditionID := d.Id()

	log.Printf("[DEBUG] Reading chaos security governance condition: %s", conditionID)

	resp, httpResp, err := c.GetconditionApi.GetCondition(ctx, c.AccountId, conditionID, &chaos.GetconditionApiGetConditionOpts{
		OrganizationIdentifier: optional.NewString(orgID),
		ProjectIdentifier:      optional.NewString(projectID),
	})
	if err != nil {
		return helpers.HandleChaosReadApiErrorWithGracefulDestroy(err, d, httpResp, []string{
			"not found",
			"no documents in result",
		})
	}

	return setConditionV3Data(d, &resp, orgID, projectID)
}

func resourceChaosSecurityGovernanceConditionV3Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)

	orgID := d.Get("org_id").(string)
	projectID := d.Get("project_id").(string)
	conditionID := d.Id()

	body := buildConditionV3(d, conditionID)

	log.Printf("[DEBUG] Updating chaos security governance condition: %s", conditionID)

	_, httpResp, err := c.UpdateconditionApi.UpdateCondition(ctx, body, c.AccountId, conditionID, &chaos.UpdateconditionApiUpdateConditionOpts{
		OrganizationIdentifier: optional.NewString(orgID),
		ProjectIdentifier:      optional.NewString(projectID),
	})
	if err != nil {
		return helpers.HandleChaosApiError(err, d, httpResp)
	}

	return resourceChaosSecurityGovernanceConditionV3Read(ctx, d, meta)
}

func resourceChaosSecurityGovernanceConditionV3Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)

	orgID := d.Get("org_id").(string)
	projectID := d.Get("project_id").(string)
	conditionID := d.Id()

	log.Printf("[DEBUG] Deleting chaos security governance condition: %s", conditionID)

	_, httpResp, err := c.DeleteconditionApi.DeleteCondition(ctx, c.AccountId, conditionID, &chaos.DeleteconditionApiDeleteConditionOpts{
		OrganizationIdentifier: optional.NewString(orgID),
		ProjectIdentifier:      optional.NewString(projectID),
	})
	if err != nil {
		// Treat a missing condition as already deleted.
		return helpers.HandleChaosReadApiErrorWithGracefulDestroy(err, d, httpResp, []string{
			"not found",
			"no documents in result",
		})
	}

	d.SetId("")
	return nil
}

// setConditionV3Data maps a GetCondition REST response onto the Terraform state.
func setConditionV3Data(d *schema.ResourceData, resp *chaos.ChaosguardconditionsGetConditionResponse, orgID, projectID string) diag.Diagnostics {
	d.Set("org_id", orgID)
	d.Set("project_id", projectID)
	d.Set("name", resp.Name)
	d.Set("description", resp.Description)

	if resp.InfraType != nil {
		d.Set("infra_type", string(*resp.InfraType))
	}

	if len(resp.Tags) > 0 {
		d.Set("tags", resp.Tags)
	}

	if resp.FaultSpec != nil {
		if err := d.Set("fault_spec", flattenFaultSpecV3(resp.FaultSpec)); err != nil {
			return diag.FromErr(fmt.Errorf("failed to set fault_spec: %w", err))
		}
	}

	if resp.K8sSpec != nil {
		if err := d.Set("k8s_spec", flattenK8sSpecV3(resp.K8sSpec)); err != nil {
			return diag.FromErr(fmt.Errorf("failed to set k8s_spec: %w", err))
		}
	}

	if resp.MachineSpec != nil {
		if err := d.Set("machine_spec", flattenMachineSpecV3(resp.MachineSpec)); err != nil {
			return diag.FromErr(fmt.Errorf("failed to set machine_spec: %w", err))
		}
	}

	return nil
}
