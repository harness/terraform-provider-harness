# nextgen{{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFeatureFlag**](FeatureFlagsApi.md#CreateFeatureFlag) | **Post** /cf/admin/features | Creates a Feature Flag
[**DeleteFeatureFlag**](FeatureFlagsApi.md#DeleteFeatureFlag) | **Delete** /cf/admin/features/{identifier} | Delete a Feature Flag
[**GetAllFeatures**](FeatureFlagsApi.md#GetAllFeatures) | **Get** /cf/admin/features | Returns all Feature Flags for the project
[**GetFeatureFlag**](FeatureFlagsApi.md#GetFeatureFlag) | **Get** /cf/admin/features/{identifier} | Returns a Feature Flag
[**PatchFeature**](FeatureFlagsApi.md#PatchFeature) | **Patch** /cf/admin/features/{identifier} | Updates a Feature Flag

# **CreateFeatureFlag**
> FeatureResponseMetadata CreateFeatureFlag(ctx, accountIdentifier, orgIdentifier, optional)
Creates a Feature Flag

Creates a Feature Flag in the Project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier | 
  **orgIdentifier** | **string**| Organization Identifier | 
 **optional** | ***FeatureFlagsApiCreateFeatureFlagOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FeatureFlagsApiCreateFeatureFlagOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of interface{}**](interface{}.md)|  | 

### Return type

[**FeatureResponseMetadata**](FeatureResponseMetadata.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFeatureFlag**
> DeleteFeatureFlag(ctx, identifier, accountIdentifier, orgIdentifier, projectIdentifier, optional)
Delete a Feature Flag

Delete Feature Flag for the given identifier and account ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| Unique identifier for the object in the API. | 
  **accountIdentifier** | **string**| Account Identifier | 
  **orgIdentifier** | **string**| Organization Identifier | 
  **projectIdentifier** | **string**| The Project identifier | 
 **optional** | ***FeatureFlagsApiDeleteFeatureFlagOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FeatureFlagsApiDeleteFeatureFlagOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **commitMsg** | **optional.String**| Git commit message | 

### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllFeatures**
> Features GetAllFeatures(ctx, accountIdentifier, orgIdentifier, projectIdentifier, optional)
Returns all Feature Flags for the project

Returns all the Feature Flag details for the given project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier | 
  **orgIdentifier** | **string**| Organization Identifier | 
  **projectIdentifier** | **string**| The Project identifier | 
 **optional** | ***FeatureFlagsApiGetAllFeaturesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FeatureFlagsApiGetAllFeaturesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **environmentIdentifier** | **optional.String**| Environment | 
 **pageNumber** | **optional.Int32**| PageNumber | 
 **pageSize** | **optional.Int32**| PageSize | 
 **sortOrder** | **optional.String**| SortOrder | 
 **sortByField** | **optional.String**| SortByField | 
 **name** | **optional.String**| Name of the field | 
 **identifier** | **optional.String**| Identifier of the field | 
 **archived** | **optional.Bool**| Status of the feature flag | 
 **kind** | **optional.String**| Kind of the feature flag | 
 **targetIdentifier** | **optional.String**| Identifier of a target | 
 **targetIdentifierFilter** | **optional.String**| Identifier of the target to filter on | 
 **metrics** | **optional.Bool**| Parameter to indicate if metrics data is requested in response | 
 **featureIdentifiers** | **optional.String**| Comma separated identifiers for multiple Features | 
 **excludedFeatures** | **optional.String**| Comma separated identifiers to exclude from the response | 
 **status** | **optional.String**| Filter for flags based on their status (active,never-requested,recently-accessed,potentially-stale) | 
 **lifetime** | **optional.String**| Filter for flags based on their lifetime (permanent/temporary) | 
 **enabled** | **optional.Bool**| Filter for flags based on if they are enabled or disabled | 
 **flagCounts** | **optional.Bool**| Returns counts for the different types of flags e.g num active, potentially-stale, recently-accessed etc | 

### Return type

[**Features**](Features.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFeatureFlag**
> Feature GetFeatureFlag(ctx, identifier, accountIdentifier, orgIdentifier, projectIdentifier, optional)
Returns a Feature Flag

Returns details such as Variation name, identifier etc for the given Feature Flag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| Unique identifier for the object in the API. | 
  **accountIdentifier** | **string**| Account Identifier | 
  **orgIdentifier** | **string**| Organization Identifier | 
  **projectIdentifier** | **string**| The Project identifier | 
 **optional** | ***FeatureFlagsApiGetFeatureFlagOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FeatureFlagsApiGetFeatureFlagOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **environmentIdentifier** | **optional.String**| Environment | 

### Return type

[**Feature**](Feature.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchFeature**
> FeatureResponseMetadata PatchFeature(ctx, identifier, accountIdentifier, orgIdentifier, projectIdentifier, optional)
Updates a Feature Flag

This operation is used to modify a Feature Flag.  The request body can include one or more instructions that can modify flag attributes such as the state (off|on), the variations that are returned and serving rules. For example if you want to turn a flag off you can use this opeartion and send the setFeatureFlagState  {   \"kind\": \"setFeatureFlagState\",   \"parameters\": {     \"state\": \"off\"   } } 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| Unique identifier for the object in the API. | 
  **accountIdentifier** | **string**| Account Identifier | 
  **orgIdentifier** | **string**| Organization Identifier | 
  **projectIdentifier** | **string**| The Project identifier | 
 **optional** | ***FeatureFlagsApiPatchFeatureOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FeatureFlagsApiPatchFeatureOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | [**optional.Interface of GitSyncPatchOperation**](GitSyncPatchOperation.md)|  | 
 **environmentIdentifier** | **optional.**| Environment | 

### Return type

[**FeatureResponseMetadata**](FeatureResponseMetadata.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

