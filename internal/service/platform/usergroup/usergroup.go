package usergroup

import (
	"context"
	"fmt"
	"net/http"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceUserGroup() *schema.Resource {
	resource := &schema.Resource{
		Description: fmt.Sprintf(`
		Resource for creating a Harness User Group. Linking SSO providers with User Groups:

		The following fields need to be populated for LDAP SSO Providers:
		
		- linked_sso_id
		
		- linked_sso_display_name
		
		- sso_group_id
		
		- sso_group_name
		
		- linked_sso_type
		
		- sso_linked
		
		The following fields need to be populated for SAML SSO Providers:
		
		- linked_sso_id
		
		- linked_sso_display_name
		
		- sso_group_name
		
		- sso_group_id // same as sso_group_name
		
		- linked_sso_type
		
		- sso_linked`),

		ReadContext:   resourceUserGroupRead,
		UpdateContext: resourceUserGroupCreateOrUpdate,
		DeleteContext: resourceUserGroupDelete,
		CreateContext: resourceUserGroupCreateOrUpdate,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"linked_sso_id": {
				Description: "The SSO account ID that the user group is linked to.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"externally_managed": {
				Description: "Whether the user group is externally managed.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"users": {
				Description:   "List of users in the UserGroup. Either provide list of users or list of user emails.",
				Type:          schema.TypeList,
				Optional:      true,
				ConflictsWith: []string{"user_emails"},
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"user_emails": {
				Description:   "List of user emails in the UserGroup. Either provide list of users or list of user emails.",
				Type:          schema.TypeList,
				Optional:      true,
				ConflictsWith: []string{"users"},
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"notification_configs": {
				Description: "List of notification settings.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Description: "Can be one of EMAIL, SLACK, PAGERDUTY, MSTEAMS.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"slack_webhook_url": {
							Description: "Url of slack webhook.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"microsoft_teams_webhook_url": {
							Description: "Url of Microsoft teams webhook.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"pager_duty_key": {
							Description: "Pager duty key.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"group_email": {
							Description: "Group email.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"send_email_to_all_users": {
							Description: "Send email to all the group members.",
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"linked_sso_display_name": {
				Description: "Name of the linked SSO.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"sso_group_id": {
				Description: "Identifier of the userGroup in SSO.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"sso_group_name": {
				Description: "Name of the SSO userGroup.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"linked_sso_type": {
				Description: "Type of linked SSO.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"sso_linked": {
				Description: "Whether sso is linked or not.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
		},
	}

	helpers.SetMultiLevelResourceSchema(resource.Schema)

	return resource
}

func resourceUserGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Id()
	resp, httpResp, err := c.UserGroupApi.GetUserGroup(ctx, c.AccountId, id, &nextgen.UserGroupApiGetUserGroupOpts{
		OrgIdentifier:     helpers.BuildField(d, "org_id"),
		ProjectIdentifier: helpers.BuildField(d, "project_id"),
	})

	if err != nil {
		return helpers.HandleReadApiError(err, d, httpResp)
	}

	// Soft delete lookup error handling
	// https://harness.atlassian.net/browse/PL-23765
	if resp.Data == nil {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	if attr, ok := d.GetOk("user_emails"); ok {
		d.Set("user_emails", attr)
	}
	if _, ok := d.GetOk("users"); ok {
		d.Set("users", resp.Data.Users)
	}
	readUserGroup(d, resp.Data)

	return nil
}

func resourceUserGroupCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	if _, ok := d.GetOk("user_emails"); !ok {
		var err error
		var resp nextgen.ResponseDtoUserGroup
		var httpResp *http.Response

		id := d.Id()
		ug := buildUserGroup(d)
		ug.AccountIdentifier = c.AccountId

		if id == "" {
			resp, httpResp, err = c.UserGroupApi.PostUserGroup(ctx, ug, c.AccountId, &nextgen.UserGroupApiPostUserGroupOpts{
				OrgIdentifier:     helpers.BuildField(d, "org_id"),
				ProjectIdentifier: helpers.BuildField(d, "project_id"),
			})
		} else {
			resp, httpResp, err = c.UserGroupApi.PutUserGroup(ctx, ug, c.AccountId, &nextgen.UserGroupApiPutUserGroupOpts{
				OrgIdentifier:     helpers.BuildField(d, "org_id"),
				ProjectIdentifier: helpers.BuildField(d, "project_id"),
			})
		}

		if err != nil {
			return helpers.HandleApiError(err, d, httpResp)
		}

		readUserGroup(d, resp.Data)

		return nil
	}

	var err error
	var resp nextgen.ResponseDtoUserGroupResponseV2
	var httpResp *http.Response

	id := d.Id()
	ug := buildUserGroupV2(d)
	ug.AccountIdentifier = c.AccountId

	if id == "" {
		resp, httpResp, err = c.UserGroupApi.PostUserGroupV2(ctx, ug, c.AccountId, &nextgen.UserGroupApiPostUserGroupV2Opts{
			OrgIdentifier:     helpers.BuildField(d, "org_id"),
			ProjectIdentifier: helpers.BuildField(d, "project_id"),
		})
	} else {
		resp, httpResp, err = c.UserGroupApi.PutUserGroupV2(ctx, ug, c.AccountId, &nextgen.UserGroupApiPutUserGroupV2Opts{
			OrgIdentifier:     helpers.BuildField(d, "org_id"),
			ProjectIdentifier: helpers.BuildField(d, "project_id"),
		})
	}

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	readUserGroupV2(d, resp.Data)

	return nil
}

func resourceUserGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	_, httpResp, err := c.UserGroupApi.DeleteUserGroup(ctx, c.AccountId, d.Id(), &nextgen.UserGroupApiDeleteUserGroupOpts{
		OrgIdentifier:     helpers.BuildField(d, "org_id"),
		ProjectIdentifier: helpers.BuildField(d, "project_id"),
	})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	return nil
}

func buildUserGroup(d *schema.ResourceData) nextgen.UserGroup {
	userGroup := &nextgen.UserGroup{}

	if attr, ok := d.GetOk("org_id"); ok {
		userGroup.OrgIdentifier = attr.(string)
	}

	if attr, ok := d.GetOk("project_id"); ok {
		userGroup.ProjectIdentifier = attr.(string)
	}

	if attr, ok := d.GetOk("description"); ok {
		userGroup.Description = attr.(string)
	}

	if attr, ok := d.GetOk("name"); ok {
		userGroup.Name = attr.(string)
	}

	if attr, ok := d.GetOk("identifier"); ok {
		userGroup.Identifier = attr.(string)
	}

	if attr := d.Get("tags").(*schema.Set).List(); len(attr) > 0 {
		userGroup.Tags = helpers.ExpandTags(attr)
	}

	if attr, ok := d.GetOk("users"); ok {
		userGroup.Users = helpers.ExpandField(attr.([]interface{}))
	}

	if attr, ok := d.GetOk("notification_configs"); ok {
		userGroup.NotificationConfigs = expandNotificationConfig(attr.([]interface{}))
	}

	if attr, ok := d.GetOk("is_sso_linked"); ok {
		userGroup.IsSsoLinked = attr.(bool)
	}

	if attr, ok := d.GetOk("linked_sso_id"); ok {
		userGroup.LinkedSsoId = attr.(string)
	}

	if attr, ok := d.GetOk("linked_sso_display_name"); ok {
		userGroup.LinkedSsoDisplayName = attr.(string)
	}

	if attr, ok := d.GetOk("sso_group_id"); ok {
		userGroup.SsoGroupId = attr.(string)
	}

	if attr, ok := d.GetOk("sso_group_name"); ok {
		userGroup.SsoGroupName = attr.(string)
	}

	if attr, ok := d.GetOk("linked_sso_type"); ok {
		userGroup.LinkedSsoType = attr.(string)
	}

	if attr, ok := d.GetOk("externally_managed"); ok {
		userGroup.ExternallyManaged = attr.(bool)
	}

	if attr, ok := d.GetOk("sso_linked"); ok {
		userGroup.SsoLinked = attr.(bool)
	}

	return *userGroup
}

func buildUserGroupV2(d *schema.ResourceData) nextgen.UserGroupRequestV2 {
	userGroup := &nextgen.UserGroupRequestV2{}

	if attr, ok := d.GetOk("org_id"); ok {
		userGroup.OrgIdentifier = attr.(string)
	}

	if attr, ok := d.GetOk("project_id"); ok {
		userGroup.ProjectIdentifier = attr.(string)
	}

	if attr, ok := d.GetOk("description"); ok {
		userGroup.Description = attr.(string)
	}

	if attr, ok := d.GetOk("name"); ok {
		userGroup.Name = attr.(string)
	}

	if attr, ok := d.GetOk("identifier"); ok {
		userGroup.Identifier = attr.(string)
	}

	if attr := d.Get("tags").(*schema.Set).List(); len(attr) > 0 {
		userGroup.Tags = helpers.ExpandTags(attr)
	}

	if attr, ok := d.GetOk("user_emails"); ok {
		userGroup.Users = helpers.ExpandField(attr.([]interface{}))
	}

	if attr, ok := d.GetOk("notification_configs"); ok {
		userGroup.NotificationConfigs = expandNotificationConfig(attr.([]interface{}))
	}

	if attr, ok := d.GetOk("is_sso_linked"); ok {
		userGroup.IsSsoLinked = attr.(bool)
	}

	if attr, ok := d.GetOk("linked_sso_id"); ok {
		userGroup.LinkedSsoId = attr.(string)
	}

	if attr, ok := d.GetOk("linked_sso_display_name"); ok {
		userGroup.LinkedSsoDisplayName = attr.(string)
	}

	if attr, ok := d.GetOk("sso_group_id"); ok {
		userGroup.SsoGroupId = attr.(string)
	}

	if attr, ok := d.GetOk("sso_group_name"); ok {
		userGroup.SsoGroupName = attr.(string)
	}

	if attr, ok := d.GetOk("linked_sso_type"); ok {
		userGroup.LinkedSsoType = attr.(string)
	}

	if attr, ok := d.GetOk("externally_managed"); ok {
		userGroup.ExternallyManaged = attr.(bool)
	}

	if attr, ok := d.GetOk("sso_linked"); ok {
		userGroup.SsoLinked = attr.(bool)
	}

	return *userGroup
}

func readUserGroupV2(d *schema.ResourceData, env *nextgen.UserGroupResponseV2) {
	d.SetId(env.Identifier)
	d.Set("identifier", env.Identifier)
	d.Set("org_id", env.OrgIdentifier)
	d.Set("project_id", env.ProjectIdentifier)
	d.Set("name", env.Name)
	d.Set("description", env.Description)
	d.Set("tags", helpers.FlattenTags(env.Tags))
	d.Set("user_emails", flattenUserInfo(env.Users))
	d.Set("notification_configs", flattenNotificationConfig(env.NotificationConfigs))
	d.Set("linked_sso_id", env.LinkedSsoId)
	d.Set("linked_sso_display_name", env.LinkedSsoDisplayName)
	d.Set("sso_group_id", env.SsoGroupId)
	d.Set("sso_group_name", env.SsoGroupName)
	d.Set("linked_sso_type", env.LinkedSsoType)
	d.Set("externally_managed", env.ExternallyManaged)
	d.Set("sso_linked", env.SsoLinked)
}

func readUserGroup(d *schema.ResourceData, env *nextgen.UserGroup) {
	d.SetId(env.Identifier)
	d.Set("identifier", env.Identifier)
	d.Set("org_id", env.OrgIdentifier)
	d.Set("project_id", env.ProjectIdentifier)
	d.Set("name", env.Name)
	d.Set("description", env.Description)
	d.Set("tags", helpers.FlattenTags(env.Tags))
	d.Set("notification_configs", flattenNotificationConfig(env.NotificationConfigs))
	d.Set("linked_sso_id", env.LinkedSsoId)
	d.Set("linked_sso_display_name", env.LinkedSsoDisplayName)
	d.Set("sso_group_id", env.SsoGroupId)
	d.Set("sso_group_name", env.SsoGroupName)
	d.Set("linked_sso_type", env.LinkedSsoType)
	d.Set("externally_managed", env.ExternallyManaged)
	d.Set("sso_linked", env.SsoLinked)
}

func fetchUserIds(userInfos []nextgen.UserInfo) []string {
	var result []string

	for _, userInfo := range userInfos {
		result = append(result, userInfo.Uuid)
	}

	return result
}

func expandNotificationConfig(notificationConfigs []interface{}) []nextgen.NotificationSettingConfigDto {
	var result []nextgen.NotificationSettingConfigDto
	for _, notificationConfig := range notificationConfigs {
		v := notificationConfig.(map[string]interface{})

		var resultNotificationConfig nextgen.NotificationSettingConfigDto
		resultNotificationConfig.Type_ = v["type"].(string)
		if resultNotificationConfig.Type_ == "SLACK" {
			resultNotificationConfig.SlackWebhookUrl = v["slack_webhook_url"].(string)
		}
		if resultNotificationConfig.Type_ == "EMAIL" {
			resultNotificationConfig.GroupEmail = v["group_email"].(string)
			resultNotificationConfig.SendEmailToAllUsers = v["send_email_to_all_users"].(bool)
		}
		if resultNotificationConfig.Type_ == "MSTEAMS" {
			resultNotificationConfig.MicrosoftTeamsWebhookUrl = v["microsoft_teams_webhook_url"].(string)
		}
		if resultNotificationConfig.Type_ == "PAGERDUTY" {
			resultNotificationConfig.PagerDutyKey = v["pager_duty_key"].(string)
		}
		result = append(result, resultNotificationConfig)
	}
	return result
}

func flattenUserInfo(userInfos []nextgen.UserInfo) []string {
	var result []string
	for _, userInfo := range userInfos {
		result = append(result, userInfo.Email)
	}
	return result
}

func flattenNotificationConfig(notificationConfigs []nextgen.NotificationSettingConfigDto) []interface{} {
	var result []interface{}
	for _, notificationConfig := range notificationConfigs {
		if notificationConfig.Type_ == "SLACK" {
			result = append(result, map[string]interface{}{
				"type":              notificationConfig.Type_,
				"slack_webhook_url": notificationConfig.SlackWebhookUrl,
			})
		}
		if notificationConfig.Type_ == "EMAIL" {
			result = append(result, map[string]interface{}{
				"type":                    notificationConfig.Type_,
				"group_email":             notificationConfig.GroupEmail,
				"send_email_to_all_users": notificationConfig.SendEmailToAllUsers,
			})
		}
		if notificationConfig.Type_ == "MSTEAMS" {
			result = append(result, map[string]interface{}{
				"type":                        notificationConfig.Type_,
				"microsoft_teams_webhook_url": notificationConfig.MicrosoftTeamsWebhookUrl,
			})
		}
		if notificationConfig.Type_ == "PAGERDUTY" {
			result = append(result, map[string]interface{}{
				"type":           notificationConfig.Type_,
				"pager_duty_key": notificationConfig.PagerDutyKey,
			})
		}
	}
	return result
}
