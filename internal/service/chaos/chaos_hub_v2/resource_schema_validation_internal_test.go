package chaos_hub_v2

import "testing"

func TestChaosHubV2_ResourceSchemaValid(t *testing.T) {
	r := ResourceChaosHubV2()
	if err := r.InternalValidate(r.Schema, true); err != nil {
		t.Fatalf("resource schema failed InternalValidate: %v", err)
	}
	if r.Importer == nil {
		t.Error("expected resource to support import (Importer must be set)")
	}
}

func TestChaosHubV2_DataSourceSchemaValid(t *testing.T) {
	ds := DataSourceChaosHubV2()
	if err := ds.InternalValidate(ds.Schema, false); err != nil {
		t.Fatalf("data source schema failed InternalValidate: %v", err)
	}
}
