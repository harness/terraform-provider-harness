# V1CmdProbeInputs

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | **string** | Command need to be executed for the probe | [optional] [default to null]
**Comparator** | [***AllOfv1CmdProbeInputsComparator**](AllOfv1CmdProbeInputsComparator.md) | Comparator check for the correctness of the probe output | [optional] [default to null]
**Source** | [***AllOfv1CmdProbeInputsSource**](AllOfv1CmdProbeInputsSource.md) | The source where we have to run the command It will run in inline(inside experiment itself) mode if source is nil | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

