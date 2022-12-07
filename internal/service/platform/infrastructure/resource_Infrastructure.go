package infrastructure

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceInfrastructure() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving a Harness Infrastructure.",

		ReadContext:   resourceInfrastructureRead,
		UpdateContext: resourceInfrastructureCreateOrUpdate,
		DeleteContext: resourceInfrastructureDelete,
		CreateContext: resourceInfrastructureCreateOrUpdate,
		Importer:      helpers.EnvRelatedResourceImporter,

		Schema: map[string]*schema.Schema{
			"identifier": {
				Description: "identifier of the Infrastructure.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"env_id": {
				Description: "environment identifier.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"type": {
				Description: fmt.Sprintf("Type of Infrastructure. Valid values are %s.", strings.Join(nextgen.InfrastructureTypeValues, ", ")),
				Type:        schema.TypeString,
				Required:    true,
			},
			"yaml": {
				Description: "Infrastructure YAML",
				Type:        schema.TypeString,
				Required:    true,
			},
			"deployment_type": {
				Description: fmt.Sprintf("Infrastructure deployment type. Valid values are %s.", strings.Join(nextgen.InfrastructureDeploymentypeValues, ", ")),
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
		},
	}
	helpers.SetProjectLevelResourceSchema(resource.Schema)

	return resource
}

func resourceInfrastructureRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	env_id := d.Get("env_id").(string)

	resp, httpResp, err := c.InfrastructuresApi.GetInfrastructure(ctx, d.Id(), c.AccountId, env_id, &nextgen.InfrastructuresApiGetInfrastructureOpts{
		OrgIdentifier:     optional.NewString(d.Get("org_id").(string)),
		ProjectIdentifier: optional.NewString(d.Get("project_id").(string)),
	})

	if err != nil {
		return helpers.HandleReadApiError(err, d, httpResp)
	}

	readInfrastructure(d, resp.Data)

	return nil
}

func resourceInfrastructureCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	var err error
	var resp nextgen.ResponseDtoInfrastructureResponse
	var httpResp *http.Response
	id := d.Id()
	infra := buildInfrastructure(d)

	if id == "" {
		resp, httpResp, err = c.InfrastructuresApi.CreateInfrastructure(ctx, c.AccountId, &nextgen.InfrastructuresApiCreateInfrastructureOpts{
			Body: optional.NewInterface(infra),
		})
	} else {
		resp, httpResp, err = c.InfrastructuresApi.UpdateInfrastructure(ctx, c.AccountId, &nextgen.InfrastructuresApiUpdateInfrastructureOpts{
			Body: optional.NewInterface(infra),
		})
	}

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	readInfrastructure(d, resp.Data)

	return nil
}

func resourceInfrastructureDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	env_id := d.Get("env_id").(string)

	_, httpResp, err := c.InfrastructuresApi.DeleteInfrastructure(ctx, d.Id(), c.AccountId, env_id, &nextgen.InfrastructuresApiDeleteInfrastructureOpts{
		OrgIdentifier:     optional.NewString(d.Get("org_id").(string)),
		ProjectIdentifier: optional.NewString(d.Get("project_id").(string)),
	})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	return nil
}

func buildInfrastructure(d *schema.ResourceData) *nextgen.InfrastructureRequest {
	return &nextgen.InfrastructureRequest{
		Identifier:        d.Get("identifier").(string),
		OrgIdentifier:     d.Get("org_id").(string),
		ProjectIdentifier: d.Get("project_id").(string),
		EnvironmentRef:    d.Get("env_id").(string),
		Name:              d.Get("name").(string),
		Description:       d.Get("description").(string),
		Tags:              helpers.ExpandTags(d.Get("tags").(*schema.Set).List()),
		Type_:             d.Get("type").(string),
		Yaml:              d.Get("yaml").(string),
	}
}

func readInfrastructure(d *schema.ResourceData, infra *nextgen.InfrastructureResponse) {
	d.SetId(infra.Infrastructure.Identifier)
	d.Set("org_id", infra.Infrastructure.OrgIdentifier)
	d.Set("project_id", infra.Infrastructure.ProjectIdentifier)
	d.Set("env_id", infra.Infrastructure.EnvironmentRef)
	d.Set("name", infra.Infrastructure.Name)
	d.Set("description", infra.Infrastructure.Description)
	d.Set("tags", helpers.FlattenTags(infra.Infrastructure.Tags))
	d.Set("type", infra.Infrastructure.Type_)
	d.Set("deployment_type", infra.Infrastructure.DeploymentType)
	d.Set("yaml", infra.Infrastructure.Yaml)
}
