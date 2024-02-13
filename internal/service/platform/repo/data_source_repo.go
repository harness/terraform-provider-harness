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

	accId := d.Get("account_id").(string)
	repoId := d.Get("identifier").(string)
	orgId := d.Get("org_id").(string)
	projId := d.Get("project_id").(string)

	repo, resp, err := c.RepositoryApi.FindRepository(
		ctx,
		accId,
		repoId,
		&code.RepositoryApiFindRepositoryOpts{
			OrgIdentifier:     optional.NewString(orgId),
			ProjectIdentifier: optional.NewString(projId),
		},
	)
	if err != nil {
		return helpers.HandleApiError(err, d, resp)
	}

	readRepo(d, &repo)

	return nil
}
