# nextgen{{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchWebhookDetails**](WebhookTriggersApi.md#FetchWebhookDetails) | **Get** /pipeline/api/webhook/triggerProcessingDetails | Gets webhook event processing details for input eventId.
[**PipelineprocessWebhookEvent**](WebhookTriggersApi.md#PipelineprocessWebhookEvent) | **Post** /pipeline/api/webhook/trigger | Handles event payload for webhook triggers.
[**ProcessCustomWebhookEvent**](WebhookTriggersApi.md#ProcessCustomWebhookEvent) | **Post** /pipeline/api/webhook/custom | Handles event payload for custom webhook triggers.

# **FetchWebhookDetails**
> ResponseDtoWebhookEventProcessingDetails FetchWebhookDetails(ctx, accountIdentifier, eventId)
Gets webhook event processing details for input eventId.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**|  | 
  **eventId** | **string**|  | 

### Return type

[**ResponseDtoWebhookEventProcessingDetails**](ResponseDTOWebhookEventProcessingDetails.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PipelineprocessWebhookEvent**
> ResponseDtoString PipelineprocessWebhookEvent(ctx, body, accountIdentifier, optional)
Handles event payload for webhook triggers.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)|  | 
  **accountIdentifier** | **string**|  | 
 **optional** | ***WebhookTriggersApiPipelineprocessWebhookEventOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WebhookTriggersApiPipelineprocessWebhookEventOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.**|  | 
 **projectIdentifier** | **optional.**|  | 

### Return type

[**ResponseDtoString**](ResponseDTOString.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml, text/plain
 - **Accept**: application/json, application/yaml, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProcessCustomWebhookEvent**
> ResponseDtoString ProcessCustomWebhookEvent(ctx, body, accountIdentifier, orgIdentifier, projectIdentifier, optional)
Handles event payload for custom webhook triggers.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)|  | 
  **accountIdentifier** | **string**|  | 
  **orgIdentifier** | **string**|  | 
  **projectIdentifier** | **string**|  | 
 **optional** | ***WebhookTriggersApiProcessCustomWebhookEventOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WebhookTriggersApiProcessCustomWebhookEventOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **pipelineIdentifier** | **optional.**|  | 
 **triggerIdentifier** | **optional.**|  | 

### Return type

[**ResponseDtoString**](ResponseDTOString.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml, text/plain
 - **Accept**: application/json, application/yaml, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

