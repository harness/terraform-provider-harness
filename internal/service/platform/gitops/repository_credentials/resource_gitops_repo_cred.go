package repository_credentials

import (
	"context"
	"net/http"

	"github.com/antihax/optional"
	hh "github.com/harness/harness-go-sdk/harness/helpers"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceGitopsRepoCred() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for managing a Harness Gitops Repository Credentials.",

		CreateContext: resourceGitopsRepoCredCreate,
		ReadContext:   resourceGitopsRepoCredRead,
		UpdateContext: resourceGitopsRepoCredUpdate,
		DeleteContext: resourceGitopsRepoCredDelete,
		Importer:      helpers.GitopsAgentResourceImporter,

		Schema: map[string]*schema.Schema{
			"identifier": {
				Description: "Identifier of the Repository Credentials.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"agent_id": {
				Description: "Agent identifier of the Repository Credentials.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"account_id": {
				Description: "Account identifier of the Repository Credentials.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"org_id": {
				Description: "Organization identifier of the Repository Credentials.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"project_id": {
				Description: "Project identifier of the Repository Credentials.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"upsert": {
				Description: "Indicates if the GitOps repository credential should be updated if existing and inserted if not.",
				Type:        schema.TypeBool,
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

func resourceGitopsRepoCredCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	ctx = context.WithValue(ctx, nextgen.ContextAccessToken, hh.EnvVars.BearerToken.Get())

	var agentIdentifier, accountIdentifier, orgIdentifier, projectIdentifier, identifier string
	accountIdentifier = c.AccountId

	if attr, ok := d.GetOk("identifier"); ok {
		identifier = attr.(string)
	}
	if attr, ok := d.GetOk("agent_id"); ok {
		agentIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("org_id"); ok {
		orgIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("project_id"); ok {
		projectIdentifier = attr.(string)
	}

	var err error
	var httpResp *http.Response
	var resp nextgen.Servicev1RepositoryCredentials

	repoCredCreateRequest := buildRepoCredCreateRequest(d)
	resp, httpResp, err = c.RepositoryCredentialsApi.AgentRepositoryCredentialsServiceCreateRepositoryCredentials(ctx, *repoCredCreateRequest, accountIdentifier, agentIdentifier,
		&nextgen.RepositoryCredentialsApiAgentRepositoryCredentialsServiceCreateRepositoryCredentialsOpts{
			Identifier:        optional.NewString(identifier),
			OrgIdentifier:     optional.NewString(orgIdentifier),
			ProjectIdentifier: optional.NewString(projectIdentifier),
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

func resourceGitopsRepoCredUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	ctx = context.WithValue(ctx, nextgen.ContextAccessToken, hh.EnvVars.BearerToken.Get())

	var agentIdentifier, accountIdentifier, orgIdentifier, projectIdentifier, identifier string
	accountIdentifier = c.AccountId

	if attr, ok := d.GetOk("identifier"); ok {
		identifier = attr.(string)
	}
	if attr, ok := d.GetOk("agent_id"); ok {
		agentIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("org_id"); ok {
		orgIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("project_id"); ok {
		projectIdentifier = attr.(string)
	}

	var err error
	var httpResp *http.Response
	var resp nextgen.Servicev1RepositoryCredentials

	repoCredUpdateRequest := buildRepoCredUpdateRequest(d)
	resp, httpResp, err = c.RepositoryCredentialsApi.AgentRepositoryCredentialsServiceUpdateRepositoryCredentials(ctx, *repoCredUpdateRequest, agentIdentifier, identifier,
		&nextgen.RepositoryCredentialsApiAgentRepositoryCredentialsServiceUpdateRepositoryCredentialsOpts{
			AccountIdentifier: optional.NewString(accountIdentifier),
			OrgIdentifier:     optional.NewString(orgIdentifier),
			ProjectIdentifier: optional.NewString(projectIdentifier),
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

func resourceGitopsRepoCredRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	ctx = context.WithValue(ctx, nextgen.ContextAccessToken, hh.EnvVars.BearerToken.Get())

	agentIdentifier := d.Get("agent_id").(string)
	identifier := d.Get("identifier").(string)

	resp, httpResp, err := c.RepositoryCredentialsApi.AgentRepositoryCredentialsServiceGetRepositoryCredentials(ctx, agentIdentifier, identifier, c.AccountId, &nextgen.RepositoryCredentialsApiAgentRepositoryCredentialsServiceGetRepositoryCredentialsOpts{
		OrgIdentifier:     optional.NewString(d.Get("org_id").(string)),
		ProjectIdentifier: optional.NewString(d.Get("project_id").(string)),
	})

	if err != nil {
		return helpers.HandleReadApiError(err, d, httpResp)
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

func resourceGitopsRepoCredDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	ctx = context.WithValue(ctx, nextgen.ContextAccessToken, hh.EnvVars.BearerToken.Get())
	agentIdentifier := d.Get("agent_id").(string)
	identifier := d.Get("identifier").(string)

	_, httpResp, err := c.RepositoryCredentialsApi.AgentRepositoryCredentialsServiceDeleteRepositoryCredentials(ctx, agentIdentifier, identifier, &nextgen.RepositoryCredentialsApiAgentRepositoryCredentialsServiceDeleteRepositoryCredentialsOpts{
		AccountIdentifier:  optional.NewString(c.AccountId),
		OrgIdentifier:      optional.NewString(d.Get("org_id").(string)),
		ProjectIdentifier:  optional.NewString(d.Get("project_id").(string)),
		QueryUrl:           optional.NewString(d.Get("creds.0.url").(string)),
		QueryRepoCredsType: optional.NewString(d.Get("creds.0.type").(string)),
	})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}
	return nil
}

func setGitopsRepositoriesCredential(d *schema.ResourceData, repoCred *nextgen.Servicev1RepositoryCredentials) {
	d.SetId(repoCred.Identifier)
	d.Set("account_id", repoCred.AccountIdentifier)
	d.Set("agent_id", repoCred.AgentIdentifier)
	d.Set("identifier", repoCred.Identifier)
	d.Set("org_id", repoCred.OrgIdentifier)
	d.Set("project_id", repoCred.ProjectIdentifier)

	if repoCred.RepoCreds != nil {
		credList := []interface{}{}
		cred := map[string]interface{}{}
		cred["url"] = repoCred.RepoCreds.Url
		cred["username"] = repoCred.RepoCreds.Username
		cred["password"] = repoCred.RepoCreds.Password
		cred["ssh_private_key"] = repoCred.RepoCreds.SshPrivateKey
		cred["tls_client_cert_data"] = repoCred.RepoCreds.TlsClientCertData
		cred["tls_client_cert_key"] = repoCred.RepoCreds.TlsClientCertKey
		cred["github_app_private_key"] = repoCred.RepoCreds.GithubAppPrivateKey
		cred["github_app_id"] = repoCred.RepoCreds.GithubAppID
		cred["github_app_installation_id"] = repoCred.RepoCreds.GithubAppInstallationID
		cred["github_app_enterprise_base_url"] = repoCred.RepoCreds.GithubAppEnterpriseBaseUrl
		cred["enable_oci"] = repoCred.RepoCreds.EnableOCI
		cred["type"] = repoCred.RepoCreds.Type_

		credList = append(credList, cred)
		d.Set("creds", credList)
	}
}

func buildRepoCredCreateRequest(d *schema.ResourceData) *nextgen.HrepocredsRepoCredsCreateRequest {
	var upsert bool
	if attr, ok := d.GetOk("upsert"); ok {
		upsert = attr.(bool)
	}

	return &nextgen.HrepocredsRepoCredsCreateRequest{
		Upsert: upsert,
		Creds:  buildRepoCred(d),
	}
}

func buildRepoCredUpdateRequest(d *schema.ResourceData) *nextgen.HrepocredsRepoCredsUpdateRequest {
	return &nextgen.HrepocredsRepoCredsUpdateRequest{
		Creds: buildRepoCred(d),
	}
}

func buildRepoCred(d *schema.ResourceData) *nextgen.HrepocredsRepoCreds {
	var repoCred nextgen.HrepocredsRepoCreds

	if attr, ok := d.GetOk("creds"); ok {
		if attr != nil && len(attr.([]interface{})) > 0 {
			var requestCreds = attr.([]interface{})[0].(map[string]interface{})

			if requestCreds["url"] != nil {
				repoCred.Url = requestCreds["url"].(string)
			}

			if requestCreds["username"] != nil {
				repoCred.Username = requestCreds["username"].(string)
			}

			if requestCreds["password"] != nil {
				repoCred.Password = requestCreds["password"].(string)
			}

			if requestCreds["ssh_private_key"] != nil {
				repoCred.SshPrivateKey = requestCreds["ssh_private_key"].(string)
			}

			if requestCreds["tls_client_cert_data"] != nil {
				repoCred.TlsClientCertData = requestCreds["tls_client_cert_data"].(string)
			}

			if requestCreds["tls_client_cert_key"] != nil {
				repoCred.TlsClientCertKey = requestCreds["tls_client_cert_key"].(string)
			}

			if requestCreds["github_app_private_key"] != nil {
				repoCred.GithubAppPrivateKey = requestCreds["github_app_private_key"].(string)
			}

			if requestCreds["github_app_id"] != nil {
				repoCred.GithubAppID = requestCreds["github_app_id"].(string)
			}

			if requestCreds["github_app_installation_id"] != nil {
				repoCred.GithubAppInstallationID = requestCreds["github_app_installation_id"].(string)
			}

			if requestCreds["github_app_enterprise_base_url"] != nil {
				repoCred.GithubAppEnterpriseBaseUrl = requestCreds["github_app_enterprise_base_url"].(string)
			}

			if requestCreds["enable_oci"] != nil {
				repoCred.EnableOCI = requestCreds["enable_oci"].(bool)
			}

			if requestCreds["type"] != nil {
				repoCred.Type_ = requestCreds["type"].(string)
			}

		}

	}
	return &repoCred
}
