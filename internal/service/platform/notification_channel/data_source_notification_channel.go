package notification_channel

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

func DataSourceNotificationRuleService() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving a Notification Rule.",

		ReadContext: dataNotificationChannelRead,

		Schema: map[string]*schema.Schema{
			"org_id": {
				Description: "Identifier of the organization in which the Notification Rule is configured.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"project_id": {
				Description: "Identifier of the project in which the Notification Rule is configured.",
				Type:        schema.TypeString,
				Required:    true,
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

func dataNotificationChannelRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	var accountIdentifier, orgIdentifier, projectIdentifier string
	accountIdentifier = c.AccountId
	if attr, ok := d.GetOk("org_id"); ok {
		orgIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("project_id"); ok {
		projectIdentifier = attr.(string)
	}
	id := d.Id()
	if id == "" {
		d.MarkNewResource()
		return nil
	}

	var resp nextgen.NotificationChannelDto
	var httpResp *http.Response
	var err error
	if orgIdentifier != "" && projectIdentifier != "" {
		resp, httpResp, err = c.NotificationChannelsApi.GetNotificationChannel(ctx, id, orgIdentifier, projectIdentifier,
			&nextgen.NotificationChannelsApiGetNotificationChannelOpts{
				HarnessAccount: optional.NewString(accountIdentifier),
			})
	} else if orgIdentifier != "" {
		resp, httpResp, err = c.NotificationChannelsApi.GetNotificationChannelOrg(ctx, id, orgIdentifier,
			&nextgen.NotificationChannelsApiGetNotificationChannelOrgOpts{
				HarnessAccount: optional.NewString(accountIdentifier),
			})
	} else {
		resp, httpResp, err = c.NotificationChannelsApi.GetNotificationChannelAccount(ctx, id,
			&nextgen.NotificationChannelsApiGetNotificationChannelAccountOpts{
				HarnessAccount: optional.NewString(accountIdentifier),
			})
	}

	if err != nil {
		return helpers.HandleReadApiError(err, d, httpResp)
	}

	readNotificationChannel(d, resp)

	return nil
}
