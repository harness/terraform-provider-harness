package ansible_playbook

import (
	"encoding/json"
	"testing"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// IAC-8160: the published spec typed playbook tags as a string, so the provider
// used to marshal its tag set into a JSON map string. The API speaks an array.
func TestBuildCreatePlaybook_TagsSentAsArray(t *testing.T) {
	d := schema.TestResourceDataRaw(t, ResourceAnsiblePlaybook().Schema, map[string]interface{}{
		"identifier":      "pb",
		"name":            "pb",
		"repository_path": "playbooks/site.yml",
		"tags":            []interface{}{"env:prod", "team:iacm"},
	})

	req := buildCreatePlaybook(d)

	assert.ElementsMatch(t, []string{"env:prod", "team:iacm"}, req.Tags)

	body, err := json.Marshal(req)
	require.NoError(t, err)
	assert.Contains(t, string(body), `"tags":[`)
}

func TestBuildCreatePlaybook_NoTagsAreOmitted(t *testing.T) {
	d := schema.TestResourceDataRaw(t, ResourceAnsiblePlaybook().Schema, map[string]interface{}{
		"identifier":      "pb",
		"name":            "pb",
		"repository_path": "playbooks/site.yml",
	})

	body, err := json.Marshal(buildCreatePlaybook(d))
	require.NoError(t, err)
	assert.NotContains(t, string(body), `"tags"`)
}

// Update is a PUT and the service writes tags unconditionally, so clearing every
// tag has to reach the API as an empty array. d.GetOk would report the empty set
// as unset and the tags would never be removed.
func TestBuildUpdatePlaybook_ClearedTagsSentAsEmptyArray(t *testing.T) {
	d := schema.TestResourceDataRaw(t, ResourceAnsiblePlaybook().Schema, map[string]interface{}{
		"identifier":      "pb",
		"name":            "pb",
		"repository_path": "playbooks/site.yml",
		"tags":            []interface{}{},
	})

	req := buildUpdatePlaybook(d)

	require.NotNil(t, req.Tags)
	assert.Empty(t, req.Tags)

	body, err := json.Marshal(req)
	require.NoError(t, err)
	assert.Contains(t, string(body), `"tags":[]`)
}

func TestReadPlaybook_TagsSetFromResponse(t *testing.T) {
	d := schema.TestResourceDataRaw(t, ResourceAnsiblePlaybook().Schema, map[string]interface{}{})

	diags := readPlaybook(d, &nextgen.ShowPlaybookResponse{
		Identifier: "pb",
		Name:       "pb",
		Tags:       []string{"env:prod", "team:iacm"},
	})

	require.Nil(t, diags)
	assert.ElementsMatch(t,
		[]interface{}{"env:prod", "team:iacm"},
		d.Get("tags").(*schema.Set).List(),
	)
}

// Read used to be gated on a non-empty tag string, which left tags removed
// outside Terraform orphaned in state.
func TestReadPlaybook_RemovedTagsAreClearedFromState(t *testing.T) {
	d := schema.TestResourceDataRaw(t, ResourceAnsiblePlaybook().Schema, map[string]interface{}{
		"identifier":      "pb",
		"name":            "pb",
		"repository_path": "playbooks/site.yml",
		"tags":            []interface{}{"env:prod"},
	})
	require.Equal(t, 1, d.Get("tags").(*schema.Set).Len())

	diags := readPlaybook(d, &nextgen.ShowPlaybookResponse{
		Identifier: "pb",
		Name:       "pb",
		Tags:       nil,
	})

	require.Nil(t, diags)
	assert.Equal(t, 0, d.Get("tags").(*schema.Set).Len())
}
