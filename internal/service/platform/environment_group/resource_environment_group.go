package environment_group

import (
	"context"

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
			"color": {
				Description: "Color of the environment group.",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}

	helpers.SetMultiLevelResourceSchemaForEnvGroup(resource.Schema)

	return resource
}

func resourceEnvironmentGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	orgIdentifier :=     (d.Get("org_id").(string))
	projectIdentifier := (d.Get("project_id").(string))

	resp, _, err := c.EnvironmentGroupApi.GetEnvironmentGroup(ctx, d.Id(), c.AccountId, orgIdentifier, projectIdentifier, &nextgen.EnvironmentGroupApiGetEnvironmentGroupOpts{
		Branch:     helpers.BuildField(d, "brach"),
		RepoIdentifier: helpers.BuildField(d, "repo_id"),
	})

	if err != nil {
		return helpers.HandleApiError(err, d)
	}

	// Soft delete lookup error handling
	// https://harness.atlassian.net/browse/PL-23765
	if resp.Data == nil || resp.Data.EnvGroup == nil {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	readEnvironmentGroup(d, resp.Data.EnvGroup)

	return nil
}

func resourceEnvironmentGroupCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	var err error
	var resp nextgen.ResponseDtoEnvironmentGroup
	id := d.Id()
	env := buildEnvironmentGroup(d)

	if id == "" {
		resp, _, err = c.EnvironmentGroupApi.PostEnvironmentGroup(ctx, c.AccountId, &nextgen.EnvironmentGroupApiPostEnvironmentGroupOpts{
			Body: optional.NewInterface(env),
		})
	} else {
		
		resp, _, err = c.EnvironmentGroupApi.UpdateEnvironmentGroup(ctx, d.Id(), c.AccountId, &nextgen.EnvironmentGroupApiUpdateEnvironmentGroupOpts{
			Body: optional.NewInterface(env),
		})
	}

	if err != nil {
		return helpers.HandleApiError(err, d)
	}

	readEnvironmentGroup(d, resp.Data.EnvGroup)

	return nil
}

func resourceEnvironmentGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	orgIdentifier :=     (d.Get("org_id").(string))
	projectIdentifier := (d.Get("project_id").(string))

	_, _, err := c.EnvironmentGroupApi.DeleteEnvironmentGroup(ctx, d.Id(), c.AccountId, orgIdentifier, projectIdentifier, &nextgen.EnvironmentGroupApiDeleteEnvironmentGroupOpts{
		Branch:     helpers.BuildField(d, "brach"),
		RepoIdentifier: helpers.BuildField(d, "repo_id"),
	})

	if err != nil {
		return diag.Errorf(err.(nextgen.GenericSwaggerError).Error())
	}

	return nil
}

func buildEnvironmentGroup(d *schema.ResourceData) *nextgen.EnvironmentGroupRequest {
	return &nextgen.EnvironmentGroupRequest{
		Identifier:        d.Get("identifier").(string),
		OrgIdentifier:     d.Get("org_id").(string),
		ProjectIdentifier: d.Get("project_id").(string),
		Color:             d.Get("color").(string),
		Yaml:             d.Get("yaml").(string),
	}
}

func readEnvironmentGroup(d *schema.ResourceData, env *nextgen.EnvironmentGroupResponse) {
	d.SetId(env.Identifier)
	d.Set("identifier", env.Identifier)
	d.Set("org_id", env.OrgIdentifier)
	d.Set("name", env.Name)
	d.Set("color", env.Color)
	d.Set("description", env.Description)
	d.Set("tags", helpers.FlattenTags(env.Tags))
}
