package agent

import (
	"context"
	"fmt"
	"time"

	"github.com/antihax/optional"
	hh "github.com/harness/harness-go-sdk/harness/helpers"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceGitopsAgentDrPromotion() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for promoting a Disaster Recovery secondary GitOps agent node to primary. " +
			"Deleting this resource is a no-op; change `triggers` to repeat a promotion.",

		CreateContext: resourceGitopsAgentDrPromotionCreate,
		ReadContext:   resourceGitopsAgentDrPromotionRead,
		DeleteContext: resourceGitopsAgentDrPromotionDelete,

		// TODO(CDS-135089): verify delayTaskCreationInSec on qa/prod before hardcoding
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(65 * time.Second),
		},

		Schema: map[string]*schema.Schema{
			"org_id": {
				Description: "Organization identifier of the GitOps agent.",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
			},
			"project_id": {
				Description: "Project identifier of the GitOps agent.",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
			},
			"agent_id": {
				Description: "Identifier of the GitOps agent to promote.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"dr_identifier": {
				Description: "Disaster Recovery identifier to promote to PRIMARY.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"disable_sync_windows": {
				Description: "Disable sync windows for the new primary agent.",
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
			},
			"triggers": {
				Description: "Map of arbitrary string values. Any change to this map forces replacement of the " +
					"resource, causing the promotion to be repeated.",
				Type:     schema.TypeMap,
				Optional: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"current_primary": {
				Description: "The dr_identifier of the DR node currently reported as PRIMARY by the server.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func resourceGitopsAgentDrPromotionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	ctx = context.WithValue(ctx, nextgen.ContextAccessToken, hh.EnvVars.BearerToken.Get())

	orgId := d.Get("org_id").(string)
	projectId := d.Get("project_id").(string)
	agentId := d.Get("agent_id").(string)
	drIdentifier := d.Get("dr_identifier").(string)

	req := nextgen.V1AgentSetPrimaryNodeRequest{
		AccountIdentifier:          c.AccountId,
		OrgIdentifier:              orgId,
		ProjectIdentifier:          projectId,
		AgentIdentifier:            agentId,
		DisasterRecoveryIdentifier: drIdentifier,
		DisableSyncWindows:         d.Get("disable_sync_windows").(bool),
	}

	_, httpResp, err := c.AgentApi.AgentServiceForServerSetPrimaryNode(ctx, req, agentId)
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	d.SetId(fmt.Sprintf("%s/%s/%s", orgId, projectId, agentId))
	return resourceGitopsAgentDrPromotionRead(ctx, d, meta)
}

func resourceGitopsAgentDrPromotionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	ctx = context.WithValue(ctx, nextgen.ContextAccessToken, hh.EnvVars.BearerToken.Get())

	orgId := d.Get("org_id").(string)
	projectId := d.Get("project_id").(string)
	agentId := d.Get("agent_id").(string)
	drIdentifier := d.Get("dr_identifier").(string)

	resp, httpResp, err := c.AgentApi.AgentServiceForServerGet(ctx, agentId, c.AccountId, &nextgen.AgentsApiAgentServiceForServerGetOpts{
		OrgIdentifier:     optional.NewString(orgId),
		ProjectIdentifier: optional.NewString(projectId),
		DrIdentifier:      optional.NewString(drIdentifier),
	})
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	if resp.DisasterRecoveryNode != nil && resp.DisasterRecoveryNode.Type_ != nil &&
		*resp.DisasterRecoveryNode.Type_ == nextgen.PRIMARY_V1DisasterRecoveryType {
		d.Set("current_primary", drIdentifier)
		return nil
	}

	listResp, listHttpResp, err := c.AgentApi.AgentServiceForServerList(ctx, c.AccountId, string(*resp.Type_), &nextgen.AgentsApiAgentServiceForServerListOpts{
		OrgIdentifier:     optional.NewString(orgId),
		ProjectIdentifier: optional.NewString(projectId),
		Identifier:        optional.NewString(agentId),
		DrIdentifier:      optional.NewString(drIdentifier),
		IncludeSecondary:  optional.NewBool(true),
	})
	if err != nil {
		return helpers.HandleApiError(err, d, listHttpResp)
	}
	for _, sibling := range listResp.Content {
		if sibling.DisasterRecoveryNode != nil && sibling.DisasterRecoveryNode.Type_ != nil &&
			*sibling.DisasterRecoveryNode.Type_ == nextgen.PRIMARY_V1DisasterRecoveryType {
			d.Set("current_primary", sibling.DisasterRecoveryNode.Identifier)
			return nil
		}
	}
	d.Set("current_primary", "")
	return nil
}

// resourceGitopsAgentDrPromotionDelete is intentionally a no-op: promoting a DR node to primary
// is not reversible via an API call, so destroying this resource only removes it from state.
func resourceGitopsAgentDrPromotionDelete(_ context.Context, d *schema.ResourceData, _ interface{}) diag.Diagnostics {
	d.SetId("")
	return nil
}
