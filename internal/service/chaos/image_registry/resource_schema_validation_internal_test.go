package image_registry

import "testing"

func TestImageRegistry_ResourceSchemaValid(t *testing.T) {
	r := ResourceChaosImageRegistry()
	if err := r.InternalValidate(r.Schema, true); err != nil {
		t.Fatalf("resource schema failed InternalValidate: %v", err)
	}
	if r.Importer == nil {
		t.Error("expected resource to support import (Importer must be set)")
	}
}

func TestImageRegistry_DataSourceSchemaValid(t *testing.T) {
	ds := DataSourceChaosImageRegistry()
	if err := ds.InternalValidate(ds.Schema, false); err != nil {
		t.Fatalf("data source schema failed InternalValidate: %v", err)
	}
}
