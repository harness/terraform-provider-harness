// Package convert provides functions to convert between API models and Terraform schema.
package convert

import (
	"fmt"

	"github.com/harness/harness-go-sdk/harness/svcdiscovery"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// // FlattenAgentToSchema converts an Agent model to a Terraform schema.
// // It handles all field conversions and nil checks.
// func FlattenAgentToSchema(d *schema.ResourceData, agent *svcdiscovery.ApiGetAgentResponse) error {
// 	if agent == nil {
// 		return fmt.Errorf("cannot flatten nil agent")
// 	}

// 	// Set basic fields
// 	if err := setIfNotEmpty(d, "id", agent.Id); err != nil {
// 		return fmt.Errorf("failed to set id: %w", err)
// 	}
// 	if err := setIfNotEmpty(d, "name", agent.Name); err != nil {
// 		return fmt.Errorf("failed to set name: %w", err)
// 	}
// 	if err := setIfNotEmpty(d, "description", agent.Description); err != nil {
// 		return fmt.Errorf("failed to set description: %w", err)
// 	}

// 	// Set identifiers
// 	if err := setIfNotEmpty(d, "account_identifier", agent.AccountIdentifier); err != nil {
// 		return fmt.Errorf("failed to set account_identifier: %w", err)
// 	}
// 	if err := setIfNotEmpty(d, "environment_identifier", agent.EnvironmentIdentifier); err != nil {
// 		return fmt.Errorf("failed to set environment_identifier: %w", err)
// 	}
// 	if err := setIfNotEmpty(d, "org_identifier", agent.OrganizationIdentifier); err != nil {
// 		return fmt.Errorf("failed to set org_identifier: %w", err)
// 	}
// 	if err := setIfNotEmpty(d, "project_identifier", agent.ProjectIdentifier); err != nil {
// 		return fmt.Errorf("failed to set project_identifier: %w", err)
// 	}
// 	if err := setIfNotEmpty(d, "identity", agent.Identity); err != nil {
// 		return fmt.Errorf("failed to set identity: %w", err)
// 	}
// 	// if err := setIfNotEmpty(d, "infra_identifier", agent.InfraIdentifier); err != nil {
// 	// 	return fmt.Errorf("failed to set infra_identifier: %w", err)
// 	// }

// 	// Set status fields
// 	if err := d.Set("permanent_installation", agent.PermanentInstallation); err != nil {
// 		return fmt.Errorf("failed to set permanent_installation: %w", err)
// 	}
// 	if err := d.Set("removed", agent.Removed); err != nil {
// 		return fmt.Errorf("failed to set removed: %w", err)
// 	}

// 	// Set timestamps
// 	if err := setIfNotEmpty(d, "created_at", agent.CreatedAt); err != nil {
// 		return fmt.Errorf("failed to set created_at: %w", err)
// 	}
// 	if err := setIfNotEmpty(d, "updated_at", agent.UpdatedAt); err != nil {
// 		return fmt.Errorf("failed to set updated_at: %w", err)
// 	}
// 	if err := setIfNotEmpty(d, "created_by", agent.CreatedBy); err != nil {
// 		return fmt.Errorf("failed to set created_by: %w", err)
// 	}
// 	if err := setIfNotEmpty(d, "updated_by", agent.UpdatedBy); err != nil {
// 		return fmt.Errorf("failed to set updated_by: %w", err)
// 	}

// 	// Set tags
// 	if agent.Tags != nil && len(agent.Tags) > 0 {
// 		tags := make(map[string]string)
// 		for _, tag := range agent.Tags {
// 			parts := strings.SplitN(tag, "=", 2)
// 			if len(parts) == 2 {
// 				tags[parts[0]] = parts[1]
// 			} else {
// 				tags[tag] = "" // Handle tags without values
// 			}
// 		}
// 		if err := d.Set("tags", tags); err != nil {
// 			return fmt.Errorf("failed to set tags: %w", err)
// 		}
// 	} else {
// 		if err := d.Set("tags", map[string]string{}); err != nil {
// 			return fmt.Errorf("failed to set empty tags: %w", err)
// 		}
// 	}

// 	// Handle configuration
// 	if agent.Config != nil {
// 		config, err := FlattenAgentConfig(agent.Config)
// 		if err != nil {
// 			return fmt.Errorf("failed to flatten agent config: %w", err)
// 		}
// 		if err := d.Set("config", []interface{}{config}); err != nil {
// 			return fmt.Errorf("failed to set config: %w", err)
// 		}
// 	} else {
// 		if err := d.Set("config", []interface{}{}); err != nil {
// 			return fmt.Errorf("failed to set empty config: %w", err)
// 		}
// 	}

// 	// Handle installation details
// 	if agent.InstallationDetails != nil {
// 		details, err := FlattenInstallationDetails(agent.InstallationDetails)
// 		if err != nil {
// 			return fmt.Errorf("failed to flatten installation details: %w", err)
// 		}
// 		if err := d.Set("installation_details", []interface{}{details}); err != nil {
// 			return fmt.Errorf("failed to set installation_details: %w", err)
// 		}
// 	} else {
// 		if err := d.Set("installation_details", []interface{}{}); err != nil {
// 			return fmt.Errorf("failed to set empty installation_details: %w", err)
// 		}
// 	}

// 	return nil
// }

// // setIfNotEmpty is a helper function to set a schema value if the input is not empty.
// func setIfNotEmpty(d *schema.ResourceData, key, value string) error {
// 	if value != "" {
// 		return d.Set(key, value)
// 	}
// 	return d.Set(key, "") // Ensure empty strings are set explicitly
// }

func FlattenAgentToSchema(d *schema.ResourceData, agent *svcdiscovery.ApiGetAgentResponse) error {
	if agent == nil {
		return fmt.Errorf("cannot flatten nil agent")
	}

	// Always set fields, even if empty
	setStringIfNotEmpty := func(key, value string) {
		if value == "" {
			d.Set(key, "")
		} else {
			d.Set(key, value)
		}
	}

	setStringIfNotEmpty("id", agent.Id)
	setStringIfNotEmpty("name", agent.Name)
	setStringIfNotEmpty("description", agent.Description)
	setStringIfNotEmpty("account_identifier", agent.AccountIdentifier)
	setStringIfNotEmpty("environment_identifier", agent.EnvironmentIdentifier)
	setStringIfNotEmpty("org_identifier", agent.OrganizationIdentifier)
	setStringIfNotEmpty("project_identifier", agent.ProjectIdentifier)
	setStringIfNotEmpty("identity", agent.Identity)
	setStringIfNotEmpty("installation_type", agent.InstallationType)
	setStringIfNotEmpty("webhook_url", agent.WebhookURL)
	setStringIfNotEmpty("correlation_id", agent.CorrelationID)
	setStringIfNotEmpty("created_at", agent.CreatedAt)
	setStringIfNotEmpty("updated_at", agent.UpdatedAt)
	setStringIfNotEmpty("created_by", agent.CreatedBy)
	setStringIfNotEmpty("updated_by", agent.UpdatedBy)
	// setStringIfNotEmpty("infra_identifier", agent.InfraIdentifier)

	d.Set("network_map_count", agent.NetworkMapCount)
	d.Set("service_count", agent.ServiceCount)

	d.Set("permanent_installation", agent.PermanentInstallation)
	d.Set("removed", agent.Removed)

	// Handle tags - ensure it's never null
	if agent.Tags != nil {
		d.Set("tags", agent.Tags)
	} else {
		d.Set("tags", map[string]string{})
	}

	// Handle config
	if agent.Config != nil {
		config, err := FlattenAgentConfig(agent.Config)
		if err != nil {
			return fmt.Errorf("failed to flatten agent config: %w", err)
		}
		d.Set("config", []interface{}{config})
	} else {
		d.Set("config", []interface{}{})
	}

	// Handle installation details
	if agent.InstallationDetails != nil {
		details, err := FlattenInstallationDetails(agent.InstallationDetails)
		if err != nil {
			return fmt.Errorf("failed to flatten installation details: %w", err)
		}
		d.Set("installation_details", []interface{}{details})
	} else {
		d.Set("installation_details", []interface{}{})
	}

	return nil
}
