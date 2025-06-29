package notification_channel

import (
	"context"
	"net/http"

	"github.com/antihax/optional"
	hh "github.com/harness/harness-go-sdk/harness/helpers"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceNotificationChannel() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for managing Harness Notification Channels.",

		CreateContext: resourceNotificationChannelCreate,
		ReadContext:   resourceNotificationChannelRead,
		UpdateContext: resourceNotificationChannelUpdate,
		DeleteContext: resourceNotificationChannelDelete,

		Schema: map[string]*schema.Schema{
			"identifier": {
				Type:     schema.TypeString,
				Required: true,
			},
			"org": {
				Type:     schema.TypeString,
				Required: true,
			},
			"project": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"notification_channel_type": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"EMAIL", "SLACK", "PAGERDUTY", "MSTEAMS", "WEBHOOK", "DATADOG",
				}, false),
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "ENABLED",
			},
			"channel": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"email_ids":                   stringListSchema(),
						"slack_webhook_urls":          stringListSchema(),
						"webhook_urls":                stringListSchema(),
						"pager_duty_integration_keys": stringListSchema(),
						"ms_team_keys":                stringListSchema(),
						"datadog_urls":                stringListSchema(),
						"delegate_selectors":          stringListSchema(),
						"api_key": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"execute_on_delegate": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"user_groups": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{
								"identifier": {Type: schema.TypeString, Required: true},
							}},
						},
						"headers": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{
								"key":   {Type: schema.TypeString, Required: true},
								"value": {Type: schema.TypeString, Required: true},
							}},
						},
					},
				},
			},
		},
	}
}

func resourceNotificationChannelCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	ctx = context.WithValue(ctx, nextgen.ContextAccessToken, hh.EnvVars.BearerToken.Get())
	var accountIdentifier, orgIdentifier, projectIdentifier string
	accountIdentifier = c.AccountId
	if attr, ok := d.GetOk("org_id"); ok {
		orgIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("project_id"); ok {
		projectIdentifier = attr.(string)
	}

	req := buildNotificationChannelRequest(d)
	var resp nextgen.NotificationChannelDto
	var httpResp *http.Response
	var err error
	if orgIdentifier != "" && projectIdentifier != "" {
		resp, httpResp, err = c.NotificationChannelsApi.CreateNotificationChannel(ctx, orgIdentifier, projectIdentifier,
			&nextgen.NotificationChannelsApiCreateNotificationChannelOpts{
				Body:           optional.NewInterface(req),
				HarnessAccount: optional.NewString(accountIdentifier),
			})
	} else if orgIdentifier != "" {
		resp, httpResp, err = c.NotificationChannelsApi.CreateNotificationChannelOrg(ctx, orgIdentifier,
			&nextgen.NotificationChannelsApiCreateNotificationChannelOrgOpts{
				Body:           optional.NewInterface(req),
				HarnessAccount: optional.NewString(accountIdentifier),
			})
	} else {
		resp, httpResp, err = c.NotificationChannelsApi.CreateNotificationChannelAccount(ctx,
			&nextgen.NotificationChannelsApiCreateNotificationChannelAccountOpts{
				Body:           optional.NewInterface(req),
				HarnessAccount: optional.NewString(accountIdentifier),
			})
	}
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}
	d.SetId(resp.Identifier)
	return readNotificationChannel(d, resp)
}

func resourceNotificationChannelRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

func resourceNotificationChannelUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return resourceNotificationChannelCreate(ctx, d, meta) // assuming PUT-like behavior
}

func resourceNotificationChannelDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	ctx = context.WithValue(ctx, nextgen.ContextAccessToken, hh.EnvVars.BearerToken.Get())
	accountIdentifier := c.AccountId
	identifier := d.Get("identifier").(string)
	orgIdentifier := d.Get("org_id").(string)
	projectIdentifier := d.Get("project_id").(string)
	var httpResp *http.Response
	var err error
	if orgIdentifier != "" && projectIdentifier != "" {
		httpResp, err = c.NotificationChannelsApi.DeleteNotificationChannel(ctx, identifier, orgIdentifier, projectIdentifier,
			&nextgen.NotificationChannelsApiDeleteNotificationChannelOpts{
				HarnessAccount: optional.NewString(accountIdentifier),
			})
	} else if orgIdentifier != "" {
		httpResp, err = c.NotificationChannelsApi.DeleteNotificationChannelOrg(ctx, identifier, orgIdentifier,
			&nextgen.NotificationChannelsApiDeleteNotificationChannelOrgOpts{
				HarnessAccount: optional.NewString(accountIdentifier),
			})
	} else {
		httpResp, err = c.NotificationChannelsApi.DeleteNotificationChannelAccount(ctx, identifier,
			&nextgen.NotificationChannelsApiDeleteNotificationChannelAccountOpts{
				HarnessAccount: optional.NewString(accountIdentifier),
			})
	}
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}
	return nil
}

func buildNotificationChannelRequest(d *schema.ResourceData) *nextgen.NotificationChannelDto {
	channelData := d.Get("channel").([]interface{})[0].(map[string]interface{})

	channelDTO := nextgen.ChannelDto{
		EmailIds:                 expandStringList(channelData["email_ids"]),
		SlackWebhookUrls:         expandStringList(channelData["slack_webhook_urls"]),
		WebhookUrls:              expandStringList(channelData["webhook_urls"]),
		PagerDutyIntegrationKeys: expandStringList(channelData["pager_duty_integration_keys"]),
		MsTeamKeys:               expandStringList(channelData["ms_team_keys"]),
		DatadogUrls:              expandStringList(channelData["datadog_urls"]),
		ApiKey:                   channelData["api_key"].(string),
		DelegateSelectors:        expandStringList(channelData["delegate_selectors"]),
		ExecuteOnDelegate:        channelData["execute_on_delegate"].(bool),
	}

	if v, ok := channelData["user_groups"]; ok {
		channelDTO.UserGroups = expandUserGroups(v.([]interface{}))
	}
	if v, ok := channelData["headers"]; ok {
		channelDTO.Headers = expandHeaders(v.([]interface{}))
	}

	return &nextgen.NotificationChannelDto{
		Identifier: d.Get("identifier").(string),
		Name:       d.Get("name").(string),
		Org:        d.Get("org").(string),
		Project:    d.Get("project").(string),
		NotificationChannelType: func() *nextgen.ChannelType {
			s := nextgen.ChannelType(d.Get("notification_channel_type").(string))
			return &s
		}(),
		Status: func() *nextgen.Status {
			s := nextgen.Status(d.Get("status").(string))
			return &s
		}(),
		Channel: &channelDTO,
	}
}

func stringListSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Elem:     &schema.Schema{Type: schema.TypeString},
	}
}

func expandStringList(raw interface{}) []string {
	if raw == nil {
		return nil
	}
	rawList := raw.([]interface{})
	strList := make([]string, len(rawList))
	for i, val := range rawList {
		strList[i] = val.(string)
	}
	return strList
}

func expandUserGroups(raw []interface{}) []nextgen.UserGroupDto {
	result := make([]nextgen.UserGroupDto, len(raw))
	for i, item := range raw {
		data := item.(map[string]interface{})
		result[i] = nextgen.UserGroupDto{
			Identifier: data["identifier"].(string),
		}
	}
	return result
}

func expandHeaders(raw []interface{}) []nextgen.WebHookHeaders {
	result := make([]nextgen.WebHookHeaders, len(raw))
	for i, item := range raw {
		data := item.(map[string]interface{})
		result[i] = nextgen.WebHookHeaders{
			Key:   data["key"].(string),
			Value: data["value"].(string),
		}
	}
	return result
}

func readNotificationChannel(d *schema.ResourceData, notificationChannelDto nextgen.NotificationChannelDto) diag.Diagnostics {
	// Implement read logic as needed
	d.SetId(notificationChannelDto.Identifier)
	d.Set("identifier", notificationChannelDto.Identifier)
	d.Set("name", notificationChannelDto.Name)
	return nil
}
