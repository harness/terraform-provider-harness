package template

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/require"
)

func TestValidateGitImportDetailsForTemplates(t *testing.T) {
	resourceSchema := ResourceTemplate().Schema

	t.Run("missing git_import_details block", func(t *testing.T) {
		d := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{
			"import_from_git": true,
			"version":         "v1",
			"identifier":      "tmpl",
			"name":            "tmpl",
		})

		diags := validateGitImportDetailsForTemplates(d)
		require.True(t, diags.HasError())
		require.Contains(t, diags[0].Summary, "git_import_details")
		require.Contains(t, diags[0].Summary, "git_details is ignored")
	})

	t.Run("only is_force_import set", func(t *testing.T) {
		d := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{
			"import_from_git": true,
			"version":         "v1",
			"identifier":      "tmpl",
			"name":            "tmpl",
			"git_import_details": []interface{}{
				map[string]interface{}{
					"is_force_import": true,
				},
			},
		})

		diags := validateGitImportDetailsForTemplates(d)
		require.True(t, diags.HasError())
		require.Contains(t, diags[0].Summary, "connector_ref")
		require.Contains(t, diags[0].Summary, "repo_name")
		require.Contains(t, diags[0].Summary, "branch_name")
		require.Contains(t, diags[0].Summary, "file_path")
	})

	t.Run("complete git_import_details", func(t *testing.T) {
		d := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{
			"import_from_git": true,
			"version":         "v1",
			"identifier":      "tmpl",
			"name":            "tmpl",
			"git_import_details": []interface{}{
				map[string]interface{}{
					"connector_ref":   "account.git",
					"repo_name":       "templates",
					"branch_name":     "main",
					"file_path":       ".harness/tmpl.yaml",
					"is_force_import": true,
				},
			},
		})

		diags := validateGitImportDetailsForTemplates(d)
		require.False(t, diags.HasError())
	})
}
