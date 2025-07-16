// Package convert provides functions to convert between API models and Terraform schema.
package convert

import (
	"fmt"

	"github.com/harness/harness-go-sdk/harness/svcdiscovery"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

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
