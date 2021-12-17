# {{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteRoleAssignment**](RoleAssignmentsApi.md#DeleteRoleAssignment) | **Delete** /authz/api/roleassignments/{identifier} | Delete an existing role assignment by identifier
[**GetFilteredRoleAssignmentList**](RoleAssignmentsApi.md#GetFilteredRoleAssignmentList) | **Post** /authz/api/roleassignments/filter | List role assignments in the scope according to the given filter
[**GetRoleAssignmentAggregateList**](RoleAssignmentsApi.md#GetRoleAssignmentAggregateList) | **Post** /authz/api/roleassignments/aggregate | List role assignments in the scope according to the given filter with added metadata
[**GetRoleAssignmentList**](RoleAssignmentsApi.md#GetRoleAssignmentList) | **Get** /authz/api/roleassignments | List role assignments in the given scope
[**PostRoleAssignment**](RoleAssignmentsApi.md#PostRoleAssignment) | **Post** /authz/api/roleassignments | Create role assignment in the given scope
[**PostRoleAssignments**](RoleAssignmentsApi.md#PostRoleAssignments) | **Post** /authz/api/roleassignments/multi | Create multiple role assignments in a scope. Returns all successfully created role assignments. Ignores failures and duplicates.
[**PutRoleAssignment**](RoleAssignmentsApi.md#PutRoleAssignment) | **Put** /authz/api/roleassignments/{identifier} | Update existing role assignment by identifier and scope. Only changing the disabled/enabled state is allowed.
[**ValidateRoleAssignment**](RoleAssignmentsApi.md#ValidateRoleAssignment) | **Post** /authz/api/roleassignments/validate | Check whether a proposed role assignment is valid.

# **DeleteRoleAssignment**
> ResponseDtoRoleAssignmentResponse DeleteRoleAssignment(ctx, identifier, optional)
Delete an existing role assignment by identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**|  | 
 **optional** | ***RoleAssignmentsApiDeleteRoleAssignmentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RoleAssignmentsApiDeleteRoleAssignmentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.String**|  | 
 **orgIdentifier** | **optional.String**|  | 
 **projectIdentifier** | **optional.String**|  | 

### Return type

[**ResponseDtoRoleAssignmentResponse**](ResponseDTORoleAssignmentResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFilteredRoleAssignmentList**
> ResponseDtoPageResponseRoleAssignmentResponse GetFilteredRoleAssignmentList(ctx, body, optional)
List role assignments in the scope according to the given filter

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RoleAssignmentFilter**](RoleAssignmentFilter.md)| Filter role assignments based on multiple parameters. | 
 **optional** | ***RoleAssignmentsApiGetFilteredRoleAssignmentListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RoleAssignmentsApiGetFilteredRoleAssignmentListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageIndex** | **optional.**|  | [default to 0]
 **pageSize** | **optional.**|  | [default to 50]
 **sortOrders** | [**optional.Interface of []SortOrder**](SortOrder.md)|  | 
 **accountIdentifier** | **optional.**|  | 
 **orgIdentifier** | **optional.**|  | 
 **projectIdentifier** | **optional.**|  | 

### Return type

[**ResponseDtoPageResponseRoleAssignmentResponse**](ResponseDTOPageResponseRoleAssignmentResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRoleAssignmentAggregateList**
> ResponseDtoRoleAssignmentAggregateResponse GetRoleAssignmentAggregateList(ctx, body, optional)
List role assignments in the scope according to the given filter with added metadata

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RoleAssignmentFilter**](RoleAssignmentFilter.md)| Filter role assignments based on multiple parameters. | 
 **optional** | ***RoleAssignmentsApiGetRoleAssignmentAggregateListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RoleAssignmentsApiGetRoleAssignmentAggregateListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.**|  | 
 **orgIdentifier** | **optional.**|  | 
 **projectIdentifier** | **optional.**|  | 

### Return type

[**ResponseDtoRoleAssignmentAggregateResponse**](ResponseDTORoleAssignmentAggregateResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRoleAssignmentList**
> ResponseDtoPageResponseRoleAssignmentResponse GetRoleAssignmentList(ctx, optional)
List role assignments in the given scope

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RoleAssignmentsApiGetRoleAssignmentListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RoleAssignmentsApiGetRoleAssignmentListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageIndex** | **optional.Int32**|  | [default to 0]
 **pageSize** | **optional.Int32**|  | [default to 50]
 **sortOrders** | [**optional.Interface of []SortOrder**](SortOrder.md)|  | 
 **accountIdentifier** | **optional.String**|  | 
 **orgIdentifier** | **optional.String**|  | 
 **projectIdentifier** | **optional.String**|  | 

### Return type

[**ResponseDtoPageResponseRoleAssignmentResponse**](ResponseDTOPageResponseRoleAssignmentResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostRoleAssignment**
> ResponseDtoRoleAssignmentResponse PostRoleAssignment(ctx, optional)
Create role assignment in the given scope

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RoleAssignmentsApiPostRoleAssignmentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RoleAssignmentsApiPostRoleAssignmentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of RoleAssignment**](RoleAssignment.md)|  | 
 **accountIdentifier** | **optional.**|  | 
 **orgIdentifier** | **optional.**|  | 
 **projectIdentifier** | **optional.**|  | 

### Return type

[**ResponseDtoRoleAssignmentResponse**](ResponseDTORoleAssignmentResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostRoleAssignments**
> ResponseDtoListRoleAssignmentResponse PostRoleAssignments(ctx, optional)
Create multiple role assignments in a scope. Returns all successfully created role assignments. Ignores failures and duplicates.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RoleAssignmentsApiPostRoleAssignmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RoleAssignmentsApiPostRoleAssignmentsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of RoleAssignmentCreateRequest**](RoleAssignmentCreateRequest.md)|  | 
 **accountIdentifier** | **optional.**|  | 
 **orgIdentifier** | **optional.**|  | 
 **projectIdentifier** | **optional.**|  | 

### Return type

[**ResponseDtoListRoleAssignmentResponse**](ResponseDTOListRoleAssignmentResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutRoleAssignment**
> ResponseDtoRoleAssignmentResponse PutRoleAssignment(ctx, identifier, optional)
Update existing role assignment by identifier and scope. Only changing the disabled/enabled state is allowed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**|  | 
 **optional** | ***RoleAssignmentsApiPutRoleAssignmentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RoleAssignmentsApiPutRoleAssignmentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of RoleAssignment**](RoleAssignment.md)|  | 
 **accountIdentifier** | **optional.**|  | 
 **orgIdentifier** | **optional.**|  | 
 **projectIdentifier** | **optional.**|  | 

### Return type

[**ResponseDtoRoleAssignmentResponse**](ResponseDTORoleAssignmentResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateRoleAssignment**
> ResponseDtoRoleAssignmentValidationResponse ValidateRoleAssignment(ctx, optional)
Check whether a proposed role assignment is valid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RoleAssignmentsApiValidateRoleAssignmentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RoleAssignmentsApiValidateRoleAssignmentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of RoleAssignmentValidationRequest**](RoleAssignmentValidationRequest.md)|  | 
 **accountIdentifier** | **optional.**|  | 
 **orgIdentifier** | **optional.**|  | 
 **projectIdentifier** | **optional.**|  | 

### Return type

[**ResponseDtoRoleAssignmentValidationResponse**](ResponseDTORoleAssignmentValidationResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

