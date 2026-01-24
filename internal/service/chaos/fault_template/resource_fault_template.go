package fault_template

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/harness/harness-go-sdk/harness/chaos"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"gopkg.in/yaml.v3"
)

// ResourceFaultTemplate returns the fault template resource
func ResourceFaultTemplate() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for managing Harness Chaos Fault Templates. " +
			"Phase 1: Core fields (identity, name, description, tags, category, infrastructure, basic spec)",

		CreateContext: resourceFaultTemplateCreate,
		ReadContext:   resourceFaultTemplateRead,
		UpdateContext: resourceFaultTemplateUpdate,
		DeleteContext: resourceFaultTemplateDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceFaultTemplateImport,
		},

		Schema: ResourceFaultTemplateSchema(),
	}
}

func resourceFaultTemplateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)

	// Get scope identifiers
	accountID := c.AccountId
	orgID := d.Get("org_id").(string)
	projectID := d.Get("project_id").(string)
	hubIdentity := d.Get("hub_identity").(string)

	// Build request with required fields
	// NOTE: Based on real Harness YAML, apiVersion and kind should be EMPTY strings
	revision := "v1" // Default revision
	if v, ok := d.GetOk("revision"); ok && v.(string) != "" {
		revision = v.(string)
	}

	req := chaos.ChaosfaulttemplateCreateFaultTemplateRequest{
		Identity:   d.Get("identity").(string),
		Name:       d.Get("name").(string),
		ApiVersion: "",       // Empty per Harness YAML format
		Kind:       "",       // Empty per Harness YAML format
		Revision:   revision, // Use revision from config or default "v1"
	}

	// Set description if provided
	if v, ok := d.GetOk("description"); ok {
		req.Description = v.(string)
	}

	// Build Phase 1 fields
	if err := buildFaultTemplateRequest(d, &req); err != nil {
		return diag.FromErr(err)
	}

	// Ensure spec is initialized (API might require it)
	if req.Spec == nil {
		req.Spec = &chaos.FaulttemplateSpec{}
	}

	// Log the request for debugging
	log.Printf("[DEBUG] Creating fault template: identity=%s, name=%s, hubIdentity=%s", req.Identity, req.Name, hubIdentity)
	log.Printf("[DEBUG] Request spec: %+v", req.Spec)

	// Log full request as JSON for debugging
	if reqJSON, err := json.MarshalIndent(req, "", "  "); err == nil {
		log.Printf("[DEBUG] Full request JSON:\n%s", string(reqJSON))
	} else {
		log.Printf("[DEBUG] Failed to marshal request to JSON: %v", err)
	}

	// Create fault template
	resp, httpResp, err := c.FaulttemplateApi.CreateFaultTemplate(ctx, req, accountID, orgID, projectID, hubIdentity, nil)
	if err != nil {
		log.Printf("[ERROR] Failed to create fault template: %v", err)
		if httpResp != nil {
			log.Printf("[ERROR] HTTP Status: %d", httpResp.StatusCode)
		}
		return helpers.HandleChaosApiError(err, d, httpResp)
	}

	// Set ID: org_id/project_id/hub_identity/identity
	d.SetId(fmt.Sprintf("%s/%s/%s/%s", orgID, projectID, hubIdentity, resp.Identity))

	return resourceFaultTemplateRead(ctx, d, meta)
}

func resourceFaultTemplateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)

	// Parse ID
	parts := strings.Split(d.Id(), "/")
	if len(parts) != 4 {
		return diag.Errorf("invalid fault template ID format: %s (expected: org_id/project_id/hub_identity/identity)", d.Id())
	}

	orgID := parts[0]
	projectID := parts[1]
	hubIdentity := parts[2]
	identity := parts[3]

	// Get revision from state, default to "v1" if not set
	revision := "v1"
	if v, ok := d.GetOk("revision"); ok && v.(string) != "" {
		revision = v.(string)
	}

	// Get fault template with the revision from state
	resp, httpResp, err := c.FaulttemplateApi.GetFaultTemplate(ctx, c.AccountId, orgID, projectID, hubIdentity, revision, identity, nil)
	if err != nil {
		return helpers.HandleChaosReadApiError(err, d, httpResp)
	}

	// Set Phase 1 fields from response
	if err := setFaultTemplateData(d, resp.Data); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceFaultTemplateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)

	// Parse ID
	parts := strings.Split(d.Id(), "/")
	if len(parts) != 4 {
		return diag.Errorf("invalid fault template ID format: %s", d.Id())
	}

	orgID := parts[0]
	projectID := parts[1]
	hubIdentity := parts[2]
	identity := parts[3]

	// Build update request
	revision := "v1" // Default revision
	if v, ok := d.GetOk("revision"); ok && v.(string) != "" {
		revision = v.(string)
	}

	req := chaos.ChaosfaulttemplateCreateFaultTemplateRequest{
		Identity:   identity,
		Name:       d.Get("name").(string),
		ApiVersion: "",       // Empty per Harness YAML format
		Kind:       "",       // Empty per Harness YAML format
		Revision:   revision, // Use revision from config or default "v1"
	}

	if err := buildFaultTemplateRequest(d, &req); err != nil {
		return diag.FromErr(err)
	}

	// Update fault template - using CreateFaultTemplate with same identity updates it
	_, httpResp, err := c.FaulttemplateApi.CreateFaultTemplate(ctx, req, c.AccountId, orgID, projectID, hubIdentity, nil)
	if err != nil {
		return helpers.HandleChaosApiError(err, d, httpResp)
	}

	return resourceFaultTemplateRead(ctx, d, meta)
}

func resourceFaultTemplateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)

	// Parse ID
	parts := strings.Split(d.Id(), "/")
	if len(parts) != 4 {
		return diag.Errorf("invalid fault template ID format: %s", d.Id())
	}

	orgID := parts[0]
	projectID := parts[1]
	hubIdentity := parts[2]
	identity := parts[3]

	// Delete fault template
	_, httpResp, err := c.FaulttemplateApi.DeleteFaultTemplate(ctx, c.AccountId, orgID, projectID, hubIdentity, identity, nil)
	if err != nil {
		return helpers.HandleChaosApiError(err, d, httpResp)
	}

	d.SetId("")
	return nil
}

func resourceFaultTemplateImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	// Import format: org_id/project_id/hub_identity/identity
	parts := strings.Split(d.Id(), "/")
	if len(parts) != 4 {
		return nil, fmt.Errorf("invalid import ID format: %s (expected: org_id/project_id/hub_identity/identity)", d.Id())
	}

	d.Set("org_id", parts[0])
	d.Set("project_id", parts[1])
	d.Set("hub_identity", parts[2])
	d.Set("identity", parts[3])

	return []*schema.ResourceData{d}, nil
}

func buildFaultTemplateRequest(d *schema.ResourceData, req *chaos.ChaosfaulttemplateCreateFaultTemplateRequest) error {
	// Basic fields
	if v, ok := d.GetOk("description"); ok {
		req.Description = v.(string)
	}

	// Tags
	if v, ok := d.GetOk("tags"); ok {
		tags := v.([]interface{})
		req.Tags = make([]string, len(tags))
		for i, tag := range tags {
			req.Tags[i] = tag.(string)
		}
	}

	// Category
	if v, ok := d.GetOk("category"); ok {
		categories := v.([]interface{})
		req.Category = make([]string, len(categories))
		for i, cat := range categories {
			req.Category[i] = cat.(string)
		}
	}

	// Infrastructure
	if v, ok := d.GetOk("infrastructures"); ok {
		infras := v.([]interface{})
		req.Infras = make([]string, len(infras))
		for i, infra := range infras {
			req.Infras[i] = infra.(string)
		}
	}

	// Type and flags
	if v, ok := d.GetOk("type"); ok {
		req.Type_ = v.(string)
	}

	if v, ok := d.GetOk("permissions_required"); ok {
		req.PermissionsRequired = v.(string)
	}

	// Links
	if v, ok := d.GetOk("links"); ok {
		linksData := v.([]interface{})
		req.Links = make([]chaos.FaulttemplateLink, len(linksData))
		for i, linkItem := range linksData {
			linkMap := linkItem.(map[string]interface{})
			link := chaos.FaulttemplateLink{}
			if name, ok := linkMap["name"].(string); ok {
				link.Name = name
			}
			if url, ok := linkMap["url"].(string); ok {
				link.Url = url
			}
			req.Links[i] = link
		}
	}

	// Variables
	if err := buildVariables(d, req); err != nil {
		return err
	}

	// Build complete spec (all phases)
	if err := buildSpecComplete(d, req); err != nil {
		return err
	}

	return nil
}

func buildVariables(d *schema.ResourceData, req *chaos.ChaosfaulttemplateCreateFaultTemplateRequest) error {
	if v, ok := d.GetOk("variables"); ok {
		varsList := v.([]interface{})
		if len(varsList) > 0 {
			req.Variables = make([]chaos.TemplateVariable, len(varsList))
			for i, varItem := range varsList {
				varConfig := varItem.(map[string]interface{})

				name := varConfig["name"].(string)
				value := varConfig["value"].(string)
				var valueInterface interface{} = value

				variable := chaos.TemplateVariable{
					Name:  name,
					Value: &valueInterface,
				}

				if desc, ok := varConfig["description"].(string); ok && desc != "" {
					variable.Description = desc
				}

				if varType, ok := varConfig["type"].(string); ok && varType != "" {
					vType := chaos.TemplateVariableType(varType)
					variable.Type_ = &vType
				}

				if required, ok := varConfig["required"].(bool); ok {
					variable.Required = required
				}

				req.Variables[i] = variable
			}
		}
	}
	return nil
}

func buildSpecComplete(d *schema.ResourceData, req *chaos.ChaosfaulttemplateCreateFaultTemplateRequest) error {
	if v, ok := d.GetOk("spec"); ok && len(v.([]interface{})) > 0 {
		specConfig := v.([]interface{})[0].(map[string]interface{})
		req.Spec = &chaos.FaulttemplateSpec{}

		// Build chaos spec - Phase 1: fault_name and params only
		if chaosV, ok := specConfig["chaos"].([]interface{}); ok && len(chaosV) > 0 {
			chaosConfig := chaosV[0].(map[string]interface{})
			req.Spec.Chaos = &chaos.FaulttemplateChaosSpec{}

			// Fault name
			if faultName, ok := chaosConfig["fault_name"].(string); ok && faultName != "" {
				req.Spec.Chaos.FaultName = faultName
			}

			// Parameters
			if paramsV, ok := chaosConfig["params"].([]interface{}); ok && len(paramsV) > 0 {
				req.Spec.Chaos.Params = make([]chaos.FaulttemplateChaosParameter, len(paramsV))
				for i, paramItem := range paramsV {
					paramConfig := paramItem.(map[string]interface{})
					req.Spec.Chaos.Params[i] = chaos.FaulttemplateChaosParameter{
						Name:  paramConfig["name"].(string),
						Value: paramConfig["value"].(string),
					}
				}
			}

			// Build complete kubernetes spec from schema
			if err := buildChaosKubernetesSpec(chaosConfig, req.Spec.Chaos); err != nil {
				return err
			}

			// TODO: Auth and TLS - SDK types don't match schema, need to fix schema first
			// See FAULT_TEMPLATE_SDK_TYPE_MISMATCH.md for details
			// if err := buildChaosAuth(chaosConfig, req.Spec.Chaos); err != nil {
			// 	return err
			// }
			// if err := buildChaosTLS(chaosConfig, req.Spec.Chaos); err != nil {
			// 	return err
			// }
		}

		// Build target spec - Phase 1: basic kubernetes targets only
		if targetV, ok := specConfig["target"].([]interface{}); ok && len(targetV) > 0 {
			targetConfig := targetV[0].(map[string]interface{})
			req.Spec.Target = &chaos.FaulttemplateTarget{}

			// Kubernetes targets
			if k8sV, ok := targetConfig["kubernetes"].([]interface{}); ok && len(k8sV) > 0 {
				req.Spec.Target.Kubernetes = make([]chaos.FaulttemplateKubernetesTarget, len(k8sV))
				for i, k8sItem := range k8sV {
					k8sConfig := k8sItem.(map[string]interface{})
					k8sTarget := chaos.FaulttemplateKubernetesTarget{}

					if kind, ok := k8sConfig["kind"].(string); ok && kind != "" {
						k8sTarget.Kind = kind
					}

					if namespace, ok := k8sConfig["namespace"].(string); ok && namespace != "" {
						k8sTarget.Namespace = namespace
					}

					if labels, ok := k8sConfig["labels"].(map[string]interface{}); ok && len(labels) > 0 {
						// Convert labels map to comma-separated key=value pairs
						labelPairs := []string{}
						for k, v := range labels {
							labelPairs = append(labelPairs, fmt.Sprintf("%s=%s", k, v.(string)))
						}
						k8sTarget.Labels = strings.Join(labelPairs, ",")
					}

					if names, ok := k8sConfig["names"].([]interface{}); ok && len(names) > 0 {
						// Convert names slice to comma-separated string
						namesList := make([]string, len(names))
						for j, name := range names {
							namesList[j] = name.(string)
						}
						k8sTarget.Names = strings.Join(namesList, ",")
					}

					req.Spec.Target.Kubernetes[i] = k8sTarget
				}
			}

			// Application target
			if appV, ok := targetConfig["application"].([]interface{}); ok && len(appV) > 0 {
				appConfig := appV[0].(map[string]interface{})
				req.Spec.Target.Application = &chaos.FaulttemplateApplicationTarget{}

				if application, ok := appConfig["application"].(string); ok && application != "" {
					req.Spec.Target.Application.Application = application
				}
				if entrypoint, ok := appConfig["entrypoint"].(string); ok && entrypoint != "" {
					req.Spec.Target.Application.Entrypoint = entrypoint
				}
			}
		}
	}
	return nil
}

// ============================================================================
// BUILD FUNCTIONS: Complete Chaos Spec
// ============================================================================

func buildChaosKubernetesSpec(chaosConfig map[string]interface{}, chaosSpec *chaos.FaulttemplateChaosSpec) error {
	if k8sV, ok := chaosConfig["kubernetes"].([]interface{}); ok && len(k8sV) > 0 {
		k8sConfig := k8sV[0].(map[string]interface{})
		chaosSpec.Kubernetes = &chaos.FaulttemplateChaosKubernetesSpec{}
		k8s := chaosSpec.Kubernetes

		// Basic fields
		if image, ok := k8sConfig["image"].(string); ok && image != "" {
			k8s.Image = image
		} else {
			// Default image if not provided
			k8s.Image = "chaosnative/chaos-go-runner:ci"
		}

		if command, ok := k8sConfig["command"].([]interface{}); ok && len(command) > 0 {
			k8s.Command = make([]string, len(command))
			for i, cmd := range command {
				k8s.Command[i] = cmd.(string)
			}
		}

		if args, ok := k8sConfig["args"].([]interface{}); ok && len(args) > 0 {
			k8s.Args = make([]string, len(args))
			for i, arg := range args {
				k8s.Args[i] = arg.(string)
			}
		}

		if hostNetwork, ok := k8sConfig["host_network"].(bool); ok {
			k8s.HostNetwork = hostNetwork
		}

		if hostPID, ok := k8sConfig["host_pid"].(bool); ok {
			k8s.HostPID = hostPID
		}

		if hostIPC, ok := k8sConfig["host_ipc"].(bool); ok {
			k8s.HostIPC = hostIPC
		}

		if imagePullPolicy, ok := k8sConfig["image_pull_policy"].(string); ok && imagePullPolicy != "" {
			policy := chaos.V1PullPolicy(imagePullPolicy)
			k8s.ImagePullPolicy = &policy
		}

		if imagePullSecrets, ok := k8sConfig["image_pull_secrets"].([]interface{}); ok && len(imagePullSecrets) > 0 {
			k8s.ImagePullSecrets = make([]string, len(imagePullSecrets))
			for i, secret := range imagePullSecrets {
				k8s.ImagePullSecrets[i] = secret.(string)
			}
		}

		// Labels
		if labels, ok := k8sConfig["labels"].(map[string]interface{}); ok && len(labels) > 0 {
			k8s.Labels = make(map[string]string)
			for k, v := range labels {
				k8s.Labels[k] = v.(string)
			}
		}

		// Annotations
		if annotations, ok := k8sConfig["annotations"].(map[string]interface{}); ok && len(annotations) > 0 {
			k8s.Annotations = make(map[string]string)
			for k, v := range annotations {
				k8s.Annotations[k] = v.(string)
			}
		}

		// Node Selector
		if nodeSelector, ok := k8sConfig["node_selector"].(map[string]interface{}); ok && len(nodeSelector) > 0 {
			k8s.NodeSelector = make(map[string]string)
			for k, v := range nodeSelector {
				k8s.NodeSelector[k] = v.(string)
			}
		}

		// Environment variables
		if envV, ok := k8sConfig["env"].([]interface{}); ok && len(envV) > 0 {
			k8s.Env = make([]chaos.K8sIoApiCoreV1EnvVar, len(envV))
			for i, envItem := range envV {
				envConfig := envItem.(map[string]interface{})
				k8s.Env[i] = chaos.K8sIoApiCoreV1EnvVar{
					Name:  envConfig["name"].(string),
					Value: envConfig["value"].(string),
				}
			}
		}

		// Resource requirements
		if resourcesV, ok := k8sConfig["resources"].([]interface{}); ok && len(resourcesV) > 0 {
			resourcesConfig := resourcesV[0].(map[string]interface{})
			k8s.ResourceRequirements = &chaos.FaulttemplateResourceRequirements{}

			if limits, ok := resourcesConfig["limits"].(map[string]interface{}); ok && len(limits) > 0 {
				k8s.ResourceRequirements.Limits = make(map[string]string)
				for k, v := range limits {
					k8s.ResourceRequirements.Limits[k] = v.(string)
				}
			}

			if requests, ok := resourcesConfig["requests"].(map[string]interface{}); ok && len(requests) > 0 {
				k8s.ResourceRequirements.Requests = make(map[string]string)
				for k, v := range requests {
					k8s.ResourceRequirements.Requests[k] = v.(string)
				}
			}
		} else {
			// Initialize empty resource requirements (required by API)
			k8s.ResourceRequirements = &chaos.FaulttemplateResourceRequirements{}
		}

		// Tolerations
		if tolerationsV, ok := k8sConfig["tolerations"].([]interface{}); ok && len(tolerationsV) > 0 {
			k8s.Toleration = make([]chaos.V1Toleration, len(tolerationsV))
			for i, tolItem := range tolerationsV {
				tolConfig := tolItem.(map[string]interface{})
				toleration := chaos.V1Toleration{}

				if key, ok := tolConfig["key"].(string); ok {
					toleration.Key = key
				}
				if operator, ok := tolConfig["operator"].(string); ok {
					tolerationOperator := chaos.AllOfv1TolerationOperator(operator)
					toleration.Operator = &tolerationOperator
				}
				if value, ok := tolConfig["value"].(string); ok {
					toleration.Value = value
				}
				if effect, ok := tolConfig["effect"].(string); ok {
					tolerationEffect := chaos.AllOfv1TolerationEffect(effect)
					toleration.Effect = &tolerationEffect
				}
				if tolerationSeconds, ok := tolConfig["toleration_seconds"].(int); ok {
					toleration.TolerationSeconds = int32(tolerationSeconds)
				}

				k8s.Toleration[i] = toleration
			}
		}

		// ConfigMap volumes (schema field: config_maps)
		if configMapVolumes, ok := k8sConfig["config_maps"].([]interface{}); ok && len(configMapVolumes) > 0 {
			k8s.ConfigMapVolume = make([]chaos.FaulttemplateConfigMapVolume, len(configMapVolumes))
			for i, volItem := range configMapVolumes {
				volConfig := volItem.(map[string]interface{})
				volume := chaos.FaulttemplateConfigMapVolume{
					Name:      volConfig["name"].(string),
					MountPath: volConfig["mount_path"].(string),
				}
				if mountMode, ok := volConfig["mount_mode"].(int); ok {
					volume.MountMode = int32(mountMode)
				}
				k8s.ConfigMapVolume[i] = volume
			}
		}

		// Secret volumes (schema field: secrets)
		if secretVolumes, ok := k8sConfig["secrets"].([]interface{}); ok && len(secretVolumes) > 0 {
			k8s.SecretVolume = make([]chaos.FaulttemplateSecretVolume, len(secretVolumes))
			for i, volItem := range secretVolumes {
				volConfig := volItem.(map[string]interface{})
				volume := chaos.FaulttemplateSecretVolume{
					Name:      volConfig["secret_name"].(string),
					MountPath: volConfig["mount_path"].(string),
				}
				if mountMode, ok := volConfig["mount_mode"].(int); ok {
					volume.MountMode = int32(mountMode)
				}
				k8s.SecretVolume[i] = volume
			}
		}

		// HostPath volumes (schema field: host_file_volumes)
		if hostPathVolumes, ok := k8sConfig["host_file_volumes"].([]interface{}); ok && len(hostPathVolumes) > 0 {
			k8s.HostPathVolume = make([]chaos.FaulttemplateHostPathVolume, len(hostPathVolumes))
			for i, volItem := range hostPathVolumes {
				volConfig := volItem.(map[string]interface{})
				volume := chaos.FaulttemplateHostPathVolume{
					Name:      volConfig["name"].(string),
					MountPath: volConfig["mount_path"].(string),
				}
				if hostPath, ok := volConfig["host_path"].(string); ok && hostPath != "" {
					volume.HostPath = hostPath
				}
				if volType, ok := volConfig["type"].(string); ok && volType != "" {
					hpType := chaos.V1HostPathType(volType)
					volume.Type_ = &hpType
				}
				k8s.HostPathVolume[i] = volume
			}
		}

		// Pod security context
		if podSecCtx, ok := k8sConfig["pod_security_context"].([]interface{}); ok && len(podSecCtx) > 0 {
			pscConfig := podSecCtx[0].(map[string]interface{})
			k8s.PodSecurityContext = &chaos.V1PodSecurityContext{}

			if runAsUser, ok := pscConfig["run_as_user"].(int); ok {
				k8s.PodSecurityContext.RunAsUser = int32(runAsUser)
			}
			if runAsGroup, ok := pscConfig["run_as_group"].(int); ok {
				k8s.PodSecurityContext.RunAsGroup = int32(runAsGroup)
			}
			if fsGroup, ok := pscConfig["fs_group"].(int); ok {
				k8s.PodSecurityContext.FsGroup = int32(fsGroup)
			}
			if runAsNonRoot, ok := pscConfig["run_as_non_root"].(bool); ok {
				k8s.PodSecurityContext.RunAsNonRoot = runAsNonRoot
			}
		}

		// Container security context
		if cscList, ok := k8sConfig["container_security_context"].([]interface{}); ok && len(cscList) > 0 {
			cscConfig := cscList[0].(map[string]interface{})
			k8s.ContainerSecurityContext = &chaos.V1SecurityContext{}

			if privileged, ok := cscConfig["privileged"].(bool); ok {
				k8s.ContainerSecurityContext.Privileged = privileged
			}
			if readOnlyRootFS, ok := cscConfig["read_only_root_filesystem"].(bool); ok {
				k8s.ContainerSecurityContext.ReadOnlyRootFilesystem = readOnlyRootFS
			}
			if allowPrivEsc, ok := cscConfig["allow_privilege_escalation"].(bool); ok {
				k8s.ContainerSecurityContext.AllowPrivilegeEscalation = allowPrivEsc
			}
			if runAsUser, ok := cscConfig["run_as_user"].(int); ok {
				k8s.ContainerSecurityContext.RunAsUser = int32(runAsUser)
			}
			if runAsGroup, ok := cscConfig["run_as_group"].(int); ok {
				k8s.ContainerSecurityContext.RunAsGroup = int32(runAsGroup)
			}
			if runAsNonRoot, ok := cscConfig["run_as_non_root"].(bool); ok {
				k8s.ContainerSecurityContext.RunAsNonRoot = runAsNonRoot
			}

			// Capabilities
			if capsList, ok := cscConfig["capabilities"].([]interface{}); ok && len(capsList) > 0 {
				capsConfig := capsList[0].(map[string]interface{})
				caps := &chaos.AllOfv1SecurityContextCapabilities{}

				if add, ok := capsConfig["add"].([]interface{}); ok && len(add) > 0 {
					caps.Add = make([]string, len(add))
					for i, cap := range add {
						caps.Add[i] = cap.(string)
					}
				}
				if drop, ok := capsConfig["drop"].([]interface{}); ok && len(drop) > 0 {
					caps.Drop = make([]string, len(drop))
					for i, cap := range drop {
						caps.Drop[i] = cap.(string)
					}
				}

				k8s.ContainerSecurityContext.Capabilities = caps
			}
		} else {
			// Initialize empty security context (required by API)
			k8s.ContainerSecurityContext = &chaos.V1SecurityContext{}
		}
	}
	return nil
}

func buildChaosAuth(chaosConfig map[string]interface{}, chaosSpec *chaos.FaulttemplateChaosSpec) error {
	if authV, ok := chaosConfig["auth"].([]interface{}); ok && len(authV) > 0 {
		authConfig := authV[0].(map[string]interface{})
		chaosSpec.Auth = &chaos.FaulttemplateAuth{}

		// AWS Auth
		if awsV, ok := authConfig["aws"].([]interface{}); ok && len(awsV) > 0 {
			awsConfig := awsV[0].(map[string]interface{})
			chaosSpec.Auth.Aws = &chaos.FaulttemplateAwsAuth{
				Identifier: awsConfig["identifier"].(string),
			}
		}

		// Azure Auth
		if azureV, ok := authConfig["azure"].([]interface{}); ok && len(azureV) > 0 {
			azureConfig := azureV[0].(map[string]interface{})
			chaosSpec.Auth.Azure = &chaos.FaulttemplateAzureAuth{
				Identifier: azureConfig["identifier"].(string),
			}
		}

		// GCP Auth
		if gcpV, ok := authConfig["gcp"].([]interface{}); ok && len(gcpV) > 0 {
			gcpConfig := gcpV[0].(map[string]interface{})
			chaosSpec.Auth.Gcp = &chaos.FaulttemplateGcpAuth{
				Identifier: gcpConfig["identifier"].(string),
			}
		}

		// Redis Auth
		if redisV, ok := authConfig["redis"].([]interface{}); ok && len(redisV) > 0 {
			redisConfig := redisV[0].(map[string]interface{})
			chaosSpec.Auth.Redis = &chaos.FaulttemplateRedisAuth{
				Password: redisConfig["password"].(string),
			}
		}

		// SSH Auth
		if sshV, ok := authConfig["ssh"].([]interface{}); ok && len(sshV) > 0 {
			sshConfig := sshV[0].(map[string]interface{})
			chaosSpec.Auth.Ssh = &chaos.FaulttemplateSshAuth{
				Key:      sshConfig["key"].(string),
				Password: sshConfig["password"].(string),
			}
		}

		// VMware Auth
		if vmwareV, ok := authConfig["vmware"].([]interface{}); ok && len(vmwareV) > 0 {
			vmwareConfig := vmwareV[0].(map[string]interface{})
			chaosSpec.Auth.Vmware = &chaos.FaulttemplateVmWareAuth{
				VCenterServer:   vmwareConfig["vcenter_server"].(string),
				VCenterUsername: vmwareConfig["vcenter_username"].(string),
				VCenterPassword: vmwareConfig["vcenter_password"].(string),
				VmPassword:      vmwareConfig["vm_password"].(string),
			}
		}
	}
	return nil
}

func buildChaosTLS(chaosConfig map[string]interface{}, chaosSpec *chaos.FaulttemplateChaosSpec) error {
	if tlsV, ok := chaosConfig["tls"].([]interface{}); ok && len(tlsV) > 0 {
		tlsConfig := tlsV[0].(map[string]interface{})
		chaosSpec.Tls = &chaos.FaulttemplateTls{}

		if caFile, ok := tlsConfig["ca_file"].(string); ok && caFile != "" {
			chaosSpec.Tls.CaFile = caFile
		}
		if clientCertFile, ok := tlsConfig["client_cert_file"].(string); ok && clientCertFile != "" {
			chaosSpec.Tls.ClientCertFile = clientCertFile
		}
		if certFile, ok := tlsConfig["cert_file"].(string); ok && certFile != "" {
			chaosSpec.Tls.CertFile = certFile
		}
		if keyFile, ok := tlsConfig["key_file"].(string); ok && keyFile != "" {
			chaosSpec.Tls.KeyFile = keyFile
		}
	}
	return nil
}

func setFaultTemplateData(d *schema.ResourceData, template *chaos.ChaosfaulttemplateChaosFaultTemplate) error {
	// Basic identifiers
	d.Set("account_id", template.AccountID)
	d.Set("org_id", template.OrgID)
	d.Set("project_id", template.ProjectID)
	d.Set("hub_ref", template.HubRef)
	d.Set("identity", template.Identity)
	d.Set("name", template.Name)
	d.Set("description", template.Description)

	// Metadata
	d.Set("tags", template.Tags)
	d.Set("category", template.Category)
	d.Set("infrastructures", template.Infras)
	d.Set("type", template.Type_)
	d.Set("is_enterprise", template.IsEnterprise)
	d.Set("is_removed", template.IsRemoved)
	d.Set("permissions_required", template.PermissionsRequired)

	// Computed fields
	d.Set("revision", template.Revision)
	d.Set("created_at", template.CreatedAt)
	d.Set("created_by", template.CreatedBy)
	d.Set("updated_at", template.UpdatedAt)
	d.Set("updated_by", template.UpdatedBy)

	// Links
	if len(template.Links) > 0 {
		links := make([]map[string]interface{}, len(template.Links))
		for i, link := range template.Links {
			links[i] = map[string]interface{}{
				"name": link.Name,
				"url":  link.Url,
			}
		}
		d.Set("links", links)
	}

	// Variables
	if len(template.Variables) > 0 {
		variables := make([]map[string]interface{}, len(template.Variables))
		for i, variable := range template.Variables {
			varMap := map[string]interface{}{
				"name":     variable.Name,
				"required": variable.Required,
			}

			if variable.Value != nil {
				if val, ok := (*variable.Value).(string); ok {
					varMap["value"] = val
				}
			}

			if variable.Description != "" {
				varMap["description"] = variable.Description
			}

			if variable.Type_ != nil {
				varMap["type"] = strings.ToLower(string(*variable.Type_))
			}

			variables[i] = varMap
		}
		d.Set("variables", variables)
	}

	// Parse spec from template YAML
	if template.Template != "" {
		if err := readSpecFromYAML(d, template.Template); err != nil {
			log.Printf("[WARN] Failed to parse template YAML: %v", err)
			// Don't fail the read, just log warning
		}
	}

	// Note: Keywords and Platforms are in the YAML template if present,
	// not as separate fields in the API response

	return nil
}

// ============================================================================
// READ FUNCTIONS: Parse Spec from YAML
// ============================================================================

func readSpecFromYAML(d *schema.ResourceData, templateYAML string) error {
	// Parse YAML
	var templateData map[string]interface{}
	if err := yaml.Unmarshal([]byte(templateYAML), &templateData); err != nil {
		return fmt.Errorf("failed to unmarshal template YAML: %w", err)
	}

	// Extract spec
	specData, ok := templateData["spec"].(map[string]interface{})
	if !ok {
		return nil // No spec in template
	}

	spec := []map[string]interface{}{}
	specMap := map[string]interface{}{}

	// Read chaos spec
	if chaosData, ok := specData["chaos"].(map[string]interface{}); ok {
		chaosSpec := readChaosSpecFromYAML(chaosData)
		if len(chaosSpec) > 0 {
			specMap["chaos"] = []map[string]interface{}{chaosSpec}
		}
	}

	// Read target spec
	if targetData, ok := specData["target"].(map[string]interface{}); ok {
		targetSpec := readTargetSpecFromYAML(targetData)
		if len(targetSpec) > 0 {
			specMap["target"] = []map[string]interface{}{targetSpec}
		}
	}

	if len(specMap) > 0 {
		spec = append(spec, specMap)
		d.Set("spec", spec)
	}

	return nil
}

func readChaosSpecFromYAML(chaosData map[string]interface{}) map[string]interface{} {
	chaosSpec := map[string]interface{}{}

	// Fault name
	if faultName, ok := chaosData["faultName"].(string); ok {
		chaosSpec["fault_name"] = faultName
	}

	// Parameters
	if paramsData, ok := chaosData["params"].([]interface{}); ok && len(paramsData) > 0 {
		params := make([]map[string]interface{}, 0, len(paramsData))
		for _, paramItem := range paramsData {
			if paramMap, ok := paramItem.(map[string]interface{}); ok {
				param := map[string]interface{}{}
				if name, ok := paramMap["name"].(string); ok {
					param["name"] = name
				}
				if value, ok := paramMap["value"]; ok {
					param["value"] = fmt.Sprintf("%v", value)
				}
				if len(param) > 0 {
					params = append(params, param)
				}
			}
		}
		if len(params) > 0 {
			chaosSpec["params"] = params
		}
	}

	// Kubernetes spec
	if k8sData, ok := chaosData["kubernetes"].(map[string]interface{}); ok {
		k8sSpec := readChaosKubernetesSpecFromYAML(k8sData)
		if len(k8sSpec) > 0 {
			chaosSpec["kubernetes"] = []map[string]interface{}{k8sSpec}
		}
	}

	return chaosSpec
}

func readChaosKubernetesSpecFromYAML(k8sData map[string]interface{}) map[string]interface{} {
	k8sSpec := map[string]interface{}{}

	// Basic fields
	if image, ok := k8sData["image"].(string); ok {
		k8sSpec["image"] = image
	}
	if command, ok := k8sData["command"].([]interface{}); ok {
		k8sSpec["command"] = command
	}
	if args, ok := k8sData["args"].([]interface{}); ok {
		k8sSpec["args"] = args
	}
	if hostNetwork, ok := k8sData["hostNetwork"].(bool); ok {
		k8sSpec["host_network"] = hostNetwork
	}
	if hostPID, ok := k8sData["hostPID"].(bool); ok {
		k8sSpec["host_pid"] = hostPID
	}
	if hostIPC, ok := k8sData["hostIPC"].(bool); ok {
		k8sSpec["host_ipc"] = hostIPC
	}
	if imagePullPolicy, ok := k8sData["imagePullPolicy"].(string); ok {
		k8sSpec["image_pull_policy"] = imagePullPolicy
	}
	if imagePullSecrets, ok := k8sData["imagePullSecrets"].([]interface{}); ok {
		k8sSpec["image_pull_secrets"] = imagePullSecrets
	}

	// Labels, Annotations, NodeSelector
	if labels, ok := k8sData["labels"].(map[string]interface{}); ok {
		k8sSpec["labels"] = labels
	}
	if annotations, ok := k8sData["annotations"].(map[string]interface{}); ok {
		k8sSpec["annotations"] = annotations
	}
	if nodeSelector, ok := k8sData["nodeSelector"].(map[string]interface{}); ok {
		k8sSpec["node_selector"] = nodeSelector
	}

	// Environment variables
	if envData, ok := k8sData["env"].([]interface{}); ok && len(envData) > 0 {
		envVars := make([]map[string]interface{}, 0, len(envData))
		for _, envItem := range envData {
			if envMap, ok := envItem.(map[string]interface{}); ok {
				envVar := map[string]interface{}{}
				if name, ok := envMap["name"].(string); ok {
					envVar["name"] = name
				}
				if value, ok := envMap["value"]; ok {
					envVar["value"] = fmt.Sprintf("%v", value)
				}
				if len(envVar) > 0 {
					envVars = append(envVars, envVar)
				}
			}
		}
		if len(envVars) > 0 {
			k8sSpec["env"] = envVars
		}
	}

	// Resource requirements
	if resourcesData, ok := k8sData["resourceRequirements"].(map[string]interface{}); ok {
		resources := map[string]interface{}{}
		if limits, ok := resourcesData["limits"].(map[string]interface{}); ok {
			resources["limits"] = limits
		}
		if requests, ok := resourcesData["requests"].(map[string]interface{}); ok {
			resources["requests"] = requests
		}
		if len(resources) > 0 {
			k8sSpec["resources"] = []map[string]interface{}{resources}
		}
	}

	// Tolerations
	if tolerationsData, ok := k8sData["toleration"].([]interface{}); ok && len(tolerationsData) > 0 {
		tolerations := make([]map[string]interface{}, 0, len(tolerationsData))
		for _, tolItem := range tolerationsData {
			if tolMap, ok := tolItem.(map[string]interface{}); ok {
				toleration := map[string]interface{}{}
				if key, ok := tolMap["key"].(string); ok {
					toleration["key"] = key
				}
				if operator, ok := tolMap["operator"].(string); ok {
					toleration["operator"] = operator
				}
				if value, ok := tolMap["value"].(string); ok {
					toleration["value"] = value
				}
				if effect, ok := tolMap["effect"].(string); ok {
					toleration["effect"] = effect
				}
				if tolerationSeconds, ok := tolMap["tolerationSeconds"].(int); ok {
					toleration["toleration_seconds"] = tolerationSeconds
				}
				if len(toleration) > 0 {
					tolerations = append(tolerations, toleration)
				}
			}
		}
		if len(tolerations) > 0 {
			k8sSpec["tolerations"] = tolerations
		}
	}

	// ConfigMap volumes
	if configMapVolumes, ok := k8sData["configMapVolume"].([]interface{}); ok && len(configMapVolumes) > 0 {
		volumes := make([]map[string]interface{}, 0, len(configMapVolumes))
		for _, volItem := range configMapVolumes {
			if volMap, ok := volItem.(map[string]interface{}); ok {
				volume := map[string]interface{}{}
				if name, ok := volMap["name"].(string); ok {
					volume["name"] = name
				}
				if mountPath, ok := volMap["mountPath"].(string); ok {
					volume["mount_path"] = mountPath
				}
				if mountMode, ok := volMap["mountMode"].(int); ok {
					volume["mount_mode"] = mountMode
				}
				if len(volume) > 0 {
					volumes = append(volumes, volume)
				}
			}
		}
		if len(volumes) > 0 {
			k8sSpec["config_map_volume"] = volumes
		}
	}

	// Secret volumes
	if secretVolumes, ok := k8sData["secretVolume"].([]interface{}); ok && len(secretVolumes) > 0 {
		volumes := make([]map[string]interface{}, 0, len(secretVolumes))
		for _, volItem := range secretVolumes {
			if volMap, ok := volItem.(map[string]interface{}); ok {
				volume := map[string]interface{}{}
				if name, ok := volMap["name"].(string); ok {
					volume["name"] = name
				}
				if mountPath, ok := volMap["mountPath"].(string); ok {
					volume["mount_path"] = mountPath
				}
				if mountMode, ok := volMap["mountMode"].(int); ok {
					volume["mount_mode"] = mountMode
				}
				if len(volume) > 0 {
					volumes = append(volumes, volume)
				}
			}
		}
		if len(volumes) > 0 {
			k8sSpec["secret_volume"] = volumes
		}
	}

	// HostPath volumes
	if hostPathVolumes, ok := k8sData["hostPathVolume"].([]interface{}); ok && len(hostPathVolumes) > 0 {
		volumes := make([]map[string]interface{}, 0, len(hostPathVolumes))
		for _, volItem := range hostPathVolumes {
			if volMap, ok := volItem.(map[string]interface{}); ok {
				volume := map[string]interface{}{}
				if name, ok := volMap["name"].(string); ok {
					volume["name"] = name
				}
				if hostPath, ok := volMap["hostPath"].(string); ok {
					volume["host_path"] = hostPath
				}
				if mountPath, ok := volMap["mountPath"].(string); ok {
					volume["mount_path"] = mountPath
				}
				if hostPathType, ok := volMap["type"].(string); ok {
					volume["type"] = hostPathType
				}
				if len(volume) > 0 {
					volumes = append(volumes, volume)
				}
			}
		}
		if len(volumes) > 0 {
			k8sSpec["host_path_volume"] = volumes
		}
	}

	return k8sSpec
}

func readTargetSpecFromYAML(targetData map[string]interface{}) map[string]interface{} {
	targetSpec := map[string]interface{}{}

	// Kubernetes targets
	if k8sTargets, ok := targetData["kubernetes"].([]interface{}); ok && len(k8sTargets) > 0 {
		targets := make([]map[string]interface{}, 0, len(k8sTargets))
		for _, targetItem := range k8sTargets {
			if targetMap, ok := targetItem.(map[string]interface{}); ok {
				target := map[string]interface{}{}
				if kind, ok := targetMap["kind"].(string); ok {
					target["kind"] = kind
				}
				if namespace, ok := targetMap["namespace"].(string); ok {
					target["namespace"] = namespace
				}
				if labels, ok := targetMap["labels"].(string); ok {
					target["labels"] = labels
				}
				if names, ok := targetMap["names"].(string); ok {
					target["names"] = names
				}
				if len(target) > 0 {
					targets = append(targets, target)
				}
			}
		}
		if len(targets) > 0 {
			targetSpec["kubernetes"] = targets
		}
	}

	// Application target
	if appData, ok := targetData["application"].(map[string]interface{}); ok {
		app := map[string]interface{}{}
		if application, ok := appData["application"].(string); ok {
			app["application"] = application
		}
		if entrypoint, ok := appData["entrypoint"].(string); ok {
			app["entrypoint"] = entrypoint
		}
		if len(app) > 0 {
			targetSpec["application"] = []map[string]interface{}{app}
		}
	}

	return targetSpec
}
