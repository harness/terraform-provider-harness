package repo

import (
	"context"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/code"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceRepo() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving a Harness repo.",

		ReadContext: dataSourceRepoRead,

		Schema: createSchema(),
	}

	helpers.SetMultiLevelDatasourceSchema(resource.Schema)

	return resource
}

func dataSourceRepoRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetCodeClientWithContext(ctx)

	repo, resp, err := c.RepositoryApi.FindRepository(
		ctx,
		d.Get("account_id").(string),
		d.Get("identifier").(string),
		&code.RepositoryApiFindRepositoryOpts{
			OrgIdentifier:     optional.NewString(d.Get("org_identifier").(string)),
			ProjectIdentifier: optional.NewString(d.Get("project_identifier").(string)),
		},
	)
	if err != nil {
		return helpers.HandleApiError(err, d, resp)
	}

	readRepo(d, &repo)

	return nil
}
