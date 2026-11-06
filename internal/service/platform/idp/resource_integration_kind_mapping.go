package idp

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceIntegrationKindMapping() *schema.Resource {
	s := mappingWritableSchema()
	for k, v := range mappingReadOnlySchema() {
		s[k] = v
	}
	s["integration_id"] = &schema.Schema{
		Description: "Identifier of the integration.", Type: schema.TypeString, Required: true, ForceNew: true,
	}
	s["org_id"] = &schema.Schema{
		Description: "Organization of the integration.", Type: schema.TypeString, Optional: true, ForceNew: true,
	}
	s["project_id"] = &schema.Schema{
		Description: "Project of the integration.", Type: schema.TypeString, Optional: true, ForceNew: true,
		RequiredWith: []string{"org_id"},
	}
	s["kind"] = &schema.Schema{
		Description: "Kind whose mapping this is.", Type: schema.TypeString, Required: true, ForceNew: true,
	}
	return &schema.Resource{
		Description: "Resource for the mapping of one kind of an IDP integration. Every kind has a mapping, so " +
			"creating this adopts it. Destroying it removes it from the state only, as a kind's mapping exists as " +
			"long as the kind does; on Airbyte types it also restores the type's default field mappings. Don't " +
			"manage a kind with both this and a mapping block on harness_platform_idp_integration, which is only " +
			"sent at create.",
		CreateContext: resourceMappingWrite,
		ReadContext:   resourceMappingRead,
		UpdateContext: resourceMappingWrite,
		DeleteContext: resourceMappingDelete,
		Importer:      &schema.ResourceImporter{StateContext: importMapping},
		Schema:        s,
	}
}

func resourceMappingWrite(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	c := newIntegrationsClient(meta.(*internal.Session))
	org, project, id, kind := mappingKeys(d)
	ws := mappingWritableSchema()
	values := make(map[string]any, len(ws))
	for k := range ws {
		values[k] = d.Get(k)
	}
	out, err := writeMapping(ctx, c, org, project, id, kind, expandObject(ws, values))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(mappingID(org, project, id, kind))
	return diag.FromErr(setMapping(d, out))
}

func resourceMappingRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	c := newIntegrationsClient(meta.(*internal.Session))
	org, project, id, kind := mappingKeys(d)
	out, err := c.getMapping(ctx, org, project, id, kind)
	if isIntegrationsStatus(err, http.StatusNotFound) {
		d.SetId("")
		return nil
	}
	if err != nil {
		return diag.FromErr(err)
	}
	return diag.FromErr(setMapping(d, out))
}

// resourceMappingDelete restores an Airbyte type's default field mappings, keeping the rest of the mapping; other
// types have no defaults to go back to, so it only removes the mapping from the state.
func resourceMappingDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	if fms, _ := d.Get("field_mapping").([]any); len(fms) == 0 {
		return nil
	}
	c := newIntegrationsClient(meta.(*internal.Session))
	org, project, id, kind := mappingKeys(d)
	in, err := c.getIntegration(ctx, org, project, id)
	if isIntegrationsStatus(err, http.StatusNotFound) {
		return nil
	}
	if err != nil {
		return diag.FromErr(err)
	}
	if status, _ := in["status"].(map[string]any); status["mode"] != "airbyte" {
		return nil
	}
	cur, err := c.getMapping(ctx, org, project, id, kind)
	if isIntegrationsStatus(err, http.StatusNotFound) {
		return nil
	}
	if err != nil {
		return diag.FromErr(err)
	}
	values := flattenObject(mappingWritableSchema(), cur)
	delete(values, "field_mapping")
	_, err = writeMapping(ctx, c, org, project, id, kind, expandObject(mappingWritableSchema(), values))
	if err != nil && !isIntegrationsStatus(err, http.StatusNotFound) {
		return diag.Errorf("restoring the default field mappings of kind %s: %s", kind, err)
	}
	return nil
}

// importMapping takes <integration_id>/<kind>, <org_id>/<integration_id>/<kind> or
// <org_id>/<project_id>/<integration_id>/<kind>.
func importMapping(_ context.Context, d *schema.ResourceData, _ any) ([]*schema.ResourceData, error) {
	parts := strings.Split(d.Id(), "/")
	if len(parts) < 2 || len(parts) > 4 {
		return nil, fmt.Errorf("invalid import ID %q", d.Id())
	}
	n := len(parts)
	values := map[string]string{"integration_id": parts[n-2], "kind": parts[n-1]}
	if n >= 3 {
		values["org_id"] = parts[0]
	}
	if n == 4 {
		values["project_id"] = parts[1]
	}
	for k, v := range values {
		if err := d.Set(k, v); err != nil {
			return nil, err
		}
	}
	return []*schema.ResourceData{d}, nil
}
