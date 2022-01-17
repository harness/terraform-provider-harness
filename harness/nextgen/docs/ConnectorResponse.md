# ConnectorResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connector** | [***ConnectorInfo**](ConnectorInfo.md) |  | [optional] [default to null]
**CreatedAt** | **int64** | This is the time at which the Connector was created. | [optional] [default to null]
**LastModifiedAt** | **int64** | This is the time at which the Connector was last modified. | [optional] [default to null]
**Status** | [***ConnectorConnectivityDetails**](ConnectorConnectivityDetails.md) |  | [optional] [default to null]
**ActivityDetails** | [***ConnectorActivityDetails**](ConnectorActivityDetails.md) |  | [optional] [default to null]
**HarnessManaged** | **bool** | This indicates if this Connector is managed by Harness or not. If True, Harness can manage and modify this Connector. | [optional] [default to null]
**GitDetails** | [***EntityGitDetails**](EntityGitDetails.md) |  | [optional] [default to null]
**EntityValidityDetails** | [***EntityValidityDetails**](EntityValidityDetails.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

