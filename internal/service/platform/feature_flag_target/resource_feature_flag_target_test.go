package feature_flag_target_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceFeatureFlagTarget(t *testing.T) {

	name := t.Name()
	targetName := name
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	resourceName := "harness_platform_environment.test"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceFeatureFlagTargetDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceFeatureFlagTarget(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
					resource.TestCheckResourceAttr(resourceName, "name", targetName),
				),
			},
			{
				Config: testAccResourceFeatureFlagTarget(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"yaml"},
				ImportStateIdFunc:       acctest.ProjectResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func testAccResourceFeatureFlagTarget(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
		}

		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_organization.test.id
			color = "#0063F7"
		}
	
		resource "harness_platform_environment" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
			tags = ["foo:bar", "baz"]
			type = "PreProduction"
			yaml = <<-EOT
			   environment:
         name: %[2]s
         identifier: %[1]s
         orgIdentifier: ${harness_platform_project.test.org_id}
         projectIdentifier: ${harness_platform_project.test.id}
         type: PreProduction
         tags:
           foo: bar
           baz: ""
         variables:
           - name: envVar1
             type: String
             value: v1
             description: ""
           - name: envVar2
             type: String
             value: v2
             description: ""
         overrides:
           manifests:
             - manifest:
                 identifier: manifestEnv
                 type: Values
                 spec:
                   store:
                     type: Git
                     spec:
                       connectorRef: <+input>
                       gitFetchType: Branch
                       paths:
                         - file1
                       repoName: <+input>
                       branch: master
           configFiles:
             - configFile:
                 identifier: configFileEnv
                 spec:
                   store:
                     type: Harness
                     spec:
                       files:
                         - account:/Add-ons/svcOverrideTest
                       secretFiles: []
      EOT
  	}

		resource "harness_platform_feature_flag_target" "target" {
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
			environment = harness_platform_environment.test.id
			account_id = harness_platform_project.test.id
		
			identifier  = "%[1]s"
			name        = "%[2]s"
		
			attributes = {
				foo : "bar"
			}
		}
`, id, name)
}

func testAccResourceFeatureFlagTargetDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		env, _ := testAccGetPlatformFeatureFlagTarget(resourceName, state)
		if env != nil {
			return fmt.Errorf("Feature Flag Target not found: %s", env.Identifier)
		}

		return nil
	}
}

func testAccGetPlatformFeatureFlagTarget(resourceName string, state *terraform.State) (*nextgen.Target, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	id := r.Primary.ID
	environment := r.Primary.Attributes["environment"]
	orgId := r.Primary.Attributes["org_id"]
	projId := r.Primary.Attributes["project_id"]

	target, resp, err := c.TargetsApi.GetTarget((ctx), id, c.AccountId, orgId, projId, environment)

	if err != nil {
		return nil, err
	}

	if resp == nil {
		return nil, nil
	}

	return &target, nil
}
