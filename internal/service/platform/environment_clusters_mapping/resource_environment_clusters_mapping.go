package cluster

import (
	"context"
	"net/http"
	"strings"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceEnvironmentClustersMapping() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for mapping environment with Harness Clusters.",

		ReadContext:   resourceEnvironmentClustersMappingRead,
		DeleteContext: resourceEnvironmentClustersMappingDelete,
		CreateContext: resourceEnvironmentClustersMappingClusterLink,
		UpdateContext: resourceEnvironmentClustersMappingClusterLink,
		// TODO: Implement Importer
		Importer: helpers.ProjectResourceImporter,

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

func resourceEnvironmentClustersMappingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

func resourceEnvironmentClustersMappingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	var err error
	var httpResp *http.Response
	env := buildEnvironmentClustersMappingCluster(d)

	_, httpResp, err = c.ClustersApi.UnlinkClustersInBatch(ctx, c.AccountId, &nextgen.ClustersApiUnlinkClustersInBatchOpts{
		Body: optional.NewInterface(env),
	})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	return nil
}

func resourceEnvironmentClustersMappingClusterLink(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	var err error
	var resp nextgen.ResponseDtoClusterBatchResponse
	var httpResp *http.Response
	env := buildEnvironmentClustersMappingCluster(d)

	resp, httpResp, err = c.ClustersApi.LinkClusters(ctx, c.AccountId, &nextgen.ClustersApiLinkClustersOpts{
		Body: optional.NewInterface(env),
	})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	// only resource id needs to be set since api from sdk doesn't return any other information during batch link clusters
	readEnvironmentClustersMappingLinkedCluster(d, resp.Data)
	return nil
}

func buildEnvironmentClustersMappingCluster(d *schema.ResourceData) *nextgen.ClusterBatchRequest {
	return &nextgen.ClusterBatchRequest{
		OrgIdentifier:     d.Get("org_id").(string),
		ProjectIdentifier: d.Get("project_id").(string),
		EnvRef:            d.Get("env_id").(string),
		Clusters:          ExpandEnvironmentClustersMappingCluster(d.Get("clusters").(*schema.Set).List()),
	}
}

func ExpandEnvironmentClustersMappingCluster(clusterBasicDTO []interface{}) []nextgen.ClusterBasicDto {
	var result []nextgen.ClusterBasicDto
	for _, cluster := range clusterBasicDTO {
		v := cluster.(map[string]interface{})

		var resultcluster nextgen.ClusterBasicDto
		// remove prefix because while deletion tf resource reads from get API
		//which returns like (account.clusterid) but (clusterid) only needs to be sent for delete

		resultcluster.Identifier = removeScopePrefix(v["identifier"].(string))
		resultcluster.AgentIdentifier = v["agent_identifier"].(string)
		resultcluster.Name = v["name"].(string)
		resultcluster.Scope = v["scope"].(string)
		result = append(result, resultcluster)
	}
	return result
}

func readEnvironmentClustersMapping(d *schema.ResourceData, clusters *[]nextgen.ClusterResponse) {
	d.SetId(d.Get("org_id").(string) + d.Get("project_id").(string) + d.Get("env_id").(string))
	if clusterSet := flattenClusters(*clusters); len(clusterSet) > 0 {
		d.Set("clusters", clusterSet)
	}
}

func flattenClusters(clusters []nextgen.ClusterResponse) []map[string]interface{} {
	if len(clusters) == 0 {
		return make([]map[string]interface{}, 0)
	}
	var results = make([]map[string]interface{}, len(clusters))
	for i, cluster := range clusters {
		results[i] = map[string]interface{}{
			"identifier":       cluster.ClusterRef,
			"name":             cluster.Name,
			"agent_identifier": cluster.AgentIdentifier,
			"scope":            cluster.Scope,
		}
	}

	return results

}

func readEnvironmentClustersMappingLinkedCluster(d *schema.ResourceData, cl *nextgen.ClusterBatchResponse) {
	d.SetId(d.Get("org_id").(string) + d.Get("project_id").(string) + d.Get("env_id").(string))
}

func removeScopePrefix(identifier string) string {
	return identifier[strings.Index(identifier, ".")+1:]
}
