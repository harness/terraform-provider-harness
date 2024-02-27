package repo_rule_branch

import (
	"context"
	"net/http"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/code"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceRepoRuleBranch() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating a Harness Repo Rule.",
		ReadContext:   resourceRepoRuleRead,
		CreateContext: resourceRepoRuleCreateOrUpdate,
		UpdateContext: resourceRepoRuleCreateOrUpdate,
		DeleteContext: resourceRepoRuleDelete,
		Importer:      helpers.RepoRuleResourceImporter,

		Schema: createSchema(),
	}

	helpers.SetMultiLevelDatasourceSchemaWithoutCommonFields(resource.Schema)

	return resource
}

func resourceRepoRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetCodeClientWithContext(ctx)

	orgID := helpers.BuildField(d, "org_id")
	projectID := helpers.BuildField(d, "project_id")
	repoIdentifier := helpers.BuildField(d, "repo_identifier")

	rule, resp, err := c.RepositoryApi.RuleGet(
		ctx,
		c.AccountId,
		repoIdentifier.Value(),
		d.Get("identifier").(string),
		&code.RepositoryApiRuleGetOpts{
			OrgIdentifier:     orgID,
			ProjectIdentifier: projectID,
		},
	)
	if err != nil {
		return helpers.HandleReadApiError(err, d, resp)
	}

	readRule(d, &rule, orgID.Value(), projectID.Value(), repoIdentifier.Value())
	return nil
}

func resourceRepoRuleCreateOrUpdate(
	ctx context.Context,
	d *schema.ResourceData,
	meta interface{},
) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetCodeClientWithContext(ctx)
	var err error
	var rule code.OpenapiRule
	var resp *http.Response

	id := d.Id()
	orgID := helpers.BuildField(d, "org_id")
	projectID := helpers.BuildField(d, "project_id")
	repoIdentifier := d.Get("repo_identifier").(string)
	body := buildRuleBody(d)
	// determine what type of write the change requires
	if id != "" {
		// update rule
		rule, resp, err = c.RepositoryApi.RuleUpdate(
			ctx,
			c.AccountId,
			repoIdentifier,
			id,
			&code.RepositoryApiRuleUpdateOpts{
				Body:              optional.NewInterface(body),
				OrgIdentifier:     orgID,
				ProjectIdentifier: projectID,
			},
		)
	} else {
		// create rule
		rule, resp, err = c.RepositoryApi.RuleAdd(
			ctx,
			c.AccountId,
			repoIdentifier,
			&code.RepositoryApiRuleAddOpts{
				Body:              optional.NewInterface(body),
				OrgIdentifier:     orgID,
				ProjectIdentifier: projectID,
			},
		)
	}
	if err != nil {
		return helpers.HandleApiError(err, d, resp)
	}

	readRule(d, &rule, orgID.Value(), projectID.Value(), repoIdentifier)
	return nil
}

func resourceRepoRuleDelete(
	ctx context.Context,
	d *schema.ResourceData,
	meta interface{},
) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetCodeClientWithContext(ctx)

	resp, err := c.RepositoryApi.RuleDelete(
		ctx,
		c.AccountId,
		d.Get("repo_identifier").(string),
		d.Get("identifier").(string),
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
	ruleType := code.BRANCH_OpenapiRuleType
	state := convertState(d.Get("state").(string))
	return &code.OpenapiRule{
		Definition:  buildRuleDef(d),
		Description: d.Get("description").(string),
		Identifier:  d.Get("identifier").(string),
		Pattern:     buildPattern(d),
		State:       &state,
		Type_:       &ruleType,
	}
}

func buildPattern(d *schema.ResourceData) *code.ProtectionPattern {
	patterns := extractList(d, "target_patterns")
	return &code.ProtectionPattern{
		Default_: patterns["default_branch"].(bool),
		Exclude:  interfaceToString(patterns["exclude"].([]interface{})),
		Include:  interfaceToString(patterns["include"].([]interface{})),
	}
}

func convertState(s string) code.EnumRuleState {
	switch s {
	case "active":
		return code.ACTIVE_EnumRuleState
	case "monitor":
		return code.MONITOR_EnumRuleState
	case "disabled":
		return code.MONITOR_EnumRuleState
	}
	return code.ACTIVE_EnumRuleState
}

func buildRuleDef(d *schema.ResourceData) *code.OpenapiRuleDefinition {
	rules := extractList(d, "rules")
	bypass := extractList(d, "bypass_list")

	// if rules != nil {
	rule := &code.OpenapiRuleDefinition{ProtectionBranch: code.ProtectionBranch(struct {
		Bypass    *code.ProtectionDefBypass
		Lifecycle *code.ProtectionDefLifecycle
		Pullreq   *code.ProtectionDefPullReq
	}{
		Bypass: &code.ProtectionDefBypass{
			RepoOwners: false,
			UserIds:    []int32{},
		},
		Lifecycle: &code.ProtectionDefLifecycle{},
		Pullreq:   &code.ProtectionDefPullReq{},
	})}
	// {
	// 	Lifecycle: &code.ProtectionDefLifecycle{
	// 		CreateForbidden: rules["block_branch_creation"].(bool),
	// 		UpdateForbidden: rules["require_pull_request"].(bool),
	// 		DeleteForbidden: rules["block_branch_deletion"].(bool),
	// 	},
	// 	Pullreq: &code.ProtectionDefPullReq{
	// 		Approvals: &code.ProtectionDefApprovals{
	// 			RequireCodeOwners:   rules["require_code_owner_review"].(bool),
	// 			RequireLatestCommit: rules["require_approval_of_new_changes"].(bool),
	// 			RequireMinimumCount: int32(rules["require_min_reviewers"].(int)),
	// 		},
	// 		Comments: &code.ProtectionDefComments{
	// 			RequireResolveAll: rules["require_comment_resolution"].(bool),
	// 		},
	// 		Merge: &code.ProtectionDefMerge{
	// 			DeleteBranch:      rules["auto_delete_branch_on_merge"].(bool),
	// 			StrategiesAllowed: convertToEnumMergeMethod(rules["limit_merge_strategies"].([]interface{})),
	// 		},
	// 		StatusChecks: &code.ProtectionDefStatusChecks{
	// 			RequireIdentifiers: convertToString(rules["require_status_check_to_pass"].([]interface{})),
	// 		},
	// 	},
	// }),
	// }
	// }
	rule.Lifecycle = &code.ProtectionDefLifecycle{
		CreateForbidden: rules["block_branch_creation"].(bool),
		UpdateForbidden: rules["require_pull_request"].(bool),
		DeleteForbidden: rules["block_branch_deletion"].(bool),
	}
	rule.Pullreq = &code.ProtectionDefPullReq{
		Approvals: &code.ProtectionDefApprovals{
			RequireCodeOwners:   rules["require_code_owner_review"].(bool),
			RequireLatestCommit: rules["require_approval_of_new_changes"].(bool),
			RequireMinimumCount: int32(rules["require_min_reviewers"].(int)),
		},
		Comments: &code.ProtectionDefComments{
			RequireResolveAll: rules["require_comment_resolution"].(bool),
		},
		Merge: &code.ProtectionDefMerge{
			DeleteBranch:      rules["auto_delete_branch_on_merge"].(bool),
			StrategiesAllowed: convertToEnumMergeMethod(rules["limit_merge_strategies"].([]interface{})),
		},
		StatusChecks: &code.ProtectionDefStatusChecks{
			RequireIdentifiers: convertToString(rules["require_status_check_to_pass"].([]interface{})),
		},
	}

	if bypass != nil {
		rule.Bypass.RepoOwners = bypass["repo_owners"].(bool)
		rule.Bypass.UserIds = convertToInt32(bypass["users"].([]interface{}))
	}
	return rule
}

func convertToEnumMergeMethod(s []interface{}) []code.EnumMergeMethod {
	list := make([]code.EnumMergeMethod, len(s))

	for _, v := range s {
		switch v.(string) {
		case "merge":
			list = append(list, code.MERGE_EnumMergeMethod)
		case "rebase":
			list = append(list, code.REBASE_EnumMergeMethod)
		case "squash":
			list = append(list, code.SQUASH_EnumMergeMethod)
		}
	}
	return list
}

func convertToInt32(i []interface{}) []int32 {
	list := make([]int32, len(i))

	for _, v := range i {
		list = append(list, v.(int32))
	}

	return list
}

func convertToString(i []interface{}) []string {
	list := make([]string, len(i))

	for _, v := range i {
		list = append(list, v.(string))
	}

	return list
}

func extractList(d *schema.ResourceData, key string) map[string]interface{} {
	set := d.Get(key).(*schema.Set)
	list := set.List()
	if len(list) == 0 {
		return *new(map[string]interface{})
	}
	elem := list[0]
	return elem.(map[string]interface{})
}

func readRule(d *schema.ResourceData, rule *code.OpenapiRule, orgId string, projectId string, repoIdentifier string) {
	d.SetId(rule.Identifier)
	d.Set("org_id", orgId)
	d.Set("project_id", projectId)
	d.Set("repo_identifier", repoIdentifier)
	d.Set("identifier", rule.Identifier)
	d.Set("created", rule.Created)
	d.Set("created_by", rule.CreatedBy.Id)
	d.Set("description", rule.Description)
	d.Set("state", rule.State)
	d.Set("updated", rule.Updated)
	if rule.Pattern != nil {
		patternList := []interface{}{}
		pattern := map[string]interface{}{}
		pattern["include"] = rule.Pattern.Include
		pattern["exclude"] = rule.Pattern.Exclude
		pattern["default_branch"] = rule.Pattern.Default_
		patternList = append(patternList, pattern)
		d.Set("target_patterns", patternList)
	}
	bypassList := []interface{}{}
	bypass := map[string]interface{}{}
	bypass["repo_owners"] = rule.Definition.Bypass.RepoOwners
	bypass["users"] = rule.Definition.Bypass.UserIds
	bypassList = append(bypassList, bypass)
	d.Set("bypass_list", bypassList)

	rulesList := []interface{}{}
	rules := map[string]interface{}{}
	rules["block_branch_creation"] = rule.Definition.Lifecycle.CreateForbidden
	rules["block_branch_deletion"] = rule.Definition.Lifecycle.DeleteForbidden
	rules["require_pull_request"] = rule.Definition.Lifecycle.UpdateForbidden
	rules["require_min_reviewers"] = rule.Definition.Pullreq.Approvals.RequireMinimumCount
	rules["require_code_owner_review"] = rule.Definition.Pullreq.Approvals.RequireCodeOwners
	rules["require_approval_of_new_changes"] = rule.Definition.Pullreq.Approvals.RequireLatestCommit
	rules["require_comment_resolution"] = rule.Definition.Pullreq.Comments.RequireResolveAll
	rules["require_status_check_to_pass"] = rule.Definition.Pullreq.StatusChecks.RequireIdentifiers
	rules["limit_merge_strategies"] = rule.Definition.Pullreq.Merge.StrategiesAllowed
	rules["auto_delete_branch_on_merge"] = rule.Definition.Pullreq.Merge.DeleteBranch
	rulesList = append(rulesList, rules)
	d.Set("rules", rulesList)
}

func interfaceToString(ds []interface{}) []string {
	list := make([]string, 0)

	for _, v := range ds {
		list = append(list, v.(string))
	}

	return list
}

func createSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"repo_identifier": {
			Description: "Repo identifier of the repository.",
			Type:        schema.TypeString,
			Required:    true,
		},
		"identifier": {
			Description: "Identifier of the rule.",
			Type:        schema.TypeString,
			Required:    true,
		},
		"description": {
			Description: "Description of the rule.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"state": {
			Description: "State of the rule (active, disable, monitor).",
			Type:        schema.TypeString,
			Required:    true,
		},
		"target_patterns": {
			Description: "Pattern of branch to which rule will apply",
			Type:        schema.TypeSet,
			Required:    true,
			//AtLeastOneOf: []string{"default_branch", "include", "exclude"},
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"default_branch": {
						Description: "Should rule apply to default branch of the repository",
						Type:        schema.TypeBool,
						Optional:    true,
					},
					"include": {
						Description: "Globstar branch patterns on which rules will be applied",
						Type:        schema.TypeList,
						Optional:    true,
						Elem:        &schema.Schema{Type: schema.TypeString},
					},
					"exclude": {
						Description: "Globstar branch patterns on which rules will NOT be applied",
						Type:        schema.TypeList,
						Optional:    true,
						Elem:        &schema.Schema{Type: schema.TypeString},
					},
				},
			},
		},
		"bypass_list": {
			Description: "List of users who can bypass this rule.",
			Type:        schema.TypeSet,
			Required:    true,
			//AtLeastOneOf: []string{"repo_owners", "users"},
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
						Elem:        &schema.Schema{Type: schema.TypeString},
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
						Elem:        &schema.Schema{Type: schema.TypeString},
					},
					"limit_merge_strategies": {
						Description: "Limit which merge strategies are available to merge a pull request(One of squash, rebase, merge).",
						Type:        schema.TypeList,
						Optional:    true,
						Elem:        &schema.Schema{Type: schema.TypeString},
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
