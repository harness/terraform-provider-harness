# Feature

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Archived** | **bool** | Indicates if the flag has been archived and is no longer used | [optional] [default to null]
**CreatedAt** | **int64** | The date the flag was created in milliseconds | [default to null]
**DefaultOffVariation** | **string** | The default value returned when a flag is off | [default to null]
**DefaultOnVariation** | **string** | The default value returned when a flag is on | [default to null]
**Description** | **string** | A description for this flag | [optional] [default to null]
**EnvProperties** | [***FeatureEnvProperties**](Feature_envProperties.md) |  | [optional] [default to null]
**Evaluation** | **string** | The value that the flag will return for the current user | [optional] [default to null]
**EvaluationIdentifier** | **string** | The identifier for the returned evaluation | [optional] [default to null]
**Identifier** | **string** | The Feature Flag identifier | [default to null]
**Kind** | **string** | The type of Feature flag | [default to null]
**ModifiedAt** | **int64** | The date the flag was last modified in milliseconds | [optional] [default to null]
**Name** | **string** | The name of the Feature Flag | [default to null]
**Owner** | **[]string** | The user who created the flag | [optional] [default to null]
**Permanent** | **bool** | Indicates if this is a permanent flag, or one that should expire | [optional] [default to null]
**Prerequisites** | [**[]Prerequisite**](Prerequisite.md) |  | [optional] [default to null]
**Project** | **string** | The project this Feature belongs to | [default to null]
**Results** | [**[]Results**](Results.md) | The results shows which variations have been evaluated, and how many times each of these have been evaluated. | [optional] [default to null]
**Status** | [***FeatureStatus**](FeatureStatus.md) |  | [optional] [default to null]
**Tags** | [**[]Tag**](Tag.md) | A list of tags for this Feature Flag | [optional] [default to null]
**Variations** | [**[]Variation**](Variation.md) | The variations that can be returned for this flag | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

