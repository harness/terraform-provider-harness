package delegatetoken

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

func DataSourceDelegateToken() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving a Harness delegate Token.",

		ReadContext: dataSourceDelegateTokenRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Description: "Name of the delegate token",
				Type:        schema.TypeString,
				Required:    true,
			},
			"account_id": {
				Description: "Account Identifier for the Entity",
				Type:        schema.TypeString,
				Required:    true,
			},
			"org_id": {
				Description: "Org Identifier for the Entity",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"project_id": {
				Description:  "Project Identifier for the Entity",
				Type:         schema.TypeString,
				Optional:     true,
				RequiredWith: []string{"org_id"},
			},
			"token_status": {
				Description:  "Status of Delegate Token (ACTIVE or REVOKED). If left empty both active and revoked tokens will be assumed",
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"ACTIVE", "REVOKED"}, false),
			},
			"value": {
				Description: "Value of the delegate token. Encoded in base64.",
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
			},
			"created_at": {
				Description: "Time when the delegate token is created. This is an epoch timestamp.",
				Type:        schema.TypeInt,
				Computed:    true,
			},
			"revoke_after": {
				Description: "Epoch time in milliseconds after which the token will be marked as revoked. There can be a delay of up to one hour from the epoch value provided and actual revoking of the token.",
				Type:        schema.TypeInt,
				Computed:    true,
			},
			"created_by": {
				Description: "created by details",
				Type:        schema.TypeMap,
				Computed:    true,
			},
		},
	}

	return resource
}

func dataSourceDelegateTokenRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	var delegateToken *nextgen.DelegateTokenDetails

	name := d.Get("name").(string)

	if name != "" {
		var err error
		var httpResp *http.Response
		resp, httpResp, err := c.DelegateTokenResourceApi.GetCgDelegateTokens(ctx, c.AccountId, &nextgen.DelegateTokenResourceApiGetCgDelegateTokensOpts{
			OrgIdentifier:     helpers.BuildField(d, "org_id"),
			ProjectIdentifier: helpers.BuildField(d, "project_id"),
			Name:              helpers.BuildField(d, "name"),
			Status:            helpers.BuildField(d, "token_status"),
		})
		if err != nil {
			return helpers.HandleApiError(err, d, httpResp)
		}
		if resp.Resource != nil && len(resp.Resource) > 0 {
			delegateToken = &resp.Resource[0]
		}

		if delegateToken == nil {
			d.SetId("")
			d.MarkNewResource()
			return nil
		}
	} else {
		return diag.FromErr(errors.New("Name must be specified"))
	}

	readDelegateToken(d, delegateToken)

	return nil
}
