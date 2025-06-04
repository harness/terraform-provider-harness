# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAgent**](AgentApi.md#CreateAgent) | **Post** /api/v1/agents | Create an agent
[**CreateInstallation**](AgentApi.md#CreateInstallation) | **Post** /api/v1/agents/{agentIdentity}/installations | Install an agent
[**DeleteAgent**](AgentApi.md#DeleteAgent) | **Delete** /api/v1/agents/{agentIdentity} | Delete an agent
[**GetAgent**](AgentApi.md#GetAgent) | **Get** /api/v1/agents/{agentIdentity} | Get an agent
[**GetAgentToken**](AgentApi.md#GetAgentToken) | **Get** /api/v1/agenttoken/{agentIdentity} | Get token for agent
[**GetInfrastructureForAgent**](AgentApi.md#GetInfrastructureForAgent) | **Get** /api/v1/agents/{agentIdentity}/infrastructure | Get infrastructure details for an agent
[**GetInstallation**](AgentApi.md#GetInstallation) | **Get** /api/v1/agents/{agentIdentity}/installations/{installationID} | Get an agent install
[**ListAgent**](AgentApi.md#ListAgent) | **Get** /api/v1/agents | Get list of agents
[**ListInfrastructure**](AgentApi.md#ListInfrastructure) | **Get** /api/v1/infrastructures | Get list of infrastructures
[**ListInstallation**](AgentApi.md#ListInstallation) | **Get** /api/v1/agents/{agentIdentity}/installations | Get list of agent installations
[**StopOngoingDiscovery**](AgentApi.md#StopOngoingDiscovery) | **Post** /api/v1/agents/{agentIdentity}/installations/{installationID}/stop | stops an ongoing discovery
[**UpdateAgent**](AgentApi.md#UpdateAgent) | **Put** /api/v1/agents/{agentIdentity} | Update an agent

# **CreateAgent**
> ApiGetAgentResponse CreateAgent(ctx, body, accountIdentifier, optional)
Create an agent

Create a new agent

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiCreateAgentRequest**](ApiCreateAgentRequest.md)| Create Agent | 
  **accountIdentifier** | **string**| account id is the account where you want to create the resource | 
 **optional** | ***AgentApiCreateAgentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AgentApiCreateAgentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **correlationID** | **optional.**| correlation id is used to debug micro svc communication | 
 **organizationIdentifier** | **optional.**| organization id is the organization where you want to create the resource | 
 **projectIdentifier** | **optional.**| project id is the project where you want to create the resource | 
 **noInstallation** | **optional.**| don&#x27;t install agent | [default to false]

### Return type

[**ApiGetAgentResponse**](api.GetAgentResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateInstallation**
> ApiGetInstallationResponse CreateInstallation(ctx, body, agentIdentity, accountIdentifier, environmentIdentifier, optional)
Install an agent

Install agent

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiCreateInstallationRequest**](ApiCreateInstallationRequest.md)| Create Installation | 
  **agentIdentity** | **string**| agent identity | 
  **accountIdentifier** | **string**| account id is the account where you want to access the resource | 
  **environmentIdentifier** | **string**| environment id is the environment where you want to access the resource | 
 **optional** | ***AgentApiCreateInstallationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AgentApiCreateInstallationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **correlationID** | **optional.**| correlation id is used to debug micro svc communication | 
 **organizationIdentifier** | **optional.**| organization id is the organization where you want to access the resource | 
 **projectIdentifier** | **optional.**| project id is the project where you want to access the resource | 

### Return type

[**ApiGetInstallationResponse**](api.GetInstallationResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAgent**
> ApiEmpty DeleteAgent(ctx, agentIdentity, accountIdentifier, environmentIdentifier, optional)
Delete an agent

Delete an agent

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentity** | **string**| agent identity | 
  **accountIdentifier** | **string**| account id is the account where you want to access the resource | 
  **environmentIdentifier** | **string**| environment id is the environment where you want to access the resource | 
 **optional** | ***AgentApiDeleteAgentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AgentApiDeleteAgentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **correlationID** | **optional.String**| correlation id is used to debug micro svc communication | 
 **organizationIdentifier** | **optional.String**| organization id is the organization where you want to access the resource | 
 **projectIdentifier** | **optional.String**| project id is the project where you want to access the resource | 

### Return type

[**ApiEmpty**](api.Empty.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAgent**
> ApiGetAgentResponse GetAgent(ctx, agentIdentity, accountIdentifier, environmentIdentifier, optional)
Get an agent

Get an agent

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentity** | **string**| agent identity | 
  **accountIdentifier** | **string**| account id is the account where you want to access the resource | 
  **environmentIdentifier** | **string**| environment id is the environment where you want to access the resource | 
 **optional** | ***AgentApiGetAgentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AgentApiGetAgentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **correlationID** | **optional.String**| correlation id is used to debug micro svc communication | 
 **organizationIdentifier** | **optional.String**| organization id is the organization where you want to access the resource | 
 **projectIdentifier** | **optional.String**| project id is the project where you want to access the resource | 

### Return type

[**ApiGetAgentResponse**](api.GetAgentResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAgentToken**
> ApiGetAgentToken GetAgentToken(ctx, agentIdentity, accountIdentifier, environmentIdentifier, optional)
Get token for agent

Get token for a given agent

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentity** | **string**| agent identity | 
  **accountIdentifier** | **string**| account id is the account where you want to access the resource | 
  **environmentIdentifier** | **string**| environment id is the environment where you want to access the resource | 
 **optional** | ***AgentApiGetAgentTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AgentApiGetAgentTokenOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **correlationID** | **optional.String**| correlation id is used to debug micro svc communication | 
 **organizationIdentifier** | **optional.String**| organization id is the organization where you want to access the resource | 
 **projectIdentifier** | **optional.String**| project id is the project where you want to access the resource | 

### Return type

[**ApiGetAgentToken**](api.GetAgentToken.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInfrastructureForAgent**
> ApiGetAgentResponse GetInfrastructureForAgent(ctx, agentIdentity, accountIdentifier, environmentIdentifier, optional)
Get infrastructure details for an agent

Get infrastructure details for an agent

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentity** | **string**| agent identity | 
  **accountIdentifier** | **string**| account id is the account where you want to access the resource | 
  **environmentIdentifier** | **string**| environment id is the environment where you want to access the resource | 
 **optional** | ***AgentApiGetInfrastructureForAgentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AgentApiGetInfrastructureForAgentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **correlationID** | **optional.String**| correlation id is used to debug micro svc communication | 
 **organizationIdentifier** | **optional.String**| organization id is the organization where you want to access the resource | 
 **projectIdentifier** | **optional.String**| project id is the project where you want to access the resource | 

### Return type

[**ApiGetAgentResponse**](api.GetAgentResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInstallation**
> ApiGetInstallationResponse GetInstallation(ctx, agentIdentity, installationID, accountIdentifier, environmentIdentifier, optional)
Get an agent install

Get an agent install

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentity** | **string**| agent identity | 
  **installationID** | **string**| installation ID | 
  **accountIdentifier** | **string**| account id is the account where you want to access the resource | 
  **environmentIdentifier** | **string**| environment id is the environment where you want to access the resource | 
 **optional** | ***AgentApiGetInstallationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AgentApiGetInstallationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **correlationID** | **optional.String**| correlation id is used to debug micro svc communication | 
 **organizationIdentifier** | **optional.String**| organization id is the organization where you want to access the resource | 
 **projectIdentifier** | **optional.String**| project id is the project where you want to access the resource | 

### Return type

[**ApiGetInstallationResponse**](api.GetInstallationResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAgent**
> ApiListAgentResponse ListAgent(ctx, accountIdentifier, environmentIdentifier, page, limit, all, optional)
Get list of agents

Get list of agents

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| account id is the account where you want to access the resource | 
  **environmentIdentifier** | **string**| environment id is the environment where you want to access the resource | 
  **page** | **int32**| page number | [default to 0]
  **limit** | **int32**| limit per page | [default to 10]
  **all** | **bool**| get all | [default to false]
 **optional** | ***AgentApiListAgentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AgentApiListAgentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **correlationID** | **optional.String**| correlation id is used to debug micro svc communication | 
 **organizationIdentifier** | **optional.String**| organization id is the organization where you want to access the resource | 
 **projectIdentifier** | **optional.String**| project id is the project where you want to access the resource | 
 **search** | **optional.String**| search based on name | 

### Return type

[**ApiListAgentResponse**](api.ListAgentResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListInfrastructure**
> ApiListInfrastructureResponse ListInfrastructure(ctx, accountIdentifier, environmentIdentifier, page, limit, optional)
Get list of infrastructures

Get list of infrastructures

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| account id is the account where you want to access the resource | 
  **environmentIdentifier** | **string**| environment id is the environment where you want to access the resource | 
  **page** | **int32**| page number | [default to 0]
  **limit** | **int32**| limit per page | [default to 10]
 **optional** | ***AgentApiListInfrastructureOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AgentApiListInfrastructureOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **correlationID** | **optional.String**| correlation id is used to debug micro svc communication | 
 **organizationIdentifier** | **optional.String**| organization id is the organization where you want to access the resource | 
 **projectIdentifier** | **optional.String**| project id is the project where you want to access the resource | 
 **search** | **optional.String**| search based on name | 

### Return type

[**ApiListInfrastructureResponse**](api.ListInfrastructureResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListInstallation**
> ApiListInstallationResponse ListInstallation(ctx, agentIdentity, accountIdentifier, environmentIdentifier, page, limit, all, optional)
Get list of agent installations

Get list of agent installations

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentity** | **string**| agent identity | 
  **accountIdentifier** | **string**| account id is the account where you want to access the resource | 
  **environmentIdentifier** | **string**| environment id is the environment where you want to access the resource | 
  **page** | **int32**| page number | [default to 0]
  **limit** | **int32**| limit per page | [default to 10]
  **all** | **bool**| get all | [default to false]
 **optional** | ***AgentApiListInstallationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AgentApiListInstallationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **correlationID** | **optional.String**| correlation id is used to debug micro svc communication | 
 **organizationIdentifier** | **optional.String**| organization id is the organization where you want to access the resource | 
 **projectIdentifier** | **optional.String**| project id is the project where you want to access the resource | 

### Return type

[**ApiListInstallationResponse**](api.ListInstallationResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StopOngoingDiscovery**
> ApiGetInstallationResponse StopOngoingDiscovery(ctx, agentIdentity, installationID, accountIdentifier, environmentIdentifier, optional)
stops an ongoing discovery

Stops ongoing discovery

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentity** | **string**| agent identity | 
  **installationID** | **string**| installation ID | 
  **accountIdentifier** | **string**| account id is the account where you want to access the resource | 
  **environmentIdentifier** | **string**| environment id is the environment where you want to access the resource | 
 **optional** | ***AgentApiStopOngoingDiscoveryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AgentApiStopOngoingDiscoveryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **correlationID** | **optional.String**| correlation id is used to debug micro svc communication | 
 **organizationIdentifier** | **optional.String**| organization id is the organization where you want to access the resource | 
 **projectIdentifier** | **optional.String**| project id is the project where you want to access the resource | 

### Return type

[**ApiGetInstallationResponse**](api.GetInstallationResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAgent**
> ApiGetAgentResponse UpdateAgent(ctx, body, agentIdentity, accountIdentifier, environmentIdentifier, optional)
Update an agent

Update an agent

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiUpdateAgentRequest**](ApiUpdateAgentRequest.md)| Update Agent | 
  **agentIdentity** | **string**| agent identity | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **environmentIdentifier** | **string**| environment id is the environment where you want to access the resource | 
 **optional** | ***AgentApiUpdateAgentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AgentApiUpdateAgentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **correlationID** | **optional.**| correlation id is used to debug micro svc communication | 
 **organizationIdentifier** | **optional.**| organization id that want to access the resource | 
 **projectIdentifier** | **optional.**| project id that want to access the resource | 

### Return type

[**ApiGetAgentResponse**](api.GetAgentResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

