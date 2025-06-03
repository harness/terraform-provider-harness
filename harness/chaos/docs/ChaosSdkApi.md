# {{classname}}

All URIs are relative to */api/manager*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteInfraV2**](ChaosSdkApi.md#DeleteInfraV2) | **Delete** /rest/v2/infrastructure/{environmentIdentifier}/{identity} | Delete a v2 infra
[**GetInfraV2**](ChaosSdkApi.md#GetInfraV2) | **Get** /rest/v2/infrastructure/{identity} | Get a new v2 infra
[**ListInfraV2**](ChaosSdkApi.md#ListInfraV2) | **Post** /rest/v2/infrastructures | List a new v2 infra
[**RegisterInfraV2**](ChaosSdkApi.md#RegisterInfraV2) | **Post** /rest/v2/infrastructure | Register a new v2 infra
[**UpdateInfraV2**](ChaosSdkApi.md#UpdateInfraV2) | **Put** /rest/v2/infrastructure | Update a new v2 infra

# **DeleteInfraV2**
> InfraV2DeleteKubernetesInfraV2Response DeleteInfraV2(ctx, identity, environmentIdentifier, accountIdentifier, organizationIdentifier, projectIdentifier)
Delete a v2 infra

Delete a new v2 infra

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identity** | **string**| Chaos V2 Infra Identity | 
  **environmentIdentifier** | **string**| Chaos V2 Environment Identity | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 

### Return type

[**InfraV2DeleteKubernetesInfraV2Response**](infra_v2.DeleteKubernetesInfraV2Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInfraV2**
> InfraV2KubernetesInfrastructureV2Details GetInfraV2(ctx, identity, accountIdentifier, organizationIdentifier, projectIdentifier, environmentIdentifier)
Get a new v2 infra

Get a new v2 infra

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identity** | **string**| Chaos V2 Infra Identity | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **environmentIdentifier** | **string**| environment identifier to filter resource | 

### Return type

[**InfraV2KubernetesInfrastructureV2Details**](infra_v2.KubernetesInfrastructureV2Details.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListInfraV2**
> InfraV2ListKubernetesInfraV2Response ListInfraV2(ctx, body, accountIdentifier, organizationIdentifier, projectIdentifier, page, limit, optional)
List a new v2 infra

List a new v2 infra

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**InfraV2ListKubernetesInfraV2Request**](InfraV2ListKubernetesInfraV2Request.md)| list Infra V2 | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **page** | **int32**| page number | [default to 0]
  **limit** | **int32**| limit per page | [default to 10]
 **optional** | ***ChaosSdkApiListInfraV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChaosSdkApiListInfraV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **includeLegacyInfra** | **optional.**| include legacy infra details | 
 **environmentIdentifier** | **optional.**| filter infra | 
 **search** | **optional.**| search based on name | 

### Return type

[**InfraV2ListKubernetesInfraV2Response**](infra_v2.ListKubernetesInfraV2Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegisterInfraV2**
> InfraV2RegisterInfrastructureV2Response RegisterInfraV2(ctx, body, accountIdentifier, organizationIdentifier, projectIdentifier, optional)
Register a new v2 infra

Register a new v2 infra

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**InfraV2RegisterInfrastructureV2Request**](InfraV2RegisterInfrastructureV2Request.md)| Register Infra V2 | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
 **optional** | ***ChaosSdkApiRegisterInfraV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChaosSdkApiRegisterInfraV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **correlationID** | **optional.**| correlation id is used to debug micro svc communication | 

### Return type

[**InfraV2RegisterInfrastructureV2Response**](infra_v2.RegisterInfrastructureV2Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateInfraV2**
> InfraV2UpdateKubernetesInfrastructureV2Response UpdateInfraV2(ctx, body, accountIdentifier, organizationIdentifier, projectIdentifier, optional)
Update a new v2 infra

Update a new v2 infra

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**InfraV2UpdateKubernetesInfrastructureV2Request**](InfraV2UpdateKubernetesInfrastructureV2Request.md)| update Infra V2 | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
 **optional** | ***ChaosSdkApiUpdateInfraV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChaosSdkApiUpdateInfraV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **correlationID** | **optional.**| correlation id is used to debug micro svc communication | 

### Return type

[**InfraV2UpdateKubernetesInfrastructureV2Response**](infra_v2.UpdateKubernetesInfrastructureV2Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

