package idp

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// identifierPattern and reservedIdentifiers are Harness's entity identifier rule, as the integrations API applies it.
var identifierPattern = regexp.MustCompile(`^[a-zA-Z_][0-9a-zA-Z_$]{0,127}$`)

var reservedIdentifiers = []string{
	"or", "and", "eq", "ne", "lt", "gt", "le", "ge", "div", "mod", "not", "null", "true", "false", "new", "var",
	"return", "step", "parallel", "stepGroup", "org", "account", "class", "shellScriptProvisioner", "organizations",
	"organization", "orgs", "projects", "project",
}

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
			"creating this adopts it; destroying it leaves the mapping as it is.",
		CreateContext: resourceMappingWrite,
		ReadContext:   resourceMappingRead,
		UpdateContext: resourceMappingWrite,
		DeleteContext: func(context.Context, *schema.ResourceData, any) diag.Diagnostics { return nil },
		Importer:      &schema.ResourceImporter{StateContext: importMapping},
		Schema:        s,
	}
}

func mappingWritableSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"unique_identifier_field": fieldPath(defaultedString("Source field that identifies an entity. Can't change on Airbyte types.")),
		"display_name_field":      fieldPath(defaultedString("Source field used as the entity's name.")),
		"field_mapping": {
			Description: "Maps a source field to an entity field. Airbyte types.",
			Type:        schema.TypeList,
			Optional:    true,
			Computed:    true,
			Elem: &schema.Resource{Schema: map[string]*schema.Schema{
				"source_field": fieldPath(requiredString("Source field.")),
				"target_field": fieldPath(requiredString("Entity field, e.g. spec.owner.")),
			}},
		},
		"idp_kind": defaultedString("Entity kind, e.g. component."),
		"idp_type": defaultedString("Entity type, e.g. service."),
		"correlation": block("Links records to existing entities.", false, map[string]*schema.Schema{
			"source_path": fieldPath(requiredString("Record field to match, e.g. name.")),
			"target_path": fieldPath(requiredString("Entity field to match, e.g. metadata.identifier.")),
			"operator": {
				Description: "Match operator; eq when omitted.", Type: schema.TypeString, Optional: true, Computed: true,
			},
		}),
		"ingestion": {
			Description: "Custom types: JQ that builds the entity from a payload.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Computed:    true,
			Elem: &schema.Resource{Schema: map[string]*schema.Schema{
				"identifier": requiredString("JQ for the entity identifier."),
				"name":       optionalString("JQ for the entity name."),
				"properties": {
					Description: "JQ per entity property.", Type: schema.TypeMap, Optional: true,
					Elem: &schema.Schema{Type: schema.TypeString},
				},
			}},
		},
	}
}

// fieldPath ignores a leading dot, which the API drops.
func fieldPath(s *schema.Schema) *schema.Schema {
	s.DiffSuppressFunc = func(_, old, new string, _ *schema.ResourceData) bool {
		return strings.TrimPrefix(old, ".") == strings.TrimPrefix(new, ".")
	}
	return s
}

func mappingReadOnlySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"emission": {
			Description: "How entities are emitted; read-only.",
			Type:        schema.TypeList,
			Computed:    true,
			Elem: &schema.Resource{Schema: map[string]*schema.Schema{
				"identifier": {Type: schema.TypeString, Computed: true},
				"name":       {Type: schema.TypeString, Computed: true},
				"properties": {Description: "JSON object.", Type: schema.TypeString, Computed: true},
			}},
		},
		"payload_schema": {Description: "JSON schema of the payload; read-only.", Type: schema.TypeString, Computed: true},
		"is_default":     {Description: "Whether the mapping is the type's default.", Type: schema.TypeBool, Computed: true},
		"revision":       {Description: "The integration's revision.", Type: schema.TypeInt, Computed: true},
		"updated":        {Description: "Last update time, in milliseconds since the epoch.", Type: schema.TypeInt, Computed: true},
	}
}

func mappingKeys(d *schema.ResourceData) (org, project, id, kind string) {
	return d.Get("org_id").(string), d.Get("project_id").(string), d.Get("integration_id").(string),
		d.Get("kind").(string)
}

func resourceMappingWrite(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	c := newIntegrationsClient(meta.(*internal.Session))
	org, project, id, kind := mappingKeys(d)
	ws := mappingWritableSchema()
	values := make(map[string]any, len(ws))
	for k := range ws {
		values[k] = d.Get(k)
	}
	body := expandObject(ws, values)

	var out map[string]any
	for attempt := 0; attempt < 2; attempt++ {
		cur, err := c.getMapping(ctx, org, project, id, kind)
		if err != nil {
			return diag.FromErr(err)
		}
		out, err = c.replaceMapping(ctx, org, project, id, kind, int64(num(cur["revision"])), body)
		if attempt == 0 && isIntegrationsStatus(err, http.StatusPreconditionFailed) {
			continue
		}
		if err != nil {
			return diag.FromErr(err)
		}
		break
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

func setMapping(d *schema.ResourceData, m map[string]any) error {
	set := flattenObject(mappingWritableSchema(), m)
	set["is_default"] = m["is_default"]
	set["revision"] = num(m["revision"])
	set["updated"] = num(m["updated"])
	set["payload_schema"] = jsonString(m["payload_schema"])
	set["emission"] = []any{}
	if e, ok := m["emission"].(map[string]any); ok {
		set["emission"] = []any{map[string]any{
			"identifier": str(e["identifier"]), "name": str(e["name"]), "properties": jsonString(e["properties"]),
		}}
	}
	for k, v := range set {
		if err := d.Set(k, v); err != nil {
			return fmt.Errorf("setting %s: %w", k, err)
		}
	}
	return nil
}

func jsonString(v any) string {
	if v == nil {
		return ""
	}
	raw, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(raw)
}

func mappingID(org, project, id, kind string) string {
	parts := []string{}
	for _, p := range []string{org, project} {
		if p != "" {
			parts = append(parts, p)
		}
	}
	return strings.Join(append(parts, id, kind), "/")
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
