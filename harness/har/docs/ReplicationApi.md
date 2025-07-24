# {{classname}}

All URIs are relative to */gateway/har/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateReplicationRule**](ReplicationApi.md#CreateReplicationRule) | **Post** /replication/rules | Create a replication rule
[**DeleteReplicationRule**](ReplicationApi.md#DeleteReplicationRule) | **Delete** /replication/rules/{id} | Delete a replication rule
[**GetMigrationLogsForImage**](ReplicationApi.md#GetMigrationLogsForImage) | **Get** /replication/rules/{id}/migration/images/{image_id}/logs | Get migration logs for an image
[**GetReplicationRule**](ReplicationApi.md#GetReplicationRule) | **Get** /replication/rules/{id} | Get a replication rule
[**ListMigrationImages**](ReplicationApi.md#ListMigrationImages) | **Get** /replication/rules/{id}/migration/images | List migration images
[**ListReplicationRules**](ReplicationApi.md#ListReplicationRules) | **Get** /replication/rules | List replication rules
[**StartMigration**](ReplicationApi.md#StartMigration) | **Post** /replication/rules/{id}/migration/start | Start migration
[**StopMigration**](ReplicationApi.md#StopMigration) | **Post** /replication/rules/{id}/migration/stop | Stop migration
[**UpdateReplicationRule**](ReplicationApi.md#UpdateReplicationRule) | **Put** /replication/rules/{id} | Update a replication rule

# **CreateReplicationRule**
> InlineResponse20023 CreateReplicationRule(ctx, optional)
Create a replication rule

Create a replication rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReplicationApiCreateReplicationRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplicationApiCreateReplicationRuleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ReplicationRuleRequest**](ReplicationRuleRequest.md)| request for create and update replication rule | 
 **spaceRef** | **optional.**| Unique space path | 

### Return type

[**InlineResponse20023**](inline_response_200_23.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteReplicationRule**
> InlineResponse200 DeleteReplicationRule(ctx, id)
Delete a replication rule

Delete a replication rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMigrationLogsForImage**
> string GetMigrationLogsForImage(ctx, id, imageId)
Get migration logs for an image

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **imageId** | **string**|  | 

### Return type

**string**

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain; charset=utf-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReplicationRule**
> InlineResponse20023 GetReplicationRule(ctx, id)
Get a replication rule

Get a replication rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**InlineResponse20023**](inline_response_200_23.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMigrationImages**
> InlineResponse20024 ListMigrationImages(ctx, id, optional)
List migration images

List migration images given an id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***ReplicationApiListMigrationImagesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplicationApiListMigrationImagesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int64**| Current page number | [default to 1]
 **size** | **optional.Int64**| Number of items per page | [default to 20]
 **sortOrder** | **optional.String**| sortOrder | 
 **sortField** | **optional.String**| sortField | 

### Return type

[**InlineResponse20024**](inline_response_200_24.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListReplicationRules**
> InlineResponse20022 ListReplicationRules(ctx, optional)
List replication rules

List all replication rules

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReplicationApiListReplicationRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplicationApiListReplicationRulesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **spaceRef** | **optional.String**| Unique space path | 

### Return type

[**InlineResponse20022**](inline_response_200_22.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StartMigration**
> InlineResponse200 StartMigration(ctx, id)
Start migration

Start migration given an id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StopMigration**
> InlineResponse200 StopMigration(ctx, id)
Stop migration

Stop migration given an id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateReplicationRule**
> InlineResponse20023 UpdateReplicationRule(ctx, id, optional)
Update a replication rule

Update a replication rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***ReplicationApiUpdateReplicationRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplicationApiUpdateReplicationRuleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ReplicationRuleRequest**](ReplicationRuleRequest.md)| request for create and update replication rule | 

### Return type

[**InlineResponse20023**](inline_response_200_23.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

