package split_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	splitpkg "github.com/harness/terraform-provider-harness/internal/service/split"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceFMERuleBasedSegment_basic(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping FME rule-based segment acceptance test in short mode")
	}
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	envName := "tf" + testAccFMEAlphanum(10)
	rbsName := "tfrbs_" + testAccFMEAlphanum(8)
	res := "harness_fme_rule_based_segment.test"
	resAssoc := "harness_fme_rule_based_segment_environment_association.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceFMERuleBasedSegment(id, envName, rbsName, "acc rbs v1", "acc comment v1"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(res, "name", rbsName),
					resource.TestCheckResourceAttr(resAssoc, "segment_name", rbsName),
					resource.TestCheckResourceAttrSet(resAssoc, "environment_id"),
				),
			},
			{
				Config: testAccResourceFMERuleBasedSegment(id, envName, rbsName, "acc rbs v2", "acc comment v2"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(res, "name", rbsName),
				),
			},
			{
				ResourceName:      res,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: fmeImportStateIDOrgProjectThird(res, "name"),
			},
			{
				ResourceName:            resAssoc,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"definition_json"},
				ImportStateCheck:        importStateCheckFMERBSAssocDefinitionJSON(testAccFMERBSAssocDefinitionJSON("acc rbs v2", "acc comment v2")),
				ImportStateIdFunc:       fmeImportStatePrimaryID(resAssoc),
			},
		},
	})
}

func testAccResourceFMERuleBasedSegment(id, envName, rbsName, title, comment string) string {
	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
		identifier = "%[1]s"
		name       = "%[1]s"
	}

	resource "harness_platform_project" "test" {
		identifier = "%[1]s"
		org_id     = harness_platform_organization.test.id
		name       = "%[1]s"
	}

	resource "harness_fme_environment" "test" {
		org_id     = harness_platform_organization.test.id
		project_id = harness_platform_project.test.id
		name       = "%[2]s"
	}

	data "harness_fme_traffic_type" "user" {
		org_id     = harness_platform_organization.test.id
		project_id = harness_platform_project.test.id
		name       = "user"
	}

	locals {
		rbs_def = jsonencode({
			title   = "%[4]s"
			comment = "%[5]s"
			rules = [{
				condition = {
					combiner = "AND"
					matchers = [{
						type    = "IN_LIST_STRING"
						strings = ["acc_rbs_key_1"]
					}]
				}
			}]
		})
	}

	resource "harness_fme_rule_based_segment" "test" {
		org_id          = harness_platform_organization.test.id
		project_id      = harness_platform_project.test.id
		traffic_type_id = data.harness_fme_traffic_type.user.traffic_type_id
		name            = "%[3]s"
		depends_on      = [harness_fme_environment.test]
	}

	resource "harness_fme_rule_based_segment_environment_association" "test" {
		org_id            = harness_platform_organization.test.id
		project_id        = harness_platform_project.test.id
		environment_id    = harness_fme_environment.test.environment_id
		segment_name      = harness_fme_rule_based_segment.test.name
		definition_json   = local.rbs_def
		depends_on        = [harness_fme_rule_based_segment.test, harness_fme_environment.test]
	}
	`, id, envName, rbsName, title, comment)
}

// testAccFMERBSAssocDefinitionJSON matches local.rbs_def in testAccResourceFMERuleBasedSegment (same rules as HCL).
func testAccFMERBSAssocDefinitionJSON(title, comment string) string {
	b, err := json.Marshal(map[string]interface{}{
		"title":   title,
		"comment": comment,
		"rules": []interface{}{
			map[string]interface{}{
				"condition": map[string]interface{}{
					"combiner": "AND",
					"matchers": []interface{}{
						map[string]interface{}{
							"type":    "IN_LIST_STRING",
							"strings": []string{"acc_rbs_key_1"},
						},
					},
				},
			},
		},
	})
	if err != nil {
		panic(err)
	}
	return string(b)
}

func importStateCheckFMERBSAssocDefinitionJSON(wantJSON string) resource.ImportStateCheckFunc {
	const wantType = "harness_fme_rule_based_segment_environment_association"
	return func(states []*terraform.InstanceState) error {
		for _, is := range states {
			if is.Ephemeral.Type != wantType {
				continue
			}
			got, ok := is.Attributes["definition_json"]
			if !ok || got == "" {
				return fmt.Errorf("%s: missing definition_json after import", wantType)
			}
			if !splitpkg.RuleBasedSegmentDefinitionJSONSemanticallyEqualIgnoreTitleComment(wantJSON, got) {
				return fmt.Errorf("%s: definition_json (ignoring title/comment) does not match applied config: got %s", wantType, got)
			}
			return nil
		}
		return fmt.Errorf("%s not found in import state", wantType)
	}
}
