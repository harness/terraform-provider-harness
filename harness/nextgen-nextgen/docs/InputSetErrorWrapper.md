# InputSetErrorWrapper

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorPipelineYaml** | **string** | If an Input Set save fails, this field contains the error fields, with the field values replaced with a UUID | [optional] [default to null]
**UuidToErrorResponseMap** | [**map[string]InputSetErrorWrapper**](InputSetErrorWrapper.md) | If an Input Set save fails, this field contains the map from FQN to why that FQN threw an error | [optional] [default to null]
**Type_** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

