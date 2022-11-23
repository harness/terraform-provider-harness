# {{classname}}

All URIs are relative to *https://app.harness.io/gateway/pm/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SystemHealth**](SystemApi.md#SystemHealth) | **Get** /api/v1/system/health | 
[**SystemVersion**](SystemApi.md#SystemVersion) | **Get** /api/v1/system/version | 

# **SystemHealth**
> SystemHealth(ctx, )


Check service health

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SystemVersion**
> ServiceVersion SystemVersion(ctx, )


Check service version

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ServiceVersion**](ServiceVersion.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

