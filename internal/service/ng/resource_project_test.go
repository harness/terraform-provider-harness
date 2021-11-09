package ng_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/antihax/optional"
	"github.com/harness-io/harness-go-sdk/harness/nextgen"
	"github.com/harness-io/harness-go-sdk/harness/utils"
	"github.com/harness-io/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceProject(t *testing.T) {

	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_project.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccProjectDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceProject(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			{
				Config: testAccResourceProject(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					primary := s.RootModule().Resources[resourceName].Primary
					id := primary.ID
					org_id := primary.Attributes["org_id"]
					return fmt.Sprintf("%s/%s", org_id, id), nil
				},
			},
		},
	})
}

func testAccGetProject(resourceName string, state *terraform.State) (*nextgen.Project, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c := acctest.TestAccGetApiClientFromProvider()
	id := r.Primary.ID
	orgId := r.Primary.Attributes["org_id"]

	resp, _, err := c.NGClient.ProjectApi.GetProject(context.Background(), id, c.AccountId, &nextgen.ProjectApiGetProjectOpts{OrgIdentifier: optional.NewString(orgId)})
	if err != nil {
		return nil, err
	}

	return resp.Data.Project, nil
}

func testAccProjectDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		project, _ := testAccGetProject(resourceName, state)
		if project != nil {
			return fmt.Errorf("Found project: %s", project.Identifier)
		}

		return nil
	}
}

func testAccResourceProject(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_project" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = "test"
			color = "#0063F7"
		}
`, id, name)
}
