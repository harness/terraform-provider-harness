package provider_registry_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/require"
)

const testGPGPublicKey = `-----BEGIN PGP PUBLIC KEY BLOCK-----

mQENBGQxKkwBCAC7VDvqVZDmVvnqLvDVqFxJvGCDhvCqJLKKvHLLVQzqwXXqzMPL
test-key-content-here
-----END PGP PUBLIC KEY BLOCK-----`

func TestAccResourceInfraProviderSigningKey(t *testing.T) {
	keyId := fmt.Sprintf("test-key-%s", utils.RandStringBytes(8))
	keyName := fmt.Sprintf("Test Key %s", utils.RandStringBytes(4))
	updatedKeyName := fmt.Sprintf("%s Updated", keyName)
	user := "test-user@example.com"
	resourceName := "harness_platform_infra_provider_signing_key.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccInfraProviderSigningKeyDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceInfraProviderSigningKey(keyId, keyName, user),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "key_id", keyId),
					resource.TestCheckResourceAttr(resourceName, "key_name", keyName),
					resource.TestCheckResourceAttr(resourceName, "user", user),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
					testAccInfraProviderSigningKeyCreation(t, resourceName, keyId),
				),
			},
			{
				Config: testAccResourceInfraProviderSigningKey(keyId, updatedKeyName, user),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "key_id", keyId),
					resource.TestCheckResourceAttr(resourceName, "key_name", updatedKeyName),
					resource.TestCheckResourceAttr(resourceName, "user", user),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"ascii_armor"},
			},
		},
	})
}

func testAccInfraProviderSigningKeyCreation(t *testing.T, resourceName string, keyId string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No signing key ID is set")
		}

		session := acctest.TestAccProvider.Meta().(*internal.Session)
		c, ctx := session.GetPlatformClientWithContext(context.Background())
		resp, httpRes, err := c.ProviderRegistryApi.ProviderRegistryListSigningKeys(
			ctx,
			session.AccountId,
			nil,
		)

		if err != nil {
			return fmt.Errorf("Failed to list signing keys: %v, status: %v", err, httpRes)
		}

		keyFound := false
		for _, key := range resp {
			if key.Id == rs.Primary.ID {
				keyFound = true
				require.Equal(t, keyId, key.KeyId)
				break
			}
		}

		require.True(t, keyFound, "Signing key not found")
		return nil
	}
}

func testAccInfraProviderSigningKeyDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return nil
		}

		if rs.Primary.ID == "" {
			return nil
		}

		session := acctest.TestAccProvider.Meta().(*internal.Session)
		c, ctx := session.GetPlatformClientWithContext(context.Background())
		resp, httpRes, err := c.ProviderRegistryApi.ProviderRegistryListSigningKeys(
			ctx,
			session.AccountId,
			nil,
		)

		if err != nil {
			return fmt.Errorf("Error listing signing keys: %v, status: %v", err, httpRes)
		}

		for _, key := range resp {
			if key.Id == rs.Primary.ID {
				return fmt.Errorf("Signing key still exists")
			}
		}

		return nil
	}
}

func testAccResourceInfraProviderSigningKey(keyId string, keyName string, user string) string {
	return fmt.Sprintf(`
		resource "harness_platform_infra_provider_signing_key" "test" {
			key_id      = "%s"
			key_name    = "%s"
			ascii_armor = <<-EOT
%s
EOT
			user        = "%s"
		}
	`, keyId, keyName, testGPGPublicKey, user)
}
