package security_governance

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/harness/harness-go-sdk/harness/chaos"
	"github.com/harness/harness-go-sdk/harness/chaos/graphql/model"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

const (
	resourceName = "harness_chaos_security_governance_condition"
)

func ResourceChaosSecurityGovernanceCondition() *schema.Resource {
	return &schema.Resource{
		Description:   "Resource for managing a Harness Chaos Security Governance Condition",
		CreateContext: resourceChaosSecurityGovernanceConditionCreate,
		ReadContext:   resourceChaosSecurityGovernanceConditionRead,
		UpdateContext: resourceChaosSecurityGovernanceConditionUpdate,
		DeleteContext: resourceChaosSecurityGovernanceConditionDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceChaosSecurityGovernanceConditionImport,
		},

		Schema: map[string]*schema.Schema{
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
							Description: "Operator for comparing faults (EQUAL_TO or NOT_EQUAL_TO)",
							Type:        schema.TypeString,
							Required:    true,
							ValidateFunc: validation.StringInSlice(
								[]string{
									string(model.OperatorEqualTo),
									string(model.OperatorNotEqualTo),
								},
								false,
							),
						},
						"faults": {
							Description: "List of fault specifications",
							Type:        schema.TypeList,
							Required:    true,
							MinItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"fault_type": {
										Description: "Type of the fault (FAULT or FAULT_GROUP)",
										Type:        schema.TypeString,
										Required:    true,
										ValidateFunc: validation.StringInSlice(
											[]string{
												string(model.FaultTypeFault),
												string(model.FaultTypeFaultGroup),
											},
											false,
										),
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
				Description: "Kubernetes specific configuration (required when infra_type is KUBERNETES or KUBERNETESV2)",
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
										Description: "Operator for comparing infrastructure IDs (EQUAL_TO or NOT_EQUAL_TO)",
										Type:        schema.TypeString,
										Required:    true,
										ValidateFunc: validation.StringInSlice(
											[]string{
												string(model.OperatorEqualTo),
												string(model.OperatorNotEqualTo),
											},
											false,
										),
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
										Description: "Operator for application matching (EQUAL_TO or NOT_EQUAL_TO)",
										Type:        schema.TypeString,
										Required:    true,
										ValidateFunc: validation.StringInSlice(
											[]string{
												string(model.OperatorEqualTo),
												string(model.OperatorNotEqualTo),
											},
											false,
										),
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
										Description: "Operator for service account matching (EQUAL_TO or NOT_EQUAL_TO)",
										Type:        schema.TypeString,
										Required:    true,
										ValidateFunc: validation.StringInSlice(
											[]string{
												string(model.OperatorEqualTo),
												string(model.OperatorNotEqualTo),
											},
											false,
										),
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
				Description: "Machine specific configuration (required when infra_type is LINUX or WINDOWS)",
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
										Description: "Operator for comparing infrastructure IDs (EQUAL_TO or NOT_EQUAL_TO)",
										Type:        schema.TypeString,
										Required:    true,
										ValidateFunc: validation.StringInSlice(
											[]string{
												string(model.OperatorEqualTo),
												string(model.OperatorNotEqualTo),
											},
											false,
										),
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
		},
	}
}

// Import formats:
// - <condition_id>
// - <org_id>/<project_id>/<condition_id>
// ScopedIdentifiersRequest represents the identifiers for a scoped resource
type ScopedIdentifiersRequest struct {
	AccountIdentifier string  `json:"accountIdentifier"`
	OrgIdentifier     *string `json:"orgIdentifier,omitempty"`
	ProjectIdentifier *string `json:"projectIdentifier,omitempty"`
}

func resourceChaosSecurityGovernanceConditionImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	c := meta.(*internal.Session).ChaosClient

	// Parse the import ID which can be in one of these formats:
	// 1. Account level: "condition_id"
	// 2. Org level: "org_id/condition_id"
	// 3. Project level: "org_id/project_id/condition_id"
	importID := d.Id()
	parts := strings.Split(importID, "/")

	var conditionID, orgID, projectID string

	switch len(parts) {
	case 1:
		// Account level: "condition_id"
		conditionID = parts[0]
	case 2:
		// Org level: "org-id/condition-name"
		orgID = parts[0]
		conditionID = parts[1]
	case 3:
		// Project level: "org-id/project-id/condition-name"
		orgID = parts[0]
		projectID = parts[1]
		conditionID = parts[2]
	default:
		return nil, fmt.Errorf("invalid import ID format. Expected \"<condition-id>\", \"<org-id>/<condition-id>\", or \"<org-id>/<project-id>/<condition-id>\"")
	}

	if conditionID == "" {
		return nil, fmt.Errorf("condition id cannot be empty")
	}

	// Create a client for the Security Governance Condition API
	client := chaos.NewSecurityGovernanceConditionClient(c)

	// Get the account ID from the provider config
	accountID := c.AccountId
	if accountID == "" {
		return nil, fmt.Errorf("account ID must be configured in the provider")
	}

	// Create identifiers for the request
	identifiers := model.IdentifiersRequest{
		AccountIdentifier: accountID,
	}

	// Set org and project identifiers if they exist
	if orgID != "" {
		identifiers.OrgIdentifier = orgID
	}
	if projectID != "" {
		identifiers.ProjectIdentifier = projectID
	}

	log.Printf("[DEBUG] Importing condition with id: %s, org: %s, project: %s", conditionID, orgID, projectID)

	// Get the condition by ID
	condition, err := client.Get(ctx, identifiers, conditionID)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			errMsg := fmt.Sprintf("no security governance condition found with id: %s", conditionID)
			if orgID != "" || projectID != "" {
				errMsg = fmt.Sprintf("%s in the specified scope (org: %s, project: %s)",
					errMsg, orgID, projectID)
			}
			return nil, fmt.Errorf("%s", errMsg)
		}
		return nil, fmt.Errorf("failed to get security governance condition: %v", err)
	}

	// Create a ScopedIdentifiersRequest for the condition
	scopedIdentifiers := ScopedIdentifiersRequest{
		AccountIdentifier: accountID,
	}

	// Set org and project identifiers if they exist
	if orgID != "" {
		scopedIdentifiers.OrgIdentifier = &orgID
	}
	if projectID != "" {
		scopedIdentifiers.ProjectIdentifier = &projectID
	}

	// Set the resource ID using the condition's ID and scope information
	d.SetId(conditionID)

	// Set the resource attributes
	d.Set("name", condition.Condition.Name)
	if condition.Condition.Description != nil {
		d.Set("description", *condition.Condition.Description)
	}
	d.Set("org_id", orgID)
	d.Set("project_id", projectID)

	// Set fault spec if it exists
	if condition.Condition.FaultSpec != nil {
		faultSpec := map[string]interface{}{
			"operator": string(condition.Condition.FaultSpec.Operator),
		}

		// Convert faults to the expected format
		if condition.Condition.FaultSpec.Faults != nil {
			faults := make([]map[string]string, len(condition.Condition.FaultSpec.Faults))
			for i, f := range condition.Condition.FaultSpec.Faults {
				faults[i] = map[string]string{
					"fault_type": string(f.FaultType),
					"name":       f.Name,
				}
			}
			faultSpec["faults"] = faults
		}

		d.Set("fault_spec", []interface{}{faultSpec})
	}

	// Set infrastructure type
	d.Set("infra_type", string(condition.Condition.InfraType))

	// Set tags if they exist
	if len(condition.Condition.Tags) > 0 {
		tags := make([]interface{}, len(condition.Condition.Tags))
		for i, tag := range condition.Condition.Tags {
			tags[i] = tag
		}
		d.Set("tags", tags)
	}

	return []*schema.ResourceData{d}, nil
}

// generateID creates a unique ID for the resource that includes scope information
func generateID(identifiers ScopedIdentifiersRequest, conditionID string) string {
	parts := []string{identifiers.AccountIdentifier}

	if identifiers.OrgIdentifier != nil && *identifiers.OrgIdentifier != "" {
		parts = append(parts, *identifiers.OrgIdentifier)
		if identifiers.ProjectIdentifier != nil && *identifiers.ProjectIdentifier != "" {
			parts = append(parts, *identifiers.ProjectIdentifier)
		}
	}

	parts = append(parts, conditionID)
	return strings.Join(parts, "/")
}

// parseID parses a resource ID into its components
func parseID(id string) (ScopedIdentifiersRequest, string, error) {
	log.Printf("[DEBUG] Parsing ID: %s", id)

	// Handle the format: org/project/condition-id
	parts := strings.Split(id, "/")

	switch len(parts) {
	case 3: // org/project/condition-id
		result := ScopedIdentifiersRequest{
			OrgIdentifier:     &parts[0],
			ProjectIdentifier: &parts[1],
		}
		return result, parts[2], nil

	case 2: // org/condition-id
		result := ScopedIdentifiersRequest{
			OrgIdentifier: &parts[0],
		}
		return result, parts[1], nil

	case 1: // condition-id
		result := ScopedIdentifiersRequest{}
		return result, parts[0], nil

	default:
		return ScopedIdentifiersRequest{}, "", fmt.Errorf("invalid ID format, expected org/project/condition-id, org/condition-id, or condition-id")
	}
}

// getIdentifiers extracts the identifiers from the resource data
func getIdentifiers(d *schema.ResourceData, accountID string) ScopedIdentifiersRequest {
	identifiers := ScopedIdentifiersRequest{
		AccountIdentifier: accountID,
	}

	if orgID, ok := d.Get("org_id").(string); ok && orgID != "" {
		identifiers.OrgIdentifier = &orgID

		if projectID, ok := d.Get("project_id").(string); ok && projectID != "" {
			identifiers.ProjectIdentifier = &projectID
		}
	}

	return identifiers
}
