# nextgen{{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnomalyFilterValues**](CloudCostAnomaliesApi.md#AnomalyFilterValues) | **Post** /ccm/api/anomaly/filter-values | Returns the list of distinct values for all the specified Anomaly fields.
[**GetAnomaliesSummary**](CloudCostAnomaliesApi.md#GetAnomaliesSummary) | **Post** /ccm/api/anomaly/summary | List Anomalies
[**ListAnomalies**](CloudCostAnomaliesApi.md#ListAnomalies) | **Post** /ccm/api/anomaly | List Anomalies
[**ListPerspectiveAnomalies**](CloudCostAnomaliesApi.md#ListPerspectiveAnomalies) | **Post** /ccm/api/anomaly/perspective/{perspectiveId} | List Anomalies for Perspective
[**ReportAnomalyFeedback**](CloudCostAnomaliesApi.md#ReportAnomalyFeedback) | **Put** /ccm/api/anomaly/feedback | Report Anomaly feedback

# **AnomalyFilterValues**
> ResponseDtoListFilterStats AnomalyFilterValues(ctx, body, accountIdentifier)
Returns the list of distinct values for all the specified Anomaly fields.

Returns the list of distinct values for all the specified Anomaly fields.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]string**](string.md)| List of Anomaly columns whose unique values will be fetched | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 

### Return type

[**ResponseDtoListFilterStats**](ResponseDTOListFilterStats.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAnomaliesSummary**
> ResponseDtoListAnomalySummary GetAnomaliesSummary(ctx, accountIdentifier, optional)
List Anomalies

Fetch the result of anomaly query

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***CloudCostAnomaliesApiGetAnomaliesSummaryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CloudCostAnomaliesApiGetAnomaliesSummaryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of AnomalyFilterProperties**](AnomalyFilterProperties.md)| Anomaly Filter Properties | 

### Return type

[**ResponseDtoListAnomalySummary**](ResponseDTOListAnomalySummary.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAnomalies**
> ResponseDtoListAnomalyData ListAnomalies(ctx, accountIdentifier, optional)
List Anomalies

Fetch the list of anomalies reported according to the filters applied

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***CloudCostAnomaliesApiListAnomaliesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CloudCostAnomaliesApiListAnomaliesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of AnomalyFilterProperties**](AnomalyFilterProperties.md)| Anomaly Filter Properties | 

### Return type

[**ResponseDtoListAnomalyData**](ResponseDTOListAnomalyData.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPerspectiveAnomalies**
> ResponseDtoListPerspectiveAnomalyData ListPerspectiveAnomalies(ctx, body, accountIdentifier, perspectiveId)
List Anomalies for Perspective

Fetch anomalies for perspective

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PerspectiveQueryDto**](PerspectiveQueryDto.md)| Perspective Query | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **perspectiveId** | **string**| Unique identifier for perspective | 

### Return type

[**ResponseDtoListPerspectiveAnomalyData**](ResponseDTOListPerspectiveAnomalyData.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReportAnomalyFeedback**
> ResponseDtoBoolean ReportAnomalyFeedback(ctx, body, accountIdentifier, anomalyId)
Report Anomaly feedback

Mark an anomaly as true/false anomaly

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AnomalyFeedback**](AnomalyFeedback.md)| Feedback | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **anomalyId** | **string**| Unique identifier for perspective | 

### Return type

[**ResponseDtoBoolean**](ResponseDTOBoolean.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

