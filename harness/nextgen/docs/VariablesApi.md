# nextgen{{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVariable**](VariablesApi.md#CreateVariable) | **Post** /ng/api/variables | Creates a Variable.
[**DeleteVariable**](VariablesApi.md#DeleteVariable) | **Delete** /ng/api/variables/{identifier} | Deletes Variable by ID.
[**GetVariable**](VariablesApi.md#GetVariable) | **Get** /ng/api/variables/{identifier} | Get the Variable by scope identifiers and variable identifier.
[**GetVariableList**](VariablesApi.md#GetVariableList) | **Get** /ng/api/variables | Fetches the list of Variables.
[**UpdateVariable**](VariablesApi.md#UpdateVariable) | **Put** /ng/api/variables | Updates the Variable.

# **CreateVariable**
> ResponseDtoVariableResponseDto CreateVariable(ctx, body, accountIdentifier)
Creates a Variable.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**VariableRequestDto**](VariableRequestDto.md)| Details of the Variable to create. | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 

### Return type

[**ResponseDtoVariableResponseDto**](ResponseDTOVariableResponseDTO.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVariable**
> ResponseDtoBoolean DeleteVariable(ctx, accountIdentifier, identifier, optional)
Deletes Variable by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **identifier** | **string**| Variable ID | 
 **optional** | ***VariablesApiDeleteVariableOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VariablesApiDeleteVariableOpts struct
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

# **GetVariable**
> ResponseDtoVariableResponseDto GetVariable(ctx, identifier, accountIdentifier, optional)
Get the Variable by scope identifiers and variable identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| Variable ID | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***VariablesApiGetVariableOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VariablesApiGetVariableOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 

### Return type

[**ResponseDtoVariableResponseDto**](ResponseDTOVariableResponseDTO.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVariableList**
> ResponseDtoPageResponseVariableResponseDto GetVariableList(ctx, accountIdentifier, optional)
Fetches the list of Variables.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***VariablesApiGetVariableListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VariablesApiGetVariableListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **pageIndex** | **optional.Int32**| Page number of navigation. The default value is 0. | [default to 0]
 **pageSize** | **optional.Int32**| Number of entries per page. The default value is 100. | [default to 100]
 **searchTerm** | **optional.String**| This would be used to filter Variables. Any Variable having the specified string in its Name or ID would be filtered. | 
 **includeVariablesFromEverySubScope** | **optional.Bool**| Specify whether or not to include all the Variables accessible at the scope. For eg if set as true, at the Project scope we will get org and account Variable also in the response. | [default to false]

### Return type

[**ResponseDtoPageResponseVariableResponseDto**](ResponseDTOPageResponseVariableResponseDTO.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateVariable**
> ResponseDtoVariableResponseDto UpdateVariable(ctx, body, accountIdentifier)
Updates the Variable.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**VariableRequestDto**](VariableRequestDto.md)| Details of the variable to update. | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 

### Return type

[**ResponseDtoVariableResponseDto**](ResponseDTOVariableResponseDTO.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

