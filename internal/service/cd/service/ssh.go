package service

import (
	"context"

	sdk "github.com/harness/harness-go-sdk"
	"github.com/harness/harness-go-sdk/harness/cd/cac"
	"github.com/harness/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceSSHService() *schema.Resource {

	sshSchema := commonServiceSchema()
	sshSchema["artifact_type"] = &schema.Schema{
		Description:  "The type of artifact to deploy.",
		Type:         schema.TypeString,
		Required:     true,
		ForceNew:     true,
		ValidateFunc: validation.StringInSlice(cac.SSHArtifactTypes, false),
	}

	return &schema.Resource{
		Description:   utils.ConfigAsCodeDescription("Resource for creating an SSH service."),
		CreateContext: resourceSSHServiceCreateOrUpdate,
		ReadContext:   resourceSSHServiceRead,
		UpdateContext: resourceSSHServiceCreateOrUpdate,
		DeleteContext: resourceServiceDelete,
		Schema:        sshSchema,
		Importer: &schema.ResourceImporter{
			State: serviceStateImporter,
		},
	}
}

func resourceSSHServiceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	return readServiceSSH(d, svc)
}

func readServiceSSH(d *schema.ResourceData, svc *cac.Service) diag.Diagnostics {
	d.SetId(svc.Id)
	d.Set("name", svc.Name)
	d.Set("app_id", svc.ApplicationId)
	d.Set("description", svc.Description)
	d.Set("artifact_type", svc.ArtifactType)

	if vars := flattenServiceVariables(svc.ConfigVariables); len(vars) > 0 {
		d.Set("variable", vars)
	}

	return nil
}

func resourceSSHServiceCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	input.Name = d.Get("name").(string)
	input.ArtifactType = cac.ArtifactType(d.Get("artifact_type").(string))
	input.DeploymentType = cac.DeploymentTypes.SSH
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

	return readServiceSSH(d, newSvc)
}
