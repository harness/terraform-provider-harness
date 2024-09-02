package repository

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

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
		Description: "Resource for managing Harness Gitops Repository.",

		CreateContext: resourceGitOpsRepositoryCreate,
		ReadContext:   resourceGitOpsRepositoryRead,
		UpdateContext: resourceGitOpsRepositoryUpdate,
		DeleteContext: resourceGitOpsRepositoryDelete,
		Importer:      helpers.GitopsAgentResourceImporter,
		Schema: map[string]*schema.Schema{
			"account_id": {
				Description: "Account identifier of the GitOps repository.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"project_id": {
				Description: "Project identifier of the GitOps repository.",
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: func() (interface{}, error) {
					return "", nil
				},
			},
			"org_id": {
				Description: "Organization identifier of the GitOps repository.",
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: func() (interface{}, error) {
					return "", nil
				},
			},
			"agent_id": {
				Description: "Agent identifier of the GitOps repository.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"identifier": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Identifier of the GitOps repository.",
			},
			"repo": {
				Description: "Repo details holding application configurations.",
				Type:        schema.TypeList,
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"repo": {
							Description: "URL to the remote repository.",
							Type:        schema.TypeString,
							Required:    true,
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
						"insecure_ignore_host_key": {
							Description: "Indicates if InsecureIgnoreHostKey should be used. Insecure is favored used only for git repos. Deprecated.",
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
							Description: "Certificate in PEM format for authenticating at the repo server. This is used for mTLS. The value should be base64 encoded.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"tls_client_cert_key": {
							Description: "Private key in PEM format for authenticating at the repo server. This is used for mTLS. The value should be base64 encoded.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"type_": {
							Description:  "Type specifies the type of the repo. Can be either \"git\" or \"helm. \"git\" is assumed if empty or absent.",
							Type:         schema.TypeString,
							Optional:     true,
							Computed:     true,
							ValidateFunc: validation.StringInSlice([]string{"git", "helm"}, false),
						},
						"name": {
							Description: "Name to be used for this repo. Only used with Helm repos.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"inherited_creds": {
							Description: "Indicates if the credentials were inherited from a repository credential.",
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
							Description: "The ArgoCD project name corresponding to this GitOps repository. An empty string means that the GitOps repository belongs to the default project created by Harness.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"connection_type": {
							Description: "Identifies the authentication method used to connect to the repository. Possible values: \"HTTPS\" \"SSH\" \"GITHUB\" \"HTTPS_ANONYMOUS_CONNECTION_TYPE\"",
							Type:        schema.TypeString,
							Required:    true,
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
			"gen_type": {
				Description: "Default: \"UNSET\"\nEnum: \"UNSET\" \"AWS_ECR\" \"GOOGLE_GCR\"",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"refresh_interval": {
				Description: "For OCI repos, this is the interval to refresh the token to access the registry.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"ecr_gen": {
				Description: "ECR access token generator specific configuration.",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"region": {
							Description: "AWS region.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"secret_ref": {
							Description: "Secret reference to the AWS credentials.",
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"aws_access_key_id": {
										Description: "AWS access key id.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"aws_secret_access_key": {
										Description: "AWS secret access key.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"aws_session_token": {
										Description: "AWS session token.",
										Type:        schema.TypeString,
										Optional:    true,
									},
								},
							},
						},
						"jwt_auth": {
							Description: "JWT authentication specific configuration.",
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Description: "The name of the ServiceAccount resource being referred to.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"namespace": {
										Description: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"audiences": {
										Description: "Audience specifies the `aud` claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
										Type:        schema.TypeList,
										Optional:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
					},
				},
			},
			"gcr_gen": {
				Description: "GCR access token generator specific configuration.",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"project_id": {
							Description: "GCP project id.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"access_key": {
							Description: "GCP access key.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"workload_identity": {
							Description: "GCP workload identity.",
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"service_account_ref": {
										Description: "Service account reference.",
										Type:        schema.TypeList,
										Optional:    true,
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Description: "The name of the ServiceAccount resource being referred to.",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"namespace": {
													Description: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"audiences": {
													Description: "Audience specifies the `aud` claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
													Type:        schema.TypeList,
													Optional:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
									"cluster_location": {
										Description: "Cluster location.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"cluster_name": {
										Description: "Cluster name.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"cluster_project_id": {
										Description: "Cluster project id.",
										Type:        schema.TypeString,
										Optional:    true,
									},
								},
							},
						},
					},
				},
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

func resourceGitOpsRepositoryRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	var orgIdentifier, projectIdentifier, agentIdentifier, identifier string
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
	resp, httpResp, err := c.RepositoriesApiService.AgentRepositoryServiceGet(ctx, agentIdentifier, identifier, c.AccountId, &nextgen.RepositoriesApiAgentRepositoryServiceGetOpts{
		OrgIdentifier:     optional.NewString(orgIdentifier),
		ProjectIdentifier: optional.NewString(projectIdentifier),
	})

	if err != nil {
		return helpers.HandleReadApiError(err, d, httpResp)
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

func resourceGitOpsRepositoryUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	var orgIdentifier, projectIdentifier, agentIdentifier, identifier string

	var e diag.Diagnostics
	if d.HasChange("identifier") {
		oldValue, newValue := d.GetChange("identifier")
		if oldValue != "" && oldValue != newValue {
			e = append(e, diag.Errorf("%s", "Field 'identifier' cannot be updated after creation.")[0])
		}

		if err := d.Set("identifier", oldValue); err != nil {
			return diag.FromErr(err)
		}
	}

	if d.HasChange("agent_id") {
		oldValue, newValue := d.GetChange("agent_id")
		if oldValue != "" && oldValue != newValue {
			e = append(e, diag.Errorf("%s", "Field 'agent_id' cannot be updated after creation.")[0])
		}

		if err := d.Set("agent_id", oldValue); err != nil {
			return diag.FromErr(err)
		}
	}

	if d.HasChange("account_id") {
		oldValue, newValue := d.GetChange("account_id")
		if oldValue != "" && oldValue != newValue {
			e = append(e, diag.Errorf("%s", "Field 'account_id' cannot be updated after creation.")[0])
		}

		if err := d.Set("account_id", oldValue); err != nil {
			return diag.FromErr(err)
		}

	}

	if d.HasChange("org_id") {
		oldValue, newValue := d.GetChange("org_id")
		if oldValue != "" && oldValue != newValue {
			e = append(e, diag.Errorf("%s", "Field 'org_id' cannot be updated after creation.")[0])
		}

		if err := d.Set("org_id", oldValue); err != nil {
			return diag.FromErr(err)
		}

	}

	if d.HasChange("project_id") {
		oldValue, newValue := d.GetChange("project_id")
		if oldValue != "" && oldValue != newValue {
			e = append(e, diag.Errorf("%s", "Field 'project_id' cannot be updated after creation.")[0])
		}

		if err := d.Set("project_id", oldValue); err != nil {
			return diag.FromErr(err)
		}
	}

	if len(e) > 0 {
		return e
	}

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

	updateRepoRequest := buildUpdateRepoRequest(d)
	resp, httpResp, err := c.RepositoriesApiService.AgentRepositoryServiceUpdateRepository(ctx, updateRepoRequest, agentIdentifier, identifier, &nextgen.RepositoriesApiAgentRepositoryServiceUpdateRepositoryOpts{
		AccountIdentifier: optional.NewString(c.AccountId),
		OrgIdentifier:     optional.NewString(orgIdentifier),
		ProjectIdentifier: optional.NewString(projectIdentifier),
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

func resourceGitOpsRepositoryDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	var orgIdentifier, projectIdentifier, agentIdentifier, identifier string
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
	_, httpResp, err := c.RepositoriesApiService.AgentRepositoryServiceDeleteRepository(ctx, agentIdentifier, identifier, &nextgen.RepositoriesApiAgentRepositoryServiceDeleteRepositoryOpts{
		AccountIdentifier: optional.NewString(c.AccountId),
		OrgIdentifier:     optional.NewString(orgIdentifier),
		ProjectIdentifier: optional.NewString(projectIdentifier),
	})
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}
	// Soft delete lookup error handling
	// https://harness.atlassian.net/browse/PL-23765
	return nil
}

func buildUpdateRepoRequest(d *schema.ResourceData) nextgen.RepositoriesRepoUpdateRequest {
	var updateMask map[string]interface{}
	if attr, ok := d.GetOk("update_mask"); ok {
		if len(attr.([]interface{})) > 0 {
			updateMask = attr.([]interface{})[0].(map[string]interface{})
		}
	}
	var updateMaskPath []string
	if updateMask != nil && updateMask["paths"] != nil && len(updateMask["paths"].([]interface{})) > 0 {
		for _, v := range updateMask["paths"].([]interface{}) {
			updateMaskPath = append(updateMaskPath, v.(string))
		}
	}
	var genType nextgen.RepositoriesEsoGeneratorType
	if attr, ok := d.GetOk("gen_type"); ok {
		genType = nextgen.RepositoriesEsoGeneratorType(attr.(string))
	}
	var refreshInterval string
	if attr, ok := d.GetOk("refresh_interval"); ok {
		refreshInterval = attr.(string)
	}
	var ecrGen *nextgen.RepositoriesEcrAuthorizationTokenGenerator
	var gcrGen *nextgen.RepositoriesGcrAccessTokenGenerator
	if genType == nextgen.GOOGLE_GCR_RepositoriesEsoGeneratorType {
		if attr, ok := d.GetOk("gcr_gen"); ok {
			gcr_gen := attr.([]interface{})
			if gcr_gen != nil && len(gcr_gen) > 0 {
				gcrGen = buildGcrGen(gcr_gen[0].(map[string]interface{}))
			}
		}
	}
	if genType == nextgen.AWS_ECR_RepositoriesEsoGeneratorType {
		if attr, ok := d.GetOk("ecr_gen"); ok {
			ecr_gen := attr.([]interface{})
			if ecr_gen != nil && len(ecr_gen) > 0 {
				ecrGen = buildEcrGen(ecr_gen[0].(map[string]interface{}))
			}
		}
	}

	request := nextgen.RepositoriesRepoUpdateRequest{
		Repo:            buildRepo(d),
		RefreshInterval: refreshInterval,
		UpdateMask: &nextgen.ProtobufFieldMask{
			Paths: updateMaskPath,
		},
	}
	if genType != "" {
		request.GenType = &genType
		request.EcrGen = ecrGen
		request.GcrGen = gcrGen
	}
	return request
}

func buildCreateRepoRequest(d *schema.ResourceData) nextgen.RepositoriesRepoCreateRequest {
	var upsert, credsOnly bool
	if attr, ok := d.GetOk("upsert"); ok {
		upsert = attr.(bool)
	}
	if attr, ok := d.GetOk("creds_only"); ok {
		credsOnly = attr.(bool)
	}
	var genType nextgen.RepositoriesEsoGeneratorType
	if attr, ok := d.GetOk("gen_type"); ok {
		genType = nextgen.RepositoriesEsoGeneratorType(attr.(string))
	}
	var refreshInterval string
	if attr, ok := d.GetOk("refresh_interval"); ok {
		refreshInterval = attr.(string)
	}
	var ecrGen *nextgen.RepositoriesEcrAuthorizationTokenGenerator
	var gcrGen *nextgen.RepositoriesGcrAccessTokenGenerator
	if genType == nextgen.GOOGLE_GCR_RepositoriesEsoGeneratorType {
		if attr, ok := d.GetOk("gcr_gen"); ok {
			gcr_gen := attr.([]interface{})
			if gcr_gen != nil && len(gcr_gen) > 0 {
				gcrGen = buildGcrGen(gcr_gen[0].(map[string]interface{}))
			}
		}
	}
	if genType == nextgen.AWS_ECR_RepositoriesEsoGeneratorType {
		if attr, ok := d.GetOk("ecr_gen"); ok {
			ecr_gen := attr.([]interface{})
			if ecr_gen != nil && len(ecr_gen) > 0 {
				ecrGen = buildEcrGen(ecr_gen[0].(map[string]interface{}))
			}
		}
	}

	request := nextgen.RepositoriesRepoCreateRequest{
		Upsert:    upsert,
		CredsOnly: credsOnly,
		Repo:      buildRepo(d),

		RefreshInterval: refreshInterval,
	}
	if genType != "" {
		request.GenType = &genType
		request.GcrGen = gcrGen
		request.EcrGen = ecrGen
	}
	return request
}

func buildEcrGen(ecrGen map[string]interface{}) *nextgen.RepositoriesEcrAuthorizationTokenGenerator {
	var ecrGenObj nextgen.RepositoriesEcrAuthorizationTokenGenerator
	if ecrGen["region"] != nil {
		ecrGenObj.Region = ecrGen["region"].(string)
	}
	if ecrGen["secret_ref"] != nil {
		attr := ecrGen["secret_ref"].([]interface{})
		if attr != nil && len(attr) > 0 {
			var secretRef nextgen.RepositoriesAwsSecretRef
			secretRefObj := attr[0].(map[string]interface{})
			if secretRefObj["aws_access_key_id"] != nil {
				secretRef.AwsAccessKeyID = secretRefObj["aws_access_key_id"].(string)
			}
			if secretRefObj["aws_secret_access_key"] != nil {
				secretRef.AwsSecretAccessKey = secretRefObj["aws_secret_access_key"].(string)
			}
			if secretRefObj["aws_session_token"] != nil {
				secretRef.AwsSessionToken = secretRefObj["aws_session_token"].(string)
			}

			ecrGenObj.SecretRef = &secretRef
		}
	}
	if ecrGen["jwt_auth"] != nil {
		attr := ecrGen["jwt_auth"].([]interface{})
		if attr != nil && len(attr) > 0 {
			var jwtAuth nextgen.RepositoriesServiceAccountSelector
			jwtAuthObj := attr[0].(map[string]interface{})
			if jwtAuthObj["name"] != nil {
				jwtAuth.Name = jwtAuthObj["name"].(string)
			}
			if jwtAuthObj["namespace"] != nil {
				jwtAuth.Namespace = jwtAuthObj["namespace"].(string)
			}
			ecrGenObj.JwtAuth = &jwtAuth
		}
	}

	return &ecrGenObj

}

func buildGcrGen(gcrGen map[string]interface{}) *nextgen.RepositoriesGcrAccessTokenGenerator {
	var gcrGenObj nextgen.RepositoriesGcrAccessTokenGenerator
	if gcrGen["project_id"] != nil {
		gcrGenObj.ProjectID = gcrGen["project_id"].(string)
	}
	if gcrGen["access_key"] != nil {
		gcrGenObj.AccessKey = gcrGen["access_key"].(string)
	}
	if gcrGen["workload_identity"] != nil {
		attr := gcrGen["workload_identity"].([]interface{})
		if attr != nil && len(attr) > 0 {
			workloadIdentity := attr[0].(map[string]interface{})
			var genWorkloadIdentity nextgen.RepositoriesGcrWorkloadIdentity
			if workloadIdentity["cluster_name"] != nil {
				genWorkloadIdentity.ClusterName = workloadIdentity["cluster_name"].(string)
			}
			if workloadIdentity["cluster_project_id"] != nil {
				genWorkloadIdentity.ClusterProjectID = workloadIdentity["cluster_project_id"].(string)
			}
			if workloadIdentity["cluster_location"] != nil {
				genWorkloadIdentity.ClusterLocation = workloadIdentity["cluster_location"].(string)
			}
			if workloadIdentity["service_account_ref"] != nil {
				attr := workloadIdentity["service_account_ref"].([]interface{})
				if attr != nil && len(attr) > 0 {
					servAccRef := attr[0].(map[string]interface{})
					var genServiceAccountRef nextgen.RepositoriesServiceAccountSelector
					if servAccRef["name"] != nil {
						genServiceAccountRef.Name = servAccRef["name"].(string)
					}
					if servAccRef["namespace"] != nil {
						genServiceAccountRef.Namespace = servAccRef["namespace"].(string)
					}
					//if servAccRef["audience"] != nil {
					//	genServiceAccountRef.Audience = servAccRef["audience"].([]string)
					//}
					genWorkloadIdentity.ServiceAccountRef = &genServiceAccountRef
				}
			}
			gcrGenObj.WorkloadIdentity = &genWorkloadIdentity
		}
	}
	return &gcrGenObj
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
	d.SetId(repo.Identifier)
	d.Set("account_id", repo.AccountIdentifier)
	d.Set("org_id", repo.OrgIdentifier)
	d.Set("project_id", repo.ProjectIdentifier)
	d.Set("agent_id", repo.AgentIdentifier)
	d.Set("identifier", repo.Identifier)
	if repo.Repository != nil {
		repoList := []interface{}{}
		repoO := map[string]interface{}{}
		repoO["repo"] = repo.Repository.Repo
		repoO["username"] = repo.Repository.Username
		repoO["password"] = repo.Repository.Password
		repoO["ssh_private_key"] = repo.Repository.SshPrivateKey
		repoO["insecure_ignore_host_key"] = repo.Repository.InsecureIgnoreHostKey
		repoO["insecure"] = repo.Repository.Insecure
		repoO["enable_lfs"] = repo.Repository.EnableLfs
		repoO["tls_client_cert_data"] = repo.Repository.TlsClientCertData
		repoO["tls_client_cert_key"] = repo.Repository.TlsClientCertKey
		repoO["type_"] = repo.Repository.Type_
		repoO["name"] = repo.Repository.Name
		repoO["inherited_creds"] = repo.Repository.InheritedCreds
		repoO["enable_oci"] = repo.Repository.EnableOCI
		repoO["github_app_private_key"] = repo.Repository.GithubAppPrivateKey
		repoO["github_app_id"] = repo.Repository.GithubAppID
		repoO["github_app_installation_id"] = repo.Repository.GithubAppInstallationID
		repoO["github_app_enterprise_base_url"] = repo.Repository.GithubAppEnterpriseBaseUrl
		repoO["proxy"] = repo.Repository.Proxy
		repoO["project"] = repo.Repository.Project
		repoO["connection_type"] = repo.Repository.ConnectionType

		repoList = append(repoList, repoO)
		d.Set("repo", repoList)
	}
}
