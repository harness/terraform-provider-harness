package action_template

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/harness/harness-go-sdk/harness/chaos"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceActionTemplate() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for managing Harness Chaos Action Templates. Action templates define reusable actions that can be used in chaos experiments.",

		CreateContext: resourceActionTemplateCreate,
		ReadContext:   resourceActionTemplateRead,
		UpdateContext: resourceActionTemplateUpdate,
		DeleteContext: resourceActionTemplateDelete,

		Schema: resourceActionTemplateSchema(),

		Importer: &schema.ResourceImporter{
			StateContext: resourceActionTemplateImport,
		},
	}
}

func resourceActionTemplateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)

	accountID := c.AccountId
	if accountID == "" {
		return diag.Errorf("account ID must be configured in the provider")
	}

	orgID := d.Get("org_id").(string)
	projectID := d.Get("project_id").(string)
	hubIdentity := d.Get("hub_identity").(string)
	identity := d.Get("identity").(string)

	log.Printf("[DEBUG] Creating action template with identity: %s in hub: %s", identity, hubIdentity)

	// Build the request
	req := chaos.ChaosfaulttemplateActionTemplate{
		Identity: identity,
		Name:     d.Get("name").(string),
		HubRef:   hubIdentity,
	}

	// Optional fields
	if v, ok := d.GetOk("description"); ok {
		req.Description = v.(string)
	}

	if v, ok := d.GetOk("tags"); ok {
		tags := []string{}
		for _, tag := range v.([]interface{}) {
			tags = append(tags, tag.(string))
		}
		req.Tags = tags
	}

	if v, ok := d.GetOk("type"); ok {
		req.Type_ = v.(string)
	}

	if v, ok := d.GetOk("infrastructure_type"); ok {
		infraType := chaos.ActionsInfrastructureType(v.(string))
		req.InfrastructureType = &infraType
	}

	// Build action properties based on type
	if err := buildActionProperties(d, &req); err != nil {
		return diag.FromErr(err)
	}

	// Build run properties
	buildRunProperties(d, &req)

	// Build variables
	buildVariables(d, &req)

	// Create the action template
	resp, httpResp, err := c.DefaultApi.CreateActionTemplate(ctx, req, accountID, orgID, projectID)
	if err != nil {
		log.Printf("[ERROR] Failed to create action template: %v", err)
		if httpResp != nil {
			log.Printf("[ERROR] HTTP Response: %+v", httpResp)
		}
		return helpers.HandleChaosApiError(err, d, httpResp)
	}

	log.Printf("[DEBUG] Action template created successfully: %s", resp.Identity)

	// Set the ID
	d.SetId(generateID(accountID, orgID, projectID, hubIdentity, resp.Identity))

	// Read back the created resource
	return resourceActionTemplateRead(ctx, d, meta)
}

func resourceActionTemplateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)

	accountID := c.AccountId
	if accountID == "" {
		return diag.Errorf("account ID must be configured in the provider")
	}

	// Parse the ID
	accountID, orgID, projectID, hubIdentity, identity, err := parseID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	log.Printf("[DEBUG] Reading action template: %s from hub: %s", identity, hubIdentity)

	// Get the action template
	resp, httpResp, err := c.DefaultApi.GetActionTemplate(ctx, accountID, orgID, projectID, hubIdentity, identity, nil)
	if err != nil {
		if httpResp != nil && httpResp.StatusCode == 404 {
			log.Printf("[WARN] Action template not found, removing from state: %s", identity)
			d.SetId("")
			return nil
		}
		return helpers.HandleChaosReadApiError(err, d, httpResp)
	}

	if resp.Data == nil {
		log.Printf("[WARN] Action template data is nil, removing from state: %s", identity)
		d.SetId("")
		return nil
	}

	// Set the resource data
	return setActionTemplateData(d, resp.Data, accountID, orgID, projectID, hubIdentity)
}

func resourceActionTemplateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)

	accountID := c.AccountId
	if accountID == "" {
		return diag.Errorf("account ID must be configured in the provider")
	}

	// Parse the ID
	accountID, orgID, projectID, hubIdentity, identity, err := parseID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	log.Printf("[DEBUG] Updating action template: %s in hub: %s", identity, hubIdentity)

	// Build the update request
	req := chaos.ChaosfaulttemplateActionTemplate{
		Identity: identity,
		Name:     d.Get("name").(string),
		HubRef:   hubIdentity,
	}

	// Optional fields
	if v, ok := d.GetOk("description"); ok {
		req.Description = v.(string)
	}

	if v, ok := d.GetOk("tags"); ok {
		tags := []string{}
		for _, tag := range v.([]interface{}) {
			tags = append(tags, tag.(string))
		}
		req.Tags = tags
	}

	if v, ok := d.GetOk("type"); ok {
		req.Type_ = v.(string)
	}

	if v, ok := d.GetOk("infrastructure_type"); ok {
		infraType := chaos.ActionsInfrastructureType(v.(string))
		req.InfrastructureType = &infraType
	}

	// Build action properties based on type
	if err := buildActionProperties(d, &req); err != nil {
		return diag.FromErr(err)
	}

	// Build run properties
	buildRunProperties(d, &req)

	// Build variables
	buildVariables(d, &req)

	// Update the action template
	_, httpResp, err := c.DefaultApi.UpdateActionTemplate(ctx, req, accountID, orgID, projectID, identity, nil)
	if err != nil {
		log.Printf("[ERROR] Failed to update action template: %v", err)
		if httpResp != nil {
			log.Printf("[ERROR] HTTP Response: %+v", httpResp)
		}
		return helpers.HandleChaosApiError(err, d, httpResp)
	}

	log.Printf("[DEBUG] Action template updated successfully: %s", identity)

	// Read back the updated resource
	return resourceActionTemplateRead(ctx, d, meta)
}

func resourceActionTemplateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)

	accountID := c.AccountId
	if accountID == "" {
		return diag.Errorf("account ID must be configured in the provider")
	}

	// Parse the ID
	accountID, orgID, projectID, hubIdentity, identity, err := parseID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	log.Printf("[DEBUG] Deleting action template: %s from hub: %s", identity, hubIdentity)

	// Delete the action template
	_, httpResp, err := c.DefaultApi.DeleteActionTemplate(ctx, accountID, orgID, projectID, hubIdentity, identity, nil)
	if err != nil {
		if httpResp != nil && httpResp.StatusCode == 404 {
			log.Printf("[WARN] Action template not found, already deleted: %s", identity)
			d.SetId("")
			return nil
		}
		return helpers.HandleChaosApiError(err, d, httpResp)
	}

	log.Printf("[DEBUG] Action template deleted successfully: %s", identity)

	d.SetId("")
	return nil
}

func resourceActionTemplateImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	// Import ID formats:
	// 1. org_id/project_id/hub_identity/identity (project level)
	// 2. org_id/hub_identity/identity (org level)
	// 3. hub_identity/identity (account level)

	parts := strings.Split(d.Id(), "/")

	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)

	accountID := c.AccountId
	if accountID == "" {
		return nil, fmt.Errorf("account ID must be configured in the provider")
	}

	var orgID, projectID, hubIdentity, identity string

	switch len(parts) {
	case 4:
		// Project level: org_id/project_id/hub_identity/identity
		orgID = parts[0]
		projectID = parts[1]
		hubIdentity = parts[2]
		identity = parts[3]
	case 3:
		// Org level: org_id/hub_identity/identity
		orgID = parts[0]
		hubIdentity = parts[1]
		identity = parts[2]
	case 2:
		// Account level: hub_identity/identity
		hubIdentity = parts[0]
		identity = parts[1]
	default:
		return nil, fmt.Errorf("invalid import ID format. Expected: org_id/project_id/hub_identity/identity or org_id/hub_identity/identity or hub_identity/identity, got: %s", d.Id())
	}

	// Set the ID in the proper format
	d.SetId(generateID(accountID, orgID, projectID, hubIdentity, identity))

	// Set the scope fields
	if orgID != "" {
		d.Set("org_id", orgID)
	}
	if projectID != "" {
		d.Set("project_id", projectID)
	}
	d.Set("hub_identity", hubIdentity)
	d.Set("identity", identity)

	log.Printf("[DEBUG] Importing action template: %s from hub: %s (org: %s, project: %s)", identity, hubIdentity, orgID, projectID)

	// Read the resource
	diags := resourceActionTemplateRead(ctx, d, meta)
	if diags.HasError() {
		return nil, fmt.Errorf("failed to read action template during import: %v", diags)
	}

	return []*schema.ResourceData{d}, nil
}

// Helper functions

func generateID(accountID, orgID, projectID, hubIdentity, identity string) string {
	if projectID != "" {
		return fmt.Sprintf("%s/%s/%s/%s/%s", accountID, orgID, projectID, hubIdentity, identity)
	} else if orgID != "" {
		return fmt.Sprintf("%s/%s/%s/%s", accountID, orgID, hubIdentity, identity)
	}
	return fmt.Sprintf("%s/%s/%s", accountID, hubIdentity, identity)
}

func parseID(id string) (accountID, orgID, projectID, hubIdentity, identity string, err error) {
	parts := strings.Split(id, "/")

	switch len(parts) {
	case 5:
		// Project level: account_id/org_id/project_id/hub_identity/identity
		return parts[0], parts[1], parts[2], parts[3], parts[4], nil
	case 4:
		// Org level: account_id/org_id/hub_identity/identity
		return parts[0], parts[1], "", parts[2], parts[3], nil
	case 3:
		// Account level: account_id/hub_identity/identity
		return parts[0], "", "", parts[1], parts[2], nil
	default:
		return "", "", "", "", "", fmt.Errorf("invalid ID format: %s", id)
	}
}

func setActionTemplateData(d *schema.ResourceData, template *chaos.ChaosactiontemplateChaosActionTemplate, accountID, orgID, projectID, hubIdentity string) diag.Diagnostics {
	d.Set("account_id", template.AccountID)
	d.Set("identity", template.Identity)
	d.Set("name", template.Name)
	d.Set("hub_identity", hubIdentity)

	if orgID != "" {
		d.Set("org_id", orgID)
	}
	if projectID != "" {
		d.Set("project_id", projectID)
	}

	// Optional fields
	if template.Description != "" {
		d.Set("description", template.Description)
	}
	if len(template.Tags) > 0 {
		d.Set("tags", template.Tags)
	}
	if template.Type_ != nil {
		d.Set("type", string(*template.Type_))
	}
	if template.InfrastructureType != nil {
		d.Set("infrastructure_type", string(*template.InfrastructureType))
	}

	// Computed fields
	d.Set("id_internal", template.Id)
	d.Set("revision", template.Revision)
	d.Set("is_default", template.IsDefault)
	d.Set("is_enterprise", template.IsEnterprise)
	d.Set("is_removed", template.IsRemoved)
	d.Set("created_at", template.CreatedAt)
	d.Set("updated_at", template.UpdatedAt)
	d.Set("created_by", template.CreatedBy)
	d.Set("updated_by", template.UpdatedBy)
	d.Set("template", template.Template)

	// Parse ActionProperties
	if template.ActionProperties != nil {
		if err := setActionPropertiesData(d, template.ActionProperties); err != nil {
			return diag.FromErr(err)
		}
	}

	// Parse RunProperties
	if template.RunProperties != nil {
		if err := setRunPropertiesData(d, template.RunProperties); err != nil {
			return diag.FromErr(err)
		}
	}

	// Parse Variables
	if len(template.Variables) > 0 {
		if err := setVariablesData(d, template.Variables); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

// Helper functions for action properties

func buildActionProperties(d *schema.ResourceData, req *chaos.ChaosfaulttemplateActionTemplate) error {
	actionType := d.Get("type").(string)

	switch actionType {
	case "delay":
		if v, ok := d.GetOk("delay_action"); ok && len(v.([]interface{})) > 0 {
			delayConfig := v.([]interface{})[0].(map[string]interface{})
			req.ActionProperties = &chaos.ActionActionTemplateProperties{
				DelayAction: &chaos.ActionDelayActionTemplate{
					Duration: delayConfig["duration"].(string),
				},
			}
		} else {
			return fmt.Errorf("delay_action block is required when type is 'delay'")
		}

	case "customScript":
		if v, ok := d.GetOk("custom_script_action"); ok && len(v.([]interface{})) > 0 {
			scriptConfig := v.([]interface{})[0].(map[string]interface{})

			customScript := &chaos.ActionCustomScriptActionTemplate{
				Command: scriptConfig["command"].(string),
			}

			// Optional args
			if args, ok := scriptConfig["args"].([]interface{}); ok && len(args) > 0 {
				argsList := make([]string, len(args))
				for i, arg := range args {
					argsList[i] = arg.(string)
				}
				customScript.Args = argsList
			}

			// Optional env vars
			if envList, ok := scriptConfig["env"].([]interface{}); ok && len(envList) > 0 {
				envVars := make([]chaos.ActionEnv, len(envList))
				for i, e := range envList {
					envMap := e.(map[string]interface{})
					envVars[i] = chaos.ActionEnv{
						Name:  envMap["name"].(string),
						Value: envMap["value"].(string),
					}
				}
				customScript.Env = envVars
			}

			req.ActionProperties = &chaos.ActionActionTemplateProperties{
				CustomScriptAction: customScript,
			}
		} else {
			return fmt.Errorf("custom_script_action block is required when type is 'customScript'")
		}

	case "container":
		if v, ok := d.GetOk("container_action"); ok && len(v.([]interface{})) > 0 {
			containerConfig := v.([]interface{})[0].(map[string]interface{})

			container := &chaos.CommonContainerTemplate{
				Image: containerConfig["image"].(string),
			}

			// Optional command
			if cmd, ok := containerConfig["command"].([]interface{}); ok && len(cmd) > 0 {
				cmdList := make([]string, len(cmd))
				for i, c := range cmd {
					cmdList[i] = c.(string)
				}
				container.Command = cmdList
			}

			// Optional args
			if args, ok := containerConfig["args"].(string); ok && args != "" {
				container.Args = args
			}

			// Optional namespace
			if ns, ok := containerConfig["namespace"].(string); ok && ns != "" {
				container.Namespace = ns
			}

			// Optional labels
			if labels, ok := containerConfig["labels"].(map[string]interface{}); ok && len(labels) > 0 {
				labelMap := make(map[string]string)
				for k, v := range labels {
					labelMap[k] = v.(string)
				}
				container.Labels = labelMap
			}

			// Optional annotations
			if annotations, ok := containerConfig["annotations"].(map[string]interface{}); ok && len(annotations) > 0 {
				annotationMap := make(map[string]string)
				for k, v := range annotations {
					annotationMap[k] = v.(string)
				}
				container.Annotations = annotationMap
			}

			// Optional image pull policy
			if policy, ok := containerConfig["image_pull_policy"].(string); ok && policy != "" {
				pullPolicy := chaos.V1PullPolicy(policy)
				container.ImagePullPolicy = &pullPolicy
			}

			// Optional image pull secrets
			if secrets, ok := containerConfig["image_pull_secrets"].([]interface{}); ok && len(secrets) > 0 {
				secretsList := make([]string, len(secrets))
				for i, s := range secrets {
					secretsList[i] = s.(string)
				}
				container.ImagePullSecrets = secretsList
			}

			// Optional environment variables
			if envList, ok := containerConfig["env"].([]interface{}); ok && len(envList) > 0 {
				envVars := make([]chaos.K8sIoApiCoreV1EnvVar, len(envList))
				for i, e := range envList {
					envMap := e.(map[string]interface{})
					envVar := chaos.K8sIoApiCoreV1EnvVar{
						Name: envMap["name"].(string),
					}
					if val, ok := envMap["value"].(string); ok {
						envVar.Value = val
					}
					envVars[i] = envVar
				}
				container.Env = envVars
			}

			// Optional resources
			if resourcesList, ok := containerConfig["resources"].([]interface{}); ok && len(resourcesList) > 0 {
				resourcesConfig := resourcesList[0].(map[string]interface{})
				resources := &chaos.CommonResourceRequirements{}

				if limits, ok := resourcesConfig["limits"].(map[string]interface{}); ok && len(limits) > 0 {
					limitsMap := make(map[string]string)
					for k, v := range limits {
						limitsMap[k] = v.(string)
					}
					resources.Limits = limitsMap
				}

				if requests, ok := resourcesConfig["requests"].(map[string]interface{}); ok && len(requests) > 0 {
					requestsMap := make(map[string]string)
					for k, v := range requests {
						requestsMap[k] = v.(string)
					}
					resources.Requests = requestsMap
				}

				container.Resources = resources
			}

			// Optional service account name
			if sa, ok := containerConfig["service_account_name"].(string); ok && sa != "" {
				container.ServiceAccountName = sa
			}

			// Optional node selector
			if nodeSelector, ok := containerConfig["node_selector"].(map[string]interface{}); ok && len(nodeSelector) > 0 {
				nodeSelectorMap := make(map[string]string)
				for k, v := range nodeSelector {
					nodeSelectorMap[k] = v.(string)
				}
				container.NodeSelector = nodeSelectorMap
			}

			// Optional host network
			if hostNet, ok := containerConfig["host_network"].(bool); ok {
				container.HostNetwork = hostNet
			}

			// Optional host PID
			if hostPID, ok := containerConfig["host_pid"].(bool); ok {
				container.HostPID = hostPID
			}

			// Optional host IPC
			if hostIPC, ok := containerConfig["host_ipc"].(bool); ok {
				container.HostIPC = hostIPC
			}

			// Optional tolerations
			if tolerationsList, ok := containerConfig["tolerations"].([]interface{}); ok && len(tolerationsList) > 0 {
				tolerations := make([]chaos.V1Toleration, len(tolerationsList))
				for i, t := range tolerationsList {
					tolMap := t.(map[string]interface{})
					toleration := chaos.V1Toleration{}

					if key, ok := tolMap["key"].(string); ok && key != "" {
						toleration.Key = key
					}

					if operator, ok := tolMap["operator"].(string); ok && operator != "" {
						operatorPtr := chaos.AllOfv1TolerationOperator(operator)
						toleration.Operator = &operatorPtr
					}

					if value, ok := tolMap["value"].(string); ok && value != "" {
						toleration.Value = value
					}

					if effect, ok := tolMap["effect"].(string); ok && effect != "" {
						effectPtr := chaos.AllOfv1TolerationEffect(effect)
						toleration.Effect = &effectPtr
					}

					if seconds, ok := tolMap["toleration_seconds"].(int); ok && seconds > 0 {
						seconds32 := int32(seconds)
						toleration.TolerationSeconds = seconds32
					}

					tolerations[i] = toleration
				}
				container.Tolerations = tolerations
			}

			// Optional volume mounts
			if volumeMountsList, ok := containerConfig["volume_mounts"].([]interface{}); ok && len(volumeMountsList) > 0 {
				volumeMounts := make([]chaos.V1VolumeMount, len(volumeMountsList))
				for i, vm := range volumeMountsList {
					vmMap := vm.(map[string]interface{})
					volumeMount := chaos.V1VolumeMount{
						Name:      vmMap["name"].(string),
						MountPath: vmMap["mount_path"].(string),
					}

					if subPath, ok := vmMap["sub_path"].(string); ok && subPath != "" {
						volumeMount.SubPath = subPath
					}

					if readOnly, ok := vmMap["read_only"].(bool); ok {
						volumeMount.ReadOnly = readOnly
					}

					volumeMounts[i] = volumeMount
				}
				container.VolumeMounts = volumeMounts
			}

			// Optional volumes
			if volumesList, ok := containerConfig["volumes"].([]interface{}); ok && len(volumesList) > 0 {
				volumes := make([]chaos.CommonVolume, len(volumesList))
				for i, v := range volumesList {
					volMap := v.(map[string]interface{})
					volume := chaos.CommonVolume{
						Name: volMap["name"].(string),
					}

					// EmptyDir volume
					if emptyDirList, ok := volMap["empty_dir"].([]interface{}); ok && len(emptyDirList) > 0 {
						emptyDir := &chaos.V1EmptyDirVolumeSource{}
						
						// Handle case where empty_dir {} is empty (nil config)
						if emptyDirList[0] != nil {
							if emptyDirConfig, ok := emptyDirList[0].(map[string]interface{}); ok {
								if medium, ok := emptyDirConfig["medium"].(string); ok && medium != "" {
									mediumPtr := chaos.AllOfv1EmptyDirVolumeSourceMedium(medium)
									emptyDir.Medium = &mediumPtr
								}

								if sizeLimit, ok := emptyDirConfig["size_limit"].(string); ok && sizeLimit != "" {
									emptyDir.SizeLimit = &sizeLimit
								}
							}
						}

						volume.EmptyDir = emptyDir
						volumeType := chaos.EMPTY_DIR_CommonVolumeType
						volume.Type_ = (*chaos.CommonVolumeType)(&volumeType)
					}

					// ConfigMap volume
					if configMapList, ok := volMap["config_map"].([]interface{}); ok && len(configMapList) > 0 && configMapList[0] != nil {
						if configMapConfig, ok := configMapList[0].(map[string]interface{}); ok {
							configMap := &chaos.CommonVolumeInputTemplate{}
							if name, ok := configMapConfig["name"].(string); ok {
								configMap.Name = name
							}
							if items, ok := configMapConfig["items"].(map[string]string); ok {
								configMap.Items = items
							}
							if defaultMode, ok := configMapConfig["default_mode"].(string); ok {
								configMap.DefaultMode = defaultMode
							}
							if optional, ok := configMapConfig["optional"].(bool); ok {
								configMap.Optional = optional
							}
							volume.ConfigMap = configMap
							volumeType := chaos.CONFIG_MAP_CommonVolumeType
							volume.Type_ = (*chaos.CommonVolumeType)(&volumeType)
						}
					}

					// Secret volume
					if secretList, ok := volMap["secret"].([]interface{}); ok && len(secretList) > 0 && secretList[0] != nil {
						if secretConfig, ok := secretList[0].(map[string]interface{}); ok {
							secret := &chaos.CommonVolumeInputTemplate{}
							if name, ok := secretConfig["secret_name"].(string); ok {
								secret.Name = name
							}
							if items, ok := secretConfig["items"].(map[string]string); ok {
								secret.Items = items
							}
							if defaultMode, ok := secretConfig["default_mode"].(string); ok {
								secret.DefaultMode = defaultMode
							}
							if optional, ok := secretConfig["optional"].(bool); ok {
								secret.Optional = optional
							}
							volume.Secret = secret
							volumeType := chaos.SECRET_CommonVolumeType
							volume.Type_ = (*chaos.CommonVolumeType)(&volumeType)
						}
					}

					// HostPath volume
					if hostPathList, ok := volMap["host_path"].([]interface{}); ok && len(hostPathList) > 0 && hostPathList[0] != nil {
						if hostPathConfig, ok := hostPathList[0].(map[string]interface{}); ok {
							hostPath := &chaos.V1HostPathVolumeSource{}
							if path, ok := hostPathConfig["path"].(string); ok {
								hostPath.Path = path
							}
							if type_, ok := hostPathConfig["type"].(string); ok && type_ != "" {
								typePtr := chaos.AllOfv1HostPathVolumeSourceType_(type_)
								hostPath.Type_ = &typePtr
							}
							volume.HostPath = hostPath
							volumeType := chaos.HOST_PATH_CommonVolumeType
							volume.Type_ = (*chaos.CommonVolumeType)(&volumeType)
						}
					}

					// PersistentVolumeClaim volume
					if pvcList, ok := volMap["persistent_volume_claim"].([]interface{}); ok && len(pvcList) > 0 && pvcList[0] != nil {
						if pvcConfig, ok := pvcList[0].(map[string]interface{}); ok {
							pvc := &chaos.V1PersistentVolumeClaimVolumeSource{}
							if claimName, ok := pvcConfig["claim_name"].(string); ok {
								pvc.ClaimName = claimName
							}

							if readOnly, ok := pvcConfig["read_only"].(bool); ok {
								pvc.ReadOnly = readOnly
							}

							volume.PersistentVolumeClaim = pvc
							volumeType := chaos.PERSISTENT_VOLUME_CLAIM_CommonVolumeType
							volume.Type_ = (*chaos.CommonVolumeType)(&volumeType)
						}
					}

					volumes[i] = volume
				}
				container.Volumes = volumes
			}

			req.ActionProperties = &chaos.ActionActionTemplateProperties{
				ContainerAction: container,
			}
		} else {
			return fmt.Errorf("container_action block is required when type is 'container'")
		}

	default:
		return fmt.Errorf("unsupported action type: %s", actionType)
	}

	return nil
}

func buildRunProperties(d *schema.ResourceData, req *chaos.ChaosfaulttemplateActionTemplate) {
	if v, ok := d.GetOk("run_properties"); ok && len(v.([]interface{})) > 0 {
		runPropsConfig := v.([]interface{})[0].(map[string]interface{})

		runProps := &chaos.ActionActionTemplateRunProperties{}

		if initialDelay, ok := runPropsConfig["initial_delay"].(string); ok && initialDelay != "" {
			runProps.InitialDelay = initialDelay
		}

		if interval, ok := runPropsConfig["interval"].(string); ok && interval != "" {
			runProps.Interval = interval
		}

		if maxRetries, ok := runPropsConfig["max_retries"].(int); ok && maxRetries > 0 {
			var retries interface{} = maxRetries
			runProps.MaxRetries = &retries
		}

		if stopOnFailure, ok := runPropsConfig["stop_on_failure"].(bool); ok {
			runProps.StopOnFailure = stopOnFailure
		}

		if timeout, ok := runPropsConfig["timeout"].(string); ok && timeout != "" {
			runProps.Timeout = timeout
		}

		if verbosity, ok := runPropsConfig["verbosity"].(string); ok && verbosity != "" {
			runProps.Verbosity = verbosity
		}

		req.RunProperties = runProps
	}
}

func buildVariables(d *schema.ResourceData, req *chaos.ChaosfaulttemplateActionTemplate) {
	if v, ok := d.GetOk("variables"); ok {
		variablesList := v.([]interface{})
		if len(variablesList) > 0 {
			variables := make([]chaos.TemplateVariable, len(variablesList))

			for i, varItem := range variablesList {
				varMap := varItem.(map[string]interface{})

				var value interface{} = varMap["value"].(string)
				variable := chaos.TemplateVariable{
					Name:  varMap["name"].(string),
					Value: &value,
				}

				if desc, ok := varMap["description"].(string); ok && desc != "" {
					variable.Description = desc
				}

				if varType, ok := varMap["type"].(string); ok && varType != "" {
					vType := chaos.TemplateVariableType(varType)
					variable.Type_ = &vType
				}

				if required, ok := varMap["required"].(bool); ok {
					variable.Required = required
				}

				variables[i] = variable
			}

			req.Variables = variables
		}
	}
}

// setActionPropertiesData populates action-specific blocks from API response
func setActionPropertiesData(d *schema.ResourceData, props *chaos.AllOfchaosactiontemplateChaosActionTemplateActionProperties) error {
	if props == nil {
		return nil
	}

	// Delay Action
	if props.DelayAction != nil {
		delayBlock := []map[string]interface{}{
			{
				"duration": props.DelayAction.Duration,
			},
		}
		if err := d.Set("delay_action", delayBlock); err != nil {
			return fmt.Errorf("error setting delay_action: %v", err)
		}
	}

	// Custom Script Action
	if props.CustomScriptAction != nil {
		scriptBlock := map[string]interface{}{
			"command": props.CustomScriptAction.Command,
		}

		if len(props.CustomScriptAction.Args) > 0 {
			scriptBlock["args"] = props.CustomScriptAction.Args
		}

		if len(props.CustomScriptAction.Env) > 0 {
			envList := make([]map[string]interface{}, len(props.CustomScriptAction.Env))
			for i, env := range props.CustomScriptAction.Env {
				envList[i] = map[string]interface{}{
					"name":  env.Name,
					"value": env.Value,
				}
			}
			scriptBlock["env"] = envList
		}

		if err := d.Set("custom_script_action", []map[string]interface{}{scriptBlock}); err != nil {
			return fmt.Errorf("error setting custom_script_action: %v", err)
		}
	}

	// Container Action
	if props.ContainerAction != nil {
		containerBlock := map[string]interface{}{
			"image": props.ContainerAction.Image,
		}

		if len(props.ContainerAction.Command) > 0 {
			containerBlock["command"] = props.ContainerAction.Command
		}
		if props.ContainerAction.Args != "" {
			containerBlock["args"] = props.ContainerAction.Args
		}
		if props.ContainerAction.Namespace != "" {
			containerBlock["namespace"] = props.ContainerAction.Namespace
		}
		if len(props.ContainerAction.Labels) > 0 {
			containerBlock["labels"] = props.ContainerAction.Labels
		}
		if len(props.ContainerAction.Annotations) > 0 {
			containerBlock["annotations"] = props.ContainerAction.Annotations
		}
		if props.ContainerAction.ImagePullPolicy != nil {
			containerBlock["image_pull_policy"] = string(*props.ContainerAction.ImagePullPolicy)
		}
		if len(props.ContainerAction.ImagePullSecrets) > 0 {
			containerBlock["image_pull_secrets"] = props.ContainerAction.ImagePullSecrets
		}
		if props.ContainerAction.ServiceAccountName != "" {
			containerBlock["service_account_name"] = props.ContainerAction.ServiceAccountName
		}
		if len(props.ContainerAction.NodeSelector) > 0 {
			containerBlock["node_selector"] = props.ContainerAction.NodeSelector
		}
		containerBlock["host_network"] = props.ContainerAction.HostNetwork
		containerBlock["host_pid"] = props.ContainerAction.HostPID
		containerBlock["host_ipc"] = props.ContainerAction.HostIPC

		// Environment variables
		if len(props.ContainerAction.Env) > 0 {
			envList := make([]map[string]interface{}, len(props.ContainerAction.Env))
			for i, env := range props.ContainerAction.Env {
				envList[i] = map[string]interface{}{
					"name":  env.Name,
					"value": env.Value,
				}
			}
			containerBlock["env"] = envList
		}

		// Resources
		if props.ContainerAction.Resources != nil {
			resourcesBlock := map[string]interface{}{}
			if len(props.ContainerAction.Resources.Limits) > 0 {
				resourcesBlock["limits"] = props.ContainerAction.Resources.Limits
			}
			if len(props.ContainerAction.Resources.Requests) > 0 {
				resourcesBlock["requests"] = props.ContainerAction.Resources.Requests
			}
			if len(resourcesBlock) > 0 {
				containerBlock["resources"] = []map[string]interface{}{resourcesBlock}
			}
		}

		// Tolerations
		if len(props.ContainerAction.Tolerations) > 0 {
			tolerationsList := make([]map[string]interface{}, len(props.ContainerAction.Tolerations))
			for i, tol := range props.ContainerAction.Tolerations {
				tolMap := map[string]interface{}{}
				if tol.Key != "" {
					tolMap["key"] = tol.Key
				}
				if tol.Operator != nil {
					tolMap["operator"] = string(*tol.Operator)
				}
				if tol.Value != "" {
					tolMap["value"] = tol.Value
				}
				if tol.Effect != nil {
					tolMap["effect"] = string(*tol.Effect)
				}
				if tol.TolerationSeconds != 0 {
					tolMap["toleration_seconds"] = tol.TolerationSeconds
				}
				tolerationsList[i] = tolMap
			}
			containerBlock["tolerations"] = tolerationsList
		}

		// Volume Mounts
		if len(props.ContainerAction.VolumeMounts) > 0 {
			volumeMountsList := make([]map[string]interface{}, len(props.ContainerAction.VolumeMounts))
			for i, vm := range props.ContainerAction.VolumeMounts {
				vmMap := map[string]interface{}{
					"name":       vm.Name,
					"mount_path": vm.MountPath,
					"read_only":  vm.ReadOnly,
				}
				if vm.SubPath != "" {
					vmMap["sub_path"] = vm.SubPath
				}
				volumeMountsList[i] = vmMap
			}
			containerBlock["volume_mounts"] = volumeMountsList
		}

		// Volumes
		if len(props.ContainerAction.Volumes) > 0 {
			volumesList := make([]map[string]interface{}, len(props.ContainerAction.Volumes))
			for i, vol := range props.ContainerAction.Volumes {
				volMap := map[string]interface{}{
					"name": vol.Name,
				}

				// EmptyDir
				if vol.EmptyDir != nil {
					emptyDirMap := map[string]interface{}{}
					if vol.EmptyDir.Medium != nil {
						emptyDirMap["medium"] = string(*vol.EmptyDir.Medium)
					}
					if vol.EmptyDir.SizeLimit != nil && *vol.EmptyDir.SizeLimit != "0" {
						// Only set size_limit if not "0" (API returns "0" for unlimited)
						emptyDirMap["size_limit"] = *vol.EmptyDir.SizeLimit
					}
					volMap["empty_dir"] = []map[string]interface{}{emptyDirMap}
				}

				// ConfigMap
				if vol.ConfigMap != nil {
					configMapMap := map[string]interface{}{
						"name":     vol.ConfigMap.Name,
						"optional": vol.ConfigMap.Optional,
					}
					if len(vol.ConfigMap.Items) > 0 {
						configMapMap["items"] = vol.ConfigMap.Items
					}
					if vol.ConfigMap.DefaultMode != "" {
						configMapMap["default_mode"] = vol.ConfigMap.DefaultMode
					}
					volMap["config_map"] = []map[string]interface{}{configMapMap}
				}

				// Secret
				if vol.Secret != nil {
					secretMap := map[string]interface{}{
						"secret_name": vol.Secret.Name,
						"optional":    vol.Secret.Optional,
					}
					if len(vol.Secret.Items) > 0 {
						secretMap["items"] = vol.Secret.Items
					}
					if vol.Secret.DefaultMode != "" {
						secretMap["default_mode"] = vol.Secret.DefaultMode
					}
					volMap["secret"] = []map[string]interface{}{secretMap}
				}

				// HostPath
				if vol.HostPath != nil {
					hostPathMap := map[string]interface{}{
						"path": vol.HostPath.Path,
					}
					if vol.HostPath.Type_ != nil {
						hostPathMap["type"] = string(*vol.HostPath.Type_)
					}
					volMap["host_path"] = []map[string]interface{}{hostPathMap}
				}

				// PersistentVolumeClaim
				if vol.PersistentVolumeClaim != nil {
					pvcMap := map[string]interface{}{
						"claim_name": vol.PersistentVolumeClaim.ClaimName,
						"read_only":  vol.PersistentVolumeClaim.ReadOnly,
					}
					volMap["persistent_volume_claim"] = []map[string]interface{}{pvcMap}
				}

				volumesList[i] = volMap
			}
			containerBlock["volumes"] = volumesList
		}

		if err := d.Set("container_action", []map[string]interface{}{containerBlock}); err != nil {
			return fmt.Errorf("error setting container_action: %v", err)
		}
	}

	return nil
}

// setRunPropertiesData populates run_properties block from API response
func setRunPropertiesData(d *schema.ResourceData, props *chaos.ActionActionTemplateRunProperties) error {
	if props == nil {
		return nil
	}

	runPropsBlock := map[string]interface{}{}

	// String fields - only set if non-empty
	if props.InitialDelay != "" {
		runPropsBlock["initial_delay"] = props.InitialDelay
	}
	if props.Interval != "" {
		runPropsBlock["interval"] = props.Interval
	}
	if props.Timeout != "" {
		runPropsBlock["timeout"] = props.Timeout
	}
	if props.Verbosity != "" {
		runPropsBlock["verbosity"] = props.Verbosity
	}
	
	// Boolean field - only set if true (non-default)
	// Following Terraform best practice: don't set default values to avoid drift
	if props.StopOnFailure {
		runPropsBlock["stop_on_failure"] = true
	}
	
	// Integer field - only set if non-zero (non-default)
	if props.MaxRetries != nil {
		retriesVal := getIntFromInterface(props.MaxRetries)
		if retriesVal > 0 {
			runPropsBlock["max_retries"] = retriesVal
		}
	}

	if len(runPropsBlock) > 0 {
		if err := d.Set("run_properties", []map[string]interface{}{runPropsBlock}); err != nil {
			return fmt.Errorf("error setting run_properties: %v", err)
		}
	}

	return nil
}

// setVariablesData populates variables block from API response
func setVariablesData(d *schema.ResourceData, vars []chaos.TemplateVariable) error {
	if len(vars) == 0 {
		return nil
	}

	varsList := make([]map[string]interface{}, len(vars))
	for i, v := range vars {
		varMap := map[string]interface{}{
			"name":  v.Name,
			"value": v.Value,
		}
		if v.Description != "" {
			varMap["description"] = v.Description
		}
		if v.Type_ != nil {
			// Normalize type to lowercase to match Terraform schema
			varMap["type"] = strings.ToLower(string(*v.Type_))
		}
		// Only set required if true (non-default)
		// Following Terraform best practice: don't set default values to avoid drift
		if v.Required {
			varMap["required"] = true
		}
		varsList[i] = varMap
	}

	if err := d.Set("variables", varsList); err != nil {
			return fmt.Errorf("error setting variables: %v", err)
		}

	return nil
}

// getIntFromInterface safely extracts int value from interface{} pointer
// Returns 0 if nil or cannot convert
func getIntFromInterface(val *interface{}) int {
	if val == nil {
		return 0
	}
	
	switch v := (*val).(type) {
	case int:
		return v
	case int32:
		return int(v)
	case int64:
		return int(v)
	case float64:
		return int(v)
	case float32:
		return int(v)
	default:
		return 0
	}
}
