package connector

import (
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DatasourceConnectorPrometheus() *schema.Resource {
	resource := &schema.Resource{
		Description: "Datasource for looking up a Prometheus connector.",
		ReadContext: resourceConnectorPrometheusRead,

		Schema: map[string]*schema.Schema{
			"url": {
				Description: "Url of the Prometheus server.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"delegate_selectors": {
				Description: "Connect using only the delegates which have these tags.",
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"user_name": {
				Description: "User name.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"password_ref": {
				Description: "Password reference.",
				Type:        schema.TypeString,
				Computed:    true,
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
							Description: "Encrypted value reference.",
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

	helpers.SetMultiLevelDatasourceSchema(resource.Schema)

	return resource
}
