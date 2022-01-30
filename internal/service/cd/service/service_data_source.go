package service

import (
	"context"
	"fmt"

	sdk "github.com/harness-io/harness-go-sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceService() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Data source for retrieving a Harness service",

		ReadContext: dataSourceServiceRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Description: "Unique identifier of the application",
				Type:        schema.TypeString,
				Required:    true,
			},
			"app_id": {
				Description: "The id of the application the service belongs to",
				Type:        schema.TypeString,
				Required:    true,
			},
			"name": {
				Description: "The name of the service",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"description": {
				Description: "The application description",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"type": {
				Description: "The type of the deployment",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"artifact_type": {
				Description: "The type of artifact deployed by the service",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"helm_version": {
				Description: "The version of Helm being used by the service.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"template_uri": {
				Description: "The path of the template used for the custom deployment",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"tags": {
				Description: "Tags for the service",
				Type:        schema.TypeMap,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceServiceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	c := meta.(*sdk.Session)

	appId := d.Get("app_id").(string)
	svcId := d.Get("id").(string)

	svc, err := c.CDClient.ConfigAsCodeClient.GetServiceById(appId, svcId)
	if err != nil {
		return diag.FromErr(err)
	}

	if svc == nil {
		return diag.FromErr(fmt.Errorf("could not find service with id '%s'", svcId))
	}

	d.SetId(svcId)
	d.Set("type", svc.DeploymentType)
	d.Set("name", svc.Name)
	d.Set("artifact_type", svc.ArtifactType)
	d.Set("description", svc.Description)
	d.Set("helm_version", svc.HelmVersion)
	d.Set("template_uri", svc.DeploymentTypeTemplateUri)
	d.Set("tags", svc.Tags)

	return nil
}
