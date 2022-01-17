# PipelineExecutionSummary

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PipelineIdentifier** | **string** |  | [optional] [default to null]
**PlanExecutionId** | **string** |  | [optional] [default to null]
**Name** | **string** |  | [optional] [default to null]
**Status** | **string** | This is the Execution Status of the entity | [optional] [default to null]
**Tags** | [**[]NgTag**](NGTag.md) |  | [optional] [default to null]
**ExecutionTriggerInfo** | [***ExecutionTriggerInfo**](ExecutionTriggerInfo.md) |  | [optional] [default to null]
**ExecutionErrorInfo** | [***ExecutionErrorInfo**](ExecutionErrorInfo.md) |  | [optional] [default to null]
**GovernanceMetadata** | [***GovernanceMetadata**](GovernanceMetadata.md) |  | [optional] [default to null]
**ModuleInfo** | [**map[string]map[string]interface{}**](map.md) |  | [optional] [default to null]
**LayoutNodeMap** | [**map[string]GraphLayoutNode**](GraphLayoutNode.md) |  | [optional] [default to null]
**Modules** | **[]string** |  | [optional] [default to null]
**StartingNodeId** | **string** |  | [optional] [default to null]
**StartTs** | **int64** |  | [optional] [default to null]
**EndTs** | **int64** |  | [optional] [default to null]
**CreatedAt** | **int64** |  | [optional] [default to null]
**CanRetry** | **bool** |  | [optional] [default to null]
**ShowRetryHistory** | **bool** |  | [optional] [default to null]
**RunSequence** | **int32** |  | [optional] [default to null]
**SuccessfulStagesCount** | **int64** |  | [optional] [default to null]
**RunningStagesCount** | **int64** |  | [optional] [default to null]
**FailedStagesCount** | **int64** |  | [optional] [default to null]
**TotalStagesCount** | **int64** |  | [optional] [default to null]
**GitDetails** | [***EntityGitDetails**](EntityGitDetails.md) |  | [optional] [default to null]
**IsStagesExecution** | **bool** |  | [optional] [default to null]
**StagesExecuted** | **[]string** |  | [optional] [default to null]
**StagesExecutedNames** | **map[string]string** |  | [optional] [default to null]
**StagesExecution** | **bool** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

