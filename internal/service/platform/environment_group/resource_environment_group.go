package environment_group

import (
	"context"
	"errors"
	"net/http"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceEnvironmentGroup() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for creating a Harness environment group.",

		ReadContext:   resourceEnvironmentGroupRead,
		UpdateContext: resourceEnvironmentGroupCreateOrUpdate,
		DeleteContext: resourceEnvironmentGroupDelete,
		CreateContext: resourceEnvironmentGroupCreateOrUpdate,
		Importer:      helpers.ProjectResourceImporter,

		Schema: map[string]*schema.Schema{
			"identifier": {
				Description: "identifier of the environment group.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"org_id": {
				Description: "org_id of the environment group.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"project_id": {
				Description: "project_id of the environment group.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"color": {
				Description: "Color of the environment group.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"yaml": {
				Description: "Env group YAML",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
	return resource
}

func resourceEnvironmentGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	var err error
	var envGroup *nextgen.EnvironmentGroupResponse
	var httpResp *http.Response

	id := d.Get("identifier").(string)

	if id != "" {
		var resp nextgen.ResponseDtoEnvironmentGroup

		orgIdentifier := (d.Get("org_id").(string))
		projectIdentifier := (d.Get("project_id").(string))

		resp, httpResp, err = c.EnvironmentGroupApi.GetEnvironmentGroup(ctx, d.Get("identifier").(string), c.AccountId, orgIdentifier, projectIdentifier, &nextgen.EnvironmentGroupApiGetEnvironmentGroupOpts{
			Branch:         helpers.BuildField(d, "branch"),
			RepoIdentifier: helpers.BuildField(d, "repo_id"),
		})
		envGroup = resp.Data.EnvGroup
	} else {
		return diag.FromErr(errors.New("identifier must be specified"))
	}

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	// Soft delete lookup error handling
	// https://harness.atlassian.net/browse/PL-23765
	if envGroup == nil {
		return nil
	}

	readEnvironmentGroup(d, envGroup)

	return nil
}

func resourceEnvironmentGroupCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	var err error
	var resp nextgen.ResponseDtoEnvironmentGroup
	var httpResp *http.Response
	id := d.Id()
	env := buildEnvironmentGroup(d)

	if id == "" {
		resp, httpResp, err = c.EnvironmentGroupApi.PostEnvironmentGroup(ctx, c.AccountId, &nextgen.EnvironmentGroupApiPostEnvironmentGroupOpts{
			Body: optional.NewInterface(env),
		})
	} else {

		resp, httpResp, err = c.EnvironmentGroupApi.UpdateEnvironmentGroup(ctx, c.AccountId, id, &nextgen.EnvironmentGroupApiUpdateEnvironmentGroupOpts{
			Body: optional.NewInterface(env),
		})
	}

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	readEnvironmentGroup(d, resp.Data.EnvGroup)

	return nil
}

func resourceEnvironmentGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	orgIdentifier := (d.Get("org_id").(string))
	projectIdentifier := (d.Get("project_id").(string))

	_, httpResp, err := c.EnvironmentGroupApi.DeleteEnvironmentGroup(ctx, d.Id(), c.AccountId, orgIdentifier, projectIdentifier, &nextgen.EnvironmentGroupApiDeleteEnvironmentGroupOpts{
		Branch:         helpers.BuildField(d, "branch"),
		RepoIdentifier: helpers.BuildField(d, "repo_id"),
	})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	return nil
}

func buildEnvironmentGroup(d *schema.ResourceData) *nextgen.EnvironmentGroupRequest {
	return &nextgen.EnvironmentGroupRequest{
		Identifier:        d.Get("identifier").(string),
		OrgIdentifier:     d.Get("org_id").(string),
		ProjectIdentifier: d.Get("project_id").(string),
		Color:             d.Get("color").(string),
		Yaml:              d.Get("yaml").(string),
	}
}

func readEnvironmentGroup(d *schema.ResourceData, env *nextgen.EnvironmentGroupResponse) {
	d.SetId(env.Identifier)
	d.Set("org_id", env.OrgIdentifier)
	d.Set("project_id", env.ProjectIdentifier)
	d.Set("identifier", env.Identifier)
	d.Set("color", env.Color)
}
