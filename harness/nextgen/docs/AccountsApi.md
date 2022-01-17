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
  **accountIdentifier** | **string**| Account Identifier for the Entity | 

### Return type

[**ResponseDtoAccount**](ResponseDTOAccount.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAccountDefaultExperienceNG**
> ResponseDtoAccount UpdateAccountDefaultExperienceNG(ctx, body, accountIdentifier)
Update Default Experience

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Account**](Account.md)| This is details of the Account. DefaultExperience is mandatory | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 

### Return type

[**ResponseDtoAccount**](ResponseDTOAccount.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAccountNameNG**
> ResponseDtoAccount UpdateAccountNameNG(ctx, body, accountIdentifier)
Update Account Name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Account**](Account.md)| This is details of the Account. Name is mandatory. | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 

### Return type

[**ResponseDtoAccount**](ResponseDTOAccount.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

