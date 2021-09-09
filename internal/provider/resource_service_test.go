package provider

import (
	"fmt"

	"github.com/harness-io/harness-go-sdk/harness/api/cac"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func testAccGetService(resourceName string, state *terraform.State) (*cac.Service, error) {
	r := testAccGetResource(resourceName, state)
	c := testAccGetApiClientFromProvider()
	svcId := r.Primary.ID
	appId := r.Primary.Attributes["app_id"]

	return c.ConfigAsCode().GetServiceById(appId, svcId)
}

func testAccServiceDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		svc, _ := testAccGetService(resourceName, state)
		if svc != nil {
			return fmt.Errorf("Found service: %s", svc.Id)
		}

		return nil
	}
}
