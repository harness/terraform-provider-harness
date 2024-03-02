package repo_rule_branch

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/code"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceRepoBranchRule() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating a Harness Repo Branch Rule.",
		ReadContext:   resourceRepoBranchRuleRead,
		CreateContext: resourceRepoBranchRuleCreateOrUpdate,
		UpdateContext: resourceRepoBranchRuleCreateOrUpdate,
		DeleteContext: resourceRepoBranchRuleDelete,
		Importer:      helpers.RepoRuleResourceImporter,

		Schema: createSchema(),
	}

	helpers.SetMultiLevelDatasourceSchemaWithoutCommonFields(resource.Schema)

	return resource
}

func resourceRepoBranchRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetCodeClientWithContext(ctx)

	orgID := helpers.BuildField(d, "org_id")
	projectID := helpers.BuildField(d, "project_id")
	repoIdentifier := d.Get("repo_identifier").(string)

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
	if err != nil {
		return helpers.HandleReadApiError(err, d, resp)
	}

	readRepoBranchRule(d, &rule, orgID.Value(), projectID.Value(), repoIdentifier)
	return nil
}

func resourceRepoBranchRuleCreateOrUpdate(
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
	body, err := buildRepoBranchRuleBody(d)
	if err != nil {
		return helpers.HandleApiError(err, d, nil)
	}
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

	readRepoBranchRule(d, &rule, orgID.Value(), projectID.Value(), repoIdentifier)
	return nil
}

func resourceRepoBranchRuleDelete(
	ctx context.Context,
	d *schema.ResourceData,
	meta interface{},
) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetCodeClientWithContext(ctx)

	resp, err := c.RepositoryApi.RuleDelete(
		ctx,
		c.AccountId,
		d.Get("repo_identifier").(string),
		d.Id(),
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

func buildRepoBranchRuleBody(d *schema.ResourceData) (*code.OpenapiRule, error) {
	ruleType := code.BRANCH_OpenapiRuleType
	state := convertState(d.Get("state").(string))
	ruleDef, err := buildRepoBranchRuleDef(d)
	if err != nil {
		return nil, err
	}
	return &code.OpenapiRule{
		Definition:  ruleDef,
		Description: d.Get("description").(string),
		Identifier:  d.Get("identifier").(string),
		Pattern:     buildPattern(d),
		State:       &state,
		Type_:       &ruleType,
	}, nil
}

func buildPattern(d *schema.ResourceData) *code.ProtectionPattern {
	patterns := extractSubSchemaSet(d, "target_patterns")
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

type branchRule struct {
	Bypass    *code.ProtectionDefBypass
	Lifecycle *code.ProtectionDefLifecycle
	Pullreq   *code.ProtectionDefPullReq
}

func buildRepoBranchRuleDef(d *schema.ResourceData) (*code.OpenapiRuleDefinition, error) {
	rules := extractSubSchemaSet(d, "policies")
	bypass := extractSubSchemaSet(d, "bypass_list")

	// if rules != nil {
	rule := &code.OpenapiRuleDefinition{ProtectionBranch: code.ProtectionBranch(branchRule{
		Bypass: &code.ProtectionDefBypass{
			RepoOwners: false,
			UserIds:    []int32{},
		},
		Lifecycle: &code.ProtectionDefLifecycle{},
		Pullreq:   &code.ProtectionDefPullReq{},
	})}

	rule.Lifecycle = &code.ProtectionDefLifecycle{
		CreateForbidden: rules["block_branch_creation"].(bool),
		UpdateForbidden: rules["require_pull_request"].(bool),
		DeleteForbidden: rules["block_branch_deletion"].(bool),
	}
	mergeMethod, err := convertToEnumMergeMethod(rules["limit_merge_strategies"].([]interface{}))
	if err != nil {
		return nil, err
	}
	rule.Pullreq = &code.ProtectionDefPullReq{
		Approvals: &code.ProtectionDefApprovals{
			RequireCodeOwners:      rules["require_code_owners"].(bool),
			RequireLatestCommit:    rules["require_latest_commit"].(bool),
			RequireMinimumCount:    int32(rules["require_minimum_count"].(int)),
			RequireNoChangeRequest: rules["require_no_change_request"].(bool),
		},
		Comments: &code.ProtectionDefComments{
			RequireResolveAll: rules["require_comment_resolution"].(bool),
		},
		Merge: &code.ProtectionDefMerge{
			DeleteBranch:      rules["auto_delete_branch_on_merge"].(bool),
			StrategiesAllowed: mergeMethod,
		},
		StatusChecks: &code.ProtectionDefStatusChecks{
			RequireIdentifiers: convertToString(rules["require_status_check_to_pass"].([]interface{})),
		},
	}

	if bypass != nil {
		rule.Bypass.RepoOwners = bypass["repo_owners"].(bool)
		rule.Bypass.UserIds = convertToInt32Slice(bypass["user_ids"].([]interface{}))
	}
	return rule, nil
}

func convertToEnumMergeMethod(s []interface{}) ([]code.EnumMergeMethod, error) {
	list := make([]code.EnumMergeMethod, len(s))

	for _, v := range s {
		switch strings.ToLower(v.(string)) {
		case "merge":
			list = append(list, code.MERGE_EnumMergeMethod)
		case "rebase":
			list = append(list, code.REBASE_EnumMergeMethod)
		case "squash":
			list = append(list, code.SQUASH_EnumMergeMethod)
		default:
			return list, fmt.Errorf("invalid merge method encountered %s", v)
		}
	}
	return list, nil
}

func convertToInt32Slice(i []interface{}) []int32 {
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

func extractSubSchemaSet(d *schema.ResourceData, key string) map[string]interface{} {
	set := d.Get(key).(*schema.Set)
	list := set.List()
	if len(list) == 0 {
		return *new(map[string]interface{})
	}
	elem := list[0]
	return elem.(map[string]interface{})
}

func readRepoBranchRule(d *schema.ResourceData, rule *code.OpenapiRule, orgId string, projectId string, repoIdentifier string) {
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
		pattern := map[string]interface{}{}
		pattern["include"] = rule.Pattern.Include
		pattern["exclude"] = rule.Pattern.Exclude
		pattern["default_branch"] = rule.Pattern.Default_
		d.Set("target_patterns", []interface{}{pattern})
	}
	bypass := map[string]interface{}{}
	bypass["repo_owners"] = rule.Definition.Bypass.RepoOwners
	bypass["user_ids"] = rule.Definition.Bypass.UserIds
	d.Set("bypass_list", []interface{}{bypass})

	rules := map[string]interface{}{}
	rules["block_branch_creation"] = rule.Definition.Lifecycle.CreateForbidden
	rules["block_branch_deletion"] = rule.Definition.Lifecycle.DeleteForbidden
	rules["require_pull_request"] = rule.Definition.Lifecycle.UpdateForbidden
	rules["require_minimum_count"] = rule.Definition.Pullreq.Approvals.RequireMinimumCount
	rules["require_code_owners"] = rule.Definition.Pullreq.Approvals.RequireCodeOwners
	rules["require_latest_commit"] = rule.Definition.Pullreq.Approvals.RequireLatestCommit
	rules["require_no_change_request"] = rule.Definition.Pullreq.Approvals.RequireNoChangeRequest
	rules["require_comment_resolution"] = rule.Definition.Pullreq.Comments.RequireResolveAll
	rules["require_status_check_to_pass"] = rule.Definition.Pullreq.StatusChecks.RequireIdentifiers
	rules["limit_merge_strategies"] = rule.Definition.Pullreq.Merge.StrategiesAllowed
	rules["auto_delete_branch_on_merge"] = rule.Definition.Pullreq.Merge.DeleteBranch
	d.Set("policies", []interface{}{rules})
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
			Optional:    true,
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
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"repo_owners": {
						Description: "Allow users with repository edit permission to bypass.",
						Type:        schema.TypeBool,
						Optional:    true,
					},
					"user_ids": {
						Description: "List of user ids with who can bypass.",
						Type:        schema.TypeList,
						Optional:    true,
						Elem:        &schema.Schema{Type: schema.TypeString},
					},
				},
			},
		},
		"policies": {
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
					"require_minimum_count": {
						Description: "Require approval on pull requests from a minimum number of reviewers.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"require_code_owners": {
						Description: "Require approval on pull requests from one reviewer for each Code Owner rule.",
						Type:        schema.TypeBool,
						Optional:    true,
					},
					"require_latest_commit": {
						Description: "Require re-approval when there are new changes in the pull request.",
						Type:        schema.TypeBool,
						Optional:    true,
					},
					"require_no_change_request": {
						Description: "Require no request changes by reviewers on pull requests.",
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
	}
}
