package apikey

import (
	"context"
	"errors"
	"net/http"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func DataSourceApiKey() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving a Harness ApiKey.",

		ReadContext: dataSourceApiKeyRead,

		Schema: map[string]*schema.Schema{
			"apikey_type": {
				Description:  "Type of the API Key",
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"USER", "SERVICE_ACCOUNT"}, false),
			},
			"parent_id": {
				Description: "Parent Identifier for the Entity",
				Type:        schema.TypeString,
				Required:    true,
			},
			"default_time_to_expire_token": {
				Description: "Expiry time of the apiKey",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"account_id": {
				Description: "Account Identifier for the Entity",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
	helpers.SetMultiLevelDatasourceSchemaIdentifierRequired(resource.Schema)

	return resource
}

func dataSourceApiKeyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	var apiKey *nextgen.ApiKey

	id := d.Get("identifier").(string)
	type_ := d.Get("apikey_type").(string)
	parentId := d.Get("parent_id").(string)

	if id != "" {
		var err error
		var httpResp *http.Response
		resp, httpResp, err := c.ApiKeyApi.GetAggregatedApiKey(ctx, c.AccountId, type_, parentId, id, &nextgen.ApiKeyApiGetAggregatedApiKeyOpts{
			OrgIdentifier:     helpers.BuildField(d, "org_id"),
			ProjectIdentifier: helpers.BuildField(d, "project_id"),
		})
		if err != nil {
			return helpers.HandleApiError(err, d, httpResp)
		}
		apiKey = resp.Data.ApiKey

		if apiKey == nil {
			d.SetId("")
			d.MarkNewResource()
			return nil
		}
	} else {
		return diag.FromErr(errors.New("Identifier must be specified"))
	}

	readApiKey(d, apiKey)

	return nil
}
