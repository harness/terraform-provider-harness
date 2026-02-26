package governance_enforcement_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceRuleEnforcement(t *testing.T) {
	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_governance_rule_enforcement.test"
		awsAccount   = os.Getenv("AWS_ACCOUNT_ID")
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceRuleEnforcement(name, awsAccount),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "cloud_provider", "AWS"),
					resource.TestCheckResourceAttr(resourceName, "execution_schedule", "0 0 * * * *"),
					resource.TestCheckResourceAttr(resourceName, "execution_timezone", "Asia/Calcutta"),
					resource.TestCheckResourceAttr(resourceName, "is_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "target_accounts.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "target_regions.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "Dummy"),
				),
			},
		},
	})
}

func testAccDataSourceRuleEnforcement(name, awsAccount string) string {
	return fmt.Sprintf(`
	resource "harness_governance_rule" "rule" {
		name           = "%[1]s_rule"
		cloud_provider = "AWS"
		description    = "Dummy"
		rules_yaml     = "policies:\n  - name: aws-list-ec2\n    resource: aws.ec2"
	}

	resource "harness_governance_rule_enforcement" "test" {
		name               = "%[1]s"
		cloud_provider     = "AWS"
		rule_ids           = [harness_governance_rule.rule.rule_id]
		rule_set_ids       = []
		execution_schedule = "0 0 * * * *"
		execution_timezone = "Asia/Calcutta"
		is_enabled         = true
		target_accounts    = ["%[2]s"]
		target_regions     = ["us-east-1"]
		is_dry_run         = true
		description        = "Dummy"
	}

	data "harness_governance_rule_enforcement" "test" {
		enforcement_id = harness_governance_rule_enforcement.test.enforcement_id
	}
	`, name, awsAccount)
}
