# dbops{{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConsumeV1StepPluginResponse**](PipelineStepApi.md#ConsumeV1StepPluginResponse) | **Post** /v1/step/plugin-response | ConsumePluginResponse
[**CreateV1StepPluginInputs**](PipelineStepApi.md#CreateV1StepPluginInputs) | **Post** /v1/step/plugin-inputs | CreatePluginInputs

# **ConsumeV1StepPluginResponse**
> ConsumeV1StepPluginResponse(ctx, optional)
ConsumePluginResponse

Consume Plugin Execution Response

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PipelineStepApiConsumeV1StepPluginResponseOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelineStepApiConsumeV1StepPluginResponseOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of interface{}**](interface{}.md)|  | 
 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateV1StepPluginInputs**
> InlineResponse200 CreateV1StepPluginInputs(ctx, body, optional)
CreatePluginInputs

Create Plugin Inputs using Step Inputs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**interface{}**](interface{}.md)| Inputs to create plugin inputs | 
 **optional** | ***PipelineStepApiCreateV1StepPluginInputsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelineStepApiCreateV1StepPluginInputsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

