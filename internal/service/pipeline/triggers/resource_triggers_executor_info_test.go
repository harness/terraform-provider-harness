package triggers_test

/*
# Trigger executor_info Acceptance Tests

These tests verify the executor_info block on harness_platform_triggers, which
routes through the v2 trigger API and attaches an executor identity to the trigger.

Test cases — executor_info resource (v2 API path):
  - TestAccResourceTriggers_ExecutorInfo_SA_ProjectScoped       — SERVICE_ACCOUNT, project-scoped
  - TestAccResourceTriggers_ExecutorInfo_SA_OrgScoped           — SERVICE_ACCOUNT, org-scoped
  - TestAccResourceTriggers_ExecutorInfo_SA_AccountScoped       — SERVICE_ACCOUNT, account-scoped
  - TestAccResourceTriggers_ExecutorInfo_User                   — USER executor
  - TestAccResourceTriggers_ExecutorInfo_NoExecutorInfo         — no executor_info block (v1 API path)
  - TestAccResourceTriggers_ExecutorInfo_OmitAfterSet           — omit block after set; executor retained (Optional+Computed)
  - TestAccResourceTriggers_ExecutorInfo_UpdateYAMLKeepExecutor — YAML update preserves executor (Harness SDK preserve-on-omission)
  - TestAccResourceTriggers_ExecutorInfo_ImportThenUpdate       — import then update executor
  - TestAccResourceTriggers_ExecutorInfo_Import                 — import populates executor_info in state
  - TestAccResourceTriggers_ExecutorInfo_UpdateExecutor         — change executor from one SA to another
  - TestAccResourceTriggers_ExecutorInfo_ExternalChange         — drift detection + apply reconciliation

Test cases — validation (CustomizeDiff + ValidateFunc, no real API call):
  - TestAccResourceTriggers_ExecutorInfo_Validation_InvalidType         — rejects invalid executor type
  - TestAccResourceTriggers_ExecutorInfo_Validation_SAProjectWithoutOrg — rejects project_id without org_id
  - TestAccResourceTriggers_ExecutorInfo_Validation_UserWithOrgID       — rejects org_id on USER executor

Test cases — v1 resource (no executor_info):
  - TestAccResourceTriggersV1_CreateUpdate             — v1 create, update, import
  - TestAccResourceTriggersV1_DeleteUnderlyingResource — out-of-band delete detected by plan

Test cases — data source:
  - TestAccDataSourceTriggersV1                            — v1 data source
  - TestAccDataSourceTriggers_ExecutorInfo_WithExecutor    — data source exposes executor_info
  - TestAccDataSourceTriggers_ExecutorInfo_NoExecutor      — data source reports no executor

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
	"regexp"
	"testing"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
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

// TestAccResourceTriggers_ExecutorInfo_OmitAfterSet creates a trigger with a
// project-scoped SA executor_info, then omits the block in a second step.
//
// executor_info is Optional+Computed. Harness preserves an existing executor
// when an update omits executorInfo (this is the default behavior of the SDK)
// Therefore, omitting the block retains the executor in
// both Terraform state and on the server — it is NOT a removal operation.
// Plans remain clean because state and server agree.
func TestAccResourceTriggers_ExecutorInfo_OmitAfterSet(t *testing.T) {
	saIdentifier := os.Getenv("TRIGGER_TEST_SA_PROJECT_IDENTIFIER")
	if saIdentifier == "" {
		t.Skip("TRIGGER_TEST_SA_PROJECT_IDENTIFIER must be set for this acceptance test")
	}

	env := getTriggerExecutorTestEnv(t)
	id := fmt.Sprintf("trgexecomit_%s", utils.RandStringBytes(5))
	resourceName := "harness_platform_triggers.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				// Step 1: create with executor_info set.
				Config: testAccTriggerWithExecutorInfo_SA_ProjectScoped(env, id, saIdentifier),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "executor_info.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "executor_info.0.identifier", saIdentifier),
				),
			},
			{
				// Step 2: omit executor_info block — executor is Optional+Computed so
				// Terraform retains what the server reports. State must still show the
				// executor and the plan must be clean.
				Config: testAccTriggerWithoutExecutorInfo(env, id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "executor_info.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "executor_info.0.identifier", saIdentifier),
					testAccCheckTriggerServerHasExecutor(env, id, saIdentifier),
				),
			},
			{
				// Step 3: confirm no perpetual diff.
				Config:   testAccTriggerWithoutExecutorInfo(env, id),
				PlanOnly: true,
			},
		},
	})
}

// TestAccResourceTriggers_ExecutorInfo_UpdateYAMLKeepExecutor creates a trigger with
// an executor, then updates the YAML without changing executor_info, and verifies
// that Harness preserves the executor (exercises the Harness SDK preserve-on-omission path).
func TestAccResourceTriggers_ExecutorInfo_UpdateYAMLKeepExecutor(t *testing.T) {
	saIdentifier := os.Getenv("TRIGGER_TEST_SA_PROJECT_IDENTIFIER")
	if saIdentifier == "" {
		t.Skip("TRIGGER_TEST_SA_PROJECT_IDENTIFIER must be set for this acceptance test")
	}

	env := getTriggerExecutorTestEnv(t)
	id := fmt.Sprintf("trgexecyaml_%s", utils.RandStringBytes(5))
	resourceName := "harness_platform_triggers.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				// Step 1: create with executor and branch filter hsf143_b01.
				Config: testAccTriggerWithExecutorInfo_SA_ProjectScoped(env, id, saIdentifier),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "executor_info.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "executor_info.0.identifier", saIdentifier),
				),
			},
			{
				// Step 2: update YAML (change branch filter) without touching executor.
				// Harness must preserve the existing executor server-side.
				Config: testAccTriggerWithExecutorInfoAndBranch(env, id, saIdentifier, "hsf143_b02"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "executor_info.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "executor_info.0.identifier", saIdentifier),
					testAccCheckTriggerServerHasExecutor(env, id, saIdentifier),
				),
			},
			{
				// Step 3: clean plan.
				Config:   testAccTriggerWithExecutorInfoAndBranch(env, id, saIdentifier, "hsf143_b02"),
				PlanOnly: true,
			},
		},
	})
}

// TestAccResourceTriggers_ExecutorInfo_ImportThenUpdate imports a trigger with an
// executor then performs an update, verifying the imported resource remains manageable.
func TestAccResourceTriggers_ExecutorInfo_ImportThenUpdate(t *testing.T) {
	saIdentifier := os.Getenv("TRIGGER_TEST_SA_PROJECT_IDENTIFIER")
	if saIdentifier == "" {
		t.Skip("TRIGGER_TEST_SA_PROJECT_IDENTIFIER must be set for this acceptance test")
	}
	saIdentifier2 := os.Getenv("TRIGGER_TEST_SA_ORG_IDENTIFIER")
	if saIdentifier2 == "" {
		t.Skip("TRIGGER_TEST_SA_ORG_IDENTIFIER must be set for this acceptance test (used as updated executor)")
	}

	env := getTriggerExecutorTestEnv(t)
	id := fmt.Sprintf("trgexecimupd_%s", utils.RandStringBytes(5))
	resourceName := "harness_platform_triggers.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				// Step 1: create with executor.
				Config: testAccTriggerWithExecutorInfo_SA_ProjectScoped(env, id, saIdentifier),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "executor_info.0.identifier", saIdentifier),
				),
			},
			{
				// Step 2: import — state rebuilt from server.
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateIdFunc: triggerImportStateID(env, id),
				ImportStateVerify: true,
			},
			{
				// Step 3: update executor after import — verifies imported resource is manageable.
				Config: testAccTriggerWithExecutorInfo_SA_OrgScoped(env, id, saIdentifier2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "executor_info.0.identifier", saIdentifier2),
					testAccCheckTriggerServerHasExecutor(env, id, saIdentifier2),
				),
			},
			{
				// Step 4: clean plan after update.
				Config:   testAccTriggerWithExecutorInfo_SA_OrgScoped(env, id, saIdentifier2),
				PlanOnly: true,
			},
		},
	})
}

// testAccCheckTriggerServerHasExecutor fetches the trigger directly from the Harness
// API and asserts that the server-side executor identifier matches the expected value.
// This verifies that removing executor_info from .tf config does NOT delete the
// executor from Harness — it only stops Terraform from tracking it.
func testAccCheckTriggerServerHasExecutor(env triggerExecutorTestEnv, id, expectedIdentifier string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		acctest.TestAccConfigureProvider()
		c, ctx := acctest.TestAccGetPlatformClientWithContext()
		resp, _, err := c.TriggersApi.GetTrigger(ctx, c.AccountId, env.orgID, env.projectID, env.pipelineID, id)
		if err != nil {
			return fmt.Errorf("error fetching trigger %s from Harness API: %w", id, err)
		}
		if resp.Data == nil {
			return fmt.Errorf("trigger %s not found on Harness server", id)
		}
		if resp.Data.ExecutorInfo == nil {
			return fmt.Errorf("expected Harness server to retain executor on trigger %s after Terraform removal, but ExecutorInfo is nil", id)
		}
		if resp.Data.ExecutorInfo.Uuid != expectedIdentifier {
			return fmt.Errorf("expected server executor identifier %q, got %q", expectedIdentifier, resp.Data.ExecutorInfo.Uuid)
		}
		return nil
	}
}

// TestAccResourceTriggers_ExecutorInfo_Import creates a trigger with a project-scoped
// SA executor_info, then imports it into a fresh state and verifies that executor_info
// is fully populated in state after import (no spurious plan diff).
func TestAccResourceTriggers_ExecutorInfo_Import(t *testing.T) {
	saIdentifier := os.Getenv("TRIGGER_TEST_SA_PROJECT_IDENTIFIER")
	if saIdentifier == "" {
		t.Skip("TRIGGER_TEST_SA_PROJECT_IDENTIFIER must be set for this acceptance test")
	}

	env := getTriggerExecutorTestEnv(t)
	id := fmt.Sprintf("trgexecimport_%s", utils.RandStringBytes(5))
	resourceName := "harness_platform_triggers.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				// Step 1: create the trigger with executor_info.
				Config: testAccTriggerWithExecutorInfo_SA_ProjectScoped(env, id, saIdentifier),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "executor_info.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "executor_info.0.type", "SERVICE_ACCOUNT"),
					resource.TestCheckResourceAttr(resourceName, "executor_info.0.identifier", saIdentifier),
				),
			},
			{
				// Step 2: import — state is rebuilt from scratch via ReadContext.
				// executor_info must be fully populated in state so the next plan is clean.
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateIdFunc: triggerImportStateID(env, id),
				ImportStateVerify: true,
			},
		},
	})
}

// TestAccResourceTriggers_ExecutorInfo_UpdateExecutor creates a trigger with one
// project-scoped SA executor, then updates it to a different SA executor identifier,
// verifying both Terraform state and the Harness API response reflect the new executor.
func TestAccResourceTriggers_ExecutorInfo_UpdateExecutor(t *testing.T) {
	saIdentifier := os.Getenv("TRIGGER_TEST_SA_PROJECT_IDENTIFIER")
	if saIdentifier == "" {
		t.Skip("TRIGGER_TEST_SA_PROJECT_IDENTIFIER must be set for this acceptance test")
	}
	// Use the org-scoped SA as the second executor to switch to.
	saIdentifier2 := os.Getenv("TRIGGER_TEST_SA_ORG_IDENTIFIER")
	if saIdentifier2 == "" {
		t.Skip("TRIGGER_TEST_SA_ORG_IDENTIFIER must be set for this acceptance test (used as second executor)")
	}

	env := getTriggerExecutorTestEnv(t)
	id := fmt.Sprintf("trgexecupd_%s", utils.RandStringBytes(5))
	resourceName := "harness_platform_triggers.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				// Step 1: create with first SA executor (project-scoped).
				Config: testAccTriggerWithExecutorInfo_SA_ProjectScoped(env, id, saIdentifier),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "executor_info.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "executor_info.0.identifier", saIdentifier),
					resource.TestCheckResourceAttr(resourceName, "executor_info.0.type", "SERVICE_ACCOUNT"),
				),
			},
			{
				// Step 2: update to second SA executor (org-scoped).
				// Verifies that executor identity can be changed after creation.
				Config: testAccTriggerWithExecutorInfo_SA_OrgScoped(env, id, saIdentifier2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "executor_info.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "executor_info.0.identifier", saIdentifier2),
					resource.TestCheckResourceAttr(resourceName, "executor_info.0.type", "SERVICE_ACCOUNT"),
					resource.TestCheckResourceAttr(resourceName, "executor_info.0.org_id", env.orgID),
					resource.TestCheckResourceAttr(resourceName, "executor_info.0.project_id", ""),
					// Verify the Harness server also reflects the updated executor.
					testAccCheckTriggerServerHasExecutor(env, id, saIdentifier2),
				),
			},
		},
	})
}

// TestAccResourceTriggers_ExecutorInfo_ExternalChange creates a trigger with a
// project-scoped SA executor, changes the executor out-of-band, verifies drift
// is detected, then applies to reconcile back to the configured executor.
func TestAccResourceTriggers_ExecutorInfo_ExternalChange(t *testing.T) {
	saIdentifier := os.Getenv("TRIGGER_TEST_SA_PROJECT_IDENTIFIER")
	if saIdentifier == "" {
		t.Skip("TRIGGER_TEST_SA_PROJECT_IDENTIFIER must be set for this acceptance test")
	}
	saIdentifier2 := os.Getenv("TRIGGER_TEST_SA_ORG_IDENTIFIER")
	if saIdentifier2 == "" {
		t.Skip("TRIGGER_TEST_SA_ORG_IDENTIFIER must be set for this acceptance test (used as out-of-band executor)")
	}

	env := getTriggerExecutorTestEnv(t)
	id := fmt.Sprintf("trgexecdrift_%s", utils.RandStringBytes(5))
	resourceName := "harness_platform_triggers.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				// Step 1: create with project-scoped SA executor.
				Config: testAccTriggerWithExecutorInfo_SA_ProjectScoped(env, id, saIdentifier),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "executor_info.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "executor_info.0.identifier", saIdentifier),
				),
			},
			{
				// Step 2: change executor out-of-band; Terraform must detect drift.
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					c, ctx := acctest.TestAccGetPlatformClientWithContext()
					getResp, _, err := c.TriggersApi.GetTrigger(ctx, c.AccountId, env.orgID, env.projectID, env.pipelineID, id)
					if err != nil || getResp.Data == nil {
						t.Fatalf("failed to fetch trigger for out-of-band update: %v", err)
					}
					reqBody := nextgen.TriggerYamlRequest{
						Yaml: getResp.Data.Yaml,
						ExecutorInfo: &nextgen.TriggerExecutorInfo{
							Type_:             "SERVICE_ACCOUNT",
							Uuid:              saIdentifier2,
							AccountIdentifier: c.AccountId,
							OrgIdentifier:     env.orgID,
						},
					}
					_, _, err = c.TriggersApi.UpdateTriggerV2(ctx, reqBody, c.AccountId, env.orgID, env.projectID, env.pipelineID, id, nil)
					if err != nil {
						t.Fatalf("out-of-band executor update failed: %v", err)
					}
				},
				Config:             testAccTriggerWithExecutorInfo_SA_ProjectScoped(env, id, saIdentifier),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
			{
				// Step 3: apply to reconcile — Terraform restores saIdentifier.
				Config: testAccTriggerWithExecutorInfo_SA_ProjectScoped(env, id, saIdentifier),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "executor_info.0.identifier", saIdentifier),
					testAccCheckTriggerServerHasExecutor(env, id, saIdentifier),
				),
			},
			{
				// Step 4: clean plan after reconciliation.
				Config:   testAccTriggerWithExecutorInfo_SA_ProjectScoped(env, id, saIdentifier),
				PlanOnly: true,
			},
		},
	})
}

// ---------------------------------------------------------------------------
// Validation tests (CustomizeDiff)
// ---------------------------------------------------------------------------

// TestAccResourceTriggers_ExecutorInfo_Validation_InvalidType verifies that an
// executor type other than USER or SERVICE_ACCOUNT is rejected at plan time by
// the schema ValidateFunc before any API call is made.
func TestAccResourceTriggers_ExecutorInfo_Validation_InvalidType(t *testing.T) {
	env := getTriggerExecutorTestEnv(t)
	id := fmt.Sprintf("trgvalinvtype_%s", utils.RandStringBytes(5))

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTriggerValidation_InvalidType(env, id),
				ExpectError: regexp.MustCompile(`expected.*to be one of.*USER.*SERVICE_ACCOUNT`),
			},
		},
	})
}

// TestAccResourceTriggers_ExecutorInfo_Validation_SAProjectWithoutOrg verifies that
// setting project_id without org_id for a SERVICE_ACCOUNT executor is rejected at
// plan time by CustomizeDiff.
func TestAccResourceTriggers_ExecutorInfo_Validation_SAProjectWithoutOrg(t *testing.T) {
	env := getTriggerExecutorTestEnv(t)
	id := fmt.Sprintf("trgvalproj_%s", utils.RandStringBytes(5))

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTriggerValidation_SAProjectWithoutOrg(env, id),
				ExpectError: regexp.MustCompile("project_id requires org_id"),
			},
		},
	})
}

// TestAccResourceTriggers_ExecutorInfo_Validation_UserWithOrgID verifies that
// setting org_id on a USER executor is rejected at plan time by CustomizeDiff.
func TestAccResourceTriggers_ExecutorInfo_Validation_UserWithOrgID(t *testing.T) {
	env := getTriggerExecutorTestEnv(t)
	id := fmt.Sprintf("trgvalusr_%s", utils.RandStringBytes(5))
	userUUID := os.Getenv("TRIGGER_TEST_USER_UUID")
	if userUUID == "" {
		t.Skip("TRIGGER_TEST_USER_UUID must be set for this acceptance test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTriggerValidation_UserWithOrgID(env, id, userUUID),
				ExpectError: regexp.MustCompile("org_id must not be set for a USER executor"),
			},
		},
	})
}

// triggerImportStateID returns the import ID in the format org/project/target/identifier.
func triggerImportStateID(env triggerExecutorTestEnv, id string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		return fmt.Sprintf("%s/%s/%s/%s", env.orgID, env.projectID, env.pipelineID, id), nil
	}
}

// ---------------------------------------------------------------------------
// V1 API path tests (no executor_info)
// These tests cover the v1 trigger API path and run against pre-existing
// org/project/pipeline resources defined via env vars (same as executor_info
// tests). Run as a group with: -run TestAccResourceTriggersV1
// ---------------------------------------------------------------------------

// TestAccResourceTriggersV1_CreateUpdate creates a trigger via the v1 API path,
// updates it in place, then imports it.
func TestAccResourceTriggersV1_CreateUpdate(t *testing.T) {
	env := getTriggerExecutorTestEnv(t)
	id := fmt.Sprintf("trgv1cu_%s", utils.RandStringBytes(5))
	updatedName := fmt.Sprintf("%s_upd", id)
	resourceName := "harness_platform_triggers.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				// Step 1: create via v1 path
				Config: testAccTriggerWithoutExecutorInfo(env, id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", id),
					resource.TestCheckResourceAttr(resourceName, "target_id", env.pipelineID),
					resource.TestCheckResourceAttr(resourceName, "executor_info.#", "0"),
				),
			},
			{
				// Step 2: update name via v1 path
				Config: testAccTriggerV1WithName(env, id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "executor_info.#", "0"),
				),
			},
			{
				// Step 3: import and verify state matches
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateIdFunc: triggerImportStateID(env, id),
				ImportStateVerify: true,
			},
		},
	})
}

// TestAccResourceTriggersV1_DeleteUnderlyingResource creates a trigger then
// deletes it out-of-band, verifying Terraform detects the drift on next plan.
func TestAccResourceTriggersV1_DeleteUnderlyingResource(t *testing.T) {
	env := getTriggerExecutorTestEnv(t)
	id := fmt.Sprintf("trgv1del_%s", utils.RandStringBytes(5))
	resourceName := "harness_platform_triggers.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTriggerWithoutExecutorInfo(env, id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
				),
			},
			{
				// Delete the trigger out-of-band then verify Terraform sees a non-empty plan.
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					c, ctx := acctest.TestAccGetPlatformClientWithContext()
					c.TriggersApi.DeleteTrigger(ctx, c.AccountId, env.orgID, env.projectID, env.pipelineID, id, nil)
				},
				Config:             testAccTriggerWithoutExecutorInfo(env, id),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

// TestAccDataSourceTriggersV1 creates a trigger via v1 path then reads it back
// via the data source, verifying all fields are populated correctly.
func TestAccDataSourceTriggersV1(t *testing.T) {
	env := getTriggerExecutorTestEnv(t)
	id := fmt.Sprintf("trgv1ds_%s", utils.RandStringBytes(5))
	dataSourceName := "data.harness_platform_triggers.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTriggerWithoutExecutorInfo(env, id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "identifier", id),
					resource.TestCheckResourceAttr(dataSourceName, "org_id", env.orgID),
					resource.TestCheckResourceAttr(dataSourceName, "project_id", env.projectID),
					resource.TestCheckResourceAttr(dataSourceName, "target_id", env.pipelineID),
					resource.TestCheckResourceAttr(dataSourceName, "executor_info.#", "0"),
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
// Data source tests
// ---------------------------------------------------------------------------

// TestAccDataSourceTriggers_ExecutorInfo_WithExecutor creates a trigger with a
// project-scoped SA executor_info, then reads it back via the data source and
// verifies executor_info is populated correctly.
func TestAccDataSourceTriggers_ExecutorInfo_WithExecutor(t *testing.T) {
	saIdentifier := os.Getenv("TRIGGER_TEST_SA_PROJECT_IDENTIFIER")
	if saIdentifier == "" {
		t.Skip("TRIGGER_TEST_SA_PROJECT_IDENTIFIER must be set for this acceptance test")
	}

	env := getTriggerExecutorTestEnv(t)
	id := fmt.Sprintf("trgdsexec_%s", utils.RandStringBytes(5))
	dataSourceName := "data.harness_platform_triggers.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTriggerWithExecutorInfo(env, id, saIdentifier),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "identifier", id),
					resource.TestCheckResourceAttr(dataSourceName, "executor_info.#", "1"),
					resource.TestCheckResourceAttr(dataSourceName, "executor_info.0.type", "SERVICE_ACCOUNT"),
					resource.TestCheckResourceAttr(dataSourceName, "executor_info.0.identifier", saIdentifier),
					resource.TestCheckResourceAttr(dataSourceName, "executor_info.0.org_id", env.orgID),
					resource.TestCheckResourceAttr(dataSourceName, "executor_info.0.project_id", env.projectID),
				),
			},
		},
	})
}

// TestAccDataSourceTriggers_ExecutorInfo_NoExecutor creates a trigger without
// executor_info, then reads it back via the data source and verifies
// executor_info is absent.
func TestAccDataSourceTriggers_ExecutorInfo_NoExecutor(t *testing.T) {
	env := getTriggerExecutorTestEnv(t)
	id := fmt.Sprintf("trgdsnoexec_%s", utils.RandStringBytes(5))
	dataSourceName := "data.harness_platform_triggers.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTriggerWithoutExecutorInfo(env, id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "identifier", id),
					resource.TestCheckResourceAttr(dataSourceName, "executor_info.#", "0"),
				),
			},
		},
	})
}

// ---------------------------------------------------------------------------
// HCL config helpers
// ---------------------------------------------------------------------------

func triggerYamlHCLWithName(env triggerExecutorTestEnv, id, name string) string {
	return fmt.Sprintf(`<<-EOT
    trigger:
      name: %[7]s
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
  EOT`, id, env.orgID, env.projectID, env.pipelineID, env.connectorRef, env.repoName, name)
}

func triggerYamlHCL(env triggerExecutorTestEnv, id string) string {
	return triggerYamlHCLWithName(env, id, id)
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

func testAccTriggerV1WithName(env triggerExecutorTestEnv, id, name string) string {
	return fmt.Sprintf(`
resource "harness_platform_triggers" "test" {
  identifier = %[1]q
  name       = %[2]q
  org_id     = %[3]q
  project_id = %[4]q
  target_id  = %[5]q

  yaml = %[6]s
}
`, id, name, env.orgID, env.projectID, env.pipelineID, triggerYamlHCLWithName(env, id, name))
}

func testAccDataSourceTriggerWithExecutorInfo(env triggerExecutorTestEnv, id, saIdentifier string) string {
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

data "harness_platform_triggers" "test" {
  identifier = harness_platform_triggers.test.identifier
  org_id     = %[2]q
  project_id = %[3]q
  target_id  = %[4]q
}
`, id, env.orgID, env.projectID, env.pipelineID, saIdentifier, triggerYamlHCL(env, id))
}

func testAccTriggerValidation_SAProjectWithoutOrg(env triggerExecutorTestEnv, id string) string {
	return fmt.Sprintf(`
resource "harness_platform_triggers" "test" {
  identifier = %[1]q
  name       = %[1]q
  org_id     = %[2]q
  project_id = %[3]q
  target_id  = %[4]q

  executor_info {
    type       = "SERVICE_ACCOUNT"
    identifier = "sa005"
    project_id = %[3]q
    # org_id intentionally omitted to trigger validation error
  }

  yaml = %[5]s
}
`, id, env.orgID, env.projectID, env.pipelineID, triggerYamlHCL(env, id))
}

func testAccTriggerValidation_UserWithOrgID(env triggerExecutorTestEnv, id, userUUID string) string {
	return fmt.Sprintf(`
resource "harness_platform_triggers" "test" {
  identifier = %[1]q
  name       = %[1]q
  org_id     = %[2]q
  project_id = %[3]q
  target_id  = %[4]q

  executor_info {
    type       = "USER"
    identifier = %[5]q
    org_id     = %[2]q
    # org_id must not be set for USER executor
  }

  yaml = %[6]s
}
`, id, env.orgID, env.projectID, env.pipelineID, userUUID, triggerYamlHCL(env, id))
}

func testAccTriggerWithExecutorInfoAndBranch(env triggerExecutorTestEnv, id, saIdentifier, branch string) string {
	yaml := fmt.Sprintf(`<<-EOT
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
                  value: %[7]s
              headerConditions: []
              actions: []
  EOT`, id, env.orgID, env.projectID, env.pipelineID, env.connectorRef, env.repoName, branch)

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
`, id, env.orgID, env.projectID, env.pipelineID, saIdentifier, yaml)
}

func testAccTriggerValidation_InvalidType(env triggerExecutorTestEnv, id string) string {
	return fmt.Sprintf(`
resource "harness_platform_triggers" "test" {
  identifier = %[1]q
  name       = %[1]q
  org_id     = %[2]q
  project_id = %[3]q
  target_id  = %[4]q

  executor_info {
    type       = "ADMIN"
    identifier = "someuser"
  }

  yaml = %[5]s
}
`, id, env.orgID, env.projectID, env.pipelineID, triggerYamlHCL(env, id))
}

func testAccDataSourceTriggerWithoutExecutorInfo(env triggerExecutorTestEnv, id string) string {
	return fmt.Sprintf(`
resource "harness_platform_triggers" "test" {
  identifier = %[1]q
  name       = %[1]q
  org_id     = %[2]q
  project_id = %[3]q
  target_id  = %[4]q

  yaml = %[5]s
}

data "harness_platform_triggers" "test" {
  identifier = harness_platform_triggers.test.identifier
  org_id     = %[2]q
  project_id = %[3]q
  target_id  = %[4]q
}
`, id, env.orgID, env.projectID, env.pipelineID, triggerYamlHCL(env, id))
}
