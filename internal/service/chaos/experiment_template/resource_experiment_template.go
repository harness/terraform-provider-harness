package experiment_template

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/chaos"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceExperimentTemplate returns the experiment template resource
func ResourceExperimentTemplate() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for managing Harness Chaos Experiment Templates. " +
			"Experiment templates define reusable chaos experiments with actions, faults, and probes.",

		CreateContext: resourceExperimentTemplateCreate,
		ReadContext:   resourceExperimentTemplateRead,
		UpdateContext: resourceExperimentTemplateUpdate,
		DeleteContext: resourceExperimentTemplateDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceExperimentTemplateImport,
		},

		Schema: ResourceExperimentTemplateSchema(),
	}
}

func resourceExperimentTemplateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)

	// Get scope identifiers
	accountID := c.AccountId
	orgID := d.Get("org_id").(string)
	projectID := d.Get("project_id").(string)
	hubIdentity := d.Get("hub_identity").(string)
	identity := d.Get("identity").(string)

	log.Printf("[DEBUG] Creating experiment template: identity=%s, hub=%s", identity, hubIdentity)

	// Build the YAML manifest
	manifest, err := buildExperimentTemplateManifest(d)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to build experiment template manifest: %w", err))
	}

	// Log the manifest for debugging
	log.Printf("[DEBUG] Experiment template manifest:\n%s", manifest)

	// Create request
	req := chaos.ChaosexperimenttemplateCreateExperimentTemplateRequest{
		Manifest: manifest,
	}

	// Create experiment template
	opts := &chaos.ExperimenttemplateApiCreateExperimentTemplateOpts{}
	opts.HubIdentity = optional.NewString(hubIdentity)
	opts.OrganizationIdentifier = optional.NewString(orgID)
	opts.ProjectIdentifier = optional.NewString(projectID)

	resp, httpResp, err := c.ExperimenttemplateApi.CreateExperimentTemplate(ctx, req, accountID, opts)
	if err != nil {
		log.Printf("[ERROR] Failed to create experiment template: %v", err)
		if httpResp != nil {
			log.Printf("[ERROR] HTTP Status: %d", httpResp.StatusCode)
			// Try to read response body for better error message
			if httpResp.Body != nil {
				body := make([]byte, 1024)
				n, _ := httpResp.Body.Read(body)
				if n > 0 {
					log.Printf("[ERROR] Response body: %s", string(body[:n]))
				}
			}
		}
		return helpers.HandleChaosApiError(err, d, httpResp)
	}

	log.Printf("[DEBUG] Experiment template created: %s", resp.Identity)

	// Set ID: org_id/project_id/hub_identity/identity
	d.SetId(fmt.Sprintf("%s/%s/%s/%s", orgID, projectID, hubIdentity, resp.Identity))

	return resourceExperimentTemplateRead(ctx, d, meta)
}

func resourceExperimentTemplateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)

	// Parse ID
	parts := strings.Split(d.Id(), "/")
	if len(parts) != 4 {
		return diag.Errorf("invalid experiment template ID format: %s (expected: org_id/project_id/hub_identity/identity)", d.Id())
	}

	accountID := c.AccountId
	orgID := parts[0]
	projectID := parts[1]
	hubIdentity := parts[2]
	identity := parts[3]

	log.Printf("[DEBUG] Reading experiment template: identity=%s, hub=%s", identity, hubIdentity)

	// Get experiment template
	opts := &chaos.ExperimenttemplateApiGetExperimentTemplateOpts{}
	opts.HubIdentity = optional.NewString(hubIdentity)
	opts.OrganizationIdentifier = optional.NewString(orgID)
	opts.ProjectIdentifier = optional.NewString(projectID)

	resp, httpResp, err := c.ExperimenttemplateApi.GetExperimentTemplate(ctx, accountID, identity, opts)
	if err != nil {
		if httpResp != nil && httpResp.StatusCode == 404 {
			log.Printf("[WARN] Experiment template not found, removing from state: %s", identity)
			d.SetId("")
			return nil
		}
		return helpers.HandleChaosReadApiError(err, d, httpResp)
	}

	// Set data from response
	if err := setExperimentTemplateData(d, &resp, orgID, projectID, hubIdentity); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceExperimentTemplateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)

	// Parse ID
	parts := strings.Split(d.Id(), "/")
	if len(parts) != 4 {
		return diag.Errorf("invalid experiment template ID format: %s", d.Id())
	}

	accountID := c.AccountId
	orgID := parts[0]
	projectID := parts[1]
	hubIdentity := parts[2]
	identity := parts[3]

	log.Printf("[DEBUG] Updating experiment template: identity=%s, hub=%s", identity, hubIdentity)

	// Build the YAML manifest
	manifest, err := buildExperimentTemplateManifest(d)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to build experiment template manifest: %w", err))
	}

	// Update request
	req := chaos.ChaosexperimenttemplateUpdateExperimentTemplateRequest{
		Manifest: manifest,
	}

	// Update experiment template
	opts := &chaos.ExperimenttemplateApiUpdateExperimentTemplateOpts{}
	opts.HubIdentity = optional.NewString(hubIdentity)
	opts.OrganizationIdentifier = optional.NewString(orgID)
	opts.ProjectIdentifier = optional.NewString(projectID)

	_, httpResp, err := c.ExperimenttemplateApi.UpdateExperimentTemplate(ctx, req, accountID, identity, opts)
	if err != nil {
		log.Printf("[ERROR] Failed to update experiment template: %v", err)
		return helpers.HandleChaosApiError(err, d, httpResp)
	}

	log.Printf("[DEBUG] Experiment template updated: %s", identity)

	return resourceExperimentTemplateRead(ctx, d, meta)
}

func resourceExperimentTemplateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)

	// Parse ID
	parts := strings.Split(d.Id(), "/")
	if len(parts) != 4 {
		return diag.Errorf("invalid experiment template ID format: %s", d.Id())
	}

	accountID := c.AccountId
	orgID := parts[0]
	projectID := parts[1]
	hubIdentity := parts[2]
	identity := parts[3]

	log.Printf("[DEBUG] Deleting experiment template: identity=%s, hub=%s", identity, hubIdentity)

	// Delete experiment template
	opts := &chaos.ExperimenttemplateApiDeleteExperimentTemplateOpts{}
	opts.HubIdentity = optional.NewString(hubIdentity)
	opts.OrganizationIdentifier = optional.NewString(orgID)
	opts.ProjectIdentifier = optional.NewString(projectID)

	_, httpResp, err := c.ExperimenttemplateApi.DeleteExperimentTemplate(ctx, accountID, identity, opts)
	if err != nil {
		if httpResp != nil && httpResp.StatusCode == 404 {
			log.Printf("[WARN] Experiment template already deleted: %s", identity)
			return nil
		}
		return helpers.HandleChaosApiError(err, d, httpResp)
	}

	log.Printf("[DEBUG] Experiment template deleted: %s", identity)
	return nil
}

func resourceExperimentTemplateImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	// Import format: org_id/project_id/hub_identity/identity
	parts := strings.Split(d.Id(), "/")
	if len(parts) != 4 {
		return nil, fmt.Errorf("invalid import ID format: %s (expected: org_id/project_id/hub_identity/identity)", d.Id())
	}

	orgID := parts[0]
	projectID := parts[1]
	hubIdentity := parts[2]
	identity := parts[3]

	// Set the scope identifiers
	d.Set("org_id", orgID)
	d.Set("project_id", projectID)
	d.Set("hub_identity", hubIdentity)
	d.Set("identity", identity)

	// Read the resource
	diags := resourceExperimentTemplateRead(ctx, d, meta)
	if diags.HasError() {
		return nil, fmt.Errorf("failed to read experiment template during import: %v", diags)
	}

	return []*schema.ResourceData{d}, nil
}

// ============================================================================
// BUILD FUNCTIONS: Create YAML Manifest
// ============================================================================

func buildExperimentTemplateManifest(d *schema.ResourceData) (string, error) {
	// Build the template structure
	template := map[string]interface{}{
		"identity":   d.Get("identity").(string),
		"name":       d.Get("name").(string),
		"kind":       "ChaosExperimentTemplate",
		"apiVersion": "litmuschaos.io/v1beta1",
		"revision":   "v1", // Default revision
		"isDefault":  true, // Default value
	}

	// Optional fields
	if v, ok := d.GetOk("description"); ok {
		template["description"] = v.(string)
	}

	if v, ok := d.GetOk("tags"); ok {
		tags := []string{}
		for _, tag := range v.([]interface{}) {
			tags = append(tags, tag.(string))
		}
		template["tags"] = tags
	}

	// Build spec
	if specList, ok := d.GetOk("spec"); ok && len(specList.([]interface{})) > 0 {
		specData := specList.([]interface{})[0].(map[string]interface{})
		spec := map[string]interface{}{}

		// Infrastructure
		// Note: infraId is NOT included in the manifest - it's only used at runtime
		if v, ok := specData["infra_type"].(string); ok && v != "" {
			spec["infraType"] = v
		}

		// Actions
		if actions, ok := specData["actions"].([]interface{}); ok && len(actions) > 0 {
			spec["actions"] = buildActions(actions)
		}

		// Faults
		if faults, ok := specData["faults"].([]interface{}); ok && len(faults) > 0 {
			spec["faults"] = buildFaults(faults)
		}

		// Probes
		if probes, ok := specData["probes"].([]interface{}); ok && len(probes) > 0 {
			spec["probes"] = buildProbes(probes)
		}

		// Vertices
		if vertices, ok := specData["vertices"].([]interface{}); ok && len(vertices) > 0 {
			spec["vertices"] = buildVertices(vertices)
		}

		// Cleanup Policy
		if v, ok := specData["cleanup_policy"].(string); ok && v != "" {
			spec["cleanupPolicy"] = v
		}

		// Status Check Timeouts
		if timeouts, ok := specData["status_check_timeouts"].([]interface{}); ok && len(timeouts) > 0 {
			timeoutData := timeouts[0].(map[string]interface{})
			timeout := map[string]interface{}{}
			if v, ok := timeoutData["delay"].(int); ok && v > 0 {
				timeout["delay"] = v
			}
			if v, ok := timeoutData["timeout"].(int); ok && v > 0 {
				timeout["timeout"] = v
			}
			if len(timeout) > 0 {
				spec["statusCheckTimeouts"] = timeout
			}
		}

		template["spec"] = spec
	}

	// Convert to JSON (API expects JSON, not YAML)
	jsonBytes, err := json.Marshal(template)
	if err != nil {
		return "", fmt.Errorf("failed to marshal template to JSON: %w", err)
	}

	return string(jsonBytes), nil
}

func buildActions(actions []interface{}) []map[string]interface{} {
	result := make([]map[string]interface{}, 0, len(actions))
	for _, actionItem := range actions {
		actionData := actionItem.(map[string]interface{})
		action := map[string]interface{}{
			"identity": actionData["identity"].(string),
			"name":     actionData["name"].(string),
		}

		if v, ok := actionData["infra_id"].(string); ok && v != "" {
			action["infraId"] = v
		}
		if v, ok := actionData["revision"].(int); ok && v > 0 {
			action["revision"] = v
		}
		if v, ok := actionData["is_enterprise"].(bool); ok {
			action["isEnterprise"] = v
		}
		if v, ok := actionData["continue_on_completion"].(bool); ok {
			action["continueOnCompletion"] = v
		}

		// Values - Always include, even if empty
		if values, ok := actionData["values"].([]interface{}); ok && len(values) > 0 {
			action["values"] = buildValues(values)
		} else {
			action["values"] = []map[string]interface{}{}
		}

		result = append(result, action)
	}
	return result
}

func buildFaults(faults []interface{}) []map[string]interface{} {
	result := make([]map[string]interface{}, 0, len(faults))
	for _, faultItem := range faults {
		faultData := faultItem.(map[string]interface{})
		fault := map[string]interface{}{
			"identity": faultData["identity"].(string),
			"name":     faultData["name"].(string),
		}

		if v, ok := faultData["infra_id"].(string); ok && v != "" {
			fault["infraId"] = v
		}
		if v, ok := faultData["revision"].(string); ok && v != "" {
			fault["revision"] = v
		}
		if v, ok := faultData["is_enterprise"].(bool); ok {
			fault["isEnterprise"] = v
		}
		if v, ok := faultData["auth_enabled"].(bool); ok {
			fault["authEnabled"] = v
		}

		// Values - Always include, even if empty
		if values, ok := faultData["values"].([]interface{}); ok && len(values) > 0 {
			fault["values"] = buildValues(values)
		} else {
			fault["values"] = []map[string]interface{}{}
		}

		result = append(result, fault)
	}
	return result
}

func buildProbes(probes []interface{}) []map[string]interface{} {
	result := make([]map[string]interface{}, 0, len(probes))
	for _, probeItem := range probes {
		probeData := probeItem.(map[string]interface{})
		probe := map[string]interface{}{
			"identity": probeData["identity"].(string),
			"name":     probeData["name"].(string),
		}

		if v, ok := probeData["infra_id"].(string); ok && v != "" {
			probe["infraId"] = v
		}
		if v, ok := probeData["revision"].(int); ok && v > 0 {
			probe["revision"] = v
		}
		if v, ok := probeData["is_enterprise"].(bool); ok {
			probe["isEnterprise"] = v
		}
		if v, ok := probeData["duration"].(string); ok && v != "" {
			probe["duration"] = v
		}
		if v, ok := probeData["weightage"].(int); ok && v > 0 {
			probe["weightage"] = v
		}
		// Note: enableDataCollection and conditions are NOT sent in the manifest
		// They are only present in the API response, not in the request

		// Values - Always include, even if empty
		if values, ok := probeData["values"].([]interface{}); ok && len(values) > 0 {
			probe["values"] = buildValues(values)
		} else {
			probe["values"] = []map[string]interface{}{}
		}

		result = append(result, probe)
	}
	return result
}

func buildVertices(vertices []interface{}) []map[string]interface{} {
	result := make([]map[string]interface{}, 0, len(vertices))
	for _, vertexItem := range vertices {
		vertexData := vertexItem.(map[string]interface{})
		vertex := map[string]interface{}{
			"name": vertexData["name"].(string),
		}

		// Start
		if start, ok := vertexData["start"].([]interface{}); ok && len(start) > 0 && start[0] != nil {
			startData := start[0].(map[string]interface{})
			startObj := map[string]interface{}{}

			if actions, ok := startData["actions"].([]interface{}); ok && len(actions) > 0 {
				actionList := make([]map[string]interface{}, 0, len(actions))
				for _, actionItem := range actions {
					actionData := actionItem.(map[string]interface{})
					actionList = append(actionList, map[string]interface{}{
						"name": actionData["name"].(string),
					})
				}
				startObj["actions"] = actionList
			}

			if faults, ok := startData["faults"].([]interface{}); ok && len(faults) > 0 {
				faultList := make([]map[string]interface{}, 0, len(faults))
				for _, faultItem := range faults {
					faultData := faultItem.(map[string]interface{})
					faultList = append(faultList, map[string]interface{}{
						"name": faultData["name"].(string),
					})
				}
				startObj["faults"] = faultList
			}

			if probes, ok := startData["probes"].([]interface{}); ok && len(probes) > 0 {
				probeList := make([]map[string]interface{}, 0, len(probes))
				for _, probeItem := range probes {
					probeData := probeItem.(map[string]interface{})
					probeList = append(probeList, map[string]interface{}{
						"name": probeData["name"].(string),
					})
				}
				startObj["probes"] = probeList
			}

			// Always include start block (even if empty)
			vertex["start"] = startObj
		} else {
			// No start block defined, add empty one
			vertex["start"] = map[string]interface{}{}
		}

		// End
		if end, ok := vertexData["end"].([]interface{}); ok && len(end) > 0 && end[0] != nil {
			endData := end[0].(map[string]interface{})
			endObj := map[string]interface{}{}

			if actions, ok := endData["actions"].([]interface{}); ok && len(actions) > 0 {
				actionList := make([]map[string]interface{}, 0, len(actions))
				for _, actionItem := range actions {
					actionData := actionItem.(map[string]interface{})
					actionList = append(actionList, map[string]interface{}{
						"name": actionData["name"].(string),
					})
				}
				endObj["actions"] = actionList
			}

			if faults, ok := endData["faults"].([]interface{}); ok && len(faults) > 0 {
				faultList := make([]map[string]interface{}, 0, len(faults))
				for _, faultItem := range faults {
					faultData := faultItem.(map[string]interface{})
					faultList = append(faultList, map[string]interface{}{
						"name": faultData["name"].(string),
					})
				}
				endObj["faults"] = faultList
			}

			if probes, ok := endData["probes"].([]interface{}); ok && len(probes) > 0 {
				probeList := make([]map[string]interface{}, 0, len(probes))
				for _, probeItem := range probes {
					probeData := probeItem.(map[string]interface{})
					probeList = append(probeList, map[string]interface{}{
						"name": probeData["name"].(string),
					})
				}
				endObj["probes"] = probeList
			}

			// Always include end block (even if empty)
			vertex["end"] = endObj
		} else {
			// No end block defined, add empty one
			vertex["end"] = map[string]interface{}{}
		}

		result = append(result, vertex)
	}
	return result
}

func buildValues(values []interface{}) []map[string]interface{} {
	result := make([]map[string]interface{}, 0, len(values))
	for _, valueItem := range values {
		valueData := valueItem.(map[string]interface{})
		value := map[string]interface{}{
			"name":  valueData["name"].(string),
			"value": valueData["value"].(string),
		}
		result = append(result, value)
	}
	return result
}

// ============================================================================
// READ FUNCTIONS: Parse Response and Set State
// ============================================================================

func setExperimentTemplateData(d *schema.ResourceData, template *chaos.ChaosexperimenttemplateGetExperimentTemplateResponse, orgID, projectID, hubIdentity string) error {
	// Set scope identifiers
	d.Set("org_id", orgID)
	d.Set("project_id", projectID)
	d.Set("hub_identity", hubIdentity)

	// Set basic fields
	d.Set("identity", template.Identity)
	d.Set("name", template.Name)
	d.Set("description", template.Description)
	d.Set("is_default", template.IsDefault)
	d.Set("revision", template.Revision)
	d.Set("api_version", template.ApiVersion)
	d.Set("kind", template.Kind)

	// Tags
	if len(template.Tags) > 0 {
		d.Set("tags", template.Tags)
	}

	// Spec
	if template.Spec != nil {
		specBlock := map[string]interface{}{}

		specBlock["infra_id"] = template.Spec.InfraId
		if template.Spec.InfraType != nil {
			specBlock["infra_type"] = string(*template.Spec.InfraType)
		}

		// Actions
		if len(template.Spec.Actions) > 0 {
			specBlock["actions"] = readActions(template.Spec.Actions)
		}

		// Faults
		if len(template.Spec.Faults) > 0 {
			specBlock["faults"] = readFaults(template.Spec.Faults)
		}

		// Probes
		if len(template.Spec.Probes) > 0 {
			specBlock["probes"] = readProbes(template.Spec.Probes)
		}

		// Vertices
		if len(template.Spec.Vertices) > 0 {
			specBlock["vertices"] = readVertices(template.Spec.Vertices)
		}

		// Cleanup Policy
		if template.Spec.CleanupPolicy != nil {
			specBlock["cleanup_policy"] = string(*template.Spec.CleanupPolicy)
		}

		// Status Check Timeouts
		if template.Spec.StatusCheckTimeouts != nil {
			timeoutBlock := map[string]interface{}{}
			if template.Spec.StatusCheckTimeouts.Delay > 0 {
				timeoutBlock["delay"] = template.Spec.StatusCheckTimeouts.Delay
			}
			if template.Spec.StatusCheckTimeouts.Timeout > 0 {
				timeoutBlock["timeout"] = template.Spec.StatusCheckTimeouts.Timeout
			}
			if len(timeoutBlock) > 0 {
				specBlock["status_check_timeouts"] = []map[string]interface{}{timeoutBlock}
			}
		}

		d.Set("spec", []map[string]interface{}{specBlock})
	}

	return nil
}

func readActions(actions []chaos.ExperimenttemplateAction) []map[string]interface{} {
	result := make([]map[string]interface{}, 0, len(actions))
	for _, action := range actions {
		actionBlock := map[string]interface{}{
			"identity": action.Identity,
			"name":     action.Name,
		}

		if action.InfraId != "" {
			actionBlock["infra_id"] = action.InfraId
		}
		if action.Revision > 0 {
			actionBlock["revision"] = action.Revision
		}
		actionBlock["is_enterprise"] = action.IsEnterprise
		actionBlock["continue_on_completion"] = action.ContinueOnCompletion

		// Values
		if len(action.Values) > 0 {
			actionBlock["values"] = readValues(action.Values)
		}

		result = append(result, actionBlock)
	}
	return result
}

func readFaults(faults []chaos.ExperimenttemplateFault) []map[string]interface{} {
	result := make([]map[string]interface{}, 0, len(faults))
	for _, fault := range faults {
		faultBlock := map[string]interface{}{
			"identity": fault.Identity,
			"name":     fault.Name,
		}

		if fault.InfraId != "" {
			faultBlock["infra_id"] = fault.InfraId
		}
		if fault.Revision != "" {
			faultBlock["revision"] = fault.Revision
		}
		faultBlock["is_enterprise"] = fault.IsEnterprise
		faultBlock["auth_enabled"] = fault.AuthEnabled

		// Values
		if len(fault.Values) > 0 {
			faultBlock["values"] = readValues(fault.Values)
		}

		result = append(result, faultBlock)
	}
	return result
}

func readProbes(probes []chaos.ExperimenttemplateProbe) []map[string]interface{} {
	result := make([]map[string]interface{}, 0, len(probes))
	for _, probe := range probes {
		probeBlock := map[string]interface{}{
			"identity": probe.Identity,
			"name":     probe.Name,
		}

		if probe.InfraId != "" {
			probeBlock["infra_id"] = probe.InfraId
		}
		if probe.Revision > 0 {
			probeBlock["revision"] = probe.Revision
		}
		probeBlock["is_enterprise"] = probe.IsEnterprise
		if probe.Duration != "" {
			probeBlock["duration"] = probe.Duration
		}
		if probe.Weightage > 0 {
			probeBlock["weightage"] = probe.Weightage
		}
		probeBlock["enable_data_collection"] = probe.EnableDataCollection

		// Conditions
		if len(probe.Conditions) > 0 {
			condList := make([]map[string]interface{}, 0, len(probe.Conditions))
			for _, cond := range probe.Conditions {
				condList = append(condList, map[string]interface{}{
					"execute_upon": cond.ExecuteUpon,
				})
			}
			probeBlock["conditions"] = condList
		}

		// Values
		if len(probe.Values) > 0 {
			probeBlock["values"] = readValues(probe.Values)
		}

		result = append(result, probeBlock)
	}
	return result
}

func readVertices(vertices []chaos.ExperimenttemplateVertex) []map[string]interface{} {
	result := make([]map[string]interface{}, 0, len(vertices))
	for _, vertex := range vertices {
		vertexBlock := map[string]interface{}{
			"name": vertex.Name,
		}

		// Start - ALWAYS include, even if empty (API behavior)
		if vertex.Start != nil {
			startBlock := map[string]interface{}{}

			if len(vertex.Start.Actions) > 0 {
				actionList := make([]map[string]interface{}, 0, len(vertex.Start.Actions))
				for _, action := range vertex.Start.Actions {
					actionList = append(actionList, map[string]interface{}{
						"name": action.Name,
					})
				}
				startBlock["actions"] = actionList
			}

			if len(vertex.Start.Faults) > 0 {
				faultList := make([]map[string]interface{}, 0, len(vertex.Start.Faults))
				for _, fault := range vertex.Start.Faults {
					faultList = append(faultList, map[string]interface{}{
						"name": fault.Name,
					})
				}
				startBlock["faults"] = faultList
			}

			if len(vertex.Start.Probes) > 0 {
				probeList := make([]map[string]interface{}, 0, len(vertex.Start.Probes))
				for _, probe := range vertex.Start.Probes {
					probeList = append(probeList, map[string]interface{}{
						"name": probe.Name,
					})
				}
				startBlock["probes"] = probeList
			}

			// Always include start block, even if empty
			vertexBlock["start"] = []map[string]interface{}{startBlock}
		}

		// End - ALWAYS include, even if empty (API behavior)
		if vertex.End != nil {
			endBlock := map[string]interface{}{}

			if len(vertex.End.Actions) > 0 {
				actionList := make([]map[string]interface{}, 0, len(vertex.End.Actions))
				for _, action := range vertex.End.Actions {
					actionList = append(actionList, map[string]interface{}{
						"name": action.Name,
					})
				}
				endBlock["actions"] = actionList
			}

			if len(vertex.End.Faults) > 0 {
				faultList := make([]map[string]interface{}, 0, len(vertex.End.Faults))
				for _, fault := range vertex.End.Faults {
					faultList = append(faultList, map[string]interface{}{
						"name": fault.Name,
					})
				}
				endBlock["faults"] = faultList
			}

			if len(vertex.End.Probes) > 0 {
				probeList := make([]map[string]interface{}, 0, len(vertex.End.Probes))
				for _, probe := range vertex.End.Probes {
					probeList = append(probeList, map[string]interface{}{
						"name": probe.Name,
					})
				}
				endBlock["probes"] = probeList
			}

			// Always include end block, even if empty
			vertexBlock["end"] = []map[string]interface{}{endBlock}
		}

		result = append(result, vertexBlock)
	}
	return result
}

func readValues(values []chaos.TemplateVariableMinimum) []map[string]interface{} {
	result := make([]map[string]interface{}, 0, len(values))
	for _, val := range values {
		valueBlock := map[string]interface{}{
			"name": val.Name,
		}

		// Handle value (which is *interface{})
		if val.Value != nil {
			// Convert interface{} to string for Terraform
			switch v := (*val.Value).(type) {
			case string:
				valueBlock["value"] = v
			case float64:
				valueBlock["value"] = fmt.Sprintf("%v", v)
			case int:
				valueBlock["value"] = fmt.Sprintf("%d", v)
			case bool:
				valueBlock["value"] = fmt.Sprintf("%t", v)
			default:
				// Try JSON encoding for complex types
				if jsonBytes, err := json.Marshal(v); err == nil {
					valueBlock["value"] = string(jsonBytes)
				} else {
					valueBlock["value"] = fmt.Sprintf("%v", v)
				}
			}
		}

		result = append(result, valueBlock)
	}
	return result
}
