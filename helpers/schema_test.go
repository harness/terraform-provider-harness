package helpers

import (
	"testing"

	"github.com/antihax/optional"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
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

func createTestResourceForBuildField() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"field_int": {
				Type: schema.TypeInt,
			},
		},
	}
}
