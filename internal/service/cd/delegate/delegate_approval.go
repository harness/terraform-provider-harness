package delegate

import (
	"context"
	"fmt"

	sdk "github.com/harness/harness-go-sdk"
	"github.com/harness/harness-go-sdk/harness/cd/graphql"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceDelegateApproval() *schema.Resource {
	return &schema.Resource{
		Description:   "Resource for approving or rejecting delegates.",
		CreateContext: resourceDelegateApprovalCreateOrUpdate,
		ReadContext:   resourceDelegateApprovalRead,
		UpdateContext: resourceDelegateApprovalCreateOrUpdate,
		DeleteContext: resourceDelegateApprovalDelete,
		Schema: map[string]*schema.Schema{
			"delegate_id": {
				Description: "The id of the delegate.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"approve": {
				Description: "Whether or not to approve the delegate.",
				Type:        schema.TypeBool,
				Required:    true,
			},
			"status": {
				Description: "The status of the delegate.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
				d.Set("delegate_id", d.Id())
				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func resourceDelegateApprovalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*sdk.Session)

	id := d.Get("delegate_id").(string)
	delegate, err := c.CDClient.DelegateClient.GetDelegateById(id)

	if err != nil {
		return diag.FromErr(err)
	}

	if delegate == nil {
		return diag.FromErr(fmt.Errorf("delegate %s not found", id))
	}

	d.SetId(delegate.UUID)

	switch delegate.Status {
	case graphql.DelegateStatusTypes.Enabled.String():
		d.Set("approve", true)
	case graphql.DelegateStatusTypes.Deleted.String():
		d.Set("approve", false)
	}

	d.Set("status", delegate.Status)

	return nil
}

func resourceDelegateApprovalCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*sdk.Session)

	// Don't attempt to do anything if we've already done this once.
	if !d.IsNewResource() {
		return diag.Errorf("the delegate approval status has already been changed.")
	}

	id := d.Get("delegate_id").(string)
	delegate, err := c.CDClient.DelegateClient.GetDelegateById(id)
	if err != nil {
		return diag.FromErr(err)
	}

	if delegate == nil {
		return diag.FromErr(fmt.Errorf("delegate %s not found", id))
	}

	if delegate.Status != graphql.DelegateStatusTypes.WaitingForApproval.String() {
		return diag.Errorf("cannot update delegate. Current status is %s", delegate.Status)
	}

	var approvaltype graphql.DelegateApprovalType = graphql.DelegateApprovalTypes.Activate
	if !d.Get("approve").(bool) {
		approvaltype = graphql.DelegateApprovalTypes.Reject
	}

	delegate, err = c.CDClient.DelegateClient.UpdateDelegateApprovalStatus(&graphql.DelegateApprovalRejectInput{
		AccountId:        c.AccountId,
		DelegateApproval: approvaltype,
		DelegateId:       delegate.UUID,
	})

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(delegate.UUID)
	d.Set("status", delegate.Status)

	return nil
}

func resourceDelegateApprovalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Nothing to do
	return nil
}
