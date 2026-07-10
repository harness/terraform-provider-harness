package experiment_template

// Schema-validation guardrails. These deterministic tests assert the resource
// and data-source schemas are internally valid (catches malformed schemas that
// would break provider startup) and that core customer-facing functionality is
// wired up: import support and the scope/identity fields.

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestExperimentTemplate_ResourceSchemaValid(t *testing.T) {
	r := ResourceExperimentTemplate()
	if err := r.InternalValidate(r.Schema, true); err != nil {
		t.Fatalf("resource schema failed InternalValidate: %v", err)
	}
	if r.Importer == nil {
		t.Error("expected resource to support import (Importer must be set)")
	}
	for _, field := range []string{"identity", "name", "hub_identity", "org_id", "project_id", "spec"} {
		if _, ok := r.Schema[field]; !ok {
			t.Errorf("expected schema to define %q", field)
		}
	}
}

func TestExperimentTemplate_DataSourceSchemaValid(t *testing.T) {
	ds := DataSourceExperimentTemplate()
	if err := ds.InternalValidate(ds.Schema, false); err != nil {
		t.Fatalf("data source schema failed InternalValidate: %v", err)
	}
}

// assertScopeForceNew is a shared helper: scope/identity fields must be
// ForceNew so a scope change recreates rather than silently corrupts state.
func assertForceNew(t *testing.T, s map[string]*schema.Schema, fields ...string) {
	t.Helper()
	for _, f := range fields {
		if sch, ok := s[f]; ok && !sch.ForceNew {
			t.Errorf("expected %q to be ForceNew", f)
		}
	}
}

func TestExperimentTemplate_ScopeFieldsForceNew(t *testing.T) {
	r := ResourceExperimentTemplate()
	assertForceNew(t, r.Schema, "identity", "org_id", "project_id", "hub_identity")
}
