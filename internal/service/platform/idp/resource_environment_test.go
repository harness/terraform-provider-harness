package idp_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/idp"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const environmentKind = "environment"

type environmentInfo struct {
	Identifier       string
	Scope            string
	BlueprintID      string
	BlueprintVersion string
	OrgID            optional.String
	ProjectID        optional.String
}

func TestAccResourceEnvironment(t *testing.T) {
	t.Skip("Skipping test as it takes 2+ minutes to run and requires elaborate Harness setup")

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	resourceName := "harness_platform_idp_environment.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccEnvironmentDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceEnvironment(id, "inactive"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "target_state", "inactive"),
				),
			},
			{
				Config: testAccResourceEnvironment(id, "running"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "target_state", "running"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateIdFunc: testAccEnvironmentImportStateIdFunc(resourceName),
			},
		},
	})
}

func testAccEnvironmentDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		entity, _ := testAccGetEnvironment(resourceName, state)
		if entity != nil {
			return fmt.Errorf("Found Environment: %s", entity.Identifier)
		}

		return nil
	}
}

func testAccGetEnvironment(resourceName string, state *terraform.State) (*idp.EntityResponse, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetIDPClientWithContext()
	id := r.Primary.ID

	info := getEnvironmentInfo(r.Primary)

	resp, _, err := c.EntitiesApi.GetEntity(ctx, info.Scope, environmentKind, id, &idp.EntitiesApiGetEntityOpts{
		HarnessAccount:    optional.NewString(c.AccountId),
		OrgIdentifier:     info.OrgID,
		ProjectIdentifier: info.ProjectID,
	})

	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func testAccResourceEnvironment(id string, targetState string) string {
	str := fmt.Sprintf(`
	resource "harness_platform_idp_environment" "test" {
	    identifier = "%[1]s"
		org_id = "default"
		project_id = "ssem"
		name = "%[1]s"
		owner = "user:account/admin@harness.io"
		blueprint_identifier = "noop"
		blueprint_version = "v1.0.0"
		target_state = "%[2]s"
		overrides = <<-EOT
        config: {}
        entities: {}
        EOT
	}
	`, id, targetState)

	return str
}

func getEnvironmentInfo(d *terraform.InstanceState) environmentInfo {
	identifier := d.Attributes["identifier"]
	orgID := d.Attributes["org_id"]
	projectID := d.Attributes["project_id"]
	blueprintId := d.Attributes["blueprint_identifier"]
	blueprintVersion := d.Attributes["blueprint_version"]

	scope := fmt.Sprintf("account.%s.%s", orgID, projectID)

	return environmentInfo{
		Identifier:       identifier,
		BlueprintID:      blueprintId,
		BlueprintVersion: blueprintVersion,
		Scope:            scope,
		OrgID:            optional.NewString(orgID),
		ProjectID:        optional.NewString(projectID),
	}
}

func testAccEnvironmentImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("not found: %s", resourceName)
		}

		info := getEnvironmentInfo(rs.Primary)
		scope, _ := strings.CutPrefix(info.Scope, "account.")

		return fmt.Sprintf("%s/%s", scope, info.Identifier), nil
	}
}
