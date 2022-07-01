# InputSetTemplateWithReplacedExpressionsResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InputSetTemplateYaml** | **string** | Runtime Input template for the Pipeline | [optional] [default to null]
**ReplacedExpressions** | **[]string** | List of Expressions that need to be replaced for running selected Stages. Empty if the full Pipeline is being run or no expressions need to be replaced | [optional] [default to null]
**Modules** | **[]string** | Modules in which the Pipeline belongs | [optional] [default to null]
**HasInputSets** | **bool** | Tells whether there are any Input Sets for this Pipeline or not. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

