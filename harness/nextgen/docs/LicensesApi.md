# {{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExtendTrialLicense**](LicensesApi.md#ExtendTrialLicense) | **Post** /ng/api/licenses/extend-trial | Extends Trial License For A Module
[**GetAccountLicenses**](LicensesApi.md#GetAccountLicenses) | **Get** /ng/api/licenses/account | Gets All Module License Information in Account
[**GetEditionActions**](LicensesApi.md#GetEditionActions) | **Get** /ng/api/licenses/actions | Get Allowed Actions Under Each Edition
[**GetLastModifiedTimeForAllModuleTypes**](LicensesApi.md#GetLastModifiedTimeForAllModuleTypes) | **Post** /ng/api/licenses/versions | Get Last Modified Time Under Each ModuleType
[**GetLicensesAndSummary**](LicensesApi.md#GetLicensesAndSummary) | **Get** /ng/api/licenses/{accountIdentifier}/summary | Gets Module Licenses With Summary By Account And ModuleType
[**GetModuleLicenseById**](LicensesApi.md#GetModuleLicenseById) | **Get** /ng/api/licenses/{identifier} | Gets Module License
[**GetModuleLicensesByAccountAndModuleType**](LicensesApi.md#GetModuleLicensesByAccountAndModuleType) | **Get** /ng/api/licenses/modules/{accountIdentifier} | Gets Module Licenses By Account And ModuleType
[**StartFreeLicense**](LicensesApi.md#StartFreeLicense) | **Post** /ng/api/licenses/free | Starts Free License For A Module
[**StartTrialLicense**](LicensesApi.md#StartTrialLicense) | **Post** /ng/api/licenses/trial | Starts Trial License For A Module

# **ExtendTrialLicense**
> ResponseDtoModuleLicense ExtendTrialLicense(ctx, body, accountIdentifier)
Extends Trial License For A Module

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**StartTrial**](StartTrial.md)| This is the details of the Trial License. ModuleType and edition are mandatory | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 

### Return type

[**ResponseDtoModuleLicense**](ResponseDTOModuleLicense.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountLicenses**
> ResponseDtoAccountLicense GetAccountLicenses(ctx, accountIdentifier)
Gets All Module License Information in Account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity | 

### Return type

[**ResponseDtoAccountLicense**](ResponseDTOAccountLicense.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEditionActions**
> ResponseDtoMapEditionSetEditionAction GetEditionActions(ctx, accountIdentifier, moduleType)
Get Allowed Actions Under Each Edition

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
  **moduleType** | **string**| A Harness Platform module. | 

### Return type

[**ResponseDtoMapEditionSetEditionAction**](ResponseDTOMapEditionSetEditionAction.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLastModifiedTimeForAllModuleTypes**
> ResponseDtoMapModuleTypeLong GetLastModifiedTimeForAllModuleTypes(ctx, accountIdentifier)
Get Last Modified Time Under Each ModuleType

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity | 

### Return type

[**ResponseDtoMapModuleTypeLong**](ResponseDTOMapModuleTypeLong.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLicensesAndSummary**
> ResponseDtoLicensesWithSummary GetLicensesAndSummary(ctx, accountIdentifier, moduleType)
Gets Module Licenses With Summary By Account And ModuleType

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
  **moduleType** | **string**| A Harness Platform module. | 

### Return type

[**ResponseDtoLicensesWithSummary**](ResponseDTOLicensesWithSummary.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetModuleLicenseById**
> ResponseDtoModuleLicense GetModuleLicenseById(ctx, identifier, accountIdentifier)
Gets Module License

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| The module license identifier | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 

### Return type

[**ResponseDtoModuleLicense**](ResponseDTOModuleLicense.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetModuleLicensesByAccountAndModuleType**
> ResponseDtoListModuleLicense GetModuleLicensesByAccountAndModuleType(ctx, accountIdentifier, moduleType)
Gets Module Licenses By Account And ModuleType

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
  **moduleType** | **string**| A Harness Platform module. | 

### Return type

[**ResponseDtoListModuleLicense**](ResponseDTOListModuleLicense.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StartFreeLicense**
> ResponseDtoModuleLicense StartFreeLicense(ctx, accountIdentifier, moduleType)
Starts Free License For A Module

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
  **moduleType** | **string**| A Harness Platform module. | 

### Return type

[**ResponseDtoModuleLicense**](ResponseDTOModuleLicense.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StartTrialLicense**
> ResponseDtoModuleLicense StartTrialLicense(ctx, body, accountIdentifier)
Starts Trial License For A Module

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**StartTrial**](StartTrial.md)| This is the details of the Trial License. ModuleType and edition are mandatory | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 

### Return type

[**ResponseDtoModuleLicense**](ResponseDTOModuleLicense.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

