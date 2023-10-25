package policyset

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

func DataSourcePolicyset() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving a Harness policyset.",

		ReadContext: dataSourceProjectRead,

		Schema: map[string]*schema.Schema{
			"action": {
				Description: "Action code for the policyset.",
				Type:        schema.TypeString,
				Optional:    false,
				Required:    true,
				Computed:    false,
			},
			"type": {
				Description: "Type of the policyset.",
				Type:        schema.TypeString,
				Optional:    false,
				Required:    true,
				Computed:    false,
			},
			"enabled": {
				Description: "Enabled for the policyset.",
				Type:        schema.TypeBool,
				Optional:    true,
				Required:    false,
				Computed:    false,
			},
			"policies": {
				Description: "List of policy identifiers / severity for the policyset.",
				Type:        schema.TypeList,
				Computed:    false,
				Optional:    true,
				Required:    false,
				MinItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"identifier": {
							Description: "Account Identifier of the account",
							Type:        schema.TypeString,
							Optional:    false,
							Required:    true,
						},
						"severity": {
							Description: "Policy failure response - 'warning' for continuation, 'error' for exit",
							Type:        schema.TypeString,
							Optional:    false,
							Required:    true,
						},
					},
				},
			},
		},
	}

	helpers.SetMultiLevelDatasourceSchemaIdentifierRequired(resource.Schema)

	return resource
}

func dataSourceProjectRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*internal.Session).GetPolicyManagementClient()
	id := d.Get("identifier").(string)

	var err error
	var policyset policymgmt.PolicySet
	var httpResp *http.Response

	if id != "" {
		policyset, _, _ = c.PolicysetsApi.PolicysetsFind(ctx, id, &policymgmt.PolicysetsApiPolicysetsFindOpts{
			AccountIdentifier: optional.NewString(meta.(*internal.Session).AccountId),
			XApiKey:           optional.NewString(meta.(*internal.Session).PLClient.ApiKey),
		})
	} else {
		return diag.FromErr(errors.New("identifier must be specified"))
	}
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	bla := policymgmt.PolicySet{}
	if policyset.Identifier == bla.Identifier {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	readPolicyset(d, policyset)
	return nil
}
