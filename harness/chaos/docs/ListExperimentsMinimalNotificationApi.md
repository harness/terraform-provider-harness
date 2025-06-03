# {{classname}}

All URIs are relative to */api/manager*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListExperimentsWithActiveInfrasMinimalNotification**](ListExperimentsMinimalNotificationApi.md#ListExperimentsWithActiveInfrasMinimalNotification) | **Get** /v1/notification-experiments | List chaos experiments with active infrastructures in minimal format for notification service

# **ListExperimentsWithActiveInfrasMinimalNotification**
> HandlersListExperimentsWithActiveInfrasMinimalNotificationResponse ListExperimentsWithActiveInfrasMinimalNotification(ctx, accountIdentifier, orgIdentifier, projectIdentifier, optional)
List chaos experiments with active infrastructures in minimal format for notification service

List all experiments with active infrastructures in minimal format for notification service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| account id is the account where you want to access the resource | 
  **orgIdentifier** | **string**| organization id is the organization where you want to access the resource | 
  **projectIdentifier** | **string**| project id is the project where you want to access the resource | 
 **optional** | ***ListExperimentsMinimalNotificationApiListExperimentsWithActiveInfrasMinimalNotificationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListExperimentsMinimalNotificationApiListExperimentsWithActiveInfrasMinimalNotificationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **page** | **optional.Int32**| page number | [default to 0]
 **limit** | **optional.Int32**| limit per page | [default to 15]
 **experimentName** | **optional.String**| search filter based on name | 

### Return type

[**HandlersListExperimentsWithActiveInfrasMinimalNotificationResponse**](handlers.ListExperimentsWithActiveInfrasMinimalNotificationResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

