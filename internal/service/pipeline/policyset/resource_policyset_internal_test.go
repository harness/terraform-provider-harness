package policyset

import (
	"strings"
	"testing"

	"github.com/harness/harness-go-sdk/harness/policymgmt"
)

func TestScopedPolicyIdentifier(t *testing.T) {
	cases := []struct {
		name                               string
		policy                             policymgmt.LinkedPolicy
		policySetOrgId, policySetProjectId string
		want                               string
	}{
		{"account policy, account policyset", policymgmt.LinkedPolicy{Identifier: "my_policy"}, "", "", "my_policy"},
		{"account policy, org policyset", policymgmt.LinkedPolicy{Identifier: "my_policy"}, "org1", "", "account.my_policy"},
		{"account policy, project policyset", policymgmt.LinkedPolicy{Identifier: "my_policy"}, "org1", "proj1", "account.my_policy"},
		{"org policy, same org policyset", policymgmt.LinkedPolicy{Identifier: "my_policy", OrgId: "org1"}, "org1", "", "my_policy"},
		{"org policy, project policyset in same org", policymgmt.LinkedPolicy{Identifier: "my_policy", OrgId: "org1"}, "org1", "proj1", "org.my_policy"},
		{"project policy", policymgmt.LinkedPolicy{Identifier: "my_policy", OrgId: "org1", ProjectId: "proj1"}, "org1", "proj1", "my_policy"},
	}

	for _, c := range cases {
		if got := scopedPolicyIdentifier(c.policy, c.policySetOrgId, c.policySetProjectId); got != c.want {
			t.Errorf("%s: scopedPolicyIdentifier() = %q, want %q", c.name, got, c.want)
		}
	}
}

func TestFlattenPoliciesPinsExistingOrder(t *testing.T) {
	// API returns policies reordered relative to the prior state - the ordered "policies" list
	// must not reflect that shuffle as long as membership is unchanged.
	previousOrder := []interface{}{
		map[string]interface{}{"identifier": "a", "severity": "error"},
		map[string]interface{}{"identifier": "b", "severity": "error"},
		map[string]interface{}{"identifier": "c", "severity": "error"},
	}
	apiOrder := []policymgmt.LinkedPolicy{
		{Identifier: "b"},
		{Identifier: "c"},
		{Identifier: "a"},
	}

	got := flattenPolicies(apiOrder, "", "", previousOrder)

	want := []string{"a", "b", "c"}
	for i, w := range want {
		if got[i]["identifier"] != w {
			t.Fatalf("position %d: got %q, want %q (full: %v)", i, got[i]["identifier"], w, got)
		}
	}
}

func TestCheckDuplicatePolicyIdentifiers(t *testing.T) {
	if err := checkDuplicatePolicyIdentifiers("policy_references", map[string][]string{
		"a": {"error"},
		"b": {"warning"},
	}); err != nil {
		t.Fatalf("expected no error for unique identifiers, got %v", err)
	}

	err := checkDuplicatePolicyIdentifiers("policy_references", map[string][]string{
		"org.policy_a948": {"error", "warning"},
	})
	if err == nil {
		t.Fatal("expected error for duplicate identifier, got nil")
	}
	if got := err.Error(); !strings.Contains(got, "org.policy_a948") || !strings.Contains(got, "policy_references") {
		t.Fatalf("error message missing identifier/field context: %v", got)
	}
}

func TestFlattenPoliciesAppendsNewEntries(t *testing.T) {
	previousOrder := []interface{}{
		map[string]interface{}{"identifier": "a", "severity": "error"},
	}
	apiOrder := []policymgmt.LinkedPolicy{
		{Identifier: "b"},
		{Identifier: "a"},
	}

	got := flattenPolicies(apiOrder, "", "", previousOrder)

	if len(got) != 2 || got[0]["identifier"] != "a" || got[1]["identifier"] != "b" {
		t.Fatalf("got %v", got)
	}
}
