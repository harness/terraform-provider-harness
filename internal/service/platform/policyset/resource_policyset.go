package policyset

import (
	"context"
	"net/http"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/policymgmt"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourcePolicyset() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for creating a Harness Policyset.",

		ReadContext:   resourcePolicysetRead,
		UpdateContext: resourcePolicysetCreateOrUpdate,
		DeleteContext: resourcePolicysetDelete,
		CreateContext: resourcePolicysetCreateOrUpdate,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"type": {
				Description: "Type for the policyset.",
				Type:        schema.TypeString,
				Required:    true,
				Computed:    false,
			},
			"action": {
				Description: "Action for the policyset.",
				Type:        schema.TypeString,
				Required:    true,
				Computed:    false,
			},
			"enabled": {
				Description: "Enabled for the policyset.",
				Type:        schema.TypeBool,
				Required:    false,
				Computed:    false,
				Optional:    true,
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

	helpers.SetMultiLevelResourceSchema(resource.Schema)

	return resource
}

func resourcePolicysetRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*internal.Session).GetPolicyManagementClient()

	id := d.Id()

	localVarOptionals := policymgmt.PolicysetsApiPolicysetsFindOpts{
		AccountIdentifier: optional.NewString(meta.(*internal.Session).AccountId),
		XApiKey:           optional.NewString(meta.(*internal.Session).PLClient.ApiKey),
	}
	// check for project and org
	if d.Get("project_id").(string) != "" {
		localVarOptionals.ProjectIdentifier = helpers.BuildField(d, "project_id")
	}
	if d.Get("org_id").(string) != "" {
		localVarOptionals.OrgIdentifier = helpers.BuildField(d, "org_id")
	}

	policy, httpResp, err := c.PolicysetsApi.PolicysetsFind(ctx, id, &localVarOptionals)

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	readPolicyset(d, policy)

	return nil
}

func resourcePolicysetCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*internal.Session).GetPolicyManagementClient()
	var err error
	var responsePolicyset policymgmt.PolicySet2
	var httpResp *http.Response
	id := d.Id()

	if id == "" {
		body := policymgmt.CreateRequestBody2{
			Name:       d.Get("name").(string),
			Identifier: d.Get("identifier").(string),
			Action:     d.Get("action").(string),
			Type_:      d.Get("type").(string),
			Enabled:    d.Get("enabled").(bool),
		}
		localVarOptionals := policymgmt.PolicysetsApiPolicysetsCreateOpts{
			AccountIdentifier: optional.NewString(meta.(*internal.Session).AccountId),

			XApiKey: optional.NewString(meta.(*internal.Session).PLClient.ApiKey),
		}
		// check for project and org
		if d.Get("project_id").(string) != "" {
			localVarOptionals.ProjectIdentifier = helpers.BuildField(d, "project_id")
		}
		if d.Get("org_id").(string) != "" {
			localVarOptionals.OrgIdentifier = helpers.BuildField(d, "org_id")
		}
		createPolicySet, _, createErr := c.PolicysetsApi.PolicysetsCreate(ctx, body, &localVarOptionals)
		if createErr != nil {
			return helpers.HandleApiError(createErr, d, httpResp)
		}
		id = createPolicySet.Identifier
	}
	// NB we can only add policies to a policyset after it has been created, so we need to do this in the update
	policies := buildPolicies(d)
	body := policymgmt.UpdateRequestBody2{
		Name:     d.Get("name").(string),
		Type_:    d.Get("type").(string),
		Action:   d.Get("action").(string),
		Enabled:  d.Get("enabled").(bool),
		Policies: policies,
	}

	localVarOptionals := policymgmt.PolicysetsApiPolicysetsUpdateOpts{
		AccountIdentifier: optional.NewString(meta.(*internal.Session).AccountId),
		XApiKey:           optional.NewString(meta.(*internal.Session).PLClient.ApiKey),
	}
	if d.Get("project_id").(string) != "" {
		localVarOptionals.ProjectIdentifier = helpers.BuildField(d, "project_id")
	}
	if d.Get("org_id").(string) != "" {
		localVarOptionals.OrgIdentifier = helpers.BuildField(d, "org_id")
	}
	httpResp, err = c.PolicysetsApi.PolicysetsUpdate(ctx, body, id, &localVarOptionals)
	if err == nil && httpResp.StatusCode == http.StatusNoContent {
		// if we get a 204, we need to get the policy again to get the updated values
		findLocalVarOptionals := policymgmt.PolicysetsApiPolicysetsFindOpts{
			AccountIdentifier: optional.NewString(meta.(*internal.Session).AccountId),
			XApiKey:           optional.NewString(meta.(*internal.Session).PLClient.ApiKey),
		}
		// check for project and org
		if d.Get("project_id").(string) != "" {
			findLocalVarOptionals.ProjectIdentifier = helpers.BuildField(d, "project_id")
		}
		if d.Get("org_id").(string) != "" {
			findLocalVarOptionals.OrgIdentifier = helpers.BuildField(d, "org_id")
		}
		responsePolicyset, httpResp, err = c.PolicysetsApi.PolicysetsFind(ctx, id, &findLocalVarOptionals)
	}

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	readPolicyset(d, responsePolicyset)
	return nil
}

func buildPolicies(d *schema.ResourceData) []policymgmt.Linkedpolicyidentifier {
	policies := []policymgmt.Linkedpolicyidentifier{}
	rawPolicies := d.Get("policies").([]interface{})
	for _, policy := range rawPolicies {
		policy := policy.(map[string]interface{})
		policies = append(policies, policymgmt.Linkedpolicyidentifier{
			Identifier: policy["identifier"].(string),
			Severity:   policy["severity"].(string),
		})
	}
	return policies
}

func resourcePolicysetDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*internal.Session).GetPolicyManagementClient()

	localVarOptionals := policymgmt.PolicysetsApiPolicysetsDeleteOpts{
		AccountIdentifier: optional.NewString(meta.(*internal.Session).AccountId),
		XApiKey:           optional.NewString(meta.(*internal.Session).PLClient.ApiKey),
	}
	// check for project and org
	if d.Get("project_id").(string) != "" {
		localVarOptionals.ProjectIdentifier = helpers.BuildField(d, "project_id")
	}
	if d.Get("org_id").(string) != "" {
		localVarOptionals.OrgIdentifier = helpers.BuildField(d, "org_id")
	}
	httpResp, err := c.PolicysetsApi.PolicysetsDelete(ctx, d.Id(), &localVarOptionals)

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	return nil
}

func readPolicyset(d *schema.ResourceData, policy policymgmt.PolicySet2) {
	d.SetId(policy.Identifier)
	_ = d.Set("identifier", policy.Identifier)
	_ = d.Set("org_id", policy.OrgId)
	_ = d.Set("project_id", policy.ProjectId)
	_ = d.Set("name", policy.Name)
	_ = d.Set("action", policy.Action)
	_ = d.Set("type", policy.Type_)
	_ = d.Set("enabled", policy.Enabled)
	_ = d.Set("policies", flattenPolicies(policy.Policies))
}

func flattenPolicies(policies []policymgmt.LinkedPolicy2) []map[string]interface{} {
	var policyList []map[string]interface{}
	for _, policy := range policies {
		policyList = append(policyList, map[string]interface{}{
			"identifier": policy.Identifier,
			"severity":   policy.Severity,
		})
	}
	return policyList
}
