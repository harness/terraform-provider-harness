# ChaoshubresourcesChaosHubResource

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountID** | **string** |  | [default to null]
**Category** | **string** | Category of the fault (Kubernetes, KubernetesV2, Linux, Windows, PCF etc) | [optional] [default to null]
**ChaosType** | **string** | ChaosType | [optional] [default to null]
**ChartServiceVersion** | **string** | chartserviceversion yml in encoded form (chartserviceversion.yaml) | [optional] [default to null]
**CreatedAt** | **int32** |  | [optional] [default to null]
**CreatedBy** | **string** |  | [optional] [default to null]
**Description** | **string** | Description of the resource | [optional] [default to null]
**DisplayName** | **string** | Display name of the resource | [optional] [default to null]
**Engine** | **string** | ChaosEngine yml in encoded form (engine.yaml) | [optional] [default to null]
**Experiment** | **string** | experiment yml in encoded form (experiment.yaml) | [optional] [default to null]
**ExperimentV2** | **string** | experiment-v2 yml in encoded form (experiment-v2.yaml) | [optional] [default to null]
**Fault** | **string** | Fault yml in encoded form (fault.yml) | [optional] [default to null]
**HubID** | **string** | ID of the chaos hub | [optional] [default to null]
**Id** | **string** | Mongo ID (primary key) | [optional] [default to null]
**Identity** | **string** | TODO: identity changes move to chaosHub Unique identifier (human-readable) immutable Initially it will be same as name Should be unique at hub level | [optional] [default to null]
**Infras** | **[]string** | Infras represents supported infrastructures | [optional] [default to null]
**IsDefaultHub** | **bool** | IsDefaultHub represents if it is a default hub | [optional] [default to null]
**IsRemoved** | **bool** |  | [default to null]
**IsTemplatised** | **bool** | IsTemplatised denotes if template is available for the fault | [optional] [default to null]
**K8SFault** | **string** | K8sFault yml in encoded form (k8s-fault.yaml) | [optional] [default to null]
**Keywords** | **[]string** | Keyword of the resource (kubernetes, VMWare etc) | [optional] [default to null]
**Kind** | **string** | Kind of the resource | [optional] [default to null]
**Links** | [**[]ChaoshubresourcesLink**](chaoshubresources.Link.md) | Links are array of Link | [optional] [default to null]
**Name** | **string** | name of the resource | [optional] [default to null]
**OrgID** | **string** |  | [optional] [default to null]
**PermissionsRequired** | [***AllOfchaoshubresourcesChaosHubResourcePermissionsRequired**](AllOfchaoshubresourcesChaosHubResourcePermissionsRequired.md) | PermissionsRequired represents the level of permissions required for the resource | [optional] [default to null]
**Plan** | **[]string** | Plan | [optional] [default to null]
**Platforms** | **[]string** | Platforms supported (GKE, Minikube, EKS, AKS etc) | [optional] [default to null]
**ProjectID** | **string** |  | [optional] [default to null]
**ResourceType** | **string** | Type: fault or experiment or probes | [optional] [default to null]
**Tags** | **[]string** | tags for the resource | [optional] [default to null]
**Template** | **string** | template yaml in encoded form (template.yaml) | [optional] [default to null]
**UpdatedAt** | **int32** |  | [optional] [default to null]
**UpdatedBy** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

