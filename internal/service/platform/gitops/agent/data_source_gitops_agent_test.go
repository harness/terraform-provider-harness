package agent_test

import (
	"fmt"
	"github.com/harness/harness-go-sdk/harness/utils"
	"os"
	"testing"

	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceGitopsAgent(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
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
					resource.TestCheckResourceAttr(resourceName, "name", id),
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
			type = "CONNECTED_ARGO_PROVIDER"
			metadata {
        		namespace = "terraform-test"
        		high_availability = true
			}
		}

		data "harness_platform_gitops_agent" "test" {
			identifier = harness_platform_gitops_agent.test.id
			account_id = "%[3]s"
			project_id = harness_platform_project.test.id
			org_id = harness_platform_organization.test.id
			name = harness_platform_gitops_agent.test.name
			type = "CONNECTED_ARGO_PROVIDER"
			metadata {
        		namespace = "terraform-test"
        		high_availability = true
			}

		}
		`, agentId, name, accountId, agentName)
}
