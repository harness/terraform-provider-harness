package ansible_inventory

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type inventoryScope struct {
	account string
	org     string
	project string
	invID   string
}

func reconcileManualInventory(ctx context.Context, c *nextgen.APIClient, d *schema.ResourceData, s inventoryScope) diag.Diagnostics {
	oldG, newG := changedGroupSets(d, "groups")
	oldMap := manualGroupsByID(oldG)
	newMap := manualGroupsByID(newG)

	// Phase A: delete removed hosts (hosts of groups still present only — groups being removed
	// are deleted as a whole in Phase B and the server cascades their hosts).
	for gid, og := range oldMap {
		ng, stillExists := newMap[gid]
		if !stillExists {
			continue
		}
		for _, host := range og.hosts {
			if !containsString(ng.hosts, host) {
				if httpResp, err := c.AnsibleApi.AnsibleDeleteInventoryHost(
					ctx, s.org, s.project, s.invID, host, s.account,
				); err != nil {
					return parseError(fmt.Errorf("delete host %s/%s: %w", gid, host, err), httpResp)
				}
			}
		}
	}

	// Phase B: delete removed groups.
	for gid := range oldMap {
		if _, ok := newMap[gid]; !ok {
			if httpResp, err := c.AnsibleApi.AnsibleDeleteInventoryGroup(
				ctx, s.org, s.project, s.invID, gid, s.account,
			); err != nil {
				return parseError(fmt.Errorf("delete group %s: %w", gid, err), httpResp)
			}
		}
	}

	// Phase C: create new groups (no hosts yet — hosts added in phase D).
	for gid, ng := range newMap {
		if _, existed := oldMap[gid]; existed {
			continue
		}
		body := nextgen.CreateInventoryGroupRequest{
			Identifier: gid,
			Name:       ng.name,
			Vars:       ng.vars,
		}
		if httpResp, err := c.AnsibleApi.AnsibleCreateInventoryGroup(
			ctx, body, s.account, s.org, s.project, s.invID,
		); err != nil {
			return parseError(fmt.Errorf("create group %s: %w", gid, err), httpResp)
		}
	}

	// Phase D: add new hosts to both newly-created and pre-existing groups.
	for gid, ng := range newMap {
		og, existed := oldMap[gid]
		for _, host := range ng.hosts {
			if existed && containsString(og.hosts, host) {
				continue
			}
			body := nextgen.CreateInventoryHostRequest{
				HostAddress: host,
				Groups:      []string{gid},
			}
			if httpResp, err := c.AnsibleApi.AnsibleCreateInventoryHost(
				ctx, body, s.account, s.org, s.project, s.invID,
			); err != nil {
				return parseError(fmt.Errorf("create host %s/%s: %w", gid, host, err), httpResp)
			}
		}
	}

	// Phase E: update groups whose vars/hosts changed.
	for gid, ng := range newMap {
		og, existed := oldMap[gid]
		if !existed {
			continue
		}
		if manualGroupEqual(og, ng) {
			continue
		}
		body := nextgen.UpdateInventoryGroupRequest{
			Name:  ng.name,
			Vars:  ng.vars,
			Hosts: ng.hosts,
		}
		if httpResp, err := c.AnsibleApi.AnsibleUpdateInventoryGroup(
			ctx, body, s.account, s.org, s.project, s.invID, gid,
		); err != nil {
			return parseError(fmt.Errorf("update group %s: %w", gid, err), httpResp)
		}
	}

	// Phase F: inventory-level vars (single idempotent PUT).
	if d.HasChange("vars") {
		if diags := updateInventoryVars(ctx, c, d, s); diags != nil {
			return diags
		}
	}

	return nil
}

func reconcileDynamicInventory(ctx context.Context, c *nextgen.APIClient, d *schema.ResourceData, s inventoryScope) diag.Diagnostics {
	oldG, newG := changedGroupSets(d, "dynamic_groups")
	oldMap := dynamicGroupsByID(oldG)
	newMap := dynamicGroupsByID(newG)

	// Delete removed groups.
	for gid := range oldMap {
		if _, ok := newMap[gid]; !ok {
			if httpResp, err := c.AnsibleApi.AnsibleDeleteInventoryGroup(
				ctx, s.org, s.project, s.invID, gid, s.account,
			); err != nil {
				return parseError(fmt.Errorf("delete group %s: %w", gid, err), httpResp)
			}
		}
	}

	// Create new groups.
	for gid, ng := range newMap {
		if _, existed := oldMap[gid]; existed {
			continue
		}
		body := nextgen.CreateInventoryGroupRequest{
			Identifier:          gid,
			Name:                ng.name,
			ConnectorIdentifier: ng.connectorIdentifier,
			ConnectorType:       ng.connectorType,
			Configuration:       ng.configuration,
			Vars:                ng.vars,
			DynamicVars:         ng.dynamicVars,
		}
		if httpResp, err := c.AnsibleApi.AnsibleCreateInventoryGroup(
			ctx, body, s.account, s.org, s.project, s.invID,
		); err != nil {
			return parseError(fmt.Errorf("create group %s: %w", gid, err), httpResp)
		}
	}

	// Update changed groups.
	for gid, ng := range newMap {
		og, existed := oldMap[gid]
		if !existed {
			continue
		}
		if dynamicGroupEqual(og, ng) {
			continue
		}
		body := nextgen.UpdateInventoryGroupRequest{
			Name:                ng.name,
			ConnectorIdentifier: ng.connectorIdentifier,
			ConnectorType:       ng.connectorType,
			Configuration:       ng.configuration,
			Vars:                ng.vars,
			DynamicVars:         ng.dynamicVars,
		}
		if httpResp, err := c.AnsibleApi.AnsibleUpdateInventoryGroup(
			ctx, body, s.account, s.org, s.project, s.invID, gid,
		); err != nil {
			return parseError(fmt.Errorf("update group %s: %w", gid, err), httpResp)
		}
	}

	// Inventory-level vars.
	if d.HasChange("vars") {
		if diags := updateInventoryVars(ctx, c, d, s); diags != nil {
			return diags
		}
	}

	return nil
}

func reconcilePluginInventory(ctx context.Context, c *nextgen.APIClient, d *schema.ResourceData, s inventoryScope) diag.Diagnostics {
	if d.HasChange("plugin_options") {
		body := nextgen.UpdateInventoryPluginRequest{
			PluginOptions: buildPluginOptions(d),
		}
		if httpResp, err := c.AnsibleApi.AnsibleUpdateInventoryPlugin(
			ctx, body, s.account, s.org, s.project, s.invID,
		); err != nil {
			return parseError(fmt.Errorf("update plugin options: %w", err), httpResp)
		}
	}

	if d.HasChange("vars") {
		if diags := updateInventoryVars(ctx, c, d, s); diags != nil {
			return diags
		}
	}

	return nil
}

func updateInventoryVars(ctx context.Context, c *nextgen.APIClient, d *schema.ResourceData, s inventoryScope) diag.Diagnostics {
	body := nextgen.UpdateInventoryVarsRequest{
		Vars: toInventoryVarCreate(buildVariableList(d, "vars")),
	}
	if body.Vars == nil {
		body.Vars = []nextgen.HarnessIacmInventoryVarCreate{}
	}
	if httpResp, err := c.AnsibleApi.AnsibleUpdateInventoryVars(
		ctx, body, s.account, s.org, s.project, s.invID,
	); err != nil {
		return parseError(fmt.Errorf("update vars: %w", err), httpResp)
	}
	return nil
}

// updateInventoryTopLevel pushes name/description/tags changes through the
// UpdateInventory endpoint. That endpoint is a PUT and the service writes name,
// description, data and tags unconditionally, so the current data document has
// to be carried through the request or the update would blank it. The document
// is read back from the show endpoint rather than rebuilt from state because the
// group/host/vars reconcile has already run by this point, which makes the
// stored document - not the config - the authoritative copy.
func updateInventoryTopLevel(ctx context.Context, c *nextgen.APIClient, d *schema.ResourceData, s inventoryScope) diag.Diagnostics {
	current, httpResp, err := c.AnsibleApi.AnsibleShowInventory(ctx, s.org, s.project, s.invID, s.account, nil)
	if err != nil {
		return parseError(fmt.Errorf("read inventory before update: %w", err), httpResp)
	}

	if httpResp, err := c.AnsibleApi.AnsibleUpdateInventory(
		ctx, buildUpdateInventory(d, current.Data), s.account, s.org, s.project, s.invID,
	); err != nil {
		return parseError(err, httpResp)
	}
	return nil
}

// buildUpdateInventory assembles the update body, carrying the inventory's
// current data document through unchanged.
func buildUpdateInventory(d *schema.ResourceData, current json.RawMessage) nextgen.UpdateInventoryRequest {
	// An absent or empty data field is rejected as "invalid inventory data", so an
	// inventory whose document is empty still needs a JSON object to update against.
	if len(current) == 0 {
		current = json.RawMessage("{}")
	}
	return nextgen.UpdateInventoryRequest{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Tags:        expandTags(d),
		Data:        current,
	}
}

// --- diff helpers ---

type manualGroup struct {
	name  string
	hosts []string
	vars  []nextgen.AnsibleVariable
}

type dynamicGroupSnapshot struct {
	name                string
	connectorIdentifier string
	connectorType       string
	configuration       *nextgen.DynamicConfig
	vars                []nextgen.AnsibleVariable
	dynamicVars         []nextgen.AnsibleVariable
}

func changedGroupSets(d *schema.ResourceData, key string) ([]interface{}, []interface{}) {
	oldRaw, newRaw := d.GetChange(key)
	return setAsList(oldRaw), setAsList(newRaw)
}

func setAsList(v interface{}) []interface{} {
	if v == nil {
		return nil
	}
	if s, ok := v.(*schema.Set); ok {
		return s.List()
	}
	return nil
}

func manualGroupsByID(list []interface{}) map[string]manualGroup {
	out := make(map[string]manualGroup, len(list))
	for _, item := range list {
		m := item.(map[string]interface{})
		id := m["identifier"].(string)
		g := manualGroup{
			name: m["name"].(string),
			vars: variablesFromSet(m["vars"]),
		}
		if hostsRaw, ok := m["hosts"].([]interface{}); ok {
			for _, h := range hostsRaw {
				g.hosts = append(g.hosts, h.(string))
			}
		}
		out[id] = g
	}
	return out
}

func dynamicGroupsByID(list []interface{}) map[string]dynamicGroupSnapshot {
	out := make(map[string]dynamicGroupSnapshot, len(list))
	for _, item := range list {
		m := item.(map[string]interface{})
		id := m["identifier"].(string)
		g := dynamicGroupSnapshot{
			name:                m["name"].(string),
			connectorIdentifier: m["connector_identifier"].(string),
			connectorType:       m["connector_type"].(string),
			vars:                variablesFromSet(m["vars"]),
			dynamicVars:         variablesFromSet(m["dynamic_vars"]),
		}
		if cfgList, ok := m["configuration"].([]interface{}); ok && len(cfgList) > 0 && cfgList[0] != nil {
			cfgMap := cfgList[0].(map[string]interface{})
			cfg := &nextgen.DynamicConfig{
				ResourceType:         cfgMap["resource_type"].(string),
				HostAddressAttribute: cfgMap["host_address_attribute"].(string),
			}
			if selRaw, ok := cfgMap["selectors"].([]interface{}); ok {
				for _, sel := range selRaw {
					sm := sel.(map[string]interface{})
					cfg.Selectors = append(cfg.Selectors, nextgen.ResourcesSelectorConditions{
						Attribute: sm["attribute"].(string),
						Operator:  sm["operator"].(string),
						Value:     sm["value"].(string),
					})
				}
			}
			g.configuration = cfg
		}
		out[id] = g
	}
	return out
}

func manualGroupEqual(a, b manualGroup) bool {
	if a.name != b.name {
		return false
	}
	if !stringSliceSetEqual(a.hosts, b.hosts) {
		return false
	}
	return variablesEqual(a.vars, b.vars)
}

func dynamicGroupEqual(a, b dynamicGroupSnapshot) bool {
	if a.name != b.name ||
		a.connectorIdentifier != b.connectorIdentifier ||
		a.connectorType != b.connectorType {
		return false
	}
	if !reflect.DeepEqual(a.configuration, b.configuration) {
		return false
	}
	if !variablesEqual(a.vars, b.vars) {
		return false
	}
	return variablesEqual(a.dynamicVars, b.dynamicVars)
}

func variablesEqual(a, b []nextgen.AnsibleVariable) bool {
	if len(a) != len(b) {
		return false
	}
	am := make(map[string]nextgen.AnsibleVariable, len(a))
	for _, v := range a {
		am[v.Key] = v
	}
	for _, v := range b {
		av, ok := am[v.Key]
		if !ok {
			return false
		}
		if av.Value != v.Value || av.ValueType != v.ValueType || av.FileName != v.FileName {
			return false
		}
	}
	return true
}

func stringSliceSetEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	seen := make(map[string]int, len(a))
	for _, v := range a {
		seen[v]++
	}
	for _, v := range b {
		if seen[v] == 0 {
			return false
		}
		seen[v]--
	}
	return true
}

func containsString(haystack []string, needle string) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}
	return false
}

func toInventoryVarCreate(vars []nextgen.AnsibleVariable) []nextgen.HarnessIacmInventoryVarCreate {
	if len(vars) == 0 {
		return nil
	}
	out := make([]nextgen.HarnessIacmInventoryVarCreate, 0, len(vars))
	for _, v := range vars {
		out = append(out, nextgen.HarnessIacmInventoryVarCreate{
			Key:       v.Key,
			Value:     v.Value,
			ValueType: v.ValueType,
			FileName:  v.FileName,
		})
	}
	return out
}
