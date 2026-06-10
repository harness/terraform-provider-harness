package security_governance_v3

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// resourceChaosSecurityGovernanceConditionV3Schema returns the schema for the
// V3 (REST-backed) security governance condition resource. The schema mirrors
// the GraphQL (V1) resource so existing configurations migrate with minimal
// changes, with the addition of native map support for namespace_labels.
func resourceChaosSecurityGovernanceConditionV3Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"org_id": {
			Description:  "The organization ID of the security governance condition",
			Type:         schema.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},
		"project_id": {
			Description:  "The project ID of the security governance condition",
			Type:         schema.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},
		"name": {
			Description:  "Name of the security governance condition",
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},
		"description": {
			Description: "Description of the security governance condition",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tags": {
			Description: "Tags for the security governance condition",
			Type:        schema.TypeList,
			Optional:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"infra_type": {
			Description: "Type of infrastructure (Kubernetes, KubernetesV2, Linux, Windows, CloudFoundry, Container)",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			ValidateFunc: validation.StringInSlice(
				[]string{
					"Kubernetes",
					"KubernetesV2",
					"Linux",
					"Windows",
					"CloudFoundry",
					"Container",
				},
				false,
			),
		},
		"fault_spec": {
			Description: "Specification for faults to be included in the condition",
			Type:        schema.TypeList,
			Required:    true,
			MaxItems:    1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"operator": {
						Description:  "Operator for comparing faults (EQUAL_TO or NOT_EQUAL_TO)",
						Type:         schema.TypeString,
						Required:     true,
						ValidateFunc: validation.StringInSlice([]string{"EQUAL_TO", "NOT_EQUAL_TO"}, false),
					},
					"faults": {
						Description: "List of fault specifications",
						Type:        schema.TypeList,
						Required:    true,
						MinItems:    1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"fault_type": {
									Description:  "Type of the fault (FAULT or FAULT_GROUP)",
									Type:         schema.TypeString,
									Required:     true,
									ValidateFunc: validation.StringInSlice([]string{"FAULT", "FAULT_GROUP"}, false),
								},
								"name": {
									Description: "Name of the fault",
									Type:        schema.TypeString,
									Required:    true,
								},
							},
						},
					},
				},
			},
		},
		"k8s_spec": {
			Description: "Kubernetes specific configuration (required when infra_type is Kubernetes or KubernetesV2)",
			Type:        schema.TypeList,
			Optional:    true,
			MaxItems:    1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"infra_spec": {
						Description: "Infrastructure specification",
						Type:        schema.TypeList,
						Optional:    true,
						MaxItems:    1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"operator": {
									Description:  "Operator for comparing infrastructure IDs (EQUAL_TO or NOT_EQUAL_TO)",
									Type:         schema.TypeString,
									Required:     true,
									ValidateFunc: validation.StringInSlice([]string{"EQUAL_TO", "NOT_EQUAL_TO"}, false),
								},
								"infra_ids": {
									Description: "List of infrastructure IDs to apply the condition to",
									Type:        schema.TypeList,
									Required:    true,
									MinItems:    1,
									Elem: &schema.Schema{
										Type: schema.TypeString,
									},
								},
							},
						},
					},
					"application_spec": {
						Description: "Application specification",
						Type:        schema.TypeList,
						Optional:    true,
						MaxItems:    1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"operator": {
									Description:  "Operator for application matching (EQUAL_TO or NOT_EQUAL_TO)",
									Type:         schema.TypeString,
									Required:     true,
									ValidateFunc: validation.StringInSlice([]string{"EQUAL_TO", "NOT_EQUAL_TO"}, false),
								},
								"workloads": {
									Description: "List of workloads to include/exclude",
									Type:        schema.TypeList,
									Optional:    true,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"namespace": {
												Description: "Namespace of the workload",
												Type:        schema.TypeString,
												Required:    true,
											},
											"kind": {
												Description: "Kind of the workload (e.g., deployment, statefulset)",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"label": {
												Description: "Label selector for the workload",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"services": {
												Description: "List of services associated with the workload",
												Type:        schema.TypeList,
												Optional:    true,
												Elem:        &schema.Schema{Type: schema.TypeString},
											},
											"application_map_id": {
												Description: "ID for the application map",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"namespace_labels": {
												Description: "Namespace labels to match against, as key-value pairs",
												Type:        schema.TypeMap,
												Optional:    true,
												Elem:        &schema.Schema{Type: schema.TypeString},
											},
										},
									},
								},
							},
						},
					},
					"chaos_service_account_spec": {
						Description: "Chaos service account specification",
						Type:        schema.TypeList,
						Optional:    true,
						MaxItems:    1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"operator": {
									Description:  "Operator for service account matching (EQUAL_TO or NOT_EQUAL_TO)",
									Type:         schema.TypeString,
									Required:     true,
									ValidateFunc: validation.StringInSlice([]string{"EQUAL_TO", "NOT_EQUAL_TO"}, false),
								},
								"service_accounts": {
									Description: "List of service accounts to include/exclude",
									Type:        schema.TypeList,
									Required:    true,
									MinItems:    1,
									Elem:        &schema.Schema{Type: schema.TypeString},
								},
							},
						},
					},
				},
			},
		},
		"machine_spec": {
			Description: "Machine specific configuration (required when infra_type is Linux or Windows)",
			Type:        schema.TypeList,
			Optional:    true,
			MaxItems:    1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"infra_spec": {
						Description: "Infrastructure specification",
						Type:        schema.TypeList,
						Optional:    true,
						MaxItems:    1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"operator": {
									Description:  "Operator for comparing infrastructure IDs (EQUAL_TO or NOT_EQUAL_TO)",
									Type:         schema.TypeString,
									Required:     true,
									ValidateFunc: validation.StringInSlice([]string{"EQUAL_TO", "NOT_EQUAL_TO"}, false),
								},
								"infra_ids": {
									Description: "List of infrastructure IDs to apply the condition to",
									Type:        schema.TypeList,
									Required:    true,
									MinItems:    1,
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
	}
}
