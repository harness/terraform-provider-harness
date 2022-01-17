# {{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConnector**](ConnectorsApi.md#CreateConnector) | **Post** /ng/api/connectors | Creates a Connector
[**DeleteConnector**](ConnectorsApi.md#DeleteConnector) | **Delete** /ng/api/connectors/{identifier} | Deletes Connector by ID
[**GetAllAllowedFieldValues**](ConnectorsApi.md#GetAllAllowedFieldValues) | **Get** /ng/api/connectors/fieldValues | Get the allowed field values by Connector Type
[**GetCEAwsTemplate**](ConnectorsApi.md#GetCEAwsTemplate) | **Post** /ng/api/connectors/getceawstemplateurl | Get the Template URL of connector
[**GetConnector**](ConnectorsApi.md#GetConnector) | **Get** /ng/api/connectors/{identifier} | Get the Connector by accountIdentifier and connectorIdentifier
[**GetConnectorCatalogue**](ConnectorsApi.md#GetConnectorCatalogue) | **Get** /ng/api/connectors/catalogue | Gets the Connector catalogue by Account Identifier
[**GetConnectorList**](ConnectorsApi.md#GetConnectorList) | **Get** /ng/api/connectors | Fetches the list of Connectors corresponding to the request&#x27;s filter criteria.
[**GetConnectorListV2**](ConnectorsApi.md#GetConnectorListV2) | **Post** /ng/api/connectors/listV2 | Fetches the list of Connectors corresponding to the request&#x27;s filter criteria.
[**GetConnectorStatistics**](ConnectorsApi.md#GetConnectorStatistics) | **Get** /ng/api/connectors/stats | Gets the connector&#x27;s statistics by Account Identifier, Project Identifier and Organization Identifier
[**GetTestConnectionResult**](ConnectorsApi.md#GetTestConnectionResult) | **Post** /ng/api/connectors/testConnection/{identifier} | Tests the connection of the Connector by ID
[**GetTestGitRepoConnectionResult**](ConnectorsApi.md#GetTestGitRepoConnectionResult) | **Post** /ng/api/connectors/testGitRepoConnection/{identifier} | Tests the Git Repo connection
[**ListConnectorByFQN**](ConnectorsApi.md#ListConnectorByFQN) | **Post** /ng/api/connectors/listbyfqn | Get the list of connectors by FQN satisfying the criteria (if any) in the request
[**UpdateConnector**](ConnectorsApi.md#UpdateConnector) | **Put** /ng/api/connectors | Updates the Connector
[**ValidateTheIdentifierIsUnique**](ConnectorsApi.md#ValidateTheIdentifierIsUnique) | **Get** /ng/api/connectors/validateUniqueIdentifier | Validate the Connector by Account Identifier and Connector Identifier

# **CreateConnector**
> ResponseDtoConnectorResponse CreateConnector(ctx, body, accountIdentifier, optional)
Creates a Connector

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Connector**](Connector.md)| Details of the Connector to create | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
 **optional** | ***ConnectorsApiCreateConnectorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConnectorsApiCreateConnectorOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **branch** | **optional.**| Branch Name | 
 **repoIdentifier** | **optional.**| Git Sync Config Id | 
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
Deletes Connector by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
  **identifier** | **string**| Connector ID | 
 **optional** | ***ConnectorsApiDeleteConnectorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConnectorsApiDeleteConnectorOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity | 
 **branch** | **optional.String**| Branch Name | 
 **repoIdentifier** | **optional.String**| Git Sync Config Id | 
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
> ResponseDtoConnectorResponse GetConnector(ctx, accountIdentifier, identifier, optional)
Get the Connector by accountIdentifier and connectorIdentifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
  **identifier** | **string**| Connector Identifier | 
 **optional** | ***ConnectorsApiGetConnectorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConnectorsApiGetConnectorOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity | 
 **branch** | **optional.String**| Branch Name | 
 **repoIdentifier** | **optional.String**| Git Sync Config Id | 
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
> ResponseDtoConnectorCatalogueResponse GetConnectorCatalogue(ctx, accountIdentifier)
Gets the Connector catalogue by Account Identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity | 

### Return type

[**ResponseDtoConnectorCatalogueResponse**](ResponseDTOConnectorCatalogueResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConnectorList**
> ResponseDtoPageResponseConnectorResponse GetConnectorList(ctx, accountIdentifier, optional)
Fetches the list of Connectors corresponding to the request's filter criteria.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
 **optional** | ***ConnectorsApiGetConnectorListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConnectorsApiGetConnectorListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageIndex** | **optional.Int32**| Page number of navigation. The default value is 0 | [default to 0]
 **pageSize** | **optional.Int32**| Number of entries per page. The default value is 100 | [default to 100]
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity | 
 **searchTerm** | **optional.String**| This would be used to filter Connectors. Any Connector having the specified string in its Name, ID and Tag would be filtered. | 
 **type_** | **optional.String**| Filter Connectors by type | 
 **category** | **optional.String**| Filter Connectors by category | 
 **sourceCategory** | **optional.String**| Filter Connectors by Source Category. Available Source Categories are CLOUD_PROVIDER, SECRET_MANAGER, CLOUD_COST, ARTIFACTORY, CODE_REPO,  MONITORING and TICKETING | 
 **branch** | **optional.String**| Branch Name | 
 **repoIdentifier** | **optional.String**| Git Sync Config Id | 
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
> ResponseDtoPageResponseConnectorResponse GetConnectorListV2(ctx, body, accountIdentifier, optional)
Fetches the list of Connectors corresponding to the request's filter criteria.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConnectorFilterProperties**](ConnectorFilterProperties.md)| Details of the filters applied | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
 **optional** | ***ConnectorsApiGetConnectorListV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConnectorsApiGetConnectorListV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageIndex** | **optional.**| Page number of navigation. The default value is 0 | [default to 0]
 **pageSize** | **optional.**| Number of entries per page. The default value is 100 | [default to 100]
 **searchTerm** | **optional.**| This would be used to filter Connectors. Any Connector having the specified string in its Name, ID and Tag would be filtered. | 
 **orgIdentifier** | **optional.**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity | 
 **filterIdentifier** | **optional.**|  | 
 **includeAllConnectorsAvailableAtScope** | **optional.**| Specify whether or not to include all the Connectors accessible at the scope. For eg if set as true, at the Project scope we will get org and account Connector also in the response | 
 **branch** | **optional.**| Branch Name | 
 **repoIdentifier** | **optional.**| Git Sync Config Id | 
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
> ResponseDtoConnectorStatistics GetConnectorStatistics(ctx, accountIdentifier, optional)
Gets the connector's statistics by Account Identifier, Project Identifier and Organization Identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
 **optional** | ***ConnectorsApiGetConnectorStatisticsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConnectorsApiGetConnectorStatisticsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity | 
 **branch** | **optional.String**| Branch Name | 
 **repoIdentifier** | **optional.String**| Git Sync Config Id | 
 **getDefaultFromOtherRepo** | **optional.Bool**| if true, return all the default entities | 

### Return type

[**ResponseDtoConnectorStatistics**](ResponseDTOConnectorStatistics.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTestConnectionResult**
> ResponseDtoConnectorValidationResult GetTestConnectionResult(ctx, accountIdentifier, identifier, optional)
Tests the connection of the Connector by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
  **identifier** | **string**| Connector ID | 
 **optional** | ***ConnectorsApiGetTestConnectionResultOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConnectorsApiGetTestConnectionResultOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity | 
 **branch** | **optional.String**| Branch Name | 
 **repoIdentifier** | **optional.String**| Git Sync Config Id | 
 **getDefaultFromOtherRepo** | **optional.Bool**| if true, return all the default entities | 

### Return type

[**ResponseDtoConnectorValidationResult**](ResponseDTOConnectorValidationResult.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTestGitRepoConnectionResult**
> ResponseDtoConnectorValidationResult GetTestGitRepoConnectionResult(ctx, accountIdentifier, identifier, optional)
Tests the Git Repo connection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
  **identifier** | **string**| Connector ID | 
 **optional** | ***ConnectorsApiGetTestGitRepoConnectionResultOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConnectorsApiGetTestGitRepoConnectionResultOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity | 
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
> ResponseDtoListConnectorResponse ListConnectorByFQN(ctx, body, accountIdentifier)
Get the list of connectors by FQN satisfying the criteria (if any) in the request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]string**](string.md)| List of ConnectorsFQN as strings | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 

### Return type

[**ResponseDtoListConnectorResponse**](ResponseDTOListConnectorResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, text/yaml, text/html, text/plain
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateConnector**
> ResponseDtoConnectorResponse UpdateConnector(ctx, body, accountIdentifier, optional)
Updates the Connector

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Connector**](Connector.md)| This is the updated Connector. Please provide values for all fields, not just the fields you are updating | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
 **optional** | ***ConnectorsApiUpdateConnectorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConnectorsApiUpdateConnectorOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **branch** | **optional.**| Branch Name | 
 **repoIdentifier** | **optional.**| Git Sync Config Id | 
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
> ResponseDtoBoolean ValidateTheIdentifierIsUnique(ctx, accountIdentifier, optional)
Validate the Connector by Account Identifier and Connector Identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
 **optional** | ***ConnectorsApiValidateTheIdentifierIsUniqueOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConnectorsApiValidateTheIdentifierIsUniqueOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity | 
 **identifier** | **optional.String**| Connector ID | 

### Return type

[**ResponseDtoBoolean**](ResponseDTOBoolean.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

