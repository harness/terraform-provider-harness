package infrastructure_v2

import "testing"

func TestInfrastructureV2_ResourceSchemaValid(t *testing.T) {
	r := ResourceChaosInfrastructureV2()
	if err := r.InternalValidate(r.Schema, true); err != nil {
		t.Fatalf("resource schema failed InternalValidate: %v", err)
	}
	if r.Importer == nil {
		t.Error("expected resource to support import (Importer must be set)")
	}
}

func TestInfrastructureV2_DataSourceSchemaValid(t *testing.T) {
	ds := DataSourceChaosInfrastructureV2()
	if err := ds.InternalValidate(ds.Schema, false); err != nil {
		t.Fatalf("data source schema failed InternalValidate: %v", err)
	}
}
