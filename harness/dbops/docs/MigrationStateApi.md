# dbops{{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1MigrationStateProjDbSchema**](MigrationStateApi.md#V1MigrationStateProjDbSchema) | **Post** /v1/orgs/{org}/projects/{project}/dbschema/{dbschema}/migrationstate | Migration state of a schema

# **V1MigrationStateProjDbSchema**
> MigrationStateOut V1MigrationStateProjDbSchema(ctx, org, project, dbschema, optional)
Migration state of a schema

Migration state of a schema

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier | 
  **project** | **string**| Project identifier | 
  **dbschema** | **string**| Identifier of the database schema | 
 **optional** | ***MigrationStateApiV1MigrationStateProjDbSchemaOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MigrationStateApiV1MigrationStateProjDbSchemaOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of MigrationStateIn**](MigrationStateIn.md)| Inputs to get migration state of schema | 
 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 
 **page** | **optional.**| Pagination page number strategy: Specify the page number within the paginated collection related to the number of items on each page. | [default to 0]
 **limit** | **optional.**| Pagination: Number of items to return. | [default to 10]
 **searchTerm** | **optional.**| This would be used to filter resources having attributes matching the search term. | 

### Return type

[**MigrationStateOut**](MigrationStateOut.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

