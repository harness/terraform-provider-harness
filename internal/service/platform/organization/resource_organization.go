package organization

import (
	"context"
	"net/http"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceOrganization() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for creating a Harness organization.",

		ReadContext:   resourceOrganizationRead,
		UpdateContext: resourceOrganizationCreateOrUpdate,
		DeleteContext: resourceOrganizationDelete,
		CreateContext: resourceOrganizationCreateOrUpdate,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{},
	}

	helpers.SetCommonResourceSchema(resource.Schema)

	return resource
}

func resourceOrganizationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Id()
	if id == "" {
		d.MarkNewResource()
		return nil
	}

	resp, httpResp, err := c.OrganizationApi.GetOrganization(ctx, d.Id(), c.AccountId)

	if err != nil {
		return helpers.HandleReadApiError(err, d, httpResp)
	}

	readOrganization(d, resp.Data.Organization)

	return nil
}

func resourceOrganizationCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Id()
	org := buildOrganization(d)

	var err error
	var resp nextgen.ResponseDtoOrganizationResponse
	var httpResp *http.Response

	if id == "" {
		resp, httpResp, err = c.OrganizationApi.PostOrganization(ctx, nextgen.OrganizationRequest{Organization: org}, c.AccountId)
	} else {
		resp, httpResp, err = c.OrganizationApi.PutOrganization(ctx, nextgen.OrganizationRequest{Organization: org}, c.AccountId, org.Identifier, nil)
	}

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	readOrganization(d, resp.Data.Organization)

	return nil
}

func resourceOrganizationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	_, httpResp, err := c.OrganizationApi.DeleteOrganization(ctx, d.Id(), c.AccountId, nil)
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	return nil
}

func buildOrganization(d *schema.ResourceData) *nextgen.Organization {
	return &nextgen.Organization{
		Identifier:  d.Get("identifier").(string),
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Tags:        helpers.ExpandTags(d.Get("tags").(*schema.Set).List()),
	}
}

func readOrganization(d *schema.ResourceData, org *nextgen.Organization) {
	d.SetId(org.Identifier)
	d.Set("identifier", org.Identifier)
	d.Set("name", org.Name)
	d.Set("description", org.Description)
	d.Set("tags", helpers.FlattenTags(org.Tags))
}
