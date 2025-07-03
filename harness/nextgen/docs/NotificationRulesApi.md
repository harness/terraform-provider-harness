# nextgen{{classname}}

All URIs are relative to *https://app.harness.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNotificationRule**](NotificationRulesApi.md#CreateNotificationRule) | **Post** /v1/orgs/{org}/projects/{project}/notification-rules | Create Notification Rule
[**CreateNotificationRuleAccount**](NotificationRulesApi.md#CreateNotificationRuleAccount) | **Post** /v1/notification-rules | Create Notification Rule
[**CreateNotificationRuleOrg**](NotificationRulesApi.md#CreateNotificationRuleOrg) | **Post** /v1/orgs/{org}/notification-rules | Create Notification Rule
[**DeleteNotificationRule**](NotificationRulesApi.md#DeleteNotificationRule) | **Delete** /v1/orgs/{org}/projects/{project}/notification-rules/{notification-rule} | Delete Notification Rule
[**DeleteNotificationRuleAccount**](NotificationRulesApi.md#DeleteNotificationRuleAccount) | **Delete** /v1/notification-rules/{notification-rule} | Delete Notification Rule
[**DeleteNotificationRuleOrg**](NotificationRulesApi.md#DeleteNotificationRuleOrg) | **Delete** /v1/orgs/{org}/notification-rules/{notification-rule} | Delete Notification Rule
[**GetNotificationRule**](NotificationRulesApi.md#GetNotificationRule) | **Get** /v1/orgs/{org}/projects/{project}/notification-rules/{notification-rule} | Get Notification Rule
[**GetNotificationRuleAccount**](NotificationRulesApi.md#GetNotificationRuleAccount) | **Get** /v1/notification-rules/{notification-rule} | Get Notification Rule
[**GetNotificationRuleOrg**](NotificationRulesApi.md#GetNotificationRuleOrg) | **Get** /v1/orgs/{org}/notification-rules/{notification-rule} | Get Notification Rule
[**ListNotificationRules**](NotificationRulesApi.md#ListNotificationRules) | **Get** /v1/orgs/{org}/projects/{project}/notification-rules | List Notification rules
[**ListNotificationRulesAccount**](NotificationRulesApi.md#ListNotificationRulesAccount) | **Get** /v1/notification-rules | List Notification rules at account level
[**ListNotificationRulesOrg**](NotificationRulesApi.md#ListNotificationRulesOrg) | **Get** /v1/orgs/{org}/notification-rules | List Notification rules at org level
[**NotificationResourceList**](NotificationRulesApi.md#NotificationResourceList) | **Get** /v1/notification-resource-list | List of notification entities and events
[**SimulateNotifications**](NotificationRulesApi.md#SimulateNotifications) | **Post** /v1/notification-rules/notifications-simulate | Simulate Notifications
[**UpdateNotificationRule**](NotificationRulesApi.md#UpdateNotificationRule) | **Put** /v1/orgs/{org}/projects/{project}/notification-rules/{notification-rule} | Update Notification Rule
[**UpdateNotificationRuleAccount**](NotificationRulesApi.md#UpdateNotificationRuleAccount) | **Put** /v1/notification-rules/{notification-rule} | Update Notification Rule
[**UpdateNotificationRuleOrg**](NotificationRulesApi.md#UpdateNotificationRuleOrg) | **Put** /v1/orgs/{org}/notification-rules/{notification-rule} | Update Notification Rule
[**ValidateNotificationRuleIdentifier**](NotificationRulesApi.md#ValidateNotificationRuleIdentifier) | **Get** /v1/orgs/{org}/projects/{project}/notification-rules/validate-rules/{notification-rule} | Validate notification rule identifier
[**ValidateNotificationRuleIdentifierAccount**](NotificationRulesApi.md#ValidateNotificationRuleIdentifierAccount) | **Get** /v1/validate-rules/{notification-rule} | Validate notification rule identifier
[**ValidateNotificationRuleIdentifierOrg**](NotificationRulesApi.md#ValidateNotificationRuleIdentifierOrg) | **Get** /v1/orgs/{org}/validate-rules/{notification-rule} | Validate notification rule identifier

# **CreateNotificationRule**
> NotificationRuleDto CreateNotificationRule(ctx, org, project, optional)
Create Notification Rule

Create Notification Rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **project** | **string**| Identifier field of the project the resource is scoped to | 
 **optional** | ***NotificationRulesApiCreateNotificationRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationRulesApiCreateNotificationRuleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of NotificationRuleDto**](NotificationRuleDto.md)| Notification rule request | 
 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**NotificationRuleDto**](NotificationRuleDTO.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateNotificationRuleAccount**
> NotificationRuleDto CreateNotificationRuleAccount(ctx, optional)
Create Notification Rule

Create Notification Rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NotificationRulesApiCreateNotificationRuleAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationRulesApiCreateNotificationRuleAccountOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of NotificationRuleDto**](NotificationRuleDto.md)| Notification rule request | 
 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**NotificationRuleDto**](NotificationRuleDTO.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateNotificationRuleOrg**
> NotificationRuleDto CreateNotificationRuleOrg(ctx, org, optional)
Create Notification Rule

Create Notification Rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
 **optional** | ***NotificationRulesApiCreateNotificationRuleOrgOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationRulesApiCreateNotificationRuleOrgOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of NotificationRuleDto**](NotificationRuleDto.md)| Notification rule request | 
 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**NotificationRuleDto**](NotificationRuleDTO.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNotificationRule**
> DeleteNotificationRule(ctx, org, project, notificationRule, optional)
Delete Notification Rule

Delete notification rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **project** | **string**| Identifier field of the project the resource is scoped to | 
  **notificationRule** | **string**| identifier | 
 **optional** | ***NotificationRulesApiDeleteNotificationRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationRulesApiDeleteNotificationRuleOpts struct
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

# **DeleteNotificationRuleAccount**
> DeleteNotificationRuleAccount(ctx, notificationRule, optional)
Delete Notification Rule

Delete notification rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **notificationRule** | **string**| identifier | 
 **optional** | ***NotificationRulesApiDeleteNotificationRuleAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationRulesApiDeleteNotificationRuleAccountOpts struct
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

# **DeleteNotificationRuleOrg**
> DeleteNotificationRuleOrg(ctx, org, notificationRule, optional)
Delete Notification Rule

Delete notification rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **notificationRule** | **string**| identifier | 
 **optional** | ***NotificationRulesApiDeleteNotificationRuleOrgOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationRulesApiDeleteNotificationRuleOrgOpts struct
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

# **GetNotificationRule**
> NotificationRuleDto GetNotificationRule(ctx, org, project, notificationRule, optional)
Get Notification Rule

Get notification rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **project** | **string**| Identifier field of the project the resource is scoped to | 
  **notificationRule** | **string**| identifier | 
 **optional** | ***NotificationRulesApiGetNotificationRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationRulesApiGetNotificationRuleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**NotificationRuleDto**](NotificationRuleDTO.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml, multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNotificationRuleAccount**
> NotificationRuleDto GetNotificationRuleAccount(ctx, notificationRule, optional)
Get Notification Rule

Get notification rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **notificationRule** | **string**| identifier | 
 **optional** | ***NotificationRulesApiGetNotificationRuleAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationRulesApiGetNotificationRuleAccountOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**NotificationRuleDto**](NotificationRuleDTO.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml, multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNotificationRuleOrg**
> NotificationRuleDto GetNotificationRuleOrg(ctx, org, notificationRule, optional)
Get Notification Rule

Get notification rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **notificationRule** | **string**| identifier | 
 **optional** | ***NotificationRulesApiGetNotificationRuleOrgOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationRulesApiGetNotificationRuleOrgOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**NotificationRuleDto**](NotificationRuleDTO.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml, multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNotificationRules**
> []NotificationRuleDto ListNotificationRules(ctx, org, project, optional)
List Notification rules

Get list of notification rules

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **project** | **string**| Identifier field of the project the resource is scoped to | 
 **optional** | ***NotificationRulesApiListNotificationRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationRulesApiListNotificationRulesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 
 **limit** | **optional.Int32**| Number of items to return per page. | [default to 30]
 **sort** | **optional.String**| Parameter on the basis of which sorting is done. | 
 **order** | **optional.String**| Order on the basis of which sorting is done. | 
 **searchTerm** | **optional.String**| This would be used to filter resources having attributes matching with search term. | 
 **page** | **optional.Int32**| Pagination page number strategy: Specify the page number within the paginated collection related to the number of items in each page  | [default to 0]
 **resource** | **optional.String**| Notification entity name | 
 **event** | **optional.String**| Notification event name | 

### Return type

[**[]NotificationRuleDto**](NotificationRuleDTO.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNotificationRulesAccount**
> []NotificationRuleDto ListNotificationRulesAccount(ctx, optional)
List Notification rules at account level

Get list of notification rules for account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NotificationRulesApiListNotificationRulesAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationRulesApiListNotificationRulesAccountOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 
 **page** | **optional.Int32**| Pagination page number strategy: Specify the page number within the paginated collection related to the number of items in each page  | [default to 0]
 **limit** | **optional.Int32**| Number of items to return per page. | [default to 30]
 **sort** | **optional.String**| Parameter on the basis of which sorting is done. | 
 **order** | **optional.String**| Order on the basis of which sorting is done. | 
 **searchTerm** | **optional.String**| This would be used to filter resources having attributes matching with search term. | 
 **resource** | **optional.String**| Notification resource name | 
 **event** | **optional.String**| Notification event name | 

### Return type

[**[]NotificationRuleDto**](NotificationRuleDTO.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNotificationRulesOrg**
> []NotificationRuleDto ListNotificationRulesOrg(ctx, org, optional)
List Notification rules at org level

Get list of notification rules

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
 **optional** | ***NotificationRulesApiListNotificationRulesOrgOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationRulesApiListNotificationRulesOrgOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 
 **page** | **optional.Int32**| Pagination page number strategy: Specify the page number within the paginated collection related to the number of items in each page  | [default to 0]
 **limit** | **optional.Int32**| Number of items to return per page. | [default to 30]
 **sort** | **optional.String**| Parameter on the basis of which sorting is done. | 
 **order** | **optional.String**| Order on the basis of which sorting is done. | 
 **searchTerm** | **optional.String**| This would be used to filter resources having attributes matching with search term. | 
 **resource** | **optional.String**| Notification entity name | 
 **event** | **optional.String**| Notification event name | 

### Return type

[**[]NotificationRuleDto**](NotificationRuleDTO.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NotificationResourceList**
> []NotificationResourceDto NotificationResourceList(ctx, optional)
List of notification entities and events

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NotificationRulesApiNotificationResourceListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationRulesApiNotificationResourceListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**[]NotificationResourceDto**](NotificationResourceDTO.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SimulateNotifications**
> SimulateNotifications(ctx, optional)
Simulate Notifications

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NotificationRulesApiSimulateNotificationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationRulesApiSimulateNotificationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of NotificationsSimulateDto**](NotificationsSimulateDto.md)|  | 

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNotificationRule**
> NotificationRuleDto UpdateNotificationRule(ctx, org, project, notificationRule, optional)
Update Notification Rule

Update Notification Rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **project** | **string**| Identifier field of the project the resource is scoped to | 
  **notificationRule** | **string**| identifier | 
 **optional** | ***NotificationRulesApiUpdateNotificationRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationRulesApiUpdateNotificationRuleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of NotificationRuleDto**](NotificationRuleDto.md)| Notification rule request | 
 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**NotificationRuleDto**](NotificationRuleDTO.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNotificationRuleAccount**
> NotificationRuleDto UpdateNotificationRuleAccount(ctx, notificationRule, optional)
Update Notification Rule

Update Notification Rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **notificationRule** | **string**| identifier | 
 **optional** | ***NotificationRulesApiUpdateNotificationRuleAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationRulesApiUpdateNotificationRuleAccountOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of NotificationRuleDto**](NotificationRuleDto.md)| Notification rule request | 
 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**NotificationRuleDto**](NotificationRuleDTO.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNotificationRuleOrg**
> NotificationRuleDto UpdateNotificationRuleOrg(ctx, org, notificationRule, optional)
Update Notification Rule

Update Notification Rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **notificationRule** | **string**| identifier | 
 **optional** | ***NotificationRulesApiUpdateNotificationRuleOrgOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationRulesApiUpdateNotificationRuleOrgOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of NotificationRuleDto**](NotificationRuleDto.md)| Notification rule request | 
 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**NotificationRuleDto**](NotificationRuleDTO.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateNotificationRuleIdentifier**
> ValidateIdentifierDto ValidateNotificationRuleIdentifier(ctx, org, project, notificationRule, optional)
Validate notification rule identifier

Validate notification rule identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **project** | **string**| Identifier field of the project the resource is scoped to | 
  **notificationRule** | **string**| Notification Rule Identifier | 
 **optional** | ***NotificationRulesApiValidateNotificationRuleIdentifierOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationRulesApiValidateNotificationRuleIdentifierOpts struct
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

# **ValidateNotificationRuleIdentifierAccount**
> ValidateIdentifierDto ValidateNotificationRuleIdentifierAccount(ctx, notificationRule, optional)
Validate notification rule identifier

Validate notification rule identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **notificationRule** | **string**| identifier | 
 **optional** | ***NotificationRulesApiValidateNotificationRuleIdentifierAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationRulesApiValidateNotificationRuleIdentifierAccountOpts struct
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

# **ValidateNotificationRuleIdentifierOrg**
> ValidateIdentifierDto ValidateNotificationRuleIdentifierOrg(ctx, org, notificationRule, optional)
Validate notification rule identifier

Validate notification rule identifier org level

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **notificationRule** | **string**| Notification Rule Identifier | 
 **optional** | ***NotificationRulesApiValidateNotificationRuleIdentifierOrgOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationRulesApiValidateNotificationRuleIdentifierOrgOpts struct
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

