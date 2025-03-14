package codeRepositories

import (
	"fmt"
	"strings"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DatasourceConnectorGitlab() *schema.Resource {
	resource := &schema.Resource{
		Description: "Datasource for looking up a Gitlab connector.",
		ReadContext: resourceConnectorGitlabRead,

		Schema: map[string]*schema.Schema{
			"url": {
				Description: "URL of the gitlab repository or account.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"connection_type": {
				Description: fmt.Sprintf("Whether the connection we're making is to a gitlab repository or a gitlab account. Valid values are %s.", strings.Join(nextgen.GitConnectorTypeValues, ", ")),
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
			"execute_on_delegate": {
				Description: "Execute on delegate or not.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"api_authentication": {
				Description: "Configuration for using the gitlab api. API Access is Computed for using “Git Experience”, for creation of Git based triggers, Webhooks management and updating Git statuses.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"token_ref": {
							Description: "Personal access token for interacting with the gitlab api." + secretRefText,
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
										Description: "Reference to a secret containing the username to use for authentication." + secretRefText,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"token_ref": {
										Description: "Reference to a secret containing the personal access to use for authentication." + secretRefText,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"password_ref": {
										Description: "Reference to a secret containing the password to use for authentication." + secretRefText,
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
										Description: "Reference to the Harness secret containing the ssh key." + secretRefText,
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

	helpers.SetMultiLevelDatasourceSchemaIdentifierRequired(resource.Schema)

	return resource
}
