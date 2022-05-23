package connector

import (
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DatasourceConnectorPagerDuty() *schema.Resource {
	resource := &schema.Resource{
		Description: "Datasource for looking up a PagerDuty connector.",
		ReadContext: resourceConnectorPagerDutyRead,

		Schema: map[string]*schema.Schema{
			"api_token_ref": {
				Description: "Reference to the Harness secret containing the api token.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"delegate_selectors": {
				Description: "Connect using only the delegates which have these tags.",
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}

	helpers.SetMultiLevelDatasourceSchema(resource.Schema)

	return resource
}
