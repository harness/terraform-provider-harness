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
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"identifier": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "CustomSecretManager",
			},
			"on_delegate": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"timeout": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"template_ref": {
				Type:     schema.TypeString,
				Required: true,
			},
			"version_label": {
				Type:     schema.TypeString,
				Required: true,
			},
			"target_host": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Host where the custom secrets manager is located, required if 'on_delegate' is false.",
			},
			"ssh_secret_ref": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "SSH secret reference for the custom secrets manager, required if 'on_delegate' is false.",
			},
			"working_directory": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The working directory for operations, required if 'on_delegate' is false.",
			},
			"template_inputs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"environment_variable": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:     schema.TypeString,
										Required: true,
									},
									"type": {
										Type:     schema.TypeString,
										Required: true,
									},
									"value": {
										Type:     schema.TypeString,
										Required: true,
									},
									"default": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  false,
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
