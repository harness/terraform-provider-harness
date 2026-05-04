package ansible_inventory

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

const (
	inventoryTypeManual  = "manual"
	inventoryTypeDynamic = "dynamic"
	inventoryTypePlugin  = "plugin"
)

func ResourceAnsibleInventory() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for managing Harness IaCM Ansible Inventories.",

		ReadContext:   resourceAnsibleInventoryRead,
		DeleteContext: resourceAnsibleInventoryDelete,
		CreateContext: resourceAnsibleInventoryCreate,
		UpdateContext: resourceAnsibleInventoryUpdate,
		Importer:      helpers.ProjectResourceImporter,

		Schema: inventorySchema(false),
	}
	resource.Schema["tags"] = helpers.GetTagsSchema(helpers.SchemaFlagTypes.Optional)
	helpers.SetProjectLevelResourceSchema(resource.Schema)
	return resource
}

func inventorySchema(dataSource bool) map[string]*schema.Schema {
	inventoryType := &schema.Schema{
		Description:  "Type of inventory. One of: manual, dynamic, plugin.",
		Type:         schema.TypeString,
		Required:     true,
		ForceNew:     true,
		ValidateFunc: validation.StringInSlice([]string{inventoryTypeManual, inventoryTypeDynamic, inventoryTypePlugin}, false),
	}
	if dataSource {
		inventoryType.Required = false
		inventoryType.Computed = true
		inventoryType.ValidateFunc = nil
	}

	s := map[string]*schema.Schema{
		"type": inventoryType,
		"groups": {
			Description: "Manual groups used when type is manual.",
			Type:        schema.TypeSet,
			Optional:    !dataSource,
			Computed:    dataSource,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"identifier": {
						Description: "Identifier of the group.",
						Type:        schema.TypeString,
						Required:    !dataSource,
						Computed:    dataSource,
					},
					"name": {
						Description: "Name of the group.",
						Type:        schema.TypeString,
						Required:    !dataSource,
						Computed:    dataSource,
					},
					"hosts": {
						Description: "List of hosts in the group.",
						Type:        schema.TypeList,
						Optional:    !dataSource,
						Computed:    dataSource,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
					},
					"vars": {
						Description: "Variables for the group.",
						Type:        schema.TypeSet,
						Optional:    !dataSource,
						Computed:    dataSource,
						Elem:        ansibleVariableSchema(dataSource),
					},
				},
			},
		},
		"dynamic_groups": {
			Description: "Dynamic groups used when type is dynamic.",
			Type:        schema.TypeSet,
			Optional:    !dataSource,
			Computed:    dataSource,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"identifier": {
						Description: "Identifier of the dynamic group.",
						Type:        schema.TypeString,
						Required:    !dataSource,
						Computed:    dataSource,
					},
					"name": {
						Description: "Name of the dynamic group.",
						Type:        schema.TypeString,
						Required:    !dataSource,
						Computed:    dataSource,
					},
					"connector_identifier": {
						Description: "Connector identifier used by the dynamic group.",
						Type:        schema.TypeString,
						Required:    !dataSource,
						Computed:    dataSource,
					},
					"connector_type": {
						Description: "Connector type (e.g. workspace, aws, gcp, azure, vault).",
						Type:        schema.TypeString,
						Required:    !dataSource,
						Computed:    dataSource,
					},
					"configuration": {
						Description: "Configuration for the dynamic group.",
						Type:        schema.TypeList,
						Optional:    !dataSource,
						Computed:    dataSource,
						MaxItems:    maxItemsOne(dataSource),
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"resource_type": {
									Description: "Resource type to select.",
									Type:        schema.TypeString,
									Optional:    !dataSource,
									Computed:    dataSource,
								},
								"host_address_attribute": {
									Description: "Host address attribute.",
									Type:        schema.TypeString,
									Optional:    !dataSource,
									Computed:    dataSource,
								},
								"selectors": {
									Description: "Resource selectors.",
									Type:        schema.TypeList,
									Optional:    !dataSource,
									Computed:    dataSource,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"attribute": {
												Description: "Attribute to filter on.",
												Type:        schema.TypeString,
												Required:    !dataSource,
												Computed:    dataSource,
											},
											"operator": {
												Description: "Operator for the filter.",
												Type:        schema.TypeString,
												Required:    !dataSource,
												Computed:    dataSource,
											},
											"value": {
												Description: "Value for the filter.",
												Type:        schema.TypeString,
												Required:    !dataSource,
												Computed:    dataSource,
											},
										},
									},
								},
							},
						},
					},
					"vars": {
						Description: "Variables for the dynamic group.",
						Type:        schema.TypeSet,
						Optional:    !dataSource,
						Computed:    dataSource,
						Elem:        ansibleVariableSchema(dataSource),
					},
					"dynamic_vars": {
						Description: "Dynamic variables for the dynamic group.",
						Type:        schema.TypeSet,
						Optional:    !dataSource,
						Computed:    dataSource,
						Elem:        ansibleVariableSchema(dataSource),
					},
				},
			},
		},
		"plugin_options": {
			Description: "Plugin options used when type is plugin.",
			Type:        schema.TypeList,
			Optional:    !dataSource,
			Computed:    dataSource,
			MaxItems:    maxItemsOne(dataSource),
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"source_type": {
						Description: "Source type for plugin inventory. One of: git, inline.",
						Type:        schema.TypeString,
						Optional:    !dataSource,
						Computed:    dataSource,
					},
					"inline_yaml": {
						Description: "Inline plugin inventory YAML content (when source_type is inline).",
						Type:        schema.TypeString,
						Optional:    !dataSource,
						Computed:    dataSource,
					},
					"provider_connector_identifier": {
						Description: "Provider connector identifier for plugin inventory execution.",
						Type:        schema.TypeString,
						Optional:    !dataSource,
						Computed:    dataSource,
					},
					"provider_connector_type": {
						Description: "Provider connector type.",
						Type:        schema.TypeString,
						Optional:    !dataSource,
						Computed:    dataSource,
					},
					"repository": {
						Description: "Git repository name (when source_type is git).",
						Type:        schema.TypeString,
						Optional:    !dataSource,
						Computed:    dataSource,
					},
					"repository_branch": {
						Description: "Git branch.",
						Type:        schema.TypeString,
						Optional:    !dataSource,
						Computed:    dataSource,
					},
					"repository_commit": {
						Description: "Git commit or tag.",
						Type:        schema.TypeString,
						Optional:    !dataSource,
						Computed:    dataSource,
					},
					"repository_connector": {
						Description: "Repository connector reference (when source_type is git).",
						Type:        schema.TypeString,
						Optional:    !dataSource,
						Computed:    dataSource,
					},
					"repository_path": {
						Description: "Path within the repository to the plugin inventory YAML.",
						Type:        schema.TypeString,
						Optional:    !dataSource,
						Computed:    dataSource,
					},
				},
			},
		},
		"vars": {
			Description: "Variables configured on the inventory.",
			Type:        schema.TypeSet,
			Optional:    !dataSource,
			Computed:    dataSource,
			Elem:        ansibleVariableSchema(dataSource),
		},
	}
	return s
}

// maxItemsOne returns the MaxItems value for a TypeList. MaxItems must be 0
// when the schema is computed-only (data source) - the Terraform SDK rejects
// MaxItems on computed-only fields.
func maxItemsOne(dataSource bool) int {
	if dataSource {
		return 0
	}
	return 1
}

func ansibleVariableSchema(dataSource bool) *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"key": {
				Description: "Key is the identifier for the variable.",
				Type:        schema.TypeString,
				Required:    !dataSource,
				Computed:    dataSource,
			},
			"value": {
				Description: "Value of the variable. For secret value types this must be a Harness secret reference.",
				Type:        schema.TypeString,
				Required:    !dataSource,
				Computed:    dataSource,
			},
			"value_type": {
				Description: "Value type. One of: string, secret.",
				Type:        schema.TypeString,
				Required:    !dataSource,
				Computed:    dataSource,
			},
			"file_name": {
				Description: "Filename to store the value in (used for file-backed variables).",
				Type:        schema.TypeString,
				Optional:    !dataSource,
				Computed:    dataSource,
			},
		},
	}
}

func resourceAnsibleInventoryRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	resp, httpResp, err := c.AnsibleApi.AnsibleShowInventory(
		ctx,
		d.Get("org_id").(string),
		d.Get("project_id").(string),
		d.Get("identifier").(string),
		c.AccountId,
		nil,
	)
	if err != nil {
		return helpers.HandleReadApiError(err, d, httpResp)
	}

	return readInventory(d, &resp)
}

func resourceAnsibleInventoryDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	if d.Id() == "" {
		return nil
	}

	httpResp, err := c.AnsibleApi.AnsibleDeleteInventory(
		ctx,
		d.Get("org_id").(string),
		d.Get("project_id").(string),
		d.Get("identifier").(string),
		c.AccountId,
	)
	if err != nil {
		return parseError(err, httpResp)
	}
	return nil
}

func resourceAnsibleInventoryCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	body, err := buildCreateInventory(d)
	if err != nil {
		return diag.Errorf("%s", err.Error())
	}

	_, httpResp, err := c.AnsibleApi.AnsibleCreateInventory(
		ctx,
		body,
		c.AccountId,
		d.Get("org_id").(string),
		d.Get("project_id").(string),
	)
	if err != nil {
		return parseError(err, httpResp)
	}

	d.SetId(body.Identifier)
	return resourceAnsibleInventoryRead(ctx, d, meta)
}

func resourceAnsibleInventoryUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	scope := inventoryScope{
		account: c.AccountId,
		org:     d.Get("org_id").(string),
		project: d.Get("project_id").(string),
		invID:   d.Get("identifier").(string),
	}

	switch d.Get("type").(string) {
	case inventoryTypeManual:
		if diags := reconcileManualInventory(ctx, c, d, scope); diags != nil {
			return diags
		}
	case inventoryTypeDynamic:
		if diags := reconcileDynamicInventory(ctx, c, d, scope); diags != nil {
			return diags
		}
	case inventoryTypePlugin:
		if diags := reconcilePluginInventory(ctx, c, d, scope); diags != nil {
			return diags
		}
	}

	if d.HasChanges("name", "tags") {
		if diags := updateInventoryTopLevel(ctx, c, d, scope); diags != nil {
			return diags
		}
	}

	return resourceAnsibleInventoryRead(ctx, d, meta)
}

func buildCreateInventory(d *schema.ResourceData) (nextgen.CreateInventoryRequest, error) {
	req := nextgen.CreateInventoryRequest{
		Identifier:    d.Get("identifier").(string),
		Name:          d.Get("name").(string),
		Type_:         d.Get("type").(string),
		Groups:        buildManualGroups(d),
		DynamicGroups: buildDynamicGroups(d),
		PluginOptions: buildPluginOptions(d),
		Vars:          buildVariableList(d, "vars"),
	}
	if tagAttr, ok := d.GetOk("tags"); ok {
		tags, err := marshalTags(tagAttr)
		if err != nil {
			return req, err
		}
		req.Tags = tags
	}
	return req, nil
}

func buildManualGroups(d *schema.ResourceData) []nextgen.CreateManualGroup {
	raw, ok := d.GetOk("groups")
	if !ok {
		return nil
	}
	var out []nextgen.CreateManualGroup
	for _, item := range raw.(*schema.Set).List() {
		m := item.(map[string]interface{})
		g := nextgen.CreateManualGroup{
			Identifier: m["identifier"].(string),
			Name:       m["name"].(string),
			Vars:       variablesFromSet(m["vars"]),
		}
		if hostsRaw, ok := m["hosts"].([]interface{}); ok {
			for _, h := range hostsRaw {
				g.Hosts = append(g.Hosts, h.(string))
			}
		}
		out = append(out, g)
	}
	return out
}

func buildDynamicGroups(d *schema.ResourceData) []nextgen.CreateDynamicGroup {
	raw, ok := d.GetOk("dynamic_groups")
	if !ok {
		return nil
	}
	var out []nextgen.CreateDynamicGroup
	for _, item := range raw.(*schema.Set).List() {
		m := item.(map[string]interface{})
		g := nextgen.CreateDynamicGroup{
			Identifier:          m["identifier"].(string),
			Name:                m["name"].(string),
			ConnectorIdentifier: m["connector_identifier"].(string),
			ConnectorType:       m["connector_type"].(string),
			Vars:                variablesFromSet(m["vars"]),
			DynamicVars:         variablesFromSet(m["dynamic_vars"]),
		}
		if cfgList, ok := m["configuration"].([]interface{}); ok && len(cfgList) > 0 && cfgList[0] != nil {
			cfgMap := cfgList[0].(map[string]interface{})
			cfg := &nextgen.DynamicConfig{
				ResourceType:         cfgMap["resource_type"].(string),
				HostAddressAttribute: cfgMap["host_address_attribute"].(string),
			}
			if selRaw, ok := cfgMap["selectors"].([]interface{}); ok {
				for _, s := range selRaw {
					sm := s.(map[string]interface{})
					cfg.Selectors = append(cfg.Selectors, nextgen.ResourcesSelectorConditions{
						Attribute: sm["attribute"].(string),
						Operator:  sm["operator"].(string),
						Value:     sm["value"].(string),
					})
				}
			}
			g.Configuration = cfg
		}
		out = append(out, g)
	}
	return out
}

func buildPluginOptions(d *schema.ResourceData) *nextgen.PluginOptions {
	raw, ok := d.GetOk("plugin_options")
	if !ok {
		return nil
	}
	list := raw.([]interface{})
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	m := list[0].(map[string]interface{})
	return &nextgen.PluginOptions{
		SourceType:                  m["source_type"].(string),
		InlineYaml:                  m["inline_yaml"].(string),
		ProviderConnectorIdentifier: m["provider_connector_identifier"].(string),
		ProviderConnectorType:       m["provider_connector_type"].(string),
		Repository:                  m["repository"].(string),
		RepositoryBranch:            m["repository_branch"].(string),
		RepositoryCommit:            m["repository_commit"].(string),
		RepositoryConnector:         m["repository_connector"].(string),
		RepositoryPath:              m["repository_path"].(string),
	}
}

func buildVariableList(d *schema.ResourceData, attr string) []nextgen.AnsibleVariable {
	raw, ok := d.GetOk(attr)
	if !ok {
		return nil
	}
	return variablesFromSet(raw)
}

func variablesFromSet(v interface{}) []nextgen.AnsibleVariable {
	if v == nil {
		return nil
	}
	set, ok := v.(*schema.Set)
	if !ok {
		return nil
	}
	var out []nextgen.AnsibleVariable
	for _, item := range set.List() {
		m := item.(map[string]interface{})
		out = append(out, nextgen.AnsibleVariable{
			Key:       m["key"].(string),
			Value:     m["value"].(string),
			ValueType: m["value_type"].(string),
			FileName:  m["file_name"].(string),
		})
	}
	return out
}

func readInventory(d *schema.ResourceData, resp *nextgen.ShowInventoryResponse) diag.Diagnostics {
	d.SetId(resp.Identifier)
	d.Set("identifier", resp.Identifier)
	d.Set("org_id", resp.Org)
	d.Set("project_id", resp.Project)
	d.Set("name", resp.Name)
	d.Set("type", resp.Type_)
	if resp.Tags != "" {
		tags, err := unmarshalTags(resp.Tags)
		if err != nil {
			return diag.Errorf("failed to decode inventory tags: %s", err.Error())
		}
		d.Set("tags", tags)
	}

	if len(resp.Data) == 0 {
		return nil
	}

	switch resp.Type_ {
	case inventoryTypeManual:
		var inv nextgen.ManualInventory
		if err := json.Unmarshal(resp.Data, &inv); err != nil {
			return diag.Errorf("failed to decode manual inventory data: %s", err.Error())
		}
		d.Set("groups", flattenManualGroups(inv.Groups))
		d.Set("vars", flattenVariableMap(inv.Vars))
	case inventoryTypeDynamic:
		var inv nextgen.DynamicInventory
		if err := json.Unmarshal(resp.Data, &inv); err != nil {
			return diag.Errorf("failed to decode dynamic inventory data: %s", err.Error())
		}
		d.Set("dynamic_groups", flattenDynamicGroups(inv.Groups))
		d.Set("vars", flattenVariableMap(inv.Vars))
	case inventoryTypePlugin:
		var inv nextgen.PluginInventory
		if err := json.Unmarshal(resp.Data, &inv); err != nil {
			return diag.Errorf("failed to decode plugin inventory data: %s", err.Error())
		}
		d.Set("plugin_options", flattenPluginOptions(inv.PluginOptions))
		d.Set("vars", flattenVariableMap(inv.Vars))
	default:
		return diag.Errorf("unsupported inventory type: %s", resp.Type_)
	}
	return nil
}

func flattenManualGroups(groups map[string]nextgen.AnsibleGroup) []interface{} {
	if len(groups) == 0 {
		return nil
	}
	out := make([]interface{}, 0, len(groups))
	for id, g := range groups {
		hosts := make([]interface{}, 0, len(g.Hosts))
		for h := range g.Hosts {
			hosts = append(hosts, h)
		}
		out = append(out, map[string]interface{}{
			"identifier": id,
			"name":       id,
			"hosts":      hosts,
			"vars":       flattenVariableMap(g.Vars),
		})
	}
	return out
}

func flattenDynamicGroups(groups []nextgen.DynamicGroup) []interface{} {
	if len(groups) == 0 {
		return nil
	}
	out := make([]interface{}, 0, len(groups))
	for _, g := range groups {
		entry := map[string]interface{}{
			"identifier":           g.Identifier,
			"name":                 g.Identifier,
			"connector_identifier": g.ConnectorIdentifier,
			"connector_type":       g.ConnectorType,
			"vars":                 flattenVariableMap(g.Vars),
			"dynamic_vars":         flattenVariableMap(g.DynamicVars),
		}
		if g.Configuration != nil {
			selectors := make([]interface{}, 0, len(g.Configuration.Selectors))
			for _, s := range g.Configuration.Selectors {
				selectors = append(selectors, map[string]interface{}{
					"attribute": s.Attribute,
					"operator":  s.Operator,
					"value":     s.Value,
				})
			}
			entry["configuration"] = []interface{}{
				map[string]interface{}{
					"resource_type":          g.Configuration.ResourceType,
					"host_address_attribute": g.Configuration.HostAddressAttribute,
					"selectors":              selectors,
				},
			}
		}
		out = append(out, entry)
	}
	return out
}

func flattenPluginOptions(opts *nextgen.PluginOptions) []interface{} {
	if opts == nil {
		return nil
	}
	return []interface{}{
		map[string]interface{}{
			"source_type":                   opts.SourceType,
			"inline_yaml":                   opts.InlineYaml,
			"provider_connector_identifier": opts.ProviderConnectorIdentifier,
			"provider_connector_type":       opts.ProviderConnectorType,
			"repository":                    opts.Repository,
			"repository_branch":             opts.RepositoryBranch,
			"repository_commit":             opts.RepositoryCommit,
			"repository_connector":          opts.RepositoryConnector,
			"repository_path":               opts.RepositoryPath,
		},
	}
}

func flattenVariableMap(vars map[string]nextgen.AnsibleVariable) []interface{} {
	if len(vars) == 0 {
		return nil
	}
	out := make([]interface{}, 0, len(vars))
	for _, v := range vars {
		out = append(out, map[string]interface{}{
			"key":        v.Key,
			"value":      v.Value,
			"value_type": v.ValueType,
			"file_name":  v.FileName,
		})
	}
	return out
}

// marshalTags renders the provider's tag set schema as the JSON map format
// expected by the IaCM ansible endpoints (wire type is a single string).
func marshalTags(v interface{}) (string, error) {
	set, ok := v.(*schema.Set)
	if !ok {
		return "", nil
	}
	tags := helpers.ExpandTags(set.List())
	if len(tags) == 0 {
		return "", nil
	}
	b, err := json.Marshal(tags)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func unmarshalTags(data string) ([]interface{}, error) {
	if data == "" {
		return nil, nil
	}
	var m map[string]string
	if err := json.Unmarshal([]byte(data), &m); err == nil {
		out := make([]interface{}, 0, len(m))
		for _, t := range helpers.FlattenTags(m) {
			out = append(out, t)
		}
		return out, nil
	}
	var list []string
	if err := json.Unmarshal([]byte(data), &list); err == nil {
		out := make([]interface{}, 0, len(list))
		for _, t := range list {
			out = append(out, t)
		}
		return out, nil
	}
	return nil, nil
}

func parseError(err error, httpResp *http.Response) diag.Diagnostics {
	if httpResp != nil && httpResp.StatusCode == 401 {
		return diag.Errorf("%s\nHint:\n1) Please check if token has expired or is wrong.\n2) Harness Provider is misconfigured.", httpResp.Status)
	}
	if httpResp != nil && httpResp.StatusCode == 403 {
		return diag.Errorf("%s\nHint:\n1) Please check if the token has required permission for this operation.\n2) Please check if the token has expired or is wrong.", httpResp.Status)
	}

	se, ok := err.(nextgen.GenericSwaggerError)
	if !ok {
		return diag.FromErr(err)
	}

	iacmErr := nextgen.IacmError{}
	if jsonErr := json.Unmarshal(se.Body(), &iacmErr); jsonErr == nil && iacmErr.Message != "" {
		if httpResp != nil {
			return diag.Errorf("%s\nHint:\n1) %s", httpResp.Status, iacmErr.Message)
		}
		return diag.Errorf("%s", iacmErr.Message)
	}
	return diag.Errorf("%s", err.Error())
}
