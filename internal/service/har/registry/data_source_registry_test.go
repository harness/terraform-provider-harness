package registry_test

import (
	"fmt"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceVirtualRegistry(t *testing.T) {
	id := fmt.Sprintf("tf_auto_virtual_registry")
	resourceName := "data.harness_platform_har_registry.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					_, _ = acctest.TestAccGetHarClientWithContext()
				},
				Config: testAccDataSourceVirtualRegistry(id, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
		},
	})
}

func TestAccDataSourceUpstreamAWSRegistry(t *testing.T) {
	id := fmt.Sprintf("tf_auto_virtual_registry")
	resourceName := "data.harness_platform_har_registry.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					_, _ = acctest.TestAccGetHarClientWithContext()
				},
				Config: testAccDataSourceUpstreamAWSRegistry(id, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
		},
	})

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					_, _ = acctest.TestAccGetHarClientWithContext()
				},
				Config: testAccDataSourceUpstreamAWSRegistry2(id, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
		},
	})
}

func testAccDataSourceVirtualRegistry(id string, accId string) string {
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

	data "harness_platform_har_registry" "test" {
			identifier = harness_platform_har_registry.test.identifier
			space_ref = "%[2]s"
	}
`, id, accId)
}

func testAccDataSourceUpstreamAWSRegistry(id string, accId string) string {
	return fmt.Sprintf(`

	 resource "harness_platform_har_registry" "test" {
	   identifier   = "%[1]s"
	   space_ref    = "%[2]s"
	   package_type = "DOCKER"
	
	   config {
		type = "UPSTREAM"
		auth_type = "AccessKeySecretKey"
		source = "AwsEcr"
        url = "https://har-registry.default.svc.cluster.local"
		auth {
			auth_type = "AccessKeySecretKey"
			access_key_identifier  = "Secret_Token"
			access_key_secret_path = "%[2]s"
			secret_key_identifier  = "Secret_Token"
			secret_key_secret_path = "%[2]s"
		}
       }
	   parent_ref = "%[2]s"
	 }

	data "harness_platform_har_registry" "test" {
			identifier = harness_platform_har_registry.test.identifier
			space_ref = "%[2]s"
	}
`, id, accId)
}

func testAccDataSourceUpstreamAWSRegistry2(id string, accId string) string {
	return fmt.Sprintf(`

	 resource "harness_platform_har_registry" "test" {
	   identifier   = "%[1]s"
	   space_ref    = "%[2]s"
	   package_type = "DOCKER"
	
	   config {
		type = "UPSTREAM"
		auth_type = "AccessKeySecretKey"
		source = "AwsEcr"
        url = "https://har-registry.default.svc.cluster.local"
		auth {
			auth_type = "AccessKeySecretKey"
			access_key  = "MY_ACCESS_KEY"
			secret_key_identifier  = "Secret_Token"
			secret_key_secret_path = "%[2]s"
		}
       }
	   parent_ref = "%[2]s"
	 }

	data "harness_platform_har_registry" "test" {
			identifier = harness_platform_har_registry.test.identifier
			space_ref = "%[2]s"
	}
`, id, accId)
}
