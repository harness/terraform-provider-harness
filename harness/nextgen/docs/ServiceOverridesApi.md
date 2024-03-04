# nextgen{{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServiceOverrideV2**](ServiceOverridesApi.md#CreateServiceOverrideV2) | **Post** /serviceOverrides | Create an ServiceOverride Entity
[**DeleteServiceOverrideV2**](ServiceOverridesApi.md#DeleteServiceOverrideV2) | **Delete** /serviceOverrides/{identifier} | Delete a Service Override entity
[**GetServiceOverridesV2**](ServiceOverridesApi.md#GetServiceOverridesV2) | **Get** /serviceOverrides/get-with-yaml/{identifier} | Gets Service Overrides by Identifier
[**UpdateServiceOverrideV2**](ServiceOverridesApi.md#UpdateServiceOverrideV2) | **Put** /serviceOverrides | Update an ServiceOverride Entity[**ImportServiceOverrides**](ServiceOverridesApi.md#ImportServiceOverrides) | **Post** /serviceOverrides/import | import Service Overrides from remote
# **CreateServiceOverrideV2**
> ResponseServiceOverridesResponseDtov2 CreateServiceOverrideV2(ctx, accountIdentifier, optional)
Create an ServiceOverride Entity

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**|  | 
 **optional** | ***ServiceOverridesApiCreateServiceOverrideV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServiceOverridesApiCreateServiceOverrideV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ServiceOverrideRequestDtov2**](ServiceOverrideRequestDtov2.md)|  | 

### Return type

[**ResponseServiceOverridesResponseDtov2**](ResponseServiceOverridesResponseDTOV2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteServiceOverrideV2**
> ResponseBoolean DeleteServiceOverrideV2(ctx, identifier, accountIdentifier, optional)
Delete a Service Override entity

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**|  | 
  **accountIdentifier** | **string**|  | 
 **optional** | ***ServiceOverridesApiDeleteServiceOverrideV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServiceOverridesApiDeleteServiceOverrideV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.String**|  | 
 **projectIdentifier** | **optional.String**|  | 

### Return type

[**ResponseBoolean**](ResponseBoolean.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceOverridesV2**
> ResponseServiceOverridesResponseDtov2 GetServiceOverridesV2(ctx, identifier, accountIdentifier, optional)
Gets Service Overrides by Identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**|  | 
  **accountIdentifier** | **string**|  | 
 **optional** | ***ServiceOverridesApiGetServiceOverridesV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServiceOverridesApiGetServiceOverridesV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.String**|  | 
 **projectIdentifier** | **optional.String**|  | 

### Return type

[**ResponseServiceOverridesResponseDtov2**](ResponseServiceOverridesResponseDTOV2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateServiceOverrideV2**
> ResponseServiceOverridesResponseDtov2 UpdateServiceOverrideV2(ctx, accountIdentifier, optional)
Update an ServiceOverride Entity

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**|  | 
 **optional** | ***ServiceOverridesApiUpdateServiceOverrideV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServiceOverridesApiUpdateServiceOverrideV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ServiceOverrideRequestDtov2**](ServiceOverrideRequestDtov2.md)|  | 

### Return type

[**ResponseServiceOverridesResponseDtov2**](ResponseServiceOverridesResponseDTOV2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImportServiceOverrides**
> ResponseServiceOverrideImportResponseDto ImportServiceOverrides(ctx, optional)
import Service Overrides from remote

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ServiceOverridesApiImportServiceOverridesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServiceOverridesApiImportServiceOverridesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ServiceOverrideImportRequestDto**](ServiceOverrideImportRequestDto.md)|  | 
 **accountIdentifier** | **optional.**|  | 
 **connectorRef** | **optional.**|  | 
 **repoName** | **optional.**|  | 
 **branch** | **optional.**|  | 
 **filePath** | **optional.**|  | 
 **isForceImport** | **optional.**|  | [default to false]
 **isHarnessCodeRepo** | **optional.**|  | 

### Return type

[**ResponseServiceOverrideImportResponseDto**](ResponseServiceOverrideImportResponseDTO.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
