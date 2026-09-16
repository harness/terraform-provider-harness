package policyset

import (
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
