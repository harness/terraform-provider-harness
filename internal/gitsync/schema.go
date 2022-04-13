package gitsync

import (
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func GetGitSyncBranchSchema(flag helpers.SchemaFlagType) *schema.Schema {
	s := &schema.Schema{
		Description: "The git branch to use for the resource.",
		Type:        schema.TypeString,
		Optional:    true,
	}
	helpers.SetSchemaFlagType(s, flag)
	return s
}

func GetGitSyncRepoIdSchema(flag helpers.SchemaFlagType) *schema.Schema {
	s := &schema.Schema{
		Description: "Git sync config Id.",
		Type:        schema.TypeString,
	}
	helpers.SetSchemaFlagType(s, flag)
	return s
}

func GetGitSyncRootFolderSchema(flag helpers.SchemaFlagType) *schema.Schema {
	s := &schema.Schema{
		Description: "The root folder path of the resource.",
		Type:        schema.TypeString,
	}
	helpers.SetSchemaFlagType(s, flag)
	return s
}

func GetGitSyncFilePathSchema(flag helpers.SchemaFlagType) *schema.Schema {
	s := &schema.Schema{
		Description: "The file path of the resource.",
		Type:        schema.TypeString,
	}
	helpers.SetSchemaFlagType(s, flag)
	return s
}

func GetGitSyncCommitMessageSchema() *schema.Schema {
	return &schema.Schema{
		Description: "The commit message to use for the resource.",
		Type:        schema.TypeString,
		Computed:    true,
	}
}

func GetGitSyncIsNewBranchSchema() *schema.Schema {
	return &schema.Schema{
		Description: "Whether to create a new branch.",
		Type:        schema.TypeBool,
		Computed:    true,
	}
}

func GetGitSyncBaseBranchSchema() *schema.Schema {
	return &schema.Schema{
		Description: "The base branch to use for the resource.",
		Type:        schema.TypeString,
		Computed:    true,
	}
}

func GetGitSyncObjectIdSchema() *schema.Schema {
	s := &schema.Schema{
		Description: "The object id of the resource.",
		Type:        schema.TypeString,
		Computed:    true,
	}
	return s
}

// GetGitSyncSchema returns the schema for a git sync resource.
// If readOnly is true, each field is marked as Computed.
// Otherwise each field is marked as Optional.
func SetGitSyncSchema(s map[string]*schema.Schema, readOnly bool) {
	flag := helpers.SchemaFlagTypes.Optional
	maxItems := 1

	if readOnly {
		flag = helpers.SchemaFlagTypes.Computed
		maxItems = 0
	}

	s["git_sync"] = &schema.Schema{
		Type:     schema.TypeList,
		MaxItems: maxItems,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"branch":      GetGitSyncBranchSchema(flag),
				"repo_id":     GetGitSyncRepoIdSchema(flag),
				"root_folder": GetGitSyncRootFolderSchema(flag),
				"file_path":   GetGitSyncFilePathSchema(flag),
				// "commit_message": GetGitSyncCommitMessageSchema(),
				// "is_new_branch":  GetGitSyncIsNewBranchSchema(),
				// "base_branch":    GetGitSyncBaseBranchSchema(),
				"object_id": GetGitSyncObjectIdSchema(),
			},
		},
	}
	helpers.SetSchemaFlagType(s["git_sync"], flag)

}
