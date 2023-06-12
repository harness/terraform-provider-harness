# nextgen{{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServiceOverrideV2**](ServiceOverridesApi.md#CreateServiceOverrideV2) | **Post** /serviceOverrides | Create an ServiceOverride Entity
[**DeleteServiceOverrideV2**](ServiceOverridesApi.md#DeleteServiceOverrideV2) | **Delete** /serviceOverrides/{identifier} | Delete a Service Override entity
[**GetServiceOverrideListV2**](ServiceOverridesApi.md#GetServiceOverrideListV2) | **Get** /serviceOverrides/list | Gets Service Override List
[**GetServiceOverridesV2**](ServiceOverridesApi.md#GetServiceOverridesV2) | **Get** /serviceOverrides/{identifier} | Gets Service Overrides by Identifier
[**MigrateServiceOverride**](ServiceOverridesApi.md#MigrateServiceOverride) | **Post** /serviceOverrides/migrate | Migrate ServiceOverride to V2
[**MigrateServiceOverrideScoped**](ServiceOverridesApi.md#MigrateServiceOverrideScoped) | **Post** /serviceOverrides/migrateScope | Migrate ServiceOverride to V2 at one scope
[**UpdateServiceOverrideV2**](ServiceOverridesApi.md#UpdateServiceOverrideV2) | **Put** /serviceOverrides | Update an ServiceOverride Entity
[**UpsertServiceOverrideV2**](ServiceOverridesApi.md#UpsertServiceOverrideV2) | **Post** /serviceOverrides/upsert | Upsert an ServiceOverride Entity

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

# **GetServiceOverrideListV2**
> ResponsePageServiceOverridesResponseDtov2 GetServiceOverrideListV2(ctx, accountIdentifier, optional)
Gets Service Override List

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**|  | 
 **optional** | ***ServiceOverridesApiGetServiceOverrideListV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServiceOverridesApiGetServiceOverrideListV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**|  | [default to 0]
 **size** | **optional.Int32**|  | [default to 100]
 **orgIdentifier** | **optional.String**|  | 
 **projectIdentifier** | **optional.String**|  | 
 **type_** | **optional.String**|  | 

### Return type

[**ResponsePageServiceOverridesResponseDtov2**](ResponsePageServiceOverridesResponseDTOV2.md)

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

# **MigrateServiceOverride**
> ResponseServiceOverrideMigrationResponseDto MigrateServiceOverride(ctx, accountIdentifier, optional)
Migrate ServiceOverride to V2

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**|  | 
 **optional** | ***ServiceOverridesApiMigrateServiceOverrideOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServiceOverridesApiMigrateServiceOverrideOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgIdentifier** | **optional.String**|  | 
 **projectIdentifier** | **optional.String**|  | 

### Return type

[**ResponseServiceOverrideMigrationResponseDto**](ResponseServiceOverrideMigrationResponseDTO.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MigrateServiceOverrideScoped**
> ResponseServiceOverrideMigrationResponseDto MigrateServiceOverrideScoped(ctx, accountIdentifier, optional)
Migrate ServiceOverride to V2 at one scope

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**|  | 
 **optional** | ***ServiceOverridesApiMigrateServiceOverrideScopedOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServiceOverridesApiMigrateServiceOverrideScopedOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgIdentifier** | **optional.String**|  | 
 **projectIdentifier** | **optional.String**|  | 

### Return type

[**ResponseServiceOverrideMigrationResponseDto**](ResponseServiceOverrideMigrationResponseDTO.md)

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

# **UpsertServiceOverrideV2**
> ResponseServiceOverridesResponseDtov2 UpsertServiceOverrideV2(ctx, accountIdentifier, optional)
Upsert an ServiceOverride Entity

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**|  | 
 **optional** | ***ServiceOverridesApiUpsertServiceOverrideV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServiceOverridesApiUpsertServiceOverrideV2Opts struct
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

