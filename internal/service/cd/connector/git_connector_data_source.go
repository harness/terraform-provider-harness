package connector

import (
	"context"
	"fmt"

	"github.com/harness-io/harness-go-sdk/harness/api"
	"github.com/harness-io/harness-go-sdk/harness/cd/graphql"
	"github.com/harness-io/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceGitConnector() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Data source for retrieving a Harness application",

		ReadContext: dataSourceGitConnectorRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Description: "Id of the encrypted text secret",
				Type:        schema.TypeString,
				Required:    true,
			},
			"name": {
				Description: "Name of the encrypted text secret",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"created_at": {
				Description: "The time the git connector was created",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"url": {
				Description: "The url of the git repository or account/organization",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"branch": {
				Description: "The branch of the git connector to use",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"commit_details": {
				Description: "Custom details to use when making commits using this git connector",
				Type:        schema.TypeSet,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"author_email_id": {
							Description: "The email id of the author",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"author_name": {
							Description: "The name of the author",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"message": {
							Description: "Commit message",
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"delegate_selectors": {
				Description: "Delegate selectors to apply to this git connector",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"generate_webhook_url": {
				Description: "Boolean indicating whether or not to generate a webhook url",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"webhook_url": {
				Description: "The generated webhook url",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"password_secret_id": {
				Description: "The id of the secret for connecting to the git repository",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"ssh_setting_id": {
				Description: "The id of the SSH secret to use",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"url_type": {
				Description: fmt.Sprintf("The type of git url being used. Options are `%s`, and `%s.`", graphql.GitUrlTypes.Account, graphql.GitUrlTypes.Repo),
				Type:        schema.TypeString,
				Computed:    true,
			},
			"username": {
				Description: "The name of the user used to connect to the git repository",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func dataSourceGitConnectorRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	c := meta.(*api.Client)

	id := d.Get("id").(string)
	conn, err := c.CDClient.ConnectorClient.GetGitConnectorById(id)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(id)
	d.Set("name", conn.Name)
	d.Set("created_at", conn.CreatedAt.String())
	d.Set("url", conn.Url)
	d.Set("branch", conn.Branch)
	d.Set("password_secret_id", conn.PasswordSecretId)
	d.Set("ssh_setting_id", conn.SSHSettingId)
	d.Set("webhook_url", conn.WebhookUrl)
	d.Set("url_type", conn.UrlType)
	d.Set("username", conn.UserName)
	d.Set("commit_details", flattenCommitDetails(conn.CustomCommitDetails))
	d.Set("delegate_selectors", utils.FlattenDelgateSelectors(conn.DelegateSelectors))

	return nil
}
