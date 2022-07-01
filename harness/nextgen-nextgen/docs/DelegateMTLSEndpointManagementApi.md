# nextgen{{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckDelegateMtlsEndpointDomainPrefixAvailability**](DelegateMTLSEndpointManagementApi.md#CheckDelegateMtlsEndpointDomainPrefixAvailability) | **Get** /ng/api/delegate-mtls/check-availability | Checks whether a given delegate mTLS endpoint domain prefix is available.
[**CreateDelegateMtlsEndpointForAccount**](DelegateMTLSEndpointManagementApi.md#CreateDelegateMtlsEndpointForAccount) | **Post** /ng/api/delegate-mtls/endpoint | Creates the delegate mTLS endpoint for an account.
[**DeleteDelegateMtlsEndpointForAccount**](DelegateMTLSEndpointManagementApi.md#DeleteDelegateMtlsEndpointForAccount) | **Delete** /ng/api/delegate-mtls/endpoint | Removes the delegate mTLS endpoint for an account.
[**GetDelegateMtlsEndpointForAccount**](DelegateMTLSEndpointManagementApi.md#GetDelegateMtlsEndpointForAccount) | **Get** /ng/api/delegate-mtls/endpoint | Gets the delegate mTLS endpoint for an account.
[**PatchDelegateMtlsEndpointForAccount**](DelegateMTLSEndpointManagementApi.md#PatchDelegateMtlsEndpointForAccount) | **Patch** /ng/api/delegate-mtls/endpoint | Updates selected properties of the existing delegate mTLS endpoint for an account.
[**UpdateDelegateMtlsEndpointForAccount**](DelegateMTLSEndpointManagementApi.md#UpdateDelegateMtlsEndpointForAccount) | **Put** /ng/api/delegate-mtls/endpoint | Updates the existing delegate mTLS endpoint for an account.

# **CheckDelegateMtlsEndpointDomainPrefixAvailability**
> RestResponseBoolean CheckDelegateMtlsEndpointDomainPrefixAvailability(ctx, accountIdentifier, domainPrefix)
Checks whether a given delegate mTLS endpoint domain prefix is available.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **domainPrefix** | **string**| The domain prefix to check. | 

### Return type

[**RestResponseBoolean**](RestResponseBoolean.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateDelegateMtlsEndpointForAccount**
> RestResponseDelegateMtlsEndpointDetails CreateDelegateMtlsEndpointForAccount(ctx, body, accountIdentifier)
Creates the delegate mTLS endpoint for an account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DelegateMtlsEndpointRequest**](DelegateMtlsEndpointRequest.md)| The details of the delegate mTLS endpoint to create. | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 

### Return type

[**RestResponseDelegateMtlsEndpointDetails**](RestResponseDelegateMtlsEndpointDetails.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDelegateMtlsEndpointForAccount**
> RestResponseBoolean DeleteDelegateMtlsEndpointForAccount(ctx, accountIdentifier)
Removes the delegate mTLS endpoint for an account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 

### Return type

[**RestResponseBoolean**](RestResponseBoolean.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDelegateMtlsEndpointForAccount**
> RestResponseDelegateMtlsEndpointDetails GetDelegateMtlsEndpointForAccount(ctx, accountIdentifier)
Gets the delegate mTLS endpoint for an account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 

### Return type

[**RestResponseDelegateMtlsEndpointDetails**](RestResponseDelegateMtlsEndpointDetails.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchDelegateMtlsEndpointForAccount**
> RestResponseDelegateMtlsEndpointDetails PatchDelegateMtlsEndpointForAccount(ctx, body, accountIdentifier)
Updates selected properties of the existing delegate mTLS endpoint for an account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DelegateMtlsEndpointRequest**](DelegateMtlsEndpointRequest.md)| A subset of the details to update for the delegate mTLS endpoint. | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 

### Return type

[**RestResponseDelegateMtlsEndpointDetails**](RestResponseDelegateMtlsEndpointDetails.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDelegateMtlsEndpointForAccount**
> RestResponseDelegateMtlsEndpointDetails UpdateDelegateMtlsEndpointForAccount(ctx, body, accountIdentifier)
Updates the existing delegate mTLS endpoint for an account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DelegateMtlsEndpointRequest**](DelegateMtlsEndpointRequest.md)| The details to update for the delegate mTLS endpoint. | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 

### Return type

[**RestResponseDelegateMtlsEndpointDetails**](RestResponseDelegateMtlsEndpointDetails.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

