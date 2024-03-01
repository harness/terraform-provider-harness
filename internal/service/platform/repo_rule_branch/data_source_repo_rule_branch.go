package repo_rule_branch

import (
	"context"

	"github.com/harness/harness-go-sdk/harness/code"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceRepoBranchRule() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving a Harness repo.",
		ReadContext: dataSourceRepoRuleRead,
		Schema:      createSchema(),
	}

	helpers.SetMultiLevelDatasourceSchemaWithoutCommonFields(resource.Schema)

	return resource
}

func dataSourceRepoRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetCodeClientWithContext(ctx)

	repoIdentifier := d.Get("repo_identifier").(string)
	orgID := helpers.BuildField(d, "org_id")
	projectID := helpers.BuildField(d, "project_id")

	rule, resp, err := c.RepositoryApi.RuleGet(
		ctx,
		c.AccountId,
		repoIdentifier,
		d.Id(),
		&code.RepositoryApiRuleGetOpts{
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

	readRepoBranchRule(d, &rule, orgID.Value(), projectID.Value(), repoIdentifier)

	return nil
}
