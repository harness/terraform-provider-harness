/*
Unit tests for AutoStopping rule data sources (all 5 types: K8s, VM, ECS, RDS, Scale Group).
Tests schema validation only — no API calls, no credentials needed.

Run all unit tests:

	go test -v ./internal/service/platform/autostopping/rule/... \
	  -run "TestUnitDataSource" -timeout 5m

*/

package as_rule_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// --- K8s ---

func TestUnitDataSourceK8sRuleSchemaNoArgs(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testDataSourceConfig("harness_autostopping_rule_k8s", "test", ""),
				ExpectError: regexp.MustCompile(`one of .identifier,name. must be specified`),
			},
		},
	})
}

// --- VM ---

func TestUnitDataSourceVMRuleSchemaNoArgs(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testDataSourceConfig("harness_autostopping_rule_vm", "test", ""),
				ExpectError: regexp.MustCompile(`one of .identifier,name. must be specified`),
			},
		},
	})
}

// --- ECS ---

func TestUnitDataSourceECSRuleSchemaNoArgs(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testDataSourceConfig("harness_autostopping_rule_ecs", "test", ""),
				ExpectError: regexp.MustCompile(`one of .identifier,name. must be specified`),
			},
		},
	})
}

// --- RDS ---

func TestUnitDataSourceRDSRuleSchemaNoArgs(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testDataSourceConfig("harness_autostopping_rule_rds", "test", ""),
				ExpectError: regexp.MustCompile(`one of .identifier,name. must be specified`),
			},
		},
	})
}

// --- Scale Group ---

func TestUnitDataSourceScaleGroupRuleSchemaNoArgs(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testDataSourceConfig("harness_autostopping_rule_scale_group", "test", ""),
				ExpectError: regexp.MustCompile(`one of .identifier,name. must be specified`),
			},
		},
	})
}

// --- Helpers ---

func testDataSourceConfig(dsType, label, body string) string {
	return fmt.Sprintf(`
data %q %q {
  %s
}
`, dsType, label, body)
}
