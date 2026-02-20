package delegatetoken

import (
	"context"
	"log"
	"net/http"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceDelegateToken() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for creating delegate tokens. Delegate tokens can be created and revoked. Revoked tokens immediately stop functioning and are only purged after 30 days, meaning you cannot recreate a token with the same name within that period.\nTo revoke a token, set token_status field to \"REVOKED\".",

		ReadContext:   resourceDelegateTokenRead,
		CreateContext: resourceDelegateTokenCreate,
		UpdateContext: resourceDelegateTokenRevoke,
		DeleteContext: resourceDelegateTokenRevoke,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"name": {
				Description: "Name of the delegate token",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"account_id": {
				Description: "Account Identifier for the Entity",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"org_id": {
				Description: "Org Identifier for the Entity",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
			},
			"project_id": {
				Description:  "Project Identifier for the Entity",
				Type:         schema.TypeString,
				Optional:     true,
				RequiredWith: []string{"org_id"},
				ForceNew:     true,
			},
			"token_status": {
				Description:  "Status of Delegate Token (ACTIVE or REVOKED). If left empty both active and revoked tokens will be assumed",
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.StringInSlice([]string{"ACTIVE", "REVOKED"}, false),
			},
			"value": {
				Description: "Value of the delegate token. Encoded in base64.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"created_at": {
				Description: "Time when the delegate token is created. This is an epoch timestamp.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"revoke_after": {
				Description: "Epoch time in milliseconds after which the token will be marked as revoked. There can be a delay of up to one hour from the epoch value provided and actual revoking of the token.",
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
			},
			"created_by": {
				Description: "created by details",
				Type:        schema.TypeMap,
				Optional:    true,
				Computed:    true,
			},
		},
	}

	return resource
}

func resourceDelegateTokenRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	resp, httpResp, err := c.DelegateTokenResourceApi.GetCgDelegateTokens(ctx, c.AccountId, &nextgen.DelegateTokenResourceApiGetCgDelegateTokensOpts{
		OrgIdentifier:     helpers.BuildField(d, "org_id"),
		ProjectIdentifier: helpers.BuildField(d, "project_id"),
		Name:              helpers.BuildField(d, "name"),
		Status:            helpers.BuildField(d, "token_status"),
	})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	if resp.Resource != nil && (len(resp.Resource) > 0) {
		readDelegateToken(d, &resp.Resource[0])
	}

	return nil
}

func resourceDelegateTokenCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	var err error
	var resp nextgen.RestResponseDelegateTokenDetails
	var httpResp *http.Response

	delegateToken := buildDelegateToken(d)

	opts := &nextgen.DelegateTokenResourceApiCreateDelegateTokenOpts{
		OrgIdentifier:     helpers.BuildField(d, "org_id"),
		ProjectIdentifier: helpers.BuildField(d, "project_id"),
	}

	if attr, ok := d.GetOk("revoke_after"); ok {
		opts.RevokeAfter = optional.NewInt64(int64(attr.(int)))
	}

	resp, httpResp, err = c.DelegateTokenResourceApi.CreateDelegateToken(ctx, c.AccountId, delegateToken.Name, opts)

	if err != nil && httpResp != nil {
		log.Printf("[INFO] Failed to create delegate token %q. This may happen if a token with the same name already exists in the scope or was recently deleted (within the 30-day purge window). Enable Terraform debug logs to view the full API error response.", delegateToken.Name)
		return helpers.HandleApiError(err, d, httpResp)
	} else if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	readDelegateToken(d, resp.Resource)

	return nil
}

func resourceDelegateTokenRevoke(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	var err error
	var resp nextgen.RestResponseDelegateTokenDetails
	var httpResp *http.Response

	delegateToken := buildDelegateToken(d)

	resp, httpResp, err = c.DelegateTokenResourceApi.RevokeCgDelegateToken(ctx, c.AccountId, delegateToken.Name, &nextgen.DelegateTokenResourceApiRevokeCgDelegateTokenOpts{
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
	d.Set("name", delegateTokenDetails.Name)
	d.Set("account_id", delegateTokenDetails.AccountId)
	d.Set("token_status", delegateTokenDetails.Status)
	d.Set("created_at", delegateTokenDetails.CreatedAt)
	d.Set("created_by", readCreatedByData(delegateTokenDetails.CreatedByNgUser.Type_, delegateTokenDetails.CreatedByNgUser.Name, delegateTokenDetails.CreatedByNgUser.Jwtclaims))
	d.Set("value", delegateTokenDetails.Value)
}

func readCreatedByData(userType string, name_ string, details map[string]string) map[string]string {
	var result = make(map[string]string)

	result["type"] = userType
	result["name"] = name_

	for key, value := range details {
		result[key] = value
	}

	return result
}
