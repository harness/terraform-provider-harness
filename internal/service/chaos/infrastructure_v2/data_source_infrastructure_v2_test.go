package infrastructure_v2_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceChaosInfrastructureV2_basic(t *testing.T) {
	// Generate unique identifiers
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	rName := id
	resourceName := "data.harness_chaos_infrastructure_v2.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceChaosInfrastructureV2Config(rName, id, "KUBERNETESV2", "NAMESPACE"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "infra_type", "KUBERNETESV2"),
					resource.TestCheckResourceAttr(resourceName, "infra_scope", "NAMESPACE"),
					resource.TestCheckResourceAttrSet(resourceName, "status"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
				),
			},
		},
	})
}

func TestAccDataSourceChaosInfrastructureV2_WithAllOptions(t *testing.T) {
	// Generate unique identifiers
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	rName := id
	resourceName := "data.harness_chaos_infrastructure_v2.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceChaosInfrastructureV2Config_WithAllOptions(rName, id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "infra_type", "KUBERNETESV2"),
					resource.TestCheckResourceAttr(resourceName, "infra_scope", "CLUSTER"),
					resource.TestCheckResourceAttr(resourceName, "ai_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "insecure_skip_verify", "true"),
					resource.TestCheckResourceAttr(resourceName, "namespace", "chaos-namespace"),
					resource.TestCheckResourceAttr(resourceName, "service_account", "litmus-admin"),
					resource.TestCheckResourceAttrSet(resourceName, "status"),
				),
			},
		},
	})
}

func TestAccDataSourceChaosInfrastructureV2_KubernetesType(t *testing.T) {
	// Generate unique identifiers
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	rName := id
	resourceName := "data.harness_chaos_infrastructure_v2.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceChaosInfrastructureV2Config(rName, id, "KUBERNETES", "NAMESPACE"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "infra_type", "KUBERNETES"),
					resource.TestCheckResourceAttr(resourceName, "infra_scope", "NAMESPACE"),
				),
			},
		},
	})
}
func testAccDataSourceChaosInfrastructureV2Config(name, id, infraType, infraScope string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name       = "%[2]s"
		}

		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name       = "%[2]s"
			org_id     = harness_platform_organization.test.id
		}

		resource "harness_platform_environment" "test" {
			identifier = "%[1]s"
			name       = "%[2]s"
			org_id     = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			type       = "PreProduction"
		}

		resource "harness_chaos_infrastructure_v2" "test" {
			org_id        = harness_platform_organization.test.id
			project_id    = harness_platform_project.test.id
			environment_id = harness_platform_environment.test.id
			name          = "%[2]s"
			infra_id      = "%[1]s"
			description   = "Test Infrastructure"
			infra_type    = "%[3]s"
			infra_scope   = "%[4]s"
			namespace     = "chaos"
			service_account = "litmus"
		}

		data "harness_chaos_infrastructure_v2" "test" {
			org_id        = harness_platform_organization.test.id
			project_id    = harness_platform_project.test.id
			environment_id = harness_platform_environment.test.id
			infra_id      = harness_chaos_infrastructure_v2.test.infra_id
		}
	`, id, name, infraType, infraScope)
}

func testAccDataSourceChaosInfrastructureV2Config_WithAllOptions(name, id string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name       = "%[2]s"
		}

		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name       = "%[2]s"
			org_id     = harness_platform_organization.test.id
		}

		resource "harness_platform_environment" "test" {
			identifier = "%[1]s"
			name       = "%[2]s"
			org_id     = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			type       = "PreProduction"
		}

		resource "harness_chaos_infrastructure_v2" "test" {
			org_id             = harness_platform_organization.test.id
			project_id         = harness_platform_project.test.id
			environment_id      = harness_platform_environment.test.id
			name               = "%[2]s"
			infra_id           = "%[1]s"
			description        = "Test Infrastructure with all options"
			infra_type         = "KUBERNETESV2"
			infra_scope        = "CLUSTER"
			namespace          = "chaos-namespace"
			service_account    = "litmus-admin"
			ai_enabled         = true
			insecure_skip_verify = true
		}

		data "harness_chaos_infrastructure_v2" "test" {
			org_id        = harness_platform_organization.test.id
			project_id    = harness_platform_project.test.id
			environment_id = harness_platform_environment.test.id
			infra_id      = harness_chaos_infrastructure_v2.test.infra_id
		}
	`, id, name)
}

func TestAccDataSourceChaosInfrastructureV2_NotFound(t *testing.T) {
	// Generate unique identifiers
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	rName := id

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceChaosInfrastructureV2Config_NonExistent(rName, id),
				ExpectError: regexp.MustCompile("not found"),
			},
		},
	})
}

func testAccDataSourceChaosInfrastructureV2Config_NonExistent(name, id string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name       = "%[2]s"
		}

		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name       = "%[2]s"
			org_id     = harness_platform_organization.test.id
		}

		resource "harness_platform_environment" "test" {
			identifier = "%[1]s"
			name       = "%[2]s"
			org_id     = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			type       = "PreProduction"
		}

		data "harness_chaos_infrastructure_v2" "test" {
			org_id        = harness_platform_organization.test.id
			project_id    = harness_platform_project.test.id
			environment_id = harness_platform_environment.test.id
			infra_id      = "nonexistent-infra"
		}
	`, id, name)
}
