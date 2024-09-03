package cluster

import (
	"context"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceEnvironmentClustersMapping() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving Harness Gitops clusters mapped to Harness Environment.",

		ReadContext: dataSourceResourceEnvironmentClustersMappingRead,

		Schema: map[string]*schema.Schema{
			"identifier": {
				// this identifier is not used anywhere and has no meaning for the resource either but keeping it maintain backward compatibility
				Description: "identifier for the cluster mapping(can be given any value).",
				Type:        schema.TypeString,
				Required:    true,
			},
			"env_id": {
				Description: "environment identifier.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"org_id": {
				Description: "org_id of the environment.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"project_id": {
				Description: "project_id of the environment.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"scope": {
				Description: "scope at which the environment exists in harness.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"clusters": {
				Description: "list of cluster identifiers and names",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"identifier": {
							Description: "identifier of the cluster",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"name": {
							Description: "name of the cluster",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"agent_identifier": {
							Description: "agent identifier of the cluster (include scope prefix)",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"scope": {
							Description: "scope at which the cluster exists in harness gitops, one of \"ACCOUNT\", \"ORGANIZATION\", \"PROJECT\". Scope of environment to which clusters are being mapped must be lower or equal to in hierarchy than the scope of the cluster",
							Type:        schema.TypeString,
							Optional:    true,
						},
					}},
			},
		},
	}
	return resource
}

func dataSourceResourceEnvironmentClustersMappingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	envId := d.Get("env_id").(string)
	clustersApiGetClusterListOpts := &nextgen.ClustersApiGetClusterListOpts{}
	if d.Get("org_id").(string) != "" {
		clustersApiGetClusterListOpts.OrgIdentifier = optional.NewString(d.Get("org_id").(string))
	}
	if d.Get("project_id").(string) != "" {
		clustersApiGetClusterListOpts.ProjectIdentifier = optional.NewString(d.Get("project_id").(string))
	}
	resp, httpResp, err := c.ClustersApi.GetClusterList(ctx, c.AccountId, envId, clustersApiGetClusterListOpts)

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	// Soft delete lookup error handling
	// https://harness.atlassian.net/browse/PL-23765
	if resp.Data == nil {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	readEnvironmentClustersMapping(d, &resp.Data.Content)

	return nil
}
