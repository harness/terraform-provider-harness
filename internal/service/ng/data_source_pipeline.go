package ng

import (
	"context"

	"github.com/antihax/optional"
	sdk "github.com/harness-io/harness-go-sdk"
	"github.com/harness-io/harness-go-sdk/harness/nextgen"
	"github.com/harness-io/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourcePipeline() *schema.Resource {
	return &schema.Resource{
		Description: utils.GetNextgenDescription("Data source for retrieving a Harness pipeline."),

		ReadContext: dataSourcePipelineRead,

		Schema: map[string]*schema.Schema{
			"identifier": {
				Description: "Unique identifier of the pipeline.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"org_id": {
				Description: "Unique identifier of the organization.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"project_id": {
				Description: "Unique identifier of the project.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"pipeline_yaml": {
				Description: "YAML of the pipeline.",
				Type:        schema.TypeString,
				Required:	 true,
			},
		},
	}
}

func dataSourcePipelineRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*sdk.Session)

	pipeline := buildPipeline(d)

	resp, _, err := c.NGClient.PipelinesApi.GetPipeline(ctx, c.AccountId, pipeline.OrgIdentifier, pipeline.ProjectIdentifier, pipeline.Identifier, &nextgen.PipelinesApiGetPipelineOpts{})
	if err != nil {
		return diag.FromErr(err)
	}

	pipeline.PipelineYAML = resp.Data.YamlPipeline

	readPipeline(d, pipeline)

	return nil
}
