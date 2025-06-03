# {{classname}}

All URIs are relative to */api/manager*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddNote**](DefaultApi.md#AddNote) | **Post** /rest/v2/notes | Add a resource note
[**BulkExperimentDelete**](DefaultApi.md#BulkExperimentDelete) | **Post** /rest/v2/bulkaction/bulkexperimentdelete | Deletes given experiments
[**BulkExperimentTagAdd**](DefaultApi.md#BulkExperimentTagAdd) | **Post** /rest/v2/bulkaction/bulkexperimenttagsadd | Adds tags in given experiments
[**CanRetryExperimentCreation**](DefaultApi.md#CanRetryExperimentCreation) | **Get** /rest/v2/applicationmaps/{applicationmapid}/canretry | can retry or not chaos experiment creation for the given target network map
[**CreateAction**](DefaultApi.md#CreateAction) | **Post** /rest/actions | Create a new action
[**CreateActionTemplate**](DefaultApi.md#CreateActionTemplate) | **Post** /rest/templates/actions | Creates the action templates in a hub based on tag
[**CreateChaosExperimentExecutionNode**](DefaultApi.md#CreateChaosExperimentExecutionNode) | **Put** /internal/execution-node/{experimentId}/{experimentRunId} | Create chaos execution node
[**CreateChaosExperimentPipeline**](DefaultApi.md#CreateChaosExperimentPipeline) | **Post** /rest/v2/experiment/bulk/run | Create pipeline to run experiments in bulk
[**CreateChaosHub**](DefaultApi.md#CreateChaosHub) | **Post** /rest/hubs | Create chaos hub
[**CreateFaultTemplate**](DefaultApi.md#CreateFaultTemplate) | **Post** /rest/templates/faults | Create fault templates in a hub based on tag
[**CreateGamedayRunV2**](DefaultApi.md#CreateGamedayRunV2) | **Post** /rest/v2/gameday/{gamedayId}/run | Create a chaos v2 gameday run
[**CreateGamedayV2**](DefaultApi.md#CreateGamedayV2) | **Post** /rest/v2/gameday | Create a chaos v2 gameday
[**CreateInputSet**](DefaultApi.md#CreateInputSet) | **Post** /rest/v2/experiments/{experimentId}/inputsets | Create an input set
[**CreateProbe**](DefaultApi.md#CreateProbe) | **Post** /rest/v2/probes | Create a new probe
[**CreateProbeTemplate**](DefaultApi.md#CreateProbeTemplate) | **Post** /rest/templates/probes | Creates the probe templates in a hub based on tag
[**CreateRecommendation**](DefaultApi.md#CreateRecommendation) | **Post** /rest/recommendations/action/create | Create new experiment
[**CreateV2Onboarding**](DefaultApi.md#CreateV2Onboarding) | **Post** /rest/v2/onboarding | Create V2 Onboarding
[**DeleteAction**](DefaultApi.md#DeleteAction) | **Delete** /rest/actions/{identity} | Delete an action
[**DeleteActionTemplate**](DefaultApi.md#DeleteActionTemplate) | **Delete** /rest/templates/actions/{identity} | Deletes action template
[**DeleteChaosV2Experiment**](DefaultApi.md#DeleteChaosV2Experiment) | **Delete** /rest/v2/experiment/{experimentId} | Delete Chaos V2 experiment
[**DeleteFaultTemplate**](DefaultApi.md#DeleteFaultTemplate) | **Delete** /rest/templates/faults/{faultName} | Delete the fault templates in a hub based on tag
[**DeleteGamedayV2**](DefaultApi.md#DeleteGamedayV2) | **Delete** /rest/v2/gameday/{gamedayId} | Get a chaos v2 Gameday
[**DeleteInputSet**](DefaultApi.md#DeleteInputSet) | **Delete** /rest/v2/experiments/{experimentId}/inputsets/{inputsetId} | Delete an input set
[**DeleteProbe**](DefaultApi.md#DeleteProbe) | **Delete** /rest/v2/probes/{probeId} | Delete a probe
[**DeleteProbeTemplate**](DefaultApi.md#DeleteProbeTemplate) | **Delete** /rest/templates/probes/{identity} | Deletes probe template
[**DeleteRecommendation**](DefaultApi.md#DeleteRecommendation) | **Delete** /rest/recommendations | Delete recommendation from db
[**DeleteTargetNetworkMap**](DefaultApi.md#DeleteTargetNetworkMap) | **Delete** /rest/v2/applicationmaps/{applicationmapid} | Delete application network map with chaos context
[**EnableProbe**](DefaultApi.md#EnableProbe) | **Post** /rest/v2/probes/{probeId}/enable | Enable and disable probes across experiments
[**ExperimentExecutionNodeDetails**](DefaultApi.md#ExperimentExecutionNodeDetails) | **Get** /internal/execution-node/{experimentId}/{experimentRunId} | Get chaos execution node
[**GetAccountServiceDetails**](DefaultApi.md#GetAccountServiceDetails) | **Get** /rest/service/{accountID} | Get service usage details for account in the requested time frame
[**GetAccountServiceUsageStats**](DefaultApi.md#GetAccountServiceUsageStats) | **Get** /rest/service/stats/{accountID} | Get the service usage stats related to requested account grouped by day or month
[**GetAction**](DefaultApi.md#GetAction) | **Get** /rest/actions/{identity} | Get an action
[**GetActionManifest**](DefaultApi.md#GetActionManifest) | **Get** /rest/actions/manifest/{identity} | Get an action manifest
[**GetActionTemplate**](DefaultApi.md#GetActionTemplate) | **Get** /rest/templates/actions/{identity} | Get the action template in a hub based on action ref
[**GetActionTemplateRevisionDifference**](DefaultApi.md#GetActionTemplateRevisionDifference) | **Get** /rest/templates/actions/{identity}/compare | Get the difference between 2 revisions of action template
[**GetChaosDashboard**](DefaultApi.md#GetChaosDashboard) | **Get** /rest/chaosDashboards | Get chaos dashboard
[**GetChaosExperimentRunReport**](DefaultApi.md#GetChaosExperimentRunReport) | **Get** /rest/chaos-experiment-run/report/{experimentRunId}/{notifyId} | Generate and return kubernetesV1 chaos experiment run report
[**GetChaosHub**](DefaultApi.md#GetChaosHub) | **Get** /rest/hubs/{hubIdentity} | Get chaos hub based on given filters
[**GetChaosPipelineExecution**](DefaultApi.md#GetChaosPipelineExecution) | **Get** /rest/v2/chaos-pipeline/{experimentId}/{experimentRunId} | Get a chaos pipeline execution
[**GetChaosPipelineStepDetails**](DefaultApi.md#GetChaosPipelineStepDetails) | **Get** /rest/v2/chaos-pipeline/step/{experimentId}/{experimentRunId}/{stepName} | Get a chaos pipeline step execution
[**GetChaosV2Experiment**](DefaultApi.md#GetChaosV2Experiment) | **Get** /rest/v2/experiments/{experimentId} | Get a chaos v2 experiment
[**GetChaosV2ExperimentRun**](DefaultApi.md#GetChaosV2ExperimentRun) | **Get** /rest/v2/experiments/{experimentId}/run | Get a chaos v2 experiment run
[**GetChaosV2ExperimentRunInternalAPI**](DefaultApi.md#GetChaosV2ExperimentRunInternalAPI) | **Post** /internal/v2/experiments/{experimentId}/run/{notifyId} | Get the chaos v2 experiment run internal API
[**GetChaosV2ExperimentVariables**](DefaultApi.md#GetChaosV2ExperimentVariables) | **Get** /rest/v2/experiments/{experimentId}/variables | Get a chaos v2 experiment
[**GetConnectorForInfra**](DefaultApi.md#GetConnectorForInfra) | **Get** /rest/v2/infrastructure/{identity}/connector | Get Connector For Infra
[**GetExperimentHelperImageVersion**](DefaultApi.md#GetExperimentHelperImageVersion) | **Get** /experimentHelperImageVersion | Get experiment helper image version
[**GetExperimentRunTimelineView**](DefaultApi.md#GetExperimentRunTimelineView) | **Get** /rest/v2/experiments/timeline/run/{experimentId} | Get a chaos v2 experiment timeline run
[**GetExperimentRunsOverviewStats**](DefaultApi.md#GetExperimentRunsOverviewStats) | **Get** /rest/overview/experiment-stats | Get chaos experiment run overview stats
[**GetFaultTemplate**](DefaultApi.md#GetFaultTemplate) | **Get** /rest/templates/faults/{faultName} | Lists all the fault templates in a hub based on tag
[**GetFaultTemplateRevisionDifference**](DefaultApi.md#GetFaultTemplateRevisionDifference) | **Get** /rest/templates/faults/{faultName}/compare | Get the difference between 2 revisions of a fault template
[**GetGamedayRunV2**](DefaultApi.md#GetGamedayRunV2) | **Get** /rest/v2/gameday/{gamedayId}/run/{gamedayRunId} | Fetch a chaos v2 gameday run
[**GetGamedayV2**](DefaultApi.md#GetGamedayV2) | **Get** /rest/v2/gameday/{gamedayId} | Get a chaos v2 Gameday
[**GetImageRegistry**](DefaultApi.md#GetImageRegistry) | **Get** /rest/imageRegistry | Get image registry
[**GetInfraToken**](DefaultApi.md#GetInfraToken) | **Get** /rest/v2/infrastructures/{infrastructureIdentity}/token | Get a v2 infra token
[**GetInputSet**](DefaultApi.md#GetInputSet) | **Get** /rest/v2/experiments/{experimentId}/inputsets/{inputsetId} | Get the input set in an experiment
[**GetNote**](DefaultApi.md#GetNote) | **Get** /rest/v2/notes | Get a chaos resource Note
[**GetOnboardingExperiments**](DefaultApi.md#GetOnboardingExperiments) | **Get** /rest/v2/onboarding/{onboardingid}/experiments | Get V2 Onboarding experiments
[**GetOverallServiceUsageStats**](DefaultApi.md#GetOverallServiceUsageStats) | **Get** /rest/service/overall/stats/{accountID} | Get the overall service usage stats by type related to requested account
[**GetProbe**](DefaultApi.md#GetProbe) | **Get** /rest/v2/probes/{probeId} | Get a probe
[**GetProbeManifest**](DefaultApi.md#GetProbeManifest) | **Get** /rest/v2/probes/manifest/{probeId} | Get a probe
[**GetProbeTemplate**](DefaultApi.md#GetProbeTemplate) | **Get** /rest/templates/probes/{identity} | Get the probe template in a hub based on probe ref
[**GetRecommendation**](DefaultApi.md#GetRecommendation) | **Get** /rest/recommendations | Get recommendation details
[**GetResourceUsage**](DefaultApi.md#GetResourceUsage) | **Get** /rest/usage | Get resource usage
[**GetSGPTemplateByGenAI**](DefaultApi.md#GetSGPTemplateByGenAI) | **Post** /rest/genai/sgp/generate | Get security governance conditions template
[**GetServiceResponse**](DefaultApi.md#GetServiceResponse) | **Get** /rest/v2/applicationmaps/{applicationmapid}/targetservices/{targetserviceid} | Get target discovered service with chaos context
[**GetServiceUsageReport**](DefaultApi.md#GetServiceUsageReport) | **Get** /rest/service/report/{accountID} | Generates service usage report in csv format
[**GetTargetNetworkMap**](DefaultApi.md#GetTargetNetworkMap) | **Get** /rest/v2/applicationmaps/{applicationmapid} | Get target network map with chaos context
[**GetV2InfrastructureYaml**](DefaultApi.md#GetV2InfrastructureYaml) | **Post** /rest/v2/infrastructure/yaml | Preview v2 infra Yaml
[**GetV2Onboarding**](DefaultApi.md#GetV2Onboarding) | **Get** /rest/v2/onboarding/{onboardingid} | Get V2 Onboarding
[**ImportAction**](DefaultApi.md#ImportAction) | **Post** /rest/actions/import | Import a new action
[**ImportProbe**](DefaultApi.md#ImportProbe) | **Post** /rest/v2/probes/import | Import a new probe
[**ListActionTemplate**](DefaultApi.md#ListActionTemplate) | **Get** /rest/templates/actions | Lists all the action templates in a hub based on tag
[**ListActionTemplateRevisions**](DefaultApi.md#ListActionTemplateRevisions) | **Get** /rest/templates/actions/{identity}/revisions | Lists all the revision of a fault template in a hub
[**ListActions**](DefaultApi.md#ListActions) | **Get** /rest/actions | List actions with filtering options
[**ListApplication**](DefaultApi.md#ListApplication) | **Get** /rest/v2/infrastructures/{infrastructureIdentity}/applications | List all applications for a given infra
[**ListChaosEnabledInfraV2**](DefaultApi.md#ListChaosEnabledInfraV2) | **Post** /rest/v2/infrastructures/chaos-enabled | List a new v2 infra
[**ListChaosHub**](DefaultApi.md#ListChaosHub) | **Get** /rest/hubs | Lists chaos hubs based on given filters
[**ListChaosV2Experiment**](DefaultApi.md#ListChaosV2Experiment) | **Get** /rest/v2/experiment | Get list of Chaos V2 Experiments
[**ListFault**](DefaultApi.md#ListFault) | **Get** /rest/hubs/faults | Lists faults in a chaos hub based on given filters
[**ListFaultTemplate**](DefaultApi.md#ListFaultTemplate) | **Get** /rest/templates/faults | Lists all the fault templates in a hub based on tag
[**ListFaultTemplateRevisions**](DefaultApi.md#ListFaultTemplateRevisions) | **Get** /rest/templates/faults/{faultName}/revisions | Lists all the revision of a fault template in a hub
[**ListFunction**](DefaultApi.md#ListFunction) | **Get** /rest/v2/infrastructures/{infrastructureIdentity}/applications/{applicationIdentity}/functions | List instrumented functions for given application
[**ListGamedayRunV2**](DefaultApi.md#ListGamedayRunV2) | **Get** /rest/v2/gameday/{gamedayId}/runs | Fetch chaos v2 gameday runs
[**ListGamedayV2**](DefaultApi.md#ListGamedayV2) | **Get** /rest/v2/gamedays | Get list of Chaos V2 Gamedays
[**ListHarnessInfra**](DefaultApi.md#ListHarnessInfra) | **Get** /rest/v2/harness-infrastructures | List harness infras
[**ListInputSet**](DefaultApi.md#ListInputSet) | **Get** /rest/v2/experiments/{experimentId}/inputsets | Get the list of input sets in an experiment
[**ListK8sInfrasV2**](DefaultApi.md#ListK8sInfrasV2) | **Post** /rest/v2/list-infras | Get list of active unused connector
[**ListProbeTemplate**](DefaultApi.md#ListProbeTemplate) | **Get** /rest/templates/probes | Lists all the probe templates in a hub based on tag
[**ListProbes**](DefaultApi.md#ListProbes) | **Get** /rest/v2/probes | List probes with filtering options
[**ListRecommendations**](DefaultApi.md#ListRecommendations) | **Post** /rest/recommendations | List recommendations
[**ListService**](DefaultApi.md#ListService) | **Get** /rest/v2/applicationmaps/{applicationmapid}/targetservices | List target discovered service with chaos context
[**ListTargetNetworkMaps**](DefaultApi.md#ListTargetNetworkMaps) | **Post** /rest/v2/applicationmaps | List target network maps with chaos context
[**ListV2Onboarding**](DefaultApi.md#ListV2Onboarding) | **Get** /rest/v2/onboarding | Get V2 Onboarding
[**ListVariablesInActionTemplate**](DefaultApi.md#ListVariablesInActionTemplate) | **Get** /rest/templates/actions/{identity}/variables | Get the list of variables in a fault template
[**ListVariablesInFaultTemplate**](DefaultApi.md#ListVariablesInFaultTemplate) | **Get** /rest/templates/faults/{faultName}/variables | Get the list of variables in a fault template
[**ListVariablesInProbeTemplate**](DefaultApi.md#ListVariablesInProbeTemplate) | **Get** /rest/templates/probes/{identity}/variables | Get the list of variables in a fault template
[**OnboardingConfirmDiscovery**](DefaultApi.md#OnboardingConfirmDiscovery) | **Post** /rest/v2/onboarding-confirm-discovery/{onboardingid} | confirm discovery step in manual onboarding
[**OnboardingConfirmExperimentCreation**](DefaultApi.md#OnboardingConfirmExperimentCreation) | **Post** /rest/v2/onboarding-confirm-experiment-creation/{onboardingid} | confirm experiment creation step in manual onboarding
[**OnboardingConfirmExperimentRun**](DefaultApi.md#OnboardingConfirmExperimentRun) | **Post** /rest/v2/onboarding-confirm-experiment-run/{onboardingid} | confirm experiment run step in manual onboarding
[**OnboardingConfirmNetworkMap**](DefaultApi.md#OnboardingConfirmNetworkMap) | **Post** /rest/v2/onboarding-confirm-networkmap/{onboardingid} | confirm network map creation step in manual onboarding
[**RecommendationEvent**](DefaultApi.md#RecommendationEvent) | **Post** /internal/recommendations/event | process the recommendation event
[**RetryExperimentCreation**](DefaultApi.md#RetryExperimentCreation) | **Post** /rest/v2/applicationmaps/{applicationmapid}/retry | retry chaos experiment creation for the given target network map
[**RunChaosV2Experiment**](DefaultApi.md#RunChaosV2Experiment) | **Post** /rest/v2/experiments/{experimentId}/run | Run a chaos v2 experiment
[**RunChaosV2InternalAPI**](DefaultApi.md#RunChaosV2InternalAPI) | **Post** /internal/v2/experiments/{experimentId}/run | Run a chaos v2 experiment internal API
[**RunRecommendation**](DefaultApi.md#RunRecommendation) | **Post** /rest/recommendations/action/run | Run the recommended experiment
[**SaveChaosV2Experiment**](DefaultApi.md#SaveChaosV2Experiment) | **Post** /rest/v2/experiment | Save a chaos v2 experiment
[**StopChaosV2Experiment**](DefaultApi.md#StopChaosV2Experiment) | **Post** /rest/v2/experiment/{experimentId}/stop | Stop Chaos V2 experiment
[**StopOnboardingExperiments**](DefaultApi.md#StopOnboardingExperiments) | **Post** /rest/v2/onboarding/{onboardingid}/stop | Stop V2 Onboarding experiments
[**UpdateAction**](DefaultApi.md#UpdateAction) | **Put** /rest/actions/{identity} | Update a new action
[**UpdateActionTemplate**](DefaultApi.md#UpdateActionTemplate) | **Put** /rest/templates/actions/{identity} | Updates the action templates in a hub
[**UpdateChaosExperimentExecutionNode**](DefaultApi.md#UpdateChaosExperimentExecutionNode) | **Post** /internal/execution-node/{experimentId}/{experimentRunId}/{name} | Update chaos execution node
[**UpdateChaosHub**](DefaultApi.md#UpdateChaosHub) | **Put** /rest/hubs/{hubIdentity} | Update chaos hub
[**UpdateChaosV2CronExperiment**](DefaultApi.md#UpdateChaosV2CronExperiment) | **Post** /rest/v2/experiment/cron | Update a chaos v2 cron experiment
[**UpdateEmissary**](DefaultApi.md#UpdateEmissary) | **Post** /rest/v2/infrastructures/{infrastructureIdentity}/updateemissary | Update emissary endpoint
[**UpdateFaultTemplate**](DefaultApi.md#UpdateFaultTemplate) | **Put** /rest/templates/faults/{faultName} | Update the fault templates in a hub based on tag
[**UpdateGamedayRunPrerequisitesV2**](DefaultApi.md#UpdateGamedayRunPrerequisitesV2) | **Put** /rest/v2/gameday/{gamedayId}/run_checklist/{gamedayRunId} | Update a chaos v2 gameday run
[**UpdateGamedayRunStakeHolderActionsV2**](DefaultApi.md#UpdateGamedayRunStakeHolderActionsV2) | **Put** /rest/v2/gameday/{gamedayId}/run_approval/{gamedayRunId} | Update a chaos v2 gameday run
[**UpdateGamedayRunV2**](DefaultApi.md#UpdateGamedayRunV2) | **Put** /rest/v2/gameday/{gamedayId}/run/{gamedayRunId} | Update a chaos v2 gameday run
[**UpdateGamedayV2**](DefaultApi.md#UpdateGamedayV2) | **Put** /rest/v2/gameday | Update a chaos v2 gameday
[**UpdateInputSet**](DefaultApi.md#UpdateInputSet) | **Put** /rest/v2/experiments/{experimentId}/inputsets/{inputsetId} | Updates an input set
[**UpdateNote**](DefaultApi.md#UpdateNote) | **Patch** /rest/v2/note | Update a resource note
[**UpdateProbe**](DefaultApi.md#UpdateProbe) | **Put** /rest/v2/probes/{probeId} | Update a new probe
[**UpdateProbeTemplate**](DefaultApi.md#UpdateProbeTemplate) | **Put** /rest/templates/probes/{identity} | Updates the probe templates in a hub
[**UpdateRecommendationStatus**](DefaultApi.md#UpdateRecommendationStatus) | **Post** /rest/recommendations/status | Update the recommendation status

# **AddNote**
> TypesCreateNoteResponse AddNote(ctx, body, accountIdentifier, organizationIdentifier, projectIdentifier)
Add a resource note

Add a resource note

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TypesCreateNoteRequest**](TypesCreateNoteRequest.md)| Create Gameday | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 

### Return type

[**TypesCreateNoteResponse**](types.CreateNoteResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BulkExperimentDelete**
> BulkactionBulkDeleteExperimetsResponse BulkExperimentDelete(ctx, body, accountIdentifier, organizationIdentifier, projectIdentifier)
Deletes given experiments

Deletes given experiments

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BulkactionBulkDeleteExperimetsInput**](BulkactionBulkDeleteExperimetsInput.md)| Retry experiment creation request | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 

### Return type

[**BulkactionBulkDeleteExperimetsResponse**](bulkaction.BulkDeleteExperimetsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BulkExperimentTagAdd**
> BulkactionBulkAddTagsInExperimetsResponse BulkExperimentTagAdd(ctx, body, accountIdentifier, organizationIdentifier, projectIdentifier)
Adds tags in given experiments

Adds tags in given experiments

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BulkactionBulkAddTagsInExperimetsInput**](BulkactionBulkAddTagsInExperimetsInput.md)| Retry experiment creation request | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 

### Return type

[**BulkactionBulkAddTagsInExperimetsResponse**](bulkaction.BulkAddTagsInExperimetsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CanRetryExperimentCreation**
> NetworkmapCanRetryExperimentCreationResponse CanRetryExperimentCreation(ctx, accountIdentifier, organizationIdentifier, projectIdentifier, environmentIdentifier, infraId, applicationmapid)
can retry or not chaos experiment creation for the given target network map

can retry or not chaos experiment creation for the given target network map

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **environmentIdentifier** | **string**| environment id that want to access the resource | 
  **infraId** | **string**| infra id that want to access the resource | 
  **applicationmapid** | **string**| Application map ID | 

### Return type

[**NetworkmapCanRetryExperimentCreationResponse**](networkmap.CanRetryExperimentCreationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateAction**
> GithubComHarnessHceSaasGraphqlServerPkgActionsAction CreateAction(ctx, body, accountIdentifier, organizationIdentifier, projectIdentifier)
Create a new action

Create a new action

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GithubComHarnessHceSaasGraphqlServerPkgActionsAction**](GithubComHarnessHceSaasGraphqlServerPkgActionsAction.md)| action configuration | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 

### Return type

[**GithubComHarnessHceSaasGraphqlServerPkgActionsAction**](github_com_harness_hce-saas_graphql_server_pkg_actions.Action.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateActionTemplate**
> ChaosfaulttemplateActionTemplate CreateActionTemplate(ctx, body, accountIdentifier, organizationIdentifier, projectIdentifier)
Creates the action templates in a hub based on tag

Creates action templates in a hub based on tag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ChaosfaulttemplateActionTemplate**](ChaosfaulttemplateActionTemplate.md)| action configuration | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 

### Return type

[**ChaosfaulttemplateActionTemplate**](chaosfaulttemplate.ActionTemplate.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateChaosExperimentExecutionNode**
> string CreateChaosExperimentExecutionNode(ctx, body, experimentId, experimentRunId)
Create chaos execution node

Create chaos execution node

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]ExecutionChaosExecutionNode**](execution.ChaosExecutionNode.md)| Create chaos execution node | 
  **experimentId** | **string**| experimentId to be fetched | 
  **experimentRunId** | **string**| experimentRunId to be fetched | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateChaosExperimentPipeline**
> PipelinesBulkExperimentRunResponse CreateChaosExperimentPipeline(ctx, body, accountIdentifier, organizationIdentifier, projectIdentifier)
Create pipeline to run experiments in bulk

Create pipeline to run experiments in bulk

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PipelinesChaosPipelineInput**](PipelinesChaosPipelineInput.md)| object containing pipeline id along with metadata | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 

### Return type

[**PipelinesBulkExperimentRunResponse**](pipelines.BulkExperimentRunResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateChaosHub**
> Chaoshubv2GetHubResponse CreateChaosHub(ctx, body, accountIdentifier, organizationIdentifier, projectIdentifier)
Create chaos hub

Create chaos hub

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Chaoshubv2CreateHubRequest**](Chaoshubv2CreateHubRequest.md)| create hub request | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 

### Return type

[**Chaoshubv2GetHubResponse**](chaoshubv2.GetHubResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateFaultTemplate**
> ChaosfaulttemplateCreateFaultTemplateResponse CreateFaultTemplate(ctx, body, accountIdentifier, organizationIdentifier, projectIdentifier, hubIdentity)
Create fault templates in a hub based on tag

Create fault templates in a hub based on tag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ChaosfaulttemplateCreateFaultTemplateRequest**](ChaosfaulttemplateCreateFaultTemplateRequest.md)| create fault request | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **hubIdentity** | **string**| chaos hub identity | 

### Return type

[**ChaosfaulttemplateCreateFaultTemplateResponse**](chaosfaulttemplate.CreateFaultTemplateResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateGamedayRunV2**
> TypesCreateGamedayRunResponse CreateGamedayRunV2(ctx, accountIdentifier, organizationIdentifier, projectIdentifier, gamedayId)
Create a chaos v2 gameday run

Create a chaos v2 gameday run

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **gamedayId** | **string**| gamedayId to be run | 

### Return type

[**TypesCreateGamedayRunResponse**](types.CreateGamedayRunResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateGamedayV2**
> TypesCreateGamedayResponse CreateGamedayV2(ctx, body, accountIdentifier, organizationIdentifier, projectIdentifier)
Create a chaos v2 gameday

Create a chaos v2 gameday

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TypesCreateGamedayRequest**](TypesCreateGamedayRequest.md)| Create Gameday | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 

### Return type

[**TypesCreateGamedayResponse**](types.CreateGamedayResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateInputSet**
> InputsetsCreateInputSetResponse CreateInputSet(ctx, body, accountIdentifier, organizationIdentifier, projectIdentifier, experimentId, isIdentity)
Create an input set

Create an input set

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**InputsetsCreateInputSetRequest**](InputsetsCreateInputSetRequest.md)| create input set request | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **experimentId** | **string**| experimentId where input set should be created | 
  **isIdentity** | **bool**| is human-readable experiment identity passed | 

### Return type

[**InputsetsCreateInputSetResponse**](inputsets.CreateInputSetResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateProbe**
> TypesCreateProbeResponse CreateProbe(ctx, body, accountIdentifier, organizationIdentifier, projectIdentifier)
Create a new probe

Create a new probe

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TypesProbeRequest**](TypesProbeRequest.md)| Probe configuration | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 

### Return type

[**TypesCreateProbeResponse**](types.CreateProbeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateProbeTemplate**
> ChaosprobetemplateProbeTemplate CreateProbeTemplate(ctx, body, accountIdentifier, organizationIdentifier, projectIdentifier)
Creates the probe templates in a hub based on tag

Creates probe templates in a hub based on tag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ChaosprobetemplateProbeTemplate**](ChaosprobetemplateProbeTemplate.md)| probe template configuration | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 

### Return type

[**ChaosprobetemplateProbeTemplate**](chaosprobetemplate.ProbeTemplate.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRecommendation**
> RecommendationsCreateActionResponse CreateRecommendation(ctx, accountIdentifier, organizationIdentifier, projectIdentifier, recommendationID)
Create new experiment

Create new experiment based on the recommendation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **recommendationID** | **string**| recommendation id | 

### Return type

[**RecommendationsCreateActionResponse**](recommendations.CreateActionResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateV2Onboarding**
> V2OnboardingV2Onboarding CreateV2Onboarding(ctx, body, accountIdentifier, organizationIdentifier, projectIdentifier, optional)
Create V2 Onboarding

Create V2 Onboarding

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V2OnboardingV2OnboardingRequest**](V2OnboardingV2OnboardingRequest.md)| Create Onboarding | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
 **optional** | ***DefaultApiCreateV2OnboardingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiCreateV2OnboardingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **corelationID** | **optional.**| corelation id is used to debug micro svc communication | 

### Return type

[**V2OnboardingV2Onboarding**](v2_onboarding.V2Onboarding.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAction**
> string DeleteAction(ctx, accountIdentifier, organizationIdentifier, projectIdentifier, identity)
Delete an action

Delete an action with a provided identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account ID to access the resource | 
  **organizationIdentifier** | **string**| Organization ID to access the resource | 
  **projectIdentifier** | **string**| Project ID to access the resource | 
  **identity** | **string**| ID of the Action to retrieve | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteActionTemplate**
> bool DeleteActionTemplate(ctx, accountIdentifier, organizationIdentifier, projectIdentifier, hubIdentity, identity, optional)
Deletes action template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **hubIdentity** | **string**| reference of the hub i.e. hub ID | 
  **identity** | **string**| name of the fault | 
 **optional** | ***DefaultApiDeleteActionTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiDeleteActionTemplateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **revision** | **optional.String**| revision of the action template to be deleted | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteChaosV2Experiment**
> TypesDeleteChaosV2ExperimentResponse DeleteChaosV2Experiment(ctx, accountIdentifier, organizationIdentifier, projectIdentifier, experimentId)
Delete Chaos V2 experiment

Delete Chaos V2 experiment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **experimentId** | **string**| experiment id that will be used to delete the experiment | 

### Return type

[**TypesDeleteChaosV2ExperimentResponse**](types.DeleteChaosV2ExperimentResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFaultTemplate**
> GithubComHarnessHceSaasGraphqlServerApiEmpty DeleteFaultTemplate(ctx, accountIdentifier, organizationIdentifier, projectIdentifier, hubIdentity, faultName)
Delete the fault templates in a hub based on tag

Delete the fault templates in a hub based on tag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **hubIdentity** | **string**| chaos hub identity | 
  **faultName** | **string**| name of the fault | 

### Return type

[**GithubComHarnessHceSaasGraphqlServerApiEmpty**](github_com_harness_hce-saas_graphql_server_api.Empty.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteGamedayV2**
> string DeleteGamedayV2(ctx, accountIdentifier, organizationIdentifier, projectIdentifier, gamedayId)
Get a chaos v2 Gameday

Get a chaos v2 Gameday

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **gamedayId** | **string**| gamedayId of gameday to be deleted | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteInputSet**
> InputsetsDeleteInputSetResponse DeleteInputSet(ctx, accountIdentifier, organizationIdentifier, projectIdentifier, experimentId, inputsetId, isIdentity)
Delete an input set

Delete an input set

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **experimentId** | **string**| experimentId where input set should be created | 
  **inputsetId** | **string**| ID of the input set | 
  **isIdentity** | **bool**| is human-readable experiment identity passed | 

### Return type

[**InputsetsDeleteInputSetResponse**](inputsets.DeleteInputSetResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProbe**
> string DeleteProbe(ctx, accountIdentifier, organizationIdentifier, projectIdentifier, probeId)
Delete a probe

Delete a probe with a provided identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account ID to access the resource | 
  **organizationIdentifier** | **string**| Organization ID to access the resource | 
  **projectIdentifier** | **string**| Project ID to access the resource | 
  **probeId** | **string**| ID of the probe to retrieve | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProbeTemplate**
> bool DeleteProbeTemplate(ctx, accountIdentifier, organizationIdentifier, projectIdentifier, hubIdentity, identity, optional)
Deletes probe template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **hubIdentity** | **string**| reference of the hub i.e. hub ID | 
  **identity** | **string**| name of the fault | 
 **optional** | ***DefaultApiDeleteProbeTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiDeleteProbeTemplateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **revision** | **optional.String**| revision of the probe template to be deleted | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRecommendation**
> DeleteRecommendation(ctx, accountIdentifier, organizationIdentifier, projectIdentifier, recommendationID)
Delete recommendation from db

Delete recommendation based on the recommendation id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **recommendationID** | **string**| recommendation id | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTargetNetworkMap**
> bool DeleteTargetNetworkMap(ctx, accountIdentifier, organizationIdentifier, projectIdentifier, environmentIdentifier, infraId, applicationmapid)
Delete application network map with chaos context

Delete application network map with chaos context

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **environmentIdentifier** | **string**| environment id that want to access the resource | 
  **infraId** | **string**| infra id that want to access the resource | 
  **applicationmapid** | **string**| Application map ID | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableProbe**
> string EnableProbe(ctx, body, accountIdentifier, organizationIdentifier, projectIdentifier, probeId)
Enable and disable probes across experiments

Enable and disable probes across experiments

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TypesProbeBulkEnableRequest**](TypesProbeBulkEnableRequest.md)| Enable Probe configuration | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **probeId** | **string**| ID of the probe to retrieve | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExperimentExecutionNodeDetails**
> []ExecutionChaosExecutionNode ExperimentExecutionNodeDetails(ctx, accountIdentifier, organizationIdentifier, projectIdentifier, name, experimentId, experimentRunId)
Get chaos execution node

Get chaos execution node

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| org id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **name** | **string**| name of the node | 
  **experimentId** | **string**| experimentId to be fetched | 
  **experimentRunId** | **string**| experimentRunId to be fetched | 

### Return type

[**[]ExecutionChaosExecutionNode**](execution.ChaosExecutionNode.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountServiceDetails**
> ChaosserviceusageServiceDataResponse GetAccountServiceDetails(ctx, accountID, page, limit, startTime, endTime, optional)
Get service usage details for account in the requested time frame

Get service details for account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountID** | **string**| ID of the account | 
  **page** | **int32**| page number | [default to 0]
  **limit** | **int32**| limit per page | [default to 10]
  **startTime** | **int32**| start time in unix format in milliseconds | 
  **endTime** | **int32**| end time in unix format in milliseconds | 
 **optional** | ***DefaultApiGetAccountServiceDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiGetAccountServiceDetailsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **service** | **optional.String**| search based on service name | 
 **serviceType** | **optional.String**| search based on service type | 
 **sortAscending** | **optional.Bool**| sort the response in ascending order | [default to false]
 **sortField** | **optional.String**| sort the response w.r.t field: faultsRan, experiments, experimentsRan | [default to faultsRan]

### Return type

[**ChaosserviceusageServiceDataResponse**](chaosserviceusage.ServiceDataResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountServiceUsageStats**
> ChaosserviceusageUsageStats GetAccountServiceUsageStats(ctx, accountID, groupBy, startTime, endTime)
Get the service usage stats related to requested account grouped by day or month

Get service usage stats for account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountID** | **string**| ID of the account | 
  **groupBy** | **string**| group by period (day/month) | 
  **startTime** | **int32**| start time in unix format in milliseconds | 
  **endTime** | **int32**| end time in unix format in milliseconds | 

### Return type

[**ChaosserviceusageUsageStats**](chaosserviceusage.UsageStats.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAction**
> ActionsActionResponse GetAction(ctx, accountIdentifier, organizationIdentifier, projectIdentifier, identity)
Get an action

Retrieve details of a specific action by its ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account ID for accessing the resource | 
  **organizationIdentifier** | **string**| Organization ID for accessing the resource | 
  **projectIdentifier** | **string**| Project ID for accessing the resource | 
  **identity** | **string**| ID of the Action to retrieve | 

### Return type

[**ActionsActionResponse**](actions.ActionResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetActionManifest**
> string GetActionManifest(ctx, accountIdentifier, organizationIdentifier, projectIdentifier, identity)
Get an action manifest

Retrieve Action manifest of a specific action by its ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account ID for accessing the resource | 
  **organizationIdentifier** | **string**| Organization ID for accessing the resource | 
  **projectIdentifier** | **string**| Project ID for accessing the resource | 
  **identity** | **string**| ID of the Action to retrieve | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetActionTemplate**
> ChaosfaulttemplateGetActionTemplateResponse GetActionTemplate(ctx, accountIdentifier, organizationIdentifier, projectIdentifier, hubIdentity, identity, optional)
Get the action template in a hub based on action ref

Get the action template in a hub based on action ref

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **hubIdentity** | **string**| chaos hub identity | 
  **identity** | **string**| name of the fault | 
 **optional** | ***DefaultApiGetActionTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiGetActionTemplateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **revision** | **optional.String**| revision of the 1st fault template | 

### Return type

[**ChaosfaulttemplateGetActionTemplateResponse**](chaosfaulttemplate.GetActionTemplateResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetActionTemplateRevisionDifference**
> ChaosfaulttemplateListActionTemplateResponse GetActionTemplateRevisionDifference(ctx, accountIdentifier, organizationIdentifier, projectIdentifier, hubIdentity, identity, revision, revisionToCompare)
Get the difference between 2 revisions of action template

Get the difference between 2 revisions of action template in a hub

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **hubIdentity** | **string**| reference of the hub i.e. hub ID | 
  **identity** | **string**| name of the fault | 
  **revision** | **string**| revision of the 1st fault template | 
  **revisionToCompare** | **string**| revision to compare | 

### Return type

[**ChaosfaulttemplateListActionTemplateResponse**](chaosfaulttemplate.ListActionTemplateResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetChaosDashboard**
> []HandlersDashboard GetChaosDashboard(ctx, )
Get chaos dashboard

Get chaos dashboard

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]HandlersDashboard**](handlers.Dashboard.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetChaosExperimentRunReport**
> *os.File GetChaosExperimentRunReport(ctx, experimentRunId, notifyId, accountIdentifier, organizationIdentifier, projectIdentifier)
Generate and return kubernetesV1 chaos experiment run report

Generate and return kubernetesV1 chaos experiment run report

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **experimentRunId** | **string**| experimentRunId for kubernetesV1 experiment run report generation | 
  **notifyId** | **string**| notifyId for kubernetesV1 experiment run report generation | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/pdf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetChaosHub**
> Chaoshubv2GetHubResponse GetChaosHub(ctx, accountIdentifier, organizationIdentifier, projectIdentifier, hubIdentity)
Get chaos hub based on given filters

Get chaos hub based on given filters

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **hubIdentity** | **string**| chaos hub identity | 

### Return type

[**Chaoshubv2GetHubResponse**](chaoshubv2.GetHubResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetChaosPipelineExecution**
> ChaosExecutionNodesChaosExecutionResponse GetChaosPipelineExecution(ctx, accountIdentifier, organizationIdentifier, projectIdentifier, experimentId, experimentRunId)
Get a chaos pipeline execution

Get a chaos pipeline execution

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **experimentId** | **string**| experimentId to be fetched | 
  **experimentRunId** | **string**| experimentRunId to be fetched | 

### Return type

[**ChaosExecutionNodesChaosExecutionResponse**](chaos_execution_nodes.ChaosExecutionResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetChaosPipelineStepDetails**
> ChaosexperimentpipelineGetChaosPipelineNodesResponse GetChaosPipelineStepDetails(ctx, accountIdentifier, organizationIdentifier, projectIdentifier, experimentId, experimentRunId, stepName)
Get a chaos pipeline step execution

Get a chaos pipeline step execution

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **experimentId** | **string**| experimentId to be fetched | 
  **experimentRunId** | **string**| experimentRunId to be fetched | 
  **stepName** | **string**| stepName to be fetched | 

### Return type

[**ChaosexperimentpipelineGetChaosPipelineNodesResponse**](chaosexperimentpipeline.GetChaosPipelineNodesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetChaosV2Experiment**
> ChaosExperimentChaosExperimentRequest GetChaosV2Experiment(ctx, accountIdentifier, organizationIdentifier, projectIdentifier, experimentId)
Get a chaos v2 experiment

Get a chaos v2 experiment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **experimentId** | **string**| experimentId to be fetched | 

### Return type

[**ChaosExperimentChaosExperimentRequest**](chaos_experiment.ChaosExperimentRequest.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetChaosV2ExperimentRun**
> ChaosExperimentRunChaosExperimentRun GetChaosV2ExperimentRun(ctx, accountIdentifier, organizationIdentifier, projectIdentifier, experimentId, optional)
Get a chaos v2 experiment run

Get a chaos v2 experiment run

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **experimentId** | **string**| experimentId to be executed | 
 **optional** | ***DefaultApiGetChaosV2ExperimentRunOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiGetChaosV2ExperimentRunOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **experimentRunId** | **optional.String**| experimentRunId to be fetched | 
 **notifyId** | **optional.String**| notifyId to be fetched | 

### Return type

[**ChaosExperimentRunChaosExperimentRun**](chaos_experiment_run.ChaosExperimentRun.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetChaosV2ExperimentRunInternalAPI**
> ModelWorkflowRun GetChaosV2ExperimentRunInternalAPI(ctx, body, accountIdentifier, organizationIdentifier, projectIdentifier, experimentId, notifyId)
Get the chaos v2 experiment run internal API

Get the execution details of a chaos v2 experiment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TypesInternalExperimentRunRequest**](TypesInternalExperimentRunRequest.md)| get Experiment | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **experimentId** | **string**| experimentId to be filtered | 
  **notifyId** | **string**| notifyId for which execution is to be fetched | 

### Return type

[**ModelWorkflowRun**](model.WorkflowRun.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetChaosV2ExperimentVariables**
> TemplateRunTimeVariables GetChaosV2ExperimentVariables(ctx, accountIdentifier, organizationIdentifier, projectIdentifier, experimentId, isIdentity)
Get a chaos v2 experiment

Get a chaos v2 experiment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **experimentId** | **string**| experimentId to be fetched | 
  **isIdentity** | **bool**| is human-readable experiment identity passed | 

### Return type

[**TemplateRunTimeVariables**](template.RunTimeVariables.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConnectorForInfra**
> ApiGetHarnessInfrastructureResponse GetConnectorForInfra(ctx, identity, accountIdentifier, organizationIdentifier, projectIdentifier, environmentIdentifier)
Get Connector For Infra

Get Connector For Infra

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identity** | **string**| Chaos V2 Infra Identity | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **environmentIdentifier** | **string**| environment id that want to access the resource | 

### Return type

[**ApiGetHarnessInfrastructureResponse**](api.GetHarnessInfrastructureResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExperimentHelperImageVersion**
> HandlersExperimentHelperImageVersion GetExperimentHelperImageVersion(ctx, )
Get experiment helper image version

Get experiment helper image version

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**HandlersExperimentHelperImageVersion**](handlers.ExperimentHelperImageVersion.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExperimentRunTimelineView**
> ChaosExecutionNodesChaosExecutionResponse GetExperimentRunTimelineView(ctx, accountIdentifier, organizationIdentifier, projectIdentifier, experimentId, optional)
Get a chaos v2 experiment timeline run

Get a chaos v2 experiment timeline run

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **experimentId** | **string**| experimentId to be executed | 
 **optional** | ***DefaultApiGetExperimentRunTimelineViewOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiGetExperimentRunTimelineViewOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **experimentRunId** | **optional.String**| experimentRunId to be fetched | 
 **notifyId** | **optional.String**| notifyId to be fetched | 

### Return type

[**ChaosExecutionNodesChaosExecutionResponse**](chaos_execution_nodes.ChaosExecutionResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExperimentRunsOverviewStats**
> HandlersChaosExperimentRunsStatsResponse GetExperimentRunsOverviewStats(ctx, optional)
Get chaos experiment run overview stats

Get resource usage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiGetExperimentRunsOverviewStatsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiGetExperimentRunsOverviewStatsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountIdentifier** | **optional.String**| Account Identifier | 
 **orgIdentifier** | **optional.String**| Organization Identifier | 
 **projectIdentifier** | **optional.String**| Project Identifier | 
 **startTime** | **optional.String**| Start Time | 
 **endTime** | **optional.String**| End Time | 
 **groupBy** | **optional.String**| Group By Parameter | 

### Return type

[**HandlersChaosExperimentRunsStatsResponse**](handlers.ChaosExperimentRunsStatsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFaultTemplate**
> ChaosfaulttemplateGetFaultTemplateResponse GetFaultTemplate(ctx, accountIdentifier, organizationIdentifier, projectIdentifier, hubIdentity, revision, faultName)
Lists all the fault templates in a hub based on tag

Lists all the fault templates in a hub based on tag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **hubIdentity** | **string**| chaos hub identity | 
  **revision** | **string**| revision of the 1st fault template | 
  **faultName** | **string**| name of the fault | 

### Return type

[**ChaosfaulttemplateGetFaultTemplateResponse**](chaosfaulttemplate.GetFaultTemplateResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFaultTemplateRevisionDifference**
> ChaosfaulttemplateListFaultTemplateResponse GetFaultTemplateRevisionDifference(ctx, accountIdentifier, organizationIdentifier, projectIdentifier, hubIdentity, faultName, revision, revisionToCompare)
Get the difference between 2 revisions of a fault template

Get the difference between 2 revisions of a fault template in a hub

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **hubIdentity** | **string**| reference of the hub i.e. hub ID | 
  **faultName** | **string**| name of the fault | 
  **revision** | **string**| revision of the 1st fault template | 
  **revisionToCompare** | **string**| revision to compare | 

### Return type

[**ChaosfaulttemplateListFaultTemplateResponse**](chaosfaulttemplate.ListFaultTemplateResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGamedayRunV2**
> TypesRun GetGamedayRunV2(ctx, accountIdentifier, organizationIdentifier, projectIdentifier, gamedayId, gamedayRunId)
Fetch a chaos v2 gameday run

Fetch a chaos v2 gameday run

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **gamedayId** | **string**| gamedayId of the run | 
  **gamedayRunId** | **string**| gamedayRunId to be run | 

### Return type

[**TypesRun**](types.Run.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGamedayV2**
> TypesGetGamedayResponse GetGamedayV2(ctx, accountIdentifier, organizationIdentifier, projectIdentifier, gamedayId)
Get a chaos v2 Gameday

Get a chaos v2 Gameday

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **gamedayId** | **string**| gamedayId to be fetched | 

### Return type

[**TypesGetGamedayResponse**](types.GetGamedayResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetImageRegistry**
> HandlersImageRegistryDetails GetImageRegistry(ctx, )
Get image registry

Get image registry

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**HandlersImageRegistryDetails**](handlers.ImageRegistryDetails.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInfraToken**
> K8sinfraGetInfraTokenResponse GetInfraToken(ctx, infrastructureIdentity, accountIdentifier, organizationIdentifier, projectIdentifier, environmentIdentifier)
Get a v2 infra token

Get a v2 infra token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **infrastructureIdentity** | **string**| Chaos V2 Infra Identity | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **environmentIdentifier** | **string**| environment identifier to filter resource | 

### Return type

[**K8sinfraGetInfraTokenResponse**](k8sinfra.GetInfraTokenResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInputSet**
> InputsetsGetInputSetResponse GetInputSet(ctx, accountIdentifier, organizationIdentifier, projectIdentifier, experimentId, inputsetId, isIdentity)
Get the input set in an experiment

Get the input set in an experiment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **experimentId** | **string**| experimentId whose input set should be retrieved | 
  **inputsetId** | **string**| ID of the input set | 
  **isIdentity** | **bool**| is human-readable experiment identity passed | 

### Return type

[**InputsetsGetInputSetResponse**](inputsets.GetInputSetResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNote**
> TypesNotes GetNote(ctx, accountIdentifier, organizationIdentifier, projectIdentifier, optional)
Get a chaos resource Note

Get a chaos resource Note

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
 **optional** | ***DefaultApiGetNoteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiGetNoteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **gamedayRunID** | **optional.String**| gamedayRunID as filter for summary notes | 
 **experimentRunId** | **optional.String**| experimentRunId as filter for experiment run notes | 
 **experimentId** | **optional.String**| experimentId as filter for gameday experiment notes | 
 **noteType** | **optional.String**| type of note | 

### Return type

[**TypesNotes**](types.Notes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOnboardingExperiments**
> V2OnboardingOnboardingExperimentResponse GetOnboardingExperiments(ctx, onboardingid, page, limit, optional)
Get V2 Onboarding experiments

Get V2 Onboarding experiments

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **onboardingid** | **string**| onboarding id | 
  **page** | **int32**| page number | [default to 0]
  **limit** | **int32**| limit per page | [default to 5]
 **optional** | ***DefaultApiGetOnboardingExperimentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiGetOnboardingExperimentsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **corelationID** | **optional.String**| corelation id is used to debug micro svc communication | 
 **accountIdentifier** | **optional.String**| account id that want to access the resource | 
 **organizationIdentifier** | **optional.String**| organization id that want to access the resource | 
 **projectIdentifier** | **optional.String**| project id that want to access the resource | 

### Return type

[**V2OnboardingOnboardingExperimentResponse**](v2_onboarding.OnboardingExperimentResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOverallServiceUsageStats**
> ChaosserviceusageOverallServiceUsageStats GetOverallServiceUsageStats(ctx, accountID, startTime, endTime)
Get the overall service usage stats by type related to requested account

Get overall service usage stats by type for account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountID** | **string**| ID of the account | 
  **startTime** | **int32**| start time in unix format in milliseconds | 
  **endTime** | **int32**| end time in unix format in milliseconds | 

### Return type

[**ChaosserviceusageOverallServiceUsageStats**](chaosserviceusage.OverallServiceUsageStats.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProbe**
> TypesGetProbeResponse GetProbe(ctx, accountIdentifier, organizationIdentifier, projectIdentifier, probeId)
Get a probe

Retrieve details of a specific probe by its ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account ID for accessing the resource | 
  **organizationIdentifier** | **string**| Organization ID for accessing the resource | 
  **projectIdentifier** | **string**| Project ID for accessing the resource | 
  **probeId** | **string**| ID of the probe to retrieve | 

### Return type

[**TypesGetProbeResponse**](types.GetProbeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProbeManifest**
> string GetProbeManifest(ctx, accountIdentifier, organizationIdentifier, projectIdentifier, probeId)
Get a probe

Retrieve probe manifest of a specific probe by its ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account ID for accessing the resource | 
  **organizationIdentifier** | **string**| Organization ID for accessing the resource | 
  **projectIdentifier** | **string**| Project ID for accessing the resource | 
  **probeId** | **string**| ID of the probe to retrieve | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProbeTemplate**
> ChaosprobetemplateGetProbeTemplateResponse GetProbeTemplate(ctx, accountIdentifier, organizationIdentifier, projectIdentifier, hubIdentity, identity, optional)
Get the probe template in a hub based on probe ref

Get the probe template in a hub based on probe ref

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **hubIdentity** | **string**| chaos hub identity | 
  **identity** | **string**| name of the fault | 
 **optional** | ***DefaultApiGetProbeTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiGetProbeTemplateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **revision** | **optional.String**| revision of the 1st fault template | 

### Return type

[**ChaosprobetemplateGetProbeTemplateResponse**](chaosprobetemplate.GetProbeTemplateResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRecommendation**
> GithubComHarnessHceSaasGraphqlServerPkgDatabaseMongodbRecommendationRecommendation GetRecommendation(ctx, accountIdentifier, organizationIdentifier, projectIdentifier, recommendationID)
Get recommendation details

Get recommendation details based on the recommendation id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **recommendationID** | **string**| recommendation id | 

### Return type

[**GithubComHarnessHceSaasGraphqlServerPkgDatabaseMongodbRecommendationRecommendation**](github_com_harness_hce-saas_graphql_server_pkg_database_mongodb_recommendation.Recommendation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetResourceUsage**
> HandlersChaosLicenseUsageDto GetResourceUsage(ctx, )
Get resource usage

Get resource usage

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**HandlersChaosLicenseUsageDto**](handlers.CHAOSLicenseUsageDTO.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSGPTemplateByGenAI**
> SecurityGovernanceCondition GetSGPTemplateByGenAI(ctx, )
Get security governance conditions template

Get security governance conditions template

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SecurityGovernanceCondition**](security_governance.Condition.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceResponse**
> NetworkmapGetTargetServiceResponse GetServiceResponse(ctx, accountIdentifier, organizationIdentifier, projectIdentifier, environmentIdentifier, infraId, applicationmapid, targetserviceid)
Get target discovered service with chaos context

Get target discovered service with chaos context

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **environmentIdentifier** | **string**| environment id that want to access the resource | 
  **infraId** | **string**| infra id that want to access the resource | 
  **applicationmapid** | **string**| Application map ID | 
  **targetserviceid** | **string**| Target discovered service ID Identity | 

### Return type

[**NetworkmapGetTargetServiceResponse**](networkmap.GetTargetServiceResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceUsageReport**
> [][]string GetServiceUsageReport(ctx, accountID, startTime, endTime)
Generates service usage report in csv format

Generates service usage report for account in a given timeframe in csv format

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountID** | **string**| ID of the account | 
  **startTime** | **int32**| start time in unix format in milliseconds | 
  **endTime** | **int32**| end time in unix format in milliseconds | 

### Return type

[**[][]string**](array.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTargetNetworkMap**
> NetworkmapGetTargetNetworkMapResponse GetTargetNetworkMap(ctx, accountIdentifier, organizationIdentifier, projectIdentifier, environmentIdentifier, infraId, applicationmapid)
Get target network map with chaos context

Get target network map with chaos context

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **environmentIdentifier** | **string**| environment id that want to access the resource | 
  **infraId** | **string**| infra id that want to access the resource | 
  **applicationmapid** | **string**| Application map ID | 

### Return type

[**NetworkmapGetTargetNetworkMapResponse**](networkmap.GetTargetNetworkMapResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV2InfrastructureYaml**
> InfraV2GetKubernetesInfrastructureV2YamlResponse GetV2InfrastructureYaml(ctx, body, accountIdentifier, organizationIdentifier, projectIdentifier)
Preview v2 infra Yaml

Preview v2 infra Yaml

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**InfraV2GetKubernetesInfrastructureV2YamlRequest**](InfraV2GetKubernetesInfrastructureV2YamlRequest.md)| preview Infra V2 yaml | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 

### Return type

[**InfraV2GetKubernetesInfrastructureV2YamlResponse**](infra_v2.GetKubernetesInfrastructureV2YamlResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV2Onboarding**
> V2OnboardingV2Onboarding GetV2Onboarding(ctx, onboardingid, accountIdentifier, organizationIdentifier, projectIdentifier, optional)
Get V2 Onboarding

Get V2 Onboarding

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **onboardingid** | **string**| onboarding id | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
 **optional** | ***DefaultApiGetV2OnboardingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiGetV2OnboardingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **corelationID** | **optional.String**| corelation id is used to debug micro svc communication | 

### Return type

[**V2OnboardingV2Onboarding**](v2_onboarding.V2Onboarding.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImportAction**
> GithubComHarnessHceSaasGraphqlServerPkgDatabaseMongodbActionsAction ImportAction(ctx, body, accountIdentifier, organizationIdentifier, projectIdentifier)
Import a new action

Import a new action

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ActionsImportActionTemplateRequest**](ActionsImportActionTemplateRequest.md)| action configuration | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 

### Return type

[**GithubComHarnessHceSaasGraphqlServerPkgDatabaseMongodbActionsAction**](github_com_harness_hce-saas_graphql_server_pkg_database_mongodb_actions.Action.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImportProbe**
> TypesCreateProbeResponse ImportProbe(ctx, body, accountIdentifier, organizationIdentifier, projectIdentifier)
Import a new probe

Import a new probe

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TypesImportProbeTemplateRequest**](TypesImportProbeTemplateRequest.md)| action configuration | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 

### Return type

[**TypesCreateProbeResponse**](types.CreateProbeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListActionTemplate**
> ChaosfaulttemplateListActionTemplateResponse ListActionTemplate(ctx, accountIdentifier, organizationIdentifier, projectIdentifier, hubIdentity, page, limit, search, optional)
Lists all the action templates in a hub based on tag

Lists all the action templates in a hub based on tag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **hubIdentity** | **string**| chaos hub identity | 
  **page** | **int32**| page number | [default to 0]
  **limit** | **int32**| limit per page | [default to 15]
  **search** | **string**| name of the action | 
 **optional** | ***DefaultApiListActionTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiListActionTemplateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------







 **infraType** | **optional.String**| infra type of the action | 
 **entityType** | **optional.String**| filter based on Action | 
 **includeAllScope** | **optional.String**| include all scope | 

### Return type

[**ChaosfaulttemplateListActionTemplateResponse**](chaosfaulttemplate.ListActionTemplateResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListActionTemplateRevisions**
> ChaosfaulttemplateListActionTemplateResponse ListActionTemplateRevisions(ctx, accountIdentifier, organizationIdentifier, projectIdentifier, hubIdentity, identity, page, limit)
Lists all the revision of a fault template in a hub

Lists all the revision of a fault template in a hub

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **hubIdentity** | **string**| chaos hub identity | 
  **identity** | **string**| name of the fault | 
  **page** | **int32**| page number | [default to 0]
  **limit** | **int32**| limit per page | [default to 15]

### Return type

[**ChaosfaulttemplateListActionTemplateResponse**](chaosfaulttemplate.ListActionTemplateResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListActions**
> ActionsListActionTemplateResponse ListActions(ctx, accountIdentifier, organizationIdentifier, projectIdentifier, search, page, limit, optional)
List actions with filtering options

Retrieve a list of actions based on various filters like name, tags, date range, and infrastructure type, with pagination support.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **search** | **string**| name of the action | 
  **page** | **int32**| page number | [default to 0]
  **limit** | **int32**| limit per page | [default to 15]
 **optional** | ***DefaultApiListActionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiListActionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **hubIdentity** | **optional.String**| chaos hub identity | 
 **infraType** | **optional.String**| infra type of the action | 
 **actionType** | **optional.String**| filter based on Action | 
 **includeAllScope** | **optional.String**| include all scope | 

### Return type

[**ActionsListActionTemplateResponse**](actions.ListActionTemplateResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListApplication**
> ApplicationchaostargetListApplicationResponse ListApplication(ctx, infrastructureIdentity, accountIdentifier, organizationIdentifier, projectIdentifier, environmentIdentifier)
List all applications for a given infra

List all applications for a given infra

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **infrastructureIdentity** | **string**| Chaos V2 Infrastructure identity | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **environmentIdentifier** | **string**| environment identifier to filter resource | 

### Return type

[**ApplicationchaostargetListApplicationResponse**](applicationchaostarget.ListApplicationResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListChaosEnabledInfraV2**
> InfraV2ListKubernetesInfraV2Response ListChaosEnabledInfraV2(ctx, body, accountIdentifier, organizationIdentifier, projectIdentifier, page, limit, optional)
List a new v2 infra

List a new v2 infra

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**InfraV2ListKubernetesInfraV2Request**](InfraV2ListKubernetesInfraV2Request.md)| list Infra V2 | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **page** | **int32**| page number | [default to 0]
  **limit** | **int32**| limit per page | [default to 10]
 **optional** | ***DefaultApiListChaosEnabledInfraV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiListChaosEnabledInfraV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **environmentIdentifier** | **optional.**| filter infra | 
 **search** | **optional.**| search based on name | 

### Return type

[**InfraV2ListKubernetesInfraV2Response**](infra_v2.ListKubernetesInfraV2Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListChaosHub**
> Chaoshubv2ListHubResponse ListChaosHub(ctx, accountIdentifier, organizationIdentifier, projectIdentifier, search, includeAllScope, page, limit)
Lists chaos hubs based on given filters

Lists chaos hubs based on given filters

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **search** | **string**| search string for chaos hub name | 
  **includeAllScope** | **bool**| get hubs from all scope | 
  **page** | **int32**| page number | [default to 0]
  **limit** | **int32**| limit per page | [default to 15]

### Return type

[**Chaoshubv2ListHubResponse**](chaoshubv2.ListHubResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListChaosV2Experiment**
> TypesListExperimentV2Response ListChaosV2Experiment(ctx, accountIdentifier, organizationIdentifier, projectIdentifier, page, limit, optional)
Get list of Chaos V2 Experiments

Get list of Chaos V2 Experiments

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **page** | **int32**| page number | [default to 0]
  **limit** | **int32**| limit per page | [default to 10]
 **optional** | ***DefaultApiListChaosV2ExperimentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiListChaosV2ExperimentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **startDate** | **optional.String**| filter based on start time | 
 **endDate** | **optional.String**| filter based on end time | 
 **experimentName** | **optional.String**| search based on experiment name | 
 **infraName** | **optional.String**| search based on infra name | 
 **infraId** | **optional.String**| filter based on infraId | 
 **infraActive** | **optional.Bool**| filter based on infra state | 
 **tags** | **optional.String**| filter based on tags | 
 **experimentIds** | **optional.String**| filter based on experimentID | 
 **environmentIdentifier** | **optional.String**| filter based on environmentID | 
 **targetNetworkMapIds** | **optional.String**| filter experiments based on experiment ids | 
 **status** | **optional.String**| filter based on status | 

### Return type

[**TypesListExperimentV2Response**](types.ListExperimentV2Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFault**
> ChaoshubListFaultsResponse ListFault(ctx, accountIdentifier, organizationIdentifier, projectIdentifier, search, page, limit, optional)
Lists faults in a chaos hub based on given filters

Lists faults in a chaos hub based on given filters

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **search** | **string**| search string for chaos hub name | 
  **page** | **int32**| page number | [default to 0]
  **limit** | **int32**| limit per page | [default to 15]
 **optional** | ***DefaultApiListFaultOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiListFaultOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **hubIdentity** | **optional.String**| identity of hub | 
 **infraType** | **optional.String**| type of the infra | 
 **permissionsRequired** | **optional.String**| permissions required for fault | 
 **entityType** | **optional.String**| category of the fault | 

### Return type

[**ChaoshubListFaultsResponse**](chaoshub.ListFaultsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFaultTemplate**
> ChaosfaulttemplateListFaultTemplateResponse ListFaultTemplate(ctx, accountIdentifier, organizationIdentifier, projectIdentifier, hubIdentity, page, limit)
Lists all the fault templates in a hub based on tag

Lists all the fault templates in a hub based on tag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **hubIdentity** | **string**| chaos hub identity | 
  **page** | **int32**| page number | [default to 0]
  **limit** | **int32**| limit per page | [default to 15]

### Return type

[**ChaosfaulttemplateListFaultTemplateResponse**](chaosfaulttemplate.ListFaultTemplateResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFaultTemplateRevisions**
> ChaosfaulttemplateListFaultTemplateResponse ListFaultTemplateRevisions(ctx, accountIdentifier, organizationIdentifier, projectIdentifier, hubIdentity, faultName, page, limit)
Lists all the revision of a fault template in a hub

Lists all the revision of a fault template in a hub

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **hubIdentity** | **string**| chaos hub identity | 
  **faultName** | **string**| name of the fault | 
  **page** | **int32**| page number | [default to 0]
  **limit** | **int32**| limit per page | [default to 15]

### Return type

[**ChaosfaulttemplateListFaultTemplateResponse**](chaosfaulttemplate.ListFaultTemplateResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFunction**
> ApplicationchaostargetListFunctionResponse ListFunction(ctx, infrastructureIdentity, applicationIdentity, accountIdentifier, organizationIdentifier, projectIdentifier, environmentIdentifier)
List instrumented functions for given application

List instrumented functions for given application

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **infrastructureIdentity** | **string**| Chaos V2 Infrastructure identity | 
  **applicationIdentity** | **string**| application identity | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **environmentIdentifier** | **string**| environment identifier to filter resource | 

### Return type

[**ApplicationchaostargetListFunctionResponse**](applicationchaostarget.ListFunctionResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListGamedayRunV2**
> TypesListGamedayRunV2Response ListGamedayRunV2(ctx, accountIdentifier, organizationIdentifier, projectIdentifier, gamedayId)
Fetch chaos v2 gameday runs

Fetch chaos v2 gameday runs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **gamedayId** | **string**| gamedayId of the run | 

### Return type

[**TypesListGamedayRunV2Response**](types.ListGamedayRunV2Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListGamedayV2**
> TypesListGamedayV2Response ListGamedayV2(ctx, accountIdentifier, organizationIdentifier, projectIdentifier, page, limit, optional)
Get list of Chaos V2 Gamedays

Get list of Chaos V2 Gamedays

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **page** | **int32**| page number | [default to 0]
  **limit** | **int32**| limit per page | [default to 10]
 **optional** | ***DefaultApiListGamedayV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiListGamedayV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **gamedayName** | **optional.String**| search based on gameday name | 
 **sortAscending** | **optional.Bool**| sort the response in ascending order | [default to false]
 **sortField** | **optional.String**| sort the response w.r.t field: CREATED_AT/UPDATED_AT/NAME | [default to UPDATED_AT]

### Return type

[**TypesListGamedayV2Response**](types.ListGamedayV2Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListHarnessInfra**
> ApiListHarnessInfrastructureResponse ListHarnessInfra(ctx, accountIdentifier, organizationIdentifier, projectIdentifier, environmentIdentifier, page, limit, optional)
List harness infras

List harness infras

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **environmentIdentifier** | **string**| environment id that want to access the resource | 
  **page** | **int32**| page number | [default to 0]
  **limit** | **int32**| limit per page | [default to 10]
 **optional** | ***DefaultApiListHarnessInfraOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiListHarnessInfraOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **search** | **optional.String**| search based on name | 

### Return type

[**ApiListHarnessInfrastructureResponse**](api.ListHarnessInfrastructureResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListInputSet**
> InputsetsListInputSetResponse ListInputSet(ctx, accountIdentifier, organizationIdentifier, projectIdentifier, experimentId, page, limit, isIdentity)
Get the list of input sets in an experiment

Get the list of input sets in an experiment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **experimentId** | **string**| experimentId for whose input sets should be listed | 
  **page** | **int32**| page number | [default to 0]
  **limit** | **int32**| limit per page | [default to 15]
  **isIdentity** | **bool**| is human-readable experiment identity passed | 

### Return type

[**InputsetsListInputSetResponse**](inputsets.ListInputSetResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListK8sInfrasV2**
> ModelListInfraResponse ListK8sInfrasV2(ctx, body, accountIdentifier, organizationIdentifier, projectIdentifier, optional)
Get list of active unused connector

Get list of active unused connector

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ModelListInfraRequest**](ModelListInfraRequest.md)| List Infra | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
 **optional** | ***DefaultApiListK8sInfrasV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiListK8sInfrasV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **correlationID** | **optional.**| correlation id is used to debug micro svc communication | 

### Return type

[**ModelListInfraResponse**](model.ListInfraResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListProbeTemplate**
> ChaosprobetemplateListProbeTemplateResponse ListProbeTemplate(ctx, accountIdentifier, organizationIdentifier, projectIdentifier, hubIdentity, page, limit, search, optional)
Lists all the probe templates in a hub based on tag

Lists all the probe templates in a hub based on tag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **hubIdentity** | **string**| chaos hub identity | 
  **page** | **int32**| page number | [default to 0]
  **limit** | **int32**| limit per page | [default to 15]
  **search** | **string**| name of the probe | 
 **optional** | ***DefaultApiListProbeTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiListProbeTemplateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------







 **infraType** | **optional.String**| infra type of the probe | 
 **entityType** | **optional.String**| filter based on probe | 
 **includeAllScope** | **optional.String**| include all scope | 

### Return type

[**ChaosprobetemplateListProbeTemplateResponse**](chaosprobetemplate.ListProbeTemplateResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListProbes**
> TypesListProbeResponse ListProbes(ctx, accountIdentifier, organizationIdentifier, projectIdentifier, optional)
List probes with filtering options

Retrieve a list of probes based on various filters like name, tags, date range, and infrastructure type, with pagination support.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account ID for accessing the resource | 
  **organizationIdentifier** | **string**| Organization ID for accessing the resource | 
  **projectIdentifier** | **string**| Project ID for accessing the resource | 
 **optional** | ***DefaultApiListProbesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiListProbesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **name** | **optional.String**| Filter probes by name | 
 **tags** | **optional.String**| Comma-separated list of tags to filter probes | 
 **startDate** | **optional.String**| Filter probes by start date (ISO 8601 format) | 
 **endDate** | **optional.String**| Filter probes by end date (ISO 8601 format) | 
 **probeIDs** | **optional.String**| Comma-separated list of probe IDs | 
 **infrastructureType** | **optional.String**| Filter probes by infrastructure type | 
 **page** | **optional.Int32**| Page index for pagination | [default to 1]
 **limit** | **optional.Int32**| Number of items per page | [default to 50]
 **sortField** | **optional.String**| Field to sort the probe list | 
 **sortAscending** | **optional.Bool**| Sort the field in ascending order | 
 **probeType** | **optional.String**| Filter based on probe | 

### Return type

[**TypesListProbeResponse**](types.ListProbeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRecommendations**
> RecommendationsListRecommendationsResponse ListRecommendations(ctx, body, accountIdentifier, organizationIdentifier, projectIdentifier)
List recommendations

List recommendations based on the filters

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RecommendationsListRecommendationsRequest**](RecommendationsListRecommendationsRequest.md)| request body | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 

### Return type

[**RecommendationsListRecommendationsResponse**](recommendations.ListRecommendationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListService**
> NetworkmapListTargetServiceResponse ListService(ctx, accountIdentifier, organizationIdentifier, projectIdentifier, environmentIdentifier, infraId, applicationmapid, page, limit, optional)
List target discovered service with chaos context

Get target discovered service with chaos context

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **environmentIdentifier** | **string**| environment id that want to access the resource | 
  **infraId** | **string**| infra id that want to access the resource | 
  **applicationmapid** | **string**| Application map ID | 
  **page** | **int32**| page number | [default to 0]
  **limit** | **int32**| limit per page | [default to 10]
 **optional** | ***DefaultApiListServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiListServiceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------








 **search** | **optional.String**| search based on name | 

### Return type

[**NetworkmapListTargetServiceResponse**](networkmap.ListTargetServiceResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTargetNetworkMaps**
> NetworkmapListTargetNetworkMapResponse ListTargetNetworkMaps(ctx, body, accountIdentifier, organizationIdentifier, projectIdentifier, page, limit, optional)
List target network maps with chaos context

List target network maps with chaos context

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NetworkmapListTargetNetworkMapRequest**](NetworkmapListTargetNetworkMapRequest.md)| list Target Network Maps request | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **page** | **int32**| page number | [default to 0]
  **limit** | **int32**| limit per page | [default to 10]
 **optional** | ***DefaultApiListTargetNetworkMapsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiListTargetNetworkMapsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **search** | **optional.**| search based on name | 

### Return type

[**NetworkmapListTargetNetworkMapResponse**](networkmap.ListTargetNetworkMapResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListV2Onboarding**
> V2OnboardingV2OnboardingList ListV2Onboarding(ctx, page, limit, accountIdentifier, organizationIdentifier, projectIdentifier, search, optional)
Get V2 Onboarding

Get V2 Onboarding

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| page number | [default to 0]
  **limit** | **int32**| limit per page | [default to 10]
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **search** | **string**| search for the onboarding | 
 **optional** | ***DefaultApiListV2OnboardingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiListV2OnboardingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **corelationID** | **optional.String**| corelation id is used to debug micro svc communication | 

### Return type

[**V2OnboardingV2OnboardingList**](v2_onboarding.V2OnboardingList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListVariablesInActionTemplate**
> ChaosfaulttemplateActionTemplateVariables ListVariablesInActionTemplate(ctx, accountIdentifier, organizationIdentifier, projectIdentifier, hubIdentity, identity, revision)
Get the list of variables in a fault template

Get the list of variables in a fault template based on revision

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **hubIdentity** | **string**| reference of the hub i.e. hub ID | 
  **identity** | **string**| name of the fault | 
  **revision** | **string**| revision of the 1st fault template | 

### Return type

[**ChaosfaulttemplateActionTemplateVariables**](chaosfaulttemplate.ActionTemplateVariables.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListVariablesInFaultTemplate**
> ChaosfaulttemplateFaultTemplateVariables ListVariablesInFaultTemplate(ctx, accountIdentifier, organizationIdentifier, projectIdentifier, hubIdentity, faultName, revision)
Get the list of variables in a fault template

Get the list of variables in a fault template based on revision

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **hubIdentity** | **string**| reference of the hub i.e. hub ID | 
  **faultName** | **string**| name of the fault | 
  **revision** | **string**| revision of the 1st fault template | 

### Return type

[**ChaosfaulttemplateFaultTemplateVariables**](chaosfaulttemplate.FaultTemplateVariables.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListVariablesInProbeTemplate**
> ChaosprobetemplateProbeTemplateVariables ListVariablesInProbeTemplate(ctx, accountIdentifier, organizationIdentifier, projectIdentifier, hubIdentity, identity, revision)
Get the list of variables in a fault template

Get the list of variables in a fault template based on revision

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **hubIdentity** | **string**| reference of the hub i.e. hub ID | 
  **identity** | **string**| name of the fault | 
  **revision** | **string**| revision of the 1st fault template | 

### Return type

[**ChaosprobetemplateProbeTemplateVariables**](chaosprobetemplate.ProbeTemplateVariables.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OnboardingConfirmDiscovery**
> V2OnboardingConfirmDiscoveryResponse OnboardingConfirmDiscovery(ctx, onboardingid, accountIdentifier, organizationIdentifier, projectIdentifier, optional)
confirm discovery step in manual onboarding

confirm discovery step in manual onboarding

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **onboardingid** | **string**| onboarding id | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
 **optional** | ***DefaultApiOnboardingConfirmDiscoveryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiOnboardingConfirmDiscoveryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **corelationID** | **optional.String**| corelation id is used to debug micro svc communication | 

### Return type

[**V2OnboardingConfirmDiscoveryResponse**](v2_onboarding.ConfirmDiscoveryResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OnboardingConfirmExperimentCreation**
> V2OnboardingConfirmExperimentCreationResponse OnboardingConfirmExperimentCreation(ctx, body, onboardingid, accountIdentifier, organizationIdentifier, projectIdentifier, optional)
confirm experiment creation step in manual onboarding

confirm experiment creation step in manual onboarding

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V2OnboardingConfirmExperimentCreationRequest**](V2OnboardingConfirmExperimentCreationRequest.md)| Onboarding Confirm ExperimentCreation | 
  **onboardingid** | **string**| onboarding id | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
 **optional** | ***DefaultApiOnboardingConfirmExperimentCreationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiOnboardingConfirmExperimentCreationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **corelationID** | **optional.**| corelation id is used to debug micro svc communication | 

### Return type

[**V2OnboardingConfirmExperimentCreationResponse**](v2_onboarding.ConfirmExperimentCreationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OnboardingConfirmExperimentRun**
> V2OnboardingConfirmExperimentRunResponse OnboardingConfirmExperimentRun(ctx, onboardingid, accountIdentifier, organizationIdentifier, projectIdentifier, optional)
confirm experiment run step in manual onboarding

confirm experiment run step in manual onboarding

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **onboardingid** | **string**| onboarding id | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
 **optional** | ***DefaultApiOnboardingConfirmExperimentRunOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiOnboardingConfirmExperimentRunOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **corelationID** | **optional.String**| corelation id is used to debug micro svc communication | 

### Return type

[**V2OnboardingConfirmExperimentRunResponse**](v2_onboarding.ConfirmExperimentRunResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OnboardingConfirmNetworkMap**
> V2OnboardingConfirmNetworkMapResponse OnboardingConfirmNetworkMap(ctx, body, onboardingid, accountIdentifier, organizationIdentifier, projectIdentifier, optional)
confirm network map creation step in manual onboarding

confirm network map creation step in manual onboarding

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V2OnboardingConfirmNetworkMapRequest**](V2OnboardingConfirmNetworkMapRequest.md)| Onboarding Confirm NetworkMap | 
  **onboardingid** | **string**| onboarding id | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
 **optional** | ***DefaultApiOnboardingConfirmNetworkMapOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiOnboardingConfirmNetworkMapOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **corelationID** | **optional.**| corelation id is used to debug micro svc communication | 

### Return type

[**V2OnboardingConfirmNetworkMapResponse**](v2_onboarding.ConfirmNetworkMapResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RecommendationEvent**
> RecommendationEvent(ctx, body)
process the recommendation event

process the recommendation event

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RecommendationEventRequest**](RecommendationEventRequest.md)| request body | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetryExperimentCreation**
> NetworkmapRetryExperimentCreationResponse RetryExperimentCreation(ctx, body, accountIdentifier, organizationIdentifier, projectIdentifier, environmentIdentifier, infraId, applicationmapid)
retry chaos experiment creation for the given target network map

retry chaos experiment creation for the given target network map

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NetworkmapRetryExperimentCreationRequest**](NetworkmapRetryExperimentCreationRequest.md)| Retry experiment creation request | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **environmentIdentifier** | **string**| environment id that want to access the resource | 
  **infraId** | **string**| infra id that want to access the resource | 
  **applicationmapid** | **string**| Application map ID | 

### Return type

[**NetworkmapRetryExperimentCreationResponse**](networkmap.RetryExperimentCreationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RunChaosV2Experiment**
> ModelReRunChaosWorkflowResponse RunChaosV2Experiment(ctx, body, accountIdentifier, organizationIdentifier, projectIdentifier, experimentId, isOnboardingRun, isIdentity)
Run a chaos v2 experiment

Run a chaos v2 experiment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TypesExperimentRunRequest**](TypesExperimentRunRequest.md)| Run Experiment | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **experimentId** | **string**| experimentId to be executed | 
  **isOnboardingRun** | **bool**| is it onboarding run | 
  **isIdentity** | **bool**| is human-readable experiment identity passed | 

### Return type

[**ModelReRunChaosWorkflowResponse**](model.ReRunChaosWorkflowResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RunChaosV2InternalAPI**
> ModelReRunChaosWorkflowResponse RunChaosV2InternalAPI(ctx, body, accountIdentifier, organizationIdentifier, projectIdentifier, experimentId)
Run a chaos v2 experiment internal API

Run a chaos v2 experiment internal API

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TypesInternalExperimentRunRequest**](TypesInternalExperimentRunRequest.md)| Run Experiment | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **experimentId** | **string**| experimentId to be executed | 

### Return type

[**ModelReRunChaosWorkflowResponse**](model.ReRunChaosWorkflowResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RunRecommendation**
> RecommendationsRunActionResponse RunRecommendation(ctx, accountIdentifier, organizationIdentifier, projectIdentifier, recommendationID)
Run the recommended experiment

Run the recommended experiment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **recommendationID** | **string**| recommendation id | 

### Return type

[**RecommendationsRunActionResponse**](recommendations.RunActionResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SaveChaosV2Experiment**
> TypesExperimentCreationResponse SaveChaosV2Experiment(ctx, body, accountIdentifier, organizationIdentifier, projectIdentifier)
Save a chaos v2 experiment

Save a chaos v2 experiment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TypesExperimentCreationRequest**](TypesExperimentCreationRequest.md)| Save Experiment | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 

### Return type

[**TypesExperimentCreationResponse**](types.ExperimentCreationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StopChaosV2Experiment**
> TypesStopChaosV2ExperimentResponse StopChaosV2Experiment(ctx, accountIdentifier, organizationIdentifier, projectIdentifier, experimentId, experimentRunId, notifyId, optional)
Stop Chaos V2 experiment

Stop Chaos V2 experiment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **experimentId** | **string**| experiment id that will be used to stop the experiment | 
  **experimentRunId** | **string**| experiment run id that will be used to stop the experiment run | 
  **notifyId** | **string**| notify id that will be used to stop the experiment run | 
 **optional** | ***DefaultApiStopChaosV2ExperimentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiStopChaosV2ExperimentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **force** | **optional.Bool**| force stop the experiment run | 

### Return type

[**TypesStopChaosV2ExperimentResponse**](types.StopChaosV2ExperimentResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StopOnboardingExperiments**
> bool StopOnboardingExperiments(ctx, onboardingid, optional)
Stop V2 Onboarding experiments

Stop V2 Onboarding experiments

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **onboardingid** | **string**| onboarding id | 
 **optional** | ***DefaultApiStopOnboardingExperimentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiStopOnboardingExperimentsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **corelationID** | **optional.String**| corelation id is used to debug micro svc communication | 
 **accountIdentifier** | **optional.String**| account id that want to access the resource | 
 **organizationIdentifier** | **optional.String**| organization id that want to access the resource | 
 **projectIdentifier** | **optional.String**| project id that want to access the resource | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAction**
> GithubComHarnessHceSaasGraphqlServerPkgActionsAction UpdateAction(ctx, body, accountIdentifier, organizationIdentifier, projectIdentifier, identity)
Update a new action

Update a new action with the specified configuration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ActionsActionResponse**](ActionsActionResponse.md)| Action configuration | 
  **accountIdentifier** | **string**| Account ID to access the resource | 
  **organizationIdentifier** | **string**| Organization ID to access the resource | 
  **projectIdentifier** | **string**| Project ID to access the resource | 
  **identity** | **string**| ID of the Action to retrieve | 

### Return type

[**GithubComHarnessHceSaasGraphqlServerPkgActionsAction**](github_com_harness_hce-saas_graphql_server_pkg_actions.Action.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateActionTemplate**
> ChaosfaulttemplateActionTemplate UpdateActionTemplate(ctx, body, accountIdentifier, organizationIdentifier, projectIdentifier, identity, optional)
Updates the action templates in a hub

Updates an existing action template in a hub with new configuration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ChaosfaulttemplateActionTemplate**](ChaosfaulttemplateActionTemplate.md)| action configuration | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **identity** | **string**| ID of the Action to edit | 
 **optional** | ***DefaultApiUpdateActionTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiUpdateActionTemplateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **isDefault** | **optional.**| mark template as default | 

### Return type

[**ChaosfaulttemplateActionTemplate**](chaosfaulttemplate.ActionTemplate.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateChaosExperimentExecutionNode**
> string UpdateChaosExperimentExecutionNode(ctx, body, accountIdentifier, name, experimentId, experimentRunId)
Update chaos execution node

Update chaos execution node

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ExecutionChaosExecutionNode**](ExecutionChaosExecutionNode.md)| Create chaos execution node | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **name** | **string**| name of the node | 
  **experimentId** | **string**| experimentId to be fetched | 
  **experimentRunId** | **string**| experimentRunId to be fetched | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateChaosHub**
> Chaoshubv2GetHubResponse UpdateChaosHub(ctx, body, accountIdentifier, organizationIdentifier, projectIdentifier, hubIdentity)
Update chaos hub

Update chaos hub

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Chaoshubv2UpdateHubRequest**](Chaoshubv2UpdateHubRequest.md)| update hub request | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **hubIdentity** | **string**| chaos hub identity | 

### Return type

[**Chaoshubv2GetHubResponse**](chaoshubv2.GetHubResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateChaosV2CronExperiment**
> TypesUpdateCronExperimentStateResponse UpdateChaosV2CronExperiment(ctx, body, accountIdentifier, organizationIdentifier, projectIdentifier)
Update a chaos v2 cron experiment

Update a chaos v2 cron experiment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TypesUpdateCronExperimentStateRequest**](TypesUpdateCronExperimentStateRequest.md)| Update Cron V2 Experiment | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 

### Return type

[**TypesUpdateCronExperimentStateResponse**](types.UpdateCronExperimentStateResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEmissary**
> K8sinfraUpdateEmissaryUrlResponse UpdateEmissary(ctx, body, infrastructureIdentity, accountIdentifier, organizationIdentifier, projectIdentifier, environmentIdentifier)
Update emissary endpoint

Update emissary endpoint

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**K8sinfraUpdateEmissaryUrlRequest**](K8sinfraUpdateEmissaryUrlRequest.md)| update emissary request | 
  **infrastructureIdentity** | **string**| Chaos V2 Infra Identity | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **environmentIdentifier** | **string**| environment identifier to filter resource | 

### Return type

[**K8sinfraUpdateEmissaryUrlResponse**](k8sinfra.UpdateEmissaryURLResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFaultTemplate**
> ChaosfaulttemplateUpdateFaultTemplateResponse UpdateFaultTemplate(ctx, accountIdentifier, organizationIdentifier, projectIdentifier, hubIdentity, faultName)
Update the fault templates in a hub based on tag

Update the fault templates in a hub based on tag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **hubIdentity** | **string**| chaos hub identity | 
  **faultName** | **string**| name of the fault | 

### Return type

[**ChaosfaulttemplateUpdateFaultTemplateResponse**](chaosfaulttemplate.UpdateFaultTemplateResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateGamedayRunPrerequisitesV2**
> string UpdateGamedayRunPrerequisitesV2(ctx, body, accountIdentifier, organizationIdentifier, projectIdentifier, gamedayId, gamedayRunId)
Update a chaos v2 gameday run

Update a chaos v2 gameday run

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TypesUpdatePrerequisiteRequest**](TypesUpdatePrerequisiteRequest.md)| Update Gameday Run Prerequisite | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **gamedayId** | **string**| gamedayId of the run | 
  **gamedayRunId** | **string**| gamedayRunId to be updated | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateGamedayRunStakeHolderActionsV2**
> string UpdateGamedayRunStakeHolderActionsV2(ctx, body, accountIdentifier, organizationIdentifier, projectIdentifier, gamedayId, gamedayRunId)
Update a chaos v2 gameday run

Update a chaos v2 gameday run

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TypesUpdateStakeHolderActionRequest**](TypesUpdateStakeHolderActionRequest.md)| Update Gameday Run Stakeholder action | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **gamedayId** | **string**| gamedayId of the run | 
  **gamedayRunId** | **string**| gamedayRunId to be updated | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateGamedayRunV2**
> string UpdateGamedayRunV2(ctx, body, accountIdentifier, organizationIdentifier, projectIdentifier, gamedayId, gamedayRunId)
Update a chaos v2 gameday run

Update a chaos v2 gameday run

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TypesUpdateGameDayRunRequest**](TypesUpdateGameDayRunRequest.md)| Update Gameday Run | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **gamedayId** | **string**| gamedayId of the run | 
  **gamedayRunId** | **string**| gamedayRunId to be updated | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateGamedayV2**
> string UpdateGamedayV2(ctx, body, accountIdentifier, organizationIdentifier, projectIdentifier)
Update a chaos v2 gameday

Update a chaos v2 gameday

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TypesUpdateGamedayRequest**](TypesUpdateGamedayRequest.md)| Update Gameday | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateInputSet**
> InputsetsUpdateInputSetResponse UpdateInputSet(ctx, body, accountIdentifier, organizationIdentifier, projectIdentifier, experimentId, inputsetId, isIdentity)
Updates an input set

Updates an input set

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**InputsetsUpdateInputSetRequest**](InputsetsUpdateInputSetRequest.md)| update input set request | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **experimentId** | **string**| experimentId where input set should be created | 
  **inputsetId** | **string**| ID of the input set | 
  **isIdentity** | **bool**| is human-readable experiment identity passed | 

### Return type

[**InputsetsUpdateInputSetResponse**](inputsets.UpdateInputSetResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNote**
> string UpdateNote(ctx, body, accountIdentifier, organizationIdentifier, projectIdentifier)
Update a resource note

Update a resource note

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TypesUpdateNoteRequest**](TypesUpdateNoteRequest.md)| Update a note | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateProbe**
> TypesCreateProbeResponse UpdateProbe(ctx, body, accountIdentifier, organizationIdentifier, projectIdentifier, probeId)
Update a new probe

Update a new probe with the specified configuration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TypesProbeRequest**](TypesProbeRequest.md)| Probe configuration | 
  **accountIdentifier** | **string**| Account ID to access the resource | 
  **organizationIdentifier** | **string**| Organization ID to access the resource | 
  **projectIdentifier** | **string**| Project ID to access the resource | 
  **probeId** | **string**| ID of the probe to retrieve | 

### Return type

[**TypesCreateProbeResponse**](types.CreateProbeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateProbeTemplate**
> ChaosprobetemplateProbeTemplate UpdateProbeTemplate(ctx, body, accountIdentifier, organizationIdentifier, projectIdentifier, identity, optional)
Updates the probe templates in a hub

Updates an existing probe template in a hub with new configuration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ChaosprobetemplateProbeTemplate**](ChaosprobetemplateProbeTemplate.md)| probe configuration | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **identity** | **string**| identity of the probe to edit | 
 **optional** | ***DefaultApiUpdateProbeTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiUpdateProbeTemplateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **isDefault** | **optional.**| mark template as default | 

### Return type

[**ChaosprobetemplateProbeTemplate**](chaosprobetemplate.ProbeTemplate.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRecommendationStatus**
> UpdateRecommendationStatus(ctx, accountIdentifier, organizationIdentifier, projectIdentifier, recommendationID, status)
Update the recommendation status

Update the recommendation status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **organizationIdentifier** | **string**| organization id that want to access the resource | 
  **projectIdentifier** | **string**| project id that want to access the resource | 
  **recommendationID** | **string**| recommendation id | 
  **status** | **string**| status | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

