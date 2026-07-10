package experiment

// White-box tests for the tag-preservation logic that prevents drift when the
// chaos API adds system-generated tags (e.g. "fault=<identity>") that are not
// part of the user's configuration. Deterministic; no live API / TF_ACC.

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestFlattenTagsPreserveConfig(t *testing.T) {
	configTags := schema.NewSet(schema.HashString, []interface{}{"env=prod", "team=core"})
	apiTags := []string{"env=prod", "team=core", "fault=pod-delete", "system=auto"}

	out := flattenTagsPreserveConfig(apiTags, configTags)

	if out.Len() != 2 {
		t.Fatalf("expected only the 2 user-configured tags, got %d: %v", out.Len(), out.List())
	}
	if !out.Contains("env=prod") || !out.Contains("team=core") {
		t.Errorf("user tags not preserved: %v", out.List())
	}
	if out.Contains("fault=pod-delete") || out.Contains("system=auto") {
		t.Errorf("system-generated tag leaked into state (would cause perpetual diff): %v", out.List())
	}
}

func TestFlattenTagsPreserveConfig_EmptyConfig(t *testing.T) {
	// With no configured tags, none of the API tags should be retained.
	out := flattenTagsPreserveConfig([]string{"fault=x"}, schema.NewSet(schema.HashString, nil))
	if out.Len() != 0 {
		t.Errorf("expected no tags when config is empty, got %v", out.List())
	}
}

func TestFlattenTags(t *testing.T) {
	out := flattenTags([]string{"a", "b", "c"})
	if out.Len() != 3 {
		t.Fatalf("expected 3 tags, got %d", out.Len())
	}
	for _, want := range []string{"a", "b", "c"} {
		if !out.Contains(want) {
			t.Errorf("missing tag %q", want)
		}
	}
}
