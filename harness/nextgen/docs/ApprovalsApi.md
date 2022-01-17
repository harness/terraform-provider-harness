# {{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddHarnessApprovalActivity**](ApprovalsApi.md#AddHarnessApprovalActivity) | **Post** /pipeline/api/approvals/{approvalInstanceId}/harness/activity | Add a new Harness Approval activity
[**GetApprovalInstance**](ApprovalsApi.md#GetApprovalInstance) | **Get** /pipeline/api/approvals/{approvalInstanceId} | Gets an Approval Instance by identifier

# **AddHarnessApprovalActivity**
> ResponseDtoApprovalInstanceResponse AddHarnessApprovalActivity(ctx, body, approvalInstanceId)
Add a new Harness Approval activity

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**HarnessApprovalActivityRequest**](HarnessApprovalActivityRequest.md)| This contains the details of Harness Approval Activity requested | 
  **approvalInstanceId** | **string**| Approval Identifier for the entity | 

### Return type

[**ResponseDtoApprovalInstanceResponse**](ResponseDTOApprovalInstanceResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApprovalInstance**
> ResponseDtoApprovalInstanceResponse GetApprovalInstance(ctx, approvalInstanceId)
Gets an Approval Instance by identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **approvalInstanceId** | **string**| Approval Identifier for the entity | 

### Return type

[**ResponseDtoApprovalInstanceResponse**](ResponseDTOApprovalInstanceResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

