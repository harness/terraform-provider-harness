# {{classname}}

All URIs are relative to */api/manager*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RegisterInfraV2**](DefaultApi.md#RegisterInfraV2) | **Post** /rest/v2/infrastructure | Register a new v2 infra

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
 **optional** | ***DefaultApiRegisterInfraV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiRegisterInfraV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **correlationID** | **optional.**| correlation id is used to debug micro svc communication | 

### Return type

[**InfraV2RegisterInfrastructureV2Response**](infra_v2.RegisterInfrastructureV2Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

