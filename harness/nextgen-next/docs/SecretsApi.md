# {{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSecretV2**](SecretsApi.md#DeleteSecretV2) | **Delete** /ng/api/v2/secrets/{identifier} | Deletes Secret by ID and Scope
[**GetSecretV2**](SecretsApi.md#GetSecretV2) | **Get** /ng/api/v2/secrets/{identifier} | Get the Secret by ID and Scope
[**ListSecretsV2**](SecretsApi.md#ListSecretsV2) | **Get** /ng/api/v2/secrets | Fetches the list of Secrets corresponding to the request&#x27;s filter criteria.
[**ListSecretsV3**](SecretsApi.md#ListSecretsV3) | **Post** /ng/api/v2/secrets/list | Fetches the list of Secrets corresponding to the request&#x27;s filter criteria.
[**PostSecret**](SecretsApi.md#PostSecret) | **Post** /ng/api/v2/secrets | Creates a Secret at given Scope
[**PostSecretFileV2**](SecretsApi.md#PostSecretFileV2) | **Post** /ng/api/v2/secrets/files | Creates a Secret File
[**PostSecretViaYaml**](SecretsApi.md#PostSecretViaYaml) | **Post** /ng/api/v2/secrets/yaml | Creates a secret via YAML
[**PutSecret**](SecretsApi.md#PutSecret) | **Put** /ng/api/v2/secrets/{identifier} | Updates the Secret by ID and Scope
[**PutSecretFileV2**](SecretsApi.md#PutSecretFileV2) | **Put** /ng/api/v2/secrets/files/{identifier} | Updates the Secret file by ID and Scope
[**PutSecretViaYaml**](SecretsApi.md#PutSecretViaYaml) | **Put** /ng/api/v2/secrets/{identifier}/yaml | Updates the Secret by ID and Scope via YAML
[**ValidateSecret**](SecretsApi.md#ValidateSecret) | **Post** /ng/api/v2/secrets/validate | Validates Secret with the provided ID and Scope
[**ValidateSecretIdentifierIsUnique**](SecretsApi.md#ValidateSecretIdentifierIsUnique) | **Get** /ng/api/v2/secrets/validateUniqueIdentifier/{identifier} | Checks whether the identifier is unique or not

# **DeleteSecretV2**
> ResponseDtoBoolean DeleteSecretV2(ctx, identifier, accountIdentifier, optional)
Deletes Secret by ID and Scope

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| Secret ID | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
 **optional** | ***SecretsApiDeleteSecretV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecretsApiDeleteSecretV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity | 

### Return type

[**ResponseDtoBoolean**](ResponseDTOBoolean.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSecretV2**
> ResponseDtoSecretResponseWrapper GetSecretV2(ctx, identifier, accountIdentifier, optional)
Get the Secret by ID and Scope

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| Secret ID | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
 **optional** | ***SecretsApiGetSecretV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecretsApiGetSecretV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity | 

### Return type

[**ResponseDtoSecretResponseWrapper**](ResponseDTOSecretResponseWrapper.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSecretsV2**
> ResponseDtoPageResponseSecretResponseWrapper ListSecretsV2(ctx, accountIdentifier, optional)
Fetches the list of Secrets corresponding to the request's filter criteria.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
 **optional** | ***SecretsApiListSecretsV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecretsApiListSecretsV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity | 
 **identifiers** | [**optional.Interface of []string**](string.md)| This is the list of Secret IDs. Details specific to these IDs would be fetched. | 
 **type_** | **optional.String**| Type of Secret whether it is SecretFile, SecretText or SSH key | 
 **searchTerm** | **optional.String**| Filter Secrets based on name, Identifier and tags by this search term | 
 **types** | [**optional.Interface of []string**](string.md)| Add multiple secret types like SecretFile, SecretText or SSH key to criteria | 
 **sourceCategory** | **optional.String**| Source Category like CLOUD_PROVIDER, SECRET_MANAGER, CLOUD_COST, ARTIFACTORY, CODE_REPO, MONITORING or TICKETING | 
 **includeSecretsFromEverySubScope** | **optional.Bool**| Specify whether or not to include secrets from all the sub-scopes of the given Scope | [default to false]
 **pageIndex** | **optional.Int32**| Page number of navigation. The default value is 0 | [default to 0]
 **pageSize** | **optional.Int32**| Number of entries per page. The default value is 100  | [default to 100]

### Return type

[**ResponseDtoPageResponseSecretResponseWrapper**](ResponseDTOPageResponseSecretResponseWrapper.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSecretsV3**
> ResponseDtoPageResponseSecretResponseWrapper ListSecretsV3(ctx, accountIdentifier, optional)
Fetches the list of Secrets corresponding to the request's filter criteria.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
 **optional** | ***SecretsApiListSecretsV3Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecretsApiListSecretsV3Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SecretResourceFilter**](SecretResourceFilter.md)|  | 
 **orgIdentifier** | **optional.**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity | 
 **pageIndex** | **optional.**| Page number of navigation. The default value of 0 | [default to 0]
 **pageSize** | **optional.**| Number of entries per page. The default value is 100 | [default to 100]

### Return type

[**ResponseDtoPageResponseSecretResponseWrapper**](ResponseDTOPageResponseSecretResponseWrapper.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSecret**
> ResponseDtoSecretResponseWrapper PostSecret(ctx, body, accountIdentifier, optional)
Creates a Secret at given Scope

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SecretRequestWrapper**](SecretRequestWrapper.md)| Details required to create the Secret | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
 **optional** | ***SecretsApiPostSecretOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecretsApiPostSecretOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity | 
 **privateSecret** | **optional.**| This is a boolean value to specify if the Secret is Private. The default value is False. | [default to false]

### Return type

[**ResponseDtoSecretResponseWrapper**](ResponseDTOSecretResponseWrapper.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSecretFileV2**
> ResponseDtoSecretResponseWrapper PostSecretFileV2(ctx, accountIdentifier, optional)
Creates a Secret File

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
 **optional** | ***SecretsApiPostSecretFileV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecretsApiPostSecretFileV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | [**optional.Interface of interface{}**](.md)|  | 
 **spec** | **optional.**|  | 
 **orgIdentifier** | **optional.**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity | 
 **privateSecret** | **optional.**| This is a boolean value to specify if the Secret is Private. The default value is False. | [default to false]

### Return type

[**ResponseDtoSecretResponseWrapper**](ResponseDTOSecretResponseWrapper.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSecretViaYaml**
> ResponseDtoSecretResponseWrapper PostSecretViaYaml(ctx, body, accountIdentifier, optional)
Creates a secret via YAML

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SecretRequestWrapper**](SecretRequestWrapper.md)| Details required to create the Secret | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
 **optional** | ***SecretsApiPostSecretViaYamlOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecretsApiPostSecretViaYamlOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity | 
 **privateSecret** | **optional.**| This is a boolean value to specify if the Secret is Private. The default value is False. | [default to false]

### Return type

[**ResponseDtoSecretResponseWrapper**](ResponseDTOSecretResponseWrapper.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutSecret**
> ResponseDtoSecretResponseWrapper PutSecret(ctx, identifier, accountIdentifier, optional)
Updates the Secret by ID and Scope

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| Secret ID | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
 **optional** | ***SecretsApiPutSecretOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecretsApiPutSecretOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of SecretRequestWrapper**](SecretRequestWrapper.md)|  | 
 **orgIdentifier** | **optional.**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity | 

### Return type

[**ResponseDtoSecretResponseWrapper**](ResponseDTOSecretResponseWrapper.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutSecretFileV2**
> ResponseDtoSecretResponseWrapper PutSecretFileV2(ctx, accountIdentifier, identifier, optional)
Updates the Secret file by ID and Scope

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
  **identifier** | **string**| Secret ID | 
 **optional** | ***SecretsApiPutSecretFileV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecretsApiPutSecretFileV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **file** | [**optional.Interface of interface{}**](.md)|  | 
 **spec** | **optional.**|  | 
 **orgIdentifier** | **optional.**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity | 

### Return type

[**ResponseDtoSecretResponseWrapper**](ResponseDTOSecretResponseWrapper.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutSecretViaYaml**
> ResponseDtoSecretResponseWrapper PutSecretViaYaml(ctx, body, identifier, accountIdentifier, optional)
Updates the Secret by ID and Scope via YAML

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SecretRequestWrapper**](SecretRequestWrapper.md)| Details of Secret to create | 
  **identifier** | **string**| Secret ID | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
 **optional** | ***SecretsApiPutSecretViaYamlOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecretsApiPutSecretViaYamlOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orgIdentifier** | **optional.**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity | 

### Return type

[**ResponseDtoSecretResponseWrapper**](ResponseDTOSecretResponseWrapper.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateSecret**
> ResponseDtoSecretValidationResultDto ValidateSecret(ctx, body, accountIdentifier, optional)
Validates Secret with the provided ID and Scope

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SecretValidationMetaData**](SecretValidationMetaData.md)| Details of the Secret type | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
 **optional** | ***SecretsApiValidateSecretOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecretsApiValidateSecretOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity | 
 **identifier** | **optional.**| Secret ID | 

### Return type

[**ResponseDtoSecretValidationResultDto**](ResponseDTOSecretValidationResultDTO.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateSecretIdentifierIsUnique**
> ResponseDtoBoolean ValidateSecretIdentifierIsUnique(ctx, identifier, accountIdentifier, optional)
Checks whether the identifier is unique or not

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| Secret Identifier | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
 **optional** | ***SecretsApiValidateSecretIdentifierIsUniqueOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecretsApiValidateSecretIdentifierIsUniqueOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity | 

### Return type

[**ResponseDtoBoolean**](ResponseDTOBoolean.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

