package fme_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccFMEIntegration_Complete(t *testing.T) {
	workspaceID := os.Getenv("SPLIT_WORKSPACE_ID")
	if workspaceID == "" {
		t.Skip("SPLIT_WORKSPACE_ID environment variable must be set for this integration test")
	}

	envName := fmt.Sprintf("test-env-%s", utils.RandStringBytes(5))
	keyName := fmt.Sprintf("test-key-%s", utils.RandStringBytes(5))
	splitName := fmt.Sprintf("test-split-%s", utils.RandStringBytes(5))

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheckFME(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccFMEIntegrationComplete(workspaceID, envName, keyName, splitName),
				Check: resource.ComposeTestCheckFunc(
					// Workspace data source
					resource.TestCheckResourceAttr("data.harness_fme_workspace.test", "id", workspaceID),
					resource.TestCheckResourceAttrSet("data.harness_fme_workspace.test", "name"),

					// Environment resource
					resource.TestCheckResourceAttr("harness_fme_environment.test", "name", envName),
					resource.TestCheckResourceAttr("harness_fme_environment.test", "production", "false"),
					resource.TestCheckResourceAttrSet("harness_fme_environment.test", "id"),

					// API Key resource
					resource.TestCheckResourceAttr("harness_fme_api_key.test", "name", keyName),
					resource.TestCheckResourceAttr("harness_fme_api_key.test", "type", "client_side"),
					resource.TestCheckResourceAttrSet("harness_fme_api_key.test", "id"),
					resource.TestCheckResourceAttrSet("harness_fme_api_key.test", "key"),

					// Split resource
					resource.TestCheckResourceAttr("harness_fme_split.test", "name", splitName),
					resource.TestCheckResourceAttr("harness_fme_split.test", "workspace_id", workspaceID),
					resource.TestCheckResourceAttrSet("harness_fme_split.test", "id"),
				),
			},
		},
	})
}

func TestAccFMEIntegration_ProductionEnvironment(t *testing.T) {
	workspaceID := os.Getenv("SPLIT_WORKSPACE_ID")
	if workspaceID == "" {
		t.Skip("SPLIT_WORKSPACE_ID environment variable must be set for this integration test")
	}

	envName := fmt.Sprintf("prod-env-%s", utils.RandStringBytes(5))
	keyName := fmt.Sprintf("prod-key-%s", utils.RandStringBytes(5))

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheckFME(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccFMEIntegrationProduction(envName, keyName),
				Check: resource.ComposeTestCheckFunc(
					// Production environment
					resource.TestCheckResourceAttr("harness_fme_environment.prod", "name", envName),
					resource.TestCheckResourceAttr("harness_fme_environment.prod", "production", "true"),

					// Server-side API key for production
					resource.TestCheckResourceAttr("harness_fme_api_key.prod", "name", keyName),
					resource.TestCheckResourceAttr("harness_fme_api_key.prod", "type", "server_side"),
				),
			},
		},
	})
}

func TestAccFMEIntegration_AllResources(t *testing.T) {
	workspaceID := os.Getenv("SPLIT_WORKSPACE_ID")
	if workspaceID == "" {
		t.Skip("SPLIT_WORKSPACE_ID environment variable must be set for this integration test")
	}

	envName := fmt.Sprintf("test-env-%s", utils.RandStringBytes(5))
	keyName := fmt.Sprintf("test-key-%s", utils.RandStringBytes(5))
	splitName := fmt.Sprintf("test-split-%s", utils.RandStringBytes(5))
	trafficTypeName := fmt.Sprintf("test-traffic-type-%s", utils.RandStringBytes(5))
	segmentName := fmt.Sprintf("test-segment-%s", utils.RandStringBytes(5))
	flagSetName := fmt.Sprintf("test-flagset-%s", utils.RandStringBytes(5))

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheckFME(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccFMEIntegrationAllResources(workspaceID, envName, keyName, splitName, trafficTypeName, segmentName, flagSetName),
				Check: resource.ComposeTestCheckFunc(
					// Workspace data source
					resource.TestCheckResourceAttr("data.harness_fme_workspace.test", "id", workspaceID),
					resource.TestCheckResourceAttrSet("data.harness_fme_workspace.test", "name"),

					// Environment resource
					resource.TestCheckResourceAttr("harness_fme_environment.test", "name", envName),
					resource.TestCheckResourceAttr("harness_fme_environment.test", "production", "false"),

					// Traffic Type resource
					resource.TestCheckResourceAttr("harness_fme_traffic_type.test", "name", trafficTypeName),
					resource.TestCheckResourceAttr("harness_fme_traffic_type.test", "workspace_id", workspaceID),
					resource.TestCheckResourceAttrSet("harness_fme_traffic_type.test", "id"),

					// Segment resource
					resource.TestCheckResourceAttr("harness_fme_segment.test", "name", segmentName),
					resource.TestCheckResourceAttr("harness_fme_segment.test", "workspace_id", workspaceID),
					resource.TestCheckResourceAttrSet("harness_fme_segment.test", "traffic_type_id"),

					// Flag Set resource
					resource.TestCheckResourceAttr("harness_fme_flag_set.test", "name", flagSetName),
					resource.TestCheckResourceAttr("harness_fme_flag_set.test", "workspace_id", workspaceID),
					resource.TestCheckResourceAttrSet("harness_fme_flag_set.test", "id"),

					// API Key resource
					resource.TestCheckResourceAttr("harness_fme_api_key.test", "name", keyName),
					resource.TestCheckResourceAttr("harness_fme_api_key.test", "type", "client_side"),
					resource.TestCheckResourceAttrSet("harness_fme_api_key.test", "key"),

					// Split resource
					resource.TestCheckResourceAttr("harness_fme_split.test", "name", splitName),
					resource.TestCheckResourceAttr("harness_fme_split.test", "workspace_id", workspaceID),
					resource.TestCheckResourceAttrSet("harness_fme_split.test", "id"),

					// Data sources
					resource.TestCheckResourceAttr("data.harness_fme_environment.test", "name", envName),
					resource.TestCheckResourceAttr("data.harness_fme_flag_set.test", "name", flagSetName),
					resource.TestCheckResourceAttr("data.harness_fme_traffic_type.test", "name", trafficTypeName),
				),
			},
		},
	})
}

func testAccFMEIntegrationComplete(workspaceID, envName, keyName, splitName string) string {
	return fmt.Sprintf(`
		data "harness_fme_workspace" "test" {
			id = "%[1]s"
		}

		resource "harness_fme_environment" "test" {
			name       = "%[2]s"
			production = false
		}

		resource "harness_fme_api_key" "test" {
			environment_id = harness_fme_environment.test.id
			name          = "%[3]s"
			type          = "client_side"
		}

		resource "harness_fme_split" "test" {
			workspace_id = data.harness_fme_workspace.test.id
			name         = "%[4]s"
			description  = "Integration test split"
		}

		output "workspace_name" {
			value = data.harness_fme_workspace.test.name
		}

		output "environment_id" {
			value = harness_fme_environment.test.id
		}

		output "api_key" {
			value     = harness_fme_api_key.test.key
			sensitive = true
		}

		output "split_id" {
			value = harness_fme_split.test.id
		}
`, workspaceID, envName, keyName, splitName)
}

func testAccFMEIntegrationProduction(envName, keyName string) string {
	return fmt.Sprintf(`
		resource "harness_fme_environment" "prod" {
			name       = "%[1]s"
			production = true
		}

		resource "harness_fme_api_key" "prod" {
			environment_id = harness_fme_environment.prod.id
			name          = "%[2]s"
			type          = "server_side"
		}
`, envName, keyName)
}

func testAccFMEIntegrationAllResources(workspaceID, envName, keyName, splitName, trafficTypeName, segmentName, flagSetName string) string {
	return fmt.Sprintf(`
		data "harness_fme_workspace" "test" {
			id = "%[1]s"
		}

		resource "harness_fme_environment" "test" {
			name       = "%[2]s"
			production = false
		}

		resource "harness_fme_traffic_type" "test" {
			workspace_id = "%[1]s"
			name         = "%[5]s"
		}

		resource "harness_fme_segment" "test" {
			workspace_id    = "%[1]s"
			traffic_type_id = harness_fme_traffic_type.test.id
			name            = "%[6]s"
			description     = "Integration test segment"
		}

		resource "harness_fme_flag_set" "test" {
			workspace_id = "%[1]s"
			name         = "%[7]s"
			description  = "Integration test flag set"
		}

		resource "harness_fme_api_key" "test" {
			environment_id = harness_fme_environment.test.id
			name          = "%[3]s"
			type          = "client_side"
		}

		resource "harness_fme_split" "test" {
			workspace_id = data.harness_fme_workspace.test.id
			name         = "%[4]s"
			description  = "Integration test split"
		}

		data "harness_fme_environment" "test" {
			id = harness_fme_environment.test.id
		}

		data "harness_fme_flag_set" "test" {
			id = harness_fme_flag_set.test.id
		}

		data "harness_fme_traffic_type" "test" {
			workspace_id = "%[1]s"
			id           = harness_fme_traffic_type.test.id
		}

		output "workspace_name" {
			value = data.harness_fme_workspace.test.name
		}

		output "environment_id" {
			value = harness_fme_environment.test.id
		}

		output "traffic_type_id" {
			value = harness_fme_traffic_type.test.id
		}

		output "segment_name" {
			value = harness_fme_segment.test.name
		}

		output "flag_set_id" {
			value = harness_fme_flag_set.test.id
		}

		output "api_key" {
			value     = harness_fme_api_key.test.key
			sensitive = true
		}

		output "split_id" {
			value = harness_fme_split.test.id
		}
`, workspaceID, envName, keyName, splitName, trafficTypeName, segmentName, flagSetName)
}