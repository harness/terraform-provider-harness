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

func TestAccResourceInfraProvider(t *testing.T) {
	providerType := fmt.Sprintf("test-provider-%s", utils.RandStringBytes(5))
	resourceName := "harness_platform_infra_provider.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccInfraProviderDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceInfraProvider(providerType, "Test provider description"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "type", providerType),
					resource.TestCheckResourceAttr(resourceName, "description", "Test provider description"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "account"),
					resource.TestCheckResourceAttrSet(resourceName, "created"),
					resource.TestCheckResourceAttrSet(resourceName, "updated"),
					testAccInfraProviderCreation(t, resourceName, providerType),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"versions"},
			},
		},
	})
}

func testAccInfraProviderCreation(t *testing.T, resourceName string, providerType string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No provider ID is set")
		}

		session := acctest.TestAccProvider.Meta().(*internal.Session)
		c, ctx := session.GetPlatformClientWithContext(context.Background())
		resp, httpRes, err := c.ProviderRegistryApi.ProviderRegistryGetProvider(
			ctx,
			rs.Primary.ID,
			session.AccountId,
		)

		if err != nil {
			return fmt.Errorf("Failed to get provider: %v, status: %v", err, httpRes)
		}

		require.Equal(t, providerType, resp.Type_)
		return nil
	}
}

func testAccInfraProviderDestroy(resourceName string) resource.TestCheckFunc {
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
		_, httpRes, err := c.ProviderRegistryApi.ProviderRegistryGetProvider(
			ctx,
			rs.Primary.ID,
			session.AccountId,
		)

		if err == nil {
			return fmt.Errorf("Provider still exists")
		}

		if httpRes != nil && httpRes.StatusCode == 404 {
			return nil
		}

		return fmt.Errorf("Unexpected error checking provider deletion: %v", err)
	}
}

func testAccResourceInfraProvider(providerType string, description string) string {
	return fmt.Sprintf(`
		resource "harness_platform_infra_provider" "test" {
			type        = "%s"
			description = "%s"
			
			lifecycle {
				ignore_changes = [versions]
			}
		}
	`, providerType, description)
}
