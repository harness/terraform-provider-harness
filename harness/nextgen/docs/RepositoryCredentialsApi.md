# nextgen{{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AgentRepositoryCredentialsServiceCreateRepositoryCredentials**](RepositoryCredentialsApi.md#AgentRepositoryCredentialsServiceCreateRepositoryCredentials) | **Post** /gitops/api/v1/agents/{agentIdentifier}/repocreds | Create creates a new repository credential
[**AgentRepositoryCredentialsServiceDeleteRepositoryCredentials**](RepositoryCredentialsApi.md#AgentRepositoryCredentialsServiceDeleteRepositoryCredentials) | **Delete** /gitops/api/v1/agents/{agentIdentifier}/repocreds/{identifier} | Delete deletes a repository credential
[**AgentRepositoryCredentialsServiceGetCredentialsForRepositoryUrl**](RepositoryCredentialsApi.md#AgentRepositoryCredentialsServiceGetCredentialsForRepositoryUrl) | **Post** /gitops/api/v1/agents/{agentIdentifier}/repocreds/get | Get returns a repository credential given its url
[**AgentRepositoryCredentialsServiceGetRepositoryCredentials**](RepositoryCredentialsApi.md#AgentRepositoryCredentialsServiceGetRepositoryCredentials) | **Get** /gitops/api/v1/agents/{agentIdentifier}/repocreds/{identifier} | Get returns a repository credential given its identifier
[**AgentRepositoryCredentialsServiceListRepositoryCredentials**](RepositoryCredentialsApi.md#AgentRepositoryCredentialsServiceListRepositoryCredentials) | **Post** /gitops/api/v1/repocreds | List repository credentials
[**AgentRepositoryCredentialsServiceUpdateRepositoryCredentials**](RepositoryCredentialsApi.md#AgentRepositoryCredentialsServiceUpdateRepositoryCredentials) | **Put** /gitops/api/v1/agents/{agentIdentifier}/repocreds/{identifier} | Update updates a repository credential

# **AgentRepositoryCredentialsServiceCreateRepositoryCredentials**
> Servicev1RepositoryCredentials AgentRepositoryCredentialsServiceCreateRepositoryCredentials(ctx, body, agentIdentifier, accountIdentifier, optional)
Create creates a new repository credential

Create creates a new repository credential.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**HrepocredsRepoCredsCreateRequest**](HrepocredsRepoCredsCreateRequest.md)|  | 
  **agentIdentifier** | **string**| Agent identifier for entity. | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***RepositoryCredentialsApiAgentRepositoryCredentialsServiceCreateRepositoryCredentialsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryCredentialsApiAgentRepositoryCredentialsServiceCreateRepositoryCredentialsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orgIdentifier** | **optional.**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity. | 
 **identifier** | **optional.**|  | 

### Return type

[**Servicev1RepositoryCredentials**](servicev1RepositoryCredentials.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentRepositoryCredentialsServiceDeleteRepositoryCredentials**
> HrepocredsRepoCredsResponse AgentRepositoryCredentialsServiceDeleteRepositoryCredentials(ctx, agentIdentifier, identifier, optional)
Delete deletes a repository credential

 Delete deletes a repository credential.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentifier** | **string**| Agent identifier for entity. | 
  **identifier** | **string**|  | 
 **optional** | ***RepositoryCredentialsApiAgentRepositoryCredentialsServiceDeleteRepositoryCredentialsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryCredentialsApiAgentRepositoryCredentialsServiceDeleteRepositoryCredentialsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accountIdentifier** | **optional.String**| Account Identifier for the Entity. | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **queryUrl** | **optional.String**| Repo URL for query. | 
 **queryRepoCredsType** | **optional.String**| RepoCreds type - git or helm. | 

### Return type

[**HrepocredsRepoCredsResponse**](hrepocredsRepoCredsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentRepositoryCredentialsServiceGetCredentialsForRepositoryUrl**
> Servicev1RepositoryCredentials AgentRepositoryCredentialsServiceGetCredentialsForRepositoryUrl(ctx, body, agentIdentifier, accountIdentifier, optional)
Get returns a repository credential given its url

Get returns a repository credential given its url.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**HrepocredsRepoCredsQuery**](HrepocredsRepoCredsQuery.md)|  | 
  **agentIdentifier** | **string**| Agent identifier for entity. | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***RepositoryCredentialsApiAgentRepositoryCredentialsServiceGetCredentialsForRepositoryUrlOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryCredentialsApiAgentRepositoryCredentialsServiceGetCredentialsForRepositoryUrlOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orgIdentifier** | **optional.**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity. | 
 **identifier** | **optional.**|  | 

### Return type

[**Servicev1RepositoryCredentials**](servicev1RepositoryCredentials.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentRepositoryCredentialsServiceGetRepositoryCredentials**
> Servicev1RepositoryCredentials AgentRepositoryCredentialsServiceGetRepositoryCredentials(ctx, agentIdentifier, identifier, accountIdentifier, optional)
Get returns a repository credential given its identifier

Get returns a repository credential given its identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentifier** | **string**| Agent identifier for entity. | 
  **identifier** | **string**|  | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***RepositoryCredentialsApiAgentRepositoryCredentialsServiceGetRepositoryCredentialsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryCredentialsApiAgentRepositoryCredentialsServiceGetRepositoryCredentialsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **queryUrl** | **optional.String**| Repo URL for query. | 
 **queryRepoCredsType** | **optional.String**| RepoCreds type - git or helm. | 

### Return type

[**Servicev1RepositoryCredentials**](servicev1RepositoryCredentials.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentRepositoryCredentialsServiceListRepositoryCredentials**
> Servicev1RepositoryCredentialsList AgentRepositoryCredentialsServiceListRepositoryCredentials(ctx, body)
List repository credentials

List repository credentials.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1RepositoryCredentialsQuery**](V1RepositoryCredentialsQuery.md)|  | 

### Return type

[**Servicev1RepositoryCredentialsList**](servicev1RepositoryCredentialsList.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentRepositoryCredentialsServiceUpdateRepositoryCredentials**
> Servicev1RepositoryCredentials AgentRepositoryCredentialsServiceUpdateRepositoryCredentials(ctx, body, agentIdentifier, identifier, optional)
Update updates a repository credential

Update updates a repository credential.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**HrepocredsRepoCredsUpdateRequest**](HrepocredsRepoCredsUpdateRequest.md)|  | 
  **agentIdentifier** | **string**| Agent identifier for entity. | 
  **identifier** | **string**|  | 
 **optional** | ***RepositoryCredentialsApiAgentRepositoryCredentialsServiceUpdateRepositoryCredentialsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryCredentialsApiAgentRepositoryCredentialsServiceUpdateRepositoryCredentialsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **accountIdentifier** | **optional.**| Account Identifier for the Entity. | 
 **orgIdentifier** | **optional.**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity. | 

### Return type

[**Servicev1RepositoryCredentials**](servicev1RepositoryCredentials.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

