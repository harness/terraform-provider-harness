package organization

import (
	"context"
	"errors"
	"net/http"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceOrganization() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving a Harness organization",

		ReadContext: dataSourceOrganizationRead,

		Schema: map[string]*schema.Schema{},
	}

	helpers.SetCommonDataSourceSchemaIdentifierRequired(resource.Schema)

	return resource
}

func dataSourceOrganizationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	var err error
	var org *nextgen.OrganizationResponse
	var httpResp *http.Response

	id := d.Get("identifier").(string)
	name := d.Get("name").(string)

	if id != "" {
		var resp nextgen.ResponseDtoOrganizationResponse
		resp, httpResp, err = c.OrganizationApi.GetOrganization(ctx, id, c.AccountId)
		org = resp.Data
	} else if name != "" {
		org, httpResp, err = c.OrganizationApi.GetOrganizationByName(ctx, c.AccountId, name)
	} else {
		return diag.FromErr(errors.New("either identifier or name must be specified"))
	}

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	if org == nil || org.Organization == nil {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	readOrganization(d, org.Organization)

	return nil
}
