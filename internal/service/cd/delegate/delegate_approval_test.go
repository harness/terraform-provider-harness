package delegate_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/harness/harness-go-sdk/harness/cd/graphql"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccApproveDelegate(t *testing.T) {
	t.Skip("Skipping until we figure out how to get the tests passing properly in CI")

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(7))
		resourceName = "harness_delegate_approval.test"
	)

	defer deleteDelegate(t, name)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			createDelegateContainer(t, name)
		},
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testaccDelegateApproval(name, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "status", graphql.DelegateStatusTypes.Enabled.String()),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					primary := s.RootModule().Resources[resourceName].Primary
					return primary.Attributes["delegate_id"], nil
				},
			},
			{
				Config:      testaccDelegateApproval(name, false),
				ExpectError: regexp.MustCompile(`.*has already been changed.*`),
			},
		},
	})
}

func TestAccApproveDelegate_DelegateNotFound(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(7))
		resourceName = "harness_delegate_approval.test"
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testaccDelegateApproval(name, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "status", graphql.DelegateStatusTypes.Enabled.String()),
				),
				ExpectError: regexp.MustCompile(`.*no delegate found.*`),
			},
		},
	})
}

func testaccDelegateApproval(name string, approve bool) string {
	return fmt.Sprintf(`
		data "harness_delegate" "test" {
			name = "%s"
		}

		resource "harness_delegate_approval" "test" {
			delegate_id = data.harness_delegate.test.id
			approve = %[2]t
		}
	`, name, approve)
}
