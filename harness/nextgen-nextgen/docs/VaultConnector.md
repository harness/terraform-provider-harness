# VaultConnector

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthToken** | **string** |  | [optional] [default to null]
**BasePath** | **string** | This is the location of the Vault directory where Secret will be stored. | [optional] [default to null]
**VaultUrl** | **string** | URL of the HashiCorp Vault. | [default to null]
**IsReadOnly** | **bool** |  | [optional] [default to null]
**RenewalIntervalMinutes** | **int64** | This is the time interval for token renewal. | [default to null]
**SecretEngineManuallyConfigured** | **bool** | Manually entered Secret Engine. | [optional] [default to null]
**SecretEngineName** | **string** | Name of the Secret Engine. | [optional] [default to null]
**AppRoleId** | **string** | ID of App Role. | [optional] [default to null]
**SecretId** | **string** |  | [optional] [default to null]
**IsDefault** | **bool** |  | [optional] [default to null]
**SecretEngineVersion** | **int32** | Version of Secret Engine. | [optional] [default to null]
**DelegateSelectors** | **[]string** | List of Delegate Selectors that belong to the same Delegate and are used to connect to the Secret Manager. | [optional] [default to null]
**Namespace** | **string** | This is the Vault namespace where Secret will be created. | [optional] [default to null]
**SinkPath** | **string** | This is the location at which auth token is to be read from. | [optional] [default to null]
**UseVaultAgent** | **bool** | Boolean value to indicate if Vault Agent is used for authentication. | [optional] [default to null]
**UseAwsIam** | **bool** | Boolean value to indicate if Aws Iam is used for authentication. | [optional] [default to null]
**AwsRegion** | **string** | This is the Aws region where aws iam auth will happen. | [optional] [default to null]
**VaultAwsIamRole** | **string** | This is the Vault role defined to bind to aws iam account/role being accessed. | [optional] [default to null]
**XvaultAwsIamServerId** | **string** |  | [optional] [default to null]
**UseK8sAuth** | **bool** | Boolean value to indicate if K8s Auth is used for authentication. | [optional] [default to null]
**VaultK8sAuthRole** | **string** | This is the role where K8s auth will happen. | [optional] [default to null]
**ServiceAccountTokenPath** | **string** | This is the SA token path where the token is mounted in the K8s Pod. | [optional] [default to null]
**AccessType** | **string** |  | [optional] [default to null]
**Default_** | **bool** |  | [optional] [default to null]
**ReadOnly** | **bool** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

