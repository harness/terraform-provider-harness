# V1DaemonSetStatus

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollisionCount** | **int32** | Count of hash collisions for the DaemonSet. The DaemonSet controller uses this field as a collision avoidance mechanism when it needs to create the name for the newest ControllerRevision. +optional | [optional] [default to null]
**Conditions** | [**[]V1DaemonSetCondition**](v1.DaemonSetCondition.md) | Represents the latest available observations of a DaemonSet&#x27;s current state. +optional +patchMergeKey&#x3D;type +patchStrategy&#x3D;merge | [optional] [default to null]
**CurrentNumberScheduled** | **int32** | The number of nodes that are running at least 1 daemon pod and are supposed to run the daemon pod. More info: https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/ | [optional] [default to null]
**DesiredNumberScheduled** | **int32** | The total number of nodes that should be running the daemon pod (including nodes correctly running the daemon pod). More info: https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/ | [optional] [default to null]
**NumberAvailable** | **int32** | The number of nodes that should be running the daemon pod and have one or more of the daemon pod running and available (ready for at least spec.minReadySeconds) +optional | [optional] [default to null]
**NumberMisscheduled** | **int32** | The number of nodes that are running the daemon pod, but are not supposed to run the daemon pod. More info: https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/ | [optional] [default to null]
**NumberReady** | **int32** | numberReady is the number of nodes that should be running the daemon pod and have one or more of the daemon pod running with a Ready Condition. | [optional] [default to null]
**NumberUnavailable** | **int32** | The number of nodes that should be running the daemon pod and have none of the daemon pod running and available (ready for at least spec.minReadySeconds) +optional | [optional] [default to null]
**ObservedGeneration** | **int32** | The most recent generation observed by the daemon set controller. +optional | [optional] [default to null]
**UpdatedNumberScheduled** | **int32** | The total number of nodes that are running updated daemon pod +optional | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

