# {{classname}}

All URIs are relative to */gateway/code/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListGitignore**](ResourceApi.md#ListGitignore) | **Get** /resources/gitignore | List available gitignore names
[**ListLicenses**](ResourceApi.md#ListLicenses) | **Get** /resources/license | List available license names

# **ListGitignore**
> []string ListGitignore(ctx, )
List available gitignore names

### Required Parameters
This endpoint does not need any parameter.

### Return type

**[]string**

### Authorization

[bearerAuth](../README.md#bearerAuth), [x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListLicenses**
> []InlineResponse200 ListLicenses(ctx, )
List available license names

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]InlineResponse200**](inline_response_200.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

