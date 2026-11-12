package loadtest

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	sdkloadtest "github.com/harness/harness-go-sdk/harness/loadtest"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceChaosLoadTestTemplate() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for a Harness load test hub template.\n\n" +
			"Creates a template and updates the revision id stored at create. " +
			"A new published revision is a harness_chaos_load_test_template_revision. " +
			"Publishing a newer revision does not change this resource's revision id. " +
			"Reads return that stored revision, including its name, tool config, tags, script, and variables. " +
			"An update is refused once that id is no longer the latest revision, so the newer revision is left as it is. " +
			"Delete returns an error when a load test still references the template. " +
			"`revision` on create is a real revision id. `latest` is reserved for syncing a load test. " +
			"Leave `org_id` and `project_id` empty for an account hub. Set `org_id` and leave `project_id` empty for an organization hub. " +
			"The identifier and hub identity allow letters, digits, underscores, hyphens, and `$`. " +
			"An empty description, environment, infrastructure, infra type, script, or script source keeps the stored value. The API does not clear those fields.",
		CreateContext: resourceChaosLoadTestTemplateCreate,
		ReadContext:   resourceChaosLoadTestTemplateRead,
		CustomizeDiff: resourceChaosLoadTestTemplateCustomizeDiff,
		UpdateContext: resourceChaosLoadTestTemplateUpdate,
		DeleteContext: resourceChaosLoadTestTemplateDelete,
		Importer:      &schema.ResourceImporter{StateContext: importChaosLoadTestTemplate},
		Schema:        resourceChaosLoadTestTemplateSchema(),
	}
}

func DataSourceChaosLoadTestTemplate() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for a Harness load test hub template. Reads the latest revision.",
		ReadContext: dataSourceChaosLoadTestTemplateRead,
		Schema:      dataSourceChaosLoadTestTemplateSchema(),
	}
}

func resourceChaosLoadTestTemplateSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"org_id": {
			Description: "Organization identifier of the template. Empty for an account hub.",
			Type:        schema.TypeString,
			Optional:    true,
			ForceNew:    true,
		},
		"project_id": {
			Description: "Project identifier of the template. Empty for an account or organization hub. Requires org_id.",
			Type:        schema.TypeString,
			Optional:    true,
			ForceNew:    true,
		},
		"hub_identity": {
			Description:  "Chaos hub that holds the template. Letters, digits, `_`, `-`, and `$` only.",
			Type:         schema.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validateLoadTestIdentifier,
		},
		"identifier": {
			Description:  "Template identifier. Letters, digits, `_`, `-`, and `$` only. Changing it removes this template.",
			Type:         schema.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validateLoadTestIdentifier,
		},
		"name": {
			Description:  "Name of this revision.",
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},
		"description": {
			Description:      "Description of this revision. An empty value leaves the stored description in place.",
			Type:             schema.TypeString,
			Optional:         true,
			DiffSuppressFunc: suppressEmptyDescription,
		},
		"revision": {
			Description:  "Revision id stored when the template is created. A later published revision does not change this value. Changing it in configuration removes this template. Use harness_chaos_load_test_template_revision to publish another revision. `latest` is rejected.",
			Type:         schema.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},
		"tool_type": {
			Description:  "Load generator. `Locust`, `K6`, or `JMeter`. Changing it removes this template.",
			Type:         schema.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringInSlice([]string{"Locust", "K6", "JMeter"}, false),
		},
		"tool_config": {
			Description:      "JSON object stored as this revision's toolConfig. JMeter plans a change only when this differs from the configuration already accepted.",
			Type:             schema.TypeString,
			Optional:         true,
			Computed:         true,
			ValidateFunc:     validateToolConfig,
			DiffSuppressFunc: suppressToolConfigDiff,
		},
		"variable": templateVariableSchema(),
		"tags": {
			Description: "Tags on this revision. An empty list clears them.",
			Type:        schema.TypeSet,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"environment_id": {
			Description:      "Environment pinned on the template. Callers still pass an environment when creating a load test from it. An empty value leaves the stored environment in place.",
			Type:             schema.TypeString,
			Optional:         true,
			DiffSuppressFunc: suppressEmptyDescription,
		},
		"infrastructure_id": {
			Description:      "Infrastructure pinned on the template. An empty value leaves the stored infrastructure in place.",
			Type:             schema.TypeString,
			Optional:         true,
			DiffSuppressFunc: suppressEmptyDescription,
		},
		"infra_type": {
			Description:      "Infrastructure type stored on the template. An empty value leaves the stored type in place.",
			Type:             schema.TypeString,
			Optional:         true,
			DiffSuppressFunc: suppressEmptyDescription,
		},
		"script_content": {
			Description:      "Base64 script stored on the template. Load Test Manager fills this from tool_config when it is omitted, and an empty value leaves that stored script in place.",
			Type:             schema.TypeString,
			Optional:         true,
			Computed:         true,
			DiffSuppressFunc: suppressEmptyDescription,
		},
		"script_source": {
			Description:      "How the script is supplied. `inline` or `image`. Load Test Manager fills this when it is omitted, and an empty value leaves the stored source in place.",
			Type:             schema.TypeString,
			Optional:         true,
			Computed:         true,
			ValidateFunc:     validateOptionalScriptSource,
			DiffSuppressFunc: suppressEmptyDescription,
		},
		"authored_tool_config": {
			Description: "Internal. The tool_config last accepted for this template. Do not set this.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"unique_id": {
			Description: "Server-generated unique id of this revision.",
			Type:        schema.TypeString,
			Computed:    true,
		},
	}
}

func dataSourceChaosLoadTestTemplateSchema() map[string]*schema.Schema {
	src := resourceChaosLoadTestTemplateSchema()
	delete(src, "authored_tool_config")
	for name, field := range src {
		if name == "org_id" || name == "project_id" || name == "identifier" || name == "hub_identity" {
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

func templateVariableSchema() *schema.Schema {
	return &schema.Schema{
		Description: "Custom variables on this revision. An empty list clears them.",
		Type:        schema.TypeList,
		Optional:    true,
		Elem: &schema.Resource{Schema: map[string]*schema.Schema{
			"name":        {Description: "Variable name.", Type: schema.TypeString, Required: true, ValidateFunc: validation.StringIsNotEmpty},
			"type":        {Description: "Variable type. `String` or `Number`.", Type: schema.TypeString, Required: true, ValidateFunc: validation.StringInSlice([]string{"String", "Number"}, false)},
			"value":       {Description: "Variable value.", Type: schema.TypeString, Required: true},
			"description": {Description: "Variable description.", Type: schema.TypeString, Optional: true},
			"required":    {Description: "Whether a caller must supply this variable.", Type: schema.TypeBool, Optional: true},
		}},
	}
}

func resourceChaosLoadTestTemplateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if err := validateTemplateScope(d.Get("org_id").(string), d.Get("project_id").(string)); err != nil {
		return diag.FromErr(err)
	}
	if strings.EqualFold(d.Get("revision").(string), "latest") {
		return diag.Errorf("revision cannot be \"latest\": that name is reserved for syncing a load test")
	}
	c, ctx := meta.(*internal.Session).GetLoadTestClientWithContext(ctx)
	req, err := templateCreateRequest(d)
	if err != nil {
		return diag.FromErr(err)
	}
	created, httpResp, err := sdkloadtest.NewLoadTestClient(c).CreateTemplate(ctx, scopeFrom(d, c.AccountId), d.Get("hub_identity").(string), req)
	if err != nil {
		return helpers.HandleChaosApiError(err, d, httpResp)
	}
	if created == nil || created.Identity == "" {
		return diag.Errorf("create load test template returned an empty identity")
	}
	d.SetId(created.Identity)
	authored, _ := d.Get("tool_config").(string)
	diags := readChaosLoadTestTemplate(ctx, d, meta)
	if diags.HasError() || d.Id() == "" {
		return diags
	}
	if err := d.Set("authored_tool_config", authored); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceChaosLoadTestTemplateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return readChaosLoadTestTemplate(ctx, d, meta)
}

func resourceChaosLoadTestTemplateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetLoadTestClientWithContext(ctx)
	client := sdkloadtest.NewLoadTestClient(c)
	scope := scopeFrom(d, c.AccountId)
	hub := d.Get("hub_identity").(string)
	current, httpResp, err := client.GetTemplate(ctx, scope, hub, d.Id())
	if err != nil {
		return helpers.HandleChaosApiError(err, d, httpResp)
	}
	stored, _ := d.Get("revision").(string)
	if current != nil && current.Revision != "" && stored != "" && current.Revision != stored {
		return diag.Errorf("revision %s is no longer the latest revision (%s). This update was not sent, so the newer revision was left unchanged", stored, current.Revision)
	}
	req, err := templateUpdateRequest(d)
	if err != nil {
		return diag.FromErr(err)
	}
	authored, _ := d.Get("tool_config").(string)
	_, httpResp, err = client.UpdateTemplate(ctx, scope, hub, d.Id(), req)
	if err != nil {
		return helpers.HandleChaosApiError(err, d, httpResp)
	}
	diags := readChaosLoadTestTemplate(ctx, d, meta)
	if diags.HasError() || d.Id() == "" {
		return diags
	}
	if err := d.Set("authored_tool_config", authored); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceChaosLoadTestTemplateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetLoadTestClientWithContext(ctx)
	httpResp, err := sdkloadtest.NewLoadTestClient(c).DeleteTemplate(ctx, scopeFrom(d, c.AccountId), d.Get("hub_identity").(string), d.Id())
	if err != nil {
		return helpers.HandleChaosReadApiError(err, d, httpResp)
	}
	d.SetId("")
	return nil
}

func readChaosLoadTestTemplate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetLoadTestClientWithContext(ctx)
	client := sdkloadtest.NewLoadTestClient(c)
	scope := scopeFrom(d, c.AccountId)
	hub := d.Get("hub_identity").(string)
	revision, _ := d.Get("revision").(string)
	var test *sdkloadtest.LoadTestTemplate
	var httpResp *http.Response
	var err error
	if revision == "" {
		test, httpResp, err = client.GetTemplate(ctx, scope, hub, d.Id())
	} else {
		test, httpResp, err = client.GetTemplateRevision(ctx, scope, hub, d.Id(), revision)
	}
	if err != nil {
		return helpers.HandleChaosReadApiError(err, d, httpResp)
	}
	variables, varResp, err := client.GetTemplateVariables(ctx, scope, hub, d.Id(), revision)
	if err != nil {
		return helpers.HandleChaosApiError(err, d, varResp)
	}
	if err := setTemplateState(d, test, variables); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func dataSourceChaosLoadTestTemplateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	d.SetId(d.Get("identifier").(string))
	diags := readChaosLoadTestTemplate(ctx, d, meta)
	if !diags.HasError() && d.Id() == "" {
		return diag.Errorf("chaos load test template %q was not found", d.Get("identifier"))
	}
	return diags
}

func importChaosLoadTestTemplate(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	orgID, projectID, hubID, identifier, err := parseTemplateImportID(d.Id())
	if err != nil {
		return nil, err
	}
	d.SetId(identifier)
	if err := d.Set("org_id", orgID); err != nil {
		return nil, err
	}
	if err := d.Set("project_id", projectID); err != nil {
		return nil, err
	}
	if err := d.Set("hub_identity", hubID); err != nil {
		return nil, err
	}
	if err := d.Set("identifier", identifier); err != nil {
		return nil, err
	}
	diags := readChaosLoadTestTemplate(ctx, d, meta)
	if diags.HasError() || d.Id() == "" {
		d.SetId("")
		return nil, fmt.Errorf("error reading chaos load test template: %v", diags)
	}
	return []*schema.ResourceData{d}, nil
}

func resourceChaosLoadTestTemplateCustomizeDiff(_ context.Context, d *schema.ResourceDiff, _ interface{}) error {
	return validateTemplateScope(fmt.Sprint(d.Get("org_id")), fmt.Sprint(d.Get("project_id")))
}

func parseTemplateImportID(id string) (string, string, string, string, error) {
	parts := strings.Split(id, "/")
	if len(parts) != 4 || parts[2] == "" || parts[3] == "" {
		return "", "", "", "", fmt.Errorf("import id must be <org_id>/<project_id>/<hub_identity>/<identifier>. Use //<hub_identity>/<identifier> for an account hub and <org_id>//<hub_identity>/<identifier> for an organization hub")
	}
	if err := validateTemplateScope(parts[0], parts[1]); err != nil {
		return "", "", "", "", err
	}
	if !loadTestIdentifierPattern.MatchString(parts[2]) || !loadTestIdentifierPattern.MatchString(parts[3]) {
		return "", "", "", "", fmt.Errorf("hub_identity and identifier must contain only letters, digits, underscores, hyphens, and $")
	}
	return parts[0], parts[1], parts[2], parts[3], nil
}

func templateCreateRequest(d *schema.ResourceData) (sdkloadtest.LoadTestTemplateRequest, error) {
	toolConfig, err := toolConfigFromResource(d)
	if err != nil {
		return sdkloadtest.LoadTestTemplateRequest{}, err
	}
	return sdkloadtest.LoadTestTemplateRequest{
		Identity:      d.Get("identifier").(string),
		Name:          d.Get("name").(string),
		Description:   d.Get("description").(string),
		Revision:      d.Get("revision").(string),
		ToolType:      d.Get("tool_type").(string),
		ToolConfig:    toolConfig,
		Variables:     expandVariables(d.Get("variable")),
		Tags:          expandStringSet(d.Get("tags")),
		EnvID:         d.Get("environment_id").(string),
		InfraID:       d.Get("infrastructure_id").(string),
		InfraType:     d.Get("infra_type").(string),
		ScriptContent: d.Get("script_content").(string),
		ScriptSource:  d.Get("script_source").(string),
	}, nil
}

func templateUpdateRequest(d *schema.ResourceData) (sdkloadtest.LoadTestTemplateUpdateRequest, error) {
	toolConfig, err := toolConfigFromResource(d)
	if err != nil {
		return sdkloadtest.LoadTestTemplateUpdateRequest{}, err
	}
	return sdkloadtest.LoadTestTemplateUpdateRequest{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		ToolConfig:  toolConfig,
		Variables:   expandVariables(d.Get("variable")),
		Tags:        expandStringSet(d.Get("tags")),
		EnvID:       d.Get("environment_id").(string),
		InfraID:     d.Get("infrastructure_id").(string),
		InfraType:   d.Get("infra_type").(string),
	}, nil
}

func setTemplateState(d *schema.ResourceData, test *sdkloadtest.LoadTestTemplate, variables []sdkloadtest.LoadTestVariable) error {
	if test == nil {
		return fmt.Errorf("load test template was empty")
	}
	d.SetId(test.Identity)
	config, err := normalizeToolConfigValue(test.ToolConfig)
	if err != nil {
		return err
	}
	revision := test.Revision
	if stored, _ := d.Get("revision").(string); stored != "" {
		revision = stored
	}
	values := map[string]interface{}{
		"identifier":        test.Identity,
		"name":              test.Name,
		"description":       test.Description,
		"revision":          revision,
		"tool_type":         test.ToolType,
		"tool_config":       config,
		"tags":              nonNilStrings(test.Tags),
		"environment_id":    test.EnvID,
		"infrastructure_id": test.InfraID,
		"infra_type":        test.InfraType,
		"script_content":    test.ScriptContent,
		"script_source":     test.ScriptSource,
		"unique_id":         test.UniqueID,
		"variable":          flattenVariables(variables),
	}
	if test.HubIdentity != "" {
		values["hub_identity"] = test.HubIdentity
	}
	for key, value := range values {
		if err := d.Set(key, value); err != nil {
			return fmt.Errorf("set %s: %w", key, err)
		}
	}
	return nil
}
