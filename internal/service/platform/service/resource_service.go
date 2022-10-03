package service

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

func ResourceService() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for creating a Harness project.",

		ReadContext:   resourceServiceRead,
		UpdateContext: resourceServiceCreateOrUpdate,
		DeleteContext: resourceServiceDelete,
		CreateContext: resourceServiceCreateOrUpdate,
		Importer:      helpers.ProjectResourceImporter,

		Schema: map[string]*schema.Schema{
			"yaml": {
				Description: "Input Set YAML",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}

	helpers.SetProjectLevelResourceSchema(resource.Schema)

	return resource
}

func resourceServiceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Id()

	resp, httpResp, err := c.ServicesApi.GetServiceV2(ctx, id, c.AccountId, &nextgen.ServicesApiGetServiceV2Opts{
		OrgIdentifier:     optional.NewString(d.Get("org_id").(string)),
		ProjectIdentifier: optional.NewString(d.Get("project_id").(string)),
	})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	if resp.Data == nil || resp.Data.Service == nil {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	readService(d, resp.Data.Service)

	return nil
}

func resourceServiceCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	var err error
	var resp nextgen.ResponseDtoServiceResponse
	var httpResp *http.Response
	id := d.Id()
	svc := buildService(d)

	if id == "" {
		resp, httpResp, err = c.ServicesApi.CreateServiceV2(ctx, c.AccountId, &nextgen.ServicesApiCreateServiceV2Opts{
			Body: optional.NewInterface(svc),
		})
	} else {
		resp, httpResp, err = c.ServicesApi.UpdateServiceV2(ctx, c.AccountId, &nextgen.ServicesApiUpdateServiceV2Opts{
			Body: optional.NewInterface(svc),
		})
	}

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	readService(d, resp.Data.Service)

	return nil
}

func resourceServiceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	_, httpResp, err := c.ProjectApi.DeleteProject(ctx, d.Id(), c.AccountId, &nextgen.ProjectApiDeleteProjectOpts{OrgIdentifier: optional.NewString(d.Get("org_id").(string))})
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	return nil
}

func buildService(d *schema.ResourceData) *nextgen.ServiceRequest {
	return &nextgen.ServiceRequest{
		Identifier:        d.Get("identifier").(string),
		OrgIdentifier:     d.Get("org_id").(string),
		ProjectIdentifier: d.Get("project_id").(string),
		Name:              d.Get("name").(string),
		Description:       d.Get("description").(string),
		Tags:              helpers.ExpandTags(d.Get("tags").(*schema.Set).List()),
		Yaml:              d.Get("yaml").(string),
	}
}

func readService(d *schema.ResourceData, project *nextgen.ServiceResponseDetails) {
	d.SetId(project.Identifier)
	d.Set("identifier", project.Identifier)
	d.Set("org_id", project.OrgIdentifier)
	d.Set("project_id", project.ProjectIdentifier)
	d.Set("name", project.Name)
	d.Set("description", project.Description)
	d.Set("tags", helpers.FlattenTags(project.Tags))
	d.Set("yaml", project.Yaml)
}
