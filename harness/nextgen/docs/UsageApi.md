# {{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLicenseUsage**](UsageApi.md#GetLicenseUsage) | **Get** /ng/api/usage/{module} | Gets License Usage By Module, Timestamp, and Account Identifier

# **GetLicenseUsage**
> ResponseDtoLicenseUsage GetLicenseUsage(ctx, module, optional)
Gets License Usage By Module, Timestamp, and Account Identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **module** | **string**| A Harness platform module. | 
 **optional** | ***UsageApiGetLicenseUsageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsageApiGetLicenseUsageOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.String**| Account id to get the license usage. | 
 **timestamp** | **optional.Int64**|  | 
 **cDLicenseType** | **optional.String**|  | 

### Return type

[**ResponseDtoLicenseUsage**](ResponseDTOLicenseUsage.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

