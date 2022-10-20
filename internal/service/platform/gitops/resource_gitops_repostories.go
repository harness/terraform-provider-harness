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
	setRespositoryDetails(d, &resp)
	return nil
}

func buildCreateRepoRequest(d *schema.ResourceData) nextgen.RepositoriesRepoCreateRequest {

}

func setRespositoryDetails(d *schema.ResourceData, repo *nextgen.Servicev1Repository) {

}
