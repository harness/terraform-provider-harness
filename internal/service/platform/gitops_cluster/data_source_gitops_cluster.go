package gitops_cluster

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
		Description: "Data source for retreiving a Harness Gitops Cluster.",
		ReadContext: dataSourceGitopsClusterRead,

		Schema: map[string]*schema.Schema{
			"account_identifier": {
				Description: "account identifier of the cluster.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"project_identifier": {
				Description: "project identifier of the cluster.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"org_identifier": {
				Description: "organization identifier of the cluster.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"agent_identifier": {
				Description: "agent identifier of the cluster.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"identifier": {
				Description: "identifier of the cluster.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"query": {
				Description: "query for cluster resources",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"server": {
							Description: "server of the cluster",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"name": {
							Description: "name of the cluster",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"id": {
							Description: "cluster server URL or cluster name",
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Description: "the type of the specified cluster identifier ( 'server' - default, 'name' )",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"value": {
										Description: "the cluster server URL or cluster name",
										Type:        schema.TypeString,
										Optional:    true,
									},
								}},
						},
					},
				},
			},
			"request": {
				Description: "Cluster create/Update request.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"upsert": {
							Description: "if the cluster should be upserted.",
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
						"update_mask": {
							Description: "Update mask of the cluster.",
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"paths": {
										Description: "The set of field mask paths.",
										Optional:    true,
										Type:        schema.TypeList,
										Elem: &schema.Schema{
											Type: schema.TypeList,
										},
									},
								},
							},
						},
						"id": {
							Description: "cluster server URL or cluster name",
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Description: "the type of the specified cluster identifier ( 'server' - default, 'name' )",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"value": {
										Description: "the cluster server URL or cluster name",
										Type:        schema.TypeString,
										Optional:    true,
									},
								}},
						},
						"cluster": {
							Description: "cluster details.",
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"server": {
										Description: "the API server URL of the Kubernetes cluster.",
										Type:        schema.TypeString,
										Required:    true,
									},
									"name": {
										Description: "Name of the cluster. If omitted, will use the server address",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"config": {
										Description: "Cluster Config",
										Type:        schema.TypeList,
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"username": {
													Description: "username for the server of the cluster.",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"password": {
													Description: "password for the server of the cluster.",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"bearer_token": {
													Description: "Bearer authentication token the cluster.",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"tls_client_config": {
													Description: "contains settings to enable transport layer security",
													Type:        schema.TypeList,
													Optional:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"insecure": {
																Description: "if the TLS connection to the cluster should be insecure.",
																Type:        schema.TypeBool,
																Optional:    true,
															},
															"server_name": {
																Description: "server name for SNI in the client to check server certificates against",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"cert_data": {
																Description: "certficate data holds PEM-encoded bytes (typically read from a client certificate file).",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"key_data": {
																Description: "key data holds PEM-encoded bytes (typically read from a client certificate key file).",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"ca_data": {
																Description: "holds PEM-encoded bytes (typically read from a root certificates bundle).",
																Type:        schema.TypeString,
																Optional:    true,
															},
														},
													},
												},
												"aws_auth_config": {
													Description: "contains IAM authentication configuration",
													Type:        schema.TypeList,
													Optional:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"cluster_name": {
																Description: "contains AWS cluster name.",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"role_a_r_n": {
																Description: "contains optional role ARN. If set then AWS IAM Authenticator.",
																Type:        schema.TypeString,
																Optional:    true,
															},
														},
													},
												},
												"exec_provider_config": {
													Description: "contains configuration for an exec provider",
													Type:        schema.TypeList,
													Optional:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"command": {
																Description: "command to execute.",
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
																Description: "additional environment variables to expose to the process.",
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
																Description: "This text is shown to the user when the executable doesn't seem to be present",
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
										Description: "list of namespaces which are accessible in that cluster. Cluster level resources will be ignored if namespace list is not empty.",
										Type:        schema.TypeList,
										Optional:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"refresh_requested_at": {
										Description: "time when cluster cache refresh has been requested.",
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
									"info": {
										Description: "information about cluster cache and state",
										Type:        schema.TypeList,
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"connection_state": {
													Description: "information about the connection to the cluster",
													Type:        schema.TypeList,
													Optional:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"status": {
																Description: "the current status indicator for the connection",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"message": {
																Description: "human readable information about the connection status",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"attempted_at": {
																Description: "time when cluster cache refresh has been requested.",
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
													Description: "information about the Kubernetes version of the cluster",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"cache_info": {
													Description: "information about the cluster cache",
													Type:        schema.TypeList,
													Optional:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"resources_count": {
																Description: "number of observed Kubernetes resources",
																Type:        schema.TypeInt,
																Optional:    true,
															},
															"apis_count": {
																Description: "number of observed Kubernetes API count",
																Type:        schema.TypeInt,
																Optional:    true,
															},
															"last_cache_sync_time": {
																Description: "time of most recent cache synchronization",
																Type:        schema.TypeString,
																Optional:    true,
															},
														},
													},
												},
												"applications_count": {
													Description: "the number of applications managed by Argo CD on the cluster",
													Type:        schema.TypeInt,
													Optional:    true,
												},
												"api_versions": {
													Description: "list of API versions supported by the cluster",
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
										Description: " optional shard number. Calculated on the fly by the application controller if not specified.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"cluster_resources": {
										Description: "Indicates if cluster level resources should be managed. This setting is used only if cluster is connected in a namespaced mode.",
										Type:        schema.TypeBool,
										Optional:    true,
									},
									"project": {
										Description: "Reference between project and cluster that allow you automatically to be added as item inside Destinations project entity",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"labels": {
										Description: "Labels for cluster secret metadata",
										Type:        schema.TypeMap,
										Optional:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"annotations": {
										Description: "Annotations for cluster secret metadata",
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
	agentIdentifier := d.Get("agent_identifier").(string)
	identifier := d.Get("identifier").(string)
	var queryName, queryServer string
	if d.Get("query") != nil && len(d.Get("query").([]interface{})) > 0 {
		query := d.Get("query").([]interface{})[0].(map[string]interface{})
		queryServer = query["server"].(string)
		queryName = query["name"].(string)
	}
	resp, httpResp, err := c.AgentClusterApi.AgentClusterServiceGet(ctx, agentIdentifier, identifier, &nextgen.AgentClusterServiceApiAgentClusterServiceGetOpts{
		AccountIdentifier: optional.NewString(c.AccountId),
		OrgIdentifier:     optional.NewString(d.Get("org_identifier").(string)),
		ProjectIdentifier: optional.NewString(d.Get("project_identifier").(string)),
		QueryServer:       optional.NewString(queryServer),
		QueryName:         optional.NewString(queryName),
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
