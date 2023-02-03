package cluster

import (
	"context"
	"net/http"

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
		DeleteContext: resourceEnvironmentClustersMappingClusterUnlink,
		CreateContext: resourceEnvironmentClustersMappingClusterLink,
		UpdateContext: resourceEnvironmentClustersMappingClusterLink,
		Importer:      helpers.ProjectResourceImporter,

		Schema: map[string]*schema.Schema{
			"identifier": {
				Description: "identifier of the cluster.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"env_id": {
				Description: "environment identifier.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"org_id": {
				Description: "org_id of the cluster.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"project_id": {
				Description: "project_id of the cluster.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"scope": {
				Description: "scope at which the cluster exists in harness gitops",
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
							Description: "account Identifier of the account",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"name": {
							Description: "name of the cluster",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"scope": {
							Description: "scope at which the cluster exists in harness gitops, project vs org vs account",
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
	resp, httpResp, err := c.ClustersApi.GetClusterList(ctx, c.AccountId, envId, &nextgen.ClustersApiGetClusterListOpts{
		OrgIdentifier:     optional.NewString(d.Get("org_id").(string)),
		ProjectIdentifier: optional.NewString(d.Get("project_id").(string)),
	})

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

	readEnvironmentClustersMappingCluster(d, &resp.Data.Content[0], d.Get("org_id").(string), d.Get("project_id").(string))

	return nil
}

func resourceEnvironmentClustersMappingClusterUnlink(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	env := buildEnvironmentClustersMappingCluster(d)

	_, httpResp, err := c.ClustersApi.UnlinkClustersInBatch(ctx, c.AccountId, &nextgen.ClustersApiUnlinkClustersInBatchOpts{
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
	// make read call
	var resptwo nextgen.ResponseDtoPageResponseClusterResponse

	resptwo, httpResp, err = c.ClustersApi.GetClusterList(ctx, c.AccountId, d.Get("env_id").(string), &nextgen.ClustersApiGetClusterListOpts{
		OrgIdentifier:     optional.NewString(d.Get("org_id").(string)),
		ProjectIdentifier: optional.NewString(d.Get("project_id").(string)),
	})

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

	readEnvironmentClustersMappingCluster(d, &resptwo.Data.Content[0], d.Get("org_id").(string), d.Get("project_id").(string))
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
		resultcluster.Identifier = v["identifier"].(string)
		resultcluster.Name = v["name"].(string)
		resultcluster.Scope = v["scope"].(string)
		result = append(result, resultcluster)
	}
	return result
}

func readEnvironmentClustersMappingCluster(d *schema.ResourceData, cl *nextgen.ClusterResponse, org_id string, project_id string) {
	d.SetId(cl.ClusterRef)
	d.Set("identifier", cl.ClusterRef)
	d.Set("org_id", org_id)
	d.Set("project_id", project_id)
	d.Set("env_id", cl.EnvRef)
	d.Set("scope", cl.Scope)
}
