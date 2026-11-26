package onboarding

import (
	"context"
	"time"

	"github.com/harness/harness-go-sdk/harness/chaos"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceChaosOnboardingService() *schema.Resource {
	return &schema.Resource{
		Description: "Creates one chaos service on an onboarding session that already exists.\n\n" +
			"`onboarding_id` comes from `harness_chaos_service_onboarding`. Set `service_name` or `discovered_service_id`. " +
			"Destroy deletes the chaos service and leaves the session in place.",

		CreateContext: resourceOnboardingServiceCreate,
		ReadContext:   resourceOnboardingServiceRead,
		UpdateContext: resourceOnboardingServiceUpdate,
		DeleteContext: resourceOnboardingServiceDelete,
		CustomizeDiff: resourceOnboardingServiceCustomizeDiff,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
		Importer: &schema.ResourceImporter{
			StateContext: resourceOnboardingServiceImport,
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
				Description: "Infrastructure identifier of the onboarding session.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"onboarding_id": {
				Description: "Onboarding session id.",
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
			"chaos_service_id": {
				Description: "Identity of the chaos service.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func resourceOnboardingServiceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)
	client := chaos.NewOnboardingClient(c)
	created, httpResp, err := client.CreateChaosService(ctx, onboardingScope(d, c.AccountId), d.Get("onboarding_id").(string), onboardingRequest(d))
	if err != nil {
		return helpers.HandleChaosApiError(err, d, httpResp)
	}
	if created == nil || created.Identity == "" {
		return diag.Errorf("create chaos service returned an empty identity")
	}
	d.SetId(created.Identity)
	if err := d.Set("chaos_service_id", created.Identity); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceOnboardingServiceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)
	_, httpResp, err := chaos.NewOnboardingClient(c).GetChaosService(ctx, onboardingScope(d, c.AccountId), d.Id())
	if err != nil {
		if chaos.OnboardingGone(err) || (httpResp != nil && httpResp.StatusCode == 404) {
			d.SetId("")
			return nil
		}
		return helpers.HandleChaosReadApiError(err, d, httpResp)
	}
	if err := d.Set("chaos_service_id", d.Id()); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceOnboardingServiceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if d.HasChange("probes") {
		c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)
		if err := d.Set("chaos_service_id", d.Id()); err != nil {
			return diag.FromErr(err)
		}
		if _, httpResp, err := chaos.NewOnboardingClient(c).UpdateChaosService(ctx, onboardingScope(d, c.AccountId), chaosServiceRequest(d)); err != nil {
			return helpers.HandleChaosApiError(err, d, httpResp)
		}
	}
	return resourceOnboardingServiceRead(ctx, d, meta)
}

func resourceOnboardingServiceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)
	httpResp, err := chaos.NewOnboardingClient(c).DeleteChaosService(ctx, onboardingScope(d, c.AccountId), d.Id())
	if err != nil && !chaos.OnboardingGone(err) {
		return helpers.HandleChaosApiError(err, d, httpResp)
	}
	d.SetId("")
	return nil
}

func resourceOnboardingServiceCustomizeDiff(ctx context.Context, d *schema.ResourceDiff, meta interface{}) error {
	return duplicateProbeIDs(d.Get("probes"))
}

func resourceOnboardingServiceImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	orgID, projectID, identity, err := parseOnboardingImportID(d.Id())
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
	return []*schema.ResourceData{d}, nil
}
