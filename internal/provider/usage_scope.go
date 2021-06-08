package provider

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func usageScopeSchema() *schema.Schema {
	return &schema.Schema{
		Description: "Usage scopes",
		Type:        schema.TypeSet,
		Optional:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"application_id": {
					Description: "Id of the application scoping",
					Type:        schema.TypeString,
					Optional:    true,
				},
				"application_filter_type": {
					Description: "Type of application filter applied. ALL if not application id supplied, otherwise NULL",
					Type:        schema.TypeString,
					Optional:    true,
				},
				"environment_id": {
					Description: "Id of the environment scoping",
					Type:        schema.TypeString,
					Optional:    true,
				},
				"environment_filter_type": {
					Description: "Type of environment filter applied. ALL if not filter applied",
					Type:        schema.TypeString,
					Optional:    true,
				},
			},
		},
	}
}
