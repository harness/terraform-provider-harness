# {{classname}}

All URIs are relative to */gateway/har/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRegistry**](RegistriesApi.md#CreateRegistry) | **Post** /registry | Create Registry.
[**DeleteRegistry**](RegistriesApi.md#DeleteRegistry) | **Delete** /registry/{registry_ref}/+ | Delete a Registry
[**GetAllArtifactsByRegistry**](RegistriesApi.md#GetAllArtifactsByRegistry) | **Get** /registry/{registry_ref}/+/artifacts | List Artifacts for Registry
[**GetClientSetupDetails**](RegistriesApi.md#GetClientSetupDetails) | **Get** /registry/{registry_ref}/+/client-setup-details | Returns CLI Client Setup Details
[**GetRegistry**](RegistriesApi.md#GetRegistry) | **Get** /registry/{registry_ref}/+ | Returns Registry Details
[**ModifyRegistry**](RegistriesApi.md#ModifyRegistry) | **Put** /registry/{registry_ref}/+ | Updates a Registry

# **CreateRegistry**
> InlineResponse201 CreateRegistry(ctx, optional)
Create Registry.

Create a Registry.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RegistriesApiCreateRegistryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RegistriesApiCreateRegistryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of RegistryRequest**](RegistryRequest.md)| request for create and update registry | 
 **spaceRef** | **optional.**| Unique space path | 

### Return type

[**InlineResponse201**](inline_response_201.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRegistry**
> InlineResponse200 DeleteRegistry(ctx, registryRef)
Delete a Registry

Delete a Registry in the account for the given key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **registryRef** | **string**| Unique registry path. | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllArtifactsByRegistry**
> InlineResponse20017 GetAllArtifactsByRegistry(ctx, registryRef, optional)
List Artifacts for Registry

Lists all the Artifacts for Registry

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **registryRef** | **string**| Unique registry path. | 
 **optional** | ***RegistriesApiGetAllArtifactsByRegistryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RegistriesApiGetAllArtifactsByRegistryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **label** | [**optional.Interface of []string**](string.md)| Label. | 
 **page** | **optional.Int64**| Current page number | [default to 1]
 **size** | **optional.Int64**| Number of items per page | [default to 20]
 **sortOrder** | **optional.String**| sortOrder | 
 **sortField** | **optional.String**| sortField | 
 **searchTerm** | **optional.String**| search Term. | 

### Return type

[**InlineResponse20017**](inline_response_200_17.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClientSetupDetails**
> InlineResponse20018 GetClientSetupDetails(ctx, registryRef, optional)
Returns CLI Client Setup Details

Returns CLI Client Setup Details based on package type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **registryRef** | **string**| Unique registry path. | 
 **optional** | ***RegistriesApiGetClientSetupDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RegistriesApiGetClientSetupDetailsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **artifact** | **optional.String**| Artifat | 
 **version** | **optional.String**| Version | 

### Return type

[**InlineResponse20018**](inline_response_200_18.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRegistry**
> InlineResponse201 GetRegistry(ctx, registryRef)
Returns Registry Details

Returns Registry Details in the account for the given key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **registryRef** | **string**| Unique registry path. | 

### Return type

[**InlineResponse201**](inline_response_201.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ModifyRegistry**
> InlineResponse201 ModifyRegistry(ctx, registryRef, optional)
Updates a Registry

Updates a Registry in the account for the given key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **registryRef** | **string**| Unique registry path. | 
 **optional** | ***RegistriesApiModifyRegistryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RegistriesApiModifyRegistryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of RegistryRequest**](RegistryRequest.md)| request for create and update registry | 

### Return type

[**InlineResponse201**](inline_response_201.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

