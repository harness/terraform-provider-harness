# nextgen{{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAPIKey**](APIKeysApi.md#AddAPIKey) | **Post** /cf/admin/apikey | Creates an API key for the given Environment
[**DeleteAPIKey**](APIKeysApi.md#DeleteAPIKey) | **Delete** /cf/admin/apikey/{identifier} | Deletes an API Key
[**GetAPIKey**](APIKeysApi.md#GetAPIKey) | **Get** /cf/admin/apikey/{identifier} | Returns API keys
[**GetAllAPIKeys**](APIKeysApi.md#GetAllAPIKeys) | **Get** /cf/admin/apikey | Returns API Keys for an Environment
[**UpdateAPIKey**](APIKeysApi.md#UpdateAPIKey) | **Put** /cf/admin/apikey/{identifier} | Updates an API Key

# **AddAPIKey**
> CfApiKey AddAPIKey(ctx, accountIdentifier, orgIdentifier, environmentIdentifier, projectIdentifier, optional)
Creates an API key for the given Environment

Creates an API key for the given Environment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier | 
  **orgIdentifier** | **string**| Organization Identifier | 
  **environmentIdentifier** | **string**| Environment Identifier | 
  **projectIdentifier** | **string**| The Project identifier | 
 **optional** | ***APIKeysApiAddAPIKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a APIKeysApiAddAPIKeyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | [**optional.Interface of interface{}**](interface{}.md)|  | 

### Return type

[**CfApiKey**](CfApiKey.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAPIKey**
> DeleteAPIKey(ctx, identifier, projectIdentifier, environmentIdentifier, accountIdentifier, orgIdentifier)
Deletes an API Key

Deletes an API key for the given identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| Unique identifier for the object in the API. | 
  **projectIdentifier** | **string**| The Project identifier | 
  **environmentIdentifier** | **string**| Environment Identifier | 
  **accountIdentifier** | **string**| Account Identifier | 
  **orgIdentifier** | **string**| Organization Identifier | 

### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAPIKey**
> CfApiKey GetAPIKey(ctx, identifier, projectIdentifier, environmentIdentifier, accountIdentifier, orgIdentifier)
Returns API keys

Returns all the API Keys for the given identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| Unique identifier for the object in the API. | 
  **projectIdentifier** | **string**| The Project identifier | 
  **environmentIdentifier** | **string**| Environment Identifier | 
  **accountIdentifier** | **string**| Account Identifier | 
  **orgIdentifier** | **string**| Organization Identifier | 

### Return type

[**CfApiKey**](CfApiKey.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllAPIKeys**
> ApiKeys GetAllAPIKeys(ctx, accountIdentifier, orgIdentifier, projectIdentifier, environmentIdentifier, optional)
Returns API Keys for an Environment

Returns all the API Keys for an Environment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier | 
  **orgIdentifier** | **string**| Organization Identifier | 
  **projectIdentifier** | **string**| The Project identifier | 
  **environmentIdentifier** | **string**| Environment Identifier | 
 **optional** | ***APIKeysApiGetAllAPIKeysOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a APIKeysApiGetAllAPIKeysOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **pageNumber** | **optional.Int32**| PageNumber | 
 **pageSize** | **optional.Int32**| PageSize | 

### Return type

[**ApiKeys**](ApiKeys.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAPIKey**
> UpdateAPIKey(ctx, identifier, projectIdentifier, environmentIdentifier, accountIdentifier, orgIdentifier, optional)
Updates an API Key

Updates an API key for the given identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| Unique identifier for the object in the API. | 
  **projectIdentifier** | **string**| The Project identifier | 
  **environmentIdentifier** | **string**| Environment Identifier | 
  **accountIdentifier** | **string**| Account Identifier | 
  **orgIdentifier** | **string**| Organization Identifier | 
 **optional** | ***APIKeysApiUpdateAPIKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a APIKeysApiUpdateAPIKeyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **body** | [**optional.Interface of interface{}**](interface{}.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

