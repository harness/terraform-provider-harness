package connector

import (
	"fmt"
	"strings"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DatasourceConnectorDocker() *schema.Resource {
	resource := &schema.Resource{
		Description: "Datasource for looking up a Datadog connector.",
		ReadContext: resourceConnectorDockerRead,

		Schema: map[string]*schema.Schema{
			"type": {
				Description: fmt.Sprintf("The type of the docker registry. Valid options are %s", strings.Join(nextgen.DockerRegistryTypesSlice, ", ")),
				Type:        schema.TypeString,
				Computed:    true,
			},
			"url": {
				Description: "The url of the docker registry.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"delegate_selectors": {
				Description: "Connect using only the delegates which have these tags.",
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"credentials": {
				Description: "The credentials to use for the docker registry. If not specified then the connection is made to the registry anonymously.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"username": {
							Description: "The username to use for the docker registry.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"username_ref": {
							Description: "The reference to the username to use for the docker registry.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"password_ref": {
							Description: "The reference to the password to use for the docker registry.",
							Type:        schema.TypeString,
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
