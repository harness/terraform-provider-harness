package provider

import (
	"context"
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/micahlmartin/terraform-provider-harness/internal/client"
	"github.com/zclconf/go-cty/cty"
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
						"commit_message": {
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
			},
			"webhook_url": {
				Description: "The generated webhook url",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"password_secret_id": {
				Description: "The id of the secret for connecting to the git repository",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"ssh_setting_id": {
				Description: "The id of the SSH secret to use",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"url_type": {
				Description: fmt.Sprintf("The type of git url being used. Options are `%s`, and `%s.`", client.GitUrlTypes.Account, client.GitUrlTypes.Repo),
				Type:        schema.TypeString,
				Optional:    true,
				// ValidateDiagFunc: validateUrlType,
			},
			"username": {
				Description: "The name of the user used to connect to the git repository",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}

func resourceGitConnectorRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.ApiClient)

	connId := d.Get("id").(string)

	conn, err := c.Connectors().GetGitConnectorById(connId)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("name", conn.Name)
	d.Set("created_at", conn.CreatedAt.String())
	d.Set("url", conn.Url)
	d.Set("branch", conn.Branch)
	d.Set("password_secret_id", conn.PasswordSecretId)
	d.Set("ssh_setting_id", conn.SSHSettingId)
	d.Set("webhook_url", conn.WebhookUrl)
	d.Set("url_type", conn.UrlType)
	d.Set("username", conn.UserName)

	// TODO: commit details
	// TODO: delegate selectors

	return nil
}

func resourceGitConnectorCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.ApiClient)

	connInput := &client.GitConnectorInput{}
	connInput.Name = d.Get("name").(string)
	connInput.Url = d.Get("url").(string)
	connInput.Branch = d.Get("branch").(string)
	connInput.UserName = d.Get("username").(string)
	connInput.GenerateWebhookUrl = d.Get("generate_webhook_url").(bool)
	connInput.PasswordSecretId = d.Get("password_secret_id").(string)
	connInput.SSHSettingId = d.Get("ssh_setting_id").(string)
	connInput.UrlType = d.Get("url_type").(string)

	// TODO: custom commit details
	// TODO: delegate selectors

	conn, err := c.Connectors().CreateGitConnector(connInput)

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(conn.Id)
	d.Set("webhook_url", conn.WebhookUrl)

	return nil
}

func resourceGitConnectorUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.ApiClient)

	// Validation
	if d.HasChange("generate_webhook_url") && !d.Get("generate_webhook_url").(bool) {
		return diag.Diagnostics{
			diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "generate_webhook_url cannot be changed",
				Detail:   "When generate_webhook_url is initially set to true it cannot be reversed. You can update it from `false` to `true` but not the other way around.",
			},
		}
	}

	if d.HasChange("url_type") {
		return diag.Diagnostics{
			diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "url_type cannot be changed",
			},
		}
	}

	id := d.Get("id").(string)
	connInput := &client.GitConnectorInput{}
	connInput.Name = d.Get("name").(string)
	connInput.Url = d.Get("url").(string)
	connInput.Branch = d.Get("branch").(string)
	connInput.UserName = d.Get("username").(string)
	connInput.PasswordSecretId = d.Get("password_secret_id").(string)
	connInput.SSHSettingId = d.Get("ssh_setting_id").(string)
	connInput.GenerateWebhookUrl = d.Get("generate_webhook_url").(bool)
	// TODO: custom commit details
	// TODO: delegate selectors

	conn, err := c.Connectors().UpdateGitConnector(id, connInput)

	if err != nil {
		return diag.FromErr(err)
	}

	// Set any computed fields
	d.Set("webhook_url", conn.WebhookUrl)

	return nil
}

func resourceGitConnectorDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.ApiClient)

	id := d.Get("id").(string)

	err := c.Connectors().DeleteConnector(id)

	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func validateUrlType(i interface{}, path cty.Path) diag.Diagnostics {
	v, ok := i.(string)
	if !ok {
		return diag.Errorf("expected type to be string")
	}

	rx := regexp.MustCompile(fmt.Sprintf("%s|%s", client.GitUrlTypes.Account, client.GitUrlTypes.Repo))

	if !rx.MatchString(v) {
		return diag.Diagnostics{
			diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "invalid url_type",
				Detail:   fmt.Sprintf("value must be either %s or %s", client.GitUrlTypes.Account, client.GitUrlTypes.Repo),
			},
		}
	}

	return nil
}
