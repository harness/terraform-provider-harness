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

	func ResourceCluster() *schema.Resource {
		resource := &schema.Resource{
			Description: "Resource for creating a Harness Cluster.",
	
			ReadContext:   resourceClusterRead,
			DeleteContext: resourceClusterDelete,
			CreateContext: resourceClusterLink,
			Importer:      helpers.ProjectResourceImporter,
	
			Schema: map[string]*schema.Schema{
				"identifier": {
					Description: "identifier of the cluster.",
					Type:        schema.TypeString,
					Required:    true,
				},
				"envref": {
					Description: "environment identifier of the cluster.",
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
					Optional:    true,
				},
			},
		}
		return resource
	}

	func resourceClusterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
		c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	
		envId := d.Get("env_id").(string)
		resp, httpResp, err := c.ClustersApi.GetCluster(ctx, d.Id(), c.AccountId, envId, &nextgen.ClustersApiGetClusterOpts{
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
	
		readCluster(d, resp.Data)
	
		return nil
	}
	
	func resourceClusterDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
		c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
		envId := d.Get("env_id").(string)
		_, httpResp, err := c.ClustersApi.DeleteCluster(ctx, d.Id(), c.AccountId, envId, &nextgen.ClustersApiDeleteClusterOpts{
			OrgIdentifier:     optional.NewString(d.Get("org_id").(string)),
			ProjectIdentifier: optional.NewString(d.Get("project_id").(string)),
		})
	
		if err != nil {
			return helpers.HandleApiError(err, d, httpResp)
		}
	
		return nil
	}

	func resourceClusterLink(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
		c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

		var err error
		var resp nextgen.ResponseDtoClusterResponse
		var httpResp *http.Response
		env := buildCluster(d)

		resp, httpResp, err = c.ClustersApi.LinkCluster(ctx, c.AccountId, &nextgen.ClustersApiLinkClusterOpts{
			Body: optional.NewInterface(env),
		})
	
		if err != nil {
			return helpers.HandleApiError(err, d, httpResp)
		}
	
		readCluster(d, resp.Data)
		return nil
	}

	func buildCluster(d *schema.ResourceData) *nextgen.ClusterRequest {
		return &nextgen.ClusterRequest{
			Identifier:        d.Get("identifier").(string),
			OrgIdentifier:     d.Get("org_id").(string),
			ProjectIdentifier: d.Get("project_id").(string),
			EnvRef:             d.Get("envref").(string),
			Scope:             d.Get("scope").(string),
		}
	}

	func readCluster(d *schema.ResourceData, cl *nextgen.ClusterResponse) {
		d.Set("clusterRef", cl.ClusterRef)
		d.Set("org_id", cl.OrgIdentifier)
		d.Set("proj_id", cl.ProjectIdentifier)
		d.Set("account_id", cl.AccountIdentifier)
		d.Set("envref", cl.EnvRef)
		d.Set("scope", cl.Scope)
		d.Set("name", cl.Name)
	}
