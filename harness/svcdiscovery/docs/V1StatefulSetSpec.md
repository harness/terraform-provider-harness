# V1StatefulSetSpec

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinReadySeconds** | **int32** | Minimum number of seconds for which a newly created pod should be ready without any of its container crashing for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready) +optional | [optional] [default to null]
**PersistentVolumeClaimRetentionPolicy** | [***V1StatefulSetPersistentVolumeClaimRetentionPolicy**](v1.StatefulSetPersistentVolumeClaimRetentionPolicy.md) |  | [optional] [default to null]
**PodManagementPolicy** | **string** | podManagementPolicy controls how pods are created during initial scale up, when replacing pods on nodes, or when scaling down. The default policy is &#x60;OrderedReady&#x60;, where pods are created in increasing order (pod-0, then pod-1, etc) and the controller will wait until each pod is ready before continuing. When scaling down, the pods are removed in the opposite order. The alternative policy is &#x60;Parallel&#x60; which will create pods in parallel to match the desired scale without waiting, and on scale down will delete all pods at once. +optional | [optional] [default to null]
**Replicas** | **int32** | replicas is the desired number of replicas of the given Template. These are replicas in the sense that they are instantiations of the same Template, but individual replicas also have a consistent identity. If unspecified, defaults to 1. TODO: Consider a rename of this field. +optional | [optional] [default to null]
**RevisionHistoryLimit** | **int32** | revisionHistoryLimit is the maximum number of revisions that will be maintained in the StatefulSet&#x27;s revision history. The revision history consists of all revisions not represented by a currently applied StatefulSetSpec version. The default value is 10. | [optional] [default to null]
**Selector** | [***V1LabelSelector**](v1.LabelSelector.md) |  | [optional] [default to null]
**ServiceName** | **string** | serviceName is the name of the service that governs this StatefulSet. This service must exist before the StatefulSet, and is responsible for the network identity of the set. Pods get DNS/hostnames that follow the pattern: pod-specific-string.serviceName.default.svc.cluster.local where \&quot;pod-specific-string\&quot; is managed by the StatefulSet controller. | [optional] [default to null]
**Template** | [***V1PodTemplateSpec**](v1.PodTemplateSpec.md) |  | [optional] [default to null]
**UpdateStrategy** | [***V1StatefulSetUpdateStrategy**](v1.StatefulSetUpdateStrategy.md) |  | [optional] [default to null]
**VolumeClaimTemplates** | [**[]V1PersistentVolumeClaim**](v1.PersistentVolumeClaim.md) | volumeClaimTemplates is a list of claims that pods are allowed to reference. The StatefulSet controller is responsible for mapping network identities to claims in a way that maintains the identity of a pod. Every claim in this list must have at least one matching (by name) volumeMount in one container in the template. A claim in this list takes precedence over any volumes in the template, with the same name. TODO: Define the behavior if a claim already exists with the same name. +optional | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

