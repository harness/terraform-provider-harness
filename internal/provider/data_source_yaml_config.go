package provider

import (
	"context"

	"github.com/harness-io/harness-go-sdk/harness/api"
	"github.com/harness-io/harness-go-sdk/harness/api/cac"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceYamlConfig() *schema.Resource {
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
	c := meta.(*api.Client)

	app_id := d.Get("app_id").(string)
	path := cac.YamlPath(d.Get("path").(string))

	entity, err := c.ConfigAsCode().FindYamlByPath(app_id, path)
	if err != nil {
		return diag.FromErr(err)
	} else if entity == nil {
		return diag.Errorf("yaml config not found at %s", path)
	}

	return readYamlConfig(d, entity)
}
