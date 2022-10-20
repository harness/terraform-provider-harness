# nextgen{{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AgentCertificateServiceCreate**](RepositoryCertificatesApi.md#AgentCertificateServiceCreate) | **Post** /gitops/api/api/v1/agents/{agentIdentifier}/certificates | Creates repository certificates on the server
[**AgentCertificateServiceDelete**](RepositoryCertificatesApi.md#AgentCertificateServiceDelete) | **Delete** /gitops/api/api/v1/agents/{agentIdentifier}/certificates | Delete the certificates that match the RepositoryCertificateQuery
[**AgentCertificateServiceList**](RepositoryCertificatesApi.md#AgentCertificateServiceList) | **Get** /gitops/api/api/v1/agents/{agentIdentifier}/certificates | List all available repository certificates

# **AgentCertificateServiceCreate**
> CertificatesRepositoryCertificateList AgentCertificateServiceCreate(ctx, body, agentIdentifier, optional)
Creates repository certificates on the server

Create repository certificates.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CertificateRepositoryCertificateCreateRequest**](CertificateRepositoryCertificateCreateRequest.md)|  | 
  **agentIdentifier** | **string**| Agent identifier for entity. | 
 **optional** | ***RepositoryCertificatesApiAgentCertificateServiceCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryCertificatesApiAgentCertificateServiceCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accountIdentifier** | **optional.**| Account Identifier for the Entity. | 
 **orgIdentifier** | **optional.**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity. | 

### Return type

[**CertificatesRepositoryCertificateList**](certificatesRepositoryCertificateList.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentCertificateServiceDelete**
> CertificatesRepositoryCertificateList AgentCertificateServiceDelete(ctx, agentIdentifier, optional)
Delete the certificates that match the RepositoryCertificateQuery

Delete repository certificates.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentifier** | **string**| Agent identifier for entity. | 
 **optional** | ***RepositoryCertificatesApiAgentCertificateServiceDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryCertificatesApiAgentCertificateServiceDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.String**| Account Identifier for the Entity. | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**|  | 
 **queryHostNamePattern** | **optional.String**| A file-glob pattern (not regular expression) the host name has to match. | 
 **queryCertType** | **optional.String**| The type of the certificate to match (ssh or https). | 
 **queryCertSubType** | **optional.String**| The sub type of the certificate to match (protocol dependent, usually only used for ssh certs). | 

### Return type

[**CertificatesRepositoryCertificateList**](certificatesRepositoryCertificateList.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentCertificateServiceList**
> CertificatesRepositoryCertificateList AgentCertificateServiceList(ctx, agentIdentifier, accountIdentifier, optional)
List all available repository certificates

List repository certificates.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentifier** | **string**| Agent identifier for entity. | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***RepositoryCertificatesApiAgentCertificateServiceListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryCertificatesApiAgentCertificateServiceListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**|  | 
 **queryHostNamePattern** | **optional.String**| A file-glob pattern (not regular expression) the host name has to match. | 
 **queryCertType** | **optional.String**| The type of the certificate to match (ssh or https). | 
 **queryCertSubType** | **optional.String**| The sub type of the certificate to match (protocol dependent, usually only used for ssh certs). | 

### Return type

[**CertificatesRepositoryCertificateList**](certificatesRepositoryCertificateList.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

