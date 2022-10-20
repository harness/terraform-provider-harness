package gitops

import (
	"context"
	"github.com/antihax/optional"
	hh "github.com/harness/harness-go-sdk/harness/helpers"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceGitopsRepositories() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for creating a Harness Gitops Cluster.",

		CreateContext: resourceGitOpsRepositoryCreate,
		//ReadContext:   resourceGitopsClusterRead,
		//UpdateContext: resourceGitopsClusterUpdate,
		//DeleteContext: resourceGitopsClusterDelete,
		//Importer:      helpers.GitopsAgentResourceImporter,
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
		},
	}
	return resource
}

func resourceGitOpsRepositoryCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	ctx = context.WithValue(ctx, nextgen.ContextAccessToken, hh.EnvVars.BearerToken.Get())
	var agentIdentifier, accountIdentifier, orgIdentifier, projectIdentifier, identifier string
	accountIdentifier = c.AccountId
	if attr, ok := d.GetOk("agent_id"); ok {
		agentIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("org_id"); ok {
		orgIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("project_id"); ok {
		projectIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("identifier"); ok {
		identifier = attr.(string)
	}

	createRepoRequest := buildCreateRepoRequest(d)
	resp, httpResp, err := c.RepositoriesApiService.AgentRepositoryServiceCreateRepository(ctx, createRepoRequest, agentIdentifier, &nextgen.RepositoriesApiAgentRepositoryServiceCreateRepositoryOpts{
		AccountIdentifier: optional.NewString(accountIdentifier),
		OrgIdentifier:     optional.NewString(orgIdentifier),
		ProjectIdentifier: optional.NewString(projectIdentifier),
		Identifier:        optional.NewString(identifier),
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

func buildCreateRepoRequest(d *schema.ResourceData) nextgen.RepositoriesRepoCreateRequest {
	var upsert, credsOnly bool
	if attr, ok := d.GetOk("upsert"); ok {
		upsert = attr.(bool)
	}
	if attr, ok := d.GetOk("creds_only"); ok {
		credsOnly = attr.(bool)
	}
	return nextgen.RepositoriesRepoCreateRequest{
		Upsert:    upsert,
		CredsOnly: credsOnly,
		Repo:      buildRepo(d),
	}
}

func buildRepo(d *schema.ResourceData) *nextgen.RepositoriesRepository {
	var repoObj = nextgen.RepositoriesRepository{}
	if attr, ok := d.GetOk("repo"); ok {
		if attr != nil && len(attr.([]interface{})) > 0 {
			var repo = attr.([]interface{})[0].(map[string]interface{})
			if repo["repo"] != nil {
				repoObj.Repo = repo["repo"].(string)
			}
			if repo["username"] != nil {
				repoObj.Username = repo["username"].(string)
			}
			if repo["password"] != nil {
				repoObj.Password = repo["password"].(string)
			}
			if repo["ssh_private_key"] != nil {
				repoObj.SshPrivateKey = repo["ssh_private_key"].(string)
			}
			if repo["insecure_ignore_host_key"] != nil {
				repoObj.InsecureIgnoreHostKey = repo["insecure_ignore_host_key"].(bool)
			}
			if repo["insecure"] != nil {
				repoObj.Insecure = repo["insecure"].(bool)
			}
			if repo["enable_lfs"] != nil {
				repoObj.EnableLfs = repo["enable_lfs"].(bool)
			}
			if repo["tls_client_cert_data"] != nil {
				repoObj.TlsClientCertData = repo["tls_client_cert_data"].(string)
			}
			if repo["tls_client_cert_key"] != nil {
				repoObj.TlsClientCertKey = repo["tls_client_cert_key"].(string)
			}
			if repo["type_"] != nil {
				repoObj.Type_ = repo["type_"].(string)
			}
			if repo["name"] != nil {
				repoObj.Name = repo["name"].(string)
			}
			if repo["inherited_creds"] != nil {
				repoObj.InheritedCreds = repo["inherited_creds"].(bool)
			}
			if repo["enable_oci"] != nil {
				repoObj.EnableOCI = repo["enable_oci"].(bool)
			}
			if repo["github_app_private_key"] != nil {
				repoObj.GithubAppPrivateKey = repo["github_app_private_key"].(string)
			}
			if repo["github_app_id"] != nil {
				repoObj.GithubAppID = repo["github_app_id"].(string)
			}
			if repo["github_app_installation_id"] != nil {
				repoObj.GithubAppInstallationID = repo["github_app_installation_id"].(string)
			}
			if repo["github_app_enterprise_base_url"] != nil {
				repoObj.GithubAppEnterpriseBaseUrl = repo["github_app_enterprise_base_url"].(string)
			}
			if repo["proxy"] != nil {
				repoObj.Proxy = repo["proxy"].(string)
			}
			if repo["project"] != nil {
				repoObj.Project = repo["project"].(string)
			}
			if repo["connection_type"] != nil {
				repoObj.ConnectionType = repo["connection_type"].(string)
			}
		}
	}
	return &repoObj
}

func setRepositoryDetails(d *schema.ResourceData, repo *nextgen.Servicev1Repository) {

}
