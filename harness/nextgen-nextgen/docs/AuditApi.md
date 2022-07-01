# nextgen{{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAuditEventList**](AuditApi.md#GetAuditEventList) | **Post** /audit/api/audits/list | List Audit Events

# **GetAuditEventList**
> ResponseDtoPageResponseAuditEvent GetAuditEventList(ctx, accountIdentifier, optional)
List Audit Events

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***AuditApiGetAuditEventListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuditApiGetAuditEventListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of AuditFilterProperties**](AuditFilterProperties.md)| This has the filter attributes for listing Audit Events | 
 **pageIndex** | **optional.**| Page Index of the results to fetch.Default Value: 0 | [default to 0]
 **pageSize** | **optional.**| Results per page(max 100)Default Value: 50 | [default to 50]
 **sortOrders** | [**optional.Interface of []SortOrder**](SortOrder.md)| Sort criteria for the elements. | 

### Return type

[**ResponseDtoPageResponseAuditEvent**](ResponseDTOPageResponseAuditEvent.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

