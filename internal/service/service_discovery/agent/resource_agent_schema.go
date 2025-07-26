package agent

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// AgentResourceSchema returns the schema for the agent resource
func AgentResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		// Required parameters
		"name": {
			Description: "The name of the agent. This is a required field.",
			Type:        schema.TypeString,
			Required:    true,
		},
		"environment_identifier": {
			Description: "The environment identifier of the agent. This is a required field.",
			Type:        schema.TypeString,
			Required:    true,
		},
		"infra_identifier": {
			Description: "The infrastructure identifier of the agent. This is a required field.",
			Type:        schema.TypeString,
			Required:    true,
		},
		"config": {
			Type:        schema.TypeList,
			Required:    true,
			Description: "Configuration for the agent. This is a required field.",
			Elem:        resourceConfigSchema(),
		},

		// Optional fields
		"org_identifier": {
			Description: "The organization identifier of the agent. Must be 1-64 characters and contain only alphanumeric characters, hyphens, or underscores.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"project_identifier": {
			Description: "The project identifier of the agent. Must be 1-64 characters and contain only alphanumeric characters, hyphens, or underscores.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"installation_type": {
			Description: "Type of installation for the agent.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"permanent_installation": {
			Description: "Whether this is a permanent installation.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"webhook_url": {
			Description: "Webhook URL for the agent.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"correlation_id": {
			Description: "Correlation ID for the agent.",
			Type:        schema.TypeString,
			Computed:    true,
			Optional:    true,
		},

		// Computed fields
		"id": {
			Description: "The unique identifier of the agent.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"identity": {
			Description: "The unique identity of the agent.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"description": {
			Description: "Description of the agent.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"tags": {
			Description: "List of resource tags for the agent.",
			Type:        schema.TypeList,
			Computed:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"network_map_count": {
			Description: "Number of network maps associated with this agent.",
			Type:        schema.TypeInt,
			Computed:    true,
		},
		"service_count": {
			Description: "Number of services managed by this agent.",
			Type:        schema.TypeInt,
			Computed:    true,
		},
		"created_at": {
			Description: "Timestamp when the agent was created.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"updated_at": {
			Description: "Timestamp when the agent was last updated.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"created_by": {
			Description: "User who created the agent.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"updated_by": {
			Description: "User who last updated the agent.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"removed": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether the agent has been removed.",
		},
		"removed_at": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Timestamp when the agent was removed.",
		},
		"installation_details": {
			Type:        schema.TypeList,
			Computed:    true,
			Description: "Installation details of the agent.",
			Elem:        installationDetailsSchema(),
		},
	}
}

func resourceConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"collector_image": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Docker image for the collector.",
			},
			"log_watcher_image": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Docker image for the log watcher.",
			},
			"skip_secure_verify": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether to skip TLS verification.",
			},
			"image_pull_secrets": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of image pull secrets.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"kubernetes": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "Kubernetes-specific configuration.",
				Elem:        resourceKubernetesSchema(),
			},
			"mtls": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "mTLS configuration.",
				Elem:        resourceMtlsSchema(),
			},
			"proxy": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Proxy configuration.",
				Elem:        resourceProxySchema(),
			},
			"data": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Description: "Data collection configuration.",
				Elem:        resourceDataSchema(),
			},
		},
	}
}

func resourceDataSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enable_node_agent": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether to enable node agent.",
			},
			"node_agent_selector": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Node selector for the node agent.",
			},
			"namespace_selector": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Namespace selector for the agent.",
			},
			"enable_orphaned_pod": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether to enable orphaned pod detection.",
			},
			"enable_batch_resources": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether to enable batch resources.",
			},
			"collection_window_in_min": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     5,
				Description: "Collection window in minutes.",
			},
			"blacklisted_namespaces": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of namespaces to exclude from discovery.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"observed_namespaces": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of namespaces to observe.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"cron": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Cron schedule for data collection.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"expression": {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "0/15 * * * *",
							Description: "Cron expression for scheduling.",
						},
					},
				},
			},
		},
	}
}

func resourceProxySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"http_proxy": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "HTTP proxy URL.",
			},
			"https_proxy": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "HTTPS proxy URL.",
			},
			"no_proxy": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Comma-separated list of hosts that should not use the proxy.",
			},
			"url": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Proxy server URL.",
			},
		},
	}
}

func resourceMtlsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cert_path": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Path to the certificate file.",
			},
			"key_path": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Path to the key file.",
			},
			"secret_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name of the Kubernetes secret containing the certificate and key.",
			},
			"url": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "URL of the mTLS server.",
			},
		},
	}
}

func resourceKubernetesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"disable_namespace_creation": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether to disable namespace creation.",
			},
			"namespaced": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether the agent is namespaced.",
			},
			"namespace": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Kubernetes namespace to use",
			},
			"service_account": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Service account to use",
			},
			"run_as_user": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     2000,
				Description: "The user ID to run as.",
			},
			"run_as_group": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     2000,
				Description: "The group ID to run as.",
			},
			"image_pull_policy": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "IfNotPresent",
				Description: "The image pull policy.",
			},
			"node_selector": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Node selector labels.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"tolerations": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Tolerations for pod assignment.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The taint key that the toleration applies to.",
						},
						"operator": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Operator represents a key's relationship to the value.",
						},
						"value": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The taint value the toleration matches to.",
						},
						"effect": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Effect indicates the taint effect to match.",
						},
						"toleration_seconds": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "TolerationSeconds represents the period of time the toleration tolerates the taint.",
						},
					},
				},
			},
			"resources": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Compute resource requirements for the agent container.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"limits": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Maximum amount of compute resources allowed.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cpu": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "CPU limit in CPU units (e.g., 500m = 0.5 CPU, 2 = 2 CPUs).",
									},
									"memory": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Memory limit in bytes (e.g., 128Mi, 1Gi).",
									},
								},
							},
						},
						"requests": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Minimum amount of compute resources required.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cpu": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "CPU request in CPU units (e.g., 100m = 0.1 CPU).",
									},
									"memory": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Memory request in bytes (e.g., 128Mi, 1Gi).",
									},
								},
							},
						},
					},
				},
			},
			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Labels to add to all resources.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"annotations": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Annotations to add to all resources.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func installationDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique identifier of the installation.",
			},
			"account_identifier": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The account identifier for the installation.",
			},
			"organization_identifier": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The organization identifier for the installation.",
			},
			"project_identifier": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The project identifier for the installation.",
			},
			"environment_identifier": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The environment identifier for the installation.",
			},
			"agent_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of the installed agent.",
			},
			"delegate_task_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of the delegate task for the installation.",
			},
			"delegate_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of the delegate used for installation.",
			},
			"delegate_task_status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the delegate task (e.g., 'SUCCESS').",
			},
			"agent_details": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Details about the installed agent.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the agent installation.",
						},
						"cluster": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Details about the cluster where the agent is installed.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of the cluster.",
									},
									"namespace": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The namespace where the agent is installed.",
									},
									"uid": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The UID of the cluster.",
									},
									"status": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The status of the cluster (e.g., 'Succeeded').",
									},
								},
							},
						},
					},
				},
			},
			"is_cron_triggered": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the installation was triggered by a cron job.",
			},
			"log_stream_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of the log stream for the installation.",
			},
			"log_stream_created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The timestamp when the log stream was created.",
			},
			"stopped": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the installation has been stopped.",
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The timestamp when the installation was created.",
			},
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The timestamp when the installation was last updated.",
			},
			"created_by": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The user who created the installation.",
			},
			"updated_by": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The user who last updated the installation.",
			},
			"removed": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the installation has been removed.",
			},
		},
	}
}

func setDataDefaults(d *schema.ResourceData) error {
	// Get the current config
	configList := d.Get("config").([]interface{})
	if len(configList) == 0 {
		return nil
	}

	config := configList[0].(map[string]interface{})

	// Ensure data block exists
	if _, ok := config["data"]; !ok {
		config["data"] = []interface{}{map[string]interface{}{}}
	}

	dataList := config["data"].([]interface{})
	if len(dataList) == 0 {
		dataList = []interface{}{map[string]interface{}{}}
	}

	data := dataList[0].(map[string]interface{})

	// Set defaults for data fields if not set
	setDefaultIfNotSet(data, "enable_node_agent", false)
	setDefaultIfNotSet(data, "enable_orphaned_pod", false)
	setDefaultIfNotSet(data, "enable_batch_resources", false)
	setDefaultIfNotSet(data, "collection_window_in_min", 5)

	// Handle nested cron
	if _, ok := data["cron"]; !ok || len(data["cron"].([]interface{})) == 0 {
		data["cron"] = []interface{}{map[string]interface{}{
			"expression": "0/15 * * * *",
		}}
	}

	// Update the config
	config["data"] = []interface{}{data}
	return d.Set("config", []interface{}{config})
}

// Helper function to set default values if not set
func setDefaultIfNotSet(m map[string]interface{}, key string, defaultValue interface{}) {
	if _, ok := m[key]; !ok {
		m[key] = defaultValue
	}
}
