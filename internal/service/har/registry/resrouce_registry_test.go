package registry_test

import (
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/harness/terraform-provider-harness/internal/service/har/registry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// randAlphanumeric generates a random alphanumeric string of specified length
// This ensures identifiers only contain lowercase letters and numbers, avoiding
// special characters like underscores that could violate API validation rules
func randAlphanumeric(n int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, n)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func testAccRegistryCheckDestroy(resourceType string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		c, ctx := acctest.TestAccGetHarClientWithContext()
		for _, rs := range s.RootModule().Resources {
			if rs.Type != resourceType {
				continue
			}
			identifier := rs.Primary.ID
			parentRef := rs.Primary.Attributes["parent_ref"]
			registryRef := parentRef + "/" + identifier
			_, _, err := c.RegistriesApi.GetRegistry(ctx, registryRef)
			if err == nil {
				return fmt.Errorf("Registry still exists: %s", identifier)
			}
		}
		return nil
	}
}

// Tests creating a virtual Docker registry at account level with import
func TestAccResourceVirtualDockerRegistry(t *testing.T) {
	id := fmt.Sprintf("tfauto_virt_docker_%s", randAlphanumeric(5))
	resourceName := "harness_platform_har_registry.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccRegistryCheckDestroy("harness_platform_har_registry"),
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					_, _ = acctest.TestAccGetHarClientWithContext()
				},
				Config: testAccResourceVirtualDockerRegistry(id, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "package_type", "DOCKER"),
					resource.TestCheckResourceAttr(resourceName, "config.0.type", "VIRTUAL"),
					resource.TestCheckResourceAttr(resourceName, "parent_ref", accountId),
					resource.TestCheckResourceAttr(resourceName, "space_ref", accountId),
					// Validate computed fields
					resource.TestCheckResourceAttrSet(resourceName, "url"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: registry.TestAccRegistryImportStateIdFunc(resourceName),
			},
		},
	})
}

// Tests creating a virtual Docker registry at organization level with import
func TestOrgResourceVirtualDockerRegistry(t *testing.T) {
	id := fmt.Sprintf("tfauto_virt_docker_%s", randAlphanumeric(5))
	resourceName := "harness_platform_har_registry.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccRegistryCheckDestroy("harness_platform_har_registry"),
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					_, _ = acctest.TestAccGetHarClientWithContext()
				},
				Config: testOrgResourceVirtualDockerRegistry(id, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "package_type", "DOCKER"),
					resource.TestCheckResourceAttr(resourceName, "config.0.type", "VIRTUAL"),
					resource.TestCheckResourceAttrSet(resourceName, "parent_ref"),
					resource.TestCheckResourceAttrSet(resourceName, "space_ref"),
					// Validate computed fields
					resource.TestCheckResourceAttrSet(resourceName, "url"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: registry.TestAccRegistryImportStateIdFunc(resourceName),
			},
		},
	})
}

// Tests creating a virtual Docker registry at project level with import
func TestProjectResourceVirtualDockerRegistry(t *testing.T) {
	id := fmt.Sprintf("tfauto_virt_docker_%s", randAlphanumeric(5))
	resourceName := "harness_platform_har_registry.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccRegistryCheckDestroy("harness_platform_har_registry"),
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					_, _ = acctest.TestAccGetHarClientWithContext()
				},
				Config: testProjResourceVirtualDockerRegistry(id, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "package_type", "DOCKER"),
					resource.TestCheckResourceAttr(resourceName, "config.0.type", "VIRTUAL"),
					resource.TestCheckResourceAttrSet(resourceName, "parent_ref"),
					resource.TestCheckResourceAttrSet(resourceName, "space_ref"),
					// Validate computed fields
					resource.TestCheckResourceAttrSet(resourceName, "url"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: registry.TestAccRegistryImportStateIdFunc(resourceName),
			},
		},
	})
}

// Tests creating an upstream Docker registry with UserPassword auth at account level
func TestAccResourceUpstreamDockerRegistry(t *testing.T) {
	id := fmt.Sprintf("tfauto_up_docker_%s", randAlphanumeric(5))
	resourceName := "harness_platform_har_registry.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccRegistryCheckDestroy("harness_platform_har_registry"),
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					_, _ = acctest.TestAccGetHarClientWithContext()
				},
				Config: testAccResourceUpstreamDockerRegistry(id, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "package_type", "DOCKER"),
					resource.TestCheckResourceAttr(resourceName, "config.0.type", "UPSTREAM"),
					resource.TestCheckResourceAttr(resourceName, "config.0.source", "Dockerhub"),
					resource.TestCheckResourceAttr(resourceName, "config.0.auth_type", "UserPassword"),
					// Validate computed fields
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
		},
	})
}

// Tests creating an upstream Docker registry with UserPassword auth at organization level
func TestOrgResourceUpstreamDockerRegistry(t *testing.T) {
	id := fmt.Sprintf("tfauto_up_docker_%s", randAlphanumeric(5))
	resourceName := "harness_platform_har_registry.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccRegistryCheckDestroy("harness_platform_har_registry"),
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					_, _ = acctest.TestAccGetHarClientWithContext()
				},
				Config: testOrgResourceUpstreamDockerRegistry(id, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "package_type", "DOCKER"),
					resource.TestCheckResourceAttr(resourceName, "config.0.type", "UPSTREAM"),
					resource.TestCheckResourceAttr(resourceName, "config.0.source", "Dockerhub"),
					resource.TestCheckResourceAttr(resourceName, "config.0.auth_type", "UserPassword"),
					// Validate computed fields
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
		},
	})
}

// Tests creating an upstream Docker registry with UserPassword auth at project level
func TestProjectResourceUpstreamDockerRegistry(t *testing.T) {
	id := fmt.Sprintf("tfauto_up_docker_%s", randAlphanumeric(5))
	resourceName := "harness_platform_har_registry.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccRegistryCheckDestroy("harness_platform_har_registry"),
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					_, _ = acctest.TestAccGetHarClientWithContext()
				},
				Config: testProjResourceUpstreamDockerRegistry(id, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "package_type", "DOCKER"),
					resource.TestCheckResourceAttr(resourceName, "config.0.type", "UPSTREAM"),
					resource.TestCheckResourceAttr(resourceName, "config.0.source", "Dockerhub"),
					resource.TestCheckResourceAttr(resourceName, "config.0.auth_type", "UserPassword"),
					// Validate computed fields
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
		},
	})
}

// Generates Terraform config for a virtual Docker registry at account level
func testAccResourceVirtualDockerRegistry(id string, accId string) string {
	return fmt.Sprintf(`

 resource "harness_platform_har_registry" "test" {
   identifier   = "%[1]s"
   space_ref    = "%[2]s"
   package_type = "DOCKER"

   config {
    type = "VIRTUAL"
   }
   parent_ref = "%[2]s"
 }
`, id, accId)
}

// Generates Terraform config for a virtual Docker registry at organization level
func testOrgResourceVirtualDockerRegistry(id string, accId string) string {
	return fmt.Sprintf(`
 resource "harness_platform_organization" "test" {
  identifier = "%[1]s_org"
  name = "%[1]s"
 }

 resource "harness_platform_har_registry" "test" {
   identifier   = "%[1]s"
   space_ref    = "%[2]s/${harness_platform_organization.test.identifier}"
   package_type = "DOCKER"

   config {
    type = "VIRTUAL"
   }
   parent_ref = "%[2]s/${harness_platform_organization.test.identifier}"
 }
`, id, accId)
}

// Generates Terraform config for a virtual Docker registry at project level
func testProjResourceVirtualDockerRegistry(id string, accId string) string {
	return fmt.Sprintf(`
 resource "harness_platform_organization" "test" {
  identifier = "%[1]s_org"
  name = "%[1]s"
 }

 resource "harness_platform_project" "test" {
  identifier = "%[1]s_project"
  name = "%[1]s"
  org_id = harness_platform_organization.test.id
  color = "#472848"
 }
 resource "harness_platform_har_registry" "test" {
   identifier   = "%[1]s"
   space_ref    = "%[2]s/${harness_platform_organization.test.identifier}/${harness_platform_project.test.identifier}"
   package_type = "DOCKER"

   config {
    type = "VIRTUAL"
   }
   parent_ref = "%[2]s/${harness_platform_organization.test.identifier}/${harness_platform_project.test.identifier}"
 }
`, id, accId)
}

// Generates Terraform config for an upstream Docker registry with UserPassword auth at account level
func testAccResourceUpstreamDockerRegistry(id string, accId string) string {
	return fmt.Sprintf(`
 resource "harness_platform_har_registry" "test" {
   identifier   = "%[1]s"
   space_ref    = "%[2]s"
   package_type = "DOCKER"

   config {
		type = "UPSTREAM"
		auth_type = "UserPassword"
		source = "Dockerhub"
		auth {
			auth_type = "UserPassword"
			user_name = "username"
			secret_identifier = "Secret_Token"
			secret_space_path = "%[2]s"
		}
   }
   parent_ref = "%[2]s"
 }
`, id, accId)
}

// Generates Terraform config for an upstream Docker registry with UserPassword auth at organization level
func testOrgResourceUpstreamDockerRegistry(id string, accId string) string {
	return fmt.Sprintf(`
 resource "harness_platform_organization" "test" {
  identifier = "%[1]s_org"
  name = "%[1]s"
 }

 resource "harness_platform_har_registry" "test" {
   identifier   = "%[1]s"
   space_ref    = "%[2]s/${harness_platform_organization.test.identifier}"
   package_type = "DOCKER"

   config {
		type = "UPSTREAM"
		auth_type = "UserPassword"
		source = "Dockerhub"
		auth {
			auth_type = "UserPassword"
			user_name = "username"
			secret_identifier = "Secret_Token"
			secret_space_path = "%[2]s"
		}
   }
   parent_ref = "%[2]s/${harness_platform_organization.test.identifier}"
 }
`, id, accId)
}

// Generates Terraform config for an upstream Docker registry with UserPassword auth at project level
func testProjResourceUpstreamDockerRegistry(id string, accId string) string {
	return fmt.Sprintf(`
 resource "harness_platform_organization" "test" {
  identifier = "%[1]s_org"
  name = "%[1]s"
 }

 resource "harness_platform_project" "test" {
  identifier = "%[1]s_project"
  name = "%[1]s"
  org_id = harness_platform_organization.test.id
  color = "#472848"
 }
 resource "harness_platform_har_registry" "test" {
   identifier   = "%[1]s"
   space_ref    = "%[2]s/${harness_platform_organization.test.identifier}/${harness_platform_project.test.identifier}"
   package_type = "DOCKER"

   config {
		type = "UPSTREAM"
		auth_type = "UserPassword"
		source = "Dockerhub"
		auth {
			auth_type = "UserPassword"
			user_name = "username"
			secret_identifier = "Secret_Token"
			secret_space_path = "%[2]s"
		}
   }
   parent_ref = "%[2]s/${harness_platform_organization.test.identifier}/${harness_platform_project.test.identifier}"
 }
`, id, accId)
}

// Tests creating a virtual Helm registry at account level
func TestAccResourceVirtualHelmRegistry(t *testing.T) {
	id := strings.ToLower(fmt.Sprintf("tfauto_virt_helm_%s", utils.RandStringBytes(5)))
	resourceName := "harness_platform_har_registry.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccRegistryCheckDestroy("harness_platform_har_registry"),
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					_, _ = acctest.TestAccGetHarClientWithContext()
				},
				Config: testAccResourceVirtualHelmRegistry(id, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "package_type", "HELM"),
					resource.TestCheckResourceAttr(resourceName, "config.0.type", "VIRTUAL"),
					resource.TestCheckResourceAttrSet(resourceName, "url"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
				),
			},
		},
	})
}

// Tests creating a virtual Helm registry at organization level
func TestOrgResourceVirtualHelmRegistry(t *testing.T) {
	id := strings.ToLower(fmt.Sprintf("tfauto_virt_helm_%s", utils.RandStringBytes(5)))
	resourceName := "harness_platform_har_registry.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccRegistryCheckDestroy("harness_platform_har_registry"),
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					_, _ = acctest.TestAccGetHarClientWithContext()
				},
				Config: testOrgResourceVirtualHelmRegistry(id, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "package_type", "HELM"),
					resource.TestCheckResourceAttr(resourceName, "config.0.type", "VIRTUAL"),
					resource.TestCheckResourceAttrSet(resourceName, "url"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
				),
			},
		},
	})
}

// Tests creating a virtual Helm registry at project level
func TestProjectResourceVirtualHelmRegistry(t *testing.T) {
	id := strings.ToLower(fmt.Sprintf("tfauto_virt_helm_%s", utils.RandStringBytes(5)))
	resourceName := "harness_platform_har_registry.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccRegistryCheckDestroy("harness_platform_har_registry"),
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					_, _ = acctest.TestAccGetHarClientWithContext()
				},
				Config: testProjResourceVirtualHelmRegistry(id, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "package_type", "HELM"),
					resource.TestCheckResourceAttr(resourceName, "config.0.type", "VIRTUAL"),
					resource.TestCheckResourceAttrSet(resourceName, "url"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
				),
			},
		},
	})
}

// Tests creating an upstream Helm registry with custom URL at account level
func TestAccResourceUpstreamHelmRegistry(t *testing.T) {
	id := fmt.Sprintf("tfauto_up_helm_%s", randAlphanumeric(5))
	resourceName := "harness_platform_har_registry.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccRegistryCheckDestroy("harness_platform_har_registry"),
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					_, _ = acctest.TestAccGetHarClientWithContext()
				},
				Config: testAccResourceUpstreamHelmRegistry(id, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "package_type", "HELM"),
					resource.TestCheckResourceAttr(resourceName, "config.0.type", "UPSTREAM"),
					resource.TestCheckResourceAttr(resourceName, "config.0.source", "Custom"),
					resource.TestCheckResourceAttrSet(resourceName, "config.0.url"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
				),
			},
		},
	})
}

// Tests creating an upstream Helm registry with custom URL at organization level
func TestOrgResourceUpstreamHelmRegistry(t *testing.T) {
	id := fmt.Sprintf("tfauto_up_helm_%s", randAlphanumeric(5))
	resourceName := "harness_platform_har_registry.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccRegistryCheckDestroy("harness_platform_har_registry"),
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					_, _ = acctest.TestAccGetHarClientWithContext()
				},
				Config: testOrgResourceUpstreamHelmRegistry(id, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "package_type", "HELM"),
					resource.TestCheckResourceAttr(resourceName, "config.0.type", "UPSTREAM"),
					resource.TestCheckResourceAttr(resourceName, "config.0.source", "Custom"),
					resource.TestCheckResourceAttrSet(resourceName, "config.0.url"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
				),
			},
		},
	})
}

// Tests creating an upstream Helm registry with custom URL at project level
func TestProjectResourceUpstreamHelmRegistry(t *testing.T) {
	id := fmt.Sprintf("tfauto_up_helm_%s", randAlphanumeric(5))
	resourceName := "harness_platform_har_registry.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccRegistryCheckDestroy("harness_platform_har_registry"),
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					_, _ = acctest.TestAccGetHarClientWithContext()
				},
				Config: testProjResourceUpstreamHelmRegistry(id, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "package_type", "HELM"),
					resource.TestCheckResourceAttr(resourceName, "config.0.type", "UPSTREAM"),
					resource.TestCheckResourceAttr(resourceName, "config.0.source", "Custom"),
					resource.TestCheckResourceAttrSet(resourceName, "config.0.url"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
				),
			},
		},
	})
}

// Generates Terraform config for a virtual Helm registry at account level
func testAccResourceVirtualHelmRegistry(id string, accId string) string {
	return fmt.Sprintf(`

 resource "harness_platform_har_registry" "test" {
   identifier   = "%[1]s"
   space_ref    = "%[2]s"
   package_type = "HELM"

   config {
    type = "VIRTUAL"
   }
   parent_ref = "%[2]s"
 }
`, id, accId)
}

// Generates Terraform config for a virtual Helm registry at organization level
func testOrgResourceVirtualHelmRegistry(id string, accId string) string {
	return fmt.Sprintf(`
 resource "harness_platform_organization" "test" {
  identifier = "%[1]s_org"
  name = "%[1]s"
 }

 resource "harness_platform_har_registry" "test" {
   identifier   = "%[1]s"
   space_ref    = "%[2]s/${harness_platform_organization.test.identifier}"
   package_type = "HELM"

   config {
    type = "VIRTUAL"
   }
   parent_ref = "%[2]s/${harness_platform_organization.test.identifier}"
 }
`, id, accId)
}

// Generates Terraform config for a virtual Helm registry at project level
func testProjResourceVirtualHelmRegistry(id string, accId string) string {
	return fmt.Sprintf(`
 resource "harness_platform_organization" "test" {
  identifier = "%[1]s_org"
  name = "%[1]s"
 }

 resource "harness_platform_project" "test" {
  identifier = "%[1]s_project"
  name = "%[1]s"
  org_id = harness_platform_organization.test.id
  color = "#472848"
 }
 resource "harness_platform_har_registry" "test" {
   identifier   = "%[1]s"
   space_ref    = "%[2]s/${harness_platform_organization.test.identifier}/${harness_platform_project.test.identifier}"
   package_type = "HELM"

   config {
    type = "VIRTUAL"
   }
   parent_ref = "%[2]s/${harness_platform_organization.test.identifier}/${harness_platform_project.test.identifier}"
 }
`, id, accId)
}

// Generates Terraform config for an upstream Helm registry with custom URL at account level
func testAccResourceUpstreamHelmRegistry(id string, accId string) string {
	return fmt.Sprintf(`
 resource "harness_platform_har_registry" "test" {
   identifier   = "%[1]s"
   space_ref    = "%[2]s"
   package_type = "HELM"

   config {
		type = "UPSTREAM"
		auth_type = "UserPassword"
		source = "Custom"
		url = "https://har-registry.default.svc.cluster.local"
		auth {
			auth_type = "UserPassword"
			user_name = "username"
			secret_identifier = "Secret_Token"
			secret_space_path = "%[2]s"
		}
   }
   parent_ref = "%[2]s"
 }
`, id, accId)
}

// Generates Terraform config for an upstream Helm registry with custom URL at organization level
func testOrgResourceUpstreamHelmRegistry(id string, accId string) string {
	return fmt.Sprintf(`
 resource "harness_platform_organization" "test" {
  identifier = "%[1]s_org"
  name = "%[1]s"
 }

 resource "harness_platform_har_registry" "test" {
   identifier   = "%[1]s"
   space_ref    = "%[2]s/${harness_platform_organization.test.identifier}"
   package_type = "HELM"

   config {
		type = "UPSTREAM"
		auth_type = "UserPassword"
		source = "Custom"
		url = "https://har-registry.default.svc.cluster.local"
		auth {
			auth_type = "UserPassword"
			user_name = "username"
			secret_identifier = "Secret_Token"
			secret_space_path = "%[2]s"
		}
   }
   parent_ref = "%[2]s/${harness_platform_organization.test.identifier}"
 }
`, id, accId)
}

// Generates Terraform config for an upstream Helm registry with custom URL at project level
func testProjResourceUpstreamHelmRegistry(id string, accId string) string {
	return fmt.Sprintf(`
 resource "harness_platform_organization" "test" {
  identifier = "%[1]s_org"
  name = "%[1]s"
 }

 resource "harness_platform_project" "test" {
  identifier = "%[1]s_project"
  name = "%[1]s"
  org_id = harness_platform_organization.test.id
  color = "#472848"
 }
 resource "harness_platform_har_registry" "test" {
   identifier   = "%[1]s"
   space_ref    = "%[2]s/${harness_platform_organization.test.identifier}/${harness_platform_project.test.identifier}"
   package_type = "HELM"

   config {
		type = "UPSTREAM"
		auth_type = "UserPassword"
		source = "Custom"
		url = "https://har-registry.default.svc.cluster.local"
		auth {
			auth_type = "UserPassword"
			user_name = "username"
			secret_identifier = "Secret_Token"
			secret_space_path = "%[2]s"
		}
   }
   parent_ref = "%[2]s/${harness_platform_organization.test.identifier}/${harness_platform_project.test.identifier}"
 }
`, id, accId)
}

// Tests creating an upstream Docker registry with Anonymous auth at account level
func TestAccResourceUpstreamDockerAnonymousRegistry(t *testing.T) {
	id := fmt.Sprintf("tfauto_up_docker_anon_%s", randAlphanumeric(5))
	resourceName := "harness_platform_har_registry.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccRegistryCheckDestroy("harness_platform_har_registry"),
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					_, _ = acctest.TestAccGetHarClientWithContext()
				},
				Config:             testAccResourceUpstreamDockerAnonymousRegistry(id, accountId),
				ExpectNonEmptyPlan: true, // TODO: Investigate why Anonymous auth causes drift
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "package_type", "DOCKER"),
					resource.TestCheckResourceAttr(resourceName, "config.0.type", "UPSTREAM"),
					resource.TestCheckResourceAttr(resourceName, "config.0.source", "Dockerhub"),
					resource.TestCheckResourceAttr(resourceName, "config.0.auth_type", "Anonymous"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
				),
			},
		},
	})
}

// Tests creating an upstream Docker registry with Anonymous auth at organization level
func TestOrgResourceUpstreamDockerAnonymousRegistry(t *testing.T) {
	id := fmt.Sprintf("tfauto_up_docker_anon_%s", randAlphanumeric(5))
	resourceName := "harness_platform_har_registry.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccRegistryCheckDestroy("harness_platform_har_registry"),
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					_, _ = acctest.TestAccGetHarClientWithContext()
				},
				Config:             testOrgResourceUpstreamDockerAnonymousRegistry(id, accountId),
				ExpectNonEmptyPlan: true, // TODO: Investigate why Anonymous auth causes drift
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "package_type", "DOCKER"),
					resource.TestCheckResourceAttr(resourceName, "config.0.type", "UPSTREAM"),
					resource.TestCheckResourceAttr(resourceName, "config.0.source", "Dockerhub"),
					resource.TestCheckResourceAttr(resourceName, "config.0.auth_type", "Anonymous"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
				),
			},
		},
	})
}

// Tests creating an upstream Docker registry with Anonymous auth at project level
func TestProjectResourceUpstreamDockerAnonymousRegistry(t *testing.T) {
	id := fmt.Sprintf("tfauto_up_docker_anon_%s", randAlphanumeric(5))
	resourceName := "harness_platform_har_registry.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccRegistryCheckDestroy("harness_platform_har_registry"),
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					_, _ = acctest.TestAccGetHarClientWithContext()
				},
				Config:             testProjResourceUpstreamDockerAnonymousRegistry(id, accountId),
				ExpectNonEmptyPlan: true, // TODO: Investigate why Anonymous auth causes drift
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "package_type", "DOCKER"),
					resource.TestCheckResourceAttr(resourceName, "config.0.type", "UPSTREAM"),
					resource.TestCheckResourceAttr(resourceName, "config.0.source", "Dockerhub"),
					resource.TestCheckResourceAttr(resourceName, "config.0.auth_type", "Anonymous"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
				),
			},
		},
	})
}

// Generates Terraform config for an upstream Docker registry with Anonymous auth at account level
func testAccResourceUpstreamDockerAnonymousRegistry(id string, accId string) string {
	return fmt.Sprintf(`
 resource "harness_platform_har_registry" "test" {
   identifier   = "%[1]s"
   space_ref    = "%[2]s"
   package_type = "DOCKER"

   config {
		type = "UPSTREAM"
		auth_type = "Anonymous"
		source = "Dockerhub"
		auth {
			auth_type = "Anonymous"
		}
   }
   parent_ref = "%[2]s"
 }
`, id, accId)
}

// Generates Terraform config for an upstream Docker registry with Anonymous auth at organization level
func testOrgResourceUpstreamDockerAnonymousRegistry(id string, accId string) string {
	return fmt.Sprintf(`
 resource "harness_platform_organization" "test" {
  identifier = "%[1]s_org"
  name = "%[1]s"
 }

 resource "harness_platform_har_registry" "test" {
   identifier   = "%[1]s"
   space_ref    = "%[2]s/${harness_platform_organization.test.identifier}"
   package_type = "DOCKER"

   config {
		type = "UPSTREAM"
		auth_type = "Anonymous"
		source = "Dockerhub"
		auth {
			auth_type = "Anonymous"
		}
   }
   parent_ref = "%[2]s/${harness_platform_organization.test.identifier}"
 }
`, id, accId)
}

// Generates Terraform config for an upstream Docker registry with Anonymous auth at project level
func testProjResourceUpstreamDockerAnonymousRegistry(id string, accId string) string {
	return fmt.Sprintf(`
 resource "harness_platform_organization" "test" {
  identifier = "%[1]s_org"
  name = "%[1]s"
 }

 resource "harness_platform_project" "test" {
  identifier = "%[1]s_project"
  name = "%[1]s"
  org_id = harness_platform_organization.test.id
  color = "#472848"
 }
 resource "harness_platform_har_registry" "test" {
   identifier   = "%[1]s"
   space_ref    = "%[2]s/${harness_platform_organization.test.identifier}/${harness_platform_project.test.identifier}"
   package_type = "DOCKER"

   config {
		type = "UPSTREAM"
		auth_type = "Anonymous"
		source = "Dockerhub"
		auth {
			auth_type = "Anonymous"
		}
   }
   parent_ref = "%[2]s/${harness_platform_organization.test.identifier}/${harness_platform_project.test.identifier}"
 }
`, id, accId)
}

// Tests creating an upstream Helm registry with Anonymous auth at account level
func TestAccResourceUpstreamHelmAnonymousRegistry(t *testing.T) {
	id := strings.ToLower(fmt.Sprintf("tfauto_up_helm_anon_%s", utils.RandStringBytes(5)))
	resourceName := "harness_platform_har_registry.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccRegistryCheckDestroy("harness_platform_har_registry"),
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					_, _ = acctest.TestAccGetHarClientWithContext()
				},
				Config:             testAccResourceUpstreamHelmAnonymousRegistry(id, accountId),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
		},
	})
}

// Tests creating an upstream Helm registry with Anonymous auth at organization level
func TestOrgResourceUpstreamHelmAnonymousRegistry(t *testing.T) {
	id := strings.ToLower(fmt.Sprintf("tfauto_up_helm_anon_%s", utils.RandStringBytes(5)))
	resourceName := "harness_platform_har_registry.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccRegistryCheckDestroy("harness_platform_har_registry"),
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					_, _ = acctest.TestAccGetHarClientWithContext()
				},
				Config:             testOrgResourceUpstreamHelmAnonymousRegistry(id, accountId),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
		},
	})
}

// Tests creating an upstream Helm registry with Anonymous auth at project level
func TestProjectResourceUpstreamHelmAnonymousRegistry(t *testing.T) {
	id := strings.ToLower(fmt.Sprintf("tfauto_up_helm_anon_%s", utils.RandStringBytes(5)))
	resourceName := "harness_platform_har_registry.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccRegistryCheckDestroy("harness_platform_har_registry"),
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					_, _ = acctest.TestAccGetHarClientWithContext()
				},
				Config:             testProjResourceUpstreamHelmAnonymousRegistry(id, accountId),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
		},
	})
}

// Generates Terraform config for an upstream Helm registry with Anonymous auth at account level
func testAccResourceUpstreamHelmAnonymousRegistry(id string, accId string) string {
	return fmt.Sprintf(`
 resource "harness_platform_har_registry" "test" {
   identifier   = "%[1]s"
   space_ref    = "%[2]s"
   package_type = "HELM"

   config {
		type = "UPSTREAM"
		auth_type = "Anonymous"
		source = "Custom"
		url = "https://har-registry.default.svc.cluster.local"
		auth {
			auth_type = "Anonymous"
		}
   }
   parent_ref = "%[2]s"
 }
`, id, accId)
}

// Generates Terraform config for an upstream Helm registry with Anonymous auth at organization level
func testOrgResourceUpstreamHelmAnonymousRegistry(id string, accId string) string {
	return fmt.Sprintf(`
 resource "harness_platform_organization" "test" {
  identifier = "%[1]s_org"
  name = "%[1]s"
 }

 resource "harness_platform_har_registry" "test" {
   identifier   = "%[1]s"
   space_ref    = "%[2]s/${harness_platform_organization.test.identifier}"
   package_type = "HELM"

   config {
		type = "UPSTREAM"
		auth_type = "Anonymous"
		source = "Custom"
		url = "https://har-registry.default.svc.cluster.local"
		auth {
			auth_type = "Anonymous"
		}
   }
   parent_ref = "%[2]s/${harness_platform_organization.test.identifier}"
 }
`, id, accId)
}

// Generates Terraform config for an upstream Helm registry with Anonymous auth at project level
func testProjResourceUpstreamHelmAnonymousRegistry(id string, accId string) string {
	return fmt.Sprintf(`
 resource "harness_platform_organization" "test" {
  identifier = "%[1]s_org"
  name = "%[1]s"
 }

 resource "harness_platform_project" "test" {
  identifier = "%[1]s_project"
  name = "%[1]s"
  org_id = harness_platform_organization.test.id
  color = "#472848"
 }
 resource "harness_platform_har_registry" "test" {
   identifier   = "%[1]s"
   space_ref    = "%[2]s/${harness_platform_organization.test.identifier}/${harness_platform_project.test.identifier}"
   package_type = "HELM"

   config {
		type = "UPSTREAM"
		auth_type = "Anonymous"
		source = "Custom"
		url = "https://har-registry.default.svc.cluster.local"
		auth {
			auth_type = "Anonymous"
		}
   }
   parent_ref = "%[2]s/${harness_platform_organization.test.identifier}/${harness_platform_project.test.identifier}"
 }
`, id, accId)
}

// Tests updating a virtual Docker registry's description field
func TestAccResourceVirtualDockerRegistryUpdate(t *testing.T) {
	id := fmt.Sprintf("tfauto_virt_docker_upd_%s", randAlphanumeric(5))
	resourceName := "harness_platform_har_registry.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccRegistryCheckDestroy("harness_platform_har_registry"),
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					_, _ = acctest.TestAccGetHarClientWithContext()
				},
				Config: testAccResourceVirtualDockerRegistryWithDescription(id, accountId, "Initial description"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "description", "Initial description"),
					resource.TestCheckResourceAttr(resourceName, "package_type", "DOCKER"),
				),
			},
			{
				Config: testAccResourceVirtualDockerRegistryWithDescription(id, accountId, "Updated description"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "description", "Updated description"),
					resource.TestCheckResourceAttr(resourceName, "package_type", "DOCKER"),
				),
			},
		},
	})
}

// Generates Terraform config for a virtual Docker registry with a custom description
func testAccResourceVirtualDockerRegistryWithDescription(id string, accId string, description string) string {
	return fmt.Sprintf(`
resource "harness_platform_har_registry" "test" {
  identifier   = "%[1]s"
  description  = "%[3]s"
  space_ref    = "%[2]s"
  package_type = "DOCKER"

  config {
    type = "VIRTUAL"
  }
  parent_ref = "%[2]s"
}
`, id, accId, description)
}

// Tests creating a virtual Maven registry with import
func TestAccResourceVirtualMavenRegistry(t *testing.T) {
	id := strings.ToLower(fmt.Sprintf("tfauto_virt_maven_%s", utils.RandStringBytes(5)))
	resourceName := "harness_platform_har_registry.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccRegistryCheckDestroy("harness_platform_har_registry"),
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					_, _ = acctest.TestAccGetHarClientWithContext()
				},
				Config: testAccResourceVirtualMavenRegistry(id, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "package_type", "MAVEN"),
					resource.TestCheckResourceAttr(resourceName, "config.0.type", "VIRTUAL"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: registry.TestAccRegistryImportStateIdFunc(resourceName),
			},
		},
	})
}

// Generates Terraform config for a virtual Maven registry
func testAccResourceVirtualMavenRegistry(id string, accId string) string {
	return fmt.Sprintf(`
 resource "harness_platform_har_registry" "test" {
   identifier   = "%[1]s"
   space_ref    = "%[2]s"
   package_type = "MAVEN"

   config {
    type = "VIRTUAL"
   }
   parent_ref = "%[2]s"
 }
`, id, accId)
}

// Tests creating a virtual NPM registry with import
func TestAccResourceVirtualNPMRegistry(t *testing.T) {
	id := strings.ToLower(fmt.Sprintf("tfauto_virt_npm_%s", utils.RandStringBytes(5)))
	resourceName := "harness_platform_har_registry.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccRegistryCheckDestroy("harness_platform_har_registry"),
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					_, _ = acctest.TestAccGetHarClientWithContext()
				},
				Config: testAccResourceVirtualNPMRegistry(id, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "package_type", "NPM"),
					resource.TestCheckResourceAttr(resourceName, "config.0.type", "VIRTUAL"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: registry.TestAccRegistryImportStateIdFunc(resourceName),
			},
		},
	})
}

// Generates Terraform config for a virtual NPM registry
func testAccResourceVirtualNPMRegistry(id string, accId string) string {
	return fmt.Sprintf(`
 resource "harness_platform_har_registry" "test" {
   identifier   = "%[1]s"
   space_ref    = "%[2]s"
   package_type = "NPM"

   config {
    type = "VIRTUAL"
   }
   parent_ref = "%[2]s"
 }
`, id, accId)
}

// Tests creating a virtual Generic registry with import
func TestAccResourceVirtualGenericRegistry(t *testing.T) {
	id := strings.ToLower(fmt.Sprintf("tfauto_virt_generic_%s", utils.RandStringBytes(5)))
	resourceName := "harness_platform_har_registry.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccRegistryCheckDestroy("harness_platform_har_registry"),
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					_, _ = acctest.TestAccGetHarClientWithContext()
				},
				Config: testAccResourceVirtualGenericRegistry(id, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "package_type", "GENERIC"),
					resource.TestCheckResourceAttr(resourceName, "config.0.type", "VIRTUAL"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: registry.TestAccRegistryImportStateIdFunc(resourceName),
			},
		},
	})
}

// Generates Terraform config for a virtual Generic registry
func testAccResourceVirtualGenericRegistry(id string, accId string) string {
	return fmt.Sprintf(`
 resource "harness_platform_har_registry" "test" {
   identifier   = "%[1]s"
   space_ref    = "%[2]s"
   package_type = "GENERIC"

   config {
    type = "VIRTUAL"
   }
   parent_ref = "%[2]s"
 }
`, id, accId)
}

// Tests creating a virtual Docker registry with an upstream proxy reference
func TestAccResourceVirtualDockerRegistryWithUpstreamProxy(t *testing.T) {
	rand := utils.RandStringBytes(5)
	id := strings.ToLower(fmt.Sprintf("tfauto_virt_up_%s", rand))
	upstreamId := strings.ToLower(fmt.Sprintf("tfauto_upsrc_%s", rand))
	resourceName := "harness_platform_har_registry.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccRegistryCheckDestroy("harness_platform_har_registry"),
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					_, _ = acctest.TestAccGetHarClientWithContext()
				},
				Config: testAccResourceVirtualDockerRegistryWithUpstreamProxy(id, upstreamId, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "package_type", "DOCKER"),
					resource.TestCheckResourceAttr(resourceName, "config.0.type", "VIRTUAL"),
					resource.TestCheckResourceAttr(resourceName, "config.0.upstream_proxies.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "config.0.upstream_proxies.0", upstreamId),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: registry.TestAccRegistryImportStateIdFunc(resourceName),
			},
		},
	})
}

// Generates Terraform config for a virtual Docker registry with an upstream proxy
func testAccResourceVirtualDockerRegistryWithUpstreamProxy(id string, upstreamId string, accId string) string {
	return fmt.Sprintf(`
 resource "harness_platform_har_registry" "upstream" {
   identifier   = "%[2]s"
   space_ref    = "%[3]s"
   package_type = "DOCKER"

   config {
		type = "UPSTREAM"
		auth_type = "Anonymous"
		source = "Dockerhub"
   }
   parent_ref = "%[3]s"
 }

 resource "harness_platform_har_registry" "test" {
   identifier   = "%[1]s"
   space_ref    = "%[3]s"
   package_type = "DOCKER"

   config {
		type = "VIRTUAL"
		upstream_proxies = [harness_platform_har_registry.upstream.identifier]
   }
   parent_ref = "%[3]s"
   
   depends_on = [harness_platform_har_registry.upstream]
 }
`, id, upstreamId, accId)
}

// Tests creating a virtual Docker registry with blocked patterns
func TestAccResourceVirtualDockerRegistryWithBlockedPatterns(t *testing.T) {
	id := strings.ToLower(fmt.Sprintf("tfauto_virt_blocked_%s", utils.RandStringBytes(5)))
	resourceName := "harness_platform_har_registry.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccRegistryCheckDestroy("harness_platform_har_registry"),
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					_, _ = acctest.TestAccGetHarClientWithContext()
				},
				Config: testAccResourceVirtualDockerRegistryWithBlockedPatterns(id, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "package_type", "DOCKER"),
					resource.TestCheckResourceAttr(resourceName, "config.0.type", "VIRTUAL"),
					resource.TestCheckResourceAttr(resourceName, "blocked_pattern.#", "2"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: registry.TestAccRegistryImportStateIdFunc(resourceName),
			},
		},
	})
}

// Generates Terraform config for a virtual Docker registry with blocked patterns
func testAccResourceVirtualDockerRegistryWithBlockedPatterns(id string, accId string) string {
	return fmt.Sprintf(`
 resource "harness_platform_har_registry" "test" {
   identifier   = "%[1]s"
   space_ref    = "%[2]s"
   package_type = "DOCKER"
   blocked_pattern = ["*-SNAPSHOT", "*.alpha"]

   config {
		type = "VIRTUAL"
   }
   parent_ref = "%[2]s"
 }
`, id, accId)
}

// Tests creating a virtual Docker registry with allowed patterns
func TestAccResourceVirtualDockerRegistryWithAllowedPatterns(t *testing.T) {
	id := strings.ToLower(fmt.Sprintf("tfauto_virt_allowed_%s", utils.RandStringBytes(5)))
	resourceName := "harness_platform_har_registry.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccRegistryCheckDestroy("harness_platform_har_registry"),
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					_, _ = acctest.TestAccGetHarClientWithContext()
				},
				Config: testAccResourceVirtualDockerRegistryWithAllowedPatterns(id, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "package_type", "DOCKER"),
					resource.TestCheckResourceAttr(resourceName, "config.0.type", "VIRTUAL"),
					resource.TestCheckResourceAttr(resourceName, "allowed_pattern.#", "2"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: registry.TestAccRegistryImportStateIdFunc(resourceName),
			},
		},
	})
}

// Generates Terraform config for a virtual Docker registry with allowed patterns
func testAccResourceVirtualDockerRegistryWithAllowedPatterns(id string, accId string) string {
	return fmt.Sprintf(`
 resource "harness_platform_har_registry" "test" {
   identifier   = "%[1]s"
   space_ref    = "%[2]s"
   package_type = "DOCKER"
   allowed_pattern = ["*-release", "*.stable"]

   config {
		type = "VIRTUAL"
   }
   parent_ref = "%[2]s"
 }
`, id, accId)
}

// Tests adding and removing upstream_proxies from a virtual registry
func TestAccResourceVirtualDockerRegistryUpdateUpstreamProxies(t *testing.T) {
	rand := utils.RandStringBytes(5)
	id := strings.ToLower(fmt.Sprintf("tfauto_virt_upprx_%s", rand))
	upstream1 := strings.ToLower(fmt.Sprintf("tfauto_up1_%s", rand))
	upstream2 := strings.ToLower(fmt.Sprintf("tfauto_up2_%s", rand))
	resourceName := "harness_platform_har_registry.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccRegistryCheckDestroy("harness_platform_har_registry"),
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					_, _ = acctest.TestAccGetHarClientWithContext()
				},
				Config: registry.TestAccResourceVirtualDockerRegistryNoUpstream(id, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "config.0.type", "VIRTUAL"),
					resource.TestCheckResourceAttr(resourceName, "config.0.upstream_proxies.#", "0"),
				),
			},
			{
				Config: registry.TestAccResourceVirtualDockerRegistryWithOneUpstream(id, upstream1, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "config.0.upstream_proxies.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "config.0.upstream_proxies.0", upstream1),
				),
			},
			{
				Config: registry.TestAccResourceVirtualDockerRegistryWithTwoUpstreams(id, upstream1, upstream2, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "config.0.upstream_proxies.#", "2"),
				),
			},
			{
				Config: registry.TestAccResourceVirtualDockerRegistryWithUpstreamsButNotUsed(id, upstream1, upstream2, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "config.0.upstream_proxies.#", "0"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: registry.TestAccRegistryImportStateIdFunc(resourceName),
			},
		},
	})
}

// Tests adding and removing allowed_pattern from a virtual registry
func TestAccResourceVirtualDockerRegistryUpdateAllowedPattern(t *testing.T) {
	id := strings.ToLower(fmt.Sprintf("tfauto_virt_updal_%s", utils.RandStringBytes(5)))
	resourceName := "harness_platform_har_registry.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccRegistryCheckDestroy("harness_platform_har_registry"),
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					_, _ = acctest.TestAccGetHarClientWithContext()
				},
				Config: registry.TestAccResourceVirtualDockerRegistryNoPatterns(id, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "allowed_pattern.#", "0"),
				),
			},
			{
				Config: testAccResourceVirtualDockerRegistryWithAllowedPatterns(id, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "allowed_pattern.#", "2"),
				),
			},
			{
				Config: registry.TestAccResourceVirtualDockerRegistryNoPatterns(id, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "allowed_pattern.#", "0"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: registry.TestAccRegistryImportStateIdFunc(resourceName),
			},
		},
	})
}

// Tests adding and removing blocked_pattern from a virtual registry
func TestAccResourceVirtualDockerRegistryUpdateBlockedPattern(t *testing.T) {
	id := strings.ToLower(fmt.Sprintf("tfauto_virt_updbl_%s", utils.RandStringBytes(5)))
	resourceName := "harness_platform_har_registry.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccRegistryCheckDestroy("harness_platform_har_registry"),
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					_, _ = acctest.TestAccGetHarClientWithContext()
				},
				Config: registry.TestAccResourceVirtualDockerRegistryNoPatterns(id, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "blocked_pattern.#", "0"),
				),
			},
			{
				Config: testAccResourceVirtualDockerRegistryWithBlockedPatterns(id, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "blocked_pattern.#", "2"),
				),
			},
			{
				Config: registry.TestAccResourceVirtualDockerRegistryNoPatterns(id, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "blocked_pattern.#", "0"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: registry.TestAccRegistryImportStateIdFunc(resourceName),
			},
		},
	})
}

// Tests virtual registry as LOCAL storage (no upstream proxies) with import
func TestAccResourceVirtualDockerRegistryAsLocal(t *testing.T) {
	id := strings.ToLower(fmt.Sprintf("tfauto_virt_local_%s", utils.RandStringBytes(5)))
	resourceName := "harness_platform_har_registry.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccRegistryCheckDestroy("harness_platform_har_registry"),
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					_, _ = acctest.TestAccGetHarClientWithContext()
				},
				Config: registry.TestAccResourceVirtualDockerRegistryNoUpstream(id, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "package_type", "DOCKER"),
					resource.TestCheckResourceAttr(resourceName, "config.0.type", "VIRTUAL"),
					resource.TestCheckResourceAttr(resourceName, "config.0.upstream_proxies.#", "0"),
					resource.TestCheckResourceAttrSet(resourceName, "url"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: registry.TestAccRegistryImportStateIdFunc(resourceName),
			},
		},
	})
}

// TestAccResourceUpstreamDockerRegistryNoAuth tests that creating an UPSTREAM registry
// without authentication should fail with validation error.
func TestAccResourceUpstreamDockerRegistryNoAuth(t *testing.T) {
	id := strings.ToLower(fmt.Sprintf("tfauto_up_noauth_%s", utils.RandStringBytes(5)))
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccRegistryCheckDestroy("harness_platform_har_registry"),
		Steps: []resource.TestStep{
			{
				Config:      testAccResourceUpstreamDockerRegistryNoAuth(id, accountId),
				ExpectError: regexp.MustCompile("authentication is required for UPSTREAM registry type"),
			},
		},
	})
}

// TestAccResourceUpstreamHelmRegistryNoAuth tests that creating an UPSTREAM Helm registry
// without authentication should fail with validation error.
func TestAccResourceUpstreamHelmRegistryNoAuth(t *testing.T) {
	id := strings.ToLower(fmt.Sprintf("tfauto_uphelm_noauth_%s", utils.RandStringBytes(5)))
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccRegistryCheckDestroy("harness_platform_har_registry"),
		Steps: []resource.TestStep{
			{
				Config:      testAccResourceUpstreamHelmRegistryNoAuth(id, accountId),
				ExpectError: regexp.MustCompile("authentication is required for UPSTREAM registry type"),
			},
		},
	})
}

// TestAccResourceUpstreamMavenRegistryNoAuth tests that creating an UPSTREAM Maven registry
// without authentication should fail with validation error.
func TestAccResourceUpstreamMavenRegistryNoAuth(t *testing.T) {
	id := strings.ToLower(fmt.Sprintf("tfauto_upmaven_noauth_%s", utils.RandStringBytes(5)))
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccRegistryCheckDestroy("harness_platform_har_registry"),
		Steps: []resource.TestStep{
			{
				Config:      testAccResourceUpstreamMavenRegistryNoAuth(id, accountId),
				ExpectError: regexp.MustCompile("authentication is required for UPSTREAM registry type"),
			},
		},
	})
}

func testAccResourceUpstreamDockerRegistryNoAuth(id string, accId string) string {
	return fmt.Sprintf(`
resource "harness_platform_har_registry" "test" {
  identifier   = "%[1]s"
  space_ref    = "%[2]s"
  package_type = "DOCKER"

  config {
    type   = "UPSTREAM"
    source = "Dockerhub"
  }
  parent_ref = "%[2]s"
}
`, id, accId)
}

func testAccResourceUpstreamHelmRegistryNoAuth(id string, accId string) string {
	return fmt.Sprintf(`
resource "harness_platform_har_registry" "test" {
  identifier   = "%[1]s"
  space_ref    = "%[2]s"
  package_type = "HELM"

  config {
    type   = "UPSTREAM"
    source = "Custom"
    url    = "https://charts.example.com"
  }
  parent_ref = "%[2]s"
}
`, id, accId)
}

func testAccResourceUpstreamMavenRegistryNoAuth(id string, accId string) string {
	return fmt.Sprintf(`
resource "harness_platform_har_registry" "test" {
  identifier   = "%[1]s"
  space_ref    = "%[2]s"
  package_type = "MAVEN"

  config {
    type   = "UPSTREAM"
    source = "MavenCentral"
  }
  parent_ref = "%[2]s"
}
`, id, accId)
}
