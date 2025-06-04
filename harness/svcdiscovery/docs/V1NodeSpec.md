# V1NodeSpec

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigSource** | [***V1NodeConfigSource**](v1.NodeConfigSource.md) |  | [optional] [default to null]
**ExternalID** | **string** | Deprecated. Not all kubelets will set this field. Remove field after 1.13. see: https://issues.k8s.io/61966 +optional | [optional] [default to null]
**PodCIDR** | **string** | PodCIDR represents the pod IP range assigned to the node. +optional | [optional] [default to null]
**PodCIDRs** | **[]string** | podCIDRs represents the IP ranges assigned to the node for usage by Pods on that node. If this field is specified, the 0th entry must match the podCIDR field. It may contain at most 1 value for each of IPv4 and IPv6. +optional +patchStrategy&#x3D;merge | [optional] [default to null]
**ProviderID** | **string** | ID of the node assigned by the cloud provider in the format: &lt;ProviderName&gt;://&lt;ProviderSpecificNodeID&gt; +optional | [optional] [default to null]
**Taints** | [**[]V1Taint**](v1.Taint.md) | If specified, the node&#x27;s taints. +optional | [optional] [default to null]
**Unschedulable** | **bool** | Unschedulable controls node schedulability of new pods. By default, node is schedulable. More info: https://kubernetes.io/docs/concepts/nodes/node/#manual-node-administration +optional | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

