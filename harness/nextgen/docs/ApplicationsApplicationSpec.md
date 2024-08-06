# ApplicationsApplicationSpec

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | [***ApplicationsApplicationSource**](applicationsApplicationSource.md) |  | [optional] [default to null]
**Destination** | [***ApplicationsApplicationDestination**](applicationsApplicationDestination.md) |  | [optional] [default to null]
**Project** | **string** | Project is a reference to the project this application belongs to. The empty string means that application belongs to the &#x27;default&#x27; project. | [optional] [default to null]
**SyncPolicy** | [***ApplicationsSyncPolicy**](applicationsSyncPolicy.md) |  | [optional] [default to null]
**IgnoreDifferences** | [**[]ApplicationsResourceIgnoreDifferences**](applicationsResourceIgnoreDifferences.md) |  | [optional] [default to null]
**Info** | [**[]ApplicationsInfo**](applicationsInfo.md) |  | [optional] [default to null]
**RevisionHistoryLimit** | **string** | RevisionHistoryLimit limits the number of items kept in the application&#x27;s revision history, which is used for informational purposes as well as for rollbacks to previous versions. This should only be changed in exceptional circumstances. Setting to zero will store no history. This will reduce storage used. Increasing will increase the space used to store the history, so we do not recommend increasing it. Default is 10. | [optional] [default to null]
**Sources** | [**[]ApplicationsApplicationSource**](applicationsApplicationSource.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

