package secret

import (
	"context"
	"fmt"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceSecretFile() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for creating a secret of type secret file in Harness.",

		ReadContext:   resourceSecretFileRead,
		UpdateContext: resourceSecretFileCreateOrUpdate,
		DeleteContext: resourceSecretDelete,
		CreateContext: resourceSecretFileCreateOrUpdate,

		Schema: map[string]*schema.Schema{
			"secret_manager_identifier": {
				Description: "Identifier of the Secret Manager used to manage the secret.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"file_path": {
				Description: "Path of the file containing secret value",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
	helpers.SetMultiLevelResourceSchema(resource.Schema)

	return resource
}

func resourceSecretFileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secret, err := resourceSecretReadBase(ctx, d, meta, nextgen.SecretTypes.SecretFile)
	if err != nil {
		return err
	}

	if err := readSecretFile(d, secret); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceSecretFileCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Id()

	var err error
	var resp nextgen.ResponseDtoSecretResponse

	if id == "" {
		resp, _, err = c.SecretsApi.PostSecretFileV2(ctx, c.AccountId, &nextgen.SecretsApiPostSecretFileV2Opts{
			// OrgIdentifier:     optional.NewString(d.Get("org_id").(string)),
			// ProjectIdentifier: optional.NewString(d.Get("project_id").(string)),
			Spec: optional.NewString(fmt.Sprintf(`{"secret":{"type":"SecretFile","name":"%[1]s","identifier":"%[2]s","description":"%[3]s","tags":{},"spec":{"secretManagerIdentifier":"%[4]s"}}}`, d.Get("name"), d.Get("identifier"), d.Get("description"), d.Get("secret_manager_identifier"))),
			File: optional.NewInterface(d.Get("file_path").(string)),
		})
	} else {
		resp, _, err = c.SecretsApi.PutSecretFileV2(ctx, c.AccountId, id, &nextgen.SecretsApiPutSecretFileV2Opts{
			// OrgIdentifier:     optional.NewString(d.Get("org_id").(string)),
			// ProjectIdentifier: optional.NewString(d.Get("project_id").(string)),
			Spec: optional.NewString(fmt.Sprintf(`{"secret":{"type":"SecretFile","name":"%[1]s","identifier":"%[2]s","description":"%[3]s","tags":{},"spec":{"secretManagerIdentifier":"%[4]s"}}}`, d.Get("name"), d.Get("identifier"), d.Get("description"), d.Get("secret_manager_identifier"))),
			File: optional.NewInterface(d.Get("file_path")),
		}, d.Get("file_name").(string))
	}

	if err != nil {
		return diag.FromErr(err)
	}

	if err := readSecretFile(d, resp.Data.Secret); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func buildSecretFile(d *schema.ResourceData) *nextgen.Secret {
	secret := &nextgen.Secret{
		Type_: nextgen.SecretTypes.SecretFile,
		File:  &nextgen.SecretFileSpe{},
	}

	if attr := d.Get("name").(string); attr != "" {
		secret.Name = attr
	}

	if attr := d.Get("identifier").(string); attr != "" {
		secret.Identifier = attr
	}

	if attr := d.Get("description").(string); attr != "" {
		secret.Description = attr
	}

	if attr := d.Get("org_id").(string); attr != "" {
		secret.OrgIdentifier = attr
	}

	if attr := d.Get("project_id").(string); attr != "" {
		secret.ProjectIdentifier = attr
	}

	if attr := d.Get("tags").(*schema.Set).List(); len(attr) > 0 {
		secret.Tags = helpers.ExpandTags(attr)
	}

	if attr, ok := d.GetOk("secret_manager_identifier"); ok {
		secret.File.SecretManagerIdentifier = attr.(string)
	}

	return secret
}

func readSecretFile(d *schema.ResourceData, secret *nextgen.Secret) error {
	d.Set("secret_manager_identifier", secret.File.SecretManagerIdentifier)
	d.SetId(secret.Identifier)
	d.Set("identifier", secret.Identifier)
	d.Set("description", secret.Description)
	d.Set("name", secret.Name)
	d.Set("org_id", secret.OrgIdentifier)
	d.Set("project_id", secret.ProjectIdentifier)
	d.Set("tags", helpers.FlattenTags(secret.Tags))
	return nil
}
