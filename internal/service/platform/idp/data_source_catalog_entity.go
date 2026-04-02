package idp

import (
	"context"
	"fmt"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/idp"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func DataSourceCatalogEntity() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving Harness catalog entities.",
		ReadContext: dataSourceCatalogEntityRead,
		Schema: map[string]*schema.Schema{
			"identifier": helpers.GetIdentifierSchema(helpers.SchemaFlagTypes.Required),
			"kind": {
				Type:        schema.TypeString,
				Description: "Kind of the catalog entity",
				Required:    true,
				ForceNew:    true,
				ValidateFunc: validation.StringInSlice([]string{
					"EnvironmentBlueprint", "Environment", "Component", "Group", "User", "Workflow", "Resource", "System",
				}, true),
			},
			"org_id":     helpers.GetOrgIdSchema(helpers.SchemaFlagTypes.Optional),
			"project_id": helpers.GetProjectIdSchema(helpers.SchemaFlagTypes.Optional),
			"yaml": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "YAML definition of the catalog entity",
			},
			"git_details": {
				Description: "Contains Git Information for importing entities from Git",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"branch_name": {
							Description: "Name of the branch.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"file_path": {
							Description: "File path of the Entity in the repository.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"commit_message": {
							Description: "Commit message used for the merge commit.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"base_branch": {
							Description: "Name of the default branch (this checks out a new branch titled by branch_name).",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"connector_ref": {
							Description: "Identifier of the Harness Connector used for importing entity from Git" + helpers.Descriptions.ConnectorRefText.String(),
							Type:        schema.TypeString,
							Computed:    true,
						},
						"store_type": {
							Description: "Specifies whether the Entity is to be stored in Git or not. Possible values: INLINE, REMOTE.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"repo_name": {
							Description: "Name of the repository.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"last_object_id": {
							Description: "Last object identifier (for Github). To be provided only when updating Pipeline.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"last_commit_id": {
							Description: "Last commit identifier (for Git Repositories other than Github). To be provided only when updating Pipeline.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"is_harness_code_repo": {
							Description: "If the repo is a Harness Code repo",
							Type:        schema.TypeBool,
							Computed:    true,
						},
					},
				},
			},
		},
	}

	return resource
}

func dataSourceCatalogEntityRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetIDPClientWithContext(ctx)

	kind := d.Get("kind").(string)
	info, err := getCatalogEntityInfo(d, kind)
	if err != nil {
		return diag.Errorf("failed to get catalog entity info: %v", err)
	}

	resp, httpResp, err := c.EntitiesApi.GetEntity(ctx, info.Scope, info.Kind, info.Identifier, &idp.EntitiesApiGetEntityOpts{
		OrgIdentifier:     info.OrgId,
		ProjectIdentifier: info.ProjectId,
		HarnessAccount:    optional.NewString(c.AccountId),
	})
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	readCatalogEntity(d, resp)

	return nil
}

func getCatalogEntityInfo(d *schema.ResourceData, kind string) (catalogEntityInfo, error) {
	identifier := d.Get("identifier").(string)
	orgId := d.Get("org_id").(string)
	projectId := d.Get("project_id").(string)

	catalogInfo := catalogEntityInfo{
		Kind:       kind,
		Scope:      "account",
		Identifier: identifier,
	}

	if orgId != "" {
		catalogInfo.OrgId = optional.NewString(orgId)
		catalogInfo.Scope = fmt.Sprintf("%s.%s", catalogInfo.Scope, orgId)
	} else {
		catalogInfo.OrgId = optional.EmptyString()
	}

	if projectId != "" {
		catalogInfo.ProjectId = optional.NewString(projectId)
		catalogInfo.Scope = fmt.Sprintf("%s.%s", catalogInfo.Scope, projectId)
	} else {
		catalogInfo.ProjectId = optional.EmptyString()
	}

	return catalogInfo, nil
}
