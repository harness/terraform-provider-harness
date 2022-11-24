package connector

import (
	"fmt"
	"strings"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DatasourceConnectorGithub() *schema.Resource {
	resource := &schema.Resource{
		Description: "Datasource for looking up a Github connector.",
		ReadContext: resourceConnectorGithubRead,

		Schema: map[string]*schema.Schema{
			"url": {
				Description: "URL of the github repository or account.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"connection_type": {
				Description: fmt.Sprintf("Whether the connection we're making is to a github repository or a github account. Valid values are %s.", strings.Join(nextgen.GitConnectorTypeValues, ", ")),
				Type:        schema.TypeString,
				Computed:    true,
			},
			"validation_repo": {
				Description: "Repository to test the connection with. This is only used when `connection_type` is `Account`.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"delegate_selectors": {
				Description: "Tags to filter delegates for connection.",
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"api_authentication": {
				Description: "Configuration for using the github api. API Access is Computed for using “Git Experience”, for creation of Git based triggers, Webhooks management and updating Git statuses.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"github_app": {
							Description: "Configuration for using the github app for interacting with the github api.",
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"installation_id": {
										Description: "Enter the Installation ID located in the URL of the installed GitHub App.",
										Type:        schema.TypeString,
										Computed:    true,
									},
									"application_id": {
										Description: "Enter the GitHub App ID from the GitHub App General tab.",
										Type:        schema.TypeString,
										Computed:    true,
									},
									"private_key_ref": {
										Description: "Reference to the secret containing the private key." + secret_ref_text,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
						"token_ref": {
							Description: "Personal access token for interacting with the github api." + secret_ref_text,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"credentials": {
				Description: "Credentials to use for the connection.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"http": {
							Description: "Authenticate using Username and password over http(s) for the connection.",
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
									"token_ref": {
										Description: "Reference to a secret containing the personal access to use for authentication." + secret_ref_text,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
						"ssh": {
							Description: "Authenticate using SSH for the connection.",
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ssh_key_ref": {
										Description: "Reference to the Harness secret containing the ssh key." + secret_ref_text,
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
