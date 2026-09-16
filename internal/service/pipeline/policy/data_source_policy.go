package policy

import (
	"context"
	"errors"
	"net/http"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/policymgmt"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourcePolicy() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving a Harness policy.",

		ReadContext: dataSourceProjectRead,

		Schema: map[string]*schema.Schema{
			"identifier": {
				Description: "Identifier of the policy.",
				Type:        schema.TypeString,
				Required:    true,
				Computed:    false,
			},
			"rego": {
				Description: "Rego code for the policy.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
		},
	}

	helpers.SetMultiLevelDatasourceSchema(resource.Schema)

	return resource
}

func dataSourceProjectRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*internal.Session).GetPolicyManagementClient()
	id := d.Get("identifier").(string)

	var err error
	var policy policymgmt.Policy
	var httpResp *http.Response

	if id != "" {
		policy, _, _ = c.PoliciesApi.PoliciesFind(ctx, id, &policymgmt.PoliciesApiPoliciesFindOpts{
			AccountIdentifier: optional.NewString(meta.(*internal.Session).AccountId),
			XApiKey:           optional.NewString(meta.(*internal.Session).PLClient.ApiKey),
		})
	} else {
		return diag.FromErr(errors.New("identifier must be specified"))
	}
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	bla := policymgmt.Policy{}
	if policy == bla {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	readPolicy(d, policy)
	return nil
}
