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
	spec := buildSpec(d)

	if id == "" {
		resp, _, err = c.SecretsApi.PostSecretFileV2(ctx, c.AccountId, &nextgen.SecretsApiPostSecretFileV2Opts{
			OrgIdentifier:     buildField(d, "org_id"),
			ProjectIdentifier: buildField(d, "project_id"),
			Spec:              optional.NewString(spec),
			File:              optional.NewInterface(d.Get("file_path").(string)),
		})
	} else {
		resp, _, err = c.SecretsApi.PutSecretFileV2(ctx, c.AccountId, id, &nextgen.SecretsApiPutSecretFileV2Opts{
			OrgIdentifier:     buildField(d, "org_id"),
			ProjectIdentifier: buildField(d, "project_id"),
			Spec:              optional.NewString(spec),
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

func buildSpec(d *schema.ResourceData) string {
	spec := fmt.Sprintf(`{"secret":{"type":"SecretFile"`)
	if attr, ok := d.GetOk("name"); ok {
		spec = spec + fmt.Sprintf(`,"name":"%[1]s"`, attr.(string))
	}
	if attr, ok := d.GetOk("identifier"); ok {
		spec = spec + fmt.Sprintf(`,"identifier":"%[1]s"`, attr.(string))
	}
	if attr, ok := d.GetOk("description"); ok {
		spec = spec + fmt.Sprintf(`,"description":"%[1]s"`, attr.(string))
	}
	if attr, ok := d.GetOk("org_id"); ok {
		spec = spec + fmt.Sprintf(`,"orgIdentifier":"%[1]s"`, attr.(string))
	}
	if attr, ok := d.GetOk("project_id"); ok {
		spec = spec + fmt.Sprintf(`,"projectIdentifier":"%[1]s"`, attr.(string))
	}
	if attr, ok := d.GetOk("tags"); ok {
		tags := attr.(*schema.Set)
		no_of_tags := tags.Len()
		var tags_string = buildTag(no_of_tags, tags)
		spec = spec + fmt.Sprintf(`,"tags":%[1]s`, tags_string)
	}
	if attr, ok := d.GetOk("secret_manager_identifier"); ok {
		spec = spec + fmt.Sprintf(`,"spec":{"secretManagerIdentifier":"%[1]s"}`, attr.(string))
	}
	return spec + "}}"
}

func buildTag(no_of_tags int, tags *schema.Set) string {
	result := "{"
	tagMap := make(map[string]string)
	for i := 0; i < tags.Len(); i++ {
		tag := fmt.Sprintf("%v", tags.List()[i])
		if strings.Contains(tag, ":") {
			splitTag := strings.Split(tag, ":")
			key := splitTag[0]
			value := splitTag[1]
			tagMap[key] = value
		} else {
			tagMap[tag] = ""
		}
	}

	first := true
	for key, value := range tagMap {
		if !first {
			result += ","
		} else {
			first = false
		}
		result += fmt.Sprintf(`"%s":"%s"`, key, value)
	}

	result += "}"
	return result
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
