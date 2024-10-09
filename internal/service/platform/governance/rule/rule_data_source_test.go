package governance_rule_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceRule(t *testing.T) {
	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_governance_rule.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceRule(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "cloud_provider", "AWS"),
					resource.TestCheckResourceAttr(resourceName, "rules_yaml", "policies:\n  - name: aws-list-ec2\n    resource: aws.ec2"),
					resource.TestCheckResourceAttr(resourceName, "description", "Dummy"),
				),
			},
		},
	})
}

func testAccDataSourceRule(name string) string {
	return fmt.Sprintf(`
	resource "harness_governance_rule" "test" {
		name               = "%[1]s"
		cloud_provider     = "AWS"
		rules_yaml 		   = "policies:\n  - name: aws-list-ec2\n    resource: aws.ec2"
		description        = "Dummy"
	}

	data "harness_governance_rule" "test" {
		rule_id     = harness_governance_rule.test.rule_id
	}
	`, name)
}
