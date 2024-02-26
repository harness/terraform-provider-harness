package repo

import (
	"context"

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

	helpers.SetMultiLevelDatasourceSchemaWithoutCommonFields(resource.Schema)

	return resource
}

func dataSourceRepoRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetCodeClientWithContext(ctx)

	repoIdentifier := d.Get("identifier").(string)
	orgID := helpers.BuildField(d, "org_id")
	projectID := helpers.BuildField(d, "project_id")

	repo, resp, err := c.RepositoryApi.GetRepository(
		ctx,
		c.AccountId,
		repoIdentifier,
		&code.RepositoryApiGetRepositoryOpts{
			OrgIdentifier:     orgID,
			ProjectIdentifier: projectID,
		},
	)
	if err != nil {
		return helpers.HandleApiError(err, d, resp)
	}

	readRepo(d, &repo, orgID.Value(), projectID.Value())

	return nil
}
