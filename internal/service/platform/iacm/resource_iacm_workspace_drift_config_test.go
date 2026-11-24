package iacm_test

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"testing"
	"time"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/harness/terraform-provider-harness/internal/acctest/contractv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const workspaceDriftConfigResourceAddress = "harness_platform_iacm_workspace_drift_config.test"

var workspaceDriftConfigDefinitions = func(p workspaceDriftConfigTests) []contractv1.ContractDefinition {
	// Each test gets a unique run_every so stale GET values can be fingerprinted
	// to the exact prior test that wrote them, isolating backend consistency issues.
	// Backend constraint: for */N cron, min(N, 60 mod N) >= 5 (wrap-around gap must be >= 5 min).
	// Invalid examples: 2,3,7,8,14,19,28,29,56,57,58,59 — all fail the cron gap check.
	mins := func(runEvery, updatedRunEvery int32) workspaceDriftConfigTests {
		q := p
		q.runEvery = runEvery
		q.updatedRunEvery = updatedRunEvery
		q.runEveryUnit = "minutes"
		return q
	}
	return []contractv1.ContractDefinition{
		{
			TestName:      contractv1.TestCreateResource,
			Supported:     true,
			Description:   "Provisions a new workspace drift config and verifies Terraform state and remote API match.",
			IsRunParallel: false,
			PreTestDelay:  5 * time.Second,
			PostTestDelay: 5 * time.Second,

			GetTerraformConfig: mins(5, 6).GetCreateResourceTfConfig,
			Verifiers:          contractv1.VerifierDefinition{VerifyCreated: mins(5, 6).verifyCreated},
		},
		{
			TestName:      contractv1.TestReadResource,
			Supported:     true,
			Description:   "Refreshes state after creation and verifies the provider's Read reconstructs all attributes.",
			IsRunParallel: false,
			PreTestDelay:  5 * time.Second,
			PostTestDelay: 5 * time.Second,

			GetTerraformConfig: mins(5, 7).GetCreateResourceTfConfig,
			Verifiers:          contractv1.VerifierDefinition{VerifyRead: mins(5, 7).verifyCreated},
		},
		{
			TestName:      contractv1.TestUpdateResource,
			Supported:     true,
			Description:   "Changes run_every and verifies the in-place update is reflected in state and remote API.",
			IsRunParallel: false,
			PreTestDelay:  5 * time.Second,
			PostTestDelay: 5 * time.Second,

			GetTerraformConfig:        mins(11, 13).GetCreateResourceTfConfig,
			GetUpdatedTerraformConfig: mins(11, 13).GetUpdatedTerraformConfig,
			Verifiers:                 contractv1.VerifierDefinition{VerifyUpdated: mins(11, 13).verifyUpdated},
		},
		{
			TestName:      contractv1.TestDeleteResource,
			Supported:     true,
			Description:   "Destroys the resource and confirms the remote drift config is disabled or absent.",
			IsRunParallel: false,
			PreTestDelay:  5 * time.Second,
			PostTestDelay: 5 * time.Second,

			GetTerraformConfig: mins(17, 19).GetCreateResourceTfConfig,
			Verifiers:          contractv1.VerifierDefinition{VerifyDeleted: mins(17, 19).verifyDeleted},
		},
		{
			TestName:      contractv1.TestCreateUpdateDeleteResource,
			Supported:     true,
			Description:   "Runs the full create → update → delete lifecycle in one scenario.",
			IsRunParallel: false,
			PreTestDelay:  5 * time.Second,
			PostTestDelay: 5 * time.Second,

			GetTerraformConfig:        mins(25, 30).GetCreateResourceTfConfig,
			GetUpdatedTerraformConfig: mins(25, 30).GetUpdatedTerraformConfig,
			Verifiers: contractv1.VerifierDefinition{
				VerifyCreated: mins(25, 30).verifyCreated,
				VerifyUpdated: mins(25, 30).verifyUpdated,
				VerifyDeleted: mins(25, 30).verifyDeleted,
			},
		},
		{
			TestName:      contractv1.TestImportResource,
			Supported:     true,
			Description:   "Imports the resource by org/project/workspace/provisioner and verifies imported state matches created state.",
			IsRunParallel: false,
			PreTestDelay:  5 * time.Second,
			PostTestDelay: 5 * time.Second,

			GetTerraformConfig:      mins(31, 37).GetCreateResourceTfConfig,
			ResourceName:            workspaceDriftConfigResourceAddress,
			GetImportID:             p.getImportID,
			ImportStateVerifyIgnore: []string{"trigger_id"},
		},
		{
			TestName:      contractv1.TestIdempotencyResource,
			Supported:     true,
			Description:   "Applies the same config twice and verifies the second plan has no changes.",
			IsRunParallel: false,
			PreTestDelay:  5 * time.Second,
			PostTestDelay: 5 * time.Second,

			GetTerraformConfig: mins(41, 43).GetCreateResourceTfConfig,
		},
		{
			TestName:      contractv1.TestExternalDeleteResource,
			Supported:     true,
			Description:   "Disables drift config via API outside Terraform and verifies a subsequent refresh detects the missing resource.",
			IsRunParallel: false,
			PreTestDelay:  5 * time.Second,
			PostTestDelay: 5 * time.Second,

			GetTerraformConfig: mins(47, 53).GetCreateResourceTfConfig,
			ExternalAction:     mins(47, 53).externalDeleteAction,
		},
		{
			TestName:      contractv1.TestExternalUpdateResource,
			Supported:     true,
			Description:   "Changes run_every via API outside Terraform and verifies a subsequent refresh detects the drift.",
			IsRunParallel: false,
			PreTestDelay:  5 * time.Second,
			PostTestDelay: 5 * time.Second,

			GetTerraformConfig: mins(55, 20).GetCreateResourceTfConfig,
			ExternalAction:     mins(55, 20).externalUpdateAction,
		},
		{
			TestName:      contractv1.TestInvalidConfigurationResource,
			Supported:     true,
			Description:   "Omits required pipeline_id and verifies Terraform returns a validation error without creating a remote resource.",
			IsRunParallel: false,
			PreTestDelay:  5 * time.Second,
			PostTestDelay: 5 * time.Second,

			GetTerraformConfig: mins(24, 26).getInvalidConfig,
			ExpectError:        regexp.MustCompile(`The argument "pipeline_id" is required`),
		},
		{
			TestName:      contractv1.TestImportNonExistentResource,
			Supported:     true,
			Description:   "Imports a syntactically valid but nonexistent org/project/workspace/provisioner ID and verifies an error is returned.",
			IsRunParallel: false,
			PreTestDelay:  5 * time.Second,
			PostTestDelay: 5 * time.Second,

			GetTerraformConfig: mins(27, 33).GetCreateResourceTfConfig,
			ResourceName:       workspaceDriftConfigResourceAddress,
			GetImportID:        p.getNonExistentImportID,
			ExpectError:        regexp.MustCompile(`(?i)(no resource|not found|does not exist|provider produced|empty|cannot import)`),
		},

		// Unsupported tests — skipped with explicit reasons.
		{
			TestName:          contractv1.TestForceNewResource,
			Supported:         false,
			UnsupportedReason: "org_id, project_id, workspace_id, and provisioner_type are ForceNew; replacement requires a second distinct fixture not available in this test environment",
			IsRunParallel:     false,
		},
		{
			TestName:          contractv1.TestAPIErrorResource,
			Supported:         false,
			UnsupportedReason: "requires controlled API failure injection not available in acceptance tests",
			IsRunParallel:     false,
		},
		{
			TestName:          contractv1.TestAsyncResource,
			Supported:         false,
			UnsupportedReason: "workspace drift PUT/DELETE returns 200 OK synchronously; no async polling path to exercise",
			IsRunParallel:     false,
		},
		{
			TestName:          contractv1.TestTimeoutResource,
			Supported:         false,
			UnsupportedReason: "requires a controlled slow API not available in acceptance tests",
			IsRunParallel:     false,
		},
	}
}

type workspaceDriftConfigTests struct {
	org, project, workspace, pipeline string
	// runEvery and updatedRunEvery are unique prime numbers per test (unit: minutes).
	// Using distinct primes means any stale GET value can be fingerprinted to the exact
	// prior test that wrote it, ruling out test-design issues as the cause of failures.
	runEvery        int32
	updatedRunEvery int32
	runEveryUnit    string
}

func (p workspaceDriftConfigTests) serializerKey() string {
	return fmt.Sprintf("%s/%s/%s/opentofu", p.org, p.project, p.workspace)
}

func (p workspaceDriftConfigTests) requiredEnv() []string {
	return []string{
		"HARNESS_ACCOUNT_ID",
		"HARNESS_PLATFORM_API_KEY",
		"IACM_WORKSPACE_DRIFT_ORG_ID",
		"IACM_WORKSPACE_DRIFT_PROJECT_ID",
		"IACM_WORKSPACE_DRIFT_WORKSPACE_ID",
		"IACM_WORKSPACE_DRIFT_PIPELINE_ID",
	}
}

func (p workspaceDriftConfigTests) GetCreateResourceTfConfig() string {
	return fmt.Sprintf(`resource "harness_platform_iacm_workspace_drift_config" "test" {
  org_id           = %q
  project_id       = %q
  workspace_id     = %q
  provisioner_type = "opentofu"
  pipeline_id      = %q
  run_every        = %d
  run_every_unit   = %q
}`, p.org, p.project, p.workspace, p.pipeline, p.runEvery, p.runEveryUnit)
}

func (p workspaceDriftConfigTests) GetUpdatedTerraformConfig() string {
	return fmt.Sprintf(`resource "harness_platform_iacm_workspace_drift_config" "test" {
  org_id           = %q
  project_id       = %q
  workspace_id     = %q
  provisioner_type = "opentofu"
  pipeline_id      = %q
  run_every        = %d
  run_every_unit   = %q
}`, p.org, p.project, p.workspace, p.pipeline, p.updatedRunEvery, p.runEveryUnit)
}

func (p workspaceDriftConfigTests) getInvalidConfig() string {
	return fmt.Sprintf(`resource "harness_platform_iacm_workspace_drift_config" "test" {
  org_id           = %q
  project_id       = %q
  workspace_id     = %q
  provisioner_type = "opentofu"
  run_every        = %d
  run_every_unit   = %q
}`, p.org, p.project, p.workspace, p.runEvery, p.runEveryUnit)
}

func (p workspaceDriftConfigTests) getImportID() string {
	return fmt.Sprintf("%s/%s/%s/opentofu", p.org, p.project, p.workspace)
}

func (p workspaceDriftConfigTests) getNonExistentImportID() string {
	return "nonexistent-org-iacm/nonexistent-project-iacm/nonexistent-workspace-iacm/opentofu"
}

// verifyCreated checks Terraform state and confirms the remote resource matches via direct API call.
func (p workspaceDriftConfigTests) verifyCreated(s *terraform.State) error {
	if err := resource.TestCheckResourceAttr(workspaceDriftConfigResourceAddress, "enabled", "true")(s); err != nil {
		return err
	}
	if err := resource.TestCheckResourceAttr(workspaceDriftConfigResourceAddress, "pipeline_id", p.pipeline)(s); err != nil {
		return err
	}
	if err := resource.TestCheckResourceAttr(workspaceDriftConfigResourceAddress, "run_every", fmt.Sprintf("%d", p.runEvery))(s); err != nil {
		return err
	}
	if err := resource.TestCheckResourceAttr(workspaceDriftConfigResourceAddress, "run_every_unit", p.runEveryUnit)(s); err != nil {
		return err
	}
	if err := resource.TestCheckResourceAttr(workspaceDriftConfigResourceAddress, "id", fmt.Sprintf("%s/%s/%s/opentofu", p.org, p.project, p.workspace))(s); err != nil {
		return err
	}
	client, ctx := acctest.TestAccGetPlatformClientWithContext()
	got, resp, err := client.IacmWorkspaceDriftConfigApi.GetIacmWorkspaceDriftConfig(ctx, client.AccountId, p.org, p.project, p.workspace, "opentofu")
	if err != nil {
		return fmt.Errorf("verify create: GET workspace drift config: %w", err)
	}
	if resp == nil || resp.StatusCode != http.StatusOK {
		return fmt.Errorf("verify create: unexpected GET response: %v", resp)
	}
	if !got.Enabled || got.PipelineId != p.pipeline || got.RunEvery != p.runEvery || got.RunEveryUnit != p.runEveryUnit {
		return fmt.Errorf("remote workspace drift config does not match: enabled=%v pipeline=%q every=%d unit=%q (expected every=%d unit=%q)",
			got.Enabled, got.PipelineId, got.RunEvery, got.RunEveryUnit, p.runEvery, p.runEveryUnit)
	}
	return nil
}

// verifyUpdated checks that run_every=updatedRunEvery is reflected in state and confirmed via direct API call.
func (p workspaceDriftConfigTests) verifyUpdated(s *terraform.State) error {
	if err := resource.TestCheckResourceAttr(workspaceDriftConfigResourceAddress, "run_every", fmt.Sprintf("%d", p.updatedRunEvery))(s); err != nil {
		return err
	}
	if err := resource.TestCheckResourceAttr(workspaceDriftConfigResourceAddress, "run_every_unit", p.runEveryUnit)(s); err != nil {
		return err
	}
	client, ctx := acctest.TestAccGetPlatformClientWithContext()
	got, resp, err := client.IacmWorkspaceDriftConfigApi.GetIacmWorkspaceDriftConfig(ctx, client.AccountId, p.org, p.project, p.workspace, "opentofu")
	if err != nil {
		return fmt.Errorf("verify update: GET workspace drift config: %w", err)
	}
	if resp == nil || resp.StatusCode != http.StatusOK {
		return fmt.Errorf("verify update: unexpected GET response: %v", resp)
	}
	if !got.Enabled || got.PipelineId != p.pipeline || got.RunEvery != p.updatedRunEvery || got.RunEveryUnit != p.runEveryUnit {
		return fmt.Errorf("remote workspace drift config does not match after update: enabled=%v pipeline=%q every=%d unit=%q (expected every=%d unit=%q)",
			got.Enabled, got.PipelineId, got.RunEvery, got.RunEveryUnit, p.updatedRunEvery, p.runEveryUnit)
	}
	return nil
}

// verifyDeleted confirms the remote resource is disabled or absent after Terraform destroy.
func (p workspaceDriftConfigTests) verifyDeleted(_ *terraform.State) error {
	client, ctx := acctest.TestAccGetPlatformClientWithContext()
	got, resp, err := client.IacmWorkspaceDriftConfigApi.GetIacmWorkspaceDriftConfig(ctx, client.AccountId, p.org, p.project, p.workspace, "opentofu")
	if resp != nil && resp.StatusCode == http.StatusNotFound {
		return nil
	}
	if err != nil {
		return fmt.Errorf("verify delete: GET workspace drift config: %w", err)
	}
	if got.Enabled {
		return fmt.Errorf("workspace drift config remains enabled after Terraform destroy")
	}
	return nil
}

// externalDeleteAction disables the drift config outside Terraform and waits for the change to converge.
func (p workspaceDriftConfigTests) externalDeleteAction() error {
	client, ctx := newWorkspaceDriftConfigClient()
	log.Printf("[TEST] externalDeleteAction: sending DELETE for org=%q project=%q workspace=%q", p.org, p.project, p.workspace)
	httpResp, err := client.IacmWorkspaceDriftConfigApi.DisableIacmWorkspaceDriftConfig(ctx, client.AccountId, p.org, p.project, p.workspace, "opentofu")
	if err != nil {
		return fmt.Errorf("external delete: %w", err)
	}
	log.Printf("[TEST] externalDeleteAction: DELETE status=%d", httpResp.StatusCode)
	httpResp.Body.Close()
	deadline := time.Now().Add(5 * time.Minute)
	for time.Now().Before(deadline) {
		got, resp, err := client.IacmWorkspaceDriftConfigApi.GetIacmWorkspaceDriftConfig(ctx, client.AccountId, p.org, p.project, p.workspace, "opentofu")
		statusCode := 0
		if resp != nil {
			statusCode = resp.StatusCode
		}
		log.Printf("[TEST] externalDeleteAction poll: status=%d err=%v enabled=%v", statusCode, err, got.Enabled)
		if (resp != nil && resp.StatusCode == http.StatusNotFound) || (err == nil && !got.Enabled) {
			log.Printf("[TEST] externalDeleteAction: confirmed disabled")
			return nil
		}
		time.Sleep(9 * time.Second)
	}
	return fmt.Errorf("external delete: disabled state not confirmed within timeout")
}

// externalUpdateAction changes run_every to p.updatedRunEvery outside Terraform and waits for convergence.
func (p workspaceDriftConfigTests) externalUpdateAction() error {
	client, ctx := newWorkspaceDriftConfigClient()
	log.Printf("[TEST] externalUpdateAction: sending PUT run_every=%d unit=%q for org=%q project=%q workspace=%q", p.updatedRunEvery, p.runEveryUnit, p.org, p.project, p.workspace)
	httpResp, err := client.IacmWorkspaceDriftConfigApi.UpsertIacmWorkspaceDriftConfig(ctx, client.AccountId, p.org, p.project, p.workspace,
		nextgen.IacmWorkspaceDriftConfigRequestBody{
			PipelineId:   p.pipeline,
			Provisioner:  "opentofu",
			RunEvery:     p.updatedRunEvery,
			RunEveryUnit: p.runEveryUnit,
		},
	)
	if err != nil {
		return fmt.Errorf("external update: %w", err)
	}
	log.Printf("[TEST] externalUpdateAction: PUT status=%d", httpResp.StatusCode)
	httpResp.Body.Close()
	deadline := time.Now().Add(5 * time.Minute)
	for time.Now().Before(deadline) {
		got, resp, err := client.IacmWorkspaceDriftConfigApi.GetIacmWorkspaceDriftConfig(ctx, client.AccountId, p.org, p.project, p.workspace, "opentofu")
		statusCode := 0
		if resp != nil {
			statusCode = resp.StatusCode
		}
		log.Printf("[TEST] externalUpdateAction poll: status=%d err=%v enabled=%v run_every=%d unit=%q", statusCode, err, got.Enabled, got.RunEvery, got.RunEveryUnit)
		if err == nil && got.RunEvery == p.updatedRunEvery {
			log.Printf("[TEST] externalUpdateAction: confirmed run_every=%d", p.updatedRunEvery)
			return nil
		}
		time.Sleep(9 * time.Second)
	}
	return fmt.Errorf("external update: run_every=%d not reflected within timeout", p.updatedRunEvery)
}

// newWorkspaceDriftConfigClient builds a nextgen API client from environment variables,
// without requiring the Terraform provider to be initialized. Used by preRun/postRun.
func newWorkspaceDriftConfigClient() (*nextgen.APIClient, context.Context) {
	cfg := nextgen.NewConfiguration()
	client := nextgen.NewAPIClient(cfg)
	ctx := context.WithValue(context.Background(), nextgen.ContextAPIKey, nextgen.APIKey{Key: cfg.ApiKey})
	return client, ctx
}

// preRun checks whether the resource exists and deletes it if so, then waits 10s for the backend to settle.
func (p workspaceDriftConfigTests) preRun(t *testing.T) error {
	client, ctx := newWorkspaceDriftConfigClient()
	got, resp, err := client.IacmWorkspaceDriftConfigApi.GetIacmWorkspaceDriftConfig(ctx, client.AccountId, p.org, p.project, p.workspace, "opentofu")
	if err != nil || (resp != nil && resp.StatusCode == http.StatusNotFound) || !got.Enabled {
		t.Logf("[contractv1] preRun: resource absent or already disabled, no cleanup needed")
		return nil
	}
	t.Logf("[contractv1] preRun: resource exists (run_every=%d), disabling before suite", got.RunEvery)
	httpResp, err := client.IacmWorkspaceDriftConfigApi.DisableIacmWorkspaceDriftConfig(ctx, client.AccountId, p.org, p.project, p.workspace, "opentofu")
	if err != nil {
		t.Logf("[contractv1] preRun: disable failed (non-fatal, continuing): %v", err)
		return nil
	}
	httpResp.Body.Close()
	t.Logf("[contractv1] preRun: disabled (status=%d), waiting 10s for backend to settle", httpResp.StatusCode)
	time.Sleep(10 * time.Second)
	return nil
}

// postRun deletes the resource unconditionally after the suite, then waits 10s for the backend to settle.
func (p workspaceDriftConfigTests) postRun(t *testing.T) error {
	client, ctx := newWorkspaceDriftConfigClient()
	t.Logf("[contractv1] postRun: disabling resource to clean up after suite")
	httpResp, err := client.IacmWorkspaceDriftConfigApi.DisableIacmWorkspaceDriftConfig(ctx, client.AccountId, p.org, p.project, p.workspace, "opentofu")
	if err != nil {
		t.Logf("[contractv1] postRun: disable returned error (may already be absent): %v", err)
		return nil
	}
	httpResp.Body.Close()
	t.Logf("[contractv1] postRun: disabled (status=%d), waiting 10s for backend to settle", httpResp.StatusCode)
	time.Sleep(10 * time.Second)
	return nil
}

var workspaceDriftConfigDebugSettings = contractv1.DebugSettings{
	RunOnlyTests: []contractv1.TestName{
		contractv1.TestDeleteResource,
		// contractv1.TestExternalUpdateResource,
	},
}

func TestAccResourceIacmWorkspaceDriftConfig(t *testing.T) {
	p := workspaceDriftConfigTests{
		org:      os.Getenv("IACM_WORKSPACE_DRIFT_ORG_ID"),
		project:  os.Getenv("IACM_WORKSPACE_DRIFT_PROJECT_ID"),
		workspace: os.Getenv("IACM_WORKSPACE_DRIFT_WORKSPACE_ID"),
		pipeline: os.Getenv("IACM_WORKSPACE_DRIFT_PIPELINE_ID"),
	}
	c := contractv1.ResourceAccTestContractV1{
		GetSerializerKey: p.serializerKey,
		GetRequiredEnv:   p.requiredEnv,
		TestCase:         acctest.GetNewResourceTestCase,
		Definitions:      workspaceDriftConfigDefinitions(p),
		PreRun:           p.preRun,
		PostRun:          p.postRun,
	}
	if single := os.Getenv("SINGLE_TEST"); single != "" {
		c.DebugSettings = contractv1.DebugSettings{
			RunOnlyTests: []contractv1.TestName{contractv1.TestName(single)},
		}
	}
	c.Run(t)
}
