package as_rule

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceECSRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return dataSourceRuleReadByIdOrName(ctx, d, meta, ECS)
}

func DataSourceECSRule() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving a Harness AutoStopping rule for ECS services.",
		ReadContext: dataSourceECSRuleRead,
		Schema:      dataSourceRuleSchema(),
	}
	return resource
}
