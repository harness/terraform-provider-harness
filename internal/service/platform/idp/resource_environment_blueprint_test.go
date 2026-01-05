package idp_test

import (
	"encoding/json"
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

func TestAccResourceEnvironmentBlueprint(t *testing.T) {
	description := t.Name()
	id := fmt.Sprintf("%s_%s", description, utils.RandStringBytes(5))
	updatedDescription := fmt.Sprintf("%s_updated", description)
	resourceName := "harness_platform_idp_environment_blueprint.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccEnvironmentBlueprintDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceEnvironmentBlueprint(id, description),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "description", description),
				),
			},
			{
				Config: testAccResourceEnvironmentBlueprint(id, updatedDescription),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "description", updatedDescription),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateIdFunc: testAccEnvironmentBlueprintImportStateIdFunc(resourceName),
			},
		},
	})
}

func testAccGetEnvironmentBlueprint(resourceName string, state *terraform.State) (*idp.EntityVersionResponse, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetIDPClientWithContext()

	id := r.Primary.ID
	version := r.Primary.Attributes["version"]

	resp, _, err := c.EntitiesApi.GetEntityVersion(ctx, "account", "environmentBlueprint", id, version, &idp.EntitiesApiGetEntityVersionOpts{
		HarnessAccount: optional.NewString(c.AccountId),
	})
	if err != nil {
		return nil, err
	}

	return &resp, nil

}

func testAccEnvironmentBlueprintDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		catalogEntity, err := testAccGetEnvironmentBlueprint(resourceName, state)

		if err != nil {
			body, _ := toIDPErrBody(err)

			fmt.Printf("%+v\n", body)

			if strings.Contains(body.Message, "not found") {
				return nil
			}

			return err
		}
		if catalogEntity != nil {
			return fmt.Errorf("Found catalog entity version: %s (version: %s)", catalogEntity.Identifier, catalogEntity.Version)
		}
		return nil
	}
}

func testAccResourceEnvironmentBlueprint(id string, description string) string {
	str := fmt.Sprintf(`
	    resource "harness_platform_idp_environment_blueprint" "test" {
		identifier = "%[1]s"
		version = "1"
		stable = true
		deprecated = false
		description = "%[2]s"
		yaml = <<-EOT
		apiVersion: harness.io/v1
		kind: EnvironmentBlueprint
		type: long-lived
		identifier: %[1]s
		name: %[1]s
		owner: group:account/_account_all_users
		metadata:
		  description: %[2]s
		spec:
		  entities:
		  - identifier: git
		    backend:
		      type: HarnessCD
		      steps:
		        apply:
		          pipeline: gittest
		          branch: main
		        destroy:
		          pipeline: gittest
		          branch: not-main
		  ownedBy:
		  - group:account/_account_all_users
		EOT
	}
	`, id, description)
	return str
}

func testAccEnvironmentBlueprintImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("not found: %s", resourceName)
		}

		return fmt.Sprintf("%s/%s", rs.Primary.ID, rs.Primary.Attributes["version"]), nil
	}
}

type IDPErrorBody struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

func toIDPErrBody(err error) (*IDPErrorBody, error) {
	var body IDPErrorBody
	apiErr, ok := err.(interface{ Body() []byte })
	if !ok {
		return nil, err
	}

	newErr := json.Unmarshal(apiErr.Body(), &body)
	if newErr != nil {
		return nil, newErr
	}

	return &body, nil
}
