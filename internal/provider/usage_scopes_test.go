package provider

const testAccDefaultUsageScope = `
	usage_scope {
		environment_filter_type = "NON_PRODUCTION_ENVIRONMENTS"
	}
`

// var testAccDefaultUsageScopeObj = &client.UsageScope{
// 	AppEnvScopes: []*client.AppEnvScope{
// 		{
// 			Application: &client.AppScopeFilter{
// 				FilterType: client.ApplicationFilterTypes.All,
// 			},
// 			Environment: &client.EnvScopeFilter{
// 				FilterType: client.EnvironmentFilterTypes.NonProduction,
// 			},
// 		},
// 	},
// }
