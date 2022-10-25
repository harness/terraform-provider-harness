package applications

import (
	"context"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceGitopsApplications() *schema.Resource {
	resource := &schema.Resource{
		Description: "Datasource for creating a Harness Gitops Application.",
		ReadContext: datasourceGitopsApplicationRead,
		Schema: map[string]*schema.Schema{
			"account_id": {
				Description: "account identifier of the Application.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"org_id": {
				Description: "org identifier of the Application.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"project_id": {
				Description: "org identifier of the Application.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"identifier": {
				Description: "identifier of the Application.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"agent_id": {
				Description: "agent identifier of the Application.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"cluster_id": {
				Description: "cluster identifier of the Application.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"repo_id": {
				Description: "Repository identifier of the Application.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"upsert": {
				Description: "Upsert the application.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"validate": {
				Description: "validate the application.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"project": {
				Description: "Project of the application.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"kind": {
				Description: "kind of resource.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"query_refresh": {
				Description: "Refresh query to get the application.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"query_project": {
				Description: "project query to get the application.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"query_resource_version": {
				Description: "resource version query to get the application.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"query_selector": {
				Description: "query selector to get the application.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"query_repo": {
				Description: "repo query to get the application.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"request_propagation_policy": {
				Description: "Request propagation policy to delete the application.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"request_cascade": {
				Description: "Request cascade to delete the application.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"options_remove_existing_finalizers": {
				Description: "Options to remove existing finalizers to delete the application.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"request_name": {
				Description: "Request name to delete the application.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"name": {
				Description: "Name of the application.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"application": {
				Description: "Application data.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"metadata": {
							Description: "Application data.",
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Description: "name of the application.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"generate_name": {
										Description: "an optional prefix, used by the server, to generate a unique name ONLY IF the Name field has not been provided. If this field is used, the name returned to the client will be different than the name passed. This value will also be combined with a unique suffix. The provided value has the same validation rules as the Name field, and may be truncated by the length of the suffix required to make the value unique on the server.  If this field is specified and the generated name exists, the server will NOT return a 409 - instead, it will either return 201 Created or 500 with Reason ServerTimeout indicating a unique name could not be found in the time allotted, and the client should retry (optionally after the time indicated in the Retry-After header).  Applied only if Name is not specified.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"namespace": {
										Description: "namespace of the application. An empty namespace is equivalent to the \"default\" namespace.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"generation": {
										Description: "generation of the application.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"uid": {
										Description: "UID of the application.",
										Type:        schema.TypeString,
										Computed:    true,
									},
									"labels": {
										Description: "labels to be tagged to the application.",
										Type:        schema.TypeMap,
										Optional:    true,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"annotations": {
										Description: "annotations of the application.",
										Type:        schema.TypeMap,
										Optional:    true,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"owner_references": {
										Description: "owner references of the application",
										Type:        schema.TypeList,
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"api_version": {
													Description: "API version of the referent.",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"kind": {
													Description: "Kind of the referent.",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"name": {
													Description: "Name of the referent.",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"uid": {
													Description: "UID of the referent.",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"controller": {
													Description: "Is the referent controller.",
													Type:        schema.TypeBool,
													Optional:    true,
												},
												"block_owner_deletion": {
													Description: "Block deletion by the owner.",
													Type:        schema.TypeBool,
													Optional:    true,
												},
											},
										},
									},
									"finalizers": {
										Description: "Finalizers of the application.",
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"cluster_name": {
										Description: "Cluster Name of the Application.",
										Type:        schema.TypeString,
										Optional:    true,
									},
								},
							},
						},
						"spec": {
							Description: "spec of the application",
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"source": {
										Description: "spec of the application",
										Type:        schema.TypeList,
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"repo_url": {
													Description: "Repo URL of the application",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"path": {
													Description: "Path in the Repo of the application",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"target_revision": {
													Description: "target revision of the Repo of the application",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"chart": {
													Description: "Helm chart name if its a Helm Repo",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"helm": {
													Description: "Helm config of the application",
													Type:        schema.TypeList,
													Optional:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"value_files": {
																Description: "value files of the helm repo application.",
																Type:        schema.TypeList,
																Optional:    true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"release_name": {
																Description: "Release name of the Helm Repo App",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"values": {
																Description: "Values of the Helm Repo App",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"version": {
																Description: "Version of the Helm Repo App",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"pass_credentials": {
																Description: "pass credentials of the Helm Repo App",
																Type:        schema.TypeBool,
																Optional:    true,
															},
															"parameters": {
																Description: "parameters of the Helm Repo App",
																Type:        schema.TypeList,
																Optional:    true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"name": {
																			Description: "Name of the parameters of the Helm Repo App",
																			Type:        schema.TypeString,
																			Optional:    true,
																		},
																		"value": {
																			Description: "Value of the parameter of the Helm Repo App",
																			Type:        schema.TypeString,
																			Optional:    true,
																		},
																		"force_string": {
																			Description: "force string value of the parameter of the Helm Repo App",
																			Type:        schema.TypeBool,
																			Optional:    true,
																		},
																	},
																},
															},
															"file_parameters": {
																Description: "File parameters of the Helm Repo App",
																Type:        schema.TypeList,
																Optional:    true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"name": {
																			Description: "Name of the file parameter of the Helm Repo App",
																			Type:        schema.TypeString,
																			Optional:    true,
																		},
																		"path": {
																			Description: "Path of the file parameter of the Helm Repo App",
																			Type:        schema.TypeString,
																			Optional:    true,
																		},
																	},
																},
															},
														},
													},
												},
												"kustomize": {
													Description: "kustomize config of the Application.",
													Type:        schema.TypeList,
													Optional:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"name_prefix": {
																Description: "Name prefix of the Kustomize Repo App",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"name_suffix": {
																Description: "Name suffix of the Kustomize Repo App",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"images": {
																Description: "List of images of the Kustomize Repo App",
																Type:        schema.TypeList,
																Optional:    true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"common_labels": {
																Description: "Common Labels of the Kustomize Repo App",
																Type:        schema.TypeMap,
																Optional:    true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"version": {
																Description: "version of the Kustomize Repo App",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"common_annotations": {
																Description: "Common Annotations of the Kustomize Repo App",
																Type:        schema.TypeMap,
																Optional:    true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"force_common_labels": {
																Description: "Force common labels of the Kustomize Repo App",
																Type:        schema.TypeBool,
																Optional:    true,
															},
															"force_common_annotations": {
																Description: "Force common annotations of the Kustomize Repo App",
																Type:        schema.TypeBool,
																Optional:    true,
															},
														},
													},
												},
												"ksonnet": {
													Description: "Ksonnet config of the application.",
													Type:        schema.TypeList,
													Optional:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"environment": {
																Description: "Environment of the Ksonnet Repo App",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"parameters": {
																Description: "Parameters of the Ksonnet Repo App",
																Type:        schema.TypeList,
																Optional:    true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"component": {
																			Description: "Component of the parameter of the Ksonnet Repo App",
																			Type:        schema.TypeString,
																			Optional:    true,
																		},
																		"name": {
																			Description: "name of the parameter of the Ksonnet Repo App",
																			Type:        schema.TypeString,
																			Optional:    true,
																		},
																		"value": {
																			Description: "value of the parameter of the Ksonnet Repo App",
																			Type:        schema.TypeString,
																			Optional:    true,
																		},
																	},
																},
															},
														},
													},
												},
												"directory": {
													Description: "Directory config of the application.",
													Type:        schema.TypeList,
													Optional:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"recurse": {
																Description: "Recurse config of the Directory Repo App",
																Type:        schema.TypeBool,
																Optional:    true,
															},
															"exclude": {
																Description: "exclude config of the Directory Repo App",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"include": {
																Description: "include config of the Directory Repo App",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"jsonnet": {
																Description: "Directory config of the application.",
																Type:        schema.TypeList,
																Optional:    true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"libs": {
																			Description: "libs of the Directory Repo App",
																			Type:        schema.TypeList,
																			Optional:    true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																		"ext_vars": {
																			Description: "external variables of the Directory Repo App",
																			Type:        schema.TypeList,
																			Optional:    true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"name": {
																						Description: "name of the external variables of Jsonnet of the Directory Repo App",
																						Type:        schema.TypeString,
																						Optional:    true,
																					},
																					"value": {
																						Description: "value of the external variables of Jsonnet of the Directory Repo App",
																						Type:        schema.TypeString,
																						Optional:    true,
																					},
																					"code": {
																						Description: "code of the external variables of Jsonnet of the Directory Repo App",
																						Type:        schema.TypeBool,
																						Optional:    true,
																					},
																				},
																			},
																		},
																		"tlas": {
																			Description: "tlas of the Directory Repo App",
																			Type:        schema.TypeList,
																			Optional:    true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"name": {
																						Description: "name of the tlas of Jsonnet of the Directory Repo App",
																						Type:        schema.TypeString,
																						Optional:    true,
																					},
																					"value": {
																						Description: "value of the tlas of Jsonnet of the Directory Repo App",
																						Type:        schema.TypeString,
																						Optional:    true,
																					},
																					"code": {
																						Description: "code of the tlas of Jsonnet of the Directory Repo App",
																						Type:        schema.TypeBool,
																						Optional:    true,
																					},
																				},
																			},
																		},
																	},
																},
															},
														},
													},
												},
												"plugin": {
													Description: "Plugin config of the Application.",
													Type:        schema.TypeList,
													Optional:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"name": {
																Description: "Name of the plugin of the Plugin Repo App",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"env": {
																Description: "List of env details of the Plugin Repo App",
																Type:        schema.TypeList,
																Optional:    true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"name": {
																			Description: "name of the environment config of the plugin of the Plugin Repo App",
																			Type:        schema.TypeString,
																			Optional:    true,
																		},
																		"value": {
																			Description: "value of the environment config of the plugin of the Plugin Repo App",
																			Type:        schema.TypeString,
																			Optional:    true,
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
									"destination": {
										Description: "destination config of the Repo App",
										Type:        schema.TypeList,
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Description: "name of the destination of the Application",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"namespace": {
													Description: "namespace of the destination of the Application",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"server": {
													Description: "server of the destination of the Application",
													Type:        schema.TypeString,
													Optional:    true,
												},
											},
										},
									},
									"sync_policy": {
										Description: "Sync Policy of the Application",
										Type:        schema.TypeList,
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"sync_options": {
													Description: "Syn Options of the sync policy of the Application",
													Type:        schema.TypeList,
													Optional:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"automated": {
													Description: "Automated Sync Policy of the Application",
													Type:        schema.TypeList,
													Optional:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"prune": {
																Description: "Enable Prune of the automated sync policy of the Application",
																Type:        schema.TypeBool,
																Optional:    true,
															},
															"self_heal": {
																Description: "Enable self heal of the automated sync policy of the Application",
																Type:        schema.TypeBool,
																Optional:    true,
															},
															"allow_empty": {
																Description: "Enable allow empty of the automated sync policy of the Application",
																Type:        schema.TypeBool,
																Optional:    true,
															},
														},
													},
												},
												"retry": {
													Description: "Retry Sync Policy of the Application",
													Type:        schema.TypeList,
													Optional:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"limit": {
																Description: "maximum number of attempts for retrying a failed sync. If set to 0, no retries will be performed.",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"backoff": {
																Description: "Automated Sync Policy of the Application",
																Type:        schema.TypeList,
																Optional:    true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"duration": {
																			Description: "Duration after which retry is to be attempted.",
																			Type:        schema.TypeString,
																			Optional:    true,
																		},
																		"factor": {
																			Description: "factor of durations to be retried.",
																			Type:        schema.TypeString,
																			Optional:    true,
																		},
																		"max_duration": {
																			Description: "max duration of the retries.",
																			Type:        schema.TypeString,
																			Optional:    true,
																		},
																	},
																},
															},
														},
													},
												},
											},
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

func datasourceGitopsApplicationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	var agentIdentifier, orgIdentifier, projectIdentifier, repoIdentifier, queryName, queryRefresh, queryProject, queryResourceVersion, querySelector string
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
		queryName = attr.(string)
	}
	if attr, ok := d.GetOk("repo_id"); ok {
		repoIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("query_refresh"); ok {
		queryRefresh = attr.(string)
	}
	if attr, ok := d.GetOk("query_project"); ok {
		queryProject = attr.(string)
	}
	if attr, ok := d.GetOk("query_resource_version"); ok {
		queryResourceVersion = attr.(string)
	}
	if attr, ok := d.GetOk("query_selector"); ok {
		querySelector = attr.(string)
	}
	resp, httpResp, err := c.ApplicationsApiService.AgentApplicationServiceGet(ctx, agentIdentifier, queryName, c.AccountId, orgIdentifier, projectIdentifier, &nextgen.ApplicationsApiAgentApplicationServiceGetOpts{
		QueryRefresh:         optional.NewString(queryRefresh),
		QueryProject:         optional.NewInterface(queryProject),
		QueryResourceVersion: optional.NewString(queryResourceVersion),
		QuerySelector:        optional.NewString(querySelector),
		QueryRepo:            optional.NewString(repoIdentifier),
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
	setApplication(d, &resp)
	return nil
}
