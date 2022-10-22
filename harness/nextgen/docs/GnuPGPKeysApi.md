# nextgen{{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AgentGPGKeyServiceCreate**](GnuPGPKeysApi.md#AgentGPGKeyServiceCreate) | **Post** /gitops/api/api/v1/agents/{agentIdentifier}/gpgkeys | Create one or more GPG public keys in the server&#x27;s configuration
[**AgentGPGKeyServiceDelete**](GnuPGPKeysApi.md#AgentGPGKeyServiceDelete) | **Delete** /gitops/api/api/v1/agents/{agentIdentifier}/gpgkeys/{query.keyID} | Delete specified GPG public key from the server&#x27;s configuration
[**AgentGPGKeyServiceGet**](GnuPGPKeysApi.md#AgentGPGKeyServiceGet) | **Get** /gitops/api/api/v1/agents/{agentIdentifier}/gpgkeys/{query.keyID} | Get information about specified GPG public key from the server

# **AgentGPGKeyServiceCreate**
> GpgkeysGnuPgPublicKeyCreateResponse AgentGPGKeyServiceCreate(ctx, body, agentIdentifier, optional)
Create one or more GPG public keys in the server's configuration

Create one or more GPG public keys in the server's configuration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GpgkeysGnuPgPublicKeyCreateRequest**](GpgkeysGnuPgPublicKeyCreateRequest.md)|  | 
  **agentIdentifier** | **string**| Agent identifier for entity. | 
 **optional** | ***GnuPGPKeysApiAgentGPGKeyServiceCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GnuPGPKeysApiAgentGPGKeyServiceCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accountIdentifier** | **optional.**| Account Identifier for the Entity. | 
 **orgIdentifier** | **optional.**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity. | 

### Return type

[**GpgkeysGnuPgPublicKeyCreateResponse**](gpgkeysGnuPGPublicKeyCreateResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentGPGKeyServiceDelete**
> GpgkeysGnuPgPublicKeyResponse AgentGPGKeyServiceDelete(ctx, agentIdentifier, queryKeyID, optional)
Delete specified GPG public key from the server's configuration

Delete specified GPG public key from the server's configuration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentifier** | **string**| Agent identifier for entity. | 
  **queryKeyID** | **string**| The GPG key ID to query for | 
 **optional** | ***GnuPGPKeysApiAgentGPGKeyServiceDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GnuPGPKeysApiAgentGPGKeyServiceDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accountIdentifier** | **optional.String**| Account Identifier for the Entity. | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 

### Return type

[**GpgkeysGnuPgPublicKeyResponse**](gpgkeysGnuPGPublicKeyResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentGPGKeyServiceGet**
> GpgkeysGnuPgPublicKey AgentGPGKeyServiceGet(ctx, agentIdentifier, queryKeyID, accountIdentifier, optional)
Get information about specified GPG public key from the server

Get information about specified GPG public key from the server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentifier** | **string**| Agent identifier for entity. | 
  **queryKeyID** | **string**| The GPG key ID to query for | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***GnuPGPKeysApiAgentGPGKeyServiceGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GnuPGPKeysApiAgentGPGKeyServiceGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 

### Return type

[**GpgkeysGnuPgPublicKey**](gpgkeysGnuPGPublicKey.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

