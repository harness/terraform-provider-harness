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

func ResourceChaosServiceOnboarding() *schema.Resource {
	return &schema.Resource{
		Description: "Onboards one discovered service onto a chaos infrastructure.\n\n" +
			"The first apply registers the session and returns `install_command`. Install that manifest, then apply again to create the chaos service. " +
			"Set `service_name` or `discovered_service_id`. `probes` lists existing project probes. " +
			"Another service on the same session uses `harness_chaos_onboarding_service`.",

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
				Description: "CD infrastructure identifier. The onboarding API registers a chaos infrastructure with this id.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"service_name": {
				Description:   "Name of the discovered service. Set this or `discovered_service_id`.",
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"discovered_service_id"},
				AtLeastOneOf:  []string{"service_name", "discovered_service_id"},
			},
			"discovered_service_id": {
				Description:   "Id of the discovered service. Set this or `service_name`.",
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"service_name"},
				AtLeastOneOf:  []string{"service_name", "discovered_service_id"},
			},
			"probes": probeSchema(),
			"install_command": {
				Description: "Install manifest returned by the first apply.",
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
				Description: "Status of the discovery step.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"service_status": {
				Description: "Status of the chaos service step.",
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
	// The next apply creates the chaos service after the manifest is installed.
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
	if id := d.Get("chaos_service_id").(string); id != "" {
		_, _, getErr := client.GetChaosService(ctx, scope, id)
		if getErr != nil && strings.Contains(getErr.Error(), "requested resource not found") {
			session.ChaosServiceID = ""
		} else {
			session.ChaosServiceID = id
		}
	}
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
		// The chaos service is created by onboarding. A saved id that the API no
		// longer has is created again, with this probe list, instead of updated.
		if err != nil && strings.Contains(err.Error(), "requested resource not found") {
			session, httpResp, getErr := client.Get(ctx, scope, d.Id())
			if getErr != nil {
				return helpers.HandleChaosApiError(getErr, d, httpResp)
			}
			finished, httpResp, createErr := client.FinishServiceOnboarding(ctx, scope, session, onboardingRequest(d))
			if createErr != nil {
				return helpers.HandleChaosApiError(createErr, d, httpResp)
			}
			return setServiceOnboarding(d, finished)
		}
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
	fields["chaos_service_id"] = session.ChaosServiceID
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
		Description: "Project probes to associate with the chaos service.",
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
