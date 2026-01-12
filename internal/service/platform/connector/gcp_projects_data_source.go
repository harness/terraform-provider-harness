package connector

import (
	"context"
	"fmt"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceGcpProjects() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for listing GCP projects using a cloud connector identifier or secret manager connector identifier.",

		ReadContext: dataSourceGcpProjectsRead,

		Schema: map[string]*schema.Schema{
			"connector_id": {
				Description: "Identifier of the GCP cloud connector or secret manager connector.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"projects": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}

	helpers.SetOptionalOrgAndProjectLevelDataSourceSchema(resource.Schema)

	return resource
}

func dataSourceGcpProjectsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	// Get connector identifier
	connectorRef := d.Get("connector_id").(string)

	// Build optional parameters for org and project scope
	opts := &nextgen.GcpProjectsApiGetGcpProjectsOpts{
		OrgIdentifier:     helpers.BuildField(d, "org_id"),
		ProjectIdentifier: helpers.BuildField(d, "project_id"),
	}

	// Call the API to get GCP projects
	// Parameters: routingId, accountIdentifier, connectorRef (path), opts (optional query params)
	resp, httpResp, err := c.GcpProjectsApi.GetGcpProjects(ctx, c.AccountId, c.AccountId, connectorRef, opts)
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	// Check if response data is nil
	if resp.Data == nil || resp.Data.Projects == nil {
		d.SetId("")
		return nil
	}

	// Convert projects to schema format
	var projects []map[string]interface{}
	for _, project := range resp.Data.Projects {
		newProject := map[string]interface{}{
			"id":   project.Id,
			"name": project.Name,
		}
		projects = append(projects, newProject)
	}

	// Set the ID and projects
	d.SetId(fmt.Sprintf("%s-%s", connectorRef, resp.CorrelationId))
	d.Set("projects", projects)

	return nil
}
