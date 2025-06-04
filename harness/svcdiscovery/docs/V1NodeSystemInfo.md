# V1NodeSystemInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Architecture** | **string** | The Architecture reported by the node | [optional] [default to null]
**BootID** | **string** | Boot ID reported by the node. | [optional] [default to null]
**ContainerRuntimeVersion** | **string** | ContainerRuntime Version reported by the node through runtime remote API (e.g. containerd://1.4.2). | [optional] [default to null]
**KernelVersion** | **string** | Kernel Version reported by the node from &#x27;uname -r&#x27; (e.g. 3.16.0-0.bpo.4-amd64). | [optional] [default to null]
**KubeProxyVersion** | **string** | KubeProxy Version reported by the node. | [optional] [default to null]
**KubeletVersion** | **string** | Kubelet Version reported by the node. | [optional] [default to null]
**MachineID** | **string** | MachineID reported by the node. For unique machine identification in the cluster this field is preferred. Learn more from man(5) machine-id: http://man7.org/linux/man-pages/man5/machine-id.5.html | [optional] [default to null]
**OperatingSystem** | **string** | The Operating System reported by the node | [optional] [default to null]
**OsImage** | **string** | OS Image reported by the node from /etc/os-release (e.g. Debian GNU/Linux 7 (wheezy)). | [optional] [default to null]
**SystemUUID** | **string** | SystemUUID reported by the node. For unique machine identification MachineID is preferred. This field is specific to Red Hat hosts https://access.redhat.com/documentation/en-us/red_hat_subscription_management/1/html/rhsm/uuid | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

