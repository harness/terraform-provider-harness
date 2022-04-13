package service_test

import (
	"fmt"

	"github.com/harness/harness-go-sdk/harness/cd/cac"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func serviceImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		primary := s.RootModule().Resources[resourceName].Primary
		id := primary.ID
		app_id := primary.Attributes["app_id"]

		return fmt.Sprintf("%s/%s", app_id, id), nil
	}
}

func testAccGetService(resourceName string, state *terraform.State) (*cac.Service, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c := acctest.TestAccGetApiClientFromProvider()
	svcId := r.Primary.ID
	appId := r.Primary.Attributes["app_id"]

	return c.CDClient.ConfigAsCodeClient.GetServiceById(appId, svcId)
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
