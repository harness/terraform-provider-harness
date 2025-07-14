# nextgen{{classname}}

All URIs are relative to *https://app.harness.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProjectDefaultNotificationTemplateSet**](ProjectDefaultNotificationTemplateSetApi.md#CreateProjectDefaultNotificationTemplateSet) | **Post** /v1/orgs/{org}/projects/{project}/default-notification-template-set | Create Default Notification Template Set
[**DeleteProjectDefaultNotificationTemplateSet**](ProjectDefaultNotificationTemplateSetApi.md#DeleteProjectDefaultNotificationTemplateSet) | **Delete** /v1/orgs/{org}/projects/{project}/default-notification-template-set/{identifier} | Delete Default Notification Template Set
[**GetProjectDefaultNotificationTemplateSet**](ProjectDefaultNotificationTemplateSetApi.md#GetProjectDefaultNotificationTemplateSet) | **Get** /v1/orgs/{org}/projects/{project}/default-notification-template-set/{identifier} | Get Default Notification Template Set
[**ListProjectDefaultNotificationTemplateSet**](ProjectDefaultNotificationTemplateSetApi.md#ListProjectDefaultNotificationTemplateSet) | **Get** /v1/orgs/{org}/projects/{project}/default-notification-template-set | List Default Notification Template Set
[**UpdateProjectDefaultNotificationTemplateSet**](ProjectDefaultNotificationTemplateSetApi.md#UpdateProjectDefaultNotificationTemplateSet) | **Put** /v1/orgs/{org}/projects/{project}/default-notification-template-set/{identifier} | Update Default Notification Template Set

# **CreateProjectDefaultNotificationTemplateSet**
> DefaultNotificationTemplateSetResponse CreateProjectDefaultNotificationTemplateSet(ctx, org, project, optional)
Create Default Notification Template Set

Create Default Notification Template Set

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to |
  **project** | **string**| Identifier field of the project the resource is scoped to |
 **optional** | ***ProjectDefaultNotificationTemplateSetApiCreateProjectDefaultNotificationTemplateSetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectDefaultNotificationTemplateSetApiCreateProjectDefaultNotificationTemplateSetOpts struct
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

# **DeleteProjectDefaultNotificationTemplateSet**
> DeleteProjectDefaultNotificationTemplateSet(ctx, identifier, org, project, optional)
Delete Default Notification Template Set

Delete Default Notification Template Set

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| Identifier of Default Notification Template Set |
  **org** | **string**| Identifier field of the organization the resource is scoped to |
  **project** | **string**| Identifier field of the project the resource is scoped to |
 **optional** | ***ProjectDefaultNotificationTemplateSetApiDeleteProjectDefaultNotificationTemplateSetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectDefaultNotificationTemplateSetApiDeleteProjectDefaultNotificationTemplateSetOpts struct
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

# **GetProjectDefaultNotificationTemplateSet**
> DefaultNotificationTemplateSetResponse GetProjectDefaultNotificationTemplateSet(ctx, identifier, org, project, optional)
Get Default Notification Template Set

Get Default Notification Template Set

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| Identifier of Default Notification Template Set |
  **org** | **string**| Identifier field of the organization the resource is scoped to |
  **project** | **string**| Identifier field of the project the resource is scoped to |
 **optional** | ***ProjectDefaultNotificationTemplateSetApiGetProjectDefaultNotificationTemplateSetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectDefaultNotificationTemplateSetApiGetProjectDefaultNotificationTemplateSetOpts struct
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

# **ListProjectDefaultNotificationTemplateSet**
> []DefaultNotificationTemplateSetResponse ListProjectDefaultNotificationTemplateSet(ctx, org, project, optional)
List Default Notification Template Set

List Default Notification Template Sets based on filter criteria.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to |
  **project** | **string**| Identifier field of the project the resource is scoped to |
 **optional** | ***ProjectDefaultNotificationTemplateSetApiListProjectDefaultNotificationTemplateSetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectDefaultNotificationTemplateSetApiListProjectDefaultNotificationTemplateSetOpts struct
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

# **UpdateProjectDefaultNotificationTemplateSet**
> DefaultNotificationTemplateSetResponse UpdateProjectDefaultNotificationTemplateSet(ctx, identifier, org, project, optional)
Update Default Notification Template Set

Update Default Notification Template Set

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| Identifier of Default Notification Template Set |
  **org** | **string**| Identifier field of the organization the resource is scoped to |
  **project** | **string**| Identifier field of the project the resource is scoped to |
 **optional** | ***ProjectDefaultNotificationTemplateSetApiUpdateProjectDefaultNotificationTemplateSetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectDefaultNotificationTemplateSetApiUpdateProjectDefaultNotificationTemplateSetOpts struct
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

