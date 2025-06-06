# nextgen{{classname}}

All URIs are relative to *https://app.harness.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloneDashboard**](DashboardsApi.md#CloneDashboard) | **Post** /dashboard/clone | Clones a Dashboard
[**DeleteDashboard**](DashboardsApi.md#DeleteDashboard) | **Delete** /dashboard/remove | Deletes a Dashboard
[**GetDashboard**](DashboardsApi.md#GetDashboard) | **Get** /dashboard/dashboards/{dashboard_id} | Gets the details of a Dashboard
[**UpdateDashboard**](DashboardsApi.md#UpdateDashboard) | **Patch** /dashboard/ | Updates a Dashboard

# **CloneDashboard**
> ClonedDashboardResponse CloneDashboard(ctx, body, optional)


Clone a dashboard.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CloneDashboardRequestBody**](CloneDashboardRequestBody.md)| Clone a Dashboard | 
 **optional** | ***DashboardsApiCloneDashboardOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DashboardsApiCloneDashboardOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountId** | **optional.**|  | 

### Return type

[**ClonedDashboardResponse**](ClonedDashboardResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDashboard**
> DeleteDashboardResponse DeleteDashboard(ctx, body, folderId, optional)


Delete a dashboard.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DeleteDashboardRequest**](DeleteDashboardRequest.md)| Delete a Dashboard by ID. | 
  **folderId** | **string**|  | 
 **optional** | ***DashboardsApiDeleteDashboardOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DashboardsApiDeleteDashboardOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accountId** | **optional.**|  | 

### Return type

[**DeleteDashboardResponse**](DeleteDashboardResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDashboard**
> GetDashboardResponse GetDashboard(ctx, dashboardId, optional)


Get all details of a dashboard by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dashboardId** | **string**|  | 
 **optional** | ***DashboardsApiGetDashboardOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DashboardsApiGetDashboardOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountId** | **optional.String**|  | 

### Return type

[**GetDashboardResponse**](GetDashboardResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDashboard**
> UpdateDashboardResponse UpdateDashboard(ctx, body, optional)


Update a dashboards name, tags or folder.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateDashboardRequest**](CreateDashboardRequest.md)| Update dashboard fields. | 
 **optional** | ***DashboardsApiUpdateDashboardOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DashboardsApiUpdateDashboardOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountId** | **optional.**|  | 

### Return type

[**UpdateDashboardResponse**](UpdateDashboardResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

