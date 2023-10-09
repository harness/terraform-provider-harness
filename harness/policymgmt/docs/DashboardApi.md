# {{classname}}

All URIs are relative to *https://app.harness.io/gateway/pm/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DashboardMetrics**](DashboardApi.md#DashboardMetrics) | **Get** /api/v1/dashboard | 

# **DashboardMetrics**
> DashboardMetrics DashboardMetrics(ctx, optional)


Get metrics about policies, policy sets and evaluations

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DashboardApiDashboardMetricsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DashboardApiDashboardMetricsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **range_** | **optional.String**| The time period over which to aggregate dashboard data. Can be: 24 hours, 7 days or 30 days | [default to 30d]
 **accountIdentifier** | **optional.String**| Harness account ID | 
 **orgIdentifier** | **optional.String**| Harness organization ID | 
 **projectIdentifier** | **optional.String**| Harness project ID | 
 **xApiKey** | **optional.String**| Harness PAT key used to perform authorization | 

### Return type

[**DashboardMetrics**](DashboardMetrics.md)

### Authorization

[api_key_header_x-api-key](../README.md#api_key_header_x-api-key), [jwt_header_Authorization](../README.md#jwt_header_Authorization)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

