# AllOfv1EnvVarValueFrom

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigMapKeyRef** | [***interface{}**](interface{}.md) | Selects a key of a ConfigMap. +optional | [optional] [default to null]
**FieldRef** | [***interface{}**](interface{}.md) | Selects a field of the pod: supports metadata.name, metadata.namespace, &#x60;metadata.labels[&#x27;&lt;KEY&gt;&#x27;]&#x60;, &#x60;metadata.annotations[&#x27;&lt;KEY&gt;&#x27;]&#x60;, spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs. +optional | [optional] [default to null]
**ResourceFieldRef** | [***interface{}**](interface{}.md) | Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported. +optional | [optional] [default to null]
**SecretKeyRef** | [***interface{}**](interface{}.md) | Selects a key of a secret in the pod&#x27;s namespace +optional | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

