package secrets

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/harness/harness-go-sdk/harness/cd/graphql"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/harness/terraform-provider-harness/internal/service/cd/usagescope"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceEncryptedText() *schema.Resource {
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
				Description:   "The value of the secret.",
				Type:          schema.TypeString,
				Sensitive:     true,
				Optional:      true,
				ConflictsWith: []string{"secret_reference"},
				ExactlyOneOf:  []string{"value", "secret_reference"},
			},
			"secret_reference": {
				Description:   "Name of the existing secret. If you already have secrets created in a secrets manager such as HashiCorp Vault or AWS Secrets Manager, you do not need to re-create the existing secrets in Harness.",
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"value"},
				ExactlyOneOf:  []string{"value", "secret_reference"},
			},
			"usage_scope": usagescope.Schema(),
		},

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceEncryptedTextRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*internal.Session)

	secretId := d.Get("id").(string)

	secret, err := c.CDClient.SecretClient.GetEncryptedTextById(secretId)
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
	d.Set("usage_scope", usagescope.FlattenUsageScope(secret.UsageScope))

	return nil
}

func resourceEncryptedTextCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*internal.Session)

	input := &graphql.CreateSecretInput{
		EncryptedText: &graphql.EncryptedTextInput{},
	}

	if attr, ok := d.GetOk("inherit_scopes_from_secret_manager"); ok {
		input.EncryptedText.InheritScopesFromSM = attr.(bool)
	}

	if attr, ok := d.GetOk("name"); ok {
		input.EncryptedText.Name = attr.(string)
	}

	if attr, ok := d.GetOk("scoped_to_account"); ok {
		input.EncryptedText.ScopedToAccount = attr.(bool)
	}

	if attr, ok := d.GetOk("secret_manager_id"); ok {
		input.EncryptedText.SecretManagerId = attr.(string)
	}

	if attr, ok := d.GetOk("value"); ok {
		input.EncryptedText.Value = attr.(string)
	}

	if attr, ok := d.GetOk("secret_reference"); ok {
		input.EncryptedText.SecretReference = attr.(string)
	}

	usageScope, err := usagescope.ExpandUsageScope(d.Get("usage_scope").(*schema.Set).List())
	if err != nil {
		return diag.FromErr(err)
	}
	input.EncryptedText.UsageScope = usageScope

	secret, err := c.CDClient.SecretClient.CreateEncryptedText(input)
	if err != nil {
		return diag.FromErr(err)
	}

	return readEncryptedText(d, secret)
}

func resourceEncryptedTextUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*internal.Session)

	input := &graphql.UpdateSecretInput{
		SecretId:      d.Id(),
		EncryptedText: &graphql.UpdateEncryptedText{},
	}

	if attr, ok := d.GetOk("inherit_scopes_from_secret_manager"); ok {
		input.EncryptedText.InheritScopesFromSM = attr.(bool)
	}

	if attr, ok := d.GetOk("name"); ok {
		input.EncryptedText.Name = attr.(string)
	}

	if attr, ok := d.GetOk("scoped_to_account"); ok {
		input.EncryptedText.ScopedToAccount = attr.(bool)
	}

	if attr, ok := d.GetOk("value"); ok {
		input.EncryptedText.Value = attr.(string)
	}

	if attr, ok := d.GetOk("secret_reference"); ok {
		input.EncryptedText.SecretReference = attr.(string)
	}

	usageScope, err := usagescope.ExpandUsageScope(d.Get("usage_scope").(*schema.Set).List())
	if err != nil {
		return diag.FromErr(err)
	}
	input.EncryptedText.UsageScope = usageScope

	secret, err := c.CDClient.SecretClient.UpdateEncryptedText(input)
	if err != nil {
		return diag.FromErr(err)
	}

	return readEncryptedText(d, secret)
}

func resourceEncryptedTextDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*internal.Session)

	if err := c.CDClient.SecretClient.DeleteSecret(d.Get("id").(string), graphql.SecretTypes.EncryptedText); err != nil {

		// There is a racecondition that happens when resources referencing a secret are deleted by the usage count
		// on the secret isn't updated yet. For now I'm putting in a simple 5s retry which works for now.
		// If it fails again we'll just assume it's actually still in use.
		if strings.Contains(err.Error(), "still being used") {
			log.Println("[WARN] Secret is still being used, cannot delete. Retrying in 5s...")
			time.Sleep(time.Second * 5)
			if err := c.CDClient.SecretClient.DeleteSecret(d.Get("id").(string), graphql.SecretTypes.EncryptedText); err != nil {
				return diag.FromErr(err)
			}
		}
	}

	return nil
}
