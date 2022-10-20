package environment

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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceEnvironment() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for creating a Harness environment.",

		ReadContext:   resourceEnvironmentRead,
		UpdateContext: resourceEnvironmentCreateOrUpdate,
		DeleteContext: resourceEnvironmentDelete,
		CreateContext: resourceEnvironmentCreateOrUpdate,
		Importer:      helpers.ProjectResourceImporter,

		Schema: map[string]*schema.Schema{
			"color": {
				Description: "Color of the environment.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"type": {
				Description:  fmt.Sprintf("The type of environment. Valid values are %s", strings.Join(nextgen.EnvironmentTypeValues, ", ")),
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice(nextgen.EnvironmentTypeValues, false),
			},
			"yaml": {
				Description: "Environment YAML",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}

	helpers.SetProjectLevelResourceSchema(resource.Schema)

	return resource
}

func resourceEnvironmentRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	resp, httpResp, err := c.EnvironmentsApi.GetEnvironmentV2(ctx, d.Id(), c.AccountId, &nextgen.EnvironmentsApiGetEnvironmentV2Opts{
		OrgIdentifier:     optional.NewString(d.Get("org_id").(string)),
		ProjectIdentifier: optional.NewString(d.Get("project_id").(string)),
	})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	// Soft delete lookup error handling
	// https://harness.atlassian.net/browse/PL-23765
	if resp.Data == nil || resp.Data.Environment == nil {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	readEnvironment(d, resp.Data.Environment)

	return nil
}

func resourceEnvironmentCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	var err error
	var resp nextgen.ResponseDtoEnvironmentResponse
	var httpResp *http.Response
	id := d.Id()
	env := buildEnvironment(d)

	if id == "" {
		resp, httpResp, err = c.EnvironmentsApi.UpsertEnvironmentV2(ctx, c.AccountId, &nextgen.EnvironmentsApiUpsertEnvironmentV2Opts{
			Body: optional.NewInterface(env),
		})
	} else {
		resp, httpResp, err = c.EnvironmentsApi.UpdateEnvironmentV2(ctx, c.AccountId, &nextgen.EnvironmentsApiUpdateEnvironmentV2Opts{
			Body: optional.NewInterface(env),
		})
	}

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	readEnvironment(d, resp.Data.Environment)

	return nil
}

func resourceEnvironmentDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	_, httpResp, err := c.EnvironmentsApi.DeleteEnvironmentV2(ctx, d.Id(), c.AccountId, &nextgen.EnvironmentsApiDeleteEnvironmentV2Opts{
		OrgIdentifier:     optional.NewString(d.Get("org_id").(string)),
		ProjectIdentifier: optional.NewString(d.Get("project_id").(string)),
	})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	return nil
}

func buildEnvironment(d *schema.ResourceData) *nextgen.EnvironmentRequest {
	return &nextgen.EnvironmentRequest{
		Identifier:        d.Get("identifier").(string),
		OrgIdentifier:     d.Get("org_id").(string),
		ProjectIdentifier: d.Get("project_id").(string),
		Name:              d.Get("name").(string),
		Color:             d.Get("color").(string),
		Description:       d.Get("description").(string),
		Tags:              helpers.ExpandTags(d.Get("tags").(*schema.Set).List()),
		Type_:             d.Get("type").(string),
		Yaml:              d.Get("yaml").(string),
	}
}

func readEnvironment(d *schema.ResourceData, env *nextgen.EnvironmentResponseDetails) {
	d.SetId(env.Identifier)
	d.Set("identifier", env.Identifier)
	d.Set("org_id", env.OrgIdentifier)
	d.Set("name", env.Name)
	d.Set("color", env.Color)
	d.Set("description", env.Description)
	d.Set("tags", helpers.FlattenTags(env.Tags))
	d.Set("type", env.Type_.String())
	if d.Get("yaml").(string) != "" {
		d.Set("yaml", env.Yaml)
	}
}
