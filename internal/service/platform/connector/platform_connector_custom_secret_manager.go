package connector

import (
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DatasourceConnectorCustomSM() *schema.Resource {
	resource := &schema.Resource{
		Description: "Datasource for looking up a Custom Secret Manager connector.",
		ReadContext: resourceConnectorCustomSMRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Description: "Name of the Custom Secret Manager Resource.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"description": {
				Description: "Description of the resource.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"identifier": {
				Description: "Unique identifier of the resource.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"type": {
				Description: "Type of the custom secrets manager, typically set to 'CustomSecretManager'.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"delegate_selectors": {
				Description: "Tags to filter delegates for connection.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"on_delegate": {
				Description: "Specifies whether the secrets manager runs on a Harness delegate.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"timeout": {
				Description: "Timeout in seconds for secrets management operations.",
				Type:        schema.TypeInt,
				Computed:    true,
			},
			"tags": {
				Description: "Tags to associate with the resource.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"template_ref": {
				Description: "Reference to the template used for managing secrets.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"version_label": {
				Description: "Version identifier of the secrets management template.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"target_host": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Host where the custom secrets manager is located, Computed if 'on_delegate' is false.",
			},
			"ssh_secret_ref": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "SSH secret reference for the custom secrets manager, Computed if 'on_delegate' is false.",
			},
			"working_directory": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The working directory for operations, Computed if 'on_delegate' is false.",
			},
			"template_inputs": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"environment_variable": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"value": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"default": {
										Type:     schema.TypeBool,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}

	helpers.SetMultiLevelDatasourceSchemaIdentifierRequired(resource.Schema)

	return resource
}
