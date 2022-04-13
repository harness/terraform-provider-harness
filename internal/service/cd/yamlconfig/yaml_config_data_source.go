package yamlconfig

import (
	"context"

	"github.com/harness/harness-go-sdk/harness/cd/cac"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceYamlConfig() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for retrieving a yaml config.",

		ReadContext: dataSourceYamlConfigRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Description: "Unique identifier of the yaml resource.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"name": {
				Description: "Name of the yaml resource.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"path": {
				Description: "Path to the yaml file.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"app_id": {
				Description: "Unique identifier of the application. This is not required for account level resources (i.e. cloud providers, connectors, etc.).",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"content": {
				Description: "Content of the yaml file.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func dataSourceYamlConfigRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*internal.Session)

	app_id := d.Get("app_id").(string)
	path := cac.YamlPath(d.Get("path").(string))

	entity, err := c.CDClient.ConfigAsCodeClient.FindYamlByPath(app_id, path)
	if err != nil {
		return diag.FromErr(err)
	} else if entity == nil {
		return diag.Errorf("yaml config not found at %s", path)
	}

	return readYamlConfig(d, entity)
}
