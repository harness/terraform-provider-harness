# RoleAssignmentFilter

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceGroupFilter** | **[]string** | Filter role assignments based on resource group identifiers | [optional] [default to null]
**RoleFilter** | **[]string** | Filter role assignments based on role identifiers | [optional] [default to null]
**PrincipalTypeFilter** | **[]string** | Filter role assignments based on principal type | [optional] [default to null]
**PrincipalFilter** | [**[]Principal**](Principal.md) | Filter role assignments based on principals | [optional] [default to null]
**HarnessManagedFilter** | **[]bool** | Filter role assignments based on role assignments being harness managed | [optional] [default to null]
**DisabledFilter** | **[]bool** | Filter role assignments based on whether they are enabled or disabled | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

