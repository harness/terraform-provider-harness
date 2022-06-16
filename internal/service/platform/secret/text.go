package secret

import (
	"context"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceSecretText() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating secret of type secret text",
		ReadContext:   resourceSecretTextRead,
		CreateContext: resourceSecretTextCreateOrUpdate,
		UpdateContext: resourceSecretTextCreateOrUpdate,
		DeleteContext: resourceSecretDelete,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"secret_manager_identifier": {
				Description: "Identifier of the Secret Manager used to manage the secret.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"value_type": {
				Description: "This has details to specify if the secret value is inline or referenced.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"value": {
				Description: "Value of the Secret",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
	helpers.SetMultiLevelResourceSchema(resource.Schema)

	return resource
}

func resourceSecretTextRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secret, err := resourceSecretReadBase(ctx, d, meta, nextgen.SecretTypes.SecretText)
	if err != nil {
		return err
	}

	if err := readSecretText(d, secret); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceSecretTextCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secret := buildSecretText(d)

	newSecret, err := resourceSecretCreateOrUpdateBase(ctx, d, meta, secret)
	if err != nil {
		return err
	}

	if err := readSecretText(d, newSecret); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func buildSecretText(d *schema.ResourceData) *nextgen.Secret {
	secret := &nextgen.Secret{
		Type_: nextgen.SecretTypes.SecretText,
		Text:  &nextgen.SecretTextSpec{},
	}

	if attr, ok := d.GetOk("secret_manager_identifier"); ok {
		secret.Text.SecretManagerIdentifier = attr.(string)
	}

	if attr, ok := d.GetOk("value_type"); ok {
		secret.Text.ValueType = nextgen.SecretTextValueType(attr.(string))
	}

	if attr, ok := d.GetOk("value"); ok {
		secret.Text.Value = attr.(string)
	}

	return secret
}

func readSecretText(d *schema.ResourceData, secret *nextgen.Secret) error {
	d.Set("secret_manager_identifier", secret.Text.SecretManagerIdentifier)
	d.Set("value_type", secret.Text.ValueType)
	d.Set("value", secret.Text.Value)

	return nil
}
