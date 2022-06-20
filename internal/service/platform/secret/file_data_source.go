package secret

import (
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceSecretFile() *schema.Resource {
	resource := &schema.Resource{
		Description: "Datasource for looking up scert file type secret.",
		ReadContext: resourceSecretFileRead,

		Schema: map[string]*schema.Schema{
			"secret_manager_identifier": {
				Description: "Identifier of the Secret Manager used to manage the secret.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"file": {
				Description: "file containing secret value",
				Type:        schema.TypeMap,
				Required:    true,
			},
		},
	}

	helpers.SetMultiLevelDatasourceSchema(resource.Schema)

	return resource
}
