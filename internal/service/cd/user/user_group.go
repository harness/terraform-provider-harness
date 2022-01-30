package user

import (
	"context"
	"fmt"
	"strings"

	sdk "github.com/harness-io/harness-go-sdk"
	"github.com/harness-io/harness-go-sdk/harness/cd/graphql"
	"github.com/harness-io/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var actionsAll = []string{"CREATE", "READ", "UPDATE", "DELETE", "EXECUTE_WORKFLOW", "EXECUTE_PIPELINE", "ROLLBACK_WORKFLOW"}
var standardActions = []string{"CREATE", "READ", "UPDATE", "DELETE"}
var deploymentActions = []string{"READ", "EXECUTE_WORKFLOW", "EXECUTE_PIPELINE", "ROLLBACK_WORKFLOW"}

func ResourceUserGroup() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for creating a Harness user group",

		CreateContext: resourceUserGroupCreate,
		ReadContext:   resourceUserGroupRead,
		UpdateContext: resourceUserGroupUpdate,
		DeleteContext: resourceUserGroupDelete,

		Schema: map[string]*schema.Schema{
			"id": {
				Description: "Unique identifier of the user group.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"name": {
				Description: "The name of the user group.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"description": {
				Description: "The description of the user group.",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
			},
			"imported_by_scim": {
				Description: "Indicates whether the user group was imported by SCIM.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"is_sso_linked": {
				Description: "Indicates whether the user group is linked to an SSO provider.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"notification_settings": {
				Description: "The notification settings of the user group.",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"group_email_addresses": {
							Description: "The email addresses of the user group.",
							Type:        schema.TypeList,
							Optional:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"microsoft_teams_webhook_url": {
							Description: "The Microsoft Teams webhook URL of the user group.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"send_mail_to_new_members": {
							Description: "Indicates whether an email is sent when a new user is added to the group.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"send_notifications_to_members": {
							Description: "Enable this setting to have notifications sent to the members of this group.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"slack_channel": {
							Description: "The Slack channel to send notifications to.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"slack_webhook_url": {
							Description: "The Slack webhook URL to send notifications to.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
			"permissions": {
				Description: "The permissions of the user group.",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account_permissions": {
							Description: fmt.Sprintf("The account permissions of the user group. Valid options are %s", strings.Join(graphql.AccountPermissionTypeValues, ", ")),
							Type:        schema.TypeSet,
							Optional:    true,
							MinItems:    1,
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"app_permissions": {
							Description: "Application specific permissions",
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"all": {
										Description: "The permission to perform actions against all resources.",
										Type:        schema.TypeSet,
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"app_ids": {
													Description: "The application IDs to which the permission applies. Leave empty to apply to all applications.",
													Type:        schema.TypeSet,
													Optional:    true,
													Elem:        &schema.Schema{Type: schema.TypeString},
												},
												"actions": {
													Description: fmt.Sprintf("The actions allowed to be performed. Valid options are %s", strings.Join(actionsAll, ", ")),
													Type:        schema.TypeSet,
													Required:    true,
													MinItems:    1,
													Elem:        &schema.Schema{Type: schema.TypeString},
												},
											},
										},
									},
									"workflow": {
										Description: "Permission configuration to perform actions against workflows.",
										Type:        schema.TypeSet,
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"app_ids": {
													Description: "The application IDs to which the permission applies. Leave empty to apply to all applications.",
													Type:        schema.TypeSet,
													Optional:    true,
													Elem:        &schema.Schema{Type: schema.TypeString},
												},
												"filters": {
													Description: fmt.Sprintf("The filters to apply to the action. Valid options are: %s.", strings.Join(graphql.WorkflowPermissionFiltersSlice, ", ")),
													Type:        schema.TypeSet,
													Required:    true,
													MinItems:    1,
													Elem:        &schema.Schema{Type: schema.TypeString},
												},
												"actions": {
													Description: fmt.Sprintf("The actions allowed to be performed. Valid options are %s", strings.Join(standardActions, ", ")),
													Type:        schema.TypeSet,
													Required:    true,
													MinItems:    1,
													Elem:        &schema.Schema{Type: schema.TypeString},
												},
											},
										},
									},
									"deployment": {
										Description: "Permission configuration to perform actions against deployments.",
										Type:        schema.TypeSet,
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"app_ids": {
													Description: "The application IDs to which the permission applies. Leave empty to apply to all applications.",
													Type:        schema.TypeSet,
													Optional:    true,
													Elem:        &schema.Schema{Type: schema.TypeString},
												},
												"env_ids": {
													Description: "The environment IDs to which the permission applies. Leave empty to apply to all environments.",
													Type:        schema.TypeSet,
													Optional:    true,
													Elem:        &schema.Schema{Type: schema.TypeString},
												},
												"filters": {
													Description: fmt.Sprintf("The filters to apply to the action. Valid options are: %s.", strings.Join(graphql.DeploymentPermissionFiltersSlice, ", ")),
													Type:        schema.TypeSet,
													Required:    true,
													MinItems:    1,
													Elem:        &schema.Schema{Type: schema.TypeString},
												},
												"actions": {
													Description: fmt.Sprintf("The actions allowed to be performed. Valid options are %s", strings.Join(deploymentActions, ", ")),
													Type:        schema.TypeSet,
													Required:    true,
													MinItems:    1,
													Elem:        &schema.Schema{Type: schema.TypeString},
												},
											},
										},
									},
									"environment": {
										Description: "Permission configuration to perform actions against workflows.",
										Type:        schema.TypeSet,
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"app_ids": {
													Description: "The application IDs to which the permission applies. Leave empty to apply to all applications.",
													Type:        schema.TypeSet,
													Optional:    true,
													Elem:        &schema.Schema{Type: schema.TypeString},
												},
												"env_ids": {
													Description: "The environment IDs to which the permission applies. Leave empty to apply to all environments.",
													Type:        schema.TypeSet,
													Optional:    true,
													Elem:        &schema.Schema{Type: schema.TypeString},
												},
												"filters": {
													Description: fmt.Sprintf("The filters to apply to the action. Valid options are: %s.", strings.Join(graphql.EnvFiltersSlice, ", ")),
													Type:        schema.TypeSet,
													Required:    true,
													MinItems:    1,
													Elem:        &schema.Schema{Type: schema.TypeString},
												},
												"actions": {
													Description: fmt.Sprintf("The actions allowed to be performed. Valid options are %s", strings.Join(standardActions, ", ")),
													Type:        schema.TypeSet,
													Required:    true,
													MinItems:    1,
													Elem:        &schema.Schema{Type: schema.TypeString},
												},
											},
										},
									},
									"pipeline": {
										Description: "Permission configuration to perform actions against pipelines.",
										Type:        schema.TypeSet,
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"app_ids": {
													Description: "The application IDs to which the permission applies. Leave empty to apply to all applications.",
													Type:        schema.TypeSet,
													Optional:    true,
													Elem:        &schema.Schema{Type: schema.TypeString},
												},
												"env_ids": {
													Description: "The environment IDs to which the permission applies. Leave empty to apply to all environments.",
													Type:        schema.TypeSet,
													Optional:    true,
													Elem:        &schema.Schema{Type: schema.TypeString},
												},
												"filters": {
													Description: fmt.Sprintf("The filters to apply to the action. Valid options are: %s.", strings.Join(graphql.PipelinePermissionFiltersSlice, ", ")),
													Type:        schema.TypeSet,
													Required:    true,
													MinItems:    1,
													Elem:        &schema.Schema{Type: schema.TypeString},
												},
												"actions": {
													Description: fmt.Sprintf("The actions allowed to be performed. Valid options are %s", standardActions),
													Type:        schema.TypeSet,
													Required:    true,
													MinItems:    1,
													Elem:        &schema.Schema{Type: schema.TypeString},
												},
											},
										},
									},
									"provisioner": {
										Description: "Permission configuration to perform actions against provisioners.",
										Type:        schema.TypeSet,
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"app_ids": {
													Description: "The application IDs to which the permission applies. Leave empty to apply to all applications.",
													Type:        schema.TypeSet,
													Optional:    true,
													Elem:        &schema.Schema{Type: schema.TypeString},
												},
												"provisioner_ids": {
													Description: "The provisioner IDs to which the permission applies. Leave empty to apply to all provisioners.",
													Type:        schema.TypeSet,
													Optional:    true,
													Elem:        &schema.Schema{Type: schema.TypeString},
												},
												"actions": {
													Description: fmt.Sprintf("The actions allowed to be performed. Valid options are %s", strings.Join(standardActions, ", ")),
													Type:        schema.TypeSet,
													Required:    true,
													MinItems:    1,
													Elem:        &schema.Schema{Type: schema.TypeString},
												},
											},
										},
									},
									"service": {
										Description: "Permission configuration to perform actions against services.",
										Type:        schema.TypeSet,
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"app_ids": {
													Description: "The application IDs to which the permission applies. Leave empty to apply to all applications.",
													Type:        schema.TypeSet,
													Optional:    true,
													Elem:        &schema.Schema{Type: schema.TypeString},
												},
												"service_ids": {
													Description: "The service IDs to which the permission applies. Leave empty to apply to all services.",
													Type:        schema.TypeSet,
													Optional:    true,
													Elem:        &schema.Schema{Type: schema.TypeString},
												},
												"actions": {
													Description: fmt.Sprintf("The actions allowed to be performed. Valid options are %s", strings.Join(standardActions, ", ")),
													Type:        schema.TypeSet,
													Required:    true,
													MinItems:    1,
													Elem:        &schema.Schema{Type: schema.TypeString},
												},
											},
										},
									},
									"template": {
										Description: "Permission configuration to perform actions against templates.",
										Type:        schema.TypeSet,
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"app_ids": {
													Description: "The application IDs to which the permission applies. Leave empty to apply to all applications.",
													Type:        schema.TypeSet,
													Optional:    true,
													Elem:        &schema.Schema{Type: schema.TypeString},
												},
												"template_ids": {
													Description: "The template IDs to which the permission applies. Leave empty to apply to all environments.",
													Type:        schema.TypeSet,
													Optional:    true,
													Elem:        &schema.Schema{Type: schema.TypeString},
												},
												"actions": {
													Description: fmt.Sprintf("The actions allowed to be performed. Valid options are %s", strings.Join(standardActions, ", ")),
													Type:        schema.TypeSet,
													Required:    true,
													MinItems:    1,
													Elem:        &schema.Schema{Type: schema.TypeString},
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
			"ldap_settings": {
				Description:   "The LDAP settings for the user group.",
				Type:          schema.TypeList,
				Optional:      true,
				MaxItems:      1,
				ConflictsWith: []string{"saml_settings"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"group_dn": {
							Description: "The group DN of the LDAP user group.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"group_name": {
							Description: "The group name of the LDAP user group.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"sso_provider_id": {
							Description: "The ID of the SSO provider.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
			"saml_settings": {
				Description:   "The SAML settings for the user group.",
				Type:          schema.TypeList,
				Optional:      true,
				MaxItems:      1,
				ConflictsWith: []string{"ldap_settings"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"group_name": {
							Description: "The group name of the SAML user group.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"sso_provider_id": {
							Description: "The ID of the SSO provider.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
		},

		Importer: &schema.ResourceImporter{StateContext: schema.ImportStatePassthroughContext},
	}
}

func resourceUserGroupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*sdk.Session)

	// if err := validateUserGroupPermissions(d); err != nil {
	// 	return diag.FromErr(err)
	// }

	input := &graphql.UserGroup{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
	}

	expandLDAPSettings(d, input)
	expandSAMLSettings(d, input)
	expandNotificationSettings(d, input)
	expandPermissions(d, input)

	userGroup, err := c.CDClient.UserClient.CreateUserGroup(input)
	if err != nil {
		return diag.FromErr(err)
	}

	// Computed fields
	d.SetId(userGroup.Id)
	d.Set("imported_by_scim", userGroup.ImportedBySCIM)
	d.Set("is_sso_linked", userGroup.IsSSOLinked)

	return nil
}

func resourceUserGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*sdk.Session)

	id := d.Get("id").(string)

	userGroup, err := c.CDClient.UserClient.GetUserGroupById(id)
	if err != nil {
		return diag.FromErr(err)
	}

	if userGroup == nil {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	return readUserGroup(d, userGroup)
}

func readUserGroup(d *schema.ResourceData, userGroup *graphql.UserGroup) diag.Diagnostics {
	d.SetId(userGroup.Id)
	d.Set("name", userGroup.Name)
	d.Set("description", userGroup.Description)
	d.Set("imported_by_scim", userGroup.ImportedBySCIM)
	d.Set("is_sso_linked", userGroup.IsSSOLinked)

	if samlSettings := flattenSAMLSettings(d, userGroup.SAMLSettings); len(samlSettings) > 0 {
		d.Set("saml_settings", samlSettings)
	}

	if ldapSettings := flattenLDAPSettings(d, userGroup.LDAPSettings); len(ldapSettings) > 0 {
		d.Set("ldap_settings", ldapSettings)
	}

	if notificationSettings := flattenNotificationSettings(d, userGroup.NotificationSettings); len(notificationSettings) > 0 {
		d.Set("notification_settings", notificationSettings)
	}

	if permissions := flattenPermissions(d, userGroup.Permissions); len(permissions) > 0 {
		d.Set("permissions", permissions)
	}

	return nil
}

func resourceUserGroupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*sdk.Session)

	input := &graphql.UserGroup{
		Id:   d.Id(),
		Name: d.Get("name").(string),
	}

	expandLDAPSettings(d, input)
	expandSAMLSettings(d, input)
	expandNotificationSettings(d, input)
	expandPermissions(d, input)

	userGroup, err := c.CDClient.UserClient.UpdateUserGroup(input)
	if err != nil {
		return diag.FromErr(err)
	}

	return readUserGroup(d, userGroup)
}

func resourceUserGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*sdk.Session)

	if err := c.CDClient.UserClient.DeleteUserGroup(d.Id()); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func expandAccountPermissions(d []interface{}, input *graphql.UserGroupPermissions) {
	if len(d) == 0 {
		return
	}

	permissionTypes := []graphql.AccountPermissionType{}
	for _, v := range d {
		permissionTypes = append(permissionTypes, graphql.AccountPermissionType(v.(string)))
	}
	input.AccountPermissions = &graphql.AccountPermissions{
		AccountPermissionTypes: permissionTypes,
	}
}

func expandTemplateAppPermissions(d []interface{}, input *graphql.UserGroupPermissions) {
	if len(d) == 0 {
		return
	}

	for _, v := range d {
		permission := &graphql.AppPermission{
			PermissionType: graphql.AppPermissionTypes.Template,
			Applications:   &graphql.AppFilter{},
			Actions:        []graphql.Action{},
			Templates:      &graphql.TemplatePermissionFilter{},
		}

		permissionConfig := v.(map[string]interface{})
		if attr := permissionConfig["app_ids"]; attr != nil {
			for _, appId := range attr.(*schema.Set).List() {
				permission.Applications.AppIds = append(permission.Applications.AppIds, appId.(string))
			}
		}

		if len(permission.Applications.AppIds) == 0 {
			permission.Applications.FilterType = graphql.FilterTypes.All
		}

		if attr := permissionConfig["actions"]; attr != nil {
			for _, action := range attr.(*schema.Set).List() {
				permission.Actions = append(permission.Actions, graphql.Action(action.(string)))
			}
		}

		if attr := permissionConfig["template_ids"]; attr != nil {
			templateIds := attr.(*schema.Set).List()
			if len(templateIds) > 0 {
				tmplIds := []string{}
				for _, templateId := range templateIds {
					tmplIds = append(tmplIds, templateId.(string))
				}
				permission.Templates.TemplateIds = tmplIds
			} else {
				permission.Templates.FilterType = graphql.FilterTypes.All
			}
		}

		input.AppPermissions = append(input.AppPermissions, permission)
	}
}

func expandWorkflowAppPermissions(d []interface{}, input *graphql.UserGroupPermissions) {
	if len(d) == 0 {
		return
	}

	for _, v := range d {
		permission := &graphql.AppPermission{
			PermissionType: graphql.AppPermissionTypes.Workflow,
			Applications:   &graphql.AppFilter{},
			Actions:        []graphql.Action{},
			Workflows:      &graphql.WorkflowPermissionFilter{},
		}

		permissionConfig := v.(map[string]interface{})
		if attr := permissionConfig["app_ids"]; attr != nil {
			for _, appId := range attr.(*schema.Set).List() {
				permission.Applications.AppIds = append(permission.Applications.AppIds, appId.(string))
			}
		}

		if len(permission.Applications.AppIds) == 0 {
			permission.Applications.FilterType = graphql.FilterTypes.All
		}

		if attr := permissionConfig["actions"]; attr != nil {
			for _, action := range attr.(*schema.Set).List() {
				permission.Actions = append(permission.Actions, graphql.Action(action.(string)))
			}
		}

		if attr := permissionConfig["filters"]; attr != nil {
			filters := attr.(*schema.Set).List()
			if len(filters) > 0 {
				filterTypes := []graphql.WorkflowPermissionFilterType{}
				for _, filter := range filters {
					filterTypes = append(filterTypes, graphql.WorkflowPermissionFilterType(filter.(string)))
				}
				permission.Workflows.FilterTypes = filterTypes
			}
		}

		input.AppPermissions = append(input.AppPermissions, permission)
	}
}

func expandServiceAppPermissions(d []interface{}, input *graphql.UserGroupPermissions) {
	if len(d) == 0 {
		return
	}

	for _, v := range d {
		permission := &graphql.AppPermission{
			PermissionType: graphql.AppPermissionTypes.Service,
			Applications:   &graphql.AppFilter{},
			Actions:        []graphql.Action{},
			Services:       &graphql.ServicePermissionFilter{},
		}

		permissionConfig := v.(map[string]interface{})
		if attr := permissionConfig["app_ids"]; attr != nil {
			for _, appId := range attr.(*schema.Set).List() {
				permission.Applications.AppIds = append(permission.Applications.AppIds, appId.(string))
			}
		}

		if len(permission.Applications.AppIds) == 0 {
			permission.Applications.FilterType = graphql.FilterTypes.All
		}

		if attr := permissionConfig["actions"]; attr != nil {
			for _, action := range attr.(*schema.Set).List() {
				permission.Actions = append(permission.Actions, graphql.Action(action.(string)))
			}
		}

		if attr := permissionConfig["service_ids"]; attr != nil {
			serviceIds := attr.(*schema.Set).List()
			if len(serviceIds) > 0 {
				svcIds := []string{}
				for _, servieId := range serviceIds {
					svcIds = append(svcIds, servieId.(string))
				}
				permission.Services.ServiceIds = svcIds
			} else {
				permission.Services.FilterType = graphql.FilterTypes.All
			}
		}

		input.AppPermissions = append(input.AppPermissions, permission)
	}
}

func expandProvisionerAppPermissions(d []interface{}, input *graphql.UserGroupPermissions) {
	if len(d) == 0 {
		return
	}

	for _, v := range d {
		permission := &graphql.AppPermission{
			PermissionType: graphql.AppPermissionTypes.Provisioner,
			Applications:   &graphql.AppFilter{},
			Actions:        []graphql.Action{},
			Provisioners:   &graphql.ProvisionerPermissionFilter{},
		}

		permissionConfig := v.(map[string]interface{})
		if attr := permissionConfig["app_ids"]; attr != nil {
			for _, appId := range attr.(*schema.Set).List() {
				permission.Applications.AppIds = append(permission.Applications.AppIds, appId.(string))
			}
		}

		if len(permission.Applications.AppIds) == 0 {
			permission.Applications.FilterType = graphql.FilterTypes.All
		}

		if attr := permissionConfig["actions"]; attr != nil {
			for _, action := range attr.(*schema.Set).List() {
				permission.Actions = append(permission.Actions, graphql.Action(action.(string)))
			}
		}

		if attr := permissionConfig["provisioner_ids"]; attr != nil {
			provisionerIds := attr.(*schema.Set).List()
			if len(provisionerIds) > 0 {
				provIds := []string{}
				for _, provisionerId := range provisionerIds {
					provIds = append(provIds, provisionerId.(string))
				}
				permission.Provisioners.ProvisionerIds = provIds
			} else {
				permission.Provisioners.FilterType = graphql.FilterTypes.All
			}
		}

		input.AppPermissions = append(input.AppPermissions, permission)
	}
}

func expandPipelineAppPermissions(d []interface{}, input *graphql.UserGroupPermissions) {
	if len(d) == 0 {
		return
	}

	for _, v := range d {
		permission := &graphql.AppPermission{
			PermissionType: graphql.AppPermissionTypes.Pipeline,
			Applications:   &graphql.AppFilter{},
			Actions:        []graphql.Action{},
			Pipelines:      &graphql.PipelinePermissionFilter{},
		}

		permissionConfig := v.(map[string]interface{})
		if attr := permissionConfig["app_ids"]; attr != nil {
			for _, appId := range attr.(*schema.Set).List() {
				permission.Applications.AppIds = append(permission.Applications.AppIds, appId.(string))
			}
		}

		if len(permission.Applications.AppIds) == 0 {
			permission.Applications.FilterType = graphql.FilterTypes.All
		}

		if attr := permissionConfig["actions"]; attr != nil {
			for _, action := range attr.(*schema.Set).List() {
				permission.Actions = append(permission.Actions, graphql.Action(action.(string)))
			}
		}

		if attr := permissionConfig["env_ids"]; attr != nil {
			environmentIds := attr.(*schema.Set).List()
			if len(environmentIds) > 0 {
				envIds := []string{}
				for _, environmentId := range environmentIds {
					envIds = append(envIds, environmentId.(string))
				}
				permission.Environments.EnvIds = envIds
			}
		}

		if attr := permissionConfig["filters"]; attr != nil {
			filters := attr.(*schema.Set).List()
			if len(filters) > 0 {
				filterTypes := []graphql.PipelinePermissionFilterType{}
				for _, filter := range filters {
					filterTypes = append(filterTypes, graphql.PipelinePermissionFilterType(filter.(string)))
				}
				permission.Pipelines.FilterTypes = filterTypes
			}
		}

		input.AppPermissions = append(input.AppPermissions, permission)
	}
}

func expandEnvironmentAppPermissions(d []interface{}, input *graphql.UserGroupPermissions) {
	if len(d) == 0 {
		return
	}

	for _, v := range d {
		permission := &graphql.AppPermission{
			PermissionType: graphql.AppPermissionTypes.Env,
			Applications:   &graphql.AppFilter{},
			Actions:        []graphql.Action{},
			Environments:   &graphql.EnvPermissionFilter{},
		}

		permissionConfig := v.(map[string]interface{})
		if attr := permissionConfig["app_ids"]; attr != nil {
			for _, appId := range attr.(*schema.Set).List() {
				permission.Applications.AppIds = append(permission.Applications.AppIds, appId.(string))
			}
		}

		if len(permission.Applications.AppIds) == 0 {
			permission.Applications.FilterType = graphql.FilterTypes.All
		}

		if attr := permissionConfig["actions"]; attr != nil {
			for _, action := range attr.(*schema.Set).List() {
				permission.Actions = append(permission.Actions, graphql.Action(action.(string)))
			}
		}

		if attr := permissionConfig["env_ids"]; attr != nil {
			environmentIds := attr.(*schema.Set).List()
			if len(environmentIds) > 0 {
				envIds := []string{}
				for _, environmentId := range environmentIds {
					envIds = append(envIds, environmentId.(string))
				}
				permission.Environments.EnvIds = envIds
			}
		}

		if attr := permissionConfig["filters"]; attr != nil {
			filters := attr.(*schema.Set).List()
			if len(filters) > 0 {
				filterTypes := []graphql.EnvFilterType{}
				for _, filter := range filters {
					filterTypes = append(filterTypes, graphql.EnvFilterType(filter.(string)))
				}
				permission.Environments.FilterTypes = filterTypes
			}
		}

		input.AppPermissions = append(input.AppPermissions, permission)
	}
}

func expandDeploymentAppPermissions(d []interface{}, input *graphql.UserGroupPermissions) {
	if len(d) == 0 {
		return
	}

	for _, v := range d {
		permission := &graphql.AppPermission{
			PermissionType: graphql.AppPermissionTypes.Deployment,
			Applications:   &graphql.AppFilter{},
			Actions:        []graphql.Action{},
			Deployments:    &graphql.DeploymentPermissionFilter{},
		}

		permissionConfig := v.(map[string]interface{})
		if attr := permissionConfig["app_ids"]; attr != nil {
			for _, appId := range attr.(*schema.Set).List() {
				permission.Applications.AppIds = append(permission.Applications.AppIds, appId.(string))
			}
		}

		if len(permission.Applications.AppIds) == 0 {
			permission.Applications.FilterType = graphql.FilterTypes.All
		}

		if attr := permissionConfig["actions"]; attr != nil {
			for _, action := range attr.(*schema.Set).List() {
				permission.Actions = append(permission.Actions, graphql.Action(action.(string)))
			}
		}

		if attr := permissionConfig["env_ids"]; attr != nil {
			environmentIds := attr.(*schema.Set).List()
			if len(environmentIds) > 0 {
				envIds := []string{}
				for _, environmentId := range environmentIds {
					envIds = append(envIds, environmentId.(string))
				}
				permission.Deployments.EnvIds = envIds
			}
		}

		if attr := permissionConfig["filters"]; attr != nil {
			filters := attr.(*schema.Set).List()
			if len(filters) > 0 {
				filterTypes := []graphql.DeploymentPermissionFilterType{}
				for _, filter := range filters {
					filterTypes = append(filterTypes, graphql.DeploymentPermissionFilterType(filter.(string)))
				}
				permission.Deployments.FilterTypes = filterTypes
			}
		}

		input.AppPermissions = append(input.AppPermissions, permission)
	}
}

func expandAllAppPermissions(d []interface{}, input *graphql.UserGroupPermissions) {
	if len(d) == 0 {
		return
	}

	for _, v := range d {
		permission := &graphql.AppPermission{
			PermissionType: graphql.AppPermissionTypes.All,
			Applications:   &graphql.AppFilter{},
			Actions:        []graphql.Action{},
		}

		permissionConfig := v.(map[string]interface{})
		if attr := permissionConfig["app_ids"]; attr != nil {
			for _, appId := range attr.(*schema.Set).List() {
				permission.Applications.AppIds = append(permission.Applications.AppIds, appId.(string))
			}
		}

		if len(permission.Applications.AppIds) == 0 {
			permission.Applications.FilterType = graphql.FilterTypes.All
		}

		if attr := permissionConfig["actions"]; attr != nil {
			for _, action := range attr.(*schema.Set).List() {
				permission.Actions = append(permission.Actions, graphql.Action(action.(string)))
			}
		}

		input.AppPermissions = append(input.AppPermissions, permission)
	}
}

func expandAppPermissions(d []interface{}, input *graphql.UserGroupPermissions) {
	if len(d) == 0 {
		return
	}

	input.AppPermissions = []*graphql.AppPermission{}

	appPermissionsConfig := d[0].(map[string]interface{})

	expandAllAppPermissions(appPermissionsConfig["all"].(*schema.Set).List(), input)
	expandDeploymentAppPermissions(appPermissionsConfig["deployment"].(*schema.Set).List(), input)
	expandEnvironmentAppPermissions(appPermissionsConfig["environment"].(*schema.Set).List(), input)
	expandPipelineAppPermissions(appPermissionsConfig["pipeline"].(*schema.Set).List(), input)
	expandProvisionerAppPermissions(appPermissionsConfig["provisioner"].(*schema.Set).List(), input)
	expandServiceAppPermissions(appPermissionsConfig["service"].(*schema.Set).List(), input)
	expandTemplateAppPermissions(appPermissionsConfig["template"].(*schema.Set).List(), input)
	expandWorkflowAppPermissions(appPermissionsConfig["workflow"].(*schema.Set).List(), input)
}

func expandPermissions(d *schema.ResourceData, input *graphql.UserGroup) {

	config := d.Get("permissions").([]interface{})
	if len(config) == 0 {
		return
	}

	permissionConfig := config[0].(map[string]interface{})
	input.Permissions = &graphql.UserGroupPermissions{}

	expandAccountPermissions(permissionConfig["account_permissions"].(*schema.Set).List(), input.Permissions)
	expandAppPermissions(permissionConfig["app_permissions"].([]interface{}), input.Permissions)

}

func expandLDAPSettings(d *schema.ResourceData, input *graphql.UserGroup) {

	config, ok := d.GetOk("ldap_settings")
	if !ok {
		return
	}

	ssoConfig := config.([]interface{})[0].(map[string]interface{})
	input.LDAPSettings = &graphql.LDAPSettings{}

	if attr := ssoConfig["group_dn"]; attr != "" {
		input.LDAPSettings.GroupDN = attr.(string)
	}

	if attr := ssoConfig["group_name"]; attr != "" {
		input.LDAPSettings.GroupName = attr.(string)
	}

	if attr := ssoConfig["sso_provider_id"]; attr != "" {
		input.LDAPSettings.SSOProviderId = attr.(string)
	}

}

func expandNotificationSettings(d *schema.ResourceData, input *graphql.UserGroup) {
	config := d.Get("notification_settings").([]interface{})
	if len(config) == 0 {
		return
	}

	notificationSettings := config[0].(map[string]interface{})
	input.NotificationSettings = &graphql.NotificationSettings{
		SlackNotificationSetting: &graphql.SlackNotificationSetting{},
	}

	if attr := notificationSettings["group_email_addresses"]; attr != nil {
		input.NotificationSettings.GroupEmailAddresses = utils.InterfaceSliceToStringSlice(attr.([]interface{}))
	}

	if attr := notificationSettings["microsoft_teams_webhook_url"]; attr != "" {
		input.NotificationSettings.MicrosoftTeamsWebhookUrl = attr.(string)
	}

	if attr := notificationSettings["send_mail_to_new_members"]; attr != "" {
		input.NotificationSettings.SendMailToNewMembers = attr.(bool)
	}

	if attr := notificationSettings["send_notifications_to_members"]; attr != nil {
		input.NotificationSettings.SendNotificationToMembers = attr.(bool)
	}

	if attr := notificationSettings["slack_webhook_url"]; attr != "" {
		input.NotificationSettings.SlackNotificationSetting.SlackWebhookUrl = attr.(string)
	}

	if attr := notificationSettings["slack_channel"]; attr != "" {
		input.NotificationSettings.SlackNotificationSetting.SlackChannelName = attr.(string)
	}

}

func flattenAccountPermissions(permissions *graphql.UserGroupPermissions) []string {
	results := []string{}

	if permissions.AccountPermissions == nil {
		return results
	}

	for _, accountPermissionType := range permissions.AccountPermissions.AccountPermissionTypes {
		results = append(results, accountPermissionType.String())
	}

	return results
}

func flattenAppPermissionAll(appPermissions *graphql.AppPermission) map[string]interface{} {
	results := map[string]interface{}{}

	if len(appPermissions.Applications.AppIds) > 0 {
		results["app_ids"] = appPermissions.Applications.AppIds
	}

	if len(appPermissions.Actions) > 0 {
		results["actions"] = appPermissions.Actions
	}

	return results
}

func flattenAppPermissionDeployment(appPermissions *graphql.AppPermission) map[string]interface{} {
	results := map[string]interface{}{}

	if len(appPermissions.Applications.AppIds) > 0 {
		results["app_ids"] = appPermissions.Applications.AppIds
	}

	if len(appPermissions.Actions) > 0 {
		results["actions"] = appPermissions.Actions
	}

	if appPermissions.Deployments != nil {
		if len(appPermissions.Deployments.EnvIds) > 0 {
			results["env_ids"] = appPermissions.Deployments.EnvIds
		}

		if len(appPermissions.Deployments.FilterTypes) > 0 {
			filters := []string{}
			for _, filterType := range appPermissions.Deployments.FilterTypes {
				filters = append(filters, filterType.String())
			}
			results["filters"] = filters
		}
	}

	return results
}

func flattenAppPermissionEnv(appPermissions *graphql.AppPermission) map[string]interface{} {
	results := map[string]interface{}{}

	if len(appPermissions.Applications.AppIds) > 0 {
		results["app_ids"] = appPermissions.Applications.AppIds
	}

	if len(appPermissions.Actions) > 0 {
		results["actions"] = appPermissions.Actions
	}

	if appPermissions.Environments != nil {
		if len(appPermissions.Environments.EnvIds) > 0 {
			results["env_ids"] = appPermissions.Environments.EnvIds
		}

		if len(appPermissions.Environments.FilterTypes) > 0 {
			filters := []string{}
			for _, filterType := range appPermissions.Environments.FilterTypes {
				filters = append(filters, filterType.String())
			}
			results["filters"] = filters
		}
	}

	return results
}

func flattenAppPermissionPipeline(appPermissions *graphql.AppPermission) map[string]interface{} {
	results := map[string]interface{}{}

	if len(appPermissions.Applications.AppIds) > 0 {
		results["app_ids"] = appPermissions.Applications.AppIds
	}

	if len(appPermissions.Actions) > 0 {
		results["actions"] = appPermissions.Actions
	}

	if appPermissions.Pipelines != nil {
		if len(appPermissions.Pipelines.EnvIds) > 0 {
			results["env_ids"] = appPermissions.Pipelines.EnvIds
		}

		if len(appPermissions.Pipelines.FilterTypes) > 0 {
			filters := []string{}
			for _, filterType := range appPermissions.Pipelines.FilterTypes {
				filters = append(filters, filterType.String())
			}
			results["filters"] = filters
		}
	}

	return results
}

func flattenAppPermissionProvisioner(appPermissions *graphql.AppPermission) map[string]interface{} {
	results := map[string]interface{}{}

	if len(appPermissions.Applications.AppIds) > 0 {
		results["app_ids"] = appPermissions.Applications.AppIds
	}

	if len(appPermissions.Actions) > 0 {
		results["actions"] = appPermissions.Actions
	}

	if appPermissions.Provisioners != nil {
		if len(appPermissions.Provisioners.ProvisionerIds) > 0 {
			results["provisioner_ids"] = appPermissions.Provisioners.ProvisionerIds
		}
	}

	return results
}

func flattenAppPermissionService(appPermissions *graphql.AppPermission) map[string]interface{} {
	results := map[string]interface{}{}

	if len(appPermissions.Applications.AppIds) > 0 {
		results["app_ids"] = appPermissions.Applications.AppIds
	}

	if len(appPermissions.Actions) > 0 {
		results["actions"] = appPermissions.Actions
	}

	if appPermissions.Services != nil {
		if len(appPermissions.Services.ServiceIds) > 0 {
			results["service_ids"] = appPermissions.Services.ServiceIds
		}
	}

	return results
}

func flattenAppPermissionTemplate(appPermissions *graphql.AppPermission) map[string]interface{} {
	results := map[string]interface{}{}

	if len(appPermissions.Applications.AppIds) > 0 {
		results["app_ids"] = appPermissions.Applications.AppIds
	}

	if len(appPermissions.Actions) > 0 {
		results["actions"] = appPermissions.Actions
	}

	if appPermissions.Templates != nil {
		if len(appPermissions.Templates.TemplateIds) > 0 {
			results["template_ids"] = appPermissions.Templates.TemplateIds
		}
	}

	return results
}

func flattenAppPermissionWorkflow(appPermissions *graphql.AppPermission) map[string]interface{} {
	results := map[string]interface{}{}

	if len(appPermissions.Applications.AppIds) > 0 {
		results["app_ids"] = appPermissions.Applications.AppIds
	}

	if len(appPermissions.Actions) > 0 {
		results["actions"] = appPermissions.Actions
	}

	if appPermissions.Workflows != nil {
		if len(appPermissions.Workflows.FilterTypes) > 0 {
			filters := []string{}
			for _, filterType := range appPermissions.Workflows.FilterTypes {
				filters = append(filters, filterType.String())
			}
			results["filters"] = filters
		}
	}

	return results
}

func flattenAppPermissions(permissions *graphql.UserGroupPermissions) []interface{} {
	results := []interface{}{}

	if len(permissions.AppPermissions) == 0 {
		return results
	}

	all := []interface{}{}
	deployments := []interface{}{}
	envs := []interface{}{}
	pipelines := []interface{}{}
	provisioners := []interface{}{}
	services := []interface{}{}
	templates := []interface{}{}
	workflows := []interface{}{}

	for _, appPermission := range permissions.AppPermissions {
		switch appPermission.PermissionType {
		case graphql.AppPermissionTypes.All:
			all = append(all, flattenAppPermissionAll(appPermission))
		case graphql.AppPermissionTypes.Deployment:
			deployments = append(deployments, flattenAppPermissionDeployment(appPermission))
		case graphql.AppPermissionTypes.Env:
			envs = append(envs, flattenAppPermissionEnv(appPermission))
		case graphql.AppPermissionTypes.Pipeline:
			pipelines = append(pipelines, flattenAppPermissionPipeline(appPermission))
		case graphql.AppPermissionTypes.Provisioner:
			provisioners = append(provisioners, flattenAppPermissionProvisioner(appPermission))
		case graphql.AppPermissionTypes.Service:
			services = append(services, flattenAppPermissionService(appPermission))
		case graphql.AppPermissionTypes.Template:
			templates = append(templates, flattenAppPermissionTemplate(appPermission))
		case graphql.AppPermissionTypes.Workflow:
			workflows = append(workflows, flattenAppPermissionWorkflow(appPermission))
		default:
			panic(fmt.Sprintf("Unhandled app permission type: %s", appPermission.PermissionType))
		}
	}

	return append(results, map[string]interface{}{
		"all":         all,
		"deployment":  deployments,
		"environment": envs,
		"pipeline":    pipelines,
		"provisioner": provisioners,
		"service":     services,
		"template":    templates,
		"workflow":    workflows,
	})
}

func flattenPermissions(d *schema.ResourceData, ugPermissions *graphql.UserGroupPermissions) []map[string]interface{} {
	results := []map[string]interface{}{}

	if ugPermissions == nil || ugPermissions.IsEmpty() {
		return results
	}

	permissions := map[string]interface{}{}

	if accountPermission := flattenAccountPermissions(ugPermissions); len(accountPermission) > 0 {
		permissions["account_permissions"] = accountPermission
	}

	if appPermissions := flattenAppPermissions(ugPermissions); len(appPermissions) > 0 {
		permissions["app_permissions"] = appPermissions
	}

	return append(results, permissions)
}

func flattenNotificationSettings(d *schema.ResourceData, settings *graphql.NotificationSettings) []map[string]interface{} {
	results := []map[string]interface{}{}

	if settings == nil || settings.IsEmpty() {
		return results
	}

	s := map[string]interface{}{
		"group_email_addresses":         settings.GroupEmailAddresses,
		"microsoft_teams_webhook_url":   settings.MicrosoftTeamsWebhookUrl,
		"send_mail_to_new_members":      settings.SendMailToNewMembers,
		"send_notifications_to_members": settings.SendNotificationToMembers,
	}

	if settings.SlackNotificationSetting != nil && !settings.SlackNotificationSetting.IsEmpty() {
		s["slack_channel"] = settings.SlackNotificationSetting.SlackChannelName
		s["slack_webhook_url"] = settings.SlackNotificationSetting.SlackWebhookUrl
	}

	return append(results, s)
}

func expandSAMLSettings(d *schema.ResourceData, input *graphql.UserGroup) {

	config, ok := d.GetOk("saml_settings")
	if !ok {
		return
	}

	ssoConfig := config.([]interface{})[0].(map[string]interface{})
	input.SAMLSettings = &graphql.SAMLSettings{}

	if attr := ssoConfig["group_name"]; attr != "" {
		input.SAMLSettings.GroupName = attr.(string)
	}

	if attr := ssoConfig["sso_provider_id"]; attr != "" {
		input.SAMLSettings.SSOProviderId = attr.(string)
	}

}

func flattenSAMLSettings(d *schema.ResourceData, settings *graphql.SAMLSettings) []map[string]interface{} {
	results := []map[string]interface{}{}

	if settings == nil || settings.IsEmpty() {
		return results
	}

	s := map[string]interface{}{
		"group_name":      settings.GroupName,
		"sso_provider_id": settings.SSOProviderId,
	}

	return append(results, s)
}

func flattenLDAPSettings(d *schema.ResourceData, settings *graphql.LDAPSettings) []map[string]interface{} {
	results := []map[string]interface{}{}

	if settings == nil || settings.IsEmpty() {
		return results
	}

	s := map[string]interface{}{
		"group_dn":        settings.GroupDN,
		"group_name":      settings.GroupName,
		"sso_provider_id": settings.SSOProviderId,
	}

	return append(results, s)
}
