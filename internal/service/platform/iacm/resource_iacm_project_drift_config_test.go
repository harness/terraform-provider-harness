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

const projectDriftConfigResourceAddress = "harness_platform_iacm_project_drift_config.test"

var projectDriftConfigDefinitions = func(p projectDriftConfigTests) []contractv1.ContractDefinition {
	// Each test gets a unique prime run_every so stale GET values can be fingerprinted
	// to the exact prior test that wrote them, isolating backend consistency issues.
	// minutes unit: valid 1-59  → primes 2,3,5,7,11,13,17,19,23,29,31,37,41,43,47,53,59
	// hours unit:   valid 1-24  → primes 2,3,5,7,11,13,17,19,23 (after minutes are exhausted)
	// A different unit makes cross-test bleed-over unambiguous even if the number matches.
	mins := func(runEvery, updatedRunEvery int32) projectDriftConfigTests {
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
			Description:   "Provisions a new project drift config and verifies Terraform state and remote API match.",
			IsRunParallel: false,
			PreTestDelay:  5 * time.Second,
			PostTestDelay: 5 * time.Second,

			GetTerraformConfig: mins(2, 3).GetCreateResourceTfConfig,
			Verifiers:          contractv1.VerifierDefinition{VerifyCreated: mins(2, 3).verifyCreated},
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

			GetTerraformConfig:        mins(23, 29).GetCreateResourceTfConfig,
			GetUpdatedTerraformConfig: mins(23, 29).GetUpdatedTerraformConfig,
			Verifiers: contractv1.VerifierDefinition{
				VerifyCreated: mins(23, 29).verifyCreated,
				VerifyUpdated: mins(23, 29).verifyUpdated,
				VerifyDeleted: mins(23, 29).verifyDeleted,
			},
		},
		{
			TestName:      contractv1.TestImportResource,
			Supported:     true,
			Description:   "Imports the resource by org/project/provisioner and verifies imported state matches created state.",
			IsRunParallel: false,
			PreTestDelay:  5 * time.Second,
			PostTestDelay: 5 * time.Second,

			GetTerraformConfig: mins(31, 37).GetCreateResourceTfConfig,
			ResourceName:       projectDriftConfigResourceAddress,
			GetImportID:        p.getImportID,
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

			GetTerraformConfig: mins(59, 22).GetCreateResourceTfConfig,
			ExternalAction:     mins(59, 22).externalUpdateAction,
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
			Description:   "Imports a syntactically valid but nonexistent org/project/provisioner ID and verifies an error is returned.",
			IsRunParallel: false,
			PreTestDelay:  5 * time.Second,
			PostTestDelay: 5 * time.Second,

			GetTerraformConfig: mins(28, 30).GetCreateResourceTfConfig,
			ResourceName:       projectDriftConfigResourceAddress,
			GetImportID:        p.getNonExistentImportID,
			ExpectError:        regexp.MustCompile(`(?i)(no resource|not found|provider produced|empty|cannot import)`),
		},

		// Unsupported tests — skipped with explicit reasons.
		{
			TestName:          contractv1.TestForceNewResource,
			Supported:         false,
			UnsupportedReason: "org_id, project_id, and provisioner_type are ForceNew; replacement requires a second distinct fixture not available in this test environment",
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
			UnsupportedReason: "async 202 polling is exercised by CreateResourceTest",
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

type projectDriftConfigTests struct {
	org, project, pipeline string
	// runEvery and updatedRunEvery are unique prime numbers per test (unit: minutes).
	// Using distinct primes means any stale GET value can be fingerprinted to the exact
	// prior test that wrote it, ruling out test-design issues as the cause of failures.
	runEvery        int32
	updatedRunEvery int32
	runEveryUnit    string
}

func (p projectDriftConfigTests) serializerKey() string {
	return fmt.Sprintf("%s/%s/opentofu", p.org, p.project)
}

func (p projectDriftConfigTests) requiredEnv() []string {
	return []string{
		"HARNESS_ACCOUNT_ID",
		"HARNESS_PLATFORM_API_KEY",
		"IACM_DRIFT_ORG_ID",
		"IACM_DRIFT_PROJECT_ID",
		"IACM_DRIFT_PIPELINE_ID",
	}
}

func (p projectDriftConfigTests) GetCreateResourceTfConfig() string {
	return fmt.Sprintf(`resource "harness_platform_iacm_project_drift_config" "test" {
  org_id           = %q
  project_id       = %q
  provisioner_type = "opentofu"
  pipeline_id      = %q
  run_every        = %d
  run_every_unit   = %q
}`, p.org, p.project, p.pipeline, p.runEvery, p.runEveryUnit)
}

func (p projectDriftConfigTests) GetUpdatedTerraformConfig() string {
	return fmt.Sprintf(`resource "harness_platform_iacm_project_drift_config" "test" {
  org_id           = %q
  project_id       = %q
  provisioner_type = "opentofu"
  pipeline_id      = %q
  run_every        = %d
  run_every_unit   = %q
}`, p.org, p.project, p.pipeline, p.updatedRunEvery, p.runEveryUnit)
}

func (p projectDriftConfigTests) getInvalidConfig() string {
	return fmt.Sprintf(`resource "harness_platform_iacm_project_drift_config" "test" {
  org_id           = %q
  project_id       = %q
  provisioner_type = "opentofu"
  run_every        = %d
  run_every_unit   = %q
}`, p.org, p.project, p.runEvery, p.runEveryUnit)
}

func (p projectDriftConfigTests) getImportID() string {
	return fmt.Sprintf("%s/%s/opentofu", p.org, p.project)
}

func (p projectDriftConfigTests) getNonExistentImportID() string {
	return "nonexistent-org-iacm/nonexistent-project-iacm/opentofu"
}

// verifyCreated checks Terraform state and confirms the remote resource matches via direct API call.
func (p projectDriftConfigTests) verifyCreated(s *terraform.State) error {
	if err := resource.TestCheckResourceAttr(projectDriftConfigResourceAddress, "enabled", "true")(s); err != nil {
		return err
	}
	if err := resource.TestCheckResourceAttr(projectDriftConfigResourceAddress, "pipeline_id", p.pipeline)(s); err != nil {
		return err
	}
	if err := resource.TestCheckResourceAttr(projectDriftConfigResourceAddress, "run_every", fmt.Sprintf("%d", p.runEvery))(s); err != nil {
		return err
	}
	if err := resource.TestCheckResourceAttr(projectDriftConfigResourceAddress, "run_every_unit", p.runEveryUnit)(s); err != nil {
		return err
	}
	if err := resource.TestCheckResourceAttr(projectDriftConfigResourceAddress, "id", fmt.Sprintf("%s/%s/opentofu", p.org, p.project))(s); err != nil {
		return err
	}
	client, ctx := acctest.TestAccGetPlatformClientWithContext()
	got, resp, err := client.IacmProjectDriftConfigApi.GetIacmProjectDriftConfig(ctx, client.AccountId, p.org, p.project, "opentofu")
	if err != nil {
		return fmt.Errorf("verify create: GET project drift config: %w", err)
	}
	if resp == nil || resp.StatusCode != http.StatusOK {
		return fmt.Errorf("verify create: unexpected GET response: %v", resp)
	}
	if !got.Enabled || got.PipelineId != p.pipeline || got.RunEvery != p.runEvery || got.RunEveryUnit != p.runEveryUnit {
		return fmt.Errorf("remote project drift config does not match: enabled=%v pipeline=%q every=%d unit=%q (expected every=%d unit=%q)",
			got.Enabled, got.PipelineId, got.RunEvery, got.RunEveryUnit, p.runEvery, p.runEveryUnit)
	}
	return nil
}

// verifyUpdated checks that run_every=2 is reflected in state and confirmed via direct API call.
func (p projectDriftConfigTests) verifyUpdated(s *terraform.State) error {
	if err := resource.TestCheckResourceAttr(projectDriftConfigResourceAddress, "run_every", fmt.Sprintf("%d", p.updatedRunEvery))(s); err != nil {
		return err
	}
	if err := resource.TestCheckResourceAttr(projectDriftConfigResourceAddress, "run_every_unit", p.runEveryUnit)(s); err != nil {
		return err
	}
	client, ctx := acctest.TestAccGetPlatformClientWithContext()
	got, resp, err := client.IacmProjectDriftConfigApi.GetIacmProjectDriftConfig(ctx, client.AccountId, p.org, p.project, "opentofu")
	if err != nil {
		return fmt.Errorf("verify update: GET project drift config: %w", err)
	}
	if resp == nil || resp.StatusCode != http.StatusOK {
		return fmt.Errorf("verify update: unexpected GET response: %v", resp)
	}
	if !got.Enabled || got.PipelineId != p.pipeline || got.RunEvery != p.updatedRunEvery || got.RunEveryUnit != p.runEveryUnit {
		return fmt.Errorf("remote project drift config does not match after update: enabled=%v pipeline=%q every=%d unit=%q (expected every=%d unit=%q)",
			got.Enabled, got.PipelineId, got.RunEvery, got.RunEveryUnit, p.updatedRunEvery, p.runEveryUnit)
	}
	return nil
}

// verifyDeleted confirms the remote resource is disabled or absent after Terraform destroy.
func (p projectDriftConfigTests) verifyDeleted(_ *terraform.State) error {
	client, ctx := acctest.TestAccGetPlatformClientWithContext()
	got, resp, err := client.IacmProjectDriftConfigApi.GetIacmProjectDriftConfig(ctx, client.AccountId, p.org, p.project, "opentofu")
	if resp != nil && resp.StatusCode == http.StatusNotFound {
		return nil
	}
	if err != nil {
		return fmt.Errorf("verify delete: GET project drift config: %w", err)
	}
	if got.Enabled {
		return fmt.Errorf("project drift config remains enabled after Terraform destroy")
	}
	return nil
}

// externalDeleteAction disables the drift config outside Terraform and waits for the change to converge.
// Used by ExternalDeleteResource to simulate out-of-band deletion so Terraform detects drift.
func (p projectDriftConfigTests) externalDeleteAction() error {
	client, ctx := acctest.TestAccGetPlatformClientWithContext()
	log.Printf("[TEST] externalDeleteAction: sending DELETE for org=%q project=%q", p.org, p.project)
	httpResp, err := client.IacmProjectDriftConfigApi.DisableIacmProjectDriftConfig(ctx, client.AccountId, p.org, p.project, "opentofu")
	if err != nil {
		return fmt.Errorf("external delete: %w", err)
	}
	log.Printf("[TEST] externalDeleteAction: DELETE status=%d", httpResp.StatusCode)
	httpResp.Body.Close()
	deadline := time.Now().Add(5 * time.Minute)
	for time.Now().Before(deadline) {
		got, resp, err := client.IacmProjectDriftConfigApi.GetIacmProjectDriftConfig(ctx, client.AccountId, p.org, p.project, "opentofu")
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
// Used by ExternalUpdateResource to verify Terraform detects the drift and plans a corrective change.
func (p projectDriftConfigTests) externalUpdateAction() error {
	client, ctx := acctest.TestAccGetPlatformClientWithContext()
	log.Printf("[TEST] externalUpdateAction: sending PUT run_every=%d unit=%q for org=%q project=%q", p.updatedRunEvery, p.runEveryUnit, p.org, p.project)
	httpResp, err := client.IacmProjectDriftConfigApi.UpsertIacmProjectDriftConfig(ctx, client.AccountId, p.org, p.project,
		nextgen.IacmProjectDriftConfigRequestBody{
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
		got, resp, err := client.IacmProjectDriftConfigApi.GetIacmProjectDriftConfig(ctx, client.AccountId, p.org, p.project, "opentofu")
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

// newDriftConfigClient builds a nextgen API client directly from environment variables,
// without requiring the Terraform provider to be initialized. Used by preRun/postRun.
func newDriftConfigClient() (*nextgen.APIClient, context.Context) {
	cfg := nextgen.NewConfiguration()
	client := nextgen.NewAPIClient(cfg)
	ctx := context.WithValue(context.Background(), nextgen.ContextAPIKey, nextgen.APIKey{Key: cfg.ApiKey})
	return client, ctx
}

// preRun checks whether the resource exists and deletes it if so, then waits 10s for the backend to settle.
func (p projectDriftConfigTests) preRun(t *testing.T) error {
	client, ctx := newDriftConfigClient()
	got, resp, err := client.IacmProjectDriftConfigApi.GetIacmProjectDriftConfig(ctx, client.AccountId, p.org, p.project, "opentofu")
	if err != nil || (resp != nil && resp.StatusCode == http.StatusNotFound) || !got.Enabled {
		t.Logf("[contractv1] preRun: resource absent or already disabled, no cleanup needed")
		return nil
	}
	t.Logf("[contractv1] preRun: resource exists (run_every=%d), disabling before suite", got.RunEvery)
	httpResp, err := client.IacmProjectDriftConfigApi.DisableIacmProjectDriftConfig(ctx, client.AccountId, p.org, p.project, "opentofu")
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
func (p projectDriftConfigTests) postRun(t *testing.T) error {
	client, ctx := newDriftConfigClient()
	t.Logf("[contractv1] postRun: disabling resource to clean up after suite")
	httpResp, err := client.IacmProjectDriftConfigApi.DisableIacmProjectDriftConfig(ctx, client.AccountId, p.org, p.project, "opentofu")
	if err != nil {
		t.Logf("[contractv1] postRun: disable returned error (may already be absent): %v", err)
		return nil
	}
	httpResp.Body.Close()
	t.Logf("[contractv1] postRun: disabled (status=%d), waiting 10s for backend to settle", httpResp.StatusCode)
	time.Sleep(10 * time.Second)
	return nil
}

var projectDriftConfigDebugSettings = contractv1.DebugSettings{
	RunOnlyTests: []contractv1.TestName{
		contractv1.TestDeleteResource,
		// contractv1.TestExternalUpdateResource,
	},
}

func TestAccResourceIacmProjectDriftConfig(t *testing.T) {
	p := projectDriftConfigTests{
		org:      os.Getenv("IACM_DRIFT_ORG_ID"),
		project:  os.Getenv("IACM_DRIFT_PROJECT_ID"),
		pipeline: os.Getenv("IACM_DRIFT_PIPELINE_ID"),
	}
	c := contractv1.ResourceAccTestContractV1{
		GetSerializerKey: p.serializerKey,
		GetRequiredEnv:   p.requiredEnv,
		TestCase:         acctest.GetNewResourceTestCase,
		Definitions:      projectDriftConfigDefinitions(p),
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
