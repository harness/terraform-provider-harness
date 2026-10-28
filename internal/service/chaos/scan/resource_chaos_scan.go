package scan

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/harness/harness-go-sdk/harness/chaos"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// scanPollInterval is how often create, retry, and abort wait for a terminal status.
var scanPollInterval = 5 * time.Second

// ResourceChaosPipelineScan returns the pipeline scan resource.
func ResourceChaosPipelineScan() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for scanning an existing Harness pipeline execution.\n\n" +
			"This is a pipeline scan. It is separate from `harness_chaos_infrastructure_scan`. " +
			"Set `pipeline_identity` and `pipeline_execution_id` from a successful execution that has a Kubernetes CD stage. " +
			"The scan records scan type PipelineExecution. Creating it requires Pipeline View permission.\n\n" +
			scanLifecycleDescription(false),

		CreateContext: resourcePipelineScanCreate,
		ReadContext:   resourcePipelineScanRead,
		UpdateContext: resourcePipelineScanUpdate,
		DeleteContext: resourceChaosScanDelete,
		CustomizeDiff: resourcePipelineScanCustomizeDiff,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(2 * time.Hour),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(2 * time.Hour),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
		Importer: &schema.ResourceImporter{
			StateContext: resourcePipelineScanImport,
		},

		Schema: resourceChaosPipelineScanSchema(),
	}
}

// ResourceChaosInfrastructureScan returns the infrastructure scan resource.
func ResourceChaosInfrastructureScan() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for scanning an existing Kubernetes chaos infrastructure.\n\n" +
			"This is an infrastructure scan. It is separate from `harness_chaos_pipeline_scan`. " +
			"Set `infrastructure_id` and `environment_id` from `harness_chaos_infrastructure_v2`. Terraform reads that chaos infrastructure and sends its discovery agent. " +
			"Set `agent_identity` and `environment_identity` instead when the discovery agent is already known, including imports. That pair is what the UI sends. " +
			"The scan records scan type DiscoveryAgent. It analyses the Kubernetes resources known to that agent, not the chaos infrastructure object. Creating it requires Service Discovery View permission.\n\n" +
			scanLifecycleDescription(true),

		CreateContext: resourceInfrastructureScanCreate,
		ReadContext:   resourceInfrastructureScanRead,
		UpdateContext: resourceInfrastructureScanUpdate,
		DeleteContext: resourceChaosScanDelete,
		CustomizeDiff: resourceInfrastructureScanCustomizeDiff,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(2 * time.Hour),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(2 * time.Hour),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
		Importer: &schema.ResourceImporter{
			StateContext: resourceInfrastructureScanImport,
		},

		Schema: resourceChaosInfrastructureScanSchema(),
	}
}

func scanLifecycleDescription(infrastructure bool) string {
	retry := "The first apply creates one scan. Later applies of the same resource do not create another scan on the UI. Change `identity` only when you want a different scan. That deletes the current scan and creates the new one.\n\n"
	if infrastructure {
		retry = "The first apply creates one scan. Later applies of the same resource do not create another scan on the UI. Terraform does not watch the cluster. To run that same scan again after this configuration creates or updates the chaos infrastructure, set `retry_token` to `harness_chaos_infrastructure_v2.<name>.updated_at`. The identifier stays the same, so the UI still shows one scan and the results are refreshed. `updated_at` is read back from the API on every refresh. If the server changes it on its own, the next apply runs the scan again. A scan created before `install_command` has been applied can finish ERRORED or NO_RESOURCES_FOUND, which fails the apply. Run the scan after the infrastructure is installed and the discovery agent has seen resources. Change `identity` only when you want a different scan. That deletes the current scan and creates the new one.\n\n"
	}
	return "Create waits until the scan is COMPLETED, ERRORED, or ABORTED. ERRORED and ABORTED fail the apply and include `error_message` when the API recorded one. " +
		"COMPLETED can still have an empty `risks` list when analysis found nothing. The default create timeout is 2 hours.\n\n" +
		"Set `detection_method` to `ai` for the UI AI Scan, or `ruleBased` for the UI Rule-based Scan. `AI` and `RiskAnalyzer` are accepted as the same choices. Leave it unset and the API uses AI Scan.\n\n" +
		retry +
		"Only `name`, `description`, and `tags` can be updated in place. " +
		"`risks` is a computed view of this scan's report, sorted by identity. A report read failure keeps the previous `risks` value. Change `abort_token` on an existing scan to abort it; `ABORTED` then succeeds."
}

func resourcePipelineScanRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return readScan(ctx, d, meta, d.Id(), chaos.ScanTypePipelineExecution, true)
}

func resourceInfrastructureScanRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return readScan(ctx, d, meta, d.Id(), chaos.ScanTypeDiscoveryAgent, true)
}

func resourcePipelineScanUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return resourceChaosScanUpdate(ctx, d, meta, chaos.ScanTypePipelineExecution)
}

func resourceInfrastructureScanUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return resourceChaosScanUpdate(ctx, d, meta, chaos.ScanTypeDiscoveryAgent)
}

func resourcePipelineScanImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	return importChaosScan(ctx, d, meta, chaos.ScanTypePipelineExecution)
}

func resourceInfrastructureScanImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	return importChaosScan(ctx, d, meta, chaos.ScanTypeDiscoveryAgent)
}

func resourcePipelineScanCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)
	client := chaos.NewScanClient(c)
	scope := scopeFrom(d, c.AccountId)
	req, err := pipelineScanRequest(d)
	if err != nil {
		return diag.FromErr(err)
	}
	log.Printf("[DEBUG] Creating chaos pipeline scan %s for pipeline %s execution %s", req.Identity, req.PipelineIdentity, req.PipelineExecutionID)
	created, httpResp, err := client.CreatePipelineScan(ctx, scope, req)
	if err != nil {
		return helpers.HandleChaosApiError(err, d, httpResp)
	}
	return finishScanCreate(ctx, d, meta, client, scope, created, chaos.ScanTypePipelineExecution)
}

func resourceInfrastructureScanCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)
	client := chaos.NewScanClient(c)
	scope := scopeFrom(d, c.AccountId)
	fromInfra, useInfra, err := infrastructureScanRequest(d)
	if err != nil {
		return diag.FromErr(err)
	}
	var created *chaos.Scan
	var httpResp *http.Response
	if useInfra {
		log.Printf("[DEBUG] Creating chaos infrastructure scan %s from infrastructure %s", fromInfra.Identity, fromInfra.InfrastructureID)
		created, httpResp, err = client.CreateInfrastructureScan(ctx, scope, *fromInfra)
	} else {
		req, reqErr := discoveryAgentScanRequest(d)
		if reqErr != nil {
			return diag.FromErr(reqErr)
		}
		log.Printf("[DEBUG] Creating chaos infrastructure scan %s for agent %s", req.Identity, req.Source.Infrastructure.AgentIdentity)
		created, httpResp, err = client.Create(ctx, scope, req)
	}
	if err != nil {
		return helpers.HandleChaosApiError(err, d, httpResp)
	}
	return finishScanCreate(ctx, d, meta, client, scope, created, chaos.ScanTypeDiscoveryAgent)
}

func finishScanCreate(ctx context.Context, d *schema.ResourceData, meta interface{}, client *chaos.ScanClient, scope chaos.Scope, created *chaos.Scan, scanType string) diag.Diagnostics {
	if created == nil || created.Identity == "" {
		return diag.Errorf("create scan returned an empty identity")
	}
	d.SetId(created.Identity)
	waitDiags := waitForTerminalScan(ctx, false, func(ctx context.Context) (*chaos.Scan, *http.Response, error) {
		return client.Get(ctx, scope, created.Identity)
	})
	diags := readScan(ctx, d, meta, created.Identity, scanType, true)
	return append(diags, waitDiags...)
}

func resourceChaosScanUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}, scanType string) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)
	client := chaos.NewScanClient(c)
	scope := scopeFrom(d, c.AccountId)
	aborting := d.HasChange("abort_token") && d.Get("abort_token").(string) != ""
	retrying := d.HasChange("retry_token") && d.Get("retry_token").(string) != ""
	if aborting && retrying {
		return diag.Errorf("change abort_token or retry_token in one apply, not both")
	}

	if d.HasChanges("name", "description", "tags") {
		log.Printf("[DEBUG] Updating chaos scan %s", d.Id())
		httpResp, err := client.Update(ctx, scope, d.Id(), scanUpdateRequest(d))
		if err != nil {
			return helpers.HandleChaosApiError(err, d, httpResp)
		}
	}
	if aborting {
		log.Printf("[DEBUG] Aborting chaos scan %s", d.Id())
		httpResp, err := client.Abort(ctx, scope, d.Id())
		if err != nil {
			return helpers.HandleChaosApiError(err, d, httpResp)
		}
	}
	if retrying {
		log.Printf("[DEBUG] Retrying chaos scan %s", d.Id())
		httpResp, err := client.Retry(ctx, scope, d.Id())
		if err != nil {
			return helpers.HandleChaosApiError(err, d, httpResp)
		}
	}
	var waitDiags diag.Diagnostics
	if aborting || retrying {
		waitDiags = waitForTerminalScan(ctx, aborting, func(ctx context.Context) (*chaos.Scan, *http.Response, error) {
			return client.Get(ctx, scope, d.Id())
		})
	}
	return append(readScan(ctx, d, meta, d.Id(), scanType, true), waitDiags...)
}

func resourceChaosScanDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)
	log.Printf("[DEBUG] Deleting chaos scan %s", d.Id())
	httpResp, err := chaos.NewScanClient(c).Delete(ctx, scopeFrom(d, c.AccountId), d.Id())
	if err != nil {
		return helpers.HandleChaosReadApiError(err, d, httpResp)
	}
	d.SetId("")
	return nil
}

func importChaosScan(ctx context.Context, d *schema.ResourceData, meta interface{}, scanType string) ([]*schema.ResourceData, error) {
	orgID, projectID, identity, err := parseScanImportID(d.Id())
	if err != nil {
		return nil, err
	}
	d.SetId(identity)
	if err := d.Set("org_id", orgID); err != nil {
		return nil, err
	}
	if err := d.Set("project_id", projectID); err != nil {
		return nil, err
	}
	if err := d.Set("identity", identity); err != nil {
		return nil, err
	}
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)
	scan, httpResp, err := chaos.NewScanClient(c).Get(ctx, scopeFrom(d, c.AccountId), identity)
	if err != nil {
		if httpResp != nil {
			return nil, fmt.Errorf("read chaos scan %s: %s: %w", identity, httpResp.Status, err)
		}
		return nil, fmt.Errorf("read chaos scan %s: %w", identity, err)
	}
	if scan.ScanType != scanType {
		return nil, fmt.Errorf("scan %s is a %s scan, not a %s scan", identity, scan.ScanType, scanType)
	}
	diags := readScan(ctx, d, meta, identity, scanType, true)
	if diags.HasError() {
		return nil, fmt.Errorf("error reading chaos scan: %v", diags)
	}
	return []*schema.ResourceData{d}, nil
}

func readScan(ctx context.Context, d *schema.ResourceData, meta interface{}, identity, scanType string, keepInfrastructureIDs bool) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)
	client := chaos.NewScanClient(c)
	scope := scopeFrom(d, c.AccountId)

	log.Printf("[DEBUG] Reading chaos scan %s", identity)
	scan, httpResp, err := client.Get(ctx, scope, identity)
	if err != nil {
		return helpers.HandleChaosReadApiError(err, d, httpResp)
	}
	if scan.ScanType != scanType {
		return diag.Errorf("scan %s is a %s scan, not a %s scan", identity, scan.ScanType, scanType)
	}

	report, reportResp, reportErr := client.GetReport(ctx, scope, identity)
	var reportDiags diag.Diagnostics
	includeRisks := reportErr == nil
	if reportErr != nil {
		report = nil
		if reportResp == nil || reportResp.StatusCode != http.StatusNotFound {
			reportDiags = diag.Diagnostics{{
				Severity: diag.Warning,
				Summary:  "Chaos scan report was not read",
				Detail:   reportErr.Error() + ". The scan stays in state and the previous risks are kept.",
			}}
		}
	}

	if err := setScanState(d, scan, report, includeRisks, keepInfrastructureIDs); err != nil {
		return diag.FromErr(err)
	}
	d.SetId(scan.Identity)
	return reportDiags
}

func resourcePipelineScanCustomizeDiff(ctx context.Context, d *schema.ResourceDiff, meta interface{}) error {
	return forceNewWhenConfigured(d, "notify_id")
}

func resourceInfrastructureScanCustomizeDiff(ctx context.Context, d *schema.ResourceDiff, meta interface{}) error {
	for _, key := range []string{"agent_identity", "environment_identity", "environment_parent_unique_id"} {
		if err := forceNewWhenConfigured(d, key); err != nil {
			return err
		}
	}
	return nil
}

func forceNewWhenConfigured(d *schema.ResourceDiff, key string) error {
	// The API fills these in after create. A value appearing where config left
	// them empty must not replace the scan. A user change between two non-empty
	// values still replaces it, because update cannot change the scan source.
	if !d.HasChange(key) {
		return nil
	}
	oldValue, newValue := d.GetChange(key)
	oldText, _ := oldValue.(string)
	newText, _ := newValue.(string)
	if oldText == "" || newText == "" || oldText == newText {
		return nil
	}
	return d.ForceNew(key)
}

func waitForTerminalScan(ctx context.Context, acceptAborted bool, get func(context.Context) (*chaos.Scan, *http.Response, error)) diag.Diagnostics {
	var lastStatus string
	for {
		scan, _, err := get(ctx)
		if err != nil {
			return diag.FromErr(err)
		}
		if scan == nil {
			return diag.Errorf("read chaos scan returned an empty body")
		}
		lastStatus = scan.Status
		switch scan.Status {
		case chaos.ScanStatusCompleted:
			return nil
		case chaos.ScanStatusAborted:
			if acceptAborted {
				return nil
			}
			fallthrough
		case chaos.ScanStatusErrored:
			message := scan.ErrorMessage
			if message == "" {
				message = "the scan ended without an error message"
			}
			return diag.Errorf("chaos scan %s is %s: %s", scan.Identity, scan.Status, message)
		}

		timer := time.NewTimer(scanPollInterval)
		select {
		case <-ctx.Done():
			timer.Stop()
			return diag.Errorf("timed out waiting for chaos scan %s to finish; last status %s", scan.Identity, lastStatus)
		case <-timer.C:
		}
	}
}
