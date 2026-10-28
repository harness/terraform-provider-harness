package onboarding

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/harness/harness-go-sdk/harness/chaos"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceChaosServiceOnboarding returns the service onboarding resource.
func ResourceChaosServiceOnboarding() *schema.Resource {
	return &schema.Resource{
		Description: "Onboards discovered services onto a chaos infrastructure.\n\n" +
			"Creates a manual onboarding session with method Service. `infrastructure_id` is the CD infrastructure id. " +
			"That CD infrastructure must already exist. The onboarding API then registers a chaos infrastructure with the same id, so that chaos infrastructure must not already exist. " +
			"Create requires Chaos Experiment Edit, Chaos Experiment Execute, Network Map Create, and Network Map Edit. " +
			"Reading `install_command` also requires Chaos Infrastructure View. " +
			"Onboard and confirm require the ChaosRTServicesEnabled feature flag.\n\n" +
			"The first apply registers the session and returns `install_command`, the same manifest `harness_chaos_infrastructure_v2` installs with. " +
			"It does not onboard services yet, because discovery cannot finish until that manifest is installed in the cluster. " +
			"Install it, then apply again. The next apply confirms discovery, onboards the discovered services, and confirms that step. " +
			"Adding services later confirms that step again. The risk-scan step is left queued, so the session status stays PENDING.\n\n" +
			"Omit `service_ids` to onboard every service discovery finds. Set it to onboard a subset. Those values are the discovered service ids. " +
			"`service_ids` is add-only. Removing an id does not delete the chaos service, and clearing it is rejected.\n\n" +
			"A retry reuses the session already created for this infrastructure. Destroy marks the session removed. " +
			"The chaos infrastructure and the discovery agent remain, so the same `infrastructure_id` cannot be onboarded again. " +
			"The next apply tries to register that chaos infrastructure and fails.",

		CreateContext: resourceServiceOnboardingCreate,
		ReadContext:   resourceServiceOnboardingRead,
		UpdateContext: resourceServiceOnboardingUpdate,
		DeleteContext: resourceServiceOnboardingDelete,
		CustomizeDiff: resourceServiceOnboardingCustomizeDiff,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(time.Hour),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(time.Hour),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
		Importer: &schema.ResourceImporter{
			StateContext: resourceServiceOnboardingImport,
		},
		Schema: map[string]*schema.Schema{
			"org_id": {
				Description: "Organization identifier.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"project_id": {
				Description: "Project identifier.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"environment_id": {
				Description: "Environment the infrastructure belongs to.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"infrastructure_id": {
				Description: "CD infrastructure identifier. It must already exist. The onboarding API registers a chaos infrastructure with this same id, so that chaos infrastructure must not already exist.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"service_ids": {
				Description: "Discovered service ids to onboard. Omit this to onboard every service discovery finds. Add-only: removing an id does not delete the chaos service, and clearing this argument is rejected.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"install_command": {
				Description: "Manifest for the chaos infrastructure this session registers. The first apply returns this value. Install it in the cluster, then apply again to confirm discovery and onboard services.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"onboarded_service_ids": {
				Description: "Discovered service ids onboarded by this session.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"onboarding_id": {
				Description: "Onboarding session id.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"status": {
				Description: "Onboarding session status.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"discovery_status": {
				Description: "Status of the discovery step. Services can be onboarded when this is COMPLETED.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"service_status": {
				Description: "Status of the chaos service onboarding step.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"onboarded_count": {
				Description: "Number of services onboarded by this session.",
				Type:        schema.TypeInt,
				Computed:    true,
			},
		},
	}
}

func resourceServiceOnboardingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)
	client := chaos.NewOnboardingClient(c)
	scope := onboardingScope(d, c.AccountId)
	req := chaos.ServiceOnboardingRequest{
		EnvironmentID:    d.Get("environment_id").(string),
		InfrastructureID: d.Get("infrastructure_id").(string),
		ServiceIDs:       serviceIDs(d),
	}
	log.Printf("[DEBUG] Creating chaos service onboarding for infrastructure %s", req.InfrastructureID)
	opened, httpResp, err := client.OpenServiceOnboarding(ctx, scope, req)
	if opened != nil && opened.OnboardingID != "" {
		d.SetId(opened.OnboardingID)
	}
	if err != nil {
		return onboardingCreateError(d, opened, httpResp, err)
	}
	// Discovery cannot finish until the user installs the manifest outside this
	// apply. Returning here saves the session and install_command. The next
	// apply onboards the services. Waiting inside create would time out and
	// Terraform would discard the session, including the manifest.
	return setServiceOnboarding(d, opened, req.ServiceIDs)
}

func onboardingCreateError(d *schema.ResourceData, session *chaos.ServiceOnboarding, httpResp *http.Response, err error) diag.Diagnostics {
	diags := helpers.HandleChaosApiError(err, d, httpResp)
	if session == nil || session.OnboardingID == "" {
		return diags
	}
	diags = append(diags, diag.Diagnostic{
		Severity: diag.Error,
		Summary: fmt.Sprintf(
			"onboarding session %s exists for this infrastructure. Retry reuses it. Import %s/%s/%s if this resource is not in state.",
			session.OnboardingID, d.Get("org_id"), d.Get("project_id"), session.OnboardingID,
		),
	})
	return diags
}

func resourceServiceOnboardingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)
	client := chaos.NewOnboardingClient(c)
	scope := onboardingScope(d, c.AccountId)
	session, httpResp, err := client.Get(ctx, scope, d.Id())
	if err != nil {
		if chaos.OnboardingGone(err) {
			log.Printf("[WARN] Chaos onboarding %s no longer exists, removing from state", d.Id())
			d.SetId("")
			return nil
		}
		return helpers.HandleChaosReadApiError(err, d, httpResp)
	}
	if session.ManualMethod != "" && session.ManualMethod != "Service" {
		return diag.Errorf("onboarding %s uses method %s, not Service", d.Id(), session.ManualMethod)
	}
	if ids, listErr := client.OnboardedServiceIDs(ctx, scope, d.Id()); listErr == nil {
		session.OnboardedServiceIDs = ids
	}
	if session.InstallCommand == "" {
		if manifest, manifestErr := client.InfrastructureManifest(ctx, scope, session.InfrastructureRef, session.EnvironmentRef); manifestErr == nil {
			session.InstallCommand = manifest
		}
	}
	return setServiceOnboarding(d, session, nil)
}

func resourceServiceOnboardingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)
	client := chaos.NewOnboardingClient(c)
	scope := onboardingScope(d, c.AccountId)
	session, httpResp, err := client.Get(ctx, scope, d.Id())
	if err != nil {
		return helpers.HandleChaosApiError(err, d, httpResp)
	}
	ids := serviceIDs(d)
	if session.ServiceStatus != "COMPLETED" {
		finished, httpResp, err := client.FinishServiceOnboarding(ctx, scope, session, ids)
		if err != nil {
			return helpers.HandleChaosApiError(err, d, httpResp)
		}
		return setServiceOnboarding(d, finished, nil)
	}
	if d.HasChange("service_ids") {
		added := addedServiceIDs(d)
		if len(added) > 0 {
			if _, httpResp, err = client.OnboardServices(ctx, scope, d.Id(), added); err != nil {
				return helpers.HandleChaosApiError(err, d, httpResp)
			}
			// Onboard sets the service step back to RUNNING. Confirm it in this apply.
			if httpResp, err = client.ConfirmServiceOnboarding(ctx, scope, d.Id()); err != nil {
				return helpers.HandleChaosApiError(err, d, httpResp)
			}
		}
	}
	return resourceServiceOnboardingRead(ctx, d, meta)
}

func resourceServiceOnboardingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)
	httpResp, err := chaos.NewOnboardingClient(c).Delete(ctx, onboardingScope(d, c.AccountId), d.Id())
	if err != nil && !chaos.OnboardingGone(err) {
		return helpers.HandleChaosApiError(err, d, httpResp)
	}
	log.Printf("[WARN] Chaos onboarding %s was marked removed. The chaos infrastructure and discovery agent remain.", d.Id())
	d.SetId("")
	return nil
}

func resourceServiceOnboardingCustomizeDiff(ctx context.Context, d *schema.ResourceDiff, meta interface{}) error {
	if d.Id() != "" && d.Get("service_status").(string) != "COMPLETED" {
		if err := d.SetNewComputed("service_status"); err != nil {
			return err
		}
	}
	previous, next := d.GetChange("service_ids")
	oldIDs := setIDs(previous)
	newIDs := setIDs(next)
	if len(oldIDs) == 0 {
		return nil
	}
	if len(newIDs) == 0 {
		return fmt.Errorf("clearing service_ids would onboard every service discovery currently finds")
	}
	for id := range oldIDs {
		if !newIDs[id] {
			return fmt.Errorf("service_ids is add-only; removing %s does not delete the chaos service", id)
		}
	}
	return nil
}

func resourceServiceOnboardingImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	orgID, projectID, onboardingID, err := parseOnboardingImportID(d.Id())
	if err != nil {
		return nil, err
	}
	d.SetId(onboardingID)
	if err := d.Set("org_id", orgID); err != nil {
		return nil, err
	}
	if err := d.Set("project_id", projectID); err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}

func setServiceOnboarding(d *schema.ResourceData, session *chaos.ServiceOnboarding, configured []string) diag.Diagnostics {
	if session == nil {
		return diag.Errorf("service onboarding response was empty")
	}
	d.SetId(session.OnboardingID)
	fields := map[string]interface{}{
		"onboarding_id":    session.OnboardingID,
		"status":           session.Status,
		"discovery_status": session.DiscoveryStatus,
		"service_status":   session.ServiceStatus,
		"onboarded_count":  session.OnboardedCount,
	}
	if session.InstallCommand != "" {
		fields["install_command"] = session.InstallCommand
	}
	if session.EnvironmentRef != "" {
		fields["environment_id"] = session.EnvironmentRef
	}
	if session.InfrastructureRef != "" {
		fields["infrastructure_id"] = session.InfrastructureRef
	}
	if configured != nil {
		fields["service_ids"] = configured
	}
	if session.OnboardedServiceIDs != nil {
		fields["onboarded_service_ids"] = session.OnboardedServiceIDs
	}
	for key, value := range fields {
		if err := d.Set(key, value); err != nil {
			return diag.FromErr(err)
		}
	}
	return nil
}

func onboardingScope(d *schema.ResourceData, accountID string) chaos.Scope {
	return chaos.Scope{
		AccountID: accountID,
		OrgID:     d.Get("org_id").(string),
		ProjectID: d.Get("project_id").(string),
	}
}

func serviceIDs(d *schema.ResourceData) []string {
	return sortedIDs(setIDs(d.Get("service_ids")))
}

func addedServiceIDs(d *schema.ResourceData) []string {
	previous, _ := d.GetChange("service_ids")
	oldIDs := setIDs(previous)
	var added []string
	for id := range setIDs(d.Get("service_ids")) {
		if !oldIDs[id] {
			added = append(added, id)
		}
	}
	sort.Strings(added)
	return added
}

func setIDs(raw interface{}) map[string]bool {
	ids := map[string]bool{}
	set, ok := raw.(*schema.Set)
	if !ok || set == nil {
		return ids
	}
	for _, value := range set.List() {
		ids[value.(string)] = true
	}
	return ids
}

func sortedIDs(ids map[string]bool) []string {
	out := make([]string, 0, len(ids))
	for id := range ids {
		out = append(out, id)
	}
	sort.Strings(out)
	return out
}

func parseOnboardingImportID(id string) (string, string, string, error) {
	parts := strings.Split(id, "/")
	if len(parts) != 3 || parts[0] == "" || parts[1] == "" || parts[2] == "" {
		return "", "", "", fmt.Errorf("import id must be org_id/project_id/onboarding_id")
	}
	return parts[0], parts[1], parts[2], nil
}
