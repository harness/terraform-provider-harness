package service

import (
	"context"

	sdk "github.com/harness-io/harness-go-sdk"
	"github.com/harness-io/harness-go-sdk/harness/cd/cac"
	"github.com/harness-io/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceKubernetesService() *schema.Resource {

	k8sSchema := commonServiceSchema()
	k8sSchema["helm_version"] = &schema.Schema{
		Description:  "The version of Helm to use. Options are `V2` and `V3`. Defaults to 'V2'. Only used when `type` is `KUBERNETES` or `HELM`.",
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: validation.StringInSlice([]string{string(cac.HelmVersions.V2), string(cac.HelmVersions.V3)}, false),
		Default:      cac.HelmVersions.V2,
	}

	return &schema.Resource{
		Description:   utils.ConfigAsCodeDescription("Resource for creating a Kubernetes service."),
		CreateContext: resourceKubernetesServiceCreateOrUpdate,
		ReadContext:   resourceKubernetesServiceKubernetesRead,
		UpdateContext: resourceKubernetesServiceCreateOrUpdate,
		DeleteContext: resourceServiceDelete,
		Schema:        k8sSchema,
		Importer: &schema.ResourceImporter{
			State: serviceStateImporter,
		},
	}
}

func resourceKubernetesServiceKubernetesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*sdk.Session)

	svcId := d.Get("id").(string)
	appId := d.Get("app_id").(string)

	var svc *cac.Service
	var err error

	if svc, err = c.CDClient.ConfigAsCodeClient.GetServiceById(appId, svcId); err != nil {
		return diag.FromErr(err)
	} else if svc == nil {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	return readServiceK8s(d, svc)
}

func readServiceK8s(d *schema.ResourceData, svc *cac.Service) diag.Diagnostics {
	d.SetId(svc.Id)
	d.Set("name", svc.Name)
	d.Set("app_id", svc.ApplicationId)
	d.Set("helm_version", svc.HelmVersion)
	d.Set("description", svc.Description)

	if vars := flattenServiceVariables(svc.ConfigVariables); len(vars) > 0 {
		d.Set("variable", vars)
	}

	return nil
}

func resourceKubernetesServiceCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*sdk.Session)

	var input *cac.Service
	var err error

	if d.IsNewResource() {
		input = cac.NewEntity(cac.ObjectTypes.Service).(*cac.Service)
	} else {
		if input, err = c.CDClient.ConfigAsCodeClient.GetServiceById(d.Get("app_id").(string), d.Id()); err != nil {
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
	input.DeploymentType = cac.DeploymentTypes.Kubernetes
	input.HelmVersion = cac.HelmVersion(d.Get("helm_version").(string))
	input.ApplicationId = d.Get("app_id").(string)
	input.Description = d.Get("description").(string)

	if vars := d.Get("variable"); vars != nil {
		input.ConfigVariables = expandServiceVariables(vars.(*schema.Set).List())
	}

	// Create Service
	newSvc, err := c.CDClient.ConfigAsCodeClient.UpsertService(input)
	if err != nil {
		return diag.FromErr(err)
	}

	return readServiceK8s(d, newSvc)
}
