package split

import (
	"encoding/json"
	"reflect"
	"sort"

	splitsdk "github.com/harness/harness-go-sdk/harness/split"
)

// RuleBasedSegmentDefinitionJSONSemanticallyEqualIgnoreTitleComment is true when both strings unmarshal to equivalent
// RuleBasedSegmentDefinition values after clearing title and comment. Use when comparing config/apply JSON to import
// refresh JSON where the Split API omits title/comment on environment-scoped definitions.
func RuleBasedSegmentDefinitionJSONSemanticallyEqualIgnoreTitleComment(a, b string) bool {
	var ar, br splitsdk.RuleBasedSegmentDefinition
	if err := json.Unmarshal([]byte(a), &ar); err != nil {
		return false
	}
	if err := json.Unmarshal([]byte(b), &br); err != nil {
		return false
	}
	ar.Title = ""
	ar.Comment = ""
	br.Title = ""
	br.Comment = ""
	normalizeRuleBasedSegmentDefinition(&ar)
	normalizeRuleBasedSegmentDefinition(&br)
	return reflect.DeepEqual(ar, br)
}

// ruleBasedSegmentDefinitionJSONSemanticallyEqual is true when both strings unmarshal to equivalent RuleBasedSegmentDefinition values.
// Used so HCL jsonencode key order differs from encoding/json field order without perpetual drift or failed import verify.
func ruleBasedSegmentDefinitionJSONSemanticallyEqual(a, b string) bool {
	var ar, br splitsdk.RuleBasedSegmentDefinition
	if err := json.Unmarshal([]byte(a), &ar); err != nil {
		return false
	}
	if err := json.Unmarshal([]byte(b), &br); err != nil {
		return false
	}
	normalizeRuleBasedSegmentDefinition(&ar)
	normalizeRuleBasedSegmentDefinition(&br)
	return reflect.DeepEqual(ar, br)
}

func normalizeRuleBasedSegmentDefinition(d *splitsdk.RuleBasedSegmentDefinition) {
	if d.Rules == nil {
		d.Rules = []splitsdk.RuleBasedSegmentRule{}
	}
	for i := range d.Rules {
		normalizeRuleBasedSegmentRule(&d.Rules[i])
	}
	if d.ExcludedKeys == nil {
		d.ExcludedKeys = []string{}
	}
	// Split and jsonencode may disagree on slice order; excluded keys behave as a set for semantic equality.
	sort.Strings(d.ExcludedKeys)
	if d.ExcludedSegments == nil {
		d.ExcludedSegments = []splitsdk.RuleBasedSegmentRef{}
	}
}

func normalizeRuleBasedSegmentRule(r *splitsdk.RuleBasedSegmentRule) {
	if r.Condition == nil {
		return
	}
	if r.Condition.Matchers == nil {
		r.Condition.Matchers = []splitsdk.Matcher{}
	}
}
