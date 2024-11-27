# dbops{{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1CreateProjDbSchemaInstance**](DatabaseInstanceApi.md#V1CreateProjDbSchemaInstance) | **Post** /v1/orgs/{org}/projects/{project}/dbschema/{dbschema}/instance | Create a database instance
[**V1DeleteProjDbSchemaInstance**](DatabaseInstanceApi.md#V1DeleteProjDbSchemaInstance) | **Delete** /v1/orgs/{org}/projects/{project}/dbschema/{dbschema}/instance/{dbinstance} | Delete a database instance
[**V1GetProjDbSchemaInstance**](DatabaseInstanceApi.md#V1GetProjDbSchemaInstance) | **Get** /v1/orgs/{org}/projects/{project}/dbschema/{dbschema}/instance/{dbinstance} | Get a database instance
[**V1ListProjDbSchemaInstance**](DatabaseInstanceApi.md#V1ListProjDbSchemaInstance) | **Post** /v1/orgs/{org}/projects/{project}/dbschema/{dbschema}/instancelist | List database instances

# **V1CreateProjDbSchemaInstance**
> DbInstanceOut V1CreateProjDbSchemaInstance(ctx, body, org, project, dbschema, optional)
Create a database instance

Create a database instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DbInstanceIn**](DbInstanceIn.md)|  | 
  **org** | **string**| Organization identifier | 
  **project** | **string**| Project identifier | 
  **dbschema** | **string**| Identifier of the database schema | 
 **optional** | ***DatabaseInstanceApiV1CreateProjDbSchemaInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DatabaseInstanceApiV1CreateProjDbSchemaInstanceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**DbInstanceOut**](DBInstanceOut.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1DeleteProjDbSchemaInstance**
> V1DeleteProjDbSchemaInstance(ctx, org, project, dbschema, dbinstance, optional)
Delete a database instance

Delete a database instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier | 
  **project** | **string**| Project identifier | 
  **dbschema** | **string**| Identifier of the database schema | 
  **dbinstance** | **string**| database instance unique id | 
 **optional** | ***DatabaseInstanceApiV1DeleteProjDbSchemaInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DatabaseInstanceApiV1DeleteProjDbSchemaInstanceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1GetProjDbSchemaInstance**
> DbInstanceOut V1GetProjDbSchemaInstance(ctx, org, project, dbschema, dbinstance, optional)
Get a database instance

Retrieves the specified database instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier | 
  **project** | **string**| Project identifier | 
  **dbschema** | **string**| Identifier of the database schema | 
  **dbinstance** | **string**| database instance unique id | 
 **optional** | ***DatabaseInstanceApiV1GetProjDbSchemaInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DatabaseInstanceApiV1GetProjDbSchemaInstanceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**DbInstanceOut**](DBInstanceOut.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ListProjDbSchemaInstance**
> []DbInstanceOut V1ListProjDbSchemaInstance(ctx, org, project, dbschema, optional)
List database instances

Retrieves the specified database instances of the database schema

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier | 
  **project** | **string**| Project identifier | 
  **dbschema** | **string**| Identifier of the database schema | 
 **optional** | ***DatabaseInstanceApiV1ListProjDbSchemaInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DatabaseInstanceApiV1ListProjDbSchemaInstanceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of DbInstanceFilterIn**](DbInstanceFilterIn.md)|  | 
 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 
 **page** | **optional.**| Pagination page number strategy: Specify the page number within the paginated collection related to the number of items on each page. | [default to 0]
 **limit** | **optional.**| Pagination: Number of items to return. | [default to 10]
 **searchTerm** | **optional.**| This would be used to filter resources having attributes matching the search term. | 
 **sort** | **optional.**| Parameter on the basis of which sorting is done. | [default to created]
 **order** | **optional.**| Order on the basis of which sorting is done. | [default to DESC]

### Return type

[**[]DbInstanceOut**](DBInstanceOut.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

