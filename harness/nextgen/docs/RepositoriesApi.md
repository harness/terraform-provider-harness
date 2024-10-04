# {{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AgentRepositoryServiceCheckExternalSecretsOperator**](RepositoriesApi.md#AgentRepositoryServiceCheckExternalSecretsOperator) | **Get** /api/v1/agents/{agentIdentifier}/repositories/eso/check | Checks whether External Secrets Operator is installed
[**AgentRepositoryServiceCheckOCIRepoType**](RepositoriesApi.md#AgentRepositoryServiceCheckOCIRepoType) | **Post** /api/v1/agents/{agentIdentifier}/repositories/oci/type | Returns the Repository type of OCI repo
[**AgentRepositoryServiceCreateRepository**](RepositoriesApi.md#AgentRepositoryServiceCreateRepository) | **Post** /api/v1/agents/{agentIdentifier}/repositories | CreateRepository creates a new repository configuration
[**AgentRepositoryServiceDeleteRepository**](RepositoriesApi.md#AgentRepositoryServiceDeleteRepository) | **Delete** /api/v1/agents/{agentIdentifier}/repositories/{identifier} | DeleteRepository deletes a repository from the configuration
[**AgentRepositoryServiceGet**](RepositoriesApi.md#AgentRepositoryServiceGet) | **Get** /api/v1/agents/{agentIdentifier}/repositories/{identifier} | Get returns a repository or its credentials
[**AgentRepositoryServiceGetAppDetails**](RepositoriesApi.md#AgentRepositoryServiceGetAppDetails) | **Get** /api/v1/agents/{agentIdentifier}/repositories/{identifier}/appdetails | GetAppDetails returns application details by given path
[**AgentRepositoryServiceGetAppDetailsByUrl**](RepositoriesApi.md#AgentRepositoryServiceGetAppDetailsByUrl) | **Get** /api/v1/agents/{agentIdentifier}/repositories_byurl/appdetails | GetAppDetailsByUrl returns application details by given path for repositories by URL
[**AgentRepositoryServiceGetByUrl**](RepositoriesApi.md#AgentRepositoryServiceGetByUrl) | **Get** /api/v1/agents/{agentIdentifier}/repository_byurl | Get a repository configured in gitops by URL from Harness
[**AgentRepositoryServiceGetHelmCharts**](RepositoriesApi.md#AgentRepositoryServiceGetHelmCharts) | **Get** /api/v1/agents/{agentIdentifier}/repositories/{identifier}/helmcharts | GetHelmCharts returns list of helm charts in the specified repository
[**AgentRepositoryServiceGetHelmChartsByUrl**](RepositoriesApi.md#AgentRepositoryServiceGetHelmChartsByUrl) | **Get** /api/v1/agents/{agentIdentifier}/repositories_byurl/helmcharts | GetHelmChartsByUrl returns list of helm charts in the repository by URL
[**AgentRepositoryServiceListApps**](RepositoriesApi.md#AgentRepositoryServiceListApps) | **Get** /api/v1/agents/{agentIdentifier}/repositories/{identifier}/apps | ListApps returns list of apps in the repo
[**AgentRepositoryServiceListAppsByUrl**](RepositoriesApi.md#AgentRepositoryServiceListAppsByUrl) | **Get** /api/v1/agents/{agentIdentifier}/repositories_byurl/apps | ListAppsByUrl returns list of apps in the repo
[**AgentRepositoryServiceListESOGenerators**](RepositoriesApi.md#AgentRepositoryServiceListESOGenerators) | **Get** /api/v1/agents/{agentIdentifier}/repositories/eso/generators | Returns a list of ESO generators installed in agent namespace.
[**AgentRepositoryServiceListRefs**](RepositoriesApi.md#AgentRepositoryServiceListRefs) | **Get** /api/v1/agents/{agentIdentifier}/repositories/{identifier}/refs | Returns a list of refs (e.g. branches and tags) in the repo
[**AgentRepositoryServiceListRefsByUrl**](RepositoriesApi.md#AgentRepositoryServiceListRefsByUrl) | **Get** /api/v1/agents/{agentIdentifier}/repositories_byurl/refs | Returns a list of refs (e.g. branches and tags) in the repo by URL
[**AgentRepositoryServiceListRepositories**](RepositoriesApi.md#AgentRepositoryServiceListRepositories) | **Get** /api/v1/agents/{agentIdentifier}/repositories | ListRepositories gets a list of all configured repositories
[**AgentRepositoryServiceUpdateRepository**](RepositoriesApi.md#AgentRepositoryServiceUpdateRepository) | **Put** /api/v1/agents/{agentIdentifier}/repositories/{identifier} | UpdateRepository updates a repository configuration
[**AgentRepositoryServiceValidateAccess**](RepositoriesApi.md#AgentRepositoryServiceValidateAccess) | **Post** /api/v1/agents/{agentIdentifier}/repositories/validate | ValidateAccess gets connection state for a repository
[**RepositoryServiceExists**](RepositoriesApi.md#RepositoryServiceExists) | **Get** /api/v1/repositories/exists | Checks whether a repository with the given name exists
[**RepositoryServiceListRepositories**](RepositoriesApi.md#RepositoryServiceListRepositories) | **Post** /api/v1/repositories | List returns list of Repositories
[**RepositoryServiceListRepositoriesByRepositoryCredentialTemplate**](RepositoriesApi.md#RepositoryServiceListRepositoriesByRepositoryCredentialTemplate) | **Get** /api/v1/repositories | List returns list of Repositories by repository credential template

# **AgentRepositoryServiceCheckExternalSecretsOperator**
> RepositoriesCheckEsoResponse AgentRepositoryServiceCheckExternalSecretsOperator(ctx, agentIdentifier, accountIdentifier, optional)
Checks whether External Secrets Operator is installed

Checks whether External Secrets Operator is installed or not

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentifier** | **string**| Agent identifier for entity. | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***RepositoriesApiAgentRepositoryServiceCheckExternalSecretsOperatorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoriesApiAgentRepositoryServiceCheckExternalSecretsOperatorOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 

### Return type

[**RepositoriesCheckEsoResponse**](repositoriesCheckESOResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentRepositoryServiceCheckOCIRepoType**
> V1OciRepoTypeResponse AgentRepositoryServiceCheckOCIRepoType(ctx, body, agentIdentifier, accountIdentifier, optional)
Returns the Repository type of OCI repo

Returns the Repository type of OCI repo

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1OciRepoType**](V1OciRepoType.md)|  | 
  **agentIdentifier** | **string**| Agent identifier for entity. | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***RepositoriesApiAgentRepositoryServiceCheckOCIRepoTypeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoriesApiAgentRepositoryServiceCheckOCIRepoTypeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orgIdentifier** | **optional.**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity. | 

### Return type

[**V1OciRepoTypeResponse**](v1OCIRepoTypeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentRepositoryServiceCreateRepository**
> Servicev1Repository AgentRepositoryServiceCreateRepository(ctx, body, agentIdentifier, accountIdentifier, optional)
CreateRepository creates a new repository configuration

CreateRepository creates a new repository configuration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RepositoriesRepoCreateRequest**](RepositoriesRepoCreateRequest.md)|  | 
  **agentIdentifier** | **string**| Agent identifier for entity. | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***RepositoriesApiAgentRepositoryServiceCreateRepositoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoriesApiAgentRepositoryServiceCreateRepositoryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orgIdentifier** | **optional.**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity. | 
 **identifier** | **optional.**|  | 
 **repoCredsId** | **optional.**|  | 

### Return type

[**Servicev1Repository**](servicev1Repository.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentRepositoryServiceDeleteRepository**
> RepositoriesRepoResponse AgentRepositoryServiceDeleteRepository(ctx, agentIdentifier, identifier, accountIdentifier, optional)
DeleteRepository deletes a repository from the configuration

DeleteRepository deletes a repository from the configuration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentifier** | **string**| Agent identifier for entity. | 
  **identifier** | **string**|  | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***RepositoriesApiAgentRepositoryServiceDeleteRepositoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoriesApiAgentRepositoryServiceDeleteRepositoryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **queryRepo** | **optional.String**| Repo URL for query. | 
 **queryForceRefresh** | **optional.Bool**| Whether to force a cache refresh on repo&#x27;s connection state. | 
 **queryProject** | **optional.String**| The associated project project. | 
 **forceDelete** | **optional.Bool**|  | 

### Return type

[**RepositoriesRepoResponse**](repositoriesRepoResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentRepositoryServiceGet**
> Servicev1Repository AgentRepositoryServiceGet(ctx, agentIdentifier, identifier, accountIdentifier, optional)
Get returns a repository or its credentials

Get returns a repository or its credentials.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentifier** | **string**| Agent identifier for entity. | 
  **identifier** | **string**|  | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***RepositoriesApiAgentRepositoryServiceGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoriesApiAgentRepositoryServiceGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **queryRepo** | **optional.String**| Repo URL for query. | 
 **queryForceRefresh** | **optional.Bool**| Whether to force a cache refresh on repo&#x27;s connection state. | 
 **queryProject** | **optional.String**| The associated project project. | 

### Return type

[**Servicev1Repository**](servicev1Repository.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentRepositoryServiceGetAppDetails**
> RepositoriesRepoAppDetailsResponse AgentRepositoryServiceGetAppDetails(ctx, agentIdentifier, identifier, accountIdentifier, optional)
GetAppDetails returns application details by given path

GetAppDetails returns application details by given path.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentifier** | **string**| Agent identifier for entity. | 
  **identifier** | **string**|  | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***RepositoriesApiAgentRepositoryServiceGetAppDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoriesApiAgentRepositoryServiceGetAppDetailsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **querySourceRepoURL** | **optional.String**| RepoURL is the URL to the repository (Git or Helm) that contains the application manifests. | 
 **querySourcePath** | **optional.String**| Path is a directory path within the Git repository, and is only valid for applications sourced from Git. | 
 **querySourceTargetRevision** | **optional.String**| TargetRevision defines the revision of the source to sync the application to. In case of Git, this can be commit, tag, or branch. If omitted, will equal to HEAD. In case of Helm, this is a semver tag for the Chart&#x27;s version. | 
 **querySourceHelmValueFiles** | [**optional.Interface of []string**](string.md)| ValuesFiles is a list of Helm value files to use when generating a template. | 
 **querySourceHelmReleaseName** | **optional.String**| ReleaseName is the Helm release name to use. If omitted it will use the application name. | 
 **querySourceHelmValues** | **optional.String**| Values specifies Helm values to be passed to helm template, typically defined as a block. | 
 **querySourceHelmVersion** | **optional.String**| Version is the Helm version to use for templating (either \&quot;2\&quot; or \&quot;3\&quot;). | 
 **querySourceHelmPassCredentials** | **optional.Bool**| PassCredentials pass credentials to all domains (Helm&#x27;s --pass-credentials). | 
 **querySourceKustomizeNamePrefix** | **optional.String**| NamePrefix is a prefix appended to resources for Kustomize apps. | 
 **querySourceKustomizeNameSuffix** | **optional.String**| NameSuffix is a suffix appended to resources for Kustomize apps. | 
 **querySourceKustomizeImages** | [**optional.Interface of []string**](string.md)| Images is a list of Kustomize image override specifications. | 
 **querySourceKustomizeVersion** | **optional.String**| Version controls which version of Kustomize to use for rendering manifests. | 
 **querySourceKustomizeForceCommonLabels** | **optional.Bool**| ForceCommonLabels specifies whether to force applying common labels to resources for Kustomize apps. | 
 **querySourceKustomizeForceCommonAnnotations** | **optional.Bool**| ForceCommonAnnotations specifies whether to force applying common annotations to resources for Kustomize apps. | 
 **querySourceKustomizeNamespace** | **optional.String**| Namespace sets the namespace that Kustomize adds to all resources. | 
 **querySourceKsonnetEnvironment** | **optional.String**| Environment is a ksonnet application environment name. | 
 **querySourceDirectoryRecurse** | **optional.Bool**| Recurse specifies whether to scan a directory recursively for manifests. | 
 **querySourceDirectoryJsonnetLibs** | [**optional.Interface of []string**](string.md)| Additional library search dirs. | 
 **querySourceDirectoryExclude** | **optional.String**| Exclude contains a glob pattern to match paths against that should be explicitly excluded from being used during manifest generation. | 
 **querySourceDirectoryInclude** | **optional.String**| Include contains a glob pattern to match paths against that should be explicitly included during manifest generation. | 
 **querySourcePluginName** | **optional.String**|  | 
 **querySourceChart** | **optional.String**| Chart is a Helm chart name, and must be specified for applications sourced from a Helm repo. | 
 **querySourceRef** | **optional.String**| Ref is reference to another source within sources field. This field will not be used if used with a &#x60;source&#x60; tag. | 
 **queryAppName** | **optional.String**|  | 
 **queryAppProject** | **optional.String**|  | 

### Return type

[**RepositoriesRepoAppDetailsResponse**](repositoriesRepoAppDetailsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentRepositoryServiceGetAppDetailsByUrl**
> RepositoriesRepoAppDetailsResponse AgentRepositoryServiceGetAppDetailsByUrl(ctx, agentIdentifier, accountIdentifier, orgIdentifier, projectIdentifier, querySourceRepoURL, optional)
GetAppDetailsByUrl returns application details by given path for repositories by URL

GetAppDetails returns application details by given path.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentifier** | **string**| Agent identifier for entity. | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the Entity. | 
  **projectIdentifier** | **string**| Project Identifier for the Entity. | 
  **querySourceRepoURL** | **string**| RepoURL is the URL to the repository (Git or Helm) that contains the application manifests. | 
 **optional** | ***RepositoriesApiAgentRepositoryServiceGetAppDetailsByUrlOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoriesApiAgentRepositoryServiceGetAppDetailsByUrlOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **identifier** | **optional.String**|  | 
 **querySourcePath** | **optional.String**| Path is a directory path within the Git repository, and is only valid for applications sourced from Git. | 
 **querySourceTargetRevision** | **optional.String**| TargetRevision defines the revision of the source to sync the application to. In case of Git, this can be commit, tag, or branch. If omitted, will equal to HEAD. In case of Helm, this is a semver tag for the Chart&#x27;s version. | 
 **querySourceHelmValueFiles** | [**optional.Interface of []string**](string.md)| ValuesFiles is a list of Helm value files to use when generating a template. | 
 **querySourceHelmReleaseName** | **optional.String**| ReleaseName is the Helm release name to use. If omitted it will use the application name. | 
 **querySourceHelmValues** | **optional.String**| Values specifies Helm values to be passed to helm template, typically defined as a block. | 
 **querySourceHelmVersion** | **optional.String**| Version is the Helm version to use for templating (either \&quot;2\&quot; or \&quot;3\&quot;). | 
 **querySourceHelmPassCredentials** | **optional.Bool**| PassCredentials pass credentials to all domains (Helm&#x27;s --pass-credentials). | 
 **querySourceKustomizeNamePrefix** | **optional.String**| NamePrefix is a prefix appended to resources for Kustomize apps. | 
 **querySourceKustomizeNameSuffix** | **optional.String**| NameSuffix is a suffix appended to resources for Kustomize apps. | 
 **querySourceKustomizeImages** | [**optional.Interface of []string**](string.md)| Images is a list of Kustomize image override specifications. | 
 **querySourceKustomizeVersion** | **optional.String**| Version controls which version of Kustomize to use for rendering manifests. | 
 **querySourceKustomizeForceCommonLabels** | **optional.Bool**| ForceCommonLabels specifies whether to force applying common labels to resources for Kustomize apps. | 
 **querySourceKustomizeForceCommonAnnotations** | **optional.Bool**| ForceCommonAnnotations specifies whether to force applying common annotations to resources for Kustomize apps. | 
 **querySourceKustomizeNamespace** | **optional.String**| Namespace sets the namespace that Kustomize adds to all resources. | 
 **querySourceKsonnetEnvironment** | **optional.String**| Environment is a ksonnet application environment name. | 
 **querySourceDirectoryRecurse** | **optional.Bool**| Recurse specifies whether to scan a directory recursively for manifests. | 
 **querySourceDirectoryJsonnetLibs** | [**optional.Interface of []string**](string.md)| Additional library search dirs. | 
 **querySourceDirectoryExclude** | **optional.String**| Exclude contains a glob pattern to match paths against that should be explicitly excluded from being used during manifest generation. | 
 **querySourceDirectoryInclude** | **optional.String**| Include contains a glob pattern to match paths against that should be explicitly included during manifest generation. | 
 **querySourcePluginName** | **optional.String**|  | 
 **querySourceChart** | **optional.String**| Chart is a Helm chart name, and must be specified for applications sourced from a Helm repo. | 
 **querySourceRef** | **optional.String**| Ref is reference to another source within sources field. This field will not be used if used with a &#x60;source&#x60; tag. | 
 **queryAppName** | **optional.String**|  | 
 **queryAppProject** | **optional.String**|  | 

### Return type

[**RepositoriesRepoAppDetailsResponse**](repositoriesRepoAppDetailsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentRepositoryServiceGetByUrl**
> V1RepoByUrlRespList AgentRepositoryServiceGetByUrl(ctx, agentIdentifier, optional)
Get a repository configured in gitops by URL from Harness

Returns a repository by URL 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentifier** | **string**| Agent identifier for entity. | 
 **optional** | ***RepositoriesApiAgentRepositoryServiceGetByUrlOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoriesApiAgentRepositoryServiceGetByUrlOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.String**| Account Identifier for the Agent. | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Agent. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Agent. | 
 **url** | [**optional.Interface of []string**](string.md)| URL by which to find the repository entity managed by agent. | 
 **fetchScopePrefixedIdentifier** | **optional.Bool**| Indicates whether the identifier of repo fetched must contain the harness scope prefix such (account./org.) | 

### Return type

[**V1RepoByUrlRespList**](v1RepoByUrlRespList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentRepositoryServiceGetHelmCharts**
> RepositoriesHelmChartsResponse AgentRepositoryServiceGetHelmCharts(ctx, agentIdentifier, identifier, accountIdentifier, optional)
GetHelmCharts returns list of helm charts in the specified repository

GetHelmCharts returns list of helm charts in the specified repository.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentifier** | **string**| Agent identifier for entity. | 
  **identifier** | **string**|  | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***RepositoriesApiAgentRepositoryServiceGetHelmChartsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoriesApiAgentRepositoryServiceGetHelmChartsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **queryRepo** | **optional.String**| Repo URL for query. | 
 **queryForceRefresh** | **optional.Bool**| Whether to force a cache refresh on repo&#x27;s connection state. | 
 **queryProject** | **optional.String**| The associated project project. | 

### Return type

[**RepositoriesHelmChartsResponse**](repositoriesHelmChartsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentRepositoryServiceGetHelmChartsByUrl**
> RepositoriesHelmChartsResponse AgentRepositoryServiceGetHelmChartsByUrl(ctx, agentIdentifier, accountIdentifier, orgIdentifier, projectIdentifier, queryRepo, optional)
GetHelmChartsByUrl returns list of helm charts in the repository by URL

GetHelmCharts returns list of helm charts in the specified repository.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentifier** | **string**| Agent identifier for entity. | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the Entity. | 
  **projectIdentifier** | **string**| Project Identifier for the Entity. | 
  **queryRepo** | **string**| Repo URL for query. | 
 **optional** | ***RepositoriesApiAgentRepositoryServiceGetHelmChartsByUrlOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoriesApiAgentRepositoryServiceGetHelmChartsByUrlOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **identifier** | **optional.String**|  | 
 **queryForceRefresh** | **optional.Bool**| Whether to force a cache refresh on repo&#x27;s connection state. | 
 **queryProject** | **optional.String**| The associated project project. | 

### Return type

[**RepositoriesHelmChartsResponse**](repositoriesHelmChartsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentRepositoryServiceListApps**
> RepositoriesRepoAppsResponse AgentRepositoryServiceListApps(ctx, agentIdentifier, identifier, accountIdentifier, orgIdentifier, optional)
ListApps returns list of apps in the repo

ListApps returns list of apps in the repo.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentifier** | **string**| Agent identifier for entity. | 
  **identifier** | **string**|  | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the Entity. | 
 **optional** | ***RepositoriesApiAgentRepositoryServiceListAppsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoriesApiAgentRepositoryServiceListAppsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **queryRepo** | **optional.String**|  | 
 **queryRevision** | **optional.String**|  | 
 **queryAppName** | **optional.String**|  | 
 **queryAppProject** | **optional.String**|  | 

### Return type

[**RepositoriesRepoAppsResponse**](repositoriesRepoAppsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentRepositoryServiceListAppsByUrl**
> RepositoriesRepoAppsResponse AgentRepositoryServiceListAppsByUrl(ctx, agentIdentifier, optional)
ListAppsByUrl returns list of apps in the repo

ListApps returns list of apps in the repo.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentifier** | **string**| Agent identifier for entity. | 
 **optional** | ***RepositoriesApiAgentRepositoryServiceListAppsByUrlOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoriesApiAgentRepositoryServiceListAppsByUrlOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.String**| Account Identifier for the Entity. | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **identifier** | **optional.String**|  | 
 **queryRepo** | **optional.String**|  | 
 **queryRevision** | **optional.String**|  | 
 **queryAppName** | **optional.String**|  | 
 **queryAppProject** | **optional.String**|  | 

### Return type

[**RepositoriesRepoAppsResponse**](repositoriesRepoAppsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentRepositoryServiceListESOGenerators**
> RepositoriesEsoGeneratorResponse AgentRepositoryServiceListESOGenerators(ctx, agentIdentifier, accountIdentifier, optional)
Returns a list of ESO generators installed in agent namespace.

Returns a list of ESO generators installed in agent namespace.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentifier** | **string**| Agent identifier for entity. | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***RepositoriesApiAgentRepositoryServiceListESOGeneratorsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoriesApiAgentRepositoryServiceListESOGeneratorsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **type_** | **optional.String**|  | [default to UNSET]
 **name** | **optional.String**|  | 
 **url** | **optional.String**|  | 

### Return type

[**RepositoriesEsoGeneratorResponse**](repositoriesESOGeneratorResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentRepositoryServiceListRefs**
> RepositoriesRefs AgentRepositoryServiceListRefs(ctx, agentIdentifier, identifier, accountIdentifier, optional)
Returns a list of refs (e.g. branches and tags) in the repo

Returns a list of refs (e.g. branches and tags) in the repo.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentifier** | **string**| Agent identifier for entity. | 
  **identifier** | **string**|  | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***RepositoriesApiAgentRepositoryServiceListRefsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoriesApiAgentRepositoryServiceListRefsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **queryRepo** | **optional.String**| Repo URL for query. | 
 **queryForceRefresh** | **optional.Bool**| Whether to force a cache refresh on repo&#x27;s connection state. | 
 **queryProject** | **optional.String**| The associated project project. | 

### Return type

[**RepositoriesRefs**](repositoriesRefs.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentRepositoryServiceListRefsByUrl**
> RepositoriesRefs AgentRepositoryServiceListRefsByUrl(ctx, agentIdentifier, accountIdentifier, orgIdentifier, projectIdentifier, optional)
Returns a list of refs (e.g. branches and tags) in the repo by URL

Returns a list of refs (e.g. branches and tags) in the repo.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentifier** | **string**| Agent identifier for entity. | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the Entity. | 
  **projectIdentifier** | **string**| Project Identifier for the Entity. | 
 **optional** | ***RepositoriesApiAgentRepositoryServiceListRefsByUrlOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoriesApiAgentRepositoryServiceListRefsByUrlOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **identifier** | **optional.String**|  | 
 **queryRepo** | **optional.String**| Repo URL for query. | 
 **queryForceRefresh** | **optional.Bool**| Whether to force a cache refresh on repo&#x27;s connection state. | 
 **queryProject** | **optional.String**| The associated project project. | 

### Return type

[**RepositoriesRefs**](repositoriesRefs.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentRepositoryServiceListRepositories**
> RepositoriesRepositoryList AgentRepositoryServiceListRepositories(ctx, agentIdentifier, accountIdentifier, optional)
ListRepositories gets a list of all configured repositories

ListRepositories gets a list of all configured repositories.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentifier** | **string**| Agent identifier for entity. | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***RepositoriesApiAgentRepositoryServiceListRepositoriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoriesApiAgentRepositoryServiceListRepositoriesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **identifier** | **optional.String**|  | 
 **queryRepo** | **optional.String**| Repo URL for query. | 
 **queryForceRefresh** | **optional.Bool**| Whether to force a cache refresh on repo&#x27;s connection state. | 
 **queryProject** | **optional.String**| The associated project project. | 

### Return type

[**RepositoriesRepositoryList**](repositoriesRepositoryList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentRepositoryServiceUpdateRepository**
> Servicev1Repository AgentRepositoryServiceUpdateRepository(ctx, body, agentIdentifier, identifier, accountIdentifier, optional)
UpdateRepository updates a repository configuration

UpdateRepository updates a repository configuration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RepositoriesRepoUpdateRequest**](RepositoriesRepoUpdateRequest.md)|  | 
  **agentIdentifier** | **string**| Agent identifier for entity. | 
  **identifier** | **string**|  | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***RepositoriesApiAgentRepositoryServiceUpdateRepositoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoriesApiAgentRepositoryServiceUpdateRepositoryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **orgIdentifier** | **optional.**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity. | 

### Return type

[**Servicev1Repository**](servicev1Repository.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentRepositoryServiceValidateAccess**
> CommonsConnectionState AgentRepositoryServiceValidateAccess(ctx, body, agentIdentifier, accountIdentifier, optional)
ValidateAccess gets connection state for a repository

ValidateAccess gets connection state for a repository.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RepositoriesRepoAccessQuery**](RepositoriesRepoAccessQuery.md)|  | 
  **agentIdentifier** | **string**| Agent identifier for entity. | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***RepositoriesApiAgentRepositoryServiceValidateAccessOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoriesApiAgentRepositoryServiceValidateAccessOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orgIdentifier** | **optional.**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity. | 
 **identifier** | **optional.**|  | 

### Return type

[**CommonsConnectionState**](commonsConnectionState.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoryServiceExists**
> bool RepositoryServiceExists(ctx, accountIdentifier, optional)
Checks whether a repository with the given name exists

Checks whether a repository with the given name exists.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***RepositoriesApiRepositoryServiceExistsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoriesApiRepositoryServiceExistsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **agentIdentifier** | **optional.String**| Agent identifier for entity. | 
 **url** | **optional.String**|  | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoryServiceListRepositories**
> V1Repositorylist RepositoryServiceListRepositories(ctx, body)
List returns list of Repositories

List returns list of Repositories

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1RepoListReq**](V1RepoListReq.md)|  | 

### Return type

[**V1Repositorylist**](v1Repositorylist.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoryServiceListRepositoriesByRepositoryCredentialTemplate**
> V1Repositorylist RepositoryServiceListRepositoriesByRepositoryCredentialTemplate(ctx, accountIdentifier, optional)
List returns list of Repositories by repository credential template

List returns list of Repositories by repositories credential template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***RepositoriesApiRepositoryServiceListRepositoriesByRepositoryCredentialTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoriesApiRepositoryServiceListRepositoriesByRepositoryCredentialTemplateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **repositoryCredentialURL** | **optional.String**|  | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **agentIdentifier** | **optional.String**| Agent identifier for entity. | 
 **pageSize** | **optional.Int32**|  | 
 **pageIndex** | **optional.Int32**|  | 

### Return type

[**V1Repositorylist**](v1Repositorylist.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

