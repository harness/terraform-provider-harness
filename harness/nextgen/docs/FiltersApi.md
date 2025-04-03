# nextgen{{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FilterServiceCreate**](FiltersApi.md#FilterServiceCreate) | **Post** /api/filters | Create a filter
[**FilterServiceDelete**](FiltersApi.md#FilterServiceDelete) | **Delete** /api/filters/{identifier} | Delete deletes a filter
[**FilterServiceGet**](FiltersApi.md#FilterServiceGet) | **Get** /api/filters/{identifier} | Get get filter details
[**FilterServiceList**](FiltersApi.md#FilterServiceList) | **Get** /api/filters | List filters
[**FilterServiceUpdate**](FiltersApi.md#FilterServiceUpdate) | **Put** /api/filters | Update a filter

# **FilterServiceCreate**
> V1Filter FilterServiceCreate(ctx, body, optional)
Create a filter

CreateFilter creates a filter

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1Filter**](V1Filter.md)|  | 
 **optional** | ***FiltersApiFilterServiceCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FiltersApiFilterServiceCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.**| Account Identifier for the Entity. | 

### Return type

[**V1Filter**](v1Filter.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FilterServiceDelete**
> bool FilterServiceDelete(ctx, identifier, optional)
Delete deletes a filter

Delete filter.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| Identifier for the filter. | 
 **optional** | ***FiltersApiFilterServiceDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FiltersApiFilterServiceDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.String**| Account Identifier for the Entity. | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **filterType** | **optional.String**| Filter type. One of {APPLICATION} | [default to FILTER_TYPE_UNSET]

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FilterServiceGet**
> V1Filter FilterServiceGet(ctx, identifier, optional)
Get get filter details

Get filter details.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| Identifier for the filter. | 
 **optional** | ***FiltersApiFilterServiceGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FiltersApiFilterServiceGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.String**| Account Identifier for the Entity. | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **filterType** | **optional.String**| Filter type. One of {APPLICATION} | [default to FILTER_TYPE_UNSET]

### Return type

[**V1Filter**](v1Filter.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FilterServiceList**
> V1FilterList FilterServiceList(ctx, optional)
List filters

List returns a list of filters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FiltersApiFilterServiceListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FiltersApiFilterServiceListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountIdentifier** | **optional.String**| Account Identifier for the Entity. | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **searchTerm** | **optional.String**|  | 
 **filterType** | **optional.String**| Filter type. One of {APPLICATION} | [default to FILTER_TYPE_UNSET]
 **pageSize** | **optional.Int32**|  | 
 **pageIndex** | **optional.Int32**|  | 

### Return type

[**V1FilterList**](v1FilterList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FilterServiceUpdate**
> V1Filter FilterServiceUpdate(ctx, body, optional)
Update a filter

Update updates a filter

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1Filter**](V1Filter.md)|  | 
 **optional** | ***FiltersApiFilterServiceUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FiltersApiFilterServiceUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.**| Account Identifier for the Entity. | 

### Return type

[**V1Filter**](v1Filter.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

