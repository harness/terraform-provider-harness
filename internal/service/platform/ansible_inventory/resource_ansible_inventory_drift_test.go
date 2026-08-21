package ansible_inventory

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

// The tests below answer "does out-of-band tag drift get corrected?" by driving
// the real SDK plan machinery rather than by inspecting our own helpers:
//
//	readInventory(server response) -> InstanceState   (the refresh)
//	Resource.SimpleDiff(state, config)                 -> InstanceDiff (the plan)
//	Resource.Apply(state, diff)                        -> the ResourceData Update sees
//
// The tag set is Optional and NOT Computed, which is what makes the config
// authoritative: an omitted tags block means "no tags", not "leave them alone".
func TestTagsSchemaIsOptionalNotComputed(t *testing.T) {
	tags := ResourceAnsibleInventory().Schema["tags"]
	require.NotNil(t, tags)

	assert.True(t, tags.Optional, "tags must be Optional")
	assert.False(t, tags.Computed,
		"tags must NOT be Computed - a Computed tag set makes an omitted config block mean "+
			"'keep whatever the server has', which would silently permit out-of-band drift")
}

// refreshedState runs the read path against a server response and returns the
// state Terraform would hold after a refresh.
func refreshedState(t *testing.T, resp *nextgen.ShowInventoryResponse) *terraform.InstanceState {
	t.Helper()

	d := schema.TestResourceDataRaw(t, ResourceAnsibleInventory().Schema, map[string]interface{}{})
	require.Nil(t, readInventory(d, resp))

	state := d.State()
	require.NotNil(t, state)
	return state
}

// planAndCaptureUpdate diffs config against state and, if the plan is non-empty,
// runs Apply with Update stubbed out so the ResourceData the real Update would
// have received can be inspected.
func planAndCaptureUpdate(
	t *testing.T,
	state *terraform.InstanceState,
	config map[string]interface{},
) (*terraform.InstanceDiff, *schema.ResourceData) {
	t.Helper()

	r := ResourceAnsibleInventory()
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
		"identifier": "inv",
		"name":       "inv",
		"type":       inventoryTypeManual,
		"org_id":     "org",
		"project_id": "proj",
	}
	if tags != nil {
		cfg["tags"] = tags
	}
	return cfg
}

// Scenario 1: config declares tags, someone edits them in the UI. The next apply
// must put the configured tags back.
func TestTagDrift_EditedOutsideTerraform_IsRestored(t *testing.T) {
	state := refreshedState(t, &nextgen.ShowInventoryResponse{
		Identifier: "inv",
		Org:        "org",
		Project:    "proj",
		Name:       "inv",
		Type_:      inventoryTypeManual,
		Tags:       []string{"env:dev", "rogue:tag"}, // what the UI left behind
	})

	diff, d := planAndCaptureUpdate(t, state, baseConfig([]interface{}{"env:prod"}))

	assert.False(t, diff.Empty(), "drift must produce a non-empty plan")
	assert.True(t, d.HasChange("tags"), "the update gate keys off HasChanges(\"tags\")")

	req := buildUpdateInventory(d, nil)
	assert.Equal(t, []string{"env:prod"}, req.Tags,
		"the configured tags must be what gets PUT, restoring the config as authoritative")
}

// Scenario 2: the config has no tags block at all and someone adds tags in the
// UI. The next apply must strip them.
func TestTagDrift_AddedOutsideTerraform_WithNoTagsInConfig_IsRemoved(t *testing.T) {
	state := refreshedState(t, &nextgen.ShowInventoryResponse{
		Identifier: "inv",
		Org:        "org",
		Project:    "proj",
		Name:       "inv",
		Type_:      inventoryTypeManual,
		Tags:       []string{"added:in-ui"},
	})
	require.Equal(t, "1", state.Attributes["tags.#"], "refresh must record the out-of-band tag")

	diff, d := planAndCaptureUpdate(t, state, baseConfig(nil)) // no tags block

	assert.False(t, diff.Empty(), "an out-of-band tag with no tags in config must plan a change")
	assert.True(t, d.HasChange("tags"))

	req := buildUpdateInventory(d, nil)
	require.NotNil(t, req.Tags)
	assert.Empty(t, req.Tags, "an empty array is what clears tags on this PUT endpoint")
}

// Same as scenario 2 but with an explicit `tags = []`, which is the other way a
// caller says "no tags". This is the case d.GetOk used to swallow.
func TestTagDrift_AddedOutsideTerraform_WithEmptyTagsInConfig_IsRemoved(t *testing.T) {
	state := refreshedState(t, &nextgen.ShowInventoryResponse{
		Identifier: "inv",
		Org:        "org",
		Project:    "proj",
		Name:       "inv",
		Type_:      inventoryTypeManual,
		Tags:       []string{"added:in-ui"},
	})

	diff, d := planAndCaptureUpdate(t, state, baseConfig([]interface{}{}))

	assert.False(t, diff.Empty())
	assert.True(t, d.HasChange("tags"))

	req := buildUpdateInventory(d, nil)
	require.NotNil(t, req.Tags)
	assert.Empty(t, req.Tags)
}

// The converse: when the server already agrees with the config there must be no
// perpetual diff. A provider that always planned a tag change would be just as
// broken, so this pins the no-op case.
func TestTagDrift_ServerMatchesConfig_PlansNothing(t *testing.T) {
	state := refreshedState(t, &nextgen.ShowInventoryResponse{
		Identifier: "inv",
		Org:        "org",
		Project:    "proj",
		Name:       "inv",
		Type_:      inventoryTypeManual,
		Tags:       []string{"env:prod", "team:iacm"},
	})

	r := ResourceAnsibleInventory()
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
	state := refreshedState(t, &nextgen.ShowInventoryResponse{
		Identifier: "inv",
		Org:        "org",
		Project:    "proj",
		Name:       "inv",
		Type_:      inventoryTypeManual,
		Tags:       nil,
	})

	r := ResourceAnsibleInventory()
	diff, err := r.SimpleDiff(
		context.Background(),
		state,
		terraform.NewResourceConfigRaw(baseConfig(nil)),
		nil,
	)
	require.NoError(t, err)
	assert.True(t, diff.Empty(), "expected no plan, got: %v", diff.Attributes)
}
