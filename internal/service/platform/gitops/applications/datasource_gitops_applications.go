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
		Description: "Datasource for fetching a Harness Gitops Application.",
		ReadContext: datasourceGitopsApplicationRead,
		Schema: map[string]*schema.Schema{
			"account_id": {
				Description: "Account Identifier for the Application.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"org_id": {
				Description: "Organization Identifier for the Application.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"project_id": {
				Description: "Project Identifier for the Application.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"identifier": {
				Description: "Identifier of the Application.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"agent_id": {
				Description: "Agent identifier for the Application.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"cluster_id": {
				Description: "Cluster identifier for the Application.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"repo_id": {
				Description: "Repository identifier for the Application.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"upsert": {
				Description: "Whether to Upsert the application.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"validate": {
				Description: "Whether to validate the application.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"project": {
				Description: "Project is a reference to the project this application belongs to. The empty string means that application belongs to the 'default' project.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"kind": {
				Description: "kind of resource.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"query_refresh": {
				Description: "forces application reconciliation if set to true.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"query_project": {
				Description: "the project names to restrict returned list applications.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"query_resource_version": {
				Description: "when specified with a watch call, shows changes that occur after that particular version of a resource.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"query_selector": {
				Description: "the selector to to restrict returned list to applications only with matched labels.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"query_repo": {
				Description: "the repoURL to restrict returned list applications.",
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
				Description: "definition of Application resource.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"metadata": {
							Description: "metadata that all persisted resources must have, which includes all objects users must create.",
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Description: "Name must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated. ",
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
										Description: "A sequence number representing a specific generation of the desired state. Populated by the system. Read-only. ",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"uid": {
										Description: "UID is the unique in time and space value for this object. It is typically generated by the server on successful creation of a resource and is not allowed to change on PUT operations.",
										Type:        schema.TypeString,
										Computed:    true,
									},
									"labels": {
										Description: "Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services.",
										Type:        schema.TypeMap,
										Optional:    true,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"annotations": {
										Description: "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.",
										Type:        schema.TypeMap,
										Optional:    true,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"owner_references": {
										Description: "List of objects depended by this object. If ALL objects in the list have been deleted, this object will be garbage collected. If this object is managed by a controller, then an entry in this list will point to this controller, with the controller field set to true. There cannot be more than one managing controller. ",
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
													Description: "If true, this reference points to the managing controller. ",
													Type:        schema.TypeBool,
													Optional:    true,
												},
												"block_owner_deletion": {
													Description: "If true, AND if the owner has the \"foregroundDeletion\" finalizer, then the owner cannot be deleted from the key-value store until this reference is removed. Defaults to false. To set this field, a user needs \"delete\" permission of the owner, otherwise 422 (Unprocessable Entity) will be returned.",
													Type:        schema.TypeBool,
													Optional:    true,
												},
											},
										},
									},
									"finalizers": {
										Description: "Must be empty before the object is deleted from the registry. Each entry is an identifier for the responsible component that will remove the entry from the list. If the deletionTimestamp of the object is non-nil, entries in this list can only be removed. Finalizers may be processed and removed in any order. Order is NOT enforced because it introduces significant risk of stuck finalizers. finalizers is a shared field, any actor with permission can reorder it. If the finalizer list is processed in order, then this can lead to a situation in which the component responsible for the first finalizer in the list is waiting for a signal (field value, external system, or other) produced by a component responsible for a finalizer later in the list, resulting in a deadlock. Without enforced ordering finalizers are free to order amongst themselves and are not vulnerable to ordering changes in the list.",
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"cluster_name": {
										Description: "The name of the cluster which the object belongs to. This is used to distinguish resources with same name and namespace in different clusters. This field is not set anywhere right now and apiserver is going to ignore it if set in create or update request.",
										Type:        schema.TypeString,
										Optional:    true,
									},
								},
							},
						},
						"spec": {
							Description: "represents desired application state. Contains link to repository with application definition and additional parameters link definition revision.",
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"source": {
										Description: "contains all information about the source of an application",
										Type:        schema.TypeList,
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"repo_url": {
													Description: "URL to the repository (Git or Helm) that contains the application manifests.",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"path": {
													Description: "directory path within the Git repository, and is only valid for applications sourced from Git.",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"target_revision": {
													Description: "the revision of the source to sync the application to. In case of Git, this can be commit, tag, or branch. If omitted, will equal to HEAD. In case of Helm, this is a semver tag for the Chart's version.",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"chart": {
													Description: "Helm chart name, and must be specified for applications sourced from a Helm repo.",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"helm": {
													Description: "holds helm specific options.",
													Type:        schema.TypeList,
													Optional:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"value_files": {
																Description: "list of Helm value files to use when generating a template",
																Type:        schema.TypeList,
																Optional:    true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"release_name": {
																Description: "Helm release name to use. If omitted it will use the application name.",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"values": {
																Description: "Helm values to be passed to helm template, typically defined as a block.",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"version": {
																Description: "Helm version to use for templating (either \"2\" or \"3\")",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"pass_credentials": {
																Description: "pass credentials to all domains (Helm's --pass-credentials)",
																Type:        schema.TypeBool,
																Optional:    true,
															},
															"parameters": {
																Description: "list of Helm parameters which are passed to the helm template command upon manifest generation.",
																Type:        schema.TypeList,
																Optional:    true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"name": {
																			Description: "the name of the Helm parameter.",
																			Type:        schema.TypeString,
																			Optional:    true,
																		},
																		"value": {
																			Description: "the value for the Helm parameter.",
																			Type:        schema.TypeString,
																			Optional:    true,
																		},
																		"force_string": {
																			Description: "determines whether to tell Helm to interpret booleans and numbers as strings.",
																			Type:        schema.TypeBool,
																			Optional:    true,
																		},
																	},
																},
															},
															"file_parameters": {
																Description: "file parameters to the helm template.",
																Type:        schema.TypeList,
																Optional:    true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"name": {
																			Description: "the name of the Helm parameter.",
																			Type:        schema.TypeString,
																			Optional:    true,
																		},
																		"path": {
																			Description: "the path to the file containing the values for the Helm parameter.",
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
													Description: "options specific to an Application source specific to Kustomize.",
													Type:        schema.TypeList,
													Optional:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"name_prefix": {
																Description: "prefix appended to resources for Kustomize apps.",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"name_suffix": {
																Description: "suffix appended to resources for Kustomize apps.",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"images": {
																Description: "List of Kustomize image override specifications.",
																Type:        schema.TypeList,
																Optional:    true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"common_labels": {
																Description: "list of additional labels to add to rendered manifests.",
																Type:        schema.TypeMap,
																Optional:    true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"version": {
																Description: "version of Kustomize to use for rendering manifests.",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"common_annotations": {
																Description: "list of additional annotations to add to rendered manifests.",
																Type:        schema.TypeMap,
																Optional:    true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"force_common_labels": {
																Description: "whether to force applying common labels to resources for Kustomize apps.",
																Type:        schema.TypeBool,
																Optional:    true,
															},
															"force_common_annotations": {
																Description: "whether to force applying common annotations to resources for Kustomize apps.",
																Type:        schema.TypeBool,
																Optional:    true,
															},
														},
													},
												},
												"ksonnet": {
													Description: "ksonnet specific options.",
													Type:        schema.TypeList,
													Optional:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"environment": {
																Description: "ksonnet application environment name",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"parameters": {
																Description: "list of ksonnet component parameter override values",
																Type:        schema.TypeList,
																Optional:    true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"component": {
																			Description: "Component of the parameter of the Ksonnet App",
																			Type:        schema.TypeString,
																			Optional:    true,
																		},
																		"name": {
																			Description: "name of the parameter of the Ksonnet App",
																			Type:        schema.TypeString,
																			Optional:    true,
																		},
																		"value": {
																			Description: "value of the parameter of the Ksonnet App",
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
													Description: "options for applications of type plain YAML or Jsonnet.",
													Type:        schema.TypeList,
													Optional:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"recurse": {
																Description: "whether to scan a directory recursively for manifests.",
																Type:        schema.TypeBool,
																Optional:    true,
															},
															"exclude": {
																Description: "a glob pattern to match paths against that should be explicitly excluded from being used during manifest generation.",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"include": {
																Description: "a glob pattern to match paths against that should be explicitly included during manifest generation.",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"jsonnet": {
																Description: "options specific to applications of type Jsonnet.",
																Type:        schema.TypeList,
																Optional:    true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"libs": {
																			Description: "Additional library search dirs.",
																			Type:        schema.TypeList,
																			Optional:    true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																		"ext_vars": {
																			Description: "list of Jsonnet External Variables.",
																			Type:        schema.TypeList,
																			Optional:    true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"name": {
																						Description: "name of the external variables of Jsonnet App",
																						Type:        schema.TypeString,
																						Optional:    true,
																					},
																					"value": {
																						Description: "value of the external variables of Jsonnet App",
																						Type:        schema.TypeString,
																						Optional:    true,
																					},
																					"code": {
																						Description: "code of the external variables of Jsonnet App",
																						Type:        schema.TypeBool,
																						Optional:    true,
																					},
																				},
																			},
																		},
																		"tlas": {
																			Description: "list of Jsonnet Top-level Arguments(TLAS).",
																			Type:        schema.TypeList,
																			Optional:    true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"name": {
																						Description: "name of the TLAS of Jsonnet App",
																						Type:        schema.TypeString,
																						Optional:    true,
																					},
																					"value": {
																						Description: "value of the TLAS of Jsonnet App",
																						Type:        schema.TypeString,
																						Optional:    true,
																					},
																					"code": {
																						Description: "code of the TLAS of Jsonnet App",
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
													Description: "options specific to config management plugins.",
													Type:        schema.TypeList,
													Optional:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"name": {
																Description: "Name of the plugin.",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"env": {
																Description: "represents an entry in the application's environment.",
																Type:        schema.TypeList,
																Optional:    true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"name": {
																			Description: "name of the variable, usually expressed in uppercase.",
																			Type:        schema.TypeString,
																			Optional:    true,
																		},
																		"value": {
																			Description: "value of the variable.",
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
										Description: "information about the application's destination.",
										Type:        schema.TypeList,
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Description: "the URL of the target cluster and must be set to the Kubernetes control plane API.",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"namespace": {
													Description: "the target namespace for the application's resources. The namespace will only be set for namespace-scoped resources that have not set a value for .metadata.namespace.",
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
										Description: "controls when a sync will be performed in response to updates in git.",
										Type:        schema.TypeList,
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"sync_options": {
													Description: "Options allow you to specify whole app sync-options.",
													Type:        schema.TypeList,
													Optional:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"automated": {
													Description: "controls the behavior of an automated sync.",
													Type:        schema.TypeList,
													Optional:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"prune": {
																Description: "specifies whether to delete resources from the cluster that are not found in the sources anymore as part of automated sync (default: false)",
																Type:        schema.TypeBool,
																Optional:    true,
															},
															"self_heal": {
																Description: "specifies whether to revert resources back to their desired state upon modification in the cluster (default: false).",
																Type:        schema.TypeBool,
																Optional:    true,
															},
															"allow_empty": {
																Description: "allows apps have zero live resources (default: false).",
																Type:        schema.TypeBool,
																Optional:    true,
															},
														},
													},
												},
												"retry": {
													Description: "contains information about the strategy to apply when a sync failed.",
													Type:        schema.TypeList,
													Optional:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"limit": {
																Description: "Limit is the maximum number of attempts for retrying a failed sync. If set to 0, no retries will be performed.",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"backoff": {
																Description: "the backoff strategy to use on subsequent retries for failing syncs.",
																Type:        schema.TypeList,
																Optional:    true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"duration": {
																			Description: "the amount to back off. Default unit is seconds, but could also be a duration (e.g. \"2m\", \"1h\")",
																			Type:        schema.TypeString,
																			Optional:    true,
																		},
																		"factor": {
																			Description: "a factor to multiply the base duration after each failed retry.",
																			Type:        schema.TypeString,
																			Optional:    true,
																		},
																		"max_duration": {
																			Description: "maximum amount of time allowed for the backoff strategy.",
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
