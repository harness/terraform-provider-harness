package provider

import (
	"context"

	"github.com/harness-io/harness-go-sdk/harness/api"
	"github.com/harness-io/harness-go-sdk/harness/api/cac"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWinRMService() *schema.Resource {

	sshSchema := commonServiceSchema()
	sshSchema["artifact_type"] = &schema.Schema{
		Description:  "The type of artifact to deploy.",
		Type:         schema.TypeString,
		Required:     true,
		ForceNew:     true,
		ValidateFunc: validation.StringInSlice(cac.WinRMArtifactTypesSlice, false),
	}

	return &schema.Resource{
		Description:   "Resource for creating an WinRM service",
		CreateContext: resourceWinRMServiceCreate,
		ReadContext:   resourceWinRMServiceRead,
		UpdateContext: resourceWinRMServiceUpdate,
		DeleteContext: resourceServiceDelete,
		Schema:        sshSchema,
	}
}

func resourceWinRMServiceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)

	svcId := d.Get("id").(string)
	appId := d.Get("app_id").(string)

	svc, err := c.ConfigAsCode().GetServiceById(appId, svcId)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("name", svc.Name)
	d.Set("app_id", svc.ApplicationId)
	d.Set("description", svc.Description)

	if vars := flattenServiceVariables(svc.ConfigVariables); len(vars) > 0 {
		d.Set("variable", vars)
	}

	return nil
}

func resourceWinRMServiceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)

	// Setup the object to be created
	svcInput := &cac.Service{
		Name:           d.Get("name").(string),
		ArtifactType:   cac.ArtifactType(d.Get("artifact_type").(string)),
		DeploymentType: cac.DeploymentTypes.WinRM,
		ApplicationId:  d.Get("app_id").(string),
		Description:    d.Get("description").(string),
	}

	if vars := d.Get("variable"); vars != nil {
		svcInput.ConfigVariables = expandServiceVariables(vars.(*schema.Set).List())
	}

	// Create Service
	newSvc, err := c.ConfigAsCode().UpsertService(svcInput)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(newSvc.Id)

	return nil
}

func resourceWinRMServiceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)

	// Setup the object to create
	svcInput := &cac.Service{
		Name:           d.Get("name").(string),
		ArtifactType:   cac.ArtifactType(d.Get("artifact_type").(string)),
		DeploymentType: cac.DeploymentTypes.WinRM,
		ApplicationId:  d.Get("app_id").(string),
		Description:    d.Get("description").(string),
	}

	if vars := d.Get("variable"); vars != nil {
		svcInput.ConfigVariables = expandServiceVariables(vars.(*schema.Set).List())
	}

	// Create Service
	newSvc, err := c.ConfigAsCode().UpsertService(svcInput)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(newSvc.Id)

	return nil
}
