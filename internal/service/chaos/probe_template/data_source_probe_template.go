package probe_template

import (
	"context"
	"log"

	"github.com/harness/harness-go-sdk/harness/chaos"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceProbeTemplate() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for retrieving a Harness Chaos Probe Template.",

		ReadContext: dataSourceProbeTemplateRead,

		Schema: dataSourceProbeTemplateSchema(),
	}
}

// dataSourceProbeTemplateSchema derives the data source schema from the
// resource schema. The data source read reuses setProbeTemplateDataSimplified
// (shared with the resource), which sets fields such as run_properties.retry,
// run_properties.verbosity, run_properties.initial_delay, variables.description,
// k8s_probe.group and the full apm_probe block. Deriving from the resource
// schema guarantees every one of those addresses is valid, preventing the
// "Invalid address to set" errors that occur when a hand-maintained data source
// schema drifts out of sync with the resource.
func dataSourceProbeTemplateSchema() map[string]*schema.Schema {
	dsSchema := resourceProbeTemplateSchema()

	// Lookup by identity (recommended) or name; both optional + computed.
	dsSchema["identity"].Required = false
	dsSchema["identity"].Optional = true
	dsSchema["identity"].Computed = true
	dsSchema["identity"].ForceNew = false

	dsSchema["name"].Required = false
	dsSchema["name"].Optional = true
	dsSchema["name"].Computed = true

	// hub_identity is required so we know which hub to search.
	dsSchema["hub_identity"].Required = true
	dsSchema["hub_identity"].Optional = false
	dsSchema["hub_identity"].ForceNew = false

	// type is read from the API, not supplied by the user.
	dsSchema["type"].Required = false
	dsSchema["type"].Optional = true
	dsSchema["type"].Computed = true

	return dsSchema
}

func dataSourceProbeTemplateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)

	accountID := c.AccountId
	if accountID == "" {
		return diag.Errorf("account ID must be configured in the provider")
	}

	orgID := d.Get("org_id").(string)
	projectID := d.Get("project_id").(string)
	hubIdentity := d.Get("hub_identity").(string)

	// If identity is provided, fetch directly
	if identity, ok := d.GetOk("identity"); ok {
		identityStr := identity.(string)
		log.Printf("[DEBUG] Fetching probe template by identity: %s", identityStr)

		resp, httpResp, apiErr := c.DefaultApi.GetProbeTemplate(ctx, accountID, orgID, projectID, hubIdentity, identityStr, nil)
		if apiErr != nil {
			return helpers.HandleChaosReadApiError(apiErr, d, httpResp)
		}

		if resp.Data == nil {
			return diag.Errorf("probe template not found: %s", identityStr)
		}

		// Set the ID
		d.SetId(generateID(accountID, orgID, projectID, hubIdentity, resp.Data.Identity))

		// Set all the data
		return setProbeTemplateDataSimplified(d, resp.Data, accountID, orgID, projectID, hubIdentity)
	} else if name, ok := d.GetOk("name"); ok {
		// If name is provided, list and filter
		nameStr := name.(string)
		log.Printf("[DEBUG] Fetching probe template by name: %s", nameStr)

		resp, httpResp, apiErr := c.DefaultApi.ListProbeTemplate(ctx, accountID, orgID, projectID, hubIdentity, 0, 100, nameStr, nil)
		if apiErr != nil {
			return helpers.HandleChaosReadApiError(apiErr, d, httpResp)
		}

		if resp.Data == nil || len(resp.Data) == 0 {
			return diag.Errorf("probe template not found with name: %s", nameStr)
		}

		// Find exact match
		var foundTemplate *chaos.GithubComHarnessHceSaasGraphqlServerPkgDatabaseMongodbChaosprobetemplateChaosProbeTemplate
		for i, t := range resp.Data {
			if t.Name == nameStr {
				foundTemplate = &resp.Data[i]
				break
			}
		}

		if foundTemplate == nil {
			return diag.Errorf("probe template not found with name: %s", nameStr)
		}

		// Set the ID
		d.SetId(generateID(accountID, orgID, projectID, hubIdentity, foundTemplate.Identity))

		// Set all the data
		return setProbeTemplateDataSimplified(d, foundTemplate, accountID, orgID, projectID, hubIdentity)
	} else {
		return diag.Errorf("either 'identity' or 'name' must be specified")
	}
}
