# nextgen{{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteRole**](RolesApi.md#DeleteRole) | **Delete** /authz/api/roles/{identifier} | Delete Role
[**GetRole**](RolesApi.md#GetRole) | **Get** /authz/api/roles/{identifier} | Get Role
[**GetRoleList**](RolesApi.md#GetRoleList) | **Get** /authz/api/roles | List Roles
[**PostRole**](RolesApi.md#PostRole) | **Post** /authz/api/roles | Create Role
[**PutRole**](RolesApi.md#PutRole) | **Put** /authz/api/roles/{identifier} | Update Role

# **DeleteRole**
> ResponseDtoRoleResponse DeleteRole(ctx, identifier, optional)
Delete Role

Delete a Custom Role in a scope

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| Identifier of the Role | 
 **optional** | ***RolesApiDeleteRoleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RolesApiDeleteRoleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.String**| Account Identifier for the Entity. | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 

### Return type

[**ResponseDtoRoleResponse**](ResponseDTORoleResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRole**
> ResponseDtoRoleResponse GetRole(ctx, identifier, optional)
Get Role

Get a Role by identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| Identifier of the Role | 
 **optional** | ***RolesApiGetRoleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RolesApiGetRoleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.String**| Account Identifier for the Entity. | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 

### Return type

[**ResponseDtoRoleResponse**](ResponseDTORoleResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRoleList**
> ResponseDtoPageResponseRoleResponse GetRoleList(ctx, optional)
List Roles

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
 **pageIndex** | **optional.Int32**| Page Index of the results to fetch.Default Value: 0 | [default to 0]
 **pageSize** | **optional.Int32**| Results per page(max 100)Default Value: 50 | [default to 50]
 **sortOrders** | [**optional.Interface of []SortOrder**](SortOrder.md)| Sort criteria for the elements. | 
 **accountIdentifier** | **optional.String**| Account Identifier for the Entity. | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **searchTerm** | **optional.String**| Search roles by name/identifier | 

### Return type

[**ResponseDtoPageResponseRoleResponse**](ResponseDTOPageResponseRoleResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostRole**
> ResponseDtoRoleResponse PostRole(ctx, body, optional)
Create Role

Create a Custom Role in a scope

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Role**](Role.md)| Role entity | 
 **optional** | ***RolesApiPostRoleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RolesApiPostRoleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.**| Account Identifier for the Entity. | 
 **orgIdentifier** | **optional.**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity. | 

### Return type

[**ResponseDtoRoleResponse**](ResponseDTORoleResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutRole**
> ResponseDtoRoleResponse PutRole(ctx, body, identifier, optional)
Update Role

Update a Custom Role by identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Role**](Role.md)| Updated Role entity | 
  **identifier** | **string**| Identifier of the Role | 
 **optional** | ***RolesApiPutRoleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RolesApiPutRoleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accountIdentifier** | **optional.**| Account Identifier for the Entity. | 
 **orgIdentifier** | **optional.**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity. | 

### Return type

[**ResponseDtoRoleResponse**](ResponseDTORoleResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

