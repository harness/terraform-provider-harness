# {{classname}}

All URIs are relative to */gateway/code/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWebhook**](WebhookApi.md#CreateWebhook) | **Post** /repos/{repo_identifier}/webhooks | Create webhook
[**DeleteWebhook**](WebhookApi.md#DeleteWebhook) | **Delete** /repos/{repo_identifier}/webhooks/{webhook_identifier} | Delete webhook
[**GetWebhook**](WebhookApi.md#GetWebhook) | **Get** /repos/{repo_identifier}/webhooks/{webhook_identifier} | Get webhook
[**GetWebhookExecution**](WebhookApi.md#GetWebhookExecution) | **Get** /repos/{repo_identifier}/webhooks/{webhook_identifier}/executions/{webhook_execution_id} | Get webhook execution
[**ListWebhookExecutions**](WebhookApi.md#ListWebhookExecutions) | **Get** /repos/{repo_identifier}/webhooks/{webhook_identifier}/executions | List webhook executions
[**ListWebhooks**](WebhookApi.md#ListWebhooks) | **Get** /repos/{repo_identifier}/webhooks | List webhooks
[**UpdateWebhook**](WebhookApi.md#UpdateWebhook) | **Patch** /repos/{repo_identifier}/webhooks/{webhook_identifier} | Update webhook

# **CreateWebhook**
> OpenapiWebhookType CreateWebhook(ctx, accountIdentifier, repoIdentifier, optional)
Create webhook

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
  **repoIdentifier** | **string**|  | 
 **optional** | ***WebhookApiCreateWebhookOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WebhookApiCreateWebhookOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of OpenapiCreateWebhookRequest**](OpenapiCreateWebhookRequest.md)|  | 
 **orgIdentifier** | **optional.**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity.. | 

### Return type

[**OpenapiWebhookType**](OpenapiWebhookType.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteWebhook**
> DeleteWebhook(ctx, accountIdentifier, repoIdentifier, webhookIdentifier, optional)
Delete webhook

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
  **repoIdentifier** | **string**|  | 
  **webhookIdentifier** | **string**|  | 
 **optional** | ***WebhookApiDeleteWebhookOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WebhookApiDeleteWebhookOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity.. | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth), [x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWebhook**
> OpenapiWebhookType GetWebhook(ctx, accountIdentifier, repoIdentifier, webhookIdentifier, optional)
Get webhook

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
  **repoIdentifier** | **string**|  | 
  **webhookIdentifier** | **string**|  | 
 **optional** | ***WebhookApiGetWebhookOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WebhookApiGetWebhookOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity.. | 

### Return type

[**OpenapiWebhookType**](OpenapiWebhookType.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWebhookExecution**
> TypesWebhookExecution GetWebhookExecution(ctx, accountIdentifier, repoIdentifier, webhookIdentifier, webhookExecutionId, optional)
Get webhook execution

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
  **repoIdentifier** | **string**|  | 
  **webhookIdentifier** | **string**|  | 
  **webhookExecutionId** | **int32**|  | 
 **optional** | ***WebhookApiGetWebhookExecutionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WebhookApiGetWebhookExecutionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity.. | 
 **page** | **optional.Int32**| The page to return. | [default to 1]
 **limit** | **optional.Int32**| The maximum number of results to return. | [default to 30]

### Return type

[**TypesWebhookExecution**](TypesWebhookExecution.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListWebhookExecutions**
> []TypesWebhookExecution ListWebhookExecutions(ctx, accountIdentifier, repoIdentifier, webhookIdentifier, optional)
List webhook executions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
  **repoIdentifier** | **string**|  | 
  **webhookIdentifier** | **string**|  | 
 **optional** | ***WebhookApiListWebhookExecutionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WebhookApiListWebhookExecutionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity.. | 
 **page** | **optional.Int32**| The page to return. | [default to 1]
 **limit** | **optional.Int32**| The maximum number of results to return. | [default to 30]

### Return type

[**[]TypesWebhookExecution**](TypesWebhookExecution.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListWebhooks**
> []OpenapiWebhookType ListWebhooks(ctx, accountIdentifier, repoIdentifier, optional)
List webhooks

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
  **repoIdentifier** | **string**|  | 
 **optional** | ***WebhookApiListWebhooksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WebhookApiListWebhooksOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity.. | 
 **query** | **optional.String**| The substring which is used to filter the webhooks by their identifier. | 
 **sort** | **optional.String**| The data by which the webhooks are sorted. | [default to identifier]
 **order** | **optional.String**| The order of the output. | [default to asc]
 **page** | **optional.Int32**| The page to return. | [default to 1]
 **limit** | **optional.Int32**| The maximum number of results to return. | [default to 30]

### Return type

[**[]OpenapiWebhookType**](OpenapiWebhookType.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateWebhook**
> OpenapiWebhookType UpdateWebhook(ctx, accountIdentifier, repoIdentifier, webhookIdentifier, optional)
Update webhook

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
  **repoIdentifier** | **string**|  | 
  **webhookIdentifier** | **string**|  | 
 **optional** | ***WebhookApiUpdateWebhookOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WebhookApiUpdateWebhookOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of OpenapiUpdateWebhookRequest**](OpenapiUpdateWebhookRequest.md)|  | 
 **orgIdentifier** | **optional.**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity.. | 

### Return type

[**OpenapiWebhookType**](OpenapiWebhookType.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

