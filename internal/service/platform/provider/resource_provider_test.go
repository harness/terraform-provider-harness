package provider_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestResourceProvider(t *testing.T) {

	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_provider.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testProviderDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testResourceProvider(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			{
				Config: testResourceProvider(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.AccountLevelResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func testGetProvider(resourceName string, state *terraform.State) (*nextgen.Provider, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	id := r.Primary.ID

	resp, _, err := c.ProviderApi.GetProvider(ctx, id, c.AccountId)
	if err != nil {
		return nil, err
	}

	if resp.Data == nil {
		return nil, nil
	}

	return resp.Data, nil
}

func testProviderDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		provider, _ := testGetProvider(resourceName, state)
		if provider != nil {
			return fmt.Errorf("found provider: %s", provider.Identifier)
		}

		return nil
	}
}

func testResourceProvider(id string, name string) string {
	return fmt.Sprintf(`
    resource "harness_platform_provider" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			spec {
				type = "BITBUCKET_SERVER"
				domain              = "https://example.com"
				secret_manager_ref  = "secret-ref"
				delegate_selectors  = ["delegate-1", "delegate-2"]
				client_id           = "client-id"
				client_secret_ref   = "client-secret-ref"
			}
		}
`, id, name)
}
