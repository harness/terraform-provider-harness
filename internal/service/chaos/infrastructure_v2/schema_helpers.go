package infrastructure_v2

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/harness/harness-go-sdk/harness/chaos"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// imageRegistrySchema defines the schema for container image registry configuration
func imageRegistrySchema() *schema.Schema {
	return &schema.Schema{
		Description: "Configuration for the container image registry.",
		Type:        schema.TypeList,
		Optional:    true,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"registry_server": {
					Description: "The container image registry server URL (e.g., docker.io, gcr.io).",
					Type:        schema.TypeString,
					Optional:    true,
					Default:     "docker.io",
				},
				"registry_account": {
					Description: "The account name for the container registry.",
					Type:        schema.TypeString,
					Default:     "harness",
					Optional:    true,
				},
				"secret_name": {
					Description: "Name of the Kubernetes secret containing registry credentials.",
					Type:        schema.TypeString,
					Optional:    true,
				},
				"is_override_allowed": {
					Description: "Whether override is allowed for this registry.",
					Type:        schema.TypeBool,
					Optional:    true,
					Default:     false,
				},
				"use_custom_images": {
					Description: "Whether to use custom images instead of default ones.",
					Type:        schema.TypeBool,
					Optional:    true,
					Default:     false,
				},
				"custom_images": {
					Description: "Custom image configurations. Required when use_custom_images is true.",
					Type:        schema.TypeList,
					Optional:    true,
					MaxItems:    1,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"ddcr": {
								Description: "Custom image for ddcr.",
								Type:        schema.TypeString,
								Optional:    true,
								Computed:    true,
							},
							"ddcr_fault": {
								Description: "Custom image for ddcr-fault.",
								Type:        schema.TypeString,
								Optional:    true,
								Computed:    true,
							},
							"ddcr_lib": {
								Description: "Custom image for ddcr-lib.",
								Type:        schema.TypeString,
								Optional:    true,
								Computed:    true,
							},
							"log_watcher": {
								Description: "Custom image for log-watcher.",
								Type:        schema.TypeString,
								Optional:    true,
								Computed:    true,
							},
						},
					},
				},
				"is_default": {
					Description: "Whether this is the default registry.",
					Type:        schema.TypeBool,
					Optional:    true,
				},
				"is_private": {
					Description: "Whether the registry is private.",
					Type:        schema.TypeBool,
					Optional:    true,
				},
				// Computed fields
				"infra_id": {
					Description: "ID of the infrastructure.",
					Type:        schema.TypeString,
					Computed:    true,
				},
				"identifier": {
					Description: "Scoped identifiers for the registry.",
					Type:        schema.TypeList,
					Computed:    true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"account_identifier": {
								Description: "Harness account identifier.",
								Type:        schema.TypeString,
								Computed:    true,
							},
							"org_identifier": {
								Description: "Harness organization identifier.",
								Type:        schema.TypeString,
								Computed:    true,
							},
							"project_identifier": {
								Description: "Harness project identifier.",
								Type:        schema.TypeString,
								Computed:    true,
							},
						},
					},
				},
				"created_at": {
					Description: "Timestamp when the registry was created.",
					Type:        schema.TypeString,
					Computed:    true,
				},
				"updated_at": {
					Description: "Timestamp when the registry was last updated.",
					Type:        schema.TypeString,
					Computed:    true,
				},
			},
		},
	}
}

// mtlsSchema defines the schema for mTLS configuration
func mtlsSchema() *schema.Schema {
	return &schema.Schema{
		Description: "mTLS configuration for the infrastructure.",
		Type:        schema.TypeList,
		Optional:    true,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"cert_path": {
					Type:         schema.TypeString,
					Optional:     true,
					Description:  "Path to the certificate file for mTLS",
					ValidateFunc: validation.StringIsNotEmpty,
				},
				"key_path": {
					Type:         schema.TypeString,
					Optional:     true,
					Description:  "Path to the private key file for mTLS",
					ValidateFunc: validation.StringIsNotEmpty,
				},
				"secret_name": {
					Type:         schema.TypeString,
					Optional:     true,
					Description:  "Name of the Kubernetes secret containing mTLS certificates",
					ValidateFunc: validation.StringIsNotEmpty,
				},
				"url": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "URL for the mTLS endpoint",
				},
			},
		},
	}
}

func identifierSchema() *schema.Schema {
	return &schema.Schema{
		Description: "Identifier for the infrastructure.",
		Type:        schema.TypeList,
		Computed:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"account_identifier": {
					Description: "Account identifier.",
					Type:        schema.TypeString,
					Computed:    true,
				},
				"org_identifier": {
					Description: "Organization identifier.",
					Type:        schema.TypeString,
					Computed:    true,
				},
				"project_identifier": {
					Description: "Project identifier.",
					Type:        schema.TypeString,
					Computed:    true,
				},
			},
		},
	}
}

// proxySchema defines the schema for proxy configuration
func proxySchema() *schema.Schema {
	return &schema.Schema{
		Description: "Proxy configuration for the infrastructure.",
		Type:        schema.TypeList,
		Optional:    true,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"url": {
					Description: "Proxy URL.",
					Type:        schema.TypeString,
					Required:    true,
				},
				"http_proxy": {
					Description: "HTTP proxy URL.",
					Type:        schema.TypeString,
					Optional:    true,
				},
				"https_proxy": {
					Description: "HTTPS proxy URL.",
					Type:        schema.TypeString,
					Optional:    true,
				},
				"no_proxy": {
					Description: "List of hosts that should not use proxy.",
					Type:        schema.TypeString,
					Optional:    true,
				},
			},
		},
	}
}

// volumesSchema defines the schema for infrastructure volumes
func volumesSchema() *schema.Schema {
	return &schema.Schema{
		Description: "Volumes to be created in the infrastructure.",
		Type:        schema.TypeList,
		Optional:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": {
					Description: "Name of the volume. Must be a DNS_LABEL and unique within the pod.",
					Type:        schema.TypeString,
					Required:    true,
					ValidateFunc: validation.StringMatch(
						regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`),
						"a lowercase RFC 1123 label must consist of alphanumeric characters or '-', and must start and end with an alphanumeric character",
					),
				},
				"size_limit": {
					Description: "Size limit of the volume. Example: '10Gi', '100Mi'",
					Type:        schema.TypeString,
					Optional:    true,
					ValidateFunc: validation.StringMatch(
						regexp.MustCompile(`^[0-9]+(E|P|T|G|M|K|Ei|Pi|Ti|Gi|Mi|Ki)$`),
						"must be a valid resource quantity (e.g., 10Gi, 100Mi)",
					),
				},
			},
		},
	}
}

// volumeMountsSchema defines the schema for volume mounts in containers
func volumeMountsSchema() *schema.Schema {
	return &schema.Schema{
		Description: "Volume mounts for the container.",
		Type:        schema.TypeList,
		Optional:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": {
					Description:  "This must match the Name of a Volume.",
					Type:         schema.TypeString,
					Required:     true,
					ValidateFunc: validation.StringIsNotEmpty,
				},
				"mount_path": {
					Description:  "Path within the container at which the volume should be mounted. Must not contain ':'.",
					Type:         schema.TypeString,
					Required:     true,
					ValidateFunc: validation.StringDoesNotContainAny(":"),
				},
				"mount_propagation": {
					Description:  "Determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationNone is used.",
					Type:         schema.TypeString,
					Optional:     true,
					Default:      "None",
					ValidateFunc: validation.StringInSlice([]string{"None", "HostToContainer", "Bidirectional"}, false),
				},
				"read_only": {
					Description: "Mounted read-only if true, read-write otherwise.",
					Type:        schema.TypeBool,
					Optional:    true,
					Default:     false,
				},
				"sub_path": {
					Description:  "Path within the volume from which the container's volume should be mounted. Mutually exclusive with sub_path_expr.",
					Type:         schema.TypeString,
					Optional:     true,
					ValidateFunc: validation.StringDoesNotContainAny(":"),
				},
				"sub_path_expr": {
					Description:  "Expanded path within the volume from which the container's volume should be mounted. Behaves similarly to sub_path but environment variable references $(VAR_NAME) are expanded using the container's environment. Mutually exclusive with sub_path.",
					Type:         schema.TypeString,
					Optional:     true,
					ValidateFunc: validation.StringDoesNotContainAny(":"),
				},
			},
		},
	}
}

// envVarsSchema defines the schema for environment variables
func envVarsSchema() *schema.Schema {
	return &schema.Schema{
		Description: "List of environment variables to set in the container.",
		Type:        schema.TypeList,
		Optional:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": {
					Description: "Name of the environment variable. Must be a C_IDENTIFIER.",
					Type:        schema.TypeString,
					Required:    true,
					ValidateFunc: validation.StringMatch(
						regexp.MustCompile(`^[A-Za-z_][A-Za-z0-9_]*$`),
						"environment variable name must be a valid C identifier (matching regex [A-Za-z_][A-Za-z0-9_]*)",
					),
				},
				"value": {
					Description: "Variable references $(VAR_NAME) are expanded using the container's environment. If the variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to \"\".",
					Type:        schema.TypeString,
					Optional:    true,
					Default:     "",
				},
				"value_from": {
					Description: "Source for the environment variable's value. Cannot be used if value is not empty.",
					Type:        schema.TypeString,
					Optional:    true,
					ValidateFunc: validation.StringInSlice(
						[]string{
							"configMapKeyRef",
							"secretKeyRef",
							"value",
							"valueFrom",
						},
						false,
					),
				},
				"key": {
					Description: "Variable name from a ConfigMap or Secret. Required when value_from is configMapKeyRef or secretKeyRef.",
					Type:        schema.TypeString,
					Optional:    true,
				},
			},
		},
	}
}

// tolerationsSchema defines the schema for pod tolerations
func tolerationsSchema() *schema.Schema {
	return &schema.Schema{
		Description: "If specified, the pod's tolerations.",
		Type:        schema.TypeList,
		Optional:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"effect": {
					Description:  "Effect indicates the taint effect to match. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
					Type:         schema.TypeString,
					Required:     true,
					ValidateFunc: validation.StringInSlice([]string{"NoSchedule", "PreferNoSchedule", "NoExecute"}, false),
				},
				"key": {
					Description: "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists.",
					Type:        schema.TypeString,
					Required:    true,
				},
				"operator": {
					Description:  "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal.",
					Type:         schema.TypeString,
					Required:     true,
					ValidateFunc: validation.StringInSlice([]string{"Exists", "Equal"}, false),
				},
				"toleration_seconds": {
					Description: "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
					Type:        schema.TypeInt,
					Optional:    true,
				},
				"value": {
					Description: "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
					Type:        schema.TypeString,
					Optional:    true,
				},
			},
		},
	}
}

// validateK8sResourceName ensures the string matches Kubernetes resource name requirements
func validateK8sResourceName() schema.SchemaValidateDiagFunc {
	return validation.ToDiagFunc(validation.All(
		validation.StringMatch(
			regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`),
			"must consist of lower case alphanumeric characters or '-', and must start and end with an alphanumeric character",
		),
		validation.StringLenBetween(1, 63),
	))
}

// sanitizeK8sResourceName converts a string to be compatible with Kubernetes resource name requirements
func sanitizeK8sResourceName(name string) string {
	// Convert to lowercase
	name = strings.ToLower(name)

	// Replace invalid characters with '-'
	re := regexp.MustCompile(`[^a-z0-9-]`)
	name = re.ReplaceAllString(name, "-")

	// Remove leading/trailing dashes
	name = strings.Trim(name, "-")

	// Ensure the name is not empty
	if name == "" {
		name = "infra"
	}

	// Truncate to 63 characters if needed
	if len(name) > 63 {
		name = name[:63]
	}

	return name
}

func ensureRegistryInfraID(reg *chaos.ImageRegistryImageRegistryV2, envID, infraID string) {
	if reg != nil && reg.InfraID == "" && envID != "" && infraID != "" {
		reg.InfraID = fmt.Sprintf("%s/%s", envID, infraID)
	}
}
