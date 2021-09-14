package provider

import (
	"context"

	"github.com/harness-io/harness-go-sdk/harness/api"
	"github.com/harness-io/harness-go-sdk/harness/api/graphql"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEncryptedText() *schema.Resource {
	return &schema.Resource{
		Description:   "Resource for creating an encrypted text secret",
		CreateContext: resourceEncryptedTextCreate,
		ReadContext:   resourceEncryptedTextRead,
		UpdateContext: resourceEncryptedTextUpdate,
		DeleteContext: resourceEncryptedTextDelete,

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
			"inherit_scopes_from_secret_manager": {
				Description: "Boolean that indicates whether or not to inherit the usage scopes from the secret manager",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"scoped_to_account": {
				Description: "Boolean that indicates whether or not the secret is scoped to the account",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"secret_manager_id": {
				Description: "The id of the secret manager to associate the secret with. Once set, this field cannot be changed.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"value": {
				Description: "The value of the secret",
				Type:        schema.TypeString,
				Sensitive:   true,
				Required:    true,
			},
			"usage_scope": usageScopeSchema(),
		},

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceEncryptedTextRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)

	secretId := d.Get("id").(string)

	secret, err := c.Secrets().GetEncryptedTextById(secretId)
	if err != nil {
		return diag.FromErr(err)
	} else if secret == nil {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	return readEncryptedText(d, secret)

}

func readEncryptedText(d *schema.ResourceData, secret *graphql.EncryptedText) diag.Diagnostics {
	d.SetId(secret.Id)
	d.Set("name", secret.Name)
	d.Set("inherit_scopes_from_secret_manager", secret.InheritScopesFromSM)
	d.Set("scoped_to_account", secret.ScopedToAccount)
	d.Set("secret_manager_id", secret.SecretManagerId)
	d.Set("usage_scope", flattenUsageScope(secret.UsageScope))

	return nil
}

func resourceEncryptedTextCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)

	input := &graphql.CreateSecretInput{
		EncryptedText: &graphql.EncryptedTextInput{
			InheritScopesFromSM: d.Get("inherit_scopes_from_secret_manager").(bool),
			Name:                d.Get("name").(string),
			ScopedToAccount:     d.Get("scoped_to_account").(bool),
			SecretManagerId:     d.Get("secret_manager_id").(string),
			Value:               d.Get("value").(string),
		},
	}

	usageScope, err := expandUsageScope(d.Get("usage_scope").(*schema.Set).List())
	if err != nil {
		return diag.FromErr(err)
	}
	input.EncryptedText.UsageScope = usageScope

	secret, err := c.Secrets().CreateEncryptedText(input)
	if err != nil {
		return diag.FromErr(err)
	}

	return readEncryptedText(d, secret)
}

func resourceEncryptedTextUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)

	input := &graphql.UpdateSecretInput{
		SecretId: d.Get("id").(string),
		EncryptedText: &graphql.UpdateEncryptedText{
			InheritScopesFromSM: d.Get("inherit_scopes_from_secret_manager").(bool),
			Name:                d.Get("name").(string),
			ScopedToAccount:     d.Get("scoped_to_account").(bool),
			Value:               d.Get("value").(string),
		},
	}

	usageScope, err := expandUsageScope(d.Get("usage_scope").(*schema.Set).List())
	if err != nil {
		return diag.FromErr(err)
	}
	input.EncryptedText.UsageScope = usageScope

	secret, err := c.Secrets().UpdateEncryptedText(input)
	if err != nil {
		return diag.FromErr(err)
	}

	return readEncryptedText(d, secret)
}

func resourceEncryptedTextDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)

	err := c.Secrets().DeleteSecret(d.Get("id").(string), graphql.SecretTypes.EncryptedText)

	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
