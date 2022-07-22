package service_account

import (
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceServiceAccount() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving service account.",

		ReadContext: resourceServiceAccountRead,
		Schema: map[string]*schema.Schema{
			"email": {
				Description: "Email of the Service Account.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"account_id": {
				Description: "Account Identifier for the Entity.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
	helpers.SetMultiLevelDatasourceSchema(resource.Schema)

	return resource
}
