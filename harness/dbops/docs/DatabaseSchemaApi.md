# nextgen{{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1CreateProjDbSchema**](DatabaseSchemaApi.md#V1CreateProjDbSchema) | **Post** /v1/orgs/{org}/projects/{project}/dbschema | Create a database schema
[**V1DeleteProjDbSchema**](DatabaseSchemaApi.md#V1DeleteProjDbSchema) | **Delete** /v1/orgs/{org}/projects/{project}/dbschema/{dbschema} | Delete a database schema
[**V1GetProjDbSchema**](DatabaseSchemaApi.md#V1GetProjDbSchema) | **Get** /v1/orgs/{org}/projects/{project}/dbschema/{dbschema} | Get a database schema
[**V1ListProjDbSchema**](DatabaseSchemaApi.md#V1ListProjDbSchema) | **Get** /v1/orgs/{org}/projects/{project}/dbschema | List database schemas
[**V1UpdateProjDbSchema**](DatabaseSchemaApi.md#V1UpdateProjDbSchema) | **Put** /v1/orgs/{org}/projects/{project}/dbschema/{dbschema} | Update a database schema
[**V1UpdateProjDbSchemaInstance**](DatabaseSchemaApi.md#V1UpdateProjDbSchemaInstance) | **Put** /v1/orgs/{org}/projects/{project}/dbschema/{dbschema}/instance/{dbinstance} | Update a database instance

# **V1CreateProjDbSchema**
> DbSchemaOut V1CreateProjDbSchema(ctx, body, org, project, optional)
Create a database schema

Create a database schema

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DbSchemaIn**](DbSchemaIn.md)|  | 
  **org** | **string**| Organization identifier | 
  **project** | **string**| Project identifier | 
 **optional** | ***DatabaseSchemaApiV1CreateProjDbSchemaOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DatabaseSchemaApiV1CreateProjDbSchemaOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**DbSchemaOut**](DBSchemaOut.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1DeleteProjDbSchema**
> V1DeleteProjDbSchema(ctx, org, project, dbschema, optional)
Delete a database schema

Delete a database schema

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier | 
  **project** | **string**| Project identifier | 
  **dbschema** | **string**| Identifier of the database schema | 
 **optional** | ***DatabaseSchemaApiV1DeleteProjDbSchemaOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DatabaseSchemaApiV1DeleteProjDbSchemaOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1GetProjDbSchema**
> DbSchemaOut V1GetProjDbSchema(ctx, org, project, dbschema, optional)
Get a database schema

Retrieves the specified database schema

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier | 
  **project** | **string**| Project identifier | 
  **dbschema** | **string**| Identifier of the database schema | 
 **optional** | ***DatabaseSchemaApiV1GetProjDbSchemaOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DatabaseSchemaApiV1GetProjDbSchemaOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**DbSchemaOut**](DBSchemaOut.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ListProjDbSchema**
> []DbSchemaOut V1ListProjDbSchema(ctx, org, project, optional)
List database schemas

List database Schemas

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier | 
  **project** | **string**| Project identifier | 
 **optional** | ***DatabaseSchemaApiV1ListProjDbSchemaOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DatabaseSchemaApiV1ListProjDbSchemaOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 
 **page** | **optional.Int64**| Pagination page number strategy: Specify the page number within the paginated collection related to the number of items on each page. | [default to 0]
 **limit** | **optional.Int32**| Pagination: Number of items to return. | [default to 10]
 **searchTerm** | **optional.String**| This would be used to filter resources having attributes matching the search term. | 
 **sort** | **optional.String**| Parameter on the basis of which sorting is done. | 
 **order** | **optional.String**| Order on the basis of which sorting is done. | 

### Return type

[**[]DbSchemaOut**](DBSchemaOut.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1UpdateProjDbSchema**
> DbSchemaOut V1UpdateProjDbSchema(ctx, body, org, project, dbschema, optional)
Update a database schema

Update a database schema

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**interface{}**](interface{}.md)| Database schema update request | 
  **org** | **string**| Organization identifier | 
  **project** | **string**| Project identifier | 
  **dbschema** | **string**| Identifier of the database schema | 
 **optional** | ***DatabaseSchemaApiV1UpdateProjDbSchemaOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DatabaseSchemaApiV1UpdateProjDbSchemaOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**DbSchemaOut**](DBSchemaOut.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1UpdateProjDbSchemaInstance**
> DbInstanceOut V1UpdateProjDbSchemaInstance(ctx, body, org, project, dbschema, dbinstance, optional)
Update a database instance

Update a database instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**interface{}**](interface{}.md)| Database instance update request | 
  **org** | **string**| Organization identifier | 
  **project** | **string**| Project identifier | 
  **dbschema** | **string**| Identifier of the database schema | 
  **dbinstance** | **string**| database instance unique id | 
 **optional** | ***DatabaseSchemaApiV1UpdateProjDbSchemaInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DatabaseSchemaApiV1UpdateProjDbSchemaInstanceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**DbInstanceOut**](DBInstanceOut.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

