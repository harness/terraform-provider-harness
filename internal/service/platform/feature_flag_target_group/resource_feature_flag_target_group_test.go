package featureflagtargetgroup_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceFeatureFlagTargetGroup(t *testing.T) {

	name := t.Name()
	targetName := name
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	resourceName := "harness_platform_environment.test"
	environment := "qa"
	environmentId := fmt.Sprintf("%s_%s", "env", utils.RandStringBytes(5))

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceFeatureFlagTargetGroupDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceFeatureFlagTarget(id, name, targetName, environmentId, environment),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
					resource.TestCheckResourceAttr(resourceName, "name", targetName),
				),
			},
			{
				Config: testAccResourceFeatureFlagTarget(id, name, name, environmentId, environment),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
					resource.TestCheckResourceAttr(resourceName, "name", targetName),
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

func testAccResourceFeatureFlagTarget(id string, name string, updatedName string, environmentId string, environment string) string {
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

		resource "harness_platform_feature_flag_target" "test" {
			identifier = "%[1]s"
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
			environment = harness_platform_environment.test.id
			account_id = harness_platform_project.test.id
			name = "%[2]s"
			attributes = {}
		}
`, id, name, updatedName, environmentId, environment)
}

func testAccResourceFeatureFlagTargetGroupDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		env, _ := testAccGetPlatformFeatureFlagTargetGroup(resourceName, state)
		if env != nil {
			return fmt.Errorf("Feature Flag Target Group not found: %s", env.Identifier)
		}

		return nil
	}
}

func testAccGetPlatformFeatureFlagTargetGroup(resourceName string, state *terraform.State) (*nextgen.Segment, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	id := r.Primary.ID
	environment := r.Primary.Attributes["environment"]
	orgId := r.Primary.Attributes["org_id"]
	projId := r.Primary.Attributes["project"]

	segment, resp, err := c.TargetGroupsApi.GetSegment((ctx), c.AccountId, orgId, id, projId, environment)

	if err != nil {
		return nil, err
	}

	if resp == nil {
		return nil, nil
	}

	return &segment, nil
}
