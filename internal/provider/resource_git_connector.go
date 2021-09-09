package provider

import (
	"context"
	"fmt"
	"regexp"

	"github.com/harness-io/harness-go-sdk/harness/api"
	"github.com/harness-io/harness-go-sdk/harness/api/graphql"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGitConnector() *schema.Resource {
	return &schema.Resource{
		Description:   "Resource for creating a git connector",
		CreateContext: resourceGitConnectorCreate,
		ReadContext:   resourceGitConnectorRead,
		UpdateContext: resourceGitConnectorUpdate,
		DeleteContext: resourceGitConnectorDelete,

		Schema: map[string]*schema.Schema{
			"id": {
				Description: "Id of the encrypted text secret",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"name": {
				Description: "Name of the encrypted text secret",
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
				Description: "Delegate selectors to apply to this git connector",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"generate_webhook_url": {
				Description: "Boolean indicating whether or not to generate a webhook url",
				Default:     true,
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
			},
			"webhook_url": {
				Description: "The generated webhook url",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"password_secret_id": {
				Description:   "The id of the secret for connecting to the git repository",
				Type:          schema.TypeString,
				Optional:      true,
				AtLeastOneOf:  []string{"password_secret_id", "ssh_setting_id"},
				ConflictsWith: []string{"ssh_setting_id"},
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
				Optional:     true,
				ValidateFunc: validateUrlType,
				ForceNew:     true,
			},
			"username": {
				Description: "The name of the user used to connect to the git repository",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}

func validateUrlType(val interface{}, key string) (warn []string, errs []error) {
	v := val.(string)

	rx, err := regexp.Compile(fmt.Sprintf("%s|%s", graphql.GitUrlTypes.Account, graphql.GitUrlTypes.Repo))
	if err != nil {
		errs = append(errs, err)
	}

	if !rx.MatchString(v) {
		errs = append(errs, fmt.Errorf("invalid value %s. Must be one of %s or %s", v, graphql.GitUrlTypes.Account, graphql.GitUrlTypes.Repo))
	}

	return warn, errs
}

func resourceGitConnectorRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)

	connId := d.Get("id").(string)

	conn, err := c.Connectors().GetGitConnectorById(connId)
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
	d.Set("ssh_setting_id", conn.SSHSettingId)
	d.Set("webhook_url", conn.WebhookUrl)
	d.Set("url_type", conn.UrlType)
	d.Set("username", conn.UserName)
	d.Set("commit_details", flattenCommitDetails(conn.CustomCommitDetails))
	d.Set("delegate_selectors", flattenDelgateSelectors(conn.DelegateSelectors))

	return nil
}

func resourceGitConnectorCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)

	connInput := &graphql.GitConnectorInput{}
	connInput.Name = d.Get("name").(string)
	connInput.Url = d.Get("url").(string)
	connInput.Branch = d.Get("branch").(string)
	connInput.UserName = d.Get("username").(string)
	connInput.GenerateWebhookUrl = d.Get("generate_webhook_url").(bool)
	connInput.PasswordSecretId = d.Get("password_secret_id").(string)
	connInput.SSHSettingId = d.Get("ssh_setting_id").(string)
	connInput.UrlType = graphql.GitUrlType(d.Get("url_type").(string))
	connInput.DelegateSelectors = expandDelegateSelectors(d.Get("delegate_selectors").([]interface{}))
	connInput.CustomCommitDetails = expandCommitDetails(d.Get("commit_details").(*schema.Set).List())

	conn, err := c.Connectors().CreateGitConnector(connInput)

	if err != nil {
		return diag.FromErr(err)
	}

	return readGitConnector(d, conn)
}

func resourceGitConnectorUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)

	id := d.Get("id").(string)
	connInput := &graphql.GitConnectorInput{}
	connInput.Name = d.Get("name").(string)
	connInput.Url = d.Get("url").(string)
	connInput.Branch = d.Get("branch").(string)
	connInput.UserName = d.Get("username").(string)
	connInput.PasswordSecretId = d.Get("password_secret_id").(string)
	connInput.SSHSettingId = d.Get("ssh_setting_id").(string)
	connInput.GenerateWebhookUrl = d.Get("generate_webhook_url").(bool)
	connInput.DelegateSelectors = expandDelegateSelectors(d.Get("delegate_selectors").([]interface{}))
	connInput.CustomCommitDetails = expandCommitDetails(d.Get("commit_details").(*schema.Set).List())

	conn, err := c.Connectors().UpdateGitConnector(id, connInput)

	if err != nil {
		return diag.FromErr(err)
	}

	return readGitConnector(d, conn)
}

func resourceGitConnectorDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)

	id := d.Get("id").(string)

	err := c.Connectors().DeleteConnector(id)

	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func flattenCommitDetails(details *graphql.CustomCommitDetails) []interface{} {

	if details.IsEmpty() {
		return nil
	}

	cd := make([]interface{}, 1)

	if details == nil {
		// Create an empty commit details to remove it
		cd[0] = &graphql.CustomCommitDetails{}
	} else {
		cd[0] = map[string]string{
			"author_email_id": details.AuthorEmailId,
			"author_name":     details.AuthorName,
			"message":         details.CommitMessage,
		}
	}

	return cd
}

func expandCommitDetails(i []interface{}) *graphql.CustomCommitDetails {
	if len(i) <= 0 {
		return &graphql.CustomCommitDetails{}
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
