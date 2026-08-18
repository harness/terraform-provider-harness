package helpers

import (
	"context"
	"testing"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func setTagsTestResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"tags": GetTagsSchema(SchemaFlagTypes.Optional),
		},
		ReadContext: func(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
			apiTags, _ := meta.(map[string]string)
			return diag.FromErr(SetOptionalCollection(d, "tags", FlattenTags(apiTags)))
		},
	}
}

func tagsRawState(tags cty.Value) cty.Value {
	return cty.ObjectVal(map[string]cty.Value{
		"id":   cty.StringVal("test"),
		"tags": tags,
	})
}

func TestSetTags(t *testing.T) {
	r := setTagsTestResource()

	tests := []struct {
		name           string
		priorAttrs     map[string]string
		rawState       cty.Value
		setRawState    bool
		apiTags        map[string]string
		wantTagsCount  string // expected tags.# value; empty string means key must be absent
		wantTagsAbsent bool
	}{
		{
			name: "skips_when_prior_null_and_api_empty",
			priorAttrs: map[string]string{
				"id": "test",
			},
			rawState:       tagsRawState(cty.NullVal(cty.Set(cty.String))),
			setRawState:    true,
			apiTags:        map[string]string{},
			wantTagsAbsent: true,
		},
		{
			name: "preserves_prior_empty_set",
			priorAttrs: map[string]string{
				"id":     "test",
				"tags.#": "0",
			},
			rawState:      tagsRawState(cty.SetValEmpty(cty.String)),
			setRawState:   true,
			apiTags:       map[string]string{},
			wantTagsCount: "0",
		},
		{
			name: "writes_when_api_has_tags",
			priorAttrs: map[string]string{
				"id": "test",
			},
			rawState:      tagsRawState(cty.NullVal(cty.Set(cty.String))),
			setRawState:   true,
			apiTags:       map[string]string{"env": "qa"},
			wantTagsCount: "1",
		},
		{
			name: "writes_empty_without_raw_state",
			priorAttrs: map[string]string{
				"id": "test",
			},
			setRawState:   false,
			apiTags:       map[string]string{},
			wantTagsCount: "0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			state := &terraform.InstanceState{
				ID:         "test",
				Attributes: tt.priorAttrs,
			}
			if tt.setRawState {
				state.RawState = tt.rawState
			}

			got, diags := r.RefreshWithoutUpgrade(context.Background(), state, tt.apiTags)
			if diags.HasError() {
				t.Fatalf("RefreshWithoutUpgrade diagnostics: %v", diags)
			}
			if got == nil {
				t.Fatal("RefreshWithoutUpgrade returned nil state")
			}

			count, ok := got.Attributes["tags.#"]
			if tt.wantTagsAbsent {
				if ok {
					t.Fatalf("tags.# present = %q, want absent (null)", count)
				}
				return
			}
			if !ok {
				t.Fatal("tags.# absent, want present")
			}
			if count != tt.wantTagsCount {
				t.Fatalf("tags.# = %q, want %q", count, tt.wantTagsCount)
			}
		})
	}
}

func optionalCollectionTestResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"permissions": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
		ReadContext: func(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
			apiPerms, _ := meta.([]string)
			return diag.FromErr(SetOptionalCollection(d, "permissions", apiPerms))
		},
	}
}

func permissionsRawState(permissions cty.Value) cty.Value {
	return cty.ObjectVal(map[string]cty.Value{
		"id":          cty.StringVal("test"),
		"permissions": permissions,
	})
}

func TestSetOptionalCollection(t *testing.T) {
	r := optionalCollectionTestResource()

	tests := []struct {
		name            string
		priorAttrs      map[string]string
		rawState        cty.Value
		setRawState     bool
		apiPerms        []string
		wantCount       string
		wantCountAbsent bool
	}{
		{
			name: "skips_when_prior_null_and_api_empty",
			priorAttrs: map[string]string{
				"id": "test",
			},
			rawState:        permissionsRawState(cty.NullVal(cty.Set(cty.String))),
			setRawState:     true,
			apiPerms:        []string{},
			wantCountAbsent: true,
		},
		{
			name: "preserves_prior_empty_set",
			priorAttrs: map[string]string{
				"id":            "test",
				"permissions.#": "0",
			},
			rawState:    permissionsRawState(cty.SetValEmpty(cty.String)),
			setRawState: true,
			apiPerms:    []string{},
			wantCount:   "0",
		},
		{
			name: "writes_when_api_has_values",
			priorAttrs: map[string]string{
				"id": "test",
			},
			rawState:    permissionsRawState(cty.NullVal(cty.Set(cty.String))),
			setRawState: true,
			apiPerms:    []string{"core_pipeline_view"},
			wantCount:   "1",
		},
		{
			name: "writes_empty_without_raw_state",
			priorAttrs: map[string]string{
				"id": "test",
			},
			setRawState: false,
			apiPerms:    []string{},
			wantCount:   "0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			state := &terraform.InstanceState{
				ID:         "test",
				Attributes: tt.priorAttrs,
			}
			if tt.setRawState {
				state.RawState = tt.rawState
			}

			got, diags := r.RefreshWithoutUpgrade(context.Background(), state, tt.apiPerms)
			if diags.HasError() {
				t.Fatalf("RefreshWithoutUpgrade diagnostics: %v", diags)
			}
			if got == nil {
				t.Fatal("RefreshWithoutUpgrade returned nil state")
			}

			count, ok := got.Attributes["permissions.#"]
			if tt.wantCountAbsent {
				if ok {
					t.Fatalf("permissions.# present = %q, want absent (null)", count)
				}
				return
			}
			if !ok {
				t.Fatal("permissions.# absent, want present")
			}
			if count != tt.wantCount {
				t.Fatalf("permissions.# = %q, want %q", count, tt.wantCount)
			}
		})
	}
}
