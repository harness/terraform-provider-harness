# nextgen{{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GenerateNgHelmValuesYaml**](DelegateSetupResourceApi.md#GenerateNgHelmValuesYaml) | **Post** /ng/api/delegate-setup/generate-helm-values | Generates helm values yaml file from the data specified in request body (Delegate setup details).

# **GenerateNgHelmValuesYaml**
> GenerateNgHelmValuesYaml(ctx, body, accountIdentifier, optional)
Generates helm values yaml file from the data specified in request body (Delegate setup details).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DelegateSetupDetails**](DelegateSetupDetails.md)| Delegate setup details, containing data to populate yaml file values. | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***DelegateSetupResourceApiGenerateNgHelmValuesYamlOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DelegateSetupResourceApiGenerateNgHelmValuesYamlOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity. | 

### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

