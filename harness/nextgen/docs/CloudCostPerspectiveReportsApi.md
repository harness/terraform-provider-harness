# {{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateReportSetting**](CloudCostPerspectiveReportsApi.md#CreateReportSetting) | **Post** /ccm/api/perspectiveReport/{accountIdentifier} | Create Report Setting
[**DeleteReportSetting**](CloudCostPerspectiveReportsApi.md#DeleteReportSetting) | **Delete** /ccm/api/perspectiveReport/{accountIdentifier} | Delete setting by Report identifier or by Perspective identifier
[**GetReportSetting**](CloudCostPerspectiveReportsApi.md#GetReportSetting) | **Get** /ccm/api/perspectiveReport/{accountIdentifier} | Get Reports by Report identifier or by Perspective identifier
[**UpdateReportSetting**](CloudCostPerspectiveReportsApi.md#UpdateReportSetting) | **Put** /ccm/api/perspectiveReport/{accountIdentifier} | Update perspective reports

# **CreateReportSetting**
> ResponseDtoListCeReportSchedule CreateReportSetting(ctx, body, accountIdentifier)
Create Report Setting

Create setting by Report identifier or by Perspective identifier, by sending CEReportSchedule as request body

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CeReportSchedule**](CeReportSchedule.md)| CEReportSchedule object to be saved | 
  **accountIdentifier** | **string**| Account Identifier for the entity | 

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
Delete setting by Report identifier or by Perspective identifier

Delete setting by Report identifier or by Perspective identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the entity | 
 **optional** | ***CloudCostPerspectiveReportsApiDeleteReportSettingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CloudCostPerspectiveReportsApiDeleteReportSettingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reportId** | **optional.String**| The Report Identifier | 
 **perspectiveId** | **optional.String**| The Perspective Identifier | 

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
Get Reports by Report identifier or by Perspective identifier

Get Reports by Report identifier or by Perspective identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the entity | 
 **optional** | ***CloudCostPerspectiveReportsApiGetReportSettingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CloudCostPerspectiveReportsApiGetReportSettingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perspectiveId** | **optional.String**| The identifier of the Perspective | 
 **reportId** | **optional.String**| The identifier of the Report | 

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
Update perspective reports

Update perspective reports by sending CEReportSchedule as request body

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CeReportSchedule**](CeReportSchedule.md)| CEReportSchedule object to be updated | 
  **accountIdentifier** | **string**| Account Identifier for the entity | 

### Return type

[**ResponseDtoListCeReportSchedule**](ResponseDTOListCEReportSchedule.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

