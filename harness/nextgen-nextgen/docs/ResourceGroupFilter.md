# ResourceGroupFilter

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountIdentifier** | **string** | Filter by account identifier | [default to null]
**OrgIdentifier** | **string** | Filter by organization identifier | [optional] [default to null]
**ProjectIdentifier** | **string** | Filter by project identifier | [optional] [default to null]
**SearchTerm** | **string** | Filter resource group matching by identifier/name | [optional] [default to null]
**IdentifierFilter** | **[]string** | Filter by resource group identifiers | [optional] [default to null]
**ResourceSelectorFilterList** | [**[]ResourceSelectorFilter**](ResourceSelectorFilter.md) | Filter based on whether it has a particular resource | [optional] [default to null]
**ManagedFilter** | **string** | Filter based on whether the resource group is Harness managed | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

