# {{classname}}

All URIs are relative to *https://app.harness.io/gateway/authz/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccessControlList**](AclApi.md#GetAccessControlList) | **Post** /acl | Check for permission on resource(s) for a principal

# **GetAccessControlList**
> ResponseDtoAccessCheckResponse GetAccessControlList(ctx, body)
Check for permission on resource(s) for a principal

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AccessCheckRequest**](AccessCheckRequest.md)|  | 

### Return type

[**ResponseDtoAccessCheckResponse**](ResponseDTOAccessCheckResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

