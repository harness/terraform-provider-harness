# {{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiKey**](ApiKeyApi.md#CreateApiKey) | **Post** /ng/api/apikey | Creates an API key
[**DeleteApiKey**](ApiKeyApi.md#DeleteApiKey) | **Delete** /ng/api/apikey/{identifier} | Deletes the API Key corresponding to the provided ID.
[**GetAggregatedApiKey**](ApiKeyApi.md#GetAggregatedApiKey) | **Get** /ng/api/apikey/aggregate/{identifier} | Fetches the API Keys details corresponding to the provided ID and Scope.
[**ListApiKeys**](ApiKeyApi.md#ListApiKeys) | **Get** /ng/api/apikey | Fetches the list of API Keys corresponding to the request&#x27;s filter criteria.
[**ListApiKeys1**](ApiKeyApi.md#ListApiKeys1) | **Get** /ng/api/apikey/aggregate | Fetches the list of Aggregated API Keys corresponding to the request&#x27;s filter criteria.
[**UpdateApiKey**](ApiKeyApi.md#UpdateApiKey) | **Put** /ng/api/apikey/{identifier} | Updates API Key for the provided ID

# **CreateApiKey**
> ResponseDtoApiKey CreateApiKey(ctx, optional)
Creates an API key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ApiKeyApiCreateApiKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApiKeyApiCreateApiKeyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ApiKey**](ApiKey.md)|  | 

### Return type

[**ResponseDtoApiKey**](ResponseDTOApiKey.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml, text/plain
 - **Accept**: application/json, application/yaml, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApiKey**
> ResponseDtoBoolean DeleteApiKey(ctx, accountIdentifier, apiKeyType, parentIdentifier, identifier, optional)
Deletes the API Key corresponding to the provided ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
  **apiKeyType** | **string**| This is the API Key type like Personal Access Key or Service Account Key. | 
  **parentIdentifier** | **string**| Id of API key&#x27;s Parent Service Account | 
  **identifier** | **string**| This is the API key ID | 
 **optional** | ***ApiKeyApiDeleteApiKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApiKeyApiDeleteApiKeyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity | 

### Return type

[**ResponseDtoBoolean**](ResponseDTOBoolean.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAggregatedApiKey**
> ResponseDtoApiKeyAggregate GetAggregatedApiKey(ctx, accountIdentifier, apiKeyType, parentIdentifier, identifier, optional)
Fetches the API Keys details corresponding to the provided ID and Scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
  **apiKeyType** | **string**| This is the API Key type like Personal Access Key or Service Account Key. | 
  **parentIdentifier** | **string**| ID of API key&#x27;s Parent Service Account | 
  **identifier** | **string**| This is the API key ID | 
 **optional** | ***ApiKeyApiGetAggregatedApiKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApiKeyApiGetAggregatedApiKeyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity | 

### Return type

[**ResponseDtoApiKeyAggregate**](ResponseDTOApiKeyAggregate.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListApiKeys**
> ResponseDtoListApiKey ListApiKeys(ctx, accountIdentifier, apiKeyType, parentIdentifier, optional)
Fetches the list of API Keys corresponding to the request's filter criteria.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
  **apiKeyType** | **string**| This is the API Key type like Personal Access Key or Service Account Key. | 
  **parentIdentifier** | **string**| ID of API key&#x27;s Parent Service Account | 
 **optional** | ***ApiKeyApiListApiKeysOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApiKeyApiListApiKeysOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity | 
 **identifiers** | [**optional.Interface of []string**](string.md)| This is the list of API Key IDs. Details specific to these IDs would be fetched. | 

### Return type

[**ResponseDtoListApiKey**](ResponseDTOListApiKey.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListApiKeys1**
> ResponseDtoPageResponseApiKeyAggregate ListApiKeys1(ctx, accountIdentifier, apiKeyType, parentIdentifier, optional)
Fetches the list of Aggregated API Keys corresponding to the request's filter criteria.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
  **apiKeyType** | **string**| This is the API Key type like Personal Access Key or Service Account Key. | 
  **parentIdentifier** | **string**| ID of API key&#x27;s Parent Service Account | 
 **optional** | ***ApiKeyApiListApiKeys1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApiKeyApiListApiKeys1Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity | 
 **identifiers** | [**optional.Interface of []string**](string.md)| This is the list of API Key IDs. Details specific to these IDs would be fetched. | 
 **pageIndex** | **optional.Int32**| Indicates the number of pages. Results for these pages will be retrieved. | [default to 0]
 **pageSize** | **optional.Int32**| The number of the elements to fetch | [default to 50]
 **sortOrders** | [**optional.Interface of []SortOrder**](SortOrder.md)| Sort criteria for the elements. | 
 **searchTerm** | **optional.String**| This would be used to filter API keys. Any API key having the specified string in its Name, ID and Tag would be filtered. | 

### Return type

[**ResponseDtoPageResponseApiKeyAggregate**](ResponseDTOPageResponseApiKeyAggregate.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateApiKey**
> ResponseDtoApiKey UpdateApiKey(ctx, identifier, optional)
Updates API Key for the provided ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| This is the API key ID | 
 **optional** | ***ApiKeyApiUpdateApiKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApiKeyApiUpdateApiKeyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ApiKey**](ApiKey.md)|  | 

### Return type

[**ResponseDtoApiKey**](ResponseDTOApiKey.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml, text/plain
 - **Accept**: application/json, application/yaml, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

