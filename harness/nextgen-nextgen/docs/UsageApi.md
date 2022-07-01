# nextgen{{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CcmgetCDLicenseUsageForServiceInstances**](UsageApi.md#CcmgetCDLicenseUsageForServiceInstances) | **Get** /ccm/api/usage/CD/serviceInstancesLicense | Gets License Usage By Module, Timestamp, and Account Identifier
[**CcmgetCDLicenseUsageForServices**](UsageApi.md#CcmgetCDLicenseUsageForServices) | **Get** /ccm/api/usage/CD/servicesLicense | Gets License Usage By Module, Timestamp, and Account Identifier
[**CcmgetLicenseUsage**](UsageApi.md#CcmgetLicenseUsage) | **Get** /ccm/api/usage/{module} | Gets License Usage By Module, Timestamp, and Account Identifier
[**GetCDLicenseUsageForServiceInstances**](UsageApi.md#GetCDLicenseUsageForServiceInstances) | **Get** /ng/api/usage/CD/serviceInstancesLicense | Gets License Usage By Module, Timestamp, and Account Identifier
[**GetCDLicenseUsageForServices**](UsageApi.md#GetCDLicenseUsageForServices) | **Get** /ng/api/usage/CD/servicesLicense | Gets License Usage By Module, Timestamp, and Account Identifier
[**GetLicenseUsage**](UsageApi.md#GetLicenseUsage) | **Get** /ng/api/usage/{module} | Gets License Usage By Module, Timestamp, and Account Identifier

# **CcmgetCDLicenseUsageForServiceInstances**
> ResponseDtoServiceInstanceUsageDto CcmgetCDLicenseUsageForServiceInstances(ctx, optional)
Gets License Usage By Module, Timestamp, and Account Identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsageApiCcmgetCDLicenseUsageForServiceInstancesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsageApiCcmgetCDLicenseUsageForServiceInstancesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountIdentifier** | **optional.String**| Account id to get the license usage. | 
 **timestamp** | **optional.Int64**|  | 

### Return type

[**ResponseDtoServiceInstanceUsageDto**](ResponseDTOServiceInstanceUsageDTO.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CcmgetCDLicenseUsageForServices**
> ResponseDtoServiceUsageDto CcmgetCDLicenseUsageForServices(ctx, optional)
Gets License Usage By Module, Timestamp, and Account Identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsageApiCcmgetCDLicenseUsageForServicesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsageApiCcmgetCDLicenseUsageForServicesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountIdentifier** | **optional.String**| Account id to get the license usage. | 
 **timestamp** | **optional.Int64**|  | 

### Return type

[**ResponseDtoServiceUsageDto**](ResponseDTOServiceUsageDTO.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CcmgetLicenseUsage**
> ResponseDtoLicenseUsage CcmgetLicenseUsage(ctx, module, optional)
Gets License Usage By Module, Timestamp, and Account Identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **module** | **string**| A Harness platform module. | 
 **optional** | ***UsageApiCcmgetLicenseUsageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsageApiCcmgetLicenseUsageOpts struct
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

# **GetCDLicenseUsageForServiceInstances**
> ResponseDtoServiceInstanceUsageDto GetCDLicenseUsageForServiceInstances(ctx, optional)
Gets License Usage By Module, Timestamp, and Account Identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsageApiGetCDLicenseUsageForServiceInstancesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsageApiGetCDLicenseUsageForServiceInstancesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountIdentifier** | **optional.String**| Account id to get the license usage. | 
 **timestamp** | **optional.Int64**|  | 

### Return type

[**ResponseDtoServiceInstanceUsageDto**](ResponseDTOServiceInstanceUsageDTO.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCDLicenseUsageForServices**
> ResponseDtoServiceUsageDto GetCDLicenseUsageForServices(ctx, optional)
Gets License Usage By Module, Timestamp, and Account Identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsageApiGetCDLicenseUsageForServicesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsageApiGetCDLicenseUsageForServicesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountIdentifier** | **optional.String**| Account id to get the license usage. | 
 **timestamp** | **optional.Int64**|  | 

### Return type

[**ResponseDtoServiceUsageDto**](ResponseDTOServiceUsageDTO.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

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

