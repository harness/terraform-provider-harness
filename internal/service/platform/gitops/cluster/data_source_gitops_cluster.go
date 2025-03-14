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

func DataSourceGitopsCluster() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for fetching a Harness GitOps Cluster.",
		ReadContext: dataSourceGitopsClusterRead,

		Schema: map[string]*schema.Schema{
			"account_id": {
				Description: "Account identifier of the GitOps cluster.",
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				Deprecated:  "This field is deprecated and will be removed in a future release.",
			},
			"project_id": {
				Description: "Project identifier of the GitOps cluster.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"org_id": {
				Description: "Organization identifier of the cluster.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"agent_id": {
				Description: "Agent identifier of the GitOps cluster.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"identifier": {
				Description: "Identifier of the GitOps cluster.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"request": {
				Description: "Cluster create or update request.",
				Type:        schema.TypeList,
				Computed:    true,
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
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"tags": {
							Description: "Tags for the GitOps cluster. These can be used to search or filter the GitOps agents.",
							Type:        schema.TypeSet,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"cluster": {
							Description: "GitOps cluster details.",
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"server": {
										Description: "API server URL of the kubernetes cluster.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"name": {
										Description: "Name of the cluster. If omitted, the server address will be used.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"config": {
										Description: "GitOps cluster config.",
										Type:        schema.TypeList,
										Optional:    true,
										MaxItems:    1,
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
													Optional:    true,
												},
												"tls_client_config": {
													Description: "Settings to enable transport layer security.",
													Type:        schema.TypeList,
													Optional:    true,
													MaxItems:    1,
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
												"disable_compression": {
													Description: "DisableCompression bypasses automatic GZip compression requests to to the cluster's API server. Corresponds to running kubectl with --disable-compression",
													Type:        schema.TypeBool,
													Optional:    true,
												},
												"proxy_url": {
													Description: "The URL to the proxy to be used for all requests send to the cluster's API server",
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
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"refresh_requested_at": {
										Description: "Time when cluster cache refresh has been requested.",
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
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
																Optional:    true,
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

func dataSourceGitopsClusterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	ctx = context.WithValue(ctx, nextgen.ContextAccessToken, hh.EnvVars.BearerToken.Get())
	agentIdentifier := d.Get("agent_id").(string)
	identifier := d.Get("identifier").(string)
	resp, httpResp, err := c.ClustersApi.AgentClusterServiceGet(ctx, agentIdentifier, identifier, c.AccountId, &nextgen.ClustersApiAgentClusterServiceGetOpts{
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
