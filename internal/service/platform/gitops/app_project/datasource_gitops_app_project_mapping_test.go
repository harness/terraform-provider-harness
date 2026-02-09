package app_project_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDatasourceGitopsAppProjectMapping(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	id = strings.ReplaceAll(id, "_", "")
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	resourceName := "harness_platform_gitops_app_project_mapping.test"
	dataSourceName := "data.harness_platform_gitops_app_project_mapping.test1"
	argoProject := "test123"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				// Test Case 1: Data source reads auto_create_service_env = true
				Config: testAccDatasourceGitopsAppProjectMappingWithAutoCreate(id, accountId, argoProject, "true"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "argo_project_name", argoProject),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
					resource.TestCheckResourceAttr(resourceName, "agent_id", id),
					resource.TestCheckResourceAttr(resourceName, "account_id", accountId),
					resource.TestCheckResourceAttr(resourceName, "auto_create_service_env", "true"),
					// Verify data source reads the same value
					resource.TestCheckResourceAttr(dataSourceName, "auto_create_service_env", "true"),
				),
			},
			{
				// Test Case 2: Data source reads auto_create_service_env = false
				Config: testAccDatasourceGitopsAppProjectMappingWithAutoCreate(id, accountId, argoProject, "false"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "auto_create_service_env", "false"),
					// Verify data source reads the same value
					resource.TestCheckResourceAttr(dataSourceName, "auto_create_service_env", "false"),
				),
			},
			{
				// Test Case 3: Data source reads default value (false) when field is omitted
				Config: testAccDatasourceGitopsAppProjectMappingWithoutAutoCreate(id, accountId, argoProject),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "auto_create_service_env", "false"),
					// Verify data source reads the same value
					resource.TestCheckResourceAttr(dataSourceName, "auto_create_service_env", "false"),
				),
			},
		},
	})
}

func testAccDatasourceGitopsAppProjectMappingWithAutoCreate(agentId string, accountId string, argoProject string, autoCreate string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
		}

		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			org_id = harness_platform_organization.test.id
		}
		resource "harness_platform_gitops_agent" "test" {
			identifier = "%[1]s"
			account_id = "%[2]s"
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			name = "%[1]s"
			type = "MANAGED_ARGO_PROVIDER"
			operator = "ARGO"
			metadata {
				namespace = "%[1]s"
        		high_availability = false
    		}
		}
		resource "harness_platform_gitops_app_project_mapping" "test" {
			depends_on = [harness_platform_gitops_agent.test]
			account_id = "%[2]s"
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			agent_id = "%[1]s"
			argo_project_name = "%[3]s"
			auto_create_service_env = %[4]s
		}

		data "harness_platform_gitops_app_project_mapping" test1 {
			depends_on = [harness_platform_gitops_app_project_mapping.test]
			
			org_id = harness_platform_organization.test.id
			account_id = "%[2]s"
			project_id = harness_platform_project.test.id
			agent_id = "%[1]s"
            argo_project_name = "%[3]s"
		}
		`, agentId, accountId, argoProject, autoCreate)
}

func testAccDatasourceGitopsAppProjectMappingWithoutAutoCreate(agentId string, accountId string, argoProject string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
		}

		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			org_id = harness_platform_organization.test.id
		}
		resource "harness_platform_gitops_agent" "test" {
			identifier = "%[1]s"
			account_id = "%[2]s"
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			name = "%[1]s"
			type = "MANAGED_ARGO_PROVIDER"
			operator = "ARGO"
			metadata {
				namespace = "%[1]s"
        		high_availability = false
    		}
		}
		resource "harness_platform_gitops_app_project_mapping" "test" {
			depends_on = [harness_platform_gitops_agent.test]
			account_id = "%[2]s"
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			agent_id = "%[1]s"
			argo_project_name = "%[3]s"
		}

		data "harness_platform_gitops_app_project_mapping" test1 {
			depends_on = [harness_platform_gitops_app_project_mapping.test]
			
			org_id = harness_platform_organization.test.id
			account_id = "%[2]s"
			project_id = harness_platform_project.test.id
			agent_id = "%[1]s"
            argo_project_name = "%[3]s"
		}
		`, agentId, accountId, argoProject)
}
