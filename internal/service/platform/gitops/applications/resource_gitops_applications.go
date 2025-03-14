package applications

import (
	"context"
	"fmt"
	"reflect"
	"sort"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceGitopsApplication() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for managing a Harness Gitops Application.",

		CreateContext: resourceGitopsApplicationCreate,
		ReadContext:   resourceGitopsApplicationRead,
		UpdateContext: resourceGitopsApplicationUpdate,
		DeleteContext: resourceGitopsApplicationDelete,
		Importer:      helpers.GitopsAgentApplicationImporter,

		CustomizeDiff: func(ctx context.Context, diff *schema.ResourceDiff, i interface{}) error {
			var e error
			if diff.HasChange("project_id") && diff.Id() != "" {
				e = fmt.Errorf("field 'project_id' cannot be changed after the resource is created")
			}
			if diff.HasChange("org_id") && diff.Id() != "" {
				if e != nil {
					e = fmt.Errorf("field 'org_id' cannot be changed after the resource is created:%w", e)
				} else {
					e = fmt.Errorf("field 'org_id' cannot be changed after the resource is created")
				}
			}
			if diff.HasChange("agent_id") && diff.Id() != "" {
				if e != nil {
					e = fmt.Errorf("field 'agent_id' cannot be changed after the resource is created:%v", e)
				} else {
					e = fmt.Errorf("field 'agent_id' cannot be changed after the resource is created")
				}
			}
			if diff.HasChange("name") && diff.Id() != "" {
				if e != nil {
					e = fmt.Errorf("field 'name' cannot be changed after the resource is created:%w", e)
				} else {
					e = fmt.Errorf("field 'name' cannot be changed after the resource is created")
				}
			}
			return e
		},

		Schema: map[string]*schema.Schema{
			"account_id": {
				Description: "Account identifier of the GitOps application.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Deprecated:  "This field is deprecated and will be removed in a future release.",
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
				Deprecated:  "This field is deprecated and will be removed in a future release.",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					// Implement logic to suppress the diff here
					return true
				},
			},
			"agent_id": {
				Description: "Agent identifier of the GitOps application.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"cluster_id": {
				Description: "Cluster identifier of the GitOps application.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"repo_id": {
				Description: "Repository identifier of the GitOps application. When using skipRepoValidation, this field is not required.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"repo_ids": {
				Description: "List of repository identifiers of the GitOps for Multi-Source application. Not required if skipRepoValidation is set to true.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					// Convert old and new string lists into slices
					oldIF, newIF := d.GetChange("repo_ids")
					oldSlice := oldIF.([]interface{})
					newSlice := newIF.([]interface{})

					// Sort the slices to ignore order differences
					sort.Slice(oldSlice, func(i, j int) bool {
						return oldSlice[i].(string) < oldSlice[j].(string)
					})
					sort.Slice(newSlice, func(i, j int) bool {
						return newSlice[i].(string) < newSlice[j].(string)
					})

					// Compare sorted slices
					return reflect.DeepEqual(oldSlice, newSlice)
				},
			},
			"upsert": {
				Description: "Indicates if the GitOps application should be updated if existing and inserted if not.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"validate": {
				Description: "Indicates if the GitOps application yaml has to be validated.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"project": {
				Description: "The ArgoCD project name corresponding to this GitOps application. An empty string means that the GitOps application belongs to the default project created by Harness.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"kind": {
				Description: "Kind of the GitOps application.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"request_propagation_policy": {
				Description: "Request propagation policy to delete the GitOps application.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"request_cascade": {
				Description: "Request cascade to delete the GitOps application.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"options_remove_existing_finalizers": {
				Description: "Options to remove existing finalizers to delete the GitOps application.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"name": {
				Description: "Name of the GitOps application.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"skip_repo_validation": {
				Description:   "Indicates if the GitOps application should skip validate repository definition exists.",
				Type:          schema.TypeBool,
				Optional:      true,
				ConflictsWith: []string{"repo_id", "repo_ids"},
			},
			"application": {
				Description: "Definition of the GitOps application resource.",
				Type:        schema.TypeList,
				Required:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"metadata": {
							Description: "Metadata corresponding to the resources. This includes all the objects a user must create.",
							Type:        schema.TypeList,
							Required:    true,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Description: "Name must be unique within a namespace. It is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Name cannot be updated.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"generate_name": {
										Description: "An optional prefix that the server will only apply if the Name field is empty to create a unique name. The name returned to the client will differ from the name passed if this field is used. A unique suffix will be added to this value as well. The supplied value must adhere to the same validation guidelines as the Name field and may be reduced by the suffix length necessary to ensure that it is unique on the server. The server will NOT return a 409 if this field is supplied and the created name already exists; instead, it will either return 201 Created or 500 with Reason ServerTimeout, indicating that a unique name could not be found in the allotted time and the client should try again later.",
										Type:        schema.TypeString,
										Optional:    true,
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
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"project": {
										Description: "The ArgoCD project name corresponding to this GitOps application. Value must match mappings of ArgoCD projects to harness project.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"source": {
										Description:   "Contains all information about the source of the GitOps application.",
										Type:          schema.TypeList,
										Optional:      true,
										ConflictsWith: []string{"application.0.spec.0.sources"},
										MaxItems:      1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"repo_url": {
													Description: "URL to the repository (git or helm) that contains the GitOps application manifests.",
													Type:        schema.TypeString,
													Required:    true,
												},
												"path": {
													Description: "Directory path within the git repository, and is only valid for the GitOps applications sourced from git.",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"target_revision": {
													Description: "Revision of the source to sync the GitOps application to. In case of git, this can be commit, tag, or branch. If omitted, will equal to HEAD. In case of Helm, this is a semver tag of the chart's version.",
													Type:        schema.TypeString,
													Required:    true,
												},
												"chart": {
													Description: "Helm chart name, and must be specified for the GitOps applications sourced from a helm repo.",
													Type:        schema.TypeString,
													Optional:    true,
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
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"release_name": {
																Description: "Helm release name to use. If omitted it will use the GitOps application name.",
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
																Description: "Indicates if to pass credentials to all domains (helm's --pass-credentials)",
																Type:        schema.TypeBool,
																Optional:    true,
															},
															"parameters": {
																Description: "List of helm parameters which are passed to the helm template command upon manifest generation.",
																Type:        schema.TypeList,
																Optional:    true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"name": {
																			Description: "Name of the helm parameter.",
																			Type:        schema.TypeString,
																			Optional:    true,
																		},
																		"value": {
																			Description: "Value of the Helm parameter.",
																			Type:        schema.TypeString,
																			Optional:    true,
																		},
																		"force_string": {
																			Description: "Indicates if helm should interpret booleans and numbers as strings.",
																			Type:        schema.TypeBool,
																			Optional:    true,
																		},
																	},
																},
															},
															"file_parameters": {
																Description: "File parameters to the helm template.",
																Type:        schema.TypeList,
																Optional:    true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"name": {
																			Description: "Name of the helm parameter.",
																			Type:        schema.TypeString,
																			Optional:    true,
																		},
																		"path": {
																			Description: "Path to the file containing the values of the helm parameter.",
																			Type:        schema.TypeString,
																			Optional:    true,
																		},
																	},
																},
															},
															"ignore_missing_value_files": {
																Description: "Prevents 'helm template' from failing when value_files do not exist locally.",
																Type:        schema.TypeBool,
																Optional:    true,
															},
															"skip_crds": {
																Description: "Indicates if to skip CRDs during helm template. Corresponds to helm --skip-crds",
																Type:        schema.TypeBool,
																Optional:    true,
															},
															"skip_tests": {
																Description: "Indicates if to skip tests during helm template. Corresponds to helm --skip-tests",
																Type:        schema.TypeBool,
																Optional:    true,
															},
															"skip_schema_validation": {
																Description: "Indicates if to skip schema validation during helm template. Corresponds to helm --skip-schema-validation",
																Type:        schema.TypeBool,
																Optional:    true,
															},
														},
													},
												},
												"kustomize": {
													Description: "Options specific to a GitOps application source specific to Kustomize.",
													Type:        schema.TypeList,
													Optional:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"name_prefix": {
																Description: "Prefix prepended to resources for kustomize apps.",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"name_suffix": {
																Description: "Suffix appended to resources for kustomize apps.",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"images": {
																Description: "List of kustomize image override specifications.",
																Type:        schema.TypeList,
																Optional:    true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"common_labels": {
																Description: "List of additional labels to add to rendered manifests.",
																Type:        schema.TypeMap,
																Optional:    true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"version": {
																Description: "Version of kustomize to use for rendering manifests.",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"common_annotations": {
																Description: "List of additional annotations to add to rendered manifests.",
																Type:        schema.TypeMap,
																Optional:    true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"force_common_labels": {
																Description: "Indicates if to force apply common labels to resources for kustomize apps.",
																Type:        schema.TypeBool,
																Optional:    true,
															},
															"force_common_annotations": {
																Description: "Indicates if to force applying common annotations to resources for kustomize apps.",
																Type:        schema.TypeBool,
																Optional:    true,
															},
														},
													},
												},
												"ksonnet": {
													Description: "Ksonnet specific options.",
													Type:        schema.TypeList,
													Optional:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"environment": {
																Description: "Ksonnet application environment name.",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"parameters": {
																Description: "List of ksonnet component parameter override values.",
																Type:        schema.TypeList,
																Optional:    true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"component": {
																			Description: "Component of the parameter of the ksonnet application.",
																			Type:        schema.TypeString,
																			Optional:    true,
																		},
																		"name": {
																			Description: "Name of the parameter of the ksonnet application.",
																			Type:        schema.TypeString,
																			Optional:    true,
																		},
																		"value": {
																			Description: "Value of the parameter of the ksonnet application.",
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
													Description: "Options for applications of type plain YAML or Jsonnet.",
													Type:        schema.TypeList,
													Optional:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"recurse": {
																Description: "Indicates to scan a directory recursively for manifests.",
																Type:        schema.TypeBool,
																Optional:    true,
															},
															"exclude": {
																Description: "Glob pattern to match paths against that should be explicitly excluded from being used during manifest generation.",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"include": {
																Description: "Glob pattern to match paths against that should be explicitly included during manifest generation.",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"jsonnet": {
																Description: "Options specific to applications of type Jsonnet.",
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
																			Description: "List of jsonnet external variables.",
																			Type:        schema.TypeList,
																			Optional:    true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"name": {
																						Description: "Name of the external variables of jsonnet application.",
																						Type:        schema.TypeString,
																						Optional:    true,
																					},
																					"value": {
																						Description: "Value of the external variables of jsonnet application.",
																						Type:        schema.TypeString,
																						Optional:    true,
																					},
																					"code": {
																						Description: "Code of the external variables of jsonnet application.",
																						Type:        schema.TypeBool,
																						Optional:    true,
																					},
																				},
																			},
																		},
																		"tlas": {
																			Description: "List of jsonnet top-level arguments(TLAS).",
																			Type:        schema.TypeList,
																			Optional:    true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"name": {
																						Description: "Name of the TLAS of the jsonnet application.",
																						Type:        schema.TypeString,
																						Optional:    true,
																					},
																					"value": {
																						Description: "Value of the TLAS of the jsonnet application.",
																						Type:        schema.TypeString,
																						Optional:    true,
																					},
																					"code": {
																						Description: "Code of the TLAS of the jsonnet application.",
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
													Description: "Options specific to config management plugins.",
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
																Description: "Entry in the GitOps application's environment.",
																Type:        schema.TypeList,
																Optional:    true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"name": {
																			Description: "Name of the variable, usually expressed in uppercase.",
																			Type:        schema.TypeString,
																			Optional:    true,
																		},
																		"value": {
																			Description: "Value of the variable.",
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
									"sources": {
										Description:   "List of sources for the GitOps application. Multi Source support",
										Type:          schema.TypeList,
										Optional:      true,
										ConflictsWith: []string{"application.0.spec.0.source"},
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"repo_url": {
													Description: "URL to the repository (git or helm) that contains the GitOps application manifests.",
													Type:        schema.TypeString,
													Required:    true,
												},
												"path": {
													Description: "Directory path within the git repository, and is only valid for the GitOps applications sourced from git.",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"target_revision": {
													Description: "Revision of the source to sync the GitOps application to. In case of git, this can be commit, tag, or branch. If omitted, will equal to HEAD. In case of Helm, this is a semver tag of the chart's version.",
													Type:        schema.TypeString,
													Required:    true,
												},
												"chart": {
													Description: "Helm chart name, and must be specified for the GitOps applications sourced from a helm repo.",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"ref": {
													Description: "Reference name to be used in other source spec, used for multi-source applications.",
													Type:        schema.TypeString,
													Optional:    true,
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
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"release_name": {
																Description: "Helm release name to use. If omitted it will use the GitOps application name.",
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
																Description: "Indicates if to pass credentials to all domains (helm's --pass-credentials)",
																Type:        schema.TypeBool,
																Optional:    true,
															},
															"parameters": {
																Description: "List of helm parameters which are passed to the helm template command upon manifest generation.",
																Type:        schema.TypeList,
																Optional:    true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"name": {
																			Description: "Name of the helm parameter.",
																			Type:        schema.TypeString,
																			Optional:    true,
																		},
																		"value": {
																			Description: "Value of the Helm parameter.",
																			Type:        schema.TypeString,
																			Optional:    true,
																		},
																		"force_string": {
																			Description: "Indicates if helm should interpret booleans and numbers as strings.",
																			Type:        schema.TypeBool,
																			Optional:    true,
																		},
																	},
																},
															},
															"file_parameters": {
																Description: "File parameters to the helm template.",
																Type:        schema.TypeList,
																Optional:    true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"name": {
																			Description: "Name of the helm parameter.",
																			Type:        schema.TypeString,
																			Optional:    true,
																		},
																		"path": {
																			Description: "Path to the file containing the values of the helm parameter.",
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
													Description: "Options specific to a GitOps application source specific to Kustomize.",
													Type:        schema.TypeList,
													Optional:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"name_prefix": {
																Description: "Prefix prepended to resources for kustomize apps.",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"name_suffix": {
																Description: "Suffix appended to resources for kustomize apps.",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"images": {
																Description: "List of kustomize image override specifications.",
																Type:        schema.TypeList,
																Optional:    true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"common_labels": {
																Description: "List of additional labels to add to rendered manifests.",
																Type:        schema.TypeMap,
																Optional:    true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"version": {
																Description: "Version of kustomize to use for rendering manifests.",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"common_annotations": {
																Description: "List of additional annotations to add to rendered manifests.",
																Type:        schema.TypeMap,
																Optional:    true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"force_common_labels": {
																Description: "Indicates if to force apply common labels to resources for kustomize apps.",
																Type:        schema.TypeBool,
																Optional:    true,
															},
															"force_common_annotations": {
																Description: "Indicates if to force applying common annotations to resources for kustomize apps.",
																Type:        schema.TypeBool,
																Optional:    true,
															},
														},
													},
												},
												"ksonnet": {
													Description: "Ksonnet specific options.",
													Type:        schema.TypeList,
													Optional:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"environment": {
																Description: "Ksonnet application environment name.",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"parameters": {
																Description: "List of ksonnet component parameter override values.",
																Type:        schema.TypeList,
																Optional:    true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"component": {
																			Description: "Component of the parameter of the ksonnet application.",
																			Type:        schema.TypeString,
																			Optional:    true,
																		},
																		"name": {
																			Description: "Name of the parameter of the ksonnet application.",
																			Type:        schema.TypeString,
																			Optional:    true,
																		},
																		"value": {
																			Description: "Value of the parameter of the ksonnet application.",
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
													Description: "Options for applications of type plain YAML or Jsonnet.",
													Type:        schema.TypeList,
													Optional:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"recurse": {
																Description: "Indicates to scan a directory recursively for manifests.",
																Type:        schema.TypeBool,
																Optional:    true,
															},
															"exclude": {
																Description: "Glob pattern to match paths against that should be explicitly excluded from being used during manifest generation.",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"include": {
																Description: "Glob pattern to match paths against that should be explicitly included during manifest generation.",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"jsonnet": {
																Description: "Options specific to applications of type Jsonnet.",
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
																			Description: "List of jsonnet external variables.",
																			Type:        schema.TypeList,
																			Optional:    true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"name": {
																						Description: "Name of the external variables of jsonnet application.",
																						Type:        schema.TypeString,
																						Optional:    true,
																					},
																					"value": {
																						Description: "Value of the external variables of jsonnet application.",
																						Type:        schema.TypeString,
																						Optional:    true,
																					},
																					"code": {
																						Description: "Code of the external variables of jsonnet application.",
																						Type:        schema.TypeBool,
																						Optional:    true,
																					},
																				},
																			},
																		},
																		"tlas": {
																			Description: "List of jsonnet top-level arguments(TLAS).",
																			Type:        schema.TypeList,
																			Optional:    true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"name": {
																						Description: "Name of the TLAS of the jsonnet application.",
																						Type:        schema.TypeString,
																						Optional:    true,
																					},
																					"value": {
																						Description: "Value of the TLAS of the jsonnet application.",
																						Type:        schema.TypeString,
																						Optional:    true,
																					},
																					"code": {
																						Description: "Code of the TLAS of the jsonnet application.",
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
													Description: "Options specific to config management plugins.",
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
																Description: "Entry in the GitOps application's environment.",
																Type:        schema.TypeList,
																Optional:    true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"name": {
																			Description: "Name of the variable, usually expressed in uppercase.",
																			Type:        schema.TypeString,
																			Optional:    true,
																		},
																		"value": {
																			Description: "Value of the variable.",
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
										Description: "Information about the GitOps application's destination.",
										Type:        schema.TypeList,
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Description: "URL of the target cluster and must be set to the kubernetes control plane API.",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"namespace": {
													Description: "Target namespace of the GitOps application's resources. The namespace will only be set for namespace-scoped resources that have not set a value for .metadata.namespace.",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"server": {
													Description: "URL of the target cluster server for the GitOps application.",
													Type:        schema.TypeString,
													Optional:    true,
												},
											},
										},
									},
									"sync_policy": {
										Description: "Controls when a sync will be performed in response to updates in git.",
										Type:        schema.TypeList,
										Optional:    true,
										MaxItems:    1,
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
													Description: "Controls the behavior of an automated sync.",
													Type:        schema.TypeList,
													Optional:    true,
													MaxItems:    1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"prune": {
																Description: "Indicates whether to delete resources from the cluster that are not found in the sources anymore as part of automated sync (default: false).",
																Type:        schema.TypeBool,
																Optional:    true,
															},
															"self_heal": {
																Description: "Indicates whether to revert resources back to their desired state upon modification in the cluster (default: false).",
																Type:        schema.TypeBool,
																Optional:    true,
															},
															"allow_empty": {
																Description: "Indicates to allows apps to have zero live resources (default: false).",
																Type:        schema.TypeBool,
																Optional:    true,
															},
														},
													},
												},
												"retry": {
													Description: "Contains information about the strategy to apply when a sync failed.",
													Type:        schema.TypeList,
													Optional:    true,
													MaxItems:    1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"limit": {
																Description: "Limit is the maximum number of attempts for retrying a failed sync. If set to 0, no retries will be performed.",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"backoff": {
																Description: "Backoff strategy to use on subsequent retries for failing syncs.",
																Type:        schema.TypeList,
																Optional:    true,
																MaxItems:    1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"duration": {
																			Description: "Amount to back off. Default unit is seconds, but could also be a duration (e.g. \"2m\", \"1h\").",
																			Type:        schema.TypeString,
																			Optional:    true,
																		},
																		"factor": {
																			Description: "Factor to multiply the base duration after each failed retry.",
																			Type:        schema.TypeString,
																			Optional:    true,
																		},
																		"max_duration": {
																			Description: "Maximum amount of time allowed of the backoff strategy.",
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
	var skipRepoValidation bool
	var repoIdentifiersList []string
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
	if attr, ok := d.GetOk("skip_repo_validation"); ok {
		skipRepoValidation = attr.(bool)
	}
	if attr, ok := d.GetOk("repo_ids"); ok {
		repoIdentifiers := attr.([]interface{})
		repoIdentifiersList = make([]string, len(repoIdentifiers))
		for i, repo := range repoIdentifiers {
			repoIdentifiersList[i] = repo.(string)
		}
	}

	resp, httpResp, err := c.ApplicationsApiService.AgentApplicationServiceCreate(ctx, createApplicationRequest, agentIdentifier, &nextgen.ApplicationsApiAgentApplicationServiceCreateOpts{
		AccountIdentifier:  optional.NewString(c.AccountId),
		OrgIdentifier:      optional.NewString(orgIdentifier),
		ProjectIdentifier:  optional.NewString(projectIdentifier),
		ClusterIdentifier:  optional.NewString(clusterIdentifier),
		RepoIdentifier:     optional.NewString(repoIdentifier),
		RepoIdentifiers:    optional.NewInterface(repoIdentifiersList),
		SkipRepoValidation: optional.NewBool(skipRepoValidation),
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
	err = setApplication(d, &resp)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGitopsApplicationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	// name is required, so must exist
	if attr, ok := d.GetOk("name"); ok {
		queryName = attr.(string)
	}

	if attr, ok := d.GetOk("repo_id"); ok {
		repoIdentifier = attr.(string)
	}
	resp, httpResp, err := c.ApplicationsApiService.AgentApplicationServiceGet(ctx, agentIdentifier, queryName, c.AccountId, orgIdentifier, projectIdentifier, &nextgen.ApplicationsApiAgentApplicationServiceGetOpts{
		QueryRepo: optional.NewString(repoIdentifier),
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
	err = setApplication(d, &resp)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGitopsApplicationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	updateApplicationRequest := buildUpdateApplicationRequest(d)
	var agentIdentifier, orgIdentifier, projectIdentifier, clusterIdentifier, repoIdentifier, appMetaDataName string
	var skipRepoValidation bool
	var repoIdentifiersList []string

	var e diag.Diagnostics
	if d.HasChange("name") {
		oldValue, newValue := d.GetChange("name")
		if oldValue != "" && oldValue != newValue {
			e = append(e, diag.Errorf("%s", "Field 'name' cannot be updated after creation.")[0])
		}
		if err := d.Set("name", oldValue); err != nil {
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

	if d.HasChange("agent_id") {
		oldValue, newValue := d.GetChange("agent_id")
		if oldValue != "" && oldValue != newValue {
			e = append(e, diag.Errorf("%s", "Field 'agent_id' cannot be updated after creation.")[0])
		}
		if err := d.Set("agent_id", oldValue); err != nil {
			return diag.FromErr(err)
		}
	}

	if len(e) > 0 {
		return e
	}

	if attr, ok := d.GetOk("agent_id"); ok {
		agentIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("name"); ok {
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
	if attr, ok := d.GetOk("skip_repo_validation"); ok {
		skipRepoValidation = attr.(bool)
	}
	if attr, ok := d.GetOk("repo_ids"); ok {
		repoIdentifiers := attr.([]interface{})
		repoIdentifiersList = make([]string, len(repoIdentifiers))
		for i, repo := range repoIdentifiers {
			repoIdentifiersList[i] = repo.(string)
		}
	}

	resp, httpResp, err := c.ApplicationsApiService.AgentApplicationServiceUpdate(ctx, updateApplicationRequest, c.AccountId, orgIdentifier, projectIdentifier, agentIdentifier, appMetaDataName, &nextgen.ApplicationsApiAgentApplicationServiceUpdateOpts{
		ClusterIdentifier:  optional.NewString(clusterIdentifier),
		RepoIdentifier:     optional.NewString(repoIdentifier),
		SkipRepoValidation: optional.NewBool(skipRepoValidation),
		RepoIdentifiers:    optional.NewInterface(repoIdentifiersList),
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
	err = setApplication(d, &resp)
	if err != nil {
		return diag.FromErr(err)
	}

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
	if attr, ok := d.GetOk("name"); ok {
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

func setApplication(d *schema.ResourceData, app *nextgen.Servicev1Application) error {
	d.SetId(app.Name)
	d.Set("org_id", app.OrgIdentifier)
	d.Set("project_id", app.ProjectIdentifier)
	d.Set("agent_id", app.AgentIdentifier)
	d.Set("account_id", app.AccountIdentifier)
	d.Set("cluster_id", app.ClusterIdentifier)
	d.Set("repo_id", app.RepoIdentifier)
	d.Set("name", app.Name)
	d.Set("skip_repo_validation", app.SkipRepoValidation)
	if app.RepoIdentifiers != nil {
		d.Set("repo_ids", app.RepoIdentifiers)
	}

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
			if app.App.Metadata.Finalizers != nil {
				metadata["finalizers"] = app.App.Metadata.Finalizers
			}
			if app.App.Metadata.OwnerReferences != nil {
				var ownerReferencesList = []interface{}{}
				for _, ownerReference := range app.App.Metadata.OwnerReferences {
					var ownerReferenceMap = map[string]interface{}{}
					ownerReferenceMap["api_version"] = ownerReference.ApiVersion
					ownerReferenceMap["kind"] = ownerReference.Kind
					ownerReferenceMap["name"] = ownerReference.Name
					ownerReferenceMap["uid"] = ownerReference.Uid
					ownerReferencesList = append(ownerReferencesList, ownerReferenceMap)
				}
				metadata["owner_references"] = ownerReferencesList
			}
			metadataList = append(metadataList, metadata)
			application["metadata"] = metadataList
		}

		if app.App.Spec != nil {
			var specList = []interface{}{}
			var spec = map[string]interface{}{}
			fmt.Println("app.App.Spec: ", app.App.Spec)
			if app.App.Spec.Project != "" {
				spec["project"] = app.App.Spec.Project
			}
			if app.App.Spec.Source != nil {
				source := getSourceForState(app.App.Spec)
				spec["source"] = []interface{}{source}
			}
			if len(app.App.Spec.Sources) > 0 {
				var sourcesList = []interface{}{}
				for _, source := range app.App.Spec.Sources {
					sourcesList = append(sourcesList, source)
				}
				spec["sources"] = sourcesList
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

		err := d.Set("application", applicationList)
		if err != nil {
			return fmt.Errorf("error setting application: %v", err)
		}
	}
	return nil
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
					metaData.Generation = meta["generation"].(string)
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
				project := specData["project"].(string)
				spec.Project = project
				if specData["source"] != nil && len(specData["source"].([]interface{})) > 0 {
					sourceMap := specData["source"].([]interface{})[0].(map[string]interface{})
					source := setSpecSourceForRequest(sourceMap)
					spec.Source = source
				}
				if specData["sources"] != nil && len(specData["sources"].([]interface{})) > 0 {
					var sources []nextgen.ApplicationsApplicationSource
					for _, v := range specData["sources"].([]interface{}) {
						source := setSpecSourceForRequest(v.(map[string]interface{}))
						sources = append(sources, *source)
					}
					spec.Sources = sources
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

func getSourceForState(appSpec *nextgen.ApplicationsApplicationSpec) map[string]interface{} {
	var source = map[string]interface{}{}
	source["repo_url"] = appSpec.Source.RepoURL
	source["path"] = appSpec.Source.Path
	source["target_revision"] = appSpec.Source.TargetRevision
	source["chart"] = appSpec.Source.Chart
	if appSpec.Source.Helm != nil {
		var helmList = []interface{}{}
		var helm = map[string]interface{}{}
		if appSpec.Source.Helm.ValueFiles != nil && len(appSpec.Source.Helm.ValueFiles) > 0 {
			helm["value_files"] = appSpec.Source.Helm.ValueFiles
		}
		helm["release_name"] = appSpec.Source.Helm.ReleaseName
		helm["values"] = appSpec.Source.Helm.Values
		helm["version"] = appSpec.Source.Helm.Version
		helm["pass_credentials"] = appSpec.Source.Helm.PassCredentials
		if appSpec.Source.Helm.Parameters != nil && len(appSpec.Source.Helm.Parameters) > 0 {
			var helmParametersList = []interface{}{}
			for _, v := range appSpec.Source.Helm.Parameters {
				var helmParam = map[string]interface{}{}
				helmParam["name"] = v.Name
				helmParam["value"] = v.Value
				helmParam["force_string"] = v.ForceString
				helmParametersList = append(helmParametersList, helmParam)
			}
			helm["parameters"] = helmParametersList
		}
		if appSpec.Source.Helm.FileParameters != nil && len(appSpec.Source.Helm.FileParameters) > 0 {
			var helmFileParametersList = []interface{}{}
			for _, v := range appSpec.Source.Helm.FileParameters {
				var helmParam = map[string]interface{}{}
				helmParam["name"] = v.Name
				helmParam["path"] = v.Path
				helmFileParametersList = append(helmFileParametersList, helmParam)
			}
			helm["file_parameters"] = helmFileParametersList
		}
		helm["ignore_missing_value_files"] = appSpec.Source.Helm.IgnoreMissingValueFiles
		helm["skip_crds"] = appSpec.Source.Helm.SkipCrds
		helm["skip_tests"] = appSpec.Source.Helm.SkipTests
		helm["skip_schema_validation"] = appSpec.Source.Helm.SkipSchemaValidation

		helmList = append(helmList, helm)
		source["helm"] = helmList
	}
	if appSpec.Source.Kustomize != nil {
		var kustomizeList = []interface{}{}
		var kustomize = map[string]interface{}{}
		kustomize["name_prefix"] = appSpec.Source.Kustomize.NamePrefix
		kustomize["name_suffix"] = appSpec.Source.Kustomize.NameSuffix
		kustomize["images"] = appSpec.Source.Kustomize.Images
		kustomize["common_labels"] = appSpec.Source.Kustomize.CommonLabels
		kustomize["version"] = appSpec.Source.Kustomize.Version
		kustomize["common_annotations"] = appSpec.Source.Kustomize.CommonAnnotations
		kustomize["force_common_labels"] = appSpec.Source.Kustomize.ForceCommonLabels
		kustomize["force_common_annotations"] = appSpec.Source.Kustomize.ForceCommonAnnotations
		kustomizeList = append(kustomizeList, kustomize)
		source["kustomize"] = kustomizeList
	}
	if appSpec.Source.Ksonnet != nil {
		var ksonnetList = []interface{}{}
		var ksonnet = map[string]interface{}{}
		ksonnet["environment"] = appSpec.Source.Ksonnet.Environment
		var ksonnetParamList = []interface{}{}
		for _, v := range appSpec.Source.Ksonnet.Parameters {
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
	if appSpec.Source.Directory != nil {
		var directoryList = []interface{}{}
		var directory = map[string]interface{}{}
		directory["recurse"] = appSpec.Source.Directory.Recurse
		directory["exclude"] = appSpec.Source.Directory.Exclude
		directory["include"] = appSpec.Source.Directory.Include
		if appSpec.Source.Directory.Jsonnet != nil {
			var jsonnetList = []interface{}{}
			var jsonnet = map[string]interface{}{}
			jsonnet["libs"] = appSpec.Source.Directory.Jsonnet.Libs
			if appSpec.Source.Directory.Jsonnet.ExtVars != nil {
				var jsonnetExtVarsList = []interface{}{}
				for _, v := range appSpec.Source.Directory.Jsonnet.ExtVars {
					var jsonnetExtVars = map[string]interface{}{}
					jsonnetExtVars["name"] = v.Name
					jsonnetExtVars["value"] = v.Value
					jsonnetExtVars["code"] = v.Code
					jsonnetExtVarsList = append(jsonnetExtVarsList, jsonnetExtVars)
				}
				jsonnet["ext_vars"] = jsonnetExtVarsList
			}
			if appSpec.Source.Directory.Jsonnet.Tlas != nil {
				var jsonnetTlasList = []interface{}{}
				for _, v := range appSpec.Source.Directory.Jsonnet.Tlas {
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
	if appSpec.Source.Plugin != nil {
		var pluginList = []interface{}{}
		var plugin = map[string]interface{}{}
		plugin["name"] = appSpec.Source.Plugin.Name
		var pluginEnvList = []interface{}{}
		for _, v := range appSpec.Source.Plugin.Env {
			var pluginEnv = map[string]interface{}{}
			pluginEnv["name"] = v.Name
			pluginEnv["value"] = v.Value
			pluginList = append(pluginEnvList, pluginEnv)
		}
		plugin["env"] = pluginEnvList
		pluginList = append(pluginList, plugin)
		source["plugin"] = pluginList
	}
	return source
}

func setSpecSourceForRequest(source map[string]interface{}) *nextgen.ApplicationsApplicationSource {
	var specSource nextgen.ApplicationsApplicationSource
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
	if source["ref"] != nil && len(source["ref"].(string)) > 0 {
		specSource.Ref = source["ref"].(string)
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
		if helm["ignore_missing_value_files"] != nil {
			helmData.IgnoreMissingValueFiles = helm["ignore_missing_value_files"].(bool)
		}
		if helm["skip_crds"] != nil {
			helmData.SkipCrds = helm["skip_crds"].(bool)
		}
		if helm["skip_tests"] != nil {
			helmData.SkipTests = helm["skip_tests"].(bool)
		}
		if helm["skip_schema_validation"] != nil {
			helmData.SkipSchemaValidation = helm["skip_schema_validation"].(bool)
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
		if kustomizeSource["common_labels"] != nil && len(kustomizeSource["common_labels"].(map[string]interface{})) > 0 {
			var kustomizeCommonLabels = map[string]string{}
			for k, v := range kustomizeSource["common_labels"].(map[string]interface{}) {
				kustomizeCommonLabels[k] = v.(string)
			}
			kustomizeData.CommonLabels = kustomizeCommonLabels
		}
		if kustomizeSource["version"] != nil && len(kustomizeSource["version"].(string)) > 0 {
			kustomizeData.Version = kustomizeSource["version"].(string)
		}
		if kustomizeSource["common_annotations"] != nil && len(kustomizeSource["common_annotations"].(map[string]interface{})) > 0 {
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
	return &specSource
}
