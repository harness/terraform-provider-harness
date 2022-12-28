# {{classname}}

All URIs are relative to *https://{{host}}/cv/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SaveSLODataNg**](SloApi.md#SaveSLODataNg) | **Post** /account/{accountIdentifier}/org/{orgIdentifier}/project/{projectIdentifier}/slo/v2 | saves slo data
[**GetServiceLevelObjectiveNg**](SloApi.md#GetServiceLevelObjectiveNg) | **Get** /account/{accountIdentifier}/org/{orgIdentifier}/project/{projectIdentifier}/slo/v2/identifier/{identifier} | get service level objective data
[**UpdateSLODataNg**](SloApi.md#UpdateSLODataNg) | **Put** /account/{accountIdentifier}/org/{orgIdentifier}/project/{projectIdentifier}/slo/v2/identifier/{identifier} | update slo data
[**DeleteSLODataNg**](SloApi.md#DeleteSLODataNg) | **Delete** /account/{accountIdentifier}/org/{orgIdentifier}/project/{projectIdentifier}/slo/v2/identifier/{identifier} | delete slo data

# **SaveSLODataNg**
> RestResponseServiceLevelObjectiveV2Response SaveSLODataNg(ctx, accountIdentifier, orgIdentifier, projectIdentifier, optional)
saves slo data

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**|  |
  **orgIdentifier** | **string**|  |
  **projectIdentifier** | **string**|  |
 **optional** | ***SloApiSaveSLODataNgOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SloApiSaveSLODataNgOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



**body** | [**optional.Interface of ServiceLevelObjectiveV2Dto**](ServiceLevelObjectiveV2Dto.md)|  |

### Return type

[**RestResponseServiceLevelObjectiveV2Response**](RestResponseServiceLevelObjectiveV2Response.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceLevelObjectiveNg**
> RestResponseServiceLevelObjectiveV2Response GetServiceLevelObjectiveNg(ctx, accountIdentifier, orgIdentifier, projectIdentifier, identifier)
get service level objective data

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**|  | 
  **orgIdentifier** | **string**|  | 
  **projectIdentifier** | **string**|  | 
  **identifier** | **string**|  | 

### Return type

[**RestResponseServiceLevelObjectiveV2Response**](RestResponseServiceLevelObjectiveV2Response.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSLODataNg**
> RestResponseServiceLevelObjectiveV2Response UpdateSLODataNg(ctx, accountIdentifier, orgIdentifier, projectIdentifier, identifier, optional)
update slo data

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**|  | 
  **orgIdentifier** | **string**|  | 
  **projectIdentifier** | **string**|  | 
  **identifier** | **string**|  | 
 **optional** | ***SloApiUpdateSLODataNgOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SloApiUpdateSLODataNgOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | [**optional.Interface of ServiceLevelObjectiveV2Dto**](ServiceLevelObjectiveV2Dto.md)|  | 

### Return type

[**RestResponseServiceLevelObjectiveV2Response**](RestResponseServiceLevelObjectiveV2Response.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSLODataNg**
> RestResponseBoolean DeleteSLODataNg(ctx, accountIdentifier, orgIdentifier, projectIdentifier, identifier)
delete slo data

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**|  |
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

