# nextgen{{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAuditFilter**](AuditFiltersApi.md#DeleteAuditFilter) | **Delete** /audit/api/auditFilters/{identifier} | Delete a Filter of type Audit by identifier
[**GetAuditFilter**](AuditFiltersApi.md#GetAuditFilter) | **Get** /audit/api/auditFilters/{identifier} | Gets a Filter of type Audit by identifier
[**GetAuditFilterList**](AuditFiltersApi.md#GetAuditFilterList) | **Get** /audit/api/auditFilters | Get the list of Filters of type Audit satisfying the criteria (if any) in the request
[**PostAuditFilter**](AuditFiltersApi.md#PostAuditFilter) | **Post** /audit/api/auditFilters | Creates a Filter
[**UpdateAuditFilter**](AuditFiltersApi.md#UpdateAuditFilter) | **Put** /audit/api/auditFilters | Updates the Filter of type Audit

# **DeleteAuditFilter**
> ResponseDtoBoolean DeleteAuditFilter(ctx, accountIdentifier, identifier, optional)
Delete a Filter of type Audit by identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **identifier** | **string**| Filter Identifier | 
 **optional** | ***AuditFiltersApiDeleteAuditFilterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuditFiltersApiDeleteAuditFilterOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 

### Return type

[**ResponseDtoBoolean**](ResponseDTOBoolean.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAuditFilter**
> ResponseDtoFilter GetAuditFilter(ctx, accountIdentifier, identifier, optional)
Gets a Filter of type Audit by identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **identifier** | **string**| Filter Identifier | 
 **optional** | ***AuditFiltersApiGetAuditFilterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuditFiltersApiGetAuditFilterOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 

### Return type

[**ResponseDtoFilter**](ResponseDTOFilter.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAuditFilterList**
> ResponseDtoPageResponseFilter GetAuditFilterList(ctx, accountIdentifier, optional)
Get the list of Filters of type Audit satisfying the criteria (if any) in the request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***AuditFiltersApiGetAuditFilterListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuditFiltersApiGetAuditFilterListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageIndex** | **optional.Int32**| Page number of navigation. If left empty, default value of 0 is assumed | [default to 0]
 **pageSize** | **optional.Int32**| Number of entries per page. If left empty, default value of 100 is assumed | [default to 100]
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 

### Return type

[**ResponseDtoPageResponseFilter**](ResponseDTOPageResponseFilter.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostAuditFilter**
> ResponseDtoFilter PostAuditFilter(ctx, body, accountIdentifier)
Creates a Filter

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Filter**](Filter.md)| Details of the Filter to create | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 

### Return type

[**ResponseDtoFilter**](ResponseDTOFilter.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, text/yaml, text/html, text/plain
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAuditFilter**
> ResponseDtoFilter UpdateAuditFilter(ctx, body, accountIdentifier)
Updates the Filter of type Audit

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Filter**](Filter.md)| This is the updated Filter. This should have all the fields not just the updated ones | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 

### Return type

[**ResponseDtoFilter**](ResponseDTOFilter.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, text/yaml, text/html, text/plain
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

