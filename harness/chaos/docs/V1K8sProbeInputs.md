# V1K8sProbeInputs

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | **string** | Data contains the manifest/data for the resource, which need to be created it supported for create operation only | [optional] [default to null]
**FieldSelector** | **string** | fieldselector to get the resource using fields selector | [optional] [default to null]
**Group** | **string** | group of the resource | [optional] [default to null]
**LabelSelector** | **string** | labelselector to get the resource using labels selector | [optional] [default to null]
**Namespace** | **string** | namespace of the resource | [optional] [default to null]
**Operation** | **string** | Operation performed by the k8s probe it can be create, delete, present, absent | [optional] [default to null]
**Resource** | **string** | kind of resource | [optional] [default to null]
**ResourceNames** | **string** | ResourceNames to get the resources using their list of comma separated names | [optional] [default to null]
**Version** | **string** | apiversion of the resource | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

