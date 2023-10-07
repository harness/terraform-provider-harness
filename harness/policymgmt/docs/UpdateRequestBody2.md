# UpdateRequestBody2

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | Action that triggers the policy set | [optional] [default to null]
**Description** | **string** | Description of the policy set | [optional] [default to null]
**Enabled** | **bool** | Only enabled policy sets are evaluated when evaluating by type/action | [optional] [default to null]
**EntitySelector** | **string** | A string enum value which determines which entities the policy set applies to during evaluation. This feature is not available for all accounts, Contact support if you wish to have it enabled. | [optional] [default to null]
**Name** | **string** | Name of the policy set | [optional] [default to null]
**Policies** | [**[]Linkedpolicyidentifier**](Linkedpolicyidentifier.md) | Policies linked to this policy set | [optional] [default to null]
**ResourceGroups** | [**[]ResourceGroupIdentifier**](ResourceGroupIdentifier.md) | Resource groups that contain the resources that this policy set should be evaluated for. Resource groups are not supported for flag or custom policy sets. This feature is not available for all accounts, Contact support if you wish to have it enabled. | [optional] [default to null]
**Type_** | **string** | Type of input suitable for the policy set | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

