package variable_set

import "testing"

func TestVariableSet_DataSourceSchemaValid(t *testing.T) {
	ds := DataSourceVariableSet()
	if err := ds.InternalValidate(ds.Schema, false); err != nil {
		t.Fatalf("data source schema failed InternalValidate: %v", err)
	}
}

// The lookup only uses identifier plus the org and project scope, so every attribute the
// API returns must be read-only. Anything else reads as a configurable input in the
// generated docs and invites users to copy blocks over from the resource schema.
func TestVariableSet_DataSourceReadOnlyAttributes(t *testing.T) {
	ds := DataSourceVariableSet()

	for _, attr := range []string{
		"connector",
		"environment_variable",
		"terraform_variable",
		"terraform_variable_file",
		"description",
		"tags",
	} {
		s, ok := ds.Schema[attr]
		if !ok {
			t.Errorf("%s: missing from data source schema", attr)
			continue
		}
		if !s.Computed {
			t.Errorf("%s: expected Computed", attr)
		}
		if s.Optional {
			t.Errorf("%s: expected read-only, but it is Optional", attr)
		}
		if s.Required {
			t.Errorf("%s: expected read-only, but it is Required", attr)
		}
	}
}

func TestVariableSet_DataSourceScopeAttributes(t *testing.T) {
	ds := DataSourceVariableSet()

	if !ds.Schema["identifier"].Required {
		t.Error("identifier: expected Required")
	}

	// org_id and project_id select the scope the Variable Set is looked up in, so both
	// must stay optional: omitting them targets the account level.
	for _, attr := range []string{"org_id", "project_id"} {
		if !ds.Schema[attr].Optional {
			t.Errorf("%s: expected Optional so the account level can be targeted", attr)
		}
		if ds.Schema[attr].Required {
			t.Errorf("%s: expected Optional, got Required", attr)
		}
	}

	// name is ignored by the lookup, but stays configurable so configs that already set
	// it keep planning rather than failing on an unconfigurable attribute.
	if !ds.Schema["name"].Optional {
		t.Error("name: expected to remain Optional for backwards compatibility")
	}
	if !ds.Schema["name"].Computed {
		t.Error("name: expected Computed")
	}
}

// Every attribute needs a description, since the registry docs are generated from them.
func TestVariableSet_DataSourceDescriptions(t *testing.T) {
	ds := DataSourceVariableSet()

	if ds.Description == "" {
		t.Error("data source is missing a description")
	}

	for name, s := range ds.Schema {
		if s.Description == "" {
			t.Errorf("%s: missing description", name)
		}
	}
}
