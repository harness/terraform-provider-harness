package module_registry_testing

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func DataSourceInfraModuleTesting() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving modules testing metadata from the module registry.",
		ReadContext: resourceInfraModuleTestingRead,
		Schema: map[string]*schema.Schema{
			"module_id": {
				Description: "Identifier of the module to enable testing for",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			// Input fields for enabling/updating testing
			"org": {
				Description: "Organization identifier",
				Type:        schema.TypeString,
				Required:    true,
			},
			"project": {
				Description: "Project identifier",
				Type:        schema.TypeString,
				Required:    true,
			},
			"provider_connector": {
				Description: "Provider connector for testing purposes",
				Type:        schema.TypeString,
				Required:    true,
			},
			"provisioner_type": {
				Description: "Provisioner type for testing purposes (e.g., terraform, tofu)",
				Type:        schema.TypeString,
				Required:    true,
			},
			"provisioner_version": {
				Description: "Provisioner version for testing purposes",
				Type:        schema.TypeString,
				Required:    true,
			},
			"pipelines": {
				Description: "List of pipeline IDs to create webhooks for triggering test executions",
				Type:        schema.TypeList,
				Required:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"release_pipeline": {
				Description: "Pipeline ID to create webhooks for releases",
				Type:        schema.TypeString,
				Optional:    true,
			},
			// Computed output fields from the module
			"description": {
				Description: "Description of the module",
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
			"repository": {
				Description: "For account connectors, the repository where the module is stored",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"repository_branch": {
				Description: "Repository Branch in which the module should be accessed",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"repository_commit": {
				Description: "Repository Commit in which the module should be accessed",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"repository_connector": {
				Description: "Repository Connector is the reference to the connector for the repository",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"repository_path": {
				Description: "Repository Path is the path in which the module resides",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"created": {
				Description: "Timestamp when the module was created",
				Type:        schema.TypeInt,
				Computed:    true,
			},
			"repository_url": {
				Description: "URL where the module is stored",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"synced": {
				Description: "Timestamp when the module was last synced",
				Type:        schema.TypeInt,
				Computed:    true,
			},
			"tags": {
				Description: "Tags associated with the module",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"account": {
				Description: "Account that owns the module",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"git_tag_style": {
				Description: "Git Tag Style",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"module_error": {
				Description: "Error while retrieving the module",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"testing_enabled": {
				Description: "Whether testing is enabled for the module",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"testing_metadata": {
				Description: "Testing metadata for the module",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account": {
							Description: "Account is the internal customer account ID",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"org": {
							Description: "Organization identifier",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"pipelines": {
							Description: "Pipelines where the testing is enabled",
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"project": {
							Description: "Project identifier",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"provider_connector": {
							Description: "Provider connector for testing purposes",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"provisioner_type": {
							Description: "Provisioner type for testing purposes",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"provisioner_version": {
							Description: "Provisioner version for testing purposes",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"release_pipeline": {
							Description: "Release pipeline",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"updated": {
				Description: "Timestamp when the module was last modified",
				Type:        schema.TypeInt,
				Computed:    true,
				Optional:    true,
			},
			"versions": {
				Description: "Versions of the module",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
	return resource
}
