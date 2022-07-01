# nextgen{{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ValidateHosts**](ValidateHostApi.md#ValidateHosts) | **Post** /ng/api/host-validation | Validates hosts connectivity credentials

# **ValidateHosts**
> ResponseDtoListHostValidationDto ValidateHosts(ctx, body, accountIdentifier, identifier, optional)
Validates hosts connectivity credentials

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**HostValidationParams**](HostValidationParams.md)| List of SSH or WinRm hosts to validate, and Delegate tags (optional) | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **identifier** | **string**| Secret Identifier | 
 **optional** | ***ValidateHostApiValidateHostsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ValidateHostApiValidateHostsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orgIdentifier** | **optional.**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity. | 

### Return type

[**ResponseDtoListHostValidationDto**](ResponseDTOListHostValidationDTO.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

