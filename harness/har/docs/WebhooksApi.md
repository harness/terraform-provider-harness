# {{classname}}

All URIs are relative to */gateway/har/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWebhook**](WebhooksApi.md#CreateWebhook) | **Post** /registry/{registry_ref}/+/webhooks | CreateWebhook
[**DeleteWebhook**](WebhooksApi.md#DeleteWebhook) | **Delete** /registry/{registry_ref}/+/webhooks/{webhook_identifier} | DeleteWebhook
[**GetWebhook**](WebhooksApi.md#GetWebhook) | **Get** /registry/{registry_ref}/+/webhooks/{webhook_identifier} | GetWebhook
[**GetWebhookExecution**](WebhooksApi.md#GetWebhookExecution) | **Get** /registry/{registry_ref}/+/webhooks/{webhook_identifier}/executions/{webhook_execution_id} | GetWebhookExecution
[**ListWebhookExecutions**](WebhooksApi.md#ListWebhookExecutions) | **Get** /registry/{registry_ref}/+/webhooks/{webhook_identifier}/executions | ListWebhookExecutions
[**ListWebhooks**](WebhooksApi.md#ListWebhooks) | **Get** /registry/{registry_ref}/+/webhooks | ListWebhooks
[**ReTriggerWebhookExecution**](WebhooksApi.md#ReTriggerWebhookExecution) | **Get** /registry/{registry_ref}/+/webhooks/{webhook_identifier}/executions/{webhook_execution_id}/retrigger | ReTriggerWebhookExecution
[**UpdateWebhook**](WebhooksApi.md#UpdateWebhook) | **Put** /registry/{registry_ref}/+/webhooks/{webhook_identifier} | UpdateWebhook

# **CreateWebhook**
> InlineResponse2011 CreateWebhook(ctx, registryRef, optional)
CreateWebhook

Returns Webhook Details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **registryRef** | **string**| Unique registry path. | 
 **optional** | ***WebhooksApiCreateWebhookOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WebhooksApiCreateWebhookOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of WebhookRequest**](WebhookRequest.md)| request for create and update webhook | 

### Return type

[**InlineResponse2011**](inline_response_201_1.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteWebhook**
> InlineResponse200 DeleteWebhook(ctx, registryRef, webhookIdentifier)
DeleteWebhook

Delete a Webhook

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **registryRef** | **string**| Unique registry path. | 
  **webhookIdentifier** | **string**| Unique webhook identifier. | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWebhook**
> InlineResponse2011 GetWebhook(ctx, registryRef, webhookIdentifier)
GetWebhook

Returns Webhook Details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **registryRef** | **string**| Unique registry path. | 
  **webhookIdentifier** | **string**| Unique webhook identifier. | 

### Return type

[**InlineResponse2011**](inline_response_201_1.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWebhookExecution**
> InlineResponse20021 GetWebhookExecution(ctx, registryRef, webhookIdentifier, webhookExecutionId)
GetWebhookExecution

Returns Webhook Execution Details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **registryRef** | **string**| Unique registry path. | 
  **webhookIdentifier** | **string**| Unique webhook identifier. | 
  **webhookExecutionId** | **string**| Unique webhook execution identifier. | 

### Return type

[**InlineResponse20021**](inline_response_200_21.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListWebhookExecutions**
> InlineResponse20020 ListWebhookExecutions(ctx, registryRef, webhookIdentifier, optional)
ListWebhookExecutions

Returns Webhook Execution Details List

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **registryRef** | **string**| Unique registry path. | 
  **webhookIdentifier** | **string**| Unique webhook identifier. | 
 **optional** | ***WebhooksApiListWebhookExecutionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WebhooksApiListWebhookExecutionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int64**| Current page number | [default to 1]
 **size** | **optional.Int64**| Number of items per page | [default to 20]

### Return type

[**InlineResponse20020**](inline_response_200_20.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListWebhooks**
> InlineResponse20019 ListWebhooks(ctx, registryRef, optional)
ListWebhooks

Returns List of Webhook Details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **registryRef** | **string**| Unique registry path. | 
 **optional** | ***WebhooksApiListWebhooksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WebhooksApiListWebhooksOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int64**| Current page number | [default to 1]
 **size** | **optional.Int64**| Number of items per page | [default to 20]
 **sortOrder** | **optional.String**| sortOrder | 
 **sortField** | **optional.String**| sortField | 
 **searchTerm** | **optional.String**| search Term. | 

### Return type

[**InlineResponse20019**](inline_response_200_19.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReTriggerWebhookExecution**
> InlineResponse20021 ReTriggerWebhookExecution(ctx, registryRef, webhookIdentifier, webhookExecutionId)
ReTriggerWebhookExecution

Retrigger Webhook Execution

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **registryRef** | **string**| Unique registry path. | 
  **webhookIdentifier** | **string**| Unique webhook identifier. | 
  **webhookExecutionId** | **string**| Unique webhook execution identifier. | 

### Return type

[**InlineResponse20021**](inline_response_200_21.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateWebhook**
> InlineResponse2011 UpdateWebhook(ctx, registryRef, webhookIdentifier, optional)
UpdateWebhook

Returns Webhook Details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **registryRef** | **string**| Unique registry path. | 
  **webhookIdentifier** | **string**| Unique webhook identifier. | 
 **optional** | ***WebhooksApiUpdateWebhookOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WebhooksApiUpdateWebhookOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of WebhookRequest**](WebhookRequest.md)| request for create and update webhook | 

### Return type

[**InlineResponse2011**](inline_response_201_1.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

