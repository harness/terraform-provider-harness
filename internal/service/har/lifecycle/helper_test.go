package lifecycle_test

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func randAlphanumeric(n int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, n)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

type ruleScope struct {
	accountId string
	orgId     string
	projectId string
}

type scopedCase struct {
	name  string
	scope ruleScope
}

func allScopeCases() []scopedCase {
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	orgId := os.Getenv("HARNESS_ORG_ID")
	projectId := os.Getenv("HARNESS_PROJECT_ID")

	cases := []scopedCase{
		{"account", ruleScope{accountId: accountId}},
	}
	if orgId != "" {
		cases = append(cases, scopedCase{"org", ruleScope{accountId: accountId, orgId: orgId}})
	}
	if orgId != "" && projectId != "" {
		cases = append(cases, scopedCase{"project", ruleScope{accountId: accountId, orgId: orgId, projectId: projectId}})
	}
	return cases
}

func (s ruleScope) hcl() string {
	out := fmt.Sprintf(`account_id = "%s"`, s.accountId)
	if s.orgId != "" {
		out += fmt.Sprintf("\n  org_id     = \"%s\"", s.orgId)
	}
	if s.projectId != "" {
		out += fmt.Sprintf("\n  project_id = \"%s\"", s.projectId)
	}
	return out
}

func testAccLifecycleRuleBasic(name string, sc ruleScope) string {
	return fmt.Sprintf(`
resource "harness_platform_har_lifecycle_rule" "test" {
  %[1]s
  name   = "%[2]s"
  action = "DELETE"

  apply_to {
    mode = "ALL_IN_SCOPE"
  }
}
`, sc.hcl(), name)
}

func testAccLifecycleRuleProtect(name string, sc ruleScope) string {
	return fmt.Sprintf(`
resource "harness_platform_har_lifecycle_rule" "test" {
  %[1]s
  name   = "%[2]s"
  action = "PROTECT"

  apply_to {
    mode = "ALL_IN_SCOPE"
  }
}
`, sc.hcl(), name)
}

func testAccLifecycleRuleWithKeepLastN(name string, sc ruleScope, keepN int) string {
	return fmt.Sprintf(`
resource "harness_platform_har_lifecycle_rule" "test" {
  %[1]s
  name        = "%[2]s"
  action      = "DELETE"
  description = "Keep last %[3]d versions"

  apply_to {
    mode = "ALL_IN_SCOPE"
  }

  criteria {
    match = "ALL"
    rules {
      type  = "KEEP_LAST_N"
      value = %[3]d
    }
  }
}
`, sc.hcl(), name, keepN)
}

func testAccLifecycleRuleWithAgeBased(name string, sc ruleScope) string {
	return fmt.Sprintf(`
resource "harness_platform_har_lifecycle_rule" "test" {
  %[1]s
  name   = "%[2]s"
  action = "DELETE"

  apply_to {
    mode = "ALL_IN_SCOPE"
  }

  criteria {
    match = "ALL"
    rules {
      type  = "AGE_BASED"
      value = 30
      unit  = "DAYS"
    }
  }
}
`, sc.hcl(), name)
}

func testAccLifecycleRuleWithMultipleCriteria(name string, sc ruleScope) string {
	return fmt.Sprintf(`
resource "harness_platform_har_lifecycle_rule" "test" {
  %[1]s
  name   = "%[2]s"
  action = "DELETE"

  apply_to {
    mode = "ALL_IN_SCOPE"
  }

  criteria {
    match = "ANY"
    rules {
      type  = "KEEP_LAST_N"
      value = 5
    }
    rules {
      type  = "AGE_BASED"
      value = 90
      unit  = "DAYS"
    }
  }
}
`, sc.hcl(), name)
}

func testAccLifecycleRuleWithSchedule(name string, sc ruleScope) string {
	return fmt.Sprintf(`
resource "harness_platform_har_lifecycle_rule" "test" {
  %[1]s
  name   = "%[2]s"
  action = "DELETE"

  apply_to {
    mode = "ALL_IN_SCOPE"
  }

  criteria {
    match = "ALL"
    rules {
      type  = "KEEP_LAST_N"
      value = 10
    }
  }

  schedule {
    expression = "0 2 * * *"
    timezone   = "UTC"
  }
}
`, sc.hcl(), name)
}

func testAccLifecycleRuleWithDockerFilter(name string, sc ruleScope) string {
	return fmt.Sprintf(`
resource "harness_platform_har_lifecycle_rule" "test" {
  %[1]s
  name         = "%[2]s"
  action       = "DELETE"
  package_type = "DOCKER"

  apply_to {
    mode = "ALL_IN_SCOPE"
  }

  criteria {
    match = "ALL"
    rules {
      type  = "KEEP_LAST_N"
      value = 5
    }
  }

  filter_config {
    package_type                 = "DOCKER"
    package_name_allowed_pattern = ["my-app*", "service-*"]
    tag_name_allowed_pattern     = ["v1.*", "release-*"]
  }
}
`, sc.hcl(), name)
}

func testAccLifecycleRuleUpdated(name string, sc ruleScope, keepN int) string {
	return fmt.Sprintf(`
resource "harness_platform_har_lifecycle_rule" "test" {
  %[1]s
  name        = "%[2]s-updated"
  action      = "DELETE"
  description = "Updated rule"

  apply_to {
    mode = "ALL_IN_SCOPE"
  }

  criteria {
    match = "ALL"
    rules {
      type  = "KEEP_LAST_N"
      value = %[3]d
    }
  }

  schedule {
    expression = "0 3 * * 0"
    timezone   = "UTC"
  }
}
`, sc.hcl(), name, keepN)
}
