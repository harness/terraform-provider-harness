package loadtest

import (
	"context"
	"fmt"
	"strings"

	sdkloadtest "github.com/harness/harness-go-sdk/harness/loadtest"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceChaosLoadTestTemplateRevision() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for one revision of a Harness load test hub template.\n\n" +
			"Creates, reads, and deletes that revision. Load Test Manager does not update a revision in place. " +
			"harness_chaos_load_test_template updates a revision only while that resource's revision is still the latest. " +
			"`revision` must be a real revision id. `latest` is reserved for syncing a load test. " +
			"Fields not set are inherited from the previous revision. An omitted inherited field keeps the stored value and does not create another revision. " +
			"Leave org_id and project_id empty for an account hub. Set org_id and leave project_id empty for an organization hub. " +
			"Delete is refused when a REFERENCE load test still pins this revision.",
		CreateContext: resourceChaosLoadTestTemplateRevisionCreate,
		ReadContext:   resourceChaosLoadTestTemplateRevisionRead,
		DeleteContext: resourceChaosLoadTestTemplateRevisionDelete,
		CustomizeDiff: resourceChaosLoadTestTemplateRevisionCustomizeDiff,
		Importer:      &schema.ResourceImporter{StateContext: importChaosLoadTestTemplateRevision},
		Schema:        resourceChaosLoadTestTemplateRevisionSchema(),
	}
}

func DataSourceChaosLoadTestTemplateRevision() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for one revision of a Harness load test hub template.",
		ReadContext: dataSourceChaosLoadTestTemplateRevisionRead,
		Schema:      dataSourceChaosLoadTestTemplateRevisionSchema(),
	}
}

func resourceChaosLoadTestTemplateRevisionSchema() map[string]*schema.Schema {
	variable := templateVariableSchema()
	variable.ForceNew = true
	variable.Computed = true
	variable.Description = "Custom variables on this revision. Omit to inherit the previous revision."
	return map[string]*schema.Schema{
		"org_id":            {Description: "Organization identifier. Empty for an account hub.", Type: schema.TypeString, Optional: true, ForceNew: true},
		"project_id":        {Description: "Project identifier. Empty for an account or organization hub. Requires org_id.", Type: schema.TypeString, Optional: true, ForceNew: true},
		"hub_identity":      {Description: "Chaos hub that holds the template.", Type: schema.TypeString, Required: true, ForceNew: true, ValidateFunc: validateLoadTestIdentifier},
		"template_identity": {Description: "Template identifier.", Type: schema.TypeString, Required: true, ForceNew: true, ValidateFunc: validateLoadTestIdentifier},
		"revision":          {Description: "Revision id to create. `latest` is rejected.", Type: schema.TypeString, Required: true, ForceNew: true, ValidateFunc: validation.StringIsNotEmpty},
		"name":              {Description: "Name stored on this revision. Omit to inherit the previous revision.", Type: schema.TypeString, Optional: true, ForceNew: true, DiffSuppressFunc: suppressEmptyDescription},
		"description":       {Description: "Description stored on this revision. Omit to inherit.", Type: schema.TypeString, Optional: true, ForceNew: true, DiffSuppressFunc: suppressEmptyDescription},
		"tool_type": {
			Description:      "Load generator, used only to compare a rewritten JMeter plan. Not sent to the revision API. `Locust`, `K6`, or `JMeter`. Omit to keep the stored type.",
			Type:             schema.TypeString,
			Optional:         true,
			ForceNew:         true,
			ValidateFunc:     validation.StringInSlice([]string{"", "Locust", "K6", "JMeter"}, false),
			DiffSuppressFunc: suppressEmptyDescription,
		},
		"tool_config": {
			Description:      "JSON object stored as this revision's toolConfig. Omit to inherit the previous revision.",
			Type:             schema.TypeString,
			Optional:         true,
			Computed:         true,
			ForceNew:         true,
			ValidateFunc:     validateToolConfig,
			DiffSuppressFunc: suppressRevisionToolConfig,
		},
		"variable": variable,
		"tags": {
			Description: "Tags stored on this revision. Omit to inherit.",
			Type:        schema.TypeSet,
			Optional:    true,
			Computed:    true,
			ForceNew:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"environment_id":    {Description: "Environment pinned on this revision. Omit to inherit.", Type: schema.TypeString, Optional: true, ForceNew: true, DiffSuppressFunc: suppressEmptyDescription},
		"infrastructure_id": {Description: "Infrastructure pinned on this revision. Omit to inherit.", Type: schema.TypeString, Optional: true, ForceNew: true, DiffSuppressFunc: suppressEmptyDescription},
		"script_content":    {Description: "Base64 script for this revision. Omit to inherit.", Type: schema.TypeString, Optional: true, Computed: true, ForceNew: true, DiffSuppressFunc: suppressEmptyDescription},
		"script_source":     {Description: "How the script is supplied. `inline` or `image`. Omit to inherit.", Type: schema.TypeString, Optional: true, Computed: true, ForceNew: true, ValidateFunc: validateOptionalScriptSource, DiffSuppressFunc: suppressEmptyDescription},
		"authored_tool_config": {
			Description: "Internal. The tool_config last accepted for this revision. Do not set this.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"unique_id": {Description: "Server-generated unique id of this revision.", Type: schema.TypeString, Computed: true},
	}
}

func dataSourceChaosLoadTestTemplateRevisionSchema() map[string]*schema.Schema {
	src := resourceChaosLoadTestTemplateRevisionSchema()
	delete(src, "authored_tool_config")
	for name, field := range src {
		if name == "org_id" || name == "project_id" || name == "hub_identity" || name == "template_identity" || name == "revision" {
			continue
		}
		field.Required = false
		field.Optional = false
		field.ForceNew = false
		field.Computed = true
		field.ValidateFunc = nil
		field.DiffSuppressFunc = nil
	}
	return src
}

func resourceChaosLoadTestTemplateRevisionCustomizeDiff(_ context.Context, d *schema.ResourceDiff, _ interface{}) error {
	if err := validateTemplateScope(fmt.Sprint(d.Get("org_id")), fmt.Sprint(d.Get("project_id"))); err != nil {
		return err
	}
	if err := keepInheritedCollection(d, "tags"); err != nil {
		return err
	}
	return keepInheritedCollection(d, "variable")
}

func keepInheritedCollection(d *schema.ResourceDiff, name string) error {
	oldValue, newValue := d.GetChange(name)
	if !collectionEmpty(newValue) || collectionEmpty(oldValue) {
		return nil
	}
	return d.SetNew(name, oldValue)
}

func collectionEmpty(value interface{}) bool {
	switch typed := value.(type) {
	case nil:
		return true
	case []interface{}:
		return len(typed) == 0
	case *schema.Set:
		return typed == nil || typed.Len() == 0
	default:
		return false
	}
}

func suppressRevisionToolConfig(key, oldValue, newValue string, d *schema.ResourceData) bool {
	if strings.TrimSpace(newValue) == "" && strings.TrimSpace(oldValue) != "" {
		return true
	}
	return suppressToolConfigDiff(key, oldValue, newValue, d)
}

func resourceChaosLoadTestTemplateRevisionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if err := validateTemplateScope(d.Get("org_id").(string), d.Get("project_id").(string)); err != nil {
		return diag.FromErr(err)
	}
	if strings.EqualFold(d.Get("revision").(string), "latest") {
		return diag.Errorf("revision must be a real revision id, not \"latest\"")
	}
	c, ctx := meta.(*internal.Session).GetLoadTestClientWithContext(ctx)
	req, err := templateRevisionRequest(d)
	if err != nil {
		return diag.FromErr(err)
	}
	hub := d.Get("hub_identity").(string)
	identity := d.Get("template_identity").(string)
	created, httpResp, err := sdkloadtest.NewLoadTestClient(c).CreateTemplateRevision(ctx, scopeFrom(d, c.AccountId), hub, identity, req)
	if err != nil {
		return helpers.HandleChaosApiError(err, d, httpResp)
	}
	if created == nil || created.Revision == "" {
		return diag.Errorf("create template revision returned an empty revision")
	}
	d.SetId(created.Revision)
	authored, _ := d.Get("tool_config").(string)
	diags := readChaosLoadTestTemplateRevision(ctx, d, meta)
	if diags.HasError() || d.Id() == "" {
		return diags
	}
	if err := d.Set("authored_tool_config", authored); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceChaosLoadTestTemplateRevisionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return readChaosLoadTestTemplateRevision(ctx, d, meta)
}

func resourceChaosLoadTestTemplateRevisionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetLoadTestClientWithContext(ctx)
	client := sdkloadtest.NewLoadTestClient(c)
	scope := scopeFrom(d, c.AccountId)
	hub := d.Get("hub_identity").(string)
	identity := d.Get("template_identity").(string)
	pinned, err := revisionPinnedBy(ctx, client, scope, hub, identity, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	if pinned != "" {
		return diag.Errorf("revision %s is still pinned by REFERENCE load test %s. Delete or repin that load test before deleting this revision", d.Id(), pinned)
	}
	httpResp, err := client.DeleteTemplateRevision(ctx, scope, hub, identity, d.Id())
	if err != nil {
		return helpers.HandleChaosReadApiError(err, d, httpResp)
	}
	d.SetId("")
	return nil
}

func readChaosLoadTestTemplateRevision(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetLoadTestClientWithContext(ctx)
	client := sdkloadtest.NewLoadTestClient(c)
	scope := scopeFrom(d, c.AccountId)
	hub := d.Get("hub_identity").(string)
	identity := d.Get("template_identity").(string)
	revision := d.Id()
	if revision == "" {
		revision = d.Get("revision").(string)
	}
	test, httpResp, err := client.GetTemplateRevision(ctx, scope, hub, identity, revision)
	if err != nil {
		return helpers.HandleChaosReadApiError(err, d, httpResp)
	}
	variables, varResp, err := client.GetTemplateVariables(ctx, scope, hub, identity, test.Revision)
	if err != nil {
		return helpers.HandleChaosApiError(err, d, varResp)
	}
	return diag.FromErr(setTemplateRevisionState(d, test, variables))
}

func dataSourceChaosLoadTestTemplateRevisionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	d.SetId(d.Get("revision").(string))
	diags := readChaosLoadTestTemplateRevision(ctx, d, meta)
	if !diags.HasError() && d.Id() == "" {
		return diag.Errorf("chaos load test template revision %q was not found", d.Get("revision"))
	}
	return diags
}

func importChaosLoadTestTemplateRevision(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	orgID, projectID, hubID, identity, revision, err := parseTemplateRevisionImportID(d.Id())
	if err != nil {
		return nil, err
	}
	d.SetId(revision)
	for key, value := range map[string]string{
		"org_id": orgID, "project_id": projectID, "hub_identity": hubID, "template_identity": identity, "revision": revision,
	} {
		if err := d.Set(key, value); err != nil {
			return nil, err
		}
	}
	diags := readChaosLoadTestTemplateRevision(ctx, d, meta)
	if diags.HasError() || d.Id() == "" {
		d.SetId("")
		return nil, fmt.Errorf("error reading chaos load test template revision: %v", diags)
	}
	return []*schema.ResourceData{d}, nil
}

func parseTemplateRevisionImportID(id string) (string, string, string, string, string, error) {
	parts := strings.Split(id, "/")
	if len(parts) != 5 || parts[2] == "" || parts[3] == "" || parts[4] == "" {
		return "", "", "", "", "", fmt.Errorf("import id must be <org_id>/<project_id>/<hub_identity>/<template_identity>/<revision>. Use //<hub_identity>/<template_identity>/<revision> for an account hub and <org_id>//<hub_identity>/<template_identity>/<revision> for an organization hub")
	}
	if err := validateTemplateScope(parts[0], parts[1]); err != nil {
		return "", "", "", "", "", err
	}
	for _, part := range parts[2:] {
		if !loadTestIdentifierPattern.MatchString(part) {
			return "", "", "", "", "", fmt.Errorf("hub_identity, template_identity, and revision must contain only letters, digits, underscores, hyphens, and $")
		}
	}
	return parts[0], parts[1], parts[2], parts[3], parts[4], nil
}

func revisionPinnedBy(ctx context.Context, client *sdkloadtest.LoadTestClient, scope sdkloadtest.Scope, hub, identity, revision string) (string, error) {
	page := 0
	for {
		listed, _, err := client.ListTemplateReferencedBy(ctx, scope, hub, identity, page, 100)
		if err != nil {
			return "", err
		}
		if listed == nil {
			return "", nil
		}
		for _, item := range listed.Data {
			if item.Kind != "" && !strings.EqualFold(item.Kind, "LoadTest") {
				continue
			}
			test, _, err := client.Get(ctx, sdkloadtest.Scope{
				AccountID: scope.AccountID,
				OrgID:     item.OrgIdentifier,
				ProjectID: item.ProjectIdentifier,
			}, item.Identity)
			if err != nil {
				return "", fmt.Errorf("read load test %s before deleting revision %s: %w", item.Identity, revision, err)
			}
			if test == nil || test.TemplateReference == nil || !strings.EqualFold(test.ImportType, "REFERENCE") {
				continue
			}
			ref := test.TemplateReference
			if ref.Identity == identity && ref.Revision == revision && (ref.HubIdentity == "" || ref.HubIdentity == hub) {
				return item.OrgIdentifier + "/" + item.ProjectIdentifier + "/" + item.Identity, nil
			}
		}
		page++
		if listed.Pagination.TotalPages == 0 || int64(page) >= listed.Pagination.TotalPages {
			return "", nil
		}
	}
}

func templateRevisionRequest(d *schema.ResourceData) (sdkloadtest.LoadTestTemplateRevisionRequest, error) {
	toolConfig, err := revisionToolConfig(d)
	if err != nil {
		return sdkloadtest.LoadTestTemplateRevisionRequest{}, err
	}
	req := sdkloadtest.LoadTestTemplateRevisionRequest{
		Name:          d.Get("name").(string),
		Description:   d.Get("description").(string),
		Revision:      d.Get("revision").(string),
		ToolConfig:    toolConfig,
		EnvID:         d.Get("environment_id").(string),
		InfraID:       d.Get("infrastructure_id").(string),
		ScriptContent: d.Get("script_content").(string),
		ScriptSource:  d.Get("script_source").(string),
	}
	if _, ok := d.GetOk("tags"); ok {
		req.Tags = expandStringSet(d.Get("tags"))
	}
	if _, ok := d.GetOk("variable"); ok {
		req.Variables = expandVariables(d.Get("variable"))
	}
	return req, nil
}

func revisionToolConfig(d *schema.ResourceData) (map[string]interface{}, error) {
	raw := strings.TrimSpace(d.Get("tool_config").(string))
	if raw == "" {
		return nil, nil
	}
	cfg, err := parseToolConfig(raw)
	if err != nil {
		return nil, err
	}
	if err := rejectAuthoredVariables(cfg); err != nil {
		return nil, err
	}
	if toolType := d.Get("tool_type").(string); toolType != "" {
		if err := validateToolMode(toolType, cfg); err != nil {
			return nil, err
		}
	}
	return cfg, nil
}

func setTemplateRevisionState(d *schema.ResourceData, test *sdkloadtest.LoadTestTemplate, variables []sdkloadtest.LoadTestVariable) error {
	if test == nil || test.Revision == "" {
		return fmt.Errorf("template revision was empty")
	}
	d.SetId(test.Revision)
	config, err := normalizeToolConfigValue(test.ToolConfig)
	if err != nil {
		return err
	}
	values := map[string]interface{}{
		"template_identity": test.Identity,
		"revision":          test.Revision,
		"name":              test.Name,
		"description":       test.Description,
		"tool_config":       config,
		"tags":              nonNilStrings(test.Tags),
		"environment_id":    test.EnvID,
		"infrastructure_id": test.InfraID,
		"script_content":    test.ScriptContent,
		"script_source":     test.ScriptSource,
		"unique_id":         test.UniqueID,
		"variable":          flattenVariables(variables),
	}
	if test.HubIdentity != "" {
		values["hub_identity"] = test.HubIdentity
	}
	if test.ToolType != "" {
		values["tool_type"] = test.ToolType
	}
	for key, value := range values {
		if err := d.Set(key, value); err != nil {
			return fmt.Errorf("set %s: %w", key, err)
		}
	}
	return nil
}
