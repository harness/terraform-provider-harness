# {{classname}}

All URIs are relative to *https://app.harness.io/gateway/ng/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TriggerFullSync**](FullSyncApi.md#TriggerFullSync) | **Post** /full-sync | Triggers Full Sync

# **TriggerFullSync**
> ResponseDtoTriggerFullSyncResponse TriggerFullSync(ctx, optional)
Triggers Full Sync

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FullSyncApiTriggerFullSyncOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FullSyncApiTriggerFullSyncOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of FullSyncRequest**](FullSyncRequest.md)| The Full Sync Request | 
 **accountIdentifier** | **optional.**| Account Identifier for the entity | 
 **orgIdentifier** | **optional.**| Organization Identifier for the entity | 
 **projectIdentifier** | **optional.**| Project Identifier for the entity | 

### Return type

[**ResponseDtoTriggerFullSyncResponse**](ResponseDTOTriggerFullSyncResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/yaml, text/html
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

