package ng

import (
	"context"

	sdk "github.com/harness-io/harness-go-sdk"
	"github.com/harness-io/harness-go-sdk/harness/nextgen"
	"github.com/harness-io/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceOrganization() *schema.Resource {
	return &schema.Resource{
		Description: utils.GetNextgenDescription("Resource for creating a Harness organization."),

		ReadContext:   resourceOrganizationRead,
		UpdateContext: resourceOrganizationUpdate,
		DeleteContext: resourceOrganizationDelete,
		CreateContext: resourceOrganizationCreate,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"identifier": {
				Description: "Unique identifier of the organization.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"name": {
				Description: "Name of the organization.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"description": {
				Description: "Description of the organization.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"tags": {
				Description: "Tags associated with the organization.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceOrganizationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*sdk.Session)

	id := d.Id()
	if id == "" {
		id = d.Get("identifier").(string)
	}

	resp, _, err := c.NGClient.OrganizationApi.GetOrganization(ctx, id, c.AccountId)

	if err != nil {
		e := err.(nextgen.GenericSwaggerError)
		if e.Code() == nextgen.ErrorCodes.ResourceNotFound {
			d.SetId("")
			d.MarkNewResource()
			return nil
		}
		return diag.Errorf(e.Error())
	}

	readOrganization(d, resp.Data.Organization)

	return nil
}

func resourceOrganizationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*sdk.Session)

	org := buildOrganization(d)

	resp, _, err := c.NGClient.OrganizationApi.PostOrganization(ctx, nextgen.OrganizationRequest{Organization: org}, c.AccountId)
	if err != nil {
		return diag.Errorf(err.(nextgen.GenericSwaggerError).Error())
	}

	readOrganization(d, resp.Data.Organization)

	return nil
}

func buildOrganization(d *schema.ResourceData) *nextgen.Organization {
	return &nextgen.Organization{
		Identifier:  d.Get("identifier").(string),
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Tags:        utils.ExpandTags(d.Get("tags").(*schema.Set).List()),
	}
}

func resourceOrganizationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*sdk.Session)

	org := buildOrganization(d)

	resp, _, err := c.NGClient.OrganizationApi.PutOrganization(ctx, nextgen.OrganizationRequest{Organization: org}, c.AccountId, org.Identifier, nil)
	if err != nil {
		return diag.Errorf(err.(nextgen.GenericSwaggerError).Error())
	}

	readOrganization(d, resp.Data.Organization)

	return nil
}

func resourceOrganizationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*sdk.Session)

	_, _, err := c.NGClient.OrganizationApi.DeleteOrganization(ctx, d.Id(), c.AccountId, nil)
	if err != nil {
		return diag.Errorf(err.(nextgen.GenericSwaggerError).Error())
	}

	return nil
}

func readOrganization(d *schema.ResourceData, org *nextgen.Organization) {
	d.SetId(org.Identifier)
	d.Set("identifier", org.Identifier)
	d.Set("name", org.Name)
	d.Set("description", org.Description)
	d.Set("tags", utils.FlattenTags(org.Tags))
}
