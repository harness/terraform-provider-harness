package token

import (
	"context"
	"net/http"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceToken() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for creating tokens.",

		ReadContext:   resourceTokenRead,
		CreateContext: resourceTokenCreateOrUpdate,
		UpdateContext: resourceTokenCreateOrUpdate,
		DeleteContext: resourceTokenDelete,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"identifier": {
				Description: "Identifier of the Token",
				Type:        schema.TypeString,
				Required:    true,
			},
			"name": {
				Description: "Name of the Token",
				Type:        schema.TypeString,
				Required:    true,
			},
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
			"org_id": {
				Description: "Organization Identifier for the Entity",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"project_id": {
				Description: "Project Identifier for the Entity",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"valid_from": {
				Description: "This is the time from which the Token is valid. The time is in milliseconds",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"valid_to": {
				Description: "This is the time till which the Token is valid. The time is in milliseconds",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"scheduled_expire_time": {
				Description: "Scheduled expiry time in milliseconds",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"valid": {
				Description: "Boolean value to indicate if Token is valid or not.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"description": {
				Description: "Description of the Token",
				Type:        schema.TypeString,
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
			"tags": {
				Description: "Tags for the Token",
				Type:        schema.TypeMap,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}

	return resource
}

func resourceTokenRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Get("identifier").(string)

	type_ := d.Get("apikey_type").(string)
	parentId := d.Get("parent_id").(string)
	apikeyId := d.Get("apikey_id").(string)

	resp, httpResp, err := c.TokenApi.ListAggregatedTokens(ctx, c.AccountId, type_, parentId, apikeyId, &nextgen.TokenApiListAggregatedTokensOpts{
		OrgIdentifier:     helpers.BuildField(d, "org_id"),
		ProjectIdentifier: helpers.BuildField(d, "project_id"),
		Identifiers:       optional.NewInterface(id),
	})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	if resp.Data.Content != nil && len(resp.Data.Content) == 1 {
		readToken(d, resp.Data.Content[0].Token)
	}

	return nil
}

func resourceTokenCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	var err error
	var resp nextgen.ResponseDtoToken
	var httpResp *http.Response

	id := d.Id()
	token := buildToken(d)

	if id == "" {
		_, httpResp, err = c.TokenApi.CreateToken(ctx, c.AccountId, &nextgen.TokenApiCreateTokenOpts{Body: optional.NewInterface(token)})
		if err == nil {
			return resourceTokenRead(ctx, d, meta)
		}
	} else {
		resp, httpResp, err = c.TokenApi.UpdateToken(ctx, c.AccountId, id, &nextgen.TokenApiUpdateTokenOpts{Body: optional.NewInterface(token)})
	}

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	readToken(d, resp.Data)

	return nil
}

func resourceTokenDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Id()

	type_ := d.Get("apikey_type").(string)
	parentId := d.Get("parent_id").(string)
	apikeyId := d.Get("apikey_id").(string)

	_, httpResp, err := c.TokenApi.DeleteToken(ctx, id, c.AccountId, type_, parentId, apikeyId, &nextgen.TokenApiDeleteTokenOpts{
		OrgIdentifier:     helpers.BuildField(d, "org_id"),
		ProjectIdentifier: helpers.BuildField(d, "project_id"),
	})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	return nil
}

func buildToken(d *schema.ResourceData) *nextgen.Token {
	token := &nextgen.Token{}

	if attr, ok := d.GetOk("identifier"); ok {
		token.Identifier = attr.(string)
	}

	if attr, ok := d.GetOk("name"); ok {
		token.Name = attr.(string)
	}

	if attr, ok := d.GetOk("description"); ok {
		token.Description = attr.(string)
	}

	if attr, ok := d.GetOk("tags"); ok {
		token.Tags = attr.(map[string]string)
	}

	if attr, ok := d.GetOk("apikey_id"); ok {
		token.ApiKeyIdentifier = attr.(string)
	}

	if attr, ok := d.GetOk("apikey_type"); ok {
		token.ApiKeyType = attr.(string)
	}

	if attr, ok := d.GetOk("parent_id"); ok {
		token.ParentIdentifier = attr.(string)
	}

	if attr, ok := d.GetOk("valid_from"); ok {
		token.ValidFrom = int64(attr.(int))
	}

	if attr, ok := d.GetOk("valid_to"); ok {
		token.ValidTo = int64(attr.(int))
	}

	if attr, ok := d.GetOk("valid"); ok {
		token.Valid = attr.(bool)
	}

	if attr, ok := d.GetOk("scheduled_expire_time"); ok {
		token.ScheduledExpireTime = int64(attr.(int))
	}

	if attr, ok := d.GetOk("account_id"); ok {
		token.AccountIdentifier = attr.(string)
	}

	if attr, ok := d.GetOk("org_id"); ok {
		token.OrgIdentifier = attr.(string)
	}

	if attr, ok := d.GetOk("project_id"); ok {
		token.ProjectIdentifier = attr.(string)
	}

	if attr, ok := d.GetOk("email"); ok {
		token.Email = attr.(string)
	}

	if attr, ok := d.GetOk("username"); ok {
		token.Username = attr.(string)
	}

	if attr, ok := d.GetOk("encodedPassword"); ok {
		token.EncodedPassword = attr.(string)
	}
	return token
}

func readToken(d *schema.ResourceData, token *nextgen.Token) {
	d.SetId(token.Identifier)
	d.Set("name", token.Name)
	d.Set("description", token.Description)
	d.Set("tags", token.Tags)
	d.Set("apikey_id", token.ApiKeyIdentifier)
	d.Set("apikey_type", token.ApiKeyType)
	d.Set("parent_id", token.ParentIdentifier)
	d.Set("valid_from", token.ValidFrom)
	d.Set("valid_to", token.ValidTo)
	d.Set("valid", token.Valid)
	d.Set("scheduled_expire_time", token.ScheduledExpireTime)
	d.Set("account_id", token.AccountIdentifier)
	d.Set("project_id", token.ProjectIdentifier)
	d.Set("org_id", token.OrgIdentifier)
	d.Set("email", token.Email)
	d.Set("username", token.Username)
	d.Set("encodedPassword", token.EncodedPassword)
}
