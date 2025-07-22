// Package main provides an example of how to use the Harness Chaos SDK to manage
// security governance rules. It demonstrates creating, listing, updating, and deleting
// security governance rules and their associated conditions.
package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/harness/harness-go-sdk/harness/chaos"
	"github.com/harness/harness-go-sdk/harness/chaos/graphql/model"
)

// Log levels for consistent logging
const (
	logLevelDebug = "DEBUG"
	logLevelInfo  = "INFO"
	logLevelWarn  = "WARN"
	logLevelError = "ERROR"
)

// Default values for configuration
const (
	defaultUserGroupID     = "all_users"
	defaultTimeZone        = "UTC"
	defaultTimeout         = 5 * time.Minute
	defaultBaseURL         = "https://app.harness.io/gateway/chaos/manager/api"
	defaultWaitBeforeClean = 60 * time.Second
)

// Error types for better error handling
var (
	ErrInvalidConfig = errors.New("invalid configuration")
	ErrAPI           = errors.New("API error")
)

// Config holds the application configuration
type Config struct {
	APIKey           string
	AccountID        string
	OrgID            string
	ProjectID        string
	InfrastructureID string
	UserGroupID      string
	BaseURL          string
	Timeout          time.Duration
	WaitBeforeClean  time.Duration
}

// Logger provides structured logging capabilities
type Logger struct {
	mu     sync.Mutex
	silent bool
}

// logMessage logs a message with the specified log level
func (l *Logger) log(level, format string, args ...interface{}) {
	if l.silent {
		return
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	timestamp := time.Now().Format(time.RFC3339)
	message := fmt.Sprintf(format, args...)
	log.Printf("[%s] [%s] %s", timestamp, level, message)
}

// Debug logs a debug message
func (l *Logger) Debug(format string, args ...interface{}) {
	l.log(logLevelDebug, format, args...)
}

// Info logs an informational message
func (l *Logger) Info(format string, args ...interface{}) {
	l.log(logLevelInfo, format, args...)
}

// Warn logs a warning message
func (l *Logger) Warn(format string, args ...interface{}) {
	l.log(logLevelWarn, format, args...)
}

// Error logs an error message
func (l *Logger) Error(format string, args ...interface{}) {
	l.log(logLevelError, format, args...)
}

var logger = &Logger{}

// loadConfig loads the application configuration from environment variables
func loadConfig() (*Config, error) {
	cfg := &Config{
		APIKey:           os.Getenv("HARNESS_API_KEY"),
		AccountID:        os.Getenv("HARNESS_ACCOUNT_ID"),
		OrgID:            os.Getenv("HARNESS_ORG_ID"),
		ProjectID:        os.Getenv("HARNESS_PROJECT_ID"),
		InfrastructureID: os.Getenv("HARNESS_INFRASTRUCTURE_ID"),
		UserGroupID:      os.Getenv("HARNESS_USER_GROUP_ID"),
		BaseURL:          defaultBaseURL,
		Timeout:          defaultTimeout,
		WaitBeforeClean:  defaultWaitBeforeClean,
	}

	// Validate required fields
	if cfg.APIKey == "" || cfg.AccountID == "" {
		return nil, fmt.Errorf("%w: HARNESS_API_KEY and HARNESS_ACCOUNT_ID are required", ErrInvalidConfig)
	}

	// Set defaults
	if cfg.UserGroupID == "" {
		cfg.UserGroupID = defaultUserGroupID
		logger.Warn("HARNESS_USER_GROUP_ID not set, using default '%s'", defaultUserGroupID)
	}

	return cfg, nil
}

// createChaosClient creates a new Chaos API client with the given configuration
func createChaosClient(cfg *Config) (*chaos.APIClient, error) {
	if cfg == nil {
		return nil, fmt.Errorf("%w: config cannot be nil", ErrInvalidConfig)
	}

	clientCfg := &chaos.Configuration{
		ApiKey:   cfg.APIKey,
		BasePath: cfg.BaseURL,
	}

	client := chaos.NewAPIClient(clientCfg)
	client.AccountId = cfg.AccountID

	return client, nil
}

// createIdentifiers creates the identifiers request from the configuration
func createIdentifiers(cfg *Config) model.IdentifiersRequest {
	identifiers := model.IdentifiersRequest{
		AccountIdentifier: cfg.AccountID,
	}

	// Set optional fields if they are not empty
	if cfg.OrgID != "" {
		orgID := cfg.OrgID
		identifiers.OrgIdentifier = orgID
	}

	if cfg.ProjectID != "" {
		projectID := cfg.ProjectID
		identifiers.ProjectIdentifier = projectID
	}

	return identifiers
}

// createExampleRule creates a new security governance rule
func createExampleRule(ctx context.Context, client *chaos.APIClient,
	identifiers model.IdentifiersRequest, conditionID string, cfg *Config) (string, error) {
	logger.Info("Creating a new security governance rule")

	// Create a unique rule ID
	ruleID := fmt.Sprintf("example-rule-%d", time.Now().Unix())
	desc := "This is an example security governance rule"

	// Create a time window for today, starting now and ending in 1 hour
	now := time.Now()
	startTime := int(now.Unix() * 1000)                  // Convert to milliseconds
	endTime := int(now.Add(1*time.Hour).Unix() * 1000)   // Convert to milliseconds
	untilTime := int(now.AddDate(0, 1, 0).Unix() * 1000) // Recur until 1 month from now
	duration := "60m"                                    // Duration in minutes

	timeWindows := []*model.TimeWindowInput{
		{
			TimeZone:  "UTC",
			StartTime: startTime,
			EndTime:   &endTime,
			Duration:  &duration,
			Recurrence: &model.RecurrenceInput{
				Type: model.RecurrenceTypeDaily,
				Spec: &model.RecurrenceSpecInput{
					Until: &untilTime,
					Value: nil, // Not all APIs require this
				},
			},
		},
	}

	req := model.RuleInput{
		Name:         "Example Security Rule",
		Description:  &desc,
		IsEnabled:    true,
		UserGroupIds: []string{cfg.UserGroupID},
		ConditionIds: []string{conditionID},
		RuleID:       ruleID,
		TimeWindows:  timeWindows,
	}

	// Log the request for debugging
	reqJSON, _ := json.MarshalIndent(req, "", "  ")
	logger.Debug("Creating rule with request: %s", string(reqJSON))

	// Create the rule
	_, err := client.SecurityGovernanceRuleApi.Create(
		ctx,
		identifiers,
		req,
	)

	if err != nil {
		logger.Error("Failed to create rule: %v", err)
		return "", fmt.Errorf("%w: failed to create rule: %v", ErrAPI, err)
	}

	logger.Info("Successfully created rule with ID: %s", ruleID)
	return ruleID, nil
}

// main is the entry point of the application
func main() {
	logger.Info("Starting security governance rule example")

	// Load configuration
	cfg, err := loadConfig()
	if err != nil {
		logger.Error("Failed to load configuration: %v", err)
		os.Exit(1)
	}

	// Set up context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), cfg.Timeout)
	defer cancel()

	// Create Chaos client
	client, err := createChaosClient(cfg)
	if err != nil {
		logger.Error("Failed to create Chaos client: %v", err)
		os.Exit(1)
	}

	// Create identifiers
	identifiers := createIdentifiers(cfg)

	// Run the example
	if err := runExample(ctx, client, identifiers, cfg); err != nil {
		logger.Error("Example failed: %v", err)
		os.Exit(1)
	}

	logger.Info("Example completed successfully")
}

// runExample executes the main example logic
func runExample(ctx context.Context, client *chaos.APIClient, identifiers model.IdentifiersRequest, cfg *Config) error {
	// 1. Create condition
	conditionID, err := createExampleCondition(ctx, client, identifiers, cfg)
	if err != nil {
		return fmt.Errorf("failed to create condition: %w", err)
	}
	defer func() {
		if conditionID != "" {
			logger.Info("Cleaning up condition: %s", conditionID)
			if err := deleteExampleCondition(ctx, client, identifiers, conditionID); err != nil {
				logger.Error("Failed to clean up condition %s: %v", conditionID, err)
			}
		}
	}()

	// 2. Create rule
	ruleID, err := createExampleRule(ctx, client, identifiers, conditionID, cfg)
	if err != nil {
		return fmt.Errorf("failed to create rule: %w", err)
	}

	// 3. List rules
	if _, err := listExampleRules(ctx, client, identifiers); err != nil {
		logger.Warn("Failed to list rules: %v", err)
	}

	// 4. Get rule details
	if _, err := getExampleRule(ctx, client, identifiers, ruleID); err != nil {
		logger.Warn("Failed to get rule details: %v", err)
	}

	// 5. Update rule
	if err := updateExampleRule(ctx, client, identifiers, ruleID, conditionID); err != nil {
		logger.Warn("Failed to update rule: %v", err)
	}

	// 6. Toggle rule state
	if err := tuneExampleRule(ctx, client, identifiers, ruleID, false); err != nil {
		logger.Warn("Failed to disable rule: %v", err)
	}

	// Wait before re-enabling
	select {
	case <-time.After(2 * time.Second):
	case <-ctx.Done():
		return ctx.Err()
	}

	if err := tuneExampleRule(ctx, client, identifiers, ruleID, true); err != nil {
		return fmt.Errorf("failed to re-enable rule: %w", err)
	}

	// 7. Wait before cleanup
	logger.Info("Waiting %v before cleanup...", cfg.WaitBeforeClean)
	select {
	case <-time.After(cfg.WaitBeforeClean):
	case <-ctx.Done():
		return ctx.Err()
	}

	// 8. Delete rule
	if err := deleteExampleRule(ctx, client, identifiers, ruleID); err != nil {
		return fmt.Errorf("failed to delete rule: %w", err)
	}

	// 9. Verify deletion
	rules, err := listExampleRules(ctx, client, identifiers)
	if err != nil {
		return fmt.Errorf("failed to verify rule deletion: %w", err)
	}

	ruleExists := false
	for _, r := range rules {
		if r.Rule.RuleID == ruleID {
			ruleExists = true
			break
		}
	}

	if ruleExists {
		logger.Warn("Rule may not have been deleted properly")
	} else {
		logger.Info("✓ Rule successfully deleted")
	}

	return nil
}

// createExampleCondition creates a new security governance condition
// It returns the created condition ID and any error encountered
func createExampleCondition(ctx context.Context, client *chaos.APIClient,
	identifiers model.IdentifiersRequest, cfg *Config) (string, error) {
	logger.Info("Creating a new condition")

	if client == nil {
		return "", fmt.Errorf("%w: client cannot be nil", ErrInvalidConfig)
	}

	if cfg == nil {
		return "", fmt.Errorf("%w: config cannot be nil", ErrInvalidConfig)
	}

	// Create a unique condition ID
	conditionID := fmt.Sprintf("example-condition-%d", time.Now().Unix())
	desc := "This condition is used in the rule example"

	// Create the condition request with all required fields
	req := model.ConditionRequest{
		ConditionID: conditionID,
		Name:        "Example Condition for Rule",
		Description: &desc,
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
					Name:      "pod-dns",
				},
				{
					FaultType: model.FaultTypeFault,
					Name:      "pod-kill",
				},
			},
		},
		K8sSpec: &model.K8sSpecInput{
			InfraSpec: &model.InfraSpecInput{
				Operator: model.OperatorEqualTo,
				InfraIds: []string{cfg.InfrastructureID},
			},
			ApplicationSpec: &model.ApplicationSpecInput{
				Operator: model.OperatorEqualTo,
				Workloads: []*model.WorkloadInput{
					{
						Namespace:        "boutique",
						Kind:             stringPtr("deployment"),
						Label:            stringPtr("app=adservice"),
						ApplicationMapID: stringPtr("boutique"),
						Services:         []string{"adservice"},
						Env: []*model.EnvInput{
							{
								Name:  "ENV",
								Value: "test",
							},
							{
								Name:  "ENV1",
								Value: "test1",
							},
						},
					},
					{
						Namespace:        "boutique",
						Kind:             stringPtr("deployment"),
						Label:            stringPtr("app=cartservice"),
						ApplicationMapID: stringPtr("boutique"),
						Services:         []string{"cartservice"},
						Env: []*model.EnvInput{
							{
								Name:  "ENV2",
								Value: "test2",
							},
							{
								Name:  "ENV21",
								Value: "test21",
							},
						},
					},
				},
			},
			ChaosServiceAccountSpec: &model.ChaosServiceAccountSpecInput{
				Operator:        model.OperatorEqualTo,
				ServiceAccounts: []string{"default", "boutique"},
			},
		},
	}

	// Log the request for debugging
	reqJSON, _ := json.MarshalIndent(req, "", "  ")
	logger.Debug("Creating condition with request: %s", string(reqJSON))

	// Make the API call with the provided context
	response, err := client.SecurityGovernanceConditionApi.Create(
		ctx,
		identifiers,
		req,
	)

	// Handle errors
	if err != nil {
		logger.Error("Failed to create condition: %v", err)
		return "", fmt.Errorf("%w: %v", ErrAPI, err)
	}

	// Log success
	if response != nil {
		respBytes, _ := json.MarshalIndent(response, "", "  ")
		logger.Debug("Condition created successfully. Response: %s", string(respBytes))
	} else {
		logger.Info("Condition created successfully, but received empty response")
	}

	return conditionID, nil
}

// deleteExampleCondition deletes a security governance condition by ID
func deleteExampleCondition(ctx context.Context, client *chaos.APIClient,
	identifiers model.IdentifiersRequest, conditionID string) error {
	if client == nil {
		return fmt.Errorf("%w: client cannot be nil", ErrInvalidConfig)
	}

	if conditionID == "" {
		return fmt.Errorf("%w: conditionID cannot be empty", ErrInvalidConfig)
	}

	logger.Info("Deleting condition: %s", conditionID)

	_, err := client.SecurityGovernanceConditionApi.Delete(
		ctx,
		identifiers,
		conditionID,
	)

	if err != nil {
		logger.Error("Failed to delete condition %s: %v", conditionID, err)
		return fmt.Errorf("%w: failed to delete condition: %v", ErrAPI, err)
	}

	logger.Info("Successfully deleted condition: %s", conditionID)
	return nil
}

// listExampleRules lists all security governance rules
// It returns a slice of RuleResponse and any error encountered
func listExampleRules(ctx context.Context, client *chaos.APIClient,
	identifiers model.IdentifiersRequest) ([]*model.RuleResponse, error) {
	logger.Info("Listing all security governance rules")

	if client == nil {
		return nil, fmt.Errorf("%w: client cannot be nil", ErrInvalidConfig)
	}

	// Create a list request with pagination
	listRequest := model.ListRuleRequest{
		Pagination: &model.Pagination{
			Page:  0,
			Limit: 100, // Adjust the limit as needed
		},
	}

	// Get the list of rules
	rules, err := client.SecurityGovernanceRuleApi.List(
		ctx,
		identifiers,
		listRequest,
	)

	if err != nil {
		logger.Error("Failed to list rules: %v", err)
		return nil, fmt.Errorf("%w: failed to list rules: %v", ErrAPI, err)
	}

	logger.Info("Successfully listed %d rules", len(rules))
	return rules, nil
}

// getExampleRule retrieves a security governance rule by ID
// It returns the rule response and any error encountered
func getExampleRule(ctx context.Context, client *chaos.APIClient,
	identifiers model.IdentifiersRequest, ruleID string) (*model.RuleResponse, error) {
	logger.Info("Getting rule with ID: %s", ruleID)

	if client == nil {
		return nil, fmt.Errorf("%w: client cannot be nil", ErrInvalidConfig)
	}

	if ruleID == "" {
		return nil, fmt.Errorf("%w: ruleID cannot be empty", ErrInvalidConfig)
	}

	// Get the rule
	rule, err := client.SecurityGovernanceRuleApi.Get(
		ctx,
		identifiers,
		ruleID,
	)

	if err != nil {
		logger.Error("Failed to get rule %s: %v", ruleID, err)
		return nil, fmt.Errorf("%w: failed to get rule: %v", ErrAPI, err)
	}

	// Log the response for debugging
	ruleJSON, _ := json.MarshalIndent(rule, "", "  ")
	logger.Debug("Retrieved rule: %s", string(ruleJSON))

	logger.Info("Successfully retrieved rule: %s", ruleID)
	return rule, nil
}

// updateExampleRule updates an existing security governance rule
func updateExampleRule(ctx context.Context, client *chaos.APIClient,
	identifiers model.IdentifiersRequest, ruleID string, conditionID string) error {
	logger.Info("Updating rule with ID: %s", ruleID)

	if client == nil {
		return fmt.Errorf("%w: client cannot be nil", ErrInvalidConfig)
	}

	if ruleID == "" {
		return fmt.Errorf("%w: ruleID cannot be empty", ErrInvalidConfig)
	}

	// First, get the current rule to update it
	rule, err := getExampleRule(ctx, client, identifiers, ruleID)
	if err != nil {
		return fmt.Errorf("failed to get rule for update: %w", err)
	}

	// Update the description
	updatedDesc := fmt.Sprintf("Updated at %s", time.Now().Format(time.RFC3339))
	desc := updatedDesc

	// Handle time windows
	var timeWindows []*model.TimeWindowInput
	if rule.Rule.TimeWindows != nil {
		for _, tw := range rule.Rule.TimeWindows {
			// Create a new recurrence input if it exists
			var recurrence *model.RecurrenceInput
			if tw.Recurrence != nil {
				recurrence = &model.RecurrenceInput{
					Type: tw.Recurrence.Type,
				}
				if tw.Recurrence.Spec != nil {
					recurrence.Spec = &model.RecurrenceSpecInput{
						Until: tw.Recurrence.Spec.Until,
					}
				}
			}

			timeWindow := model.TimeWindowInput{
				TimeZone:   tw.TimeZone,
				Recurrence: recurrence,
				StartTime:  tw.StartTime, // StartTime is an int
			}

			// Handle end time
			if tw.EndTime != nil {
				endTimeVal := *tw.EndTime
				timeWindow.EndTime = &endTimeVal
			}

			timeWindows = append(timeWindows, &timeWindow)
		}
	}

	// Handle tags
	var tags []*string
	if rule.Rule.Tags != nil {
		tags = make([]*string, len(rule.Rule.Tags))
		for i, tag := range rule.Rule.Tags {
			tagCopy := tag
			tags[i] = &tagCopy
		}
	}

	// Create the update request
	updateReq := model.RuleInput{
		Name:         rule.Rule.Name + " (Updated)",
		Description:  &desc,
		IsEnabled:    rule.Rule.IsEnabled,
		UserGroupIds: rule.Rule.UserGroupIds,
		ConditionIds: []string{conditionID},
		TimeWindows:  timeWindows,
		Tags:         tags,
		RuleID:       ruleID,
	}

	// Log the update request
	reqJSON, _ := json.MarshalIndent(updateReq, "", "  ")
	logger.Debug("Updating rule with request: %s", string(reqJSON))

	// Update the rule
	_, err = client.SecurityGovernanceRuleApi.Update(
		ctx,
		identifiers,
		updateReq,
	)

	if err != nil {
		logger.Error("Failed to update rule %s: %v", ruleID, err)
		return fmt.Errorf("%w: failed to update rule: %v", ErrAPI, err)
	}

	logger.Info("Successfully updated rule: %s", ruleID)
	return nil
}

// tuneExampleRule enables or disables a security governance rule
// It returns any error encountered
func tuneExampleRule(ctx context.Context, client *chaos.APIClient,
	identifiers model.IdentifiersRequest, ruleID string, enable bool) error {
	action := "disable"
	if enable {
		action = "enable"
	}

	logger.Info("%sing rule with ID: %s", action, ruleID)

	if client == nil {
		return fmt.Errorf("%w: client cannot be nil", ErrInvalidConfig)
	}

	if ruleID == "" {
		return fmt.Errorf("%w: ruleID cannot be empty", ErrInvalidConfig)
	}

	// Tune the rule
	response, err := client.SecurityGovernanceRuleApi.Tune(
		ctx,
		identifiers,
		ruleID,
		enable,
	)

	if err != nil {
		logger.Error("Failed to %s rule %s: %v", action, ruleID, err)
		return fmt.Errorf("%w: failed to %s rule: %v", ErrAPI, action, err)
	}

	// Log the response for debugging
	respJSON, _ := json.MarshalIndent(response, "", "  ")
	logger.Debug("Rule %s response: %s", action, string(respJSON))

	logger.Info("Successfully %sd rule: %s", action, ruleID)
	return nil
}

// deleteExampleRule deletes a security governance rule by ID
// It returns any error encountered
func deleteExampleRule(ctx context.Context, client *chaos.APIClient,
	identifiers model.IdentifiersRequest, ruleID string) error {
	logger.Info("Deleting rule with ID: %s", ruleID)

	if client == nil {
		return fmt.Errorf("%w: client cannot be nil", ErrInvalidConfig)
	}

	if ruleID == "" {
		return fmt.Errorf("%w: ruleID cannot be empty", ErrInvalidConfig)
	}

	// Delete the rule
	response, err := client.SecurityGovernanceRuleApi.Delete(
		ctx,
		identifiers,
		ruleID,
	)

	if err != nil {
		logger.Error("Failed to delete rule %s: %v", ruleID, err)
		return fmt.Errorf("%w: failed to delete rule: %v", ErrAPI, err)
	}

	// Log the response for debugging
	respJSON, _ := json.MarshalIndent(response, "", "  ")
	logger.Debug("Delete rule response: %s", string(respJSON))

	logger.Info("Successfully deleted rule: %s", ruleID)
	return nil
}

// Helper function to pretty print JSON output
// Helper function to create string pointers
func stringPtr(s string) *string {
	return &s
}

// prettyPrint logs a value in a pretty-printed JSON format
func prettyPrint(prefix string, v interface{}) {
	logger.Info(prefix)
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		logger.Error("Error marshaling to JSON: %v", err)
		return
	}
	logger.Info(string(b))
}
