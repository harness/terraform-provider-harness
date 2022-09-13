# nextgen{{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteEnvironmentGroup**](EnvironmentGroupApi.md#DeleteEnvironmentGroup) | **Delete** /ng/api/environmentGroup/{envGroupIdentifier} | Delete en Environment Group by Identifier
[**GetEnvironmentGroup**](EnvironmentGroupApi.md#GetEnvironmentGroup) | **Get** /ng/api/environmentGroup/{envGroupIdentifier} | Gets an Environment Group by identifier
[**GetEnvironmentGroupList**](EnvironmentGroupApi.md#GetEnvironmentGroupList) | **Post** /ng/api/environmentGroup/list | Gets Environment Group list for a Project
[**PostEnvironmentGroup**](EnvironmentGroupApi.md#PostEnvironmentGroup) | **Post** /ng/api/environmentGroup | Create an Environment Group
[**UpdateEnvironmentGroup**](EnvironmentGroupApi.md#UpdateEnvironmentGroup) | **Put** /ng/api/environmentGroup/{envGroupIdentifier} | Update an Environment Group by Identifier

# **DeleteEnvironmentGroup**
> ResponseDtoEnvironmentGroupDelete DeleteEnvironmentGroup(ctx, envGroupIdentifier, accountIdentifier, orgIdentifier, projectIdentifier, optional)
Delete en Environment Group by Identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **envGroupIdentifier** | **string**| Environment Group Identifier for the entity | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the Entity. | 
  **projectIdentifier** | **string**| Project Identifier for the Entity. | 
 **optional** | ***EnvironmentGroupApiDeleteEnvironmentGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EnvironmentGroupApiDeleteEnvironmentGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **ifMatch** | **optional.String**|  | 
 **branch** | **optional.String**| Name of the branch. | 
 **repoIdentifier** | **optional.String**| Git Sync Config Id. | 
 **rootFolder** | **optional.String**| Path to the root folder of the Entity. | 
 **filePath** | **optional.String**| File Path of the Entity. | 
 **commitMsg** | **optional.String**| Commit Message to use for the merge commit. | 
 **lastObjectId** | **optional.String**| Last Object Id | 

### Return type

[**ResponseDtoEnvironmentGroupDelete**](ResponseDTOEnvironmentGroupDelete.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEnvironmentGroup**
> ResponseDtoEnvironmentGroup GetEnvironmentGroup(ctx, envGroupIdentifier, accountIdentifier, orgIdentifier, projectIdentifier, optional)
Gets an Environment Group by identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **envGroupIdentifier** | **string**| Environment Group Identifier for the entity | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the Entity. | 
  **projectIdentifier** | **string**| Project Identifier for the Entity. | 
 **optional** | ***EnvironmentGroupApiGetEnvironmentGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EnvironmentGroupApiGetEnvironmentGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **deleted** | **optional.Bool**| Specify whether Environment is deleted or not | [default to false]
 **branch** | **optional.String**| Name of the branch. | 
 **repoIdentifier** | **optional.String**| Git Sync Config Id. | 
 **getDefaultFromOtherRepo** | **optional.Bool**| if true, return all the default entities | 

### Return type

[**ResponseDtoEnvironmentGroup**](ResponseDTOEnvironmentGroup.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEnvironmentGroupList**
> ResponseDtoPageResponseEnvironmentGroup GetEnvironmentGroupList(ctx, accountIdentifier, orgIdentifier, projectIdentifier, optional)
Gets Environment Group list for a Project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the Entity. | 
  **projectIdentifier** | **string**| Project Identifier for the Entity. | 
 **optional** | ***EnvironmentGroupApiGetEnvironmentGroupListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EnvironmentGroupApiGetEnvironmentGroupListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of FilterProperties**](FilterProperties.md)| This is the body for the filter properties for listing Environment Groups | 
 **envGroupIdentifiers** | [**optional.Interface of []string**](string.md)|  | 
 **searchTerm** | **optional.**| The word to be searched and included in the list response | 
 **page** | **optional.**| Page Index of the results to fetch.Default Value: 0 | [default to 0]
 **size** | **optional.**| Results per page | [default to 25]
 **sort** | [**optional.Interface of []string**](string.md)| Sort criteria for the elements. | 
 **filterIdentifier** | **optional.**| Filter identifier | 
 **branch** | **optional.**| Name of the branch. | 
 **repoIdentifier** | **optional.**| Git Sync Config Id. | 
 **getDefaultFromOtherRepo** | **optional.**| if true, return all the default entities | 

### Return type

[**ResponseDtoPageResponseEnvironmentGroup**](ResponseDTOPageResponseEnvironmentGroup.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostEnvironmentGroup**
> ResponseDtoEnvironmentGroup PostEnvironmentGroup(ctx, accountIdentifier, optional)
Create an Environment Group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***EnvironmentGroupApiPostEnvironmentGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EnvironmentGroupApiPostEnvironmentGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of EnvironmentGroupRequest**](EnvironmentGroupRequest.md)| Details of the Environment Group to be created | 
 **branch** | **optional.**| Name of the branch. | 
 **repoIdentifier** | **optional.**| Git Sync Config Id. | 
 **getDefaultFromOtherRepo** | **optional.**| if true, return all the default entities | 

### Return type

[**ResponseDtoEnvironmentGroup**](ResponseDTOEnvironmentGroup.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEnvironmentGroup**
> ResponseDtoEnvironmentGroup UpdateEnvironmentGroup(ctx, envGroupIdentifier, accountIdentifier, optional)
Update an Environment Group by Identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **envGroupIdentifier** | **string**| Environment Group Identifier for the entity | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***EnvironmentGroupApiUpdateEnvironmentGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EnvironmentGroupApiUpdateEnvironmentGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of EnvironmentGroupRequest**](EnvironmentGroupRequest.md)| Details of the Environment Group to be updated | 
 **ifMatch** | **optional.**|  | 
 **branch** | **optional.**| Name of the branch. | 
 **repoIdentifier** | **optional.**| Git Sync Config Id. | 
 **rootFolder** | **optional.**| Path to the root folder of the Entity. | 
 **filePath** | **optional.**| Path to the root folder of the Entity. | 
 **commitMsg** | **optional.**| Commit Message to use for the merge commit. | 
 **lastObjectId** | **optional.**| Last Object Id | 
 **resolvedConflictCommitId** | **optional.**| If the entity is git-synced, this parameter represents the commit id against which file conflicts are resolved | 
 **baseBranch** | **optional.**| Name of the default branch. | 
 **connectorRef** | **optional.**| Identifier of Connector needed for CRUD operations on the respective Entity | 

### Return type

[**ResponseDtoEnvironmentGroup**](ResponseDTOEnvironmentGroup.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

