# {{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CopyUserGroup**](UserGroupApi.md#CopyUserGroup) | **Put** /ng/api/user-groups/copy | Get a User Group in an account/org/project
[**DeleteMember**](UserGroupApi.md#DeleteMember) | **Delete** /ng/api/user-groups/{identifier}/member/{userIdentifier} | Remove a user from the user group in an account/org/project
[**DeleteUserGroup**](UserGroupApi.md#DeleteUserGroup) | **Delete** /ng/api/user-groups/{identifier} | Delete a User Group in an account/org/project
[**GetBatchUsersGroupList**](UserGroupApi.md#GetBatchUsersGroupList) | **Post** /ng/api/user-groups/batch | List the User Groups selected by a filter in an account/org/project
[**GetMember**](UserGroupApi.md#GetMember) | **Get** /ng/api/user-groups/{identifier}/member/{userIdentifier} | Check if the user is part of the user group in an account/org/project
[**GetUserGroup**](UserGroupApi.md#GetUserGroup) | **Get** /ng/api/user-groups/{identifier} | Get a User Group in an account/org/project
[**GetUserGroupList**](UserGroupApi.md#GetUserGroupList) | **Get** /ng/api/user-groups | List the User Groups in an account/org/project
[**GetUserListInUserGroup**](UserGroupApi.md#GetUserListInUserGroup) | **Post** /ng/api/user-groups/{identifier}/users | List the users in a User Group in an account/org/project
[**LinkUserGroupToSAML**](UserGroupApi.md#LinkUserGroupToSAML) | **Put** /ng/api/user-groups/{userGroupId}/link/saml/{samlId} | Link SAML Group to the User Group in an account/org/project
[**PostUserGroup**](UserGroupApi.md#PostUserGroup) | **Post** /ng/api/user-groups | Create a User Group in an account/org/project
[**PutMember**](UserGroupApi.md#PutMember) | **Put** /ng/api/user-groups/{identifier}/member/{userIdentifier} | Add a user to the user group in an account/org/project
[**PutUserGroup**](UserGroupApi.md#PutUserGroup) | **Put** /ng/api/user-groups | Update a User Group in an account/org/project
[**UnlinkUserGroupfromSSO**](UserGroupApi.md#UnlinkUserGroupfromSSO) | **Put** /ng/api/user-groups/{userGroupId}/unlink | Unlink SSO Group from the User Group in an account/org/project

# **CopyUserGroup**
> ResponseDtoBoolean CopyUserGroup(ctx, body, accountIdentifier, groupIdentifier)
Get a User Group in an account/org/project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]Scope**](Scope.md)|  | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
  **groupIdentifier** | **string**| groupIdentifier | 

### Return type

[**ResponseDtoBoolean**](ResponseDTOBoolean.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMember**
> ResponseDtoUserGroup DeleteMember(ctx, accountIdentifier, identifier, userIdentifier, optional)
Remove a user from the user group in an account/org/project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
  **identifier** | **string**| Identifier of the user group | 
  **userIdentifier** | **string**| Identifier of the user | 
 **optional** | ***UserGroupApiDeleteMemberOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserGroupApiDeleteMemberOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity | 

### Return type

[**ResponseDtoUserGroup**](ResponseDTOUserGroup.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUserGroup**
> ResponseDtoUserGroup DeleteUserGroup(ctx, accountIdentifier, identifier, optional)
Delete a User Group in an account/org/project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
  **identifier** | **string**| Identifier of the user group | 
 **optional** | ***UserGroupApiDeleteUserGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserGroupApiDeleteUserGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity | 

### Return type

[**ResponseDtoUserGroup**](ResponseDTOUserGroup.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBatchUsersGroupList**
> ResponseDtoListUserGroup GetBatchUsersGroupList(ctx, body)
List the User Groups selected by a filter in an account/org/project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UserGroupFilter**](UserGroupFilter.md)| User Group Filter | 

### Return type

[**ResponseDtoListUserGroup**](ResponseDTOListUserGroup.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMember**
> ResponseDtoBoolean GetMember(ctx, accountIdentifier, identifier, userIdentifier, optional)
Check if the user is part of the user group in an account/org/project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
  **identifier** | **string**| Identifier of the user group | 
  **userIdentifier** | **string**| Identifier of the user | 
 **optional** | ***UserGroupApiGetMemberOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserGroupApiGetMemberOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity | 

### Return type

[**ResponseDtoBoolean**](ResponseDTOBoolean.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserGroup**
> ResponseDtoUserGroup GetUserGroup(ctx, accountIdentifier, identifier, optional)
Get a User Group in an account/org/project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
  **identifier** | **string**| Identifier of the user group | 
 **optional** | ***UserGroupApiGetUserGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserGroupApiGetUserGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity | 

### Return type

[**ResponseDtoUserGroup**](ResponseDTOUserGroup.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserGroupList**
> ResponseDtoPageResponseUserGroup GetUserGroupList(ctx, accountIdentifier, optional)
List the User Groups in an account/org/project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
 **optional** | ***UserGroupApiGetUserGroupListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserGroupApiGetUserGroupListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity | 
 **searchTerm** | **optional.String**| Search filter which matches by user group name/identifier | 
 **pageIndex** | **optional.Int32**| Indicates the number of pages. Results for these pages will be retrieved. | [default to 0]
 **pageSize** | **optional.Int32**| The number of the elements to fetch | [default to 50]
 **sortOrders** | [**optional.Interface of []SortOrder**](SortOrder.md)| Sort criteria for the elements. | 

### Return type

[**ResponseDtoPageResponseUserGroup**](ResponseDTOPageResponseUserGroup.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserListInUserGroup**
> ResponseDtoPageResponseUserMetadata GetUserListInUserGroup(ctx, identifier, accountIdentifier, optional)
List the users in a User Group in an account/org/project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| Identifier of the user group | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
 **optional** | ***UserGroupApiGetUserListInUserGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserGroupApiGetUserListInUserGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of UserFilter**](UserFilter.md)| Filter users based on multiple parameters | 
 **orgIdentifier** | **optional.**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity | 
 **pageIndex** | **optional.**| Indicates the number of pages. Results for these pages will be retrieved. | [default to 0]
 **pageSize** | **optional.**| The number of the elements to fetch | [default to 50]
 **sortOrders** | [**optional.Interface of []SortOrder**](SortOrder.md)| Sort criteria for the elements. | 

### Return type

[**ResponseDtoPageResponseUserMetadata**](ResponseDTOPageResponseUserMetadata.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LinkUserGroupToSAML**
> RestResponseUserGroup LinkUserGroupToSAML(ctx, body, userGroupId, samlId, accountIdentifier, optional)
Link SAML Group to the User Group in an account/org/project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SamlLinkGroupRequest**](SamlLinkGroupRequest.md)| Saml Link Group Request | 
  **userGroupId** | **string**| Identifier of the user group | 
  **samlId** | **string**| Saml Group entity identifier | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
 **optional** | ***UserGroupApiLinkUserGroupToSAMLOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserGroupApiLinkUserGroupToSAMLOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **orgIdentifier** | **optional.**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity | 

### Return type

[**RestResponseUserGroup**](RestResponseUserGroup.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostUserGroup**
> ResponseDtoUserGroup PostUserGroup(ctx, body, accountIdentifier, optional)
Create a User Group in an account/org/project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UserGroup**](UserGroup.md)| User Group entity to be created | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
 **optional** | ***UserGroupApiPostUserGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserGroupApiPostUserGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity | 

### Return type

[**ResponseDtoUserGroup**](ResponseDTOUserGroup.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutMember**
> ResponseDtoUserGroup PutMember(ctx, accountIdentifier, identifier, userIdentifier, optional)
Add a user to the user group in an account/org/project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
  **identifier** | **string**| Identifier of the user group | 
  **userIdentifier** | **string**| Identifier of the user | 
 **optional** | ***UserGroupApiPutMemberOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserGroupApiPutMemberOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity | 

### Return type

[**ResponseDtoUserGroup**](ResponseDTOUserGroup.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutUserGroup**
> ResponseDtoUserGroup PutUserGroup(ctx, body, accountIdentifier, optional)
Update a User Group in an account/org/project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UserGroup**](UserGroup.md)| User Group entity with the updates | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
 **optional** | ***UserGroupApiPutUserGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserGroupApiPutUserGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity | 

### Return type

[**ResponseDtoUserGroup**](ResponseDTOUserGroup.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnlinkUserGroupfromSSO**
> RestResponseUserGroup UnlinkUserGroupfromSSO(ctx, userGroupId, accountIdentifier, optional)
Unlink SSO Group from the User Group in an account/org/project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userGroupId** | **string**| Identifier of the user group | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
 **optional** | ***UserGroupApiUnlinkUserGroupfromSSOOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserGroupApiUnlinkUserGroupfromSSOOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **retainMembers** | **optional.Bool**| Retain currently synced members of the user group | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity | 

### Return type

[**RestResponseUserGroup**](RestResponseUserGroup.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

