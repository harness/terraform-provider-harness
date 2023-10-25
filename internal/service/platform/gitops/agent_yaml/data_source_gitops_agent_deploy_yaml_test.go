package agent_yaml_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"

	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceGitopsAgentDeployYaml(t *testing.T) {
	agentId := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	agentId = strings.ReplaceAll(agentId, "_", "")
	namespace := "ns-" + agentId
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	resourceName := "data.harness_platform_gitops_agent_deploy_yaml.test"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceGitopsAgentDeployYaml(agentId, accountId, agentId, namespace),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "yaml"),
				),
			},
		},
	})

}

func testAccDataSourceGitopsAgentDeployYaml(agentId string, accountId string, agentName string, namespace string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[3]s"
		}

		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name = "%[3]s"
			org_id = harness_platform_organization.test.id
		}
		resource "harness_platform_gitops_agent" "test" {
			identifier = "%[1]s"
			account_id = "%[2]s"
			project_id = harness_platform_project.test.id
			org_id = harness_platform_organization.test.id
			name = "%[3]s"
			type = "MANAGED_ARGO_PROVIDER"
			operator = "ARGO"
			metadata {
        		namespace = "%[4]s"
        		high_availability = false
    		}
		}
		
		data "harness_platform_gitops_agent_deploy_yaml" "test" {
			depends_on = [harness_platform_gitops_agent.test]
			identifier = "%[1]s"
			account_id = "%[2]s"
			project_id = harness_platform_project.test.id
			org_id = harness_platform_organization.test.id
			namespace = "%[4]s"
		}
		`, agentId, accountId, agentName, namespace)
}
