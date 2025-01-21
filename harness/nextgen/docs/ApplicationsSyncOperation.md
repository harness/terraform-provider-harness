# ApplicationsSyncOperation

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Revision** | **string** | Revision is the revision (Git) or chart version (Helm) which to sync the application to If omitted, will use the revision specified in app spec. | [optional] [default to null]
**Prune** | **bool** |  | [optional] [default to null]
**DryRun** | **bool** |  | [optional] [default to null]
**SyncStrategy** | [***ApplicationsSyncStrategy**](applicationsSyncStrategy.md) |  | [optional] [default to null]
**Resources** | [**[]ApplicationsSyncOperationResource**](applicationsSyncOperationResource.md) |  | [optional] [default to null]
**Source** | [***ApplicationsApplicationSource**](applicationsApplicationSource.md) |  | [optional] [default to null]
**Manifests** | **[]string** |  | [optional] [default to null]
**SyncOptions** | **[]string** |  | [optional] [default to null]
**Sources** | [**[]ApplicationsApplicationSource**](applicationsApplicationSource.md) |  | [optional] [default to null]
**Revisions** | **[]string** | Revisions is the list of revision (Git) or chart version (Helm) which to sync each source in sources field for the application to If omitted, will use the revision specified in app spec. | [optional] [default to null]
**AutoHealAttemptsCount** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

