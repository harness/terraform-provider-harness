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
	argoProject := "test123"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDatasourceGitopsAppProjectMapping(id, accountId, argoProject),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "argo_project_name", argoProject),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
					resource.TestCheckResourceAttr(resourceName, "agent_id", id),
					resource.TestCheckResourceAttr(resourceName, "account_id", accountId),
				),
			},
		},
	})
}

func testAccDatasourceGitopsAppProjectMapping(agentId string, accountId string, argoProject string) string {
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
