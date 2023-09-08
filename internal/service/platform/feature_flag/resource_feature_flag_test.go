package feature_flag_test

import (
	"fmt"
	"testing"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceFeatureFlag(t *testing.T) {

	name := t.Name()
	targetName := name
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	resourceName := "harness_platform_feature_flag.test"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceFeatureFlagDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceFeatureFlag(id, name, targetName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
					resource.TestCheckResourceAttr(resourceName, "name", targetName),
				),
			},
			{
				Config: testAccResourceFeatureFlag(id, name, name),
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

func testAccResourceFeatureFlag(id string, name string, updatedName string) string {
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

		resource "harness_platform_feature_flag" "test" {
			identifier = "%[1]s"
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
			name = "%[2]s"
			kind       = "boolean"
			permanent  = false
		  
			default_on_variation  = "Enabled"
			default_off_variation = "Disabled"
		  
			variation {
			  identifier  = "Enabled"
			  name        = "Enabled"
			  description = "The feature is enabled"
			  value       = "true"
			}
		  
			variation {
			  identifier  = "Disabled"
			  name        = "Disabled"
			  description = "The feature is disabled"
			  value       = "false"
			}
		}
`, id, name, updatedName)
}

func testAccResourceFeatureFlagDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		env, _ := testAccGetPlatformFeatureFlag(resourceName, state)
		if env != nil {
			return fmt.Errorf("Feature Flag Target not found: %s", env.Identifier)
		}

		return nil
	}
}

func testAccGetPlatformFeatureFlag(resourceName string, state *terraform.State) (*nextgen.Feature, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	id := r.Primary.ID
	orgId := r.Primary.Attributes["org_id"]
	projId := r.Primary.Attributes["project_id"]
	readOpts := &nextgen.FeatureFlagsApiGetFeatureFlagOpts{
		EnvironmentIdentifier: optional.EmptyString(),
	}

	featureFlag, resp, err := c.FeatureFlagsApi.GetFeatureFlag(ctx, id, c.AccountId, orgId, projId, readOpts)

	if err != nil {
		return nil, err
	}

	if resp == nil {
		return nil, nil
	}

	return &featureFlag, nil
}
