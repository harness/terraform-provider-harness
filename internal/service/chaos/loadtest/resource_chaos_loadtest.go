package loadtest

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	sdkloadtest "github.com/harness/harness-go-sdk/harness/loadtest"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceChaosLoadTest() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for a Harness load test.\n\n" +
			"Creates and updates the load test definition in Load Test Manager. " +
			"It does not start a run. `tool_config` is the toolConfig JSON the API stores. " +
			"K6 is a flat object. Locust is nested under `locust`. JMeter is not supported. " +
			"K6 and managed infrastructure require target_type kubernetes. Locust may omit target_type; the API stores linux. " +
			"REFERENCE template imports cannot be managed here. " +
			"`tool_config` is saved in Terraform state and includes the script. " +
			"Custom variables belong in `variable`. A variables key inside tool_config is rejected. " +
			"The identifier allows letters, digits, underscores, and hyphens. " +
			"Changing identifier, org_id, project_id, or tool_type removes the load test. Recorded runs stay stored under the old unique id and are not listed on a recreated test. " +
			"Destroy deletes the load test definition. It does not stop a run that is already in progress, and recorded runs stay stored. " +
			"Only Locust and K6 can be imported. The data source can still read a JMeter test or a REFERENCE import. " +
			"The first apply after import sends the configuration as the full desired state. Copy variable, resources, max_duration_sec, tags, tool_config, cleanup_policy, and managed_infra from state first. An omitted variable list, resources block, or tags set clears the stored value. cleanup_policy defaults to delete and managed_infra defaults to false.",

		CreateContext: resourceChaosLoadTestCreate,
		ReadContext:   resourceChaosLoadTestRead,
		UpdateContext: resourceChaosLoadTestUpdate,
		DeleteContext: resourceChaosLoadTestDelete,
		CustomizeDiff: resourceChaosLoadTestCustomizeDiff,
		Importer: &schema.ResourceImporter{
			StateContext: importChaosLoadTest,
		},
		Schema: resourceChaosLoadTestSchema(),
	}
}

func resourceChaosLoadTestCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetLoadTestClientWithContext(ctx)
	req, err := loadTestCreateRequest(d)
	if err != nil {
		return diag.FromErr(err)
	}
	log.Printf("[DEBUG] Creating chaos load test %s", req.Identity)
	created, httpResp, err := sdkloadtest.NewLoadTestClient(c).Create(ctx, scopeFrom(d, c.AccountId), req)
	if err != nil {
		if loadTestCreateCommitted(httpResp) {
			return rollbackCreatedLoadTest(ctx, c, d, req.Identity, diag.FromErr(err))
		}
		diags := helpers.HandleChaosApiError(err, d, httpResp)
		if httpResp == nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary: fmt.Sprintf(
					"Create of load test %s returned no response, so it was not deleted. If Harness saved it, delete that load test before the next apply, or the apply fails with a duplicate identity.",
					req.Identity,
				),
			})
		}
		return diags
	}
	if created == nil || created.Identity == "" {
		if loadTestCreateCommitted(httpResp) {
			return rollbackCreatedLoadTest(ctx, c, d, req.Identity, diag.Errorf("create load test returned an empty identity"))
		}
		return diag.Errorf("create load test returned an empty identity")
	}
	d.SetId(created.Identity)
	diags := readChaosLoadTest(ctx, d, meta, true, false)
	if !diags.HasError() && d.Id() != "" {
		return nil
	}
	if !diags.HasError() {
		diags = diag.Errorf("load test %s could not be read after create", created.Identity)
	}
	return rollbackCreatedLoadTest(ctx, c, d, created.Identity, diags)
}

func rollbackCreatedLoadTest(ctx context.Context, c *sdkloadtest.APIClient, d *schema.ResourceData, identity string, readErr diag.Diagnostics) diag.Diagnostics {
	reason := diagnosticSummary(readErr)
	log.Printf("[WARN] Deleting load test %s because reading it after create failed: %s", identity, reason)
	_, httpResp, err := sdkloadtest.NewLoadTestClient(c).Delete(ctx, scopeFrom(d, c.AccountId), identity)
	if err != nil && (httpResp == nil || httpResp.StatusCode != http.StatusNotFound) {
		if d.Id() != "" {
			return diag.Errorf("%s. Load test %s is stored and this resource stays tainted because deleting it failed: %s. Do not apply: a tainted apply deletes the load test. Recorded runs stay stored.", reason, identity, err)
		}
		return diag.Errorf("%s. Load test %s may still exist because deleting it failed: %s. Delete it before the next apply, or that apply fails with a duplicate identity.", reason, identity, err)
	}
	d.SetId("")
	return diag.Errorf("%s. The load test created by this apply was deleted, so nothing was saved in state. The next apply creates it again.", reason)
}

func loadTestCreateCommitted(resp *http.Response) bool {
	return resp != nil && resp.StatusCode >= http.StatusOK && resp.StatusCode < http.StatusMultipleChoices
}

func diagnosticSummary(diags diag.Diagnostics) string {
	parts := make([]string, 0, len(diags))
	for _, item := range diags {
		if item.Summary != "" {
			parts = append(parts, item.Summary)
		}
	}
	return strings.Join(parts, "; ")
}

func resourceChaosLoadTestRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return readChaosLoadTest(ctx, d, meta, true, false)
}

func readChaosLoadTest(ctx context.Context, d *schema.ResourceData, meta interface{}, rejectReference, managedToolsOnly bool) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetLoadTestClientWithContext(ctx)
	identity := d.Id()
	log.Printf("[DEBUG] Reading chaos load test %s", identity)
	client := sdkloadtest.NewLoadTestClient(c)
	scope := scopeFrom(d, c.AccountId)
	test, httpResp, err := client.Get(ctx, scope, identity)
	if err != nil {
		return helpers.HandleChaosReadApiError(err, d, httpResp)
	}
	if managedToolsOnly && test.ToolType != sdkloadtest.LoadTestToolLocust && test.ToolType != sdkloadtest.LoadTestToolK6 {
		d.SetId("")
		return diag.Errorf("load test %s has tool type %q; only Locust and K6 can be imported", identity, test.ToolType)
	}
	if rejectReference && strings.EqualFold(test.ImportType, "REFERENCE") {
		return diag.Errorf("load test %s is a REFERENCE import owned by its template and cannot be managed by this resource", identity)
	}
	if err := serviceReferencesDrift(d, test.ServiceReferences); err != nil {
		return diag.FromErr(err)
	}
	variables, varResp, err := client.GetVariables(ctx, scope, identity)
	if err != nil {
		return helpers.HandleChaosApiError(err, d, varResp)
	}
	custom := []sdkloadtest.LoadTestVariable{}
	if variables != nil {
		custom = variables.Variables
	}
	if err := setLoadTestState(d, test, custom); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceChaosLoadTestUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetLoadTestClientWithContext(ctx)
	req, err := loadTestUpdateRequest(d)
	if err != nil {
		return diag.FromErr(err)
	}
	log.Printf("[DEBUG] Updating chaos load test %s", d.Id())
	_, httpResp, err := sdkloadtest.NewLoadTestClient(c).Update(ctx, scopeFrom(d, c.AccountId), d.Id(), req)
	if err != nil {
		return helpers.HandleChaosApiError(err, d, httpResp)
	}
	return readChaosLoadTest(ctx, d, meta, true, false)
}

func resourceChaosLoadTestDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetLoadTestClientWithContext(ctx)
	log.Printf("[DEBUG] Deleting chaos load test %s", d.Id())
	_, httpResp, err := sdkloadtest.NewLoadTestClient(c).Delete(ctx, scopeFrom(d, c.AccountId), d.Id())
	if err != nil {
		return helpers.HandleChaosReadApiError(err, d, httpResp)
	}
	d.SetId("")
	return nil
}

func resourceChaosLoadTestCustomizeDiff(_ context.Context, d *schema.ResourceDiff, _ interface{}) error {
	if d.NewValueKnown("service_references") {
		oldRefs, newRefs := d.GetChange("service_references")
		oldSet, _ := oldRefs.(*schema.Set)
		newSet, _ := newRefs.(*schema.Set)
		if newSet != nil && newSet.Len() == 0 && oldSet != nil && oldSet.Len() > 0 {
			if err := d.SetNew("service_references", oldSet); err != nil {
				return err
			}
		}
	}
	return validateLoadTestPlan(d)
}

func importChaosLoadTest(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
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
	diags := readChaosLoadTest(ctx, d, meta, true, true)
	if diags.HasError() {
		d.SetId("")
		return nil, fmt.Errorf("error reading chaos load test: %v", diags)
	}
	if d.Id() == "" {
		return nil, fmt.Errorf("chaos load test %s was not found", identifier)
	}
	return []*schema.ResourceData{d}, nil
}
