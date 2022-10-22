# nextgen{{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GnuPGKeyServiceListGPGKeys**](GPGKeysApi.md#GnuPGKeyServiceListGPGKeys) | **Get** /gitops/api/api/v1/gpgkeys | List all available repository certificates

# **GnuPGKeyServiceListGPGKeys**
> Servicev1GnuPgPublicKeyList GnuPGKeyServiceListGPGKeys(ctx, accountIdentifier, optional)
List all available repository certificates

List all available repository certificates

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***GPGKeysApiGnuPGKeyServiceListGPGKeysOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GPGKeysApiGnuPGKeyServiceListGPGKeysOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **gnuPG** | **optional.String**|  | 
 **searchTerm** | **optional.String**|  | 
 **pageSize** | **optional.Int32**|  | 
 **pageIndex** | **optional.Int32**|  | 
 **agentIdentifier** | **optional.String**| Agent identifier for entity. | 

### Return type

[**Servicev1GnuPgPublicKeyList**](servicev1GnuPGPublicKeyList.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

