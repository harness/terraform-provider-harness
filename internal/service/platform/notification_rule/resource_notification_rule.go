package notification_rule

import (
	"context"

	"github.com/antihax/optional"
	hh "github.com/harness/harness-go-sdk/harness/helpers"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceNotificationRuleService() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for creating a Notification Rule.",

		CreateContext: resourceNotificationRuleCreate,
		ReadContext:   resourceNotificationRuleRead,
		UpdateContext: resourceNotificationRuleUpdate,
		DeleteContext: resourceNotificationRuleDelete,
		Importer:      helpers.MultiLevelResourceImporter,

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
			"request": {
				Description: "Request for creating or updating Notification Rule.",
				Type:        schema.TypeList,
				MinItems:    1,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Description: "Name for the Notification Rule.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"conditions": {
							Description: "Notification Rule conditions specification.",
							Type:        schema.TypeList,
							Required:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Description: "Type of the condition.",
										Type:        schema.TypeString,
										Required:    true,
									},
									"spec": {
										Description: "Specification of the notification condition. Depends on the type of the notification condition.",
										Type:        schema.TypeString,
										Optional:    true,
									},
								},
							},
						},
						"notification_method": {
							Description: "Notification Method specifications.",
							Type:        schema.TypeSet,
							MinItems:    1,
							MaxItems:    1,
							Required:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Description: "Type of the Notification Method.",
										Type:        schema.TypeString,
										Required:    true,
									},
									"spec": {
										Description: "Specification of the notification method. Depends on the type of the notification method.",
										Type:        schema.TypeString,
										Optional:    true,
									},
								},
							},
						},
						"type": {
							Description: "Type of the Notification Rule.",
							Type:        schema.TypeString,
							Required:    true,
						},
					},
				},
			},
		},
	}

	return resource
}

func resourceNotificationRuleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	ctx = context.WithValue(ctx, nextgen.ContextAccessToken, hh.EnvVars.BearerToken.Get())
	var accountIdentifier, identifier string
	accountIdentifier = c.AccountId

	if attr, ok := d.GetOk("identifier"); ok {
		identifier = attr.(string)
	}
	createNotificationRuleRequest, err := buildNotificationRuleRequest(d, identifier)
	if err != nil {
		return diag.Errorf(err.Error())
	}
	respCreate, httpRespCreate, errCreate := c.SrmNotificationApiService.SaveSrmNotification(ctx, accountIdentifier,
		&nextgen.SrmNotificationServiceSaveSrmNotificationOpts{
			Body: optional.NewInterface(createNotificationRuleRequest),
		})
	if errCreate != nil {
		return helpers.HandleApiError(errCreate, d, httpRespCreate)
	}

	readNotificationRule(d, &respCreate.Resource)
	return nil
}

func resourceNotificationRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	ctx = context.WithValue(ctx, nextgen.ContextAccessToken, hh.EnvVars.BearerToken.Get())
	var accountIdentifier, orgIdentifier, projectIdentifier string
	identifier := d.Get("identifier").(string)
	accountIdentifier = c.AccountId
	if attr, ok := d.GetOk("org_id"); ok {
		orgIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("project_id"); ok {
		projectIdentifier = attr.(string)
	}
	resp, httpResp, err := c.SrmNotificationApiService.GetSrmNotification(ctx, identifier, accountIdentifier, orgIdentifier, projectIdentifier)
	if err != nil {
		if err.Error() == "404 Not Found" || err.Error() == "400 Bad Request" {
			d.SetId("")
			d.MarkNewResource()
			return nil
		}
		return helpers.HandleReadApiError(err, d, httpResp)
	}

	// Soft delete lookup error handling
	// https://harness.atlassian.net/browse/PL-23765
	if &resp == nil {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	readNotificationRule(d, &resp.Resource)
	return nil
}

func resourceNotificationRuleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	ctx = context.WithValue(ctx, nextgen.ContextAccessToken, hh.EnvVars.BearerToken.Get())
	var accountIdentifier, orgIdentifier, projectIdentifier, identifier string
	accountIdentifier = c.AccountId
	if attr, ok := d.GetOk("org_id"); ok {
		orgIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("project_id"); ok {
		projectIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("identifier"); ok {
		identifier = attr.(string)
	}
	updateNotifRequest, err := buildNotificationRuleRequest(d, identifier)
	if err != nil {
		return diag.Errorf(err.Error())
	}
	respCreate, httpRespCreate, errCreate := c.SrmNotificationApiService.UpdateSrmNotification(ctx, accountIdentifier, orgIdentifier, projectIdentifier, identifier,
		&nextgen.SrmNotificationApiUpdateSrmNotificationOpts{
			Body: optional.NewInterface(updateNotifRequest),
		})

	if errCreate != nil {
		return helpers.HandleApiError(errCreate, d, httpRespCreate)
	}

	readNotificationRule(d, &respCreate.Resource)
	return nil
}

func resourceNotificationRuleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	ctx = context.WithValue(ctx, nextgen.ContextAccessToken, hh.EnvVars.BearerToken.Get())
	var accountIdentifier, orgIdentifier, projectIdentifier string
	identifier := d.Get("identifier").(string)
	accountIdentifier = c.AccountId
	if attr, ok := d.GetOk("org_id"); ok {
		orgIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("project_id"); ok {
		projectIdentifier = attr.(string)
	}
	_, httpResp, err := c.SrmNotificationApiService.DeleteSrmNotification(ctx, accountIdentifier, orgIdentifier, projectIdentifier, identifier)
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}
	return nil
}

func buildNotificationRuleRequest(d *schema.ResourceData, identifier string) (*nextgen.NotificationRule, error) {
	notificationRuleDto := &nextgen.NotificationRule{}

	if attr, ok := d.GetOk("org_id"); ok {
		notificationRuleDto.OrgIdentifier = attr.(string)
	}

	if attr, ok := d.GetOk("project_id"); ok {
		notificationRuleDto.ProjectIdentifier = attr.(string)
	}

	notificationRuleDto.Identifier = identifier

	if attr, ok := d.GetOk("request"); ok {
		request := attr.([]interface{})[0].(map[string]interface{})

		notificationRuleDto.Name = request["name"].(string)
		notificationRuleDto.Type_ = request["type"].(string)

		notificationMethod := request["notification_method"].(*schema.Set).List()[0].(map[string]interface{})

		notifChannel, err := getNotificationChannelByType(notificationMethod)
		if err != nil {
			return notificationRuleDto, err
		}
		notificationRuleDto.NotificationMethod = &notifChannel

		conditions := request["conditions"].([]interface{})
		notificationRuleConditions := make([]nextgen.NotificationRuleCondition, len(conditions))
		for i, condition := range conditions {
			item := condition.(map[string]interface{})
			nrc, err := getNotificationRuleConditionByType(item)
			if err != nil {
				return notificationRuleDto, err
			}
			notificationRuleConditions[i] = nrc
		}
		notificationRuleDto.Conditions = notificationRuleConditions
	}

	return notificationRuleDto, nil
}

func readNotificationRule(d *schema.ResourceData, notificationRuleResponse **nextgen.NotificationRuleResponse) {
	notificationRuleDto := &(*notificationRuleResponse).NotificationRule

	d.SetId((*notificationRuleDto).Identifier)

	d.Set("org_id", (*notificationRuleDto).OrgIdentifier)
	d.Set("project_id", (*notificationRuleDto).ProjectIdentifier)
	d.Set("identifier", (*notificationRuleDto).Identifier)
}
