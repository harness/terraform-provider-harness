package monitored_service

import (
	"context"
	"github.com/antihax/optional"
	hh "github.com/harness/harness-go-sdk/harness/helpers"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strings"
)

func ResourceMonitoredService() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for creating a monitored service.",

		CreateContext: resourceMonitoredServiceCreate,
		ReadContext:   resourceMonitoredServiceRead,
		UpdateContext: resourceMonitoredServiceUpdate,
		DeleteContext: resourceMonitoredServiceDelete,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"org_id": {
				Description: "Identifier of the organization in which the monitored service is configured.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"project_id": {
				Description: "Identifier of the project in which the monitored service is configured.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"identifier": {
				Description: "Identifier of the monitored service.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"request": {
				Description: "Request for creating or updating a monitored service.",
				Type:        schema.TypeList,
				MinItems:    1,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Description: "Name for the monitored service.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"type": {
							Description: "Type of the monitored service.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"description": {
							Description: "Description for the monitored service.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"service_ref": {
							Description: "Service reference for the monitored service.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"environment_ref": {
							Description: "Environment in which the service is deployed.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"environment_ref_list": {
							Description: "Environment reference list for the monitored service.",
							Type:        schema.TypeList,
							Optional:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"tags": {
							Description: "Tags for the monitored service. comma-separated key value string pairs.",
							Type:        schema.TypeSet,
							Optional:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"health_sources": {
							Description: "Set of health sources for the monitored service.",
							Type:        schema.TypeSet,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Description: "Name of the health source.",
										Type:        schema.TypeString,
										Required:    true,
									},
									"identifier": {
										Description: "Identifier of the health source.",
										Type:        schema.TypeString,
										Required:    true,
									},
									"type": {
										Description: "Type of the health source.",
										Type:        schema.TypeString,
										Required:    true,
									},
									"version": {
										Description: "Version of the health source.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"spec": {
										Description: "Specification of the health source. Depends on the type of the health source.",
										Type:        schema.TypeString,
										Required:    true,
									},
								},
							},
						},
						"change_sources": {
							Description: "Set of change sources for the monitored service.",
							Type:        schema.TypeSet,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Description: "Name of the change source.",
										Type:        schema.TypeString,
										Required:    true,
									},
									"identifier": {
										Description: "Identifier of the change source.",
										Type:        schema.TypeString,
										Required:    true,
									},
									"type": {
										Description: "Type of the change source.",
										Type:        schema.TypeString,
										Required:    true,
									},
									"enabled": {
										Description: "Enable or disable the change source.",
										Type:        schema.TypeBool,
										Optional:    true,
									},
									"spec": {
										Description: "Specification of the change source. Depends on the type of the change source.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"category": {
										Description: "Category of the change source.",
										Type:        schema.TypeString,
										Required:    true,
									},
								},
							},
						},
						"dependencies": {
							Description: "Dependencies of the monitored service.",
							Type:        schema.TypeSet,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"monitored_service_identifier": {
										Description: "Monitored service identifier of the dependency.",
										Type:        schema.TypeString,
										Required:    true,
									},
									"type": {
										Description: "Type of the service dependency.",
										Type:        schema.TypeString,
										Required:    true,
									},
									"dependency_metadata": {
										Description: "Dependency metadata for the monitored service.",
										Type:        schema.TypeString,
										Optional:    true,
									},
								},
							},
						},
						"notification_rule_refs": {
							Description: "Notification rule references for the monitored service.",
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"notification_rule_ref": {
										Description: "Notification rule reference for the monitored service.",
										Type:        schema.TypeString,
										Required:    true,
									},
									"enabled": {
										Description: "Enable or disable notification rule reference for the monitored service.",
										Type:        schema.TypeBool,
										Required:    true,
									},
								},
							},
						},
						"template_ref": {
							Description: "Template reference for the monitored service.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"version_label": {
							Description: "Template version label for the monitored service.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"enabled": {
							Description: "Enable or disable the monitored service.",
							Type:        schema.TypeBool,
							Optional:    true,
							Deprecated:  "enabled field is deprecated",
						},
					},
				},
			},
		},
	}

	return resource
}

func resourceMonitoredServiceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	ctx = context.WithValue(ctx, nextgen.ContextAccessToken, hh.EnvVars.BearerToken.Get())
	var accountIdentifier string
	accountIdentifier = c.AccountId
	createMonitoredServiceRequest := buildMonitoredServiceRequest(d)
	respCreate, httpRespCreate, errCreate := c.MonitoredServiceApi.SaveMonitoredService(ctx, accountIdentifier,
		&nextgen.MonitoredServiceApiSaveMonitoredServiceOpts{
			Body: optional.NewInterface(createMonitoredServiceRequest),
		})

	if errCreate != nil {
		return helpers.HandleApiError(errCreate, d, httpRespCreate)
	}

	readMonitoredService(d, &respCreate.Resource)
	return nil
}

func resourceMonitoredServiceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	resp, httpResp, err := c.MonitoredServiceApi.GetMonitoredService(ctx, identifier, accountIdentifier, orgIdentifier, projectIdentifier)

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

	readMonitoredService(d, &resp.Data)
	return nil
}

func resourceMonitoredServiceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	ctx = context.WithValue(ctx, nextgen.ContextAccessToken, hh.EnvVars.BearerToken.Get())
	var accountIdentifier string
	accountIdentifier = c.AccountId
	identifier := d.Get("identifier").(string)
	updateMonitoredServiceRequest := buildMonitoredServiceRequest(d)
	respCreate, httpRespCreate, errCreate := c.MonitoredServiceApi.UpdateMonitoredService(ctx, accountIdentifier, identifier,
		&nextgen.MonitoredServiceApiUpdateMonitoredServiceOpts{
			Body: optional.NewInterface(updateMonitoredServiceRequest),
		})

	if errCreate != nil {
		return helpers.HandleApiError(errCreate, d, httpRespCreate)
	}

	readMonitoredService(d, &respCreate.Resource)
	return nil
}

func resourceMonitoredServiceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	_, httpResp, err := c.MonitoredServiceApi.DeleteMonitoredService(ctx, accountIdentifier, orgIdentifier, projectIdentifier, identifier)

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}
	return nil
}

func buildMonitoredServiceRequest(d *schema.ResourceData) *nextgen.MonitoredServiceDto {
	monitoredService := &nextgen.MonitoredServiceDto{}

	if attr, ok := d.GetOk("org_id"); ok {
		monitoredService.OrgIdentifier = attr.(string)
	}

	if attr, ok := d.GetOk("project_id"); ok {
		monitoredService.ProjectIdentifier = attr.(string)
	}

	if attr, ok := d.GetOk("identifier"); ok {
		monitoredService.Identifier = attr.(string)
	}

	if attr, ok := d.GetOk("request"); ok {
		request := attr.([]interface{})[0].(map[string]interface{})

		monitoredService.Name = request["name"].(string)
		monitoredService.Type_ = request["type"].(string)
		monitoredService.Description = request["description"].(string)
		monitoredService.ServiceRef = request["service_ref"].(string)
		monitoredService.EnvironmentRef = request["environment_ref"].(string)

		environmentRefListReq := request["environment_ref_list"].([]interface{})
		environmentRefList := make([]string, len(environmentRefListReq))
		for i, environmentRef := range environmentRefListReq {
			environmentRefList[i] = environmentRef.(string)
		}
		monitoredService.EnvironmentRefList = environmentRefList

		tags := map[string]string{}
		for _, t := range request["tags"].(*schema.Set).List() {
			tagStr := t.(string)
			parts := strings.Split(tagStr, ":")
			tags[parts[0]] = parts[1]
		}
		monitoredService.Tags = tags

		healthSources := request["health_sources"].(*schema.Set).List()
		hss := make([]nextgen.HealthSource, len(healthSources))
		for i, healthSource := range healthSources {
			hs := healthSource.(map[string]interface{})
			healthSourceDto := getHealthSourceByType(hs)
			hss[i] = healthSourceDto
		}

		changeSources := request["change_sources"].(*schema.Set).List()
		csDto := make([]nextgen.ChangeSourceDto, len(changeSources))
		for i, changeSource := range changeSources {
			cs := changeSource.(map[string]interface{})
			changeSourceDto := getChangeSourceByType(cs)
			csDto[i] = changeSourceDto
		}

		monitoredService.Sources = &nextgen.Sources{
			HealthSources: hss,
			ChangeSources: csDto,
		}

		dependencies := request["dependencies"].(*schema.Set).List()
		serviceDependencyDto := make([]nextgen.ServiceDependencyDto, len(dependencies))
		for i, dependency := range dependencies {
			sd := dependency.(map[string]interface{})
			serviceDependency := getServiceDependencyByType(sd)
			serviceDependencyDto[i] = serviceDependency
		}
		monitoredService.Dependencies = serviceDependencyDto

		notificationRuleRefsReq := request["notification_rule_refs"].([]interface{})
		notificationRuleRefs := make([]nextgen.NotificationRuleRefDto, len(notificationRuleRefsReq))
		for i, notificationRuleRef := range notificationRuleRefsReq {
			test := notificationRuleRef.(map[string]interface{})
			notificationRuleRefDto := &nextgen.NotificationRuleRefDto{
				NotificationRuleRef: test["notification_rule_ref"].(string),
				Enabled:             test["enabled"].(bool),
			}
			notificationRuleRefs[i] = *notificationRuleRefDto
		}
		monitoredService.NotificationRuleRefs = notificationRuleRefs

		monitoredService.Template = &nextgen.TemplateDto{
			TemplateRef:  request["template_ref"].(string),
			VersionLabel: request["version_label"].(string),
		}

	}

	return monitoredService
}

func readMonitoredService(d *schema.ResourceData, monitoredServiceResponse **nextgen.MonitoredServiceResponse) {
	monitoredService := &(*monitoredServiceResponse).MonitoredService

	d.SetId((*monitoredService).Identifier)

	d.Set("org_id", (*monitoredService).OrgIdentifier)
	d.Set("project_id", (*monitoredService).ProjectIdentifier)
	d.Set("identifier", (*monitoredService).Identifier)
}
