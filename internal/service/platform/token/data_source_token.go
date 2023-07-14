package token

import (
	"context"
	"errors"
	"net/http"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func DataSourceToken() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving a Harness ApiKey Token.",

		ReadContext: dataSourceTokenRead,

		Schema: map[string]*schema.Schema{
			"apikey_id": {
				Description: "Identifier of the API Key",
				Type:        schema.TypeString,
				Required:    true,
			},
			"apikey_type": {
				Description:  "Type of the API Key",
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"USER", "SERVICE_ACCOUNT"}, false),
			},
			"parent_id": {
				Description: "Parent Entity Identifier of the API Key",
				Type:        schema.TypeString,
				Required:    true,
			},
			"account_id": {
				Description: "Account Identifier for the Entity",
				Type:        schema.TypeString,
				Required:    true,
			},
			"valid_from": {
				Description: "This is the time from which the Token is valid. The time is in milliseconds",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"valid_to": {
				Description: "This is the time till which the Token is valid. The time is in milliseconds",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"scheduled_expire_time": {
				Description: "Scheduled expiry time in milliseconds",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"valid": {
				Description: "Boolean value to indicate if Token is valid or not.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"email": {
				Description: "Email Id of the user who created the Token",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"username": {
				Description: "Name of the user who created the Token",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"encoded_password": {
				Description: "Encoded password of the Token",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
	helpers.SetMultiLevelDatasourceSchemaIdentifierRequired(resource.Schema)

	return resource
}

func dataSourceTokenRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	var token *nextgen.Token

	id := d.Get("identifier").(string)
	apikey_type := d.Get("apikey_type").(string)
	parentId := d.Get("parent_id").(string)
	apikey_id := d.Get("apikey_id").(string)

	if id != "" {
		var err error
		var httpResp *http.Response
		resp, httpResp, err := c.TokenApi.ListAggregatedTokens(ctx, c.AccountId, apikey_type, parentId, apikey_id, &nextgen.TokenApiListAggregatedTokensOpts{
			OrgIdentifier:     helpers.BuildField(d, "org_id"),
			ProjectIdentifier: helpers.BuildField(d, "project_id"),
			Identifiers:       optional.NewInterface(id),
		})
		tokenList := resp.Data.Content
		if err != nil {
			return helpers.HandleApiError(err, d, httpResp)
		}

		if tokenList == nil || len(tokenList) == 0 {
			d.SetId("")
			d.MarkNewResource()
			return nil
		} else {
			token = tokenList[0].Token
		}

	} else {
		return diag.FromErr(errors.New("Identifier must be specified"))
	}

	readToken(d, token)

	return nil
}
