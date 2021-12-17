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
> ResponseDtoResourceGroupResponse CreateResourceGroup(ctx, accountIdentifier, optional)
Create a resource group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the entity | 
 **optional** | ***HarnessResourceGroupApiCreateResourceGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HarnessResourceGroupApiCreateResourceGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ResourceGroupRequest**](ResourceGroupRequest.md)| This contains the details required to create a Resource Group | 
 **orgIdentifier** | **optional.**| Organization Identifier for the entity | 
 **projectIdentifier** | **optional.**| Project Identifier for the entity | 

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
  **accountIdentifier** | **string**| Account Identifier for the entity | 
 **optional** | ***HarnessResourceGroupApiDeleteResourceGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HarnessResourceGroupApiDeleteResourceGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.String**| Organization Identifier for the entity | 
 **projectIdentifier** | **optional.String**| Project Identifier for the entity | 

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
  **body** | [**ResourceGroupFilter**](ResourceGroupFilter.md)| Filter Resource Group Entity based on multiple parameters | 
 **optional** | ***HarnessResourceGroupApiGetFilterResourceGroupListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HarnessResourceGroupApiGetFilterResourceGroupListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageIndex** | **optional.**|  | [default to 0]
 **pageSize** | **optional.**|  | [default to 50]
 **sortOrders** | [**optional.Interface of []SortOrder**](SortOrder.md)|  | 

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
  **identifier** | **string**| This is the ID of the Entity | 
  **accountIdentifier** | **string**| Account Identifier for the entity | 
 **optional** | ***HarnessResourceGroupApiGetResourceGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HarnessResourceGroupApiGetResourceGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.String**| Organization Identifier for the entity | 
 **projectIdentifier** | **optional.String**| Project Identifier for the entity | 

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
  **accountIdentifier** | **string**| Account Identifier for the entity | 
 **optional** | ***HarnessResourceGroupApiGetResourceGroupListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HarnessResourceGroupApiGetResourceGroupListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgIdentifier** | **optional.String**| Organization Identifier for the entity | 
 **projectIdentifier** | **optional.String**| Project Identifier for the entity | 
 **searchTerm** | **optional.String**| Search Term | 
 **pageIndex** | **optional.Int32**|  | [default to 0]
 **pageSize** | **optional.Int32**|  | [default to 50]
 **sortOrders** | [**optional.Interface of []SortOrder**](SortOrder.md)|  | 

### Return type

[**ResponseDtoPageResponseResourceGroupResponse**](ResponseDTOPageResponseResourceGroupResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateResourceGroup**
> ResponseDtoResourceGroupResponse UpdateResourceGroup(ctx, identifier, accountIdentifier, optional)
Update a resource group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| Identifier for the entity | 
  **accountIdentifier** | **string**| Account Identifier for the entity | 
 **optional** | ***HarnessResourceGroupApiUpdateResourceGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HarnessResourceGroupApiUpdateResourceGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ResourceGroupRequest**](ResourceGroupRequest.md)| This contains the details required to create a Resource Group | 
 **orgIdentifier** | **optional.**| Organization Identifier for the entity | 
 **projectIdentifier** | **optional.**| Project Identifier for the entity | 

### Return type

[**ResponseDtoResourceGroupResponse**](ResponseDTOResourceGroupResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

