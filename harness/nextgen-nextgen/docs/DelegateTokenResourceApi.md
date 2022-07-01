# nextgen{{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDelegateToken**](DelegateTokenResourceApi.md#CreateDelegateToken) | **Post** /ng/api/delegate-token-ng | Creates Delegate Token.
[**GetDelegateGroupsUsingToken**](DelegateTokenResourceApi.md#GetDelegateGroupsUsingToken) | **Get** /ng/api/delegate-token-ng/delegate-groups | Lists delegate groups that are using the specified delegate token.
[**GetDelegateTokens**](DelegateTokenResourceApi.md#GetDelegateTokens) | **Get** /ng/api/delegate-token-ng | Retrieves Delegate Tokens by Account, Organization, Project and status.
[**RevokeDelegateToken**](DelegateTokenResourceApi.md#RevokeDelegateToken) | **Put** /ng/api/delegate-token-ng | Revokes Delegate Token.

# **CreateDelegateToken**
> RestResponseDelegateTokenDetails CreateDelegateToken(ctx, accountIdentifier, tokenName, optional)
Creates Delegate Token.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **tokenName** | **string**| Delegate Token name | 
 **optional** | ***DelegateTokenResourceApiCreateDelegateTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DelegateTokenResourceApiCreateDelegateTokenOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 

### Return type

[**RestResponseDelegateTokenDetails**](RestResponseDelegateTokenDetails.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDelegateGroupsUsingToken**
> RestResponseDelegateGroupListing GetDelegateGroupsUsingToken(ctx, accountIdentifier, optional)
Lists delegate groups that are using the specified delegate token.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***DelegateTokenResourceApiGetDelegateGroupsUsingTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DelegateTokenResourceApiGetDelegateGroupsUsingTokenOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **delegateTokenName** | **optional.String**| Delegate Token name | 

### Return type

[**RestResponseDelegateGroupListing**](RestResponseDelegateGroupListing.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDelegateTokens**
> RestResponseListDelegateTokenDetails GetDelegateTokens(ctx, accountIdentifier, optional)
Retrieves Delegate Tokens by Account, Organization, Project and status.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***DelegateTokenResourceApiGetDelegateTokensOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DelegateTokenResourceApiGetDelegateTokensOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **status** | **optional.String**| Status of Delegate Token (ACTIVE or REVOKED). If left empty both active and revoked tokens will be retrieved | 

### Return type

[**RestResponseListDelegateTokenDetails**](RestResponseListDelegateTokenDetails.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RevokeDelegateToken**
> RestResponseDelegateTokenDetails RevokeDelegateToken(ctx, accountIdentifier, tokenName, optional)
Revokes Delegate Token.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **tokenName** | **string**| Delegate Token name | 
 **optional** | ***DelegateTokenResourceApiRevokeDelegateTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DelegateTokenResourceApiRevokeDelegateTokenOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 

### Return type

[**RestResponseDelegateTokenDetails**](RestResponseDelegateTokenDetails.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

