package pipeline

import (
	"testing"

	"github.com/antihax/optional"
	"github.com/harness/harness-openapi-go-client/nextgen"
)

func TestReadGitDetailsUsesApiValues(t *testing.T) {
	got := readGitDetails(nextgen.PipelineGetResponseBody{
		GitDetails: &nextgen.GitDetails{
			BranchName:        "master",
			FilePath:          ".harness/pipeline.yaml",
			RepoName:          "gittest",
			CommitId:          "abc123",
			ConnectorRef:      "account.bitbucket_connector",
			IsHarnessCodeRepo: false,
		},
	}, optional.EmptyString(), optional.EmptyString(), optional.EmptyString(), optional.EmptyString())

	if got["branch_name"] != "master" {
		t.Fatalf("branch_name = %v", got["branch_name"])
	}
	if got["file_path"] != ".harness/pipeline.yaml" {
		t.Fatalf("file_path = %v", got["file_path"])
	}
	if got["repo_name"] != "gittest" {
		t.Fatalf("repo_name = %v", got["repo_name"])
	}
	if got["last_commit_id"] != "abc123" {
		t.Fatalf("last_commit_id = %v", got["last_commit_id"])
	}
	if got["store_type"] != "REMOTE" {
		t.Fatalf("store_type = %v", got["store_type"])
	}
	if got["connector_ref"] != "account.bitbucket_connector" {
		t.Fatalf("connector_ref = %v", got["connector_ref"])
	}
	if got["is_harness_code_repo"] != false {
		t.Fatalf("is_harness_code_repo = %v", got["is_harness_code_repo"])
	}
}

func TestReadGitDetailsKeepsConfiguredValues(t *testing.T) {
	got := readGitDetails(nextgen.PipelineGetResponseBody{
		GitDetails: &nextgen.GitDetails{
			ConnectorRef:      "account.from_api",
			IsHarnessCodeRepo: true,
		},
	}, optional.NewString("INLINE"), optional.EmptyString(), optional.EmptyString(), optional.NewString("account.from_config"))

	if got["store_type"] != "INLINE" {
		t.Fatalf("store_type = %v", got["store_type"])
	}
	if got["connector_ref"] != "account.from_config" {
		t.Fatalf("connector_ref = %v", got["connector_ref"])
	}
	if got["is_harness_code_repo"] != true {
		t.Fatalf("is_harness_code_repo = %v", got["is_harness_code_repo"])
	}
}
