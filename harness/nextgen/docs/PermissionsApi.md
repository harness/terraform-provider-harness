# {{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPermissionList**](PermissionsApi.md#GetPermissionList) | **Get** /authz/api/permissions | Get all permissions in a scope or all permissions in the system.
[**GetPermissionResourceTypesList**](PermissionsApi.md#GetPermissionResourceTypesList) | **Get** /authz/api/permissions/resourcetypes | Get all resource types for permissions in a scope or in the system.

# **GetPermissionList**
> ResponseDtoListPermissionResponse GetPermissionList(ctx, optional)
Get all permissions in a scope or all permissions in the system.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PermissionsApiGetPermissionListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PermissionsApiGetPermissionListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountIdentifier** | **optional.String**| Account Identifier for the Entity | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity | 
 **scopeFilterDisabled** | **optional.Bool**| This is to enable or disable filtering by scope. The default value is false. If the value is true, all the permissions in the system are fetched. | 

### Return type

[**ResponseDtoListPermissionResponse**](ResponseDTOListPermissionResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPermissionResourceTypesList**
> ResponseDtoSetString GetPermissionResourceTypesList(ctx, optional)
Get all resource types for permissions in a scope or in the system.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PermissionsApiGetPermissionResourceTypesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PermissionsApiGetPermissionResourceTypesListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountIdentifier** | **optional.String**| Account Identifier for the Entity | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity | 
 **scopeFilterDisabled** | **optional.Bool**| This is to enable or disable filtering by scope. The default value is false. If the value is true, all the permissions in the system are fetched. | 

### Return type

[**ResponseDtoSetString**](ResponseDTOSetString.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

