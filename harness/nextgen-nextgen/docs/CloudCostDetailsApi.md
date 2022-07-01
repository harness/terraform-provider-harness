# nextgen{{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Costdetailoverview**](CloudCostDetailsApi.md#Costdetailoverview) | **Post** /ccm/api/costdetails/overview | Returns an overview of the cost
[**Costdetailtabular**](CloudCostDetailsApi.md#Costdetailtabular) | **Post** /ccm/api/costdetails/tabularformat | Returns cost details in a tabular format
[**Costdetailttimeseries**](CloudCostDetailsApi.md#Costdetailttimeseries) | **Post** /ccm/api/costdetails/timeseriesformat | Returns cost details in a time series format

# **Costdetailoverview**
> ResponseDtoCostOverview Costdetailoverview(ctx, accountIdentifier, perspectiveId, optional)
Returns an overview of the cost

Returns total cost, cost trend, and the time period based on the specified query parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **perspectiveId** | **string**| Perspective identifier of the cost details | 
 **optional** | ***CloudCostDetailsApiCostdetailoverviewOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CloudCostDetailsApiCostdetailoverviewOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of CostDetailsQueryParams**](CostDetailsQueryParams.md)| Cost details query parameters. | 
 **startTime** | **optional.**| Start time of the cost details. Should use org.joda.time.DateTime parsable format. Example, &#x27;2022-01-31&#x27;, &#x27;2022-01-31T07:54Z&#x27; or &#x27;2022-01-31T07:54:51.264Z&#x27;.  Defaults to Today - 7days | 
 **endTime** | **optional.**| End time of the cost details. Should use org.joda.time.DateTime parsable format. Example, &#x27;2022-01-31&#x27;, &#x27;2022-01-31T07:54Z&#x27; or &#x27;2022-01-31T07:54:51.264Z&#x27;.  Defaults to Today | 

### Return type

[**ResponseDtoCostOverview**](ResponseDTOCostOverview.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Costdetailtabular**
> ResponseDtoPerspectiveEntityStatsData Costdetailtabular(ctx, body, accountIdentifier, perspectiveId, optional)
Returns cost details in a tabular format

Returns cost details in a tabular format based on the specified query parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CostDetailsQueryParams**](CostDetailsQueryParams.md)| Cost details query parameters. | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **perspectiveId** | **string**| Perspective identifier of the cost details | 
 **optional** | ***CloudCostDetailsApiCostdetailtabularOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CloudCostDetailsApiCostdetailtabularOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **startTime** | **optional.**| Start time of the cost details. Should use org.joda.time.DateTime parsable format. Example, &#x27;2022-01-31&#x27;, &#x27;2022-01-31T07:54Z&#x27; or &#x27;2022-01-31T07:54:51.264Z&#x27;.  Defaults to Today - 7days | 
 **endTime** | **optional.**| End time of the cost details. Should use org.joda.time.DateTime parsable format. Example, &#x27;2022-01-31&#x27;, &#x27;2022-01-31T07:54Z&#x27; or &#x27;2022-01-31T07:54:51.264Z&#x27;.  Defaults to Today | 

### Return type

[**ResponseDtoPerspectiveEntityStatsData**](ResponseDTOPerspectiveEntityStatsData.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Costdetailttimeseries**
> ResponseDtoPerspectiveTimeSeriesData Costdetailttimeseries(ctx, body, accountIdentifier, perspectiveId, optional)
Returns cost details in a time series format

Returns cost details in a time series format based on the specified query parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CostDetailsQueryParams**](CostDetailsQueryParams.md)| Cost details query parameters. | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **perspectiveId** | **string**| Perspective identifier of the cost details | 
 **optional** | ***CloudCostDetailsApiCostdetailttimeseriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CloudCostDetailsApiCostdetailttimeseriesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **startTime** | **optional.**| Start time of the cost details. Should use org.joda.time.DateTime parsable format. Example, &#x27;2022-01-31&#x27;, &#x27;2022-01-31T07:54Z&#x27; or &#x27;2022-01-31T07:54:51.264Z&#x27;.  Defaults to Today - 7days | 
 **endTime** | **optional.**| End time of the cost details. Should use org.joda.time.DateTime parsable format. Example, &#x27;2022-01-31&#x27;, &#x27;2022-01-31T07:54Z&#x27; or &#x27;2022-01-31T07:54:51.264Z&#x27;.  Defaults to Today | 

### Return type

[**ResponseDtoPerspectiveTimeSeriesData**](ResponseDTOPerspectiveTimeSeriesData.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

