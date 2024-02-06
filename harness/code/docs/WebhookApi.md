# code{{classname}}

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWebhook**](WebhookApi.md#CreateWebhook) | **Post** /repos/{repo_identifier}/webhooks | 
[**DeleteWebhook**](WebhookApi.md#DeleteWebhook) | **Delete** /repos/{repo_identifier}/webhooks/{webhook_identifier} | 
[**GetWebhook**](WebhookApi.md#GetWebhook) | **Get** /repos/{repo_identifier}/webhooks/{webhook_identifier} | 
[**GetWebhookExecution**](WebhookApi.md#GetWebhookExecution) | **Get** /repos/{repo_identifier}/webhooks/{webhook_identifier}/executions/{webhook_execution_id} | 
[**ListWebhookExecutions**](WebhookApi.md#ListWebhookExecutions) | **Get** /repos/{repo_identifier}/webhooks/{webhook_identifier}/executions | 
[**ListWebhooks**](WebhookApi.md#ListWebhooks) | **Get** /repos/{repo_identifier}/webhooks | 
[**UpdateWebhook**](WebhookApi.md#UpdateWebhook) | **Patch** /repos/{repo_identifier}/webhooks/{webhook_identifier} | 

# **CreateWebhook**
> OpenapiWebhookType CreateWebhook(ctx, repoIdentifier, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoIdentifier** | **string**|  | 
 **optional** | ***WebhookApiCreateWebhookOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WebhookApiCreateWebhookOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of OpenapiCreateWebhookRequest**](OpenapiCreateWebhookRequest.md)|  | 

### Return type

[**OpenapiWebhookType**](OpenapiWebhookType.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteWebhook**
> DeleteWebhook(ctx, repoIdentifier, webhookIdentifier)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoIdentifier** | **string**|  | 
  **webhookIdentifier** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWebhook**
> OpenapiWebhookType GetWebhook(ctx, repoIdentifier, webhookIdentifier)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoIdentifier** | **string**|  | 
  **webhookIdentifier** | **int32**|  | 

### Return type

[**OpenapiWebhookType**](OpenapiWebhookType.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWebhookExecution**
> TypesWebhookExecution GetWebhookExecution(ctx, repoIdentifier, webhookIdentifier, webhookExecutionId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoIdentifier** | **string**|  | 
  **webhookIdentifier** | **int32**|  | 
  **webhookExecutionId** | **int32**|  | 
 **optional** | ***WebhookApiGetWebhookExecutionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WebhookApiGetWebhookExecutionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **page** | **optional.Int32**| The page to return. | [default to 1]
 **limit** | **optional.Int32**| The maximum number of results to return. | [default to 30]

### Return type

[**TypesWebhookExecution**](TypesWebhookExecution.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListWebhookExecutions**
> []TypesWebhookExecution ListWebhookExecutions(ctx, repoIdentifier, webhookIdentifier, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoIdentifier** | **string**|  | 
  **webhookIdentifier** | **int32**|  | 
 **optional** | ***WebhookApiListWebhookExecutionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WebhookApiListWebhookExecutionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| The page to return. | [default to 1]
 **limit** | **optional.Int32**| The maximum number of results to return. | [default to 30]

### Return type

[**[]TypesWebhookExecution**](TypesWebhookExecution.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListWebhooks**
> []OpenapiWebhookType ListWebhooks(ctx, repoIdentifier, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoIdentifier** | **string**|  | 
 **optional** | ***WebhookApiListWebhooksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WebhookApiListWebhooksOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **query** | **optional.String**| The substring which is used to filter the webhooks by their identifier. | 
 **sort** | **optional.String**| The data by which the webhooks are sorted. | [default to identifier]
 **order** | **optional.String**| The order of the output. | [default to asc]
 **page** | **optional.Int32**| The page to return. | [default to 1]
 **limit** | **optional.Int32**| The maximum number of results to return. | [default to 30]

### Return type

[**[]OpenapiWebhookType**](OpenapiWebhookType.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateWebhook**
> OpenapiWebhookType UpdateWebhook(ctx, repoIdentifier, webhookIdentifier, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoIdentifier** | **string**|  | 
  **webhookIdentifier** | **int32**|  | 
 **optional** | ***WebhookApiUpdateWebhookOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WebhookApiUpdateWebhookOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of OpenapiUpdateWebhookRequest**](OpenapiUpdateWebhookRequest.md)|  | 

### Return type

[**OpenapiWebhookType**](OpenapiWebhookType.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

