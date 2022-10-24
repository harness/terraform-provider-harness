package gnupg

import (
	"context"

	"github.com/antihax/optional"
	hh "github.com/harness/harness-go-sdk/harness/helpers"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceGitopsGnupg() *schema.Resource {
	resource := &schema.Resource{
		Description: "GPG public key in the server's configuration.",

		CreateContext: resourceGitopsGnupgCreate,
		ReadContext:   resourceGitopsGnupgRead,
		UpdateContext: resourceGitopsGnupgCreate,
		DeleteContext: resourceGitopsGnupgDelete,
		Importer:      helpers.GitopsAgentResourceImporter,

		Schema: map[string]*schema.Schema{
			"account_id": {
				Description: "account Identifier for the Entity.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"org_id": {
				Description: "organization Identifier for the Entity.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"project_id": {
				Description: "project Identifier for the Entity.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"agent_id": {
				Description: "agent identifier of the cluster.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"request": {
				Description: "GnuPGPublicKey is a representation of a GnuPG public key",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"upsert": {
							Description: "if the gnupg should be upserted.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"publickey": {
							Description: "publickey details.",
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key_id": {
										Description: "KeyID specifies the key ID, in hexadecimal string format.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"fingerprint": {
										Description: "Fingerprint is the fingerprint of the key",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"owner": {
										Description: "Owner holds the owner identification, e.g. a name and e-mail address",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"trust": {
										Description: "Trust holds the level of trust assigned to this key",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"sub_type": {
										Description: "SubType holds the key's sub type",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"key_data": {
										Description: "KeyData holds the raw key data, in base64 encoded format",
										Type:        schema.TypeString,
										Optional:    true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
	return resource
}

func resourceGitopsGnupgCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	ctx = context.WithValue(ctx, nextgen.ContextAccessToken, hh.EnvVars.BearerToken.Get())
	var agentIdentifier, accountIdentifier string
	accountIdentifier = c.AccountId
	if attr, ok := d.GetOk("agent_id"); ok {
		agentIdentifier = attr.(string)
	}
	// if attr, ok := d.GetOk("org_id"); ok {
	// 	orgIdentifier = attr.(string)
	// }
	// if attr, ok := d.GetOk("project_id"); ok {
	// 	projectIdentifier = attr.(string)
	// }

	createGnupgRequest := buildGnupgCreateRequest(d)
	respCreate, httpRespCreate, errCreate := c.GnuPGPKeysApi.AgentGPGKeyServiceCreate(ctx, *createGnupgRequest, agentIdentifier,
		&nextgen.GnuPGPKeysApiAgentGPGKeyServiceCreateOpts{
			AccountIdentifier: optional.NewString(accountIdentifier),
			// OrgIdentifier:     optional.NewString(orgIdentifier),
			// ProjectIdentifier: optional.NewString(projectIdentifier),
		})

	if errCreate != nil {
		return helpers.HandleApiError(errCreate, d, httpRespCreate)
	}

	// Soft delete lookup error handling
	// https://harness.atlassian.net/browse/PL-23765
	if &respCreate == nil {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	// respRead, httpRespRead, errRead := c.GnuPGPKeysApi.GnuPGKeyServiceListGPGKeys(ctx, c.AccountId, &nextgen.GPGKeysApiGnuPGKeyServiceListGPGKeysOpts{})

	// if errRead != nil {
	// 	return helpers.HandleApiError(errRead, d, httpRespRead)
	// }
	// // Soft delete lookup error handling
	// // https://harness.atlassian.net/browse/PL-23765
	// if &respRead == nil || respRead.Content == nil || &respRead.Content[0] == nil {
	// 	d.SetId("")
	// 	d.MarkNewResource()
	// 	return nil
	// }
	// readGnupgKey(d, respRead.Content[0].GnuPGPublicKey)
	return nil
}

func resourceGitopsGnupgRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	ctx = context.WithValue(ctx, nextgen.ContextAccessToken, hh.EnvVars.BearerToken.Get())
	// agentIdentifier := d.Get("agent_id").(string)

	resp, httpResp, err := c.GnuPGPKeysApi.GnuPGKeyServiceListGPGKeys(ctx, c.AccountId, &nextgen.GPGKeysApiGnuPGKeyServiceListGPGKeysOpts{})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	// Soft delete lookup error handling
	// https://harness.atlassian.net/browse/PL-23765
	if &resp == nil {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}
	readGnupgKey(d, resp.Content[0].GnuPGPublicKey)
	return nil
}

func resourceGitopsGnupgDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	ctx = context.WithValue(ctx, nextgen.ContextAccessToken, hh.EnvVars.BearerToken.Get())
	agentIdentifier := d.Get("agent_id").(string)

	_, httpResp, err := c.GnuPGPKeysApi.AgentGPGKeyServiceDelete(ctx, agentIdentifier, d.Id(), &nextgen.GnuPGPKeysApiAgentGPGKeyServiceDeleteOpts{
		AccountIdentifier: optional.NewString(c.AccountId),
		// OrgIdentifier:     optional.NewString(d.Get("org_id").(string)),
		// ProjectIdentifier: optional.NewString(d.Get("project_id").(string)),
	})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}
	return nil
}

func readGnupgKey(d *schema.ResourceData, gpgkey *nextgen.GpgkeysGnuPgPublicKey) {
	d.SetId(gpgkey.KeyID)
	request := map[string]interface{}{}
	requestList := []interface{}{}
	publickey := map[string]interface{}{}
	publickeyList := []interface{}{}

	publickey["fingerprint"] = gpgkey.Fingerprint
	publickey["owner"] = gpgkey.Owner
	publickey["trust"] = gpgkey.Trust
	publickey["sub_type"] = gpgkey.SubType
	publickey["key_data"] = gpgkey.KeyData

	publickeyList = append(publickeyList, publickey)
	request["publickey"] = publickeyList
	requestList = append(requestList, request)
	d.Set("request", requestList)
}

func buildGnupgCreateRequest(d *schema.ResourceData) *nextgen.GpgkeysGnuPgPublicKeyCreateRequest {
	var upsert bool
	if attr, ok := d.GetOk("request"); ok {
		request := attr.([]interface{})[0].(map[string]interface{})
		upsert = request["upsert"].(bool)
	}
	return &nextgen.GpgkeysGnuPgPublicKeyCreateRequest{
		Upsert:    upsert,
		Publickey: buildPublickeyDetails(d),
	}
}

func buildPublickeyDetails(d *schema.ResourceData) *nextgen.GpgkeysGnuPgPublicKey {
	var publickeyDetails nextgen.GpgkeysGnuPgPublicKey
	var request map[string]interface{}
	if attr, ok := d.GetOk("request"); ok {
		request = attr.([]interface{})[0].(map[string]interface{})
		if request["publickey"] != nil && len(request["publickey"].([]interface{})) > 0 {

			requestPublicKey := request["publickey"].([]interface{})[0].(map[string]interface{})

			if requestPublicKey["key_ID"] != nil {
				publickeyDetails.KeyID = requestPublicKey["key_ID"].(string)
			}

			if requestPublicKey["fingerprint"] != nil {
				publickeyDetails.Fingerprint = requestPublicKey["fingerprint"].(string)
			}

			if requestPublicKey["owner"] != nil {
				publickeyDetails.Owner = requestPublicKey["owner"].(string)
			}

			if requestPublicKey["trust"] != nil {
				publickeyDetails.Trust = requestPublicKey["trust"].(string)
			}

			if requestPublicKey["sub_type"] != nil {
				publickeyDetails.SubType = requestPublicKey["sub_type"].(string)
			}

			if requestPublicKey["key_data"] != nil {
				publickeyDetails.KeyData = requestPublicKey["key_data"].(string)
			}

		}
	}
	return &publickeyDetails
}
