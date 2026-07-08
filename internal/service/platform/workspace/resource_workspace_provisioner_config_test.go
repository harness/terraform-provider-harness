package workspace

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestBuildProvisionerConfig(t *testing.T) {
	tests := []struct {
		name       string
		rawConfig  map[string]interface{}
		wantNil    bool
		wantLang   string
		wantLangV  string
		wantPkgMgr string
		wantPkgV   string
	}{
		{
			name: "valid python config",
			rawConfig: map[string]interface{}{
				"provisioner_config": []interface{}{
					map[string]interface{}{
						"language":                "python",
						"language_version":        "3.12",
						"package_manager":         "pip",
						"package_manager_version": "25.3",
					},
				},
			},
			wantNil:    false,
			wantLang:   "python",
			wantLangV:  "3.12",
			wantPkgMgr: "pip",
			wantPkgV:   "25.3",
		},
		{
			name: "valid typescript config",
			rawConfig: map[string]interface{}{
				"provisioner_config": []interface{}{
					map[string]interface{}{
						"language":                "typescript",
						"language_version":        "5.4",
						"package_manager":         "npm",
						"package_manager_version": "10.2",
					},
				},
			},
			wantNil:    false,
			wantLang:   "typescript",
			wantLangV:  "5.4",
			wantPkgMgr: "npm",
			wantPkgV:   "10.2",
		},
		{
			name:      "no provisioner_config set",
			rawConfig: map[string]interface{}{},
			wantNil:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resourceSchema := ResourceWorkspace().Schema
			d := schema.TestResourceDataRaw(t, resourceSchema, tt.rawConfig)

			result := buildProvisionerConfig(d)

			if tt.wantNil {
				if result != nil {
					t.Fatalf("expected nil, got %v", result)
				}
				return
			}

			if result == nil {
				t.Fatal("expected non-nil result")
			}

			if result.Language != tt.wantLang {
				t.Errorf("Language = %v, want %v", result.Language, tt.wantLang)
			}
			if result.LanguageVersion != tt.wantLangV {
				t.Errorf("LanguageVersion = %v, want %v", result.LanguageVersion, tt.wantLangV)
			}
			if result.PackageManager != tt.wantPkgMgr {
				t.Errorf("PackageManager = %v, want %v", result.PackageManager, tt.wantPkgMgr)
			}
			if result.PackageManagerVersion != tt.wantPkgV {
				t.Errorf("PackageManagerVersion = %v, want %v", result.PackageManagerVersion, tt.wantPkgV)
			}
		})
	}
}

func TestProvisionerConfigSchema(t *testing.T) {
	resourceSchema := ResourceWorkspace().Schema

	pcSchema, ok := resourceSchema["provisioner_config"]
	if !ok {
		t.Fatal("provisioner_config not found in schema")
	}

	if pcSchema.Type != schema.TypeSet {
		t.Errorf("expected TypeSet, got %v", pcSchema.Type)
	}

	if pcSchema.MaxItems != 1 {
		t.Errorf("expected MaxItems=1, got %d", pcSchema.MaxItems)
	}

	if !pcSchema.Optional {
		t.Error("expected provisioner_config to be Optional")
	}

	elem := pcSchema.Elem.(*schema.Resource)
	requiredFields := []string{"language", "language_version", "package_manager", "package_manager_version"}
	for _, field := range requiredFields {
		fieldSchema, ok := elem.Schema[field]
		if !ok {
			t.Errorf("field %q not found in provisioner_config schema", field)
			continue
		}
		if !fieldSchema.Required {
			t.Errorf("field %q should be Required", field)
		}
		if fieldSchema.Type != schema.TypeString {
			t.Errorf("field %q should be TypeString, got %v", field, fieldSchema.Type)
		}
	}
}
