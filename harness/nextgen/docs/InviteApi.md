# {{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteInvite**](InviteApi.md#DeleteInvite) | **Delete** /ng/api/invites/{inviteId} | Delete an Invite by Identifier
[**GetInvite**](InviteApi.md#GetInvite) | **Get** /ng/api/invites/invite | Gets an Invite by either Invite Id or JwtToken
[**GetInvites**](InviteApi.md#GetInvites) | **Get** /ng/api/invites | List all the Invites for a Project or Organization
[**GetPendingUsersAggregated**](InviteApi.md#GetPendingUsersAggregated) | **Post** /ng/api/invites/aggregate | List of all the Invites pending users
[**SendInvite**](InviteApi.md#SendInvite) | **Post** /ng/api/invites | Send a user Invite to either Project or Organization
[**UpdateInvite**](InviteApi.md#UpdateInvite) | **Put** /ng/api/invites/{inviteId} | Resend the Invite email

# **DeleteInvite**
> ResponseDtoOptionalInvite DeleteInvite(ctx, inviteId)
Delete an Invite by Identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inviteId** | **string**| Invite Id | 

### Return type

[**ResponseDtoOptionalInvite**](ResponseDTOOptionalInvite.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInvite**
> ResponseDtoInvite GetInvite(ctx, optional)
Gets an Invite by either Invite Id or JwtToken

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***InviteApiGetInviteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InviteApiGetInviteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inviteId** | **optional.String**| Invitation Id | 
 **jwttoken** | **optional.String**| JWT Token | 

### Return type

[**ResponseDtoInvite**](ResponseDTOInvite.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, text/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInvites**
> ResponseDtoPageResponseInvite GetInvites(ctx, accountIdentifier, optional)
List all the Invites for a Project or Organization

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
 **optional** | ***InviteApiGetInvitesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InviteApiGetInvitesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity | 
 **pageIndex** | **optional.Int32**| Indicates the number of pages. Results for these pages will be retrieved. | [default to 0]
 **pageSize** | **optional.Int32**| The number of the elements to fetch | [default to 50]
 **sortOrders** | [**optional.Interface of []SortOrder**](SortOrder.md)| Sort criteria for the elements. | 

### Return type

[**ResponseDtoPageResponseInvite**](ResponseDTOPageResponseInvite.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, text/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPendingUsersAggregated**
> ResponseDtoPageResponseInvite GetPendingUsersAggregated(ctx, accountIdentifier, optional)
List of all the Invites pending users

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
 **optional** | ***InviteApiGetPendingUsersAggregatedOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InviteApiGetPendingUsersAggregatedOpts struct
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

[**ResponseDtoPageResponseInvite**](ResponseDTOPageResponseInvite.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, text/yaml
 - **Accept**: application/json, application/yaml, text/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendInvite**
> ResponseDtoListInviteOperationResponse SendInvite(ctx, body, accountIdentifier, optional)
Send a user Invite to either Project or Organization

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateInvite**](CreateInvite.md)| Details of the Invite to create | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
 **optional** | ***InviteApiSendInviteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InviteApiSendInviteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity | 

### Return type

[**ResponseDtoListInviteOperationResponse**](ResponseDTOListInviteOperationResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, text/yaml
 - **Accept**: application/json, application/yaml, text/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateInvite**
> ResponseDtoOptionalInvite UpdateInvite(ctx, body, inviteId, accountIdentifier)
Resend the Invite email

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Invite**](Invite.md)| Details of the Updated Invite | 
  **inviteId** | **string**| Invite id | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 

### Return type

[**ResponseDtoOptionalInvite**](ResponseDTOOptionalInvite.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, text/yaml
 - **Accept**: application/json, application/yaml, text/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

