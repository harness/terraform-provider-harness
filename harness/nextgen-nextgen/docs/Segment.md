# Segment

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **int64** | The data and time in milliseconds when the group was created | [optional] [default to null]
**Environment** | **string** | The environment this target group belongs to | [optional] [default to null]
**Excluded** | [**[]Target**](Target.md) | A list of Targets who are excluded from this target group | [optional] [default to null]
**Identifier** | **string** | Unique identifier for the target group. | [default to null]
**Included** | [**[]Target**](Target.md) | A list of Targets who belong to this target group | [optional] [default to null]
**ModifiedAt** | **int64** | The data and time in milliseconds when the group was last modified | [optional] [default to null]
**Name** | **string** | Name of the target group. | [default to null]
**Rules** | [**[]Clause**](Clause.md) | An array of rules that can cause a user to be included in this segment. | [optional] [default to null]
**Tags** | [**[]Tag**](Tag.md) | Tags for this target group | [optional] [default to null]
**Version** | **int64** | The version of this group.  Each time it is modified the version is incremented | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

