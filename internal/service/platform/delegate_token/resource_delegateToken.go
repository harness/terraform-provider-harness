package delegatetoken

import (
	"context"
	"net/http"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceDelegateToken() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for creating delegate tokens.",

		ReadContext:   resourceDelegateTokenRead,
		CreateContext: resourceDelegateTokenCreateOrUpdate,
		UpdateContext: resourceDelegateTokenCreateOrUpdate,
		DeleteContext: resourceDelegateTokenDelete,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"identifier": {
				Description: "Identifier of the delegate token",
				Type:        schema.TypeString,
				Required:    true,
			},
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
				Optional:    true,
				Computed:    true,
			},
			"created_at": {
				Description: "Time when the delegate token is created. This is an epoch timestamp.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
		},
	}

	return resource
}

func resourceDelegateTokenRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	resp, httpResp, err := c.DelegateTokenResourceApi.GetDelegateTokens(ctx, c.AccountId, &nextgen.DelegateTokenResourceApiGetDelegateTokensOpts{
		OrgIdentifier:     helpers.BuildField(d, "org_id"),
		ProjectIdentifier: helpers.BuildField(d, "project_id"),
		Name:              helpers.BuildField(d, "name"),
		Status:            helpers.BuildField(d, "token_status"),
	})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	if resp.Resource != nil {
		readDelegateToken(d, &resp.Resource[0])
	}

	return nil
}

func resourceDelegateTokenCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	var err error
	var resp nextgen.RestResponseDelegateTokenDetails
	var httpResp *http.Response

	delegateToken := buildDelegateToken(d)

	if delegateToken.Value == "" {
		resp, httpResp, err = c.DelegateTokenResourceApi.CreateDelegateToken(ctx, c.AccountId, delegateToken.Name, &nextgen.DelegateTokenResourceApiCreateDelegateTokenOpts{
			OrgIdentifier:     helpers.BuildField(d, "org_id"),
			ProjectIdentifier: helpers.BuildField(d, "project_id"),
		})
	} else {
		resp, httpResp, err = c.DelegateTokenResourceApi.RevokeDelegateToken(ctx, c.AccountId, delegateToken.Name, &nextgen.DelegateTokenResourceApiRevokeDelegateTokenOpts{
			OrgIdentifier:     helpers.BuildField(d, "org_id"),
			ProjectIdentifier: helpers.BuildField(d, "project_id"),
		})
	}

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	readDelegateToken(d, resp.Resource)

	return nil
}

func resourceDelegateTokenDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	var err error
	var resp nextgen.RestResponseDelegateTokenDetails
	var httpResp *http.Response

	delegateToken := buildDelegateToken(d)

	resp, httpResp, err = c.DelegateTokenResourceApi.RevokeDelegateToken(ctx, c.AccountId, delegateToken.Name, &nextgen.DelegateTokenResourceApiRevokeDelegateTokenOpts{
		OrgIdentifier:     helpers.BuildField(d, "org_id"),
		ProjectIdentifier: helpers.BuildField(d, "project_id"),
	})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	readDelegateToken(d, resp.Resource)

	return nil

}

func buildDelegateToken(d *schema.ResourceData) *nextgen.DelegateTokenDetails {
	delegateToken := &nextgen.DelegateTokenDetails{}

	if attr, ok := d.GetOk("account_id"); ok {
		delegateToken.AccountId = attr.(string)
	}

	if attr, ok := d.GetOk("name"); ok {
		delegateToken.Name = attr.(string)
	}

	if attr, ok := d.GetOk("created_at"); ok {
		delegateToken.CreatedAt = int64(attr.(int))
	}

	if attr, ok := d.GetOk("token_status"); ok {
		delegateToken.Status = attr.(string)
	}

	if attr, ok := d.GetOk("value"); ok {
		delegateToken.Value = attr.(string)
	}

	return delegateToken
}

func readDelegateToken(d *schema.ResourceData, delegateTokenDetails *nextgen.DelegateTokenDetails) {
	d.SetId(delegateTokenDetails.Name)
	d.Set("identifier", delegateTokenDetails.Name)
	d.Set("name", delegateTokenDetails.Name)
	d.Set("account_id", delegateTokenDetails.AccountId)
	d.Set("token_status", delegateTokenDetails.Status)
	d.Set("created_at", delegateTokenDetails.CreatedAt)
	d.Set("value", delegateTokenDetails.Value)
}
