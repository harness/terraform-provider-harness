# {{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSmtpConfig**](SmtpConfigApi.md#CreateSmtpConfig) | **Post** /ng/api/smtpConfig | Creates SMTP config
[**DeleteSmtpConfig**](SmtpConfigApi.md#DeleteSmtpConfig) | **Delete** /ng/api/smtpConfig/{identifier} | Delete Smtp Config by identifier
[**GetSmtpConfig**](SmtpConfigApi.md#GetSmtpConfig) | **Get** /ng/api/smtpConfig | Gets Smtp config by accountId
[**UpdateSmtp**](SmtpConfigApi.md#UpdateSmtp) | **Put** /ng/api/smtpConfig | Updates the Smtp Config
[**ValidateConnectivity**](SmtpConfigApi.md#ValidateConnectivity) | **Post** /ng/api/smtpConfig/validate-connectivity | Tests the config&#x27;s connectivity by sending a test email
[**ValidateName**](SmtpConfigApi.md#ValidateName) | **Post** /ng/api/smtpConfig/validateName | Checks whether other connectors exist with the same name

# **CreateSmtpConfig**
> ResponseDtoNgSmtp CreateSmtpConfig(ctx, optional)
Creates SMTP config

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SmtpConfigApiCreateSmtpConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SmtpConfigApiCreateSmtpConfigOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of NgSmtp**](NgSmtp.md)|  | 

### Return type

[**ResponseDtoNgSmtp**](ResponseDTONgSmtp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSmtpConfig**
> ResponseDtoBoolean DeleteSmtpConfig(ctx, identifier)
Delete Smtp Config by identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| Config identifier | 

### Return type

[**ResponseDtoBoolean**](ResponseDTOBoolean.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSmtpConfig**
> ResponseDtoNgSmtp GetSmtpConfig(ctx, optional)
Gets Smtp config by accountId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SmtpConfigApiGetSmtpConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SmtpConfigApiGetSmtpConfigOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **optional.String**| Account Identifier for the Entity | 

### Return type

[**ResponseDtoNgSmtp**](ResponseDTONgSmtp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSmtp**
> ResponseDtoNgSmtp UpdateSmtp(ctx, optional)
Updates the Smtp Config

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SmtpConfigApiUpdateSmtpOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SmtpConfigApiUpdateSmtpOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of NgSmtp**](NgSmtp.md)|  | 

### Return type

[**ResponseDtoNgSmtp**](ResponseDTONgSmtp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateConnectivity**
> ResponseDtoValidationResult ValidateConnectivity(ctx, identifier, accountId, to, subject, body)
Tests the config's connectivity by sending a test email

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| Attribute uuid | 
  **accountId** | **string**| Account Identifier for the Entity | 
  **to** | **string**|  | 
  **subject** | **string**|  | 
  **body** | **string**|  | 

### Return type

[**ResponseDtoValidationResult**](ResponseDTOValidationResult.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateName**
> ResponseDtoValidationResult ValidateName(ctx, accountId, optional)
Checks whether other connectors exist with the same name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Account Identifier for the Entity | 
 **optional** | ***SmtpConfigApiValidateNameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SmtpConfigApiValidateNameOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| The name of Config | 

### Return type

[**ResponseDtoValidationResult**](ResponseDTOValidationResult.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

