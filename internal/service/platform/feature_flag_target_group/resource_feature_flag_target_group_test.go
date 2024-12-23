package featureflagtargetgroup_test

import (
	"fmt"
	featureflagtargetgroup "github.com/harness/terraform-provider-harness/internal/service/platform/feature_flag_target_group"
	"reflect"
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

		resource "harness_platform_feature_flag_target_group" "test" {
			identifier = "%[1]s"
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
			environment = harness_platform_environment.test.id
			account_id = harness_platform_project.test.id
			name = "%[2]s"
			included = []
			excluded = []
			rule {
				attribute = "identifier"
				op        = "equal"
				values    = [harness_platform_feature_flag_target.target.id]
			}
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
	projId := r.Primary.Attributes["project_id"]

	segment, resp, err := c.TargetGroupsApi.GetSegment((ctx), c.AccountId, orgId, id, projId, environment)

	if err != nil {
		return nil, err
	}

	if resp == nil {
		return nil, nil
	}

	return &segment, nil
}

func TestRuleDiffs(t *testing.T) {
	type args struct {
		first  *[]nextgen.Clause
		second *[]nextgen.Clause
	}
	type expected struct {
		extraRules   []nextgen.Clause
		missingRules []nextgen.Clause
	}
	tests := map[string]struct {
		args     args
		expected expected
	}{
		"no rules have no changes": {
			args: args{},
			expected: expected{
				extraRules:   nil,
				missingRules: nil,
			},
		},
		"identical rules have no changes": {
			args: args{
				first:  &[]nextgen.Clause{{Attribute: "foo", Op: "equal", Values: []string{"bar"}}},
				second: &[]nextgen.Clause{{Attribute: "foo", Op: "equal", Values: []string{"bar"}}},
			},
			expected: expected{
				extraRules:   nil,
				missingRules: nil,
			},
		},
		"order of rules doesn't matter": {
			args: args{
				first:  &[]nextgen.Clause{{Attribute: "foo", Op: "equal", Values: []string{"bar"}}, {Attribute: "foo2", Op: "equal", Values: []string{"bar2"}}},
				second: &[]nextgen.Clause{{Attribute: "foo2", Op: "equal", Values: []string{"bar2"}}, {Attribute: "foo", Op: "equal", Values: []string{"bar"}}},
			},
			expected: expected{
				extraRules:   nil,
				missingRules: nil,
			},
		},
		"extra rule gets returned": {
			args: args{
				first:  &[]nextgen.Clause{{Attribute: "foo", Op: "equal", Values: []string{"bar"}}, {Attribute: "foo2", Op: "equal", Values: []string{"bar2"}}},
				second: &[]nextgen.Clause{{Attribute: "foo", Op: "equal", Values: []string{"bar"}}},
			},
			expected: expected{
				extraRules:   []nextgen.Clause{{Attribute: "foo2", Op: "equal", Values: []string{"bar2"}}},
				missingRules: nil,
			},
		},
		"missing rule gets returned": {
			args: args{
				first:  &[]nextgen.Clause{{Attribute: "foo", Op: "equal", Values: []string{"bar"}}},
				second: &[]nextgen.Clause{{Attribute: "foo", Op: "equal", Values: []string{"bar"}}, {Attribute: "foo2", Op: "equal", Values: []string{"bar2"}}},
			},
			expected: expected{
				extraRules:   nil,
				missingRules: []nextgen.Clause{{Attribute: "foo2", Op: "equal", Values: []string{"bar2"}}},
			},
		},
		"extra and missing rule gets returned": {
			args: args{
				first:  &[]nextgen.Clause{{Attribute: "foo", Op: "equal", Values: []string{"bar"}}, {Attribute: "foo2", Op: "equal", Values: []string{"bar2"}}},
				second: &[]nextgen.Clause{{Attribute: "foo", Op: "equal", Values: []string{"bar"}}, {Attribute: "foo3", Op: "equal", Values: []string{"bar3"}}},
			},
			expected: expected{
				extraRules:   []nextgen.Clause{{Attribute: "foo2", Op: "equal", Values: []string{"bar2"}}},
				missingRules: []nextgen.Clause{{Attribute: "foo3", Op: "equal", Values: []string{"bar3"}}},
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			extraRules, missingRules := featureflagtargetgroup.RuleDiffs(tt.args.first, tt.args.second)
			if !reflect.DeepEqual(extraRules, tt.expected.extraRules) {
				t.Errorf("Expected extraRules to be %v, got %v", tt.expected.extraRules, extraRules)
			}
			if !reflect.DeepEqual(missingRules, tt.expected.missingRules) {
				t.Errorf("Expected missingRules to be %v, got %v", tt.expected.missingRules, missingRules)
			}
		})
	}
}

func TestIncludeRuleDiffs(t *testing.T) {
	type args struct {
		first  []string
		second []string
	}
	type expected struct {
		extraRules   []string
		missingRules []string
	}
	tests := map[string]struct {
		args     args
		expected expected
	}{
		"no rules have no changes": {
			args: args{},
			expected: expected{
				extraRules:   nil,
				missingRules: nil,
			},
		},
		"identical rules have no changes": {
			args: args{
				first:  []string{"foo"},
				second: []string{"foo"},
			},
			expected: expected{
				extraRules:   nil,
				missingRules: nil,
			},
		},
		"order of rules doesn't matter": {
			args: args{
				first:  []string{"foo", "foo2"},
				second: []string{"foo2", "foo"},
			},
			expected: expected{
				extraRules:   nil,
				missingRules: nil,
			},
		},
		"extra rule gets returned": {
			args: args{
				first:  []string{"foo", "foo2"},
				second: []string{"foo"},
			},
			expected: expected{
				extraRules:   []string{"foo2"},
				missingRules: nil,
			},
		},
		"missing rule gets returned": {
			args: args{
				first:  []string{"foo"},
				second: []string{"foo", "foo2"},
			},
			expected: expected{
				extraRules:   nil,
				missingRules: []string{"foo2"},
			},
		},
		"extra and missing rule gets returned": {
			args: args{
				first:  []string{"foo", "foo2"},
				second: []string{"foo", "foo3"},
			},
			expected: expected{
				extraRules:   []string{"foo2"},
				missingRules: []string{"foo3"},
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			extraRules, missingRules := featureflagtargetgroup.IncludeRuleDiffs(tt.args.first, tt.args.second)
			if !reflect.DeepEqual(extraRules, tt.expected.extraRules) {
				t.Errorf("Expected extraRules to be %v, got %v", tt.expected.extraRules, extraRules)
			}
			if !reflect.DeepEqual(missingRules, tt.expected.missingRules) {
				t.Errorf("Expected missingRules to be %v, got %v", tt.expected.missingRules, missingRules)
			}
		})
	}
}
