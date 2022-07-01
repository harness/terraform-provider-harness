# nextgen{{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPipelineExecution**](PipelineDashboardApi.md#GetPipelineExecution) | **Get** /pipeline/api/pipelines/pipelineExecution | Fetch Pipeline Execution Details

# **GetPipelineExecution**
> ResponseDtoDashboardPipelineExecution GetPipelineExecution(ctx, accountIdentifier, orgIdentifier, projectIdentifier, pipelineIdentifier, moduleInfo, startTime, endTime)
Fetch Pipeline Execution Details

Returns Pipeline Execution Details for a Given Interval (Presented in Day Wise Format)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the Entity. | 
  **projectIdentifier** | **string**| Project Identifier for the Entity. | 
  **pipelineIdentifier** | **string**| Pipeline Identifier | 
  **moduleInfo** | **string**| The module from which execution was triggered. | 
  **startTime** | **int64**| Start Date Epoch time in ms | 
  **endTime** | **int64**| End Date Epoch time in ms | 

### Return type

[**ResponseDtoDashboardPipelineExecution**](ResponseDTODashboardPipelineExecution.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

