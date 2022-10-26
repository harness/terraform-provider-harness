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
		Description: "Data Source for Harness Gitops Repositories.",

		ReadContext: dataSourceGitOpsRepositoryRead,

		Schema: map[string]*schema.Schema{
			"account_id": {
				Description: "account identifier of the cluster.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"project_id": {
				Description: "project identifier of the cluster.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"org_id": {
				Description: "organization identifier of the cluster.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"agent_id": {
				Description: "agent identifier of the cluster.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"identifier": {
				Description: "identifier of the cluster.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"repo": {
				Description: "Repo Details that need to be stored.",
				Type:        schema.TypeList,
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"repo": {
							Description: "Repo Url.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"username": {
							Description: "Username of the repo.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"password": {
							Description: "Password of the repo.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"ssh_private_key": {
							Description: "ssh private key of the repo.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"insecure_ignore_host_key": {
							Description: "insecure ignore for host key of the repo.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"insecure": {
							Description: "insecure connection of the repo.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"enable_lfs": {
							Description: "is lfs enabled of the repo.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"tls_client_cert_data": {
							Description: "tls client certificate data of the repo.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"tls_client_cert_key": {
							Description: "tls client certificate key of the repo.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"type_": {
							Description: "Type of the repo.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Description: "Name of the repo.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"inherited_creds": {
							Description: "are the credentials inherited for the repo.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"enable_oci": {
							Description: "enable OCI for the repo.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"github_app_private_key": {
							Description: "GitHub App Private Key for the repo.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"github_app_id": {
							Description: "GitHub App ID for the repo.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"github_app_installation_id": {
							Description: "GitHub App ID for the repo.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"github_app_enterprise_base_url": {
							Description: "GitHub App Enterprise base Url for the repo.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"proxy": {
							Description: "Proxy used to connect to the repo if any.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"project": {
							Description: "Project of the Repo.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"connection_type": {
							Description: "Connection type for connecting to the Repo.",
							Type:        schema.TypeString,
							Required:    true,
						},
					},
				},
			},
			"upsert": {
				Description: "Upsert the Repo Details.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"creds_only": {
				Description: "Credentials only of the Repo.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"query_repo": {
				Description: "Repo to Query.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"query_project": {
				Description: "Project to Query for Repo.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"query_force_refresh": {
				Description: "Force refresh query for Repo.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"update_mask": {
				Description: "Update mask of the Repository.",
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
