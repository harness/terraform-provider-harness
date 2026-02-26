package as_rule

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"regexp"
	"strings"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceRules() *schema.Resource {
	resource := &schema.Resource{
		Description: "Datasource for listing Harness AutoStopping rules.",
		ReadContext: readAutoStopRules,
		Schema: map[string]*schema.Schema{
			"kind": {
				Description: "Return rules matching provided kind.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"name": {
				Description: "Filter by rule name using a regular expression. e.g. \"^myname-.*\" or \"^(app|svc).*\".",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"rules": {
				Description: "List of AutoStopping rules.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"identifier": {
							Description: "Unique identifier of the rule.",
							Type:        schema.TypeFloat,
							Computed:    true,
						},
						"name": {
							Description: "Name of the rule.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"kind": {
							Description: "Kind of the rule.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"cloud_connector_id": {
							Description: "Id of the cloud connector.",
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}

	return resource
}

func readAutoStopRules(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	resp, httpResp, err := c.CloudCostAutoStoppingRulesApi.ListAutoStoppingRules(ctx, c.AccountId, c.AccountId)
	if err != nil {
		return helpers.HandleReadApiError(err, d, httpResp)
	}

	kindFilter := ""
	if v, ok := d.GetOk("kind"); ok {
		kindFilter = strings.TrimSpace(v.(string))
		if !validKinds[kindFilter] {
			return diag.Errorf("invalid kind %q: must be one of %v", kindFilter, validKindKeys())
		}
	}

	nameFilter := ""
	if v, ok := d.GetOk("name"); ok {
		nameFilter = strings.TrimSpace(v.(string))
	}

	rules := make([]interface{}, 0)
	for _, rule := range resp.Response {
		if !matchesKind(rule.Kind, kindFilter) {
			continue
		}
		ok, err := matchesName(rule.Name, nameFilter)
		if err != nil {
			return diag.FromErr(err)
		}
		if !ok {
			continue
		}
		rules = append(rules, flattenService(rule))
	}

	if err := d.Set("rules", rules); err != nil {
		return diag.FromErr(err)
	}

	h := sha256.Sum256([]byte("kind=" + kindFilter + ";name=" + nameFilter))
	d.SetId(hex.EncodeToString(h[:]))
	return nil
}

func flattenService(s nextgen.Service) map[string]interface{} {
	return map[string]interface{}{
		"identifier":         float64(s.Id),
		"name":               s.Name,
		"kind":               s.Kind,
		"cloud_connector_id": s.CloudAccountId,
	}
}

var validKinds = map[string]bool{
	"instance":   true,
	"k8s":        true,
	"containers": true,
	"database":   true,
	"clusters":   true,
}

func matchesKind(ruleKind string, filter string) bool {
	if filter == "" {
		return true
	}
	return ruleKind == filter
}

func validKindKeys() []string {
	keys := make([]string, 0, len(validKinds))
	for k := range validKinds {
		keys = append(keys, k)
	}
	return keys
}

func matchesName(ruleName string, pattern string) (bool, error) {
	if pattern == "" {
		return true, nil
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		return false, fmt.Errorf("invalid name filter regex %q: %w", pattern, err)
	}
	return re.MatchString(ruleName), nil
}
