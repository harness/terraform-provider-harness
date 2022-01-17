# {{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddUsers**](UserApi.md#AddUsers) | **Post** /ng/api/user/users | Add user(s) to given scope
[**ChangeUserPassword**](UserApi.md#ChangeUserPassword) | **Put** /ng/api/user/password | Updates the User password
[**CheckIfLastAdmin**](UserApi.md#CheckIfLastAdmin) | **Get** /ng/api/user/last-admin | Boolean status whether the user is last admin at scope or not
[**DisableTTwoFactorAuth**](UserApi.md#DisableTTwoFactorAuth) | **Put** /ng/api/user/disable-two-factor-auth | Disables two-factor-auth for an user in an account
[**EnableTwoFactorAuth**](UserApi.md#EnableTwoFactorAuth) | **Put** /ng/api/user/enable-two-factor-auth | Enables two-factor-auth for an user in an account
[**GetAccessibleProjectsCount**](UserApi.md#GetAccessibleProjectsCount) | **Get** /ng/api/user/projects-count | Count of projects that are accessible to a user filtered by CreatedAt time
[**GetAggregatedUser**](UserApi.md#GetAggregatedUser) | **Get** /ng/api/user/aggregate/{userId} | Returns the user metadata along with rolesAssignments by userId and scope
[**GetAggregatedUsers**](UserApi.md#GetAggregatedUsers) | **Post** /ng/api/user/aggregate | List of all the user&#x27;s metadata along with rolesAssignments who have access to given scope
[**GetCurrentGenUsers**](UserApi.md#GetCurrentGenUsers) | **Get** /ng/api/user/currentgen | List of current gen users with the given Account Identifier
[**GetCurrentUserInfo**](UserApi.md#GetCurrentUserInfo) | **Get** /ng/api/user/currentUser | Gets current logged in User information
[**GetTwoFactorAuthSettings**](UserApi.md#GetTwoFactorAuthSettings) | **Get** /ng/api/user/two-factor-auth/{authMechanism} | Gets two factor authentication settings information of the current logged in user
[**GetUserAllProjectsInfo**](UserApi.md#GetUserAllProjectsInfo) | **Get** /ng/api/user/all-projects | list of project(s) of current user in the passed account Id in form of List
[**GetUserProjectInfo**](UserApi.md#GetUserProjectInfo) | **Get** /ng/api/user/projects | Retrieves the list of projects of the current user corresponding to the specified Account Identifier.
[**GetUsers**](UserApi.md#GetUsers) | **Post** /ng/api/user/batch | List of user&#x27;s Metadata for a given scope
[**RemoveUser**](UserApi.md#RemoveUser) | **Delete** /ng/api/user/{userId} | Remove user as the collaborator from the scope
[**UnlockUser**](UserApi.md#UnlockUser) | **Put** /ng/api/user/unlock-user/{userId} | unlock user in a given scope
[**UpdateUserInfo**](UserApi.md#UpdateUserInfo) | **Put** /ng/api/user | Updates the User information

# **AddUsers**
> ResponseDtoAddUsersResponse AddUsers(ctx, body, accountIdentifier, optional)
Add user(s) to given scope

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AddUsersDto**](AddUsersDto.md)|  | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
 **optional** | ***UserApiAddUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiAddUsersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity | 

### Return type

[**ResponseDtoAddUsersResponse**](ResponseDTOAddUsersResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChangeUserPassword**
> ResponseDtoPasswordChangeResponse ChangeUserPassword(ctx, optional)
Updates the User password

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UserApiChangeUserPasswordOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiChangeUserPasswordOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of PasswordChange**](PasswordChange.md)|  | 

### Return type

[**ResponseDtoPasswordChangeResponse**](ResponseDTOPasswordChangeResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CheckIfLastAdmin**
> ResponseDtoBoolean CheckIfLastAdmin(ctx, accountIdentifier, optional)
Boolean status whether the user is last admin at scope or not

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
 **optional** | ***UserApiCheckIfLastAdminOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiCheckIfLastAdminOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **optional.String**| User identifier | 
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

# **DisableTTwoFactorAuth**
> ResponseDtoUserInfo DisableTTwoFactorAuth(ctx, optional)
Disables two-factor-auth for an user in an account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UserApiDisableTTwoFactorAuthOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiDisableTTwoFactorAuthOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **routingId** | **optional.String**| Account Identifier for the Entity | 

### Return type

[**ResponseDtoUserInfo**](ResponseDTOUserInfo.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableTwoFactorAuth**
> ResponseDtoUserInfo EnableTwoFactorAuth(ctx, optional)
Enables two-factor-auth for an user in an account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UserApiEnableTwoFactorAuthOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiEnableTwoFactorAuthOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of TwoFactorAuthSettingsInfo**](TwoFactorAuthSettingsInfo.md)|  | 
 **routingId** | **optional.**| Account Identifier for the Entity | 

### Return type

[**ResponseDtoUserInfo**](ResponseDTOUserInfo.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccessibleProjectsCount**
> ResponseDtoActiveProjectsCount GetAccessibleProjectsCount(ctx, optional)
Count of projects that are accessible to a user filtered by CreatedAt time

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UserApiGetAccessibleProjectsCountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiGetAccessibleProjectsCountOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountIdentifier** | **optional.String**| Account Identifier for the Entity | 
 **userId** | **optional.String**| user Identifier | 
 **startTime** | **optional.Int64**| Start time to Filter projects by CreatedAt time | 
 **endTime** | **optional.Int64**| End time to Filter projects by CreatedAt time | 

### Return type

[**ResponseDtoActiveProjectsCount**](ResponseDTOActiveProjectsCount.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAggregatedUser**
> ResponseDtoUserAggregate GetAggregatedUser(ctx, userId, accountIdentifier, optional)
Returns the user metadata along with rolesAssignments by userId and scope

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**| user Identifier | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
 **optional** | ***UserApiGetAggregatedUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiGetAggregatedUserOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity | 

### Return type

[**ResponseDtoUserAggregate**](ResponseDTOUserAggregate.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAggregatedUsers**
> ResponseDtoPageResponseUserAggregate GetAggregatedUsers(ctx, accountIdentifier, optional)
List of all the user's metadata along with rolesAssignments who have access to given scope

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
 **optional** | ***UserApiGetAggregatedUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiGetAggregatedUsersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of AclAggregateFilter**](AclAggregateFilter.md)|  | 
 **orgIdentifier** | **optional.**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity | 
 **searchTerm** | **optional.**| Search term | 
 **pageIndex** | **optional.**| Indicates the number of pages. Results for these pages will be retrieved. | [default to 0]
 **pageSize** | **optional.**| The number of the elements to fetch | [default to 50]
 **sortOrders** | [**optional.Interface of []SortOrder**](SortOrder.md)| Sort criteria for the elements. | 

### Return type

[**ResponseDtoPageResponseUserAggregate**](ResponseDTOPageResponseUserAggregate.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCurrentGenUsers**
> ResponseDtoPageResponseUserMetadata GetCurrentGenUsers(ctx, accountIdentifier, optional)
List of current gen users with the given Account Identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| This is the Account Identifier. Users corresponding to this Account will be retrieved. | 
 **optional** | ***UserApiGetCurrentGenUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiGetCurrentGenUsersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **searchString** | **optional.String**| This string will be used to filter the search results. Details of all the users having this string in their name or email address will be filtered. | 
 **pageIndex** | **optional.Int32**| Indicates the number of pages. Results for these pages will be retrieved. | [default to 0]
 **pageSize** | **optional.Int32**| The number of the elements to fetch | [default to 50]
 **sortOrders** | [**optional.Interface of []SortOrder**](SortOrder.md)| Sort criteria for the elements. | 

### Return type

[**ResponseDtoPageResponseUserMetadata**](ResponseDTOPageResponseUserMetadata.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCurrentUserInfo**
> ResponseDtoUserInfo GetCurrentUserInfo(ctx, )
Gets current logged in User information

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ResponseDtoUserInfo**](ResponseDTOUserInfo.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTwoFactorAuthSettings**
> ResponseDtoTwoFactorAuthSettingsInfo GetTwoFactorAuthSettings(ctx, authMechanism)
Gets two factor authentication settings information of the current logged in user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authMechanism** | **string**| This is the authentication mechanism for the logged-in User. Two-Factor Authentication settings will be fetched for this mechanism. | 

### Return type

[**ResponseDtoTwoFactorAuthSettingsInfo**](ResponseDTOTwoFactorAuthSettingsInfo.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserAllProjectsInfo**
> ResponseDtoListProject GetUserAllProjectsInfo(ctx, optional)
list of project(s) of current user in the passed account Id in form of List

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UserApiGetUserAllProjectsInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiGetUserAllProjectsInfoOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **optional.String**| Account Identifier for the Entity | 
 **userId** | **optional.String**| User Identifier | 

### Return type

[**ResponseDtoListProject**](ResponseDTOListProject.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserProjectInfo**
> ResponseDtoPageResponseProject GetUserProjectInfo(ctx, optional)
Retrieves the list of projects of the current user corresponding to the specified Account Identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UserApiGetUserProjectInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiGetUserProjectInfoOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **optional.String**| This is the Account Identifier. Details of all the Projects within the scope of this Account will be fetched. | 
 **pageIndex** | **optional.Int32**| Indicates the number of pages. Results for these pages will be retrieved. | [default to 0]
 **pageSize** | **optional.Int32**| The number of the elements to fetch | [default to 50]
 **sortOrders** | [**optional.Interface of []SortOrder**](SortOrder.md)| Sort criteria for the elements. | 

### Return type

[**ResponseDtoPageResponseProject**](ResponseDTOPageResponseProject.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsers**
> ResponseDtoPageResponseUserMetadata GetUsers(ctx, accountIdentifier, optional)
List of user's Metadata for a given scope

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
 **optional** | ***UserApiGetUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiGetUsersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UserFilter**](UserFilter.md)|  | 
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

# **RemoveUser**
> ResponseDtoBoolean RemoveUser(ctx, userId, accountIdentifier, optional)
Remove user as the collaborator from the scope

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**| user Identifier | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
 **optional** | ***UserApiRemoveUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiRemoveUserOpts struct
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

# **UnlockUser**
> ResponseDtoUserInfo UnlockUser(ctx, userId, accountIdentifier, optional)
unlock user in a given scope

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**| user Identifier | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
 **optional** | ***UserApiUnlockUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiUnlockUserOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity | 

### Return type

[**ResponseDtoUserInfo**](ResponseDTOUserInfo.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUserInfo**
> ResponseDtoUserInfo UpdateUserInfo(ctx, optional)
Updates the User information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UserApiUpdateUserInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiUpdateUserInfoOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of UserInfo**](UserInfo.md)|  | 

### Return type

[**ResponseDtoUserInfo**](ResponseDTOUserInfo.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

