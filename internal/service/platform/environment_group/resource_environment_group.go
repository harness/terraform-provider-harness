package environment_group

import (
	"context"
	"fmt"
	"strings"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
			"type": {
				Description:  fmt.Sprintf("The type of environment group. Valid values are %s", strings.Join(nextgen.EnvironmentGroupTypeValues, ", ")),
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice(nextgen.EnvironmentGroupTypeValues, false),
			},
		},
	}

	helpers.SetMultiLevelResourceSchema(resource.Schema)

	return resource
}

func resourceEnvironmentGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	resp, _, err := c.EnvironmentGroupApi.GetEnvironmentGroupV2(ctx, d.Id(), c.AccountId, &nextgen.EnvironmentGroupApiGetEnvironmentGroupV2Opts{
		OrgIdentifier:     optional.NewString(d.Get("org_id").(string)),
		ProjectIdentifier: optional.NewString(d.Get("project_id").(string)),
	})

	if err != nil {
		return helpers.HandleApiError(err, d)
	}

	// Soft delete lookup error handling
	// https://harness.atlassian.net/browse/PL-23765
	if resp.Data == nil || resp.Data.EnvironmentGroup == nil {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	readEnvironmentGroup(d, resp.Data.EnvironmentGroup)

	return nil
}

func resourceEnvironmentGroupCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	var err error
	var resp nextgen.ResponseDtoEnvironmentGroupResponse
	id := d.Id()
	env := buildEnvironmentGroup(d)

	if id == "" {
		resp, _, err = c.EnvironmentGroupApi.CreateEnvironmentGroupV2(ctx, c.AccountId, &nextgen.EnvironmentGroupApiCreateEnvironmentGroupV2Opts{
			Body: optional.NewInterface(env),
		})
	} else {
		resp, _, err = c.EnvironmentGroupApi.UpdateEnvironmentGroupV2(ctx, c.AccountId, &nextgen.EnvironmentGroupApiUpdateEnvironmentGroupV2Opts{
			Body: optional.NewInterface(env),
		})
	}

	if err != nil {
		return helpers.HandleApiError(err, d)
	}

	readEnvironmentGroup(d, resp.Data.EnvironmentGroup)

	return nil
}

func resourceEnvironmentGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	_, _, err := c.EnvironmentGroupApi.DeleteEnvironmentGroupV2(ctx, d.Id(), c.AccountId, &nextgen.EnvironmentGroupApiDeleteEnvironmentGroupV2Opts{
		OrgIdentifier:     optional.NewString(d.Get("org_id").(string)),
		ProjectIdentifier: optional.NewString(d.Get("project_id").(string)),
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
		yaml:             d.Get("yaml").(string),
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
	d.Set("type", env.Type_.String())
}
