package triggers_test

/*
# Trigger executor_info Acceptance Tests

These tests verify the executor_info block on harness_platform_triggers, which
routes through the v2 trigger API and attaches an executor identity to the trigger.

Test cases:
  - TestAccResourceTriggers_ExecutorInfo_SA_ProjectScoped  — SERVICE_ACCOUNT, project-scoped
  - TestAccResourceTriggers_ExecutorInfo_SA_OrgScoped      — SERVICE_ACCOUNT, org-scoped
  - TestAccResourceTriggers_ExecutorInfo_SA_AccountScoped  — SERVICE_ACCOUNT, account-scoped
  - TestAccResourceTriggers_ExecutorInfo_User              — USER executor
  - TestAccResourceTriggers_ExecutorInfo_NoExecutorInfo    — no executor_info block (v1 API path)

## Required env vars (set all to run every test in one shot)

	export TF_ACC=1
	export HARNESS_ACCOUNT_ID='<your-account-id>'

	# Provider auth key used for SA executor tests (any key with core_trigger_edit permission).
	export HARNESS_PLATFORM_API_KEY='<sa-or-personal-api-key>'

	# Existing org and project that the pipeline lives in
	export TRIGGER_TEST_ORG_ID='hsforg01'
	export TRIGGER_TEST_PROJECT_ID='hsfproj01'
	export TRIGGER_TEST_PIPELINE_ID='hsf143test01'
	export TRIGGER_TEST_CONNECTOR_REF='githubconn003'
	export TRIGGER_TEST_REPO_NAME='hands-on-workbook'

	# SERVICE_ACCOUNT executor — project-scoped SA (org_id + project_id set)
	export TRIGGER_TEST_SA_PROJECT_IDENTIFIER='sa005'

	# SERVICE_ACCOUNT executor — org-scoped SA (org_id set, no project_id)
	export TRIGGER_TEST_SA_ORG_IDENTIFIER='hsforg01sa006'

	# SERVICE_ACCOUNT executor — account-scoped SA (no org_id or project_id)
	export TRIGGER_TEST_SA_ACCOUNT_IDENTIFIER='hsftftestsvcacc01'

	# USER executor — must be a personal API key (Harness rejects USER executor
	# when the auth key is an SA token).
	export TRIGGER_TEST_USER_API_KEY='<personal-api-key>'
	export TRIGGER_TEST_USER_UUID='<user-uuid-from-harness-api>'

## Run all executor_info acceptance tests

	go test ./internal/service/pipeline/triggers/... \
	  -run TestAccResourceTriggers_ExecutorInfo \
	  -v -count=1 -timeout 120s
*/

import (
	"fmt"
	"os"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// triggerExecutorTestEnv holds the pipeline-level env vars shared across all tests.
type triggerExecutorTestEnv struct {
	accountID    string
	orgID        string
	projectID    string
	pipelineID   string
	connectorRef string
	repoName     string
}

func getTriggerExecutorTestEnv(t *testing.T) triggerExecutorTestEnv {
	t.Helper()
	env := triggerExecutorTestEnv{
		accountID:    os.Getenv("HARNESS_ACCOUNT_ID"),
		orgID:        os.Getenv("TRIGGER_TEST_ORG_ID"),
		projectID:    os.Getenv("TRIGGER_TEST_PROJECT_ID"),
		pipelineID:   os.Getenv("TRIGGER_TEST_PIPELINE_ID"),
		connectorRef: os.Getenv("TRIGGER_TEST_CONNECTOR_REF"),
		repoName:     os.Getenv("TRIGGER_TEST_REPO_NAME"),
	}
	if env.accountID == "" {
		t.Skip("HARNESS_ACCOUNT_ID must be set for executor_info acceptance tests")
	}
	missing := []string{}
	if env.orgID == "" {
		missing = append(missing, "TRIGGER_TEST_ORG_ID")
	}
	if env.projectID == "" {
		missing = append(missing, "TRIGGER_TEST_PROJECT_ID")
	}
	if env.pipelineID == "" {
		missing = append(missing, "TRIGGER_TEST_PIPELINE_ID")
	}
	if env.connectorRef == "" {
		missing = append(missing, "TRIGGER_TEST_CONNECTOR_REF")
	}
	if env.repoName == "" {
		missing = append(missing, "TRIGGER_TEST_REPO_NAME")
	}
	if len(missing) > 0 {
		t.Skipf("missing required env vars: %v", missing)
	}
	return env
}

// ---------------------------------------------------------------------------
// Acceptance tests
// ---------------------------------------------------------------------------

// TestAccResourceTriggers_ExecutorInfo_SA_ProjectScoped creates a trigger with a
// project-scoped SERVICE_ACCOUNT executor (org_id + project_id both set).
func TestAccResourceTriggers_ExecutorInfo_SA_ProjectScoped(t *testing.T) {
	saIdentifier := os.Getenv("TRIGGER_TEST_SA_PROJECT_IDENTIFIER")
	if saIdentifier == "" {
		t.Skip("TRIGGER_TEST_SA_PROJECT_IDENTIFIER must be set for this acceptance test")
	}

	env := getTriggerExecutorTestEnv(t)
	id := fmt.Sprintf("trgexecproj_%s", utils.RandStringBytes(5))
	resourceName := "harness_platform_triggers.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTriggerWithExecutorInfo_SA_ProjectScoped(env, id, saIdentifier),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "executor_info.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "executor_info.0.type", "SERVICE_ACCOUNT"),
					resource.TestCheckResourceAttr(resourceName, "executor_info.0.identifier", saIdentifier),
					resource.TestCheckResourceAttr(resourceName, "executor_info.0.org_id", env.orgID),
					resource.TestCheckResourceAttr(resourceName, "executor_info.0.project_id", env.projectID),
				),
			},
		},
	})
}

// TestAccResourceTriggers_ExecutorInfo_SA_OrgScoped creates a trigger with an
// org-scoped SERVICE_ACCOUNT executor (org_id set, project_id omitted).
func TestAccResourceTriggers_ExecutorInfo_SA_OrgScoped(t *testing.T) {
	saIdentifier := os.Getenv("TRIGGER_TEST_SA_ORG_IDENTIFIER")
	if saIdentifier == "" {
		t.Skip("TRIGGER_TEST_SA_ORG_IDENTIFIER must be set for this acceptance test")
	}

	env := getTriggerExecutorTestEnv(t)
	id := fmt.Sprintf("trgexecorg_%s", utils.RandStringBytes(5))
	resourceName := "harness_platform_triggers.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTriggerWithExecutorInfo_SA_OrgScoped(env, id, saIdentifier),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "executor_info.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "executor_info.0.type", "SERVICE_ACCOUNT"),
					resource.TestCheckResourceAttr(resourceName, "executor_info.0.identifier", saIdentifier),
					resource.TestCheckResourceAttr(resourceName, "executor_info.0.org_id", env.orgID),
					resource.TestCheckResourceAttr(resourceName, "executor_info.0.project_id", ""),
				),
			},
		},
	})
}

// TestAccResourceTriggers_ExecutorInfo_SA_AccountScoped creates a trigger with an
// account-scoped SERVICE_ACCOUNT executor (org_id and project_id both omitted).
func TestAccResourceTriggers_ExecutorInfo_SA_AccountScoped(t *testing.T) {
	saIdentifier := os.Getenv("TRIGGER_TEST_SA_ACCOUNT_IDENTIFIER")
	if saIdentifier == "" {
		t.Skip("TRIGGER_TEST_SA_ACCOUNT_IDENTIFIER must be set for this acceptance test")
	}

	env := getTriggerExecutorTestEnv(t)
	id := fmt.Sprintf("trgexecacc_%s", utils.RandStringBytes(5))
	resourceName := "harness_platform_triggers.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTriggerWithExecutorInfo_SA_AccountScoped(env, id, saIdentifier),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "executor_info.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "executor_info.0.type", "SERVICE_ACCOUNT"),
					resource.TestCheckResourceAttr(resourceName, "executor_info.0.identifier", saIdentifier),
					resource.TestCheckResourceAttr(resourceName, "executor_info.0.org_id", ""),
					resource.TestCheckResourceAttr(resourceName, "executor_info.0.project_id", ""),
				),
			},
		},
	})
}

// TestAccResourceTriggers_ExecutorInfo_User creates a trigger with a USER executor.
// Uses TRIGGER_TEST_USER_API_KEY as provider auth — Harness rejects USER executor
// when the auth key is an SA token; personal API key is required.
func TestAccResourceTriggers_ExecutorInfo_User(t *testing.T) {
	userAPIKey := os.Getenv("TRIGGER_TEST_USER_API_KEY")
	if userAPIKey == "" {
		t.Skip("TRIGGER_TEST_USER_API_KEY must be set for this acceptance test (must be a personal API key, not an SA token)")
	}
	userUUID := os.Getenv("TRIGGER_TEST_USER_UUID")
	if userUUID == "" {
		t.Skip("TRIGGER_TEST_USER_UUID must be set for this acceptance test")
	}

	env := getTriggerExecutorTestEnv(t)
	id := fmt.Sprintf("trgexecuser_%s", utils.RandStringBytes(5))
	resourceName := "harness_platform_triggers.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				// The HCL embeds the user API key directly in the provider block,
				// overriding HARNESS_PLATFORM_API_KEY for this resource.
				Config: testAccTriggerWithExecutorInfo_User(env, id, userUUID, userAPIKey),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "executor_info.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "executor_info.0.type", "USER"),
					resource.TestCheckResourceAttr(resourceName, "executor_info.0.identifier", userUUID),
				),
			},
		},
	})
}

// TestAccResourceTriggers_ExecutorInfo_NoExecutorInfo creates a trigger without
// an executor_info block, verifying the v1 API path is still used and the resource
// works correctly when executor_info is omitted.
func TestAccResourceTriggers_ExecutorInfo_NoExecutorInfo(t *testing.T) {
	env := getTriggerExecutorTestEnv(t)
	id := fmt.Sprintf("trgnoexec_%s", utils.RandStringBytes(5))
	resourceName := "harness_platform_triggers.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTriggerWithoutExecutorInfo(env, id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "executor_info.#", "0"),
				),
			},
		},
	})
}

// ---------------------------------------------------------------------------
// HCL config helpers
// ---------------------------------------------------------------------------

func triggerYamlHCL(env triggerExecutorTestEnv, id string) string {
	return fmt.Sprintf(`<<-EOT
    trigger:
      name: %[1]s
      identifier: %[1]s
      enabled: true
      stagesToExecute: []
      tags: {}
      orgIdentifier: %[2]s
      projectIdentifier: %[3]s
      pipelineIdentifier: %[4]s
      source:
        type: Webhook
        spec:
          type: Github
          spec:
            type: Push
            spec:
              connectorRef: %[5]s
              repoName: %[6]s
              autoAbortPreviousExecutions: false
              payloadConditions:
                - key: targetBranch
                  operator: Contains
                  value: hsf143_b01
              headerConditions: []
              actions: []
  EOT`, id, env.orgID, env.projectID, env.pipelineID, env.connectorRef, env.repoName)
}

func testAccTriggerWithExecutorInfo_SA_ProjectScoped(env triggerExecutorTestEnv, id, saIdentifier string) string {
	return fmt.Sprintf(`
resource "harness_platform_triggers" "test" {
  identifier = %[1]q
  name       = %[1]q
  org_id     = %[2]q
  project_id = %[3]q
  target_id  = %[4]q

  executor_info {
    type       = "SERVICE_ACCOUNT"
    identifier = %[5]q
    org_id     = %[2]q
    project_id = %[3]q
  }

  yaml = %[6]s
}
`, id, env.orgID, env.projectID, env.pipelineID, saIdentifier, triggerYamlHCL(env, id))
}

func testAccTriggerWithExecutorInfo_SA_OrgScoped(env triggerExecutorTestEnv, id, saIdentifier string) string {
	return fmt.Sprintf(`
resource "harness_platform_triggers" "test" {
  identifier = %[1]q
  name       = %[1]q
  org_id     = %[2]q
  project_id = %[3]q
  target_id  = %[4]q

  executor_info {
    type       = "SERVICE_ACCOUNT"
    identifier = %[5]q
    org_id     = %[2]q
  }

  yaml = %[6]s
}
`, id, env.orgID, env.projectID, env.pipelineID, saIdentifier, triggerYamlHCL(env, id))
}

func testAccTriggerWithExecutorInfo_SA_AccountScoped(env triggerExecutorTestEnv, id, saIdentifier string) string {
	return fmt.Sprintf(`
resource "harness_platform_triggers" "test" {
  identifier = %[1]q
  name       = %[1]q
  org_id     = %[2]q
  project_id = %[3]q
  target_id  = %[4]q

  executor_info {
    type       = "SERVICE_ACCOUNT"
    identifier = %[5]q
  }

  yaml = %[6]s
}
`, id, env.orgID, env.projectID, env.pipelineID, saIdentifier, triggerYamlHCL(env, id))
}

func testAccTriggerWithExecutorInfo_User(env triggerExecutorTestEnv, id, userUUID, userAPIKey string) string {
	return fmt.Sprintf(`
provider "harness" {
  endpoint         = "https://app.harness.io/gateway"
  account_id       = %[6]q
  platform_api_key = %[7]q
}

resource "harness_platform_triggers" "test" {
  identifier = %[1]q
  name       = %[1]q
  org_id     = %[2]q
  project_id = %[3]q
  target_id  = %[4]q

  executor_info {
    type       = "USER"
    identifier = %[5]q
  }

  yaml = %[8]s
}
`, id, env.orgID, env.projectID, env.pipelineID, userUUID, env.accountID, userAPIKey, triggerYamlHCL(env, id))
}

func testAccTriggerWithoutExecutorInfo(env triggerExecutorTestEnv, id string) string {
	return fmt.Sprintf(`
resource "harness_platform_triggers" "test" {
  identifier = %[1]q
  name       = %[1]q
  org_id     = %[2]q
  project_id = %[3]q
  target_id  = %[4]q

  yaml = %[5]s
}
`, id, env.orgID, env.projectID, env.pipelineID, triggerYamlHCL(env, id))
}
