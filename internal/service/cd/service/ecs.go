package service

import (
	"context"

	"github.com/harness/harness-go-sdk/harness/cd"
	"github.com/harness/harness-go-sdk/harness/cd/cac"
	"github.com/harness/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceECSService() *schema.Resource {
	return &schema.Resource{
		Description:   utils.ConfigAsCodeDescription("Resource for creating an AWS ECS service."),
		CreateContext: resourceECSServiceCreateOrUpdate,
		ReadContext:   resourceECSServiceRead,
		UpdateContext: resourceECSServiceCreateOrUpdate,
		DeleteContext: resourceServiceDelete,
		Schema:        commonServiceSchema(),
		Importer: &schema.ResourceImporter{
			State: serviceStateImporter,
		},
	}
}

func resourceECSServiceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*cd.ApiClient)

	svcId := d.Get("id").(string)
	appId := d.Get("app_id").(string)

	var svc *cac.Service
	var err error

	if svc, err = c.ConfigAsCodeClient.GetServiceById(appId, svcId); err != nil {
		return diag.FromErr(err)
	} else if svc == nil {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	return readServiceEcs(d, svc)
}

func readServiceEcs(d *schema.ResourceData, svc *cac.Service) diag.Diagnostics {
	d.SetId(svc.Id)
	d.Set("name", svc.Name)
	d.Set("app_id", svc.ApplicationId)
	d.Set("description", svc.Description)

	if vars := flattenServiceVariables(svc.ConfigVariables); len(vars) > 0 {
		d.Set("variable", vars)
	}

	return nil
}

func resourceECSServiceCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*cd.ApiClient)

	var input *cac.Service
	var err error

	if d.IsNewResource() {
		input = cac.NewEntity(cac.ObjectTypes.Service).(*cac.Service)
	} else {
		if input, err = c.ConfigAsCodeClient.GetServiceById(d.Get("app_id").(string), d.Id()); err != nil {
			return diag.FromErr(err)
		} else if input == nil {
			d.SetId("")
			d.MarkNewResource()
			return nil
		}
	}

	// Setup the object to be created
	input.Name = d.Get("name").(string)
	input.ArtifactType = cac.ArtifactTypes.Docker
	input.DeploymentType = cac.DeploymentTypes.ECS
	input.ApplicationId = d.Get("app_id").(string)
	input.Description = d.Get("description").(string)

	if vars := d.Get("variable"); vars != nil {
		input.ConfigVariables = expandServiceVariables(vars.(*schema.Set).List())
	}

	// Create Service
	newSvc, err := c.ConfigAsCodeClient.UpsertService(input)
	if err != nil {
		return diag.FromErr(err)
	}

	return readServiceEcs(d, newSvc)
}
