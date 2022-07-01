# AzureRepoConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | SSH | HTTP URL based on type of connection | [default to null]
**ValidationProject** | **string** | The project to validate AzureRepo credentials. Only valid for Account type connector | [optional] [default to null]
**ValidationRepo** | **string** | The repo to validate AzureRepo credentials. Only valid for Account type connector | [optional] [default to null]
**Authentication** | [***AzureRepoAuthentication**](AzureRepoAuthentication.md) |  | [default to null]
**ApiAccess** | [***AzureRepoApiAccess**](AzureRepoApiAccess.md) |  | [optional] [default to null]
**DelegateSelectors** | **[]string** | Selected Connectivity Modes | [optional] [default to null]
**Type_** | **string** | Account | Repository connector type | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

