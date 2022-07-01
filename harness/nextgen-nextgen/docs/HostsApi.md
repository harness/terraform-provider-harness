# nextgen{{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FilterHostsByConnector**](HostsApi.md#FilterHostsByConnector) | **Post** /ng/api/hosts/filter | Gets the list of hosts filtered by accountIdentifier and connectorIdentifier

# **FilterHostsByConnector**
> ResponseDtoPageResponseHostDto FilterHostsByConnector(ctx, accountIdentifier, optional)
Gets the list of hosts filtered by accountIdentifier and connectorIdentifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***HostsApiFilterHostsByConnectorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HostsApiFilterHostsByConnectorOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of HostFilterDto**](HostFilterDto.md)| Details of the filters applied | 
 **pageIndex** | **optional.**| Page Index of the results to fetch.Default Value: 0 | [default to 0]
 **pageSize** | **optional.**| Results per page(max 100)Default Value: 50 | [default to 50]
 **sortOrders** | [**optional.Interface of []SortOrder**](SortOrder.md)| Sort criteria for the elements. | 
 **orgIdentifier** | **optional.**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity. | 
 **identifier** | **optional.**| Connector Identifier | 

### Return type

[**ResponseDtoPageResponseHostDto**](ResponseDTOPageResponseHostDTO.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, text/yaml, text/html, text/plain
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

