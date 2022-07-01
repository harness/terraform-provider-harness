# AuditFilterProperties

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scopes** | [**[]AuditResourceScope**](AuditResourceScope.md) | List of Resource Scopes | [optional] [default to null]
**Resources** | [**[]AuditResource**](AuditResource.md) | List of Resources | [optional] [default to null]
**Modules** | **[]string** | List of Module Types | [optional] [default to null]
**Actions** | **[]string** | List of Actions | [optional] [default to null]
**Environments** | [**[]Environment**](Environment.md) | List of Environments | [optional] [default to null]
**Principals** | [**[]AuditPrincipal**](AuditPrincipal.md) | List of Principals | [optional] [default to null]
**StaticFilter** | **string** | Pre-defined Filter | [optional] [default to null]
**StartTime** | **int64** | Used to specify a start time for retrieving Audit events that occurred at or after the time indicated. | [optional] [default to null]
**EndTime** | **int64** | Used to specify the end time for retrieving Audit events that occurred at or before the time indicated. | [optional] [default to null]
**Tags** | **map[string]string** | Filter tags as a key-value pair. | [optional] [default to null]
**FilterType** | **string** | This specifies the corresponding Entity of the filter. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

