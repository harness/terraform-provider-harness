package provider

import (
	"errors"

	"github.com/harness-io/harness-go-sdk/harness/api"
	"github.com/harness-io/harness-go-sdk/harness/api/cac"
	"github.com/harness-io/harness-go-sdk/harness/api/graphql"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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

func expandUsageScope(d []interface{}) (*graphql.UsageScope, error) {

	us := &graphql.UsageScope{}
	scopes := make([]*graphql.AppEnvScope, 0)

	for _, appScope := range d {
		scopeData := appScope.(map[string]interface{})
		scope := &graphql.AppEnvScope{
			Application: &graphql.AppScopeFilter{},
			Environment: &graphql.EnvScopeFilter{},
		}

		if attr, ok := scopeData["application_filter_type"]; ok && attr != "" {
			scope.Application.FilterType = graphql.ApplicationFilterType(attr.(string))
		}

		if attr, ok := scopeData["application_id"]; ok && attr != "" {
			scope.Application.AppId = attr.(string)
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

		if attr, ok := scopeData["application_filter_type"]; ok && attr != "" {
			if graphql.ApplicationFilterType(attr.(string)) == graphql.ApplicationFilterTypes.All {
				scope.AppFilter.FilterType = cac.ApplicationFilterTypes.All
			} else {
				scope.AppFilter.FilterType = cac.ApplicationFilterTypes.Selected
			}
		}

		var app *cac.Application
		if attr, ok := scopeData["application_id"]; ok && attr != "" {
			app, err := c.Applications().GetApplicationById(attr.(string))
			if err != nil {
				return nil
			} else if app == nil {
				return errors.New("application not found")
			}
			scope.AppFilter.EntityNames = []string{app.Name}
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

			if attr, ok := scopeData["environment_id"]; ok && attr != "" {
				env, err := c.ConfigAsCode().GetEnvironmentById(app.Id, attr.(string))
				if err != nil {
					return err
				} else if env.IsEmpty() {
					return errors.New("environment not found")
				}
				scope.EnvFilter.EntityNames = []string{env.Name}
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
			"application_filter_type": scope.Application.FilterType,
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
			"application_filter_type": flattenAppFilterType(scope.AppFilter),
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

func flattenAppFilterType(filter *cac.AppFilter) string {
	switch filter.FilterType {
	case cac.ApplicationFilterTypes.All:
		return string(graphql.ApplicationFilterTypes.All)
	default:
		return ""
	}
}

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
