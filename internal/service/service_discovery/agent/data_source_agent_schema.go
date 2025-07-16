package agent

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// AgentDataSourceSchema returns the schema for the agent data source
func AgentDataSourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		// Query parameters
		"name": {
			Description:  "The name of the agent. Either this or 'identity' must be provided.",
			Type:         schema.TypeString,
			Optional:     true,
			AtLeastOneOf: []string{"name", "identity"},
		},
		"identity": {
			Description:  "The unique identity of the agent. Either this or 'name' must be provided.",
			Type:         schema.TypeString,
			Optional:     true,
			AtLeastOneOf: []string{"name", "identity"},
		},

		// Required fields
		"environment_identifier": {
			Description: "The environment identifier of the agent. This is a required field.",
			Type:        schema.TypeString,
			Required:    true,
		},

		// Optional fields
		"org_identifier": {
			Description: "The organization identifier of the agent (optional). Must be 1-64 characters and contain only alphanumeric characters, hyphens, or underscores.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"project_identifier": {
			Description: "The project identifier of the agent (optional). Must be 1-64 characters and contain only alphanumeric characters, hyphens, or underscores.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"infra_identifier": {
			Description: "The infrastructure identifier of the agent (optional).",
			Type:        schema.TypeString,
			Optional:    true,
		},

		// Computed fields
		"id": {
			Description: "The unique identifier of the agent.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"description": {
			Description: "Description of the agent.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"tags": {
			Description: "Key-value map of resource tags for the agent.",
			Type:        schema.TypeMap,
			Computed:    true,
		},
		"installation_type": {
			Description: "Type of installation for the agent.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"network_map_count": {
			Description: "Number of network maps associated with this agent.",
			Type:        schema.TypeInt,
			Computed:    true,
		},
		"permanent_installation": {
			Description: "Whether this is a permanent installation.",
			Type:        schema.TypeBool,
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
		"webhook_url": {
			Description: "Webhook URL for the agent.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"correlation_id": {
			Description: "Correlation ID for the agent.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"config": {
			Type:        schema.TypeList,
			Computed:    true,
			Description: "Configuration for the agent.",
			Elem:        dataSourceConfigSchema(),
		},
		"installation_details": {
			Type:        schema.TypeList,
			Computed:    true,
			Description: "Installation details of the agent.",
			Elem:        dataSourceInstallationDetailsSchema(),
		},
	}
}

func dataSourceConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"collector_image": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Docker image for the collector.",
			},
			"log_watcher_image": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Docker image for the log watcher.",
			},
			"skip_secure_verify": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether to skip TLS verification.",
			},
			"image_pull_secrets": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of image pull secrets.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"kubernetes": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Kubernetes-specific configuration.",
				Elem:        dataSourceKubernetesSchema(),
			},
			"mtls": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "mTLS configuration.",
				Elem:        dataSourceMtlsSchema(),
			},
			"proxy": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Proxy configuration.",
				Elem:        dataSourceProxySchema(),
			},
			"data": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Data collection configuration.",
				Elem:        dataSourceDataSchema(),
			},
		},
	}
}

func dataSourceDataSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enable_node_agent": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether to enable node agent.",
			},
			"node_agent_selector": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Node selector for the node agent.",
			},
			"namespace_selector": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Namespace selector for the agent.",
			},
			"enable_orphaned_pod": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether to enable orphaned pod detection.",
			},
			"enable_batch_resources": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether to enable batch resources.",
			},
			"collection_window_in_min": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Collection window in minutes.",
			},
			"blacklisted_namespaces": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of namespaces to exclude from discovery.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"observed_namespaces": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of namespaces to observe.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"cron": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Cron schedule for data collection.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"expression": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Cron expression for scheduling.",
						},
					},
				},
			},
		},
	}
}

func dataSourceProxySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"http_proxy": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "HTTP proxy URL.",
			},
			"https_proxy": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "HTTPS proxy URL.",
			},
			"no_proxy": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Comma-separated list of hosts that should not use the proxy.",
			},
			"url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Proxy server URL.",
			},
		},
	}
}

func dataSourceMtlsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cert_path": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Path to the certificate file.",
			},
			"key_path": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Path to the key file.",
			},
			"secret_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of the Kubernetes secret containing the certificate and key.",
			},
			"url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URL of the mTLS server.",
			},
		},
	}
}

func dataSourceKubernetesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"disable_namespace_creation": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether to disable namespace creation.",
			},
			"namespaced": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the agent is namespaced.",
			},
			"namespace": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Kubernetes namespace to use",
			},
			"service_account": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Service account to use",
			},
			"run_as_user": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The user ID to run as.",
			},
			"run_as_group": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The group ID to run as.",
			},
			"image_pull_policy": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The image pull policy.",
			},
			"node_selector": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Node selector labels.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"tolerations": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Tolerations for pod assignment.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The taint key that the toleration applies to.",
						},
						"operator": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Operator represents a key's relationship to the value.",
						},
						"value": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The taint value the toleration matches to.",
						},
						"effect": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Effect indicates the taint effect to match.",
						},
						"toleration_seconds": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "TolerationSeconds represents the period of time the toleration tolerates the taint.",
						},
					},
				},
			},
			"resources": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Compute resource requirements for the agent container.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"limits": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Maximum amount of compute resources allowed.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cpu": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "CPU limit in CPU units (e.g., 500m = 0.5 CPU, 2 = 2 CPUs).",
									},
									"memory": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Memory limit in bytes (e.g., 128Mi, 1Gi).",
									},
								},
							},
						},
						"requests": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Minimum amount of compute resources required.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cpu": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "CPU request in CPU units (e.g., 100m = 0.1 CPU).",
									},
									"memory": {
										Type:        schema.TypeString,
										Computed:    true,
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
				Computed:    true,
				Description: "Labels to add to all resources.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"annotations": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Annotations to add to all resources.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

// dataSourceInstallationDetailsSchema defines the schema for installation details
func dataSourceInstallationDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Installation ID.",
			},
			"account_identifier": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Account identifier of the installation.",
			},
			"org_identifier": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Organization identifier of the installation.",
			},
			"project_identifier": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Project identifier of the installation.",
			},
			"environment_identifier": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Environment identifier of the installation.",
			},
			"delegate_task_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "ID of the delegate task.",
			},
			"delegate_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "ID of the delegate.",
			},
			"delegate_task_status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Status of the delegate task.",
			},
			"agent_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "ID of the agent.",
			},
			"is_cron_triggered": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the installation was triggered by a cron job.",
			},
			"log_stream_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "ID of the log stream.",
			},
			"log_stream_created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "When the log stream was created.",
			},
			"stopped": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the installation has been stopped.",
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "When the installation was created.",
			},
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "When the installation was last updated.",
			},
			"created_by": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Who created the installation.",
			},
			"updated_by": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Who last updated the installation.",
			},
			"removed_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "When the installation was removed.",
			},
			"removed": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the installation has been removed.",
			},
			"agent_details": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Details about the agent installation.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Status of the agent.",
						},
						"cluster": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Cluster information.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Name of the cluster.",
									},
									"namespace": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Namespace of the cluster.",
									},
									"uid": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "UID of the cluster.",
									},
									"status": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Status of the cluster.",
									},
								},
							},
						},
						"node": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "List of nodes in the cluster.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Name of the node.",
									},
									"namespace": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Namespace of the node.",
									},
									"uid": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "UID of the node.",
									},
									"status": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Status of the node.",
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
