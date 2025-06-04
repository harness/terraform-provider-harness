# V1StatefulSetStatus

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailableReplicas** | **int32** | Total number of available pods (ready for at least minReadySeconds) targeted by this statefulset. +optional | [optional] [default to null]
**CollisionCount** | **int32** | collisionCount is the count of hash collisions for the StatefulSet. The StatefulSet controller uses this field as a collision avoidance mechanism when it needs to create the name for the newest ControllerRevision. +optional | [optional] [default to null]
**Conditions** | [**[]V1StatefulSetCondition**](v1.StatefulSetCondition.md) | Represents the latest available observations of a statefulset&#x27;s current state. +optional +patchMergeKey&#x3D;type +patchStrategy&#x3D;merge | [optional] [default to null]
**CurrentReplicas** | **int32** | currentReplicas is the number of Pods created by the StatefulSet controller from the StatefulSet version indicated by currentRevision. | [optional] [default to null]
**CurrentRevision** | **string** | currentRevision, if not empty, indicates the version of the StatefulSet used to generate Pods in the sequence [0,currentReplicas). | [optional] [default to null]
**ObservedGeneration** | **int32** | observedGeneration is the most recent generation observed for this StatefulSet. It corresponds to the StatefulSet&#x27;s generation, which is updated on mutation by the API Server. +optional | [optional] [default to null]
**ReadyReplicas** | **int32** | readyReplicas is the number of pods created for this StatefulSet with a Ready Condition. | [optional] [default to null]
**Replicas** | **int32** | replicas is the number of Pods created by the StatefulSet controller. | [optional] [default to null]
**UpdateRevision** | **string** | updateRevision, if not empty, indicates the version of the StatefulSet used to generate Pods in the sequence [replicas-updatedReplicas,replicas) | [optional] [default to null]
**UpdatedReplicas** | **int32** | updatedReplicas is the number of Pods created by the StatefulSet controller from the StatefulSet version indicated by updateRevision. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

