# har{{classname}}

All URIs are relative to */gateway/har/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllHarnessArtifacts**](SpacesApi.md#GetAllHarnessArtifacts) | **Get** /spaces/{space_ref}/+/artifacts | List Harness Artifacts
[**GetAllRegistries**](SpacesApi.md#GetAllRegistries) | **Get** /spaces/{space_ref}/+/registries | List Registries

# **GetAllHarnessArtifacts**
> InlineResponse20016 GetAllHarnessArtifacts(ctx, spaceRef, optional)
List Harness Artifacts

Lists all the Harness Artifacts.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **spaceRef** | **string**| Unique space path. | 
 **optional** | ***SpacesApiGetAllHarnessArtifactsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpacesApiGetAllHarnessArtifactsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **regIdentifier** | [**optional.Interface of []string**](string.md)| Registry Identifier | 
 **page** | **optional.Int64**| Current page number | [default to 1]
 **size** | **optional.Int64**| Number of items per page | [default to 20]
 **sortOrder** | **optional.String**| sortOrder | 
 **sortField** | **optional.String**| sortField | 
 **searchTerm** | **optional.String**| search Term. | 
 **latestVersion** | **optional.Bool**| Latest Version Filter. | 
 **deployedArtifact** | **optional.Bool**| Deployed Artifact Filter. | 
 **packageType** | [**optional.Interface of []string**](string.md)| Registry Package Type | 

### Return type

[**InlineResponse20016**](inline_response_200_16.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllRegistries**
> InlineResponse20017 GetAllRegistries(ctx, spaceRef, optional)
List Registries

Lists all the Registries.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **spaceRef** | **string**| Unique space path. | 
 **optional** | ***SpacesApiGetAllRegistriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpacesApiGetAllRegistriesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **packageType** | [**optional.Interface of []string**](string.md)| Registry Package Type | 
 **type_** | **optional.String**| Registry Type | 
 **page** | **optional.Int64**| Current page number | [default to 1]
 **size** | **optional.Int64**| Number of items per page | [default to 20]
 **sortOrder** | **optional.String**| sortOrder | 
 **sortField** | **optional.String**| sortField | 
 **searchTerm** | **optional.String**| search Term. | 
 **recursive** | **optional.Bool**| Whether to list registries recursively. | [default to false]

### Return type

[**InlineResponse20017**](inline_response_200_17.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

