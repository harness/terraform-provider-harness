package central_notification_rule

import (
	"context"
	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"net/http"
)

func DataSourceCentralNotificationRuleService() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving a Notification Rule.",

		ReadContext: dataCentralNotificationRuleRead,

		Schema: map[string]*schema.Schema{
			"org_id": {
				Description: "Identifier of the organization in which the Notification Rule is configured.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"project_id": {
				Description: "Identifier of the project in which the Notification Rule is configured.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"identifier": {
				Description: "Identifier of the Notification Rule.",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}

	return resource
}

func dataCentralNotificationRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	accountID := c.AccountId
	orgID := d.Get("org_id").(string)
	projectID := d.Get("project_id").(string)
	identifier := d.Id()
	var resp nextgen.NotificationRuleDto
	var httpResp *http.Response
	var err error
	if orgID != "" && projectID != "" {
		resp, httpResp, err = c.NotificationRulesApi.GetNotificationRule(ctx, orgID, projectID,
			identifier,
			&nextgen.NotificationRulesApiGetNotificationRuleOpts{
				HarnessAccount: optional.NewString(accountID),
			})
	} else if orgID != "" {
		resp, httpResp, err = c.NotificationRulesApi.GetNotificationRuleOrg(ctx, orgID, identifier,
			&nextgen.NotificationRulesApiGetNotificationRuleOrgOpts{
				HarnessAccount: optional.NewString(accountID),
			})
	} else {
		resp, httpResp, err = c.NotificationRulesApi.GetNotificationRuleAccount(ctx, identifier,
			&nextgen.NotificationRulesApiGetNotificationRuleAccountOpts{
				HarnessAccount: optional.NewString(accountID),
			})
	}

	if err != nil {
		return helpers.HandleReadApiError(err, d, httpResp)
	}

	readCentralNotificationRule(d, resp)
	return nil
}
