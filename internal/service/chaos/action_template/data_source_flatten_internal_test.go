package action_template

// White-box test for the data source read path. The data source reuses the
// resource-derived schema (dataSourceActionTemplateSchema) together with the
// shared setActionTemplateData flatten. This asserts that every field the read
// writes has a valid address in the *data source* schema, guarding against the
// "Invalid address to set" / dropped-field class of bug (the same class fixed
// in the probe_template data source). Deterministic; no live API / TF_ACC.

import (
	"testing"

	"github.com/harness/harness-go-sdk/harness/chaos"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestDataSourceActionTemplate_FlattenRoundTrip(t *testing.T) {
	actionType := chaos.DELAY_ActionsActionType
	infraType := chaos.KUBERNETES_V2_ActionsInfrastructureType
	retries := interface{}(float64(2))

	tmpl := &chaos.ChaosactiontemplateChaosActionTemplate{
		AccountID:          "acc",
		Identity:           "e2e-action",
		Name:               "E2E Action",
		Description:        "desc",
		Tags:               []string{"a", "b"},
		Type_:              &actionType,
		InfrastructureType: &infraType,
		Revision:           3,
		IsDefault:          true,
		ActionProperties: &chaos.AllOfchaosactiontemplateChaosActionTemplateActionProperties{
			DelayAction: &chaos.ActionDelayActionTemplate{Duration: "30s"},
		},
		RunProperties: &chaos.ActionActionTemplateRunProperties{
			InitialDelay:  "5s",
			Interval:      "10s",
			Timeout:       "60s",
			Verbosity:     "info",
			StopOnFailure: true,
			MaxRetries:    &retries,
		},
		Variables: []chaos.TemplateVariable{
			{Name: "STR_VAR", Value: ptrIface("hello"), Required: true},
		},
	}

	// Build ResourceData from the DATA SOURCE schema (not the resource schema),
	// so a data-source-only field gap surfaces as an error here.
	d := schema.TestResourceDataRaw(t, dataSourceActionTemplateSchema(), map[string]interface{}{})

	if diags := setActionTemplateData(d, tmpl, "acc", "org", "proj", "hub"); diags.HasError() {
		t.Fatalf("setActionTemplateData returned diagnostics against data source schema: %v", diags)
	}

	checks := map[string]interface{}{
		"identity":                       "e2e-action",
		"name":                           "E2E Action",
		"hub_identity":                   "hub",
		"org_id":                         "org",
		"project_id":                     "proj",
		"description":                    "desc",
		"type":                           "delay",
		"infrastructure_type":            "KubernetesV2",
		"revision":                       3,
		"is_default":                     true,
		"delay_action.0.duration":        "30s",
		"run_properties.0.initial_delay": "5s",
		"run_properties.0.verbosity":     "info",
		"run_properties.0.max_retries":   2,
		"variables.0.name":               "STR_VAR",
		"variables.0.value":              "hello",
		"variables.0.required":           true,
	}
	for path, want := range checks {
		if got := d.Get(path); got != want {
			t.Errorf("%s: got %#v (%T), want %#v", path, got, got, want)
		}
	}
}
