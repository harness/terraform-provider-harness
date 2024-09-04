package cluster

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

func ResourceGitopsCluster() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for managing a Harness Gitops Cluster.",

		CreateContext: resourceGitopsClusterCreate,
		ReadContext:   resourceGitopsClusterRead,
		UpdateContext: resourceGitopsClusterUpdate,
		DeleteContext: resourceGitopsClusterDelete,
		Importer:      helpers.GitopsAgentResourceImporter,

		Schema: map[string]*schema.Schema{
			"account_id": {
				Description: "Account identifier of the GitOps cluster.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"project_id": {
				Description: "Project identifier of the GitOps cluster.",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
			},
			"org_id": {
				Description: "Organization identifier of the cluster.",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
			},
			"agent_id": {
				Description: "Agent identifier of the GitOps cluster.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"identifier": {
				Description: "Identifier of the GitOps cluster.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"request": {
				Description: "Cluster create or update request.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"upsert": {
							Description: "Indicates if the GitOps cluster should be updated if existing and inserted if not.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"updated_fields": {
							Description: "Fields which are updated.",
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"tags": {
							Description: "Tags for the GitOps cluster. These can be used to search or filter the GitOps agents.",
							Type:        schema.TypeSet,
							Optional:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"cluster": {
							Description: "GitOps cluster details.",
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"server": {
										Description: "API server URL of the kubernetes cluster.",
										Type:        schema.TypeString,
										Required:    true,
									},
									"name": {
										Description: "Name of the cluster. If omitted, the server address will be used.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"config": {
										Description: "GitOps cluster config.",
										Type:        schema.TypeList,
										Required:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"username": {
													Description: "Username of the server of the cluster.",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"password": {
													Description: "Password of the server of the cluster.",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"bearer_token": {
													Description: "Bearer authentication token the cluster.",
													Type:        schema.TypeString,
													Sensitive:   true,
													Optional:    true,
												},
												"tls_client_config": {
													Description: "Settings to enable transport layer security.",
													Type:        schema.TypeList,
													Optional:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"insecure": {
																Description: "Indicates if the TLS connection to the cluster should be insecure.",
																Type:        schema.TypeBool,
																Optional:    true,
															},
															"server_name": {
																Description: "Server name for SNI in the client to check server certificates against. If ServerName is empty, the hostname used to contact the server is used.",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"cert_data": {
																Description: "Certificate data holds PEM-encoded bytes (typically read from a client certificate file). CertData takes precedence over CertFile. Use this if you are using mTLS. The value should be base64 encoded.",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"key_data": {
																Description: "Key data holds PEM-encoded bytes (typically read from a client certificate key file). KeyData takes precedence over KeyFile. Use this if you are using mTLS. The value should be base64 encoded.",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"ca_data": {
																Description: "CA data holds PEM-encoded bytes (typically read from a root certificates bundle). Use this if you are using self-signed certificates. CAData takes precedence over CAFile. The value should be base64 encoded.",
																Type:        schema.TypeString,
																Optional:    true,
															},
														},
													},
												},
												"role_a_r_n": {
													Description: "Optional role ARN. If set then used for AWS IAM Authenticator.",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"aws_cluster_name": {
													Description: "AWS Cluster name. If set then AWS CLI EKS token command will be used to access cluster.",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"exec_provider_config": {
													Description: "Configuration for an exec provider.",
													Type:        schema.TypeList,
													Optional:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"command": {
																Description: "Command to execute.",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"args": {
																Description: "Arguments to pass to the command when executing it.",
																Type:        schema.TypeList,
																Optional:    true,
																Computed:    true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"env": {
																Description: "Additional environment variables to expose to the process.",
																Type:        schema.TypeMap,
																Optional:    true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"api_version": {
																Description: "Preferred input version of the ExecInfo.",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"install_hint": {
																Description: "Message displayed when the executable is not found.",
																Type:        schema.TypeString,
																Optional:    true,
															},
														},
													},
												},
												"cluster_connection_type": {
													Description: "Identifies the authentication method used to connect to the cluster.",
													Type:        schema.TypeString,
													Optional:    true,
												},
											},
										},
									},
									"namespaces": {
										Description: "List of namespaces which are accessible in that cluster. Cluster level resources will be ignored if namespace list is not empty.",
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"refresh_requested_at": {
										Description: "Time when cluster cache refresh has been requested.",
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"seconds": {
													Description: "Represents seconds of UTC time since Unix epoch 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59Z inclusive.",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"nanos": {
													Description: "Non-negative fractions of a second at nanosecond resolution. Negative second values with fractions must still have non-negative nanos values that count forward in time. Must be from 0 to 999,999,999 inclusive. This field may be limited in precision depending on context.",
													Type:        schema.TypeInt,
													Optional:    true,
												},
											},
										},
									},
									"info": {
										Description: "Information about cluster cache and state.",
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"connection_state": {
													Description: "Information about the connection to the cluster.",
													Type:        schema.TypeList,
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"status": {
																Description: "Current status indicator of the connection.",
																Type:        schema.TypeString,
																Computed:    true,
															},
															"message": {
																Description: "Information about the connection status.",
																Type:        schema.TypeString,
																Computed:    true,
															},
															"attempted_at": {
																Description: "Time when cluster cache refresh has been requested.",
																Type:        schema.TypeList,
																Computed:    true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"seconds": {
																			Description: "Represents seconds of UTC time since Unix epoch 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59Z inclusive.",
																			Type:        schema.TypeString,
																			Optional:    true,
																		},
																		"nanos": {
																			Description: "Non-negative fractions of a second at nanosecond resolution. Negative second values with fractions must still have non-negative nanos values that count forward in time. Must be from 0 to 999,999,999 inclusive. This field may be limited in precision depending on context.",
																			Type:        schema.TypeInt,
																			Optional:    true,
																		},
																	},
																},
															},
														},
													},
												},
												"server_version": {
													Description: "Kubernetes version of the cluster.",
													Type:        schema.TypeString,
													Computed:    true,
												},
												"cache_info": {
													Description: "Information about the cluster cache.",
													Type:        schema.TypeList,
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"resources_count": {
																Description: "Number of observed kubernetes resources.",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"apis_count": {
																Description: "Number of observed kubernetes API count.",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"last_cache_sync_time": {
																Description: "Time of most recent cache synchronization.",
																Type:        schema.TypeString,
																Optional:    true,
															},
														},
													},
												},
												"applications_count": {
													Description: "Number of applications managed by Argo CD on the cluster.",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"api_versions": {
													Description: "List of API versions supported by the cluster.",
													Type:        schema.TypeList,
													Optional:    true,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
									"shard": {
										Description: "Shard number to be managed by a specific application controller pod. Calculated on the fly by the application controller if not specified.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"cluster_resources": {
										Description: "Indicates if cluster level resources should be managed. This setting is used only if cluster is connected in a namespaced mode.",
										Type:        schema.TypeBool,
										Optional:    true,
									},
									"project": {
										Description: "The ArgoCD project name corresponding to this GitOps cluster. An empty string means that the GitOps cluster belongs to the default project created by Harness.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"labels": {
										Description: "Labels for cluster secret metadata.",
										Type:        schema.TypeMap,
										Optional:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"annotations": {
										Description: "Annotations for cluster secret metadata.",
										Type:        schema.TypeMap,
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
		},
	}
	return resource
}

func resourceGitopsClusterCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	createClusterRequest := buildCreateClusterRequest(d)
	resp, httpResp, err := c.ClustersApi.AgentClusterServiceCreate(ctx, *createClusterRequest, agentIdentifier,
		&nextgen.ClustersApiAgentClusterServiceCreateOpts{
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
	if resp.Cluster == nil {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}
	setClusterDetails(d, &resp)
	return nil
}

func resourceGitopsClusterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	ctx = context.WithValue(ctx, nextgen.ContextAccessToken, hh.EnvVars.BearerToken.Get())
	agentIdentifier := d.Get("agent_id").(string)
	identifier := d.Get("identifier").(string)
	resp, httpResp, err := c.ClustersApi.AgentClusterServiceGet(ctx, agentIdentifier, identifier, c.AccountId, &nextgen.ClustersApiAgentClusterServiceGetOpts{
		OrgIdentifier:     optional.NewString(d.Get("org_id").(string)),
		ProjectIdentifier: optional.NewString(d.Get("project_id").(string)),
	})

	if err != nil {
		return helpers.HandleReadApiError(err, d, httpResp)
	}

	// Soft delete lookup error handling
	// https://harness.atlassian.net/browse/PL-23765
	if resp.Cluster == nil {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}
	setClusterDetails(d, &resp)
	return nil
}

func resourceGitopsClusterUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	ctx = context.WithValue(ctx, nextgen.ContextAccessToken, hh.EnvVars.BearerToken.Get())

	agentIdentifier := d.Get("agent_id").(string)
	identifier := d.Get("identifier").(string)
	updateClusterRequest := buildUpdateClusterRequest(d)
	resp, httpResp, err := c.ClustersApi.AgentClusterServiceUpdate(ctx, *updateClusterRequest, agentIdentifier, identifier,
		&nextgen.ClustersApiAgentClusterServiceUpdateOpts{
			AccountIdentifier: optional.NewString(c.AccountId),
			OrgIdentifier:     optional.NewString(d.Get("org_id").(string)),
			ProjectIdentifier: optional.NewString(d.Get("project_id").(string)),
		})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}
	// Soft delete lookup error handling
	// https://harness.atlassian.net/browse/PL-23765
	if resp.Cluster == nil {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}
	setClusterDetails(d, &resp)
	return nil
}

func resourceGitopsClusterDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	ctx = context.WithValue(ctx, nextgen.ContextAccessToken, hh.EnvVars.BearerToken.Get())
	agentIdentifier := d.Get("agent_id").(string)
	identifier := d.Get("identifier").(string)
	_, httpResp, err := c.ClustersApi.AgentClusterServiceDelete(ctx, agentIdentifier, identifier, &nextgen.ClustersApiAgentClusterServiceDeleteOpts{
		AccountIdentifier: optional.NewString(c.AccountId),
		OrgIdentifier:     optional.NewString(d.Get("org_id").(string)),
		ProjectIdentifier: optional.NewString(d.Get("project_id").(string)),
	})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}
	return nil
}

func setClusterDetails(d *schema.ResourceData, cl *nextgen.Servicev1Cluster) {
	d.SetId(cl.Identifier)
	d.Set("account_id", cl.AccountIdentifier)
	d.Set("org_id", cl.OrgIdentifier)
	d.Set("project_id", cl.ProjectIdentifier)
	d.Set("agent_id", cl.AgentIdentifier)
	d.Set("identifier", cl.Identifier)
	// d.Set("created_at", cl.CreatedAt)
	// d.Set("last_modified_at", cl.LastModifiedAt)
	if cl.Cluster != nil {
		requestList := []interface{}{}
		request := map[string]interface{}{}
		clusterList := []interface{}{}
		cluster := map[string]interface{}{}
		cluster["server"] = cl.Cluster.Server
		cluster["name"] = cl.Cluster.Name
		if cl.Cluster.Config != nil {
			configList := []interface{}{}
			config := map[string]interface{}{}
			config["username"] = cl.Cluster.Config.Username
			config["password"] = cl.Cluster.Config.Password
			config["bearer_token"] = cl.Cluster.Config.BearerToken
			if cl.Cluster.Config.TlsClientConfig != nil {
				tlsClientConfigList := []interface{}{}
				tlsClientConfig := map[string]interface{}{}
				tlsClientConfig["insecure"] = cl.Cluster.Config.TlsClientConfig.Insecure
				tlsClientConfig["server_name"] = cl.Cluster.Config.TlsClientConfig.ServerName
				tlsClientConfig["cert_data"] = cl.Cluster.Config.TlsClientConfig.CertData
				tlsClientConfig["key_data"] = cl.Cluster.Config.TlsClientConfig.KeyData
				tlsClientConfig["ca_data"] = cl.Cluster.Config.TlsClientConfig.CaData
				tlsClientConfigList = append(tlsClientConfigList, tlsClientConfig)
				config["tls_client_config"] = tlsClientConfigList
			}
			if cl.Cluster.Config.AwsAuthConfig != nil {
				awsAuthConfigList := []interface{}{}
				awsAuthConfig := map[string]interface{}{}
				awsAuthConfig["cluster_name"] = cl.Cluster.Config.AwsAuthConfig.ClusterName
				awsAuthConfig["role_a_r_n"] = cl.Cluster.Config.AwsAuthConfig.RoleARN
				awsAuthConfigList = append(awsAuthConfigList, awsAuthConfig)
			}
			config["role_a_r_n"] = cl.Cluster.Config.RoleARN
			config["aws_cluster_name"] = cl.Cluster.Config.AwsClusterName
			if cl.Cluster.Config.ExecProviderConfig != nil {
				execProviderConfigList := []interface{}{}
				execProviderConfig := map[string]interface{}{}
				execProviderConfig["command"] = cl.Cluster.Config.ExecProviderConfig.Command
				execProviderConfig["args"] = cl.Cluster.Config.ExecProviderConfig.Args
				execProviderConfig["env"] = cl.Cluster.Config.ExecProviderConfig.Env
				execProviderConfig["api_version"] = cl.Cluster.Config.ExecProviderConfig.ApiVersion
				execProviderConfig["install_hint"] = cl.Cluster.Config.ExecProviderConfig.InstallHint

				execProviderConfigList = append(execProviderConfigList, execProviderConfig)
				config["exec_provider_config"] = execProviderConfigList
			}
			config["cluster_connection_type"] = cl.Cluster.Config.ClusterConnectionType

			configList = append(configList, config)
			cluster["config"] = configList
		}
		cluster["namespaces"] = cl.Cluster.Namespaces
		cluster["refresh_requested_at"] = cl.Cluster.RefreshRequestedAt

		if cl.Cluster.Info != nil {
			clusterInfoList := []interface{}{}
			clusterInfo := map[string]interface{}{}
			if cl.Cluster.Info.ConnectionState != nil {
				connectionStateList := []interface{}{}
				connectionState := map[string]interface{}{}
				connectionState["status"] = cl.Cluster.Info.ConnectionState.Status
				connectionState["message"] = cl.Cluster.Info.ConnectionState.Message
				connectionStateList = append(connectionStateList, connectionState)
				clusterInfo["connection_state"] = connectionStateList
			}
			clusterInfo["server_version"] = cl.Cluster.Info.ServerVersion
			if cl.Cluster.Info.CacheInfo != nil {
				clusterInfoCacheList := []interface{}{}
				clusterInfoCache := map[string]interface{}{}
				clusterInfoCache["resources_count"] = cl.Cluster.Info.CacheInfo.ResourcesCount
				clusterInfoCache["apis_count"] = cl.Cluster.Info.CacheInfo.ApisCount

				clusterInfoCacheList = append(clusterInfoCacheList, clusterInfoCache)
				clusterInfo["cache_info"] = clusterInfoCacheList
			}
			clusterInfo["applications_count"] = cl.Cluster.Info.ApplicationsCount
			clusterInfo["api_versions"] = cl.Cluster.Info.ApiVersions
			clusterInfoList = append(clusterInfoList, clusterInfo)
			cluster["info"] = clusterInfoList
		}
		cluster["project"] = cl.Cluster.Project
		if cl.Cluster.Annotations != nil {
			cluster["annotations"] = cl.Cluster.Annotations
		}
		if cl.Cluster.Labels != nil {
			cluster["labels"] = cl.Cluster.Labels
		}
		clusterList = append(clusterList, cluster)
		request["cluster"] = clusterList
		request["tags"] = helpers.FlattenTags(cl.Tags)
		requestList = append(requestList, request)
		d.Set("request", requestList)
	}
}

func buildCreateClusterRequest(d *schema.ResourceData) *nextgen.ClustersClusterCreateRequest {
	var upsert bool
	var tags map[string]string
	if attr, ok := d.GetOk("request"); ok {
		request := attr.([]interface{})[0].(map[string]interface{})
		upsert = request["upsert"].(bool)
		if tag := request["tags"].(*schema.Set).List(); len(tag) > 0 {
			tags = helpers.ExpandTags(tag)
		}
	}
	return &nextgen.ClustersClusterCreateRequest{
		Upsert:  upsert,
		Tags:    tags,
		Cluster: buildClusterDetails(d),
	}
}

func buildUpdateClusterRequest(d *schema.ResourceData) *nextgen.ClustersClusterUpdateRequest {
	var request map[string]interface{}
	var tags map[string]string
	if attr, ok := d.GetOk("request"); ok {
		request = attr.([]interface{})[0].(map[string]interface{})
	}
	var updatedFields []string
	if request["updated_fields"] != nil && len(request["updated_fields"].([]interface{})) > 0 {
		for _, v := range request["updated_fields"].([]interface{}) {
			updatedFields = append(updatedFields, v.(string))
		}
		if tag := request["tags"].(*schema.Set).List(); len(tag) > 0 {
			tags = helpers.ExpandTags(tag)
		}
	}

	return &nextgen.ClustersClusterUpdateRequest{
		Cluster:       buildClusterDetails(d),
		UpdatedFields: updatedFields,
		Tags:          tags,
	}
}

func buildClusterDetails(d *schema.ResourceData) *nextgen.ClustersCluster {
	var clusterDetails nextgen.ClustersCluster
	var request map[string]interface{}
	if attr, ok := d.GetOk("request"); ok {
		request = attr.([]interface{})[0].(map[string]interface{})
		if request["cluster"] != nil && len(request["cluster"].([]interface{})) > 0 {
			requestCluster := request["cluster"].([]interface{})[0].(map[string]interface{})

			if requestCluster["server"] != nil {
				clusterDetails.Server = requestCluster["server"].(string)
			}
			if requestCluster["name"] != nil {
				clusterDetails.Name = requestCluster["name"].(string)
			}

			if requestCluster["config"] != nil && len(requestCluster["config"].([]interface{})) > 0 {
				clusterConfig := requestCluster["config"].([]interface{})[0].(map[string]interface{})
				clusterDetails.Config = &nextgen.ClustersClusterConfig{}
				if clusterConfig["username"] != nil {
					clusterDetails.Config.Username = clusterConfig["username"].(string)
				}
				if clusterConfig["password"] != nil {
					clusterDetails.Config.Password = clusterConfig["password"].(string)
				}
				if clusterConfig["bearer_token"] != nil {
					clusterDetails.Config.BearerToken = clusterConfig["bearer_token"].(string)
				}

				if clusterConfig["tls_client_config"] != nil && len(clusterConfig["tls_client_config"].([]interface{})) > 0 {
					clusterDetails.Config.TlsClientConfig = &nextgen.ClustersTlsClientConfig{}
					configTlsClientConfig := clusterConfig["tls_client_config"].([]interface{})[0].(map[string]interface{})
					if configTlsClientConfig["insecure"] != nil {
						clusterDetails.Config.TlsClientConfig.Insecure = configTlsClientConfig["insecure"].(bool)
					}
					if configTlsClientConfig["server_name"] != nil {
						clusterDetails.Config.TlsClientConfig.ServerName = configTlsClientConfig["server_name"].(string)
					}
					if configTlsClientConfig["cert_data"] != nil {
						clusterDetails.Config.TlsClientConfig.CertData = configTlsClientConfig["cert_data"].(string)
					}
					if configTlsClientConfig["key_data"] != nil {
						clusterDetails.Config.TlsClientConfig.KeyData = configTlsClientConfig["key_data"].(string)
					}
					if configTlsClientConfig["ca_data"] != nil {
						clusterDetails.Config.TlsClientConfig.CaData = configTlsClientConfig["ca_data"].(string)
					}
				}

				if clusterConfig["role_a_r_n"] != nil {
					clusterDetails.Config.RoleARN = clusterConfig["role_a_r_n"].(string)
				}

				if clusterConfig["aws_cluster_name"] != nil {
					clusterDetails.Config.AwsClusterName = clusterConfig["aws_cluster_name"].(string)
				}

				if clusterConfig["exec_provider_config"] != nil && len(clusterConfig["exec_provider_config"].([]interface{})) > 0 {
					clusterDetails.Config.ExecProviderConfig = &nextgen.ClustersExecProviderConfig{}
					configExecProviderConfig := clusterConfig["exec_provider_config"].([]interface{})[0].(map[string]interface{})
					if configExecProviderConfig["command"] != nil {
						clusterDetails.Config.ExecProviderConfig.Command = configExecProviderConfig["command"].(string)
					}
					if configExecProviderConfig["args"] != nil {
						argsString := make([]string, len(configExecProviderConfig["args"].([]interface{})))
						for _, v := range configExecProviderConfig["args"].([]interface{}) {
							argsString = append(argsString, v.(string))
						}
						clusterDetails.Config.ExecProviderConfig.Args = argsString
					}
					if configExecProviderConfig["env"] != nil {
						var envMap = map[string]string{}
						for k, v := range configExecProviderConfig["env"].(map[string]interface{}) {
							envMap[k] = v.(string)
						}
						clusterDetails.Config.ExecProviderConfig.Env = envMap
					}
					if configExecProviderConfig["api_version"] != nil {
						clusterDetails.Config.ExecProviderConfig.ApiVersion = configExecProviderConfig["api_version"].(string)
					}
					if configExecProviderConfig["install_hint"] != nil {
						clusterDetails.Config.ExecProviderConfig.InstallHint = configExecProviderConfig["install_hint"].(string)
					}
				}

				if clusterConfig["cluster_connection_type"] != nil {
					clusterDetails.Config.ClusterConnectionType = clusterConfig["cluster_connection_type"].(string)
				}
			}

			if requestCluster["namespaces"] != nil {
				namespaces := make([]string, len(requestCluster["namespaces"].([]interface{})))
				for _, v := range requestCluster["namespaces"].([]interface{}) {
					namespaces = append(namespaces, v.(string))
				}
				clusterDetails.Namespaces = namespaces
			}
			if requestCluster["refresh_requested_at"] != nil && len(requestCluster["refresh_requested_at"].([]interface{})) > 0 {
				clusterDetails.RefreshRequestedAt = &nextgen.V1Time{}
				refreshRequestedAt := requestCluster["refresh_requested_at"].([]interface{})[0].(map[string]interface{})
				if refreshRequestedAt["seconds"] != nil {
					clusterDetails.RefreshRequestedAt.Seconds = refreshRequestedAt["seconds"].(string)
				}
				if refreshRequestedAt["nanos"] != nil {
					clusterDetails.RefreshRequestedAt.Nanos = int32(refreshRequestedAt["nanos"].(int))
				}
			}

			if requestCluster["info"] != nil && len(requestCluster["info"].([]interface{})) > 0 {
				clusterDetails.Info = &nextgen.ClustersClusterInfo{}
				clusterInfo := requestCluster["info"].([]interface{})[0].(map[string]interface{})
				if clusterInfo["connection_state"] != nil && len(clusterInfo["connection_state"].([]interface{})) > 0 {
					clusterDetails.Info.ConnectionState = &nextgen.CommonsConnectionState{}
					connectionState := clusterInfo["connection_state"].([]interface{})[0].(map[string]interface{})
					if connectionState["status"] != nil {
						clusterDetails.Info.ConnectionState.Status = connectionState["status"].(string)
					}
					if connectionState["message"] != nil {
						clusterDetails.Info.ConnectionState.Message = connectionState["message"].(string)
					}
					if clusterInfo["attempted_at"] != nil && len(clusterInfo["attempted_at"].([]interface{})) > 0 {
						clusterDetails.Info.ConnectionState.AttemptedAt = &nextgen.V1Time{}
						attemptedAt := clusterInfo["attempted_at"].([]interface{})[0].(map[string]interface{})
						if attemptedAt["seconds"] != nil {
							clusterDetails.Info.ConnectionState.AttemptedAt.Seconds = attemptedAt["seconds"].(string)
						}
						if attemptedAt["nanos"] != nil {
							clusterDetails.Info.ConnectionState.AttemptedAt.Nanos = attemptedAt["nanos"].(int32)
						}
					}
					if clusterInfo["server_version"] != nil {
						clusterDetails.Info.ServerVersion = clusterInfo["server_version"].(string)
					}
				}
			}

			if requestCluster["shard"] != nil {
				clusterDetails.Shard = requestCluster["shard"].(string)
			}
			if requestCluster["cluster_resources"] != nil {
				clusterDetails.ClusterResources = requestCluster["cluster_resources"].(bool)
			}

			if requestCluster["labels"] != nil && len(requestCluster["labels"].(map[string]interface{})) > 0 {
				var labelMap = map[string]string{}
				for k, v := range requestCluster["labels"].(map[string]interface{}) {
					labelMap[k] = v.(string)
				}
				clusterDetails.Labels = labelMap
			}
			if requestCluster["annotations"] != nil && len(requestCluster["annotations"].(map[string]interface{})) > 0 {
				var annotationMap = map[string]string{}
				for k, v := range requestCluster["annotations"].(map[string]interface{}) {
					annotationMap[k] = v.(string)
				}
				clusterDetails.Annotations = annotationMap
			}
		}
	}
	return &clusterDetails
}
