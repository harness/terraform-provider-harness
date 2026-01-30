package ip_allowlist_test

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceIPAllowlist(t *testing.T) {
	id := fmt.Sprintf("ip_%s", utils.RandStringBytes(5))
	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_ip_allowlist.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccIPAllowlistDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccIPAllowlist(id, name, "0.0.0.0/0", "API"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "allowlist test"),
					resource.TestCheckResourceAttr(resourceName, "ip_address", "0.0.0.0/0"),
					resource.TestCheckResourceAttr(resourceName, "enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "allowed_source_type.#", "1"),
					resource.TestCheckTypeSetElemAttr(resourceName, "allowed_source_type.*", "API"),
					resource.TestCheckResourceAttrSet(resourceName, "created"),
					resource.TestCheckResourceAttrSet(resourceName, "updated"),
				),
			},
			{
				PreConfig: func() { time.Sleep(60 * time.Second) },
				Config:    testAccIPAllowlist(id, updatedName, "10.0.0.0/8", "UI"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "ip_address", "10.0.0.0/8"),
					resource.TestCheckResourceAttr(resourceName, "allowed_source_type.#", "1"),
					resource.TestCheckTypeSetElemAttr(resourceName, "allowed_source_type.*", "UI"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateIdFunc: acctest.AccountLevelResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func testAccIPAllowlistDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		config, httpResp, err := testAccGetIPAllowlist(resourceName, state)
		if err != nil && httpResp != nil && httpResp.StatusCode == 404 {
			return nil
		}
		if err != nil {
			return err
		}
		if config != nil {
			return fmt.Errorf("found IP allowlist config: %s", config.Identifier)
		}
		return nil
	}
}

func testAccGetIPAllowlist(resourceName string, state *terraform.State) (*nextgen.IpAllowlistConfig, *http.Response, error) {
	rm := state.RootModule()
	r, ok := rm.Resources[resourceName]
	if !ok {
		return nil, nil, nil
	}
	c, ctx := acctest.TestAccGetPlatformClientWithContext()

	resp, httpResp, err := c.IPAllowlistApiService.GetIpAllowlistConfig(ctx, r.Primary.ID, &nextgen.IPAllowlistApiGetIpAllowlistConfigOpts{
		HarnessAccount: optional.NewString(c.AccountId),
	})
	if err != nil {
		return nil, httpResp, err
	}

	return resp.IpAllowlistConfig, httpResp, nil
}

func testAccIPAllowlist(id, name, ipAddress, sourceType string) string {
	return fmt.Sprintf(`
		resource "harness_platform_ip_allowlist" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "allowlist test"
			ip_address = "%[3]s"
			allowed_source_type = ["%[4]s"]
			enabled = false
		}
	`, id, name, ipAddress, sourceType)
}
