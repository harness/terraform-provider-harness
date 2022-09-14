package usergroup

import (
	"context"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceUserGroup() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for creating a Harness User Group.",

		ReadContext:   resourceUserGroupRead,
		UpdateContext: resourceUserGroupCreateOrUpdate,
		DeleteContext: resourceUserGroupDelete,
		CreateContext: resourceUserGroupCreateOrUpdate,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			// "is_sso_linked": {
			// 	Description: "Whether the user group is linked to an SSO account.",
			// 	Type:        schema.TypeBool,
			// 	Optional:    true,
			// 	Computed:    true,
			// },
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
				Description: "List of users in the UserGroup.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"notification_configs": {
				Description: "List of notification settings.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Description: "Can be one of EMAIL, SLACK, PAGERDUTY, MSTEAMS",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"slack_webhook_url": {
							Description: "Url of slack webhook",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"microsoft_teams_webhook_url": {
							Description: "Url of Microsoft teams webhook",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"pager_duty_key": {
							Description: "Pager duty key",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"group_email": {
							Description: "Group email",
							Type:        schema.TypeString,
							Optional:    true,
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
				Description: "Type of linked SSO",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"sso_linked": {
				Description: "Whether sso is linked or not",
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
	resp, _, err := c.UserGroupApi.GetUserGroup(ctx, c.AccountId, id, &nextgen.UserGroupApiGetUserGroupOpts{
		OrgIdentifier:     helpers.BuildField(d, "org_id"),
		ProjectIdentifier: helpers.BuildField(d, "project_id"),
	})

	if err != nil {
		return helpers.HandleApiError(err, d)
	}

	// Soft delete lookup error handling
	// https://harness.atlassian.net/browse/PL-23765
	if resp.Data == nil {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	readUserGroup(d, resp.Data)

	return nil
}

func resourceUserGroupCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	var err error
	var resp nextgen.ResponseDtoUserGroup

	id := d.Id()
	ug := buildUserGroup(d)
	ug.AccountIdentifier = c.AccountId

	if id == "" {
		resp, _, err = c.UserGroupApi.PostUserGroup(ctx, ug, c.AccountId, &nextgen.UserGroupApiPostUserGroupOpts{
			OrgIdentifier:     helpers.BuildField(d, "org_id"),
			ProjectIdentifier: helpers.BuildField(d, "project_id"),
		})
	} else {
		resp, _, err = c.UserGroupApi.PutUserGroup(ctx, ug, c.AccountId, &nextgen.UserGroupApiPutUserGroupOpts{
			OrgIdentifier:     helpers.BuildField(d, "org_id"),
			ProjectIdentifier: helpers.BuildField(d, "project_id"),
		})
	}

	if err != nil {
		return helpers.HandleApiError(err, d)
	}

	readUserGroup(d, resp.Data)

	return nil
}

func resourceUserGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	_, _, err := c.UserGroupApi.DeleteUserGroup(ctx, c.AccountId, d.Id(), &nextgen.UserGroupApiDeleteUserGroupOpts{
		OrgIdentifier:     optional.NewString(d.Get("org_id").(string)),
		ProjectIdentifier: optional.NewString(d.Get("project_id").(string)),
	})

	if err != nil {
		return diag.Errorf(err.(nextgen.GenericSwaggerError).Error())
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
		userGroup.Users = helpers.ExpandField(attr.(*schema.Set).List())
	}

	if attr, ok := d.GetOk("notification_configs"); ok {
		userGroup.NotificationConfigs = expandNotificationConfig(attr.(*schema.Set).List())
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

func readUserGroup(d *schema.ResourceData, env *nextgen.UserGroup) {
	d.SetId(env.Identifier)
	d.Set("identifier", env.Identifier)
	d.Set("org_id", env.OrgIdentifier)
	d.Set("project_id", env.ProjectIdentifier)
	d.Set("name", env.Name)
	d.Set("description", env.Description)
	d.Set("tags", helpers.FlattenTags(env.Tags))
	d.Set("users", env.Users)
	d.Set("notification_configs", flattenNotificationConfig(env.NotificationConfigs))
	d.Set("linked_sso_id", env.LinkedSsoId)
	d.Set("linked_sso_display_name", env.LinkedSsoDisplayName)
	d.Set("sso_group_id", env.SsoGroupId)
	d.Set("sso_group_name", env.SsoGroupName)
	d.Set("linked_sso_type", env.LinkedSsoType)
	d.Set("externally_managed", env.ExternallyManaged)
	d.Set("sso_linked", env.SsoLinked)
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
				"type":        notificationConfig.Type_,
				"group_email": notificationConfig.GroupEmail,
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
