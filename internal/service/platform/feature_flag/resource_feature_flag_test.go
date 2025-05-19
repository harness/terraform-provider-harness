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
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	flagResourceName := "harness_platform_feature_flag.test"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceFeatureFlagDestroy(flagResourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceFeatureFlag(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(flagResourceName, "identifier", id),
					resource.TestCheckResourceAttr(flagResourceName, "org_id", id),
					resource.TestCheckResourceAttr(flagResourceName, "project_id", id),
					resource.TestCheckResourceAttr(flagResourceName, "name", name),
				),
			},
			{
				Config: testAccResourceFeatureFlag(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(flagResourceName, "identifier", id),
					resource.TestCheckResourceAttr(flagResourceName, "org_id", id),
					resource.TestCheckResourceAttr(flagResourceName, "project_id", id),
					resource.TestCheckResourceAttr(flagResourceName, "name", name),
				),
			},
			{
				ResourceName:            flagResourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"yaml"},
				ImportStateIdFunc:       acctest.ProjectResourceImportStateIdFunc(flagResourceName),
			},
		},
	})
}

func testAccResourceFeatureFlag(id string, name string) string {
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
			type = "PreProduction"
  	}

		resource "harness_platform_feature_flag_target" "target1" {
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
			environment = harness_platform_environment.test.id
			account_id = harness_platform_project.test.id
		
			identifier  = "target1"
			name        = "target1"
		
			attributes = {
				foo : "bar"
			}
		}

		resource "harness_platform_feature_flag_target_group" "targetgroup1" {
			identifier = "targetgroup1"
			name = "targetgroup1"
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
			environment = harness_platform_environment.test.id
			account_id = harness_platform_project.test.id
			included = []
			excluded = []
			rule {
				attribute = "identifier"
				op        = "equal"
				values    = [harness_platform_feature_flag_target.target1.id]
			}
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
			  value       = "true"
			}

			tags {
				identifier = "bar"
			}
		}
`, id, name)
}

func testAccResourceFeatureFlagDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		env, _ := testAccGetPlatformFeatureFlag(resourceName, state)
		if env != nil {
			return fmt.Errorf("Feature Flag not found: %s", env.Identifier)
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
