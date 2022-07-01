# AwsKmsCredentialSpecAssumeSts

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DelegateSelectors** | **[]string** | List of Delegate Selectors that belong to the same Delegate and are used to connect to the Secret Manager. | [default to null]
**RoleArn** | **string** | Role ARN for the Delegate with STS Role. | [default to null]
**ExternalName** | **string** | External Name. | [optional] [default to null]
**AssumeStsRoleDuration** | **int32** | This is the time duration for STS Role. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

