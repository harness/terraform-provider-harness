# nextgen{{classname}}

All URIs are relative to *https://app.harness.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccountScopedDataSinks**](AccountDataSinksApi.md#CreateAccountScopedDataSinks) | **Post** /v1/data-sinks | Create Data Sink
[**DeleteAccountScopedDataSink**](AccountDataSinksApi.md#DeleteAccountScopedDataSink) | **Delete** /v1/data-sinks/{data-sink} | Delete Data Sink
[**GetAccountScopedDataSink**](AccountDataSinksApi.md#GetAccountScopedDataSink) | **Get** /v1/data-sinks/{data-sink} | Get Data Sink
[**GetAccountScopedDataSinks**](AccountDataSinksApi.md#GetAccountScopedDataSinks) | **Get** /v1/data-sinks | Get Data Sinks
[**UpdateAccountScopedDataSink**](AccountDataSinksApi.md#UpdateAccountScopedDataSink) | **Put** /v1/data-sinks/{data-sink} | Update Data Sink
[**ValidateAccountScopedDataSinkIdentifier**](AccountDataSinksApi.md#ValidateAccountScopedDataSinkIdentifier) | **Get** /v1/data-sinks/validate-unique-identifier/{data-sink} | Your GET endpoint

# **CreateAccountScopedDataSinks**
> DataSinkResponseDto CreateAccountScopedDataSinks(ctx, optional)
Create Data Sink

Creates a data sinks in account scope

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountDataSinksApiCreateAccountScopedDataSinksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountDataSinksApiCreateAccountScopedDataSinksOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of DataSinkDto**](DataSinkDto.md)| Request body to create a Data Sink | 
 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**DataSinkResponseDto**](DataSinkResponseDTO.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAccountScopedDataSink**
> DeleteAccountScopedDataSink(ctx, dataSink, optional)
Delete Data Sink

Deletes the specified data-sink in account scope

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dataSink** | **string**| Identifier field for the data-sink | 
 **optional** | ***AccountDataSinksApiDeleteAccountScopedDataSinkOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountDataSinksApiDeleteAccountScopedDataSinkOpts struct
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

# **GetAccountScopedDataSink**
> DataSinkResponseDto GetAccountScopedDataSink(ctx, dataSink, optional)
Get Data Sink

Retrieves the specified data-sink in account scope

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dataSink** | **string**| Identifier field for the data-sink | 
 **optional** | ***AccountDataSinksApiGetAccountScopedDataSinkOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountDataSinksApiGetAccountScopedDataSinkOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**DataSinkResponseDto**](DataSinkResponseDTO.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountScopedDataSinks**
> []DataSinkResponseDto GetAccountScopedDataSinks(ctx, optional)
Get Data Sinks

Retreives the list of data sinks in account scope

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountDataSinksApiGetAccountScopedDataSinksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountDataSinksApiGetAccountScopedDataSinksOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 
 **page** | **optional.Int32**| Pagination page number strategy: Specify the page number within the paginated collection related to the number of items on each page. | [default to 0]
 **limit** | **optional.Int32**| Pagination: Number of items to return. | [default to 30]
 **sort** | **optional.String**| Parameter on the basis of which sorting is done. | 
 **order** | **optional.String**| Order on the basis of which sorting is done. | 
 **searchTerm** | **optional.String**| This would be used to filter resources having attributes matching the search term. | 
 **status** | [**optional.Interface of Status**](.md)| Filter data-sinks by status | 
 **type_** | [**optional.Interface of []string**](string.md)| This would be used to filter data-sinks having type matching either of the values of this attribute. | 

### Return type

[**[]DataSinkResponseDto**](DataSinkResponseDTO.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAccountScopedDataSink**
> DataSinkResponseDto UpdateAccountScopedDataSink(ctx, dataSink, optional)
Update Data Sink

Updates the specified data-sink in account scope

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dataSink** | **string**| Identifier field for the data-sink | 
 **optional** | ***AccountDataSinksApiUpdateAccountScopedDataSinkOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountDataSinksApiUpdateAccountScopedDataSinkOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of DataSinkDto**](DataSinkDto.md)| Request body to update a Data Sink | 
 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**DataSinkResponseDto**](DataSinkResponseDTO.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateAccountScopedDataSinkIdentifier**
> bool ValidateAccountScopedDataSinkIdentifier(ctx, dataSink, optional)
Your GET endpoint

Validate if the specified data-sink identifier is available for use in account scope

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dataSink** | **string**| Identifier field for the data-sink | 
 **optional** | ***AccountDataSinksApiValidateAccountScopedDataSinkIdentifierOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountDataSinksApiValidateAccountScopedDataSinkIdentifierOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

**bool**

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

