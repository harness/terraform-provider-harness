package role_assignments_test

import (
	"testing"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/internal/service/platform/role_assignments"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
)

// TestBuildRoleAssignment_NoRoleReference tests that when role_reference is not set,
// the RoleReference field should be nil (not an empty struct)
func TestBuildRoleAssignment_NoRoleReference(t *testing.T) {
	d := schema.TestResourceDataRaw(t, role_assignments.ResourceRoleAssignments().Schema, map[string]interface{}{
		"identifier":                "test_id",
		"role_identifier":           "test_role",
		"resource_group_identifier": "test_rg",
		"disabled":                  false,
		"managed":                   false,
		"principal": []interface{}{
			map[string]interface{}{
				"identifier": "test_principal",
				"type":       "SERVICE_ACCOUNT",
			},
		},
		// role_reference is NOT set
	})

	result := role_assignments.BuildRoleAssignment(d)

	assert.NotNil(t, result, "RoleAssignment should not be nil")
	assert.Equal(t, "test_id", result.Identifier)
	assert.Equal(t, "test_role", result.RoleIdentifier)
	assert.NotNil(t, result.Principal, "Principal should not be nil")
	assert.Nil(t, result.RoleReference, "RoleReference should be nil when not provided")
}

// TestBuildRoleAssignment_WithRoleReference tests that when role_reference is set,
// the RoleReference field should be populated correctly
func TestBuildRoleAssignment_WithRoleReference(t *testing.T) {
	d := schema.TestResourceDataRaw(t, role_assignments.ResourceRoleAssignments().Schema, map[string]interface{}{
		"identifier":                "test_id",
		"role_identifier":           "test_role",
		"resource_group_identifier": "test_rg",
		"disabled":                  false,
		"managed":                   false,
		"principal": []interface{}{
			map[string]interface{}{
				"identifier": "test_principal",
				"type":       "SERVICE_ACCOUNT",
			},
		},
		"role_reference": []interface{}{
			map[string]interface{}{
				"identifier":  "org_role",
				"scope_level": "organization",
			},
		},
	})

	result := role_assignments.BuildRoleAssignment(d)

	assert.NotNil(t, result, "RoleAssignment should not be nil")
	assert.NotNil(t, result.RoleReference, "RoleReference should not be nil when provided")
	assert.Equal(t, "org_role", result.RoleReference.Identifier)
	assert.Equal(t, "organization", result.RoleReference.ScopeLevel)
}

// TestBuildRoleAssignment_EmptyRoleReferenceBlock tests that even if role_reference block
// exists but has no fields set, it should still create the RoleReference
func TestBuildRoleAssignment_EmptyRoleReferenceBlock(t *testing.T) {
	d := schema.TestResourceDataRaw(t, role_assignments.ResourceRoleAssignments().Schema, map[string]interface{}{
		"identifier":                "test_id",
		"role_identifier":           "test_role",
		"resource_group_identifier": "test_rg",
		"principal": []interface{}{
			map[string]interface{}{
				"identifier": "test_principal",
				"type":       "SERVICE_ACCOUNT",
			},
		},
		"role_reference": []interface{}{
			map[string]interface{}{
				// Empty block - no fields set
			},
		},
	})

	result := role_assignments.BuildRoleAssignment(d)

	// Even with empty block, GetOk will detect it exists, so RoleReference will be created
	assert.NotNil(t, result.RoleReference, "RoleReference should be created even for empty block")
	assert.Equal(t, "", result.RoleReference.Identifier)
	assert.Equal(t, "", result.RoleReference.ScopeLevel)
}

// TestReadRoleAssignments_WithRoleReference tests reading when RoleReference is present
func TestReadRoleAssignments_WithRoleReference(t *testing.T) {
	d := schema.TestResourceDataRaw(t, role_assignments.ResourceRoleAssignments().Schema, map[string]interface{}{})

	roleAssignment := &nextgen.RoleAssignment{
		Identifier:              "test_id",
		RoleIdentifier:          "test_role",
		ResourceGroupIdentifier: "test_rg",
		Disabled:                false,
		Managed:                 false,
		Principal: &nextgen.AuthzPrincipal{
			Identifier: "test_principal",
			Type_:      "SERVICE_ACCOUNT",
			ScopeLevel: "project",
		},
		RoleReference: &nextgen.RoleReference{
			Identifier: "org_role",
			ScopeLevel: "organization",
		},
	}

	role_assignments.ReadRoleAssignments(d, roleAssignment)

	assert.Equal(t, "test_id", d.Id())
	assert.Equal(t, "test_role", d.Get("role_identifier"))

	roleRefList := d.Get("role_reference").([]interface{})
	assert.Len(t, roleRefList, 1)
	roleRef := roleRefList[0].(map[string]interface{})
	assert.Equal(t, "org_role", roleRef["identifier"])
	assert.Equal(t, "organization", roleRef["scope_level"])
}

// TestReadRoleAssignments_NilRoleReference tests that reading with nil RoleReference doesn't crash
func TestReadRoleAssignments_NilRoleReference(t *testing.T) {
	d := schema.TestResourceDataRaw(t, role_assignments.ResourceRoleAssignments().Schema, map[string]interface{}{})

	roleAssignment := &nextgen.RoleAssignment{
		Identifier:              "test_id",
		RoleIdentifier:          "test_role",
		ResourceGroupIdentifier: "test_rg",
		Disabled:                false,
		Managed:                 false,
		Principal: &nextgen.AuthzPrincipal{
			Identifier: "test_principal",
			Type_:      "SERVICE_ACCOUNT",
			ScopeLevel: "project",
		},
		RoleReference: nil, // No role reference
	}

	// Should not panic
	assert.NotPanics(t, func() {
		role_assignments.ReadRoleAssignments(d, roleAssignment)
	})

	assert.Equal(t, "test_id", d.Id())
	assert.Equal(t, "test_role", d.Get("role_identifier"))

	// role_reference should not be set in state
	roleRefList := d.Get("role_reference").([]interface{})
	assert.Len(t, roleRefList, 0, "role_reference should be empty when RoleReference is nil")
}
