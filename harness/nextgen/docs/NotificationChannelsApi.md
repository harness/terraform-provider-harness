# nextgen{{classname}}

All URIs are relative to *https://app.harness.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNotificationChannel**](NotificationChannelsApi.md#CreateNotificationChannel) | **Post** /v1/orgs/{org}/projects/{project}/notification-channels | Create Notification channel
[**CreateNotificationChannelAccount**](NotificationChannelsApi.md#CreateNotificationChannelAccount) | **Post** /v1/notification-channels | Create Notification channel
[**CreateNotificationChannelOrg**](NotificationChannelsApi.md#CreateNotificationChannelOrg) | **Post** /v1/orgs/{org}/notification-channels | Create Notification channel
[**DeleteNotificationChannel**](NotificationChannelsApi.md#DeleteNotificationChannel) | **Delete** /v1/orgs/{org}/projects/{project}/notification-channels/{notification-channel} | Delete Notification Channel
[**DeleteNotificationChannelAccount**](NotificationChannelsApi.md#DeleteNotificationChannelAccount) | **Delete** /v1/notification-channels/{notification-channel} | Delete Notification Channel
[**DeleteNotificationChannelOrg**](NotificationChannelsApi.md#DeleteNotificationChannelOrg) | **Delete** /v1/orgs/{org}/notification-channels/{notification-channel} | Delete Notification Channel
[**GetNotificationChannel**](NotificationChannelsApi.md#GetNotificationChannel) | **Get** /v1/orgs/{org}/projects/{project}/notification-channels/{notification-channel} | Get Notification channel
[**GetNotificationChannelAccount**](NotificationChannelsApi.md#GetNotificationChannelAccount) | **Get** /v1/notification-channels/{notification-channel} | Get Notification channel
[**GetNotificationChannelOrg**](NotificationChannelsApi.md#GetNotificationChannelOrg) | **Get** /v1/orgs/{org}/notification-channels/{notification-channel} | Get Notification channel
[**ListNotificationChannels**](NotificationChannelsApi.md#ListNotificationChannels) | **Get** /v1/orgs/{org}/projects/{project}/notification-channels | List Notification channels
[**ListNotificationChannelsAccount**](NotificationChannelsApi.md#ListNotificationChannelsAccount) | **Get** /v1/notification-channels | List Notification channels at account level
[**ListNotificationChannelsOrg**](NotificationChannelsApi.md#ListNotificationChannelsOrg) | **Get** /v1/orgs/{org}/notification-channels | List Notification channels at org level
[**UpdateNotificationChannel**](NotificationChannelsApi.md#UpdateNotificationChannel) | **Put** /v1/orgs/{org}/projects/{project}/notification-channels/{notification-channel} | Update Notification Channel
[**UpdateNotificationChannelAccount**](NotificationChannelsApi.md#UpdateNotificationChannelAccount) | **Put** /v1/notification-channels/{notification-channel} | Update Notification Channel
[**UpdateNotificationChannelOrg**](NotificationChannelsApi.md#UpdateNotificationChannelOrg) | **Put** /v1/orgs/{org}/notification-channels/{notification-channel} | Update Notification Channel
[**ValidateNotificationChannelIdentifier**](NotificationChannelsApi.md#ValidateNotificationChannelIdentifier) | **Get** /v1/orgs/{org}/projects/{project}/validate-channels/{notification-channel} | Validate Notification Channel Identifier
[**ValidateNotificationChannelIdentifierAccount**](NotificationChannelsApi.md#ValidateNotificationChannelIdentifierAccount) | **Get** /v1/validate-channels/{notification-channel} | Validate Notification channel identifier
[**ValidateNotificationChannelIdentifierOrg**](NotificationChannelsApi.md#ValidateNotificationChannelIdentifierOrg) | **Get** /v1/orgs/{org}/validate-channels/{notification-channel} | Validate unique identifier for notification channel

# **CreateNotificationChannel**
> NotificationChannelDto CreateNotificationChannel(ctx, org, project, optional)
Create Notification channel

Create Notification Channel

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **project** | **string**| Identifier field of the project the resource is scoped to | 
 **optional** | ***NotificationChannelsApiCreateNotificationChannelOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationChannelsApiCreateNotificationChannelOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of NotificationChannelDto**](NotificationChannelDto.md)| Notification channel request | 
 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**NotificationChannelDto**](NotificationChannelDTO.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateNotificationChannelAccount**
> NotificationChannelDto CreateNotificationChannelAccount(ctx, optional)
Create Notification channel

Create Notification Channel

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NotificationChannelsApiCreateNotificationChannelAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationChannelsApiCreateNotificationChannelAccountOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of NotificationChannelDto**](NotificationChannelDto.md)| Notification channel request | 
 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**NotificationChannelDto**](NotificationChannelDTO.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateNotificationChannelOrg**
> NotificationChannelDto CreateNotificationChannelOrg(ctx, org, optional)
Create Notification channel

Create Notification Channel

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
 **optional** | ***NotificationChannelsApiCreateNotificationChannelOrgOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationChannelsApiCreateNotificationChannelOrgOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of NotificationChannelDto**](NotificationChannelDto.md)| Notification channel request | 
 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**NotificationChannelDto**](NotificationChannelDTO.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNotificationChannel**
> DeleteNotificationChannel(ctx, notificationChannel, org, project, optional)
Delete Notification Channel

Delete notification channel

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **notificationChannel** | **string**| identifier | 
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **project** | **string**| Identifier field of the project the resource is scoped to | 
 **optional** | ***NotificationChannelsApiDeleteNotificationChannelOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationChannelsApiDeleteNotificationChannelOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNotificationChannelAccount**
> DeleteNotificationChannelAccount(ctx, notificationChannel, optional)
Delete Notification Channel

Delete notification channel

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **notificationChannel** | **string**| identifier | 
 **optional** | ***NotificationChannelsApiDeleteNotificationChannelAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationChannelsApiDeleteNotificationChannelAccountOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNotificationChannelOrg**
> DeleteNotificationChannelOrg(ctx, notificationChannel, org, optional)
Delete Notification Channel

Delete notification channel

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **notificationChannel** | **string**| identifier | 
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
 **optional** | ***NotificationChannelsApiDeleteNotificationChannelOrgOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationChannelsApiDeleteNotificationChannelOrgOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNotificationChannel**
> NotificationChannelDto GetNotificationChannel(ctx, notificationChannel, org, project, optional)
Get Notification channel

Get Notification Channel

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **notificationChannel** | **string**| identifier | 
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **project** | **string**| Identifier field of the project the resource is scoped to | 
 **optional** | ***NotificationChannelsApiGetNotificationChannelOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationChannelsApiGetNotificationChannelOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**NotificationChannelDto**](NotificationChannelDTO.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml, multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNotificationChannelAccount**
> NotificationChannelDto GetNotificationChannelAccount(ctx, notificationChannel, optional)
Get Notification channel

Get Notification Channel

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **notificationChannel** | **string**| identifier | 
 **optional** | ***NotificationChannelsApiGetNotificationChannelAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationChannelsApiGetNotificationChannelAccountOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**NotificationChannelDto**](NotificationChannelDTO.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml, multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNotificationChannelOrg**
> NotificationChannelDto GetNotificationChannelOrg(ctx, notificationChannel, org, optional)
Get Notification channel

Get Notification Channel

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **notificationChannel** | **string**| identifier | 
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
 **optional** | ***NotificationChannelsApiGetNotificationChannelOrgOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationChannelsApiGetNotificationChannelOrgOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**NotificationChannelDto**](NotificationChannelDTO.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml, multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNotificationChannels**
> []NotificationChannelDto ListNotificationChannels(ctx, org, project, optional)
List Notification channels

Returns a list of notification channels for the scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **project** | **string**| Identifier field of the project the resource is scoped to | 
 **optional** | ***NotificationChannelsApiListNotificationChannelsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationChannelsApiListNotificationChannelsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 
 **limit** | **optional.Int32**| Number of items to return per page. | [default to 30]
 **searchTerm** | **optional.String**| This would be used to filter resources having attributes matching with search term. | 
 **sort** | **optional.String**| Parameter on the basis of which sorting is done. | 
 **order** | **optional.String**| Order on the basis of which sorting is done. | 
 **notificationChannelType** | **optional.String**| Notification Channel Type | 
 **page** | **optional.Int32**| Pagination page number strategy: Specify the page number within the paginated collection related to the number of items in each page  | [default to 0]
 **includeParentScope** | **optional.Bool**| Include entities from current and parent scopes. | [default to false]

### Return type

[**[]NotificationChannelDto**](NotificationChannelDTO.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml, multipart/form-data, text/html, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNotificationChannelsAccount**
> []NotificationChannelDto ListNotificationChannelsAccount(ctx, optional)
List Notification channels at account level

Returns a list of notification channels for the scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NotificationChannelsApiListNotificationChannelsAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationChannelsApiListNotificationChannelsAccountOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 
 **limit** | **optional.Int32**| Number of items to return per page. | [default to 30]
 **searchTerm** | **optional.String**| This would be used to filter resources having attributes matching with search term. | 
 **sort** | **optional.String**| Parameter on the basis of which sorting is done. | 
 **order** | **optional.String**| Order on the basis of which sorting is done. | 
 **notificationChannelType** | **optional.String**| Notification Channel Type | 
 **page** | **optional.Int32**| Pagination page number strategy: Specify the page number within the paginated collection related to the number of items in each page  | [default to 0]
 **includeParentScope** | **optional.Bool**| Include entities from current and parent scopes. | [default to false]

### Return type

[**[]NotificationChannelDto**](NotificationChannelDTO.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml, multipart/form-data, text/html, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNotificationChannelsOrg**
> []NotificationChannelDto ListNotificationChannelsOrg(ctx, org, optional)
List Notification channels at org level

Returns a list of notification channels for the scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
 **optional** | ***NotificationChannelsApiListNotificationChannelsOrgOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationChannelsApiListNotificationChannelsOrgOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 
 **limit** | **optional.Int32**| Number of items to return per page. | [default to 30]
 **searchTerm** | **optional.String**| This would be used to filter resources having attributes matching with search term. | 
 **sort** | **optional.String**| Parameter on the basis of which sorting is done. | 
 **order** | **optional.String**| Order on the basis of which sorting is done. | 
 **notificationChannelType** | **optional.String**| Notification Channel Type | 
 **page** | **optional.Int32**| Pagination page number strategy: Specify the page number within the paginated collection related to the number of items in each page  | [default to 0]
 **includeParentScope** | **optional.Bool**| Include entities from current and parent scopes. | [default to false]

### Return type

[**[]NotificationChannelDto**](NotificationChannelDTO.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml, multipart/form-data, text/html, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNotificationChannel**
> NotificationChannelDto UpdateNotificationChannel(ctx, notificationChannel, org, project, optional)
Update Notification Channel

Update Notification Channel

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **notificationChannel** | **string**| identifier | 
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **project** | **string**| Identifier field of the project the resource is scoped to | 
 **optional** | ***NotificationChannelsApiUpdateNotificationChannelOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationChannelsApiUpdateNotificationChannelOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of NotificationChannelDto**](NotificationChannelDto.md)| Notification channel request | 
 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**NotificationChannelDto**](NotificationChannelDTO.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNotificationChannelAccount**
> NotificationChannelDto UpdateNotificationChannelAccount(ctx, notificationChannel, optional)
Update Notification Channel

Update Notification Channel

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **notificationChannel** | **string**| identifier | 
 **optional** | ***NotificationChannelsApiUpdateNotificationChannelAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationChannelsApiUpdateNotificationChannelAccountOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of NotificationChannelDto**](NotificationChannelDto.md)| Notification channel request | 
 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**NotificationChannelDto**](NotificationChannelDTO.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNotificationChannelOrg**
> NotificationChannelDto UpdateNotificationChannelOrg(ctx, notificationChannel, org, optional)
Update Notification Channel

Update Notification Channel

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **notificationChannel** | **string**| identifier | 
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
 **optional** | ***NotificationChannelsApiUpdateNotificationChannelOrgOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationChannelsApiUpdateNotificationChannelOrgOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of NotificationChannelDto**](NotificationChannelDto.md)| Notification channel request | 
 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**NotificationChannelDto**](NotificationChannelDTO.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateNotificationChannelIdentifier**
> ValidateIdentifierDto ValidateNotificationChannelIdentifier(ctx, org, project, notificationChannel, optional)
Validate Notification Channel Identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **project** | **string**| Identifier field of the project the resource is scoped to | 
  **notificationChannel** | **string**| identifier | 
 **optional** | ***NotificationChannelsApiValidateNotificationChannelIdentifierOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationChannelsApiValidateNotificationChannelIdentifierOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**ValidateIdentifierDto**](ValidateIdentifierDTO.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateNotificationChannelIdentifierAccount**
> ValidateIdentifierDto ValidateNotificationChannelIdentifierAccount(ctx, notificationChannel, optional)
Validate Notification channel identifier

Validate Notification Channel Indetifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **notificationChannel** | **string**| identifier | 
 **optional** | ***NotificationChannelsApiValidateNotificationChannelIdentifierAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationChannelsApiValidateNotificationChannelIdentifierAccountOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**ValidateIdentifierDto**](ValidateIdentifierDTO.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateNotificationChannelIdentifierOrg**
> ValidateIdentifierDto ValidateNotificationChannelIdentifierOrg(ctx, org, notificationChannel, optional)
Validate unique identifier for notification channel

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **notificationChannel** | **string**| identifier | 
 **optional** | ***NotificationChannelsApiValidateNotificationChannelIdentifierOrgOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationChannelsApiValidateNotificationChannelIdentifierOrgOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**ValidateIdentifierDto**](ValidateIdentifierDTO.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

