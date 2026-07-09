package experiment

import "testing"

func TestChaosExperiment_ResourceSchemaValid(t *testing.T) {
	r := ResourceChaosExperiment()
	if err := r.InternalValidate(r.Schema, true); err != nil {
		t.Fatalf("resource schema failed InternalValidate: %v", err)
	}
	if r.Importer == nil {
		t.Error("expected resource to support import (Importer must be set)")
	}
}

func TestChaosExperiment_DataSourceSchemaValid(t *testing.T) {
	ds := DataSourceChaosExperiment()
	if err := ds.InternalValidate(ds.Schema, false); err != nil {
		t.Fatalf("data source schema failed InternalValidate: %v", err)
	}
}
