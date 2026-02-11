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
			"connector_org": {
				Description: "Repository connector orgoanization",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"connector_project": {
				Description: "Repository connector project",
				Type:        schema.TypeString,
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
			"account": {
				Description: "Account that owns the module",
				Type:        schema.TypeString,
				Required:    true,
			},
			"git_tag_style": {
				Description: "Git Tag Style",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"module_error": {
				Description: "Error while retrieving the module",
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
			},
			"org": {
				Description: "Organization that owns the module",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"project": {
				Description: "Project that owns the module",
				Type:        schema.TypeString,
				Optional:    true,
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
			"onboarding_pipeline": {
				Description: "Onboarding Pipeline identifier.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"onboarding_pipeline_org": {
				Description: "Onboarding Pipeline organization.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"onboarding_pipeline_project": {
				Description: "Onboarding Pipeline project.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"onboarding_pipeline_sync": {
				Description: "Sync the project automatically.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"storage_type": {
				Description: "How to store the artifact.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
		},
	}
	return resource
}
