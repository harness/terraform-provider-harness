# AuditEvent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuditId** | **string** | Identifier of the Audit. | [optional] [default to null]
**InsertId** | **string** | Insert Identifier of the Audit. | [default to null]
**ResourceScope** | [***AuditResourceScope**](AuditResourceScope.md) |  | [default to null]
**HttpRequestInfo** | [***HttpRequestInfo**](HttpRequestInfo.md) |  | [optional] [default to null]
**RequestMetadata** | [***RequestMetadata**](RequestMetadata.md) |  | [optional] [default to null]
**Timestamp** | **int64** |  | [default to null]
**AuthenticationInfo** | [***AuthenticationInfo**](AuthenticationInfo.md) |  | [default to null]
**Module** | **string** | Type of module associated with the Audit. | [default to null]
**Environment** | [***Environment**](Environment.md) |  | [optional] [default to null]
**Resource** | [***AuditResource**](AuditResource.md) |  | [default to null]
**YamlDiffRecord** | [***YamlDiffRecord**](YamlDiffRecord.md) |  | [optional] [default to null]
**Action** | **string** | Action type associated with the Audit. | [default to null]
**AuditEventData** | [***AuditEventData**](AuditEventData.md) |  | [optional] [default to null]
**InternalInfo** | **map[string]string** | Internal information. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

