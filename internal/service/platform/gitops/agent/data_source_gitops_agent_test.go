package agent_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"

	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceGitopsAgent(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	id = strings.ReplaceAll(id, "_", "")
	name := id
	agentId := id
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	resourceName := "data.harness_platform_gitops_agent.test"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceGitopsAgent(agentId, name, accountId, agentId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
				),
			},
		},
	})

}

// FLAMINGO
func TestAccDataSourceGitopsAgentFlamingo(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	id = strings.ReplaceAll(id, "_", "")
	name := id
	agentId := id
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	resourceName := "data.harness_platform_gitops_agent.test"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceGitopsAgentFlamingo(agentId, name, accountId, agentId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
				),
			},
		},
	})

}

func testAccDataSourceGitopsAgent(agentId string, name string, accountId string, agentName string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
		}

		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_organization.test.id
		}
		resource "harness_platform_gitops_agent" "test" {
			identifier = "%[1]s"
			account_id = "%[3]s"
			project_id = harness_platform_project.test.id
			org_id = harness_platform_organization.test.id
			name = "%[4]s"
			type = "MANAGED_ARGO_PROVIDER"
			metadata {
        		namespace = "terraform-test"
        		high_availability = false
			}
		}

		data "harness_platform_gitops_agent" "test" {
			depends_on = [harness_platform_gitops_agent.test]
			identifier = harness_platform_gitops_agent.test.id
			account_id = "%[3]s"
			project_id = harness_platform_project.test.id
			org_id = harness_platform_organization.test.id
		}
		`, agentId, name, accountId, agentName)
}

func testAccDataSourceGitopsAgentFlamingo(agentId string, name string, accountId string, agentName string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
		}

		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_organization.test.id
		}
		resource "harness_platform_gitops_agent" "test" {
			identifier = "%[1]s"
			account_id = "%[3]s"
			project_id = harness_platform_project.test.id
			org_id = harness_platform_organization.test.id
			name = "%[4]s"
			type = "MANAGED_ARGO_PROVIDER"
			metadata {
        		namespace = "terraform-test"
        		high_availability = false
			}
			operator = "FLAMINGO"	
		}

		data "harness_platform_gitops_agent" "test" {
			depends_on = [harness_platform_gitops_agent.test]
			identifier = harness_platform_gitops_agent.test.id
			account_id = "%[3]s"
			project_id = harness_platform_project.test.id
			org_id = harness_platform_organization.test.id
		}
		`, agentId, name, accountId, agentName)
}
