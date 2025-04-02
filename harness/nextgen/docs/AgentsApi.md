# {{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AgentServiceForServerCreate**](AgentsApi.md#AgentServiceForServerCreate) | **Post** /gitops/api/v1/agents | 
[**AgentServiceForServerDelete**](AgentsApi.md#AgentServiceForServerDelete) | **Delete** /gitops/api/v1/agents/{identifier} | 
[**AgentServiceForServerGet**](AgentsApi.md#AgentServiceForServerGet) | **Get** /gitops/api/v1/agents/{identifier} | 
[**AgentServiceForServerGetDeployYaml**](AgentsApi.md#AgentServiceForServerGetDeployYaml) | **Get** /gitops/api/v1/agents/{agentIdentifier}/deploy.yaml | 
[**AgentServiceForServerList**](AgentsApi.md#AgentServiceForServerList) | **Get** /gitops/api/v1/agents | 
[**AgentServiceForServerRegenerateCredentials**](AgentsApi.md#AgentServiceForServerRegenerateCredentials) | **Post** /gitops/api/v1/agents/{identifier}/credentials | 
[**AgentServiceForServerUnique**](AgentsApi.md#AgentServiceForServerUnique) | **Get** /gitops/api/v1/agents/{identifier}/unique | 
[**AgentServiceForServerUpdate**](AgentsApi.md#AgentServiceForServerUpdate) | **Put** /gitops/api/v1/agents/{agent.identifier} | 

# **AgentServiceForServerCreate**
> V1Agent AgentServiceForServerCreate(ctx, body)


Create agent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1Agent**](V1Agent.md)|  | 

### Return type

[**V1Agent**](v1Agent.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentServiceForServerDelete**
> V1Agent AgentServiceForServerDelete(ctx, identifier, optional)


Delete agents.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**|  | 
 **optional** | ***AgentsApiAgentServiceForServerDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AgentsApiAgentServiceForServerDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.String**| Account Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **name** | **optional.String**|  | 
 **type_** | **optional.String**|  | [default to AGENT_TYPE_UNSET]
 **tags** | [**optional.Interface of []string**](string.md)|  | 
 **searchTerm** | **optional.String**|  | 
 **pageSize** | **optional.Int32**|  | 
 **pageIndex** | **optional.Int32**|  | 
 **scope** | **optional.String**|  | [default to AGENT_SCOPE_UNSET]
 **drIdentifier** | **optional.String**|  | 
 **sortBy** | **optional.String**|  | [default to SORT_BY_NOT_SET]
 **sortOrder** | **optional.String**|  | [default to SORT_ORDER_NOT_SET]
 **metadataOnly** | **optional.Bool**|  | 
 **ignoreScope** | **optional.Bool**|  | 
 **connectedStatus** | **optional.String**|  | [default to CONNECTED_STATUS_UNSET]
 **healthStatus** | **optional.String**|  | [default to HEALTH_STATUS_UNSET]
 **withCredentials** | **optional.Bool**| Applicable when trying to retrieve an agent. Set to true to include the credentials for the agent in the response. (Private key may not be included in response if agent is already connected to harness). NOTE: Setting this to true requires the user to have edit permissions on Agent. | 

### Return type

[**V1Agent**](v1Agent.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentServiceForServerGet**
> V1Agent AgentServiceForServerGet(ctx, identifier, accountIdentifier, optional)


Get agents.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**|  | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***AgentsApiAgentServiceForServerGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AgentsApiAgentServiceForServerGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **name** | **optional.String**|  | 
 **type_** | **optional.String**|  | [default to AGENT_TYPE_UNSET]
 **tags** | [**optional.Interface of []string**](string.md)|  | 
 **searchTerm** | **optional.String**|  | 
 **pageSize** | **optional.Int32**|  | 
 **pageIndex** | **optional.Int32**|  | 
 **scope** | **optional.String**|  | [default to AGENT_SCOPE_UNSET]
 **withCredentials** | **optional.Bool**| Applicable when trying to retrieve an agent. Set to true to include the credentials for the agent in the response. (Private key may not be included in response if agent is already connected to harness). NOTE: Setting this to true requires the user to have edit permissions on Agent. |

### Return type

[**V1Agent**](v1Agent.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentServiceForServerGetDeployHelmChart**
> StreamResultOfV1DownloadResponse AgentServiceForServerGetDeployHelmChart(ctx, agentIdentifier, optional)


GetDeployHelmChart returns the Helm Chart for depoying the agents.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentifier** | **string**| Agent identifier for entity. | 
 **optional** | ***AgentsApiAgentServiceForServerGetDeployHelmChartOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AgentsApiAgentServiceForServerGetDeployHelmChartOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.String**| Account Identifier for the Entity. | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **namespace** | **optional.String**|  | 
 **disasterRecoveryIdentifier** | **optional.String**| Disaster Recovery Identifier for entity. | 
 **skipCrds** | **optional.Bool**|  | 
 **caData** | **optional.String**| Certificate chain for the agent, must be base64 encoded. | 
 **proxyHttp** | **optional.String**|  | 
 **proxyHttps** | **optional.String**|  | 
 **proxyUsername** | **optional.String**|  | 
 **proxyPassword** | **optional.String**|  | 
 **proxySkipSSLVerify** | **optional.Bool**|  | 
 **privateKey** | **optional.String**|  | 
 **argocdSettingsEnableHelmPathTraversal** | **optional.Bool**| Controls the Environment variable HELM_SECRETS_VALUES_ALLOW_PATH_TRAVERSAL to allow or deny dot-dot-slash values file paths. Disabled by default for security reasons. This config is pushed as an env variable to the repo-server. | 

### Return type

[**StreamResultOfV1DownloadResponse**](Stream result of v1DownloadResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
# **AgentServiceForServerGetDeployYaml**
> string AgentServiceForServerGetDeployYaml(ctx, agentIdentifier, optional)


GetDeployYaml returns depoyment yamls for agents.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentifier** | **string**| Agent identifier for entity. | 
 **optional** | ***AgentsApiAgentServiceForServerGetDeployYamlOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AgentsApiAgentServiceForServerGetDeployYamlOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accountIdentifier** | **optional.String**| Account Identifier for the Entity. | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **namespace** | **optional.String**|  | 
 **disasterRecoveryIdentifier** | **optional.String**| Disaster Recovery Identifier for entity. | 
 **skipCrds** | **optional.Bool**|  | 
 **caData** | **optional.String**| Certificate chain for the agent, must be base64 encoded. | 
 **proxyHttp** | **optional.String**|  | 
 **proxyHttps** | **optional.String**|  | 
 **proxyUsername** | **optional.String**|  | 
 **proxyPassword** | **optional.String**|  | 
 **proxySkipSSLVerify** | **optional.Bool**|  | 
 **privateKey** | **optional.String**|  | 
 **argocdSettingsEnableHelmPathTraversal** | **optional.Bool**| Controls the Environment variable HELM_SECRETS_VALUES_ALLOW_PATH_TRAVERSAL to allow or deny dot-dot-slash values file paths. Disabled by default for security reasons. This config is pushed as an env variable to the repo-server. | 

### Return type

**string**

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentServiceForServerGetOperatorYaml**
> string AgentServiceForServerGetOperatorYaml(ctx, agentIdentifier, optional)


GetOperatorYaml returns operator yaml for deploying the agents.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentifier** | **string**| Agent identifier for entity. | 
 **optional** | ***AgentsApiAgentServiceForServerGetOperatorYamlOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AgentsApiAgentServiceForServerGetOperatorYamlOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.String**| Account Identifier for the Entity. | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **namespace** | **optional.String**|  | 
 **disasterRecoveryIdentifier** | **optional.String**| Disaster Recovery Identifier for entity. | 
 **skipCrds** | **optional.Bool**|  | 
 **caData** | **optional.String**| Certificate chain for the agent, must be base64 encoded. | 
 **proxyHttp** | **optional.String**|  | 
 **proxyHttps** | **optional.String**|  | 
 **proxyUsername** | **optional.String**|  | 
 **proxyPassword** | **optional.String**|  | 
 **proxySkipSSLVerify** | **optional.Bool**|  | 
 **privateKey** | **optional.String**|  | 
 **argocdSettingsEnableHelmPathTraversal** | **optional.Bool**| Controls the Environment variable HELM_SECRETS_VALUES_ALLOW_PATH_TRAVERSAL to allow or deny dot-dot-slash values file paths. Disabled by default for security reasons. This config is pushed as an env variable to the repo-server. | 

### Return type

**string**

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


# **AgentServiceForServerList**
> V1AgentList AgentServiceForServerList(ctx, optional)


List agents.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AgentsApiAgentServiceForServerListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AgentsApiAgentServiceForServerListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accountIdentifier** | **optional.String**| Account Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **identifier** | **optional.String**|  | 
 **name** | **optional.String**|  | 
 **type_** | **optional.String**|  | [default to AGENT_TYPE_UNSET]
 **tags** | [**optional.Interface of []string**](string.md)|  | 
 **searchTerm** | **optional.String**|  | 
 **pageSize** | **optional.Int32**|  | 
 **pageIndex** | **optional.Int32**|  | 
 **scope** | **optional.String**|  | [default to AGENT_SCOPE_UNSET]
 **drIdentifier** | **optional.String**|  | 
 **sortBy** | **optional.String**|  | [default to SORT_BY_NOT_SET]
 **sortOrder** | **optional.String**|  | [default to SORT_ORDER_NOT_SET]
 **metadataOnly** | **optional.Bool**|  | 
 **ignoreScope** | **optional.Bool**|  | 
 **connectedStatus** | **optional.String**|  | [default to CONNECTED_STATUS_UNSET]
 **healthStatus** | **optional.String**|  | [default to HEALTH_STATUS_UNSET]
 **withCredentials** | **optional.Bool**| Applicable when trying to retrieve an agent. Set to true to include the credentials for the agent in the response. (Private key may not be included in response if agent is already connected to harness). NOTE: Setting this to true requires the user to have edit permissions on Agent. | 

### Return type

[**V1AgentList**](v1AgentList.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentServiceForServerPostDeployHelmChart**
> StreamResultOfV1DownloadResponse AgentServiceForServerPostDeployHelmChart(ctx, body, agentIdentifier)


PostDeployHelmChart returns the Helm Chart for deploying the agents.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1AgentYamlQuery**](V1AgentYamlQuery.md)|  |
  **agentIdentifier** | **string**| Agent identifier for entity. |

### Return type

[**StreamResultOfV1DownloadResponse**](Stream result of v1DownloadResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentServiceForServerPostDeployYaml**
> string AgentServiceForServerPostDeployYaml(ctx, body, agentIdentifier)


PostDeployYaml returns deployment yamls for agents.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1AgentYamlQuery**](V1AgentYamlQuery.md)|  |
  **agentIdentifier** | **string**| Agent identifier for entity. |

### Return type

**string**

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
 - **Accept**: application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

PostOperatorYaml returns operator yaml for deploying the agents.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1AgentYamlQuery**](V1AgentYamlQuery.md)|  | 
  **agentIdentifier** | **string**| Agent identifier for entity. | 

### Return type

**string**

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


# **AgentServiceForServerRegenerateCredentials**
> V1Agent AgentServiceForServerRegenerateCredentials(ctx, identifier)


Regenerate credentials for agents.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**|  | 

### Return type

[**V1Agent**](v1Agent.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentServiceForServerUnique**
> V1UniqueMessage AgentServiceForServerUnique(ctx, identifier, accountIdentifier, optional)


Unique returns unique agents.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**|  | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***AgentsApiAgentServiceForServerUniqueOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AgentsApiAgentServiceForServerUniqueOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **name** | **optional.String**|  | 
 **type_** | **optional.String**|  | [default to AGENT_TYPE_UNSET]
 **tags** | [**optional.Interface of []string**](string.md)|  | 
 **searchTerm** | **optional.String**|  | 
 **pageSize** | **optional.Int32**|  | 
 **pageIndex** | **optional.Int32**|  | 
 **scope** | **optional.String**|  | [default to AGENT_SCOPE_UNSET]
 **drIdentifier** | **optional.String**|  | 
 **sortBy** | **optional.String**|  | [default to SORT_BY_NOT_SET]
 **sortOrder** | **optional.String**|  | [default to SORT_ORDER_NOT_SET]
 **metadataOnly** | **optional.Bool**|  | 
 **ignoreScope** | **optional.Bool**|  | 
 **connectedStatus** | **optional.String**|  | [default to CONNECTED_STATUS_UNSET]
 **healthStatus** | **optional.String**|  | [default to HEALTH_STATUS_UNSET]
 **withCredentials** | **optional.Bool**| Applicable when trying to retrieve an agent. Set to true to include the credentials for the agent in the response. (Private key may not be included in response if agent is already connected to harness). NOTE: Setting this to true requires the user to have edit permissions on Agent. | 

### Return type

[**V1UniqueMessage**](v1UniqueMessage.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentServiceForServerUpdate**
> V1Agent AgentServiceForServerUpdate(ctx, body, agentIdentifier)


Update agents.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1Agent**](V1Agent.md)|  | 
  **agentIdentifier** | **string**| The gitops-server generated ID for this gitops-agent | 

### Return type

[**V1Agent**](v1Agent.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

