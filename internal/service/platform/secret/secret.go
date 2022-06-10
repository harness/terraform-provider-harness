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

type ReadSecretData func(*schema.ResourceData, *nextgen.Secret) error

func resourceSecretReadBase(ctx context.Context, d *schema.ResourceData, meta interface{}, secretType nextgen.SecretType) (*nextgen.Secret, diag.Diagnostics) {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Id()
	if id == "" {
		id = d.Get("identifier").(string)
	}

	resp, _, err := c.SecretsApi.GetSecretV2(ctx, id, c.AccountId, getReadSecretOpts(d))
	if err != nil {
		return nil, helpers.HandleApiError(err, d)
	}

	if secretType != resp.Data.Secret.Type_ {
		return nil, diag.FromErr(fmt.Errorf("expected secret to be of type %s, but got %s", secretType, resp.Data.Secret.Type_))
	}

	readCommonSecretData(d, resp.Data.Secret)

	return resp.Data.Secret, nil
}

func getReadSecretOpts(d *schema.ResourceData) *nextgen.SecretsApiGetSecretV2Opts {
	secretOpts := &nextgen.SecretsApiGetSecretV2Opts{}

	if attr, ok := d.GetOk("org_id"); ok {
		secretOpts.OrgIdentifier = optional.NewString(attr.(string))
	}
	if attr, ok := d.GetOk("project_id"); ok {
		secretOpts.ProjectIdentifier = optional.NewString(attr.(string))
	}

	return secretOpts
}

func resourceSecretCreateOrUpdateBase(ctx context.Context, d *schema.ResourceData, meta interface{}, secret *nextgen.Secret) (*nextgen.Secret, diag.Diagnostics) {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Id()
	buildSecret(d, secret)

	var err error
	var resp nextgen.ResponseDtoSecretResponse

	if id == "" {
		resp, _, err = c.SecretsApi.PostSecret(ctx, nextgen.SecretRequestWrapper{Secret: secret}, c.AccountId, &nextgen.SecretsApiPostSecretOpts{})
	} else {
		resp, _, err = c.SecretsApi.PutSecret(ctx, c.AccountId, d.Id(), &nextgen.SecretsApiPutSecretOpts{Body: optional.NewInterface(secret)})
	}

	if err != nil {
		return nil, helpers.HandleApiError(err, d)
	}

	readCommonSecretData(d, resp.Data.Secret)

	return resp.Data.Secret, nil
}

func resourceSecretDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	_, _, err := c.SecretsApi.DeleteSecretV2(ctx, d.Id(), c.AccountId, &nextgen.SecretsApiDeleteSecretV2Opts{})
	if err != nil {
		return diag.Errorf(err.(nextgen.GenericSwaggerError).Error())
	}

	return nil
}

func buildSecret(d *schema.ResourceData, secret *nextgen.Secret) {
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
}

func readCommonSecretData(d *schema.ResourceData, secret *nextgen.Secret) {
	d.SetId(secret.Identifier)
	d.Set("identifier", secret.Identifier)
	d.Set("description", secret.Description)
	d.Set("name", secret.Name)
	d.Set("org_id", secret.OrgIdentifier)
	d.Set("project_id", secret.ProjectIdentifier)
	d.Set("tags", helpers.FlattenTags(secret.Tags))
}
