package registry_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/harness/harness-go-sdk/harness/har"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/harness/terraform-provider-harness/internal/service/har/registry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// allPackageTypes is every value of the har.PackageType enum. The SDK does not
// expose an iterable list, so it is mirrored here on purpose: the unit tests
// below fail when the SDK gains a package type that the provider schema has not
// been taught to accept yet.
var allPackageTypes = []har.PackageType{
	har.DOCKER_PackageType,
	har.MAVEN_PackageType,
	har.PYTHON_PackageType,
	har.GENERIC_PackageType,
	har.HELM_PackageType,
	har.HELM_HTTP_PackageType,
	har.NUGET_PackageType,
	har.NPM_PackageType,
	har.RPM_PackageType,
	har.CARGO_PackageType,
	har.COMPOSER_PackageType,
	har.GO_PackageType,
	har.HUGGINGFACE_PackageType,
	har.CONDA_PackageType,
	har.DART_PackageType,
	har.SWIFT_PackageType,
	har.PUPPET_PackageType,
	har.RAW_PackageType,
	har.DEBIAN_PackageType,
	har.CONAN_PackageType,
	har.RUBY_PackageType,
	har.TERRAFORM_PackageType,
	har.CRAN_PackageType,
	har.ALPINE_PackageType,
	har.WOLFI_PackageType,
}

// newPackageTypes are the package types added in AH-5667.
var newPackageTypes = []har.PackageType{
	har.COMPOSER_PackageType,
	har.HUGGINGFACE_PackageType,
	har.DART_PackageType,
	har.SWIFT_PackageType,
}

func packageTypeValidateFunc(t *testing.T, s *schema.Resource) schema.SchemaValidateFunc {
	t.Helper()
	packageType, ok := s.Schema["package_type"]
	if !ok {
		t.Fatal("schema has no package_type attribute")
	}
	if packageType.ValidateFunc == nil {
		t.Fatal("package_type has no ValidateFunc")
	}
	return packageType.ValidateFunc
}

// Verifies the resource schema accepts every package type the SDK knows about.
func TestResourceRegistryPackageTypeValidation(t *testing.T) {
	validate := packageTypeValidateFunc(t, registry.ResourceRegistry())

	for _, packageType := range allPackageTypes {
		t.Run(string(packageType), func(t *testing.T) {
			_, errs := validate((string)(packageType), "package_type")
			if len(errs) > 0 {
				t.Errorf("package_type %q rejected by resource schema: %v", packageType, errs)
			}
		})
	}
}

// Verifies the read-only (data source) schema accepts the same package types as
// the resource schema. The two lists are maintained separately in schema.go, so
// they drift easily.
func TestDataSourceRegistryPackageTypeValidation(t *testing.T) {
	validate := packageTypeValidateFunc(t, registry.DataSourceRegistry())

	for _, packageType := range allPackageTypes {
		t.Run(string(packageType), func(t *testing.T) {
			_, errs := validate((string)(packageType), "package_type")
			if len(errs) > 0 {
				t.Errorf("package_type %q rejected by data source schema: %v", packageType, errs)
			}
		})
	}
}

// Verifies an unknown package type is still rejected.
func TestResourceRegistryPackageTypeRejectsUnknown(t *testing.T) {
	validate := packageTypeValidateFunc(t, registry.ResourceRegistry())

	_, errs := validate("NOT_A_PACKAGE_TYPE", "package_type")
	if len(errs) == 0 {
		t.Error("expected an unknown package_type to be rejected")
	}
}

// Tests creating a virtual registry for each package type added in AH-5667,
// including an import round trip.
func TestAccResourceVirtualRegistryNewPackageTypes(t *testing.T) {
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	resourceName := "harness_platform_har_registry.test"

	for _, packageType := range newPackageTypes {
		packageType := packageType
		t.Run(string(packageType), func(t *testing.T) {
			id := fmt.Sprintf("tfauto_virt_%s_%s", randAlphanumeric(4), randAlphanumeric(5))

			// resource.Test (not UnitTest) so this is skipped unless TF_ACC is set -
			// it creates real registries.
			resource.Test(t, resource.TestCase{
				PreCheck:          func() { acctest.TestAccPreCheck(t) },
				ProviderFactories: acctest.ProviderFactories,
				CheckDestroy:      testAccRegistryCheckDestroy("harness_platform_har_registry"),
				Steps: []resource.TestStep{
					{
						PreConfig: func() {
							acctest.TestAccConfigureProvider()
							_, _ = acctest.TestAccGetHarClientWithContext()
						},
						Config: testAccResourceVirtualRegistryWithPackageType(id, accountId, (string)(packageType)),
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttr(resourceName, "identifier", id),
							resource.TestCheckResourceAttr(resourceName, "package_type", (string)(packageType)),
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
		})
	}
}

// Generates Terraform config for a virtual registry of the given package type
func testAccResourceVirtualRegistryWithPackageType(id string, accId string, packageType string) string {
	return fmt.Sprintf(`
 resource "harness_platform_har_registry" "test" {
   identifier   = "%[1]s"
   space_ref    = "%[2]s"
   package_type = "%[3]s"

   config {
    type = "VIRTUAL"
   }
   parent_ref = "%[2]s"
 }
`, id, accId, packageType)
}
