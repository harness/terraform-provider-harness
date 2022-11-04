package repository

import (
	"context"
	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceGitopsRepository() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data Source for fetching a Harness GitOps Repository.",

		ReadContext: dataSourceGitOpsRepositoryRead,

		Schema: map[string]*schema.Schema{
			"account_id": {
				Description: "Account identifier of the GitOps repository.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"project_id": {
				Description: "Project identifier of the GitOps repository.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"org_id": {
				Description: "Organization identifier of the GitOps repository.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"agent_id": {
				Description: "Agent identifier of the GitOps repository.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"identifier": {
				Description: "Identifier of the GitOps repository.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"repo": {
				Description: "Repo details holding application configurations.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"repo": {
							Description: "URL to the remote repository.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"username": {
							Description: "Username used for authenticating at the remote repository.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"password": {
							Description: "Password or PAT used for authenticating at the remote repository.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"ssh_private_key": {
							Description: "PEM data for authenticating at the repo server. Only used with Git repos.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"insecure_ignore_host_key": {
							Description: "Indicates if InsecureIgnoreHostKey should be used. Insecure is favored used only for git repos.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"insecure": {
							Description: "Indicates if the connection to the repository ignores any errors when verifying TLS certificates or SSH host keys.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"enable_lfs": {
							Description: "Indicates if git-lfs support must be enabled for this repo. This is valid only for Git repositories.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"tls_client_cert_data": {
							Description: "Certificate in PEM format for authenticating at the repo server.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"tls_client_cert_key": {
							Description: "Private key in PEM format for authenticating at the repo server.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"type_": {
							Description: "Type specifies the type of the repo. Can be either \"git\" or \"helm. \"git\" is assumed if empty or absent.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Description: "Name to be used for this repo. Only used with Helm repos.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"inherited_creds": {
							Description: "Indicates if the credentials were inherited from a credential set.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"enable_oci": {
							Description: "Indicates if helm-oci support must be enabled for this repo.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"github_app_private_key": {
							Description: "GitHub app private key PEM data.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"github_app_id": {
							Description: "Id of the GitHub app used to access the repo.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"github_app_installation_id": {
							Description: "Installation id of the GitHub app used to access the repo.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"github_app_enterprise_base_url": {
							Description: "Base URL of GitHub Enterprise installation. If left empty, this defaults to https://api.github.com.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"proxy": {
							Description: "The HTTP/HTTPS proxy used to access the repo.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"project": {
							Description: "Reference between project and repository that allow you automatically to be added as item inside SourceRepos project entity.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"connection_type": {
							Description: "Identifies the authentication method used to connect to the repository.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
			"upsert": {
				Description: "Indicates if the GitOps repository should be updated if existing and inserted if not.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"creds_only": {
				Description: "Indicates if to operate on credential set instead of repository.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"query_repo": {
				Description: "GitOps repository to query.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"query_project": {
				Description: "Project to query for the GitOps repo.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"query_force_refresh": {
				Description: "Indicates to force refresh query for repository.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"update_mask": {
				Description: "Update mask of the repository.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"paths": {
							Description: "The set of field mask paths.",
							Optional:    true,
							Type:        schema.TypeList,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
	return resource
}

func dataSourceGitOpsRepositoryRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	var orgIdentifier, projectIdentifier, agentIdentifier, identifier, queryRepo, queryProject string
	var queryForceRefresh bool
	if attr, ok := d.GetOk("org_id"); ok {
		orgIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("project_id"); ok {
		projectIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("agent_id"); ok {
		agentIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("identifier"); ok {
		identifier = attr.(string)
	}
	if attr, ok := d.GetOk("query_repo"); ok {
		queryRepo = attr.(string)
	}
	if attr, ok := d.GetOk("query_project"); ok {
		queryProject = attr.(string)
	}
	if attr, ok := d.GetOk("query_force_refresh"); ok {
		queryForceRefresh = attr.(bool)
	}
	resp, httpResp, err := c.RepositoriesApiService.AgentRepositoryServiceGet(ctx, agentIdentifier, identifier, c.AccountId, &nextgen.RepositoriesApiAgentRepositoryServiceGetOpts{
		OrgIdentifier:     optional.NewString(orgIdentifier),
		ProjectIdentifier: optional.NewString(projectIdentifier),
		QueryRepo:         optional.NewString(queryRepo),
		QueryForceRefresh: optional.NewBool(queryForceRefresh),
		QueryProject:      optional.NewString(queryProject),
	})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}
	// Soft delete lookup error handling
	// https://harness.atlassian.net/browse/PL-23765
	if resp.Repository == nil {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}
	setRepositoryDetails(d, &resp)
	return nil

}
