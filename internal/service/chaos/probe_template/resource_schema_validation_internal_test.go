package probe_template

import "testing"

func TestProbeTemplate_ResourceSchemaValid(t *testing.T) {
	r := ResourceProbeTemplate()
	if err := r.InternalValidate(r.Schema, true); err != nil {
		t.Fatalf("resource schema failed InternalValidate: %v", err)
	}
	if r.Importer == nil {
		t.Error("expected resource to support import (Importer must be set)")
	}
}

func TestProbeTemplate_DataSourceSchemaValid(t *testing.T) {
	ds := DataSourceProbeTemplate()
	if err := ds.InternalValidate(ds.Schema, false); err != nil {
		t.Fatalf("data source schema failed InternalValidate: %v", err)
	}
}
