package onboarding

import (
	"context"
	"fmt"
	"log"
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
		Description: "Onboards one discovered service onto a chaos infrastructure and associates probes that already exist in the project.\n\n" +
			"Creates a manual onboarding session with method Service. `infrastructure_id` is the CD infrastructure id. " +
			"That CD infrastructure must already exist. The onboarding API then registers a chaos infrastructure with the same id, so that chaos infrastructure must not already exist. " +
			"Create requires Chaos Experiment Edit, Chaos Experiment Execute, Network Map Create, and Network Map Edit. " +
			"Reading `install_command` also requires Chaos Infrastructure View.\n\n" +
			"The first apply registers the session and returns `install_command`, the same manifest `harness_chaos_infrastructure_v2` installs with. " +
			"It does not create the chaos service yet, because discovery cannot finish until that manifest is installed in the cluster. " +
			"Install it, then apply again. The next apply confirms discovery and creates one chaos service for `service_name` or `discovered_service_id`. " +
			"`probes` lists existing project probe ids to associate with that service. Several probes are valid, and the same probe cannot be added twice. " +
			"An empty list associates no probes. Later applies update that probe list in place. Changing `service_name` or `discovered_service_id` creates a new resource. " +
			"The risk-scan step is left queued, so the session status stays PENDING.\n\n" +
			"This resource owns the onboarding session for the infrastructure. Create returns an error when that session already exists; import `org_id/project_id/onboarding_id` to manage it. " +
			"Another discovered service on the same infrastructure uses `harness_chaos_onboarding_service` with this `onboarding_id`, which creates the chaos service and leaves the session in place. " +
			"Destroy deletes the chaos service created here, then marks the session removed. " +
			"The chaos infrastructure and the discovery agent remain, so the same `infrastructure_id` cannot be onboarded again.",

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
			"service_name": {
				Description:   "Name of the discovered service, as shown in the service list, such as `bank-of-anthos-config`. Set this or `discovered_service_id`. When several services share the name, set `discovered_service_id` instead.",
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"discovered_service_id"},
				AtLeastOneOf:  []string{"service_name", "discovered_service_id"},
			},
			"discovered_service_id": {
				Description:   "Id shown under the discovered service name, such as `bankofanthosconfig`. The discovery record id is also accepted. Set this or `service_name`. Changing it creates a new resource.",
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"service_name"},
				AtLeastOneOf:  []string{"service_name", "discovered_service_id"},
			},
			"probes": probeSchema(),
			"install_command": {
				Description: "Manifest for the chaos infrastructure this session registers. The first apply returns this value. Install it in the cluster, then apply again to confirm discovery and create the chaos service.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"chaos_service_id": {
				Description: "Identity of the chaos service created for this discovered service.",
				Type:        schema.TypeString,
				Computed:    true,
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
				Description: "Status of the discovery step. The chaos service can be created when this is COMPLETED.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"service_status": {
				Description: "Status of the chaos service onboarding step. Creating the service through the chaos service API leaves this step queued.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func resourceServiceOnboardingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)
	client := chaos.NewOnboardingClient(c)
	scope := onboardingScope(d, c.AccountId)
	req := onboardingRequest(d)
	log.Printf("[DEBUG] Creating chaos service onboarding for infrastructure %s", req.InfrastructureID)
	existing, httpResp, err := client.FindServiceOnboarding(ctx, scope, req.InfrastructureID)
	if err != nil {
		return helpers.HandleChaosApiError(err, d, httpResp)
	}
	if existing != nil && existing.OnboardingID != "" {
		return diag.Errorf(
			"infrastructure %s already has onboarding session %s. Import %s/%s/%s to manage that session. Another discovered service uses harness_chaos_onboarding_service with onboarding_id %s.",
			req.InfrastructureID, existing.OnboardingID, d.Get("org_id"), d.Get("project_id"), existing.OnboardingID, existing.OnboardingID,
		)
	}
	opened, httpResp, err := client.OpenServiceOnboarding(ctx, scope, req)
	if opened != nil && opened.OnboardingID != "" {
		d.SetId(opened.OnboardingID)
	}
	if err != nil {
		diags := helpers.HandleChaosApiError(err, d, httpResp)
		if opened != nil && opened.OnboardingID != "" {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary: fmt.Sprintf(
					"onboarding session %s exists for this infrastructure. Import %s/%s/%s to manage that session.",
					opened.OnboardingID, d.Get("org_id"), d.Get("project_id"), opened.OnboardingID,
				),
			})
		}
		return diags
	}
	// Discovery cannot finish until the user installs the manifest outside this
	// apply. Returning here saves the session and install_command. The next
	// apply creates the chaos service. Waiting inside create would time out and
	// Terraform would discard the session, including the manifest.
	return setServiceOnboarding(d, opened)
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
	session.ChaosServiceID = d.Get("chaos_service_id").(string)
	if session.InstallCommand == "" {
		if manifest, manifestErr := client.InfrastructureManifest(ctx, scope, session.InfrastructureRef, session.EnvironmentRef); manifestErr == nil {
			session.InstallCommand = manifest
		}
	}
	return setServiceOnboarding(d, session)
}

func resourceServiceOnboardingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)
	client := chaos.NewOnboardingClient(c)
	scope := onboardingScope(d, c.AccountId)
	if d.Get("chaos_service_id").(string) == "" {
		session, httpResp, err := client.Get(ctx, scope, d.Id())
		if err != nil {
			return helpers.HandleChaosApiError(err, d, httpResp)
		}
		finished, httpResp, err := client.FinishServiceOnboarding(ctx, scope, session, onboardingRequest(d))
		if err != nil {
			return helpers.HandleChaosApiError(err, d, httpResp)
		}
		return setServiceOnboarding(d, finished)
	}
	if d.HasChange("probes") {
		updated, httpResp, err := client.UpdateChaosService(ctx, scope, chaosServiceRequest(d))
		if err != nil {
			return helpers.HandleChaosApiError(err, d, httpResp)
		}
		if err := d.Set("chaos_service_id", updated.Identity); err != nil {
			return diag.FromErr(err)
		}
	}
	return resourceServiceOnboardingRead(ctx, d, meta)
}

func resourceServiceOnboardingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)
	client := chaos.NewOnboardingClient(c)
	scope := onboardingScope(d, c.AccountId)
	if identity := d.Get("chaos_service_id").(string); identity != "" {
		httpResp, err := client.DeleteChaosService(ctx, scope, identity)
		if err != nil && !chaos.OnboardingGone(err) {
			return helpers.HandleChaosApiError(err, d, httpResp)
		}
	}
	httpResp, err := client.Delete(ctx, scope, d.Id())
	if err != nil && !chaos.OnboardingGone(err) {
		return helpers.HandleChaosApiError(err, d, httpResp)
	}
	log.Printf("[WARN] Chaos onboarding %s was marked removed. The chaos infrastructure and discovery agent remain.", d.Id())
	d.SetId("")
	return nil
}

func resourceServiceOnboardingCustomizeDiff(ctx context.Context, d *schema.ResourceDiff, meta interface{}) error {
	if err := duplicateProbeIDs(d.Get("probes")); err != nil {
		return err
	}
	if d.Id() != "" && d.Get("chaos_service_id").(string) == "" {
		return d.SetNewComputed("chaos_service_id")
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

func setServiceOnboarding(d *schema.ResourceData, session *chaos.ServiceOnboarding) diag.Diagnostics {
	if session == nil {
		return diag.Errorf("service onboarding response was empty")
	}
	d.SetId(session.OnboardingID)
	fields := map[string]interface{}{
		"onboarding_id":    session.OnboardingID,
		"status":           session.Status,
		"discovery_status": session.DiscoveryStatus,
		"service_status":   session.ServiceStatus,
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
	if session.ChaosServiceID != "" {
		fields["chaos_service_id"] = session.ChaosServiceID
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

func probeSchema() *schema.Schema {
	return &schema.Schema{
		Description: "Project probes to associate with the chaos service. The same probe cannot be added twice. Leave this empty to associate no probes.",
		Type:        schema.TypeList,
		Optional:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"probe_id": {
					Description: "Identifier of a probe that already exists in the project.",
					Type:        schema.TypeString,
					Required:    true,
				},
				"inputs": {
					Description: "Runtime input values for the probe, keyed by input name.",
					Type:        schema.TypeMap,
					Optional:    true,
					Elem:        &schema.Schema{Type: schema.TypeString},
				},
			},
		},
	}
}

func onboardingRequest(d *schema.ResourceData) chaos.ServiceOnboardingRequest {
	return chaos.ServiceOnboardingRequest{
		EnvironmentID:    d.Get("environment_id").(string),
		InfrastructureID: d.Get("infrastructure_id").(string),
		ServiceID:        d.Get("discovered_service_id").(string),
		ServiceName:      d.Get("service_name").(string),
		Probes:           probesFrom(d.Get("probes")),
	}
}

func chaosServiceRequest(d *schema.ResourceData) chaos.ChaosServiceRequest {
	return chaos.ChaosServiceRequest{
		Identity:          d.Get("chaos_service_id").(string),
		EnvironmentID:     d.Get("environment_id").(string),
		InfrastructureID:  d.Get("infrastructure_id").(string),
		ExternalServiceID: d.Get("discovered_service_id").(string),
		Probes:            probesFrom(d.Get("probes")),
	}
}

func probesFrom(raw interface{}) []chaos.ServiceProbe {
	items, ok := raw.([]interface{})
	if !ok {
		return nil
	}
	out := make([]chaos.ServiceProbe, 0, len(items))
	for _, item := range items {
		block, ok := item.(map[string]interface{})
		if !ok {
			continue
		}
		probe := chaos.ServiceProbe{ProbeID: block["probe_id"].(string)}
		if inputs, ok := block["inputs"].(map[string]interface{}); ok && len(inputs) > 0 {
			probe.Inputs = map[string]string{}
			for name, value := range inputs {
				probe.Inputs[name] = value.(string)
			}
		}
		out = append(out, probe)
	}
	return out
}

func duplicateProbeIDs(raw interface{}) error {
	seen := map[string]struct{}{}
	for _, probe := range probesFrom(raw) {
		if probe.ProbeID == "" {
			continue
		}
		if _, ok := seen[probe.ProbeID]; ok {
			return fmt.Errorf("probe %s is already associated", probe.ProbeID)
		}
		seen[probe.ProbeID] = struct{}{}
	}
	return nil
}

func parseOnboardingImportID(id string) (string, string, string, error) {
	parts := strings.Split(id, "/")
	if len(parts) != 3 || parts[0] == "" || parts[1] == "" || parts[2] == "" {
		return "", "", "", fmt.Errorf("import id must be org_id/project_id/id")
	}
	return parts[0], parts[1], parts[2], nil
}
