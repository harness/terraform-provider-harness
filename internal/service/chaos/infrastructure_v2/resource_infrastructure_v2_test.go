package infrastructure_v2_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceChaosInfrastructureV2_basic(t *testing.T) {
	// Check for required environment variables
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	if accountId == "" {
		t.Skip("Skipping test because HARNESS_ACCOUNT_ID is not set")
	}

	// Generate unique identifiers
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	rName := id
	resourceName := "harness_chaos_infrastructure_v2.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccChaosInfrastructureV2Destroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceChaosInfrastructureV2ConfigBasic(rName, id, "KUBERNETESV2", "NAMESPACE"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "infra_type", "KUBERNETESV2"),
					resource.TestCheckResourceAttr(resourceName, "infra_scope", "NAMESPACE"),
					resource.TestCheckResourceAttr(resourceName, "namespace", "chaos"),
					resource.TestCheckResourceAttr(resourceName, "service_account", "litmus"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccResourceChaosInfrastructureV2ImportStateIdFunc(resourceName),
			},
		},
	})
}

func TestAccResourceChaosInfrastructureV2_Update(t *testing.T) {
	// Check for required environment variables
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	if accountId == "" {
		t.Skip("Skipping test because HARNESS_ACCOUNT_ID is not set")
	}

	// Generate unique identifiers
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	rName := id
	updatedName := fmt.Sprintf("%s_updated", rName)
	resourceName := "harness_chaos_infrastructure_v2.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccChaosInfrastructureV2Destroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceChaosInfrastructureV2ConfigBasic(rName, id, "KUBERNETESV2", "NAMESPACE"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rName),
				),
			},
			{
				Config: testAccResourceChaosInfrastructureV2ConfigUpdate(updatedName, id, "KUBERNETESV2", "CLUSTER"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "Updated Test Infrastructure"),
					resource.TestCheckResourceAttr(resourceName, "infra_scope", "CLUSTER"),
					resource.TestCheckResourceAttr(resourceName, "namespace", "chaos-updated"),
					resource.TestCheckResourceAttr(resourceName, "service_account", "litmus-admin"),
					resource.TestCheckResourceAttr(resourceName, "run_as_user", "1001"),
					resource.TestCheckResourceAttr(resourceName, "insecure_skip_verify", "true"),
				),
			},
		},
	})
}

func TestAccResourceChaosInfrastructureV2_KubernetesType(t *testing.T) {
	// Check for required environment variables
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	if accountId == "" {
		t.Skip("Skipping test because HARNESS_ACCOUNT_ID is not set")
	}

	// Generate unique identifiers
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	rName := id
	resourceName := "harness_chaos_infrastructure_v2.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccChaosInfrastructureV2Destroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceChaosInfrastructureV2ConfigBasic(rName, id, "KUBERNETES", "NAMESPACE"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "infra_type", "KUBERNETES"),
					resource.TestCheckResourceAttr(resourceName, "infra_scope", "NAMESPACE"),
				),
			},
		},
	})
}

func testAccChaosInfrastructureV2Destroy(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// No-op for now as we don't have a direct way to verify deletion
		return nil
	}
}

func testAccResourceChaosInfrastructureV2ImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		// Format: org_id/project_id/environment_id/infra_id
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("resource not found: %s", resourceName)
		}

		orgID := rs.Primary.Attributes["org_id"]
		projectID := rs.Primary.Attributes["project_id"]
		envID := rs.Primary.Attributes["environment_id"]
		infraID := rs.Primary.Attributes["infra_id"]

		return fmt.Sprintf("%s/%s/%s/%s", orgID, projectID, envID, infraID), nil
	}
}

func testAccResourceChaosInfrastructureV2ConfigBasic(name, id, infraType, infraScope string) string {
	// Use the account ID from environment variables
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	if accountId == "" {
		accountId = "test" // Default for test cases when not set
	}

	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
		identifier = "%[1]s"
		name       = "%[2]s"
		account_id = "%[3]s"
	}

	resource "harness_platform_project" "test" {
		identifier  = "%[1]s"
		name        = "%[2]s"
		org_id      = harness_platform_organization.test.id
		account_id  = "%[3]s"
		color       = "#0063F7"
		description = "Test project for Chaos Infrastructure"
		tags        = ["foo:bar", "baz:qux"]
	}

	resource "harness_platform_environment" "test" {
		identifier  = "%[1]s"
		name        = "%[2]s"
		org_id      = harness_platform_organization.test.id
		project_id  = harness_platform_project.test.id
		account_id  = "%[3]s"
		type        = "PreProduction"
		description = "Test environment for Chaos Infrastructure"
		tags        = ["foo:bar", "baz:qux"]
	}

	resource "harness_platform_environment" "test" {
		identifier = "%[1]s"
		name       = "%[2]s"
		org_id     = harness_platform_organization.test.id
		project_id = harness_platform_project.test.id
		account_id = "%[3]s"
		type       = "PreProduction"
	}

	resource "harness_chaos_infrastructure_v2" "test" {
		org_id              = harness_platform_organization.test.id
		project_id          = harness_platform_project.test.id
		environment_id       = harness_platform_environment.test.id
		account_id          = "%[3]s"
		name                = "%[2]s"
		infra_id            = "%[1]s"
		description         = "Test Infrastructure"
		infra_type          = "%[4]s"
		infra_scope         = "%[5]s"
		namespace           = "chaos"
		service_account     = "litmus"
		tags                = ["test:true", "chaos:true"]
		run_as_user         = 1000
		insecure_skip_verify = true
	}
	`, id, name, accountId, infraType, infraScope)
}

func testAccResourceChaosInfrastructureV2ConfigUpdate(name, id, infraType, infraScope string) string {
	// Use the account ID from environment variables
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	if accountId == "" {
		accountId = "test" // Default for test cases when not set
	}

	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
		identifier = "%[1]s"
		name       = "%[2]s"
		account_id = "%[3]s"
	}

	resource "harness_platform_project" "test" {
		identifier  = "%[1]s"
		name        = "%[2]s"
		org_id      = harness_platform_organization.test.id
		account_id  = "%[3]s"
		color       = "#0063F7"
		description = "Test project for Chaos Infrastructure"
		tags        = ["foo:bar", "baz:qux"]
	}

	resource "harness_platform_environment" "test" {
		identifier  = "%[1]s"
		name        = "%[2]s"
		org_id      = harness_platform_organization.test.id
		project_id  = harness_platform_project.test.id
		account_id  = "%[3]s"
		type        = "PreProduction"
		description = "Test environment for Chaos Infrastructure"
		tags        = ["foo:bar", "baz:qux"]
	}

	resource "harness_chaos_infrastructure_v2" "test" {
		org_id              = harness_platform_organization.test.id
		project_id          = harness_platform_project.test.id
		environment_id      = harness_platform_environment.test.id
		account_id          = "%[3]s"
		name                = "%[2]s"
		infra_id            = "%[1]s"
		description         = "Updated Test Infrastructure"
		infra_type          = "%[4]s"
		infra_scope         = "%[5]s"
		namespace           = "chaos-updated"
		service_account     = "litmus-admin"
		tags                = ["test:true", "chaos:true", "updated:true"]
		run_as_user         = 1001
		insecure_skip_verify = true
	}
	`, id, name, accountId, infraType, infraScope)
}

func TestAccResourceChaosInfrastructureV2_Import(t *testing.T) {
	// Check for required environment variables
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	if accountId == "" {
		t.Skip("Skipping test because HARNESS_ACCOUNT_ID is not set")
	}

	// Generate unique identifiers
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	rName := id
	resourceName := "harness_chaos_infrastructure_v2.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceChaosInfrastructureV2ConfigBasic(rName, id, "KUBERNETESV2", "NAMESPACE"),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccResourceChaosInfrastructureV2ImportStateIdFunc(resourceName),
			},
		},
	})
}
