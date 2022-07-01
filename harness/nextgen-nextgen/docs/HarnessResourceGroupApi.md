# nextgen{{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateResourceGroup**](HarnessResourceGroupApi.md#CreateResourceGroup) | **Post** /resourcegroup/api/resourcegroup | Create a resource group
[**CreateResourceGroupV2**](HarnessResourceGroupApi.md#CreateResourceGroupV2) | **Post** /resourcegroup/api/v2/resourcegroup | Create Resource Group
[**DeleteResourceGroup**](HarnessResourceGroupApi.md#DeleteResourceGroup) | **Delete** /resourcegroup/api/resourcegroup/{identifier} | Delete a resource group
[**DeleteResourceGroupV2**](HarnessResourceGroupApi.md#DeleteResourceGroupV2) | **Delete** /resourcegroup/api/v2/resourcegroup/{identifier} | Delete Resource Group
[**GetFilterResourceGroupList**](HarnessResourceGroupApi.md#GetFilterResourceGroupList) | **Post** /resourcegroup/api/resourcegroup/filter | This fetches a filtered list of Resource Groups
[**GetFilterResourceGroupListV2**](HarnessResourceGroupApi.md#GetFilterResourceGroupListV2) | **Post** /resourcegroup/api/v2/resourcegroup/filter | List Resource Groups by filter
[**GetResourceGroup**](HarnessResourceGroupApi.md#GetResourceGroup) | **Get** /resourcegroup/api/resourcegroup/{identifier} | Get a resource group by identifier
[**GetResourceGroupList**](HarnessResourceGroupApi.md#GetResourceGroupList) | **Get** /resourcegroup/api/resourcegroup | Get list of resource groups
[**GetResourceGroupListV2**](HarnessResourceGroupApi.md#GetResourceGroupListV2) | **Get** /resourcegroup/api/v2/resourcegroup | List Resource Groups
[**GetResourceGroupV2**](HarnessResourceGroupApi.md#GetResourceGroupV2) | **Get** /resourcegroup/api/v2/resourcegroup/{identifier} | Get Resource Group
[**UpdateResourceGroup**](HarnessResourceGroupApi.md#UpdateResourceGroup) | **Put** /resourcegroup/api/resourcegroup/{identifier} | Update a resource group
[**UpdateResourceGroup1**](HarnessResourceGroupApi.md#UpdateResourceGroup1) | **Put** /resourcegroup/api/v2/resourcegroup/{identifier} | Update Resource Group

# **CreateResourceGroup**
> ResponseDtoResourceGroupResponse CreateResourceGroup(ctx, body, accountIdentifier, optional)
Create a resource group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ResourceGroupRequest**](ResourceGroupRequest.md)| This contains the details required to create a Resource Group | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***HarnessResourceGroupApiCreateResourceGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HarnessResourceGroupApiCreateResourceGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity. | 

### Return type

[**ResponseDtoResourceGroupResponse**](ResponseDTOResourceGroupResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateResourceGroupV2**
> ResponseDtoResourceGroupV2Response CreateResourceGroupV2(ctx, body, accountIdentifier, optional)
Create Resource Group

Create a resource group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ResourceGroupV2Request**](ResourceGroupV2Request.md)| This contains the details required to create a Resource Group | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***HarnessResourceGroupApiCreateResourceGroupV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HarnessResourceGroupApiCreateResourceGroupV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity. | 

### Return type

[**ResponseDtoResourceGroupV2Response**](ResponseDTOResourceGroupV2Response.md)

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
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***HarnessResourceGroupApiDeleteResourceGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HarnessResourceGroupApiDeleteResourceGroupOpts struct
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
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteResourceGroupV2**
> ResponseDtoBoolean DeleteResourceGroupV2(ctx, identifier, accountIdentifier, optional)
Delete Resource Group

Delete a resource group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| Identifier for the Entity. | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***HarnessResourceGroupApiDeleteResourceGroupV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HarnessResourceGroupApiDeleteResourceGroupV2Opts struct
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
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFilterResourceGroupList**
> ResponseDtoPageResponseResourceGroupResponse GetFilterResourceGroupList(ctx, body, accountIdentifier, optional)
This fetches a filtered list of Resource Groups

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ResourceGroupFilter**](ResourceGroupFilter.md)| Filter Resource Groups based on multiple parameters | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***HarnessResourceGroupApiGetFilterResourceGroupListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HarnessResourceGroupApiGetFilterResourceGroupListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageIndex** | **optional.**| Page Index of the results to fetch.Default Value: 0 | [default to 0]
 **pageSize** | **optional.**| Results per page(max 100)Default Value: 50 | [default to 50]
 **sortOrders** | [**optional.Interface of []SortOrder**](SortOrder.md)| Sort criteria for the elements. | 

### Return type

[**ResponseDtoPageResponseResourceGroupResponse**](ResponseDTOPageResponseResourceGroupResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFilterResourceGroupListV2**
> ResponseDtoPageResponseResourceGroupV2Response GetFilterResourceGroupListV2(ctx, body, accountIdentifier, optional)
List Resource Groups by filter

This fetches a filtered list of Resource Groups

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ResourceGroupFilter**](ResourceGroupFilter.md)| Filter Resource Groups based on multiple parameters | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***HarnessResourceGroupApiGetFilterResourceGroupListV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HarnessResourceGroupApiGetFilterResourceGroupListV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageIndex** | **optional.**| Page Index of the results to fetch.Default Value: 0 | [default to 0]
 **pageSize** | **optional.**| Results per page(max 100)Default Value: 50 | [default to 50]
 **sortOrders** | [**optional.Interface of []SortOrder**](SortOrder.md)| Sort criteria for the elements. | 

### Return type

[**ResponseDtoPageResponseResourceGroupV2Response**](ResponseDTOPageResponseResourceGroupV2Response.md)

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
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***HarnessResourceGroupApiGetResourceGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HarnessResourceGroupApiGetResourceGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 

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
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***HarnessResourceGroupApiGetResourceGroupListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HarnessResourceGroupApiGetResourceGroupListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **searchTerm** | **optional.String**| Details of all the resource groups having this string in their name or identifier will be returned. | 
 **pageIndex** | **optional.Int32**| Page Index of the results to fetch.Default Value: 0 | [default to 0]
 **pageSize** | **optional.Int32**| Results per page(max 100)Default Value: 50 | [default to 50]
 **sortOrders** | [**optional.Interface of []SortOrder**](SortOrder.md)| Sort criteria for the elements. | 

### Return type

[**ResponseDtoPageResponseResourceGroupResponse**](ResponseDTOPageResponseResourceGroupResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetResourceGroupListV2**
> ResponseDtoPageResponseResourceGroupV2Response GetResourceGroupListV2(ctx, accountIdentifier, optional)
List Resource Groups

Get list of resource groups

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***HarnessResourceGroupApiGetResourceGroupListV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HarnessResourceGroupApiGetResourceGroupListV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **searchTerm** | **optional.String**| Details of all the resource groups having this string in their name or identifier will be returned. | 
 **pageIndex** | **optional.Int32**| Page Index of the results to fetch.Default Value: 0 | [default to 0]
 **pageSize** | **optional.Int32**| Results per page(max 100)Default Value: 50 | [default to 50]
 **sortOrders** | [**optional.Interface of []SortOrder**](SortOrder.md)| Sort criteria for the elements. | 

### Return type

[**ResponseDtoPageResponseResourceGroupV2Response**](ResponseDTOPageResponseResourceGroupV2Response.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetResourceGroupV2**
> ResponseDtoResourceGroupV2Response GetResourceGroupV2(ctx, identifier, accountIdentifier, optional)
Get Resource Group

Get a resource group by identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| Identifier for the Entity. | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***HarnessResourceGroupApiGetResourceGroupV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HarnessResourceGroupApiGetResourceGroupV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 

### Return type

[**ResponseDtoResourceGroupV2Response**](ResponseDTOResourceGroupV2Response.md)

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
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***HarnessResourceGroupApiUpdateResourceGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HarnessResourceGroupApiUpdateResourceGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orgIdentifier** | **optional.**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity. | 

### Return type

[**ResponseDtoResourceGroupResponse**](ResponseDTOResourceGroupResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateResourceGroup1**
> ResponseDtoResourceGroupV2Response UpdateResourceGroup1(ctx, body, identifier, accountIdentifier, optional)
Update Resource Group

Update a resource group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ResourceGroupV2Request**](ResourceGroupV2Request.md)| This contains the details required to create a Resource Group | 
  **identifier** | **string**| Identifier for the Entity. | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***HarnessResourceGroupApiUpdateResourceGroup1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HarnessResourceGroupApiUpdateResourceGroup1Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orgIdentifier** | **optional.**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity. | 

### Return type

[**ResponseDtoResourceGroupV2Response**](ResponseDTOResourceGroupV2Response.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

