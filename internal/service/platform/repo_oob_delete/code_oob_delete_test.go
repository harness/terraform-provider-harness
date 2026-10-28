// Package repo_oob_delete_test verifies, without a Harness account, how the Harness Code
// resources (harness_platform_repo, harness_platform_repo_webhook and
// harness_platform_repo_rule_branch) behave when the underlying object was deleted
// outside Terraform (CODE-6342).
//
// The tests run the real provider code. The Terraform-flow tests also drive a real
// Terraform CLI (downloaded by the plugin testing framework), while the Harness Code
// API is replaced by acctest.FakeCodeAPI.
package repo_oob_delete_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/harness/harness-go-sdk/harness/code"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/harness/terraform-provider-harness/internal/service/platform/repo"
	"github.com/harness/terraform-provider-harness/internal/service/platform/repo_rule_branch"
	"github.com/harness/terraform-provider-harness/internal/service/platform/repo_webhook"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const (
	repoID    = "proof_repo"
	webhookID = "proof_hook"
	ruleID    = "proof_rule"

	repoAddr    = "harness_platform_repo.test"
	webhookAddr = "harness_platform_repo_webhook.test"
	ruleAddr    = "harness_platform_repo_rule_branch.test"

	repoPath    = "repos/" + repoID
	webhookPath = "repos/" + repoID + "/webhooks/" + webhookID
	rulePath    = "repos/" + repoID + "/rules/" + ruleID
)

const config = `
resource "harness_platform_repo" "test" {
	identifier     = "` + repoID + `"
	default_branch = "main"
	description    = "proof"
	readme         = true
}

resource "harness_platform_repo_webhook" "test" {
	identifier      = "` + webhookID + `"
	repo_identifier = harness_platform_repo.test.identifier
	description     = "proof"
	url             = "http://harness.io"
	enabled         = true
	insecure        = true
	triggers        = ["branch_deleted"]
}

resource "harness_platform_repo_rule_branch" "test" {
	identifier      = "` + ruleID + `"
	repo_identifier = harness_platform_repo.test.identifier
	description     = "proof"
	state           = "active"
	pattern {
		default_branch = true
	}
	policies {
		require_pull_request   = true
		require_status_checks  = ["check-stage1"]
		allow_merge_strategies = ["merge"]
	}
	bypass {}
}
`

// scenario describes one out-of-band deletion and which resources must be dropped from
// Terraform state as a result (the rest must be kept).
type scenario struct {
	name    string
	delete  string
	dropped []string
	kept    []string
}

var scenarios = []scenario{
	{
		// The ticket's case. Deleting a repo takes its webhooks and rules with it.
		name:    "repo_deleted",
		delete:  repoPath,
		dropped: []string{repoAddr, webhookAddr, ruleAddr},
	},
	{
		name:    "webhook_deleted",
		delete:  webhookPath,
		dropped: []string{webhookAddr},
		kept:    []string{repoAddr, ruleAddr},
	},
	{
		name:    "rule_deleted",
		delete:  rulePath,
		dropped: []string{ruleAddr},
		kept:    []string{repoAddr, webhookAddr},
	},
}

func newFake(t *testing.T) *acctest.FakeCodeAPI {
	fake := acctest.NewFakeCodeAPI(t)
	fake.SetProviderEnv(t)
	t.Cleanup(func() {
		t.Log("Code API request log:\n  " + strings.Join(fake.Requests(), "\n  "))
	})
	return fake
}

func applyAll() resource.TestStep {
	return resource.TestStep{
		Config: config,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(repoAddr, "identifier", repoID),
			resource.TestCheckResourceAttr(webhookAddr, "identifier", webhookID),
			resource.TestCheckResourceAttr(ruleAddr, "identifier", ruleID),
		),
	}
}

func inState(addr string, s *terraform.State) bool {
	_, ok := s.RootModule().Resources[addr]
	return ok
}

// TestFake_PlanAndRefresh_DropDeletedObjectsFromState checks that, after an object is
// deleted out-of-band, `terraform plan` and `terraform refresh` succeed and remove
// exactly the deleted objects from state.
func TestFake_PlanAndRefresh_DropDeletedObjectsFromState(t *testing.T) {
	for _, sc := range scenarios {
		t.Run(sc.name, func(t *testing.T) {
			fake := newFake(t)

			resource.UnitTest(t, resource.TestCase{
				ProviderFactories: acctest.ProviderFactories,
				Steps: []resource.TestStep{
					applyAll(),
					{
						// `terraform plan` must not fail with "404 Not Found".
						PreConfig:          func() { fake.DeleteOutOfBand(sc.delete) },
						Config:             config,
						PlanOnly:           true,
						ExpectNonEmptyPlan: true,
					},
					{
						// `terraform refresh` must drop the deleted objects from state.
						RefreshState:       true,
						ExpectNonEmptyPlan: true,
						Check: func(s *terraform.State) error {
							for _, addr := range sc.dropped {
								if inState(addr, s) {
									return fmt.Errorf("%s should have been removed from state after refresh", addr)
								}
							}
							for _, addr := range sc.kept {
								if !inState(addr, s) {
									return fmt.Errorf("%s should still be in state after refresh", addr)
								}
							}
							return nil
						},
					},
					{
						// Whatever is left must still destroy cleanly.
						Config:  config,
						Destroy: true,
					},
				},
			})

			if left := fake.Has(repoPath); left {
				t.Errorf("repo %s still exists in the fake after destroy", repoPath)
			}
		})
	}
}

// TestFake_DestroyRetry_StaleState reproduces the ticket's teardown: the state still
// references objects that were already deleted, and `terraform destroy` runs against it.
func TestFake_DestroyRetry_StaleState(t *testing.T) {
	for _, sc := range scenarios {
		t.Run(sc.name, func(t *testing.T) {
			fake := newFake(t)

			resource.UnitTest(t, resource.TestCase{
				ProviderFactories: acctest.ProviderFactories,
				Steps: []resource.TestStep{
					applyAll(),
					{
						PreConfig: func() { fake.DeleteOutOfBand(sc.delete) },
						Config:    config,
						Destroy:   true,
					},
				},
			})
		})
	}
}

func codeSession(t *testing.T, fake *acctest.FakeCodeAPI) *internal.Session {
	t.Helper()
	cfg := code.NewConfiguration()
	cfg.AccountId = "fake_account"
	cfg.ApiKey = "fake_key"
	cfg.BasePath = fake.Server.URL + "/code/api/v1"
	return &internal.Session{AccountId: cfg.AccountId, CodeClient: code.NewAPIClient(cfg)}
}

// TestFake_Delete_AlreadyDeletedIsSuccess covers `terraform destroy -refresh=false` and
// objects that vanish between refresh and delete: Delete is called for an object that
// no longer exists, and must succeed.
func TestFake_Delete_AlreadyDeletedIsSuccess(t *testing.T) {
	cases := []struct {
		name string
		run  func(ctx context.Context, meta *internal.Session) (hasError bool, summary string)
	}{
		{
			name: "harness_platform_repo",
			run: func(ctx context.Context, meta *internal.Session) (bool, string) {
				r := repo.ResourceRepo()
				d := r.TestResourceData()
				d.SetId(repoID)
				diags := r.DeleteContext(ctx, d, meta)
				return diags.HasError(), fmt.Sprint(diags)
			},
		},
		{
			name: "harness_platform_repo_webhook",
			run: func(ctx context.Context, meta *internal.Session) (bool, string) {
				r := repo_webhook.ResourceRepoWebhook()
				d := r.TestResourceData()
				d.SetId(webhookID)
				_ = d.Set("repo_identifier", repoID)
				diags := r.DeleteContext(ctx, d, meta)
				return diags.HasError(), fmt.Sprint(diags)
			},
		},
		{
			name: "harness_platform_repo_rule_branch",
			run: func(ctx context.Context, meta *internal.Session) (bool, string) {
				r := repo_rule_branch.ResourceRepoBranchRule()
				d := r.TestResourceData()
				d.SetId(ruleID)
				_ = d.Set("repo_identifier", repoID)
				diags := r.DeleteContext(ctx, d, meta)
				return diags.HasError(), fmt.Sprint(diags)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			fake := acctest.NewFakeCodeAPI(t) // empty: every object is "already deleted"
			meta := codeSession(t, fake)

			hasError, summary := tc.run(context.Background(), meta)

			t.Logf("Code API request log: %v", fake.Requests())
			if hasError {
				t.Fatalf("Delete of an already-deleted object must succeed, got: %s", summary)
			}
		})
	}
}
