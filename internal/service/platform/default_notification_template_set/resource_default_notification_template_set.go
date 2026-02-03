package default_notification_template_set

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

func ResourceDefaultNotificationTemplateSet() *schema.Resource {
	return &schema.Resource{
		Description:   "Resource for creating a Harness Default Notification Template Set",
		CreateContext: resourceDefaultNotificationTemplateSetCreate,
		ReadContext:   resourceDefaultNotificationTemplateSetRead,
		UpdateContext: resourceDefaultNotificationTemplateSetUpdate,
		DeleteContext: resourceDefaultNotificationTemplateSetDelete,
		Schema: map[string]*schema.Schema{
			"org_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Unique identifier of the organization.",
			},
			"project_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Unique identifier of the project.",
			},
			"org": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Unique identifier of the organization. Deprecated: Use org_id instead.",
				Deprecated:  "This field is deprecated and will be removed in a future release. Please use 'org_id' instead.",
			},
			"project": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Unique identifier of the project. Deprecated: Use project_id instead.",
				Deprecated:  "This field is deprecated and will be removed in a future release. Please use 'project_id' instead.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of Default Notification Template Set",
			},
			"identifier": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Identifier of Default Notification Template Set",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description for Default Notification Template Set",
			},
			"notification_entity": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Type of the entity (e.g. PIPELINE, SERVICE, etc.)",
			},
			"notification_channel_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Type of channel (e.g. SLACK, EMAIL, etc.)",
			},
			"event_template_configuration_set": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "Set of event-template configurations",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"notification_events": {
							Type:        schema.TypeList,
							Required:    true,
							Description: "List of notification events like PIPELINE_START",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"template": {
							Type:        schema.TypeList,
							Required:    true,
							MaxItems:    1,
							Description: "Template reference configuration",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"template_ref": {
										Type:     schema.TypeString,
										Required: true,
									},
									"version_label": {
										Type:     schema.TypeString,
										Required: true,
									},
									"variables": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of variables passed to the template",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:     schema.TypeString,
													Required: true,
												},
												"value": {
													Type:     schema.TypeString,
													Required: true,
												},
												"type": {
													Type:     schema.TypeString,
													Required: true,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"tags": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Key-value tags",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"last_modified": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Timestamp when the default notification template set was last modified.",
			},
			"created": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Timestamp when the default notification template set was created.",
			},
		},
	}
}

func resourceDefaultNotificationTemplateSetCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	accountID := c.AccountId
	scope := getScope(d)
	defaultNotificationTemplateSetDto := buildDefaultNotificationTemplateSetRequest(d)
	var resp nextgen.DefaultNotificationTemplateSetResponse
	var httpResp *http.Response
	var err error
	switch scope.scope {
	case Project:
		resp, httpResp, err = c.ProjectDefaultNotificationTemplateSetApi.CreateProjectDefaultNotificationTemplateSet(ctx, scope.org, scope.project,
			&nextgen.ProjectDefaultNotificationTemplateSetApiCreateProjectDefaultNotificationTemplateSetOpts{
				HarnessAccount: optional.NewString(accountID),
				Body:           optional.NewInterface(defaultNotificationTemplateSetDto),
			})
	case Org:
		resp, httpResp, err = c.OrgDefaultNotificationTemplateSetApi.CreateOrgDefaultNotificationTemplateSet(ctx, scope.org,
			&nextgen.OrgDefaultNotificationTemplateSetApiCreateOrgDefaultNotificationTemplateSetOpts{
				HarnessAccount: optional.NewString(accountID),
				Body:           optional.NewInterface(defaultNotificationTemplateSetDto),
			})
	default:
		resp, httpResp, err = c.AccountDefaultNotificationTemplateSetApi.CreateAccountDefaultNotificationTemplateSet(ctx,
			&nextgen.AccountDefaultNotificationTemplateSetApiCreateAccountDefaultNotificationTemplateSetOpts{
				HarnessAccount: optional.NewString(accountID),
				Body:           optional.NewInterface(defaultNotificationTemplateSetDto),
			})
	}
	if err != nil {
		return helpers.HandleReadApiError(err, d, httpResp)
	}

	readDefaultNotificationTemplateSet(d, resp)

	return nil
}

func resourceDefaultNotificationTemplateSetUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	accountID := c.AccountId
	identifier := d.Get("identifier").(string)
	scope := getScope(d)
	defaultNotificationTemplateSetDto := buildDefaultNotificationTemplateSetRequest(d)
	var resp nextgen.DefaultNotificationTemplateSetResponse
	var httpResp *http.Response
	var err error
	switch scope.scope {
	case Project:
		resp, httpResp, err = c.ProjectDefaultNotificationTemplateSetApi.UpdateProjectDefaultNotificationTemplateSet(ctx, identifier, scope.org, scope.project,
			&nextgen.ProjectDefaultNotificationTemplateSetApiUpdateProjectDefaultNotificationTemplateSetOpts{
				HarnessAccount: optional.NewString(accountID),
				Body:           optional.NewInterface(defaultNotificationTemplateSetDto),
			})
	case Org:
		resp, httpResp, err = c.OrgDefaultNotificationTemplateSetApi.UpdateOrgDefaultNotificationTemplateSet(ctx, identifier, scope.org,
			&nextgen.OrgDefaultNotificationTemplateSetApiUpdateOrgDefaultNotificationTemplateSetOpts{
				HarnessAccount: optional.NewString(accountID),
				Body:           optional.NewInterface(defaultNotificationTemplateSetDto),
			})
	default:
		resp, httpResp, err = c.AccountDefaultNotificationTemplateSetApi.UpdateAccountDefaultNotificationTemplateSet(ctx, identifier,
			&nextgen.AccountDefaultNotificationTemplateSetApiUpdateAccountDefaultNotificationTemplateSetOpts{
				HarnessAccount: optional.NewString(accountID),
				Body:           optional.NewInterface(defaultNotificationTemplateSetDto),
			})
	}
	if err != nil {
		return helpers.HandleReadApiError(err, d, httpResp)
	}

	readDefaultNotificationTemplateSet(d, resp)

	return nil
}

func resourceDefaultNotificationTemplateSetDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	accountID := c.AccountId
	identifier := d.Get("identifier").(string)
	scope := getScope(d)
	var httpResp *http.Response
	var err error
	switch scope.scope {
	case Project:
		httpResp, err = c.ProjectDefaultNotificationTemplateSetApi.DeleteProjectDefaultNotificationTemplateSet(ctx, identifier, scope.org, scope.project,
			&nextgen.ProjectDefaultNotificationTemplateSetApiDeleteProjectDefaultNotificationTemplateSetOpts{
				HarnessAccount: optional.NewString(accountID),
			})
	case Org:
		httpResp, err = c.OrgDefaultNotificationTemplateSetApi.DeleteOrgDefaultNotificationTemplateSet(ctx, identifier, scope.org,
			&nextgen.OrgDefaultNotificationTemplateSetApiDeleteOrgDefaultNotificationTemplateSetOpts{
				HarnessAccount: optional.NewString(accountID),
			})
	default:
		httpResp, err = c.AccountDefaultNotificationTemplateSetApi.DeleteAccountDefaultNotificationTemplateSet(ctx, identifier,
			&nextgen.AccountDefaultNotificationTemplateSetApiDeleteAccountDefaultNotificationTemplateSetOpts{
				HarnessAccount: optional.NewString(accountID),
			})
	}
	if err != nil {
		return helpers.HandleReadApiError(err, d, httpResp)
	}

	return nil
}

func resourceDefaultNotificationTemplateSetRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	accountID := c.AccountId
	identifier := d.Get("identifier").(string)
	scope := getScope(d)
	var resp nextgen.DefaultNotificationTemplateSetResponse
	var httpResp *http.Response
	var err error
	switch scope.scope {
	case Project:
		resp, httpResp, err = c.ProjectDefaultNotificationTemplateSetApi.GetProjectDefaultNotificationTemplateSet(ctx, identifier, scope.org, scope.project,
			&nextgen.ProjectDefaultNotificationTemplateSetApiGetProjectDefaultNotificationTemplateSetOpts{
				HarnessAccount: optional.NewString(accountID),
			})
	case Org:
		resp, httpResp, err = c.OrgDefaultNotificationTemplateSetApi.GetOrgDefaultNotificationTemplateSet(ctx, identifier, scope.org,
			&nextgen.OrgDefaultNotificationTemplateSetApiGetOrgDefaultNotificationTemplateSetOpts{
				HarnessAccount: optional.NewString(accountID),
			})
	default:
		resp, httpResp, err = c.AccountDefaultNotificationTemplateSetApi.GetAccountDefaultNotificationTemplateSet(ctx, identifier,
			&nextgen.AccountDefaultNotificationTemplateSetApiGetAccountDefaultNotificationTemplateSetOpts{
				HarnessAccount: optional.NewString(accountID),
			})
	}
	if err != nil {
		return helpers.HandleReadApiError(err, d, httpResp)
	}

	readDefaultNotificationTemplateSet(d, resp)

	return nil
}

func readDefaultNotificationTemplateSet(d *schema.ResourceData, response nextgen.DefaultNotificationTemplateSetResponse) diag.Diagnostics {
	dto := response.DefaultNotificationTemplateSet

	d.SetId(dto.Identifier)

	d.Set("name", dto.Name)
	d.Set("identifier", dto.Identifier)
	d.Set("description", dto.Description)
	d.Set("notification_entity", dto.NotificationEntity)
	d.Set("notification_channel_type", dto.NotificationChannelType)
	d.Set("created", response.Created)
	d.Set("last_modified", response.Updated)

	// Set tags
	if dto.Tags != nil {
		d.Set("tags", dto.Tags)
	}

	// Map event_template_configuration_set
	var eventTemplates []map[string]interface{}
	for _, cfg := range dto.EventTemplateConfigurationSet {
		template := map[string]interface{}{
			"template_ref":  cfg.Template.TemplateRef,
			"version_label": cfg.Template.VersionLabel,
		}

		// Map variables
		if len(cfg.Template.Variables) > 0 {
			var variables []map[string]interface{}
			for _, v := range cfg.Template.Variables {
				variables = append(variables, map[string]interface{}{
					"name":  v.Name,
					"value": v.Value,
					"type":  v.Type_,
				})
			}
			template["variables"] = variables
		}

		eventTemplates = append(eventTemplates, map[string]interface{}{
			"notification_events": cfg.NotificationEvents,
			"template":            []interface{}{template},
		})
	}
	d.Set("event_template_configuration_set", eventTemplates)

	return nil
}

func buildDefaultNotificationTemplateSetRequest(d *schema.ResourceData) *nextgen.DefaultNotificationTemplateSetDto {
	eventTemplateConfigsRaw := d.Get("event_template_configuration_set").([]interface{})
	var eventTemplateConfigs []nextgen.EventTemplateConfigurationDto

	for _, etc := range eventTemplateConfigsRaw {
		etcMap := etc.(map[string]interface{})

		// Parse notification events
		events := expandStringList(etcMap["notification_events"])

		// Parse template object
		templateList := etcMap["template"].([]interface{})
		var customTemplateDto nextgen.CustomNotificationTemplateDto
		if len(templateList) > 0 {
			tpl := templateList[0].(map[string]interface{})

			// Parse variables if any
			var vars []nextgen.NotificationTemplateInputsDto
			if rawVars, ok := tpl["variables"]; ok {
				for _, v := range rawVars.([]interface{}) {
					vMap := v.(map[string]interface{})
					vars = append(vars, nextgen.NotificationTemplateInputsDto{
						Name:  vMap["name"].(string),
						Value: vMap["value"].(string),
						Type_: vMap["type"].(string),
					})
				}
			}

			customTemplateDto = nextgen.CustomNotificationTemplateDto{
				TemplateRef:  tpl["template_ref"].(string),
				VersionLabel: tpl["version_label"].(string),
				Variables:    vars,
			}
		}

		eventTemplateConfigs = append(eventTemplateConfigs, nextgen.EventTemplateConfigurationDto{
			NotificationEvents: events,
			Template:           &customTemplateDto,
		})
	}

	return &nextgen.DefaultNotificationTemplateSetDto{
		Name:                          d.Get("name").(string),
		Identifier:                    d.Get("identifier").(string),
		Description:                   d.Get("description").(string),
		NotificationEntity:            ptrToNotificationEntity(d.Get("notification_entity").(string)),
		NotificationChannelType:       ptrToChannelType(d.Get("notification_channel_type").(string)),
		EventTemplateConfigurationSet: eventTemplateConfigs,
		Tags:                          expandTags(d.Get("tags").(map[string]interface{})),
	}
}

func ptrToNotificationEntity(s string) *nextgen.NotificationEntity {
	ne := nextgen.NotificationEntity(s)
	return &ne
}

func ptrToChannelType(s string) *nextgen.ChannelType {
	ct := nextgen.ChannelType(s)
	return &ct
}

func expandTags(tagMap map[string]interface{}) map[string]string {
	tags := make(map[string]string)
	for k, v := range tagMap {
		tags[k] = v.(string)
	}
	return tags
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

	// Support both org_id (preferred) and deprecated org
	if attr, ok := d.GetOk("org_id"); ok {
		org = (attr.(string))
	} else if attr, ok := d.GetOk("org"); ok {
		org = (attr.(string))
	}

	// Support both project_id (preferred) and deprecated project
	if attr, ok := d.GetOk("project_id"); ok {
		project = (attr.(string))
	} else if attr, ok := d.GetOk("project"); ok {
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
