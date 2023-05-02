package connector

import (
	"fmt"
	"strings"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DatasourceConnectorNexus() *schema.Resource {
	resource := &schema.Resource{
		Description: "Datasource for looking up a Nexus connector.",
		ReadContext: resourceConnectorNexusRead,

		Schema: map[string]*schema.Schema{
			"url": {
				Description: "URL of the Nexus server.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"delegate_selectors": {
				Description: "Tags to filter delegates for connection.",
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"version": {
				Description: fmt.Sprintf("Version of the Nexus server. Valid values are %s", strings.Join(nextgen.NexusVersionSlice, ", ")),
				Type:        schema.TypeString,
				Computed:    true,
			},
			"credentials": {
				Description: "Credentials to use for authentication.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"username": {
							Description: "Username to use for authentication.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"username_ref": {
							Description: "Reference to a secret containing the username to use for authentication." + secret_ref_text,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"password_ref": {
							Description: "Reference to a secret containing the password to use for authentication." + secret_ref_text,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}

	helpers.SetMultiLevelDatasourceSchemaIdentifierRequired(resource.Schema)

	return resource
}
