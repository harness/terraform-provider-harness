# nextgen{{classname}}

All URIs are relative to *https://ng-manager-596c9d4b85-8dcb2/gateway/ng/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDelegateToken**](DelegateTokenResourceApi.md#CreateDelegateToken) | **Post** /delegate-token-ng | Creates Delegate Token.
[**GetCgDelegateTokens**](DelegateTokenResourceApi.md#GetCgDelegateTokens) | **Get** /delegate-token-ng | Retrieves Delegate Tokens by Account, Organization, Project and status.
[**GetDelegateGroupsUsingToken**](DelegateTokenResourceApi.md#GetDelegateGroupsUsingToken) | **Get** /delegate-token-ng/delegate-groups | Lists delegate groups that are using the specified delegate token.
[**RevokeCgDelegateToken**](DelegateTokenResourceApi.md#RevokeCgDelegateToken) | **Put** /delegate-token-ng | Revokes Delegate Token.

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
 **revokeAfter** | **optional.Int64**| Epoch time in milliseconds after which the token will be marked as revoked. There can be a delay of upto one hour from the epoch value provided and actual revoking of the token. | 

### Return type

[**RestResponseDelegateTokenDetails**](RestResponseDelegateTokenDetails.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCgDelegateTokens**
> RestResponseListDelegateTokenDetails GetCgDelegateTokens(ctx, accountIdentifier, optional)
Retrieves Delegate Tokens by Account, Organization, Project and status.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***DelegateTokenResourceApiGetCgDelegateTokensOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DelegateTokenResourceApiGetCgDelegateTokensOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| Name of Delegate Token (ACTIVE or REVOKED). | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **status** | **optional.String**| Status of Delegate Token (ACTIVE or REVOKED). If left empty both active and revoked tokens will be retrieved | 

### Return type

[**RestResponseListDelegateTokenDetails**](RestResponseListDelegateTokenDetails.md)

### Authorization

No authorization required

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

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RevokeCgDelegateToken**
> RestResponseDelegateTokenDetails RevokeCgDelegateToken(ctx, accountIdentifier, tokenName, optional)
Revokes Delegate Token.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **tokenName** | **string**| Delegate Token name | 
 **optional** | ***DelegateTokenResourceApiRevokeCgDelegateTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DelegateTokenResourceApiRevokeCgDelegateTokenOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 

### Return type

[**RestResponseDelegateTokenDetails**](RestResponseDelegateTokenDetails.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

