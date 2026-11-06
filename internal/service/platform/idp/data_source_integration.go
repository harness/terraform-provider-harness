package idp

import (
	"context"

	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
			return diag.FromErr(setIntegration(d, out, false))
		},
		Schema: s,
	}
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
