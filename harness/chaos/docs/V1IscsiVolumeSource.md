# V1IscsiVolumeSource

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChapAuthDiscovery** | **bool** | whether support iSCSI Discovery CHAP authentication +optional | [optional] [default to null]
**ChapAuthSession** | **bool** | whether support iSCSI Session CHAP authentication +optional | [optional] [default to null]
**FsType** | **string** | Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: \&quot;ext4\&quot;, \&quot;xfs\&quot;, \&quot;ntfs\&quot;. Implicitly inferred to be \&quot;ext4\&quot; if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi TODO: how do we prevent errors in the filesystem from compromising the machine +optional | [optional] [default to null]
**InitiatorName** | **string** | Custom iSCSI Initiator Name. If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface &lt;target portal&gt;:&lt;volume name&gt; will be created for the connection. +optional | [optional] [default to null]
**Iqn** | **string** | Target iSCSI Qualified Name. | [optional] [default to null]
**IscsiInterface** | **string** | iSCSI Interface Name that uses an iSCSI transport. Defaults to &#x27;default&#x27; (tcp). +optional | [optional] [default to null]
**Lun** | **int32** | iSCSI Target Lun number. | [optional] [default to null]
**Portals** | **[]string** | iSCSI Target Portal List. The portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260). +optional | [optional] [default to null]
**ReadOnly** | **bool** | ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. +optional | [optional] [default to null]
**SecretRef** | [***AllOfv1IscsiVolumeSourceSecretRef**](AllOfv1IscsiVolumeSourceSecretRef.md) | CHAP Secret for iSCSI target and initiator authentication +optional | [optional] [default to null]
**TargetPortal** | **string** | iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260). | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

