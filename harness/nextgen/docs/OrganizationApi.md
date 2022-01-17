# {{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteOrganization**](OrganizationApi.md#DeleteOrganization) | **Delete** /ng/api/organizations/{identifier} | Deletes the Organization corresponding to the specified Organization ID.
[**GetOrganization**](OrganizationApi.md#GetOrganization) | **Get** /ng/api/organizations/{identifier} | Get the Organization by accountIdentifier and orgIdentifier
[**GetOrganizationList**](OrganizationApi.md#GetOrganizationList) | **Get** /ng/api/organizations | Get the list of Organizations satisfying the criteria (if any) in the request
[**PostOrganization**](OrganizationApi.md#PostOrganization) | **Post** /ng/api/organizations | Creates an Organization
[**PutOrganization**](OrganizationApi.md#PutOrganization) | **Put** /ng/api/organizations/{identifier} | Updates the Organization

# **DeleteOrganization**
> ResponseDtoBoolean DeleteOrganization(ctx, identifier, accountIdentifier, optional)
Deletes the Organization corresponding to the specified Organization ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| Organization Identifier for the Entity | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
 **optional** | ***OrganizationApiDeleteOrganizationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrganizationApiDeleteOrganizationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifMatch** | **optional.String**| Version number of the Organization | 

### Return type

[**ResponseDtoBoolean**](ResponseDTOBoolean.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrganization**
> ResponseDtoOrganizationResponse GetOrganization(ctx, identifier, accountIdentifier)
Get the Organization by accountIdentifier and orgIdentifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| Organization Identifier for the Entity | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 

### Return type

[**ResponseDtoOrganizationResponse**](ResponseDTOOrganizationResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrganizationList**
> ResponseDtoPageResponseOrganizationResponse GetOrganizationList(ctx, accountIdentifier, optional)
Get the list of Organizations satisfying the criteria (if any) in the request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
 **optional** | ***OrganizationApiGetOrganizationListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrganizationApiGetOrganizationListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **identifiers** | [**optional.Interface of []string**](string.md)| This is the list of Org Key IDs. Details specific to these IDs would be fetched. | 
 **searchTerm** | **optional.String**| This would be used to filter Organizations. Any Organization having the specified string in its Name, ID and Tag would be filtered. | 
 **pageIndex** | **optional.Int32**| Indicates the number of pages. Results for these pages will be retrieved. | [default to 0]
 **pageSize** | **optional.Int32**| The number of the elements to fetch | [default to 50]
 **sortOrders** | [**optional.Interface of []SortOrder**](SortOrder.md)| Sort criteria for the elements. | 

### Return type

[**ResponseDtoPageResponseOrganizationResponse**](ResponseDTOPageResponseOrganizationResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostOrganization**
> ResponseDtoOrganizationResponse PostOrganization(ctx, body, accountIdentifier)
Creates an Organization

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**OrganizationRequest**](OrganizationRequest.md)| Details of the Organization to create | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 

### Return type

[**ResponseDtoOrganizationResponse**](ResponseDTOOrganizationResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutOrganization**
> ResponseDtoOrganizationResponse PutOrganization(ctx, body, identifier, accountIdentifier, optional)
Updates the Organization

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**OrganizationRequest**](OrganizationRequest.md)| This is the updated Organization. Please provide values for all fields, not just the fields you are updating | 
  **identifier** | **string**| Organization Identifier for the Entity | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
 **optional** | ***OrganizationApiPutOrganizationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrganizationApiPutOrganizationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ifMatch** | **optional.**| Version number of the Organization | 

### Return type

[**ResponseDtoOrganizationResponse**](ResponseDTOOrganizationResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

