package module_registry

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func DataSourceInfraModule() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving modules from the module registry.",
		ReadContext: resourceInfraModuleRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "Identifier of the module",
				Type:        schema.TypeString,
				Required:    true,
			},
			"description": {
				Description: "Description of the module",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Description: "Name of the module",
				Type:        schema.TypeString,
				Required:    true,
			},
			"system": {
				Description: "Provider of the module",
				Type:        schema.TypeString,
				Required:    true,
			},
			"repository": {
				Description: "For account connectors, the repository where the module is stored",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"repository_branch": {
				Description: "Repository Branch in which the module should be accessed",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"repository_commit": {
				Description: "Repository Commit in which the module should be accessed",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"repository_connector": {
				Description: "Repository Connector is the reference to the connector for the repository",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"repository_path": {
				Description: "Repository Path is the path in which the module resides",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"created": {
				Description: "Timestamp when the module was created",
				Type:        schema.TypeInt,
				Computed:    true,
				Optional:    true,
			},
			"repository_url": {
				Description: "URL where the module is stored",
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
			},
			"synced": {
				Description: "Timestamp when the module was last synced",
				Type:        schema.TypeInt,
				Computed:    true,
				Optional:    true,
			},
			"tags": {
				Description: "Tags associated with the module",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
		},
	}
	return resource
}
