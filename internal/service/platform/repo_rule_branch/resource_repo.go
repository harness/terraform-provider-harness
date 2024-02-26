package repo

import (
	"context"
	"net/http"
	"strings"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/code"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceRepoRule() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for creating a Harness Repo Rule.",

		ReadContext:   resourceRepoRuleRead,
		CreateContext: resourceRepoRuleCreateOrUpdate,
		UpdateContext: resourceRepoRuleCreateOrUpdate,
		DeleteContext: resourceRepoRuleDelete,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: createSchema(),
	}

	helpers.SetMultiLevelDatasourceSchemaWithoutCommonFields(resource.Schema)

	return resource
}

func resourceRepoRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetCodeClientWithContext(ctx)

	id := d.Id()
	orgID := helpers.BuildField(d, "org_id")
	projectID := helpers.BuildField(d, "project_id")
	repoIdentifier := helpers.BuildField(d, "repo_identifier")

	rule, resp, err := c.RepositoryApi.RuleGet(
		ctx,
		c.AccountId,
		repoIdentifier.Value(),
		id,
		&code.RepositoryApiRuleGetOpts{
			OrgIdentifier:     orgID,
			ProjectIdentifier: projectID,
		},
	)
	if err != nil {
		return helpers.HandleReadApiError(err, d, resp)
	}

	readRule(d, &rule[0], orgID.Value(), projectID.Value())

	return nil
}

func resourceRepoRuleCreateOrUpdate(
	ctx context.Context,
	d *schema.ResourceData,
	meta interface{},
) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetCodeClientWithContext(ctx)
	var err error
	var repo code.TypesRepository
	var resp *http.Response

	id := d.Id()
	orgID := helpers.BuildField(d, "org_id")
	projectID := helpers.BuildField(d, "project_id")

	// check
	var sourceRepo string
	source := getSource(d)
	if source != nil {
		sourceRepo = source["repo"].(string)
	}

	// determine what type of write the change requires
	switch {
	case id != "":
		// update repo
		body := buildRuleBody(d)
		body.DefaultBranch = ""
		repo, resp, err = c.RepositoryApi.UpdateRepository(
			ctx,
			c.AccountId,
			id,
			&code.RepositoryApiUpdateRepositoryOpts{
				Body:              optional.NewInterface(body),
				OrgIdentifier:     orgID,
				ProjectIdentifier: projectID,
			},
		)
	case sourceRepo != "":
		// import repo
		repo, resp, err = c.RepositoryApi.ImportRepository(
			ctx,
			c.AccountId,
			&code.RepositoryApiImportRepositoryOpts{
				Body:              optional.NewInterface(buildRepoImportBody(d)),
				OrgIdentifier:     orgID,
				ProjectIdentifier: projectID,
			},
		)
	default:
		// create repo
		repo, resp, err = c.RepositoryApi.CreateRepository(
			ctx,
			c.AccountId,
			&code.RepositoryApiCreateRepositoryOpts{
				Body:              optional.NewInterface(buildRuleBody(d)),
				OrgIdentifier:     orgID,
				ProjectIdentifier: projectID,
			},
		)
	}
	if err != nil {
		return helpers.HandleApiError(err, d, resp)
	}

	readRule(d, &repo, orgID.Value(), projectID.Value())
	return nil
}

func resourceRepoRuleDelete(
	ctx context.Context,
	d *schema.ResourceData,
	meta interface{},
) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetCodeClientWithContext(ctx)

	id := d.Id()

	resp, err := c.RepositoryApi.RuleDelete(
		ctx,
		c.AccountId,
		helpers.BuildField(d, "repo_identifier").Value(),
		id,
		&code.RepositoryApiRuleDeleteOpts{
			OrgIdentifier:     helpers.BuildField(d, "org_id"),
			ProjectIdentifier: helpers.BuildField(d, "project_id"),
		},
	)
	if err != nil {
		return helpers.HandleApiError(err, d, resp)
	}

	return nil
}

func buildRuleBody(d *schema.ResourceData) *code.OpenapiRule {
	return &code.OpenapiRule{
		Definition:  buildRuleDef(d),
		Description: d.Get("description").(string),
		Identifier:  d.Id(),
		Pattern:     buildPattern(),
		//State:       code.ACTIVE_EnumRuleState.(code.EnumRuleState),
		Type_:   code.BRANCH_OpenapiRuleType,
		Updated: 0,
		Users:   nil,
	}
}

func buildPattern() *code.ProtectionPattern {

}

func buildRuleDef(d *schema.ResourceData) *code.OpenapiRuleDefinition {
	rulesSet := d.Get("rules").(*schema.Set)
	rulesList := rulesSet.List()
	if len(rulesList) == 0 {
		return nil
	}
	rulesElem := rulesList[0]
	rules := rulesElem.(map[string]interface{})

	return &code.OpenapiRuleDefinition{ProtectionBranch: code.ProtectionBranch(struct {
		Bypass    *code.ProtectionDefBypass
		Lifecycle *code.ProtectionDefLifecycle
		Pullreq   *code.ProtectionDefPullReq
	}{
		Bypass: &code.ProtectionDefBypass{
			RepoOwners: false,
			UserIds:    []int32{},
		},
		Lifecycle: &code.ProtectionDefLifecycle{
			CreateForbidden: rules["block_branch_creation"].(bool),
			UpdateForbidden: false,
			DeleteForbidden: rules["block_branch_deletion"].(bool),
		},
		Pullreq: &code.ProtectionDefPullReq{
			Approvals: &code.ProtectionDefApprovals{
				RequireCodeOwners:   rules["require_code_owner_review"].(bool),
				RequireLatestCommit: rules["require_approval_of_new_changes"].(bool),
				RequireMinimumCount: rules["require_min_reviewers"].(int32),
			},
			Comments: &code.ProtectionDefComments{
				RequireResolveAll: rules["require_comment_resolution"].(bool),
			},
			Merge: &code.ProtectionDefMerge{
				DeleteBranch: rules["auto_delete_branch_on_merge"].(bool),
			},
			StatusChecks: &code.ProtectionDefStatusChecks{
				RequireIdentifiers: []string{},
			},
		},
	}),
	}
}

func buildRepoImportBody(d *schema.ResourceData) *code.ReposImportBody {
	importBody := &code.ReposImportBody{
		Description: d.Get("description").(string),
		Identifier:  d.Get("identifier").(string),
	}

	source := getSource(d)
	if source != nil {
		providerType := code.ImporterProviderType(strings.ToLower(source["type"].(string)))
		importBody.Provider = &code.ImporterProvider{
			Host:     source["host"].(string),
			Password: source["password"].(string),
			Type_:    &providerType,
			Username: source["username"].(string),
		}
		importBody.ProviderRepo = source["repo"].(string)
	}

	return importBody
}

func getSource(d *schema.ResourceData) map[string]interface{} {
	srcSet := d.Get("source").(*schema.Set)
	srcList := srcSet.List()
	if len(srcList) == 0 {
		return nil
	}
	srcElem := srcList[0]
	return srcElem.(map[string]interface{})
}

func readRule(d *schema.ResourceData, repo *code.OpenapiRule, orgId string, projectId string) {
	d.SetId(repo.Identifier)
	d.Set("org_id", orgId)
	d.Set("project_id", projectId)
	d.Set("name", repo.Identifier)
	d.Set("identifier", repo.Identifier)
	d.Set("created", repo.Created)
	d.Set("created_by", repo.CreatedBy)
	d.Set("default_branch", repo.DefaultBranch)
	d.Set("description", repo.Description)
	d.Set("git_url", repo.GitUrl)
	d.Set("importing", repo.Importing)
	// d.Set("is_public", repo.IsPublic)
	d.Set("path", repo.Path)
	d.Set("updated", repo.Updated)
}

func createSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"repoIdentifier": {
			Description: "Repo identifier of the repository.",
			Type:        schema.TypeString,
			Required:    true,
		},
		"name": {
			Description: "Name of the rule.",
			Type:        schema.TypeString,
			Required:    true,
		},
		"description": {
			Description: "Description of the rule.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"target_patterns": {
			Description:  "Pattern of branch to which rule will apply",
			Type:         schema.TypeSet,
			Required:     true,
			AtLeastOneOf: []string{"default_branch", "include", "exclude"},
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"default_branch": {
						Description: "Should rule apply to default branch of the repository",
						Type:        schema.TypeBool,
						Optional:    true,
					},
					"include": {
						Description: "Globstar branch patterns on which rules will be applied",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"exclude": {
						Description: "Globstar branch patterns on which rules will NOT be applied",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"bypass_list": {
			Description:  "List of users who can bypass this rule.",
			Type:         schema.TypeSet,
			Required:     true,
			AtLeastOneOf: []string{"repo_owners", "users"},
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"repo_owners": {
						Description: "Allow users with repository edit permission to bypass.",
						Type:        schema.TypeBool,
						Optional:    true,
					},
					"users": {
						Description: "List of user ids with who can bypass.",
						Type:        schema.TypeList,
						Optional:    true,
					},
				},
			},
		},
		"rules": {
			Description: "Rules to be applied on the repository.",
			Type:        schema.TypeSet,
			Required:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"block_branch_creation": {
						Description: "Only allow users with bypass permission to create matching branches.",
						Type:        schema.TypeBool,
						Optional:    true,
					},
					"block_branch_deletion": {
						Description: "Only allow users with bypass permission to delete matching branches.",
						Type:        schema.TypeBool,
						Optional:    true,
					},
					"require_pull_request": {
						Description: "Do not allow any changes to matching branches without a pull request.",
						Type:        schema.TypeBool,
						Optional:    true,
					},
					"require_min_reviewers": {
						Description: "Require approval on pull requests from a minimum number of reviewers.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"require_code_owner_review": {
						Description: "Require approval on pull requests from one reviewer for each Code Owner rule.",
						Type:        schema.TypeBool,
						Optional:    true,
					},
					"require_approval_of_new_changes": {
						Description: "Require re-approval when there are new changes in the pull request.",
						Type:        schema.TypeBool,
						Optional:    true,
					},
					"require_comment_resolution": {
						Description: "All comments on a pull request must be resolved before it can be merged.",
						Type:        schema.TypeBool,
						Optional:    true,
					},
					"require_status_check_to_pass": {
						Description: "Selected status checks must pass before a pull request can be merged.",
						Type:        schema.TypeList,
						Optional:    true,
					},
					"limit_merge_strategies": {
						Description: "Limit which merge strategies are available to merge a pull request(One of squash, rebase, merge).",
						Type:        schema.TypeList,
						Optional:    true,
					},
					"auto_delete_branch_on_merge": {
						Description: "Automatically delete the source branch of a pull request after it is merged.",
						Type:        schema.TypeBool,
						Optional:    true,
					},
				},
			},
		},
		"created_by": {
			Description: "ID of the user who created the rule.",
			Type:        schema.TypeInt,
			Computed:    true,
		},
		"created": {
			Description: "Timestamp when the rule was created.",
			Type:        schema.TypeInt,
			Computed:    true,
		},
		"updated_by": {
			Description: "ID of the user who updated the rule.",
			Type:        schema.TypeInt,
			Computed:    true,
		},
		"updated": {
			Description: "Timestamp when the rule was updated.",
			Type:        schema.TypeInt,
			Computed:    true,
		},
		"id": {
			Description: "Internal ID of the rule.",
			Type:        schema.TypeString,
			Computed:    true,
		},
	}
}
