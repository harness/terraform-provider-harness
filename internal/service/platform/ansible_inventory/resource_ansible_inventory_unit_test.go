package ansible_inventory

import (
	"encoding/json"
	"testing"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// IAC-8160: the published spec typed inventory tags as a string, so the provider
// used to marshal its tag set into a JSON map string. The API speaks an array.
func TestBuildCreateInventory_TagsSentAsArray(t *testing.T) {
	d := schema.TestResourceDataRaw(t, ResourceAnsibleInventory().Schema, map[string]interface{}{
		"identifier":  "inv",
		"name":        "inv",
		"description": "hosts for the web tier",
		"type":        inventoryTypeManual,
		"tags":        []interface{}{"env:prod", "team:iacm"},
	})

	req := buildCreateInventory(d)

	assert.ElementsMatch(t, []string{"env:prod", "team:iacm"}, req.Tags)
	assert.Equal(t, "hosts for the web tier", req.Description)

	body, err := json.Marshal(req)
	require.NoError(t, err)
	assert.Contains(t, string(body), `"tags":[`)
}

func TestBuildCreateInventory_NoTagsAreOmitted(t *testing.T) {
	d := schema.TestResourceDataRaw(t, ResourceAnsibleInventory().Schema, map[string]interface{}{
		"identifier": "inv",
		"name":       "inv",
		"type":       inventoryTypeManual,
	})

	body, err := json.Marshal(buildCreateInventory(d))
	require.NoError(t, err)
	assert.NotContains(t, string(body), `"tags"`)
}

// Update is a PUT and the service writes tags unconditionally, so clearing every
// tag has to reach the API as an empty array. d.GetOk would report the empty set
// as unset and the tags would never be removed.
func TestBuildUpdateInventory_ClearedTagsSentAsEmptyArray(t *testing.T) {
	d := schema.TestResourceDataRaw(t, ResourceAnsibleInventory().Schema, map[string]interface{}{
		"identifier": "inv",
		"name":       "inv",
		"type":       inventoryTypeManual,
		"tags":       []interface{}{},
	})

	req := buildUpdateInventory(d, json.RawMessage(`{"groups":{}}`))

	require.NotNil(t, req.Tags)
	assert.Empty(t, req.Tags)

	body, err := json.Marshal(req)
	require.NoError(t, err)
	assert.Contains(t, string(body), `"tags":[]`)
}

// The 400 "invalid inventory data" this ticket fixes: the update PUT writes data
// unconditionally, so the current document must be carried through as a JSON
// object rather than an empty string.
func TestBuildUpdateInventory_CarriesCurrentDataDocument(t *testing.T) {
	d := schema.TestResourceDataRaw(t, ResourceAnsibleInventory().Schema, map[string]interface{}{
		"identifier": "inv",
		"name":       "renamed",
		"type":       inventoryTypeManual,
	})
	current := json.RawMessage(`{"groups":{"web":{"hosts":{"web-1":{}}}}}`)

	body, err := json.Marshal(buildUpdateInventory(d, current))
	require.NoError(t, err)

	assert.Contains(t, string(body), `"data":{"groups":{"web":{"hosts":{"web-1":{}}}}}`)
	assert.NotContains(t, string(body), `"data":""`)
}

func TestBuildUpdateInventory_EmptyDataBecomesEmptyObject(t *testing.T) {
	d := schema.TestResourceDataRaw(t, ResourceAnsibleInventory().Schema, map[string]interface{}{
		"identifier": "inv",
		"name":       "inv",
		"type":       inventoryTypeManual,
	})

	for name, current := range map[string]json.RawMessage{
		"nil":   nil,
		"empty": json.RawMessage(""),
	} {
		t.Run("Given "+name+" data - then an empty JSON object is sent", func(t *testing.T) {
			body, err := json.Marshal(buildUpdateInventory(d, current))
			require.NoError(t, err)
			assert.Contains(t, string(body), `"data":{}`)
		})
	}
}

func TestReadInventory_TagsAndDescriptionSetFromResponse(t *testing.T) {
	d := schema.TestResourceDataRaw(t, ResourceAnsibleInventory().Schema, map[string]interface{}{})

	diags := readInventory(d, &nextgen.ShowInventoryResponse{
		Identifier:  "inv",
		Name:        "inv",
		Type_:       inventoryTypeManual,
		Description: "hosts for the web tier",
		Tags:        []string{"env:prod", "team:iacm"},
	})

	require.Nil(t, diags)
	assert.Equal(t, "hosts for the web tier", d.Get("description").(string))
	assert.ElementsMatch(t,
		[]interface{}{"env:prod", "team:iacm"},
		d.Get("tags").(*schema.Set).List(),
	)
}

// Read used to be gated on a non-empty tag string, which left tags removed
// outside Terraform orphaned in state.
func TestReadInventory_RemovedTagsAreClearedFromState(t *testing.T) {
	d := schema.TestResourceDataRaw(t, ResourceAnsibleInventory().Schema, map[string]interface{}{
		"identifier": "inv",
		"name":       "inv",
		"type":       inventoryTypeManual,
		"tags":       []interface{}{"env:prod"},
	})
	require.Equal(t, 1, d.Get("tags").(*schema.Set).Len())

	diags := readInventory(d, &nextgen.ShowInventoryResponse{
		Identifier: "inv",
		Name:       "inv",
		Type_:      inventoryTypeManual,
		Tags:       nil,
	})

	require.Nil(t, diags)
	assert.Equal(t, 0, d.Get("tags").(*schema.Set).Len())
}
