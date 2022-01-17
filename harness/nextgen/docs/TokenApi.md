# {{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateToken**](TokenApi.md#CreateToken) | **Post** /ng/api/token | Creates a Token
[**DeleteToken**](TokenApi.md#DeleteToken) | **Delete** /ng/api/token/{identifier} | Deletes a Token by ID
[**ListAggregatedTokens**](TokenApi.md#ListAggregatedTokens) | **Get** /ng/api/token/aggregate | Fetches the list of Aggregated Tokens corresponding to the request&#x27;s filter criteria.
[**RotateToken**](TokenApi.md#RotateToken) | **Post** /ng/api/token/rotate/{identifier} | Rotates a Token by ID
[**UpdateToken**](TokenApi.md#UpdateToken) | **Put** /ng/api/token/{identifier} | Updates a Token by ID

# **CreateToken**
> ResponseDtoString CreateToken(ctx, optional)
Creates a Token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TokenApiCreateTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TokenApiCreateTokenOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of TokenDto**](TokenDto.md)|  | 

### Return type

[**ResponseDtoString**](ResponseDTOString.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml, text/plain
 - **Accept**: application/json, application/yaml, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteToken**
> ResponseDtoBoolean DeleteToken(ctx, identifier, accountIdentifier, apiKeyType, parentIdentifier, apiKeyIdentifier, optional)
Deletes a Token by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| Token ID | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
  **apiKeyType** | **string**| This is the API Key type like Personal Access Key or Service Account Key. | 
  **parentIdentifier** | **string**| ID of API key&#x27;s Parent Service Account | 
  **apiKeyIdentifier** | **string**| API key ID | 
 **optional** | ***TokenApiDeleteTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TokenApiDeleteTokenOpts struct
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

# **ListAggregatedTokens**
> ResponseDtoPageResponseTokenAggregateDto ListAggregatedTokens(ctx, accountIdentifier, apiKeyType, parentIdentifier, apiKeyIdentifier, optional)
Fetches the list of Aggregated Tokens corresponding to the request's filter criteria.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
  **apiKeyType** | **string**| This is the API Key type like Personal Access Key or Service Account Key. | 
  **parentIdentifier** | **string**| ID of API key&#x27;s Parent Service Account | 
  **apiKeyIdentifier** | **string**| API key ID | 
 **optional** | ***TokenApiListAggregatedTokensOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TokenApiListAggregatedTokensOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity | 
 **identifiers** | [**optional.Interface of []string**](string.md)| This is the list of Token IDs. Details specific to these IDs would be fetched. | 
 **pageIndex** | **optional.Int32**| Indicates the number of pages. Results for these pages will be retrieved. | [default to 0]
 **pageSize** | **optional.Int32**| The number of the elements to fetch | [default to 50]
 **sortOrders** | [**optional.Interface of []SortOrder**](SortOrder.md)| Sort criteria for the elements. | 
 **searchTerm** | **optional.String**| This would be used to filter Tokens. Any Token having the specified string in its Name, ID and Tag would be filtered. | 

### Return type

[**ResponseDtoPageResponseTokenAggregateDto**](ResponseDTOPageResponseTokenAggregateDTO.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RotateToken**
> ResponseDtoString RotateToken(ctx, identifier, accountIdentifier, apiKeyType, parentIdentifier, apiKeyIdentifier, optional)
Rotates a Token by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| Token Identifier | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
  **apiKeyType** | **string**| This is the API Key type like Personal Access Key or Service Account Key. | 
  **parentIdentifier** | **string**| ID of API key&#x27;s Parent Service Account | 
  **apiKeyIdentifier** | **string**| API key ID | 
 **optional** | ***TokenApiRotateTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TokenApiRotateTokenOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **rotateTimestamp** | **optional.Int64**| Time stamp when the Token is to be rotated | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity | 

### Return type

[**ResponseDtoString**](ResponseDTOString.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateToken**
> ResponseDtoTokenDto UpdateToken(ctx, identifier, optional)
Updates a Token by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| Token ID | 
 **optional** | ***TokenApiUpdateTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TokenApiUpdateTokenOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of TokenDto**](TokenDto.md)|  | 

### Return type

[**ResponseDtoTokenDto**](ResponseDTOTokenDTO.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml, text/plain
 - **Accept**: application/json, application/yaml, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

