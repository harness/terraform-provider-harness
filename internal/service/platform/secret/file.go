package secret

import (
	"context"
	"fmt"
	"strings"

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
		Importer:      helpers.MultiLevelResourceImporter,

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

	tags := d.Get("tags").(*schema.Set)
	no_of_tags := tags.Len()
	var tags_string = buildTag(no_of_tags, tags)

	if id == "" {
		resp, _, err = c.SecretsApi.PostSecretFileV2(ctx, c.AccountId, &nextgen.SecretsApiPostSecretFileV2Opts{
			OrgIdentifier:     buildField(d, "org_id"),
			ProjectIdentifier: buildField(d, "project_id"),
			Spec:              optional.NewString(fmt.Sprintf(`{"secret":{"type":"SecretFile","name":"%[1]s","identifier":"%[2]s","description":"%[3]s","tags":%[4]s,"spec":{"secretManagerIdentifier":"%[5]s"}}}`, d.Get("name"), d.Get("identifier"), d.Get("description"), strings.Join(tags_string, ","), d.Get("secret_manager_identifier"))),
			File:              optional.NewInterface(d.Get("file_path").(string)),
		})
	} else {
		resp, _, err = c.SecretsApi.PutSecretFileV2(ctx, c.AccountId, id, &nextgen.SecretsApiPutSecretFileV2Opts{
			OrgIdentifier:     buildField(d, "org_id"),
			ProjectIdentifier: buildField(d, "project_id"),
			Spec:              optional.NewString(fmt.Sprintf(`{"secret":{"type":"SecretFile","name":"%[1]s","identifier":"%[2]s","description":"%[3]s","tags":%[4]s,"spec":{"secretManagerIdentifier":"%[5]s"}}}`, d.Get("name"), d.Get("identifier"), d.Get("description"), strings.Join(tags_string, ","), d.Get("secret_manager_identifier"))),
			File:              optional.NewInterface(d.Get("file_path").(string)),
		})
	}

	if err != nil {
		return diag.FromErr(err)
	}

	if err := readSecretFile(d, resp.Data.Secret); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func buildTag(no_of_tags int, tags *schema.Set) []string {
	var tags_string = make([]string, no_of_tags)
	for i := 0; i < tags.Len(); i++ {
		tag := fmt.Sprintf("%v", tags.List()[i])
		split_tag := strings.Split(tag, ":")
		key := split_tag[0]
		value := split_tag[1]
		tags_string[i] = fmt.Sprintf(`{"%[1]s":"%[2]s"}`, key, value)
	}
	return tags_string
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
