package user

import (
	"context"
	"strings"

	"github.com/harness/harness-go-sdk/harness/cd"
	"github.com/harness/harness-go-sdk/harness/cd/graphql"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceUserGroup() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Data source for retrieving a Harness user group",

		ReadContext: dataSourceUserGroupRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Description:   "Unique identifier of the user group",
				Type:          schema.TypeString,
				Optional:      true,
				AtLeastOneOf:  []string{"id", "name"},
				ConflictsWith: []string{"name"},
			},
			"name": {
				Description:   "The name of the user group.",
				Type:          schema.TypeString,
				Optional:      true,
				AtLeastOneOf:  []string{"id", "name"},
				ConflictsWith: []string{"id"},
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return strings.EqualFold(old, new)
				},
			},
		},
	}
}

func dataSourceUserGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	c := meta.(*cd.ApiClient)

	var userGroup *graphql.UserGroup
	var err error

	if id := d.Get("id").(string); id != "" {
		// Try lookup by Id first
		userGroup, err = c.UserClient.GetUserGroupById(id)
		if err != nil {
			return diag.FromErr(err)
		}
	} else if name := d.Get("name").(string); name != "" {
		// Fallback to lookup by name
		userGroup, err = c.UserClient.GetUserGroupByName(name)
		if err != nil {
			return diag.FromErr(err)
		}
	} else {
		// Throw error if neither are set
		return diag.Errorf("id or name must be set")
	}

	if userGroup == nil {
		return diag.Errorf("user group not found")
	}

	d.SetId(userGroup.Id)
	d.Set("name", userGroup.Name)

	return nil
}
