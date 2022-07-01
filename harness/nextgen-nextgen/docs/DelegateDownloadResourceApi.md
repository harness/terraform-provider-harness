# nextgen{{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadDockerDelegateYaml**](DelegateDownloadResourceApi.md#DownloadDockerDelegateYaml) | **Post** /ng/api/download-delegates/docker | Downloads a docker delegate yaml file.
[**DownloadKubernetesDelegateYaml**](DelegateDownloadResourceApi.md#DownloadKubernetesDelegateYaml) | **Post** /ng/api/download-delegates/kubernetes | Downloads a kubernetes delegate yaml file.

# **DownloadDockerDelegateYaml**
> DownloadDockerDelegateYaml(ctx, body, optional)
Downloads a docker delegate yaml file.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DelegateDownloadRequest**](DelegateDownloadRequest.md)| Parameters needed for downloading docker delegate yaml | 
 **optional** | ***DelegateDownloadResourceApiDownloadDockerDelegateYamlOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DelegateDownloadResourceApiDownloadDockerDelegateYamlOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.**| Account Identifier for the Entity. | 
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

# **DownloadKubernetesDelegateYaml**
> DownloadKubernetesDelegateYaml(ctx, body, optional)
Downloads a kubernetes delegate yaml file.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DelegateDownloadRequest**](DelegateDownloadRequest.md)| Parameters needed for downloading kubernetes delegate yaml | 
 **optional** | ***DelegateDownloadResourceApiDownloadKubernetesDelegateYamlOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DelegateDownloadResourceApiDownloadKubernetesDelegateYamlOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.**| Account Identifier for the Entity. | 
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

