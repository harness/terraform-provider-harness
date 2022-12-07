# {{classname}}

All URIs are relative to *https://app.harness.io/gateway*
 
Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMonitoredService**](MonitoredServiceApi.md#DeleteMonitoredService) | **Delete** /monitored-service/{identifier} | delete monitored service data 
[**GetMonitoredService**](MonitoredServiceApi.md#GetMonitoredService) | **Get** /monitored-service/{identifier} | get monitored service data 
[**SaveMonitoredService**](MonitoredServiceApi.md#SaveMonitoredService) | **Post** /monitored-service | saves monitored service data
[**UpdateMonitoredService**](MonitoredServiceApi.md#UpdateMonitoredService) | **Put** /monitored-service/{identifier} | updates monitored service data

# **DeleteMonitoredService**
> RestResponseBoolean DeleteMonitoredService(ctx, accountId, orgIdentifier, projectIdentifier, identifier)
delete monitored service data 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**|  | 
  **orgIdentifier** | **string**|  | 
  **projectIdentifier** | **string**|  | 
  **identifier** | **string**|  | 

### Return type

[**RestResponseBoolean**](RestResponseBoolean.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMonitoredService**
> ResponseMonitoredServiceResponse GetMonitoredService(ctx, identifier, accountId, orgIdentifier, projectIdentifier)
get monitored service data 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**|  | 
  **accountId** | **string**|  | 
  **orgIdentifier** | **string**|  | 
  **projectIdentifier** | **string**|  | 

### Return type

[**ResponseMonitoredServiceResponse**](ResponseMonitoredServiceResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SaveMonitoredService**
> RestResponseMonitoredServiceResponse SaveMonitoredService(ctx, accountId, optional)
saves monitored service data

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**|  | 
 **optional** | ***MonitoredServiceApiSaveMonitoredServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MonitoredServiceApiSaveMonitoredServiceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of MonitoredServiceDto**](MonitoredServiceDto.md)|  | 

### Return type

[**RestResponseMonitoredServiceResponse**](RestResponseMonitoredServiceResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMonitoredService**
> RestResponseMonitoredServiceResponse UpdateMonitoredService(ctx, identifier, accountId, optional)
updates monitored service data

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**|  | 
  **accountId** | **string**|  | 
 **optional** | ***MonitoredServiceApiUpdateMonitoredServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MonitoredServiceApiUpdateMonitoredServiceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of MonitoredServiceDto**](MonitoredServiceDto.md)|  | 

### Return type

[**RestResponseMonitoredServiceResponse**](RestResponseMonitoredServiceResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

