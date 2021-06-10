package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/micahlmartin/terraform-provider-harness/internal/client"
)

func usageScopeSchema() *schema.Schema {
	return &schema.Schema{
		Description: "Usage scopes",
		Type:        schema.TypeSet,
		Optional:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"application_id": {
					Description: "Id of the application scoping",
					Type:        schema.TypeString,
					Optional:    true,
				},
				"application_filter_type": {
					Description: "Type of application filter applied. ALL if not application id supplied, otherwise NULL",
					Type:        schema.TypeString,
					Optional:    true,
				},
				"environment_id": {
					Description: "Id of the environment scoping",
					Type:        schema.TypeString,
					Optional:    true,
				},
				"environment_filter_type": {
					Description: "Type of environment filter applied. ALL if not filter applied",
					Type:        schema.TypeString,
					Optional:    true,
				},
			},
		},
	}
}

func expandUsageScope(d []interface{}) (*client.UsageScope, error) {

	us := &client.UsageScope{}
	scopes := make([]*client.AppEnvScope, 0)

	for _, appScope := range d {
		scopeData := appScope.(map[string]interface{})
		scope := &client.AppEnvScope{
			Application: &client.AppScopeFilter{},
			Environment: &client.EnvScopeFilter{},
		}

		if attr, ok := scopeData["application_filter_type"]; ok && attr != "" {
			scope.Application.FilterType = attr.(string)
		}

		if attr, ok := scopeData["application_id"]; ok && attr != "" {
			scope.Application.AppId = attr.(string)
		}

		if attr, ok := scopeData["environment_filter_type"]; ok && attr != "" {
			scope.Environment.FilterType = attr.(string)
		}

		if attr, ok := scopeData["environment_id"]; ok && attr != "" {
			scope.Environment.EnvId = attr.(string)
		}

		scopes = append(scopes, scope)
	}

	us.AppEnvScopes = scopes

	return us, nil
}

func flattenUsageScope(uc *client.UsageScope) []map[string]interface{} {
	if uc == nil {
		return make([]map[string]interface{}, 0)
	}

	results := make([]map[string]interface{}, len(uc.AppEnvScopes))

	for i, scope := range uc.AppEnvScopes {
		results[i] = map[string]interface{}{
			"application_id":          scope.Application.AppId,
			"application_filter_type": scope.Application.FilterType,
			"environment_id":          scope.Environment.EnvId,
			"environment_filter_type": scope.Environment.FilterType,
		}
	}

	return results
}

// func flattenAppEnvScopes(appEnvScopes []*client.AppEnvScope) []interface{} {
// 	if appEnvScopes == nil {
// 		return make([]interface{}, 0)
// 	}

// 	scopes := make([]interface{}, len(appEnvScopes))

// 	for i, scope := range appEnvScopes {
// 		s := map[string]string{
// 			"application_filter_type": scope.Application.FilterType,
// 			"application_id":          scope.Application.AppId,
// 			"environment_filter_type": scope.Environment.FilterType,
// 			"environment_id":          scope.Environment.EnvId,
// 		}

// 		scopes[i] = s
// 	}

// 	return scopes
// }

// func expandUsageScopeObject(scope interface{}) *client.AppEnvScope {
// 	sc := scope.(map[string]interface{})

// 	opts := &client.AppEnvScope{
// 		Application: &client.AppScopeFilter{},
// 		Environment: &client.EnvScopeFilter{},
// 	}

// 	if attr, ok := sc["application_id"]; ok && attr != "" {
// 		opts.Application.AppId = attr.(string)
// 	}

// 	if attr, ok := sc["application_filter_type"]; ok && attr != "" {
// 		opts.Application.FilterType = attr.(string)
// 	}

// 	if attr, ok := sc["environment_id"]; ok && attr != "" {
// 		opts.Environment.EnvId = attr.(string)
// 	}

// 	if attr, ok := sc["environment_filter_type"]; ok && attr != "" {
// 		opts.Environment.FilterType = attr.(string)
// 	}

// 	return opts
// }
