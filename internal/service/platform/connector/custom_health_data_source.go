package connector

import (
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DatasourceConnectorCustomHealthSource() *schema.Resource {
	resource := &schema.Resource{
		Description: "Datasource for looking up a Custom Health source connector.",
		ReadContext: resourceConnectorCustomHealthSourceRead,

		Schema: map[string]*schema.Schema{
			"url": {
				Description: "URL of the Custom Health source server.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"method": {
				Description: "HTTP Verb Method for the API Call",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"validation_body": {
				Description: "Body to be sent with the API Call",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"validation_path": {
				Description: "Path to be added to the base URL for the API Call",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"delegate_selectors": {
				Description: "Tags to filter delegates for connection.",
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"headers": {
				Description: "Headers.",
				Type:        schema.TypeSet,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Description: "Key.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"encrypted_value_ref": {
							Description: "Reference to the Harness secret containing the encrypted value." + secret_ref_text,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"value": {
							Description: "Value.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"value_encrypted": {
							Description: "Encrypted value.",
							Type:        schema.TypeBool,
							Computed:    true,
						},
					}},
			},
			"params": {
				Description: "Parameters.",
				Type:        schema.TypeSet,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Description: "Key.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"encrypted_value_ref": {
							Description: "Reference to the Harness secret containing the encrypted value." + secret_ref_text,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"value": {
							Description: "Value.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"value_encrypted": {
							Description: "Encrypted value.",
							Type:        schema.TypeBool,
							Computed:    true,
						},
					}},
			},
		},
	}

	helpers.SetMultiLevelDatasourceSchemaIdentifierRequired(resource.Schema)

	return resource
}
