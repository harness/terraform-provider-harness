# nextgen{{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddUsers**](UserApi.md#AddUsers) | **Post** /ng/api/user/users | Add user(s) to scope
[**ChangeUserPassword**](UserApi.md#ChangeUserPassword) | **Put** /ng/api/user/password | Change user password
[**CheckIfLastAdmin**](UserApi.md#CheckIfLastAdmin) | **Get** /ng/api/user/last-admin | Check if user is last admin
[**DisableTTwoFactorAuth**](UserApi.md#DisableTTwoFactorAuth) | **Put** /ng/api/user/disable-two-factor-auth | Disable two factor authentication
[**EnableTwoFactorAuth**](UserApi.md#EnableTwoFactorAuth) | **Put** /ng/api/user/enable-two-factor-auth | Enable two factor authentication
[**GetAggregatedUser**](UserApi.md#GetAggregatedUser) | **Get** /ng/api/user/aggregate/{userId} | Get detailed user information
[**GetAggregatedUsers**](UserApi.md#GetAggregatedUsers) | **Post** /ng/api/user/aggregate | Get list of users
[**GetCurrentUserInfo**](UserApi.md#GetCurrentUserInfo) | **Get** /ng/api/user/currentUser | Get Current User Info
[**GetTwoFactorAuthSettings**](UserApi.md#GetTwoFactorAuthSettings) | **Get** /ng/api/user/two-factor-auth/{authMechanism} | Gets Two Factor Auth Settings
[**GetUsers**](UserApi.md#GetUsers) | **Post** /ng/api/user/batch | Get users list
[**RemoveUser**](UserApi.md#RemoveUser) | **Delete** /ng/api/user/{userId} | Remove user from scope
[**UnlockUser**](UserApi.md#UnlockUser) | **Put** /ng/api/user/unlock-user/{userId} | Unlock user
[**UpdateUserInfo**](UserApi.md#UpdateUserInfo) | **Put** /ng/api/user | Update User

# **AddUsers**
> ResponseDtoAddUsersResponse AddUsers(ctx, body, accountIdentifier, optional)
Add user(s) to scope

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AddUsersDto**](AddUsersDto.md)|  | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***UserApiAddUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiAddUsersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity. | 

### Return type

[**ResponseDtoAddUsersResponse**](ResponseDTOAddUsersResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChangeUserPassword**
> ResponseDtoPasswordChangeResponse ChangeUserPassword(ctx, accountIdentifier, optional)
Change user password

Updates the User password

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
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
Check if user is last admin

Check whether the user is last admin at scope or not

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***UserApiCheckIfLastAdminOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiCheckIfLastAdminOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **optional.String**| User identifier | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 

### Return type

[**ResponseDtoBoolean**](ResponseDTOBoolean.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisableTTwoFactorAuth**
> ResponseDtoUserInfo DisableTTwoFactorAuth(ctx, accountIdentifier)
Disable two factor authentication

Disables two-factor-auth for an user in an account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 

### Return type

[**ResponseDtoUserInfo**](ResponseDTOUserInfo.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableTwoFactorAuth**
> ResponseDtoUserInfo EnableTwoFactorAuth(ctx, accountIdentifier, optional)
Enable two factor authentication

Enables two-factor-auth for an user in an account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***UserApiEnableTwoFactorAuthOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiEnableTwoFactorAuthOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of TwoFactorAuthSettingsInfo**](TwoFactorAuthSettingsInfo.md)|  | 

### Return type

[**ResponseDtoUserInfo**](ResponseDTOUserInfo.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAggregatedUser**
> ResponseDtoUserAggregate GetAggregatedUser(ctx, userId, accountIdentifier, optional)
Get detailed user information

Returns the user metadata along with rolesAssignments by userId and scope

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**| user Identifier | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***UserApiGetAggregatedUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiGetAggregatedUserOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 

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
Get list of users

List of all the user's metadata along with rolesAssignments who have access to given scope

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***UserApiGetAggregatedUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiGetAggregatedUsersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of AclAggregateFilter**](AclAggregateFilter.md)|  | 
 **orgIdentifier** | **optional.**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity. | 
 **searchTerm** | **optional.**| Search term | 
 **pageIndex** | **optional.**| Page Index of the results to fetch.Default Value: 0 | [default to 0]
 **pageSize** | **optional.**| Results per page(max 100)Default Value: 50 | [default to 50]
 **sortOrders** | [**optional.Interface of []SortOrder**](SortOrder.md)| Sort criteria for the elements. | 

### Return type

[**ResponseDtoPageResponseUserAggregate**](ResponseDTOPageResponseUserAggregate.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCurrentUserInfo**
> ResponseDtoUserInfo GetCurrentUserInfo(ctx, accountIdentifier)
Get Current User Info

Gets current logged in User information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 

### Return type

[**ResponseDtoUserInfo**](ResponseDTOUserInfo.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTwoFactorAuthSettings**
> ResponseDtoTwoFactorAuthSettingsInfo GetTwoFactorAuthSettings(ctx, accountIdentifier, authMechanism)
Gets Two Factor Auth Settings

Gets two factor authentication settings information of the current logged in user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **authMechanism** | **string**| This is the authentication mechanism for the logged-in User. Two-Factor Authentication settings will be fetched for this mechanism. | 

### Return type

[**ResponseDtoTwoFactorAuthSettingsInfo**](ResponseDTOTwoFactorAuthSettingsInfo.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsers**
> ResponseDtoPageResponseUserMetadata GetUsers(ctx, accountIdentifier, optional)
Get users list

Get list of user's for a given scope

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***UserApiGetUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiGetUsersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UserFilter**](UserFilter.md)|  | 
 **orgIdentifier** | **optional.**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity. | 
 **pageIndex** | **optional.**| Page Index of the results to fetch.Default Value: 0 | [default to 0]
 **pageSize** | **optional.**| Results per page(max 100)Default Value: 50 | [default to 50]
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
Remove user from scope

Remove user as the collaborator from the scope

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**| user Identifier | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***UserApiRemoveUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiRemoveUserOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 

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
Unlock user

unlock user in a given scope

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**| user Identifier | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***UserApiUnlockUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiUnlockUserOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 

### Return type

[**ResponseDtoUserInfo**](ResponseDTOUserInfo.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUserInfo**
> ResponseDtoUserInfo UpdateUserInfo(ctx, accountIdentifier, optional)
Update User

Updates the User information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
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

