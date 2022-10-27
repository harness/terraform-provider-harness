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

func ResourceGitopsApplication() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for creating a Harness Gitops Application.",

		CreateContext: resourceGitopsApplicationCreate,
		ReadContext:   resourceGitopsApplicationRead,
		UpdateContext: resourceGitopsApplicationUpdate,
		DeleteContext: resourceGitopsApplicationDelete,
		Importer:      helpers.GitopsAgentResourceImporter,

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
				Required:    true,
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
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"metadata": {
							Description: "metadata that all persisted resources must have, which includes all objects users must create.",
							Type:        schema.TypeList,
							Required:    true,
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
													Required:    true,
												},
												"path": {
													Description: "directory path within the Git repository, and is only valid for applications sourced from Git.",
													Type:        schema.TypeString,
													Required:    true,
												},
												"target_revision": {
													Description: "the revision of the source to sync the application to. In case of Git, this can be commit, tag, or branch. If omitted, will equal to HEAD. In case of Helm, this is a semver tag for the Chart's version.",
													Type:        schema.TypeString,
													Required:    true,
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

func resourceGitopsApplicationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	createApplicationRequest := buildCreateApplicationRequest(d)
	var agentIdentifier, orgIdentifier, projectIdentifier, clusterIdentifier, repoIdentifier string
	if attr, ok := d.GetOk("agent_id"); ok {
		agentIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("org_id"); ok {
		orgIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("project_id"); ok {
		projectIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("cluster_id"); ok {
		clusterIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("repo_id"); ok {
		repoIdentifier = attr.(string)
	}

	resp, httpResp, err := c.ApplicationsApiService.AgentApplicationServiceCreate(ctx, createApplicationRequest, agentIdentifier, &nextgen.ApplicationsApiAgentApplicationServiceCreateOpts{
		AccountIdentifier: optional.NewString(c.AccountId),
		OrgIdentifier:     optional.NewString(orgIdentifier),
		ProjectIdentifier: optional.NewString(projectIdentifier),
		ClusterIdentifier: optional.NewString(clusterIdentifier),
		RepoIdentifier:    optional.NewString(repoIdentifier),
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

func resourceGitopsApplicationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

func resourceGitopsApplicationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	updateApplicationRequest := buildUpdateApplicationRequest(d)
	var agentIdentifier, orgIdentifier, projectIdentifier, clusterIdentifier, repoIdentifier, appMetaDataName string
	if attr, ok := d.GetOk("agent_id"); ok {
		agentIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("identifier"); ok {
		appMetaDataName = attr.(string)
	}
	if attr, ok := d.GetOk("org_id"); ok {
		orgIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("project_id"); ok {
		projectIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("cluster_id"); ok {
		clusterIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("repo_id"); ok {
		repoIdentifier = attr.(string)
	}

	resp, httpResp, err := c.ApplicationsApiService.AgentApplicationServiceUpdate(ctx, updateApplicationRequest, c.AccountId, orgIdentifier, projectIdentifier, agentIdentifier, appMetaDataName, &nextgen.ApplicationsApiAgentApplicationServiceUpdateOpts{
		ClusterIdentifier: optional.NewString(clusterIdentifier),
		RepoIdentifier:    optional.NewString(repoIdentifier),
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

func resourceGitopsApplicationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	var agentIdentifier, orgIdentifier, projectIdentifier, requestName, requestPropagationPolicy string
	var requestCascade, optionsRemoveExistingFinalizers bool
	if attr, ok := d.GetOk("agent_id"); ok {
		agentIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("org_id"); ok {
		orgIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("project_id"); ok {
		projectIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("request_propagation_policy"); ok {
		requestPropagationPolicy = attr.(string)
	}
	if attr, ok := d.GetOk("request_cascade"); ok {
		requestCascade = attr.(bool)
	}
	if attr, ok := d.GetOk("options_remove_existing_finalizers"); ok {
		optionsRemoveExistingFinalizers = attr.(bool)
	}
	if attr, ok := d.GetOk("identifier"); ok {
		requestName = attr.(string)
	}

	_, httpResp, err := c.ApplicationsApiService.AgentApplicationServiceDelete(ctx, agentIdentifier, requestName, &nextgen.ApplicationsApiAgentApplicationServiceDeleteOpts{
		AccountIdentifier:               optional.NewString(c.AccountId),
		OrgIdentifier:                   optional.NewString(orgIdentifier),
		ProjectIdentifier:               optional.NewString(projectIdentifier),
		RequestCascade:                  optional.NewBool(requestCascade),
		RequestPropagationPolicy:        optional.NewString(requestPropagationPolicy),
		OptionsRemoveExistingFinalizers: optional.NewBool(optionsRemoveExistingFinalizers),
	})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}
	return nil
}

func setApplication(d *schema.ResourceData, app *nextgen.Servicev1Application) {
	d.SetId(app.Name)
	d.Set("org_id", app.OrgIdentifier)
	d.Set("project_id", app.ProjectIdentifier)
	d.Set("agent_id", app.AgentIdentifier)
	d.Set("account_id", app.AccountIdentifier)
	d.Set("cluster_id", app.ClusterIdentifier)
	d.Set("repo_id", app.RepoIdentifier)
	d.Set("name", app.Name)

	if app.App != nil {
		var applicationList = []interface{}{}
		var application = map[string]interface{}{}

		if app.App.Metadata != nil {
			var metadataList = []interface{}{}
			var metadata = map[string]interface{}{}
			metadata["name"] = app.App.Metadata.Name
			metadata["namespace"] = app.App.Metadata.Namespace
			metadata["uid"] = app.App.Metadata.Uid
			metadata["generation"] = app.App.Metadata.Generation
			if app.App.Metadata.Labels != nil {
				metadata["labels"] = app.App.Metadata.Labels
			}
			if app.App.Metadata.Annotations != nil {
				metadata["annotations"] = app.App.Metadata.Annotations
			}
			metadataList = append(metadataList, metadata)
			application["metadata"] = metadataList
		}
		if app.App.Spec != nil {
			var specList = []interface{}{}
			var spec = map[string]interface{}{}
			if app.App.Spec.Source != nil {
				var sourceList = []interface{}{}
				var source = map[string]interface{}{}
				source["repo_url"] = app.App.Spec.Source.RepoURL
				source["path"] = app.App.Spec.Source.Path
				source["target_revision"] = app.App.Spec.Source.TargetRevision
				source["chart"] = app.App.Spec.Source.Chart
				if app.App.Spec.Source.Helm != nil {
					var helmList = []interface{}{}
					var helm = map[string]interface{}{}
					if app.App.Spec.Source.Helm.ValueFiles != nil && len(app.App.Spec.Source.Helm.ValueFiles) > 0 {
						helm["value_files"] = app.App.Spec.Source.Helm.ValueFiles
					}
					helm["release_name"] = app.App.Spec.Source.Helm.ReleaseName
					helm["values"] = app.App.Spec.Source.Helm.Values
					helm["version"] = app.App.Spec.Source.Helm.Version
					helm["pass_credentials"] = app.App.Spec.Source.Helm.PassCredentials
					if app.App.Spec.Source.Helm.Parameters != nil && len(app.App.Spec.Source.Helm.Parameters) > 0 {
						var helmParametersList = []interface{}{}
						for _, v := range app.App.Spec.Source.Helm.Parameters {
							var helmParam = map[string]interface{}{}
							helmParam["name"] = v.Name
							helmParam["value"] = v.Value
							helmParam["force_string"] = v.ForceString
							helmParametersList = append(helmParametersList, helmParam)
						}
						helm["parameters"] = helmParametersList
					}
					if app.App.Spec.Source.Helm.FileParameters != nil && len(app.App.Spec.Source.Helm.FileParameters) > 0 {
						var helmFileParametersList = []interface{}{}
						for _, v := range app.App.Spec.Source.Helm.FileParameters {
							var helmParam = map[string]interface{}{}
							helmParam["name"] = v.Name
							helmParam["path"] = v.Path
							helmFileParametersList = append(helmFileParametersList, helmParam)
						}
						helm["file_parameters"] = helmFileParametersList
					}

					helmList = append(helmList, helm)
					source["helm"] = helmList
				}
				if app.App.Spec.Source.Kustomize != nil {
					var kustomizeList = []interface{}{}
					var kustomize = map[string]interface{}{}
					kustomize["name_prefix"] = app.App.Spec.Source.Kustomize.NamePrefix
					kustomize["name_suffix"] = app.App.Spec.Source.Kustomize.NameSuffix
					kustomize["images"] = app.App.Spec.Source.Kustomize.Images
					kustomize["common_labels"] = app.App.Spec.Source.Kustomize.CommonLabels
					kustomize["version"] = app.App.Spec.Source.Kustomize.Version
					kustomize["common_annotations"] = app.App.Spec.Source.Kustomize.CommonAnnotations
					kustomize["force_common_labels"] = app.App.Spec.Source.Kustomize.ForceCommonLabels
					kustomize["force_common_annotations"] = app.App.Spec.Source.Kustomize.ForceCommonAnnotations
					kustomizeList = append(kustomizeList, kustomize)
					source["kustomize"] = kustomizeList
				}
				if app.App.Spec.Source.Ksonnet != nil {
					var ksonnetList = []interface{}{}
					var ksonnet = map[string]interface{}{}
					ksonnet["environment"] = app.App.Spec.Source.Ksonnet.Environment
					var ksonnetParamList = []interface{}{}
					for _, v := range app.App.Spec.Source.Ksonnet.Parameters {
						var ksonnetParam = map[string]interface{}{}
						ksonnetParam["component"] = v.Component
						ksonnetParam["name"] = v.Name
						ksonnetParam["value"] = v.Value
						ksonnetParamList = append(ksonnetParamList, ksonnetParam)
					}
					ksonnet["parameters"] = ksonnetParamList
					ksonnetList = append(ksonnetList, ksonnet)
					source["ksonnet"] = ksonnetList
				}
				if app.App.Spec.Source.Directory != nil {
					var directoryList = []interface{}{}
					var directory = map[string]interface{}{}
					directory["recurse"] = app.App.Spec.Source.Directory.Recurse
					directory["exclude"] = app.App.Spec.Source.Directory.Exclude
					directory["include"] = app.App.Spec.Source.Directory.Include
					if app.App.Spec.Source.Directory.Jsonnet != nil {
						var jsonnetList = []interface{}{}
						var jsonnet = map[string]interface{}{}
						jsonnet["libs"] = app.App.Spec.Source.Directory.Jsonnet.Libs
						if app.App.Spec.Source.Directory.Jsonnet.ExtVars != nil {
							var jsonnetExtVarsList = []interface{}{}
							for _, v := range app.App.Spec.Source.Directory.Jsonnet.ExtVars {
								var jsonnetExtVars = map[string]interface{}{}
								jsonnetExtVars["name"] = v.Name
								jsonnetExtVars["value"] = v.Value
								jsonnetExtVars["code"] = v.Code
								jsonnetExtVarsList = append(jsonnetExtVarsList, jsonnetExtVars)
							}
							jsonnet["ext_vars"] = jsonnetExtVarsList
						}
						if app.App.Spec.Source.Directory.Jsonnet.Tlas != nil {
							var jsonnetTlasList = []interface{}{}
							for _, v := range app.App.Spec.Source.Directory.Jsonnet.Tlas {
								var jsonnetTlas = map[string]interface{}{}
								jsonnetTlas["name"] = v.Name
								jsonnetTlas["value"] = v.Value
								jsonnetTlas["code"] = v.Code
								jsonnetTlasList = append(jsonnetTlasList, jsonnetTlas)
							}
							jsonnet["tlas"] = jsonnetTlasList
						}
						jsonnetList = append(jsonnetList, jsonnet)
						directory["jsonnet"] = jsonnetList
					}
					directoryList = append(directoryList, directory)
					source["directory"] = directoryList
				}
				if app.App.Spec.Source.Plugin != nil {
					var pluginList = []interface{}{}
					var plugin = map[string]interface{}{}
					plugin["name"] = app.App.Spec.Source.Plugin.Name
					var pluginEnvList = []interface{}{}
					for _, v := range app.App.Spec.Source.Plugin.Env {
						var pluginEnv = map[string]interface{}{}
						pluginEnv["name"] = v.Name
						pluginEnv["value"] = v.Value
						pluginList = append(pluginEnvList, pluginEnv)
					}
					plugin["env"] = pluginEnvList
					pluginList = append(pluginList, plugin)
					source["plugin"] = pluginList
				}
				sourceList = append(sourceList, source)
				spec["source"] = sourceList
			}
			//destination
			if app.App.Spec.Destination != nil {
				var destinationList = []interface{}{}
				var destination = map[string]interface{}{}
				destination["name"] = app.App.Spec.Destination.Name
				destination["namespace"] = app.App.Spec.Destination.Namespace
				destination["server"] = app.App.Spec.Destination.Server
				destinationList = append(destinationList, destination)
				spec["destination"] = destinationList
			}
			//sync policy
			if app.App.Spec.SyncPolicy != nil {
				var syncPolicyList = []interface{}{}
				var syncPolicy = map[string]interface{}{}
				syncPolicy["sync_options"] = app.App.Spec.SyncPolicy.SyncOptions
				if app.App.Spec.SyncPolicy.Automated != nil {
					var syncPolicyAutomatedList = []interface{}{}
					var syncPolicyAutomated = map[string]interface{}{}
					syncPolicyAutomated["prune"] = app.App.Spec.SyncPolicy.Automated.Prune
					syncPolicyAutomated["self_heal"] = app.App.Spec.SyncPolicy.Automated.SelfHeal
					syncPolicyAutomated["allow_empty"] = app.App.Spec.SyncPolicy.Automated.AllowEmpty
					syncPolicyAutomatedList = append(syncPolicyAutomatedList, syncPolicyAutomated)
					syncPolicy["automated"] = syncPolicyAutomatedList
				}
				if app.App.Spec.SyncPolicy.Retry != nil {
					var syncPolicyRetryList = []interface{}{}
					var syncPolicyRetry = map[string]interface{}{}
					syncPolicyRetry["limit"] = app.App.Spec.SyncPolicy.Retry.Limit
					if app.App.Spec.SyncPolicy.Retry.Backoff != nil {
						var syncPolicyRetryBackoffList = []interface{}{}
						var syncPolicyRetryBackoff = map[string]interface{}{}
						syncPolicyRetryBackoff["duration"] = app.App.Spec.SyncPolicy.Retry.Backoff.Duration
						syncPolicyRetryBackoff["factor"] = app.App.Spec.SyncPolicy.Retry.Backoff.Factor
						syncPolicyRetryBackoff["max_duration"] = app.App.Spec.SyncPolicy.Retry.Backoff.MaxDuration
						syncPolicyRetryBackoffList = append(syncPolicyRetryBackoffList, syncPolicyRetryBackoff)
						syncPolicyRetry["backoff"] = syncPolicyRetryBackoffList
					}
					syncPolicyRetryList = append(syncPolicyRetryList, syncPolicyRetry)
					syncPolicy["retry"] = syncPolicyRetryList
				}

				syncPolicyList = append(syncPolicyList, syncPolicy)
				spec["sync_policy"] = syncPolicyList
			}

			specList = append(specList, spec)
			application["spec"] = specList
		}
		applicationList = append(applicationList, application)
		d.Set("application", applicationList)
	}

}

func buildCreateApplicationRequest(d *schema.ResourceData) nextgen.ApplicationsApplicationCreateRequest {
	var upsert, validate bool
	if attr, ok := d.GetOk("upsert"); ok {
		upsert = attr.(bool)
	}
	if attr, ok := d.GetOk("validate"); ok {
		validate = attr.(bool)
	}
	return nextgen.ApplicationsApplicationCreateRequest{
		Application: buildApplicationRequest(d),
		Upsert:      upsert,
		Validate:    validate,
	}
}

func buildUpdateApplicationRequest(d *schema.ResourceData) nextgen.ApplicationsApplicationUpdateRequest {
	var validate bool
	if attr, ok := d.GetOk("validate"); ok {
		validate = attr.(bool)
	}
	return nextgen.ApplicationsApplicationUpdateRequest{
		Application: buildApplicationRequest(d),
		Validate:    validate,
	}
}

func buildApplicationRequest(d *schema.ResourceData) *nextgen.ApplicationsApplication {
	var metaData nextgen.V1ObjectMeta
	var spec nextgen.ApplicationsApplicationSpec
	var application map[string]interface{}
	if attr, ok := d.GetOk("application"); ok {
		if attr != nil && len(attr.([]interface{})) > 0 {
			application = attr.([]interface{})[0].(map[string]interface{})
			if application["metadata"] != nil && len(application["metadata"].([]interface{})) > 0 {
				var meta = application["metadata"].([]interface{})[0].(map[string]interface{})
				if meta["name"] != nil && len(meta["name"].(string)) > 0 {
					metaData.Name = meta["name"].(string)
				}
				if meta["namespace"] != nil && len(meta["namespace"].(string)) > 0 {
					metaData.Namespace = meta["namespace"].(string)
				}
				if meta["generation"] != nil && len(meta["generation"].(string)) > 0 {
					metaData.Namespace = meta["generation"].(string)
				}
				if meta["cluster_name"] != nil && len(meta["cluster_name"].(string)) > 0 {
					metaData.ClusterName = meta["cluster_name"].(string)
				}
				if meta["finalizers"] != nil && len(meta["finalizers"].([]interface{})) > 0 {
					var finalizers []string
					for _, v := range meta["finalizers"].([]interface{}) {
						finalizers = append(finalizers, v.(string))
					}
					metaData.Finalizers = finalizers
				}
				if meta["labels"] != nil && len(meta["labels"].(map[string]interface{})) > 0 {
					var labelMap = map[string]string{}
					for k, v := range meta["labels"].(map[string]interface{}) {
						labelMap[k] = v.(string)
					}
					metaData.Labels = labelMap
				}
				if meta["annotations"] != nil && len(meta["annotations"].(map[string]interface{})) > 0 {
					var annotationMap = map[string]string{}
					for k, v := range meta["annotations"].(map[string]interface{}) {
						annotationMap[k] = v.(string)
					}
					metaData.Annotations = annotationMap
				}
				if meta["owner_references"] != nil && len(meta["owner_references"].([]interface{})) > 0 {
					var ownerReferences []nextgen.V1OwnerReference
					for _, v := range meta["owner_references"].([]interface{}) {
						if v != nil {
							var vMap = v.(map[string]interface{})
							var ownerRef nextgen.V1OwnerReference
							if vMap["api_version"] != nil && len(vMap["api_version"].(string)) > 0 {
								ownerRef.ApiVersion = vMap["api_version"].(string)
							}
							if vMap["kind"] != nil && len(vMap["kind"].(string)) > 0 {
								ownerRef.Kind = vMap["kind"].(string)
							}
							if vMap["name"] != nil && len(vMap["name"].(string)) > 0 {
								ownerRef.Name = vMap["name"].(string)
							}
							if vMap["uid"] != nil && len(vMap["uid"].(string)) > 0 {
								ownerRef.Uid = vMap["uid"].(string)
							}
							if vMap["controller"] != nil {
								ownerRef.Controller = vMap["controller"].(bool)
							}
							if vMap["block_owner_deletion"] != nil {
								ownerRef.BlockOwnerDeletion = vMap["block_owner_deletion"].(bool)
							}
							ownerReferences = append(ownerReferences, ownerRef)
						}
					}
					metaData.OwnerReferences = ownerReferences
				}
			}

			if application["spec"] != nil && len(application["spec"].([]interface{})) > 0 {
				var specData map[string]interface{}
				specData = application["spec"].([]interface{})[0].(map[string]interface{})
				//Spec Source
				if specData["source"] != nil && len(specData["source"].([]interface{})) > 0 {
					var specSource nextgen.ApplicationsApplicationSource
					var source = specData["source"].([]interface{})[0].(map[string]interface{})
					if source["repo_url"] != nil && len(source["repo_url"].(string)) > 0 {
						specSource.RepoURL = source["repo_url"].(string)
					}
					if source["path"] != nil && len(source["path"].(string)) > 0 {
						specSource.Path = source["path"].(string)
					}
					if source["target_revision"] != nil && len(source["target_revision"].(string)) > 0 {
						specSource.TargetRevision = source["target_revision"].(string)
					}
					if source["chart"] != nil && len(source["chart"].(string)) > 0 {
						specSource.Chart = source["chart"].(string)
					}
					//Helm Source Details
					if source["helm"] != nil && len(source["helm"].([]interface{})) > 0 {
						var helm = source["helm"].([]interface{})[0].(map[string]interface{})
						var helmData nextgen.ApplicationsApplicationSourceHelm
						if helm["value_files"] != nil && len(helm["value_files"].([]interface{})) > 0 {
							var valueFiles []string
							for _, v := range helm["value_files"].([]interface{}) {
								valueFiles = append(valueFiles, v.(string))
							}
							helmData.ValueFiles = valueFiles
						}
						if helm["release_name"] != nil && len(helm["release_name"].(string)) > 0 {
							helmData.ReleaseName = helm["release_name"].(string)
						}
						if helm["values"] != nil && len(helm["values"].(string)) > 0 {
							helmData.Values = helm["values"].(string)
						}
						if helm["version"] != nil && len(helm["version"].(string)) > 0 {
							helmData.Version = helm["version"].(string)
						}
						if helm["pass_credentials"] != nil {
							helmData.PassCredentials = helm["pass_credentials"].(bool)
						}
						if helm["parameters"] != nil && len(helm["parameters"].([]interface{})) > 0 {
							var helmParams []nextgen.ApplicationsHelmParameter
							for _, v := range helm["parameters"].([]interface{}) {
								if v != nil {
									var helmParam = v.(map[string]interface{})
									var helmParamD nextgen.ApplicationsHelmParameter
									if helmParam["name"] != nil && len(helmParam["name"].(string)) > 0 {
										helmParamD.Name = helmParam["name"].(string)
									}
									if helmParam["value"] != nil && len(helmParam["value"].(string)) > 0 {
										helmParamD.Value = helmParam["value"].(string)
									}
									if helmParam["force_string"] != nil {
										helmParamD.ForceString = helmParam["force_string"].(bool)
									}
									helmParams = append(helmParams, helmParamD)
								}
							}
							helmData.Parameters = helmParams
						}
						if helm["file_parameters"] != nil && len(helm["file_parameters"].([]interface{})) > 0 {
							var helmFileParams []nextgen.ApplicationsHelmFileParameter
							for _, v := range helm["file_parameters"].([]interface{}) {
								if v != nil {
									var helmFileParam = v.(map[string]interface{})
									var helmFileParamD nextgen.ApplicationsHelmFileParameter
									if helmFileParam["name"] != nil && len(helmFileParam["name"].(string)) > 0 {
										helmFileParamD.Name = helmFileParam["name"].(string)
									}
									if helmFileParam["path"] != nil && len(helmFileParam["path"].(string)) > 0 {
										helmFileParamD.Path = helmFileParam["path"].(string)
									}
									helmFileParams = append(helmFileParams, helmFileParamD)
								}
								helmData.FileParameters = helmFileParams
							}
						}
						specSource.Helm = &helmData
					}

					//Kustomize Source details
					if source["kustomize"] != nil && len(source["kustomize"].([]interface{})) > 0 {
						var kustomizeSource = source["kustomize"].([]interface{})[0].(map[string]interface{})
						var kustomizeData nextgen.ApplicationsApplicationSourceKustomize
						if kustomizeSource["name_prefix"] != nil && len(kustomizeSource["name_prefix"].(string)) > 0 {
							kustomizeData.NamePrefix = kustomizeSource["name_prefix"].(string)
						}
						if kustomizeSource["name_suffix"] != nil && len(kustomizeSource["name_suffix"].(string)) > 0 {
							kustomizeData.NameSuffix = kustomizeSource["name_suffix"].(string)
						}
						if kustomizeSource["images"] != nil && len(kustomizeSource["images"].([]interface{})) > 0 {
							var kustomizeImages []string
							for _, v := range kustomizeSource["images"].([]interface{}) {
								kustomizeImages = append(kustomizeImages, v.(string))
							}
							kustomizeData.Images = kustomizeImages
						}
						if kustomizeSource["common_labels"] != nil && len(kustomizeSource["common_labels"].([]interface{})) > 0 {
							var kustomizeCommonLabels = map[string]string{}
							for k, v := range kustomizeSource["common_labels"].(map[string]interface{}) {
								kustomizeCommonLabels[k] = v.(string)
							}
							kustomizeData.CommonLabels = kustomizeCommonLabels
						}
						if kustomizeSource["version"] != nil && len(kustomizeSource["version"].(string)) > 0 {
							kustomizeData.Version = kustomizeSource["version"].(string)
						}
						if kustomizeSource["common_annotations"] != nil && len(kustomizeSource["common_annotations"].([]interface{})) > 0 {
							var kustomizeCommonAnnotations = map[string]string{}
							for k, v := range kustomizeSource["common_annotations"].(map[string]interface{}) {
								kustomizeCommonAnnotations[k] = v.(string)
							}
							kustomizeData.CommonAnnotations = kustomizeCommonAnnotations
						}
						if kustomizeSource["force_common_labels"] != nil {
							kustomizeData.ForceCommonLabels = kustomizeSource["force_common_labels"].(bool)
						}
						if kustomizeSource["force_common_annotations"] != nil {
							kustomizeData.ForceCommonAnnotations = kustomizeSource["force_common_annotations"].(bool)
						}

						specSource.Kustomize = &kustomizeData
					}

					//Ksonnet
					if source["ksonnet"] != nil && len(source["ksonnet"].([]interface{})) > 0 {
						var ksonnetSource = source["ksonnet"].([]interface{})[0].(map[string]interface{})
						var ksonnetData nextgen.ApplicationsApplicationSourceKsonnet
						if ksonnetSource["environment"] != nil && len(ksonnetSource["environment"].(string)) > 0 {
							ksonnetData.Environment = ksonnetSource["environment"].(string)
						}
						if ksonnetSource["parameters"] != nil && len(ksonnetSource["parameters"].([]interface{})) > 0 {
							var ksonnetParams []nextgen.ApplicationsKsonnetParameter
							for _, v := range ksonnetSource["parameters"].([]interface{}) {
								if v != nil {
									var ksonnetParamSource = v.(map[string]interface{})
									var ksonnetParam nextgen.ApplicationsKsonnetParameter
									if ksonnetParamSource["component"] != nil && len(ksonnetParamSource["component"].(string)) > 0 {
										ksonnetParam.Component = ksonnetParamSource["component"].(string)
									}
									if ksonnetParamSource["name"] != nil && len(ksonnetParamSource["name"].(string)) > 0 {
										ksonnetParam.Name = ksonnetParamSource["name"].(string)
									}
									if ksonnetParamSource["value"] != nil && len(ksonnetParamSource["value"].(string)) > 0 {
										ksonnetParam.Value = ksonnetParamSource["value"].(string)
									}
									ksonnetParams = append(ksonnetParams, ksonnetParam)
								}
								ksonnetData.Parameters = ksonnetParams
							}
						}
						specSource.Ksonnet = &ksonnetData
					}
					//Directory
					if source["directory"] != nil && len(source["directory"].([]interface{})) > 0 {
						var directorySource = source["directory"].([]interface{})[0].(map[string]interface{})
						var directoryData nextgen.ApplicationsApplicationSourceDirectory
						if directorySource["recurse"] != nil {
							directoryData.Recurse = directorySource["recurse"].(bool)
						}
						if directorySource["exclude"] != nil && len(directorySource["exclude"].(string)) > 0 {
							directoryData.Exclude = directorySource["exclude"].(string)
						}
						if directorySource["include"] != nil && len(directorySource["include"].(string)) > 0 {
							directoryData.Exclude = directorySource["include"].(string)
						}

						if directorySource["jsonnet"] != nil && len(directorySource["jsonnet"].([]interface{})) > 0 {
							var directoryJsonnet = directorySource["jsonnet"].([]interface{})[0].(map[string]interface{})
							var jsonnetData nextgen.ApplicationsApplicationSourceJsonnet
							if directoryJsonnet["libs"] != nil && len(directoryJsonnet["libs"].([]interface{})) > 0 {
								var jsonnetLibs []string
								for _, v := range directoryJsonnet["libs"].([]interface{}) {
									jsonnetLibs = append(jsonnetLibs, v.(string))
								}
								jsonnetData.Libs = jsonnetLibs
							}
							if directoryJsonnet["ext_vars"] != nil && len(directoryJsonnet["ext_vars"].([]interface{})) > 0 {
								var jsonnetExtVars []nextgen.ApplicationsJsonnetVar
								for _, v := range directoryJsonnet["ext_vars"].([]interface{}) {
									if v != nil {
										var jsonnetExtVar = v.(map[string]interface{})
										var jsonnetExtVarData nextgen.ApplicationsJsonnetVar
										if jsonnetExtVar["name"] != nil && len(jsonnetExtVar["name"].(string)) > 0 {
											jsonnetExtVarData.Name = jsonnetExtVar["name"].(string)
										}
										if jsonnetExtVar["value"] != nil && len(jsonnetExtVar["value"].(string)) > 0 {
											jsonnetExtVarData.Value = jsonnetExtVar["value"].(string)
										}
										if jsonnetExtVar["code"] != nil {
											jsonnetExtVarData.Code = jsonnetExtVar["code"].(bool)
										}
										jsonnetExtVars = append(jsonnetExtVars, jsonnetExtVarData)
									}
								}
								jsonnetData.ExtVars = jsonnetExtVars
							}
							if directoryJsonnet["tlas"] != nil && len(directoryJsonnet["tlas"].([]interface{})) > 0 {
								var jsonnetTlasVars []nextgen.ApplicationsJsonnetVar
								for _, v := range directoryJsonnet["ext_vars"].([]interface{}) {
									if v != nil {
										var jsonnetTlasVar = v.(map[string]interface{})
										var jsonnetTlasVarData nextgen.ApplicationsJsonnetVar
										if jsonnetTlasVar["name"] != nil && len(jsonnetTlasVar["name"].(string)) > 0 {
											jsonnetTlasVarData.Name = jsonnetTlasVar["name"].(string)
										}
										if jsonnetTlasVar["value"] != nil && len(jsonnetTlasVar["value"].(string)) > 0 {
											jsonnetTlasVarData.Value = jsonnetTlasVar["value"].(string)
										}
										if jsonnetTlasVar["code"] != nil {
											jsonnetTlasVarData.Code = jsonnetTlasVar["code"].(bool)
										}
										jsonnetTlasVars = append(jsonnetTlasVars, jsonnetTlasVarData)
									}
								}
								jsonnetData.Tlas = jsonnetTlasVars
							}

							directoryData.Jsonnet = &jsonnetData

						}
						specSource.Directory = &directoryData
					}

					//Plugin
					if source["plugin"] != nil && len(source["plugin"].([]interface{})) > 0 {
						var pluginSource = source["plugin"].([]interface{})[0].(map[string]interface{})
						var pluginData nextgen.ApplicationsApplicationSourcePlugin
						if pluginSource["name"] != nil && len(pluginSource["name"].(string)) > 0 {
							pluginData.Name = pluginSource["name"].(string)
						}
						if pluginSource["env"] != nil && len(pluginSource["env"].([]interface{})) > 0 {
							var pluginEnvs []nextgen.ApplicationsEnvEntry
							for _, v := range pluginSource["env"].([]interface{}) {
								if v != nil {
									var pluginEnv = v.(map[string]interface{})
									var pluginEnvData nextgen.ApplicationsEnvEntry
									if pluginEnv["name"] != nil && len(pluginEnv["name"].(string)) > 0 {
										pluginEnvData.Name = pluginEnv["name"].(string)
									}
									if pluginEnv["value"] != nil && len(pluginEnv["value"].(string)) > 0 {
										pluginEnvData.Value = pluginEnv["value"].(string)
									}
									pluginEnvs = append(pluginEnvs, pluginEnvData)
								}
								pluginData.Env = pluginEnvs
							}
						}
						specSource.Plugin = &pluginData
					}
					spec.Source = &specSource
				}

				//Destination
				if specData["destination"] != nil && len(specData["destination"].([]interface{})) > 0 {
					var specDestinationData nextgen.ApplicationsApplicationDestination
					var specDestination = specData["destination"].([]interface{})[0].(map[string]interface{})
					if specDestination["name"] != nil && len(specDestination["name"].(string)) > 0 {
						specDestinationData.Name = specDestination["name"].(string)
					}
					if specDestination["namespace"] != nil && len(specDestination["namespace"].(string)) > 0 {
						specDestinationData.Namespace = specDestination["namespace"].(string)
					}
					if specDestination["server"] != nil && len(specDestination["server"].(string)) > 0 {
						specDestinationData.Server = specDestination["server"].(string)
					}
					spec.Destination = &specDestinationData
				}
				//sync policy
				if specData["sync_policy"] != nil && len(specData["sync_policy"].([]interface{})) > 0 {
					var syncPolicyData nextgen.ApplicationsSyncPolicy
					var syncPolicy = specData["sync_policy"].([]interface{})[0].(map[string]interface{})
					if syncPolicy["sync_options"] != nil && len(syncPolicy["sync_options"].([]interface{})) > 0 {
						var syncOptions []string
						for _, v := range syncPolicy["sync_options"].([]interface{}) {
							syncOptions = append(syncOptions, v.(string))
						}
						syncPolicyData.SyncOptions = syncOptions
					}
					if syncPolicy["automated"] != nil && len(syncPolicy["automated"].([]interface{})) > 0 {
						var automatedSyncPolicyData nextgen.ApplicationsSyncPolicyAutomated
						var automatedSyncPolicy = syncPolicy["automated"].([]interface{})[0].(map[string]interface{})
						if automatedSyncPolicy["prune"] != nil {
							automatedSyncPolicyData.Prune = automatedSyncPolicy["prune"].(bool)
						}
						if automatedSyncPolicy["self_heal"] != nil {
							automatedSyncPolicyData.SelfHeal = automatedSyncPolicy["self_heal"].(bool)
						}
						if automatedSyncPolicy["allow_empty"] != nil {
							automatedSyncPolicyData.AllowEmpty = automatedSyncPolicy["allow_empty"].(bool)
						}
						syncPolicyData.Automated = &automatedSyncPolicyData
					}
					if syncPolicy["retry"] != nil && len(syncPolicy["retry"].([]interface{})) > 0 {
						var retrySync = syncPolicy["retry"].([]interface{})[0].(map[string]interface{})
						var retrySyncData nextgen.ApplicationsRetryStrategy
						if retrySync["limit"] != nil && len(retrySync["limit"].(string)) > 0 {
							retrySyncData.Limit = retrySync["limit"].(string)
						}
						if retrySync["backoff"] != nil && len(retrySync["backoff"].([]interface{})) > 0 {
							var syncBackoff = retrySync["backoff"].([]interface{})[0].(map[string]interface{})
							var syncBackoffData nextgen.ApplicationsBackoff
							if syncBackoff["duration"] != nil && len(syncBackoff["duration"].(string)) > 0 {
								syncBackoffData.Duration = syncBackoff["duration"].(string)
							}
							if syncBackoff["factor"] != nil && len(syncBackoff["factor"].(string)) > 0 {
								syncBackoffData.Factor = syncBackoff["factor"].(string)
							}
							if syncBackoff["max_duration"] != nil && len(syncBackoff["max_duration"].(string)) > 0 {
								syncBackoffData.MaxDuration = syncBackoff["max_duration"].(string)
							}
							retrySyncData.Backoff = &syncBackoffData
						}
						syncPolicyData.Retry = &retrySyncData
					}
					spec.SyncPolicy = &syncPolicyData
				}
			}
		}
	}
	return &nextgen.ApplicationsApplication{
		Metadata: &metaData,
		Spec:     &spec,
	}
}
