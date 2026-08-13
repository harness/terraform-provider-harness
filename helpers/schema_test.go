package helpers

import (
	"testing"

	"github.com/antihax/optional"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBuildFieldForInt32_Int(t *testing.T) {

	var value = 14
	expected := optional.NewInt32(int32(value))

	resource := createTestResourceForBuildField()
	data := map[string]interface{}{
		"field_int": value,
	}

	d := schema.TestResourceDataRaw(t, resource.Schema, data)

	assert.Equal(t, expected, BuildFieldInt32(d, "field_int"))
}

func TestBuildFieldForInt32_Missing(t *testing.T) {

	expected := optional.EmptyInt32()

	resource := createTestResourceForBuildField()
	data := map[string]interface{}{}

	d := schema.TestResourceDataRaw(t, resource.Schema, data)

	assert.Equal(t, expected, BuildFieldInt32(d, "field_int"))
}

// To run: go test -vet=off ./helpers/... -run TestMultiLevelTemplateImporter -v
func TestMultiLevelTemplateImporter(t *testing.T) {
	templateResource := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"identifier": {Type: schema.TypeString, Optional: true},
			"org_id":     {Type: schema.TypeString, Optional: true},
			"project_id": {Type: schema.TypeString, Optional: true},
			"version":    {Type: schema.TypeString, Optional: true},
		},
	}

	tests := []struct {
		name        string
		id          string
		wantID      string
		wantOrg     string
		wantProject string
		wantVersion string
		wantErr     bool
	}{
		{
			name:   "account level stable",
			id:     "mytmpl",
			wantID: "mytmpl",
		},
		{
			name:        "account level specific version",
			id:          "mytmpl/versions/v1",
			wantID:      "mytmpl",
			wantVersion: "v1",
		},
		{
			name:    "org level stable",
			id:      "myorg/mytmpl",
			wantID:  "mytmpl",
			wantOrg: "myorg",
		},
		{
			name:        "org level specific version",
			id:          "myorg/mytmpl/versions/v1",
			wantID:      "mytmpl",
			wantOrg:     "myorg",
			wantVersion: "v1",
		},
		{
			name:        "project level stable",
			id:          "myorg/myproject/mytmpl",
			wantID:      "mytmpl",
			wantOrg:     "myorg",
			wantProject: "myproject",
		},
		{
			name:        "project level specific version",
			id:          "myorg/myproject/mytmpl/versions/v1",
			wantID:      "mytmpl",
			wantOrg:     "myorg",
			wantProject: "myproject",
			wantVersion: "v1",
		},
		{
			name:    "invalid id",
			id:      "a/b/c/d",
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			d := schema.TestResourceDataRaw(t, templateResource.Schema, map[string]interface{}{})
			d.SetId(tc.id)

			results, err := MultiLevelTemplateImporter.State(d, nil)
			if tc.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.Len(t, results, 1)
			got := results[0]

			assert.Equal(t, tc.wantID, got.Id())
			assert.Equal(t, tc.wantID, got.Get("identifier"))
			assert.Equal(t, tc.wantOrg, got.Get("org_id"))
			assert.Equal(t, tc.wantProject, got.Get("project_id"))
			assert.Equal(t, tc.wantVersion, got.Get("version"))
		})
	}
}

func createTestResourceForBuildField() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"field_int": {
				Type: schema.TypeInt,
			},
		},
	}
}
