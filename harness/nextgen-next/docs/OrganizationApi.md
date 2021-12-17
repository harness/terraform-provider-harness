# {{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganization**](OrganizationApi.md#CreateOrganization) | **Post** /ng/api/organizations | Creates an Organization
[**DeleteOrganization**](OrganizationApi.md#DeleteOrganization) | **Delete** /ng/api/organizations/{identifier} | Deletes Organization by identifier
[**GetOrganization**](OrganizationApi.md#GetOrganization) | **Get** /ng/api/organizations/{identifier} | Get the Organization by accountIdentifier and orgIdentifier
[**GetOrganizationList**](OrganizationApi.md#GetOrganizationList) | **Get** /ng/api/organizations | Get the list of organizations satisfying the criteria (if any) in the request
[**ListAllOrganizations**](OrganizationApi.md#ListAllOrganizations) | **Post** /ng/api/organizations/all-organizations | 
[**PutOrganization**](OrganizationApi.md#PutOrganization) | **Put** /ng/api/organizations/{identifier} | Updates the Organization

# **CreateOrganization**
> ResponseDtoOrganizationResponse CreateOrganization(ctx, body, accountIdentifier)
Creates an Organization

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**OrganizationRequest**](OrganizationRequest.md)| Details of the Organization to create | 
  **accountIdentifier** | **string**| Account Identifier for the entity | 

### Return type

[**ResponseDtoOrganizationResponse**](ResponseDTOOrganizationResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrganization**
> ResponseDtoBoolean DeleteOrganization(ctx, identifier, accountIdentifier, optional)
Deletes Organization by identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| Organization Identifier for the entity | 
  **accountIdentifier** | **string**| Account Identifier for the entity | 
 **optional** | ***OrganizationApiDeleteOrganizationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrganizationApiDeleteOrganizationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifMatch** | **optional.String**|  | 

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
  **identifier** | **string**| Organization Identifier for the entity | 
  **accountIdentifier** | **string**| Account Identifier for the entity | 

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
Get the list of organizations satisfying the criteria (if any) in the request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the entity | 
 **optional** | ***OrganizationApiGetOrganizationListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrganizationApiGetOrganizationListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **identifiers** | [**optional.Interface of []string**](string.md)| list of Project Ids for filtering results | 
 **searchTerm** | **optional.String**| Search Term | 
 **pageIndex** | **optional.Int32**|  | [default to 0]
 **pageSize** | **optional.Int32**|  | [default to 50]
 **sortOrders** | [**optional.Interface of []SortOrder**](SortOrder.md)|  | 

### Return type

[**ResponseDtoPageResponseOrganizationResponse**](ResponseDTOPageResponseOrganizationResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAllOrganizations**
> ListAllOrganizations(ctx, body, accountIdentifier, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]string**](string.md)| list of ProjectIdentifiers to filter results by | 
  **accountIdentifier** | **string**| Account Identifier for the entity | 
 **optional** | ***OrganizationApiListAllOrganizationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrganizationApiListAllOrganizationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **searchTerm** | **optional.**| Search term | 

### Return type

 (empty response body)

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
  **identifier** | **string**| Organization Identifier for the entity | 
  **accountIdentifier** | **string**| Account Identifier for the entity | 
 **optional** | ***OrganizationApiPutOrganizationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrganizationApiPutOrganizationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ifMatch** | **optional.**|  | 

### Return type

[**ResponseDtoOrganizationResponse**](ResponseDTOOrganizationResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

