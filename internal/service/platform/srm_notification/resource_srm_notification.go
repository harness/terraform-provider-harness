package srm_notification

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

func ResourceSrmNotification() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for creating a srm notificatione.",

		CreateContext: resourceSrmNotificationCreate,
		ReadContext:   resourceSrmNotificationRead,
		UpdateContext: resourceSrmNotificationUpdate,
		DeleteContext: resourceSrmNotificationDelete,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"org_id": {
				Description: "Identifier of the organization in which the Srm Notification is configured.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"project_id": {
				Description: "Identifier of the project in which the Srm Notification is configured.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"identifier": {
				Description: "Identifier of the Srm Notification",
				Type:        schema.TypeString,
				Required:    true,
			},
			"request": {
				Description: "Request for creating or updating a Srm Notification.",
				Type:        schema.TypeList,
				MinItems:    1,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Description: "Name for the Srm Notification.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"type": {
							Description: "Type of the Srm Notification.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"conditions": {
							Description: "Set of notification conditions for SRM Notifications.",
							Type:        schema.TypeSet,
							Required:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Description: "Type of the notification condition.",
										Type:        schema.TypeString,
										Required:    true,
									},
									"spec": {
										Description: "Specification of the condition. Depends on the type of the condition.",
										Type:        schema.TypeString,
										Required:    true,
									},
								},
							},
						},
						"notification_method": {
							Description: "Notification channel for the SRM Notification.",
							Type:        schema.TypeSet,
							MinItems:    1,
							MaxItems:    1,
							Required:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Description: "Type of the Notification channel.",
										Type:        schema.TypeString,
										Required:    true,
									},
									"spec": {
										Description: "Specification of the Notification Channel. Depends on the type of the Notification channel.",
										Type:        schema.TypeString,
										Optional:    true,
									},
								},
							},
						},
					},
				},
			},
		},
	}

	return resource
}

func resourceSrmNotificationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	ctx = context.WithValue(ctx, nextgen.ContextAccessToken, hh.EnvVars.BearerToken.Get())
	var accountIdentifier string
	accountIdentifier = c.AccountId
	createSrmNotificationRequest := buildSrmNotificationRequest(d)
	respCreate, httpRespCreate, errCreate := c.SrmNotificationApiService.SaveSrmNotification(ctx, accountIdentifier,
		&nextgen.SrmNotificationServiceSaveSrmNotificationOpts{
			Body: optional.NewInterface(createSrmNotificationRequest),
		})

	if errCreate != nil {
		return helpers.HandleApiError(errCreate, d, httpRespCreate)
	}

	readSrmNotification(d, &respCreate.Resource.NotificationRule)
	return nil
}

func resourceSrmNotificationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
		return helpers.HandleReadApiError(err, d, httpResp)
	}

	// Soft delete lookup error handling
	// https://harness.atlassian.net/browse/PL-23765
	if &resp == nil {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	readSrmNotification(d, &resp.Resource.NotificationRule)
	return nil
}

func resourceSrmNotificationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	ctx = context.WithValue(ctx, nextgen.ContextAccessToken, hh.EnvVars.BearerToken.Get())
	var accountIdentifier string
	accountIdentifier = c.AccountId
	identifier := d.Get("identifier").(string)
	orgIdentifier := d.Get("org_id").(string)
	projectIdentifier := d.Get("project_id").(string)
	updateSrmNotificationRequest := buildSrmNotificationRequest(d)
	respCreate, httpRespCreate, errCreate := c.SrmNotificationApiService.UpdateSrmNotification(ctx, accountIdentifier, orgIdentifier, projectIdentifier, identifier,
		&nextgen.SrmNotificationApiUpdateSrmNotificationOpts{
			Body: optional.NewInterface(updateSrmNotificationRequest),
		})

	if errCreate != nil {
		return helpers.HandleApiError(errCreate, d, httpRespCreate)
	}

	readSrmNotification(d, &respCreate.Resource.NotificationRule)
	return nil
}

func resourceSrmNotificationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

func buildSrmNotificationRequest(d *schema.ResourceData) *nextgen.NotificationRule {
	srmNotification := &nextgen.NotificationRule{}

	if attr, ok := d.GetOk("org_id"); ok {
		srmNotification.OrgIdentifier = attr.(string)
	}

	if attr, ok := d.GetOk("project_id"); ok {
		srmNotification.ProjectIdentifier = attr.(string)
	}

	if attr, ok := d.GetOk("identifier"); ok {
		srmNotification.Identifier = attr.(string)
	}

	if attr, ok := d.GetOk("request"); ok {
		request := attr.([]interface{})[0].(map[string]interface{})

		srmNotification.Name = request["name"].(string)
		srmNotification.Type_ = request["type"].(string)

		conditions := request["conditions"].(*schema.Set).List()
		hss := make([]nextgen.NotificationRuleCondition, len(conditions))
		for i, condition := range conditions {
			hs := condition.(map[string]interface{})
			notificationRuleConditionDto := getNotificationRuleConditionByType(hs)
			hss[i] = notificationRuleConditionDto
		}
		srmNotification.Conditions = hss

		notificationMethod := getNotificationChannelByType(request["notification_method"].(*schema.Set).List()[0].(map[string]interface{}))
		srmNotification.NotificationMethod = &notificationMethod
	}

	return srmNotification
}

func readSrmNotification(d *schema.ResourceData, srmNotificationResponse **nextgen.NotificationRule) {
	notificationRule := &(*srmNotificationResponse)

	d.SetId((*notificationRule).Identifier)

	d.Set("org_id", (*notificationRule).OrgIdentifier)
	d.Set("project_id", (*notificationRule).ProjectIdentifier)
	d.Set("identifier", (*notificationRule).Identifier)
}
