package provider

import (
	"context"

	"github.com/harness-io/harness-go-sdk/harness/api"
	"github.com/harness-io/harness-go-sdk/harness/api/cac"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceHelmService() *schema.Resource {
	return &schema.Resource{
		Description:   configAsCodeDescription("Resource for creating a Kubernetes Helm service."),
		CreateContext: resourceHelmServiceCreateOrUpdate,
		ReadContext:   resourceHelmServiceRead,
		UpdateContext: resourceHelmServiceCreateOrUpdate,
		DeleteContext: resourceServiceDelete,
		Schema:        commonServiceSchema(),
		Importer: &schema.ResourceImporter{
			State: serviceStateImporter,
		},
	}
}

func resourceHelmServiceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)

	svcId := d.Get("id").(string)
	appId := d.Get("app_id").(string)

	var svc *cac.Service
	var err error

	if svc, err = c.ConfigAsCode().GetServiceById(appId, svcId); err != nil {
		return diag.FromErr(err)
	} else if svc == nil {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	return readServiceHelm(d, svc)
}

func readServiceHelm(d *schema.ResourceData, svc *cac.Service) diag.Diagnostics {
	d.SetId(svc.Id)
	d.Set("name", svc.Name)
	d.Set("app_id", svc.ApplicationId)
	d.Set("description", svc.Description)

	if vars := flattenServiceVariables(svc.ConfigVariables); len(vars) > 0 {
		d.Set("variable", vars)
	}

	return nil
}

func resourceHelmServiceCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)

	var input *cac.Service
	var err error

	if d.IsNewResource() {
		input = cac.NewEntity(cac.ObjectTypes.Service).(*cac.Service)
	} else {
		if input, err = c.ConfigAsCode().GetServiceById(d.Get("app_id").(string), d.Id()); err != nil {
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
	input.DeploymentType = cac.DeploymentTypes.Helm
	input.ApplicationId = d.Get("app_id").(string)
	input.Description = d.Get("description").(string)

	if vars := d.Get("variable"); vars != nil {
		input.ConfigVariables = expandServiceVariables(vars.(*schema.Set).List())
	}

	// Create Service
	newSvc, err := c.ConfigAsCode().UpsertService(input)
	if err != nil {
		return diag.FromErr(err)
	}

	return readServiceHelm(d, newSvc)
}
