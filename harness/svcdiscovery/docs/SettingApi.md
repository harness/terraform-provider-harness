# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSetting**](SettingApi.md#GetSetting) | **Get** /api/v1/settings | Gat setting
[**ResetImageRegistrySetting**](SettingApi.md#ResetImageRegistrySetting) | **Post** /api/v1/resetimageregistrysettings | Reset image registry setting
[**SaveSetting**](SettingApi.md#SaveSetting) | **Post** /api/v1/settings | Save setting

# **GetSetting**
> ServiceGetSettingResponse GetSetting(ctx, accountIdentifier, optional)
Gat setting

Get setting

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| account id is the account where you want to create the resource | 
 **optional** | ***SettingApiGetSettingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SettingApiGetSettingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **correlationID** | **optional.String**| correlation id is used to debug micro svc communication | 
 **organizationIdentifier** | **optional.String**| organization id is the organization where you want to create the resource | 
 **projectIdentifier** | **optional.String**| project id is the project where you want to create the resource | 

### Return type

[**ServiceGetSettingResponse**](service.GetSettingResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResetImageRegistrySetting**
> ServiceGetSettingResponse ResetImageRegistrySetting(ctx, body, accountIdentifier, optional)
Reset image registry setting

Reset image registry setting

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServiceEmpty**](ServiceEmpty.md)| Reset image registry setting | 
  **accountIdentifier** | **string**| account id is the account where you want to create the resource | 
 **optional** | ***SettingApiResetImageRegistrySettingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SettingApiResetImageRegistrySettingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **correlationID** | **optional.**| correlation id is used to debug micro svc communication | 
 **organizationIdentifier** | **optional.**| organization id is the organization where you want to create the resource | 
 **projectIdentifier** | **optional.**| project id is the project where you want to create the resource | 

### Return type

[**ServiceGetSettingResponse**](service.GetSettingResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SaveSetting**
> ServiceGetSettingResponse SaveSetting(ctx, body, accountIdentifier, optional)
Save setting

Save setting

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServiceSaveSettingRequest**](ServiceSaveSettingRequest.md)| Save Setting | 
  **accountIdentifier** | **string**| account id is the account where you want to create the resource | 
 **optional** | ***SettingApiSaveSettingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SettingApiSaveSettingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **correlationID** | **optional.**| correlation id is used to debug micro svc communication | 
 **organizationIdentifier** | **optional.**| organization id is the organization where you want to create the resource | 
 **projectIdentifier** | **optional.**| project id is the project where you want to create the resource | 

### Return type

[**ServiceGetSettingResponse**](service.GetSettingResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

