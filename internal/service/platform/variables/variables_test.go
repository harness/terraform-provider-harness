package variables_test

import (
	"fmt"
	"testing"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/require"
)

func TestAccResourceVariables(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)

	variableValue := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	updatedValue := variableValue + "updated"

	resourceName := "harness_platform_variables.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccVariablesDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceVariables(id, name, variableValue),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "spec.0.fixed_value", variableValue),
				),
			},
			{
				Config: testAccResourceVariables(id, updatedName, updatedValue),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "spec.0.fixed_value", updatedValue),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.ProjectResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func TestAccResourceVariablesOrgLevel(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)

	variableValue := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	updatedValue := variableValue + "updated"

	resourceName := "harness_platform_variables.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccVariablesOrgLevelDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceVariablesOrgLevel(id, name, variableValue),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "spec.0.fixed_value", variableValue),
				),
			},
			{
				Config: testAccResourceVariablesOrgLevel(id, updatedName, updatedValue),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "spec.0.fixed_value", updatedValue),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.OrgResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func TestAccResourceVariables_DeleteUnderlyingResource(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	resourceName := "harness_platform_variables.test"
	variableValue := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceVariables(id, name, variableValue),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "spec.0.fixed_value", variableValue),
				),
			},
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					c, ctx := acctest.TestAccGetPlatformClientWithContext()
					_, _, err := c.VariablesApi.DeleteVariable(ctx, c.AccountId, id, &nextgen.VariablesApiDeleteVariableOpts{
						OrgIdentifier:     optional.NewString(id),
						ProjectIdentifier: optional.NewString(id),
					})
					require.NoError(t, err)
				},
				Config:             testAccResourceVariables(id, name, variableValue),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccGetResourceVariables(resourceName string, state *terraform.State) (*nextgen.VariableDto, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	id := r.Primary.ID

	resp, _, err := c.VariablesApi.GetVariable(ctx, id, c.AccountId, &nextgen.VariablesApiGetVariableOpts{
		OrgIdentifier:     buildField(r, "org_id"),
		ProjectIdentifier: buildField(r, "project_id"),
	})

	if err != nil {
		return nil, err
	}

	if resp.Data == nil {
		return nil, nil
	}

	return resp.Data.Variable, nil
}

func testAccVariablesOrgLevelDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		variable, _ := testAccGetResourceVariables(resourceName, state)
		if variable != nil {
			return fmt.Errorf("Found variable: %s", variable.Identifier)
		}

		return nil
	}
}

func testAccVariablesDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		variable, _ := testAccGetResourceVariables(resourceName, state)
		if variable != nil {
			return fmt.Errorf("Found variable: %s", variable.Identifier)
		}

		return nil
	}
}

func buildField(r *terraform.ResourceState, field string) optional.String {
	if attr, ok := r.Primary.Attributes[field]; ok {
		return optional.NewString(attr)
	}
	return optional.EmptyString()
}

func testAccResourceVariables(id string, name string, variableValue string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
		}

		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_organization.test.id
			color = "#472848"
		}

		resource "harness_platform_variables" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
			type = "String"
			spec {
				value_type = "FIXED"
				fixed_value = "%[3]s"
			}
		}
`, id, name, variableValue)
}

func testAccResourceVariablesOrgLevel(id string, name string, variableValue string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
		}

		resource "harness_platform_variables" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_organization.test.id
			type = "String"
			spec {
				value_type = "FIXED"
				fixed_value = "%[3]s"
			}
		}
`, id, name, variableValue)
}
