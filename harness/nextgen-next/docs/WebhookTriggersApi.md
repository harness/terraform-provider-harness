# {{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchWebhookDetails**](WebhookTriggersApi.md#FetchWebhookDetails) | **Get** /pipeline/api/webhook/triggerProcessingDetails | Gets webhook event processing details for input eventId.
[**GetActionsList**](WebhookTriggersApi.md#GetActionsList) | **Get** /pipeline/api/webhook/actions | Get all supported actions for event type and source.
[**GetBitbucketPRActions**](WebhookTriggersApi.md#GetBitbucketPRActions) | **Get** /pipeline/api/webhook/bitbucketPRActions | Gets all supported Bitbucket PR event actions.
[**GetBitbucketTriggerEvents**](WebhookTriggersApi.md#GetBitbucketTriggerEvents) | **Get** /pipeline/api/webhook/bitbucketTriggerEvents | Gets all supported Bitbucket trigger events.
[**GetGitTriggerEventDetails**](WebhookTriggersApi.md#GetGitTriggerEventDetails) | **Get** /pipeline/api/webhook/gitTriggerEventDetails | Gets trigger git actions for each supported event.
[**GetGithubIssueCommentActions**](WebhookTriggersApi.md#GetGithubIssueCommentActions) | **Get** /pipeline/api/webhook/githubIssueCommentActions | Gets all supported Github Issue comment event actions
[**GetGithubPRActions**](WebhookTriggersApi.md#GetGithubPRActions) | **Get** /pipeline/api/webhook/githubPRActions | Gets all supported Github PR event actions
[**GetGithubTriggerEvents**](WebhookTriggersApi.md#GetGithubTriggerEvents) | **Get** /pipeline/api/webhook/githubTriggerEvents | Gets all supported Github trigger events.
[**GetGitlabPRActions**](WebhookTriggersApi.md#GetGitlabPRActions) | **Get** /pipeline/api/webhook/gitlabPRActions | Gets all supported GitLab PR event actions.
[**GetGitlabTriggerEvents**](WebhookTriggersApi.md#GetGitlabTriggerEvents) | **Get** /pipeline/api/webhook/gitlabTriggerEvents | Gets all supported Gitlab trigger events.
[**GetSourceRepos**](WebhookTriggersApi.md#GetSourceRepos) | **Get** /pipeline/api/webhook/sourceRepos | Gets source repo types with all supported events.
[**GetWebhookTriggerTypes**](WebhookTriggersApi.md#GetWebhookTriggerTypes) | **Get** /pipeline/api/webhook/webhookTriggerTypes | Gets all supported scm webhook type.
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

# **GetActionsList**
> ResponseDtoListWebhookAction GetActionsList(ctx, sourceRepo, event)
Get all supported actions for event type and source.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sourceRepo** | **string**|  | 
  **event** | **string**|  | 

### Return type

[**ResponseDtoListWebhookAction**](ResponseDTOListWebhookAction.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBitbucketPRActions**
> ResponseDtoListBitbucketPrAction GetBitbucketPRActions(ctx, )
Gets all supported Bitbucket PR event actions.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ResponseDtoListBitbucketPrAction**](ResponseDTOListBitbucketPRAction.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBitbucketTriggerEvents**
> ResponseDtoListBitbucketTriggerEvent GetBitbucketTriggerEvents(ctx, )
Gets all supported Bitbucket trigger events.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ResponseDtoListBitbucketTriggerEvent**](ResponseDTOListBitbucketTriggerEvent.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGitTriggerEventDetails**
> ResponseDtoMapStringMapStringListString GetGitTriggerEventDetails(ctx, )
Gets trigger git actions for each supported event.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ResponseDtoMapStringMapStringListString**](ResponseDTOMapStringMapStringListString.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGithubIssueCommentActions**
> ResponseDtoListGithubIssueCommentAction GetGithubIssueCommentActions(ctx, )
Gets all supported Github Issue comment event actions

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ResponseDtoListGithubIssueCommentAction**](ResponseDTOListGithubIssueCommentAction.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGithubPRActions**
> ResponseDtoListGithubPrAction GetGithubPRActions(ctx, )
Gets all supported Github PR event actions

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ResponseDtoListGithubPrAction**](ResponseDTOListGithubPRAction.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGithubTriggerEvents**
> ResponseDtoListGithubTriggerEvent GetGithubTriggerEvents(ctx, )
Gets all supported Github trigger events.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ResponseDtoListGithubTriggerEvent**](ResponseDTOListGithubTriggerEvent.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGitlabPRActions**
> ResponseDtoListGitlabPrAction GetGitlabPRActions(ctx, )
Gets all supported GitLab PR event actions.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ResponseDtoListGitlabPrAction**](ResponseDTOListGitlabPRAction.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGitlabTriggerEvents**
> ResponseDtoListGitlabTriggerEvent GetGitlabTriggerEvents(ctx, )
Gets all supported Gitlab trigger events.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ResponseDtoListGitlabTriggerEvent**](ResponseDTOListGitlabTriggerEvent.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSourceRepos**
> ResponseDtoMapWebhookSourceRepoListWebhookEvent GetSourceRepos(ctx, )
Gets source repo types with all supported events.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ResponseDtoMapWebhookSourceRepoListWebhookEvent**](ResponseDTOMapWebhookSourceRepoListWebhookEvent.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWebhookTriggerTypes**
> ResponseDtoListWebhookTriggerType GetWebhookTriggerTypes(ctx, )
Gets all supported scm webhook type.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ResponseDtoListWebhookTriggerType**](ResponseDTOListWebhookTriggerType.md)

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

