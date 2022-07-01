# nextgen{{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateReportSetting**](CloudCostPerspectiveReportsApi.md#CreateReportSetting) | **Post** /ccm/api/perspectiveReport/{accountIdentifier} | Create a schedule for a Report
[**DeleteReportSetting**](CloudCostPerspectiveReportsApi.md#DeleteReportSetting) | **Delete** /ccm/api/perspectiveReport/{accountIdentifier} | Delete cost Perspective report
[**GetReportSetting**](CloudCostPerspectiveReportsApi.md#GetReportSetting) | **Get** /ccm/api/perspectiveReport/{accountIdentifier} | Fetch details of a cost Report
[**UpdateReportSetting**](CloudCostPerspectiveReportsApi.md#UpdateReportSetting) | **Put** /ccm/api/perspectiveReport/{accountIdentifier} | Update a cost Perspective Report

# **CreateReportSetting**
> ResponseDtoListCeReportSchedule CreateReportSetting(ctx, body, accountIdentifier)
Create a schedule for a Report

Create a report schedule for the given Report ID or a Perspective ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CeReportSchedule**](CeReportSchedule.md)| CEReportSchedule object to be saved | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 

### Return type

[**ResponseDtoListCeReportSchedule**](ResponseDTOListCEReportSchedule.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteReportSetting**
> ResponseDtoString DeleteReportSetting(ctx, accountIdentifier, optional)
Delete cost Perspective report

Delete cost Perspective Report for the given Report ID or a Perspective ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***CloudCostPerspectiveReportsApiDeleteReportSettingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CloudCostPerspectiveReportsApiDeleteReportSettingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reportId** | **optional.String**| Unique identifier for the Report | 
 **perspectiveId** | **optional.String**| Unique identifier for the Perspective | 

### Return type

[**ResponseDtoString**](ResponseDTOString.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportSetting**
> ResponseDtoListCeReportSchedule GetReportSetting(ctx, accountIdentifier, optional)
Fetch details of a cost Report

Fetch cost Report details for the given Report ID or a Perspective ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***CloudCostPerspectiveReportsApiGetReportSettingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CloudCostPerspectiveReportsApiGetReportSettingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perspectiveId** | **optional.String**| Unique identifier for the Perspective | 
 **reportId** | **optional.String**| Unique identifier for the Report | 

### Return type

[**ResponseDtoListCeReportSchedule**](ResponseDTOListCEReportSchedule.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateReportSetting**
> ResponseDtoListCeReportSchedule UpdateReportSetting(ctx, body, accountIdentifier)
Update a cost Perspective Report

Update cost Perspective Reports.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CeReportSchedule**](CeReportSchedule.md)| CEReportSchedule object to be updated | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 

### Return type

[**ResponseDtoListCeReportSchedule**](ResponseDTOListCEReportSchedule.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

