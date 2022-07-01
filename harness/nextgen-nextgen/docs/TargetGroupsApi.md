# nextgen{{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSegment**](TargetGroupsApi.md#CreateSegment) | **Post** /cf/admin/segments | Creates a Target Group
[**DeleteSegment**](TargetGroupsApi.md#DeleteSegment) | **Delete** /cf/admin/segments/{identifier} | Deletes a Target Group
[**GetAllSegments**](TargetGroupsApi.md#GetAllSegments) | **Get** /cf/admin/segments | Returns all Target Groups
[**GetAvailableFlagsForSegment**](TargetGroupsApi.md#GetAvailableFlagsForSegment) | **Get** /cf/admin/segments/{identifier}/available_flags | Returns Feature Flags that are available to be added to the given Target Group
[**GetSegment**](TargetGroupsApi.md#GetSegment) | **Get** /cf/admin/segments/{identifier} | Returns Target Group details for the given identifier
[**GetSegmentFlags**](TargetGroupsApi.md#GetSegmentFlags) | **Get** /cf/admin/segments/{identifier}/flags | Returns Feature Flags in a Target Group
[**PatchSegment**](TargetGroupsApi.md#PatchSegment) | **Patch** /cf/admin/segments/{identifier} | Updates a Target Group

# **CreateSegment**
> CreateSegment(ctx, body, accountIdentifier, orgIdentifier)
Creates a Target Group

Creates a Target Group in the given Project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**interface{}**](interface{}.md)|  | 
  **accountIdentifier** | **string**| Account Identifier | 
  **orgIdentifier** | **string**| Organization Identifier | 

### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSegment**
> DeleteSegment(ctx, accountIdentifier, orgIdentifier, identifier, projectIdentifier, environmentIdentifier)
Deletes a Target Group

Deletes a Target Group for the given ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier | 
  **orgIdentifier** | **string**| Organization Identifier | 
  **identifier** | **string**| Unique identifier for the object in the API. | 
  **projectIdentifier** | **string**| The Project identifier | 
  **environmentIdentifier** | **string**| Environment Identifier | 

### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllSegments**
> Segments GetAllSegments(ctx, accountIdentifier, orgIdentifier, environmentIdentifier, projectIdentifier, optional)
Returns all Target Groups

Returns Target Group details for the given account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier | 
  **orgIdentifier** | **string**| Organization Identifier | 
  **environmentIdentifier** | **string**| Environment Identifier | 
  **projectIdentifier** | **string**| The Project identifier | 
 **optional** | ***TargetGroupsApiGetAllSegmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TargetGroupsApiGetAllSegmentsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **pageNumber** | **optional.Int32**| PageNumber | 
 **pageSize** | **optional.Int32**| PageSize | 
 **sortOrder** | **optional.String**| SortOrder | 
 **sortByField** | **optional.String**| SortByField | 
 **name** | **optional.String**| Name of the field | 
 **identifier** | **optional.String**| Identifier of the field | 

### Return type

[**Segments**](Segments.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAvailableFlagsForSegment**
> FlagBasicInfos GetAvailableFlagsForSegment(ctx, identifier, accountIdentifier, orgIdentifier, projectIdentifier, environmentIdentifier, optional)
Returns Feature Flags that are available to be added to the given Target Group

Returns the list of Feature Flags that the Target Group can be added to.  This list will exclude any Feature Flag that the Target Group is already part of.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| Unique identifier for the object in the API. | 
  **accountIdentifier** | **string**| Account Identifier | 
  **orgIdentifier** | **string**| Organization Identifier | 
  **projectIdentifier** | **string**| The Project identifier | 
  **environmentIdentifier** | **string**| Environment Identifier | 
 **optional** | ***TargetGroupsApiGetAvailableFlagsForSegmentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TargetGroupsApiGetAvailableFlagsForSegmentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **pageNumber** | **optional.Int32**| PageNumber | 
 **pageSize** | **optional.Int32**| PageSize | 
 **sortOrder** | **optional.String**| SortOrder | 
 **sortByField** | **optional.String**| SortByField | 
 **flagNameIdentifier** | **optional.String**| Identifier of the feature flag | 

### Return type

[**FlagBasicInfos**](FlagBasicInfos.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSegment**
> Segment GetSegment(ctx, accountIdentifier, orgIdentifier, identifier, projectIdentifier, environmentIdentifier)
Returns Target Group details for the given identifier

Returns Target Group details for the given ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier | 
  **orgIdentifier** | **string**| Organization Identifier | 
  **identifier** | **string**| Unique identifier for the object in the API. | 
  **projectIdentifier** | **string**| The Project identifier | 
  **environmentIdentifier** | **string**| Environment Identifier | 

### Return type

[**Segment**](Segment.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSegmentFlags**
> []SegmentFlag GetSegmentFlags(ctx, accountIdentifier, orgIdentifier, identifier, projectIdentifier, environmentIdentifier)
Returns Feature Flags in a Target Group

Returns the details of a Feature Flag in a Target Group for the given identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier | 
  **orgIdentifier** | **string**| Organization Identifier | 
  **identifier** | **string**| Unique identifier for the object in the API. | 
  **projectIdentifier** | **string**| The Project identifier | 
  **environmentIdentifier** | **string**| Environment Identifier | 

### Return type

[**[]SegmentFlag**](SegmentFlag.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchSegment**
> Segment PatchSegment(ctx, accountIdentifier, orgIdentifier, projectIdentifier, environmentIdentifier, identifier, optional)
Updates a Target Group

Updates a Target Group for the given identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier | 
  **orgIdentifier** | **string**| Organization Identifier | 
  **projectIdentifier** | **string**| The Project identifier | 
  **environmentIdentifier** | **string**| Environment Identifier | 
  **identifier** | **string**| Unique identifier for the object in the API. | 
 **optional** | ***TargetGroupsApiPatchSegmentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TargetGroupsApiPatchSegmentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **body** | [**optional.Interface of GitSyncPatchOperation**](GitSyncPatchOperation.md)|  | 

### Return type

[**Segment**](Segment.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

