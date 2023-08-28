# {{classname}}

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDefaultMonitoredService**](MonitoredServicesApi.md#CreateDefaultMonitoredService) | **Post** /monitored-service/create-default | 
[**DeleteMonitoredService**](MonitoredServicesApi.md#DeleteMonitoredService) | **Delete** /monitored-service/{identifier} | Delete monitored service data
[**GetMonitoredService**](MonitoredServicesApi.md#GetMonitoredService) | **Get** /monitored-service/{identifier} | Get monitored service data
[**SaveMonitoredService**](MonitoredServicesApi.md#SaveMonitoredService) | **Post** /monitored-service | Saves monitored service data
[**UpdateMonitoredService**](MonitoredServicesApi.md#UpdateMonitoredService) | **Put** /monitored-service/{identifier} | Updates monitored service data

# **CreateDefaultMonitoredService**
> CreateDefaultMonitoredService(ctx, accountId, orgIdentifier, projectIdentifier, environmentIdentifier, serviceIdentifier)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Account Identifier for the Entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the Entity. | 
  **projectIdentifier** | **string**| Project Identifier for the Entity. | 
  **environmentIdentifier** | **string**|  | 
  **serviceIdentifier** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMonitoredService**
> RestResponseBoolean DeleteMonitoredService(ctx, accountId, orgIdentifier, projectIdentifier, identifier)
Delete monitored service data

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Account Identifier for the Entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the Entity. | 
  **projectIdentifier** | **string**| Project Identifier for the Entity. | 
  **identifier** | **string**| Identifier for the Entity. | 

### Return type

[**RestResponseBoolean**](RestResponseBoolean.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMonitoredService**
> ResponseDtoMonitoredServiceResponse GetMonitoredService(ctx, identifier, accountId, orgIdentifier, projectIdentifier)
Get monitored service data

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**|  | 
  **accountId** | **string**| Account Identifier for the Entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the Entity. | 
  **projectIdentifier** | **string**| Project Identifier for the Entity. | 

### Return type

[**ResponseDtoMonitoredServiceResponse**](ResponseDTOMonitoredServiceResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SaveMonitoredService**
> RestResponseMonitoredServiceResponse SaveMonitoredService(ctx, body, accountId)
Saves monitored service data

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MonitoredService**](MonitoredService.md)|  | 
  **accountId** | **string**|  | 

### Return type

[**RestResponseMonitoredServiceResponse**](RestResponseMonitoredServiceResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMonitoredService**
> RestResponseMonitoredServiceResponse UpdateMonitoredService(ctx, body, identifier, accountId)
Updates monitored service data

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MonitoredService**](MonitoredService.md)|  | 
  **identifier** | **string**| Identifier for the Entity. | 
  **accountId** | **string**| Account Identifier for the Entity. | 

### Return type

[**RestResponseMonitoredServiceResponse**](RestResponseMonitoredServiceResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

