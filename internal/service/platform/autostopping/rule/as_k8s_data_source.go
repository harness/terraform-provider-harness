package as_rule

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceK8sRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return dataSourceRuleReadByIdOrName(ctx, d, meta, K8s)
}

func DataSourceK8sRule() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving a Harness AutoStopping rule for K8s services.",
		ReadContext: dataSourceK8sRuleRead,
		Schema:      dataSourceRuleSchema(),
	}
	return resource
}
