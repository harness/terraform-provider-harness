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

func ResourceChaosScan() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for creating a Harness Chaos scan.\n\n" +
			"A scan is one object. Set `pipeline_scan` to scan a successful pipeline execution, which is the execution the UI lists. " +
			"Set `infrastructure_scan` with `infrastructure_id` and `environment_id` from `harness_chaos_infrastructure_v2` to scan that Kubernetes infrastructure. Terraform reads the linked discovery agent and sends the same agent, environment, and parent id as the UI Scan infrastructure action. " +
			"`agent_identity` and `environment_identity` remain supported when the discovery agent is already known, including imports. It is not a Harness service. Collection fails when a pipeline execution has no Kubernetes CD stage, or when the infrastructure has no discovery agent.\n\n" +
			"Create waits until the scan is COMPLETED, ERRORED, or ABORTED. ERRORED and ABORTED fail the apply and include `error_message` when the API recorded one. " +
			"COMPLETED can still have an empty `risks` list when analysis found nothing. The default create timeout is 30 minutes.\n\n" +
			"Set `detection_method` to `ai` for the UI AI Scan, or `ruleBased` for the UI Rule-based Scan. `AI` and `RiskAnalyzer` are accepted as the same choices. Leave it unset and the API uses AI Scan.\n\n" +
			"The first apply creates one scan. Later applies of the same resource do not create another scan on the UI. Terraform does not watch the cluster. To run that same scan again after this configuration creates or updates the chaos infrastructure, set `retry_token` to `harness_chaos_infrastructure_v2.<name>.updated_at`. The identifier stays the same, so the UI still shows one scan and the results are refreshed. Change `identity` only when you want a different scan. That deletes the current scan and creates the new one.\n\n" +
			"Only `name`, `description`, and `tags` can be updated in place. Changing identity, organization, project, detection method, pipeline identity, pipeline execution, chaos infrastructure, agent identity, or environment identity replaces the scan. " +
			"Values the API writes back (`infrastructures`, `notify_id`, `agent_identity` when it was resolved, `environment_parent_unique_id`) do not replace it.\n\n" +
			"`risks` is a computed view of this scan's report, sorted by identity. It is not the durable scanned risk. A report read failure keeps the previous `risks` value. Change `abort_token` on an existing scan to abort it; `ABORTED` then succeeds. Create holds the state lock, so another apply cannot abort a create that is still waiting.",

		CreateContext: resourceChaosScanCreate,
		ReadContext:   resourceChaosScanRead,
		UpdateContext: resourceChaosScanUpdate,
		DeleteContext: resourceChaosScanDelete,
		CustomizeDiff: resourceChaosScanCustomizeDiff,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
		Importer: &schema.ResourceImporter{
			StateContext: resourceChaosScanImport,
		},
		Schema: resourceChaosScanSchema(),
	}
}

func resourceChaosScanCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)
	client := chaos.NewScanClient(c)
	scope := scopeFrom(d, c.AccountId)
	fromInfra, useInfra, err := infrastructureScanFromResource(d)
	if err != nil {
		return diag.FromErr(err)
	}
	var created *chaos.Scan
	var httpResp *http.Response
	if useInfra {
		log.Printf("[DEBUG] Creating chaos infrastructure scan %s from infrastructure %s", fromInfra.Identity, fromInfra.InfrastructureID)
		created, httpResp, err = client.CreateInfrastructureScan(ctx, scope, *fromInfra)
	} else {
		req, reqErr := scanCreateRequest(d)
		if reqErr != nil {
			return diag.FromErr(reqErr)
		}
		log.Printf("[DEBUG] Creating chaos scan %s (%s)", req.Identity, req.ScanType)
		created, httpResp, err = client.Create(ctx, scope, req)
	}
	if err != nil {
		return helpers.HandleChaosApiError(err, d, httpResp)
	}
	if created == nil || created.Identity == "" {
		return diag.Errorf("create scan returned an empty identity")
	}
	d.SetId(created.Identity)
	waitDiags := waitForTerminalScan(ctx, false, func(ctx context.Context) (*chaos.Scan, *http.Response, error) {
		return client.Get(ctx, scope, created.Identity)
	})
	diags := readScan(ctx, d, meta, created.Identity)
	return append(diags, waitDiags...)
}

func resourceChaosScanRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return readScan(ctx, d, meta, d.Id())
}

func resourceChaosScanUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	return append(readScan(ctx, d, meta, d.Id()), waitDiags...)
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

func resourceChaosScanImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
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
	diags := readScan(ctx, d, meta, identity)
	if diags.HasError() {
		return nil, fmt.Errorf("error reading chaos scan: %v", diags)
	}
	return []*schema.ResourceData{d}, nil
}

func readScan(ctx context.Context, d *schema.ResourceData, meta interface{}, identity string) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)
	client := chaos.NewScanClient(c)
	scope := scopeFrom(d, c.AccountId)

	log.Printf("[DEBUG] Reading chaos scan %s", identity)
	scan, httpResp, err := client.Get(ctx, scope, identity)
	if err != nil {
		return helpers.HandleChaosReadApiError(err, d, httpResp)
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

	if err := setScanState(d, scan, report, includeRisks); err != nil {
		return diag.FromErr(err)
	}
	d.SetId(scan.Identity)
	return reportDiags
}

func resourceChaosScanCustomizeDiff(_ context.Context, d *schema.ResourceDiff, _ interface{}) error {
	// The API fills these in after create. A value appearing where config left
	// them empty must not replace the scan. A user change between two non-empty
	// values still replaces it, because update cannot change the scan source.
	for _, key := range []string{
		"pipeline_scan.0.notify_id",
		"infrastructure_scan.0.agent_identity",
		"infrastructure_scan.0.environment_identity",
		"infrastructure_scan.0.environment_parent_unique_id",
	} {
		if !d.HasChange(key) {
			continue
		}
		oldValue, newValue := d.GetChange(key)
		oldText, _ := oldValue.(string)
		newText, _ := newValue.(string)
		if oldText == "" || newText == "" || oldText == newText {
			continue
		}
		if err := d.ForceNew(key); err != nil {
			return err
		}
	}
	return nil
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
