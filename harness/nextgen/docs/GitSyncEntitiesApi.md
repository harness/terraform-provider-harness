# {{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListGitSyncEntitiesByType**](GitSyncEntitiesApi.md#ListGitSyncEntitiesByType) | **Get** /ng/api/git-sync-entities/{entityType} | Lists Git Sync Entity filtered by their Type for the given scope
[**ListGitSyncEntitiesSummaryForRepoAndBranch**](GitSyncEntitiesApi.md#ListGitSyncEntitiesSummaryForRepoAndBranch) | **Post** /ng/api/git-sync-entities/branch/{branch} | Lists Git Sync Entity by product for the given Repo, Branch and list of Entity Types
[**ListGitSyncEntitiesSummaryForRepoAndTypes**](GitSyncEntitiesApi.md#ListGitSyncEntitiesSummaryForRepoAndTypes) | **Post** /ng/api/git-sync-entities/summary | Lists Git Sync Entity by product for the given list of Repos and Entity Types

# **ListGitSyncEntitiesByType**
> ResponseDtoPageResponseGitSyncEntityList ListGitSyncEntitiesByType(ctx, entityType, optional)
Lists Git Sync Entity filtered by their Type for the given scope

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **entityType** | **string**| Entity Type | 
 **optional** | ***GitSyncEntitiesApiListGitSyncEntitiesByTypeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GitSyncEntitiesApiListGitSyncEntitiesByTypeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectIdentifier** | **optional.String**| Project Identifier for the Entity | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity | 
 **accountIdentifier** | **optional.String**| Account Identifier for the Entity | 
 **gitSyncConfigId** | **optional.String**| Git Sync Config Id | 
 **branch** | **optional.String**| Branch Name | 
 **page** | **optional.Int32**| Indicates the number of pages. Results for these pages will be retrieved. | [default to 0]
 **size** | **optional.Int32**| The number of the elements to fetch | 
 **moduleType** | **optional.String**| Module Type | 

### Return type

[**ResponseDtoPageResponseGitSyncEntityList**](ResponseDTOPageResponseGitSyncEntityList.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListGitSyncEntitiesSummaryForRepoAndBranch**
> ResponseDtoListGitSyncEntityList ListGitSyncEntitiesSummaryForRepoAndBranch(ctx, body, branch, optional)
Lists Git Sync Entity by product for the given Repo, Branch and list of Entity Types

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GitEntityBranchSummaryFilter**](GitEntityBranchSummaryFilter.md)| This filters the Git Sync Entity based on multiple parameters | 
  **branch** | **string**| Branch Name | 
 **optional** | ***GitSyncEntitiesApiListGitSyncEntitiesSummaryForRepoAndBranchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GitSyncEntitiesApiListGitSyncEntitiesSummaryForRepoAndBranchOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **projectIdentifier** | **optional.**| Project Identifier for the Entity | 
 **orgIdentifier** | **optional.**| Organization Identifier for the Entity | 
 **accountIdentifier** | **optional.**| Account Identifier for the Entity | 
 **size** | **optional.**| The number of the elements to fetch | 
 **gitSyncConfigId** | **optional.**| Git Sync Config Id | 

### Return type

[**ResponseDtoListGitSyncEntityList**](ResponseDTOListGitSyncEntityList.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, text/yaml, text/html
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListGitSyncEntitiesSummaryForRepoAndTypes**
> ResponseDtoGitSyncRepoFilesList ListGitSyncEntitiesSummaryForRepoAndTypes(ctx, body, optional)
Lists Git Sync Entity by product for the given list of Repos and Entity Types

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GitEntitySummaryFilter**](GitEntitySummaryFilter.md)| Filter Git Sync Entity based on multiple parameters | 
 **optional** | ***GitSyncEntitiesApiListGitSyncEntitiesSummaryForRepoAndTypesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GitSyncEntitiesApiListGitSyncEntitiesSummaryForRepoAndTypesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectIdentifier** | **optional.**| Project Identifier for the Entity | 
 **orgIdentifier** | **optional.**| Organization Identifier for the Entity | 
 **accountIdentifier** | **optional.**| Account Identifier for the Entity | 
 **size** | **optional.**| The number of the elements to fetch | 

### Return type

[**ResponseDtoGitSyncRepoFilesList**](ResponseDTOGitSyncRepoFilesList.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, text/yaml, text/html
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

