package fault_template

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceFaultTemplateSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		// Identifiers
		"account_id": {
			Description: "Account identifier",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"org_id": {
			Description: "Organization identifier",
			Type:        schema.TypeString,
			Optional:    true,
			ForceNew:    true,
		},
		"project_id": {
			Description: "Project identifier",
			Type:        schema.TypeString,
			Optional:    true,
			ForceNew:    true,
		},
		"hub_identity": {
			Description: "Hub identity reference",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"hub_ref": {
			Description: "Hub reference (computed)",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"identity": {
			Description: "Unique identifier for the fault template (immutable)",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"name": {
			Description: "Name of the fault template",
			Type:        schema.TypeString,
			Required:    true,
		},
		"description": {
			Description: "Description of the fault template",
			Type:        schema.TypeString,
			Optional:    true,
		},

		// Metadata
		"tags": {
			Description: "Tags for the fault template",
			Type:        schema.TypeList,
			Optional:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"category": {
			Description: "Fault categories",
			Type:        schema.TypeList,
			Optional:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"infrastructures": {
			Description: "List of supported infrastructures",
			Type:        schema.TypeList,
			Optional:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"type": {
			Description: "Fault type",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"permissions_required": {
			Description: "Required permissions for the fault",
			Type:        schema.TypeString,
			Optional:    true,
		},

		// API metadata
		"api_version": {
			Description: "API version",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"kind": {
			Description: "Resource kind",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"keywords": {
			Description: "Search keywords",
			Type:        schema.TypeList,
			Optional:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"platforms": {
			Description: "Supported platforms",
			Type:        schema.TypeList,
			Optional:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},

		// Computed fields
		"revision": {
			Description: "Template revision (defaults to v1 if not specified)",
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
		},
		"is_enterprise": {
			Description: "Whether this is an enterprise-only template",
			Type:        schema.TypeBool,
			Computed:    true,
		},
		"is_removed": {
			Description: "Soft delete flag",
			Type:        schema.TypeBool,
			Computed:    true,
		},
		"created_at": {
			Description: "Creation timestamp",
			Type:        schema.TypeInt,
			Computed:    true,
		},
		"created_by": {
			Description: "Creator user ID",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"updated_at": {
			Description: "Update timestamp",
			Type:        schema.TypeInt,
			Computed:    true,
		},
		"updated_by": {
			Description: "Updater user ID",
			Type:        schema.TypeString,
			Computed:    true,
		},

		// Links
		"links": {
			Description: "Related links",
			Type:        schema.TypeList,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"name": {
						Description: "Link name",
						Type:        schema.TypeString,
						Required:    true,
					},
					"url": {
						Description: "Link URL",
						Type:        schema.TypeString,
						Required:    true,
					},
				},
			},
		},

		// Variables
		"variables": {
			Description: "Template variables",
			Type:        schema.TypeList,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"name": {
						Description: "Variable name",
						Type:        schema.TypeString,
						Required:    true,
					},
					"value": {
						Description: "Variable value",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"type": {
						Description:  "Variable type",
						Type:         schema.TypeString,
						Optional:     true,
						ValidateFunc: validation.StringInSlice([]string{"string", "number", "boolean", "secret"}, false),
					},
					"required": {
						Description: "Whether the variable is required",
						Type:        schema.TypeBool,
						Optional:    true,
						Default:     false,
					},
					"description": {
						Description: "Variable description",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},

		// Fault Specification
		"spec": {
			Description: "Fault specification",
			Type:        schema.TypeList,
			Optional:    true,
			MaxItems:    1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					// spec.chaos
					"chaos": {
						Description: "Chaos configuration",
						Type:        schema.TypeList,
						Optional:    true,
						MaxItems:    1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"fault_name": {
									Description: "Name of the fault",
									Type:        schema.TypeString,
									Optional:    true,
								},

								// Chaos parameters
								"params": {
									Description: "Fault parameters",
									Type:        schema.TypeList,
									Optional:    true,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"name": {
												Description: "Parameter name",
												Type:        schema.TypeString,
												Required:    true,
											},
											"value": {
												Description: "Parameter value",
												Type:        schema.TypeString,
												Required:    true,
											},
										},
									},
								},

								// Status check timeouts
								"status_check_timeouts": {
									Description: "Status check timeout configuration",
									Type:        schema.TypeList,
									Optional:    true,
									MaxItems:    1,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"delay": {
												Description: "Delay before status check (seconds)",
												Type:        schema.TypeInt,
												Optional:    true,
											},
											"timeout": {
												Description: "Timeout for status check (seconds)",
												Type:        schema.TypeInt,
												Optional:    true,
											},
										},
									},
								},

								// TLS configuration
								"tls": {
									Description: "TLS configuration",
									Type:        schema.TypeList,
									Optional:    true,
									MaxItems:    1,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"ca_certificate": {
												Description: "CA certificate",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"client_certificate": {
												Description: "Client certificate",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"client_key": {
												Description: "Client key",
												Type:        schema.TypeString,
												Optional:    true,
												Sensitive:   true,
											},
										},
									},
								},

								// Authentication
								"auth": {
									Description: "Authentication configuration",
									Type:        schema.TypeList,
									Optional:    true,
									MaxItems:    1,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											// AWS Auth
											"aws": {
												Description: "AWS authentication",
												Type:        schema.TypeList,
												Optional:    true,
												MaxItems:    1,
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														"access_key_id": {
															Description: "AWS access key ID",
															Type:        schema.TypeString,
															Required:    true,
														},
														"secret_access_key": {
															Description: "AWS secret access key",
															Type:        schema.TypeString,
															Required:    true,
															Sensitive:   true,
														},
														"region": {
															Description: "AWS region",
															Type:        schema.TypeString,
															Required:    true,
														},
													},
												},
											},

											// Azure Auth
											"azure": {
												Description: "Azure authentication",
												Type:        schema.TypeList,
												Optional:    true,
												MaxItems:    1,
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														"client_id": {
															Description: "Azure client ID",
															Type:        schema.TypeString,
															Required:    true,
														},
														"client_secret": {
															Description: "Azure client secret",
															Type:        schema.TypeString,
															Required:    true,
															Sensitive:   true,
														},
														"tenant_id": {
															Description: "Azure tenant ID",
															Type:        schema.TypeString,
															Required:    true,
														},
														"subscription_id": {
															Description: "Azure subscription ID",
															Type:        schema.TypeString,
															Required:    true,
														},
													},
												},
											},

											// GCP Auth
											"gcp": {
												Description: "GCP authentication",
												Type:        schema.TypeList,
												Optional:    true,
												MaxItems:    1,
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														"service_account_key": {
															Description: "GCP service account key (JSON)",
															Type:        schema.TypeString,
															Required:    true,
															Sensitive:   true,
														},
														"project_id": {
															Description: "GCP project ID",
															Type:        schema.TypeString,
															Required:    true,
														},
													},
												},
											},

											// Redis Auth
											"redis": {
												Description: "Redis authentication",
												Type:        schema.TypeList,
												Optional:    true,
												MaxItems:    1,
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														"username": {
															Description: "Redis username",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"password": {
															Description: "Redis password",
															Type:        schema.TypeString,
															Required:    true,
															Sensitive:   true,
														},
													},
												},
											},

											// SSH Auth
											"ssh": {
												Description: "SSH authentication",
												Type:        schema.TypeList,
												Optional:    true,
												MaxItems:    1,
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														"username": {
															Description: "SSH username",
															Type:        schema.TypeString,
															Required:    true,
														},
														"password": {
															Description: "SSH password",
															Type:        schema.TypeString,
															Optional:    true,
															Sensitive:   true,
														},
														"private_key": {
															Description: "SSH private key",
															Type:        schema.TypeString,
															Optional:    true,
															Sensitive:   true,
														},
													},
												},
											},

											// VMware Auth
											"vmware": {
												Description: "VMware authentication",
												Type:        schema.TypeList,
												Optional:    true,
												MaxItems:    1,
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														"username": {
															Description: "VMware username",
															Type:        schema.TypeString,
															Required:    true,
														},
														"password": {
															Description: "VMware password",
															Type:        schema.TypeString,
															Required:    true,
															Sensitive:   true,
														},
														"vcenter_server": {
															Description: "vCenter server address",
															Type:        schema.TypeString,
															Required:    true,
														},
													},
												},
											},
										},
									},
								},

								// Kubernetes-specific configuration
								"kubernetes": {
									Description: "Kubernetes-specific chaos configuration",
									Type:        schema.TypeList,
									Optional:    true,
									MaxItems:    1,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"image": {
												Description: "Container image for chaos experiment",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"image_pull_policy": {
												Description:  "Image pull policy",
												Type:         schema.TypeString,
												Optional:     true,
												ValidateFunc: validation.StringInSlice([]string{"Always", "IfNotPresent", "Never"}, false),
											},
											"image_pull_secrets": {
												Description: "Image pull secrets",
												Type:        schema.TypeList,
												Optional:    true,
												Elem: &schema.Schema{
													Type: schema.TypeString,
												},
											},
											"host_pid": {
												Description: "Use host PID namespace",
												Type:        schema.TypeBool,
												Optional:    true,
												Default:     false,
											},

											// Resources
											"resources": {
												Description: "Resource requirements",
												Type:        schema.TypeList,
												Optional:    true,
												MaxItems:    1,
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														"limits": {
															Description: "Resource limits",
															Type:        schema.TypeMap,
															Optional:    true,
															Elem: &schema.Schema{
																Type: schema.TypeString,
															},
														},
														"requests": {
															Description: "Resource requests",
															Type:        schema.TypeMap,
															Optional:    true,
															Elem: &schema.Schema{
																Type: schema.TypeString,
															},
														},
													},
												},
											},

											// ConfigMaps
											"config_maps": {
												Description: "ConfigMap volumes",
												Type:        schema.TypeList,
												Optional:    true,
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														"name": {
															Description: "ConfigMap name",
															Type:        schema.TypeString,
															Required:    true,
														},
														"mount_path": {
															Description: "Mount path",
															Type:        schema.TypeString,
															Required:    true,
														},
														"mount_mode": {
															Description: "Mount mode (0-3)",
															Type:        schema.TypeInt,
															Optional:    true,
														},
													},
												},
											},

											// Secrets
											"secrets": {
												Description: "Secret volumes",
												Type:        schema.TypeList,
												Optional:    true,
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														"secret_name": {
															Description: "Secret name",
															Type:        schema.TypeString,
															Required:    true,
														},
														"mount_path": {
															Description: "Mount path",
															Type:        schema.TypeString,
															Required:    true,
														},
														"mount_mode": {
															Description: "Mount mode (0-3)",
															Type:        schema.TypeInt,
															Optional:    true,
														},
													},
												},
											},

											// Host path volumes
											"host_file_volumes": {
												Description: "Host path volumes",
												Type:        schema.TypeList,
												Optional:    true,
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														"name": {
															Description: "Volume name",
															Type:        schema.TypeString,
															Required:    true,
														},
														"mount_path": {
															Description: "Mount path",
															Type:        schema.TypeString,
															Required:    true,
														},
														"host_path": {
															Description: "Host path on the node",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"type": {
															Description: "Host path type (e.g., Directory, File, BlockDevice, CharDevice)",
															Type:        schema.TypeString,
															Optional:    true,
														},
													},
												},
											},

											// Command and args
											"command": {
												Description: "Container command",
												Type:        schema.TypeList,
												Optional:    true,
												Elem: &schema.Schema{
													Type: schema.TypeString,
												},
											},
											"args": {
												Description: "Container arguments",
												Type:        schema.TypeList,
												Optional:    true,
												Elem: &schema.Schema{
													Type: schema.TypeString,
												},
											},

											// Host namespace settings
											"host_network": {
												Description: "Use host network namespace",
												Type:        schema.TypeBool,
												Optional:    true,
												Default:     false,
											},
											"host_ipc": {
												Description: "Use host IPC namespace",
												Type:        schema.TypeBool,
												Optional:    true,
												Default:     false,
											},

											// Labels, annotations, node selector
											"labels": {
												Description: "Pod labels",
												Type:        schema.TypeMap,
												Optional:    true,
												Elem: &schema.Schema{
													Type: schema.TypeString,
												},
											},
											"annotations": {
												Description: "Pod annotations",
												Type:        schema.TypeMap,
												Optional:    true,
												Elem: &schema.Schema{
													Type: schema.TypeString,
												},
											},
											"node_selector": {
												Description: "Node selector for pod scheduling",
												Type:        schema.TypeMap,
												Optional:    true,
												Elem: &schema.Schema{
													Type: schema.TypeString,
												},
											},

											// Environment variables
											"env": {
												Description: "Environment variables",
												Type:        schema.TypeList,
												Optional:    true,
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														"name": {
															Description: "Environment variable name",
															Type:        schema.TypeString,
															Required:    true,
														},
														"value": {
															Description: "Environment variable value",
															Type:        schema.TypeString,
															Optional:    true,
														},
													},
												},
											},

											// Tolerations
											"tolerations": {
												Description: "Pod tolerations",
												Type:        schema.TypeList,
												Optional:    true,
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														"key": {
															Description: "Toleration key",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"operator": {
															Description: "Toleration operator (Equal, Exists)",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"value": {
															Description: "Toleration value",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"effect": {
															Description: "Toleration effect (NoSchedule, PreferNoSchedule, NoExecute)",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"toleration_seconds": {
															Description: "Toleration seconds",
															Type:        schema.TypeInt,
															Optional:    true,
														},
													},
												},
											},

											// Pod security context
											"pod_security_context": {
												Description: "Pod security context",
												Type:        schema.TypeList,
												Optional:    true,
												MaxItems:    1,
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														"run_as_user": {
															Description: "User ID to run as",
															Type:        schema.TypeInt,
															Optional:    true,
														},
														"run_as_group": {
															Description: "Group ID to run as",
															Type:        schema.TypeInt,
															Optional:    true,
														},
														"fs_group": {
															Description: "Filesystem group ID",
															Type:        schema.TypeInt,
															Optional:    true,
														},
														"run_as_non_root": {
															Description: "Run as non-root user",
															Type:        schema.TypeBool,
															Optional:    true,
														},
													},
												},
											},

											// Container security context
											"container_security_context": {
												Description: "Container security context",
												Type:        schema.TypeList,
												Optional:    true,
												MaxItems:    1,
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														"privileged": {
															Description: "Run container in privileged mode",
															Type:        schema.TypeBool,
															Optional:    true,
														},
														"read_only_root_filesystem": {
															Description: "Mount root filesystem as read-only",
															Type:        schema.TypeBool,
															Optional:    true,
														},
														"allow_privilege_escalation": {
															Description: "Allow privilege escalation",
															Type:        schema.TypeBool,
															Optional:    true,
														},
														"run_as_user": {
															Description: "User ID to run as",
															Type:        schema.TypeInt,
															Optional:    true,
														},
														"run_as_group": {
															Description: "Group ID to run as",
															Type:        schema.TypeInt,
															Optional:    true,
														},
														"run_as_non_root": {
															Description: "Run as non-root user",
															Type:        schema.TypeBool,
															Optional:    true,
														},
														"capabilities": {
															Description: "Linux capabilities",
															Type:        schema.TypeList,
															Optional:    true,
															MaxItems:    1,
															Elem: &schema.Resource{
																Schema: map[string]*schema.Schema{
																	"add": {
																		Description: "Capabilities to add",
																		Type:        schema.TypeList,
																		Optional:    true,
																		Elem: &schema.Schema{
																			Type: schema.TypeString,
																		},
																	},
																	"drop": {
																		Description: "Capabilities to drop",
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
										},
									},
								},
							},
						},
					},

					// spec.target
					"target": {
						Description: "Target configuration",
						Type:        schema.TypeList,
						Optional:    true,
						MaxItems:    1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								// Kubernetes targets
								"kubernetes": {
									Description: "Kubernetes target configuration",
									Type:        schema.TypeList,
									Optional:    true,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"kind": {
												Description: "Resource kind (e.g., deployment, pod)",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"namespace": {
												Description: "Target namespace",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"labels": {
												Description: "Label selectors",
												Type:        schema.TypeMap,
												Optional:    true,
												Elem: &schema.Schema{
													Type: schema.TypeString,
												},
											},
											"annotations": {
												Description: "Annotation selectors",
												Type:        schema.TypeMap,
												Optional:    true,
												Elem: &schema.Schema{
													Type: schema.TypeString,
												},
											},
											"names": {
												Description: "Specific resource names",
												Type:        schema.TypeList,
												Optional:    true,
												Elem: &schema.Schema{
													Type: schema.TypeString,
												},
											},
											"annotation_check": {
												Description: "Annotation check expression",
												Type:        schema.TypeString,
												Optional:    true,
											},
										},
									},
								},

								// Application target
								"application": {
									Description: "Application target configuration",
									Type:        schema.TypeList,
									Optional:    true,
									MaxItems:    1,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"app_ns": {
												Description: "Application namespace",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"app_kind": {
												Description: "Application kind",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"app_label": {
												Description: "Application label",
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
	}
}
