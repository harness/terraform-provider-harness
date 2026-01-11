package workspace

import (
	"context"
	"fmt"

	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// DataSourceWorkspaces lists workspaces using the List Workspaces API and exposes
// identifiers along with all relevant summary fields returned by that API.
func DataSourceWorkspaces() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for listing workspaces.",

		ReadContext: dataSourceWorkspacesRead,

		Schema: map[string]*schema.Schema{
			"org_id": {
				Description: "Organization Identifier",
				Type:        schema.TypeString,
				Required:    true,
			},
			"project_id": {
				Description: "Project Identifier",
				Type:        schema.TypeString,
				Required:    true,
			},
			"search_term": {
				Description: "Filter results by partial name match when listing workspaces.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"identifiers": {
				Description: "List of workspace identifiers matching the filters.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"workspaces": {
				Description: "List of workspaces matching the filters.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"identifier": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"org_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"project_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"created": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"updated": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceWorkspacesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	orgID := d.Get("org_id").(string)
	projectID := d.Get("project_id").(string)
	searchTerm := d.Get("search_term").(string)
	limit := int32(defaultLimit)
	if v, ok := d.GetOk("limit"); ok {
		limit = int32(v.(int))
	}

	workspaces, httpResp, err := findWorkspaces(ctx, orgID, projectID, c.AccountId, c.WorkspaceApi, searchTerm, limit)
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	var (
		identifiers []string
		items       []map[string]interface{}
	)

	for _, ws := range workspaces {
		identifiers = append(identifiers, ws.Identifier)
		items = append(items, map[string]interface{}{
			"account_id":  ws.Account,
			"identifier":  ws.Identifier,
			"name":        ws.Name,
			"org_id":      ws.Org,
			"project_id":  ws.Project,
			"description": ws.Description,
			"status":      ws.Status,
			"created":     ws.Created,
			"updated":     ws.Updated,
		})
	}

	// synthetic ID so repeated calls with same filters are stable
	d.SetId(fmt.Sprintf("%s/%s/%s", orgID, projectID, searchTerm))
	d.Set("identifiers", identifiers)
	d.Set("workspaces", items)

	return nil
}
