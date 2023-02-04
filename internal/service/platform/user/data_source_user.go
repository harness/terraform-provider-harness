package user

import (
	"context"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/harness/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceUser() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for retrieving the user based on the API key.",

		ReadContext: dataSourceUserRead,

		Schema: map[string]*schema.Schema{
			"identifier": {
				Description: "Unique identifier of the user.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"org_id": {
				Description: "Organization identifier of the user.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"project_id": {
				Description: "Project identifier of the user.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"name": {
				Description: "Name of the user.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"emails": {
				Description: "The email of the user.",
				Type:        schema.TypeSet,
				Required:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"email": {
				Description: "The email of the user.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"disabled": {
				Description: "Whether or not the user account is disabled.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"locked": {
				Description: "Whether or not the user account is locked.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"externally_managed": {
				Description: "Whether or not the user account is externally managed.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
		},
	}
}

func dataSourceUserRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	emails := []string{}
	var email = ""
	if attr, ok := d.GetOk("emails"); ok {
		emails = utils.InterfaceSliceToStringSlice(attr.(*schema.Set).List())
		email = emails[0]
	}

	resp, httpResp, err := c.UserApi.GetAggregatedUsers(ctx, c.AccountId, &nextgen.UserApiGetAggregatedUsersOpts{
		OrgIdentifier:     helpers.BuildField(d, "org_id"),
		ProjectIdentifier: helpers.BuildField(d, "project_id"),
		SearchTerm:        optional.NewString(email),
	})

	if &resp == nil || resp.Data == nil || resp.Data.Empty {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	if err != nil {
		return helpers.HandleReadApiError(err, d, httpResp)
	}

	readUserList(d, resp.Data)

	return nil
}
