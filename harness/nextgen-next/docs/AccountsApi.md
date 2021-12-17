# {{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccountNG**](AccountsApi.md#GetAccountNG) | **Get** /ng/api/accounts/{accountIdentifier} | Gets an account
[**UpdateAccountDefaultExperienceNG**](AccountsApi.md#UpdateAccountDefaultExperienceNG) | **Put** /ng/api/accounts/{accountIdentifier}/default-experience | Update Default Experience
[**UpdateAccountNameNG**](AccountsApi.md#UpdateAccountNameNG) | **Put** /ng/api/accounts/{accountIdentifier}/name | Update Account Name

# **GetAccountNG**
> ResponseDtoAccount GetAccountNG(ctx, accountIdentifier)
Gets an account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account id to get an account. | 

### Return type

[**ResponseDtoAccount**](ResponseDTOAccount.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAccountDefaultExperienceNG**
> ResponseDtoAccount UpdateAccountDefaultExperienceNG(ctx, accountIdentifier, optional)
Update Default Experience

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account id to update the default experience. | 
 **optional** | ***AccountsApiUpdateAccountDefaultExperienceNGOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountsApiUpdateAccountDefaultExperienceNGOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Account**](Account.md)|  | 

### Return type

[**ResponseDtoAccount**](ResponseDTOAccount.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAccountNameNG**
> ResponseDtoAccount UpdateAccountNameNG(ctx, accountIdentifier, optional)
Update Account Name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account id to update an account name. | 
 **optional** | ***AccountsApiUpdateAccountNameNGOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountsApiUpdateAccountNameNGOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Account**](Account.md)|  | 

### Return type

[**ResponseDtoAccount**](ResponseDTOAccount.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

