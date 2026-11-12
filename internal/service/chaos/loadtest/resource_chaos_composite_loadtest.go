package loadtest

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"regexp"

	sdkloadtest "github.com/harness/harness-go-sdk/harness/loadtest"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// A pipeline identifier starts with a letter or underscore and then allows letters, digits, underscores, and $.
// The composite pipeline id is the stage id plus "_pipeline", and that whole value must fit this pattern.
var pipelineIdentifierPattern = regexp.MustCompile(`^[a-zA-Z_][0-9a-zA-Z_$]{0,127}$`)

func ResourceChaosCompositeLoadTest() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for a Harness composite load test.\n\n" +
			"Creates the pipeline that runs one load test and one chaos probe in parallel. " +
			"It does not start a run. `identifier` is the stage id. It must start with a letter or underscore and contain only letters, digits, underscores, and `$`. A hyphen is not allowed. The pipeline id is that value with `_pipeline` appended, and that pipeline id must be at most 128 characters. " +
			"Changing `identifier`, `org_id`, or `project_id` removes the composite load test. " +
			"Destroy deletes that pipeline. The load test and the probe stay in place. " +
			"`load_test_ref` is a load test identity in this project. A `LOCAL` copy and a `REFERENCE` template import are both valid.",
		CreateContext: resourceChaosCompositeLoadTestCreate,
		ReadContext:   resourceChaosCompositeLoadTestRead,
		UpdateContext: resourceChaosCompositeLoadTestUpdate,
		DeleteContext: resourceChaosCompositeLoadTestDelete,
		Importer:      &schema.ResourceImporter{StateContext: importChaosCompositeLoadTest},
		Schema:        resourceChaosCompositeLoadTestSchema(),
	}
}

func DataSourceChaosCompositeLoadTest() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for a Harness composite load test. Reads the pipeline stored for the stage id.",
		ReadContext: dataSourceChaosCompositeLoadTestRead,
		Schema:      dataSourceChaosCompositeLoadTestSchema(),
	}
}

func resourceChaosCompositeLoadTestSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"org_id": {
			Description: "Organization identifier of the composite load test.",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"project_id": {
			Description: "Project identifier of the composite load test.",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"identifier": {
			Description:  "Stage identifier. Must start with a letter or underscore and contain only letters, digits, underscores, and `$`. A hyphen is not allowed. Changing it removes this composite load test. The pipeline id is this value with `_pipeline` appended and must be at most 128 characters.",
			Type:         schema.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validateCompositeIdentifier,
		},
		"name": {
			Description:  "Name of the composite load test.",
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},
		"description": {
			Description: "Description stored on the composite load test.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"objective": {
			Description: "Objective stored on the composite load test.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tags": {
			Description: "Tags stored on the composite load test. This is a map. An empty map clears them.",
			Type:        schema.TypeMap,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"load_test_ref": {
			Description:  "Identity of the load test in this project. A load test created from a template is valid.",
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},
		"probe": {
			Description: "Chaos probe that runs in parallel with the load test.",
			Type:        schema.TypeList,
			Required:    true,
			MaxItems:    1,
			Elem: &schema.Resource{Schema: map[string]*schema.Schema{
				"identity": {
					Description:  "Probe identity.",
					Type:         schema.TypeString,
					Required:     true,
					ValidateFunc: validation.StringIsNotEmpty,
				},
				"infra_reference": {
					Description:      "Infrastructure reference for the probe. Empty leaves the stored step as a runtime input. `<+input>` is that same empty value.",
					Type:             schema.TypeString,
					Optional:         true,
					DiffSuppressFunc: suppressProbePlaceholder,
				},
				"duration": {
					Description:      "Probe duration, for example `5m`. Empty leaves the default timeout. `10m` is that default and is stored empty.",
					Type:             schema.TypeString,
					Optional:         true,
					DiffSuppressFunc: suppressProbeDuration,
				},
			}},
		},
		"pipeline_identifier": {
			Description: "Pipeline id stored for this stage. `{identifier}_pipeline`.",
			Type:        schema.TypeString,
			Computed:    true,
		},
	}
}

func dataSourceChaosCompositeLoadTestSchema() map[string]*schema.Schema {
	src := resourceChaosCompositeLoadTestSchema()
	for name, field := range src {
		if name == "org_id" || name == "project_id" || name == "identifier" {
			continue
		}
		field.Required = false
		field.Optional = false
		field.ForceNew = false
		field.Computed = true
		field.ValidateFunc = nil
		if field.Type == schema.TypeList || field.Type == schema.TypeSet {
			field.MaxItems = 0
			field.MinItems = 0
		}
	}
	return src
}

func resourceChaosCompositeLoadTestCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetLoadTestClientWithContext(ctx)
	req := compositeLoadTestRequest(d)
	created, httpResp, err := sdkloadtest.NewLoadTestClient(c).CreateCompositeLoadTest(ctx, scopeFrom(d, c.AccountId), req)
	if err != nil {
		return helpers.HandleChaosApiError(err, d, httpResp)
	}
	identifier := ""
	if created != nil {
		identifier = created.Identifier
		if identifier == "" {
			identifier = created.StageIdentifier
		}
	}
	if identifier == "" {
		if loadTestCreateCommitted(httpResp) {
			requested := d.Get("identifier").(string)
			_, deleteErr := sdkloadtest.NewLoadTestClient(c).DeleteCompositeLoadTest(ctx, scopeFrom(d, c.AccountId), requested)
			if deleteErr != nil {
				return diag.Errorf("composite load test was created without a stage identifier, and deleting %s failed: %s. Delete that pipeline before the next apply.", requested, deleteErr)
			}
		}
		return diag.Errorf("composite load test was created without a stage identifier")
	}
	d.SetId(identifier)
	diags := resourceChaosCompositeLoadTestRead(ctx, d, meta)
	if !diags.HasError() && d.Id() != "" {
		return nil
	}
	if !diags.HasError() {
		diags = diag.Errorf("composite load test %s could not be read after create", identifier)
	}
	log.Printf("[WARN] Deleting composite load test %s because reading it after create failed", identifier)
	httpResp, err = sdkloadtest.NewLoadTestClient(c).DeleteCompositeLoadTest(ctx, scopeFrom(d, c.AccountId), identifier)
	if err != nil && (httpResp == nil || httpResp.StatusCode != http.StatusNotFound) {
		return diag.Errorf("%s. Composite load test %s is stored and this resource stays tainted because deleting it failed: %s. Do not apply: a tainted apply deletes the pipeline. The load test and the probe stay in place.", diagnosticSummary(diags), identifier, err)
	}
	d.SetId("")
	return diag.Errorf("%s. The composite load test created by this apply was deleted, so nothing was saved in state. The next apply creates it again.", diagnosticSummary(diags))
}

func resourceChaosCompositeLoadTestRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetLoadTestClientWithContext(ctx)
	got, httpResp, err := sdkloadtest.NewLoadTestClient(c).GetCompositeLoadTest(ctx, scopeFrom(d, c.AccountId), d.Id())
	if err != nil {
		return helpers.HandleChaosReadApiError(err, d, httpResp)
	}
	return setCompositeLoadTestState(d, got)
}

func resourceChaosCompositeLoadTestUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetLoadTestClientWithContext(ctx)
	req := compositeLoadTestRequest(d)
	_, httpResp, err := sdkloadtest.NewLoadTestClient(c).UpdateCompositeLoadTest(ctx, scopeFrom(d, c.AccountId), d.Id(), sdkloadtest.CompositeLoadTestUpdateRequest{
		Name:        req.Name,
		Description: req.Description,
		Objective:   req.Objective,
		Tags:        req.Tags,
		LoadTest:    req.LoadTest,
		Probe:       req.Probe,
	})
	if err != nil {
		return helpers.HandleChaosApiError(err, d, httpResp)
	}
	return resourceChaosCompositeLoadTestRead(ctx, d, meta)
}

func resourceChaosCompositeLoadTestDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetLoadTestClientWithContext(ctx)
	httpResp, err := sdkloadtest.NewLoadTestClient(c).DeleteCompositeLoadTest(ctx, scopeFrom(d, c.AccountId), d.Id())
	if err != nil {
		return helpers.HandleChaosReadApiError(err, d, httpResp)
	}
	return nil
}

func dataSourceChaosCompositeLoadTestRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	identity := d.Get("identifier").(string)
	d.SetId(identity)
	diags := resourceChaosCompositeLoadTestRead(ctx, d, meta)
	if !diags.HasError() && d.Id() == "" {
		return diag.Errorf("composite load test %q was not found", identity)
	}
	return diags
}

func importChaosCompositeLoadTest(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	orgID, projectID, identifier, err := parseLoadTestImportID(d.Id())
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
	if err := d.Set("identifier", identifier); err != nil {
		return nil, err
	}
	if err := compositeIdentifierError(identifier); err != nil {
		return nil, fmt.Errorf("invalid import ID %q: %w", d.Id(), err)
	}
	diags := resourceChaosCompositeLoadTestRead(ctx, d, meta)
	if diags.HasError() {
		d.SetId("")
		return nil, fmt.Errorf("error reading composite load test: %v", diags)
	}
	if d.Id() == "" {
		return nil, fmt.Errorf("composite load test %s was not found", identifier)
	}
	return []*schema.ResourceData{d}, nil
}

func validateCompositeIdentifier(v interface{}, k string) (ws []string, errs []error) {
	value, _ := v.(string)
	if err := compositeIdentifierError(value); err != nil {
		errs = append(errs, fmt.Errorf("%s %s", k, err.Error()))
	}
	return
}

func compositeIdentifierError(value string) error {
	pipelineID := value + "_pipeline"
	if pipelineIdentifierPattern.MatchString(value) && pipelineIdentifierPattern.MatchString(pipelineID) {
		return nil
	}
	return fmt.Errorf("must start with a letter or underscore and contain only letters, digits, underscores, and $. A hyphen is not allowed, and %q must be at most 128 characters", pipelineID)
}

func suppressProbePlaceholder(_, oldValue, newValue string, _ *schema.ResourceData) bool {
	return probePlaceholder(oldValue) == probePlaceholder(newValue)
}

func suppressProbeDuration(_, oldValue, newValue string, _ *schema.ResourceData) bool {
	return probeDuration(oldValue) == probeDuration(newValue)
}

func probePlaceholder(value string) string {
	if value == "<+input>" {
		return ""
	}
	return value
}

func probeDuration(value string) string {
	if value == "10m" {
		return ""
	}
	return value
}

func compositeLoadTestRequest(d *schema.ResourceData) sdkloadtest.CompositeLoadTestRequest {
	return sdkloadtest.CompositeLoadTestRequest{
		Name:        d.Get("name").(string),
		Identifier:  d.Get("identifier").(string),
		Description: d.Get("description").(string),
		Objective:   d.Get("objective").(string),
		Tags:        compositeTags(d),
		LoadTest:    sdkloadtest.CompositeLoadTestRef{LoadTestRef: d.Get("load_test_ref").(string)},
		Probe:       compositeProbe(d),
	}
}

func compositeTags(d *schema.ResourceData) map[string]string {
	raw, ok := d.GetOk("tags")
	if !ok {
		return map[string]string{}
	}
	tags := map[string]string{}
	for key, value := range raw.(map[string]interface{}) {
		tags[key], _ = value.(string)
	}
	return tags
}

func compositeProbe(d *schema.ResourceData) sdkloadtest.CompositeLoadTestProbe {
	raw, ok := d.GetOk("probe")
	if !ok {
		return sdkloadtest.CompositeLoadTestProbe{}
	}
	items := raw.([]interface{})
	if len(items) == 0 || items[0] == nil {
		return sdkloadtest.CompositeLoadTestProbe{}
	}
	block := items[0].(map[string]interface{})
	identity, _ := block["identity"].(string)
	infra, _ := block["infra_reference"].(string)
	duration, _ := block["duration"].(string)
	infra = probePlaceholder(infra)
	duration = probeDuration(duration)
	return sdkloadtest.CompositeLoadTestProbe{
		Identity:       identity,
		InfraReference: infra,
		Duration:       duration,
	}
}

func setCompositeLoadTestState(d *schema.ResourceData, got *sdkloadtest.CompositeLoadTest) diag.Diagnostics {
	if got == nil {
		return diag.Errorf("composite load test response was empty")
	}
	identifier := got.Identifier
	if identifier == "" {
		identifier = d.Id()
	}
	pipelineID := got.PipelineIdentifier
	if pipelineID == "" {
		pipelineID = identifier + "_pipeline"
	}
	tags := got.Tags
	if tags == nil {
		tags = map[string]string{}
	}
	values := map[string]interface{}{
		"identifier":          identifier,
		"name":                got.Name,
		"description":         got.Description,
		"objective":           got.Objective,
		"tags":                tags,
		"load_test_ref":       got.LoadTest.LoadTestRef,
		"pipeline_identifier": pipelineID,
		"probe": []interface{}{map[string]interface{}{
			"identity":        got.Probe.Identity,
			"infra_reference": got.Probe.InfraReference,
			"duration":        got.Probe.Duration,
		}},
	}
	for key, value := range values {
		if err := d.Set(key, value); err != nil {
			return diag.FromErr(err)
		}
	}
	return nil
}
