package central_notification_channel

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

func ResourceCentralNotificationChannel() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for managing Harness Notification Channels.",

		CreateContext: resourceCentralNotificationChannelCreate,
		ReadContext:   resourceCentralNotificationChannelRead,
		UpdateContext: resourceCentralNotificationChannelUpdate,
		DeleteContext: resourceCentralNotificationChannelDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"identifier": {
				Type:     schema.TypeString,
				Required: true,
			},
			"org": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
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
			"account": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Account identifier associated with this notification channel.",
			},
			"created": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Timestamp when the notification channel was created.",
			},
			"last_modified": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Timestamp when the notification channel was last modified.",
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

func resourceCentralNotificationChannelCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	ctx = context.WithValue(ctx, nextgen.ContextAccessToken, hh.EnvVars.BearerToken.Get())
	accountID := c.AccountId
	id := d.Id()
	if id == "" {
		d.MarkNewResource()
	}
	scope := getScope(d)

	req := buildCentralNotificationChannelRequest(d, accountID)
	var resp nextgen.NotificationChannelDto
	var httpResp *http.Response
	var err error
	switch scope.scope {
	case Project:
		resp, httpResp, err = c.NotificationChannelsApi.CreateNotificationChannel(ctx, scope.org, scope.project,
			&nextgen.NotificationChannelsApiCreateNotificationChannelOpts{
				Body:           optional.NewInterface(req),
				HarnessAccount: optional.NewString(accountID),
			})
	case Org:
		resp, httpResp, err = c.NotificationChannelsApi.CreateNotificationChannelOrg(ctx, scope.org,
			&nextgen.NotificationChannelsApiCreateNotificationChannelOrgOpts{
				Body:           optional.NewInterface(req),
				HarnessAccount: optional.NewString(accountID),
			})
	default:

		resp, httpResp, err = c.NotificationChannelsApi.CreateNotificationChannelAccount(ctx,
			&nextgen.NotificationChannelsApiCreateNotificationChannelAccountOpts{
				Body:           optional.NewInterface(req),
				HarnessAccount: optional.NewString(accountID),
			})
	}
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}
	d.SetId(resp.Identifier)
	return readCentralNotificationChannel(accountID, d, resp)
}

func resourceCentralNotificationChannelRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	accountID := c.AccountId
	identifier := d.Get("identifier").(string)
	scope := getScope(d)

	var resp nextgen.NotificationChannelDto
	var httpResp *http.Response
	var err error
	switch scope.scope {
	case Project:
		resp, httpResp, err = c.NotificationChannelsApi.GetNotificationChannel(ctx, identifier, scope.org, scope.project,
			&nextgen.NotificationChannelsApiGetNotificationChannelOpts{
				HarnessAccount: optional.NewString(accountID),
			})
	case Org:
		resp, httpResp, err = c.NotificationChannelsApi.GetNotificationChannelOrg(ctx, identifier, scope.org,
			&nextgen.NotificationChannelsApiGetNotificationChannelOrgOpts{
				HarnessAccount: optional.NewString(accountID),
			})
	default:
		resp, httpResp, err = c.NotificationChannelsApi.GetNotificationChannelAccount(ctx, identifier,
			&nextgen.NotificationChannelsApiGetNotificationChannelAccountOpts{
				HarnessAccount: optional.NewString(accountID),
			})
	}

	if err != nil {
		return helpers.HandleReadApiError(err, d, httpResp)
	}

	readCentralNotificationChannel(accountID, d, resp)

	return nil
}

func resourceCentralNotificationChannelUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	ctx = context.WithValue(ctx, nextgen.ContextAccessToken, hh.EnvVars.BearerToken.Get())
	accountID := c.AccountId
	identifier := d.Get("identifier").(string)
	scope := getScope(d)

	req := buildCentralNotificationChannelRequest(d, accountID)
	var resp nextgen.NotificationChannelDto
	var httpResp *http.Response
	var err error
	switch scope.scope {
	case Project:
		resp, httpResp, err = c.NotificationChannelsApi.UpdateNotificationChannel(ctx, identifier, scope.org, scope.project,
			&nextgen.NotificationChannelsApiUpdateNotificationChannelOpts{
				Body:           optional.NewInterface(req),
				HarnessAccount: optional.NewString(accountID),
			})
	case Org:
		resp, httpResp, err = c.NotificationChannelsApi.UpdateNotificationChannelOrg(ctx, identifier, scope.org,
			&nextgen.NotificationChannelsApiUpdateNotificationChannelOrgOpts{
				Body:           optional.NewInterface(req),
				HarnessAccount: optional.NewString(accountID),
			})
	default:
		resp, httpResp, err = c.NotificationChannelsApi.UpdateNotificationChannelAccount(ctx, identifier,
			&nextgen.NotificationChannelsApiUpdateNotificationChannelAccountOpts{
				Body:           optional.NewInterface(req),
				HarnessAccount: optional.NewString(accountID),
			})
	}
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}
	d.SetId(resp.Identifier)
	return readCentralNotificationChannel(accountID, d, resp)
}

func resourceCentralNotificationChannelDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	ctx = context.WithValue(ctx, nextgen.ContextAccessToken, hh.EnvVars.BearerToken.Get())
	accountID := c.AccountId
	identifier := d.Get("identifier").(string)
	scope := getScope(d)
	var httpResp *http.Response
	var err error
	switch scope.scope {

	case Project:
		httpResp, err = c.NotificationChannelsApi.DeleteNotificationChannel(ctx, identifier, scope.org, scope.project,
			&nextgen.NotificationChannelsApiDeleteNotificationChannelOpts{
				HarnessAccount: optional.NewString(accountID),
			})
	case Org:
		httpResp, err = c.NotificationChannelsApi.DeleteNotificationChannelOrg(ctx, identifier, scope.org,
			&nextgen.NotificationChannelsApiDeleteNotificationChannelOrgOpts{
				HarnessAccount: optional.NewString(accountID),
			})
	default:
		httpResp, err = c.NotificationChannelsApi.DeleteNotificationChannelAccount(ctx, identifier,
			&nextgen.NotificationChannelsApiDeleteNotificationChannelAccountOpts{
				HarnessAccount: optional.NewString(accountID),
			})
	}
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}
	return nil
}

func buildCentralNotificationChannelRequest(d *schema.ResourceData, accountIdentifier string) *nextgen.NotificationChannelDto {
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
		Account:    accountIdentifier,
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

func readCentralNotificationChannel(accountIdentifier string, d *schema.ResourceData, notificationChannelDto nextgen.NotificationChannelDto) diag.Diagnostics {
	d.SetId(notificationChannelDto.Identifier)
	if notificationChannelDto.Org != "" {
		d.Set("org", notificationChannelDto.Org)
	}
	if notificationChannelDto.Project != "" {
		d.Set("project", notificationChannelDto.Project)
	}
	d.Set("identifier", notificationChannelDto.Identifier)
	d.Set("name", notificationChannelDto.Name)
	d.Set("notification_channel_type", notificationChannelDto.NotificationChannelType)
	d.Set("status", notificationChannelDto.Status)
	d.Set("last_modified", notificationChannelDto.LastModified)
	d.Set("created", notificationChannelDto.Created)
	d.Set("account", accountIdentifier)

	channelDTO := notificationChannelDto.Channel
	if channelDTO == nil {
		return nil
	}
	channel := map[string]interface{}{
		"slack_webhook_urls":          channelDTO.SlackWebhookUrls,
		"webhook_urls":                channelDTO.WebhookUrls,
		"email_ids":                   channelDTO.EmailIds,
		"pager_duty_integration_keys": channelDTO.PagerDutyIntegrationKeys,
		"ms_team_keys":                channelDTO.MsTeamKeys,
		"datadog_urls":                channelDTO.DatadogUrls,
		"user_groups":                 flattenUserGroups(channelDTO.UserGroups),
		"headers":                     flattenHeaders(channelDTO.Headers),
		"delegate_selectors":          channelDTO.DelegateSelectors,
		"execute_on_delegate":         channelDTO.ExecuteOnDelegate,
	}
	if val := channelDTO.ApiKey; val != "" {
		channel["api_key"] = val
	}
	d.Set("channel", []interface{}{channel})

	return nil
}

func flattenUserGroups(input []nextgen.UserGroupDto) []interface{} {
	var result []interface{}
	for _, ug := range input {
		result = append(result, map[string]interface{}{"identifier": ug.Identifier})
	}
	return result
}

func flattenHeaders(input []nextgen.WebHookHeaders) []interface{} {
	var result []interface{}
	for _, h := range input {
		result = append(result, map[string]interface{}{
			"key":   h.Key,
			"value": h.Value,
		})
	}
	return result
}

type Scope struct {
	org     string
	project string
	scope   ScopeLevel
}

type ScopeLevel string

const (
	Account ScopeLevel = "account"
	Org     ScopeLevel = "org"
	Project ScopeLevel = "project"
)

func getScope(d *schema.ResourceData) *Scope {
	org := ""
	project := ""

	if attr, ok := d.GetOk("org"); ok {
		org = (attr.(string))
	}

	if attr, ok := d.GetOk("project"); ok {
		project = (attr.(string))
	}

	var scope ScopeLevel
	if org == "" {
		scope = Account
	} else if project == "" {
		scope = Org
	} else {
		scope = Project
	}

	return &Scope{
		org:     org,
		project: project,
		scope:   scope,
	}
}
