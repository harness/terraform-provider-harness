# {{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateResourceGroup**](HarnessResourceGroupApi.md#CreateResourceGroup) | **Post** /resourcegroup/api/resourcegroup | Create a resource group
[**DeleteResourceGroup**](HarnessResourceGroupApi.md#DeleteResourceGroup) | **Delete** /resourcegroup/api/resourcegroup/{identifier} | Delete a resource group
[**GetFilterResourceGroupList**](HarnessResourceGroupApi.md#GetFilterResourceGroupList) | **Post** /resourcegroup/api/resourcegroup/filter | 
[**GetResourceGroup**](HarnessResourceGroupApi.md#GetResourceGroup) | **Get** /resourcegroup/api/resourcegroup/{identifier} | Get a resource group by identifier
[**GetResourceGroupList**](HarnessResourceGroupApi.md#GetResourceGroupList) | **Get** /resourcegroup/api/resourcegroup | Get list of resource groups
[**UpdateResourceGroup**](HarnessResourceGroupApi.md#UpdateResourceGroup) | **Put** /resourcegroup/api/resourcegroup/{identifier} | Update a resource group

# **CreateResourceGroup**
> ResponseDtoResourceGroupResponse CreateResourceGroup(ctx, body, accountIdentifier, optional)
Create a resource group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ResourceGroupRequest**](ResourceGroupRequest.md)| This contains the details required to create a Resource Group | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
 **optional** | ***HarnessResourceGroupApiCreateResourceGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HarnessResourceGroupApiCreateResourceGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity | 

### Return type

[**ResponseDtoResourceGroupResponse**](ResponseDTOResourceGroupResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteResourceGroup**
> ResponseDtoBoolean DeleteResourceGroup(ctx, identifier, accountIdentifier, optional)
Delete a resource group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**|  | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
 **optional** | ***HarnessResourceGroupApiDeleteResourceGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HarnessResourceGroupApiDeleteResourceGroupOpts struct
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
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFilterResourceGroupList**
> ResponseDtoPageResponseResourceGroupResponse GetFilterResourceGroupList(ctx, body, optional)


This fetches a filtered list of Resource Groups

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ResourceGroupFilter**](ResourceGroupFilter.md)| Filter Resource Groups based on multiple parameters | 
 **optional** | ***HarnessResourceGroupApiGetFilterResourceGroupListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HarnessResourceGroupApiGetFilterResourceGroupListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageIndex** | **optional.**| Indicates the number of pages. Results for these pages will be retrieved. | [default to 0]
 **pageSize** | **optional.**| The number of the elements to fetch | [default to 50]
 **sortOrders** | [**optional.Interface of []SortOrder**](SortOrder.md)| Sort criteria for the elements. | 

### Return type

[**ResponseDtoPageResponseResourceGroupResponse**](ResponseDTOPageResponseResourceGroupResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetResourceGroup**
> ResponseDtoResourceGroupResponse GetResourceGroup(ctx, identifier, accountIdentifier, optional)
Get a resource group by identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| This is the Identifier of the Entity | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
 **optional** | ***HarnessResourceGroupApiGetResourceGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HarnessResourceGroupApiGetResourceGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity | 

### Return type

[**ResponseDtoResourceGroupResponse**](ResponseDTOResourceGroupResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetResourceGroupList**
> ResponseDtoPageResponseResourceGroupResponse GetResourceGroupList(ctx, accountIdentifier, optional)
Get list of resource groups

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
 **optional** | ***HarnessResourceGroupApiGetResourceGroupListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HarnessResourceGroupApiGetResourceGroupListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity | 
 **searchTerm** | **optional.String**| Details of all the resource groups having this string in their name or identifier will be returned. | 
 **pageIndex** | **optional.Int32**| Indicates the number of pages. Results for these pages will be retrieved. | [default to 0]
 **pageSize** | **optional.Int32**| The number of the elements to fetch | [default to 50]
 **sortOrders** | [**optional.Interface of []SortOrder**](SortOrder.md)| Sort criteria for the elements. | 

### Return type

[**ResponseDtoPageResponseResourceGroupResponse**](ResponseDTOPageResponseResourceGroupResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateResourceGroup**
> ResponseDtoResourceGroupResponse UpdateResourceGroup(ctx, body, identifier, accountIdentifier, optional)
Update a resource group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ResourceGroupRequest**](ResourceGroupRequest.md)| This contains the details required to create a Resource Group | 
  **identifier** | **string**| Identifier for the entity | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
 **optional** | ***HarnessResourceGroupApiUpdateResourceGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HarnessResourceGroupApiUpdateResourceGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orgIdentifier** | **optional.**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity | 

### Return type

[**ResponseDtoResourceGroupResponse**](ResponseDTOResourceGroupResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

