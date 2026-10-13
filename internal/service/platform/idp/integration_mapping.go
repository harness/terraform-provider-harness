package idp

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// identifierPattern and reservedIdentifiers are Harness's entity identifier rule, as the integrations API applies it.
var identifierPattern = regexp.MustCompile(`^[a-zA-Z_][0-9a-zA-Z_$]{0,127}$`)

var reservedIdentifiers = []string{
	"or", "and", "eq", "ne", "lt", "gt", "le", "ge", "div", "mod", "not", "null", "true", "false", "new", "var",
	"return", "step", "parallel", "stepGroup", "org", "account", "class", "shellScriptProvisioner", "organizations",
	"organization", "orgs", "projects", "project",
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
			"source_path": fieldPath(requiredString("Field to match. Airbyte types: a path in the emitted entity " +
				"data, e.g. metadata.integration_properties.ServiceNow.name; Custom: a payload field.")),
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

// writeMapping replaces a kind's mapping at the current revision, retrying once if another write moved it.
func writeMapping(
	ctx context.Context, c *integrationsClient, org, project, id, kind string, body map[string]any,
) (map[string]any, error) {
	for attempt := 0; ; attempt++ {
		cur, err := c.getMapping(ctx, org, project, id, kind)
		if err != nil {
			return nil, err
		}
		out, err := c.replaceMapping(ctx, org, project, id, kind, int64(num(cur["revision"])), body)
		if attempt == 0 && isIntegrationsStatus(err, http.StatusPreconditionFailed) {
			continue
		}
		return out, err
	}
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

func mappingKeys(d *schema.ResourceData) (org, project, id, kind string) {
	return d.Get("org_id").(string), d.Get("project_id").(string), d.Get("integration_id").(string),
		d.Get("kind").(string)
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
