package split

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRuleBasedSegmentDefinitionJSONSemanticallyEqual_keyOrder(t *testing.T) {
	a := `{"title":"t","comment":"c","rules":[{"condition":{"combiner":"AND","matchers":[{"type":"IN_LIST_STRING","strings":["x"]}]}}]}`
	b := `{"comment":"c","title":"t","rules":[{"condition":{"matchers":[{"strings":["x"],"type":"IN_LIST_STRING"}],"combiner":"AND"}}]}`
	require.True(t, ruleBasedSegmentDefinitionJSONSemanticallyEqual(a, b))
}

func TestRuleBasedSegmentDefinitionJSONSemanticallyEqual_notEqual(t *testing.T) {
	a := `{"title":"t1","rules":[]}`
	b := `{"title":"t2","rules":[]}`
	require.False(t, ruleBasedSegmentDefinitionJSONSemanticallyEqual(a, b))
}

func TestRuleBasedSegmentDefinitionJSONSemanticallyEqualIgnoreTitleComment_ignoresPresentationFields(t *testing.T) {
	withMeta := `{"title":"acc rbs v2","comment":"acc comment v2","rules":[{"condition":{"combiner":"AND","matchers":[{"type":"IN_LIST_STRING","strings":["acc_rbs_key_1"]}]}}]}`
	rulesOnly := `{"rules":[{"condition":{"combiner":"AND","matchers":[{"type":"IN_LIST_STRING","strings":["acc_rbs_key_1"]}]}}]}`
	require.True(t, RuleBasedSegmentDefinitionJSONSemanticallyEqualIgnoreTitleComment(withMeta, rulesOnly))
}

func TestRuleBasedSegmentDefinitionJSONSemanticallyEqualIgnoreTitleComment_stillComparesRules(t *testing.T) {
	a := `{"title":"t","rules":[{"condition":{"combiner":"AND","matchers":[{"type":"IN_LIST_STRING","strings":["x"]}]}}]}`
	b := `{"rules":[{"condition":{"combiner":"AND","matchers":[{"type":"IN_LIST_STRING","strings":["y"]}]}}]}`
	require.False(t, RuleBasedSegmentDefinitionJSONSemanticallyEqualIgnoreTitleComment(a, b))
}
