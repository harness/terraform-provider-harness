package helpers

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestParseCtyPath_FlatAttr(t *testing.T) {
	p := parseCtyPath("password_wo")
	if len(p) != 1 {
		t.Fatalf("expected 1-step path, got %d", len(p))
	}
}

func TestParseCtyPath_NestedList(t *testing.T) {
	p := parseCtyPath("repo.0.password_wo")
	if len(p) != 3 {
		t.Fatalf("expected 3-step path, got %d", len(p))
	}
}

func TestParseCtyPath_DeeplyNested(t *testing.T) {
	// request.0.cluster.0.config.0.tls_client_config.0.cert_data_wo → 8 steps
	p := parseCtyPath("request.0.cluster.0.config.0.tls_client_config.0.cert_data_wo")
	if len(p) != 9 {
		t.Fatalf("expected 9-step path, got %d", len(p))
	}
}

func TestWoStringValue_ReturnsFalseWithoutRawConfigContext(t *testing.T) {
	d := schema.TestResourceDataRaw(t, map[string]*schema.Schema{
		"password_wo": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}, map[string]interface{}{
		"password_wo": "secret",
	})

	_, ok := WoStringValue(d, "password_wo")
	if ok {
		t.Fatal("expected WoStringValue to return ok=false in unit-test ResourceData without raw config context")
	}
}

func TestWoActive_UsesVersionFallback(t *testing.T) {
	d := schema.TestResourceDataRaw(t, map[string]*schema.Schema{
		"repo": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"password_wo_version": {
						Type:     schema.TypeInt,
						Optional: true,
					},
				},
			},
		},
	}, map[string]interface{}{
		"repo": []interface{}{
			map[string]interface{}{
				"password_wo_version": 1,
			},
		},
	})

	if !WoActive(d, "repo.0.password_wo", "repo.0.password_wo_version") {
		t.Fatal("expected WoActive=true when version key is present")
	}
}

func TestWoActive_FalseWhenUnset(t *testing.T) {
	d := schema.TestResourceDataRaw(t, map[string]*schema.Schema{
		"repo": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"password_wo_version": {
						Type:     schema.TypeInt,
						Optional: true,
					},
				},
			},
		},
	}, map[string]interface{}{})

	if WoActive(d, "repo.0.password_wo", "repo.0.password_wo_version") {
		t.Fatal("expected WoActive=false when wo and version are both absent")
	}
}
