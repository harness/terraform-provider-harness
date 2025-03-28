package policy

import (
	"context"
	"net/http"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/policymgmt"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourcePolicy() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for creating a Harness Policy.",

		ReadContext:   resourcePolicyRead,
		UpdateContext: resourcePolicyCreateOrUpdate,
		DeleteContext: resourcePolicyDelete,
		CreateContext: resourcePolicyCreateOrUpdate,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"name": {
				Description: "Name of the policy.",
				Type:        schema.TypeString,
				Required:    true,
				Computed:    false,
			},
			"identifier": {
				Description: "Identifier of the policy.",
				Type:        schema.TypeString,
				Required:    true,
				Computed:    false,
			},
			"rego": {
				Description: "Rego code for the policy.",
				Type:        schema.TypeString,
				Required:    true,
				Computed:    false,
			},
			"git_connector_ref": {
				Description: "Git connector reference for the policy.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"git_path": {
				Description: "Git path for the policy.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"git_repo": {
				Description: "Git repository for the policy.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"git_commit_msg": {
				Description: "Git commit message for the policy.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"git_import": {
				Description: "Flag to import the policy from git.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"git_branch": {
				Description: "Git branch for the policy.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"git_is_new_branch": {
				Description: "Flag to create a new branch for the policy.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"git_base_branch": {
				Description: "Base branch for the new git branch.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"git_commit_sha": {
				Description: "The existing commit sha of the file being updated",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"git_file_id": {
				Description: "The existing file id of the file being updated, not required for bitbucket files",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
		},
	}

	helpers.SetMultiLevelResourceSchema(resource.Schema)

	return resource
}

func resourcePolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*internal.Session).GetPolicyManagementClient()

	id := d.Id()

	localVarOptionals := policymgmt.PoliciesApiPoliciesFindOpts{
		AccountIdentifier: optional.NewString(meta.(*internal.Session).AccountId),
		XApiKey:           optional.NewString(meta.(*internal.Session).PLClient.ApiKey),
	}
	// check for project and org
	if d.Get("project_id").(string) != "" {
		localVarOptionals.ProjectIdentifier = helpers.BuildField(d, "project_id")
	}
	if d.Get("org_id").(string) != "" {
		localVarOptionals.OrgIdentifier = helpers.BuildField(d, "org_id")
	}

	policy, httpResp, err := c.PoliciesApi.PoliciesFind(ctx, id, &localVarOptionals)

	if err != nil {
		return helpers.HandlePolicyApiError(err, d, httpResp)
	}

	readPolicy(d, policy)

	return nil
}

func resourcePolicyCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*internal.Session).GetPolicyManagementClient()
	var err error
	var responsePolicy policymgmt.Policy

	id := d.Id()

	if id == "" {
		body := policymgmt.CreateRequestBody{
			Name:            d.Get("name").(string),
			Rego:            d.Get("rego").(string),
			Identifier:      d.Get("identifier").(string),
			GitConnectorRef: d.Get("git_connector_ref").(string),
			GitPath:         d.Get("git_path").(string),
			GitRepo:         d.Get("git_repo").(string),
		}
		localVarOptionals := policymgmt.PoliciesApiPoliciesCreateOpts{
			AccountIdentifier: optional.NewString(meta.(*internal.Session).AccountId),
			XApiKey:           optional.NewString(meta.(*internal.Session).PLClient.ApiKey),
			GitCommitMsg:      optional.NewString(d.Get("git_commit_msg").(string)),
			GitImport:         optional.NewBool(d.Get("git_import").(bool)),
			GitBranch:         optional.NewString(d.Get("git_branch").(string)),
			GitIsNewBranch:    optional.NewBool(d.Get("git_is_new_branch").(bool)),
			GitBaseBranch:     optional.NewString(d.Get("git_base_branch").(string)),
		}
		// check for project and org
		if d.Get("project_id").(string) != "" {
			localVarOptionals.ProjectIdentifier = helpers.BuildField(d, "project_id")
		}
		if d.Get("org_id").(string) != "" {
			localVarOptionals.OrgIdentifier = helpers.BuildField(d, "org_id")
		}
		responsePolicy, _, err = c.PoliciesApi.PoliciesCreate(ctx, body, &localVarOptionals)
		if err != nil {
			// For create operations, we should return the error directly
			return diag.Errorf("error creating policy: %v", err)
		}
	} else {
		body := policymgmt.UpdateRequestBody{
			Name: d.Get("name").(string),
			Rego: d.Get("rego").(string),
		}
		localVarOptionals := policymgmt.PoliciesApiPoliciesUpdateOpts{
			AccountIdentifier: optional.NewString(meta.(*internal.Session).AccountId),
			XApiKey:           optional.NewString(meta.(*internal.Session).PLClient.ApiKey),
			GitCommitMsg:      optional.NewString(d.Get("git_commit_msg").(string)),
			GitBranch:         optional.NewString(d.Get("git_branch").(string)),
			GitIsNewBranch:    optional.NewBool(d.Get("git_is_new_branch").(bool)),
			GitBaseBranch:     optional.NewString(d.Get("git_base_branch").(string)),
			GitCommitSha:      optional.NewString(d.Get("git_commit_sha").(string)),
			GitFileId:         optional.NewString(d.Get("git_file_id").(string)),
		}
		if d.Get("project_id").(string) != "" {
			localVarOptionals.ProjectIdentifier = helpers.BuildField(d, "project_id")
		}
		if d.Get("org_id").(string) != "" {
			localVarOptionals.OrgIdentifier = helpers.BuildField(d, "org_id")
		}
		updateResp, err := c.PoliciesApi.PoliciesUpdate(ctx, body, id, &localVarOptionals)
		if err != nil {
			// For update operations, if we get a 404, we should trigger recreation
			if updateResp != nil && updateResp.StatusCode == http.StatusNotFound {
				d.SetId("")
				return nil
			}
			return diag.Errorf("error updating policy: %v", err)
		}

		if updateResp.StatusCode == http.StatusNoContent {
			// if we get a 204, we need to get the policy again to get the updated values
			findLocalVarOptionals := policymgmt.PoliciesApiPoliciesFindOpts{
				AccountIdentifier: optional.NewString(meta.(*internal.Session).AccountId),
				XApiKey:           optional.NewString(meta.(*internal.Session).PLClient.ApiKey),
			}
			// check for project and org
			if d.Get("project_id").(string) != "" {
				findLocalVarOptionals.ProjectIdentifier = helpers.BuildField(d, "project_id")
			}
			if d.Get("org_id").(string) != "" {
				findLocalVarOptionals.OrgIdentifier = helpers.BuildField(d, "org_id")
			}
			var findResp *http.Response
			responsePolicy, findResp, err = c.PoliciesApi.PoliciesFind(ctx, id, &findLocalVarOptionals)
			if err != nil {
				// Check error message first
				if err.Error() == "Not Found" {
					d.SetId("")
					return nil
				}
				// Only access findResp if it's not nil
				if findResp != nil && findResp.StatusCode == http.StatusNotFound {
					d.SetId("")
					return nil
				}
				return diag.Errorf("error reading policy after update: %v", err)
			}

		}
	}

	readPolicy(d, responsePolicy)
	return nil
}

func resourcePolicyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*internal.Session).GetPolicyManagementClient()

	localVarOptionals := policymgmt.PoliciesApiPoliciesDeleteOpts{
		AccountIdentifier: optional.NewString(meta.(*internal.Session).AccountId),
		XApiKey:           optional.NewString(meta.(*internal.Session).PLClient.ApiKey),
	}
	// check for project and org
	if d.Get("project_id").(string) != "" {
		localVarOptionals.ProjectIdentifier = helpers.BuildField(d, "project_id")
	}
	if d.Get("org_id").(string) != "" {
		localVarOptionals.OrgIdentifier = helpers.BuildField(d, "org_id")
	}
	httpResp, err := c.PoliciesApi.PoliciesDelete(ctx, d.Id(), &localVarOptionals)

	if err != nil {
		return helpers.HandlePolicyApiError(err, d, httpResp)
	}

	return nil
}

func readPolicy(d *schema.ResourceData, policy policymgmt.Policy) {
	d.SetId(policy.Identifier)
	_ = d.Set("identifier", policy.Identifier)
	_ = d.Set("org_id", policy.OrgId)
	_ = d.Set("project_id", policy.ProjectId)
	_ = d.Set("name", policy.Name)
	_ = d.Set("rego", policy.Rego)

	if policy.GitCommitSha != "" {
		_ = d.Set("git_commit_sha", policy.GitCommitSha)
	}
	if policy.GitConnectorRef != "" {
		_ = d.Set("git_connector_ref", policy.GitConnectorRef)
	}
	if policy.GitFileId != "" {
		_ = d.Set("git_file_id", policy.GitFileId)
	}
	if policy.GitPath != "" {
		_ = d.Set("git_path", policy.GitPath)
	}
	if policy.GitRepo != "" {
		_ = d.Set("git_repo", policy.GitRepo)
	}
}
