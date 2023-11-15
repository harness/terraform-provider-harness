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
				Description: "Account Identifier for the GnuPG Key.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"agent_id": {
				Description: "Agent identifier for the GnuPG Key.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"org_id": {
				Description: "Organization Identifier for the GnuPG Key.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"project_id": {
				Description: "Project Identifier for the GnuPG Key.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"identifier": {
				Description: "Identifier for the GnuPG Key.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"request": {
				Description: "GnuPGPublicKey is a representation of a GnuPG public key",
				Type:        schema.TypeList,
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"upsert": {
							Description: "Indicates if the GnuPG Key should be inserted if not present or updated if present.",
							Type:        schema.TypeBool,
							Required:    true,
						},
						"publickey": {
							Description: "Public key details.",
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key_id": {
										Description: "KeyID specifies the key ID, in hexadecimal string format.",
										Type:        schema.TypeString,
										Computed:    true,
									},
									"fingerprint": {
										Description: "Fingerprint is the fingerprint of the key",
										Type:        schema.TypeString,
										Computed:    true,
									},
									"owner": {
										Description: "Owner holds the owner identification, e.g. a name and e-mail address",
										Type:        schema.TypeString,
										Computed:    true,
									},
									"trust": {
										Description: "Trust holds the level of trust assigned to this key",
										Type:        schema.TypeString,
										Computed:    true,
									},
									"sub_type": {
										Description: "SubType holds the key's sub type",
										Type:        schema.TypeString,
										Computed:    true,
									},
									"key_data": {
										Description: "KeyData holds the raw key data, in base64 encoded format.",
										Type:        schema.TypeString,
										Required:    true,
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
	var agentIdentifier, accountIdentifier, orgIdentifier, projectIdentifier string
	accountIdentifier = c.AccountId
	if attr, ok := d.GetOk("agent_id"); ok {
		agentIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("project_id"); ok {
		projectIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("org_id"); ok {
		orgIdentifier = attr.(string)
	}
	createGnupgRequest := buildGnupgCreateRequest(d)
	respCreate, httpRespCreate, errCreate := c.GnuPGPKeysApi.AgentGPGKeyServiceCreate(ctx, *createGnupgRequest, agentIdentifier,
		&nextgen.GnuPGPKeysApiAgentGPGKeyServiceCreateOpts{
			AccountIdentifier: optional.NewString(accountIdentifier),
			OrgIdentifier:     optional.NewString(orgIdentifier),
			ProjectIdentifier: optional.NewString(projectIdentifier),
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

	readGnupgKeyCreate(d, &respCreate, accountIdentifier, agentIdentifier, orgIdentifier, projectIdentifier)
	return nil
}

func resourceGitopsGnupgRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	ctx = context.WithValue(ctx, nextgen.ContextAccessToken, hh.EnvVars.BearerToken.Get())
	var agentIdentifier, orgIdentifier, projectIdentifier string
	keyId := d.Get("identifier").(string)
	if attr, ok := d.GetOk("agent_id"); ok {
		agentIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("project_id"); ok {
		projectIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("org_id"); ok {
		orgIdentifier = attr.(string)
	}
	resp, httpResp, err := c.GnuPGPKeysApi.AgentGPGKeyServiceGet(ctx, agentIdentifier, keyId, c.AccountId, &nextgen.GnuPGPKeysApiAgentGPGKeyServiceGetOpts{
		OrgIdentifier:     optional.NewString(orgIdentifier),
		ProjectIdentifier: optional.NewString(projectIdentifier),
	})

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
	readGnupgKey(d, &resp, c.AccountId, agentIdentifier, orgIdentifier, projectIdentifier)
	return nil
}

func resourceGitopsGnupgDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	ctx = context.WithValue(ctx, nextgen.ContextAccessToken, hh.EnvVars.BearerToken.Get())
	var agentIdentifier, orgIdentifier, projectIdentifier string
	keyId := d.Get("identifier").(string)
	if attr, ok := d.GetOk("agent_id"); ok {
		agentIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("project_id"); ok {
		projectIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("org_id"); ok {
		orgIdentifier = attr.(string)
	}

	_, httpResp, err := c.GnuPGPKeysApi.AgentGPGKeyServiceDelete(ctx, agentIdentifier, keyId, &nextgen.GnuPGPKeysApiAgentGPGKeyServiceDeleteOpts{
		AccountIdentifier: optional.NewString(c.AccountId),
		OrgIdentifier:     optional.NewString(orgIdentifier),
		ProjectIdentifier: optional.NewString(projectIdentifier),
	})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}
	return nil
}

func readGnupgKeyCreate(d *schema.ResourceData, gpgKey *nextgen.GpgkeysGnuPgPublicKeyCreateResponse, accountIdentifier string, agentIdentifier string, orgIdentifier string, projectIdentifier string) {
	readGnupgKey(d, &gpgKey.Created.Items[0], accountIdentifier, agentIdentifier, orgIdentifier, projectIdentifier)
}

func readGnupgKey(d *schema.ResourceData, gpgkey *nextgen.GpgkeysGnuPgPublicKey, accountIdentifier string, agentIdentifier string, orgIdentifier string, projectIdentifier string) {
	d.SetId(gpgkey.KeyID)
	d.Set("identifier", gpgkey.KeyID)
	d.Set("account_id", accountIdentifier)
	d.Set("agent_id", agentIdentifier)
	d.Set("org_id", orgIdentifier)
	d.Set("project_id", projectIdentifier)
	request := map[string]interface{}{}
	requestList := []interface{}{}
	publickey := map[string]interface{}{}
	publickeyList := []interface{}{}

	publickey["key_id"] = gpgkey.KeyID
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

			if requestPublicKey["key_id"] != nil {
				publickeyDetails.KeyID = requestPublicKey["key_id"].(string)
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
