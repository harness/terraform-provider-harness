# AllOfgithubComHarnessHceSaasHceSdkCommonProbeV1ProbeRunProperties

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attempt** | **int32** | Attempt contains the total attempt count for the probe | [optional] [default to null]
**InitialDelay** | **string** | InitialDelay time interval for which probe will wait before run | [optional] [default to null]
**InitialDelaySeconds** | **int32** | InitialDelaySeconds time interval for which probe will wait before run | [optional] [default to null]
**Interval** | **string** | Interval contains the interval for the probe | [optional] [default to null]
**ProbePollingInterval** | **string** | ProbePollingInterval contains time interval, for which continuous probe should be sleep after each iteration | [optional] [default to null]
**ProbeTimeout** | **string** | ProbeTimeout contains timeout for the probe | [optional] [default to null]
**Retry** | **int32** | Retry contains the retry count for the probe | [optional] [default to null]
**StopOnFailure** | **bool** | StopOnFailure contains flag to stop/continue experiment execution, if probe fails it will stop the experiment execution, if provided true it will continue the experiment execution, if provided false or not provided(default case) | [optional] [default to null]
**Verbosity** | **string** | Verbosity contains flag to set the verbosity of probe it support info and debug verbosity | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

