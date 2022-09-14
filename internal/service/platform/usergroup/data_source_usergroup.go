package usergroup

import (
	"context"
	"errors"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceUserGroup() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving a Harness User Group.",

		ReadContext: dataSourceUserGroupRead,

		Schema: map[string]*schema.Schema{
			"linked_sso_id": {
				Description: "The SSO account ID that the user group is linked to.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"externally_managed": {
				Description: "Whether the user group is externally managed.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"users": {
				Description: "List of users in the UserGroup.",
				Type:        schema.TypeSet,
				Computed:    true,
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
							Computed:    true,
						},
						"slack_webhook_url": {
							Description: "Url of slack webhook",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"microsoft_teams_webhook_url": {
							Description: "Url of Microsoft teams webhook",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"pager_duty_key": {
							Description: "Pager duty key",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"group_email": {
							Description: "Group email",
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"linked_sso_display_name": {
				Description: "Name of the linked SSO.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"sso_group_id": {
				Description: "Identifier of the userGroup in SSO.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"sso_group_name": {
				Description: "Name of the SSO userGroup.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"linked_sso_type": {
				Description: "Type of linked SSO",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"sso_linked": {
				Description: "Whether sso is linked or not",
				Type:        schema.TypeBool,
				Computed:    true,
			},
		},
	}

	helpers.SetMultiLevelDatasourceSchema(resource.Schema)

	return resource
}

func dataSourceUserGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	var ug *nextgen.UserGroup
	var err error

	id := d.Get("identifier").(string)
	name := d.Get("name").(string)

	if id != "" {
		var resp nextgen.ResponseDtoUserGroup
		resp, _, err = c.UserGroupApi.GetUserGroup(ctx, c.AccountId, id, &nextgen.UserGroupApiGetUserGroupOpts{
			OrgIdentifier:     optional.NewString(d.Get("org_id").(string)),
			ProjectIdentifier: optional.NewString(d.Get("project_id").(string)),
		})
		ug = resp.Data
	} else if name != "" {
		ug, err = c.UserGroupApi.GetUserGroupByName(ctx, c.AccountId, name, &nextgen.UserGroupApiGetUserGroupByNameOpts{
			OrgIdentifier:     optional.NewString(d.Get("org_id").(string)),
			ProjectIdentifier: optional.NewString(d.Get("project_id").(string)),
		})
	} else {
		return diag.FromErr(errors.New("either identifier or name must be specified"))
	}

	if err != nil {
		return helpers.HandleApiError(err, d)
	}

	if ug == nil {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	readUserGroup(d, ug)

	return nil
}
