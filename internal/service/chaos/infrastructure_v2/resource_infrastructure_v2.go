package infrastructure_v2

import (
	"context"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/chaos"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceChaosInfrastructureV2() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for managing Harness Chaos Infrastructure V2.",

		CreateContext: resourceInfrastructureV2Create,
		ReadContext:   resourceInfrastructureV2Read,
		UpdateContext: resourceInfrastructureV2Update,
		DeleteContext: resourceInfrastructureV2Delete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceInfrastructureV2Import,
		},
		Schema: resourceChaosInfrastructureV2Schema(),
	}
}

// resourceInfrastructureV2Create creates a new infrastructure
func resourceInfrastructureV2Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)

	// Build the create request
	req, err := buildRegisterInfrastructureV2Request(d, c.AccountId)
	if err != nil {
		return diag.FromErr(err)
	}

	// Extract identifiers
	orgID := d.Get("org_id").(string)
	projectID := d.Get("project_id").(string)
	environmentID := d.Get("environment_id").(string)

	// Make the API call with correct parameter order
	resp, httpResp, err := c.ChaosSdkApi.RegisterInfraV2(
		ctx,
		*req,
		c.AccountId,
		orgID,
		projectID,
		&chaos.ChaosSdkApiRegisterInfraV2Opts{
			CorrelationID: optional.NewString(d.Get("correlation_id").(string)),
		},
	)
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	// Set the ID
	d.SetId(fmt.Sprintf("%s/%s/%s/%s", orgID, projectID, environmentID, resp.Identity))

	log.Printf("Created infrastructure with ID: %s", d.Id())

	// Check if we need to perform an immediate update for fields not supported in create
	needsUpdate := d.HasChanges(
		"volumes", "volume_mounts", "env", "image_registry", "label",
		"annotation", "containers", "insecure_skip_verify",
	)
	if needsUpdate {
		log.Printf("[DEBUG] Performing immediate update after creation for additional fields")
		return resourceInfrastructureV2Update(ctx, d, meta)
	}

	return resourceInfrastructureV2Read(ctx, d, meta)
}

// resourceInfrastructureV2Read reads the infrastructure details
func resourceInfrastructureV2Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)

	// Parse the ID
	orgID, projectID, environmentID, infraID, err := parseInfrastructureV2ID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	// Get the infrastructure with all required parameters
	infra, httpResp, err := c.ChaosSdkApi.GetInfraV2(
		ctx,
		infraID,
		c.AccountId,
		orgID,
		projectID,
		environmentID,
	)
	if err != nil {
		if httpResp != nil && httpResp.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return helpers.HandleApiError(err, d, httpResp)
	}

	// Set the fields
	if err := d.Set("name", infra.Name); err != nil {
		return diag.FromErr(fmt.Errorf("failed to set name: %v", err))
	}
	if err := d.Set("infra_id", infra.InfraID); err != nil {
		return diag.FromErr(fmt.Errorf("failed to set infra_id: %v", err))
	}
	if err := d.Set("infra_namespace", infra.InfraNamespace); err != nil {
		return diag.FromErr(fmt.Errorf("failed to set infra_namespace: %v", err))
	}
	if err := d.Set("service_account", infra.ServiceAccount); err != nil {
		return diag.FromErr(fmt.Errorf("failed to set service_account: %v", err))
	}
	if err := d.Set("description", infra.Description); err != nil {
		return diag.FromErr(fmt.Errorf("failed to set description: %v", err))
	}
	if err := d.Set("k8s_connector_id", infra.K8sConnectorID); err != nil {
		return diag.FromErr(fmt.Errorf("failed to set k8s_connector_id: %v", err))
	}
	if err := d.Set("created_at", infra.CreatedAt); err != nil {
		return diag.FromErr(fmt.Errorf("failed to set created_at: %v", err))
	}
	if err := d.Set("updated_at", infra.UpdatedAt); err != nil {
		return diag.FromErr(fmt.Errorf("failed to set updated_at: %v", err))
	}
	if err := d.Set("status", infra.Status); err != nil {
		return diag.FromErr(fmt.Errorf("failed to set status: %v", err))
	}
	if err := d.Set("identity", infra.Identity); err != nil {
		return diag.FromErr(fmt.Errorf("failed to set identity: %v", err))
	}
	if err := d.Set("environment_id", environmentID); err != nil {
		return diag.FromErr(fmt.Errorf("failed to set environment_id: %v", err))
	}

	// Set maps
	if err := d.Set("label", infra.Label); err != nil {
		return diag.FromErr(fmt.Errorf("failed to set label: %v", err))
	}
	if err := d.Set("annotation", infra.Annotation); err != nil {
		return diag.FromErr(fmt.Errorf("failed to set annotation: %v", err))
	}
	if err := d.Set("node_selector", infra.NodeSelector); err != nil {
		return diag.FromErr(fmt.Errorf("failed to set node_selector: %v", err))
	}

	// Debug log the entire infra object to see what we got from the API
	log.Printf("[DEBUG] Infrastructure details from API: %+v", infra)

	// Set the image registry if it exists
	if infra.ImageRegistry != nil {
		// Ensure the infra_id is properly formatted
		if infra.ImageRegistry.InfraID == "" {
			ensureRegistryInfraID(infra.ImageRegistry, environmentID, infraID)
		}
		log.Printf("[DEBUG] Setting image_registry from API: %+v", infra.ImageRegistry)
		if err := setImageRegistry(d, infra.ImageRegistry); err != nil {
			return diag.FromErr(fmt.Errorf("failed to set image_registry: %v", err))
		}
	} else {
		log.Printf("[DEBUG] No image_registry in API response, using default values")
		// If no image_registry in API response, set default values to match the schema
		defaultRegistry := &chaos.ImageRegistryImageRegistryV2{
			RegistryServer:    "docker.io",
			RegistryAccount:   "harness",
			IsPrivate:         false,
			UseCustomImages:   false,
			IsOverrideAllowed: false,
		}
		ensureRegistryInfraID(defaultRegistry, environmentID, infraID)
		if err := setImageRegistry(d, defaultRegistry); err != nil {
			return diag.FromErr(fmt.Errorf("failed to set default image_registry: %v", err))
		}
	}

	// Set nested objects
	if err := setMtls(d, infra.Mtls); err != nil {
		return diag.FromErr(err)
	}
	if err := setProxy(d, infra.Proxy); err != nil {
		return diag.FromErr(err)
	}
	if err := setTolerations(d, infra.Tolerations); err != nil {
		return diag.FromErr(err)
	}
	if err := setVolumes(d, infra.Volumes); err != nil {
		return diag.FromErr(err)
	}
	if err := setVolumeMounts(d, infra.VolumeMounts); err != nil {
		return diag.FromErr(err)
	}
	if err := setEnvVars(d, infra.Env); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceInfrastructureV2Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)

	// Parse the ID
	orgID, projectID, _, infraID, err := parseInfrastructureV2ID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	// Build the update request
	req, err := buildUpdateInfrastructureV2Request(d, c.AccountId)
	if err != nil {
		return diag.FromErr(err)
	}
	req.Identity = infraID

	// Make the API call
	_, httpResp, err := c.ChaosSdkApi.UpdateInfraV2(
		ctx,
		*req,
		c.AccountId,
		orgID,
		projectID,
		&chaos.ChaosSdkApiUpdateInfraV2Opts{
			CorrelationID: optional.NewString(d.Get("correlation_id").(string)),
		},
	)
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	return resourceInfrastructureV2Read(ctx, d, meta)
}

func resourceInfrastructureV2Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)

	// Parse the ID
	orgID, projectID, environmentID, infraID, err := parseInfrastructureV2ID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	// Make the API call
	_, httpResp, err := c.ChaosSdkApi.DeleteInfraV2(
		ctx,
		infraID,
		environmentID,
		c.AccountId,
		orgID,
		projectID,
	)
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	d.SetId("")
	return nil
}

// parseInfrastructureV2ID parses the resource ID into its components
// Format: orgID/projectID/environmentID/infraID
func parseInfrastructureV2ID(id string) (string, string, string, string, error) {
	parts := strings.Split(id, "/")
	if len(parts) != 4 {
		return "", "", "", "", fmt.Errorf("invalid ID format, expected org_id/project_id/environment_id/infra_id, got %s", id)
	}
	return parts[0], parts[1], parts[2], parts[3], nil
}

// Setter functions for nested objects
func setImageRegistry(d *schema.ResourceData, reg *chaos.ImageRegistryImageRegistryV2) error {
	if reg == nil {
		log.Printf("[DEBUG] No image registry data provided, setting to nil")
		return d.Set("image_registry", []interface{}{})
	}

	log.Printf("[DEBUG] Setting image registry data: %+v", reg)

	// Initialize the result with default values that match the schema
	result := map[string]interface{}{
		"registry_server":     "docker.io",
		"registry_account":    "harness",
		"is_private":          false,
		"use_custom_images":   false,
		"is_default":          true,
		"is_override_allowed": false,
		"infra_id":            reg.InfraID,
	}

	// Override with values from the API response if they exist
	if reg.RegistryServer != "" {
		result["registry_server"] = reg.RegistryServer
	}
	if reg.RegistryAccount != "" {
		result["registry_account"] = reg.RegistryAccount
	}
	if reg.IsPrivate {
		result["is_private"] = true
	}
	if reg.SecretName != "" {
		result["secret_name"] = reg.SecretName
	}
	if reg.UseCustomImages {
		result["use_custom_images"] = true
	}
	if !reg.IsDefault {
		result["is_default"] = false
	}
	if reg.IsOverrideAllowed {
		result["is_override_allowed"] = true
	}

	// Handle custom images if they exist
	if reg.CustomImages != nil {
		customImages := map[string]interface{}{
			"ddcr":        "",
			"ddcr_fault":  "",
			"ddcr_lib":    "",
			"log_watcher": "",
		}

		if reg.CustomImages.Ddcr != "" {
			customImages["ddcr"] = reg.CustomImages.Ddcr
		}
		if reg.CustomImages.DdcrFault != "" {
			customImages["ddcr_fault"] = reg.CustomImages.DdcrFault
		}
		if reg.CustomImages.DdcrLib != "" {
			customImages["ddcr_lib"] = reg.CustomImages.DdcrLib
		}
		if reg.CustomImages.LogWatcher != "" {
			customImages["log_watcher"] = reg.CustomImages.LogWatcher
		}

		result["custom_images"] = []map[string]interface{}{customImages}
	}

	// Handle identifier if it exists
	if reg.Identifier != nil {
		identifier := map[string]interface{}{
			"account_identifier": "",
			"org_identifier":     "",
			"project_identifier": "",
		}

		if reg.Identifier.AccountIdentifier != "" {
			identifier["account_identifier"] = reg.Identifier.AccountIdentifier
		}
		if reg.Identifier.OrgIdentifier != "" {
			identifier["org_identifier"] = reg.Identifier.OrgIdentifier
		}
		if reg.Identifier.ProjectIdentifier != "" {
			identifier["project_identifier"] = reg.Identifier.ProjectIdentifier
		}

		result["identifier"] = []map[string]interface{}{identifier}
	}

	log.Printf("[DEBUG] Final image registry state: %+v", result)
	return d.Set("image_registry", []interface{}{result})
}

func setMtls(d *schema.ResourceData, mtls *chaos.InfraV2MtlsConfiguration) error {
	if mtls == nil {
		return d.Set("mtls", nil)
	}

	return d.Set("mtls", []interface{}{map[string]interface{}{
		"cert_path":   mtls.CertPath,
		"key_path":    mtls.KeyPath,
		"secret_name": mtls.SecretName,
		"url":         mtls.Url,
	}})
}

func setProxy(d *schema.ResourceData, proxy *chaos.InfraV2ProxyConfiguration) error {
	if proxy == nil {
		return d.Set("proxy", nil)
	}

	result := map[string]interface{}{
		"url":         proxy.Url,
		"http_proxy":  proxy.HttpProxy,
		"https_proxy": proxy.HttpsProxy,
		"no_proxy":    proxy.NoProxy,
	}

	return d.Set("proxy", []interface{}{result})
}

func setTolerations(d *schema.ResourceData, tolerations []chaos.V1Toleration) error {
	if tolerations == nil {
		return d.Set("tolerations", nil)
	}

	result := make([]interface{}, len(tolerations))
	for i, t := range tolerations {
		toleration := map[string]interface{}{
			"key":                t.Key,
			"value":              t.Value,
			"toleration_seconds": t.TolerationSeconds,
		}

		// Set operator if it exists
		if t.Operator != nil {
			toleration["operator"] = string(*t.Operator)
		}

		// Set effect if it exists
		if t.Effect != nil {
			toleration["effect"] = string(*t.Effect)
		}

		result[i] = toleration
	}

	return d.Set("tolerations", result)
}

func setVolumes(d *schema.ResourceData, volumes []chaos.InfraV2Volumes) error {
	if volumes == nil {
		return d.Set("volumes", nil)
	}

	result := make([]interface{}, len(volumes))
	for i, v := range volumes {
		result[i] = map[string]interface{}{
			"name":       v.Name,
			"size_limit": v.SizeLimit,
		}
	}

	return d.Set("volumes", result)
}

func setVolumeMounts(d *schema.ResourceData, mounts []chaos.V1VolumeMount) error {
	if mounts == nil {
		return d.Set("volume_mounts", nil)
	}

	result := make([]interface{}, len(mounts))
	for i, m := range mounts {
		mount := map[string]interface{}{
			"mount_path":    m.MountPath,
			"name":          m.Name,
			"read_only":     m.ReadOnly,
			"sub_path":      m.SubPath,
			"sub_path_expr": m.SubPathExpr,
		}

		// Set mount propagation if it exists
		if m.MountPropagation != nil {
			mount["mount_propagation"] = string(*m.MountPropagation)
		}

		result[i] = mount
	}

	return d.Set("volume_mounts", result)
}

func setEnvVars(d *schema.ResourceData, envVars []chaos.InfraV2Env) error {
	if envVars == nil {
		return d.Set("env", nil)
	}

	result := make([]interface{}, len(envVars))
	for i, e := range envVars {
		env := map[string]interface{}{
			"name":  e.Name,
			"value": e.Value,
		}

		if e.ValueFrom != nil {
			env["value_from"] = string(*e.ValueFrom)
			if e.Key != "" {
				env["key"] = e.Key
			}
		}

		result[i] = env
	}

	return d.Set("env", result)
}

// buildRegisterInfrastructureV2Request builds the request object for create operations
func buildRegisterInfrastructureV2Request(d *schema.ResourceData, accountID string) (*chaos.InfraV2RegisterInfrastructureV2Request, error) {
	infraType := chaos.InfraV2InfraType(d.Get("infra_type").(string))
	infraScope := chaos.InfraV2InfraScope(d.Get("infra_scope").(string))

	// In resourceInfrastructureV2Create and resourceInfrastructureV2Update functions
	name := d.Get("name").(string)
	sanitizedName := sanitizeK8sResourceName(name)

	// If the sanitized name is different, log a warning
	if name != sanitizedName {
		log.Printf("[WARN] Infrastructure name '%s' has been sanitized to '%s' to comply with Kubernetes naming requirements", name, sanitizedName)
	}

	req := &chaos.InfraV2RegisterInfrastructureV2Request{
		Name:               sanitizedName,
		Identity:           d.Get("infra_id").(string), // set identity from infra_id
		InfraID:            d.Get("infra_id").(string),
		InfraType:          &infraType,
		InfraScope:         &infraScope,
		Description:        d.Get("description").(string),
		K8sConnectorID:     d.Get("k8s_connector_id").(string),
		ServiceAccount:     d.Get("service_account").(string),
		InfraNamespace:     d.Get("namespace").(string),
		EnvironmentID:      d.Get("environment_id").(string),
		AiEnabled:          d.Get("ai_enabled").(bool),
		InsecureSkipVerify: d.Get("insecure_skip_verify").(bool),
		Containers:         d.Get("containers").(string),
		RunAsUser:          int32(d.Get("run_as_user").(int)),
		RunAsGroup:         int32(d.Get("run_as_group").(int)),
	}

	// Set maps
	if v, ok := d.GetOk("label"); ok {
		req.Label = expandStringMap(v.(map[string]interface{}))
	}
	if v, ok := d.GetOk("annotation"); ok {
		req.Annotation = expandStringMap(v.(map[string]interface{}))
	}
	if v, ok := d.GetOk("node_selector"); ok {
		req.NodeSelector = expandStringMap(v.(map[string]interface{}))
	}

	// Set nested objects
	req.Identifier = &chaos.InfraV2Identifiers{
		AccountIdentifier: accountID,
		OrgIdentifier:     d.Get("org_id").(string),
		ProjectIdentifier: d.Get("project_id").(string),
	}
	if v, ok := d.GetOk("image_registry"); ok {
		req.ImageRegistry = expandImageRegistry(v.([]interface{}), d, accountID)
	}
	if v, ok := d.GetOk("mtls"); ok {
		req.Mtls = expandMtls(v.([]interface{}))
	}
	if v, ok := d.GetOk("proxy"); ok {
		req.Proxy = expandProxy(v.([]interface{}))
	}
	if v, ok := d.GetOk("tolerations"); ok {
		req.Tolerations = expandTolerations(v.([]interface{}))
	}
	if v, ok := d.GetOk("volumes"); ok {
		req.Volumes = expandVolumes(v.([]interface{}))
	}
	if v, ok := d.GetOk("volume_mounts"); ok {
		req.VolumeMounts = expandVolumeMounts(v.([]interface{}))
	}
	if v, ok := d.GetOk("env"); ok {
		req.Env = expandEnvVars(v.([]interface{}))
	}

	log.Printf("[DEBUG] buildRegisterInfrastructureV2Request request: %+v", req)

	return req, nil
}

// buildUpdateInfrastructureV2Request builds the request object for update operations
func buildUpdateInfrastructureV2Request(d *schema.ResourceData, accountID string) (*chaos.InfraV2UpdateKubernetesInfrastructureV2Request, error) {
	// In resourceInfrastructureV2Create and resourceInfrastructureV2Update functions
	name := d.Get("name").(string)
	sanitizedName := sanitizeK8sResourceName(name)

	// If the sanitized name is different, log a warning
	if name != sanitizedName {
		log.Printf("[WARN] Infrastructure name '%s' has been sanitized to '%s' to comply with Kubernetes naming requirements", name, sanitizedName)
	}

	req := &chaos.InfraV2UpdateKubernetesInfrastructureV2Request{
		Name:               sanitizedName,
		Identity:           d.Get("infra_id").(string), // set identity from infra_id
		Description:        d.Get("description").(string),
		ServiceAccount:     d.Get("service_account").(string),
		InfraNamespace:     d.Get("namespace").(string),
		EnvironmentID:      d.Get("environment_id").(string),
		AiEnabled:          d.Get("ai_enabled").(bool),
		InsecureSkipVerify: d.Get("insecure_skip_verify").(bool),
		Containers:         d.Get("containers").(string),
		RunAsUser:          int32(d.Get("run_as_user").(int)),
		RunAsGroup:         int32(d.Get("run_as_group").(int)),
	}

	// Set maps
	if v, ok := d.GetOk("label"); ok {
		req.Label = expandStringMap(v.(map[string]interface{}))
	}
	if v, ok := d.GetOk("annotation"); ok {
		req.Annotation = expandStringMap(v.(map[string]interface{}))
	}
	if v, ok := d.GetOk("node_selector"); ok {
		req.NodeSelector = expandStringMap(v.(map[string]interface{}))
	}
	if v, ok := d.GetOk("tags"); ok {
		req.Tags = expandStringList(v.([]interface{}))
	}

	// Set nested objects
	log.Printf("[DEBUG] buildUpdateInfrastructureV2Request image_registry: %+v",
		d.Get("image_registry"))
	if v, ok := d.GetOk("image_registry"); ok {
		req.ImageRegistry = expandImageRegistry(v.([]interface{}), d, accountID)
	}
	if v, ok := d.GetOk("mtls"); ok {
		req.Mtls = expandMtls(v.([]interface{}))
	}
	if v, ok := d.GetOk("proxy"); ok {
		req.Proxy = expandProxy(v.([]interface{}))
	}
	if v, ok := d.GetOk("tolerations"); ok {
		req.Tolerations = expandTolerations(v.([]interface{}))
	}
	if v, ok := d.GetOk("volumes"); ok {
		req.Volumes = expandVolumes(v.([]interface{}))
	}
	if v, ok := d.GetOk("volume_mounts"); ok {
		req.VolumeMounts = expandVolumeMounts(v.([]interface{}))
	}
	if v, ok := d.GetOk("env"); ok {
		req.Env = expandEnvVars(v.([]interface{}))
	}

	log.Printf("[DEBUG] buildUpdateInfrastructureV2Request request: %+v, %+v", req, req.ImageRegistry)

	return req, nil
}

// expandStringList converts a Terraform schema list of strings to a Go string slice
func expandStringList(in []interface{}) []string {
	if len(in) == 0 {
		return nil
	}

	result := make([]string, len(in))
	for i, v := range in {
		result[i] = v.(string)
	}
	return result
}

// Helper function to expand string maps
func expandStringMap(m map[string]interface{}) map[string]string {
	result := make(map[string]string, len(m))
	for k, v := range m {
		if v == nil {
			continue
		}
		if s, ok := v.(string); ok {
			result[k] = s
		}
	}
	return result
}

func expandImageRegistry(in []interface{}, d *schema.ResourceData, accountID string) *chaos.ImageRegistryImageRegistryV2 {
	if len(in) == 0 || in[0] == nil {
		return nil
	}

	m, ok := in[0].(map[string]interface{})
	if !ok {
		return nil
	}

	reg := &chaos.ImageRegistryImageRegistryV2{
		RegistryServer:    getString(m, "registry_server"),
		RegistryAccount:   getString(m, "registry_account"),
		IsPrivate:         getBool(m, "is_private"),
		SecretName:        getString(m, "secret_name"),
		UseCustomImages:   getBool(m, "use_custom_images"),
		IsDefault:         getBool(m, "is_default"),
		IsOverrideAllowed: getBool(m, "is_override_allowed"),
	}

	// Always use the infrastructure's infra_id, format is env_id/infra_id
	if envID := d.Get("environment_id").(string); envID != "" {
		if infraID := d.Get("infra_id").(string); infraID != "" {
			ensureRegistryInfraID(reg, envID, infraID)
		} else {
			log.Printf("[WARN] image_registry not configured, infra_id is not set in the configuration")
		}
	} else {
		log.Printf("[WARN] image_registry not configured, environment_id is not set in the configuration")
	}

	// Handle custom images
	if v, ok := m["custom_images"].([]interface{}); ok && len(v) > 0 {
		reg.CustomImages = expandCustomImages(v)
	}

	// Handle identifier - use the one from registry config or create from infrastructure details
	if v, ok := m["identifier"].([]interface{}); ok && len(v) > 0 {
		if idMap, ok := v[0].(map[string]interface{}); ok {
			reg.Identifier = &chaos.GithubComHarnessHceSaasHceSdkTypesApiK8sifsImageRegistryScopedIdentifiers{
				AccountIdentifier: getString(idMap, "account_identifier"),
				OrgIdentifier:     getString(idMap, "org_identifier"),
				ProjectIdentifier: getString(idMap, "project_identifier"),
			}
		}
	} else {
		// Create identifier from infrastructure details if not set in the registry config
		orgID := d.Get("org_id").(string)
		projectID := d.Get("project_id").(string)

		if accountID != "" {
			reg.Identifier = &chaos.GithubComHarnessHceSaasHceSdkTypesApiK8sifsImageRegistryScopedIdentifiers{
				AccountIdentifier: accountID,
				OrgIdentifier:     orgID,
				ProjectIdentifier: projectID,
			}
		}
	}

	return reg
}

// Helper functions for safe type assertions
func getString(m map[string]interface{}, key string) string {
	if v, ok := m[key]; ok && v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

func getBool(m map[string]interface{}, key string) bool {
	if v, ok := m[key]; ok && v != nil {
		if b, ok := v.(bool); ok {
			return b
		}
	}
	return false
}

func expandIdentifier(in []interface{}) *chaos.InfraV2Identifiers {
	if len(in) == 0 || in[0] == nil {
		return nil
	}

	m := in[0].(map[string]interface{})
	return &chaos.InfraV2Identifiers{
		AccountIdentifier: m["account_identifier"].(string),
		OrgIdentifier:     m["org_identifier"].(string),
		ProjectIdentifier: m["project_identifier"].(string),
	}
}

func expandCustomImages(in []interface{}) *chaos.ImageRegistryCustomImagesRequest {
	if len(in) == 0 || in[0] == nil {
		return nil
	}

	m := in[0].(map[string]interface{})
	return &chaos.ImageRegistryCustomImagesRequest{
		Ddcr:       m["ddcr"].(string),
		DdcrFault:  m["ddcr_fault"].(string),
		DdcrLib:    m["ddcr_lib"].(string),
		LogWatcher: m["log_watcher"].(string),
	}
}

func expandMtls(in []interface{}) *chaos.InfraV2MtlsConfiguration {
	if len(in) == 0 || in[0] == nil {
		return nil
	}

	m := in[0].(map[string]interface{})
	return &chaos.InfraV2MtlsConfiguration{
		CertPath:   m["cert_path"].(string),
		KeyPath:    m["key_path"].(string),
		SecretName: m["secret_name"].(string),
		Url:        m["url"].(string),
	}
}

func expandProxy(in []interface{}) *chaos.InfraV2ProxyConfiguration {
	if len(in) == 0 || in[0] == nil {
		return nil
	}

	m := in[0].(map[string]interface{})
	proxy := &chaos.InfraV2ProxyConfiguration{
		HttpProxy:  m["http_proxy"].(string),
		HttpsProxy: m["https_proxy"].(string),
		NoProxy:    m["no_proxy"].(string),
		Url:        m["url"].(string),
	}

	return proxy
}

func expandTolerations(in []interface{}) []chaos.V1Toleration {
	if len(in) == 0 {
		return nil
	}

	result := make([]chaos.V1Toleration, len(in))
	for i, v := range in {
		m := v.(map[string]interface{})
		toleration := chaos.V1Toleration{
			Key:   m["key"].(string),
			Value: m["value"].(string),
		}

		// Handle Operator
		if op, ok := m["operator"].(string); ok && op != "" {
			opValue := chaos.AllOfv1TolerationOperator(op)
			toleration.Operator = &opValue
		}

		// Handle Effect
		if effect, ok := m["effect"].(string); ok && effect != "" {
			effectValue := chaos.AllOfv1TolerationEffect(effect)
			toleration.Effect = &effectValue
		}

		// Handle TolerationSeconds if provided
		if ts, ok := m["toleration_seconds"].(int); ok {
			toleration.TolerationSeconds = int32(ts)
		}

		result[i] = toleration
	}

	return result
}

func expandVolumes(in []interface{}) []chaos.InfraV2Volumes {
	if len(in) == 0 {
		return nil
	}

	result := make([]chaos.InfraV2Volumes, len(in))
	for i, v := range in {
		m := v.(map[string]interface{})
		result[i] = chaos.InfraV2Volumes{
			Name:      m["name"].(string),
			SizeLimit: m["size_limit"].(string),
		}
	}

	return result
}

func expandVolumeMounts(in []interface{}) []chaos.V1VolumeMount {
	if len(in) == 0 {
		return nil
	}

	result := make([]chaos.V1VolumeMount, len(in))
	for i, v := range in {
		m := v.(map[string]interface{})
		mount := chaos.V1VolumeMount{
			MountPath:   m["mount_path"].(string),
			Name:        m["name"].(string),
			ReadOnly:    m["read_only"].(bool),
			SubPath:     m["sub_path"].(string),
			SubPathExpr: m["sub_path_expr"].(string),
		}

		// Handle MountPropagation if provided
		if prop, ok := m["mount_propagation"].(string); ok && prop != "" {
			propValue := chaos.AllOfv1VolumeMountMountPropagation(prop)
			mount.MountPropagation = &propValue
		}

		result[i] = mount
	}

	return result
}

func expandEnvVars(in []interface{}) []chaos.InfraV2Env {
	if len(in) == 0 {
		return nil
	}

	result := make([]chaos.InfraV2Env, len(in))
	for i, v := range in {
		m := v.(map[string]interface{})
		env := chaos.InfraV2Env{
			Name:  m["name"].(string),
			Value: m["value"].(string),
		}

		if v, ok := m["value_from"].(string); ok && v != "" {
			valueFrom := chaos.InfraV2EnvValueFrom(v)
			env.ValueFrom = &valueFrom
			if k, ok := m["key"].(string); ok && k != "" {
				env.Key = k
			}
		}

		result[i] = env
	}

	return result
}

// ImportStateContext handles importing an existing infrastructure
func resourceInfrastructureV2Import(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	// The import ID is expected to be in the format: "org_id/project_id/environment_id/infra_id"
	parts := strings.Split(d.Id(), "/")
	if len(parts) != 4 {
		return nil, fmt.Errorf("invalid import ID format. Expected: org_id/project_id/environment_id/infra_id, got: %s", d.Id())
	}

	orgID := parts[0]
	projectID := parts[1]
	environmentID := parts[2]
	infraID := parts[3]

	log.Printf("[DEBUG] Importing infrastructure with ID: %s/%s/%s/%s", orgID, projectID, environmentID, infraID)

	// Set the ID in the format that our Read function expects
	d.SetId(fmt.Sprintf("%s/%s/%s/%s", orgID, projectID, environmentID, infraID))

	// Set the individual ID fields
	if err := d.Set("org_id", orgID); err != nil {
		return nil, fmt.Errorf("failed to set org_id: %v", err)
	}
	if err := d.Set("project_id", projectID); err != nil {
		return nil, fmt.Errorf("failed to set project_id: %v", err)
	}
	if err := d.Set("environment_id", environmentID); err != nil {
		return nil, fmt.Errorf("failed to set environment_id: %v", err)
	}
	if err := d.Set("infra_id", infraID); err != nil {
		return nil, fmt.Errorf("failed to set infra_id: %v", err)
	}
	if err := d.Set("identity", infraID); err != nil {
		return nil, fmt.Errorf("failed to set identity: %v", err)
	}

	// Get the client
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)

	// Verify the infrastructure exists
	log.Printf("[DEBUG] Verifying infrastructure exists with params: account_id=%s, org_id=%s, project_id=%s, environment_id=%s, infra_id=%s",
		c.AccountId, orgID, projectID, environmentID, infraID)

	// Log the request details for debugging
	log.Printf("[DEBUG] Request details for GetInfraV2:")
	log.Printf("[DEBUG]   - infraID: %s", infraID)
	log.Printf("[DEBUG]   - accountID: %s", c.AccountId)
	log.Printf("[DEBUG]   - orgID: %s", orgID)
	log.Printf("[DEBUG]   - projectID: %s", projectID)
	log.Printf("[DEBUG]   - environmentID: %s", environmentID)

	// First, try to get the infrastructure
	infra, httpResp, err := c.ChaosSdkApi.GetInfraV2(
		ctx,
		infraID,
		c.AccountId,
		orgID,
		projectID,
		environmentID,
	)

	if err != nil {
		// Log detailed error information
		errMsg := fmt.Sprintf("Error getting infrastructure: %v", err)
		log.Printf("[ERROR] %s", errMsg)

		// If we have an HTTP response, log its details
		if httpResp != nil {
			log.Printf("[ERROR] Response status: %s", httpResp.Status)
			log.Printf("[ERROR] Response headers: %v", httpResp.Header)

			// Try to read the response body for more details
			if httpResp.Body != nil {
				defer httpResp.Body.Close()
				body, readErr := io.ReadAll(httpResp.Body)
				if readErr != nil {
					log.Printf("[ERROR] Failed to read response body: %v", readErr)
				} else {
					log.Printf("[ERROR] Response body: %s", string(body))
					errMsg = fmt.Sprintf("%s\nResponse body: %s", errMsg, string(body))
				}
			}

			// Handle specific status codes
			switch httpResp.StatusCode {
			case 401:
				errMsg = fmt.Sprintf("authentication failed - please check your API key and account ID\n%s", errMsg)
			case 403:
				errMsg = fmt.Sprintf("permission denied - your API key doesn't have access to this resource\n%s", errMsg)
			case 404:
				return nil, fmt.Errorf("infrastructure not found with ID: %s (account_id: %s, org_id: %s, project_id: %s, environment_id: %s)",
					infraID, c.AccountId, orgID, projectID, environmentID)
			case 500:
				errMsg = fmt.Sprintf("internal server error - please check the Harness API status\n%s", errMsg)
			}

			return nil, fmt.Errorf("%s (status: %d)", errMsg, httpResp.StatusCode)
		}

		// If we get here, we don't have an HTTP response
		return nil, fmt.Errorf("%s (no response from server)", errMsg)
	}

	log.Printf("[DEBUG] Found infrastructure: %+v", infra)

	// Call the read function to populate the rest of the state
	log.Printf("[DEBUG] Calling read function to populate state...")
	diags := resourceInfrastructureV2Read(ctx, d, meta)
	if diags.HasError() {
		return nil, fmt.Errorf("error reading infrastructure: %v", diags)
	}

	log.Printf("[DEBUG] Successfully imported infrastructure: %s", d.Id())
	return []*schema.ResourceData{d}, nil
}
