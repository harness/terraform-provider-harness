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
		Description: "Datasource for fetching a Harness GitOps Application.",
		ReadContext: datasourceGitopsApplicationRead,
		Schema: map[string]*schema.Schema{
			"account_id": {
				Description: "Account identifier of the GitOps application.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"org_id": {
				Description: "Organization identifier of the GitOps application.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"project_id": {
				Description: "Project identifier of the GitOps application.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"identifier": {
				Description: "Identifier of the GitOps application.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"agent_id": {
				Description: "Agent identifier of the GitOps application.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"cluster_id": {
				Description: "Cluster identifier of the GitOps application.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"repo_id": {
				Description: "Repository identifier of the GitOps application.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"upsert": {
				Description: "Indicates if the GitOps application should be updated if existing and inserted if not.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"validate": {
				Description: "Indicates if the GitOps application yaml has to be validated.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"project": {
				Description: "The ArgoCD project name corresponding to this GitOps application. An empty string means that the GitOps application belongs to the default project created by Harness.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"kind": {
				Description: "Kind of the GitOps application.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"request_propagation_policy": {
				Description: "Request propagation policy to delete the GitOps application.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"request_cascade": {
				Description: "Request cascade to delete the GitOps application.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"options_remove_existing_finalizers": {
				Description: "Options to remove existing finalizers to delete the GitOps application.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"name": {
				Description: "Name of the GitOps application.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"application": {
				Description: "Definition of the GitOps application resource.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"metadata": {
							Description: "Metadata corresponding to the resources. This includes all the objects a user must create.",
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Description: "Name must be unique within a namespace. It is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Name cannot be updated.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"generate_name": {
										Description: "An optional prefix that the server will only apply if the Name field is empty to create a unique name. The name returned to the client will differ from the name passed if this field is used. A unique suffix will be added to this value as well. The supplied value must adhere to the same validation guidelines as the Name field and may be reduced by the suffix length necessary to ensure that it is unique on the server. The server will NOT return a 409 if this field is supplied and the created name already exists; instead, it will either return 201 Created or 500 with Reason ServerTimeout, indicating that a unique name could not be found in the allotted time and the client should try again later.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"namespace": {
										Description: "Namespace of the GitOps application. An empty namespace is equivalent to the namespace of the GitOps agent.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"generation": {
										Description: "A sequence number representing a specific generation of the desired state. This is a read-only value populated by the system.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"uid": {
										Description: "UID is the unique identifier in time and space value for this object. It is generated by the server on successful creation of a resource and is not allowed to change on PUT operations.",
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
										Description: "Annotations are unstructured key value pairs corresponding to a resource. External tools set these to store and retrieve arbitrary metadata.",
										Type:        schema.TypeMap,
										Optional:    true,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"owner_references": {
										Description: "List of objects depended by this object. If ALL objects in the list have been deleted, this object will be garbage collected. If this object is managed by a controller, then an entry in this list will point to this controller, with the controller field set to true. There cannot be more than one managing controller.",
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"api_version": {
													Description: "API version of the referent.",
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"kind": {
													Description: "Kind of the referent.",
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"name": {
													Description: "Name of the referent.",
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"uid": {
													Description: "UID of the referent.",
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"controller": {
													Description: "Indicates if the reference points to the managing controller.",
													Type:        schema.TypeBool,
													Optional:    true,
													Computed:    true,
												},
												"block_owner_deletion": {
													Description: "If true, AND if the owner has the \"foregroundDeletion\" finalizer, then the owner cannot be deleted from the key-value store until this reference is removed. Defaults to false. To set this field, a user needs \"delete\" permission of the owner, otherwise 422 (Unprocessable Entity) will be returned.",
													Type:        schema.TypeBool,
													Optional:    true,
													Computed:    true,
												},
											},
										},
									},
									"finalizers": {
										Description: "Before the object is removed from the register, it must be empty. Each element serves as a unique identifier for the component that is accountable for removing that entry from the list. Entries in this list can only be removed if the object's deletionTimestamp is not null. The processing and removal of finalizers can happen in any sequence. No order is enforced as it may block the finalizers. Finalizers is a shared field that can be reordered by any actor with authority. If the finalizer list is processed in order, this could result in a scenario where the component in charge of the list's first finalizer is waiting for a signal (generated by a field value, an external system, or another) produced by a component in charge of the list's later finalizer.",
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"cluster_name": {
										Description: "Name of the cluster corresponding to the object. API server ignores this if set in any create or update request.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"spec": {
							Description: "Specifications of the GitOps application. This includes the repository URL, application definition, source, destination and sync policy.",
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"source": {
										Description: "Contains all information about the source of a GitOps application.",
										Type:        schema.TypeList,
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"repo_url": {
													Description: "URL to the repository (git or helm) that contains the GitOps application manifests.",
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"path": {
													Description: "Directory path within the git repository, and is only valid for the GitOps applications sourced from git.",
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"target_revision": {
													Description: "Revision of the source to sync the GitOps application to. In case of git, this can be commit, tag, or branch. If omitted, will equal to HEAD. In case of Helm, this is a semver tag of the chart's version.",
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"chart": {
													Description: "Helm chart name, and must be specified for the GitOps applications sourced from a helm repo.",
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"helm": {
													Description: "Holds helm specific options.",
													Type:        schema.TypeList,
													Optional:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"value_files": {
																Description: "List of helm value files to use when generating a template.",
																Type:        schema.TypeList,
																Optional:    true,
																Computed:    true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"release_name": {
																Description: "Helm release name to use. If omitted it will use the GitOps application name.",
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
															},
															"values": {
																Description: "Helm values to be passed to helm template, typically defined as a block.",
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
															},
															"version": {
																Description: "Helm version to use for templating (either \"2\" or \"3\")",
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
															},
															"pass_credentials": {
																Description: "Indicates if to pass credentials to all domains (helm's --pass-credentials)",
																Type:        schema.TypeBool,
																Optional:    true,
																Computed:    true,
															},
															"parameters": {
																Description: "List of helm parameters which are passed to the helm template command upon manifest generation.",
																Type:        schema.TypeList,
																Optional:    true,
																Computed:    true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"name": {
																			Description: "Name of the helm parameter.",
																			Type:        schema.TypeString,
																			Optional:    true,
																			Computed:    true,
																		},
																		"value": {
																			Description: "Value of the helm parameter.",
																			Type:        schema.TypeString,
																			Optional:    true,
																			Computed:    true,
																		},
																		"force_string": {
																			Description: "Indicates if helm should interpret booleans and numbers as strings.",
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Computed:    true,
																		},
																	},
																},
															},
															"file_parameters": {
																Description: "File parameters to the helm template.",
																Type:        schema.TypeList,
																Optional:    true,
																Computed:    true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"name": {
																			Description: "Name of the helm parameter.",
																			Type:        schema.TypeString,
																			Optional:    true,
																			Computed:    true,
																		},
																		"path": {
																			Description: "Path to the file containing the values of the helm parameter.",
																			Type:        schema.TypeString,
																			Optional:    true,
																			Computed:    true,
																		},
																	},
																},
															},
														},
													},
												},
												"kustomize": {
													Description: "Options specific to a GitOps application source specific to Kustomize.",
													Type:        schema.TypeList,
													Optional:    true,
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"name_prefix": {
																Description: "Prefix prepended to resources for kustomize apps.",
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
															},
															"name_suffix": {
																Description: "Suffix appended to resources for kustomize apps.",
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
															},
															"images": {
																Description: "List of kustomize image override specifications.",
																Type:        schema.TypeList,
																Optional:    true,
																Computed:    true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"common_labels": {
																Description: "List of additional labels to add to rendered manifests.",
																Type:        schema.TypeMap,
																Optional:    true,
																Computed:    true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"version": {
																Description: "Version of kustomize to use for rendering manifests.",
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
															},
															"common_annotations": {
																Description: "List of additional annotations to add to rendered manifests.",
																Type:        schema.TypeMap,
																Optional:    true,
																Computed:    true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"force_common_labels": {
																Description: "Indicates if to force apply common labels to resources for kustomize apps.",
																Type:        schema.TypeBool,
																Optional:    true,
																Computed:    true,
															},
															"force_common_annotations": {
																Description: "Indicates if to force applying common annotations to resources for kustomize apps.",
																Type:        schema.TypeBool,
																Optional:    true,
																Computed:    true,
															},
														},
													},
												},
												"ksonnet": {
													Description: "Ksonnet specific options.",
													Type:        schema.TypeList,
													Optional:    true,
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"environment": {
																Description: "Ksonnet application environment name.",
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
															},
															"parameters": {
																Description: "List of ksonnet component parameter override values.",
																Type:        schema.TypeList,
																Optional:    true,
																Computed:    true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"component": {
																			Description: "Component of the parameter of the ksonnet application.",
																			Type:        schema.TypeString,
																			Optional:    true,
																			Computed:    true,
																		},
																		"name": {
																			Description: "Name of the parameter of the ksonnet application.",
																			Type:        schema.TypeString,
																			Optional:    true,
																			Computed:    true,
																		},
																		"value": {
																			Description: "Value of the parameter of the ksonnet application.",
																			Type:        schema.TypeString,
																			Optional:    true,
																			Computed:    true,
																		},
																	},
																},
															},
														},
													},
												},
												"directory": {
													Description: "Options for applications of type plain YAML or Jsonnet.",
													Type:        schema.TypeList,
													Optional:    true,
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"recurse": {
																Description: "Indicates to scan a directory recursively for manifests.",
																Type:        schema.TypeBool,
																Optional:    true,
																Computed:    true,
															},
															"exclude": {
																Description: "Glob pattern to match paths against that should be explicitly excluded from being used during manifest generation.",
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
															},
															"include": {
																Description: "Glob pattern to match paths against that should be explicitly included during manifest generation.",
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
															},
															"jsonnet": {
																Description: "Options specific to applications of type Jsonnet.",
																Type:        schema.TypeList,
																Optional:    true,
																Computed:    true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"libs": {
																			Description: "Additional library search dirs.",
																			Type:        schema.TypeList,
																			Optional:    true,
																			Computed:    true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																		"ext_vars": {
																			Description: "List of jsonnet external variables.",
																			Type:        schema.TypeList,
																			Optional:    true,
																			Computed:    true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"name": {
																						Description: "Name of the external variables of jsonnet application.",
																						Type:        schema.TypeString,
																						Optional:    true,
																						Computed:    true,
																					},
																					"value": {
																						Description: "Value of the external variables of jsonnet application.",
																						Type:        schema.TypeString,
																						Optional:    true,
																						Computed:    true,
																					},
																					"code": {
																						Description: "Code of the external variables of jsonnet application.",
																						Type:        schema.TypeBool,
																						Optional:    true,
																						Computed:    true,
																					},
																				},
																			},
																		},
																		"tlas": {
																			Description: "List of jsonnet top-level arguments(TLAS).",
																			Type:        schema.TypeList,
																			Optional:    true,
																			Computed:    true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"name": {
																						Description: "Name of the TLAS of the jsonnet application.",
																						Type:        schema.TypeString,
																						Optional:    true,
																						Computed:    true,
																					},
																					"value": {
																						Description: "Value of the TLAS of the jsonnet application.",
																						Type:        schema.TypeString,
																						Optional:    true,
																						Computed:    true,
																					},
																					"code": {
																						Description: "Code of the TLAS of the jsonnet application.",
																						Type:        schema.TypeBool,
																						Optional:    true,
																						Computed:    true,
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
													Description: "Options specific to config management plugins.",
													Type:        schema.TypeList,
													Optional:    true,
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"name": {
																Description: "Name of the plugin.",
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
															},
															"env": {
																Description: "Entry in the GitOps application's environment.",
																Type:        schema.TypeList,
																Optional:    true,
																Computed:    true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"name": {
																			Description: "Name of the variable, usually expressed in uppercase.",
																			Type:        schema.TypeString,
																			Optional:    true,
																			Computed:    true,
																		},
																		"value": {
																			Description: "Value of the variable.",
																			Type:        schema.TypeString,
																			Optional:    true,
																			Computed:    true,
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
										Description: "Information about the GitOps application's destination.",
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Description: "URL of the target cluster and must be set to the kubernetes control plane API.",
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"namespace": {
													Description: "Target namespace of the GitOps application's resources. The namespace will only be set for namespace-scoped resources that have not set a value for .metadata.namespace.",
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"server": {
													Description: "URL of the target cluster server for the GitOps application.",
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
											},
										},
									},
									"sync_policy": {
										Description: "Controls when a sync will be performed in response to updates in git.",
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"sync_options": {
													Description: "Options allow you to specify whole app sync-options.",
													Type:        schema.TypeList,
													Optional:    true,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"automated": {
													Description: "Controls the behavior of an automated sync.",
													Type:        schema.TypeList,
													Optional:    true,
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"prune": {
																Description: "Indicates whether to delete resources from the cluster that are not found in the sources anymore as part of automated sync (default: false).",
																Type:        schema.TypeBool,
																Optional:    true,
																Computed:    true,
															},
															"self_heal": {
																Description: "Indicates whether to revert resources back to their desired state upon modification in the cluster (default: false).",
																Type:        schema.TypeBool,
																Optional:    true,
																Computed:    true,
															},
															"allow_empty": {
																Description: "Indicates to allows apps to have zero live resources (default: false).",
																Type:        schema.TypeBool,
																Optional:    true,
																Computed:    true,
															},
														},
													},
												},
												"retry": {
													Description: "Contains information about the strategy to apply when a sync failed.",
													Type:        schema.TypeList,
													Optional:    true,
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"limit": {
																Description: "Limit is the maximum number of attempts for retrying a failed sync. If set to 0, no retries will be performed.",
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
															},
															"backoff": {
																Description: "Backoff strategy to use on subsequent retries for failing syncs.",
																Type:        schema.TypeList,
																Optional:    true,
																Computed:    true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"duration": {
																			Description: "Amount to back off. Default unit is seconds, but could also be a duration (e.g. \"2m\", \"1h\").",
																			Type:        schema.TypeString,
																			Optional:    true,
																			Computed:    true,
																		},
																		"factor": {
																			Description: "Factor to multiply the base duration after each failed retry.",
																			Type:        schema.TypeString,
																			Optional:    true,
																			Computed:    true,
																		},
																		"max_duration": {
																			Description: "Maximum amount of time allowed of the backoff strategy.",
																			Type:        schema.TypeString,
																			Optional:    true,
																			Computed:    true,
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
	var agentIdentifier, orgIdentifier, projectIdentifier, repoIdentifier, queryName string
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

	resp, httpResp, err := c.ApplicationsApiService.AgentApplicationServiceGet(ctx, agentIdentifier, queryName, c.AccountId, orgIdentifier, projectIdentifier, &nextgen.ApplicationsApiAgentApplicationServiceGetOpts{
		QueryRepo: optional.NewString(repoIdentifier),
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
