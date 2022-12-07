# nextgen{{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTrigger**](TriggersApi.md#CreateTrigger) | **Post** /pipeline/api/triggers | Creates Trigger for triggering target pipeline identifier.
[**DeleteTrigger**](TriggersApi.md#DeleteTrigger) | **Delete** /pipeline/api/triggers/{triggerIdentifier} | Deletes Trigger by identifier.
[**GetListForTarget**](TriggersApi.md#GetListForTarget) | **Get** /pipeline/api/triggers | Gets the paginated list of triggers for accountIdentifier, orgIdentifier, projectIdentifier, targetIdentifier.
[**GetTrigger**](TriggersApi.md#GetTrigger) | **Get** /pipeline/api/triggers/{triggerIdentifier} | Gets the trigger by accountIdentifier, orgIdentifier, projectIdentifier, targetIdentifier and triggerIdentifier.
[**GetTriggerCatalog**](TriggersApi.md#GetTriggerCatalog) | **Get** /pipeline/api/triggers/catalog | Lists all Triggers
[**GetTriggerDetails**](TriggersApi.md#GetTriggerDetails) | **Get** /pipeline/api/triggers/{triggerIdentifier}/details | Fetches Trigger details for a specific accountIdentifier, orgIdentifier, projectIdentifier, targetIdentifier, triggerIdentifier.
[**UpdateTrigger**](TriggersApi.md#UpdateTrigger) | **Put** /pipeline/api/triggers/{triggerIdentifier} | Updates trigger for pipeline with target pipeline identifier.

# **CreateTrigger**
> ResponseDtongTriggerResponse CreateTrigger(ctx, body, accountIdentifier, orgIdentifier, projectIdentifier, targetIdentifier, optional)
Creates Trigger for triggering target pipeline identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)|  | 
  **accountIdentifier** | **string**|  | 
  **orgIdentifier** | **string**|  | 
  **projectIdentifier** | **string**|  | 
  **targetIdentifier** | **string**| Identifier of the target pipeline | 
 **optional** | ***TriggersApiCreateTriggerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TriggersApiCreateTriggerOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **ignoreError** | **optional.**|  | [default to false]
 **withServiceV2** | **optional.**|  | [default to false]

### Return type

[**ResponseDtongTriggerResponse**](ResponseDTONGTriggerResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTrigger**
> ResponseDtoBoolean DeleteTrigger(ctx, accountIdentifier, orgIdentifier, projectIdentifier, targetIdentifier, triggerIdentifier, optional)
Deletes Trigger by identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**|  | 
  **orgIdentifier** | **string**|  | 
  **projectIdentifier** | **string**|  | 
  **targetIdentifier** | **string**| Identifier of the target pipeline under which trigger resides. | 
  **triggerIdentifier** | **string**|  | 
 **optional** | ***TriggersApiDeleteTriggerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TriggersApiDeleteTriggerOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **ifMatch** | **optional.String**|  | 

### Return type

[**ResponseDtoBoolean**](ResponseDTOBoolean.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListForTarget**
> ResponseDtoPageResponseNgTriggerDetailsResponseDto GetListForTarget(ctx, accountIdentifier, orgIdentifier, projectIdentifier, targetIdentifier, optional)
Gets the paginated list of triggers for accountIdentifier, orgIdentifier, projectIdentifier, targetIdentifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**|  | 
  **orgIdentifier** | **string**|  | 
  **projectIdentifier** | **string**|  | 
  **targetIdentifier** | **string**| Identifier of the target pipeline | 
 **optional** | ***TriggersApiGetListForTargetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TriggersApiGetListForTargetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **filter** | **optional.String**|  | 
 **page** | **optional.Int32**|  | [default to 0]
 **size** | **optional.Int32**|  | [default to 25]
 **sort** | [**optional.Interface of []string**](string.md)|  | 
 **searchTerm** | **optional.String**|  | 

### Return type

[**ResponseDtoPageResponseNgTriggerDetailsResponseDto**](ResponseDTOPageResponseNGTriggerDetailsResponseDTO.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTrigger**
> ResponseDtongTriggerResponse GetTrigger(ctx, accountIdentifier, orgIdentifier, projectIdentifier, targetIdentifier, triggerIdentifier)
Gets the trigger by accountIdentifier, orgIdentifier, projectIdentifier, targetIdentifier and triggerIdentifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**|  | 
  **orgIdentifier** | **string**|  | 
  **projectIdentifier** | **string**|  | 
  **targetIdentifier** | **string**| Identifier of the target pipeline under which trigger resides | 
  **triggerIdentifier** | **string**|  | 

### Return type

[**ResponseDtongTriggerResponse**](ResponseDTONGTriggerResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTriggerCatalog**
> ResponseDtoTriggerCatalogResponse GetTriggerCatalog(ctx, accountIdentifier)
Lists all Triggers

Lists all the Triggers for the given Account ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 

### Return type

[**ResponseDtoTriggerCatalogResponse**](ResponseDTOTriggerCatalogResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTriggerDetails**
> ResponseDtongTriggerDetailsResponseDto GetTriggerDetails(ctx, accountIdentifier, orgIdentifier, projectIdentifier, triggerIdentifier, targetIdentifier)
Fetches Trigger details for a specific accountIdentifier, orgIdentifier, projectIdentifier, targetIdentifier, triggerIdentifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**|  | 
  **orgIdentifier** | **string**|  | 
  **projectIdentifier** | **string**|  | 
  **triggerIdentifier** | **string**| Identifier of the target pipeline | 
  **targetIdentifier** | **string**|  | 

### Return type

[**ResponseDtongTriggerDetailsResponseDto**](ResponseDTONGTriggerDetailsResponseDTO.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTrigger**
> ResponseDtongTriggerResponse UpdateTrigger(ctx, body, accountIdentifier, orgIdentifier, projectIdentifier, targetIdentifier, triggerIdentifier, optional)
Updates trigger for pipeline with target pipeline identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)|  | 
  **accountIdentifier** | **string**|  | 
  **orgIdentifier** | **string**|  | 
  **projectIdentifier** | **string**|  | 
  **targetIdentifier** | **string**| Identifier of the target pipeline under which trigger resides | 
  **triggerIdentifier** | **string**|  | 
 **optional** | ***TriggersApiUpdateTriggerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TriggersApiUpdateTriggerOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **ifMatch** | **optional.**|  | 
 **ignoreError** | **optional.**|  | [default to false]

### Return type

[**ResponseDtongTriggerResponse**](ResponseDTONGTriggerResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

