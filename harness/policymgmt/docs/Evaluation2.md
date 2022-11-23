# Evaluation2

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The Harness account in which the evaluation was performed | [default to null]
**Action** | **string** | The action that triggered evaluation | [default to null]
**Created** | **int64** | The time at which the evaluation was performed in Unix time millseconds | [default to null]
**Details** | [**[]EvaluationDetail2**](EvaluationDetail2.md) | The detailed results of te evaluation | [default to null]
**Entity** | **string** | An arbtrary user-supplied string that globally identifies the entity under evaluation | [default to null]
**EntityMetadata** | **string** | Additional arbtrary user-supplied metadta about the entity under evaluation | [default to null]
**Id** | **int64** | The ID of this evaluation | [default to null]
**Input** | [****os.File**](*os.File.md) | The input provided at evaluation time | [default to null]
**OrgId** | **string** | The Harness organisation in which the evaluation was performed | [default to null]
**ProjectId** | **string** | The Harness project in which the evaluation was performed | [default to null]
**Status** | **string** | The overall status of the evaluation indicating whether it passed | [default to null]
**Type_** | **string** | The type of the entity under evaluation | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

