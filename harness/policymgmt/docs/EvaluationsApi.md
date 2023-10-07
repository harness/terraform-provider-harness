# {{classname}}

All URIs are relative to *https://app.harness.io/gateway/pm/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EvaluationsFind**](EvaluationsApi.md#EvaluationsFind) | **Get** /api/v1/evaluations/{id} | 
[**EvaluationsList**](EvaluationsApi.md#EvaluationsList) | **Get** /api/v1/evaluations | 

# **EvaluationsFind**
> Evaluation EvaluationsFind(ctx, id, optional)


Find an evaluation by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| The ID of the evaluation to retrieve | 
 **optional** | ***EvaluationsApiEvaluationsFindOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EvaluationsApiEvaluationsFindOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.String**| Harness account ID | 
 **orgIdentifier** | **optional.String**| Harness organization ID | 
 **projectIdentifier** | **optional.String**| Harness project ID | 
 **xApiKey** | **optional.String**| Harness PAT key used to perform authorization | 

### Return type

[**Evaluation**](Evaluation.md)

### Authorization

[api_key_header_x-api-key](../README.md#api_key_header_x-api-key), [jwt_header_Authorization](../README.md#jwt_header_Authorization)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EvaluationsList**
> []Evaluation EvaluationsList(ctx, optional)


List evaluations

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EvaluationsApiEvaluationsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EvaluationsApiEvaluationsListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountIdentifier** | **optional.String**| Harness account ID | 
 **orgIdentifier** | **optional.String**| Harness organization ID | 
 **projectIdentifier** | **optional.String**| Harness project ID | 
 **perPage** | **optional.Int32**| Number of results per page | [default to 50]
 **page** | **optional.Int32**| Page number (starting from 0) | [default to 0]
 **entity** | **optional.String**| Filter by the entity associated with the evaluation | 
 **type_** | **optional.String**| Filter by the type associated with the evaluation | 
 **action** | **optional.String**| Filter by the action associated with the evaluation | 
 **lastSeen** | **optional.Int64**| Retrieve results starting after this last-seen result | 
 **createdDateFrom** | **optional.Int64**| Retrieve results created from this date | 
 **createdDateTo** | **optional.Int64**| Retrieve results created up to this date | 
 **status** | **optional.String**| Retrieve results with these statuses | 
 **includeChildScopes** | **optional.Bool**| When true, evaluations from child scopes will be inculded in the results | [default to false]
 **xApiKey** | **optional.String**| Harness PAT key used to perform authorization | 

### Return type

[**[]Evaluation**](Evaluation.md)

### Authorization

[api_key_header_x-api-key](../README.md#api_key_header_x-api-key), [jwt_header_Authorization](../README.md#jwt_header_Authorization)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

