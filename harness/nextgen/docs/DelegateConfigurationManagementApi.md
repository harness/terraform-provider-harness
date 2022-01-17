# {{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDelegateConfiguration**](DelegateConfigurationManagementApi.md#CreateDelegateConfiguration) | **Post** /ng/api/delegate-profiles/ng | Creates Delegate Configuration specified by Configuration details in body
[**DeleteDelegateConfig**](DelegateConfigurationManagementApi.md#DeleteDelegateConfig) | **Delete** /ng/api/delegate-profiles/ng/{delegateProfileId} | Deletes Delegate Configuration specified by Id
[**GetDelegateConfigrationDetails**](DelegateConfigurationManagementApi.md#GetDelegateConfigrationDetails) | **Get** /ng/api/delegate-profiles/ng/{delegateProfileId} | Retrieves Delegate Configuration details for given Delegate Configuration Id.
[**GetDelegateConfigurationsForAccount**](DelegateConfigurationManagementApi.md#GetDelegateConfigurationsForAccount) | **Get** /ng/api/delegate-profiles/ng | Lists Delegate Configuration for specified Account, Organization and Project
[**UpdateDelegateConfiguration**](DelegateConfigurationManagementApi.md#UpdateDelegateConfiguration) | **Put** /ng/api/delegate-profiles/ng/{delegateProfileId} | Updates Delegate Configuration specified by Id
[**UpdateDelegateSelectors**](DelegateConfigurationManagementApi.md#UpdateDelegateSelectors) | **Put** /ng/api/delegate-profiles/ng/{delegateProfileId}/selectors | Updates Delegate Selectors for Delegate Configuration specified by Id
[**UpdateScopingRules**](DelegateConfigurationManagementApi.md#UpdateScopingRules) | **Put** /ng/api/delegate-profiles/ng/{delegateProfileId}/scoping-rules | Updates Scoping Rules for the Delegate Configuration specified by Id

# **CreateDelegateConfiguration**
> RestResponseDelegateProfileDetailsNg CreateDelegateConfiguration(ctx, body, optional)
Creates Delegate Configuration specified by Configuration details in body

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DelegateProfileDetailsNg**](DelegateProfileDetailsNg.md)| Delegate Configuration to be created. These include uuid, identifier, accountId, orgId, projId, name, startupScript, scopingRules, selectors... | 
 **optional** | ***DelegateConfigurationManagementApiCreateDelegateConfigurationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DelegateConfigurationManagementApiCreateDelegateConfigurationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountId** | **optional.**| Account id | 
 **orgId** | **optional.**| Organization Id | 
 **projectId** | **optional.**| Project Id | 

### Return type

[**RestResponseDelegateProfileDetailsNg**](RestResponseDelegateProfileDetailsNg.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDelegateConfig**
> RestResponseVoid DeleteDelegateConfig(ctx, delegateProfileId, optional)
Deletes Delegate Configuration specified by Id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **delegateProfileId** | **string**| Delegate Configuration Id | 
 **optional** | ***DelegateConfigurationManagementApiDeleteDelegateConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DelegateConfigurationManagementApiDeleteDelegateConfigOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountId** | **optional.String**| Account id | 
 **orgId** | **optional.String**| Organization Id | 
 **projectId** | **optional.String**| Project Id | 

### Return type

[**RestResponseVoid**](RestResponseVoid.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDelegateConfigrationDetails**
> RestResponseDelegateProfileDetailsNg GetDelegateConfigrationDetails(ctx, delegateProfileId, optional)
Retrieves Delegate Configuration details for given Delegate Configuration Id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **delegateProfileId** | **string**| Delegate Configuration Id | 
 **optional** | ***DelegateConfigurationManagementApiGetDelegateConfigrationDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DelegateConfigurationManagementApiGetDelegateConfigrationDetailsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountId** | **optional.String**| Account id | 
 **orgId** | **optional.String**| Organization Id | 
 **projectId** | **optional.String**| Project Id | 

### Return type

[**RestResponseDelegateProfileDetailsNg**](RestResponseDelegateProfileDetailsNg.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDelegateConfigurationsForAccount**
> RestResponsePageResponseDelegateProfileDetailsNg GetDelegateConfigurationsForAccount(ctx, optional)
Lists Delegate Configuration for specified Account, Organization and Project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DelegateConfigurationManagementApiGetDelegateConfigurationsForAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DelegateConfigurationManagementApiGetDelegateConfigurationsForAccountOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **optional.String**|  | [default to 0]
 **limit** | **optional.String**|  | 
 **fieldsIncluded** | [**optional.Interface of []string**](string.md)|  | 
 **fieldsExcluded** | [**optional.Interface of []string**](string.md)|  | 
 **accountId** | **optional.String**| Account id | 
 **orgId** | **optional.String**| Organization Id | 
 **projectId** | **optional.String**| Project Id | 

### Return type

[**RestResponsePageResponseDelegateProfileDetailsNg**](RestResponsePageResponseDelegateProfileDetailsNg.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDelegateConfiguration**
> RestResponseDelegateProfileDetailsNg UpdateDelegateConfiguration(ctx, body, delegateProfileId, optional)
Updates Delegate Configuration specified by Id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DelegateProfileDetailsNg**](DelegateProfileDetailsNg.md)| Delegate Configuration to be created. These include uuid, identifier, accountId, orgId, projId, name, startupScript, scopingRules, selectors... | 
  **delegateProfileId** | **string**| Delegate Configuration Id | 
 **optional** | ***DelegateConfigurationManagementApiUpdateDelegateConfigurationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DelegateConfigurationManagementApiUpdateDelegateConfigurationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accountId** | **optional.**| Account id | 
 **orgId** | **optional.**| Organization Id | 
 **projectId** | **optional.**| Project Id | 

### Return type

[**RestResponseDelegateProfileDetailsNg**](RestResponseDelegateProfileDetailsNg.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDelegateSelectors**
> RestResponseDelegateProfileDetailsNg UpdateDelegateSelectors(ctx, delegateProfileId, optional)
Updates Delegate Selectors for Delegate Configuration specified by Id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **delegateProfileId** | **string**| Delegate Configuration Id | 
 **optional** | ***DelegateConfigurationManagementApiUpdateDelegateSelectorsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DelegateConfigurationManagementApiUpdateDelegateSelectorsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of []string**](string.md)| Delegate Selectors to be updated | 
 **accountId** | **optional.**| Account Id | 
 **orgId** | **optional.**| Organization Id | 
 **projectId** | **optional.**| Project Id | 

### Return type

[**RestResponseDelegateProfileDetailsNg**](RestResponseDelegateProfileDetailsNg.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateScopingRules**
> RestResponseDelegateProfileDetailsNg UpdateScopingRules(ctx, delegateProfileId, optional)
Updates Scoping Rules for the Delegate Configuration specified by Id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **delegateProfileId** | **string**| Delegate Configuration Id | 
 **optional** | ***DelegateConfigurationManagementApiUpdateScopingRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DelegateConfigurationManagementApiUpdateScopingRulesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of []ScopingRuleDetailsNg**](ScopingRuleDetailsNg.md)| Delegate Scoping Rules to be updated | 
 **accountId** | **optional.**| Account id | 
 **orgId** | **optional.**| Organization Id | 
 **projectId** | **optional.**| Project Id | 

### Return type

[**RestResponseDelegateProfileDetailsNg**](RestResponseDelegateProfileDetailsNg.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

