# {{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AgentServiceForServerCreate**](AgentsApi.md#AgentServiceForServerCreate) | **Post** /gitops/api/api/v1/agents | 
[**AgentServiceForServerDelete**](AgentsApi.md#AgentServiceForServerDelete) | **Delete** /gitops/api/api/v1/agents/{identifier} | 
[**AgentServiceForServerGet**](AgentsApi.md#AgentServiceForServerGet) | **Get** /gitops/api/api/v1/agents/{identifier} | 
[**AgentServiceForServerGetDeployYaml**](AgentsApi.md#AgentServiceForServerGetDeployYaml) | **Get** /gitops/api/api/v1/agents/{agentIdentifier}/deploy.yaml | 
[**AgentServiceForServerList**](AgentsApi.md#AgentServiceForServerList) | **Get** /gitops/api/api/v1/agents | 
[**AgentServiceForServerRegenerateCredentials**](AgentsApi.md#AgentServiceForServerRegenerateCredentials) | **Post** /gitops/api/api/v1/agents/{identifier}/credentials | 
[**AgentServiceForServerUnique**](AgentsApi.md#AgentServiceForServerUnique) | **Get** /gitops/api/api/v1/agents/{identifier}/unique | 
[**AgentServiceForServerUpdate**](AgentsApi.md#AgentServiceForServerUpdate) | **Put** /gitops/api/api/v1/agents/{agent.identifier} | 

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

### Return type

[**V1Agent**](v1Agent.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentServiceForServerGetDeployYaml**
> string AgentServiceForServerGetDeployYaml(ctx, agentIdentifier, accountIdentifier, optional)


GetDeployYaml returns depoyment yamls for agents.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentifier** | **string**| Agent identifier for entity. | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***AgentsApiAgentServiceForServerGetDeployYamlOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AgentsApiAgentServiceForServerGetDeployYamlOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **namespace** | **optional.String**|  | 

### Return type

**string**

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/x-yml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentServiceForServerList**
> V1AgentList AgentServiceForServerList(ctx, accountIdentifier, type_, optional)


List agents.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **type_** | **string**| MANAGED_ARGO_PROVIDER | [default to AGENT_TYPE_UNSET]
 **optional** | ***AgentsApiAgentServiceForServerListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AgentsApiAgentServiceForServerListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **identifier** | **optional.String**|  | 
 **name** | **optional.String**|  | 
 **tags** | [**optional.Interface of []string**](string.md)|  | 
 **searchTerm** | **optional.String**|  | 
 **pageSize** | **optional.Int32**|  | 
 **pageIndex** | **optional.Int32**|  | 
 **scope** | **optional.String**|  | [default to AGENT_SCOPE_UNSET]

### Return type

[**V1AgentList**](v1AgentList.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

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

