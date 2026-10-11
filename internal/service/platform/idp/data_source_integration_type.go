package idp

import (
	"context"
	"fmt"
	"strings"

	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func computedField(t schema.ValueType) *schema.Schema {
	return &schema.Schema{Type: t, Computed: true}
}

func computedStrings() *schema.Schema {
	return &schema.Schema{Type: schema.TypeList, Computed: true, Elem: &schema.Schema{Type: schema.TypeString}}
}

func integrationTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"display_name":   computedField(schema.TypeString),
		"description":    computedField(schema.TypeString),
		"category":       computedField(schema.TypeString),
		"mode":           computedField(schema.TypeString),
		"connector":      {Description: "required or none.", Type: schema.TypeString, Computed: true},
		"allowed_scopes": computedStrings(),
		"singleton":      computedField(schema.TypeBool),
		"kinds": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{Schema: map[string]*schema.Schema{
				"kind":             computedField(schema.TypeString),
				"secondary":        computedField(schema.TypeBool),
				"primary_kind":     computedField(schema.TypeString),
				"optional":         computedField(schema.TypeBool),
				"category":         computedField(schema.TypeString),
				"sync_modes":       computedStrings(),
				"default_action":   computedField(schema.TypeString),
				"default_idp_kind": computedField(schema.TypeString),
				"default_idp_type": computedField(schema.TypeString),
				"mapping_fields": {
					Type:     schema.TypeList,
					Computed: true,
					Elem: &schema.Resource{Schema: map[string]*schema.Schema{
						"field_name":   computedField(schema.TypeString),
						"display_name": computedField(schema.TypeString),
						"description":  computedField(schema.TypeString),
						"target_path":  computedField(schema.TypeString),
						"required":     computedField(schema.TypeBool),
					}},
				},
			}},
		},
		"default_field_mappings": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{Schema: map[string]*schema.Schema{
				"source_field": computedField(schema.TypeString),
				"target_field": computedField(schema.TypeString),
				"required":     computedField(schema.TypeBool),
			}},
		},
	}
}

func DataSourceIntegrationType() *schema.Resource {
	s := integrationTypeSchema()
	s["type"] = &schema.Schema{Description: "Integration type, e.g. ServiceNow.", Type: schema.TypeString, Required: true}
	return &schema.Resource{
		Description: "Data source for an IDP integration type: its kinds, sync modes and mappable fields.",
		ReadContext: func(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
			c := newIntegrationsClient(meta.(*internal.Session))
			typ := d.Get("type").(string)
			out, err := c.getIntegrationType(ctx, typ)
			if err != nil {
				return diag.FromErr(err)
			}
			d.SetId(typ)
			for k, v := range flattenObject(integrationTypeSchema(), out) {
				if err := d.Set(k, v); err != nil {
					return diag.FromErr(fmt.Errorf("setting %s: %w", k, err))
				}
			}
			return nil
		},
		Schema: s,
	}
}

func DataSourceIntegrationCatalog() *schema.Resource {
	s := map[string]*schema.Schema{
		"integration_id": {Description: "Identifier of the integration.", Type: schema.TypeString, Required: true},
		"version":        computedField(schema.TypeString),
		"streams": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{Schema: map[string]*schema.Schema{
				"name":                 computedField(schema.TypeString),
				"supported_sync_modes": computedStrings(),
				"sync_mode":            computedField(schema.TypeString),
				"cursor_field":         computedStrings(),
				"primary_key": {
					Description: "Primary key paths, each joined with \".\".",
					Type:        schema.TypeList, Computed: true, Elem: &schema.Schema{Type: schema.TypeString},
				},
				"json_schema": {Description: "JSON schema of the stream.", Type: schema.TypeString, Computed: true},
			}},
		},
	}
	helpers.SetMultiLevelDatasourceSchemaWithoutCommonFields(s)
	return &schema.Resource{
		Description: "Data source for the Airbyte streams of an IDP integration.",
		ReadContext: dataSourceIntegrationCatalogRead,
		Schema:      s,
	}
}

func dataSourceIntegrationCatalogRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	c := newIntegrationsClient(meta.(*internal.Session))
	org, project := scope(d)
	id := d.Get("integration_id").(string)
	out, err := c.getCatalog(ctx, org, project, id)
	if err != nil {
		return diag.FromErr(err)
	}
	raw, _ := out["streams"].([]any)
	streams := make([]any, 0, len(raw))
	for _, r := range raw {
		st, _ := r.(map[string]any)
		var keys []any
		pk, _ := st["primary_key"].([]any)
		for _, path := range pk {
			parts, _ := path.([]any)
			names := make([]string, 0, len(parts))
			for _, p := range parts {
				names = append(names, str(p))
			}
			keys = append(keys, strings.Join(names, "."))
		}
		streams = append(streams, map[string]any{
			"name": str(st["name"]), "supported_sync_modes": st["supported_sync_modes"],
			"sync_mode": str(st["sync_mode"]), "cursor_field": st["cursor_field"], "primary_key": keys,
			"json_schema": jsonString(st["json_schema"]),
		})
	}
	d.SetId(mappingID(org, project, id, "catalog"))
	if err := d.Set("version", str(out["version"])); err != nil {
		return diag.FromErr(err)
	}
	return diag.FromErr(d.Set("streams", streams))
}
