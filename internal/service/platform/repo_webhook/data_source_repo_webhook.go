package repo_webhook

import (
	"context"

	"github.com/harness/harness-go-sdk/harness/code"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceRepoWebhook() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving a Harness repo webhook.",
		ReadContext: dataSourceRepoWebhookRead,
		Schema:      createSchema(),
	}

	helpers.SetMultiLevelDatasourceSchemaWithoutCommonFields(resource.Schema)

	return resource
}

func dataSourceRepoWebhookRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetCodeClientWithContext(ctx)

	repoIdentifier := d.Get("repo_identifier").(string)
	orgID := helpers.BuildField(d, "org_id")
	projectID := helpers.BuildField(d, "project_id")

	rule, resp, err := c.WebhookApi.GetWebhook(
		ctx,
		c.AccountId,
		repoIdentifier,
		d.Id(),
		&code.WebhookApiGetWebhookOpts{
			OrgIdentifier:     orgID,
			ProjectIdentifier: projectID,
		},
	)
	if resp.StatusCode == 404 {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	if err != nil {
		return helpers.HandleApiError(err, d, resp)
	}

	readRepoWebhook(d, &rule, orgID.Value(), projectID.Value(), repoIdentifier)

	return nil
}
