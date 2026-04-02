package governance_rule_set_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceRuleSet(t *testing.T) {
	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_governance_rule_set.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceRuleSet(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "cloud_provider", "AWS"),
					resource.TestCheckResourceAttr(resourceName, "description", "Dummy"),
					resource.TestCheckResourceAttr(resourceName, "rule_ids.#", "1"),
				),
			},
		},
	})
}

func testAccDataSourceRuleSet(name string) string {
	return fmt.Sprintf(`
	resource "harness_governance_rule" "rule" {
		name           = "%[1]s_rule"
		cloud_provider = "AWS"
		description    = "Dummy"
		rules_yaml     = "policies:\n  - name: aws-list-ec2\n    resource: aws.ec2"
	}

	resource "harness_governance_rule_set" "test" {
		name           = "%[1]s"
		cloud_provider = "AWS"
		description    = "Dummy"
		rule_ids       = [harness_governance_rule.rule.rule_id]
	}

	data "harness_governance_rule_set" "test" {
		rule_set_id = harness_governance_rule_set.test.rule_set_id
	}
	`, name)
}
