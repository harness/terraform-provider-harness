# InputSetSummaryResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | **string** | Input Set Identifier | [optional] [default to null]
**Name** | **string** | Input Set Name | [optional] [default to null]
**PipelineIdentifier** | **string** | Pipeline Identifier for the entity. | [optional] [default to null]
**Description** | **string** | Input Set description | [optional] [default to null]
**InputSetType** | **string** | Type of Input Set. The default value is ALL. | [optional] [default to null]
**Tags** | **map[string]string** | Input Set tags | [optional] [default to null]
**GitDetails** | [***PipelineEntityGitDetails**](PipelineEntityGitDetails.md) |  | [optional] [default to null]
**CreatedAt** | **int64** | Time at which the entity was created | [optional] [default to null]
**LastUpdatedAt** | **int64** | Time at which the entity was last updated | [optional] [default to null]
**IsOutdated** | **bool** | This field is true if a Pipeline update has made this Input Set invalid, and cannot be used for Pipeline Execution | [optional] [default to null]
**InputSetErrorDetails** | [***InputSetErrorWrapper**](InputSetErrorWrapper.md) |  | [optional] [default to null]
**OverlaySetErrorDetails** | **map[string]string** | This contains the invalid references in the Overlay Input Set, along with a message why they are invalid | [optional] [default to null]
**EntityValidityDetails** | [***PipelineEntityGitDetails**](PipelineEntityGitDetails.md) |  | [optional] [default to null]
**Modules** | **[]string** | Modules in which the Pipeline belongs | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

