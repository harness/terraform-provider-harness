package repository_credentials

import (
	"context"
	"net/http"

	hcty "github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

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
				ForceNew:    true,
			},
			"agent_id": {
				Description: "Agent identifier of the Repository Credentials.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"account_id": {
				Description: "Account identifier of the Repository Credentials.",
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				Deprecated:  "This field is deprecated and will be removed in a future release.",
			},
			"org_id": {
				Description: "Organization identifier of the Repository Credentials.",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
			},
			"project_id": {
				Description: "Project identifier of the Repository Credentials.",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
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
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"url": {
							Description: "URL of the remote repository. Make sure you pass at least an org, this will not work if you just provide the host, for eg. \"https://github.com\"",
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"username": {
							Description:   "Username to be used for authenticating the remote repository.",
							Type:          schema.TypeString,
							Optional:      true,
							ConflictsWith: []string{"creds.0.ssh_private_key", "creds.0.tls_client_cert_data", "creds.0.tls_client_cert_key", "creds.0.github_app_private_key", "creds.0.github_app_id", "creds.0.github_app_installation_id", "creds.0.github_app_enterprise_base_url"},
						},
						"password": {
							Description:   "Password or PAT to be used for authenticating the remote repository. Use password_wo for write-only support (Terraform >= 1.11).",
							Type:          schema.TypeString,
							Optional:      true,
							Computed:      true,
							Sensitive:     true,
							ConflictsWith: []string{"creds.0.ssh_private_key", "creds.0.tls_client_cert_data", "creds.0.tls_client_cert_key", "creds.0.github_app_private_key", "creds.0.github_app_id", "creds.0.github_app_installation_id", "creds.0.github_app_enterprise_base_url", "creds.0.password_wo"},
						},
						"password_wo": {
							Description:   "Password or PAT for authenticating the remote repository. Write-only: never stored in state. Requires Terraform >= 1.11.",
							Type:          schema.TypeString,
							Optional:      true,
							WriteOnly:     true,
							Sensitive:     true,
							ConflictsWith: []string{"creds.0.password"},
							RequiredWith:  []string{"creds.0.password_wo_version"},
						},
						"password_wo_version": {
							Description:  "Increment to rotate the credential when using password_wo.",
							Type:         schema.TypeInt,
							Optional:     true,
							RequiredWith: []string{"creds.0.password_wo"},
						},
						"ssh_private_key": {
							Description:   "SSH Key in PEM format for authenticating the repository. Used only for Git repository. Use ssh_private_key_wo for write-only support (Terraform >= 1.11).",
							Type:          schema.TypeString,
							Optional:      true,
							Computed:      true,
							Sensitive:     true,
							ConflictsWith: []string{"creds.0.username", "creds.0.password", "creds.0.tls_client_cert_data", "creds.0.tls_client_cert_key", "creds.0.github_app_private_key", "creds.0.github_app_id", "creds.0.github_app_installation_id", "creds.0.github_app_enterprise_base_url", "creds.0.ssh_private_key_wo"},
						},
						"ssh_private_key_wo": {
							Description:   "SSH Key in PEM format for authenticating the repository. Write-only: never stored in state. Requires Terraform >= 1.11.",
							Type:          schema.TypeString,
							Optional:      true,
							WriteOnly:     true,
							Sensitive:     true,
							ConflictsWith: []string{"creds.0.ssh_private_key"},
							RequiredWith:  []string{"creds.0.ssh_private_key_wo_version"},
						},
						"ssh_private_key_wo_version": {
							Description:  "Increment to rotate the credential when using ssh_private_key_wo.",
							Type:         schema.TypeInt,
							Optional:     true,
							RequiredWith: []string{"creds.0.ssh_private_key_wo"},
						},
						"tls_client_cert_data": {
							Description:   "Certificate in PEM format for authenticating at the repo server. This is used for mTLS. Use tls_client_cert_data_wo for write-only support (Terraform >= 1.11).",
							Type:          schema.TypeString,
							Optional:      true,
							Computed:      true,
							Sensitive:     true,
							ConflictsWith: []string{"creds.0.username", "creds.0.password", "creds.0.github_app_private_key", "creds.0.github_app_id", "creds.0.github_app_installation_id", "creds.0.github_app_enterprise_base_url", "creds.0.tls_client_cert_data_wo"},
						},
						"tls_client_cert_data_wo": {
							Description:   "Certificate in PEM format for authenticating at the repo server (mTLS). Write-only: never stored in state. Requires Terraform >= 1.11.",
							Type:          schema.TypeString,
							Optional:      true,
							WriteOnly:     true,
							Sensitive:     true,
							ConflictsWith: []string{"creds.0.tls_client_cert_data"},
							RequiredWith:  []string{"creds.0.tls_client_cert_data_wo_version"},
						},
						"tls_client_cert_data_wo_version": {
							Description:  "Increment to rotate the credential when using tls_client_cert_data_wo.",
							Type:         schema.TypeInt,
							Optional:     true,
							RequiredWith: []string{"creds.0.tls_client_cert_data_wo"},
						},
						"tls_client_cert_key": {
							Description:   "Private key in PEM format for authenticating at the repo server. This is used for mTLS. Use tls_client_cert_key_wo for write-only support (Terraform >= 1.11).",
							Type:          schema.TypeString,
							Optional:      true,
							Computed:      true,
							Sensitive:     true,
							ConflictsWith: []string{"creds.0.username", "creds.0.password", "creds.0.github_app_private_key", "creds.0.github_app_id", "creds.0.github_app_installation_id", "creds.0.github_app_enterprise_base_url", "creds.0.tls_client_cert_key_wo"},
						},
						"tls_client_cert_key_wo": {
							Description:   "Private key in PEM format for authenticating at the repo server (mTLS). Write-only: never stored in state. Requires Terraform >= 1.11.",
							Type:          schema.TypeString,
							Optional:      true,
							WriteOnly:     true,
							Sensitive:     true,
							ConflictsWith: []string{"creds.0.tls_client_cert_key"},
							RequiredWith:  []string{"creds.0.tls_client_cert_key_wo_version"},
						},
						"tls_client_cert_key_wo_version": {
							Description:  "Increment to rotate the credential when using tls_client_cert_key_wo.",
							Type:         schema.TypeInt,
							Optional:     true,
							RequiredWith: []string{"creds.0.tls_client_cert_key_wo"},
						},
						"github_app_private_key": {
							Description:   "github_app_private_key specifies the private key PEM data for authentication via GitHub app. Use github_app_private_key_wo for write-only support (Terraform >= 1.11).",
							Type:          schema.TypeString,
							Optional:      true,
							Computed:      true,
							Sensitive:     true,
							ConflictsWith: []string{"creds.0.username", "creds.0.password", "creds.0.ssh_private_key", "creds.0.tls_client_cert_data", "creds.0.tls_client_cert_key", "creds.0.github_app_private_key_wo"},
						},
						"github_app_private_key_wo": {
							Description:   "GitHub app private key PEM data. Write-only: never stored in state. Requires Terraform >= 1.11.",
							Type:          schema.TypeString,
							Optional:      true,
							WriteOnly:     true,
							Sensitive:     true,
							ConflictsWith: []string{"creds.0.github_app_private_key"},
							RequiredWith:  []string{"creds.0.github_app_private_key_wo_version"},
						},
						"github_app_private_key_wo_version": {
							Description:  "Increment to rotate the credential when using github_app_private_key_wo.",
							Type:         schema.TypeInt,
							Optional:     true,
							RequiredWith: []string{"creds.0.github_app_private_key_wo"},
						},
						"github_app_id": {
							Description:   "Specifies the Github App ID of the app used to access the repo for GitHub app authentication.",
							Type:          schema.TypeString,
							Sensitive:     true,
							Computed:      true,
							Optional:      true,
							ConflictsWith: []string{"creds.0.username", "creds.0.password", "creds.0.ssh_private_key", "creds.0.tls_client_cert_data", "creds.0.tls_client_cert_key"},
						},
						"github_app_installation_id": {
							Description:   "Specifies the ID of the installed GitHub App for GitHub app authentication.",
							Type:          schema.TypeString,
							Optional:      true,
							Sensitive:     true,
							Computed:      true,
							ConflictsWith: []string{"creds.0.username", "creds.0.password", "creds.0.ssh_private_key", "creds.0.tls_client_cert_data", "creds.0.tls_client_cert_key"},
						},
						"github_app_enterprise_base_url": {
							Description:   "Specifies the GitHub API URL for GitHub app authentication.",
							Type:          schema.TypeString,
							Optional:      true,
							ConflictsWith: []string{"creds.0.username", "creds.0.password", "creds.0.ssh_private_key", "creds.0.tls_client_cert_data", "creds.0.tls_client_cert_key"},
						},
						"enable_oci": {
							Description: "Specifies whether helm-oci support should be enabled for this repo.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"type": {
							Description:  "Type specifies the type of the repoCreds.Can be either 'git' or 'helm. 'git' is assumed if empty or absent",
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice([]string{"git", "helm"}, false),
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
	if woVal, diags := d.GetRawConfigAt(hcty.GetAttrPath("creds").IndexInt(0).GetAttr("password_wo")); !diags.HasError() && woVal.IsKnown() && !woVal.IsNull() {
		resp.RepoCreds.Password = ""
	} else if attr, ok := d.GetOk("creds.0.password"); ok {
		resp.RepoCreds.Password = attr.(string)
	}
	if woVal, diags := d.GetRawConfigAt(hcty.GetAttrPath("creds").IndexInt(0).GetAttr("ssh_private_key_wo")); !diags.HasError() && woVal.IsKnown() && !woVal.IsNull() {
		resp.RepoCreds.SshPrivateKey = ""
	} else if attr, ok := d.GetOk("creds.0.ssh_private_key"); ok {
		resp.RepoCreds.SshPrivateKey = attr.(string)
	}
	if woVal, diags := d.GetRawConfigAt(hcty.GetAttrPath("creds").IndexInt(0).GetAttr("tls_client_cert_data_wo")); !diags.HasError() && woVal.IsKnown() && !woVal.IsNull() {
		resp.RepoCreds.TlsClientCertData = ""
	} else if attr, ok := d.GetOk("creds.0.tls_client_cert_data"); ok {
		resp.RepoCreds.TlsClientCertData = attr.(string)
	}
	if woVal, diags := d.GetRawConfigAt(hcty.GetAttrPath("creds").IndexInt(0).GetAttr("tls_client_cert_key_wo")); !diags.HasError() && woVal.IsKnown() && !woVal.IsNull() {
		resp.RepoCreds.TlsClientCertKey = ""
	} else if attr, ok := d.GetOk("creds.0.tls_client_cert_key"); ok {
		resp.RepoCreds.TlsClientCertKey = attr.(string)
	}
	if woVal, diags := d.GetRawConfigAt(hcty.GetAttrPath("creds").IndexInt(0).GetAttr("github_app_private_key_wo")); !diags.HasError() && woVal.IsKnown() && !woVal.IsNull() {
		resp.RepoCreds.GithubAppPrivateKey = ""
	} else if attr, ok := d.GetOk("creds.0.github_app_private_key"); ok {
		resp.RepoCreds.GithubAppPrivateKey = attr.(string)
	}
	if attr, ok := d.GetOk("creds.0.github_app_installation_id"); ok {
		resp.RepoCreds.GithubAppInstallationID = attr.(string)
	}

	setGitopsRepositoriesCredential(d, &resp)
	preserveRepoCredWoVersions(d)
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

	if woVal, diags := d.GetRawConfigAt(hcty.GetAttrPath("creds").IndexInt(0).GetAttr("password_wo")); !diags.HasError() && woVal.IsKnown() && !woVal.IsNull() {
		resp.RepoCreds.Password = ""
	} else if attr, ok := d.GetOk("creds.0.password"); ok {
		resp.RepoCreds.Password = attr.(string)
	}
	if woVal, diags := d.GetRawConfigAt(hcty.GetAttrPath("creds").IndexInt(0).GetAttr("ssh_private_key_wo")); !diags.HasError() && woVal.IsKnown() && !woVal.IsNull() {
		resp.RepoCreds.SshPrivateKey = ""
	} else if attr, ok := d.GetOk("creds.0.ssh_private_key"); ok {
		resp.RepoCreds.SshPrivateKey = attr.(string)
	}
	if woVal, diags := d.GetRawConfigAt(hcty.GetAttrPath("creds").IndexInt(0).GetAttr("tls_client_cert_data_wo")); !diags.HasError() && woVal.IsKnown() && !woVal.IsNull() {
		resp.RepoCreds.TlsClientCertData = ""
	} else if attr, ok := d.GetOk("creds.0.tls_client_cert_data"); ok {
		resp.RepoCreds.TlsClientCertData = attr.(string)
	}
	if woVal, diags := d.GetRawConfigAt(hcty.GetAttrPath("creds").IndexInt(0).GetAttr("tls_client_cert_key_wo")); !diags.HasError() && woVal.IsKnown() && !woVal.IsNull() {
		resp.RepoCreds.TlsClientCertKey = ""
	} else if attr, ok := d.GetOk("creds.0.tls_client_cert_key"); ok {
		resp.RepoCreds.TlsClientCertKey = attr.(string)
	}
	if woVal, diags := d.GetRawConfigAt(hcty.GetAttrPath("creds").IndexInt(0).GetAttr("github_app_private_key_wo")); !diags.HasError() && woVal.IsKnown() && !woVal.IsNull() {
		resp.RepoCreds.GithubAppPrivateKey = ""
	} else if attr, ok := d.GetOk("creds.0.github_app_private_key"); ok {
		resp.RepoCreds.GithubAppPrivateKey = attr.(string)
	}
	if attr, ok := d.GetOk("creds.0.github_app_installation_id"); ok {
		resp.RepoCreds.GithubAppInstallationID = attr.(string)
	}

	setGitopsRepositoriesCredential(d, &resp)
	preserveRepoCredWoVersions(d)
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

	if _, ok := d.GetOk("creds.0.password_wo_version"); ok {
		resp.RepoCreds.Password = ""
	} else if attr, ok := d.GetOk("creds.0.password"); ok {
		if len(resp.RepoCreds.Password) != 0 {
			resp.RepoCreds.Password = attr.(string)
		}
	}
	if _, ok := d.GetOk("creds.0.ssh_private_key_wo_version"); ok {
		resp.RepoCreds.SshPrivateKey = ""
	} else if attr, ok := d.GetOk("creds.0.ssh_private_key"); ok {
		if len(resp.RepoCreds.SshPrivateKey) != 0 {
			resp.RepoCreds.SshPrivateKey = attr.(string)
		}
	}
	if _, ok := d.GetOk("creds.0.tls_client_cert_data_wo_version"); ok {
		resp.RepoCreds.TlsClientCertData = ""
	} else if attr, ok := d.GetOk("creds.0.tls_client_cert_data"); ok {
		if len(resp.RepoCreds.TlsClientCertData) != 0 {
			resp.RepoCreds.TlsClientCertData = attr.(string)
		}
	}
	if _, ok := d.GetOk("creds.0.tls_client_cert_key_wo_version"); ok {
		resp.RepoCreds.TlsClientCertKey = ""
	} else if attr, ok := d.GetOk("creds.0.tls_client_cert_key"); ok {
		if len(resp.RepoCreds.TlsClientCertKey) != 0 {
			resp.RepoCreds.TlsClientCertKey = attr.(string)
		}
	}
	if _, ok := d.GetOk("creds.0.github_app_private_key_wo_version"); ok {
		resp.RepoCreds.GithubAppPrivateKey = ""
	} else if attr, ok := d.GetOk("creds.0.github_app_private_key"); ok {
		if len(resp.RepoCreds.GithubAppPrivateKey) != 0 {
			resp.RepoCreds.GithubAppPrivateKey = attr.(string)
		}
	}
	if attr, ok := d.GetOk("creds.0.github_app_installation_id"); ok {
		if len(resp.RepoCreds.GithubAppInstallationID) != 0 {
			resp.RepoCreds.GithubAppInstallationID = attr.(string)
		}
	}

	setGitopsRepositoriesCredential(d, &resp)
	preserveRepoCredWoVersions(d)
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

		// Preserve _wo_version integers: d.Set("creds", ...) below would zero
		// them out because the API never returns write-only credential values.
		if v, ok := d.GetOk("creds.0.password_wo_version"); ok {
			cred["password_wo_version"] = v.(int)
		}
		if v, ok := d.GetOk("creds.0.ssh_private_key_wo_version"); ok {
			cred["ssh_private_key_wo_version"] = v.(int)
		}
		if v, ok := d.GetOk("creds.0.tls_client_cert_data_wo_version"); ok {
			cred["tls_client_cert_data_wo_version"] = v.(int)
		}
		if v, ok := d.GetOk("creds.0.tls_client_cert_key_wo_version"); ok {
			cred["tls_client_cert_key_wo_version"] = v.(int)
		}
		if v, ok := d.GetOk("creds.0.github_app_private_key_wo_version"); ok {
			cred["github_app_private_key_wo_version"] = v.(int)
		}

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

			if woVal, diags := d.GetRawConfigAt(hcty.GetAttrPath("creds").IndexInt(0).GetAttr("password_wo")); !diags.HasError() && woVal.IsKnown() && !woVal.IsNull() {
				repoCred.Password = woVal.AsString()
			} else if requestCreds["password"] != nil {
				repoCred.Password = requestCreds["password"].(string)
			}

			if woVal, diags := d.GetRawConfigAt(hcty.GetAttrPath("creds").IndexInt(0).GetAttr("ssh_private_key_wo")); !diags.HasError() && woVal.IsKnown() && !woVal.IsNull() {
				repoCred.SshPrivateKey = woVal.AsString()
			} else if requestCreds["ssh_private_key"] != nil {
				repoCred.SshPrivateKey = requestCreds["ssh_private_key"].(string)
			}

			if woVal, diags := d.GetRawConfigAt(hcty.GetAttrPath("creds").IndexInt(0).GetAttr("tls_client_cert_data_wo")); !diags.HasError() && woVal.IsKnown() && !woVal.IsNull() {
				repoCred.TlsClientCertData = woVal.AsString()
			} else if requestCreds["tls_client_cert_data"] != nil {
				repoCred.TlsClientCertData = requestCreds["tls_client_cert_data"].(string)
			}

			if woVal, diags := d.GetRawConfigAt(hcty.GetAttrPath("creds").IndexInt(0).GetAttr("tls_client_cert_key_wo")); !diags.HasError() && woVal.IsKnown() && !woVal.IsNull() {
				repoCred.TlsClientCertKey = woVal.AsString()
			} else if requestCreds["tls_client_cert_key"] != nil {
				repoCred.TlsClientCertKey = requestCreds["tls_client_cert_key"].(string)
			}

			if woVal, diags := d.GetRawConfigAt(hcty.GetAttrPath("creds").IndexInt(0).GetAttr("github_app_private_key_wo")); !diags.HasError() && woVal.IsKnown() && !woVal.IsNull() {
				repoCred.GithubAppPrivateKey = woVal.AsString()
			} else if requestCreds["github_app_private_key"] != nil {
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

func preserveRepoCredWoVersions(d *schema.ResourceData) {
	if v, ok := d.GetOk("creds.0.password_wo_version"); ok {
		d.Set("creds.0.password_wo_version", v)
	}
	if v, ok := d.GetOk("creds.0.ssh_private_key_wo_version"); ok {
		d.Set("creds.0.ssh_private_key_wo_version", v)
	}
	if v, ok := d.GetOk("creds.0.tls_client_cert_data_wo_version"); ok {
		d.Set("creds.0.tls_client_cert_data_wo_version", v)
	}
	if v, ok := d.GetOk("creds.0.tls_client_cert_key_wo_version"); ok {
		d.Set("creds.0.tls_client_cert_key_wo_version", v)
	}
	if v, ok := d.GetOk("creds.0.github_app_private_key_wo_version"); ok {
		d.Set("creds.0.github_app_private_key_wo_version", v)
	}
}
