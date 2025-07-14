# nextgen{{classname}}

All URIs are relative to *https://app.harness.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccountDefaultNotificationTemplateSet**](AccountDefaultNotificationTemplateSetApi.md#CreateAccountDefaultNotificationTemplateSet) | **Post** /v1/default-notification-template-set | Create Default Notification Template Set
[**DeleteAccountDefaultNotificationTemplateSet**](AccountDefaultNotificationTemplateSetApi.md#DeleteAccountDefaultNotificationTemplateSet) | **Delete** /v1/default-notification-template-set/{identifier} | Delete Default Notification Template Set
[**GetAccountDefaultNotificationTemplateSet**](AccountDefaultNotificationTemplateSetApi.md#GetAccountDefaultNotificationTemplateSet) | **Get** /v1/default-notification-template-set/{identifier} | Get Default Notification Template Set
[**ListAccountDefaultNotificationTemplateSet**](AccountDefaultNotificationTemplateSetApi.md#ListAccountDefaultNotificationTemplateSet) | **Get** /v1/default-notification-template-set | List Default Notification Template Set
[**UpdateAccountDefaultNotificationTemplateSet**](AccountDefaultNotificationTemplateSetApi.md#UpdateAccountDefaultNotificationTemplateSet) | **Put** /v1/default-notification-template-set/{identifier} | Update Default Notification Template Set

# **CreateAccountDefaultNotificationTemplateSet**
> DefaultNotificationTemplateSetResponse CreateAccountDefaultNotificationTemplateSet(ctx, optional)
Create Default Notification Template Set

Create Default Notification Template Set

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountDefaultNotificationTemplateSetApiCreateAccountDefaultNotificationTemplateSetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountDefaultNotificationTemplateSetApiCreateAccountDefaultNotificationTemplateSetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**optional.Interface of DefaultNotificationTemplateSetDto**](DefaultNotificationTemplateSetDto.md)| Default Notification Template Set Request |
**harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. |

### Return type

[**DefaultNotificationTemplateSetResponse**](DefaultNotificationTemplateSetResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAccountDefaultNotificationTemplateSet**
> DeleteAccountDefaultNotificationTemplateSet(ctx, identifier, optional)
Delete Default Notification Template Set

Delete Default Notification Template Set

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| Identifier of Default Notification Template Set |
 **optional** | ***AccountDefaultNotificationTemplateSetApiDeleteAccountDefaultNotificationTemplateSetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountDefaultNotificationTemplateSetApiDeleteAccountDefaultNotificationTemplateSetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

**harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. |

### Return type

(empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountDefaultNotificationTemplateSet**
> DefaultNotificationTemplateSetResponse GetAccountDefaultNotificationTemplateSet(ctx, identifier, optional)
Get Default Notification Template Set

Get Default Notification Template Set

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| Identifier of Default Notification Template Set |
 **optional** | ***AccountDefaultNotificationTemplateSetApiGetAccountDefaultNotificationTemplateSetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountDefaultNotificationTemplateSetApiGetAccountDefaultNotificationTemplateSetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

**harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. |

### Return type

[**DefaultNotificationTemplateSetResponse**](DefaultNotificationTemplateSetResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAccountDefaultNotificationTemplateSet**
> []DefaultNotificationTemplateSetResponse ListAccountDefaultNotificationTemplateSet(ctx, optional)
List Default Notification Template Set

List Default Notification Template Sets based on filter criteria.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountDefaultNotificationTemplateSetApiListAccountDefaultNotificationTemplateSetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountDefaultNotificationTemplateSetApiListAccountDefaultNotificationTemplateSetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**searchTerm** | **optional.String**| This would be used to filter resources having attributes matching with search term. |
**identifiers** | [**optional.Interface of []string**](string.md)| Filter by default notification template set identifiers. |
**notificationChannelTypes** | [**optional.Interface of []string**](string.md)| Filter by one or more notification channel types. |
**notificationEvents** | [**optional.Interface of []string**](string.md)| Filter by one or more notification event types. |
**notificationEntities** | [**optional.Interface of []string**](string.md)| Filter by one or more notification entities. |
**harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. |
**page** | **optional.Int32**| Pagination page number strategy: Specify the page number within the paginated collection related to the number of items in each page  | [default to 0]
**limit** | **optional.Int32**| Number of items to return per page. | [default to 30]
**sort** | **optional.String**| Parameter on the basis of which sorting is done. |
**order** | **optional.String**| Order on the basis of which sorting is done. |

### Return type

[**[]DefaultNotificationTemplateSetResponse**](DefaultNotificationTemplateSetResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAccountDefaultNotificationTemplateSet**
> DefaultNotificationTemplateSetResponse UpdateAccountDefaultNotificationTemplateSet(ctx, identifier, optional)
Update Default Notification Template Set

Update Default Notification Template Set

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| Identifier of Default Notification Template Set |
 **optional** | ***AccountDefaultNotificationTemplateSetApiUpdateAccountDefaultNotificationTemplateSetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountDefaultNotificationTemplateSetApiUpdateAccountDefaultNotificationTemplateSetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

**body** | [**optional.Interface of DefaultNotificationTemplateSetDto**](DefaultNotificationTemplateSetDto.md)| Default Notification Template Set Request |
**harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. |

### Return type

[**DefaultNotificationTemplateSetResponse**](DefaultNotificationTemplateSetResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

