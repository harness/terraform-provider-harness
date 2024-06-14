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
			ca_data = "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURuekNDQW9lZ0F3SUJBZ0lVYm8wQmJSU2IrYWE3OGhyWWFueDdUangxdE1Vd0RRWUpLb1pJaHZjTkFRRUwKQlFBd1h6RUxNQWtHQTFVRUJoTUNWVk14Q3pBSkJnTlZCQWdNQWtOQk1Rc3dDUVlEVlFRSERBSlRSakVoTUI4RwpBMVVFQ2d3WVNXNTBaWEp1WlhRZ1YybGtaMmwwY3lCUWRIa2dUSFJrTVJNd0VRWURWUVFEREFwa2IyMWhhVzR1ClkyOXRNQjRYRFRJek1URXhOVEl4TkRrME5Gb1hEVEkwTVRFeE5ESXhORGswTkZvd1h6RUxNQWtHQTFVRUJoTUMKVlZNeEN6QUpCZ05WQkFnTUFrTkJNUXN3Q1FZRFZRUUhEQUpUUmpFaE1COEdBMVVFQ2d3WVNXNTBaWEp1WlhRZwpWMmxrWjJsMGN5QlFkSGtnVEhSa01STXdFUVlEVlFRRERBcGtiMjFoYVc0dVkyOXRNSUlCSWpBTkJna3Foa2lHCjl3MEJBUUVGQUFPQ0FROEFNSUlCQ2dLQ0FRRUF2UXJRR3JpdXN5OG5Hbk1hWjhvUk9nd1NiN05OdDNkM3llb1oKV0JmV2ZNU0xhWXpwdjcvL0Noc2lSdzlFUTNKcFN1SlF0bUx4SDdsZHcwNVZyY0M4VTBQOWFEWlZ1Q1ljOStSTwpiRUF6MmtwVUFkcUw4N29uYkN6OVkweERwTmJIZDJOaGtkZGF6ME9DbDJJOU10MGdSTk9ZT0N1RXliZS90TStvCmR3WVdMTnYrMXJGb2NKbEFJYjZ0Z2t1MldoZUNNYXlsV1Jqc1U1VldkNXdUTitUWW9GN25YWVhjWmY3cHhtekYKUytTY3l3NDN4M2hsU0E3RzNsZnY2Ri9VOWY5YVpZU0ZVUktMenByQllXaWpUR3F6Mm94M3VDUWF0MFlpTkdxMwpoS2RZY2N4UVJJQnl3dDEyR3RmaUFja0l6NnpKVVdoZzJDN1cxZkIzN2ZBNVBWOFRsUUlEQVFBQm8xTXdVVEFkCkJnTlZIUTRFRmdRVXY2ZUx3QytURjBHSzlyamx0TVEwMzh3NFQzb3dId1lEVlIwakJCZ3dGb0FVdjZlTHdDK1QKRjBHSzlyamx0TVEwMzh3NFQzb3dEd1lEVlIwVEFRSC9CQVV3QXdFQi96QU5CZ2txaGtpRzl3MEJBUXNGQUFPQwpBUUVBSGVyWGc1a2hEVkxpWG9ZSmpRMnhTQ2xoQlVIdGdSTGJ6R25ZekJ1R3VseTh1UW9BZ1dLZU1kM0pjSk93CmJ4K3c3NzFsUzFNbmdENEhiK0ZXWWxkdE5xUHZQa2c3RXZKb2lFMHQzSElzck02WXdyNDUvNHBSZVBMWU1paSsKV0FqTFhOZGVuUUUwVlFvY2pzKzN4M0QyK0FOYitRUTFxVTAzYVhiSEVRQzdmU2k1Y2pPMjd1aWRocVoyNEtHbQpJOE0vN0VmRWhWR01LeStKYnd3WGdYaVZvQ2FHK3QyUTRHS3NETlNJRlNWVEpsT1JPbXBYUUVZdmxIMmRndzdkClVieWRZR1l0TXlBY3hSSHVTWCtwMVNXSFZabDJ1TjMvd1AzWWt1M1kwbXFkdzVPdk1Lc1htbkJFNGFtb3BvQm0KWFV4V2F6U3YxaG5iaGdMWkRWVHk4VmRYY0E9PQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg=="
            proxy {
				http = "http://proxy.com"
				https = "https://proxy.com"
				username = "user"
				password = "pass"
			}
		}
		`, agentId, accountId, agentName, namespace)
}
