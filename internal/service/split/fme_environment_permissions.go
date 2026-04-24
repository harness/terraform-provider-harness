package split

import (
	splitsdk "github.com/harness/harness-go-sdk/harness/split"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func permissionEntityListSchema(description string) *schema.Schema {
	return &schema.Schema{
		Description: description,
		Type:        schema.TypeList,
		Optional:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"id": {
					Description: "Identifier of the user, group, or API key.",
					Type:        schema.TypeString,
					Required:    true,
				},
				"name": {
					Description: "Display name (resolved by the API; may differ from the value provided at creation).",
					Type:        schema.TypeString,
					Optional:    true,
					Computed:    true,
				},
				"type": {
					Description: "Entity type: `user`, `group`, or `api_key` (a Harness service account).",
					Type:        schema.TypeString,
					Required:    true,
				},
			},
		},
	}
}

func expandChangePermissions(d *schema.ResourceData) *splitsdk.ChangePermissions {
	raw := d.Get("change_permissions").([]interface{})
	if len(raw) == 0 || raw[0] == nil {
		return nil
	}
	m := raw[0].(map[string]interface{})
	cp := &splitsdk.ChangePermissions{}

	if v, ok := m["allow_kills"]; ok {
		b := v.(bool)
		cp.AllowKills = &b
	}
	if v, ok := m["are_approvals_required"]; ok {
		b := v.(bool)
		cp.AreApprovalsRequired = &b
	}
	if v, ok := m["are_approvers_restricted"]; ok {
		b := v.(bool)
		cp.AreApproversRestricted = &b
	}
	if v, ok := m["approvers"]; ok {
		cp.Approvers = expandPermissionEntities(v.([]interface{}))
	}
	if v, ok := m["approval_skippable_by"]; ok {
		cp.ApprovalSkippableBy = expandPermissionEntities(v.([]interface{}))
	}
	return cp
}

func expandPermissionEntities(raw []interface{}) []splitsdk.PermissionEntity {
	if len(raw) == 0 {
		return nil
	}
	out := make([]splitsdk.PermissionEntity, 0, len(raw))
	for _, item := range raw {
		m := item.(map[string]interface{})
		out = append(out, splitsdk.PermissionEntity{
			ID:   m["id"].(string),
			Name: m["name"].(string),
			Type: m["type"].(string),
		})
	}
	return out
}

func flattenChangePermissions(cp *splitsdk.ChangePermissions) []interface{} {
	if cp == nil {
		return nil
	}
	m := map[string]interface{}{}
	if cp.AllowKills != nil {
		m["allow_kills"] = *cp.AllowKills
	}
	if cp.AreApprovalsRequired != nil {
		m["are_approvals_required"] = *cp.AreApprovalsRequired
	}
	if cp.AreApproversRestricted != nil {
		m["are_approvers_restricted"] = *cp.AreApproversRestricted
	}
	m["approvers"] = flattenPermissionEntities(cp.Approvers)
	m["approval_skippable_by"] = flattenPermissionEntities(cp.ApprovalSkippableBy)
	return []interface{}{m}
}

func flattenPermissionEntities(entities []splitsdk.PermissionEntity) []interface{} {
	if len(entities) == 0 {
		return []interface{}{}
	}
	out := make([]interface{}, 0, len(entities))
	for _, e := range entities {
		out = append(out, map[string]interface{}{
			"id":   e.ID,
			"name": e.Name,
			"type": e.Type,
		})
	}
	return out
}
