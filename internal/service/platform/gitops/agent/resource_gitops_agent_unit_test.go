package agent

import (
	"testing"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
)

func TestAgentTokenSensitive(t *testing.T) {
	assert.True(t, ResourceGitopsAgent().Schema["agent_token"].Sensitive)
	assert.True(t, DataSourceGitopsAgent().Schema["agent_token"].Sensitive)
}

func TestDisasterRecoveryFieldsRequiredWith(t *testing.T) {
	s := ResourceGitopsAgent().Schema
	assert.Equal(t, []string{"disaster_recovery_type"}, s["disaster_recovery_identifier"].RequiredWith)
	assert.Equal(t, []string{"disaster_recovery_identifier"}, s["disaster_recovery_type"].RequiredWith)
}

func TestDisasterRecoveryIdentifierValidateFunc(t *testing.T) {
	validateFn := ResourceGitopsAgent().Schema["disaster_recovery_identifier"].ValidateFunc

	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{name: "plain identifier is valid", value: "dr-east", wantErr: false},
		{name: "contains colon is invalid", value: "dr:east", wantErr: true},
		{name: "contains slash is invalid", value: "dr/east", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, errs := validateFn(tt.value, "disaster_recovery_identifier")
			if tt.wantErr {
				assert.NotEmpty(t, errs)
			} else {
				assert.Empty(t, errs)
			}
		})
	}
}

func TestDisasterRecoveryTypeValidateFunc(t *testing.T) {
	validateFn := ResourceGitopsAgent().Schema["disaster_recovery_type"].ValidateFunc

	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{name: "PRIMARY is valid", value: "PRIMARY", wantErr: false},
		{name: "SECONDARY is valid", value: "SECONDARY", wantErr: false},
		{name: "invalid value is rejected", value: "TERTIARY", wantErr: true},
		{name: "lowercase is rejected", value: "primary", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, errs := validateFn(tt.value, "disaster_recovery_type")
			if tt.wantErr {
				assert.NotEmpty(t, errs)
			} else {
				assert.Empty(t, errs)
			}
		})
	}
}

func TestDisasterRecoveryTypeDiffSuppress(t *testing.T) {
	suppressFn := ResourceGitopsAgent().Schema["disaster_recovery_type"].DiffSuppressFunc

	assert.False(t, suppressFn("disaster_recovery_type", "", "PRIMARY", nil),
		"must not suppress when old is empty, including non-DR -> DR conversion")
	assert.True(t, suppressFn("disaster_recovery_type", "PRIMARY", "SECONDARY", nil),
		"must suppress once a role is already recorded in state, regardless of new value")
	assert.True(t, suppressFn("disaster_recovery_type", "SECONDARY", "", nil),
		"must suppress when old is non-empty even if new is empty")
}

func TestBuildCreateUpdateAgentRequestDisasterRecoveryNodeName(t *testing.T) {
	d := schema.TestResourceDataRaw(t, ResourceGitopsAgent().Schema, map[string]interface{}{
		"identifier":                   "dragent1",
		"disaster_recovery_identifier": "dragent1",
		"disaster_recovery_type":       "PRIMARY",
	})

	node := buildCreateUpdateAgentRequest(d).DisasterRecoveryNode
	assert.NotNil(t, node)
	assert.Equal(t, "dragent1", node.Identifier)
	assert.Equal(t, "dragent1", node.Name, "DR node name must default to the identifier")
	assert.Equal(t, "PRIMARY", string(*node.Type_))
}

func TestHasLiveSecondarySibling(t *testing.T) {
	primaryType := nextgen.PRIMARY_V1DisasterRecoveryType
	secondaryType := nextgen.SECONDARY_V1DisasterRecoveryType

	tests := []struct {
		name     string
		siblings []nextgen.V1Agent
		want     bool
	}{
		{name: "no siblings", siblings: nil, want: false},
		{
			name: "only primary siblings",
			siblings: []nextgen.V1Agent{
				{DisasterRecoveryNode: &nextgen.V1DisasterRecoveryNode{Type_: &primaryType}},
			},
			want: false,
		},
		{
			name: "secondary sibling present",
			siblings: []nextgen.V1Agent{
				{DisasterRecoveryNode: &nextgen.V1DisasterRecoveryNode{Type_: &primaryType}},
				{DisasterRecoveryNode: &nextgen.V1DisasterRecoveryNode{Type_: &secondaryType}},
			},
			want: true,
		},
		{
			name:     "nil disaster recovery node is skipped",
			siblings: []nextgen.V1Agent{{}},
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, hasLiveSecondarySibling(tt.siblings))
		})
	}
}

func TestValidateDisasterRecoveryDiff(t *testing.T) {
	tests := []struct {
		name                string
		oldIdentifier       string
		newIdentifier       string
		oldType             string
		newType             string
		hasId               bool
		wantForceIdentifier bool
		wantForceType       bool
		wantErr             bool
	}{
		{
			name:                "identifier changed between two non-empty values forces new",
			oldIdentifier:       "dr-east",
			newIdentifier:       "dr-west",
			oldType:             "PRIMARY",
			newType:             "PRIMARY",
			wantForceIdentifier: true,
		},
		{
			name:                "DR to non-DR forces new on both fields",
			oldIdentifier:       "dr-east",
			newIdentifier:       "",
			oldType:             "PRIMARY",
			newType:             "",
			wantForceIdentifier: true,
			wantForceType:       true,
		},
		{
			name:          "non-DR to PRIMARY on existing resource is a plain update",
			oldIdentifier: "",
			newIdentifier: "dr-east",
			oldType:       "",
			newType:       "PRIMARY",
			hasId:         true,
			wantErr:       false,
		},
		{
			name:          "non-DR to SECONDARY on existing resource is a plan-time error",
			oldIdentifier: "",
			newIdentifier: "dr-east",
			oldType:       "",
			newType:       "SECONDARY",
			hasId:         true,
			wantErr:       true,
		},
		{
			name:          "brand-new create as SECONDARY stays allowed",
			oldIdentifier: "",
			newIdentifier: "dr-east",
			oldType:       "",
			newType:       "SECONDARY",
			hasId:         false,
			wantErr:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			forceIdentifier, forceType, err := validateDisasterRecoveryDiff(
				tt.oldIdentifier, tt.newIdentifier, tt.oldType, tt.newType, tt.hasId,
			)
			assert.Equal(t, tt.wantForceIdentifier, forceIdentifier)
			assert.Equal(t, tt.wantForceType, forceType)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
