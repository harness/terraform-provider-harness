package infrastructure_v2

// White-box unit tests for the newly exposed Chaos Infrastructure V2 fields:
//   - resources (cpu/memory requests & limits)
//   - autopilot_enabled
//   - discovery_agent_id
//
// These are deterministic and require no live API / TF_ACC. They lock in the
// expand/set round-trip and the request-build wiring so the fields cannot
// silently regress (the historic discovery_agent_id bug was exactly a
// build-wiring gap that no test caught).

import (
	"testing"

	"github.com/harness/harness-go-sdk/harness/chaos"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestExpandResourceList(t *testing.T) {
	tests := []struct {
		name string
		in   []interface{}
		want *chaos.InfraV2ResourceList
	}{
		{name: "nil", in: nil, want: nil},
		{name: "empty list", in: []interface{}{}, want: nil},
		{name: "nil element", in: []interface{}{nil}, want: nil},
		{name: "empty cpu and memory", in: []interface{}{map[string]interface{}{"cpu": "", "memory": ""}}, want: nil},
		{
			name: "cpu only",
			in:   []interface{}{map[string]interface{}{"cpu": "500m", "memory": ""}},
			want: &chaos.InfraV2ResourceList{Cpu: "500m"},
		},
		{
			name: "memory only",
			in:   []interface{}{map[string]interface{}{"cpu": "", "memory": "256Mi"}},
			want: &chaos.InfraV2ResourceList{Memory: "256Mi"},
		},
		{
			name: "cpu and memory",
			in:   []interface{}{map[string]interface{}{"cpu": "1", "memory": "1Gi"}},
			want: &chaos.InfraV2ResourceList{Cpu: "1", Memory: "1Gi"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := expandResourceList(tt.in)
			switch {
			case tt.want == nil && got != nil:
				t.Fatalf("expected nil, got %+v", got)
			case tt.want != nil && got == nil:
				t.Fatalf("expected %+v, got nil", tt.want)
			case tt.want != nil && (got.Cpu != tt.want.Cpu || got.Memory != tt.want.Memory):
				t.Fatalf("got %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestExpandResources(t *testing.T) {
	tests := []struct {
		name         string
		in           []interface{}
		wantNil      bool
		wantLimits   *chaos.InfraV2ResourceList
		wantRequests *chaos.InfraV2ResourceList
	}{
		{name: "nil", in: nil, wantNil: true},
		{name: "empty element", in: []interface{}{map[string]interface{}{}}, wantNil: true},
		{
			name: "both empty lists collapse to nil",
			in: []interface{}{map[string]interface{}{
				"limits":   []interface{}{},
				"requests": []interface{}{},
			}},
			wantNil: true,
		},
		{
			name: "limits only",
			in: []interface{}{map[string]interface{}{
				"limits": []interface{}{map[string]interface{}{"cpu": "2", "memory": "2Gi"}},
			}},
			wantLimits: &chaos.InfraV2ResourceList{Cpu: "2", Memory: "2Gi"},
		},
		{
			name: "requests and limits",
			in: []interface{}{map[string]interface{}{
				"limits":   []interface{}{map[string]interface{}{"cpu": "2", "memory": "2Gi"}},
				"requests": []interface{}{map[string]interface{}{"cpu": "500m", "memory": "512Mi"}},
			}},
			wantLimits:   &chaos.InfraV2ResourceList{Cpu: "2", Memory: "2Gi"},
			wantRequests: &chaos.InfraV2ResourceList{Cpu: "500m", Memory: "512Mi"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := expandResources(tt.in)
			if tt.wantNil {
				if got != nil {
					t.Fatalf("expected nil, got %+v", got)
				}
				return
			}
			if got == nil {
				t.Fatalf("expected non-nil resources")
			}
			assertResourceList(t, "limits", got.Limits, tt.wantLimits)
			assertResourceList(t, "requests", got.Requests, tt.wantRequests)
		})
	}
}

func assertResourceList(t *testing.T, name string, got, want *chaos.InfraV2ResourceList) {
	t.Helper()
	switch {
	case want == nil && got != nil:
		t.Fatalf("%s: expected nil, got %+v", name, got)
	case want != nil && got == nil:
		t.Fatalf("%s: expected %+v, got nil", name, want)
	case want != nil && (got.Cpu != want.Cpu || got.Memory != want.Memory):
		t.Fatalf("%s: got %+v, want %+v", name, got, want)
	}
}

func TestSetResourcesRoundTrip(t *testing.T) {
	d := schema.TestResourceDataRaw(t, resourceChaosInfrastructureV2Schema(), map[string]interface{}{})

	in := &chaos.InfraV2ResourceRequirements{
		Limits:   &chaos.InfraV2ResourceList{Cpu: "2", Memory: "2Gi"},
		Requests: &chaos.InfraV2ResourceList{Cpu: "500m", Memory: "512Mi"},
	}
	if err := setResources(d, in); err != nil {
		t.Fatalf("setResources returned error: %v", err)
	}

	got := expandResources(d.Get("resources").([]interface{}))
	if got == nil {
		t.Fatalf("round-trip produced nil resources")
	}
	assertResourceList(t, "limits", got.Limits, in.Limits)
	assertResourceList(t, "requests", got.Requests, in.Requests)
}

func TestSetResourcesNilClearsState(t *testing.T) {
	d := schema.TestResourceDataRaw(t, resourceChaosInfrastructureV2Schema(), map[string]interface{}{})
	if err := setResources(d, nil); err != nil {
		t.Fatalf("setResources(nil) returned error: %v", err)
	}
	if v := d.Get("resources").([]interface{}); len(v) != 0 {
		t.Fatalf("expected empty resources for nil input, got %+v", v)
	}
}

func TestBuildRegisterRequestWiresNewFields(t *testing.T) {
	raw := map[string]interface{}{
		"infra_id":           "infra-1",
		"name":               "my-infra",
		"autopilot_enabled":  true,
		"discovery_agent_id": "agent-123",
		"resources": []interface{}{map[string]interface{}{
			"limits":   []interface{}{map[string]interface{}{"cpu": "2", "memory": "2Gi"}},
			"requests": []interface{}{map[string]interface{}{"cpu": "500m", "memory": "512Mi"}},
		}},
	}
	d := schema.TestResourceDataRaw(t, resourceChaosInfrastructureV2Schema(), raw)

	req, err := buildRegisterInfrastructureV2Request(d, "acct-1")
	if err != nil {
		t.Fatalf("buildRegisterInfrastructureV2Request error: %v", err)
	}

	if !req.AutopilotEnabled {
		t.Errorf("AutopilotEnabled = false, want true")
	}
	if req.DiscoveryAgentID != "agent-123" {
		t.Errorf("DiscoveryAgentID = %q, want %q", req.DiscoveryAgentID, "agent-123")
	}
	if req.Resources == nil || req.Resources.Limits == nil || req.Resources.Requests == nil {
		t.Fatalf("Resources not populated: %+v", req.Resources)
	}
	if req.Resources.Limits.Cpu != "2" || req.Resources.Limits.Memory != "2Gi" {
		t.Errorf("limits = %+v, want cpu=2 memory=2Gi", req.Resources.Limits)
	}
	if req.Resources.Requests.Cpu != "500m" || req.Resources.Requests.Memory != "512Mi" {
		t.Errorf("requests = %+v, want cpu=500m memory=512Mi", req.Resources.Requests)
	}
}

func TestBuildRequestsWireSecurityContext(t *testing.T) {
	raw := map[string]interface{}{
		"infra_id":        "infra-1",
		"name":            "my-infra",
		"service_account": "litmus-admin",
		"run_as_user":     1500,
		"run_as_group":    2500,
	}
	d := schema.TestResourceDataRaw(t, resourceChaosInfrastructureV2Schema(), raw)

	reg, err := buildRegisterInfrastructureV2Request(d, "acct-1")
	if err != nil {
		t.Fatalf("buildRegisterInfrastructureV2Request error: %v", err)
	}
	if reg.RunAsUser != 1500 || reg.RunAsGroup != 2500 {
		t.Errorf("register run_as: got user=%d group=%d, want 1500/2500", reg.RunAsUser, reg.RunAsGroup)
	}
	if reg.ServiceAccount != "litmus-admin" {
		t.Errorf("register service_account = %q, want %q", reg.ServiceAccount, "litmus-admin")
	}

	upd, err := buildUpdateInfrastructureV2Request(d, "acct-1")
	if err != nil {
		t.Fatalf("buildUpdateInfrastructureV2Request error: %v", err)
	}
	if upd.RunAsUser != 1500 || upd.RunAsGroup != 2500 {
		t.Errorf("update run_as: got user=%d group=%d, want 1500/2500", upd.RunAsUser, upd.RunAsGroup)
	}
	if upd.ServiceAccount != "litmus-admin" {
		t.Errorf("update service_account = %q, want %q", upd.ServiceAccount, "litmus-admin")
	}
}

func TestBuildUpdateRequestWiresResourcesAndAutopilot(t *testing.T) {
	raw := map[string]interface{}{
		"infra_id":          "infra-1",
		"name":              "my-infra",
		"autopilot_enabled": true,
		"resources": []interface{}{map[string]interface{}{
			"requests": []interface{}{map[string]interface{}{"cpu": "250m", "memory": "128Mi"}},
		}},
	}
	d := schema.TestResourceDataRaw(t, resourceChaosInfrastructureV2Schema(), raw)

	req, err := buildUpdateInfrastructureV2Request(d, "acct-1")
	if err != nil {
		t.Fatalf("buildUpdateInfrastructureV2Request error: %v", err)
	}

	if !req.AutopilotEnabled {
		t.Errorf("AutopilotEnabled = false, want true")
	}
	if req.Resources == nil || req.Resources.Requests == nil {
		t.Fatalf("Resources.Requests not populated: %+v", req.Resources)
	}
	if req.Resources.Requests.Cpu != "250m" || req.Resources.Requests.Memory != "128Mi" {
		t.Errorf("requests = %+v, want cpu=250m memory=128Mi", req.Resources.Requests)
	}
	if req.Resources.Limits != nil {
		t.Errorf("limits = %+v, want nil (not configured)", req.Resources.Limits)
	}
}
