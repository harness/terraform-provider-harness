# GcpKmsConnector

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** | ID of the project on GCP. | [default to null]
**Region** | **string** | Region for GCP KMS | [default to null]
**KeyRing** | **string** | Name of the Key Ring where Google Cloud Symmetric Key is created. | [default to null]
**KeyName** | **string** | Name of the Google Cloud Symmetric Key. | [default to null]
**Credentials** | **string** |  | [default to null]
**DelegateSelectors** | **[]string** | List of Delegate Selectors that belong to the same Delegate and are used to connect to the Secret Manager. | [optional] [default to null]
**Default_** | **bool** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

