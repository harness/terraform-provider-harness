package cd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/harness-io/harness-go-sdk/harness/api"
	"github.com/harness-io/harness-go-sdk/harness/api/cac"
	"github.com/harness-io/harness-go-sdk/harness/api/graphql"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func usageScopeSchema() *schema.Schema {
	return &schema.Schema{
		Description: "This block is used for scoping the resource to a specific set of applications or environments.",
		Type:        schema.TypeSet,
		Optional:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"application_id": {
					Description: "Id of the application to scope to. If empty then this scope applies to all applications.",
					Type:        schema.TypeString,
					Optional:    true,
				},
				"environment_id": {
					Description: "Id of the id of the specific environment to scope to. Cannot be used with `environment_filter_type`.",
					Type:        schema.TypeString,
					Optional:    true,
					// ConflictsWith: []string{"usage_scope.0.environment_filter_type"},
					// ExactlyOneOf: []string{"usage_scope.0.environment_id", "usage_scope.0.environment_filter_type"},
				},
				"environment_filter_type": {
					Description: fmt.Sprintf("Type of environment filter applied. Cannot be used with `environment_id`. Valid options are %s.", strings.Join(graphql.EnvFiltersSlice, ", ")),
					Type:        schema.TypeString,
					Optional:    true,
					// ConflictsWith: []string{"usage_scope.0.environment_id"},
					// ExactlyOneOf: []string{"usage_scope.0.environment_id", "usage_scope.0.environment_filter_type"},
				},
			},
		},
	}
}

func expandUsageScope(d []interface{}) (*graphql.UsageScope, error) {

	us := &graphql.UsageScope{}
	scopes := make([]*graphql.AppEnvScope, 0)

	for _, appScope := range d {
		scopeData := appScope.(map[string]interface{})
		scope := &graphql.AppEnvScope{
			Application: &graphql.AppScopeFilter{},
			Environment: &graphql.EnvScopeFilter{},
		}

		if ok, err := validateUsageScopeSettings(scopeData); !ok {
			return nil, err
		}

		if attr, ok := scopeData["application_id"]; ok && attr != "" {
			scope.Application.AppId = attr.(string)
		} else {
			scope.Application.FilterType = graphql.ApplicationFilterTypes.All
		}

		if attr, ok := scopeData["environment_filter_type"]; ok && attr != "" {
			scope.Environment.FilterType = graphql.EnvironmentFilterType(attr.(string))
		}

		if attr, ok := scopeData["environment_id"]; ok && attr != "" {
			scope.Environment.EnvId = attr.(string)
		}

		scopes = append(scopes, scope)
	}

	us.AppEnvScopes = scopes

	return us, nil
}

func validateUsageScopeSettings(scopeData map[string]interface{}) (bool, error) {
	if scopeData["environment_id"] != "" && scopeData["application_id"] == "" {
		return false, errors.New("`application_id` must be set when using `environment_id`")
	}

	if scopeData["environment_filter_type"] != "" && scopeData["environment_id"] != "" {
		return false, fmt.Errorf("only one of environment_filter_type or environment_id must be set")
	}

	if scopeData["environment_filter_type"] == "" && scopeData["environment_id"] == "" {
		return false, fmt.Errorf("at least one of environment_filter_type or environment_id must be set")
	}

	return true, nil
}

func expandUsageRestrictions(c *api.Client, d []interface{}, ur *cac.UsageRestrictions) error {
	if len(d) == 0 {
		return nil
	}

	restrictions := make([]*cac.AppEnvRestriction, 0)

	for _, appScope := range d {
		scopeData := appScope.(map[string]interface{})
		scope := &cac.AppEnvRestriction{
			AppFilter: &cac.AppFilter{},
			EnvFilter: &cac.EnvFilter{},
		}

		if ok, err := validateUsageScopeSettings(scopeData); !ok {
			return err
		}

		var app *graphql.Application
		var err error

		if attr, ok := scopeData["application_id"]; ok && attr != "" {
			app, err = c.Applications().GetApplicationById(attr.(string))
			if err != nil {
				return nil
			}

			if app == nil {
				return errors.New("application not found")
			}

			scope.AppFilter.EntityNames = []string{app.Name}
			scope.AppFilter.FilterType = cac.ApplicationFilterTypes.Selected
		} else {
			scope.AppFilter.FilterType = cac.ApplicationFilterTypes.All
		}

		if attr, ok := scopeData["environment_id"]; ok && attr != "" {
			env, err := c.ConfigAsCode().GetEnvironmentById(app.Id, attr.(string))
			if err != nil {
				return err
			} else if env.IsEmpty() {
				return errors.New("environment not found")
			}
			scope.EnvFilter.EntityNames = []string{env.Name}
			scope.EnvFilter.FilterTypes = []cac.EnvironmentFilterType{cac.EnvironmentFilterTypes.Selected}
		}

		if attr, ok := scopeData["environment_filter_type"]; ok && attr != "" {
			switch graphql.EnvironmentFilterType(attr.(string)) {
			case graphql.EnvironmentFilterTypes.NonProduction:
				scope.EnvFilter.FilterTypes = []cac.EnvironmentFilterType{cac.EnvironmentFilterTypes.NonProd}
			case graphql.EnvironmentFilterTypes.Production:
				scope.EnvFilter.FilterTypes = []cac.EnvironmentFilterType{cac.EnvironmentFilterTypes.Prod}
			default:
				return errors.New("could not parse environment_filter_type '" + attr.(string) + "'")
			}
		}

		restrictions = append(restrictions, scope)
	}

	if len(restrictions) > 0 {
		ur.AppEnvRestrictions = restrictions
	}
	return nil
}

func flattenUsageScope(uc *graphql.UsageScope) []map[string]interface{} {
	if uc == nil {
		return make([]map[string]interface{}, 0)
	}

	results := make([]map[string]interface{}, len(uc.AppEnvScopes))

	for i, scope := range uc.AppEnvScopes {
		results[i] = map[string]interface{}{
			"application_id":          scope.Application.AppId,
			"environment_id":          scope.Environment.EnvId,
			"environment_filter_type": scope.Environment.FilterType,
		}
	}

	return results
}

func flattenUsageRestrictions(c *api.Client, ur *cac.UsageRestrictions) ([]map[string]interface{}, error) {
	if ur == nil {
		return make([]map[string]interface{}, 0), nil
	}

	results := make([]map[string]interface{}, len(ur.AppEnvRestrictions))

	for i, scope := range ur.AppEnvRestrictions {
		appId, err := flattenAppFilterEntityName(c, scope.AppFilter)
		if err != nil {
			return nil, err
		}

		envId, err := flattenEnvFilterEntityName(c, scope.EnvFilter, appId)
		if err != nil {
			return nil, err
		}

		results[i] = map[string]interface{}{
			"application_id":          appId,
			"environment_id":          envId,
			"environment_filter_type": flattenEnvFilterTypes(scope.EnvFilter),
		}
	}

	return results, nil
}

func flattenAppFilterEntityName(c *api.Client, filter *cac.AppFilter) (string, error) {
	if len(filter.EntityNames) == 0 {
		return "", nil
	}

	name := filter.EntityNames[0]
	app, err := c.Applications().GetApplicationByName(name)
	if err != nil {
		return "", err
	}
	return app.Id, nil
}

// func flattenAppFilterType(filter *cac.AppFilter) string {
// 	switch filter.FilterType {
// 	case cac.ApplicationFilterTypes.All:
// 		return string(graphql.ApplicationFilterTypes.All)
// 	default:
// 		return ""
// 	}
// }

func flattenEnvFilterEntityName(c *api.Client, filter *cac.EnvFilter, applicationId string) (string, error) {
	if len(filter.EntityNames) == 0 {
		return "", nil
	}

	name := filter.EntityNames[0]
	env, err := c.ConfigAsCode().GetEnvironmentByName(applicationId, name)
	if err != nil {
		return "", err
	}
	return env.Id, nil
}

func flattenEnvFilterTypes(filter *cac.EnvFilter) string {
	switch filter.FilterTypes[0] {
	case cac.EnvironmentFilterTypes.Prod:
		return string(graphql.EnvironmentFilterTypes.Production)
	case cac.EnvironmentFilterTypes.NonProd:
		return string(graphql.EnvironmentFilterTypes.NonProduction)
	case cac.EnvironmentFilterTypes.Selected:
		return ""
	default:
		panic("Unknown environment filter type")
	}
}
