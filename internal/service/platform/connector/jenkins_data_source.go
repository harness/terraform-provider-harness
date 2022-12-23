package connector

import (
	"context"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceConnectorJenkins() *schema.Resource {
	resource := &schema.Resource{
		Description: "Datasource for looking up a Jenkins connector.",
		ReadContext: dataConnectorJenkinsRead,

		Schema: map[string]*schema.Schema{
			"jenkins_url": {
				Description: "Jenkins Url.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"delegate_selectors": {
				Description: "Tags to filter delegates for connection.",
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"auth": {
				Description: "This entity contains the details for Jenkins Authentication.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Description: "Can be one of UsernamePassword, Anonymous, BearerToken",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"jenkins_bearer_token": {
							Description: "Authenticate to App Dynamics using bearer token.",
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"token_ref": {
										Description: "Reference of the token." + secret_ref_text,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
						"jenkins_user_name_password": {
							Description: "Authenticate to App Dynamics using user name and password.",
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
										Description: "Username reference to use for authentication.",
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
				},
			},
		},
	}

	helpers.SetMultiLevelDatasourceSchema(resource.Schema)

	return resource
}

func dataConnectorJenkinsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn, err := dataConnectorReadBase(ctx, d, meta, nextgen.ConnectorTypes.Jenkins)
	if err != nil {
		return err
	}

	if err := readConnectorJenkins(d, conn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}
