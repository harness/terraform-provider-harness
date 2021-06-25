package provider

import (
	"context"
	"errors"

	"github.com/harness-io/harness-go-sdk/harness/api"
	"github.com/harness-io/harness-go-sdk/harness/api/cac"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceKubernetesService() *schema.Resource {

	k8sSchema := commonServiceSchema()
	k8sSchema["helm_version"] = &schema.Schema{
		Description:  "The version of Helm to use. Options are `V2` and `V3`. Defaults to 'V2'. Only used when `type` is `KUBERNETES` or `HELM`.",
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: validation.StringInSlice([]string{cac.HelmVersions.V2, cac.HelmVersions.V3}, false),
		Default:      cac.HelmVersions.V2,
	}

	return &schema.Resource{
		Description:   "Resource for creating a Kubernetes service",
		CreateContext: resourceKubernetesServiceCreate,
		ReadContext:   resourceKubernetesServiceKubernetesRead,
		UpdateContext: resourceKubernetesServiceUpdate,
		DeleteContext: resourceServiceDelete,
		Schema:        k8sSchema,
	}
}

func resourceKubernetesServiceKubernetesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)

	svcId := d.Get("id").(string)
	appId := d.Get("app_id").(string)

	svc, err := c.Services().GetServiceById(appId, svcId)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("name", svc.Name)
	d.Set("app_id", svc.ApplicationId)
	d.Set("helm_version", svc.HelmVersion)
	d.Set("description", svc.Description)
	return nil
}

func resourceKubernetesServiceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)

	// Setup the object to be created
	svcInput := &cac.Service{
		Name:           d.Get("name").(string),
		ArtifactType:   cac.ArtifactTypes.Docker,
		DeploymentType: cac.DeploymentTypes.Kubernetes,
		HelmVersion:    d.Get("helm_version").(string),
		ApplicationId:  d.Get("app_id").(string),
		Description:    d.Get("description").(string),
	}

	// Create Service
	newSvc, err := c.Services().UpsertService(svcInput)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(newSvc.Id)

	return nil
}

func resourceKubernetesServiceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)

	if d.HasChange("app_id") {
		return diag.FromErr(errors.New("app_id cannot be changed"))
	}

	// Setup the object to create
	svcInput := &cac.Service{
		Name:           d.Get("name").(string),
		ArtifactType:   cac.ArtifactTypes.Docker,
		DeploymentType: cac.DeploymentTypes.Kubernetes,
		HelmVersion:    d.Get("helm_version").(string),
		ApplicationId:  d.Get("app_id").(string),
		Description:    d.Get("description").(string),
	}

	// Create Service
	newSvc, err := c.Services().UpsertService(svcInput)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(newSvc.Id)

	return nil
}
