package fault_template

import "testing"

func TestFaultTemplate_ResourceSchemaValid(t *testing.T) {
	r := ResourceFaultTemplate()
	if err := r.InternalValidate(r.Schema, true); err != nil {
		t.Fatalf("resource schema failed InternalValidate: %v", err)
	}
	if r.Importer == nil {
		t.Error("expected resource to support import (Importer must be set)")
	}
}

func TestFaultTemplate_DataSourceSchemaValid(t *testing.T) {
	ds := DataSourceFaultTemplate()
	if err := ds.InternalValidate(ds.Schema, false); err != nil {
		t.Fatalf("data source schema failed InternalValidate: %v", err)
	}
}
