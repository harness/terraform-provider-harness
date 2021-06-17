package provider

// import (
// 	"context"

// 	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
// 	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
// 	"github.com/micahlmartin/terraform-provider-harness/harness/api"
// 	"github.com/micahlmartin/terraform-provider-harness/harness/api/graphql"
// )

// func resourceServiceKubernetes() *schema.Resource {
// 	return &schema.Resource{
// 		Description:   "Resource for creating a git connector",
// 		CreateContext: resourceServiceKubernetesCreate,
// 		ReadContext:   resourceServiceKubernetesRead,
// 		UpdateContext: resourceServiceKubernetesUpdate,
// 		DeleteContext: resourceServiceKubernetesDelete,

// 		Schema: map[string]*schema.Schema{
// 			"id": {
// 				Description: "Id of the service",
// 				Type:        schema.TypeString,
// 				Computed:    true,
// 			},
// 			"app_id": {
// 				Description: "The id of the application the service belongs to",
// 				Type:        schema.TypeString,
// 				Required:    true,
// 			},
// 			"name": {
// 				Description: "Name of the service",
// 				Type:        schema.TypeString,
// 				Required:    true,
// 			},
// 			"helm_version": {
// 				Description: "The version of Helm to use. Options are `V2` and `V3`. Defaults to 'V2'",
// 				Type:        schema.TypeString,
// 				Required:    false,
// 			},
// 		},
// 	}
// }

// func resourceServiceKubernetesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
// 	c := meta.(*api.Client)

// 	svc := d.Get("id").(string)

// 	if err != nil {
// 		return diag.FromErr(err)
// 	}

// 	d.Set("name", conn.Name)
// 	d.Set("created_at", conn.CreatedAt.String())
// 	d.Set("url", conn.Url)
// 	d.Set("branch", conn.Branch)
// 	d.Set("password_secret_id", conn.PasswordSecretId)
// 	d.Set("ssh_setting_id", conn.SSHSettingId)
// 	d.Set("webhook_url", conn.WebhookUrl)
// 	d.Set("url_type", conn.UrlType)
// 	d.Set("username", conn.UserName)
// 	d.Set("commit_details", flattenCommitDetails(conn.CustomCommitDetails))
// 	d.Set("delegate_selectors", flattenDelgateSelectors(conn.DelegateSelectors))

// 	return nil
// }

// func resourceServiceKubernetesCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
// 	c := meta.(*api.Client)

// 	// Validation
// 	if err := validateGitConnectorSecret(d.Get("ssh_setting_id").(string), d.Get("password_secret_id").(string)); err != nil {
// 		return diag.FromErr(err)
// 	}

// 	connInput := &graphql.GitConnectorInput{}
// 	connInput.Name = d.Get("name").(string)
// 	connInput.Url = d.Get("url").(string)
// 	connInput.Branch = d.Get("branch").(string)
// 	connInput.UserName = d.Get("username").(string)
// 	connInput.GenerateWebhookUrl = d.Get("generate_webhook_url").(bool)
// 	connInput.PasswordSecretId = d.Get("password_secret_id").(string)
// 	connInput.SSHSettingId = d.Get("ssh_setting_id").(string)
// 	connInput.UrlType = d.Get("url_type").(string)
// 	connInput.DelegateSelectors = expandDelegateSelectors(d.Get("delegate_selectors").([]interface{}))
// 	connInput.CustomCommitDetails = expandCommitDetails(d.Get("commit_details").(*schema.Set).List())

// 	conn, err := c.Connectors().CreateGitConnector(connInput)

// 	if err != nil {
// 		return diag.FromErr(err)
// 	}

// 	d.SetId(conn.Id)
// 	d.Set("webhook_url", conn.WebhookUrl)

// 	return nil
// }

// func resourceServiceKubernetesUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
// 	c := meta.(*api.Client)

// 	// Validation
// 	if d.HasChange("generate_webhook_url") && !d.Get("generate_webhook_url").(bool) {
// 		return diag.Diagnostics{
// 			diag.Diagnostic{
// 				Severity: diag.Error,
// 				Summary:  "generate_webhook_url cannot be changed",
// 				Detail:   "When generate_webhook_url is initially set to true it cannot be reversed. You can update it from `false` to `true` but not the other way around.",
// 			},
// 		}
// 	}

// 	if err := validateGitConnectorSecret(d.Get("ssh_setting_id").(string), d.Get("password_secret_id").(string)); err != nil {
// 		return diag.FromErr(err)
// 	}

// 	if d.HasChange("url_type") {
// 		return diag.Diagnostics{
// 			diag.Diagnostic{
// 				Severity: diag.Error,
// 				Summary:  "url_type cannot be changed",
// 			},
// 		}
// 	}

// 	id := d.Get("id").(string)
// 	connInput := &graphql.GitConnectorInput{}
// 	connInput.Name = d.Get("name").(string)
// 	connInput.Url = d.Get("url").(string)
// 	connInput.Branch = d.Get("branch").(string)
// 	connInput.UserName = d.Get("username").(string)
// 	connInput.PasswordSecretId = d.Get("password_secret_id").(string)
// 	connInput.SSHSettingId = d.Get("ssh_setting_id").(string)
// 	connInput.GenerateWebhookUrl = d.Get("generate_webhook_url").(bool)
// 	connInput.DelegateSelectors = expandDelegateSelectors(d.Get("delegate_selectors").([]interface{}))
// 	connInput.CustomCommitDetails = expandCommitDetails(d.Get("commit_details").(*schema.Set).List())

// 	conn, err := c.Connectors().UpdateGitConnector(id, connInput)

// 	if err != nil {
// 		return diag.FromErr(err)
// 	}

// 	// Set any computed fields
// 	d.Set("webhook_url", conn.WebhookUrl)

// 	return nil
// }

// func resourceServiceKubernetesDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
// 	c := meta.(*api.Client)

// 	id := d.Get("id").(string)

// 	err := c.Connectors().DeleteConnector(id)

// 	if err != nil {
// 		return diag.FromErr(err)
// 	}

// 	return nil
// }
