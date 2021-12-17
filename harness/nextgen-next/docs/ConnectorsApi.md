# {{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConnector**](ConnectorsApi.md#CreateConnector) | **Post** /ng/api/connectors | Creates a Connector
[**DeleteConnector**](ConnectorsApi.md#DeleteConnector) | **Delete** /ng/api/connectors/{identifier} | Deletes Connector by identifier
[**GetAllAllowedFieldValues**](ConnectorsApi.md#GetAllAllowedFieldValues) | **Get** /ng/api/connectors/fieldValues | Get the allowed field values by Connector Type
[**GetCEAwsTemplate**](ConnectorsApi.md#GetCEAwsTemplate) | **Post** /ng/api/connectors/getceawstemplateurl | Get the Template URL of connector
[**GetConnector**](ConnectorsApi.md#GetConnector) | **Get** /ng/api/connectors/{identifier} | get the Connector by accountIdentifier and connectorIdentifier
[**GetConnectorCatalogue**](ConnectorsApi.md#GetConnectorCatalogue) | **Get** /ng/api/connectors/catalogue | gets the connector catalogue by accountIdentifier
[**GetConnectorList**](ConnectorsApi.md#GetConnectorList) | **Get** /ng/api/connectors | Get the list of Connectors satisfying the criteria (if any) in the request
[**GetConnectorListV2**](ConnectorsApi.md#GetConnectorListV2) | **Post** /ng/api/connectors/listV2 | Get the list of Connectors satisfying the criteria (if any) in the request
[**GetConnectorStatistics**](ConnectorsApi.md#GetConnectorStatistics) | **Get** /ng/api/connectors/stats | gets the connector&#x27;s statistics by accountIdentifier, projectIdentifier and orgIdentifier
[**GetConnectorValidationParams**](ConnectorsApi.md#GetConnectorValidationParams) | **Get** /ng/api/connectors/{identifier}/validation-params | 
[**GetTestConnectionResult**](ConnectorsApi.md#GetTestConnectionResult) | **Post** /ng/api/connectors/testConnection/{identifier} | Tests the connection of the Connector by Identifier
[**GetTestConnectionResultInternal**](ConnectorsApi.md#GetTestConnectionResultInternal) | **Post** /ng/api/connectors/testConnectionInternal/{identifier} | Tests the connection of the connector by Identifier
[**GetTestGitRepoConnectionResult**](ConnectorsApi.md#GetTestGitRepoConnectionResult) | **Post** /ng/api/connectors/testGitRepoConnection/{identifier} | Tests the created Connector&#x27;s connection
[**ListConnectorByFQN**](ConnectorsApi.md#ListConnectorByFQN) | **Post** /ng/api/connectors/listbyfqn | Get the list of connectors by FQN satisfying the criteria (if any) in the request
[**PutConnector**](ConnectorsApi.md#PutConnector) | **Put** /ng/api/connectors | Updates the Connector
[**ValidateTheIdentifierIsUnique**](ConnectorsApi.md#ValidateTheIdentifierIsUnique) | **Get** /ng/api/connectors/validateUniqueIdentifier | validate the Connector by accountIdentifier and connectorIdentifier

# **CreateConnector**
> ResponseDtoConnectorResponse CreateConnector(ctx, body, optional)
Creates a Connector

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Connector**](Connector.md)| Details of the Connector to create | 
 **optional** | ***ConnectorsApiCreateConnectorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConnectorsApiCreateConnectorOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.**| Account Identifier for the entity | 
 **branch** | **optional.**| Branch Name | 
 **repoIdentifier** | **optional.**| Git Sync Config Identifier | 
 **rootFolder** | **optional.**| Default Folder Path | 
 **filePath** | **optional.**| File Path | 
 **commitMsg** | **optional.**| File Path | 
 **isNewBranch** | **optional.**| Checks the new branch | [default to false]
 **baseBranch** | **optional.**| Default Branch | 

### Return type

[**ResponseDtoConnectorResponse**](ResponseDTOConnectorResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, text/yaml, text/html, text/plain
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteConnector**
> ResponseDtoBoolean DeleteConnector(ctx, accountIdentifier, identifier, optional)
Deletes Connector by identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the entity | 
  **identifier** | **string**| Connector Identifier | 
 **optional** | ***ConnectorsApiDeleteConnectorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConnectorsApiDeleteConnectorOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.String**| Organization Identifier for the entity | 
 **projectIdentifier** | **optional.String**| Project Identifier for the entity | 
 **branch** | **optional.String**| Branch Name | 
 **repoIdentifier** | **optional.String**| Git Sync Config Identifier | 
 **rootFolder** | **optional.String**| Default Folder Path | 
 **filePath** | **optional.String**| File Path | 
 **commitMsg** | **optional.String**| Commit Message | 
 **lastObjectId** | **optional.String**| Last Object Id | 

### Return type

[**ResponseDtoBoolean**](ResponseDTOBoolean.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllAllowedFieldValues**
> ResponseDtoFieldValues GetAllAllowedFieldValues(ctx, connectorType)
Get the allowed field values by Connector Type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **connectorType** | **string**| Connector type | 

### Return type

[**ResponseDtoFieldValues**](ResponseDTOFieldValues.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCEAwsTemplate**
> ResponseDtoString GetCEAwsTemplate(ctx, optional)
Get the Template URL of connector

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConnectorsApiGetCEAwsTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConnectorsApiGetCEAwsTemplateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventsEnabled** | **optional.Bool**| Specify whether or not to enable events | 
 **curEnabled** | **optional.Bool**| Specify whether or not to enable CUR | 
 **optimizationEnabled** | **optional.Bool**| Specify whether or not to enable optimization | 

### Return type

[**ResponseDtoString**](ResponseDTOString.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConnector**
> ResponseDtoConnectorResponse GetConnector(ctx, identifier, optional)
get the Connector by accountIdentifier and connectorIdentifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| Connector Identifier | 
 **optional** | ***ConnectorsApiGetConnectorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConnectorsApiGetConnectorOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.String**| Account Identifier for the entity | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the entity | 
 **projectIdentifier** | **optional.String**| Project Identifier for the entity | 
 **branch** | **optional.String**| Branch Name | 
 **repoIdentifier** | **optional.String**| Git Sync Config Identifier | 
 **getDefaultFromOtherRepo** | **optional.Bool**| if true, return all the default entities | 

### Return type

[**ResponseDtoConnectorResponse**](ResponseDTOConnectorResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConnectorCatalogue**
> ResponseDtoConnectorCatalogueRespone GetConnectorCatalogue(ctx, optional)
gets the connector catalogue by accountIdentifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConnectorsApiGetConnectorCatalogueOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConnectorsApiGetConnectorCatalogueOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountIdentifier** | **optional.String**| Account Identifier for the entity | 

### Return type

[**ResponseDtoConnectorCatalogueRespone**](ResponseDTOConnectorCatalogueRespone.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConnectorList**
> ResponseDtoPageResponseConnectorResponse GetConnectorList(ctx, optional)
Get the list of Connectors satisfying the criteria (if any) in the request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConnectorsApiGetConnectorListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConnectorsApiGetConnectorListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageIndex** | **optional.Int32**| Page number of navigation. If left empty, default value of 0 is assumed | [default to 0]
 **pageSize** | **optional.Int32**| Number of entries per page. If left empty, default value of 100 is assumed  | [default to 100]
 **accountIdentifier** | **optional.String**| Account Identifier for the entity | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the entity | 
 **projectIdentifier** | **optional.String**| Project Identifier for the entity | 
 **searchTerm** | **optional.String**| Filter Connectors by searching for this word in Name, Id, and Tag | 
 **type_** | **optional.String**| Filter Connectors by type | 
 **category** | **optional.String**| Filter Connectors by category | 
 **sourceCategory** | **optional.String**| Filter Connectors by Source Category. Available Source Categories are CLOUD_PROVIDER, SECRET_MANAGER, CLOUD_COST, ARTIFACTORY, CODE_REPO,  MONITORING and TICKETING | 
 **branch** | **optional.String**| Branch Name | 
 **repoIdentifier** | **optional.String**| Git Sync Config Identifier | 
 **getDefaultFromOtherRepo** | **optional.Bool**| if true, return all the default entities | 

### Return type

[**ResponseDtoPageResponseConnectorResponse**](ResponseDTOPageResponseConnectorResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConnectorListV2**
> ResponseDtoPageResponseConnectorResponse GetConnectorListV2(ctx, body, optional)
Get the list of Connectors satisfying the criteria (if any) in the request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConnectorFilterProperties**](ConnectorFilterProperties.md)| Details of the filters applied | 
 **optional** | ***ConnectorsApiGetConnectorListV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConnectorsApiGetConnectorListV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageIndex** | **optional.**| Page number of navigation. If left empty, default value of 0 is assumed | [default to 0]
 **pageSize** | **optional.**| Number of entries per page. If left empty, default value of 100 is assumed | [default to 100]
 **accountIdentifier** | **optional.**| Account Identifier for the entity | 
 **searchTerm** | **optional.**| Filter Connectors based on this word in Connectors name, id and tag | 
 **orgIdentifier** | **optional.**| Organization Identifier for the entity | 
 **projectIdentifier** | **optional.**| Project Identifier for the entity | 
 **filterIdentifier** | **optional.**|  | 
 **includeAllConnectorsAvailableAtScope** | **optional.**| Specify whether or not to include all the Connectors accessible at the scope. For eg if set as true, at the Project scope we will get org and account Connector also in the response | 
 **branch** | **optional.**| Branch Name | 
 **repoIdentifier** | **optional.**| Git Sync Config Identifier | 
 **getDefaultFromOtherRepo** | **optional.**| if true, return all the default entities | 
 **getDistinctFromBranches** | **optional.**| This when set to true along with GitSync enabled for the Connector, you can to get other Connectors too which are not from same repo - branch but different repo&#x27;s default branch | 

### Return type

[**ResponseDtoPageResponseConnectorResponse**](ResponseDTOPageResponseConnectorResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, text/yaml, text/html, text/plain
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConnectorStatistics**
> ResponseDtoConnectorStatistics GetConnectorStatistics(ctx, optional)
gets the connector's statistics by accountIdentifier, projectIdentifier and orgIdentifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConnectorsApiGetConnectorStatisticsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConnectorsApiGetConnectorStatisticsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountIdentifier** | **optional.String**| Account Identifier for the entity | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the entity | 
 **projectIdentifier** | **optional.String**| Project Identifier for the entity | 
 **branch** | **optional.String**| Branch Name | 
 **repoIdentifier** | **optional.String**| Git Sync Config Identifier | 
 **getDefaultFromOtherRepo** | **optional.Bool**| if true, return all the default entities | 

### Return type

[**ResponseDtoConnectorStatistics**](ResponseDTOConnectorStatistics.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConnectorValidationParams**
> GetConnectorValidationParams(ctx, identifier, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| Connector Identifier | 
 **optional** | ***ConnectorsApiGetConnectorValidationParamsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConnectorsApiGetConnectorValidationParamsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.String**| Account Identifier for the entity | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the entity | 
 **projectIdentifier** | **optional.String**| Project Identifier for the entity | 

### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTestConnectionResult**
> ResponseDtoConnectorValidationResult GetTestConnectionResult(ctx, identifier, optional)
Tests the connection of the Connector by Identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| Connector Identifier | 
 **optional** | ***ConnectorsApiGetTestConnectionResultOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConnectorsApiGetTestConnectionResultOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.String**| Account Identifier for the entity | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the entity | 
 **projectIdentifier** | **optional.String**| Project Identifier for the entity | 
 **branch** | **optional.String**| Branch Name | 
 **repoIdentifier** | **optional.String**| Git Sync Config Identifier | 
 **getDefaultFromOtherRepo** | **optional.Bool**| if true, return all the default entities | 

### Return type

[**ResponseDtoConnectorValidationResult**](ResponseDTOConnectorValidationResult.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTestConnectionResultInternal**
> ResponseDtoConnectorValidationResult GetTestConnectionResultInternal(ctx, identifier, optional)
Tests the connection of the connector by Identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**|  | 
 **optional** | ***ConnectorsApiGetTestConnectionResultInternalOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConnectorsApiGetTestConnectionResultInternalOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.String**|  | 
 **orgIdentifier** | **optional.String**|  | 
 **projectIdentifier** | **optional.String**|  | 

### Return type

[**ResponseDtoConnectorValidationResult**](ResponseDTOConnectorValidationResult.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTestGitRepoConnectionResult**
> ResponseDtoConnectorValidationResult GetTestGitRepoConnectionResult(ctx, identifier, optional)
Tests the created Connector's connection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| Connector Identifier | 
 **optional** | ***ConnectorsApiGetTestGitRepoConnectionResultOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConnectorsApiGetTestGitRepoConnectionResultOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.String**| Account Identifier for the entity | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the entity | 
 **projectIdentifier** | **optional.String**| Project Identifier for the entity | 
 **repoURL** | **optional.String**| URL of the repository, specify only in the case of Account Type Git Connector | 

### Return type

[**ResponseDtoConnectorValidationResult**](ResponseDTOConnectorValidationResult.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListConnectorByFQN**
> ResponseDtoListConnectorResponse ListConnectorByFQN(ctx, body, optional)
Get the list of connectors by FQN satisfying the criteria (if any) in the request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]string**](string.md)| List of ConnectorsFQN as strings | 
 **optional** | ***ConnectorsApiListConnectorByFQNOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConnectorsApiListConnectorByFQNOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.**| Account Identifier for the entity | 

### Return type

[**ResponseDtoListConnectorResponse**](ResponseDTOListConnectorResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, text/yaml, text/html, text/plain
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutConnector**
> ResponseDtoConnectorResponse PutConnector(ctx, body, optional)
Updates the Connector

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Connector**](Connector.md)| This is the updated Connector. Please provide values for all fields, not just the fields you are updating | 
 **optional** | ***ConnectorsApiPutConnectorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConnectorsApiPutConnectorOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.**| Account Identifier for the entity | 
 **branch** | **optional.**| Branch Name | 
 **repoIdentifier** | **optional.**| Git Sync Config Identifier | 
 **rootFolder** | **optional.**| Default Folder Path | 
 **filePath** | **optional.**| Default Folder Path | 
 **commitMsg** | **optional.**| Commit Message | 
 **lastObjectId** | **optional.**| Last Object Id | 
 **baseBranch** | **optional.**| Default Branch | 

### Return type

[**ResponseDtoConnectorResponse**](ResponseDTOConnectorResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, text/yaml, text/html, text/plain
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateTheIdentifierIsUnique**
> ResponseDtoBoolean ValidateTheIdentifierIsUnique(ctx, optional)
validate the Connector by accountIdentifier and connectorIdentifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConnectorsApiValidateTheIdentifierIsUniqueOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConnectorsApiValidateTheIdentifierIsUniqueOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountIdentifier** | **optional.String**| Account Identifier for the entity | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the entity | 
 **projectIdentifier** | **optional.String**| Project Identifier for the entity | 
 **identifier** | **optional.String**| Connector Identifier | 

### Return type

[**ResponseDtoBoolean**](ResponseDTOBoolean.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

