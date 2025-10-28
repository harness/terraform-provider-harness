package chaos_hub_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceChaosHub_ProjectLevel(t *testing.T) {
	chaosGithubToken := os.Getenv("CHAOS_GITHUB_TOKEN")
	if chaosGithubToken == "" {
		t.Skip("Skipping test because CHAOS_GITHUB_TOKEN is not set")
	}

	// Generate unique identifiers
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	rName := id
	dataSourceName := "data.harness_chaos_hub.test"
	resourceName := "harness_chaos_hub.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceChaosHubProjectLevelConfig(rName, id, "master", chaosGithubToken),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "name", rName),
					resource.TestCheckResourceAttrPair(dataSourceName, "project_id", resourceName, "project_id"),
					resource.TestCheckResourceAttrPair(dataSourceName, "org_id", resourceName, "org_id"),
					resource.TestCheckResourceAttrPair(dataSourceName, "connector_id", resourceName, "connector_id"),
					resource.TestCheckResourceAttrPair(dataSourceName, "repo_branch", resourceName, "repo_branch"),
					resource.TestCheckResourceAttrSet(dataSourceName, "created_at"),
					resource.TestCheckResourceAttrSet(dataSourceName, "is_available"),
					resource.TestCheckResourceAttrSet(dataSourceName, "last_synced_at"),
					resource.TestCheckResourceAttrSet(dataSourceName, "total_experiments"),
					resource.TestCheckResourceAttrSet(dataSourceName, "total_faults"),
					resource.TestCheckResourceAttrSet(dataSourceName, "updated_at"),
				),
			},
		},
	})
}

// Terraform Configurations

func testAccDataSourceChaosHubProjectLevelConfig(name, id, branch, githubToken string) string {
	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
		identifier = "%[2]s"
		name       = "%[1]s"
	}

	resource "harness_platform_project" "test" {
		identifier  = "%[2]s"
		name        = "%[1]s"
		org_id      = harness_platform_organization.test.id
		color       = "#0063F7"
		description = "Test project for Chaos Hub"
		tags        = ["foo:bar", "baz:qux"]
	}
	
	resource "harness_platform_secret_text" "test" {
		identifier = "%[2]s"
		name = "%[1]s"
		description = "test"
		tags = ["foo:bar"]

		secret_manager_identifier = "harnessSecretManager"
		value_type = "Inline"
		value = "%[4]s"
	}

	resource "harness_platform_connector_github" "test" {
		identifier  = "%[2]s"
		name        = "%[1]s"
		description = "Test connector"
		url         = "https://github.com/litmuschaos/chaos-charts"
		connection_type = "Repo"
		project_id  = harness_platform_project.test.id
		org_id      = harness_platform_organization.test.id
	
		credentials {
			http {
				username = "admin"
				token_ref = "account.${harness_platform_secret_text.test.id}"
			}
		}
		depends_on = [time_sleep.wait_4_seconds]
	}

	resource "time_sleep" "wait_4_seconds" {
		depends_on = [harness_platform_secret_text.test]
		destroy_duration = "4s"
	}

	resource "harness_chaos_hub" "test" {
		org_id        = harness_platform_organization.test.id
		project_id    = harness_platform_project.test.id
		name          = "%[1]s"
		connector_id  = harness_platform_connector_github.test.id
		connector_scope = "PROJECT"
		repo_branch   = "%[3]s"
		description   = "Test chaos hub in project"
		tags          = ["test:true", "chaos:true"]
	}

	data "harness_chaos_hub" "test" {
		name       = harness_chaos_hub.test.name
		org_id     = harness_platform_organization.test.id
		project_id = harness_platform_project.test.id
	}
	`, name, id, branch, githubToken)
}
