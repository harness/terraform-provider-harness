# V1ManagedFieldsEntry

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the version of this resource that this field set applies to. The format is \&quot;group/version\&quot; just like the top-level APIVersion field. It is necessary to track the version of a field set because it cannot be automatically converted. | [optional] [default to null]
**FieldsType** | **string** | FieldsType is the discriminator for the different fields format and version. There is currently only one possible value: \&quot;FieldsV1\&quot; | [optional] [default to null]
**FieldsV1** | [***AllOfv1ManagedFieldsEntryFieldsV1**](AllOfv1ManagedFieldsEntryFieldsV1.md) | FieldsV1 holds the first JSON version format as described in the \&quot;FieldsV1\&quot; type. +optional | [optional] [default to null]
**Manager** | **string** | Manager is an identifier of the workflow managing these fields. | [optional] [default to null]
**Operation** | [***AllOfv1ManagedFieldsEntryOperation**](AllOfv1ManagedFieldsEntryOperation.md) | Operation is the type of operation which lead to this ManagedFieldsEntry being created. The only valid values for this field are &#x27;Apply&#x27; and &#x27;Update&#x27;. | [optional] [default to null]
**Time** | **string** | Time is timestamp of when these fields were set. It should always be empty if Operation is &#x27;Apply&#x27; +optional | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

