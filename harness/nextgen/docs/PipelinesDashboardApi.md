# {{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPipelineDashboardExecution**](PipelinesDashboardApi.md#GetPipelineDashboardExecution) | **Get** /pipeline/api/dashboard/pipelineExecution | Fetches Pipeline Executions details for a given Interval and will be presented in day wise format
[**GetPipelineExecution**](PipelinesDashboardApi.md#GetPipelineExecution) | **Get** /pipeline/api/pipelines/pipelineExecution | Fetches Pipeline Executions details for a given Interval and will be presented in day wise format
[**GetPipelinedHealth**](PipelinesDashboardApi.md#GetPipelinedHealth) | **Get** /pipeline/api/pipelines/pipelineHealth | Fetches Pipeline Health data for a given Interval and will be presented in day wise format 
[**GetPipelinedHealth1**](PipelinesDashboardApi.md#GetPipelinedHealth1) | **Get** /pipeline/api/dashboard/pipelineHealth | Fetches Pipeline Health data for a given Interval and will be presented in day wise format

# **GetPipelineDashboardExecution**
> ResponseDtoDashboardPipelineExecution GetPipelineDashboardExecution(ctx, accountIdentifier, orgIdentifier, projectIdentifier, pipelineIdentifier, moduleInfo, startTime, endTime)
Fetches Pipeline Executions details for a given Interval and will be presented in day wise format

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the entity. | 
  **projectIdentifier** | **string**| Project Identifier for the entity. | 
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

# **GetPipelineExecution**
> ResponseDtoDashboardPipelineExecution GetPipelineExecution(ctx, accountIdentifier, orgIdentifier, projectIdentifier, pipelineIdentifier, moduleInfo, startTime, endTime)
Fetches Pipeline Executions details for a given Interval and will be presented in day wise format

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the entity. | 
  **projectIdentifier** | **string**| Project Identifier for the entity. | 
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

# **GetPipelinedHealth**
> ResponseDtoDashboardPipelineHealth GetPipelinedHealth(ctx, accountIdentifier, orgIdentifier, projectIdentifier, pipelineIdentifier, moduleInfo, startTime, endTime)
Fetches Pipeline Health data for a given Interval and will be presented in day wise format 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the entity. | 
  **projectIdentifier** | **string**| Project Identifier for the entity. | 
  **pipelineIdentifier** | **string**| Pipeline Identifier | 
  **moduleInfo** | **string**| The module from which execution was triggered. | 
  **startTime** | **int64**| Start Date Epoch time in ms | 
  **endTime** | **int64**| End Date Epoch time in ms | 

### Return type

[**ResponseDtoDashboardPipelineHealth**](ResponseDTODashboardPipelineHealth.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPipelinedHealth1**
> ResponseDtoDashboardPipelineHealth GetPipelinedHealth1(ctx, accountIdentifier, orgIdentifier, projectIdentifier, pipelineIdentifier, moduleInfo, startTime, endTime)
Fetches Pipeline Health data for a given Interval and will be presented in day wise format

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the entity. | 
  **projectIdentifier** | **string**| Project Identifier for the entity. | 
  **pipelineIdentifier** | **string**| Pipeline Identifier | 
  **moduleInfo** | **string**| The module from which execution was triggered. | 
  **startTime** | **int64**| Start Date Epoch time in ms | 
  **endTime** | **int64**| End Date Epoch time in ms | 

### Return type

[**ResponseDtoDashboardPipelineHealth**](ResponseDTODashboardPipelineHealth.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

