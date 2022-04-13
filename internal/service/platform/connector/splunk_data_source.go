package connector

import (
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal/gitsync"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DatasourceConnectorSplunk() *schema.Resource {
	resource := &schema.Resource{
		Description: "Datasource for looking up a Splunk connector.",
		ReadContext: resourceConnectorSplunkRead,

		Schema: map[string]*schema.Schema{
			"url": {
				Description: "Url of the Splunk server.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"username": {
				Description: "The username used for connecting to Splunk.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"account_id": {
				Description: "Splunk account id.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"password_ref": {
				Description: "The reference to the Harness secret containing the Splunk password.",
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
	gitsync.SetGitSyncSchema(resource.Schema, true)

	return resource
}
