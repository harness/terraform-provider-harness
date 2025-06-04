# V1Volume

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsElasticBlockStore** | [***V1AwsElasticBlockStoreVolumeSource**](v1.AWSElasticBlockStoreVolumeSource.md) |  | [optional] [default to null]
**AzureDisk** | [***V1AzureDiskVolumeSource**](v1.AzureDiskVolumeSource.md) |  | [optional] [default to null]
**AzureFile** | [***V1AzureFileVolumeSource**](v1.AzureFileVolumeSource.md) |  | [optional] [default to null]
**Cephfs** | [***V1CephFsVolumeSource**](v1.CephFSVolumeSource.md) |  | [optional] [default to null]
**Cinder** | [***V1CinderVolumeSource**](v1.CinderVolumeSource.md) |  | [optional] [default to null]
**ConfigMap** | [***V1ConfigMapVolumeSource**](v1.ConfigMapVolumeSource.md) |  | [optional] [default to null]
**Csi** | [***V1CsiVolumeSource**](v1.CSIVolumeSource.md) |  | [optional] [default to null]
**DownwardAPI** | [***V1DownwardApiVolumeSource**](v1.DownwardAPIVolumeSource.md) |  | [optional] [default to null]
**EmptyDir** | [***V1EmptyDirVolumeSource**](v1.EmptyDirVolumeSource.md) |  | [optional] [default to null]
**Ephemeral** | [***V1EphemeralVolumeSource**](v1.EphemeralVolumeSource.md) |  | [optional] [default to null]
**Fc** | [***V1FcVolumeSource**](v1.FCVolumeSource.md) |  | [optional] [default to null]
**FlexVolume** | [***V1FlexVolumeSource**](v1.FlexVolumeSource.md) |  | [optional] [default to null]
**Flocker** | [***V1FlockerVolumeSource**](v1.FlockerVolumeSource.md) |  | [optional] [default to null]
**GcePersistentDisk** | [***V1GcePersistentDiskVolumeSource**](v1.GCEPersistentDiskVolumeSource.md) |  | [optional] [default to null]
**GitRepo** | [***V1GitRepoVolumeSource**](v1.GitRepoVolumeSource.md) |  | [optional] [default to null]
**Glusterfs** | [***V1GlusterfsVolumeSource**](v1.GlusterfsVolumeSource.md) |  | [optional] [default to null]
**HostPath** | [***V1HostPathVolumeSource**](v1.HostPathVolumeSource.md) |  | [optional] [default to null]
**Iscsi** | [***V1IscsiVolumeSource**](v1.ISCSIVolumeSource.md) |  | [optional] [default to null]
**Name** | **string** | name of the volume. Must be a DNS_LABEL and unique within the pod. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names | [optional] [default to null]
**Nfs** | [***V1NfsVolumeSource**](v1.NFSVolumeSource.md) |  | [optional] [default to null]
**PersistentVolumeClaim** | [***V1PersistentVolumeClaimVolumeSource**](v1.PersistentVolumeClaimVolumeSource.md) |  | [optional] [default to null]
**PhotonPersistentDisk** | [***V1PhotonPersistentDiskVolumeSource**](v1.PhotonPersistentDiskVolumeSource.md) |  | [optional] [default to null]
**PortworxVolume** | [***V1PortworxVolumeSource**](v1.PortworxVolumeSource.md) |  | [optional] [default to null]
**Projected** | [***V1ProjectedVolumeSource**](v1.ProjectedVolumeSource.md) |  | [optional] [default to null]
**Quobyte** | [***V1QuobyteVolumeSource**](v1.QuobyteVolumeSource.md) |  | [optional] [default to null]
**Rbd** | [***V1RbdVolumeSource**](v1.RBDVolumeSource.md) |  | [optional] [default to null]
**ScaleIO** | [***V1ScaleIoVolumeSource**](v1.ScaleIOVolumeSource.md) |  | [optional] [default to null]
**Secret** | [***V1SecretVolumeSource**](v1.SecretVolumeSource.md) |  | [optional] [default to null]
**Storageos** | [***V1StorageOsVolumeSource**](v1.StorageOSVolumeSource.md) |  | [optional] [default to null]
**VsphereVolume** | [***V1VsphereVirtualDiskVolumeSource**](v1.VsphereVirtualDiskVolumeSource.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

