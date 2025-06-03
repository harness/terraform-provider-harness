# ModelKubernetesInfra

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterType** | [***AllOfmodelKubernetesInfraClusterType**](AllOfmodelKubernetesInfraClusterType.md) | Cluster type Indicates the type on infrastructure (Kubernetes/openshift) | [optional] [default to null]
**CreatedAt** | **string** | Timestamp when the infra was created | [optional] [default to null]
**CreatedBy** | [***AllOfmodelKubernetesInfraCreatedBy**](AllOfmodelKubernetesInfraCreatedBy.md) | User who created the infra | [optional] [default to null]
**Description** | **string** | Description of the infra | [optional] [default to null]
**EnvironmentID** | **string** | Environment ID for the infra | [optional] [default to null]
**InfraID** | **string** | ID of the infra | [optional] [default to null]
**InfraNamespace** | **string** | Namespace where the infra is being installed | [optional] [default to null]
**InfraNsExists** | **bool** | Bool value indicating whether infra ns used already exists on infra or not | [optional] [default to null]
**InfraSaExists** | **bool** | Bool value indicating whether service account used already exists on infra or not | [optional] [default to null]
**InfraScope** | [***AllOfmodelKubernetesInfraInfraScope**](AllOfmodelKubernetesInfraInfraScope.md) | Scope of the infra : ns or cluster | [optional] [default to null]
**InfraType** | [***AllOfmodelKubernetesInfraInfraType**](AllOfmodelKubernetesInfraInfraType.md) | Type of the infrastructure | [optional] [default to null]
**InstallationType** | [***AllOfmodelKubernetesInfraInstallationType**](AllOfmodelKubernetesInfraInstallationType.md) | InstallationType connector/manifest | [optional] [default to null]
**IsActive** | **bool** | Boolean value indicating if chaos infrastructure is active or not | [optional] [default to null]
**IsInfraConfirmed** | **bool** | Boolean value indicating if chaos infrastructure is confirmed or not | [optional] [default to null]
**IsRemoved** | **bool** | Boolean value indicating if chaos infrastructure is removed or not | [optional] [default to null]
**IsSecretEnabled** | **bool** | Tune secret for infra | [optional] [default to null]
**K8sConnectorID** | **string** | K8sConnectorID | [optional] [default to null]
**LastHeartbeat** | **string** | Last Heartbeat status sent by the infra | [optional] [default to null]
**LastWorkflowTimestamp** | **string** | Timestamp of the last workflow run in the infra | [optional] [default to null]
**Name** | **string** | Name of the infra | [optional] [default to null]
**NoOfSchedules** | **int32** | Number of schedules created in the infra | [optional] [default to null]
**NoOfWorkflows** | **int32** | Number of workflows run in the infra | [optional] [default to null]
**PlatformName** | **string** | Infra Platform Name eg. GKE,AWS, Others | [optional] [default to null]
**RunAsGroup** | **int32** | set the user group for security context in pod | [optional] [default to null]
**RunAsUser** | **int32** | set the user for security context in pod | [optional] [default to null]
**ServiceAccount** | **string** | Name of service account used by infra | [optional] [default to null]
**StartTime** | **string** | Timestamp when the infra got connected | [optional] [default to null]
**Tags** | **[]string** | Tags of the infra | [optional] [default to null]
**Token** | **string** | Token used to verify and retrieve the infra manifest | [optional] [default to null]
**UpdateStatus** | [***AllOfmodelKubernetesInfraUpdateStatus**](AllOfmodelKubernetesInfraUpdateStatus.md) | update status of infra | [optional] [default to null]
**UpdatedAt** | **string** | Timestamp when the infra was last updated | [optional] [default to null]
**UpdatedBy** | [***AllOfmodelKubernetesInfraUpdatedBy**](AllOfmodelKubernetesInfraUpdatedBy.md) | User who has updated the infra | [optional] [default to null]
**Upgrade** | [***AllOfmodelKubernetesInfraUpgrade**](AllOfmodelKubernetesInfraUpgrade.md) | Upgrade struct for the chaos infrastructure | [optional] [default to null]
**Version** | **string** | Version of the infra | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

