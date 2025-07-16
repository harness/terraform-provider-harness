package agent

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/antihax/optional"
	"github.com/google/uuid"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/harness/terraform-provider-harness/internal/service/service_discovery/agent/convert"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	svcdiscovery "github.com/harness/harness-go-sdk/harness/svcdiscovery"
)

// ResourceServiceDiscoveryAgent returns a new Terraform resource for Harness Service Discovery Agent
func ResourceServiceDiscoveryAgent() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for managing a Harness Service Discovery Agent.\n\n" +
			"This resource allows you to create, read, update, and delete a Service Discovery Agent in Harness.\n\n" +
			"## Example Usage\n\n" +
			"```hcl\n" +
			`resource "harness_platform_service_discovery_agent" "example" {
  identifier             = "example_agent"
  name                   = "Example Agent"
  description            = "Example Service Discovery Agent"
  org_identifier         = "your_org_id"
  project_identifier     = "your_project_id"
  environment_identifier = "your_environment_id"
  permanent_installation = false
  
  config {
    collector_image    = "harness/service-discovery-collector:main-latest"
    log_watcher_image  = "harness/chaos-log-watcher:main-latest"
    skip_secure_verify = false
    
    kubernetes {
      resources {
        limits {
          cpu    = "500m"
          memory = "512Mi"
        }
        requests {
          cpu    = "250m"
          memory = "512Mi"
        }
      }
    }
  }
}` + "\n```",

		CreateContext: resourceServiceDiscoveryAgentCreate,
		ReadContext:   resourceServiceDiscoveryAgentRead,
		UpdateContext: resourceServiceDiscoveryAgentUpdate,
		DeleteContext: resourceServiceDiscoveryAgentDelete,
		// Importer: &schema.ResourceImporter{
		// 	StateContext: resourceServiceDiscoveryAgentImport,
		// },
		Importer: MultiLevelResourceImporter,
		Schema:   AgentResourceSchema(),
	}
}

var MultiLevelResourceImporter = &schema.ResourceImporter{
	State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
		parts := strings.Split(d.Id(), "/")
		partCount := len(parts)

		// Handle format: org_id/project_id/environment_identifier/infra_identifier
		if partCount == 4 {
			d.SetId(parts[3]) // infra_identifier
			d.Set("org_identifier", parts[0])
			d.Set("project_identifier", parts[1])
			d.Set("environment_identifier", parts[2])
			d.Set("infra_identifier", parts[3])
			return []*schema.ResourceData{d}, nil
		}

		// Original handling for other formats
		isAccountConnector := partCount == 2
		isOrgConnector := partCount == 3

		if isAccountConnector {
			d.SetId(parts[1])
			d.Set("environment_identifier", parts[0])
			d.Set("infra_identifier", parts[1])
			return []*schema.ResourceData{d}, nil
		}

		if isOrgConnector {
			d.SetId(parts[2])
			d.Set("org_identifier", parts[0])
			d.Set("environment_identifier", parts[1])
			d.Set("infra_identifier", parts[2])
			return []*schema.ResourceData{d}, nil
		}

		return nil, fmt.Errorf("invalid import format. Expected one of:\n" +
			"- <environment_identifier>/<infra_identifier> (account level)\n" +
			"- <org_id>/<environment_identifier>/<infra_identifier> (org level)\n" +
			"- <org_id>/<project_id>/<environment_identifier>/<infra_identifier> (project level)")
	},
}

// validateAgentInput validates the required fields for agent operations
func validateAgentInput(d *schema.ResourceData) diag.Diagnostics {
	var diags diag.Diagnostics

	if v, ok := d.Get("environment_identifier").(string); !ok || v == "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Missing required argument",
			Detail:   "environment_identifier is required and cannot be empty",
		})
	}

	if v, ok := d.Get("infra_identifier").(string); !ok || v == "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Missing required argument",
			Detail:   "infra_identifier is required and cannot be empty",
		})
	}

	if v, ok := d.Get("name").(string); !ok || v == "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Missing required argument",
			Detail:   "name is required and cannot be empty",
		})
	}

	// Validate configuration if provided
	if v, ok := d.GetOk("config"); ok {
		configs := v.([]interface{})
		if len(configs) > 0 {
			if config, ok := configs[0].(map[string]interface{}); ok {
				if _, ok := config["kubernetes"].([]interface{})[0].(map[string]interface{}); !ok {
					diags = append(diags, diag.Diagnostic{
						Severity: diag.Error,
						Summary:  "Invalid configuration",
						Detail:   "config.kubernetes is required when config block is specified",
					})
				}
				if _, ok := config["kubernetes"].([]interface{})[0].(map[string]interface{})["namespace"].(string); !ok {
					diags = append(diags, diag.Diagnostic{
						Severity: diag.Error,
						Summary:  "Invalid configuration",
						Detail:   "config.kubernetes.namespace is required when config block is specified",
					})
				}
			}
		}
	}

	return diags
}

func resourceServiceDiscoveryAgentCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Validate input before making any API calls
	if diags := validateAgentInput(d); diags.HasError() {
		return diags
	}

	c := meta.(*internal.Session).SDClient
	if c == nil {
		return diag.Errorf("service discovery client is not properly configured")
	}

	// Extract parameters from the resource data
	accountIdentifier := c.AccountId
	environmentIdentifier := d.Get("environment_identifier").(string)
	infraIdentifier := d.Get("infra_identifier").(string)

	// Apply defaults
	if err := setDataDefaults(d); err != nil {
		return diag.FromErr(err)
	}

	// Prepare the create request
	req := svcdiscovery.ApiCreateAgentRequest{
		EnvironmentIdentifier: environmentIdentifier,
		InfraIdentifier:       infraIdentifier,
		Name:                  d.Get("name").(string),
	}

	// Set permanent installation flag if provided
	if v, ok := d.GetOk("permanent_installation"); ok {
		req.PermanentInstallation = v.(bool)
	}

	// Set webhook URL if provided
	if v, ok := d.GetOk("webhook_url"); ok {
		req.WebhookURL = v.(string)
	}

	// Expand the configuration if provided
	if v, ok := d.GetOk("config"); ok {
		config, err := convert.ExpandAgentConfig(v.([]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		req.Config = config
	}

	// Set up API options
	opts := &svcdiscovery.AgentApiCreateAgentOpts{
		CorrelationID: optional.NewString(uuid.New().String()),
	}

	// Set organization and project identifiers if provided
	if v, ok := d.GetOk("org_identifier"); ok {
		opts.OrganizationIdentifier = optional.NewString(v.(string))
	}

	if v, ok := d.GetOk("project_identifier"); ok {
		opts.ProjectIdentifier = optional.NewString(v.(string))
	}

	// Create the agent
	agent, httpResp, err := c.AgentApi.CreateAgent(ctx, req, accountIdentifier, opts)
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	// Set the ID and read back the agent
	d.SetId(agent.Id)
	return resourceServiceDiscoveryAgentRead(ctx, d, meta)
}

func resourceServiceDiscoveryAgentRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*internal.Session).SDClient
	if c == nil {
		return diag.Errorf("service discovery client is not configured")
	}

	// Set up API options
	opts := &svcdiscovery.AgentApiGetAgentOpts{
		CorrelationID: optional.NewString(uuid.New().String()),
	}

	// Set account identifier
	accountIdentifier := c.AccountId
	environmentIdentifier := d.Get("environment_identifier").(string)
	agentIdentity := d.Get("infra_identifier").(string)

	// Set organization and project identifiers if provided
	if v, ok := d.GetOk("org_identifier"); ok {
		opts.OrganizationIdentifier = optional.NewString(v.(string))
	}

	if v, ok := d.GetOk("project_identifier"); ok {
		opts.ProjectIdentifier = optional.NewString(v.(string))
	}

	// Read the agent
	agent, httpResp, err := c.AgentApi.GetAgent(
		ctx,
		agentIdentity,
		accountIdentifier,
		environmentIdentifier,
		opts,
	)

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	// Set the basic fields
	d.Set("name", agent.Name)
	d.Set("description", agent.Description)
	d.Set("identity", agent.Identity)
	d.Set("installation_type", agent.InstallationType)
	d.Set("permanent_installation", agent.PermanentInstallation)
	d.Set("service_count", agent.ServiceCount)
	d.Set("network_map_count", agent.NetworkMapCount)
	d.Set("webhook_url", agent.WebhookURL)
	d.Set("correlation_id", agent.CorrelationID)

	// Handle tags
	if len(agent.Tags) > 0 {
		tags := make(map[string]string)
		for _, tag := range agent.Tags {
			tags[tag] = tag // Using the same key-value pair for simplicity
		}
		if err := d.Set("tags", tags); err != nil {
			return diag.FromErr(fmt.Errorf("error setting tags: %w", err))
		}
	} else {
		d.Set("tags", nil)
	}

	// Handle timestamps
	if agent.CreatedAt != "" {
		if err := d.Set("created_at", agent.CreatedAt); err != nil {
			return diag.FromErr(fmt.Errorf("error setting created_at: %w", err))
		}
	}

	if agent.UpdatedAt != "" {
		if err := d.Set("updated_at", agent.UpdatedAt); err != nil {
			return diag.FromErr(fmt.Errorf("error setting updated_at: %w", err))
		}
	}

	// Handle created_by and updated_by
	if agent.CreatedBy != "" {
		if err := d.Set("created_by", agent.CreatedBy); err != nil {
			return diag.FromErr(fmt.Errorf("error setting created_by: %w", err))
		}
	}

	if agent.UpdatedBy != "" {
		if err := d.Set("updated_by", agent.UpdatedBy); err != nil {
			return diag.FromErr(fmt.Errorf("error setting updated_by: %w", err))
		}
	}

	if agent.RemovedAt != "" {
		if err := d.Set("removed_at", agent.RemovedAt); err != nil {
			return diag.FromErr(fmt.Errorf("error setting removed_at: %w", err))
		}
	}

	// Handle removed flag
	if agent.Removed {
		if err := d.Set("removed", agent.Removed); err != nil {
			return diag.FromErr(fmt.Errorf("error setting removed: %w", err))
		}
	}

	// Handle configuration
	if agent.Config != nil {
		configState, err := convert.FlattenAgentConfig(agent.Config)
		if err != nil {
			return diag.FromErr(fmt.Errorf("flattening agent config: %w", err))
		}
		d.Set("config", configState)
	}

	// Handle installation details
	if agent.InstallationDetails != nil {
		log.Printf("[DEBUG] Raw installation details: %+v", agent.InstallationDetails)
		installationDetails, err := convert.FlattenInstallationDetails(agent.InstallationDetails)
		if err != nil {
			return diag.FromErr(fmt.Errorf("failed to flatten installation details: %w", err))
		}
		log.Printf("[DEBUG] Flattened installation details: %+v", installationDetails)

		// Only set if we have data
		if len(installationDetails) > 0 {
			if err := d.Set("installation_details", []interface{}{installationDetails}); err != nil {
				return diag.FromErr(fmt.Errorf("error setting installation_details: %w", err))
			}
		} else {
			d.Set("installation_details", nil)
		}
	} else {
		d.Set("installation_details", nil)
	}

	return nil
}

func resourceServiceDiscoveryAgentUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Validate input before making any API calls
	if diags := validateAgentInput(d); diags.HasError() {
		return diags
	}

	c := meta.(*internal.Session).SDClient
	if c == nil {
		return diag.Errorf("service discovery client is not properly configured")
	}

	// Extract parameters from the resource data
	accountIdentifier := c.AccountId
	environmentIdentifier := d.Get("environment_identifier").(string)
	agentIdentity := d.Get("infra_identifier").(string)

	// Apply defaults to config
	if err := setDataDefaults(d); err != nil {
		return diag.FromErr(err)
	}

	// Check if any of the immutable fields have changed
	if d.HasChanges("org_identifier", "project_identifier", "environment_identifier") {
		return diag.Errorf("cannot update immutable fields: org_identifier, project_identifier, environment_identifier")
	}

	// Read the current agent state to get the current config
	readOpts := &svcdiscovery.AgentApiGetAgentOpts{
		CorrelationID: optional.NewString(uuid.New().String()),
	}

	if v, ok := d.GetOk("org_identifier"); ok {
		readOpts.OrganizationIdentifier = optional.NewString(v.(string))
	}

	if v, ok := d.GetOk("project_identifier"); ok {
		readOpts.ProjectIdentifier = optional.NewString(v.(string))
	}

	currentAgent, httpResp, err := c.AgentApi.GetAgent(
		ctx,
		agentIdentity,
		accountIdentifier,
		environmentIdentifier,
		readOpts,
	)
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	// Prepare the update request with the current config
	req := svcdiscovery.ApiUpdateAgentRequest{
		Name:   d.Get("name").(string),
		Config: currentAgent.Config,
	}

	// If config is explicitly provided in the update, use it
	if d.HasChange("config") {
		if v, ok := d.GetOk("config"); ok && len(v.([]interface{})) > 0 {
			log.Printf("[DEBUG] Config provided in update: %+v", v)
			config, err := convert.ExpandAgentConfig(v.([]interface{}))
			if err != nil {
				return diag.FromErr(err)
			}
			req.Config = config
		}
	}

	// Set up API options
	opts := &svcdiscovery.AgentApiUpdateAgentOpts{
		CorrelationID: optional.NewString(uuid.New().String()),
	}

	// Set organization and project identifiers if provided
	if v, ok := d.GetOk("org_identifier"); ok {
		opts.OrganizationIdentifier = optional.NewString(v.(string))
	}

	if v, ok := d.GetOk("project_identifier"); ok {
		opts.ProjectIdentifier = optional.NewString(v.(string))
	}

	log.Printf("[DEBUG] Update agent request: %+v, opts: %+v, config: %+v", req, opts, req.Config)

	// Update the agent
	_, updateHttpResp, updateErr := c.AgentApi.UpdateAgent(ctx, req, accountIdentifier, environmentIdentifier,
		agentIdentity, opts)
	if updateErr != nil {
		return helpers.HandleApiError(updateErr, d, updateHttpResp)
	}

	// Read back the agent
	return resourceServiceDiscoveryAgentRead(ctx, d, meta)
}

func resourceServiceDiscoveryAgentDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*internal.Session).SDClient
	if c == nil {
		return diag.Errorf("service discovery client is not configured")
	}

	// Get the agent ID before it's removed from state
	agentIdentity := d.Get("infra_identifier").(string)
	accountID := c.AccountId
	envID := d.Get("environment_identifier").(string)

	// Set up API options for delete
	deleteOpts := &svcdiscovery.AgentApiDeleteAgentOpts{
		CorrelationID: optional.NewString(uuid.New().String()),
	}

	// Set organization and project identifiers if provided
	if v, ok := d.GetOk("org_identifier"); ok {
		deleteOpts.OrganizationIdentifier = optional.NewString(v.(string))
	}

	if v, ok := d.GetOk("project_identifier"); ok {
		deleteOpts.ProjectIdentifier = optional.NewString(v.(string))
	}

	// Set up API options for get
	getOpts := &svcdiscovery.AgentApiGetAgentOpts{
		CorrelationID: optional.NewString(""),
	}

	if v, ok := d.GetOk("org_identifier"); ok {
		getOpts.OrganizationIdentifier = optional.NewString(v.(string))
	}

	if v, ok := d.GetOk("project_identifier"); ok {
		getOpts.ProjectIdentifier = optional.NewString(v.(string))
	}

	// Try to get the agent first to check if it exists
	_, getHttpResp, getErr := c.AgentApi.GetAgent(ctx, agentIdentity, accountID, envID, getOpts)
	if getErr != nil {
		// If the agent is already deleted, we're done
		if getHttpResp != nil && getHttpResp.StatusCode == http.StatusNotFound {
			d.SetId("")
			return nil
		}
		return helpers.HandleApiError(getErr, d, getHttpResp)
	}

	// Delete the agent
	_, httpResp, err := c.AgentApi.DeleteAgent(
		ctx,
		agentIdentity,
		accountID,
		envID,
		deleteOpts,
	)

	if err != nil {
		// If the agent is already deleted, that's fine
		if httpResp != nil && httpResp.StatusCode == http.StatusNotFound {
			d.SetId("")
			return nil
		}
		return helpers.HandleApiError(err, d, httpResp)
	}

	// Clear the ID to remove from state
	d.SetId("")

	return nil
}
