package repository_credentials

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

func DataSourceGitOpsRepoCred() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for fetching a GitOps Repository Credential.",

		ReadContext: dataSourceGitopsRepoCredRead,

		Schema: map[string]*schema.Schema{
			"agent_id": {
				Description: "Agent identifier of the Repository Credential.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"account_id": {
				Description: "Account identifier of the Repository Credential.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"identifier": {
				Description: "Identifier of the Repository Credential.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"org_id": {
				Description: "Organization identifier of the Repository Credential.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"project_id": {
				Description: "Project identifier of the Repository Credential.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"creds": {
				Description: "credential details.",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"url": {
							Description: "URL of the remote repository. Make sure you pass at least an org, this will not work if you just provide the host, for eg. \"https://github.com\"",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"username": {
							Description: "Username to be used for authenticating the remote repository.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"password": {
							Description: "Password or PAT to be used for authenticating the remote repository.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"ssh_private_key": {
							Description: "SSH Key in PEM format for authenticating the repository. Used only for Git repository.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"tls_client_cert_data": {
							Description: "Certificate in PEM format for authenticating at the repo server. This is used for mTLS.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"tls_client_cert_key": {
							Description: "Private key in PEM format for authenticating at the repo server. This is used for mTLS.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"github_app_private_key": {
							Description: "github_app_private_key specifies the private key PEM data for authentication via GitHub app.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"github_app_id": {
							Description: "Specifies the Github App ID of the app used to access the repo for GitHub app authentication.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"github_app_installation_id": {
							Description: "Specifies the ID of the installed GitHub App for GitHub app authentication.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"github_app_enterprise_base_url": {
							Description: "Specifies the GitHub API URL for GitHub app authentication.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"enable_oci": {
							Description: "Specifies whether helm-oci support should be enabled for this repo.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"type": {
							Description: "Type specifies the type of the repoCreds.Can be either 'git' or 'helm. 'git' is assumed if empty or absent",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
		},
	}

	return resource
}

func dataSourceGitopsRepoCredRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	ctx = context.WithValue(ctx, nextgen.ContextAccessToken, hh.EnvVars.BearerToken.Get())

	agentIdentifier := d.Get("agent_id").(string)
	identifier := d.Get("identifier").(string)

	resp, httpResp, err := c.RepositoryCredentialsApi.AgentRepositoryCredentialsServiceGetRepositoryCredentials(ctx, agentIdentifier, identifier, c.AccountId, &nextgen.RepositoryCredentialsApiAgentRepositoryCredentialsServiceGetRepositoryCredentialsOpts{
		OrgIdentifier:     optional.NewString(d.Get("org_id").(string)),
		ProjectIdentifier: optional.NewString(d.Get("project_id").(string)),
	})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	// Soft delete lookup error handling
	// https://harness.atlassian.net/browse/PL-23765
	if &resp == nil {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}
	setGitopsRepositoriesCredential(d, &resp)
	return nil
}
