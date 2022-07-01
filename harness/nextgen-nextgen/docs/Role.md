# Role

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | **string** | Unique identifier of the role | [default to null]
**Name** | **string** | Name of the role | [default to null]
**Permissions** | **[]string** | List of the permission identifiers (Subset of the list returned by GET /authz/api/permissions) | [optional] [default to null]
**AllowedScopeLevels** | **[]string** | The scope levels at which this role can be used | [optional] [default to null]
**Description** | **string** | Description of the role | [optional] [default to null]
**Tags** | **map[string]string** | Tags | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

