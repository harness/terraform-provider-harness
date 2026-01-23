package action_template

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceActionTemplateSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		// Required Fields
		"identity": {
			Description:  "Unique identifier for the action template (immutable).",
			Type:         schema.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},
		"name": {
			Description:  "Name of the action template.",
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},
		"hub_identity": {
			Description:  "Identity of the chaos hub this action template belongs to.",
			Type:         schema.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},
		"type": {
			Description: "Type of the action template. Valid values: delay, customScript, container.",
			Type:        schema.TypeString,
			Required:    true,
		},

		// Optional Fields
		"org_id": {
			Description:  "Organization identifier.",
			Type:         schema.TypeString,
			Optional:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},
		"project_id": {
			Description:  "Project identifier.",
			Type:         schema.TypeString,
			Optional:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},
		"description": {
			Description: "Description of the action template.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tags": {
			Description: "Tags to associate with the action template.",
			Type:        schema.TypeList,
			Optional:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"infrastructure_type": {
			Description: "Infrastructure type for the action template. Valid values: Kubernetes, KubernetesV2, Windows, Linux, CloudFoundry, Container. Supports runtime inputs like <+input>.",
			Type:        schema.TypeString,
			Optional:    true,
		},

		// Action Properties - Conditional based on type
		"delay_action": {
			Description: "Delay action configuration. Required when type is 'delay'.",
			Type:        schema.TypeList,
			Optional:    true,
			MaxItems:    1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"duration": {
						Description:  "Duration of the delay (e.g., '30s', '5m', '1h').",
						Type:         schema.TypeString,
						Required:     true,
						ValidateFunc: validation.StringIsNotEmpty,
					},
				},
			},
		},
		"custom_script_action": {
			Description: "Custom script action configuration. Required when type is 'customScript'.",
			Type:        schema.TypeList,
			Optional:    true,
			MaxItems:    1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"command": {
						Description:  "Command to execute (e.g., 'bash', 'python', 'sh').",
						Type:         schema.TypeString,
						Required:     true,
						ValidateFunc: validation.StringIsNotEmpty,
					},
					"args": {
						Description: "Arguments to pass to the command.",
						Type:        schema.TypeList,
						Optional:    true,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
					},
					"env": {
						Description: "Environment variables for the script.",
						Type:        schema.TypeList,
						Optional:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"name": {
									Description:  "Environment variable name.",
									Type:         schema.TypeString,
									Required:     true,
									ValidateFunc: validation.StringIsNotEmpty,
								},
								"value": {
									Description: "Environment variable value.",
									Type:        schema.TypeString,
									Required:    true,
								},
							},
						},
					},
				},
			},
		},
		"container_action": {
			Description: "Container action configuration. Required when type is 'container'.",
			Type:        schema.TypeList,
			Optional:    true,
			MaxItems:    1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"image": {
						Description:  "Container image to use (e.g., 'busybox:latest').",
						Type:         schema.TypeString,
						Required:     true,
						ValidateFunc: validation.StringIsNotEmpty,
					},
					"command": {
						Description: "Command to run in the container.",
						Type:        schema.TypeList,
						Optional:    true,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
					},
					"args": {
						Description: "Arguments to pass to the container command.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"namespace": {
						Description: "Kubernetes namespace for the container.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"labels": {
						Description: "Labels to apply to the container pod.",
						Type:        schema.TypeMap,
						Optional:    true,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
					},
					"annotations": {
						Description: "Annotations to apply to the container pod.",
						Type:        schema.TypeMap,
						Optional:    true,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
					},
					"image_pull_policy": {
						Description: "Image pull policy (Always, IfNotPresent, Never). Supports runtime inputs like <+input>.allowedValues(...).",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"image_pull_secrets": {
						Description: "List of image pull secrets for private registries.",
						Type:        schema.TypeList,
						Optional:    true,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
					},
					"env": {
						Description: "Environment variables for the container.",
						Type:        schema.TypeList,
						Optional:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"name": {
									Description:  "Environment variable name.",
									Type:         schema.TypeString,
									Required:     true,
									ValidateFunc: validation.StringIsNotEmpty,
								},
								"value": {
									Description: "Environment variable value.",
									Type:        schema.TypeString,
									Optional:    true,
								},
							},
						},
					},
					"resources": {
						Description: "Resource requirements for the container.",
						Type:        schema.TypeList,
						Optional:    true,
						MaxItems:    1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"limits": {
									Description: "Resource limits.",
									Type:        schema.TypeMap,
									Optional:    true,
									Elem: &schema.Schema{
										Type: schema.TypeString,
									},
								},
								"requests": {
									Description: "Resource requests.",
									Type:        schema.TypeMap,
									Optional:    true,
									Elem: &schema.Schema{
										Type: schema.TypeString,
									},
								},
							},
						},
					},
					"service_account_name": {
						Description: "Kubernetes service account name.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"node_selector": {
						Description: "Node selector for pod scheduling.",
						Type:        schema.TypeMap,
						Optional:    true,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
					},
					"host_network": {
						Description: "Use host network namespace.",
						Type:        schema.TypeBool,
						Optional:    true,
						Default:     false,
					},
					"host_pid": {
						Description: "Use host PID namespace.",
						Type:        schema.TypeBool,
						Optional:    true,
						Default:     false,
					},
					"host_ipc": {
						Description: "Use host IPC namespace.",
						Type:        schema.TypeBool,
						Optional:    true,
						Default:     false,
					},
					"tolerations": {
						Description: "Tolerations for pod scheduling on tainted nodes.",
						Type:        schema.TypeList,
						Optional:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"key": {
									Description: "Taint key to tolerate.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"operator": {
									Description:  "Operator (Exists, Equal).",
									Type:         schema.TypeString,
									Optional:     true,
									Default:      "Equal",
									ValidateFunc: validation.StringInSlice([]string{"Exists", "Equal"}, false),
								},
								"value": {
									Description: "Taint value to tolerate.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"effect": {
									Description:  "Taint effect (NoSchedule, PreferNoSchedule, NoExecute).",
									Type:         schema.TypeString,
									Optional:     true,
									ValidateFunc: validation.StringInSlice([]string{"NoSchedule", "PreferNoSchedule", "NoExecute"}, false),
								},
								"toleration_seconds": {
									Description: "Toleration seconds for NoExecute effect.",
									Type:        schema.TypeInt,
									Optional:    true,
								},
							},
						},
					},
					"volume_mounts": {
						Description: "Volume mounts for the container.",
						Type:        schema.TypeList,
						Optional:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"name": {
									Description:  "Volume name to mount.",
									Type:         schema.TypeString,
									Required:     true,
									ValidateFunc: validation.StringIsNotEmpty,
								},
								"mount_path": {
									Description:  "Path to mount the volume in the container.",
									Type:         schema.TypeString,
									Required:     true,
									ValidateFunc: validation.StringIsNotEmpty,
								},
								"sub_path": {
									Description: "Sub-path within the volume.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"read_only": {
									Description: "Mount as read-only.",
									Type:        schema.TypeBool,
									Optional:    true,
									Default:     false,
								},
							},
						},
					},
					"volumes": {
						Description: "Volumes to attach to the pod.",
						Type:        schema.TypeList,
						Optional:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"name": {
									Description:  "Volume name.",
									Type:         schema.TypeString,
									Required:     true,
									ValidateFunc: validation.StringIsNotEmpty,
								},
								"empty_dir": {
									Description: "EmptyDir volume configuration.",
									Type:        schema.TypeList,
									Optional:    true,
									MaxItems:    1,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"medium": {
												Description: "Storage medium (empty string for default, Memory for tmpfs).",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"size_limit": {
												Description: "Size limit (e.g., '1Gi').",
												Type:        schema.TypeString,
												Optional:    true,
											},
										},
									},
								},
								"config_map": {
									Description: "ConfigMap volume configuration.",
									Type:        schema.TypeList,
									Optional:    true,
									MaxItems:    1,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"name": {
												Description:  "ConfigMap name.",
												Type:         schema.TypeString,
												Required:     true,
												ValidateFunc: validation.StringIsNotEmpty,
											},
											"optional": {
												Description: "Whether the ConfigMap is optional.",
												Type:        schema.TypeBool,
												Optional:    true,
												Default:     false,
											},
										},
									},
								},
								"secret": {
									Description: "Secret volume configuration.",
									Type:        schema.TypeList,
									Optional:    true,
									MaxItems:    1,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"secret_name": {
												Description:  "Secret name.",
												Type:         schema.TypeString,
												Required:     true,
												ValidateFunc: validation.StringIsNotEmpty,
											},
											"optional": {
												Description: "Whether the Secret is optional.",
												Type:        schema.TypeBool,
												Optional:    true,
												Default:     false,
											},
										},
									},
								},
								"host_path": {
									Description: "HostPath volume configuration.",
									Type:        schema.TypeList,
									Optional:    true,
									MaxItems:    1,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"path": {
												Description:  "Host path.",
												Type:         schema.TypeString,
												Required:     true,
												ValidateFunc: validation.StringIsNotEmpty,
											},
											"type": {
												Description: "Host path type (Directory, File, etc.).",
												Type:        schema.TypeString,
												Optional:    true,
											},
										},
									},
								},
								"persistent_volume_claim": {
									Description: "PersistentVolumeClaim configuration.",
									Type:        schema.TypeList,
									Optional:    true,
									MaxItems:    1,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"claim_name": {
												Description:  "PVC name.",
												Type:         schema.TypeString,
												Required:     true,
												ValidateFunc: validation.StringIsNotEmpty,
											},
											"read_only": {
												Description: "Mount as read-only.",
												Type:        schema.TypeBool,
												Optional:    true,
												Default:     false,
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

		// Run Properties
		"run_properties": {
			Description: "Run properties for the action template execution.",
			Type:        schema.TypeList,
			Optional:    true,
			MaxItems:    1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"initial_delay": {
						Description: "Initial delay before action execution (e.g., '5s', '1m').",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"interval": {
						Description: "Interval between retries (e.g., '10s', '30s').",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"max_retries": {
						Description: "Maximum number of retries.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"stop_on_failure": {
						Description: "Whether to stop on failure.",
						Type:        schema.TypeBool,
						Optional:    true,
						Default:     false,
					},
					"timeout": {
						Description: "Timeout for action execution (e.g., '5m', '10m').",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"verbosity": {
						Description: "Verbosity level for logging.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},

		// Variables
		"variables": {
			Description: "Template variables that can be used in the action.",
			Type:        schema.TypeList,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"name": {
						Description:  "Variable name.",
						Type:         schema.TypeString,
						Required:     true,
						ValidateFunc: validation.StringIsNotEmpty,
					},
					"value": {
						Description: "Variable value.",
						Type:        schema.TypeString,
						Required:    true,
					},
					"description": {
						Description: "Variable description.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"type": {
						Description: "Variable type (e.g., 'string', 'number', 'boolean').",
						Type:        schema.TypeString,
						Optional:    true,
						Default:     "string",
					},
					"required": {
						Description: "Whether the variable is required.",
						Type:        schema.TypeBool,
						Optional:    true,
						Default:     false,
					},
				},
			},
		},

		// Computed Fields
		"account_id": {
			Description: "Account identifier.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"id_internal": {
			Description: "Internal ID of the action template.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"revision": {
			Description: "Revision number of the action template.",
			Type:        schema.TypeInt,
			Computed:    true,
		},
		"is_default": {
			Description: "Whether this is the default version for predefined actions.",
			Type:        schema.TypeBool,
			Computed:    true,
		},
		"is_enterprise": {
			Description: "Whether this is an enterprise action template.",
			Type:        schema.TypeBool,
			Computed:    true,
		},
		"is_removed": {
			Description: "Whether the action template has been removed.",
			Type:        schema.TypeBool,
			Computed:    true,
		},
		"created_at": {
			Description: "Creation timestamp (Unix epoch).",
			Type:        schema.TypeInt,
			Computed:    true,
		},
		"updated_at": {
			Description: "Last update timestamp (Unix epoch).",
			Type:        schema.TypeInt,
			Computed:    true,
		},
		"created_by": {
			Description: "User who created the action template.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"updated_by": {
			Description: "User who last updated the action template.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"template": {
			Description: "Template content/definition.",
			Type:        schema.TypeString,
			Computed:    true,
		},
	}
}
