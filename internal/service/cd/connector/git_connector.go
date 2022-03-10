package connector

import (
	"context"
	"fmt"

	"github.com/harness/harness-go-sdk/harness/cd"
	"github.com/harness/harness-go-sdk/harness/cd/graphql"
	"github.com/harness/terraform-provider-harness/internal/service/cd/usagescope"
	"github.com/harness/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceGitConnector() *schema.Resource {
	return &schema.Resource{
		Description:   "Resource for creating a git connector",
		CreateContext: resourceGitConnectorCreate,
		ReadContext:   resourceGitConnectorRead,
		UpdateContext: resourceGitConnectorUpdate,
		DeleteContext: resourceGitConnectorDelete,

		Schema: map[string]*schema.Schema{
			"id": {
				Description: "Id of the git connector.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"name": {
				Description: "Name of the git connector.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"created_at": {
				Description: "The time the git connector was created",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"url": {
				Description: "The url of the git repository or account/organization",
				Type:        schema.TypeString,
				Required:    true,
			},
			"branch": {
				Description: "The branch of the git connector to use",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"commit_details": {
				Description: "Custom details to use when making commits using this git connector",
				Type:        schema.TypeSet,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"author_email_id": {
							Description: "The email id of the author",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"author_name": {
							Description: "The name of the author",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"message": {
							Description: "Commit message",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
			"delegate_selectors": {
				Description: "Delegate selectors to apply to this git connector.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"generate_webhook_url": {
				Description: "Boolean indicating whether or not to generate a webhook url.",
				// Default:     false,
				Type:     schema.TypeBool,
				Optional: true,
				// ForceNew:    true,
			},
			"webhook_url": {
				Description: "The generated webhook url",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"password_secret_id": {
				Description:   "The id of the secret for connecting to the git repository.",
				Type:          schema.TypeString,
				Optional:      true,
				AtLeastOneOf:  []string{"password_secret_id", "ssh_setting_id"},
				ConflictsWith: []string{"ssh_setting_id", "usage_scope"},
			},
			"ssh_setting_id": {
				Description:   "The id of the SSH secret to use",
				Type:          schema.TypeString,
				Optional:      true,
				AtLeastOneOf:  []string{"password_secret_id", "ssh_setting_id"},
				ConflictsWith: []string{"password_secret_id"},
			},
			"url_type": {
				Description:  fmt.Sprintf("The type of git url being used. Options are `%s`, and `%s.`", graphql.GitUrlTypes.Account, graphql.GitUrlTypes.Repo),
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{graphql.GitUrlTypes.Account.String(), graphql.GitUrlTypes.Repo.String()}, false),
				// ForceNew:     true,
			},
			"username": {
				Description: "The name of the user used to connect to the git repository",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"usage_scope": usagescope.Schema(),
		},

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceGitConnectorRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*cd.ApiClient)

	connId := d.Get("id").(string)

	conn, err := c.ConnectorClient.GetGitConnectorById(connId)
	if err != nil {
		return diag.FromErr(err)
	} else if conn == nil {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	return readGitConnector(d, conn)
}

func readGitConnector(d *schema.ResourceData, conn *graphql.GitConnector) diag.Diagnostics {
	d.SetId(conn.Id)
	d.Set("name", conn.Name)
	d.Set("created_at", conn.CreatedAt.String())
	d.Set("url", conn.Url)
	d.Set("branch", conn.Branch)
	d.Set("password_secret_id", conn.PasswordSecretId)
	d.Set("generate_webhook_url", conn.GenerateWebhookUrl || conn.WebhookUrl != "")
	d.Set("ssh_setting_id", conn.SSHSettingId)
	d.Set("webhook_url", conn.WebhookUrl)
	d.Set("url_type", conn.UrlType)
	d.Set("username", conn.UserName)

	if details := flattenCommitDetails(conn.CustomCommitDetails); len(details) > 0 {
		d.Set("commit_details", details)
	}

	if selectors := utils.FlattenDelgateSelectors(conn.DelegateSelectors); len(selectors) > 0 {
		d.Set("delegate_selectors", selectors)
	}

	if scope := usagescope.FlattenUsageScope(conn.UsageScope); len(scope) > 0 {
		d.Set("usage_scope", scope)
	}

	return nil
}

func setGitConnectorConfig(d *schema.ResourceData, connInput *graphql.GitConnectorInput, isUpdate bool) error {
	connInput.Name = d.Get("name").(string)
	connInput.Url = d.Get("url").(string)
	connInput.Branch = d.Get("branch").(string)
	connInput.UserName = d.Get("username").(string)
	connInput.GenerateWebhookUrl = d.Get("generate_webhook_url").(bool)
	connInput.PasswordSecretId = d.Get("password_secret_id").(string)
	connInput.SSHSettingId = d.Get("ssh_setting_id").(string)

	if !isUpdate {
		connInput.UrlType = graphql.GitUrlType(d.Get("url_type").(string))
	}

	if selectors := utils.ExpandDelegateSelectors(d.Get("delegate_selectors").(*schema.Set).List()); len(selectors) > 0 {
		connInput.DelegateSelectors = selectors
	}

	if details := expandCommitDetails(d.Get("commit_details").(*schema.Set).List()); details != nil {
		connInput.CustomCommitDetails = details
	}

	scope, err := usagescope.ExpandUsageScope(d.Get("usage_scope").(*schema.Set).List())
	if err != nil {
		return err
	}
	if scope != nil {
		connInput.UsageScope = scope
	}

	return nil
}

func resourceGitConnectorCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*cd.ApiClient)

	connInput := &graphql.GitConnectorInput{}
	err := setGitConnectorConfig(d, connInput, false)
	if err != nil {
		return diag.FromErr(err)
	}

	conn, err := c.ConnectorClient.CreateGitConnector(connInput)
	if err != nil {
		return diag.FromErr(err)
	}

	return readGitConnector(d, conn)
}

func resourceGitConnectorUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*cd.ApiClient)

	id := d.Get("id").(string)

	connInput := &graphql.GitConnectorInput{}
	err := setGitConnectorConfig(d, connInput, true)
	if err != nil {
		return diag.FromErr(err)
	}

	conn, err := c.ConnectorClient.UpdateGitConnector(id, connInput)
	if err != nil {
		return diag.FromErr(err)
	}

	return readGitConnector(d, conn)
}

func resourceGitConnectorDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*cd.ApiClient)

	id := d.Get("id").(string)

	err := c.ConnectorClient.DeleteConnector(id)

	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func flattenCommitDetails(details *graphql.CustomCommitDetails) []interface{} {
	results := []interface{}{}

	if details == nil || details.IsEmpty() {
		return results
	}

	cd := map[string]string{
		"author_email_id": details.AuthorEmailId,
		"author_name":     details.AuthorName,
		"message":         details.CommitMessage,
	}

	return append(results, cd)
}

func expandCommitDetails(i []interface{}) *graphql.CustomCommitDetails {
	if len(i) <= 0 {
		return nil
	}

	cd := i[0].(map[string]interface{})

	commitDetails := &graphql.CustomCommitDetails{}

	if attr, ok := cd["author_email_id"]; ok && attr != "" {
		commitDetails.AuthorEmailId = attr.(string)
	}

	if attr, ok := cd["author_name"]; ok && attr != "" {
		commitDetails.AuthorName = attr.(string)
	}

	if attr, ok := cd["message"]; ok && attr != "" {
		commitDetails.CommitMessage = attr.(string)
	}

	return commitDetails
}
