# V1LabelSelectorRequirement

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | key is the label key that the selector applies to. +patchMergeKey&#x3D;key +patchStrategy&#x3D;merge | [optional] [default to null]
**Operator** | **string** | operator represents a key&#x27;s relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist. | [optional] [default to null]
**Values** | **[]string** | values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch. +optional | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

