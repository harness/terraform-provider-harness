package delegatetoken

import (
	"testing"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
)

// TestReadDelegateTokenWithNilCreatedByNgUser tests the fix for DEL-6930
// where default tokens have nil CreatedByNgUser and caused panic
func TestReadDelegateTokenWithNilCreatedByNgUser(t *testing.T) {
	d := schema.TestResourceDataRaw(t, ResourceDelegateToken().Schema, map[string]interface{}{})

	// Simulate a default token with nil CreatedByNgUser
	delegateToken := &nextgen.DelegateTokenDetails{
		Name:             "default_token_test",
		AccountId:        "test_account",
		Status:           "ACTIVE",
		CreatedAt:        1620000000000,
		Value:            "test_token_value",
		CreatedByNgUser:  nil, // This is nil for default tokens
	}

	// This should not panic
	assert.NotPanics(t, func() {
		readDelegateToken(d, delegateToken)
	})

	// Verify the fields are set correctly
	assert.Equal(t, "default_token_test", d.Id())
	assert.Equal(t, "default_token_test", d.Get("name").(string))
	assert.Equal(t, "test_account", d.Get("account_id").(string))
	assert.Equal(t, "ACTIVE", d.Get("token_status").(string))
	assert.Equal(t, int64(1620000000000), int64(d.Get("created_at").(int)))
	assert.Equal(t, "test_token_value", d.Get("value").(string))

	// Verify created_by is set to empty map when CreatedByNgUser is nil
	createdBy := d.Get("created_by").(map[string]interface{})
	assert.NotNil(t, createdBy)
	assert.Equal(t, 0, len(createdBy))
}

// TestReadDelegateTokenWithCreatedByNgUser tests that non-default tokens
// with CreatedByNgUser still work correctly
func TestReadDelegateTokenWithCreatedByNgUser(t *testing.T) {
	d := schema.TestResourceDataRaw(t, ResourceDelegateToken().Schema, map[string]interface{}{})

	// Simulate a user-created token with CreatedByNgUser
	delegateToken := &nextgen.DelegateTokenDetails{
		Name:      "user_token_test",
		AccountId: "test_account",
		Status:    "ACTIVE",
		CreatedAt: 1620000000000,
		Value:     "test_token_value",
		CreatedByNgUser: &nextgen.Principal{
			Type_:     "USER",
			Name:      "John Doe",
			Jwtclaims: map[string]string{"email": "john@example.com"},
		},
	}

	// This should not panic
	assert.NotPanics(t, func() {
		readDelegateToken(d, delegateToken)
	})

	// Verify the fields are set correctly
	assert.Equal(t, "user_token_test", d.Id())
	assert.Equal(t, "user_token_test", d.Get("name").(string))
	assert.Equal(t, "test_account", d.Get("account_id").(string))
	assert.Equal(t, "ACTIVE", d.Get("token_status").(string))

	// Verify created_by is populated correctly
	createdBy := d.Get("created_by").(map[string]interface{})
	assert.NotNil(t, createdBy)
	assert.Equal(t, "USER", createdBy["type"])
	assert.Equal(t, "John Doe", createdBy["name"])
	assert.Equal(t, "john@example.com", createdBy["email"])
}
