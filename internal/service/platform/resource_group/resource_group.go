package resource_group

import (
	"context"
	"net/http"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceResourceGroup() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for creating a Harness Resource Group",

		ReadContext:   resourceResourceGroupRead,
		UpdateContext: resourceResourceGroupCreateOrUpdate,
		CreateContext: resourceResourceGroupCreateOrUpdate,
		DeleteContext: resourceResourceGroupDelete,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"color": {
				Description: "Color of the environment.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"account_id": {
				Description: "Account Identifier of the account",
				Type:        schema.TypeString,
				Required:    true,
			},
			"allowed_scope_levels": {
				Description: "The scope levels at which this resource group can be used",
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"included_scopes": {
				Description: "Included scopes",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account_id": {
							Description: "Account Identifier of the account",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"org_id": {
							Description: "Organization Identifier",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"project_id": {
							Description: "Project Identifier",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"filter": {
							Description: "Can be one of these 2 EXCLUDING_CHILD_SCOPES or INCLUDING_CHILD_SCOPES",
							Type:        schema.TypeString,
							Required:    true,
						},
					}},
			},
			"resource_filter": {
				Description: "Contains resource filter for a resource group",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"include_all_resources": {
							Description: "Include all resource or not",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"resources": {
							Description: "Resources for a resource group",
							Type:        schema.TypeSet,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"resource_type": {
										Description: "Type of the resource",
										Type:        schema.TypeString,
										Required:    true,
									},
									"identifiers": {
										Description: "List of the identifiers",
										Type:        schema.TypeSet,
										Optional:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"attribute_filter": {
										Description: "Used to filter resources on their attributes",
										Type:        schema.TypeList,
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"attribute_name": {
													Description: "Name of the attribute",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"attribute_values": {
													Description: "Value of the attributes",
													Type:        schema.TypeSet,
													Optional:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	helpers.SetMultiLevelResourceSchema(resource.Schema)

	return resource
}

func resourceResourceGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Get("identifier").(string)

	resp, httpResp, err := c.HarnessResourceGroupApi.GetResourceGroupV2(ctx, id, c.AccountId, &nextgen.HarnessResourceGroupApiGetResourceGroupV2Opts{
		OrgIdentifier:     helpers.BuildField(d, "org_id"),
		ProjectIdentifier: helpers.BuildField(d, "project_id"),
	})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	if resp.Data == nil {
		return nil
	}

	readResourceGroup(d, resp.Data.ResourceGroup)

	return nil

}

func resourceResourceGroupCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	var err error
	var resp nextgen.ResponseDtoResourceGroupV2Response
	var httpResp *http.Response

	id := d.Id()
	resourceGroup := buildResourceGroup(d)

	if id == "" {
		resp, httpResp, err = c.HarnessResourceGroupApi.CreateResourceGroupV2(ctx, nextgen.ResourceGroupV2Request{ResourceGroup: resourceGroup}, c.AccountId, &nextgen.HarnessResourceGroupApiCreateResourceGroupV2Opts{
			OrgIdentifier:     helpers.BuildField(d, "org_id"),
			ProjectIdentifier: helpers.BuildField(d, "project_id"),
		})
	} else {
		resp, httpResp, err = c.HarnessResourceGroupApi.UpdateResourceGroup1(ctx, nextgen.ResourceGroupV2Request{ResourceGroup: resourceGroup}, c.AccountId, d.Id(), &nextgen.HarnessResourceGroupApiUpdateResourceGroup1Opts{
			OrgIdentifier:     helpers.BuildField(d, "org_id"),
			ProjectIdentifier: helpers.BuildField(d, "project_id"),
		})
	}

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	readResourceGroup(d, resp.Data.ResourceGroup)

	return nil
}

func resourceResourceGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	_, httpResp, err := c.HarnessResourceGroupApi.DeleteResourceGroupV2(ctx, d.Id(), c.AccountId, &nextgen.HarnessResourceGroupApiDeleteResourceGroupV2Opts{
		OrgIdentifier:     helpers.BuildField(d, "org_id"),
		ProjectIdentifier: helpers.BuildField(d, "project_id"),
	})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	return nil
}

func buildResourceGroup(d *schema.ResourceData) *nextgen.ResourceGroupV2 {
	resourceGroup := &nextgen.ResourceGroupV2{
		ResourceFilter: &nextgen.ResourceFilter{},
	}

	if attr, ok := d.GetOk("account_id"); ok {
		resourceGroup.AccountIdentifier = attr.(string)
	}

	if attr, ok := d.GetOk("org_id"); ok {
		resourceGroup.OrgIdentifier = attr.(string)
	}

	if attr, ok := d.GetOk("project_id"); ok {
		resourceGroup.ProjectIdentifier = attr.(string)
	}

	if attr, ok := d.GetOk("color"); ok {
		resourceGroup.Color = attr.(string)
	}

	if attr, ok := d.GetOk("description"); ok {
		resourceGroup.Description = attr.(string)
	}

	if attr, ok := d.GetOk("name"); ok {
		resourceGroup.Name = attr.(string)
	}

	if attr, ok := d.GetOk("identifier"); ok {
		resourceGroup.Identifier = attr.(string)
	}

	if attr := d.Get("tags").(*schema.Set).List(); len(attr) > 0 {
		resourceGroup.Tags = helpers.ExpandTags(attr)
	}

	if attr, ok := d.GetOk("allowed_scope_levels"); ok {
		resourceGroup.AllowedScopeLevels = helpers.ExpandField(attr.(*schema.Set).List())
	}

	if attr, ok := d.GetOk("included_scopes"); ok {
		resourceGroup.IncludedScopes = helpers.ExpandScopeSelector(attr.(*schema.Set).List())
	}

	if attr, ok := d.GetOk("resource_filter"); ok {
		config := attr.([]interface{})[0].(map[string]interface{})
		if attr, ok := config["include_all_resources"]; ok {
			resourceGroup.ResourceFilter.IncludeAllResources = attr.(bool)
		}

		if attr, ok := config["resources"]; ok {
			resourceGroup.ResourceFilter.Resources = expandResources(attr.(*schema.Set).List())
		}
	}

	return resourceGroup
}

func readResourceGroup(d *schema.ResourceData, resourceGroup *nextgen.ResourceGroupV2) {
	d.SetId(resourceGroup.Identifier)
	d.Set("identifier", resourceGroup.Identifier)
	d.Set("name", resourceGroup.Name)
	d.Set("org_id", resourceGroup.OrgIdentifier)
	d.Set("project_id", resourceGroup.ProjectIdentifier)
	d.Set("account_id", resourceGroup.AccountIdentifier)
	d.Set("color", resourceGroup.Color)
	d.Set("tags", helpers.FlattenTags(resourceGroup.Tags))
	d.Set("description", resourceGroup.Description)
	d.Set("allowed_scope_levels", resourceGroup.AllowedScopeLevels)
	d.Set("included_scopes", expandIncludedScope(resourceGroup.IncludedScopes))
	d.Set("resource_filter", []interface{}{
		map[string]interface{}{
			"include_all_resources": resourceGroup.ResourceFilter.IncludeAllResources,
			"resources":             expandResourceFilter(resourceGroup.ResourceFilter.Resources),
		},
	})
}

func expandIncludedScope(includedScopes []nextgen.ScopeSelector) []interface{} {
	var result []interface{}
	for _, scopeSelector := range includedScopes {
		result = append(result, map[string]interface{}{
			"filter":     scopeSelector.Filter,
			"account_id": scopeSelector.AccountIdentifier,
			"org_id":     scopeSelector.OrgIdentifier,
			"project_id": scopeSelector.ProjectIdentifier,
		})
	}

	return result
}

func expandResourceFilter(resourceSelector []nextgen.ResourceSelectorV2) []interface{} {
	var result []interface{}
	for _, selector := range resourceSelector {
		result = append(result, map[string]interface{}{
			"resource_type":    selector.ResourceType,
			"identifiers":      selector.Identifiers,
			"attribute_filter": expandAttributeFilter(selector),
		})
	}

	return result
}

func expandAttributeFilter(resourceSelector nextgen.ResourceSelectorV2) []interface{} {
	var result []interface{}
	if resourceSelector.AttributeFilter != nil {
		result = append(result, map[string]interface{}{
			"attribute_name":   resourceSelector.AttributeFilter.AttributeName,
			"attribute_values": resourceSelector.AttributeFilter.AttributeValues,
		})
	}
	return result
}

func expandResources(resources []interface{}) []nextgen.ResourceSelectorV2 {
	var result []nextgen.ResourceSelectorV2
	for _, resourceSelector := range resources {
		r := nextgen.ResourceSelectorV2{}
		v := resourceSelector.(map[string]interface{})
		r.ResourceType = v["resource_type"].(string)
		r.Identifiers = helpers.ExpandField(v["identifiers"].(*schema.Set).List())
		if attr, ok := v["attribute_filter"]; ok {
			if len(attr.([]interface{})) != 0 {
				config := attr.([]interface{})[0].(map[string]interface{})
				r.AttributeFilter = &nextgen.AttributeFilter{}
				r.AttributeFilter.AttributeName = config["attribute_name"].(string)
				r.AttributeFilter.AttributeValues = helpers.ExpandField(config["attribute_values"].(*schema.Set).List())
			}
		}
		result = append(result, r)
	}
	return result
}
