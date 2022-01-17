# {{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServiceAccount**](ServiceAccountApi.md#CreateServiceAccount) | **Post** /ng/api/serviceaccount | Creates a Service Account
[**DeleteServiceAccount**](ServiceAccountApi.md#DeleteServiceAccount) | **Delete** /ng/api/serviceaccount/{identifier} | Deletes Service Account by ID
[**GetAggregatedServiceAccount**](ServiceAccountApi.md#GetAggregatedServiceAccount) | **Get** /ng/api/serviceaccount/aggregate/{identifier} | Get the Service Account by accountIdentifier and Service Account ID and Scope.
[**ListAggregatedServiceAccounts**](ServiceAccountApi.md#ListAggregatedServiceAccounts) | **Get** /ng/api/serviceaccount/aggregate | Fetches the list of Aggregated Service Accounts corresponding to the request&#x27;s filter criteria.
[**ListServiceAccount**](ServiceAccountApi.md#ListServiceAccount) | **Get** /ng/api/serviceaccount | Fetches the list of Service Accounts corresponding to the request&#x27;s filter criteria.
[**UpdateServiceAccount**](ServiceAccountApi.md#UpdateServiceAccount) | **Put** /ng/api/serviceaccount/{identifier} | Updates the Service Account.

# **CreateServiceAccount**
> ResponseDtoServiceAccountDto CreateServiceAccount(ctx, body, accountIdentifier, optional)
Creates a Service Account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServiceAccountDto**](ServiceAccountDto.md)| Details required to create Service Account | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
 **optional** | ***ServiceAccountApiCreateServiceAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServiceAccountApiCreateServiceAccountOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity | 

### Return type

[**ResponseDtoServiceAccountDto**](ResponseDTOServiceAccountDTO.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml, text/plain
 - **Accept**: application/json, application/yaml, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteServiceAccount**
> ResponseDtoBoolean DeleteServiceAccount(ctx, accountIdentifier, identifier, optional)
Deletes Service Account by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
  **identifier** | **string**| Service Account ID | 
 **optional** | ***ServiceAccountApiDeleteServiceAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServiceAccountApiDeleteServiceAccountOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity | 

### Return type

[**ResponseDtoBoolean**](ResponseDTOBoolean.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAggregatedServiceAccount**
> ResponseDtoServiceAccountAggregateDto GetAggregatedServiceAccount(ctx, accountIdentifier, identifier, optional)
Get the Service Account by accountIdentifier and Service Account ID and Scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
  **identifier** | **string**| Service Account IDr | 
 **optional** | ***ServiceAccountApiGetAggregatedServiceAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServiceAccountApiGetAggregatedServiceAccountOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity | 

### Return type

[**ResponseDtoServiceAccountAggregateDto**](ResponseDTOServiceAccountAggregateDTO.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAggregatedServiceAccounts**
> ResponseDtoPageResponseServiceAccountAggregateDto ListAggregatedServiceAccounts(ctx, accountIdentifier, optional)
Fetches the list of Aggregated Service Accounts corresponding to the request's filter criteria.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
 **optional** | ***ServiceAccountApiListAggregatedServiceAccountsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServiceAccountApiListAggregatedServiceAccountsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity | 
 **identifiers** | [**optional.Interface of []string**](string.md)| This is the list of Service Account IDs. Details specific to these IDs would be fetched. | 
 **pageIndex** | **optional.Int32**| Indicates the number of pages. Results for these pages will be retrieved. | [default to 0]
 **pageSize** | **optional.Int32**| The number of the elements to fetch | [default to 50]
 **sortOrders** | [**optional.Interface of []SortOrder**](SortOrder.md)| Sort criteria for the elements. | 
 **searchTerm** | **optional.String**| This would be used to filter Service Accounts. Any Service Account having the specified string in its Name, ID and Tag would be filtered. | 

### Return type

[**ResponseDtoPageResponseServiceAccountAggregateDto**](ResponseDTOPageResponseServiceAccountAggregateDTO.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListServiceAccount**
> ResponseDtoListServiceAccountDto ListServiceAccount(ctx, accountIdentifier, optional)
Fetches the list of Service Accounts corresponding to the request's filter criteria.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
 **optional** | ***ServiceAccountApiListServiceAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServiceAccountApiListServiceAccountOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity | 
 **identifiers** | [**optional.Interface of []string**](string.md)| This is the list of Service Account IDs. Details specific to these IDs would be fetched. | 

### Return type

[**ResponseDtoListServiceAccountDto**](ResponseDTOListServiceAccountDTO.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateServiceAccount**
> ResponseDtoServiceAccountDto UpdateServiceAccount(ctx, body, accountIdentifier, identifier, optional)
Updates the Service Account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServiceAccountDto**](ServiceAccountDto.md)| Details of the updated Service Account | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
  **identifier** | **string**| Service Account ID | 
 **optional** | ***ServiceAccountApiUpdateServiceAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServiceAccountApiUpdateServiceAccountOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orgIdentifier** | **optional.**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity | 

### Return type

[**ResponseDtoServiceAccountDto**](ResponseDTOServiceAccountDTO.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml, text/plain
 - **Accept**: application/json, application/yaml, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

