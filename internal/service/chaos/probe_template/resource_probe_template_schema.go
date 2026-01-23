package probe_template

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceProbeTemplateSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		// Required Fields
		"identity": {
			Description:  "Unique identifier for the probe template (immutable).",
			Type:         schema.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},
		"name": {
			Description:  "Name of the probe template.",
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},
		"hub_identity": {
			Description:  "Identity of the chaos hub this probe template belongs to.",
			Type:         schema.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},
		"type": {
			Description: "Type of the probe template. Valid values: httpProbe, cmdProbe, k8sProbe, promProbe, sloProbe, datadogProbe, dynatraceProbe, containerProbe, apmProbe.",
			Type:        schema.TypeString,
			Required:    true,
			ValidateFunc: validation.StringInSlice([]string{
				"httpProbe",
				"cmdProbe",
				"k8sProbe",
				"promProbe",
				"sloProbe",
				"datadogProbe",
				"dynatraceProbe",
				"containerProbe",
				"apmProbe",
			}, false),
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
			Description: "Description of the probe template.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tags": {
			Description: "Tags to associate with the probe template.",
			Type:        schema.TypeList,
			Optional:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"infrastructure_type": {
			Description: "Infrastructure type for the probe template. Valid values: Kubernetes, KubernetesV2, Windows, Linux, CloudFoundry, Container.",
			Type:        schema.TypeString,
			Optional:    true,
		},

		// Probe Properties - HTTP Probe
		"http_probe": {
			Description: "HTTP probe configuration. Required when type is 'httpProbe'.",
			Type:        schema.TypeList,
			Optional:    true,
			MaxItems:    1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"url": {
						Description:  "URL to probe.",
						Type:         schema.TypeString,
						Required:     true,
						ValidateFunc: validation.StringIsNotEmpty,
					},
					"method": {
						Description: "HTTP method configuration with GET or POST.",
						Type:        schema.TypeList,
						Optional:    true,
						MaxItems:    1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"get": {
									Description: "GET method configuration.",
									Type:        schema.TypeList,
									Optional:    true,
									MaxItems:    1,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"criteria": {
												Description: "Response criteria (e.g., '==', '!=', 'contains').",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"response_body": {
												Description: "Expected response body.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"response_code": {
												Description: "Expected HTTP response code (e.g., '200', '404').",
												Type:        schema.TypeString,
												Optional:    true,
											},
										},
									},
								},
								"post": {
									Description: "POST method configuration.",
									Type:        schema.TypeList,
									Optional:    true,
									MaxItems:    1,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"criteria": {
												Description: "Response criteria (e.g., '==', '!=', 'contains').",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"response_body": {
												Description: "Expected response body.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"response_code": {
												Description: "Expected HTTP response code (e.g., '200', '404').",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"body": {
												Description: "POST request body.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"body_path": {
												Description: "Path to file containing POST body.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"content_type": {
												Description: "Content-Type header for POST request.",
												Type:        schema.TypeString,
												Optional:    true,
											},
										},
									},
								},
							},
						},
					},
					"headers": {
						Description: "HTTP headers.",
						Type:        schema.TypeMap,
						Optional:    true,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
					},
					"auth": {
						Description: "Authentication configuration.",
						Type:        schema.TypeList,
						Optional:    true,
						MaxItems:    1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"type": {
									Description: "Auth type (basic, bearer, etc.).",
									Type:        schema.TypeString,
									Required:    true,
								},
								"username": {
									Description: "Username for basic auth.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"password": {
									Description: "Password for basic auth.",
									Type:        schema.TypeString,
									Optional:    true,
									Sensitive:   true,
								},
								"token": {
									Description: "Token for bearer auth.",
									Type:        schema.TypeString,
									Optional:    true,
									Sensitive:   true,
								},
							},
						},
					},
					"tls_config": {
						Description: "TLS configuration.",
						Type:        schema.TypeList,
						Optional:    true,
						MaxItems:    1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"insecure_skip_verify": {
									Description: "Skip TLS certificate verification.",
									Type:        schema.TypeBool,
									Optional:    true,
									Default:     false,
								},
								"ca_cert": {
									Description: "CA certificate.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"client_cert": {
									Description: "Client certificate.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"client_key": {
									Description: "Client key.",
									Type:        schema.TypeString,
									Optional:    true,
									Sensitive:   true,
								},
							},
						},
					},
				},
			},
		},

		// Probe Properties - CMD Probe
		"cmd_probe": {
			Description: "Command probe configuration. Required when type is 'cmdProbe'.",
			Type:        schema.TypeList,
			Optional:    true,
			MaxItems:    1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"command": {
						Description:  "Command to execute.",
						Type:         schema.TypeString,
						Required:     true,
						ValidateFunc: validation.StringIsNotEmpty,
					},
					"source": {
						Description: "Source of the command (inline, configMap, secret).",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"env": {
						Description: "Environment variables for the command.",
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
					"comparator": {
						Description: "Comparator for command output validation.",
						Type:        schema.TypeList,
						Optional:    true,
						MaxItems:    1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"type": {
									Description: "Comparator type (string, int, float).",
									Type:        schema.TypeString,
									Required:    true,
								},
								"criteria": {
									Description: "Comparison criteria (==, !=, <, >, <=, >=, contains, matches, notMatches, oneOf).",
									Type:        schema.TypeString,
									Required:    true,
								},
								"value": {
									Description: "Expected value.",
									Type:        schema.TypeString,
									Required:    true,
								},
							},
						},
					},
				},
			},
		},

		// Probe Properties - K8s Probe
		"k8s_probe": {
			Description: "Kubernetes probe configuration. Required when type is 'k8sProbe'.",
			Type:        schema.TypeList,
			Optional:    true,
			MaxItems:    1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"group": {
						Description: "API group (e.g., 'apps', 'batch').",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"version": {
						Description: "API version (e.g., 'v1', 'v1beta1').",
						Type:        schema.TypeString,
						Required:    true,
					},
					"resource": {
						Description:  "Resource type (e.g., 'pods', 'deployments').",
						Type:         schema.TypeString,
						Required:     true,
						ValidateFunc: validation.StringIsNotEmpty,
					},
					"namespace": {
						Description: "Kubernetes namespace.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"field_selector": {
						Description: "Field selector for filtering resources.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"label_selector": {
						Description: "Label selector for filtering resources.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"resource_names": {
						Description: "Comma-separated list of resource names.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"operation": {
						Description: "Operation to perform (create, delete, present, absent, etc.).",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},

		// Probe Properties - APM Probe
		"apm_probe": {
			Description: "APM probe configuration. Required when type is 'apmProbe'.",
			Type:        schema.TypeList,
			Optional:    true,
			MaxItems:    1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"apm_type": {
						Description: "APM provider type. Valid values: Prometheus, AppDynamics, SplunkObservability, Dynatrace, NewRelic, Datadog, GCPCloudMonitoring.",
						Type:        schema.TypeString,
						Required:    true,
						ValidateFunc: validation.StringInSlice([]string{
							"Prometheus",
							"AppDynamics",
							"SplunkObservability",
							"Dynatrace",
							"NewRelic",
							"Datadog",
							"GCPCloudMonitoring",
						}, false),
					},
					"comparator": {
						Description: "Comparator for APM metric validation.",
						Type:        schema.TypeList,
						Optional:    true,
						MaxItems:    1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"type": {
									Description: "Comparator type (string, int, float).",
									Type:        schema.TypeString,
									Required:    true,
								},
								"criteria": {
									Description: "Comparison criteria (==, !=, <, >, <=, >=, contains, matches, notMatches, oneOf).",
									Type:        schema.TypeString,
									Required:    true,
								},
								"value": {
									Description: "Expected value.",
									Type:        schema.TypeString,
									Required:    true,
								},
							},
						},
					},
					"prometheus_inputs": {
						Description: "Prometheus-specific inputs. Required when apm_type is 'Prometheus'.",
						Type:        schema.TypeList,
						Optional:    true,
						MaxItems:    1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"connector_id": {
									Description: "Harness connector ID for Prometheus.",
									Type:        schema.TypeString,
									Required:    true,
								},
								"query": {
									Description: "PromQL query string.",
									Type:        schema.TypeString,
									Required:    true,
								},
								"tls_config": {
									Description: "TLS configuration for Prometheus connection.",
									Type:        schema.TypeList,
									Optional:    true,
									MaxItems:    1,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"ca_cert_secret": {
												Description: "Harness secret identifier for CA certificate.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"client_cert_secret": {
												Description: "Harness secret identifier for client certificate.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"client_key_secret": {
												Description: "Harness secret identifier for client key.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"insecure_skip_verify": {
												Description: "Skip TLS certificate verification.",
												Type:        schema.TypeBool,
												Optional:    true,
											},
										},
									},
								},
							},
						},
					},
					"datadog_inputs": {
						Description: "Datadog-specific inputs. Required when apm_type is 'Datadog'.",
						Type:        schema.TypeList,
						Optional:    true,
						MaxItems:    1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"connector_id": {
									Description: "Harness connector ID for Datadog.",
									Type:        schema.TypeString,
									Required:    true,
								},
								"query": {
									Description: "Datadog query string.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"duration_in_min": {
									Description: "Duration in minutes for the Datadog query.",
									Type:        schema.TypeInt,
									Optional:    true,
								},
								"synthetics_test": {
									Description: "Datadog Synthetics test configuration.",
									Type:        schema.TypeList,
									Optional:    true,
									MaxItems:    1,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"public_id": {
												Description: "Public ID of the Datadog Synthetics test.",
												Type:        schema.TypeString,
												Required:    true,
											},
											"test_type": {
												Description: "Type of Synthetics test (api, browser).",
												Type:        schema.TypeString,
												Optional:    true,
											},
										},
									},
								},
							},
						},
					},
					"dynatrace_inputs": {
						Description: "Dynatrace-specific inputs. Required when apm_type is 'Dynatrace'.",
						Type:        schema.TypeList,
						Optional:    true,
						MaxItems:    1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"connector_id": {
									Description: "Harness connector ID for Dynatrace.",
									Type:        schema.TypeString,
									Required:    true,
								},
								"duration_in_min": {
									Description: "Duration in minutes for the Dynatrace query.",
									Type:        schema.TypeInt,
									Optional:    true,
								},
								"metrics": {
									Description: "Dynatrace metrics configuration.",
									Type:        schema.TypeList,
									Optional:    true,
									MaxItems:    1,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"entity_selector": {
												Description: "Dynatrace entity selector.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"metrics_selector": {
												Description: "Dynatrace metrics selector.",
												Type:        schema.TypeString,
												Optional:    true,
											},
										},
									},
								},
							},
						},
					},
					"app_dynamics_inputs": {
						Description: "AppDynamics-specific inputs. Required when apm_type is 'AppDynamics'.",
						Type:        schema.TypeList,
						Optional:    true,
						MaxItems:    1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"connector_id": {
									Description: "Harness connector ID for AppDynamics.",
									Type:        schema.TypeString,
									Required:    true,
								},
								"appd_metrics": {
									Description: "AppDynamics metrics configuration.",
									Type:        schema.TypeList,
									Optional:    true,
									MaxItems:    1,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"application_name": {
												Description: "AppDynamics application name.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"metrics_full_path": {
												Description: "Full path to the AppDynamics metric.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"duration_in_min": {
												Description: "Duration in minutes for the AppDynamics query.",
												Type:        schema.TypeInt,
												Optional:    true,
											},
										},
									},
								},
							},
						},
					},
					"new_relic_inputs": {
						Description: "NewRelic-specific inputs. Required when apm_type is 'NewRelic'.",
						Type:        schema.TypeList,
						Optional:    true,
						MaxItems:    1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"connector_id": {
									Description: "Harness connector ID for NewRelic.",
									Type:        schema.TypeString,
									Required:    true,
								},
								"new_relic_metric": {
									Description: "NewRelic metric configuration.",
									Type:        schema.TypeList,
									Optional:    true,
									MaxItems:    1,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"query": {
												Description: "NRQL query string.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"query_metric": {
												Description: "NewRelic query metric name.",
												Type:        schema.TypeString,
												Optional:    true,
											},
										},
									},
								},
							},
						},
					},
					"splunk_observability_inputs": {
						Description: "SplunkObservability-specific inputs. Required when apm_type is 'SplunkObservability'.",
						Type:        schema.TypeList,
						Optional:    true,
						MaxItems:    1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"connector_id": {
									Description: "Harness connector ID for Splunk Observability.",
									Type:        schema.TypeString,
									Required:    true,
								},
								"splunk_observability_metrics": {
									Description: "Splunk Observability metrics configuration.",
									Type:        schema.TypeList,
									Optional:    true,
									MaxItems:    1,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"query": {
												Description: "Splunk Observability query string.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"duration_in_min": {
												Description: "Duration in minutes for the Splunk query.",
												Type:        schema.TypeInt,
												Optional:    true,
											},
										},
									},
								},
							},
						},
					},
					"gcp_cloud_monitoring_inputs": {
						Description: "GCP Cloud Monitoring-specific inputs. Required when apm_type is 'GCPCloudMonitoring'.",
						Type:        schema.TypeList,
						Optional:    true,
						MaxItems:    1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"project_id": {
									Description: "GCP project ID.",
									Type:        schema.TypeString,
									Required:    true,
								},
								"query": {
									Description: "GCP monitoring query string.",
									Type:        schema.TypeString,
									Required:    true,
								},
								"service_account_key": {
									Description: "GCP service account key (JSON).",
									Type:        schema.TypeString,
									Required:    true,
									Sensitive:   true,
								},
							},
						},
					},
				},
			},
		},

		// Run Properties
		"run_properties": {
			Description: "Run properties for the probe template execution.",
			Type:        schema.TypeList,
			Optional:    true,
			MaxItems:    1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"initial_delay": {
						Description: "Initial delay before probe execution (e.g., '5s', '1m').",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"interval": {
						Description: "Interval between probe executions (e.g., '10s', '30s').",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"timeout": {
						Description: "Timeout for probe execution (e.g., '30s', '5m').",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"polling_interval": {
						Description: "Polling interval for continuous probes (e.g., '2s', '5s').",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"attempt": {
						Description: "Number of attempts.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"retry": {
						Description: "Number of retries.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"stop_on_failure": {
						Description: "Whether to stop on failure.",
						Type:        schema.TypeBool,
						Optional:    true,
						Default:     false,
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
			Description: "Template variables that can be used in the probe.",
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
		"revision": {
			Description: "Revision number of the probe template.",
			Type:        schema.TypeInt,
			Computed:    true,
		},
		"is_default": {
			Description: "Whether this is the default version for predefined probes.",
			Type:        schema.TypeBool,
			Computed:    true,
		},
		"hub_ref": {
			Description: "Hub reference.",
			Type:        schema.TypeString,
			Computed:    true,
		},
	}
}
