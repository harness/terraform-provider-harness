package chaos_hub_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceChaosHub(t *testing.T) {

	// Generate unique identifiers
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	rName := id
	dataSourceName := "data.harness_chaos_hub.test"
	resourceName := "harness_chaos_hub.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceChaosHubConfig(rName, id, "main"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "name", rName),
					resource.TestCheckResourceAttrPair(dataSourceName, "id", resourceName, "id"),
					resource.TestCheckResourceAttrPair(dataSourceName, "connector_id", resourceName, "connector_id"),
					resource.TestCheckResourceAttrPair(dataSourceName, "repo_branch", resourceName, "repo_branch"),
					resource.TestCheckResourceAttr(dataSourceName, "connector_scope", "ACCOUNT"),
					resource.TestCheckResourceAttrSet(dataSourceName, "created_at"),
					resource.TestCheckResourceAttrSet(dataSourceName, "is_available"),
				),
			},
		},
	})
}

func TestAccDataSourceChaosHub_ProjectLevel(t *testing.T) {

	// Generate unique identifiers
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	rName := id
	dataSourceName := "data.harness_chaos_hub.test"
	resourceName := "harness_chaos_hub.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceChaosHubProjectLevelConfig(rName, id, "develop"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "name", rName),
					resource.TestCheckResourceAttrPair(dataSourceName, "id", resourceName, "id"),
					resource.TestCheckResourceAttrPair(dataSourceName, "project_id", resourceName, "project_id"),
					resource.TestCheckResourceAttrPair(dataSourceName, "org_id", resourceName, "org_id"),
					resource.TestCheckResourceAttr(dataSourceName, "connector_scope", "PROJECT"),
				),
			},
		},
	})
}

// Terraform Configurations

func testAccDataSourceChaosHubConfig(name, id, branch string) string {

	return fmt.Sprintf(`
	resource "harness_platform_connector_github" "test" {
		identifier  = "%[2]s"
		name        = "%[1]s"
		description = "Test connector"
		url         = "https://github.com"
		connection_type = "Repo"
		validation_repo = "harness/chaos-hub"
	
		credentials {
			http {
				username  = "test"
				token_ref = "account.do_not_delete_harness_platform_qa_token"
			}
		}
	}

	resource "harness_chaos_hub" "test" {
		name           = "%[1]s"
		connector_id   = harness_platform_connector_github.test.id
		connector_scope = "ACCOUNT"
		repo_branch    = "%[3]s"
		description    = "Test chaos hub"
		tags           = ["test:true", "chaos:true"]
	}

	data "harness_chaos_hub" "test" {
		name = harness_chaos_hub.test.name
	}
	`, name, id, branch)
}

func testAccDataSourceChaosHubProjectLevelConfig(name, id, branch string) string {
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

	resource "harness_platform_connector_github" "test" {
		identifier  = "%[2]s"
		name        = "%[1]s"
		description = "Test connector"
		url         = "https://github.com"
		connection_type = "Repo"
		validation_repo = "harness/chaos-hub"
		project_id  = harness_platform_project.test.id
		org_id      = harness_platform_organization.test.id
	
		credentials {
			http {
				username  = "test"
				token_ref = "account.do_not_delete_harness_platform_qa_token"
			}
		}
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
	`, name, id, branch)
}
