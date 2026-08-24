package idp

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/require"
)

func TestExpandScorecardRequest(t *testing.T) {
	resource := ResourceScorecard()
	data := schema.TestResourceDataRaw(t, resource.Schema, map[string]interface{}{
		"identifier":         "gold",
		"name":               "Gold",
		"description":        "Gold standard",
		"published":          true,
		"weightage_strategy": "EQUAL_WEIGHTS",
		"filter": []interface{}{
			map[string]interface{}{
				"kind":      "component",
				"type":      "service",
				"lifecycle": []interface{}{"production"},
			},
		},
		"checks": []interface{}{
			map[string]interface{}{
				"identifier": "readme",
				"custom":     true,
				"weightage":  1.0,
			},
		},
	})

	req := expandScorecardRequest(data)
	require.Equal(t, "gold", req.Scorecard.Identifier)
	require.Equal(t, "Gold", req.Scorecard.Name)
	require.True(t, req.Scorecard.Published)
	require.NotNil(t, req.Scorecard.Filter)
	require.Equal(t, "component", req.Scorecard.Filter.Kind)
	require.Equal(t, "service", req.Scorecard.Filter.Type_)
	require.Equal(t, []string{"production"}, req.Scorecard.Filter.Lifecycle)
	require.Len(t, req.Checks, 1)
	require.Equal(t, "readme", req.Checks[0].Identifier)
	require.True(t, req.Checks[0].Custom)
}

func TestScorecardCheckInputValuesSchemaTracksRemoval(t *testing.T) {
	resourceRules := ResourceScorecardCheck().Schema["rules"]
	require.True(t, resourceRules.Optional)
	require.False(t, resourceRules.Computed)
	resourceInputValues := resourceRules.Elem.(*schema.Resource).Schema["input_values"]
	require.True(t, resourceInputValues.Optional)
	require.False(t, resourceInputValues.Computed)

	dataSourceRules := DataSourceScorecardCheck().Schema["rules"]
	require.False(t, dataSourceRules.Optional)
	require.True(t, dataSourceRules.Computed)
	dataSourceInputValues := dataSourceRules.Elem.(*schema.Resource).Schema["input_values"]
	require.False(t, dataSourceInputValues.Optional)
	require.True(t, dataSourceInputValues.Computed)
}

func TestExpandCheckDetails(t *testing.T) {
	resource := ResourceScorecardCheck()
	data := schema.TestResourceDataRaw(t, resource.Schema, map[string]interface{}{
		"identifier":        "readme",
		"name":              "README exists",
		"description":       "Ensure README",
		"rule_strategy":     "ALL_OF",
		"default_behaviour": "FAIL",
		"rules": []interface{}{
			map[string]interface{}{
				"data_source_identifier": "github",
				"data_point_identifier":  "isFileExists",
				"operator":               "==",
				"value":                  "true",
			},
		},
	})

	details := expandCheckDetails(data)
	require.Equal(t, "readme", details.Identifier)
	require.Equal(t, "README exists", details.Name)
	require.True(t, details.Custom)
	require.Equal(t, "ALL_OF", details.RuleStrategy)
	require.Equal(t, "FAIL", details.DefaultBehaviour)
	require.Len(t, details.Rules, 1)
	require.Equal(t, "github", details.Rules[0].DataSourceIdentifier)
	require.Equal(t, "isFileExists", details.Rules[0].DataPointIdentifier)
}

func TestValidateScorecardCheckConfig(t *testing.T) {
	rule := map[string]interface{}{
		"data_source_identifier": "github",
		"data_point_identifier":  "isFileExists",
		"operator":               "==",
		"value":                  "true",
	}

	tests := []struct {
		name         string
		ruleStrategy string
		expression   string
		rules        interface{}
		wantErr      string
	}{
		{
			name:         "all_of requires rules",
			ruleStrategy: "ALL_OF",
			wantErr:      "rules is required when rule_strategy is ALL_OF",
		},
		{
			name:         "any_of requires rules",
			ruleStrategy: "ANY_OF",
			rules:        []interface{}{},
			wantErr:      "rules is required when rule_strategy is ANY_OF",
		},
		{
			name:         "all_of with rules",
			ruleStrategy: "ALL_OF",
			rules:        []interface{}{rule},
		},
		{
			name:         "advanced requires expression",
			ruleStrategy: "ADVANCED",
			wantErr:      "expression is required when rule_strategy is ADVANCED",
		},
		{
			name:         "advanced with blank expression",
			ruleStrategy: "ADVANCED",
			expression:   "   ",
			wantErr:      "expression is required when rule_strategy is ADVANCED",
		},
		{
			name:         "advanced with expression",
			ruleStrategy: "ADVANCED",
			expression:   "catalog.metadata.name != null",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateScorecardCheckConfig(tt.ruleStrategy, tt.expression, tt.rules)
			if tt.wantErr == "" {
				require.NoError(t, err)
				return
			}
			require.EqualError(t, err, tt.wantErr)
		})
	}
}

func TestScorecardSchemasInternalValidate(t *testing.T) {
	scorecard := ResourceScorecard()
	require.NoError(t, scorecard.InternalValidate(scorecard.Schema, true))

	scorecardDS := DataSourceScorecard()
	require.NoError(t, scorecardDS.InternalValidate(scorecardDS.Schema, false))

	check := ResourceScorecardCheck()
	require.NoError(t, check.InternalValidate(check.Schema, true))

	checkDS := DataSourceScorecardCheck()
	require.NoError(t, checkDS.InternalValidate(checkDS.Schema, false))
}
