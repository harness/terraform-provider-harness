package module_registry

import (
	"context"

	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceInfraModules() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for retrieving a list of modules from the module registry.",
		ReadContext: dataSourceInfraModulesRead,
		Schema: map[string]*schema.Schema{
			"modules": {
				Description: "List of modules",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Description: "Identifier of the module",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"name": {
							Description: "Name of the module",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"system": {
							Description: "Provider of the module",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"description": {
							Description: "Description of the module",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"account": {
							Description: "Account that owns the module",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"org": {
							Description: "Organization that owns the module",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"project": {
							Description: "Project that owns the module",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"repository": {
							Description: "Repository where the module is stored",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"repository_branch": {
							Description: "Repository branch",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"repository_commit": {
							Description: "Repository commit",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"repository_connector": {
							Description: "Repository connector reference",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"repository_path": {
							Description: "Path within repository",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"repository_url": {
							Description: "Repository URL",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"tags": {
							Description: "Tags associated with the module",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"testing_enabled": {
							Description: "Whether testing is enabled for the module",
							Type:        schema.TypeBool,
							Computed:    true,
						},
						"created": {
							Description: "Timestamp when the module was created",
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"updated": {
							Description: "Timestamp when the module was last modified",
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"synced": {
							Description: "Timestamp when the module was last synced",
							Type:        schema.TypeInt,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceInfraModulesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c, ctx := m.(*internal.Session).GetPlatformClientWithContext(ctx)

	resp, httpRes, err := c.ModuleRegistryApi.ModuleRegistryListModulesByAccount(
		ctx,
		c.AccountId,
		nil,
	)
	if err != nil {
		return helpers.HandleApiError(err, d, httpRes)
	}

	// Flatten the modules list
	modules := make([]interface{}, 0, len(resp))
	for _, module := range resp {
		moduleMap := map[string]interface{}{
			"id":                   module.Id,
			"name":                 module.Name,
			"system":               module.System,
			"description":          module.Description,
			"account":              module.Account,
			"org":                  module.Org,
			"project":              module.Project,
			"repository":           module.Repository,
			"repository_branch":    module.RepositoryBranch,
			"repository_commit":    module.RepositoryCommit,
			"repository_connector": module.RepositoryConnector,
			"repository_path":      module.RepositoryPath,
			"repository_url":       module.RepositoryUrl,
			"tags":                 module.Tags,
			"testing_enabled":      module.TestingEnabled,
			"created":              module.Created,
			"updated":              module.Updated,
			"synced":               module.Synced,
		}
		modules = append(modules, moduleMap)
	}

	d.Set("modules", modules)
	
	d.SetId(c.AccountId)

	return nil
}
