# {{classname}}

All URIs are relative to */gateway/har/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllHarnessArtifacts**](SpacesApi.md#GetAllHarnessArtifacts) | **Get** /spaces/{space_ref}/artifacts | List Harness Artifacts
[**GetAllRegistries**](SpacesApi.md#GetAllRegistries) | **Get** /spaces/{space_ref}/registries | List registries
[**GetArtifactStatsForSpace**](SpacesApi.md#GetArtifactStatsForSpace) | **Get** /spaces/{space_ref}/artifact/stats | Get artifact stats
[**GetStorageDetails**](SpacesApi.md#GetStorageDetails) | **Get** /spaces/{space_ref}/details | Get storage details for given space

# **GetAllHarnessArtifacts**
> InlineResponse20025 GetAllHarnessArtifacts(ctx, spaceRef, optional)
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

[**InlineResponse20025**](inline_response_200_25.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllRegistries**
> InlineResponse20027 GetAllRegistries(ctx, spaceRef, optional)
List registries

Lists all the registries.

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

[**InlineResponse20027**](inline_response_200_27.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetArtifactStatsForSpace**
> InlineResponse2002 GetArtifactStatsForSpace(ctx, spaceRef, optional)
Get artifact stats

Get artifact stats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **spaceRef** | **string**| Unique space path. | 
 **optional** | ***SpacesApiGetArtifactStatsForSpaceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpacesApiGetArtifactStatsForSpaceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **optional.String**| Date. Format - MM/DD/YYYY | 
 **to** | **optional.String**| Date. Format - MM/DD/YYYY | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStorageDetails**
> InlineResponse20026 GetStorageDetails(ctx, spaceRef)
Get storage details for given space

Get storage details for given space

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **spaceRef** | **string**| Unique space path. | 

### Return type

[**InlineResponse20026**](inline_response_200_26.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

