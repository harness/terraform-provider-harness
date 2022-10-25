package gnupg

import (
	"context"

	hh "github.com/harness/harness-go-sdk/harness/helpers"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceGitopsGnupg() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving a GPG public key in the server's configuration.",

		ReadContext: dataSourceGitopsGnupgRead,

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
				Description: "agent identifier.",
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

func dataSourceGitopsGnupgRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	ctx = context.WithValue(ctx, nextgen.ContextAccessToken, hh.EnvVars.BearerToken.Get())
	// agentIdentifier := d.Get("agent_id").(string)

	resp, httpResp, err := c.GnuPGPKeysApi.GnuPGKeyServiceListGPGKeys(ctx, c.AccountId, &nextgen.GPGKeysApiGnuPGKeyServiceListGPGKeysOpts{})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	// Soft delete lookup error handling
	// https://harness.atlassian.net/browse/PL-23765
	if &resp == nil || resp.Content == nil || &resp.Content[0] == nil {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	readGnupgKey(d, resp.Content[0].GnuPGPublicKey)
	return nil
}
