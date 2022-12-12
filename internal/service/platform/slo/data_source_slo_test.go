package slo_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceSlo(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	orgId := "default"
	projectId := "default_project"
	name := id
	resourceName := "data.harness_platform_slo.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSlo(id, name, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "org_id", orgId),
					resource.TestCheckResourceAttr(resourceName, "account_id", accountId),
					resource.TestCheckResourceAttr(resourceName, "project_id", projectId),
				),
			},
		},
	})
}

func testAccDataSourceSlo(id string, name string, accountId string) string {
	return fmt.Sprintf(`
	resource "harness_platform_slo" "test" {
		account_id = "%[3]s"
		org_id     = "default"
		project_id = "default_project"
		identifier = "%[1]s"
		request {
			  name = "%[2]s"
			  description = "description"
			  tags = ["foo:bar", "bar:foo"]
			  user_journey_refs = ["one", "two"]
			  slo_target {
					type = "Rolling"
					slo_target_percentage = 10.0
					spec = jsonencode({
						periodLength = "28d"
					})
			  }
			  type = "Simple"
			  spec = jsonencode({
					monitoredServiceRef = "monitoredServiceRef"
					healthSourceRef = "healthSourceRef"
					serviceLevelIndicatorType = "serviceLevelIndicatorType"
			  })
			  notification_rule_refs {
					notification_rule_ref = "notification_rule_ref"
					enabled = true
			  }
		}
	}

	data "harness_platform_slo" "test" {
		account_id = harness_platform_slo.test.account_id
		identifier = harness_platform_slo.test.identifier
		org_id = harness_platform_slo.test.org_id
		project_id = harness_platform_slo.test.project_id
	}
`, id, name, accountId)
}
