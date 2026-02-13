package provider_registry_test

import (
	"archive/zip"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/require"
)

func TestAccResourceInfraProviderVersionFile(t *testing.T) {
	providerType := fmt.Sprintf("test-provider-%s", utils.RandStringBytes(5))
	version := "1.0.0"
	keyId := fmt.Sprintf("test-key-%s", utils.RandStringBytes(8))
	filename := "terraform-provider-test_1.0.0_linux_amd64.zip"
	resourceName := "harness_platform_infra_provider_version_file.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccInfraProviderVersionFileDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceInfraProviderVersionFile(providerType, version, keyId, filename),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "version", version),
					resource.TestCheckResourceAttr(resourceName, "filename", filename),
					resource.TestCheckResourceAttrSet(resourceName, "file_path"),
					testAccInfraProviderVersionFileCreation(t, resourceName, filename),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdFunc:       providerVersionFileImportStateIdFunc(resourceName),
				ImportStateVerifyIgnore: []string{"file_path"},
			},
		},
	})
}

func TestAccResourceInfraProviderVersionFile_DerivedFilename(t *testing.T) {
	providerType := fmt.Sprintf("test-provider-%s", utils.RandStringBytes(5))
	version := "1.0.0"
	keyId := fmt.Sprintf("test-key-%s", utils.RandStringBytes(8))
	filename := "terraform-provider-test_1.0.0_darwin_amd64.zip"
	resourceName := "harness_platform_infra_provider_version_file.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccInfraProviderVersionFileDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceInfraProviderVersionFileDerivedFilename(providerType, version, keyId, filename),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "version", version),
					resource.TestCheckResourceAttr(resourceName, "filename", filename),
					resource.TestCheckResourceAttrSet(resourceName, "file_path"),
					testAccInfraProviderVersionFileCreation(t, resourceName, filename),
				),
			},
		},
	})
}

func providerVersionFileImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		primary := s.RootModule().Resources[resourceName].Primary
		providerId := primary.Attributes["provider_id"]
		version := primary.Attributes["version"]
		filename := primary.Attributes["filename"]
		return fmt.Sprintf("%s/%s/%s", providerId, version, filename), nil
	}
}

func testAccInfraProviderVersionFileCreation(t *testing.T, resourceName string, filename string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No provider version file ID is set")
		}

		providerId := rs.Primary.Attributes["provider_id"]
		if providerId == "" {
			return fmt.Errorf("No provider_id is set")
		}

		version := rs.Primary.Attributes["version"]
		if version == "" {
			return fmt.Errorf("No version is set")
		}

		session := acctest.TestAccProvider.Meta().(*internal.Session)
		c, ctx := session.GetPlatformClientWithContext(context.Background())
		resp, httpRes, err := c.ProviderRegistryApi.ProviderRegistryGetProviderVersion(
			ctx,
			providerId,
			version,
			session.AccountId,
		)

		if err != nil {
			return fmt.Errorf("Failed to get provider version: %v, status: %v", err, httpRes)
		}

		require.NotNil(t, resp)
		return nil
	}
}

func testAccInfraProviderVersionFileDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return nil
		}

		if rs.Primary.ID == "" {
			return nil
		}

		providerId := rs.Primary.Attributes["provider_id"]
		version := rs.Primary.Attributes["version"]

		session := acctest.TestAccProvider.Meta().(*internal.Session)
		c, ctx := session.GetPlatformClientWithContext(context.Background())
		_, httpRes, err := c.ProviderRegistryApi.ProviderRegistryGetProviderVersion(
			ctx,
			providerId,
			version,
			session.AccountId,
		)

		if err == nil {
			return fmt.Errorf("Provider version still exists")
		}

		if httpRes != nil && httpRes.StatusCode == 404 {
			return nil
		}

		return nil
	}
}

func testAccResourceInfraProviderVersionFile(providerType string, version string, keyId string, filename string) string {
	filePath := createTestFile(filename)
	return fmt.Sprintf(`
		resource "harness_platform_infra_provider_signing_key" "test" {
			key_id      = "%[3]s"
			key_name    = "Test Key"
			ascii_armor = <<-EOT
-----BEGIN PGP PUBLIC KEY BLOCK-----

mQINBGDUMYYBEADExampleDummyKeyDataForTestingPurposesOnly123456789ABCDEF
GHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/Example
DummyKeyDataForTestingPurposesOnly123456789ABCDEFGHIJKLMNOPQRSTUVWXYZa
bcdefghijklmnopqrstuvwxyz0123456789+/ExampleDummyKeyDataForTesting
PurposesOnly123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrs
tuvwxyz0123456789+/ExampleDummyKeyDataForTestingPurposesOnly1234567
89ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789
+/ExampleDummyKeyDataForTestingPurposesOnly123456789ABCDEFGHIJKLMNOP
QRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/ExampleDummyKeyDa
taForTestingPurposesOnly123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefg
hijklmnopqrstuvwxyz0123456789+/ExampleDummyKeyDataForTestingPurpos
esOnly123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwx
yz0123456789+/ExampleDummyKeyDataForTestingPurposesOnly123456789ABCD
EFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/Examp
leARAQABtCVUZXN0IFVzZXIgPHRlc3RAZXhhbXBsZS5jb20+iQJOBBMBCAA4Fh
kEExampleDummyKeyDataForTestingPurposesOnly123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz==
=ABCD
-----END PGP PUBLIC KEY BLOCK-----
EOT
			user        = "test@example.com"
		}

		resource "harness_platform_infra_provider" "test" {
			type        = "%[1]s"
			description = "Test provider for version file testing"
			
			lifecycle {
				ignore_changes = [versions]
			}
		}

		resource "harness_platform_infra_provider_version" "test" {
			provider_id = harness_platform_infra_provider.test.id
			version     = "%[2]s"
			gpg_key_id  = harness_platform_infra_provider_signing_key.test.id
			protocols   = ["5.0"]
			
			depends_on = [harness_platform_infra_provider_signing_key.test]
		}

		resource "harness_platform_infra_provider_version_file" "test" {
			provider_id = harness_platform_infra_provider.test.id
			version     = "%[2]s"
			filename    = "%[4]s"
			file_path   = "%[5]s"
			
			depends_on = [harness_platform_infra_provider_version.test]
		}
	`, providerType, version, keyId, filename, filePath)
}

func testAccResourceInfraProviderVersionFileDerivedFilename(providerType string, version string, keyId string, filename string) string {
	filePath := createTestFile(filename)
	return fmt.Sprintf(`
		resource "harness_platform_infra_provider_signing_key" "test" {
			key_id      = "%[3]s"
			key_name    = "Test Key"
			ascii_armor = <<-EOT
-----BEGIN PGP PUBLIC KEY BLOCK-----

mQINBGDUMYYBEADExampleDummyKeyDataForTestingPurposesOnly123456789ABCDEF
GHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/Example
DummyKeyDataForTestingPurposesOnly123456789ABCDEFGHIJKLMNOPQRSTUVWXYZa
bcdefghijklmnopqrstuvwxyz0123456789+/ExampleDummyKeyDataForTesting
PurposesOnly123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrs
tuvwxyz0123456789+/ExampleDummyKeyDataForTestingPurposesOnly1234567
89ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789
+/ExampleDummyKeyDataForTestingPurposesOnly123456789ABCDEFGHIJKLMNOP
QRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/ExampleDummyKeyDa
taForTestingPurposesOnly123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefg
hijklmnopqrstuvwxyz0123456789+/ExampleDummyKeyDataForTestingPurpos
esOnly123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwx
yz0123456789+/ExampleDummyKeyDataForTestingPurposesOnly123456789ABCD
EFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/Examp
leARAQABtCVUZXN0IFVzZXIgPHRlc3RAZXhhbXBsZS5jb20+iQJOBBMBCAA4Fh
kEExampleDummyKeyDataForTestingPurposesOnly123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz==
=ABCD
-----END PGP PUBLIC KEY BLOCK-----
EOT
			user        = "test@example.com"
		}

		resource "harness_platform_infra_provider" "test" {
			type        = "%[1]s"
			description = "Test provider for version file testing"
			
			lifecycle {
				ignore_changes = [versions]
			}
		}

		resource "harness_platform_infra_provider_version" "test" {
			provider_id = harness_platform_infra_provider.test.id
			version     = "%[2]s"
			gpg_key_id  = harness_platform_infra_provider_signing_key.test.id
			protocols   = ["5.0"]
			
			depends_on = [harness_platform_infra_provider_signing_key.test]
		}

		resource "harness_platform_infra_provider_version_file" "test" {
			provider_id = harness_platform_infra_provider.test.id
			version     = "%[2]s"
			filename    = "%[4]s"
			file_path   = "%[5]s"
			
			depends_on = [harness_platform_infra_provider_version.test]
		}
	`, providerType, version, keyId, filename, filePath)
}

func createTestFile(filename string) string {
	tmpDir := os.TempDir()
	tmpFile := filepath.Join(tmpDir, filename)

	zipFile, err := os.Create(tmpFile)
	if err != nil {
		panic(fmt.Sprintf("Failed to create test file: %v", err))
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	writer, err := zipWriter.Create("terraform-provider-test")
	if err != nil {
		panic(fmt.Sprintf("Failed to create zip entry: %v", err))
	}

	_, err = writer.Write([]byte("test provider binary content"))
	if err != nil {
		panic(fmt.Sprintf("Failed to write zip content: %v", err))
	}

	return tmpFile
}
