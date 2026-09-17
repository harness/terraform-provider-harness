package policyset

import (
	"bytes"
	"context"
	"errors"
	"fmt"
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
				Description: "List of policy identifiers / severity for the policyset. Deprecated: use 'policy_references' instead - this field is order-sensitive and the underlying API does not guarantee a stable order for linked policies across reads.",
				Type:        schema.TypeList,
				Computed:    true,
				Optional:    true,
				MinItems:    1,
				Deprecated:  "The 'policies' field is deprecated. Use 'policy_references' instead. This field will be removed in a future version.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"identifier": {
							Description: "Identifier of the policy. Policies at a broader scope than the policyset are returned with a scope-qualified identifier (e.g. 'account.my_policy' or 'org.my_policy').",
							Type:        schema.TypeString,
							Required:    true,
						},
						"severity": {
							Description: "Policy failure response - 'warning' for continuation, 'error' for exit",
							Type:        schema.TypeString,
							Required:    true,
						},
					},
				},
			},
			"policy_references": {
				Description: "Set of policy identifiers / severity for the policyset. Order is not significant. Preferred over the deprecated 'policies' field.",
				Type:        schema.TypeSet,
				Computed:    true,
				Optional:    true,
				MinItems:    1,
				Set: func(v interface{}) int {
					m := v.(map[string]interface{})
					var buf bytes.Buffer
					buf.WriteString(fmt.Sprintf("%s-", m["identifier"].(string)))
					buf.WriteString(fmt.Sprintf("%s-", m["severity"].(string)))
					return hashcode(buf.String())
				},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"identifier": {
							Description: "Identifier of the policy. Policies at a broader scope than the policyset are returned with a scope-qualified identifier (e.g. 'account.my_policy' or 'org.my_policy').",
							Type:        schema.TypeString,
							Required:    true,
						},
						"severity": {
							Description: "Policy failure response - 'warning' for continuation, 'error' for exit",
							Type:        schema.TypeString,
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
