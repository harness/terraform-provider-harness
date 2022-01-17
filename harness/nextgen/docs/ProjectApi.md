# {{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteProject**](ProjectApi.md#DeleteProject) | **Delete** /ng/api/projects/{identifier} | Deletes the Project corresponding to the specified Project ID.
[**GetProject**](ProjectApi.md#GetProject) | **Get** /ng/api/projects/{identifier} | Gets a Project by ID
[**GetProjectList**](ProjectApi.md#GetProjectList) | **Get** /ng/api/projects | List user&#x27;s project
[**PostProject**](ProjectApi.md#PostProject) | **Post** /ng/api/projects | Creates a Project
[**PutProject**](ProjectApi.md#PutProject) | **Put** /ng/api/projects/{identifier} | Update Project by ID

# **DeleteProject**
> ResponseDtoBoolean DeleteProject(ctx, identifier, accountIdentifier, optional)
Deletes the Project corresponding to the specified Project ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| Project Identifier for the Entity | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
 **optional** | ***ProjectApiDeleteProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectApiDeleteProjectOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifMatch** | **optional.String**| Version number of Project | 
 **orgIdentifier** | **optional.String**| This is the Organization Identifier for the Project. By default, the Default Organization&#x27;s Identifier is considered. | [default to default]

### Return type

[**ResponseDtoBoolean**](ResponseDTOBoolean.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProject**
> ResponseDtoProjectResponse GetProject(ctx, identifier, accountIdentifier, optional)
Gets a Project by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| Project Identifier for the Entity | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
 **optional** | ***ProjectApiGetProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectApiGetProjectOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.String**| Organization identifier for the project. If left empty, Default Organization is assumed | [default to default]

### Return type

[**ResponseDtoProjectResponse**](ResponseDTOProjectResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectList**
> ResponseDtoPageResponseProjectResponse GetProjectList(ctx, accountIdentifier, optional)
List user's project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
 **optional** | ***ProjectApiGetProjectListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectApiGetProjectListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity | 
 **hasModule** | **optional.Bool**| This boolean specifies whether to Filter Projects which has the Module of type passed in the module type parameter or to Filter Projects which does not has the Module of type passed in the module type parameter | [default to true]
 **identifiers** | [**optional.Interface of []string**](string.md)| This is the list of Project IDs. Details specific to these IDs would be fetched. | 
 **moduleType** | **optional.String**| Filter Projects by module type | 
 **searchTerm** | **optional.String**| This would be used to filter Projects. Any Project having the specified string in its Name, ID and Tag would be filtered. | 
 **pageIndex** | **optional.Int32**| Indicates the number of pages. Results for these pages will be retrieved. | [default to 0]
 **pageSize** | **optional.Int32**| The number of the elements to fetch | [default to 50]
 **sortOrders** | [**optional.Interface of []SortOrder**](SortOrder.md)| Sort criteria for the elements. | 

### Return type

[**ResponseDtoPageResponseProjectResponse**](ResponseDTOPageResponseProjectResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostProject**
> ResponseDtoProjectResponse PostProject(ctx, body, accountIdentifier, optional)
Creates a Project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ProjectRequest**](ProjectRequest.md)| Details of the Project to create | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
 **optional** | ***ProjectApiPostProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectApiPostProjectOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.**| Organization identifier for the Project. If left empty, the Project is created under Default Organization | [default to default]

### Return type

[**ResponseDtoProjectResponse**](ResponseDTOProjectResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutProject**
> ResponseDtoProjectResponse PutProject(ctx, body, identifier, accountIdentifier, optional)
Update Project by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ProjectRequest**](ProjectRequest.md)| This is the updated Project. Please provide values for all fields, not just the fields you are updating | 
  **identifier** | **string**| Project Identifier for the Entity | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
 **optional** | ***ProjectApiPutProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectApiPutProjectOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ifMatch** | **optional.**| Version number of Project | 
 **orgIdentifier** | **optional.**| Organization identifier for the Project. If left empty, Default Organization is assumed | [default to default]

### Return type

[**ResponseDtoProjectResponse**](ResponseDTOProjectResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

