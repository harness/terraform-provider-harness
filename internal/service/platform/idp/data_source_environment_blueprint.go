package idp

import (
	"context"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/idp"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceEnvironmentBlueprint() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving Harness environment blueprints.",
		ReadContext: dataSourceEnvironmentBlueprintRead,
		Schema: map[string]*schema.Schema{
			"identifier": helpers.GetIdentifierSchema(helpers.SchemaFlagTypes.Required),
			"version": {
				Type:        schema.TypeString,
				Description: "Version of the environment blueprint",
				Required:    true,
				ForceNew:    true,
			},
			"yaml": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "YAML definition of the catalog entity",
			},
			"description": {
				Type:        schema.TypeString,
				Description: "Description of the catalog entity",
				Computed:    true,
			},
			"deprecated": {
				Type:        schema.TypeBool,
				Description: "Whether the catalog entity is deprecated",
				Computed:    true,
			},
			"stable": {
				Type:        schema.TypeBool,
				Description: "Whether the catalog entity is stable",
				Computed:    true,
			},
		},
	}

	return resource
}

func dataSourceEnvironmentBlueprintRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetIDPClientWithContext(ctx)

	version := d.Get("version").(string)

	id := d.Id()
	if id == "" {
		id = d.Get("identifier").(string)
	}

	resp, httpResp, err := c.EntitiesApi.GetEntityVersion(ctx, "account", environmentBlueprintKind, id, version, &idp.EntitiesApiGetEntityVersionOpts{
		HarnessAccount: optional.NewString(c.AccountId),
	})
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	readEnvironmentBlueprint(d, resp)

	return nil
}
