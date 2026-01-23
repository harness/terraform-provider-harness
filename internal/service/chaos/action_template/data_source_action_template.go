package action_template

import (
	"context"
	"log"

	"github.com/harness/harness-go-sdk/harness/chaos"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceActionTemplate() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for reading Harness Chaos Action Templates.",

		ReadContext: dataSourceActionTemplateRead,

		Schema: dataSourceActionTemplateSchema(),
	}
}

func dataSourceActionTemplateSchema() map[string]*schema.Schema {
	// Start with the resource schema
	dsSchema := resourceActionTemplateSchema()

	// Make identity and hub_identity optional for data source (can lookup by name)
	dsSchema["identity"].Required = false
	dsSchema["identity"].Optional = true
	dsSchema["identity"].Computed = true

	// Add name as optional (can lookup by name instead of identity)
	dsSchema["name"].Required = false
	dsSchema["name"].Optional = true
	dsSchema["name"].Computed = true

	// hub_identity is still required as we need to know which hub to search in
	dsSchema["hub_identity"].Required = true
	dsSchema["hub_identity"].Optional = false

	return dsSchema
}

func dataSourceActionTemplateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)

	accountID := c.AccountId
	if accountID == "" {
		return diag.Errorf("account ID must be configured in the provider")
	}

	orgID := d.Get("org_id").(string)
	projectID := d.Get("project_id").(string)
	hubIdentity := d.Get("hub_identity").(string)

	// Check if identity is provided
	if identity, ok := d.GetOk("identity"); ok {
		// Direct lookup by identity
		identityStr := identity.(string)
		log.Printf("[DEBUG] Looking up action template by identity: %s in hub: %s", identityStr, hubIdentity)

		resp, httpResp, err := c.DefaultApi.GetActionTemplate(ctx, accountID, orgID, projectID, hubIdentity, identityStr, nil)
		if err != nil {
			if httpResp != nil && httpResp.StatusCode == 404 {
				return diag.Errorf("action template with identity '%s' not found in hub '%s'", identityStr, hubIdentity)
			}
			return helpers.HandleChaosReadApiError(err, d, httpResp)
		}

		if resp.Data == nil {
			return diag.Errorf("action template data is nil for identity '%s'", identityStr)
		}

		d.SetId(generateID(accountID, orgID, projectID, hubIdentity, resp.Data.Identity))
		return setActionTemplateData(d, resp.Data, accountID, orgID, projectID, hubIdentity)
	}

	// Lookup by name
	if name, ok := d.GetOk("name"); ok {
		nameStr := name.(string)
		log.Printf("[DEBUG] Looking up action template by name: %s in hub: %s", nameStr, hubIdentity)

		// List action templates and find by name
		opts := &chaos.DefaultApiListActionTemplateOpts{}
		resp, httpResp, err := c.DefaultApi.ListActionTemplate(ctx, accountID, orgID, projectID, hubIdentity, 0, 100, nameStr, opts)
		if err != nil {
			return helpers.HandleChaosReadApiError(err, d, httpResp)
		}

		if httpResp != nil && httpResp.StatusCode != 200 {
			return diag.Errorf("failed to list action templates: HTTP %d", httpResp.StatusCode)
		}

		// Find the template with matching name
		var foundTemplate *chaos.ChaosactiontemplateChaosActionTemplate
		if resp.Data != nil {
			for _, template := range resp.Data {
				if template.Name == nameStr {
					foundTemplate = &template
					break
				}
			}
		}

		if foundTemplate == nil {
			return diag.Errorf("action template with name '%s' not found in hub '%s'", nameStr, hubIdentity)
		}

		d.SetId(generateID(accountID, orgID, projectID, hubIdentity, foundTemplate.Identity))
		return setActionTemplateData(d, foundTemplate, accountID, orgID, projectID, hubIdentity)
	}

	return diag.Errorf("either 'identity' or 'name' must be specified")
}
