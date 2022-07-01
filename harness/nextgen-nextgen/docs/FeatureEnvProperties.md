# FeatureEnvProperties

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultServe** | [***Serve**](Serve.md) |  | [default to null]
**Environment** | **string** | The environment identifier | [default to null]
**ModifiedAt** | **int64** | The last time the flag was modified in this environment | [optional] [default to null]
**OffVariation** | **string** | The variation to serve for this flag in this environment when the flag is off | [default to null]
**Rules** | [**[]ServingRule**](ServingRule.md) | A list of rules to use when evaluating this flag in this environment | [optional] [default to null]
**State** | [***FeatureState**](FeatureState.md) |  | [default to null]
**VariationMap** | [**[]VariationMap**](VariationMap.md) | A list of the variations that will be served to specific targets or target groups in an environment. | [optional] [default to null]
**Version** | **int64** | The version of the flag.  This is incremented each time it is changed | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

