# nextgen{{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTarget**](TargetsApi.md#CreateTarget) | **Post** /cf/admin/targets | Creates a Target
[**DeleteTarget**](TargetsApi.md#DeleteTarget) | **Delete** /cf/admin/targets/{identifier} | Deletes a Target
[**GetAllTargets**](TargetsApi.md#GetAllTargets) | **Get** /cf/admin/targets | Returns all Targets
[**GetTarget**](TargetsApi.md#GetTarget) | **Get** /cf/admin/targets/{identifier} | Returns details of a Target
[**GetTargetSegments**](TargetsApi.md#GetTargetSegments) | **Get** /cf/admin/targets/{identifier}/segments | Returns Target Groups for the given Target
[**ModifyTarget**](TargetsApi.md#ModifyTarget) | **Put** /cf/admin/targets/{identifier} | Modifies a Target
[**PatchTarget**](TargetsApi.md#PatchTarget) | **Patch** /cf/admin/targets/{identifier} | Updates a Target
[**UploadTargets**](TargetsApi.md#UploadTargets) | **Post** /cf/admin/targets/upload | Add Target details

# **CreateTarget**
> CreateTarget(ctx, body, accountIdentifier, orgIdentifier)
Creates a Target

Create Targets for the given identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Target**](Target.md)|  | 
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

# **DeleteTarget**
> DeleteTarget(ctx, identifier, accountIdentifier, orgIdentifier, projectIdentifier, environmentIdentifier)
Deletes a Target

Deletes a Target for the given identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| Unique identifier for the object in the API. | 
  **accountIdentifier** | **string**| Account Identifier | 
  **orgIdentifier** | **string**| Organization Identifier | 
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

# **GetAllTargets**
> Targets GetAllTargets(ctx, accountIdentifier, orgIdentifier, projectIdentifier, environmentIdentifier, optional)
Returns all Targets

Returns all the Targets for the given Account ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier | 
  **orgIdentifier** | **string**| Organization Identifier | 
  **projectIdentifier** | **string**| The Project identifier | 
  **environmentIdentifier** | **string**| Environment Identifier | 
 **optional** | ***TargetsApiGetAllTargetsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TargetsApiGetAllTargetsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **pageNumber** | **optional.Int32**| PageNumber | 
 **pageSize** | **optional.Int32**| PageSize | 
 **sortOrder** | **optional.String**| SortOrder | 
 **sortByField** | **optional.String**| SortByField | 
 **targetName** | **optional.String**| Name of the target | 
 **targetIdentifier** | **optional.String**| Identifier of the target | 

### Return type

[**Targets**](Targets.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTarget**
> Target GetTarget(ctx, identifier, accountIdentifier, orgIdentifier, projectIdentifier, environmentIdentifier)
Returns details of a Target

Returns details of a Target for the given identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| Unique identifier for the object in the API. | 
  **accountIdentifier** | **string**| Account Identifier | 
  **orgIdentifier** | **string**| Organization Identifier | 
  **projectIdentifier** | **string**| The Project identifier | 
  **environmentIdentifier** | **string**| Environment Identifier | 

### Return type

[**Target**](Target.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTargetSegments**
> TargetDetail GetTargetSegments(ctx, identifier, accountIdentifier, orgIdentifier, projectIdentifier, environmentIdentifier)
Returns Target Groups for the given Target

Returns the Target Groups that the specified Target belongs to.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| Unique identifier for the object in the API. | 
  **accountIdentifier** | **string**| Account Identifier | 
  **orgIdentifier** | **string**| Organization Identifier | 
  **projectIdentifier** | **string**| The Project identifier | 
  **environmentIdentifier** | **string**| Environment Identifier | 

### Return type

[**TargetDetail**](TargetDetail.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ModifyTarget**
> Target ModifyTarget(ctx, body, identifier, accountIdentifier, orgIdentifier, projectIdentifier, environmentIdentifier)
Modifies a Target

Modifies a Target for the given account identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Target**](Target.md)|  | 
  **identifier** | **string**| Unique identifier for the object in the API. | 
  **accountIdentifier** | **string**| Account Identifier | 
  **orgIdentifier** | **string**| Organization Identifier | 
  **projectIdentifier** | **string**| The Project identifier | 
  **environmentIdentifier** | **string**| Environment Identifier | 

### Return type

[**Target**](Target.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchTarget**
> Target PatchTarget(ctx, accountIdentifier, orgIdentifier, projectIdentifier, environmentIdentifier, identifier, optional)
Updates a Target

Updates a Target for the given identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier | 
  **orgIdentifier** | **string**| Organization Identifier | 
  **projectIdentifier** | **string**| The Project identifier | 
  **environmentIdentifier** | **string**| Environment Identifier | 
  **identifier** | **string**| Unique identifier for the object in the API. | 
 **optional** | ***TargetsApiPatchTargetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TargetsApiPatchTargetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **body** | [**optional.Interface of GitSyncPatchOperation**](GitSyncPatchOperation.md)|  | 

### Return type

[**Target**](Target.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadTargets**
> UploadTargets(ctx, accountIdentifier, orgIdentifier, projectIdentifier, environmentIdentifier, optional)
Add Target details

Add targets by uploading a CSV file

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier | 
  **orgIdentifier** | **string**| Organization Identifier | 
  **projectIdentifier** | **string**| The Project identifier | 
  **environmentIdentifier** | **string**| Environment Identifier | 
 **optional** | ***TargetsApiUploadTargetsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TargetsApiUploadTargetsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **fileName** | **optional.Interface of *os.File****optional.**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

