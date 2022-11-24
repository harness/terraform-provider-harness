package connector

import (
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DatasourceConnectorSpot() *schema.Resource {
	resource := &schema.Resource{
		Description: "Datasource for looking up an Spot connector.",
		ReadContext: resourceConnectorSpotRead,

		Schema: map[string]*schema.Schema{
			"permanent_token": {
				Description: "Authenticate to Spot using account id and permanent token.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"spot_account_id": {
							Description: "Spot account id.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"spot_account_id_ref": {
							Description: "Reference to the Harness secret containing the spot account id." + secret_ref_text,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"api_token_ref": {
							Description: "Reference to the Harness secret containing the permanent api token." + secret_ref_text,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"delegate_selectors": {
							Description: "Connect only using delegates with these tags.",
							Type:        schema.TypeSet,
							Computed:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"execute_on_delegate": {
							Description: "Execute on delegate or not.",
							Type:        schema.TypeBool,
							Computed:    true,
						},
					},
				},
			},
		},
	}

	helpers.SetMultiLevelDatasourceSchema(resource.Schema)

	return resource
}
