# {{classname}}

All URIs are relative to *https://app.harness.io/gateway/authz/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteRole**](RolesApi.md#DeleteRole) | **Delete** /roles/{identifier} | Delete a Custom Role in a scope
[**GetRole**](RolesApi.md#GetRole) | **Get** /roles/{identifier} | Get a Role by identifier
[**GetRoleList**](RolesApi.md#GetRoleList) | **Get** /roles | List roles in the given scope
[**PostRole**](RolesApi.md#PostRole) | **Post** /roles | Create a Custom Role in a scope
[**PutRole**](RolesApi.md#PutRole) | **Put** /roles/{identifier} | Update a Custom Role by identifier

# **DeleteRole**
> ResponseDtoRoleResponse DeleteRole(ctx, identifier, optional)
Delete a Custom Role in a scope

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**|  | 
 **optional** | ***RolesApiDeleteRoleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RolesApiDeleteRoleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.String**|  | 
 **orgIdentifier** | **optional.String**|  | 
 **projectIdentifier** | **optional.String**|  | 

### Return type

[**ResponseDtoRoleResponse**](ResponseDTORoleResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRole**
> ResponseDtoRoleResponse GetRole(ctx, identifier, optional)
Get a Role by identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**|  | 
 **optional** | ***RolesApiGetRoleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RolesApiGetRoleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.String**|  | 
 **orgIdentifier** | **optional.String**|  | 
 **projectIdentifier** | **optional.String**|  | 

### Return type

[**ResponseDtoRoleResponse**](ResponseDTORoleResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRoleList**
> ResponseDtoPageResponseRoleResponse GetRoleList(ctx, optional)
List roles in the given scope

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RolesApiGetRoleListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RolesApiGetRoleListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageIndex** | **optional.Int32**|  | [default to 0]
 **pageSize** | **optional.Int32**|  | [default to 50]
 **sortOrders** | [**optional.Interface of []SortOrder**](SortOrder.md)|  | 
 **accountIdentifier** | **optional.String**|  | 
 **orgIdentifier** | **optional.String**|  | 
 **projectIdentifier** | **optional.String**|  | 
 **searchTerm** | **optional.String**|  | 

### Return type

[**ResponseDtoPageResponseRoleResponse**](ResponseDTOPageResponseRoleResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostRole**
> ResponseDtoRoleResponse PostRole(ctx, optional)
Create a Custom Role in a scope

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RolesApiPostRoleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RolesApiPostRoleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Role**](Role.md)|  | 
 **accountIdentifier** | **optional.**|  | 
 **orgIdentifier** | **optional.**|  | 
 **projectIdentifier** | **optional.**|  | 

### Return type

[**ResponseDtoRoleResponse**](ResponseDTORoleResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutRole**
> ResponseDtoRoleResponse PutRole(ctx, identifier, optional)
Update a Custom Role by identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**|  | 
 **optional** | ***RolesApiPutRoleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RolesApiPutRoleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Role**](Role.md)|  | 
 **accountIdentifier** | **optional.**|  | 
 **orgIdentifier** | **optional.**|  | 
 **projectIdentifier** | **optional.**|  | 

### Return type

[**ResponseDtoRoleResponse**](ResponseDTORoleResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

