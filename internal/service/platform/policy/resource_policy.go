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
		Importer:      helpers.OrgResourceImporter,

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
				Description: "The harness connector used for authenticating on the git provider.",
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
			},
			"git_path": {
				Description: "The path to the file in the git repo.",
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
			},
			"git_repo": {
				Description: "The git repo the policy resides in.",
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
			},
			"git_commit_msg": {
				Description: "The commit message used in git when creating the policy.",
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
			},
			"git_import": {
				Description: "A flag to determine if the api should try and import and existing policy from git.",
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
			},
			"git_branch": {
				Description: "The git branch the policy will be created in.",
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
			},
			"git_is_new_branch": {
				Description: "A flag to determine if the api should try and commit to a new branch.",
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
			},
			"git_base_branch": {
				Description: "If committing to a new branch, git_base_branch tells the api which branch to base the new branch from.",
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
			},
			"git_commit_sha": {
				Description: "The existing commit sha of the file being updated.",
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
			},
			"git_file_id": {
				Description: "The existing file id of the file being updated, not required for bitbucket files.",
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
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
		return helpers.HandleApiError(err, d, httpResp)
	}

	readPolicy(d, policy)

	return nil
}

func resourcePolicyCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*internal.Session).GetPolicyManagementClient()
	var err error
	var responsePolicy policymgmt.Policy
	var httpResp *http.Response
	id := d.Id()

	if id == "" {
		body := policymgmt.CreateRequestBody{
			Name:       d.Get("name").(string),
			Rego:       d.Get("rego").(string),
			Identifier: d.Get("identifier").(string),
			GitConnectorRef: d.Get("git_connector_ref").(string),
			GitPath: d.Get("git_path").(string),
			GitRepo: d.Get("git_repo").(string),
		}
		localVarOptionals := policymgmt.PoliciesApiPoliciesCreateOpts{
			AccountIdentifier: optional.NewString(meta.(*internal.Session).AccountId),

			XApiKey: optional.NewString(meta.(*internal.Session).PLClient.ApiKey),
		}
		// check for project and org
		if d.Get("project_id").(string) != "" {
			localVarOptionals.ProjectIdentifier = helpers.BuildField(d, "project_id")
		}
		if d.Get("org_id").(string) != "" {
			localVarOptionals.OrgIdentifier = helpers.BuildField(d, "org_id")
		}
		//check for git details
	    if d.Get("git_commit_msg").(string) != "" {
		    localVarOptionals.OrgIdentifier = helpers.BuildField(d, "git_commit_msg")
	    }
	    if d.Get("git_import").(string) != "" {
		    localVarOptionals.OrgIdentifier = helpers.BuildField(d, "git_import")
	    }
	    if d.Get("git_branch").(string) != "" {
		    localVarOptionals.OrgIdentifier = helpers.BuildField(d, "git_branch")
	    }
	    if d.Get("git_is_new_branch").(string) != "" {
		    localVarOptionals.OrgIdentifier = helpers.BuildField(d, "git_is_new_branch")
	    }
	    if d.Get("git_base_branch").(string) != "" {
		    localVarOptionals.OrgIdentifier = helpers.BuildField(d, "git_base_branch")
	    }

		responsePolicy, httpResp, err = c.PoliciesApi.PoliciesCreate(ctx, body, &localVarOptionals)
	} else {
		body := policymgmt.UpdateRequestBody{
			Name: d.Get("name").(string),
			Rego: d.Get("rego").(string),
		}
		localVarOptionals := policymgmt.PoliciesApiPoliciesUpdateOpts{
			AccountIdentifier: optional.NewString(meta.(*internal.Session).AccountId),
			XApiKey:           optional.NewString(meta.(*internal.Session).PLClient.ApiKey),
		}
		if d.Get("project_id").(string) != "" {
			localVarOptionals.ProjectIdentifier = helpers.BuildField(d, "project_id")
		}
		if d.Get("org_id").(string) != "" {
			localVarOptionals.OrgIdentifier = helpers.BuildField(d, "org_id")
		}
		//check for git details
	    if d.Get("git_commit_msg").(string) != "" {
		    localVarOptionals.OrgIdentifier = helpers.BuildField(d, "git_commit_msg")
	    }
	    if d.Get("git_commit_sha").(string) != "" {
		    localVarOptionals.OrgIdentifier = helpers.BuildField(d, "git_commit_sha")
	    }
	    if d.Get("git_branch").(string) != "" {
		    localVarOptionals.OrgIdentifier = helpers.BuildField(d, "git_branch")
	    }
	    if d.Get("git_is_new_branch").(string) != "" {
		    localVarOptionals.OrgIdentifier = helpers.BuildField(d, "git_is_new_branch")
	    }
	    if d.Get("git_base_branch").(string) != "" {
		    localVarOptionals.OrgIdentifier = helpers.BuildField(d, "git_base_branch")
	    }
		if d.Get("git_file_id").(string) != "" {
		    localVarOptionals.OrgIdentifier = helpers.BuildField(d, "git_file_id")
	    }

		httpResp, err = c.PoliciesApi.PoliciesUpdate(ctx, body, id, &localVarOptionals)
		if err == nil && httpResp.StatusCode == http.StatusNoContent {
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
			responsePolicy, httpResp, err = c.PoliciesApi.PoliciesFind(ctx, id, &findLocalVarOptionals)
		}
	}
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
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
		return helpers.HandleApiError(err, d, httpResp)
	}

	return nil
}

func readPolicy(d *schema.ResourceData, policy policymgmt.Policy) {
	d.SetId(policy.Identifier)
	_ = d.Set("identifier", policy.Identifier)
	_ = d.Set("org_id", policy.OrgId)
	_ = d.Set("account_id", policy.AccountId)
	_ = d.Set("project_id", policy.ProjectId)
	_ = d.Set("name", policy.Name)
	_ = d.Set("rego", policy.Rego)
}
