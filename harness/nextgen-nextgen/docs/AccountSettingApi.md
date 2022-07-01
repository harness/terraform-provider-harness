# nextgen{{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccountSetting**](AccountSettingApi.md#GetAccountSetting) | **Get** /ng/api/account-setting | Get the AccountSetting by accountIdentifier
[**ListAccountSetting**](AccountSettingApi.md#ListAccountSetting) | **Get** /ng/api/account-setting/list | Get the AccountSetting by accountIdentifier
[**UpdateAccountSetting**](AccountSettingApi.md#UpdateAccountSetting) | **Put** /ng/api/account-setting | Updates account settings

# **GetAccountSetting**
> ResponseDtoAccountSettingResponse GetAccountSetting(ctx, accountIdentifier, type_, optional)
Get the AccountSetting by accountIdentifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **type_** | **string**|  | 
 **optional** | ***AccountSettingApiGetAccountSettingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountSettingApiGetAccountSettingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 

### Return type

[**ResponseDtoAccountSettingResponse**](ResponseDTOAccountSettingResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAccountSetting**
> ResponseDtoListAccountSettings ListAccountSetting(ctx, accountIdentifier, optional)
Get the AccountSetting by accountIdentifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***AccountSettingApiListAccountSettingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountSettingApiListAccountSettingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **type_** | **optional.String**|  | 

### Return type

[**ResponseDtoListAccountSettings**](ResponseDTOListAccountSettings.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAccountSetting**
> ResponseDtoAccountSettingResponse UpdateAccountSetting(ctx, body, accountIdentifier)
Updates account settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AccountSettings**](AccountSettings.md)| Details of the AccountSetting to create | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 

### Return type

[**ResponseDtoAccountSettingResponse**](ResponseDTOAccountSettingResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, text/yaml, text/html, text/plain
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

