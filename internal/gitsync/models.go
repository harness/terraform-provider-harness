package gitsync

import (
	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const commit_message = "Managed by the Harness Platform Terraform provider."

func GetGitSyncOptions(d *schema.ResourceData) *nextgen.EntityGitDetailsOptions {
	attr, ok := d.GetOk("git_sync")
	if !ok {
		return &nextgen.EntityGitDetailsOptions{}
	}

	opts := attr.([]interface{})[0].(map[string]interface{})

	return &nextgen.EntityGitDetailsOptions{
		Branch:         optional.NewString(opts["branch"].(string)),
		RepoIdentifier: optional.NewString(opts["repo_id"].(string)),
		RootFolder:     optional.NewString(opts["root_folder"].(string)),
		FilePath:       optional.NewString(opts["file_path"].(string)),
		// CommitMsg:      optional.NewString(opts["commit_message"].(string)),
		// IsNewBranch:    optional.NewBool(opts["is_new_branch"].(bool)),
		// BaseBranch:     optional.NewString(opts["base_branch"].(string)),
	}
}

func SetGitSyncDetails(d *schema.ResourceData, opts *nextgen.EntityGitDetails) {
	if opts.IsEmpty() {
		return
	}

	d.Set("git_sync", []interface{}{
		map[string]interface{}{
			"branch":      opts.Branch,
			"repo_id":     opts.RepoIdentifier,
			"root_folder": opts.RootFolder,
			"file_path":   opts.FilePath,
			// "commit_message": commit_message,
			// "is_new_branch":  false,
			// "base_branch":    "",
		},
	})
}
