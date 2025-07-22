package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/harness/harness-go-sdk/harness/chaos"
	"github.com/harness/harness-go-sdk/harness/chaos/graphql/model"
)

// Helper function to create string pointers
func stringPtr(s string) *string {
	return &s
}

func main() {
	// Initialize the Chaos API client
	apiKey := os.Getenv("HARNESS_API_KEY")
	accountID := os.Getenv("HARNESS_ACCOUNT_ID")
	orgID := os.Getenv("HARNESS_ORG_ID")
	projectID := os.Getenv("HARNESS_PROJECT_ID")
	infraID := os.Getenv("HARNESS_INFRASTRUCTURE_ID")

	if apiKey == "" || accountID == "" {
		log.Fatal("HARNESS_API_KEY and HARNESS_ACCOUNT_ID environment variables are required")
	}

	// Create a new API client
	cfg := &chaos.Configuration{
		ApiKey:   apiKey,
		BasePath: "https://app.harness.io/gateway/chaos/manager/api",
	}

	client := chaos.NewAPIClient(cfg)
	client.AccountId = accountID

	// Set up common identifiers
	identifiers := model.IdentifiersRequest{
		AccountIdentifier: accountID,
		OrgIdentifier:     orgID,     // Replace with your org ID
		ProjectIdentifier: projectID, // Replace with your project ID
	}

	// Example 1: Create a new condition
	fmt.Println("=== Creating a new Security Governance Condition ===")
	conditionID, err := createExample(client, identifiers, infraID)
	if err != nil {
		log.Fatalf("Failed to create condition: %v", err)
	}

	// Ensure cleanup happens even if there's an error
	defer func() {
		if conditionID != "" {
			fmt.Println("\n=== Cleaning up condition ===")
			if err := deleteExample(client, identifiers, conditionID); err != nil {
				log.Printf("Warning: Failed to delete condition: %v", err)
			} else {
				log.Println("Successfully cleaned up condition")
			}
		}
	}()

	// Wait a bit for the condition to be created
	time.Sleep(2 * time.Second)

	// Example 2: List all conditions
	listExample(client, identifiers)

	// Example 3: Get the created condition
	getExample(client, identifiers, conditionID)

	// Example 4: Update the condition
	updateExample(client, identifiers, conditionID)

	// Cleanup will happen automatically via defer
	// Wait a bit before deleting the condition
	time.Sleep(20 * time.Second)
}

// createKubernetesCondition creates a security governance condition for Kubernetes infrastructure
func createKubernetesCondition(client *chaos.APIClient, identifiers model.IdentifiersRequest, infraID string) (string, error) {
	fmt.Println("=== Creating a new Kubernetes Security Governance Condition ===")

	// Validate infrastructure IDs
	if infraID == "" {
		return "", fmt.Errorf("HARNESS_INFRASTRUCTURE_ID environment variable is required")
	}

	// Split the infraID string to support multiple infra IDs
	infraIDs := []string{infraID}
	// Add a second infrastructure ID for demonstration (you can add more as needed)
	if os.Getenv("HARNESS_INFRASTRUCTURE_ID_2") != "" {
		infraIDs = append(infraIDs, os.Getenv("HARNESS_INFRASTRUCTURE_ID_2"))
	}

	// Define a new condition
	conditionID := fmt.Sprintf("k8s-condition-%d", time.Now().Unix())
	description := "Kubernetes security governance condition with multiple workloads"

	request := model.ConditionRequest{
		ConditionID: conditionID,
		Name:        "Kubernetes Security Condition",
		Description: &description,
		InfraType:   model.InfrastructureTypeKubernetesV2,
		FaultSpec: &model.FaultSpecInput{
			Operator: model.OperatorEqualTo,
			Faults: []*model.Fault{
				{
					FaultType: model.FaultTypeFault,
					Name:      "pod-delete",
				},
				{
					FaultType: model.FaultTypeFault,
					Name:      "pod-cpu-hog",
				},
				{
					FaultType: model.FaultTypeFault,
					Name:      "pod-io-stress",
				},
			},
		},
		K8sSpec: &model.K8sSpecInput{
			InfraSpec: &model.InfraSpecInput{
				Operator: model.OperatorEqualTo,
				InfraIds: infraIDs,
			},
			ApplicationSpec: &model.ApplicationSpecInput{
				Operator: model.OperatorNotEqualTo,
				Workloads: []*model.WorkloadInput{
					{
						Namespace:        "boutique",
						Kind:             stringPtr("statefulset"),
						Label:            stringPtr("app=boutique"),
						Services:         []string{"adservice"},
						ApplicationMapID: stringPtr("boutique"),
						Env: []*model.EnvInput{
							{
								Name:  "ENV",
								Value: "production",
							},
						},
					},
					{
						Namespace:        "kube-system",
						Kind:             stringPtr("deployment"),
						Label:            stringPtr("k8s-app=metrics-server"),
						Services:         []string{"metrics-server"},
						ApplicationMapID: stringPtr("metrics-server-application"),
						Env: []*model.EnvInput{
							{
								Name:  "METRICS_SERVER_NAMESPACE",
								Value: "kube-system",
							},
						},
					},
				},
			},
			ChaosServiceAccountSpec: &model.ChaosServiceAccountSpecInput{
				Operator: model.OperatorEqualTo,
				ServiceAccounts: []string{
					"default",
					"chaos-service-account",
				},
			},
		},
	}

	// Create the condition
	response, err := client.SecurityGovernanceConditionApi.Create(
		context.Background(),
		identifiers,
		request,
	)

	if err != nil {
		return "", fmt.Errorf("failed to create Kubernetes condition: %w", err)
	}

	prettyPrint("Kubernetes condition created successfully:", response)
	return conditionID, nil
}

// createWindowsCondition creates a security governance condition for Windows infrastructure
func createWindowsCondition(client *chaos.APIClient, identifiers model.IdentifiersRequest, infraID string) (string, error) {
	fmt.Println("=== Creating a new Windows Security Governance Condition ===")

	if infraID == "" {
		return "", fmt.Errorf("HARNESS_WINDOWS_INFRASTRUCTURE_ID environment variable is required")
	}

	conditionID := fmt.Sprintf("win-condition-%d", time.Now().Unix())
	description := "Windows security governance condition for service faults"

	request := model.ConditionRequest{
		ConditionID: conditionID,
		Name:        "Windows Security Condition",
		Description: &description,
		InfraType:   model.InfrastructureTypeWindows,
		FaultSpec: &model.FaultSpecInput{
			Operator: model.OperatorEqualTo,
			Faults: []*model.Fault{
				{
					FaultType: model.FaultTypeFault,
					Name:      "windows-service-stop",
				},
				{
					FaultType: model.FaultTypeFault,
					Name:      "windows-cpu-hog",
				},
			},
		},
		MachineSpec: &model.MachineSpecInput{
			InfraSpec: &model.InfraSpecInput{
				Operator: model.OperatorEqualTo,
				InfraIds: []string{infraID},
			},
		},
	}

	response, err := client.SecurityGovernanceConditionApi.Create(
		context.Background(),
		identifiers,
		request,
	)

	if err != nil {
		return "", fmt.Errorf("failed to create Windows condition: %w", err)
	}

	prettyPrint("Windows condition created successfully:", response)
	return conditionID, nil
}

// createLinuxCondition creates a security governance condition for Linux infrastructure
func createLinuxCondition(client *chaos.APIClient, identifiers model.IdentifiersRequest, infraID string) (string, error) {
	fmt.Println("=== Creating a new Linux Security Governance Condition ===")

	if infraID == "" {
		return "", fmt.Errorf("HARNESS_LINUX_INFRASTRUCTURE_ID environment variable is required")
	}

	conditionID := fmt.Sprintf("linux-condition-%d", time.Now().Unix())
	description := "Linux security governance condition for process and network faults"

	request := model.ConditionRequest{
		ConditionID: conditionID,
		Name:        "Linux Security Condition",
		Description: &description,
		InfraType:   model.InfrastructureTypeLinux,
		FaultSpec: &model.FaultSpecInput{
			Operator: model.OperatorEqualTo,
			Faults: []*model.Fault{
				{
					FaultType: model.FaultTypeFault,
					Name:      "process-kill",
				},
				{
					FaultType: model.FaultTypeFault,
					Name:      "network-loss",
				},
			},
		},
		MachineSpec: &model.MachineSpecInput{
			InfraSpec: &model.InfraSpecInput{
				Operator: model.OperatorEqualTo,
				InfraIds: []string{infraID},
			},
		},
	}

	response, err := client.SecurityGovernanceConditionApi.Create(
		context.Background(),
		identifiers,
		request,
	)

	if err != nil {
		return "", fmt.Errorf("failed to create Linux condition: %w", err)
	}

	prettyPrint("Linux condition created successfully:", response)
	return conditionID, nil
}

// createExample creates a condition based on the infrastructure type
func createExample(client *chaos.APIClient, identifiers model.IdentifiersRequest, infraID string) (string, error) {
	// Default to Kubernetes if no specific type is specified
	infraType := os.Getenv("HARNESS_INFRASTRUCTURE_TYPE")
	if infraType == "" {
		infraType = "kubernetes"
	}

	switch strings.ToLower(infraType) {
	case "windows":
		return createWindowsCondition(client, identifiers, infraID)
	case "linux":
		return createLinuxCondition(client, identifiers, infraID)
	case "kubernetes":
		fallthrough
	default:
		return createKubernetesCondition(client, identifiers, infraID)
	}
}

func listExample(client *chaos.APIClient, identifiers model.IdentifiersRequest) {
	fmt.Println("\n=== Listing Security Governance Conditions ===")

	// List all conditions with pagination
	page := 0
	pageSize := 10
	var allConditions []*model.ConditionResponse

	for {
		listRequest := model.ListConditionRequest{
			Pagination: &model.Pagination{
				Page:  page,
				Limit: pageSize,
			},
		}

		response, err := client.SecurityGovernanceConditionApi.List(
			context.Background(),
			identifiers,
			listRequest,
		)

		if err != nil {
			log.Printf("Failed to list conditions: %v", err)
			return
		}

		allConditions = append(allConditions, response.Conditions...)

		// Check if we've retrieved all conditions
		if len(response.Conditions) < pageSize {
			break
		}

		page++
	}

	prettyPrint(fmt.Sprintf("Found %d conditions:", len(allConditions)), allConditions)
}

func getExample(client *chaos.APIClient, identifiers model.IdentifiersRequest, conditionID string) {
	fmt.Printf("\n=== Getting Security Governance Condition %s ===\n", conditionID)

	condition, err := client.SecurityGovernanceConditionApi.Get(
		context.Background(),
		identifiers,
		conditionID,
	)

	if err != nil {
		log.Printf("Failed to get condition: %v", err)
		return
	}

	prettyPrint("Condition details:", condition)
}

func updateExample(client *chaos.APIClient, identifiers model.IdentifiersRequest, conditionID string) {
	fmt.Printf("\n=== Updating Security Governance Condition %s ===\n", conditionID)

	// First, get the existing condition
	condition, err := client.SecurityGovernanceConditionApi.Get(
		context.Background(),
		identifiers,
		conditionID,
	)

	if err != nil {
		log.Printf("Failed to get condition for update: %v", err)
		return
	}

	description := "This condition was updated at " + time.Now().Format(time.RFC3339)

	// Create the update request with basic fields
	updateRequest := model.ConditionRequest{
		ConditionID: conditionID,
		Name:        condition.Condition.Name + " (Updated)",
		Description: &description,
		InfraType:   condition.Condition.InfraType,
	}

	// Copy fault spec if it exists
	if condition.Condition.FaultSpec != nil {
		faults := make([]*model.Fault, len(condition.Condition.FaultSpec.Faults))
		for i, f := range condition.Condition.FaultSpec.Faults {
			faults[i] = &model.Fault{
				FaultType: f.FaultType,
				Name:      f.Name,
			}
		}
		updateRequest.FaultSpec = &model.FaultSpecInput{
			Operator: condition.Condition.FaultSpec.Operator,
			Faults:   faults,
		}
	}

	// Copy K8s spec if it exists
	if condition.Condition.K8sSpec != nil {
		k8sSpec := &model.K8sSpecInput{
			InfraSpec: &model.InfraSpecInput{
				Operator: condition.Condition.K8sSpec.InfraSpec.Operator,
				InfraIds: append([]string{}, condition.Condition.K8sSpec.InfraSpec.InfraIds...),
			},
		}

		// Copy application spec if it exists
		if condition.Condition.K8sSpec.ApplicationSpec != nil {
			workloads := make([]*model.WorkloadInput, len(condition.Condition.K8sSpec.ApplicationSpec.Workloads))
			for i, w := range condition.Condition.K8sSpec.ApplicationSpec.Workloads {
				envs := make([]*model.EnvInput, len(w.Env))
				for j, e := range w.Env {
					envs[j] = &model.EnvInput{
						Name:  e.Name,
						Value: e.Value,
					}
				}
				workloads[i] = &model.WorkloadInput{
					Label:            w.Label,
					Namespace:        w.Namespace,
					Kind:             w.Kind,
					Services:         append([]string{}, w.Services...),
					ApplicationMapID: w.ApplicationMapID,
					Env:              envs,
				}
			}

			k8sSpec.ApplicationSpec = &model.ApplicationSpecInput{
				Operator:  condition.Condition.K8sSpec.ApplicationSpec.Operator,
				Workloads: workloads,
			}
		}

		// Copy chaos service account spec if it exists
		if condition.Condition.K8sSpec.ChaosServiceAccountSpec != nil {
			k8sSpec.ChaosServiceAccountSpec = &model.ChaosServiceAccountSpecInput{
				Operator:        condition.Condition.K8sSpec.ChaosServiceAccountSpec.Operator,
				ServiceAccounts: append([]string{}, condition.Condition.K8sSpec.ChaosServiceAccountSpec.ServiceAccounts...),
			}
		}

		updateRequest.K8sSpec = k8sSpec
	}

	// Copy machine spec if it exists
	if condition.Condition.MachineSpec != nil {
		updateRequest.MachineSpec = &model.MachineSpecInput{
			InfraSpec: &model.InfraSpecInput{
				Operator: condition.Condition.MachineSpec.InfraSpec.Operator,
				InfraIds: append([]string{}, condition.Condition.MachineSpec.InfraSpec.InfraIds...),
			},
		}
	}

	// Add or update tags
	var tags []*string
	for _, t := range condition.Condition.Tags {
		tag := t // Create a new variable to avoid taking the address of the loop variable
		tags = append(tags, &tag)
	}
	updateRequest.Tags = tags

	response, err := client.SecurityGovernanceConditionApi.Update(
		context.Background(),
		identifiers,
		updateRequest,
	)

	if err != nil {
		log.Printf("Failed to update condition: %v", err)
		return
	}

	prettyPrint("Condition updated successfully:", response)
}

func deleteExample(client *chaos.APIClient, identifiers model.IdentifiersRequest, conditionID string) error {
	fmt.Printf("\n=== Deleting Security Governance Condition %s ===\n", conditionID)

	_, err := client.SecurityGovernanceConditionApi.Delete(
		context.Background(),
		identifiers,
		conditionID,
	)

	if err != nil {
		return fmt.Errorf("failed to delete condition: %w", err)
	}

	fmt.Printf("Condition %s deleted successfully\n", conditionID)
	return nil
}

// Helper function to pretty print JSON output
func prettyPrint(prefix string, v interface{}) {
	fmt.Println(prefix)
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Printf("Error marshaling to JSON: %v", err)
		return
	}
	fmt.Println(string(b))
}
