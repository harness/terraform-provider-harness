package security_governance_v3

import "testing"

func TestSecurityGovernanceV3_ConditionSchemaValid(t *testing.T) {
	r := ResourceChaosSecurityGovernanceConditionV3()
	if err := r.InternalValidate(r.Schema, true); err != nil {
		t.Fatalf("condition resource schema failed InternalValidate: %v", err)
	}
	if r.Importer == nil {
		t.Error("expected condition resource to support import (Importer must be set)")
	}
	ds := DataSourceChaosSecurityGovernanceConditionV3()
	if err := ds.InternalValidate(ds.Schema, false); err != nil {
		t.Fatalf("condition data source schema failed InternalValidate: %v", err)
	}
}

func TestSecurityGovernanceV3_RuleSchemaValid(t *testing.T) {
	r := ResourceChaosSecurityGovernanceRuleV3()
	if err := r.InternalValidate(r.Schema, true); err != nil {
		t.Fatalf("rule resource schema failed InternalValidate: %v", err)
	}
	if r.Importer == nil {
		t.Error("expected rule resource to support import (Importer must be set)")
	}
	ds := DataSourceChaosSecurityGovernanceRuleV3()
	if err := ds.InternalValidate(ds.Schema, false); err != nil {
		t.Fatalf("rule data source schema failed InternalValidate: %v", err)
	}
}
