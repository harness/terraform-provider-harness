package secret

import (
	"context"
	"fmt"
	"net/http"

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

	resp, httpResp, err := c.SecretsApi.GetSecretV2(ctx, id, c.AccountId, getReadSecretOpts(d))
	if err != nil {
		return nil, helpers.HandleReadApiError(err, d, httpResp)
	}

	if resp.Data == nil {
		return nil, nil
	}

	if secretType != resp.Data.Secret.Type_ {
		return nil, diag.FromErr(fmt.Errorf("expected secret to be of type %s, but got %s", secretType, resp.Data.Secret.Type_))
	}

	readCommonSecretData(d, resp.Data.Secret)

	return resp.Data.Secret, nil
}

func getReadSecretOpts(d *schema.ResourceData) *nextgen.SecretsApiGetSecretV2Opts {
	secretOpts := &nextgen.SecretsApiGetSecretV2Opts{}

	secretOpts.OrgIdentifier = buildField(d, "org_id")
	secretOpts.ProjectIdentifier = buildField(d, "project_id")

	return secretOpts
}

func resourceSecretCreateOrUpdateBase(ctx context.Context, d *schema.ResourceData, meta interface{}, secret *nextgen.Secret) (*nextgen.Secret, diag.Diagnostics) {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Id()
	buildSecret(d, secret)

	var err error
	var resp nextgen.ResponseDtoSecretResponse
	var httpResp *http.Response

	if id == "" {
		resp, httpResp, err = c.SecretsApi.PostSecret(ctx, nextgen.SecretRequestWrapper{Secret: secret}, c.AccountId, &nextgen.SecretsApiPostSecretOpts{
			OrgIdentifier:     buildField(d, "org_id"),
			ProjectIdentifier: buildField(d, "project_id"),
		})
	} else {
		resp, httpResp, err = c.SecretsApi.PutSecret(ctx, c.AccountId, d.Id(), &nextgen.SecretsApiPutSecretOpts{
			OrgIdentifier:     buildField(d, "org_id"),
			ProjectIdentifier: buildField(d, "project_id"),
			Body:              optional.NewInterface(nextgen.SecretRequestWrapper{Secret: secret})})
	}

	if err != nil {
		return nil, helpers.HandleApiError(err, d, httpResp)
	}

	readCommonSecretData(d, resp.Data.Secret)

	return resp.Data.Secret, nil
}

func buildField(d *schema.ResourceData, field string) optional.String {
	if arr, ok := d.GetOk(field); ok {
		return optional.NewString(arr.(string))
	}
	return optional.EmptyString()
}

func resourceSecretDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	_, httpResp, err := c.SecretsApi.DeleteSecretV2(ctx, d.Id(), c.AccountId, &nextgen.SecretsApiDeleteSecretV2Opts{
		OrgIdentifier:     buildField(d, "org_id"),
		ProjectIdentifier: buildField(d, "project_id"),
	})
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
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
