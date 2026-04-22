package split_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// TestAccResourceFMEApiKey_invalidType expects schema validation to reject an unsupported api_key_type.
func TestAccResourceFMEApiKey_invalidType(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping FME API key negative acceptance test in short mode")
	}
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	envName := "tf" + testAccFMEAlphanum(10)
	keyName := "tf_bad_key_" + testAccFMEAlphanum(8)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccResourceFMEApiKeyInvalidType(id, envName, keyName),
				ExpectError: regexp.MustCompile(`(?i)server_side|client_side|expected`),
			},
		},
	})
}

func testAccResourceFMEApiKeyInvalidType(id, envName, keyName string) string {
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

	resource "harness_fme_api_key" "test" {
		org_id         = harness_platform_organization.test.id
		project_id     = harness_platform_project.test.id
		name           = "%[3]s"
		api_key_type   = "not_a_supported_key_type"
		environment_id = harness_fme_environment.test.environment_id
		depends_on     = [harness_fme_environment.test]
	}
	`, id, envName, keyName)
}

// TestAccResourceFMERuleBasedSegment_invalidDefinitionJSON expects apply to fail when harness_fme_rule_based_segment_environment_association.definition_json is not valid JSON.
func TestAccResourceFMERuleBasedSegment_invalidDefinitionJSON(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping FME RBS negative acceptance test in short mode")
	}
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	envName := "tf" + testAccFMEAlphanum(10)
	rbsName := "tfrbsbad_" + testAccFMEAlphanum(8)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccResourceFMERuleBasedSegmentBadJSON(id, envName, rbsName),
				ExpectError: regexp.MustCompile(`(?i)invalid|unmarshal|json|character`),
			},
		},
	})
}

func testAccResourceFMERuleBasedSegmentBadJSON(id, envName, rbsName string) string {
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

	resource "harness_fme_rule_based_segment" "test" {
		org_id          = harness_platform_organization.test.id
		project_id      = harness_platform_project.test.id
		traffic_type_id = data.harness_fme_traffic_type.user.traffic_type_id
		name            = "%[3]s"
		depends_on      = [harness_fme_environment.test]
	}

	resource "harness_fme_rule_based_segment_environment_association" "test" {
		org_id          = harness_platform_organization.test.id
		project_id      = harness_platform_project.test.id
		environment_id  = harness_fme_environment.test.environment_id
		segment_name    = harness_fme_rule_based_segment.test.name
		definition_json = "NOT VALID JSON {{{"
		depends_on      = [harness_fme_rule_based_segment.test, harness_fme_environment.test]
	}
	`, id, envName, rbsName)
}

// TestAccResourceFMEFeatureFlagDefinition_invalidDefinition expects create to fail when definition is not valid JSON.
func TestAccResourceFMEFeatureFlagDefinition_invalidDefinition(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping FME feature flag definition negative acceptance test in short mode")
	}
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	envName := "tf" + testAccFMEAlphanum(10)
	flagName := "tfflagbad_" + testAccFMEAlphanum(8)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccResourceFMEFeatureFlagDefinitionInvalidJSON(id, envName, flagName),
				ExpectError: regexp.MustCompile(`(?i)invalid|unmarshal|json|character`),
			},
		},
	})
}

func testAccResourceFMEFeatureFlagDefinitionInvalidJSON(id, envName, flagName string) string {
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

	resource "harness_fme_feature_flag" "test" {
		org_id          = harness_platform_organization.test.id
		project_id      = harness_platform_project.test.id
		traffic_type_id = data.harness_fme_traffic_type.user.traffic_type_id
		name            = "%[3]s"
	}

	resource "harness_fme_feature_flag_definition" "test" {
		org_id           = harness_platform_organization.test.id
		project_id       = harness_platform_project.test.id
		environment_id   = harness_fme_environment.test.environment_id
		flag_name        = harness_fme_feature_flag.test.name
		definition       = "NOT VALID JSON {{{"
		depends_on       = [harness_fme_feature_flag.test, harness_fme_environment.test]
	}
	`, id, envName, flagName)
}

// TestAccResourceFMETrafficTypeAttribute_invalidSearchableType expects Terraform to reject a non-bool for is_searchable before any API call.
func TestAccResourceFMETrafficTypeAttribute_invalidSearchableType(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping FME traffic type attribute negative acceptance test in short mode")
	}
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	attrID := "tfattrbad_" + testAccFMEAlphanum(8)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccResourceFMETrafficTypeAttributeInvalidSearchable(id, attrID),
				ExpectError: regexp.MustCompile(`(?i)bool|unsuitable|incorrect`),
			},
		},
	})
}

func testAccResourceFMETrafficTypeAttributeInvalidSearchable(id, attrIdentifier string) string {
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

	data "harness_fme_traffic_type" "user" {
		org_id     = harness_platform_organization.test.id
		project_id = harness_platform_project.test.id
		name       = "user"
	}

	resource "harness_fme_traffic_type_attribute" "test" {
		org_id             = harness_platform_organization.test.id
		project_id         = harness_platform_project.test.id
		traffic_type_id    = data.harness_fme_traffic_type.user.traffic_type_id
		identifier         = "%[2]s"
		display_name       = "bad searchable type"
		data_type          = "string"
		is_searchable      = "not_a_boolean"
	}
	`, id, attrIdentifier)
}
