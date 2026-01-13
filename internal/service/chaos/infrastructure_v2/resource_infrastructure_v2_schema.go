package infrastructure_v2

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceChaosInfrastructureV2Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		// Required Fields
		"org_id": {
			Description: "The ID of the organization.",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"project_id": {
			Description: "The ID of the project.",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"environment_id": {
			Description: "The ID of the environment.",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"name": {
			Description: "Name of the infrastructure.",
			Type:        schema.TypeString,
			//ValidateDiagFunc: validateK8sResourceName(),
			StateFunc: func(val interface{}) string {
				return sanitizeK8sResourceName(val.(string))
			},
			Required: true,
		},
		"infra_id": {
			Description: "ID of the infrastructure.",
			Type:        schema.TypeString,
			Required:    true,
		},

		// Optional Fields
		"service_account": {
			Description: "Service account used by the infrastructure.",
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "litmus",
		},
		"namespace": {
			Description: "Kubernetes namespace where the infrastructure will be installed. Maps to the infrastructure namespace.",
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "hce",
		},
		"infra_type": {
			Description:  "Type of the infrastructure. Valid values: KUBERNETES, KUBERNETESV2",
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "KUBERNETESV2",
			ValidateFunc: validation.StringInSlice([]string{"KUBERNETES", "KUBERNETESV2"}, true),
			StateFunc: func(val interface{}) string {
				return strings.ToUpper(val.(string))
			},
		},
		"infra_scope": {
			Description:  "Scope of the infrastructure. Valid values: NAMESPACE, CLUSTER",
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ForceNew:     true, // Cannot be changed after creation
			ValidateFunc: validation.StringInSlice([]string{"NAMESPACE", "CLUSTER"}, true),
			StateFunc: func(val interface{}) string {
				return strings.ToUpper(val.(string))
			},
		},
		"containers": {
			Description: "Container configurations.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"description": {
			Description: "Description of the infrastructure.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"ai_enabled": {
			Description: "Enable AI features for the infrastructure.",
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
		},
		"insecure_skip_verify": {
			Description: "Skip TLS verification for the infrastructure.",
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
		},
		"run_as_user": {
			Description: "User ID to run the infrastructure as.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"run_as_group": {
			Description: "Group ID to run the infrastructure as.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"tags": {
			Description: "Tags for the infrastructure.",
			Type:        schema.TypeList,
			Optional:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"node_selector": {
			Description: "Node selector for the infrastructure pods.",
			Type:        schema.TypeMap,
			Optional:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},

		// Topology
		"label": {
			Description: "Labels to apply to the infrastructure pods.",
			Type:        schema.TypeMap,
			Optional:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"annotation": {
			Description: "Annotations to apply to the infrastructure pods.",
			Type:        schema.TypeMap,
			Optional:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},

		"discovery_agent_id": {
			Description: "ID of the discovery agent to use.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"correlation_id": {
			Description: "Correlation ID for the request.",
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
		},

		// Nested Structures
		"identifier":     identifierSchema(),
		"image_registry": imageRegistrySchema(),
		"mtls":           mtlsSchema(),
		"proxy":          proxySchema(),
		"volumes":        volumesSchema(),
		"volume_mounts":  volumeMountsSchema(),
		"env":            envVarsSchema(),
		"tolerations":    tolerationsSchema(),

		// Computed Fields
		"identity": {
			Description: "Identity for the infrastructure.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"k8s_connector_id": {
			Description: "Kubernetes connector identifier.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"status": {
			Description: "Status of the infrastructure.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"created_at": {
			Description: "Creation timestamp.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"updated_at": {
			Description: "Last update timestamp.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"install_command": {
			Description: "Installation command for the infrastructure.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"infra_namespace": {
			Description: "Namespace where the infrastructure is installed.",
			Type:        schema.TypeString,
			Computed:    true,
		},
	}
}
