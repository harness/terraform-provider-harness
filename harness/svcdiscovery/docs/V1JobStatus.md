# V1JobStatus

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **int32** | The number of pending and running pods. +optional | [optional] [default to null]
**CompletedIndexes** | **string** | CompletedIndexes holds the completed indexes when .spec.completionMode &#x3D; \&quot;Indexed\&quot; in a text format. The indexes are represented as decimal integers separated by commas. The numbers are listed in increasing order. Three or more consecutive numbers are compressed and represented by the first and last element of the series, separated by a hyphen. For example, if the completed indexes are 1, 3, 4, 5 and 7, they are represented as \&quot;1,3-5,7\&quot;. +optional | [optional] [default to null]
**CompletionTime** | **string** | Represents time when the job was completed. It is not guaranteed to be set in happens-before order across separate operations. It is represented in RFC3339 form and is in UTC. The completion time is only set when the job finishes successfully. +optional | [optional] [default to null]
**Conditions** | [**[]V1JobCondition**](v1.JobCondition.md) | The latest available observations of an object&#x27;s current state. When a Job fails, one of the conditions will have type \&quot;Failed\&quot; and status true. When a Job is suspended, one of the conditions will have type \&quot;Suspended\&quot; and status true; when the Job is resumed, the status of this condition will become false. When a Job is completed, one of the conditions will have type \&quot;Complete\&quot; and status true. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/ +optional +patchMergeKey&#x3D;type +patchStrategy&#x3D;merge +listType&#x3D;atomic | [optional] [default to null]
**Failed** | **int32** | The number of pods which reached phase Failed. +optional | [optional] [default to null]
**Ready** | **int32** | The number of pods which have a Ready condition.  This field is beta-level. The job controller populates the field when the feature gate JobReadyPods is enabled (enabled by default). +optional | [optional] [default to null]
**StartTime** | **string** | Represents time when the job controller started processing a job. When a Job is created in the suspended state, this field is not set until the first time it is resumed. This field is reset every time a Job is resumed from suspension. It is represented in RFC3339 form and is in UTC. +optional | [optional] [default to null]
**Succeeded** | **int32** | The number of pods which reached phase Succeeded. +optional | [optional] [default to null]
**UncountedTerminatedPods** | [***V1UncountedTerminatedPods**](v1.UncountedTerminatedPods.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

