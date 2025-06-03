# ModelWorkflowRun

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **string** | Timestamp at which workflow run was created | [optional] [default to null]
**CreatedBy** | [***AllOfmodelWorkflowRunCreatedBy**](AllOfmodelWorkflowRunCreatedBy.md) | User who has created the experiment run | [optional] [default to null]
**CronSyntax** | **string** | Cron syntax of the workflow schedule | [optional] [default to null]
**ErrorResponse** | **string** | Error Response is the reason why experiment failed to run | [optional] [default to null]
**ExecutionData** | **string** | Stores all the workflow run details related to the nodes of DAG graph and chaos results of the experiments | [optional] [default to null]
**ExperimentType** | **string** | experimentType is the type of experiment run | [optional] [default to null]
**ExperimentsAwaited** | **int32** | Number of experiments awaited | [optional] [default to null]
**ExperimentsFailed** | **int32** | Number of experiments failed | [optional] [default to null]
**ExperimentsNa** | **int32** | Number of experiments which are not available | [optional] [default to null]
**ExperimentsPassed** | **int32** | Number of experiments passed | [optional] [default to null]
**ExperimentsStopped** | **int32** | Number of experiments stopped | [optional] [default to null]
**Identifiers** | [***AllOfmodelWorkflowRunIdentifiers**](AllOfmodelWorkflowRunIdentifiers.md) | Harness identifiers | [optional] [default to null]
**Infra** | [***AllOfmodelWorkflowRunInfra**](AllOfmodelWorkflowRunInfra.md) | Target infra in which the workflow will run | [optional] [default to null]
**IsCronEnabled** | **bool** | If cron is enabled or disabled | [optional] [default to null]
**IsRemoved** | **bool** | Bool value indicating if the workflow run has removed | [optional] [default to null]
**IsSingleRunCronEnabled** | **bool** | Flag to check if single run status is enabled or not | [optional] [default to null]
**NotifyID** | **string** | Notify ID of the experiment run | [optional] [default to null]
**Phase** | [***AllOfmodelWorkflowRunPhase**](AllOfmodelWorkflowRunPhase.md) | Phase of the workflow run | [optional] [default to null]
**Probe** | [**[]ModelProbeMap**](model.ProbeMap.md) | Probe object containing reference of probeIDs | [optional] [default to null]
**ResiliencyScore** | **float64** | Resiliency score of the workflow | [optional] [default to null]
**RunSequence** | **int32** | runSequence is the sequence number of experiment run | [optional] [default to null]
**SecurityGovernance** | [***AllOfmodelWorkflowRunSecurityGovernance**](AllOfmodelWorkflowRunSecurityGovernance.md) | Security Governance details of the workflow run | [optional] [default to null]
**TotalExperiments** | **int32** | Total number of experiments | [optional] [default to null]
**UpdatedAt** | **string** | Timestamp at which workflow run was last updated | [optional] [default to null]
**UpdatedBy** | [***AllOfmodelWorkflowRunUpdatedBy**](AllOfmodelWorkflowRunUpdatedBy.md) | User who has updated the workflow | [optional] [default to null]
**Weightages** | [**[]ModelWeightages**](model.Weightages.md) | Array containing weightage and name of each chaos experiment in the workflow | [optional] [default to null]
**WorkflowDescription** | **string** | Description of the workflow | [optional] [default to null]
**WorkflowID** | **string** | ID of the workflow | [optional] [default to null]
**WorkflowManifest** | **string** | Manifest of the workflow run | [optional] [default to null]
**WorkflowName** | **string** | Name of the workflow | [optional] [default to null]
**WorkflowRunID** | **string** | ID of the workflow run which is to be queried | [optional] [default to null]
**WorkflowTags** | **[]string** | Tag of the workflow | [optional] [default to null]
**WorkflowType** | **string** | Type of the workflow | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

