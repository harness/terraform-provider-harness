package user

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	sdk "github.com/harness-io/harness-go-sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceAddUserToGroup() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Resource for adding a user to a group.",

		CreateContext: resourceAddUserToGroupCreate,
		ReadContext:   resourceAddUserToGroupRead,
		DeleteContext: resourceAddUserToGroupDelete,

		Schema: map[string]*schema.Schema{
			"user_id": {
				Description: "Unique identifier of the user.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"group_id": {
				Description: "The name of the user.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
		},

		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
				// <user_id>/<group_id>
				parts := strings.Split(d.Id(), "/")
				d.Set("user_id", parts[0])
				d.Set("group_id", parts[1])
				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func resourceAddUserToGroupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*sdk.Session)

	userId := d.Get("user_id").(string)
	groupId := d.Get("group_id").(string)

	log.Printf("[DEBUG] Check if user is already in group")
	ok, err := c.CDClient.UserClient.IsUserInGroup(userId, groupId)
	if err != nil {
		return diag.FromErr(err)
	}
	if !ok {
		log.Printf("[DEBUG] User is not in group, adding user to group")
		ok, err := c.CDClient.UserClient.AddUserToGroup(userId, groupId)
		if err != nil {
			return diag.FromErr(err)
		}
		if !ok {
			return diag.FromErr(errors.New("failed to add user to group"))
		}
	}

	d.SetId(fmt.Sprintf("%s:%s", userId, groupId))

	return nil
}

func resourceAddUserToGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*sdk.Session)

	userId := d.Get("user_id").(string)
	groupId := d.Get("group_id").(string)

	ok, err := c.CDClient.UserClient.IsUserInGroup(userId, groupId)
	if err != nil {
		return diag.FromErr(err)
	}
	if !ok {
		log.Printf("[WARN] Removing from state because user %s is no longer in group %s", userId, groupId)
		d.SetId("")
		return nil
	}

	d.SetId(fmt.Sprintf("%s:%s", userId, groupId))

	return nil
}

func resourceAddUserToGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*sdk.Session)

	userId := d.Get("user_id").(string)
	groupId := d.Get("group_id").(string)

	ok, err := c.CDClient.UserClient.RemoveUserFromGroup(userId, groupId)
	if err != nil {
		return diag.FromErr(err)
	}
	if !ok {
		return diag.FromErr(errors.New("failed to remove user from group"))
	}

	d.SetId("")

	return nil
}
