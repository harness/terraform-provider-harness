package agent

import (
	"context"
	"time"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/svcdiscovery"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/harness/terraform-provider-harness/internal/service/service_discovery/agent/convert"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// DataSourceServiceDiscoveryAgent returns the data source for a Harness Service Discovery Agent
func DataSourceServiceDiscoveryAgent() *schema.Resource {
	return &schema.Resource{
		Description: `Data source for retrieving a Harness Service Discovery Agent.

This data source allows you to fetch details of a Service Discovery Agent using either its unique identifier or name.`,

		ReadContext: dataSourceServiceDiscoveryAgentRead,

		Schema: AgentDataSourceSchema(),
	}
}

func dataSourceServiceDiscoveryAgentRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*internal.Session)
	client := c.SDClient

	// Validate required fields
	if v, ok := d.GetOk("account_identifier"); !ok || v.(string) == "" {
		return diag.Errorf("account_identifier is required and cannot be empty")
	}

	if v, ok := d.GetOk("environment_identifier"); !ok || v.(string) == "" {
		return diag.Errorf("environment_identifier is required and cannot be empty")
	}

	accountID := d.Get("account_identifier").(string)
	envID := d.Get("environment_identifier").(string)

	// listAll := d.Get("list_all").(bool)
	// // Handle list_all case
	// if listAll {
	// 	// Get all agents with enhanced pagination
	// 	allAgents, diags := listAllAgents(
	// 		ctx,
	// 		client,
	// 		accountID,
	// 		envID,
	// 		d.Get("search").(string),
	// 		d.Get("org_identifier").(string),
	// 		d.Get("project_identifier").(string),
	// 		d.Get("max_pages").(int),
	// 		d.Get("page_size").(int),
	// 		d.Get("timeout_seconds").(int),
	// 	)
	// 	if diags.HasError() {
	// 		return diags
	// 	}

	// 	// Set the agents list
	// 	if err := d.Set("agents", allAgents); err != nil {
	// 		return diag.Errorf("failed to set agents: %v", err)
	// 	}

	// 	d.SetId(fmt.Sprintf("%s/%s/agents", accountID, envID))
	// 	return nil
	// }

	// Handle single agent lookup
	return getSingleAgent(ctx, d, client, accountID, envID)
}

// getSingleAgent retrieves a single agent by ID or name with proper pagination
func getSingleAgent(ctx context.Context, d *schema.ResourceData, client *svcdiscovery.APIClient, accountID, envID string) diag.Diagnostics {
	// Get timeout from context or use default
	timeout := 30 * time.Second
	if deadline, ok := ctx.Deadline(); ok {
		timeout = time.Until(deadline)
	}

	if identity, ok := d.Get("identity").(string); ok && identity != "" {
		// Lookup by identity (direct lookup, no pagination needed)
		agent, _, err := client.AgentApi.GetAgent(
			ctx,
			identity,
			accountID,
			envID,
			&svcdiscovery.AgentApiGetAgentOpts{
				OrganizationIdentifier: optional.NewString(d.Get("org_identifier").(string)),
				ProjectIdentifier:      optional.NewString(d.Get("project_identifier").(string)),
			},
		)
		if err != nil {
			return diag.Errorf("failed to get agent: %v", err)
		}

		// Then set it in the schema
		if err := convert.FlattenAgentToSchema(d, &agent); err != nil {
			return diag.Errorf("failed to set agent data: %v", err)
		}

		d.SetId(agent.Id)
		return nil
	}

	// Lookup by name with pagination
	name := d.Get("name").(string)
	if name == "" {
		return diag.Errorf("either 'identity' or 'name' must be provided")
	}

	// Create a context with timeout for the search operation
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	var foundAgent *svcdiscovery.ApiGetAgentResponse
	page := int32(0)
	pageSize := int32(50) // Smaller page size for more responsive search

	searchOpts := &svcdiscovery.AgentApiListAgentOpts{
		Search:                 optional.NewString(name),
		OrganizationIdentifier: optional.NewString(d.Get("org_identifier").(string)),
		ProjectIdentifier:      optional.NewString(d.Get("project_identifier").(string)),
	}

	for {
		select {
		case <-ctx.Done():
			if ctx.Err() == context.DeadlineExceeded {
				return diag.Errorf("timeout while searching for agent with name '%s'", name)
			}
			return diag.FromErr(ctx.Err())
		default:
		}

		// Search for agents with the given name
		response, _, err := client.AgentApi.ListAgent(
			ctx,
			accountID,
			envID,
			page,
			pageSize,
			false,
			searchOpts,
		)
		if err != nil {
			return diag.Errorf("failed to search for agents: %v", err)
		}

		// Check for exact name match in the current page
		for i, agent := range response.Items {
			if agent.Name == name {
				foundAgent = &response.Items[i]
				break
			}
		}

		// If we found the agent or reached the end of results, break the loop
		if foundAgent != nil || len(response.Items) == 0 || len(response.Items) < int(pageSize) {
			break
		}

		page++
	}

	if foundAgent == nil {
		return diag.Errorf("no agent found with name '%s'", name)
	}

	// First, get the flattened agent data
	if err := convert.FlattenAgentToSchema(d, foundAgent); err != nil {
		return diag.Errorf("failed to flatten agent: %v", err)
	}

	// // Then set it in the schema
	// if err := d.Set("agent", []interface{}{foundAgent}); err != nil {
	// 	return diag.Errorf("failed to set agent data: %v", err)
	// }

	d.SetId(foundAgent.Id)
	return nil
}
