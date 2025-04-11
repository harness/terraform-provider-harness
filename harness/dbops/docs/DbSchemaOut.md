# DbSchemaOut

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | **string** | identifier of the database schema | [default to null]
**Name** | **string** | name of the database schema | [default to null]
**Created** | **int64** | epoch seconds when the database schema was created | [default to null]
**Updated** | **int64** | epoch seconds when the database schema was last updated | [optional] [default to null]
**Tags** | **map[string]string** | tags attached to the database schema | [optional] [default to null]
**Changelog** | [***Changelog**](Changelog.md) |  | [optional] [default to null]
**Service** | **string** | harness service corresponding to database schema | [optional] [default to null]
**InstanceCount** | **int64** | number of database instances corresponding to database schema | [default to null]
**SchemaSourceType** | **string** |  | [optional] [default to null]
**ChangeLogScript** | [***ChangeLogScript**](ChangeLogScript.md) |  | [optional] [default to null]
**Type_** | [***DbSchemaType**](DBSchemaType.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

