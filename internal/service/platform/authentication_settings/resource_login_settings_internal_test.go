package authentication_settings

import (
	"testing"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFlattenLastUpdatedByNil(t *testing.T) {
	assert.Nil(t, flattenLastUpdatedBy(nil))
}

func TestFlattenLastUpdatedByFullObject(t *testing.T) {
	out := flattenLastUpdatedBy(&nextgen.EmbeddedUser{
		Uuid:           "u-1",
		Name:           "Test User",
		Email:          "test@harness.io",
		ExternalUserId: "ext-1",
	})

	require.Len(t, out, 1)
	m := out[0].(map[string]interface{})
	assert.Equal(t, "u-1", m["uuid"])
	assert.Equal(t, "Test User", m["name"])
	assert.Equal(t, "test@harness.io", m["email"])
	assert.Equal(t, "ext-1", m["external_user_id"])
}

func TestFlattenLastUpdatedByPartialObject(t *testing.T) {
	out := flattenLastUpdatedBy(&nextgen.EmbeddedUser{Email: "test@harness.io"})

	require.Len(t, out, 1)
	m := out[0].(map[string]interface{})
	assert.Equal(t, "", m["uuid"])
	assert.Equal(t, "test@harness.io", m["email"])
}
