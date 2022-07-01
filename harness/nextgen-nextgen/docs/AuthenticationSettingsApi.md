# nextgen{{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSamlMetaData**](AuthenticationSettingsApi.md#DeleteSamlMetaData) | **Delete** /ng/api/authentication-settings/delete-saml-metadata | Delete SAML meta data
[**GetAuthenticationSettings**](AuthenticationSettingsApi.md#GetAuthenticationSettings) | **Get** /ng/api/authentication-settings | Gets authentication settings for the given Account ID
[**GetPasswordStrengthSettings**](AuthenticationSettingsApi.md#GetPasswordStrengthSettings) | **Get** /ng/api/authentication-settings/login-settings/password-strength | Get password strength
[**GetSamlLoginTest**](AuthenticationSettingsApi.md#GetSamlLoginTest) | **Get** /ng/api/authentication-settings/saml-login-test | Test SAML connectivity
[**RemoveOauthMechanism**](AuthenticationSettingsApi.md#RemoveOauthMechanism) | **Delete** /ng/api/authentication-settings/oauth/remove-mechanism | Delete OAuth Setting
[**SetTwoFactorAuthAtAccountLevel**](AuthenticationSettingsApi.md#SetTwoFactorAuthAtAccountLevel) | **Put** /ng/api/authentication-settings/two-factor-admin-override-settings | Set two factor authorization
[**UpdateAuthMechanism**](AuthenticationSettingsApi.md#UpdateAuthMechanism) | **Put** /ng/api/authentication-settings/update-auth-mechanism | Update Auth mechanism
[**UpdateOauthProviders**](AuthenticationSettingsApi.md#UpdateOauthProviders) | **Put** /ng/api/authentication-settings/oauth/update-providers | Update Oauth providers
[**UpdateSamlMetaData**](AuthenticationSettingsApi.md#UpdateSamlMetaData) | **Put** /ng/api/authentication-settings/saml-metadata-upload | Update SAML metadata
[**UpdateWhitelistedDomains**](AuthenticationSettingsApi.md#UpdateWhitelistedDomains) | **Put** /ng/api/authentication-settings/whitelisted-domains | Updates the whitelisted domains
[**UploadSamlMetaData**](AuthenticationSettingsApi.md#UploadSamlMetaData) | **Post** /ng/api/authentication-settings/saml-metadata-upload | Upload SAML metadata

# **DeleteSamlMetaData**
> RestResponseSsoConfig DeleteSamlMetaData(ctx, accountIdentifier)
Delete SAML meta data

Deletes SAML metadata for the given Account ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 

### Return type

[**RestResponseSsoConfig**](RestResponseSSOConfig.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAuthenticationSettings**
> RestResponseAuthenticationSettingsResponse GetAuthenticationSettings(ctx, accountIdentifier)
Gets authentication settings for the given Account ID

Gets authentication settings for the given Account ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 

### Return type

[**RestResponseAuthenticationSettingsResponse**](RestResponseAuthenticationSettingsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPasswordStrengthSettings**
> RestResponsePasswordStrengthPolicy GetPasswordStrengthSettings(ctx, accountIdentifier)
Get password strength

Gets password strength for the given Account ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 

### Return type

[**RestResponsePasswordStrengthPolicy**](RestResponsePasswordStrengthPolicy.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSamlLoginTest**
> RestResponseLoginTypeResponse GetSamlLoginTest(ctx, accountId)
Test SAML connectivity

Tests SAML connectivity for the given Account ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Account Identifier for the Entity. | 

### Return type

[**RestResponseLoginTypeResponse**](RestResponseLoginTypeResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveOauthMechanism**
> RestResponseBoolean RemoveOauthMechanism(ctx, accountIdentifier)
Delete OAuth Setting

Deletes OAuth settings for a given Account ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 

### Return type

[**RestResponseBoolean**](RestResponseBoolean.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetTwoFactorAuthAtAccountLevel**
> RestResponseBoolean SetTwoFactorAuthAtAccountLevel(ctx, body, accountIdentifier)
Set two factor authorization

Sets Two-Factor authorization for the given Account ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TwoFactorAdminOverrideSettings**](TwoFactorAdminOverrideSettings.md)| Boolean that specify whether or not to override two factor enabled setting | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 

### Return type

[**RestResponseBoolean**](RestResponseBoolean.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAuthMechanism**
> RestResponseBoolean UpdateAuthMechanism(ctx, accountIdentifier, optional)
Update Auth mechanism

Updates the authentication mechanism for the given Account ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***AuthenticationSettingsApiUpdateAuthMechanismOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthenticationSettingsApiUpdateAuthMechanismOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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
> RestResponseBoolean UpdateOauthProviders(ctx, body, accountIdentifier)
Update Oauth providers

Updates OAuth providers for the given Account ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**OAuthSettings**](OAuthSettings.md)| This is the updated OAuthSettings. Please provide values for all fields, not just the fields you are updating | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 

### Return type

[**RestResponseBoolean**](RestResponseBoolean.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSamlMetaData**
> RestResponseSsoConfig UpdateSamlMetaData(ctx, accountId, optional)
Update SAML metadata

Updates SAML metadata of the SAML configuration configured for an account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Account Identifier for the Entity. | 
 **optional** | ***AuthenticationSettingsApiUpdateSamlMetaDataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthenticationSettingsApiUpdateSamlMetaDataOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | [**optional.Interface of interface{}**](.md)|  | 
 **fileMetadata** | [**optional.Interface of FormDataContentDisposition**](.md)|  | 
 **displayName** | **optional.**|  | 
 **groupMembershipAttr** | **optional.**|  | 
 **authorizationEnabled** | **optional.**|  | 
 **logoutUrl** | **optional.**|  | 
 **entityIdentifier** | **optional.**|  | 
 **samlProviderType** | **optional.**|  | 
 **clientId** | **optional.**|  | 
 **clientSecret** | **optional.**|  | 

### Return type

[**RestResponseSsoConfig**](RestResponseSSOConfig.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateWhitelistedDomains**
> RestResponseBoolean UpdateWhitelistedDomains(ctx, accountIdentifier, optional)
Updates the whitelisted domains

Updates whitelisted domains configured for an account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***AuthenticationSettingsApiUpdateWhitelistedDomainsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthenticationSettingsApiUpdateWhitelistedDomainsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of []string**](string.md)| Set of whitelisted domains and IPs for the account | 

### Return type

[**RestResponseBoolean**](RestResponseBoolean.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadSamlMetaData**
> RestResponseSsoConfig UploadSamlMetaData(ctx, accountId, optional)
Upload SAML metadata

Updates the SAML metadata for the given Account ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Account Identifier for the Entity. | 
 **optional** | ***AuthenticationSettingsApiUploadSamlMetaDataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthenticationSettingsApiUploadSamlMetaDataOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | [**optional.Interface of interface{}**](.md)|  | 
 **fileMetadata** | [**optional.Interface of FormDataContentDisposition**](.md)|  | 
 **displayName** | **optional.**|  | 
 **groupMembershipAttr** | **optional.**|  | 
 **authorizationEnabled** | **optional.**|  | 
 **logoutUrl** | **optional.**|  | 
 **entityIdentifier** | **optional.**|  | 
 **samlProviderType** | **optional.**|  | 
 **clientId** | **optional.**|  | 
 **clientSecret** | **optional.**|  | 

### Return type

[**RestResponseSsoConfig**](RestResponseSSOConfig.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

