package slo_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/require"
)

func TestAccResourceSlo(t *testing.T) {

	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	org := "default"
	project := "default_project"
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_slo.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccSloDestroy(resourceName, org, project),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceSlo(id, name, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			{
				Config: testAccResourceSlo(id, updatedName, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccResourceSlo_DeleteUnderlyingResource(t *testing.T) {
	t.Skip()
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	org := "default"
	project := "default_project"
	resourceName := "harness_platform_slo.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccSloDestroy(resourceName, org, project),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceSlo(id, name, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					c, ctx := acctest.TestAccGetPlatformClientWithContext()
					resp, _, err := c.SloApi.DeleteSLODataNg(ctx, id, c.AccountId, org, project)
					require.NoError(t, err)
					require.True(t, resp.Resource)
				},
				Config:             testAccResourceSlo(id, name, accountId),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccGetSlo(resourceName string, org string, project string, state *terraform.State) (*nextgen.ServiceLevelObjectiveV2Dto, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	id := r.Primary.ID

	resp, _, err := c.SloApi.GetServiceLevelObjectiveNg(ctx, id, c.AccountId, org, project)
	if err != nil {
		return nil, err
	}

	return resp.Resource.ServiceLevelObjectiveV2, nil
}

func testAccSloDestroy(resourceName string, org string, project string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		slo, _ := testAccGetSlo(resourceName, org, project, state)
		if slo != nil {
			return fmt.Errorf("Found SLO: %s", slo.Identifier)
		}

		return nil
	}
}

func testAccResourceSlo(id string, name string, accountId string) string {
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
`, id, name, accountId)
}
