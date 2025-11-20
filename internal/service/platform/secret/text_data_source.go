package secret

import (
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceSecretText() *schema.Resource {
	resource := &schema.Resource{
		Description: "DataSource for looking up secret of type secret text.",
		ReadContext: resourceSecretTextRead,

		Schema: map[string]*schema.Schema{
			"secret_manager_identifier": {
				Description: "Identifier of the Secret Manager used to manage the secret.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"value_type": {
				Description: "This has details to specify if the secret value is Inline or Reference.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"value": {
				Description: "Value of the Secret",
				Type:        schema.TypeString,
				Sensitive:   true,
				Computed:    true,
			},
			"additional_metadata": {
				Description: "Additional Metadata for the Secret",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"values": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"version": {
										Type:     schema.TypeString,
										Optional: true,
									},
									// Add other fields for the inner map as needed
								},
							},
						},
						// Add other fields for the outer map as needed
					},
				},
			},
		},
	}

	helpers.SetMultiLevelDatasourceSchemaIdentifierRequired(resource.Schema)

	return resource
}
