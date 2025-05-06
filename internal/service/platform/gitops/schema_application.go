package gitops

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

// ArgoAppSpecSchemaV1 is what is our current terraform application resource's schema, we have many required fields as optional, etc
// which can be handled at provider level itself with validation. ArgoAppSpecSchemaV2 is the schema from ArgoCD with all validations in place.
func ArgoAppSpecSchemaV1() *schema.Schema {
	return &schema.Schema{
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
	}
}

// ArgoAppSpecSchemaV2 is a newer app spec schema. This has proper optional/required checks at provider level, we should eventually move to this
func ArgoAppSpecSchemaV2(allOptional bool) *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		MinItems:    1,
		MaxItems:    1,
		Description: "The application specification.",
		Optional:    allOptional,
		Required:    !allOptional,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"destination": {
					Type:        schema.TypeList,
					Description: "Reference to the Kubernetes server and namespace in which the application will be deployed.",
					Optional:    allOptional,
					Required:    !allOptional,
					MinItems:    1,
					MaxItems:    1,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"server": {
								Type:        schema.TypeString,
								Description: "URL of the target cluster and must be set to the Kubernetes control plane API.",
								Optional:    true,
							},
							"namespace": {
								Type:        schema.TypeString,
								Description: "Target namespace for the application's resources. The namespace will only be set for namespace-scoped resources that have not set a value for .metadata.namespace.",
								Optional:    true,
							},
							"name": {
								Type:        schema.TypeString,
								Description: "Name of the target cluster. Can be used instead of `server`.",
								Optional:    true,
							},
						},
					},
				},
				"sources": {
					Type:        schema.TypeList,
					Description: "Location of the application's manifests or chart. Use when specifying multiple fields",
					Optional:    allOptional,
					Required:    !allOptional,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"repo_url": {
								Type:        schema.TypeString,
								Description: "URL to the repository (Git or Helm) that contains the application manifests.",
								Optional:    allOptional,
								Required:    !allOptional,
							},
							"path": {
								Type:        schema.TypeString,
								Description: "Directory path within the repository. Only valid for applications sourced from Git.",
								Optional:    true,
							},
							"target_revision": {
								Type:        schema.TypeString,
								Description: "Revision of the source to sync the application to. In case of Git, this can be commit, tag, or branch. If omitted, will equal to HEAD. In case of Helm, this is a semver tag for the Chart's version.",
								Optional:    true,
							},
							"ref": {
								Type:        schema.TypeString,
								Description: "Reference to another `source` within defined sources. See associated documentation on [Helm value files from external Git repository](https://argo-cd.readthedocs.io/en/stable/user-guide/multiple_sources/#helm-value-files-from-external-git-repository) regarding combining `ref` with `path` and/or `chart`.",
								Optional:    true,
							},
							"chart": {
								Type:        schema.TypeString,
								Description: "Helm chart name. Must be specified for applications sourced from a Helm repo.",
								Optional:    true,
							},
							"helm": {
								Type:        schema.TypeList,
								Description: "Helm specific options.",
								MaxItems:    1,
								MinItems:    1,
								Optional:    true,
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"value_files": {
											Type:        schema.TypeList,
											Description: "List of Helm value files to use when generating a template.",
											Optional:    true,
											Elem: &schema.Schema{
												Type: schema.TypeString,
											},
										},
										"values": {
											Type:        schema.TypeString,
											Description: "Helm values to be passed to 'helm template', typically defined as a block.",
											Optional:    true,
										},
										"ignore_missing_value_files": {
											Type:        schema.TypeBool,
											Description: "Prevents 'helm template' from failing when `value_files` do not exist locally by not appending them to 'helm template --values'.",
											Optional:    true,
										},
										"parameters": {
											Type:        schema.TypeList,
											Description: "Helm parameters which are passed to the helm template command upon manifest generation.",
											Optional:    true,
											Elem: &schema.Resource{
												Schema: map[string]*schema.Schema{
													"name": {
														Type:        schema.TypeString,
														Description: "Name of the Helm parameter.",
														Optional:    true,
													},
													"value": {
														Type:        schema.TypeString,
														Description: "Value of the Helm parameter.",
														Optional:    true,
													},
													"force_string": {
														Type:        schema.TypeBool,
														Optional:    true,
														Description: "Determines whether to tell Helm to interpret booleans and numbers as strings.",
													},
												},
											},
										},
										"file_parameters": {
											Type:        schema.TypeList,
											Description: "File parameters for the helm template.",
											Optional:    true,
											Elem: &schema.Resource{
												Schema: map[string]*schema.Schema{
													"name": {
														Type:        schema.TypeString,
														Description: "Name of the Helm parameter.",
														Required:    true,
													},
													"path": {
														Type:        schema.TypeString,
														Description: "Path to the file containing the values for the Helm parameter.",
														Required:    true,
													},
												},
											},
										},
										"release_name": {
											Type:        schema.TypeString,
											Description: "Helm release name. If omitted it will use the application name.",
											Optional:    true,
										},
										"skip_crds": {
											Type:        schema.TypeBool,
											Description: "Whether to skip custom resource definition installation step (Helm's [--skip-crds](https://helm.sh/docs/chart_best_practices/custom_resource_definitions/)).",
											Optional:    true,
										},
										"pass_credentials": {
											Type:        schema.TypeBool,
											Description: "If true then adds '--pass-credentials' to Helm commands to pass credentials to all domains.",
											Optional:    true,
										},
									},
								},
							},
							"kustomize": {
								Type:        schema.TypeList,
								Description: "Kustomize specific options.",
								MaxItems:    1,
								MinItems:    1,
								Optional:    true,
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"name_prefix": {
											Type:        schema.TypeString,
											Description: "Prefix appended to resources for Kustomize apps.",
											Optional:    true,
										},
										"name_suffix": {
											Type:        schema.TypeString,
											Description: "Suffix appended to resources for Kustomize apps.",
											Optional:    true,
										},
										"version": {
											Type:        schema.TypeString,
											Description: "Version of Kustomize to use for rendering manifests.",
											Optional:    true,
										},
										"images": {
											Type:        schema.TypeList,
											Description: "List of Kustomize image override specifications.",
											Optional:    true,
											Elem: &schema.Schema{
												Type: schema.TypeString,
											},
										},
										"common_labels": {
											Type:        schema.TypeMap,
											Description: "List of additional labels to add to rendered manifests.",
											Optional:    true,
											Elem:        &schema.Schema{Type: schema.TypeString},
										},
										"common_annotations": {
											Type:        schema.TypeMap,
											Description: "List of additional annotations to add to rendered manifests.",
											Optional:    true,
											Elem:        &schema.Schema{Type: schema.TypeString},
										},
									},
								},
							},
							"directory": {
								Type:        schema.TypeList,
								Description: "Path/directory specific options.",
								MaxItems:    1,
								MinItems:    1,
								Optional:    true,
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"recurse": {
											Type:        schema.TypeBool,
											Description: "Whether to scan a directory recursively for manifests.",
											Optional:    true,
										},
										"jsonnet": {
											Type:        schema.TypeList,
											Description: "Jsonnet specific options.",
											Optional:    true,
											MaxItems:    1,
											MinItems:    1,
											Elem: &schema.Resource{
												Schema: map[string]*schema.Schema{
													"ext_var": {
														Type:        schema.TypeList,
														Description: "List of Jsonnet External Variables.",
														Optional:    true,
														Elem: &schema.Resource{
															Schema: map[string]*schema.Schema{
																"name": {
																	Type:        schema.TypeString,
																	Description: "Name of Jsonnet variable.",
																	Optional:    true,
																},
																"value": {
																	Type:        schema.TypeString,
																	Description: "Value of Jsonnet variable.",
																	Optional:    true,
																},
																"code": {
																	Type:        schema.TypeBool,
																	Description: "Determines whether the variable should be evaluated as jsonnet code or treated as string.",
																	Optional:    true,
																},
															},
														},
													},
													"tlas": {
														Type:        schema.TypeList,
														Description: "List of Jsonnet Top-level Arguments",
														Optional:    true,
														Elem: &schema.Resource{
															Schema: map[string]*schema.Schema{
																"name": {
																	Type:        schema.TypeString,
																	Description: "Name of Jsonnet variable.",
																	Optional:    true,
																},
																"value": {
																	Type:        schema.TypeString,
																	Description: "Value of Jsonnet variable.",
																	Optional:    true,
																},
																"code": {
																	Type:        schema.TypeBool,
																	Description: "Determines whether the variable should be evaluated as jsonnet code or treated as string.",
																	Optional:    true,
																},
															},
														},
													},
													"libs": {
														Type:        schema.TypeList,
														Description: "Additional library search dirs.",
														Optional:    true,
														Elem: &schema.Schema{
															Type: schema.TypeString,
														},
													},
												},
											},
										},
										"exclude": {
											Type:        schema.TypeString,
											Description: "Glob pattern to match paths against that should be explicitly excluded from being used during manifest generation. This takes precedence over the `include` field. To match multiple patterns, wrap the patterns in {} and separate them with commas. For example: '{config.yaml,env-use2/*}'",
											Optional:    true,
										},
										"include": {
											Type:        schema.TypeString,
											Description: "Glob pattern to match paths against that should be explicitly included during manifest generation. If this field is set, only matching manifests will be included. To match multiple patterns, wrap the patterns in {} and separate them with commas. For example: '{*.yml,*.yaml}'",
											Optional:    true,
										},
									},
								},
							},
							"plugin": {
								Type:        schema.TypeList,
								Description: "Config management plugin specific options.",
								MaxItems:    1,
								MinItems:    1,
								Optional:    true,
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"name": {
											Type:        schema.TypeString,
											Description: "Name of the plugin. Only set the plugin name if the plugin is defined in `argocd-cm`. If the plugin is defined as a sidecar, omit the name. The plugin will be automatically matched with the Application according to the plugin's discovery rules.",
											Optional:    true,
										},
										"env": {
											Type:        schema.TypeList,
											Description: "Environment variables passed to the plugin.",
											Optional:    true,
											Elem: &schema.Resource{
												Schema: map[string]*schema.Schema{
													"name": {
														Type:        schema.TypeString,
														Description: "Name of the environment variable.",
														Optional:    true,
													},
													"value": {
														Type:        schema.TypeString,
														Description: "Value of the environment variable.",
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
				"source": {
					Type:        schema.TypeList,
					Description: "Location of the application's manifests or chart.",
					Optional:    allOptional,
					Required:    !allOptional,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"repo_url": {
								Type:        schema.TypeString,
								Description: "URL to the repository (Git or Helm) that contains the application manifests.",
								Optional:    allOptional,
								Required:    !allOptional,
							},
							"path": {
								Type:        schema.TypeString,
								Description: "Directory path within the repository. Only valid for applications sourced from Git.",
								Optional:    true,
							},
							"target_revision": {
								Type:        schema.TypeString,
								Description: "Revision of the source to sync the application to. In case of Git, this can be commit, tag, or branch. If omitted, will equal to HEAD. In case of Helm, this is a semver tag for the Chart's version.",
								Optional:    true,
							},
							"ref": {
								Type:        schema.TypeString,
								Description: "Reference to another `source` within defined sources. See associated documentation on [Helm value files from external Git repository](https://argo-cd.readthedocs.io/en/stable/user-guide/multiple_sources/#helm-value-files-from-external-git-repository) regarding combining `ref` with `path` and/or `chart`.",
								Optional:    true,
							},
							"chart": {
								Type:        schema.TypeString,
								Description: "Helm chart name. Must be specified for applications sourced from a Helm repo.",
								Optional:    true,
							},
							"helm": {
								Type:        schema.TypeList,
								Description: "Helm specific options.",
								MaxItems:    1,
								MinItems:    1,
								Optional:    true,
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"value_files": {
											Type:        schema.TypeList,
											Description: "List of Helm value files to use when generating a template.",
											Optional:    true,
											Elem: &schema.Schema{
												Type: schema.TypeString,
											},
										},
										"values": {
											Type:        schema.TypeString,
											Description: "Helm values to be passed to 'helm template', typically defined as a block.",
											Optional:    true,
										},
										"ignore_missing_value_files": {
											Type:        schema.TypeBool,
											Description: "Prevents 'helm template' from failing when `value_files` do not exist locally by not appending them to 'helm template --values'.",
											Optional:    true,
										},
										"parameters": {
											Type:        schema.TypeList,
											Description: "Helm parameters which are passed to the helm template command upon manifest generation.",
											Optional:    true,
											Elem: &schema.Resource{
												Schema: map[string]*schema.Schema{
													"name": {
														Type:        schema.TypeString,
														Description: "Name of the Helm parameter.",
														Optional:    true,
													},
													"value": {
														Type:        schema.TypeString,
														Description: "Value of the Helm parameter.",
														Optional:    true,
													},
													"force_string": {
														Type:        schema.TypeBool,
														Optional:    true,
														Description: "Determines whether to tell Helm to interpret booleans and numbers as strings.",
													},
												},
											},
										},
										"file_parameters": {
											Type:        schema.TypeList,
											Description: "File parameters for the helm template.",
											Optional:    true,
											Elem: &schema.Resource{
												Schema: map[string]*schema.Schema{
													"name": {
														Type:        schema.TypeString,
														Description: "Name of the Helm parameter.",
														Required:    true,
													},
													"path": {
														Type:        schema.TypeString,
														Description: "Path to the file containing the values for the Helm parameter.",
														Required:    true,
													},
												},
											},
										},
										"release_name": {
											Type:        schema.TypeString,
											Description: "Helm release name. If omitted it will use the application name.",
											Optional:    true,
										},
										"skip_crds": {
											Type:        schema.TypeBool,
											Description: "Whether to skip custom resource definition installation step (Helm's [--skip-crds](https://helm.sh/docs/chart_best_practices/custom_resource_definitions/)).",
											Optional:    true,
										},
										"pass_credentials": {
											Type:        schema.TypeBool,
											Description: "If true then adds '--pass-credentials' to Helm commands to pass credentials to all domains.",
											Optional:    true,
										},
									},
								},
							},
							"kustomize": {
								Type:        schema.TypeList,
								Description: "Kustomize specific options.",
								MaxItems:    1,
								MinItems:    1,
								Optional:    true,
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"name_prefix": {
											Type:        schema.TypeString,
											Description: "Prefix appended to resources for Kustomize apps.",
											Optional:    true,
										},
										"name_suffix": {
											Type:        schema.TypeString,
											Description: "Suffix appended to resources for Kustomize apps.",
											Optional:    true,
										},
										"version": {
											Type:        schema.TypeString,
											Description: "Version of Kustomize to use for rendering manifests.",
											Optional:    true,
										},
										"images": {
											Type:        schema.TypeList,
											Description: "List of Kustomize image override specifications.",
											Optional:    true,
											Elem: &schema.Schema{
												Type: schema.TypeString,
											},
										},
										"common_labels": {
											Type:        schema.TypeMap,
											Description: "List of additional labels to add to rendered manifests.",
											Optional:    true,
											Elem:        &schema.Schema{Type: schema.TypeString},
										},
										"common_annotations": {
											Type:        schema.TypeMap,
											Description: "List of additional annotations to add to rendered manifests.",
											Optional:    true,
											Elem:        &schema.Schema{Type: schema.TypeString},
										},
									},
								},
							},
							"directory": {
								Type:        schema.TypeList,
								Description: "Path/directory specific options.",
								MaxItems:    1,
								MinItems:    1,
								Optional:    true,
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"recurse": {
											Type:        schema.TypeBool,
											Description: "Whether to scan a directory recursively for manifests.",
											Optional:    true,
										},
										"jsonnet": {
											Type:        schema.TypeList,
											Description: "Jsonnet specific options.",
											Optional:    true,
											MaxItems:    1,
											MinItems:    1,
											Elem: &schema.Resource{
												Schema: map[string]*schema.Schema{
													"ext_var": {
														Type:        schema.TypeList,
														Description: "List of Jsonnet External Variables.",
														Optional:    true,
														Elem: &schema.Resource{
															Schema: map[string]*schema.Schema{
																"name": {
																	Type:        schema.TypeString,
																	Description: "Name of Jsonnet variable.",
																	Optional:    true,
																},
																"value": {
																	Type:        schema.TypeString,
																	Description: "Value of Jsonnet variable.",
																	Optional:    true,
																},
																"code": {
																	Type:        schema.TypeBool,
																	Description: "Determines whether the variable should be evaluated as jsonnet code or treated as string.",
																	Optional:    true,
																},
															},
														},
													},
													"tlas": {
														Type:        schema.TypeList,
														Description: "List of Jsonnet Top-level Arguments",
														Optional:    true,
														Elem: &schema.Resource{
															Schema: map[string]*schema.Schema{
																"name": {
																	Type:        schema.TypeString,
																	Description: "Name of Jsonnet variable.",
																	Optional:    true,
																},
																"value": {
																	Type:        schema.TypeString,
																	Description: "Value of Jsonnet variable.",
																	Optional:    true,
																},
																"code": {
																	Type:        schema.TypeBool,
																	Description: "Determines whether the variable should be evaluated as jsonnet code or treated as string.",
																	Optional:    true,
																},
															},
														},
													},
													"libs": {
														Type:        schema.TypeList,
														Description: "Additional library search dirs.",
														Optional:    true,
														Elem: &schema.Schema{
															Type: schema.TypeString,
														},
													},
												},
											},
										},
										"exclude": {
											Type:        schema.TypeString,
											Description: "Glob pattern to match paths against that should be explicitly excluded from being used during manifest generation. This takes precedence over the `include` field. To match multiple patterns, wrap the patterns in {} and separate them with commas. For example: '{config.yaml,env-use2/*}'",
											Optional:    true,
										},
										"include": {
											Type:        schema.TypeString,
											Description: "Glob pattern to match paths against that should be explicitly included during manifest generation. If this field is set, only matching manifests will be included. To match multiple patterns, wrap the patterns in {} and separate them with commas. For example: '{*.yml,*.yaml}'",
											Optional:    true,
										},
									},
								},
							},
							"plugin": {
								Type:        schema.TypeList,
								Description: "Config management plugin specific options.",
								MaxItems:    1,
								MinItems:    1,
								Optional:    true,
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"name": {
											Type:        schema.TypeString,
											Description: "Name of the plugin. Only set the plugin name if the plugin is defined in `argocd-cm`. If the plugin is defined as a sidecar, omit the name. The plugin will be automatically matched with the Application according to the plugin's discovery rules.",
											Optional:    true,
										},
										"env": {
											Type:        schema.TypeList,
											Description: "Environment variables passed to the plugin.",
											Optional:    true,
											Elem: &schema.Resource{
												Schema: map[string]*schema.Schema{
													"name": {
														Type:        schema.TypeString,
														Description: "Name of the environment variable.",
														Optional:    true,
													},
													"value": {
														Type:        schema.TypeString,
														Description: "Value of the environment variable.",
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
				"project": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "The project the application belongs to.",
				},
				"sync_policy": {
					Type:        schema.TypeList,
					Description: "Controls when and how a sync will be performed.",
					Optional:    true,
					MaxItems:    1,
					MinItems:    1,
					//DiffSuppressFunc: func(k, oldValue, newValue string, d *schema.ResourceData) bool {
					//	// Avoid drift when sync_policy is empty
					//	if k == "spec.0.sync_policy.#" {
					//		_, hasAutomated := d.GetOk("spec.0.sync_policy.0.automated")
					//		_, hasSyncOptions := d.GetOk("spec.0.sync_policy.0.sync_options")
					//		_, hasRetry := d.GetOk("spec.0.sync_policy.0.retry")
					//		_, hasManagedNamespaceMetadata := d.GetOk("spec.0.sync_policy.0.managed_namespace_metadata")
					//
					//		if !hasAutomated && !hasSyncOptions && !hasRetry && !hasManagedNamespaceMetadata {
					//			return true
					//		}
					//	}
					//	return false
					//},
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"automated": {
								Type:        schema.TypeList,
								Description: "Whether to automatically keep an application synced to the target revision.",
								MaxItems:    1,
								Optional:    true,
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"prune": {
											Type:        schema.TypeBool,
											Description: "Whether to delete resources from the cluster that are not found in the sources anymore as part of automated sync.",
											Optional:    true,
										},
										"self_heal": {
											Type:        schema.TypeBool,
											Description: "Whether to revert resources back to their desired state upon modification in the cluster.",
											Optional:    true,
										},
										"allow_empty": {
											Type:        schema.TypeBool,
											Description: "Allows apps have zero live resources.",
											Optional:    true,
										},
									},
								},
							},
							"sync_options": {
								Type:        schema.TypeList,
								Description: "List of sync options. More info: https://argo-cd.readthedocs.io/en/stable/user-guide/sync-options/.",
								Optional:    true,
								Elem: &schema.Schema{
									Type: schema.TypeString,
									// TODO: add a validator
								},
							},
							"retry": {
								Type:        schema.TypeList,
								Description: "Controls failed sync retry behavior.",
								MaxItems:    1,
								Optional:    true,
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"limit": {
											Type:        schema.TypeString,
											Description: "Maximum number of attempts for retrying a failed sync. If set to 0, no retries will be performed.",
											Optional:    true,
										},
										"backoff": {
											Type:        schema.TypeList,
											MaxItems:    1,
											Description: "Controls how to backoff on subsequent retries of failed syncs.",
											Optional:    true,
											Elem: &schema.Resource{
												Schema: map[string]*schema.Schema{
													"duration": {
														Type:        schema.TypeString,
														Description: "Duration is the amount to back off. Default unit is seconds, but could also be a duration (e.g. `2m`, `1h`), as a string.",
														Optional:    true,
													},
													"factor": {
														Type:        schema.TypeString,
														Description: "Factor to multiply the base duration after each failed retry.",
														Optional:    true,
													},
													"max_duration": {
														Type:        schema.TypeString,
														Description: "Maximum amount of time allowed for the backoff strategy. Default unit is seconds, but could also be a duration (e.g. `2m`, `1h`), as a string.",
														Optional:    true,
													},
												},
											},
										},
									},
								},
							},
							"managed_namespace_metadata": {
								Type:        schema.TypeList,
								MaxItems:    1,
								Description: "Controls metadata in the given namespace (if `CreateNamespace=true`).",
								Optional:    true,
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"annotations": {
											Type:        schema.TypeMap,
											Description: "Annotations to apply to the namespace.",
											Optional:    true,
											Elem:        &schema.Schema{Type: schema.TypeString},
										},
										"labels": {
											Type:        schema.TypeMap,
											Description: "Labels to apply to the namespace.",
											Optional:    true,
											Elem:        &schema.Schema{Type: schema.TypeString},
										},
									},
								},
							},
						},
					},
				},
				"ignore_difference": {
					Type:        schema.TypeList,
					Description: "Resources and their fields which should be ignored during comparison. More info: https://argo-cd.readthedocs.io/en/stable/user-guide/diffing/#application-level-configuration.",
					Optional:    true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"group": {
								Type:        schema.TypeString,
								Description: "The Kubernetes resource Group to match for.",
								Optional:    true,
							},
							"kind": {
								Type:        schema.TypeString,
								Description: "The Kubernetes resource Kind to match for.",
								Optional:    true,
							},
							"name": {
								Type:        schema.TypeString,
								Description: "The Kubernetes resource Name to match for.",
								Optional:    true,
							},
							"namespace": {
								Type:        schema.TypeString,
								Description: "The Kubernetes resource Namespace to match for.",
								Optional:    true,
							},
							"json_pointers": {
								Type:        schema.TypeList,
								Description: "List of JSONPaths strings targeting the field(s) to ignore.",
								Set:         schema.HashString,
								Optional:    true,
								Elem: &schema.Schema{
									Type: schema.TypeString,
								},
							},
							"jq_path_expressions": {
								Type:        schema.TypeList,
								Description: "List of JQ path expression strings targeting the field(s) to ignore.",
								Set:         schema.HashString,
								Optional:    true,
								Elem: &schema.Schema{
									Type: schema.TypeString,
								},
							},
						},
					},
				},
				"info": {
					Type:        schema.TypeList,
					Description: "List of information (URLs, email addresses, and plain text) that relates to the application.",
					Optional:    true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"name": {
								Type:        schema.TypeString,
								Description: "Name of the information.",
								Optional:    true,
							},
							"value": {
								Type:        schema.TypeString,
								Description: "Value of the information.",
								Optional:    true,
							},
						},
					},
				},
				"revision_history_limit": {
					Type:        schema.TypeInt,
					Description: "Limits the number of items kept in the application's revision history, which is used for informational purposes as well as for rollbacks to previous versions. This should only be changed in exceptional circumstances. Setting to zero will store no history. This will reduce storage used. Increasing will increase the space used to store the history, so we do not recommend increasing it. Default is 10.",
					Optional:    true,
					Default:     10,
				},
			},
		},
	}
}
