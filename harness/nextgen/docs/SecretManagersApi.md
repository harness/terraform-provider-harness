# {{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMetadata**](SecretManagersApi.md#GetMetadata) | **Post** /ng/api/secret-managers/meta-data | Gets the metadata of Secret Manager

# **GetMetadata**
> ResponseDtoSecretManagerMetadataDto GetMetadata(ctx, body, accountIdentifier)
Gets the metadata of Secret Manager

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Invite**](Invite.md)| Details required for the creation of the Secret Manager | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 

### Return type

[**ResponseDtoSecretManagerMetadataDto**](ResponseDTOSecretManagerMetadataDTO.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

