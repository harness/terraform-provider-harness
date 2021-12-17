# {{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddDelegateConfigurationForAccount**](DelegateConfigurationResourceApi.md#AddDelegateConfigurationForAccount) | **Post** /ng/api/v2/accounts/{accountId}/delegate-configs | Creates Delegate Configuration specified by config details for specified account
[**CreateDelegateConfigurationV2**](DelegateConfigurationResourceApi.md#CreateDelegateConfigurationV2) | **Post** /ng/api/v2/delegate-configs | Creates Delegate Configuration specified by config details
[**DeleteDelegateConfigV2**](DelegateConfigurationResourceApi.md#DeleteDelegateConfigV2) | **Delete** /ng/api/v2/accounts/{accountId}/delegate-configs/{delegateConfigIdentifier} | Deletes Delegate Configuration specified by identifier
[**GetDelegateConfigrationDetailsV2**](DelegateConfigurationResourceApi.md#GetDelegateConfigrationDetailsV2) | **Get** /ng/api/v2/accounts/{accountId}/delegate-configs/{delegateConfigIdentifier} | Retrieves Delegate Configuration details for given Delegate Configuration identifier.
[**GetDelegateConfigurationsForAccountV2**](DelegateConfigurationResourceApi.md#GetDelegateConfigurationsForAccountV2) | **Get** /ng/api/v2/accounts/{accountId}/delegate-configs | Lists Delegate Configuration for specified account, org and project
[**GetDelegateConfigurationsWithFiltering**](DelegateConfigurationResourceApi.md#GetDelegateConfigurationsWithFiltering) | **Post** /ng/api/v2/accounts/{accountId}/delegate-configs/listV2 | Lists Delegate Configuration for specified account, org and project and filter applied
[**UpdateDelegateConfigurationV2**](DelegateConfigurationResourceApi.md#UpdateDelegateConfigurationV2) | **Put** /ng/api/v2/accounts/{accountId}/delegate-configs/{delegateConfigIdentifier} | Updates Delegate Configuration specified by Identifier
[**UpdateDelegateSelectorsV2**](DelegateConfigurationResourceApi.md#UpdateDelegateSelectorsV2) | **Put** /ng/api/v2/accounts/{accountId}/delegate-configs/{delegateConfigIdentifier}/selectors | Updates Delegate selectors for Delegate Configuration specified by identifier
[**UpdateScopingRulesV2**](DelegateConfigurationResourceApi.md#UpdateScopingRulesV2) | **Put** /ng/api/v2/accounts/{accountId}/delegate-configs/{delegateConfigIdentifier}/scoping-rules | Updates Scoping Rules for the Delegate Configuration specified by identifier

# **AddDelegateConfigurationForAccount**
> RestResponseDelegateProfileDetailsNg AddDelegateConfigurationForAccount(ctx, body, accountId)
Creates Delegate Configuration specified by config details for specified account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DelegateProfileDetailsNg**](DelegateProfileDetailsNg.md)| Delegate Configuration to be created. These include uuid, identifier, accountId, orgId, projId, name, startupScript, scopingRules, selectors... | 
  **accountId** | **string**| Account id | 

### Return type

[**RestResponseDelegateProfileDetailsNg**](RestResponseDelegateProfileDetailsNg.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateDelegateConfigurationV2**
> RestResponseDelegateProfileDetailsNg CreateDelegateConfigurationV2(ctx, body)
Creates Delegate Configuration specified by config details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DelegateProfileDetailsNg**](DelegateProfileDetailsNg.md)| Delegate Configuration to be created. These include uuid, identifier, accountId, orgId, projId, name, startupScript, scopingRules, selectors... | 

### Return type

[**RestResponseDelegateProfileDetailsNg**](RestResponseDelegateProfileDetailsNg.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDelegateConfigV2**
> ResponseDtoBoolean DeleteDelegateConfigV2(ctx, delegateConfigIdentifier, accountId, optional)
Deletes Delegate Configuration specified by identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **delegateConfigIdentifier** | **string**| Delegate Configuration identifier | 
  **accountId** | **string**| Account id | 
 **optional** | ***DelegateConfigurationResourceApiDeleteDelegateConfigV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DelegateConfigurationResourceApiDeleteDelegateConfigV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgId** | **optional.String**| Organization Id | 
 **projectId** | **optional.String**| Project Id | 

### Return type

[**ResponseDtoBoolean**](ResponseDTOBoolean.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDelegateConfigrationDetailsV2**
> RestResponseDelegateProfileDetailsNg GetDelegateConfigrationDetailsV2(ctx, delegateConfigIdentifier, accountId, optional)
Retrieves Delegate Configuration details for given Delegate Configuration identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **delegateConfigIdentifier** | **string**| Delegate Configuration identifier | 
  **accountId** | **string**| Account id | 
 **optional** | ***DelegateConfigurationResourceApiGetDelegateConfigrationDetailsV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DelegateConfigurationResourceApiGetDelegateConfigrationDetailsV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

# **GetDelegateConfigurationsForAccountV2**
> RestResponsePageResponseDelegateProfileDetailsNg GetDelegateConfigurationsForAccountV2(ctx, accountId, optional)
Lists Delegate Configuration for specified account, org and project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Account id | 
 **optional** | ***DelegateConfigurationResourceApiGetDelegateConfigurationsForAccountV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DelegateConfigurationResourceApiGetDelegateConfigurationsForAccountV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **optional.String**|  | [default to 0]
 **limit** | **optional.String**|  | 
 **fieldsIncluded** | [**optional.Interface of []string**](string.md)|  | 
 **fieldsExcluded** | [**optional.Interface of []string**](string.md)|  | 
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

# **GetDelegateConfigurationsWithFiltering**
> RestResponsePageResponseDelegateProfileDetailsNg GetDelegateConfigurationsWithFiltering(ctx, accountId, optional)
Lists Delegate Configuration for specified account, org and project and filter applied

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Account id | 
 **optional** | ***DelegateConfigurationResourceApiGetDelegateConfigurationsWithFilteringOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DelegateConfigurationResourceApiGetDelegateConfigurationsWithFilteringOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of DelegateProfileFilterPropertiesDto**](DelegateProfileFilterPropertiesDto.md)| Delegate Configuration filter properties: name, identifier, description, approvalRequired, list of selectors  | 
 **orgId** | **optional.**| Organization Id | 
 **projectId** | **optional.**| Project Id | 
 **filterIdentifier** | **optional.**| Filter identifier | 
 **searchTerm** | **optional.**| Search term | 
 **offset** | **optional.**|  | [default to 0]
 **limit** | **optional.**|  | 
 **fieldsIncluded** | [**optional.Interface of []string**](string.md)|  | 
 **fieldsExcluded** | [**optional.Interface of []string**](string.md)|  | 

### Return type

[**RestResponsePageResponseDelegateProfileDetailsNg**](RestResponsePageResponseDelegateProfileDetailsNg.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDelegateConfigurationV2**
> RestResponseDelegateProfileDetailsNg UpdateDelegateConfigurationV2(ctx, body, delegateConfigIdentifier, accountId, optional)
Updates Delegate Configuration specified by Identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DelegateProfileDetailsNg**](DelegateProfileDetailsNg.md)| Delegate configuration details to be updated. These include name, startupScript, scopingRules, selectors | 
  **delegateConfigIdentifier** | **string**| Delegate Configuration identifier | 
  **accountId** | **string**| Account id | 
 **optional** | ***DelegateConfigurationResourceApiUpdateDelegateConfigurationV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DelegateConfigurationResourceApiUpdateDelegateConfigurationV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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

# **UpdateDelegateSelectorsV2**
> RestResponseDelegateProfileDetailsNg UpdateDelegateSelectorsV2(ctx, delegateConfigIdentifier, accountId, optional)
Updates Delegate selectors for Delegate Configuration specified by identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **delegateConfigIdentifier** | **string**| Delegate Configuration identifier | 
  **accountId** | **string**| Account id | 
 **optional** | ***DelegateConfigurationResourceApiUpdateDelegateSelectorsV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DelegateConfigurationResourceApiUpdateDelegateSelectorsV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of []string**](string.md)| List of Delegate selectors to be updated | 
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

# **UpdateScopingRulesV2**
> RestResponseDelegateProfileDetailsNg UpdateScopingRulesV2(ctx, body, delegateConfigIdentifier, accountId, optional)
Updates Scoping Rules for the Delegate Configuration specified by identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]ScopingRuleDetailsNg**](ScopingRuleDetailsNg.md)| List of Delegate Scoping Rules to be updated | 
  **delegateConfigIdentifier** | **string**| Delegate Configuration identifier | 
  **accountId** | **string**| Account id | 
 **optional** | ***DelegateConfigurationResourceApiUpdateScopingRulesV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DelegateConfigurationResourceApiUpdateScopingRulesV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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

