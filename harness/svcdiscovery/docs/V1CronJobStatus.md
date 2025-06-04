# V1CronJobStatus

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | [**[]V1ObjectReference**](v1.ObjectReference.md) | A list of pointers to currently running jobs. +optional +listType&#x3D;atomic | [optional] [default to null]
**LastScheduleTime** | **string** | Information when was the last time the job was successfully scheduled. +optional | [optional] [default to null]
**LastSuccessfulTime** | **string** | Information when was the last time the job successfully completed. +optional | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

