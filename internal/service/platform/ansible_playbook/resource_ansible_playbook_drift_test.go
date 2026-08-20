package ansible_playbook

import (
	"context"
	"testing"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// These drive the real SDK plan machinery to prove out-of-band tag drift is
// corrected. See the equivalent file in the ansible_inventory package; the two
// resources carry separate copies of expandTags, so both need pinning.
func TestTagsSchemaIsOptionalNotComputed(t *testing.T) {
	tags := ResourceAnsiblePlaybook().Schema["tags"]
	require.NotNil(t, tags)

	assert.True(t, tags.Optional, "tags must be Optional")
	assert.False(t, tags.Computed,
		"tags must NOT be Computed - a Computed tag set makes an omitted config block mean "+
			"'keep whatever the server has', which would silently permit out-of-band drift")
}

func refreshedState(t *testing.T, resp *nextgen.ShowPlaybookResponse) *terraform.InstanceState {
	t.Helper()

	d := schema.TestResourceDataRaw(t, ResourceAnsiblePlaybook().Schema, map[string]interface{}{})
	require.Nil(t, readPlaybook(d, resp))

	state := d.State()
	require.NotNil(t, state)
	return state
}

func planAndCaptureUpdate(
	t *testing.T,
	state *terraform.InstanceState,
	config map[string]interface{},
) (*terraform.InstanceDiff, *schema.ResourceData) {
	t.Helper()

	r := ResourceAnsiblePlaybook()
	diff, err := r.SimpleDiff(context.Background(), state, terraform.NewResourceConfigRaw(config), nil)
	require.NoError(t, err)

	var captured *schema.ResourceData
	r.UpdateContext = func(_ context.Context, d *schema.ResourceData, _ interface{}) diag.Diagnostics {
		captured = d
		return nil
	}
	_, diags := r.Apply(context.Background(), state, diff, nil)
	require.False(t, diags.HasError(), "apply reported errors: %v", diags)
	require.NotNil(t, captured, "Update was never called, so no drift correction would happen")

	return diff, captured
}

func baseConfig(tags []interface{}) map[string]interface{} {
	cfg := map[string]interface{}{
		"identifier":      "pb",
		"name":            "pb",
		"org_id":          "org",
		"project_id":      "proj",
		"repository_path": "playbooks/site.yml",
	}
	if tags != nil {
		cfg["tags"] = tags
	}
	return cfg
}

func serverPlaybook(tags []string) *nextgen.ShowPlaybookResponse {
	return &nextgen.ShowPlaybookResponse{
		Identifier:     "pb",
		Org:            "org",
		Project:        "proj",
		Name:           "pb",
		RepositoryPath: "playbooks/site.yml",
		Tags:           tags,
	}
}

// Config declares tags, someone edits them in the UI: the next apply must put
// the configured tags back.
func TestTagDrift_EditedOutsideTerraform_IsRestored(t *testing.T) {
	state := refreshedState(t, serverPlaybook([]string{"env:dev", "rogue:tag"}))

	diff, d := planAndCaptureUpdate(t, state, baseConfig([]interface{}{"env:prod"}))

	assert.False(t, diff.Empty(), "drift must produce a non-empty plan")
	assert.True(t, d.HasChange("tags"))
	assert.Equal(t, []string{"env:prod"}, buildUpdatePlaybook(d).Tags,
		"the configured tags must be what gets PUT, restoring the config as authoritative")
}

// No tags block in config, someone adds tags in the UI: the next apply must
// strip them.
func TestTagDrift_AddedOutsideTerraform_WithNoTagsInConfig_IsRemoved(t *testing.T) {
	state := refreshedState(t, serverPlaybook([]string{"added:in-ui"}))
	require.Equal(t, "1", state.Attributes["tags.#"], "refresh must record the out-of-band tag")

	diff, d := planAndCaptureUpdate(t, state, baseConfig(nil))

	assert.False(t, diff.Empty(), "an out-of-band tag with no tags in config must plan a change")
	assert.True(t, d.HasChange("tags"))

	tags := buildUpdatePlaybook(d).Tags
	require.NotNil(t, tags)
	assert.Empty(t, tags, "an empty array is what clears tags on this PUT endpoint")
}

// The case d.GetOk used to swallow: an explicit `tags = []`.
func TestTagDrift_AddedOutsideTerraform_WithEmptyTagsInConfig_IsRemoved(t *testing.T) {
	state := refreshedState(t, serverPlaybook([]string{"added:in-ui"}))

	diff, d := planAndCaptureUpdate(t, state, baseConfig([]interface{}{}))

	assert.False(t, diff.Empty())
	assert.True(t, d.HasChange("tags"))

	tags := buildUpdatePlaybook(d).Tags
	require.NotNil(t, tags)
	assert.Empty(t, tags)
}

// No perpetual diff when the server already agrees with the config.
func TestTagDrift_ServerMatchesConfig_PlansNothing(t *testing.T) {
	state := refreshedState(t, serverPlaybook([]string{"env:prod", "team:iacm"}))

	r := ResourceAnsiblePlaybook()
	diff, err := r.SimpleDiff(
		context.Background(),
		state,
		terraform.NewResourceConfigRaw(baseConfig([]interface{}{"team:iacm", "env:prod"})), // order must not matter
		nil,
	)
	require.NoError(t, err)
	assert.True(t, diff.Empty(), "matching tags must not plan a change, got: %v", diff.Attributes)
}

func TestTagDrift_NoTagsAnywhere_PlansNothing(t *testing.T) {
	state := refreshedState(t, serverPlaybook(nil))

	r := ResourceAnsiblePlaybook()
	diff, err := r.SimpleDiff(
		context.Background(),
		state,
		terraform.NewResourceConfigRaw(baseConfig(nil)),
		nil,
	)
	require.NoError(t, err)
	assert.True(t, diff.Empty(), "expected no plan, got: %v", diff.Attributes)
}
