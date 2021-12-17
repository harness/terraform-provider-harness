# {{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSamlMetaData**](AuthenticationSettingsApi.md#DeleteSamlMetaData) | **Delete** /ng/api/authentication-settings/delete-saml-metadata | Deletes Saml meta data by accountIdentifier
[**GetAuthenticationSettings**](AuthenticationSettingsApi.md#GetAuthenticationSettings) | **Get** /ng/api/authentication-settings | Get the authentication settings by accountIdentifier
[**GetPasswordStrengthSettings**](AuthenticationSettingsApi.md#GetPasswordStrengthSettings) | **Get** /ng/api/authentication-settings/login-settings/password-strength | Get the password strength settings by accountIdentifier
[**GetSamlLoginTest**](AuthenticationSettingsApi.md#GetSamlLoginTest) | **Get** /ng/api/authentication-settings/saml-login-test | Get the Saml login test by accountId
[**PutLoginSettings**](AuthenticationSettingsApi.md#PutLoginSettings) | **Put** /ng/api/authentication-settings/login-settings/{loginSettingsId} | Updates the login settings
[**RemoveOauthMechanism**](AuthenticationSettingsApi.md#RemoveOauthMechanism) | **Delete** /ng/api/authentication-settings/oauth/remove-mechanism | Deletes Oauth mechanism by accountIdentifier
[**SetTwoFactorAuthAtAccountLevel**](AuthenticationSettingsApi.md#SetTwoFactorAuthAtAccountLevel) | **Put** /ng/api/authentication-settings/two-factor-admin-override-settings | set two factor auth at account lever by accountIdentifier
[**UpdateAuthMechanism**](AuthenticationSettingsApi.md#UpdateAuthMechanism) | **Put** /ng/api/authentication-settings/update-auth-mechanism | Updates the Auth mechanism by accountIdentifier
[**UpdateOauthProviders**](AuthenticationSettingsApi.md#UpdateOauthProviders) | **Put** /ng/api/authentication-settings/oauth/update-providers | Updates the Oauth providers by accountIdentifier
[**UpdateSamlMetaData**](AuthenticationSettingsApi.md#UpdateSamlMetaData) | **Put** /ng/api/authentication-settings/saml-metadata-upload | Updates the saml metadata by accountId
[**UpdateWhitelistedDomains**](AuthenticationSettingsApi.md#UpdateWhitelistedDomains) | **Put** /ng/api/authentication-settings/whitelisted-domains | Updates the Whitelisted domains by accountIdentifier
[**UploadSamlMetaData**](AuthenticationSettingsApi.md#UploadSamlMetaData) | **Post** /ng/api/authentication-settings/saml-metadata-upload | Uploads the saml metadata by accountId

# **DeleteSamlMetaData**
> RestResponseSsoConfig DeleteSamlMetaData(ctx, optional)
Deletes Saml meta data by accountIdentifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AuthenticationSettingsApiDeleteSamlMetaDataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthenticationSettingsApiDeleteSamlMetaDataOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountIdentifier** | **optional.String**| Account Identifier for the entity | 

### Return type

[**RestResponseSsoConfig**](RestResponseSSOConfig.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAuthenticationSettings**
> RestResponseAuthenticationSettingsResponse GetAuthenticationSettings(ctx, optional)
Get the authentication settings by accountIdentifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AuthenticationSettingsApiGetAuthenticationSettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthenticationSettingsApiGetAuthenticationSettingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountIdentifier** | **optional.String**| Account Identifier for the entity | 

### Return type

[**RestResponseAuthenticationSettingsResponse**](RestResponseAuthenticationSettingsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPasswordStrengthSettings**
> RestResponsePasswordStrengthPolicy GetPasswordStrengthSettings(ctx, optional)
Get the password strength settings by accountIdentifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AuthenticationSettingsApiGetPasswordStrengthSettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthenticationSettingsApiGetPasswordStrengthSettingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountIdentifier** | **optional.String**| Account Identifier for the entity | 

### Return type

[**RestResponsePasswordStrengthPolicy**](RestResponsePasswordStrengthPolicy.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSamlLoginTest**
> RestResponseLoginTypeResponse GetSamlLoginTest(ctx, optional)
Get the Saml login test by accountId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AuthenticationSettingsApiGetSamlLoginTestOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthenticationSettingsApiGetSamlLoginTestOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **optional.String**| Account Identifier for the entity | 

### Return type

[**RestResponseLoginTypeResponse**](RestResponseLoginTypeResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutLoginSettings**
> RestResponseLoginSettings PutLoginSettings(ctx, body, loginSettingsId, optional)
Updates the login settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LoginSettings**](LoginSettings.md)| This is the updated Login Settings. Please provide values for all fields, not just the fields you are updating | 
  **loginSettingsId** | **string**| Login Settings Identifier | 
 **optional** | ***AuthenticationSettingsApiPutLoginSettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthenticationSettingsApiPutLoginSettingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accountIdentifier** | **optional.**| Account Identifier for the entity | 

### Return type

[**RestResponseLoginSettings**](RestResponseLoginSettings.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveOauthMechanism**
> RestResponseBoolean RemoveOauthMechanism(ctx, optional)
Deletes Oauth mechanism by accountIdentifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AuthenticationSettingsApiRemoveOauthMechanismOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthenticationSettingsApiRemoveOauthMechanismOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountIdentifier** | **optional.String**| Account Identifier for the entity | 

### Return type

[**RestResponseBoolean**](RestResponseBoolean.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetTwoFactorAuthAtAccountLevel**
> RestResponseBoolean SetTwoFactorAuthAtAccountLevel(ctx, optional)
set two factor auth at account lever by accountIdentifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AuthenticationSettingsApiSetTwoFactorAuthAtAccountLevelOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthenticationSettingsApiSetTwoFactorAuthAtAccountLevelOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of TwoFactorAdminOverrideSettings**](TwoFactorAdminOverrideSettings.md)|  | 
 **accountIdentifier** | **optional.**| Account Identifier for the entity | 

### Return type

[**RestResponseBoolean**](RestResponseBoolean.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAuthMechanism**
> RestResponseBoolean UpdateAuthMechanism(ctx, optional)
Updates the Auth mechanism by accountIdentifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AuthenticationSettingsApiUpdateAuthMechanismOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthenticationSettingsApiUpdateAuthMechanismOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountIdentifier** | **optional.String**| Account Identifier for the entity | 
 **authenticationMechanism** | **optional.String**| Type of Authentication Mechanism SSO or NON_SSO | 

### Return type

[**RestResponseBoolean**](RestResponseBoolean.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOauthProviders**
> RestResponseBoolean UpdateOauthProviders(ctx, optional)
Updates the Oauth providers by accountIdentifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AuthenticationSettingsApiUpdateOauthProvidersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthenticationSettingsApiUpdateOauthProvidersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of OAuthSettings**](OAuthSettings.md)|  | 
 **accountIdentifier** | **optional.**| Account Identifier for the entity | 

### Return type

[**RestResponseBoolean**](RestResponseBoolean.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSamlMetaData**
> RestResponseSsoConfig UpdateSamlMetaData(ctx, optional)
Updates the saml metadata by accountId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AuthenticationSettingsApiUpdateSamlMetaDataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthenticationSettingsApiUpdateSamlMetaDataOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | [**optional.Interface of FormDataContentDisposition**](.md)|  | 
 **displayName** | **optional.**|  | 
 **groupMembershipAttr** | **optional.**|  | 
 **authorizationEnabled** | **optional.**|  | 
 **logoutUrl** | **optional.**|  | 
 **entityIdentifier** | **optional.**|  | 
 **accountId** | **optional.**| Account Identifier for the entity | 

### Return type

[**RestResponseSsoConfig**](RestResponseSSOConfig.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateWhitelistedDomains**
> RestResponseBoolean UpdateWhitelistedDomains(ctx, optional)
Updates the Whitelisted domains by accountIdentifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AuthenticationSettingsApiUpdateWhitelistedDomainsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthenticationSettingsApiUpdateWhitelistedDomainsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of []string**](string.md)| Set of whitelisted domains and IPs for the account | 
 **accountIdentifier** | **optional.**| Account Identifier for the entity | 

### Return type

[**RestResponseBoolean**](RestResponseBoolean.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadSamlMetaData**
> RestResponseSsoConfig UploadSamlMetaData(ctx, optional)
Uploads the saml metadata by accountId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AuthenticationSettingsApiUploadSamlMetaDataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthenticationSettingsApiUploadSamlMetaDataOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | [**optional.Interface of FormDataContentDisposition**](.md)|  | 
 **displayName** | **optional.**|  | 
 **groupMembershipAttr** | **optional.**|  | 
 **authorizationEnabled** | **optional.**|  | 
 **logoutUrl** | **optional.**|  | 
 **entityIdentifier** | **optional.**|  | 
 **accountId** | **optional.**| Account Identifier for the entity | 

### Return type

[**RestResponseSsoConfig**](RestResponseSSOConfig.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

