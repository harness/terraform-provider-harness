# EvaluatedPolicy

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DenyMessages** | **[]string** | The values of any &#x60;deny&#x60; rego rules as returned by the rego engine | [default to null]
**Error_** | **string** | Any errors returned by the rego engine when this policy was evaluated | [default to null]
**Output** | [****os.File**](*os.File.md) | The output returned by the rego engine when this policy was evaluated | [default to null]
**Policy** | [***Policy**](Policy.md) |  | [default to null]
**Status** | **string** | The overall status for this individual policy indicating whether it passed | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

