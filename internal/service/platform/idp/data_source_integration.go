package idp

import (
	"context"
	"net/url"

	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// computedCopy turns a resource schema into its data source form.
func computedCopy(m map[string]*schema.Schema) map[string]*schema.Schema {
	out := make(map[string]*schema.Schema, len(m))
	for k, s := range m {
		c := &schema.Schema{Type: s.Type, Description: s.Description, Computed: true}
		switch e := s.Elem.(type) {
		case *schema.Resource:
			c.Elem = &schema.Resource{Schema: computedCopy(e.Schema)}
		case *schema.Schema:
			c.Elem = &schema.Schema{Type: e.Type}
		}
		out[k] = c
	}
	return out
}

func DataSourceIntegration() *schema.Resource {
	s := computedCopy(integrationSpecSchema())
	for k, v := range computedCopy(integrationStatusSchema()) {
		s[k] = v
	}
	helpers.SetMultiLevelDatasourceSchemaIdentifierRequired(s)
	s["type"] = &schema.Schema{Description: "Integration type.", Type: schema.TypeString, Computed: true}
	s["enabled"] = &schema.Schema{Description: "Whether the integration syncs.", Type: schema.TypeBool, Computed: true}
	return &schema.Resource{
		Description: "Data source for an IDP integration.",
		ReadContext: func(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
			c := newIntegrationsClient(meta.(*internal.Session))
			org, project := scope(d)
			out, err := c.getIntegration(ctx, org, project, d.Get("identifier").(string))
			if err != nil {
				return diag.FromErr(err)
			}
			return diag.FromErr(setIntegration(d, out))
		},
		Schema: s,
	}
}

func DataSourceIntegrations() *schema.Resource {
	s := map[string]*schema.Schema{
		"scope": {
			Description:  "self (default), ancestors or descendants of the given scope.",
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringInSlice([]string{"self", "ancestors", "descendants"}, false),
		},
		"type":        {Description: "Only these types.", Type: schema.TypeSet, Optional: true, Elem: &schema.Schema{Type: schema.TypeString}},
		"health":      {Description: "healthy, degraded, failing or unknown.", Type: schema.TypeString, Optional: true},
		"search_term": {Description: "Matches name and identifier.", Type: schema.TypeString, Optional: true},
		"integrations": {
			Description: "Matching integrations.",
			Type:        schema.TypeList,
			Computed:    true,
			Elem: &schema.Resource{Schema: map[string]*schema.Schema{
				"identifier": {Type: schema.TypeString, Computed: true},
				"name":       {Type: schema.TypeString, Computed: true},
				"type":       {Type: schema.TypeString, Computed: true},
				"mode":       {Type: schema.TypeString, Computed: true},
				"enabled":    {Type: schema.TypeBool, Computed: true},
				"org_id":     {Type: schema.TypeString, Computed: true},
				"project_id": {Type: schema.TypeString, Computed: true},
				"kinds":      {Type: schema.TypeList, Computed: true, Elem: &schema.Schema{Type: schema.TypeString}},
				"health":     {Type: schema.TypeString, Computed: true},
				"last_sync_status": {
					Description: "Status of the last sync, if any.", Type: schema.TypeString, Computed: true,
				},
			}},
		},
	}
	helpers.SetMultiLevelDatasourceSchemaWithoutCommonFields(s)
	return &schema.Resource{
		Description: "Data source for listing IDP integrations.",
		ReadContext: dataSourceIntegrationsRead,
		Schema:      s,
	}
}

func dataSourceIntegrationsRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	c := newIntegrationsClient(meta.(*internal.Session))
	org, project := scope(d)
	q := url.Values{}
	if v := d.Get("scope").(string); v == "ancestors" || v == "descendants" {
		q.Set("scope", v)
	}
	for _, t := range d.Get("type").(*schema.Set).List() {
		q.Add("type", t.(string))
	}
	for _, k := range []string{"health", "search_term"} {
		if v := d.Get(k).(string); v != "" {
			q.Set(k, v)
		}
	}
	items, err := c.listIntegrations(ctx, org, project, q)
	if err != nil {
		return diag.FromErr(err)
	}
	out := make([]any, 0, len(items))
	for _, it := range items {
		lastSync, _ := it["last_sync"].(map[string]any)
		out = append(out, map[string]any{
			"identifier": str(it["identifier"]), "name": str(it["name"]), "type": str(it["type"]),
			"mode": str(it["mode"]), "enabled": it["enabled"] == true, "org_id": str(it["org"]),
			"project_id": str(it["project"]), "kinds": it["kinds"], "health": str(it["health"]),
			"last_sync_status": str(lastSync["status"]),
		})
	}
	d.SetId(mappingID(org, project, "integrations", q.Encode()))
	return diag.FromErr(d.Set("integrations", out))
}

func DataSourceIntegrationKindMapping() *schema.Resource {
	s := computedCopy(mappingWritableSchema())
	for k, v := range mappingReadOnlySchema() {
		s[k] = v
	}
	s["integration_id"] = &schema.Schema{Description: "Identifier of the integration.", Type: schema.TypeString, Required: true}
	s["kind"] = &schema.Schema{Description: "Kind.", Type: schema.TypeString, Required: true}
	helpers.SetMultiLevelDatasourceSchemaWithoutCommonFields(s)
	return &schema.Resource{
		Description: "Data source for the mapping of one kind of an IDP integration.",
		ReadContext: func(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
			c := newIntegrationsClient(meta.(*internal.Session))
			org, project, id, kind := mappingKeys(d)
			out, err := c.getMapping(ctx, org, project, id, kind)
			if err != nil {
				return diag.FromErr(err)
			}
			d.SetId(mappingID(org, project, id, kind))
			return diag.FromErr(setMapping(d, out))
		},
		Schema: s,
	}
}
