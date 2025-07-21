package model

import (
	"fmt"
	"io"
	"strconv"
)

type Audit interface {
	IsAudit()
}

type AuditV2 interface {
	IsAuditV2()
}

// Defines the common probe properties shared across different ProbeTypes
type CommonProbeProperties interface {
	IsCommonProbeProperties()
}

type ResourceDetails interface {
	IsResourceDetails()
}

// Defines the APM appdynamics probe properties
type APMAppDynamicsProbe struct {
	// connectorID of the apm appdynamics probe
	ConnectorID string `json:"connectorID"`
	// appdynamics metrics properties
	AppdMetrics *AppdMetrics `json:"appdMetrics"`
}

// Defines the APM dynatrace probe properties
type APMDynatraceProbe struct {
	// connectorID of the apm splunk observability probe
	ConnectorID string `json:"connectorID"`
	// duration in minutes for splunk observability probe
	DurationInMin int `json:"durationInMin"`
	// dynatrace metric
	DynatraceMetrics *DynatraceMetrics `json:"dynatraceMetrics"`
}

// Defines the APM probe properties
type APMProbe struct {
	// Timeout of the Probe
	ProbeTimeout string `json:"probeTimeout"`
	// Interval of the Probe
	Interval string `json:"interval"`
	// Retry interval of the Probe
	Retry *int `json:"retry"`
	// Attempt contains the total attempt count for the probe
	Attempt int `json:"attempt"`
	// Polling interval of the Probe
	ProbePollingInterval *string `json:"probePollingInterval"`
	// Initial delay interval of the Probe in seconds
	InitialDelay *string `json:"initialDelay"`
	// Is stop on failure enabled in the Probe
	StopOnFailure *bool `json:"stopOnFailure"`
	// type of the apm probe
	Type string `json:"type"`
	// apm Prometheus probe properties
	ApmPrometheusProbe *APMPrometheusProbe `json:"apmPrometheusProbe"`
	// apm AppDynamics probe properties
	ApmAppDynamicsProbe *APMAppDynamicsProbe `json:"apmAppDynamicsProbe"`
	// apm SplunkObservability probe properties
	ApmSplunkObservabilityProbe *APMSplunkObservabilityProbe `json:"apmSplunkObservabilityProbe"`
	// apm Dynatrace probe properties
	ApmDynatraceProbe *APMDynatraceProbe `json:"apmDynatraceProbe"`
	// Comparator of the Probe
	Comparator *Comparator `json:"comparator"`
	// Verbosity for the probe logging
	Verbosity *string `json:"verbosity"`
}

func (APMProbe) IsCommonProbeProperties() {}

// Defines the APM prometheus probe properties
type APMPrometheusProbe struct {
	// query of the apm probe
	Query string `json:"query"`
	// connectorID of the apm prometheus probe
	ConnectorID string `json:"connectorID"`
	// apm prometheus probe tls configurations
	TLSConfig *APMTLSConfig `json:"tlsConfig"`
}

// Defines the APM Splunk Observability probe properties
type APMSplunkObservabilityProbe struct {
	// connectorID of the apm splunk observability probe
	ConnectorID string `json:"connectorID"`
	// splunk observability metrics properties
	SplunkObservabilityMetrics *SplunkObservabilityMetrics `json:"splunkObservabilityMetrics"`
}

// APMTLSConfig configures the options for TLS connections
type APMTLSConfig struct {
	// Flag to hold the ca file path
	CaCrt *SecretManager `json:"caCrt"`
	// Flag to hold the client cert file path
	ClientCrt *SecretManager `json:"clientCrt"`
	// Flag to hold the client key file path
	Key *SecretManager `json:"key"`
	// Flag to skip the tls certificates checks
	InsecureSkipVerify *bool `json:"insecureSkipVerify"`
}

type ActionItemRequest struct {
	Item   string `json:"item"`
	IsDone bool   `json:"isDone"`
}

type ActionItems struct {
	Item   string `json:"item"`
	IsDone bool   `json:"isDone"`
}

type ActionPayload struct {
	RequestID    string  `json:"requestID"`
	RequestType  string  `json:"requestType"`
	K8sManifest  string  `json:"k8sManifest"`
	Namespace    string  `json:"namespace"`
	ExternalData *string `json:"externalData"`
	UUID         *string `json:"uuid"`
}

type Annotation struct {
	Categories       string `json:"categories"`
	Vendor           string `json:"vendor"`
	CreatedAt        string `json:"createdAt"`
	Repository       string `json:"repository"`
	Support          string `json:"support"`
	ChartDescription string `json:"chartDescription"`
}

// Defines the APM appdynamics metrics properties
type AppdMetrics struct {
	// full path of the metrics for appdynamics probe
	MetricsFullPath string `json:"metricsFullPath"`
	// application name for the appdynamics probe
	ApplicationName string `json:"applicationName"`
	// duration in minutes for appdynamics probe
	DurationInMin int `json:"durationInMin"`
}

type ApplicationSpec struct {
	Operator  Operator    `json:"operator"`
	Workloads []*Workload `json:"workloads"`
}

type ApplicationSpecInput struct {
	Operator  Operator         `json:"operator"`
	Workloads []*WorkloadInput `json:"workloads"`
}

// Defines the details for a chaos experiment
type ChaosExperimentRequest struct {
	// ID of the experiment
	ID string `json:"id"`
	// Identity of the experiment
	Identity *string `json:"identity"`
	// Name of the experiment
	Name string `json:"name"`
	// Description of the experiment
	Description *string `json:"description"`
	// Manifest of the experiment
	Manifest string `json:"manifest"`
	// Array containing service identifier and environment identifier
	// for SRM change source events
	EventsMetadata []*EventMetadataInput `json:"eventsMetadata"`
	// ID of the target infrastructure in which the experiment will run
	InfraID string `json:"infraID"`
	// Tags of the infrastructure
	Tags []string `json:"tags"`
	// Type of the infrastructure
	InfraType *InfrastructureType `json:"infraType"`
	// Validate the experiment manifest
	ValidateManifest *bool `json:"validateManifest"`
	// Cron syntax of the workflow schedule
	CronSyntax *string `json:"cronSyntax"`
	// Flag to check if single run status is enabled or not
	IsSingleRunCronEnabled *bool `json:"isSingleRunCronEnabled"`
	// type of experiment
	ExperimentType *WorkflowType `json:"experimentType"`
}

type ChaosHub struct {
	// ID of the chaos hub
	ID string `json:"id"`
	// Harness identifiers
	Identifiers *Identifiers `json:"identifiers"`
	// identity of the hub
	Identity string `json:"identity"`
	// Name of the repository if connector is of Account type
	RepoName *string `json:"repoName"`
	// URL of the git repository
	RepoURL string `json:"repoURL"`
	// Branch of the git repository
	RepoBranch string `json:"repoBranch"`
	// AuthType
	AuthType ChaosHubAuthType `json:"AuthType"`
	// Name of the GitConnectorId
	ConnectorID string `json:"connectorId"`
	// Name of the ConnectorScope
	ConnectorScope ConnectorScope `json:"connectorScope"`
	// Reference of the connector
	ConnectorRef string `json:"connectorRef"`
	// Name of the chaos hub
	Name string `json:"name"`
	// Timestamp when the chaos hub was created
	CreatedAt string `json:"createdAt"`
	// Timestamp when the chaos hub was last updated
	UpdatedAt string `json:"updatedAt"`
	// Timestamp when the chaos hub was last synced
	LastSyncedAt string `json:"lastSyncedAt"`
	// Default Hub Identifier
	IsDefault bool `json:"isDefault"`
	// Tags of the ChaosHub
	Tags []string `json:"tags"`
	// User who created the ChaosHub
	CreatedBy *UserDetails `json:"createdBy"`
	// User who has updated the ChaosHub
	UpdatedBy *UserDetails `json:"updatedBy"`
	// Description of ChaosHub
	Description *string `json:"description"`
	// Connectivity status of the ChaosHub
	Status string `json:"status"`
}

func (ChaosHub) IsResourceDetails() {}
func (ChaosHub) IsAudit()           {}

// Defines filter options for ChaosHub
type ChaosHubFilterInput struct {
	// Name of the ChaosHub
	ChaosHubName *string `json:"chaosHubName"`
	// Tags of a chaos hub
	Tags []string `json:"tags"`
	// Description of a chaos hub
	Description *string `json:"description"`
}

// Defines the details required for creating a chaos hub
type ChaosHubRequest struct {
	// Name of the chaos hub
	HubName string `json:"hubName"`
	// Name of the GitConnectorId
	ConnectorID string `json:"connectorId"`
	// Name of the ConnectorScope
	ConnectorScope ConnectorScope `json:"connectorScope"`
	// Repo name of the git repository
	RepoName *string `json:"repoName"`
	// Branch of the git repository
	RepoBranch string `json:"repoBranch"`
	// Tags of the ChaosHub
	Tags []string `json:"tags"`
	// Description of ChaosHub
	Description *string `json:"description"`
}

type ChaosHubStatus struct {
	// ID of the hub
	ID string `json:"id"`
	// Identity of the hub
	Identity string `json:"identity"`
	// Name of the repository if connector is of Account type
	RepoName *string `json:"repoName"`
	// URL of the git repository
	RepoURL string `json:"repoURL"`
	// Branch of the git repository
	RepoBranch string `json:"repoBranch"`
	// Name of the GitConnectorId
	ConnectorID string `json:"connectorId"`
	// Name of the ConnectorScope
	ConnectorScope ConnectorScope `json:"connectorScope"`
	// AuthType
	AuthType ChaosHubAuthType `json:"AuthType"`
	// Bool value indicating whether the hub is available or not.
	IsAvailable bool `json:"isAvailable"`
	// Total number of experiments in the hub
	TotalFaults int `json:"totalFaults"`
	// Total workflows
	TotalExperiments int `json:"totalExperiments"`
	// Name of the chaos hub
	Name string `json:"name"`
	// Timestamp when the chaos hub was last synced
	LastSyncedAt string `json:"lastSyncedAt"`
	// Default Hub Identifier
	IsDefault bool `json:"isDefault"`
	// Tags of the ChaosHub
	Tags []string `json:"tags"`
	// User who created the ChaosHub
	CreatedBy *UserDetails `json:"createdBy"`
	// User who has updated the ChaosHub
	UpdatedBy *UserDetails `json:"updatedBy"`
	// Created at timestamp
	CreatedAt string `json:"createdAt"`
	// Updated at timestamp
	UpdatedAt string `json:"updatedAt"`
	// Description of ChaosHub
	Description *string `json:"description"`
}

func (ChaosHubStatus) IsResourceDetails() {}
func (ChaosHubStatus) IsAudit()           {}

type ChaosServiceAccountSpec struct {
	Operator        Operator `json:"operator"`
	ServiceAccounts []string `json:"serviceAccounts"`
}

type ChaosServiceAccountSpecInput struct {
	Operator        Operator `json:"operator"`
	ServiceAccounts []string `json:"serviceAccounts"`
}

// Defines the details for a chaos workflow
type ChaosWorkFlowRequest struct {
	// ID of the workflow
	WorkflowID *string `json:"workflowID"`
	// Identity of the experiment
	Identity *string `json:"identity"`
	// Boolean check indicating if the created scenario will be executed or not
	RunExperiment *bool `json:"runExperiment"`
	// Manifest of the workflow
	WorkflowManifest string `json:"workflowManifest"`
	// Type of the workflow
	WorkflowType *WorkflowType `json:"workflowType"`
	// Cron syntax of the workflow schedule
	CronSyntax string `json:"cronSyntax"`
	// Flag to check if single run status is enabled or not
	IsSingleRunCronEnabled *bool `json:"isSingleRunCronEnabled"`
	// Name of the workflow
	WorkflowName string `json:"workflowName"`
	// Description of the workflow
	WorkflowDescription string `json:"workflowDescription"`
	// Array containing weightage and name of each chaos experiment in the workflow
	Weightages []*WeightagesInput `json:"weightages"`
	// Array containing service identifier and environment identifier
	// for SRM change source events
	EventsMetadata []*EventMetadataInput `json:"eventsMetadata"`
	// Bool value indicating whether the workflow is a custom workflow or not
	IsCustomWorkflow bool `json:"isCustomWorkflow"`
	// ID of the target infra in which the workflow will run
	InfraID string `json:"infraID"`
	// Tags of the infra
	Tags []string `json:"tags"`
	// Type of the infra
	InfraType *InfrastructureType `json:"infraType"`
}

// Defines the response received for querying the details of chaos workflow
type ChaosWorkFlowResponse struct {
	// ID of the workflow
	WorkflowID string `json:"workflowID"`
	// Harness Identifiers
	Identifiers *Identifiers `json:"identifiers"`
	// Cron syntax of the workflow schedule
	CronSyntax string `json:"cronSyntax"`
	// Name of the workflow
	WorkflowName string `json:"workflowName"`
	// Description of the workflow
	WorkflowDescription string `json:"workflowDescription"`
	// Bool value indicating whether the workflow is a custom workflow or not
	IsCustomWorkflow bool `json:"isCustomWorkflow"`
	// Tags of the infra
	Tags []string `json:"tags"`
}

type Chart struct {
	APIVersion  string              `json:"apiVersion"`
	Kind        string              `json:"kind"`
	Metadata    *Metadata           `json:"metadata"`
	Spec        *Spec               `json:"spec"`
	PackageInfo *PackageInformation `json:"packageInfo"`
}

type CheckImageRegistryOverrideResponse struct {
	OverrideBlockedByScope string                 `json:"OverrideBlockedByScope"`
	ImageRegistry          *ImageRegistryResponse `json:"ImageRegistry"`
}

type CheckResourceIDRequest struct {
	ResourceName ResourceType `json:"resourceName"`
	ID           string       `json:"ID"`
}

// Defines the details for a infra
type CloudFoundryInfra struct {
	// ID of the infra
	InfraID string `json:"infraID"`
	// Name of the infra
	Name string `json:"name"`
	// Description of the infra
	Description *string `json:"description"`
	// Tags of the infra
	Tags []string `json:"tags"`
	// Environment ID for the infra
	EnvironmentID string `json:"environmentID"`
	// Boolean value indicating if chaos infrastructure is active or not
	IsActive bool `json:"isActive"`
	// Boolean value indicating if chaos infrastructure is confirmed or not
	IsInfraConfirmed bool `json:"isInfraConfirmed"`
	// Timestamp when the infra was last updated
	UpdatedAt string `json:"updatedAt"`
	// Timestamp when the infra was created
	CreatedAt string `json:"createdAt"`
	// Number of schedules created in the infra
	NoOfSchedules *int `json:"noOfSchedules"`
	// Number of workflows run in the infra
	NoOfWorkflows *int `json:"noOfWorkflows"`
	// Timestamp of the last workflow run in the infra
	LastWorkflowTimestamp *string `json:"lastWorkflowTimestamp"`
	// Timestamp when the infra got connected
	StartTime string `json:"startTime"`
	// Version of the infra
	Version string `json:"version"`
	// User who created the infra
	CreatedBy *UserDetails `json:"createdBy"`
	// User who has updated the infra
	UpdatedBy *UserDetails `json:"updatedBy"`
	// Last Heartbeat status sent by the infra
	LastHeartbeat *string `json:"lastHeartbeat"`
	// hostname of the infra
	Hostname *string `json:"hostname"`
}

func (CloudFoundryInfra) IsResourceDetails() {}
func (CloudFoundryInfra) IsAudit()           {}

// Defines filter options for infras
type CloudFoundryInfraFilterInput struct {
	// Name of the infra
	Name *string `json:"name"`
	// ID of the infra
	InfraID *string `json:"infraID"`
	// ID of the infra
	Description *string `json:"description"`
	// Status of infra
	IsActive *bool `json:"isActive"`
	// Tags of an infra
	Tags []*string `json:"tags"`
}

// Defines the properties of the comparator
type Comparator struct {
	// Type of the Comparator
	Type string `json:"type"`
	// Value of the Comparator
	Value string `json:"value"`
	// Operator of the Comparator
	Criteria string `json:"criteria"`
}

// Defines the input properties of the comparator
type ComparatorInput struct {
	// Type of the Comparator
	Type string `json:"type"`
	// Value of the Comparator
	Value string `json:"value"`
	// Operator of the Comparator
	Criteria string `json:"criteria"`
}

type Condition struct {
	Name        string             `json:"name"`
	Description *string            `json:"description"`
	Tags        []string           `json:"tags"`
	ConditionID string             `json:"conditionId"`
	FaultSpec   *FaultSpec         `json:"faultSpec"`
	InfraType   InfrastructureType `json:"infraType"`
	K8sSpec     *K8sSpec           `json:"k8sSpec"`
	MachineSpec *MachineSpec       `json:"machineSpec"`
	Rules       []*ConditionRule   `json:"rules"`
}

func (Condition) IsResourceDetails() {}

type ConditionDetails struct {
	ConditionID   *string                  `json:"conditionId"`
	ConditionName *string                  `json:"conditionName"`
	Message       *string                  `json:"message"`
	Phase         *SecurityGovernancePhase `json:"phase"`
}

type ConditionFilterInput struct {
	Name        *string   `json:"name"`
	Description *string   `json:"description"`
	Tags        []*string `json:"tags"`
}

type ConditionRequest struct {
	ConditionID string             `json:"conditionId"`
	Description *string            `json:"description"`
	Name        string             `json:"name"`
	Tags        []*string          `json:"tags"`
	InfraType   InfrastructureType `json:"infraType"`
	FaultSpec   *FaultSpecInput    `json:"faultSpec"`
	K8sSpec     *K8sSpecInput      `json:"k8sSpec"`
	MachineSpec *MachineSpecInput  `json:"machineSpec"`
}

type ConditionResponse struct {
	UpdatedAt   int          `json:"updatedAt"`
	CreatedAt   int          `json:"createdAt"`
	UpdatedBy   *UserDetails `json:"updatedBy"`
	CreatedBy   *UserDetails `json:"createdBy"`
	Identifiers *Identifiers `json:"identifiers"`
	Condition   *Condition   `json:"condition"`
}

func (ConditionResponse) IsAuditV2() {}

type ConditionRule struct {
	RuleID string `json:"ruleId"`
	Name   string `json:"name"`
}

type ConfirmInfraRegistrationResponse struct {
	IsInfraConfirmed bool    `json:"isInfraConfirmed"`
	NewAccessKey     *string `json:"newAccessKey"`
	InfraID          *string `json:"infraID"`
}

type CreateGameDayExperimentRequest struct {
	GameDayRunID           string `json:"gameDayRunID"`
	ExperimentID           string `json:"experimentID"`
	HubID                  string `json:"hubID"`
	ExperimentTemplateName string `json:"experimentTemplateName"`
	ChaosInfraID           string `json:"chaosInfraID"`
}

type CreateGameDayRequest struct {
	ID          string               `json:"ID"`
	Name        string               `json:"name"`
	Experiments []*NonCronExperiment `json:"experiments"`
	Objective   *string              `json:"objective"`
	Description *string              `json:"description"`
}

type CreateGameDayRunRequest struct {
	GamedayID string `json:"gamedayID"`
}

type CustomImages struct {
	LogWatcher *string `json:"logWatcher"`
	Ddcr       *string `json:"ddcr"`
	DdcrLib    *string `json:"ddcrLib"`
	DdcrFault  *string `json:"ddcrFault"`
}

type CustomImagesRequest struct {
	LogWatcher *string `json:"logWatcher"`
	Ddcr       *string `json:"ddcr"`
	DdcrLib    *string `json:"ddcrLib"`
	DdcrFault  *string `json:"ddcrFault"`
}

// Raw metrics details of the datadog probe
type DatadogMetrics struct {
	// Query to get Datadog metrics
	DatadogQuery string `json:"datadogQuery"`
	// Timeframe of the metric
	TimeFrame string `json:"timeFrame"`
	// Comparator check for the correctness of the probe output
	Comparator *Comparator `json:"comparator"`
}

// Defines the input for Raw metrics details of the datadog probe
type DatadogMetricsInput struct {
	// Query to get Datadog metrics
	DatadogQuery string `json:"datadogQuery"`
	// Timeframe of the metric
	TimeFrame string `json:"timeFrame"`
	// Comparator check for the correctness of the probe output
	Comparator *ComparatorInput `json:"comparator"`
}

// Defines the start date and end date for the filtering the data
type DateRange struct {
	// Start date
	StartDate string `json:"startDate"`
	// End date
	EndDate *string `json:"endDate"`
}

type DeleteGameDayRequest struct {
	GamedayID string `json:"gamedayID"`
}

type DuplicateGameDayRequest struct {
	GamedayID string `json:"gamedayID"`
}

// Defines the dynatrace metrics properties
type DynatraceMetrics struct {
	// dynatrace metrcis selector
	MetricsSelector string `json:"metricsSelector"`
	// dynatarce entity selector
	EntitySelector string `json:"entitySelector"`
}

type Env struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type EnvInput struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// EvaluationWindow is the time period for which the SLO probe will work
type EvaluationWindow struct {
	// Start time of evaluation
	EvaluationStartTime int `json:"evaluationStartTime"`
	// End time of evaluation
	EvaluationEndTime int `json:"evaluationEndTime"`
}

// Defines the input properties of EvaluationWindow
type EvaluationWindowInput struct {
	// Start time of evaluation
	EvaluationStartTime int `json:"evaluationStartTime"`
	// End time of evaluation
	EvaluationEndTime int `json:"evaluationEndTime"`
}

type EventMetadata struct {
	FaultName             string   `json:"faultName"`
	ServiceIdentifier     []string `json:"serviceIdentifier"`
	EnvironmentIdentifier []string `json:"environmentIdentifier"`
}

type EventMetadataInput struct {
	FaultName             string   `json:"faultName"`
	ServiceIdentifier     []string `json:"serviceIdentifier"`
	EnvironmentIdentifier []string `json:"environmentIdentifier"`
}

// Defines the Executed by which experiment details for Probes
type ExecutedByExperiment struct {
	// Experiment ID
	ExperimentID string `json:"experimentID"`
	// Experiment Run ID
	ExperimentRunID string `json:"experimentRunID"`
	// Notify ID
	NotifyID string `json:"notifyID"`
	// Experiment Name
	ExperimentName string `json:"experimentName"`
	// Type of the experiment i.e. CRON, NON_CRON or Gameday
	ExperimentType ScenarioType `json:"experimentType"`
	// Timestamp at which the experiment was last updated
	UpdatedAt int `json:"updatedAt"`
	// User who has updated the experiment
	UpdatedBy *UserDetails `json:"updatedBy"`
}

// Defines the Execution History of experiment referenced by the Probe
type ExecutionHistory struct {
	// Probe Mode
	Mode Mode `json:"mode"`
	// Fault Name
	FaultName string `json:"faultName"`
	// Fault Status
	Status *Status `json:"status"`
	// Fault executed by which experiment
	ExecutedByExperiment *ExecutedByExperiment `json:"executedByExperiment"`
}

type ExperimentInfoInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ExperimentRequest struct {
	// Name of the chart being used
	Category string `json:"category"`
	// Name of the experiment
	ExperimentName string `json:"experimentName"`
	// Name of the hub
	HubID string `json:"hubID"`
	// To fetch experiment details from only engine if infra compatible with new format
	CompatibleWithNewExp *bool `json:"compatibleWithNewExp"`
}

type ExperimentResponse struct {
	ExperimentID           string `json:"experimentID"`
	ExperimentTemplateName string `json:"experimentTemplateName"`
	ChaosInfraID           string `json:"chaosInfraID"`
	HubID                  string `json:"hubID"`
}

// Defines the details of a experiment run report
type ExperimentRunReport struct {
	// Harness identifiers
	Identifiers *Identifiers `json:"identifiers"`
	// ID of the workflow run which is to be queried
	WorkflowRunID string `json:"workflowRunID"`
	// Type of the workflow
	WorkflowType *string `json:"workflowType"`
	// ID of the workflow
	WorkflowID string `json:"workflowID"`
	// Array containing weightage and name of each chaos experiment in the workflow
	Weightages []*Weightages `json:"weightages"`
	// Timestamp at which workflow run was last updated
	UpdatedAt string `json:"updatedAt"`
	// Timestamp at which workflow run was created
	CreatedAt string `json:"createdAt"`
	// Cron syntax of the workflow schedule
	CronSyntax string `json:"cronSyntax"`
	// Target infra in which the workflow will run
	Infra []*Infrastructure `json:"infra"`
	// Name of the workflow
	WorkflowName string `json:"workflowName"`
	// Description of the workflow
	WorkflowDescription *string `json:"workflowDescription"`
	// Tag of the workflow
	WorkflowTags []string `json:"workflowTags"`
	// Manifest of the workflow run
	WorkflowManifest string `json:"workflowManifest"`
	// If cron is enabled or disabled
	IsCronEnabled *bool `json:"isCronEnabled"`
	// Flag to check if single run status is enabled or not
	IsSingleRunCronEnabled *bool `json:"isSingleRunCronEnabled"`
	// Probe object containing reference of probeIDs
	Probe []*ProbeMap `json:"probe"`
	// Phase of the workflow run
	Phase WorkflowRunStatus `json:"phase"`
	// Resiliency score of the workflow
	ResiliencyScore *float64 `json:"resiliencyScore"`
	// Number of experiments passed
	ExperimentsPassed *int `json:"experimentsPassed"`
	// Number of experiments failed
	ExperimentsFailed *int `json:"experimentsFailed"`
	// Number of experiments awaited
	ExperimentsAwaited *int `json:"experimentsAwaited"`
	// Number of experiments stopped
	ExperimentsStopped *int `json:"experimentsStopped"`
	// Number of experiments which are not available
	ExperimentsNa *int `json:"experimentsNa"`
	// Total number of experiments
	TotalExperiments *int `json:"totalExperiments"`
	// Stores all the workflow run details related to the nodes of DAG graph and chaos results of the experiments
	ExecutionData string `json:"executionData"`
	// Bool value indicating if the workflow run has removed
	IsRemoved *bool `json:"isRemoved"`
	// User who has updated the workflow
	UpdatedBy *UserDetails `json:"updatedBy"`
	// User who has created the experiment run
	CreatedBy *UserDetails `json:"createdBy"`
	// Notify ID of the experiment run
	NotifyID *string `json:"notifyID"`
	// Error Response is the reason why experiment failed to run
	ErrorResponse *string `json:"errorResponse"`
	// Security Governance details of the workflow run
	SecurityGovernance *SecurityGovernance `json:"securityGovernance"`
	// runSequence is the sequence number of experiment run
	RunSequence int `json:"runSequence"`
}

func (ExperimentRunReport) IsAudit() {}

// Defines the experiment variables
type ExperimentVariables struct {
	// Identity of the input set
	InputSetIdentity *string `json:"inputSetIdentity"`
	// Run time inputs
	RunTimeInputs *RunTimeInputs `json:"runTimeInputs"`
}

type Experiments struct {
	Name string `json:"name"`
	Csv  string `json:"CSV"`
	Desc string `json:"desc"`
}

type Fault struct {
	FaultType FaultType `json:"faultType"`
	Name      string    `json:"name"`
}

// Fault Detail consists of all the fault related details
type FaultDetails struct {
	// fault consists of fault.yaml
	Fault string `json:"fault"`
	// engine consists engine.yaml
	Engine string `json:"engine"`
	// csv consists chartserviceversion.yaml
	Csv string `json:"csv"`
	// k8sSpec
	K8sSpec string `json:"k8sSpec"`
}

type FaultList struct {
	Name             string   `json:"name"`
	DisplayName      string   `json:"displayName"`
	Description      string   `json:"description"`
	Plan             []string `json:"plan"`
	SupportedVersion []string `json:"supportedVersion"`
	ChaosType        *string  `json:"chaosType"`
}

type FaultResponse struct {
	FaultType FaultType `json:"faultType"`
	Name      string    `json:"name"`
}

type FaultSpec struct {
	Operator Operator         `json:"operator"`
	Faults   []*FaultResponse `json:"faults"`
}

type FaultSpecInput struct {
	Faults   []*Fault `json:"faults"`
	Operator Operator `json:"operator"`
}

// Details of GET request
type Get struct {
	// Criteria of the response
	Criteria string `json:"criteria"`
	// Response Code of the response
	ResponseCode *string `json:"responseCode"`
	// Response Body of the response
	ResponseBody *string `json:"responseBody"`
}

// Details for input of GET request
type GETRequest struct {
	// Criteria of the response
	Criteria string `json:"criteria"`
	// Response Code of the response
	ResponseCode *string `json:"responseCode"`
	// Response Body of the response
	ResponseBody *string `json:"responseBody"`
}

type GameDay struct {
	GameDayID *string `json:"gameDayID"`
	// Harness identifiers
	Identifiers *Identifiers          `json:"identifiers"`
	Name        string                `json:"name"`
	Experiments []*ExperimentResponse `json:"experiments"`
	Objective   *string               `json:"objective"`
	Description *string               `json:"description"`
	CreatedBy   string                `json:"createdBy"`
	CreatedAt   string                `json:"createdAt"`
	UpdatedAt   string                `json:"updatedAt"`
	Summary     *GameDaySummary       `json:"summary"`
	IsRemoved   bool                  `json:"isRemoved"`
}

type GameDayResponse struct {
	GameDayID string `json:"gameDayID"`
	// Harness identifiers
	Identifiers *Identifiers          `json:"identifiers"`
	Name        string                `json:"name"`
	Tags        []string              `json:"tags"`
	Experiments []*ExperimentResponse `json:"experiments"`
	Objective   *string               `json:"objective"`
	Description *string               `json:"description"`
	CreatedBy   *UserDetails          `json:"createdBy"`
	UpdatedBy   *UserDetails          `json:"updatedBy"`
	CreatedAt   string                `json:"createdAt"`
	UpdatedAt   string                `json:"updatedAt"`
	Summary     *GameDaySummary       `json:"summary"`
	IsRemoved   bool                  `json:"isRemoved"`
}

func (GameDayResponse) IsAudit()           {}
func (GameDayResponse) IsResourceDetails() {}

type GameDayRunExperiments struct {
	ExperimentID           string    `json:"experimentID"`
	ExperimentTemplateName string    `json:"experimentTemplateName"`
	ChaosInfraID           string    `json:"chaosInfraID"`
	HubID                  string    `json:"hubID"`
	ExperimentRunIDs       []*string `json:"experimentRunIDs"`
	ExperimentNotes        *string   `json:"experimentNotes"`
}

type GameDayRunResponse struct {
	GameDayRunID string `json:"gameDayRunID"`
	Name         string `json:"name"`
	GameDayID    string `json:"gameDayID"`
	// Harness identifiers
	Identifiers *Identifiers             `json:"identifiers"`
	Experiments []*GameDayRunExperiments `json:"experiments"`
	StartTime   *string                  `json:"startTime"`
	EndTime     *string                  `json:"endTime"`
	Notes       *string                  `json:"notes"`
	Summary     *GameDaySummary          `json:"summary"`
	CreatedBy   *UserDetails             `json:"createdBy"`
	UpdatedBy   *UserDetails             `json:"updatedBy"`
	UpdatedAt   *string                  `json:"updatedAt"`
	CreatedAt   *string                  `json:"createdAt"`
	Completed   *bool                    `json:"completed"`
}

type GameDaySummary struct {
	Notes      *string        `json:"notes"`
	Qna        []*QnAs        `json:"qna"`
	ActionItem []*ActionItems `json:"actionItem"`
}

type GamedayFilterInput struct {
	GamedayName *string `json:"gamedayName"`
}

type GamedayInfraDetails struct {
	ID            string      `json:"ID"`
	Type          *string     `json:"type"`
	Name          *string     `json:"name"`
	EnvironmentID *string     `json:"environmentID"`
	Namespace     *string     `json:"namespace"`
	Scope         *InfraScope `json:"scope"`
	IsActive      *bool       `json:"isActive"`
}

// Defines sorting options for workflow runs
type GamedaySortInput struct {
	// Field in which sorting will be done
	Field GamedaySortingField `json:"field"`
	// Bool value indicating whether the sorting will be done in ascending order
	Ascending *bool `json:"ascending"`
}

type GetChaosHubStatsResponse struct {
	// Total number of chaoshubs
	TotalChaosHubs int `json:"totalChaosHubs"`
}

// Defines the details for a given experiment with some additional data
type GetExperimentResponse struct {
	// Details of experiment
	ExperimentDetails *Workflow `json:"experimentDetails"`
	// Average resiliency score of the experiment
	AverageResiliencyScore *float64 `json:"averageResiliencyScore"`
}

type GetExperimentRunStatsResponse struct {
	// Total number of experiment runs
	TotalExperimentRuns int `json:"totalExperimentRuns"`
	// Total number of completed experiments runs
	TotalCompletedExperimentRuns int `json:"totalCompletedExperimentRuns"`
	// Total number of stopped experiment runs
	TotalTerminatedExperimentRuns int `json:"totalTerminatedExperimentRuns"`
	// Total number of running experiment runs
	TotalRunningExperimentRuns int `json:"totalRunningExperimentRuns"`
	// Total number of stopped experiment runs
	TotalStoppedExperimentRuns int `json:"totalStoppedExperimentRuns"`
	// Total number of errored experiment runs
	TotalErroredExperimentRuns int `json:"totalErroredExperimentRuns"`
}

type GetExperimentStatsResponse struct {
	// Total number of experiments
	TotalExperiments int `json:"totalExperiments"`
	// Total number of cron experiments
	TotalExpCategorizedByResiliencyScore []*ResilienceScoreCategory `json:"totalExpCategorizedByResiliencyScore"`
}

type GetGameDayResponse struct {
	GameDayID string `json:"gameDayID"`
	// Harness identifiers
	Identifiers      *Identifiers                    `json:"identifiers"`
	Name             string                          `json:"name"`
	Experiments      []*GetGamedayExperimentResponse `json:"experiments"`
	Objective        *string                         `json:"objective"`
	Description      *string                         `json:"description"`
	CreatedBy        *UserDetails                    `json:"createdBy"`
	CreatedAt        *string                         `json:"createdAt"`
	UpdatedAt        *string                         `json:"updatedAt"`
	Summary          *GameDaySummary                 `json:"summary"`
	IsRemoved        *bool                           `json:"isRemoved"`
	TotalGamedayRuns *int                            `json:"totalGamedayRuns"`
}

type GetGameDayRunResponse struct {
	GameDayRunID string `json:"gameDayRunID"`
	Name         string `json:"name"`
	GameDayID    string `json:"gameDayID"`
	// Harness identifiers
	Identifiers *Identifiers             `json:"identifiers"`
	Experiments []*GameDayRunExperiments `json:"experiments"`
	StartTime   *string                  `json:"startTime"`
	EndTime     *string                  `json:"endTime"`
	Notes       *string                  `json:"notes"`
	Summary     *GameDaySummary          `json:"summary"`
	CreatedBy   *UserDetails             `json:"createdBy"`
	UpdatedBy   *UserDetails             `json:"updatedBy"`
	UpdatedAt   *string                  `json:"updatedAt"`
	CreatedAt   *string                  `json:"createdAt"`
	Completed   *bool                    `json:"completed"`
}

type GetGamedayExperimentResponse struct {
	ExperimentID           string               `json:"experimentID"`
	ExperimentTemplateName string               `json:"experimentTemplateName"`
	ChaosInfra             *GamedayInfraDetails `json:"chaosInfra"`
	HubID                  string               `json:"hubID"`
	ExperimentNotes        *string              `json:"experimentNotes"`
	ExperimentCsv          *string              `json:"experimentCSV"`
	ExperimentManifest     *string              `json:"experimentManifest"`
}

type GetInfraStatsResponse struct {
	// Total number of infrastructures
	TotalInfrastructures int `json:"totalInfrastructures"`
	// Total number of active infrastructures
	TotalActiveInfrastructure int `json:"totalActiveInfrastructure"`
	// Total number of inactive infrastructures
	TotalInactiveInfrastructures int `json:"totalInactiveInfrastructures"`
	// Total number of confirmed infrastructures
	TotalConfirmedInfrastructure int `json:"totalConfirmedInfrastructure"`
	// Total number of non confirmed infrastructures
	TotalNonConfirmedInfrastructures int `json:"totalNonConfirmedInfrastructures"`
}

type GetProbeDetails struct {
	// Name of the probe
	ProbeName string `json:"probeName"`
	// Type of Probe
	ProbeType ProbeType `json:"probeType"`
	// Infra type of Probe
	InfraType InfrastructureType `json:"infraType"`
	// Description of Probe
	Description *string `json:"description"`
	// Tags of probe
	Tags []string `json:"tags"`
	// YAML spec of probe
	ProbeSpec *string `json:"probeSpec"`
}

// Defines the response of the Probe reference API
type GetProbeReferenceResponse struct {
	// Harness identifiers
	Identifiers *Identifiers `json:"identifiers"`
	// ID of the Probe
	ProbeID string `json:"probeID"`
	// Total Runs
	TotalRuns int `json:"totalRuns"`
	// Recent Executions of the probe
	RecentExecutions []*RecentExecutions `json:"recentExecutions"`
}

// Defines the input requests for GetProbeYAML query
type GetProbeYAMLRequest struct {
	// Probe name of the probe
	ProbeID string `json:"probeID"`
	// Mode of the Probe (SoT, EoT, Edge, Continuous or OnChaos)
	Mode Mode `json:"mode"`
}

// Defines the response for Get Probe In Experiment Run Query
type GetProbesInExperimentRunResponse struct {
	// Probe Object
	Probe *Probe `json:"probe"`
	// Mode of the probe
	Mode Mode `json:"mode"`
	// Status of the Probe
	Status *Status `json:"status"`
}

// Auth contains the authentication details for the HTTP probe
type HTTPAuthorization struct {
	// Flag to hold the authentication type
	AuthType *string `json:"authType"`
	// Flag to hold the basic auth credentials in base64 format
	Credentials *string `json:"credentials"`
	// Flag to hold the basic auth credentials file path
	CredentialsFile *string `json:"credentialsFile"`
}

// Defines the input for HTTP authentication details
type HTTPAuthorizationInput struct {
	// Flag to hold the authentication type
	AuthType *string `json:"authType"`
	// Flag to hold the basic auth credentials in base64 format
	Credentials *string `json:"credentials"`
	// Flag to hold the basic auth credentials file path
	CredentialsFile *string `json:"credentialsFile"`
}

// Defines the structure for HTTP Headers
type Headers struct {
	// Key of the header
	Key string `json:"key"`
	// Value of the header
	Value string `json:"value"`
}

// Defines the input for HTTP Headers
type HeadersRequest struct {
	// Key of the header
	Key string `json:"key"`
	// Value of the header
	Value string `json:"value"`
}

// Defines the common identifiers for API operations
type Identifiers struct {
	// Harness OrgID
	OrgIdentifier string `json:"orgIdentifier"`
	// Harness AccountID
	AccountIdentifier string `json:"accountIdentifier"`
	// Harness projectID
	ProjectIdentifier string `json:"projectIdentifier"`
}

type IdentifiersRequest struct {
	// Harness OrgID
	OrgIdentifier string `json:"orgIdentifier"`
	// Harness AccountID
	AccountIdentifier string `json:"accountIdentifier"`
	// Harness projectID
	ProjectIdentifier string `json:"projectIdentifier"`
}

type ImageRegistryRequest struct {
	RegistryServer    string        `json:"registryServer"`
	RegistryAccount   string        `json:"registryAccount"`
	InfraID           *string       `json:"infraID"`
	IsPrivate         bool          `json:"isPrivate"`
	SecretName        *string       `json:"secretName"`
	IsDefault         bool          `json:"isDefault"`
	IsOverrideAllowed bool          `json:"isOverrideAllowed"`
	UseCustomImages   bool          `json:"useCustomImages"`
	CustomImages      *CustomImages `json:"customImages"`
}

type ImageRegistryResponse struct {
	Identifier        *ScopedIdentifiers   `json:"identifier"`
	InfraID           *string              `json:"infraID"`
	RegistryServer    string               `json:"registryServer"`
	RegistryAccount   string               `json:"registryAccount"`
	IsOverrideAllowed bool                 `json:"isOverrideAllowed"`
	IsPrivate         bool                 `json:"isPrivate"`
	SecretName        *string              `json:"secretName"`
	IsDefault         bool                 `json:"isDefault"`
	CreatedBy         *UserDetails         `json:"createdBy"`
	UpdatedBy         *UserDetails         `json:"updatedBy"`
	CreatedAt         string               `json:"createdAt"`
	UpdatedAt         string               `json:"updatedAt"`
	UseCustomImages   bool                 `json:"useCustomImages"`
	CustomImages      *CustomImagesRequest `json:"customImages"`
}

func (ImageRegistryResponse) IsAudit() {}

// Import Probes from ChaosHub request
type ImportProbeRequest struct {
	// Hub ID for reference
	HubID string `json:"hubID"`
	// Probe Name from hub to be imported
	ProbeName string `json:"probeName"`
}

type InfraActionResponse struct {
	Action *ActionPayload `json:"action"`
}

type InfraEventResponse struct {
	EventID     string           `json:"eventID"`
	EventType   string           `json:"eventType"`
	EventName   string           `json:"eventName"`
	Description string           `json:"description"`
	Infra       *KubernetesInfra `json:"infra"`
}

// Defines filter options for infras
type InfraFilterInput struct {
	// Name of the infra
	Name *string `json:"name"`
	// ID of the infra
	InfraID *string `json:"infraID"`
	// ID of the infra
	Description *string `json:"description"`
	// Platform name of infra
	PlatformName *string `json:"platformName"`
	// Scope of infra
	InfraScope *InfraScope `json:"infraScope"`
	// Status of infra
	IsActive *bool `json:"isActive"`
	// Tags of an infra
	Tags []*string `json:"tags"`
	// To filter out infras that are compatible with the new experiment manifest format: without experiment CRs
	CompatibleWithNewExp *bool `json:"compatibleWithNewExp"`
}

type InfraIdentity struct {
	InfraID   string `json:"infraID"`
	AccessKey string `json:"accessKey"`
	Version   string `json:"version"`
}

type InfraSpec struct {
	Operator Operator `json:"operator"`
	InfraIds []string `json:"infraIds"`
}

type InfraSpecInput struct {
	InfraIds []string `json:"infraIds"`
	Operator Operator `json:"operator"`
}

// InfraVersionDetails returns the details of compatible infra versions and the latest infra version supported
type InfraVersionDetails struct {
	// Latest infra version supported
	LatestVersion string `json:"latestVersion"`
	// List of all infra versions supported
	CompatibleVersions []string `json:"compatibleVersions"`
}

// InfraWithExperimentStats defines experiment stats for a given infra
type InfraWithExperimentStats struct {
	// ID of the chaos infrastructure
	InfraID string `json:"infraID"`
	// Experiments count
	ExperimentsCount int `json:"experimentsCount"`
	// Experiment Runs count
	ExperimentRunsCount int `json:"experimentRunsCount"`
}

type Infrastructure struct {
	// Harness identifiers
	Identifiers *Identifiers `json:"identifiers"`
	// Environment ID where infra is installed
	EnvironmentID string `json:"environmentID"`
	// Type of infrastructure
	InfraType InfrastructureType `json:"infraType"`
	// ID of the infra
	InfraID string `json:"infraID"`
	// Name of the workflow
	Name string `json:"name"`
	// Description of the workflow
	Description *string `json:"description"`
	// Version of infrastructure
	Version string `json:"version"`
	// Timestamp when the workflow was last updated
	UpdatedAt string `json:"updatedAt"`
	// Timestamp when the workflow was created
	CreatedAt string `json:"createdAt"`
	// Bool value indicating if the workflow has removed
	IsRemoved bool `json:"isRemoved"`
	// Tags of the workflow
	Tags []string `json:"tags"`
	// User who created the workflow
	CreatedBy *UserDetails `json:"createdBy"`
	// Details of the user who updated the workflow
	UpdatedBy *UserDetails `json:"updatedBy"`
	// Last received heartbeat of infrastructure
	LastHeartbeat string `json:"lastHeartbeat"`
	// Time when infra became active
	StartTime string `json:"startTime"`
	// Bool value to check if infra is active
	IsActive bool `json:"isActive"`
	// Bool value to check if infra is confirmed
	IsInfraConfirmed bool `json:"isInfraConfirmed"`
	// Namespace where the infra is being installed
	InfraNamespace *string `json:"infraNamespace"`
	// Cluster type Indicates the type on infrastructure (Kubernetes/openshift)
	ClusterType *ClusterType `json:"clusterType"`
	// Scope of the infra : ns or cluster
	InfraScope *InfraScope `json:"infraScope"`
}

func (Infrastructure) IsResourceDetails() {}
func (Infrastructure) IsAudit()           {}

// Defines the K8S probe properties
type K8SProbe struct {
	// Timeout of the Probe
	ProbeTimeout string `json:"probeTimeout"`
	// Interval of the Probe
	Interval string `json:"interval"`
	// Retry interval of the Probe
	Retry *int `json:"retry"`
	// Attempt contains the total attempt count for the probe
	Attempt int `json:"attempt"`
	// Polling interval of the Probe
	ProbePollingInterval *string `json:"probePollingInterval"`
	// Initial delay interval of the Probe in seconds
	InitialDelay *string `json:"initialDelay"`
	// Is stop on failure enabled in the Probe
	StopOnFailure *bool `json:"stopOnFailure"`
	// Group of the Probe
	Group *string `json:"group"`
	// Version of the Probe
	Version string `json:"version"`
	// Resource of the Probe
	Resource string `json:"resource"`
	// Namespace of the Probe
	Namespace *string `json:"namespace"`
	// Resource Names of the Probe
	ResourceNames *string `json:"resourceNames"`
	// Field Selector of the Probe
	FieldSelector *string `json:"fieldSelector"`
	// Label Selector of the Probe
	LabelSelector *string `json:"labelSelector"`
	// Operation of the Probe
	Operation string `json:"operation"`
	// Verbosity for the probe logging
	Verbosity *string `json:"verbosity"`
}

func (K8SProbe) IsCommonProbeProperties() {}

// Defines the input for K8S probe properties
type K8SProbeRequest struct {
	// Timeout of the Probe
	ProbeTimeout string `json:"probeTimeout"`
	// Interval of the Probe
	Interval string `json:"interval"`
	// Retry interval of the Probe
	Retry *int `json:"retry"`
	// Attempt contains the total attempt count for the probe
	Attempt int `json:"attempt"`
	// Polling interval of the Probe
	ProbePollingInterval *string `json:"probePollingInterval"`
	// Initial delay interval of the Probe in seconds
	InitialDelay *string `json:"initialDelay"`
	// Is stop on failure enabled in the Probe
	StopOnFailure *bool `json:"stopOnFailure"`
	// Group of the Probe
	Group *string `json:"group"`
	// Version of the Probe
	Version string `json:"version"`
	// Resource of the Probe
	Resource string `json:"resource"`
	// Namespace of the Probe
	Namespace *string `json:"namespace"`
	// Resource Names of the Probe
	ResourceNames *string `json:"resourceNames"`
	// Field Selector of the Probe
	FieldSelector *string `json:"fieldSelector"`
	// Label Selector of the Probe
	LabelSelector *string `json:"labelSelector"`
	// Operation of the Probe
	Operation string `json:"operation"`
	// Verbosity for the probe logging
	Verbosity *string `json:"verbosity"`
}

type K8sSpec struct {
	InfraSpec               *InfraSpec               `json:"infraSpec"`
	ApplicationSpec         *ApplicationSpec         `json:"applicationSpec"`
	ChaosServiceAccountSpec *ChaosServiceAccountSpec `json:"chaosServiceAccountSpec"`
}

type K8sSpecInput struct {
	InfraSpec               *InfraSpecInput               `json:"infraSpec"`
	ApplicationSpec         *ApplicationSpecInput         `json:"applicationSpec"`
	ChaosServiceAccountSpec *ChaosServiceAccountSpecInput `json:"chaosServiceAccountSpec"`
}

type KubeGVRRequest struct {
	Group    string `json:"group"`
	Version  string `json:"version"`
	Resource string `json:"resource"`
}

// KubeObject consists of the namespace and the available resources in the same
type KubeObject struct {
	// Namespace of the resource
	Namespace string `json:"namespace"`
	// Details of the resource
	Data []*ObjectData `json:"data"`
}

// Defines details for fetching Kubernetes object data
type KubeObjectRequest struct {
	// ID of the infra in which the Kubernetes object is present
	RequestID string `json:"requestID"`
	// ID of the infra in which the Kubernetes object is present
	InfraID string `json:"infraID"`
	// GVR Request
	KubeObjRequest *KubeGVRRequest `json:"kubeObjRequest"`
	// Request Namespace
	Namespace *string `json:"namespace"`
}

// Response received for querying Kubernetes Object
type KubeObjectResponse struct {
	// ID of the infra in which the Kubernetes object is present
	InfraID string `json:"infraID"`
	// Type of the Kubernetes object
	KubeObj []*KubeObject `json:"kubeObj"`
}

// Defines the Kubernetes CMD probe properties
type KubernetesCMDProbe struct {
	// Timeout of the Probe
	ProbeTimeout string `json:"probeTimeout"`
	// Interval of the Probe
	Interval string `json:"interval"`
	// Retry interval of the Probe
	Retry *int `json:"retry"`
	// Attempt contains the total attempt count for the probe
	Attempt int `json:"attempt"`
	// Polling interval of the Probe
	ProbePollingInterval *string `json:"probePollingInterval"`
	// Initial delay interval of the Probe in seconds
	InitialDelay *string `json:"initialDelay"`
	// Is stop on failure enabled in the Probe
	StopOnFailure *bool `json:"stopOnFailure"`
	// Command of the Probe
	Command string `json:"command"`
	// Comparator of the Probe
	Comparator *Comparator `json:"comparator"`
	// Source of the Probe
	Source *string `json:"source"`
	// Verbosity for the probe logging
	Verbosity *string `json:"verbosity"`
	// ENV for the probe
	Env []*Env `json:"env"`
}

func (KubernetesCMDProbe) IsCommonProbeProperties() {}

// Defines the input for Kubernetes CMD probe properties
type KubernetesCMDProbeRequest struct {
	// Timeout of the Probe
	ProbeTimeout string `json:"probeTimeout"`
	// Interval of the Probe
	Interval string `json:"interval"`
	// Retry interval of the Probe
	Retry *int `json:"retry"`
	// Attempt contains the total attempt count for the probe
	Attempt int `json:"attempt"`
	// Polling interval of the Probe
	ProbePollingInterval *string `json:"probePollingInterval"`
	// Initial delay interval of the Probe in seconds
	InitialDelay *string `json:"initialDelay"`
	// Is stop on failure enabled in the Probe
	StopOnFailure *bool `json:"stopOnFailure"`
	// Command of the Probe
	Command string `json:"command"`
	// Comparator of the Probe
	Comparator *ComparatorInput `json:"comparator"`
	// Source of the Probe
	Source *string `json:"source"`
	// Verbosity for the probe logging
	Verbosity *string `json:"verbosity"`
}

// Defines the Kubernetes Datadog probe properties
type KubernetesDatadogProbe struct {
	// Timeout of the Probe
	ProbeTimeout string `json:"probeTimeout"`
	// Interval of the Probe
	Interval string `json:"interval"`
	// Retry interval of the Probe
	Retry *int `json:"retry"`
	// Attempt contains the total attempt count for the probe
	Attempt int `json:"attempt"`
	// Polling interval of the Probe
	ProbePollingInterval *string `json:"probePollingInterval"`
	// Initial delay interval of the Probe in seconds
	InitialDelay *string `json:"initialDelay"`
	// Is stop on failure enabled in the Probe
	StopOnFailure *bool `json:"stopOnFailure"`
	// Datadog site URL identifier
	DatadogSite string `json:"datadogSite"`
	// Synthetics test parameters
	SyntheticsTest *SyntheticsTest `json:"syntheticsTest"`
	// Metrics parameters
	Metrics *DatadogMetrics `json:"metrics"`
	// Name of the kubernetes secret containing the Datadog credentials
	DatadogCredentialsSecretName string `json:"datadogCredentialsSecretName"`
	// Verbosity for the probe logging
	Verbosity *string `json:"verbosity"`
}

func (KubernetesDatadogProbe) IsCommonProbeProperties() {}

// Defines the input for Kubernetes Datadog probe
type KubernetesDatadogProbeRequest struct {
	// Timeout of the Probe
	ProbeTimeout string `json:"probeTimeout"`
	// Interval of the Probe
	Interval string `json:"interval"`
	// Retry interval of the Probe
	Retry *int `json:"retry"`
	// Attempt contains the total attempt count for the probe
	Attempt int `json:"attempt"`
	// Polling interval of the Probe
	ProbePollingInterval *string `json:"probePollingInterval"`
	// Initial delay interval of the Probe in seconds
	InitialDelay *string `json:"initialDelay"`
	// Is stop on failure enabled in the Probe
	StopOnFailure *bool `json:"stopOnFailure"`
	// Datadog site URL identifier
	DatadogSite string `json:"datadogSite"`
	// Synthetics test parameters
	SyntheticsTest *SyntheticsTestRequest `json:"syntheticsTest"`
	// Metrics parameters
	Metrics *DatadogMetricsInput `json:"metrics"`
	// Name of the kubernetes secret containing the Datadog credentials
	DatadogCredentialsSecretName string `json:"datadogCredentialsSecretName"`
	// Verbosity for the probe logging
	Verbosity *string `json:"verbosity"`
}

// Defines the Kubernetes Dynatrace probe properties
type KubernetesDynatraceProbe struct {
	// Timeout of the Probe
	ProbeTimeout string `json:"probeTimeout"`
	// Interval of the Probe
	Interval string `json:"interval"`
	// Retry interval of the Probe
	Retry *int `json:"retry"`
	// Attempt contains the total attempt count for the probe
	Attempt int `json:"attempt"`
	// Polling interval of the Probe
	ProbePollingInterval *string `json:"probePollingInterval"`
	// Initial delay interval of the Probe in seconds
	InitialDelay *string `json:"initialDelay"`
	// Is stop on failure enabled in the Probe
	StopOnFailure *bool `json:"stopOnFailure"`
	// Endpoint of the dynatrace probe
	Endpoint string `json:"endpoint"`
	// Raw metrics details of the dynatrace probe
	Metrics *Metrics `json:"metrics"`
	// Timeframe of the metrics
	TimeFrame string `json:"timeFrame"`
	// APITokenSecretName for authenticating with the Dynatrace platform
	APITokenSecretName string `json:"apiTokenSecretName"`
	// Comparator check for the correctness of the probe output
	Comparator *Comparator `json:"comparator"`
	// Verbosity for the probe logging
	Verbosity *string `json:"verbosity"`
}

func (KubernetesDynatraceProbe) IsCommonProbeProperties() {}

// Defines the input for Kubernetes Dynatrace probe properties
type KubernetesDynatraceProbeRequest struct {
	// Timeout of the Probe
	ProbeTimeout string `json:"probeTimeout"`
	// Interval of the Probe
	Interval string `json:"interval"`
	// Retry interval of the Probe
	Retry *int `json:"retry"`
	// Attempt contains the total attempt count for the probe
	Attempt int `json:"attempt"`
	// Polling interval of the Probe
	ProbePollingInterval *string `json:"probePollingInterval"`
	// Initial delay interval of the Probe in seconds
	InitialDelay *string `json:"initialDelay"`
	// Is stop on failure enabled in the Probe
	StopOnFailure *bool `json:"stopOnFailure"`
	// Endpoint of the dynatrace probe
	Endpoint string `json:"endpoint"`
	// Raw metrics details of the dynatrace probe
	Metrics *MetricsInput `json:"metrics"`
	// Timeframe of the metrics
	TimeFrame string `json:"timeFrame"`
	// APITokenSecretName for authenticating with the Dynatrace platform
	APITokenSecretName string `json:"apiTokenSecretName"`
	// Comparator check for the correctness of the probe output
	Comparator *ComparatorInput `json:"comparator"`
	// Verbosity for the probe logging
	Verbosity *string `json:"verbosity"`
}

// Defines the Kubernetes HTTP probe properties
type KubernetesHTTPProbe struct {
	// Timeout of the Probe
	ProbeTimeout string `json:"probeTimeout"`
	// Interval of the Probe
	Interval string `json:"interval"`
	// Retry interval of the Probe
	Retry *int `json:"retry"`
	// Attempt contains the total attempt count for the probe
	Attempt int `json:"attempt"`
	// Polling interval of the Probe
	ProbePollingInterval *string `json:"probePollingInterval"`
	// Initial delay interval of the Probe in seconds
	InitialDelay *string `json:"initialDelay"`
	// Is stop on failure enabled in the Probe
	StopOnFailure *bool `json:"stopOnFailure"`
	// URL of the Probe
	URL string `json:"url"`
	// HTTP method of the Probe
	Method *Method `json:"method"`
	// Verbosity for the probe logging
	Verbosity *string `json:"verbosity"`
	// Auth contains the authentication details for the prometheus probe
	Auth *HTTPAuthorization `json:"auth"`
	// TLSConfig contains the tls configuration for the prometheus probe
	TLSConfig *TLSConfig `json:"tlsConfig"`
	// Headers contains the request headers
	Headers []*Headers `json:"headers"`
}

func (KubernetesHTTPProbe) IsCommonProbeProperties() {}

// Defines the input for Kubernetes HTTP probe properties
type KubernetesHTTPProbeRequest struct {
	// Timeout of the Probe
	ProbeTimeout string `json:"probeTimeout"`
	// Interval of the Probe
	Interval string `json:"interval"`
	// Retry interval of the Probe
	Retry *int `json:"retry"`
	// Attempt contains the total attempt count for the probe
	Attempt int `json:"attempt"`
	// Polling interval of the Probe
	ProbePollingInterval *string `json:"probePollingInterval"`
	// Initial delay interval of the Probe in seconds
	InitialDelay *string `json:"initialDelay"`
	// Is stop on failure enabled in the Probe
	StopOnFailure *bool `json:"stopOnFailure"`
	// URL of the Probe
	URL string `json:"url"`
	// HTTP method of the Probe
	Method *MethodRequest `json:"method"`
	// Verbosity for the probe logging
	Verbosity *string `json:"verbosity"`
	// Auth contains the authentication details for the HTTP probe
	Auth *HTTPAuthorizationInput `json:"auth"`
	// TLSConfig contains the tls configuration for the prometheus probe
	TLSConfig *TLSConfigInput `json:"tlsConfig"`
	// Headers contains the request headers
	Headers []*HeadersRequest `json:"headers"`
}

// Defines the details for a infra
type KubernetesInfra struct {
	// ID of the infra
	InfraID string `json:"infraID"`
	// Name of the infra
	Name string `json:"name"`
	// Description of the infra
	Description *string `json:"description"`
	// Tags of the infra
	Tags []string `json:"tags"`
	// Environment ID for the infra
	EnvironmentID string `json:"environmentID"`
	// Infra Platform Name eg. GKE,AWS, Others
	PlatformName string `json:"platformName"`
	// Boolean value indicating if chaos infrastructure is active or not
	IsActive bool `json:"isActive"`
	// Boolean value indicating if chaos infrastructure is confirmed or not
	IsInfraConfirmed bool `json:"isInfraConfirmed"`
	// Boolean value indicating if chaos infrastructure is removed or not
	IsRemoved bool `json:"isRemoved"`
	// Timestamp when the infra was last updated
	UpdatedAt string `json:"updatedAt"`
	// Timestamp when the infra was created
	CreatedAt string `json:"createdAt"`
	// Number of schedules created in the infra
	NoOfSchedules *int `json:"noOfSchedules"`
	// Number of workflows run in the infra
	NoOfWorkflows *int `json:"noOfWorkflows"`
	// Token used to verify and retrieve the infra manifest
	Token string `json:"token"`
	// Namespace where the infra is being installed
	InfraNamespace *string `json:"infraNamespace"`
	// Name of service account used by infra
	ServiceAccount *string `json:"serviceAccount"`
	// Scope of the infra : ns or cluster
	InfraScope InfraScope `json:"infraScope"`
	// Bool value indicating whether infra ns used already exists on infra or not
	InfraNsExists *bool `json:"infraNsExists"`
	// Bool value indicating whether service account used already exists on infra or not
	InfraSaExists *bool `json:"infraSaExists"`
	// InstallationType connector/manifest
	InstallationType InstallationType `json:"installationType"`
	// K8sConnectorID
	K8sConnectorID *string `json:"k8sConnectorID"`
	// Timestamp of the last workflow run in the infra
	LastWorkflowTimestamp *string `json:"lastWorkflowTimestamp"`
	// Timestamp when the infra got connected
	StartTime string `json:"startTime"`
	// Version of the infra
	Version string `json:"version"`
	// User who created the infra
	CreatedBy *UserDetails `json:"createdBy"`
	// User who has updated the infra
	UpdatedBy *UserDetails `json:"updatedBy"`
	// Last Heartbeat status sent by the infra
	LastHeartbeat *string `json:"lastHeartbeat"`
	// Type of the infrastructure
	InfraType *InfrastructureType `json:"infraType"`
	// update status of infra
	UpdateStatus UpdateStatus `json:"updateStatus"`
	// Tune secret for infra
	IsSecretEnabled *bool `json:"isSecretEnabled"`
	// set the user for security context in pod
	RunAsUser *int `json:"runAsUser"`
	// set the user group for security context in pod
	RunAsGroup *int `json:"runAsGroup"`
	// Upgrade struct for the chaos infrastructure
	Upgrade *Upgrade `json:"upgrade"`
	// Cluster type Indicates the type on infrastructure (Kubernetes/openshift)
	ClusterType ClusterType `json:"clusterType"`
}

func (KubernetesInfra) IsResourceDetails() {}
func (KubernetesInfra) IsAudit()           {}

type Link struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Defines the Linux CMD probe properties
type LinuxCMDProbe struct {
	// Timeout of the Probe
	ProbeTimeout string `json:"probeTimeout"`
	// Interval of the Probe
	Interval string `json:"interval"`
	// Attempt contains the total attempt count for the probe
	Attempt int `json:"attempt"`
	// Polling interval of the Probe
	ProbePollingInterval *string `json:"probePollingInterval"`
	// Initial delay interval of the Probe in seconds
	InitialDelay *string `json:"initialDelay"`
	// Is stop on failure enabled in the Probe
	StopOnFailure *bool `json:"stopOnFailure"`
	// Command of the Probe
	Command string `json:"command"`
	// Comparator of the Probe
	Comparator *Comparator `json:"comparator"`
	// Source of the Probe
	Source *string `json:"source"`
}

// Defines the input for Linux CMD probe properties
type LinuxCMDProbeRequest struct {
	// Timeout of the Probe
	ProbeTimeout string `json:"probeTimeout"`
	// Interval of the Probe
	Interval string `json:"interval"`
	// Attempt contains the total attempt count for the probe
	Attempt int `json:"attempt"`
	// Polling interval of the Probe
	ProbePollingInterval *string `json:"probePollingInterval"`
	// Initial delay interval of the Probe in seconds
	InitialDelay *string `json:"initialDelay"`
	// Is stop on failure enabled in the Probe
	StopOnFailure *bool `json:"stopOnFailure"`
	// Command of the Probe
	Command string `json:"command"`
	// Comparator of the Probe
	Comparator *ComparatorInput `json:"comparator"`
	// Source of the Probe
	Source *string `json:"source"`
}

// Defines the Linux Datadog probe properties
type LinuxDatadogProbe struct {
	// Timeout of the Probe
	ProbeTimeout string `json:"probeTimeout"`
	// Interval of the Probe
	Interval string `json:"interval"`
	// Attempt contains the total attempt count for the probe
	Attempt int `json:"attempt"`
	// Polling interval of the Probe
	ProbePollingInterval *string `json:"probePollingInterval"`
	// Initial delay interval of the Probe in seconds
	InitialDelay *string `json:"initialDelay"`
	// Is stop on failure enabled in the Probe
	StopOnFailure *bool `json:"stopOnFailure"`
	// Datadog site URL identifier
	DatadogSite string `json:"datadogSite"`
	// Synthetics test parameters
	SyntheticsTest *SyntheticsTest `json:"syntheticsTest"`
	// Metrics parameters
	Metrics *DatadogMetrics `json:"metrics"`
}

// Defines the input for Linux Datadog probe
type LinuxDatadogProbeRequest struct {
	// Timeout of the Probe
	ProbeTimeout string `json:"probeTimeout"`
	// Interval of the Probe
	Interval string `json:"interval"`
	// Attempt contains the total attempt count for the probe
	Attempt int `json:"attempt"`
	// Polling interval of the Probe
	ProbePollingInterval *string `json:"probePollingInterval"`
	// Initial delay interval of the Probe in seconds
	InitialDelay *string `json:"initialDelay"`
	// Is stop on failure enabled in the Probe
	StopOnFailure *bool `json:"stopOnFailure"`
	// Datadog site URL identifier
	DatadogSite string `json:"datadogSite"`
	// Synthetics test parameters
	SyntheticsTest *SyntheticsTestRequest `json:"syntheticsTest"`
	// Metrics parameters
	Metrics *DatadogMetricsInput `json:"metrics"`
}

// Defines the Linux Dynatrace probe properties
type LinuxDynatraceProbe struct {
	// Timeout of the Probe
	ProbeTimeout string `json:"probeTimeout"`
	// Interval of the Probe
	Interval string `json:"interval"`
	// Attempt contains the total attempt count for the probe
	Attempt int `json:"attempt"`
	// Polling interval of the Probe
	ProbePollingInterval *string `json:"probePollingInterval"`
	// Initial delay interval of the Probe in seconds
	InitialDelay *string `json:"initialDelay"`
	// Is stop on failure enabled in the Probe
	StopOnFailure *bool `json:"stopOnFailure"`
	// Endpoint of the dynatrace probe
	Endpoint string `json:"endpoint"`
	// Raw metrics details of the dynatrace probe
	Metrics *Metrics `json:"metrics"`
	// Timeframe of the metrics
	TimeFrame string `json:"timeFrame"`
	// Comparator check for the correctness of the probe output
	Comparator *Comparator `json:"comparator"`
}

// Defines the input for Linux Dynatrace probe properties
type LinuxDynatraceProbeRequest struct {
	// Timeout of the Probe
	ProbeTimeout string `json:"probeTimeout"`
	// Interval of the Probe
	Interval string `json:"interval"`
	// Attempt contains the total attempt count for the probe
	Attempt int `json:"attempt"`
	// Polling interval of the Probe
	ProbePollingInterval *string `json:"probePollingInterval"`
	// Initial delay interval of the Probe in seconds
	InitialDelay *string `json:"initialDelay"`
	// Is stop on failure enabled in the Probe
	StopOnFailure *bool `json:"stopOnFailure"`
	// Endpoint of the dynatrace probe
	Endpoint string `json:"endpoint"`
	// Raw metrics details of the dynatrace probe
	Metrics *MetricsInput `json:"metrics"`
	// Timeframe of the metrics
	TimeFrame string `json:"timeFrame"`
	// Comparator check for the correctness of the probe output
	Comparator *ComparatorInput `json:"comparator"`
}

// Defines the Linux HTTP probe properties
type LinuxHTTPProbe struct {
	// Timeout of the Probe
	ProbeTimeout string `json:"probeTimeout"`
	// Interval of the Probe
	Interval string `json:"interval"`
	// Attempt contains the total attempt count for the probe
	Attempt int `json:"attempt"`
	// Polling interval of the Probe
	ProbePollingInterval *string `json:"probePollingInterval"`
	// Initial delay interval of the Probe in seconds
	InitialDelay *string `json:"initialDelay"`
	// Is stop on failure enabled in the Probe
	StopOnFailure *bool `json:"stopOnFailure"`
	// URL of the Probe
	URL string `json:"url"`
	// HTTP method of the Probe
	Method *Method `json:"method"`
	// If Insecure HTTP verification should  be skipped
	InsecureSkipVerify *bool `json:"insecureSkipVerify"`
	// Auth contains the authentication details for the prometheus probe
	Auth *HTTPAuthorization `json:"auth"`
	// TLSConfig contains the tls configuration for the prometheus probe
	TLSConfig *TLSConfig `json:"tlsConfig"`
	// Headers contains the request headers
	Headers []*Headers `json:"headers"`
}

// Defines the input for Linux HTTP probe properties
type LinuxHTTPProbeRequest struct {
	// Timeout of the Probe
	ProbeTimeout string `json:"probeTimeout"`
	// Interval of the Probe
	Interval string `json:"interval"`
	// Attempt contains the total attempt count for the probe
	Attempt int `json:"attempt"`
	// Polling interval of the Probe
	ProbePollingInterval *string `json:"probePollingInterval"`
	// Initial delay interval of the Probe in seconds
	InitialDelay *string `json:"initialDelay"`
	// Is stop on failure enabled in the Probe
	StopOnFailure *bool `json:"stopOnFailure"`
	// URL of the Probe
	URL string `json:"url"`
	// HTTP method of the Probe
	Method *MethodRequest `json:"method"`
	// Auth contains the authentication details for the HTTP probe
	Auth *HTTPAuthorizationInput `json:"auth"`
	// TLSConfig contains the tls configuration for the prometheus probe
	TLSConfig *TLSConfigInput `json:"tlsConfig"`
	// Headers contains the request headers
	Headers []*HeadersRequest `json:"headers"`
}

// Defines the details for a infra
type LinuxInfra struct {
	// ID of the infra
	InfraID string `json:"infraID"`
	// Name of the infra
	Name string `json:"name"`
	// Description of the infra
	Description *string `json:"description"`
	// Tags of the infra
	Tags []string `json:"tags"`
	// Environment ID for the infra
	EnvironmentID string `json:"environmentID"`
	// Boolean value indicating if chaos infrastructure is active or not
	IsActive bool `json:"isActive"`
	// Boolean value indicating if chaos infrastructure is confirmed or not
	IsInfraConfirmed bool `json:"isInfraConfirmed"`
	// Boolean value indicating if chaos infrastructure is removed or not
	IsRemoved bool `json:"isRemoved"`
	// Timestamp when the infra was last updated
	UpdatedAt string `json:"updatedAt"`
	// Timestamp when the infra was created
	CreatedAt string `json:"createdAt"`
	// Number of schedules created in the infra
	NoOfSchedules *int `json:"noOfSchedules"`
	// Number of workflows run in the infra
	NoOfWorkflows *int `json:"noOfWorkflows"`
	// Timestamp of the last workflow run in the infra
	LastWorkflowTimestamp *string `json:"lastWorkflowTimestamp"`
	// Timestamp when the infra got connected
	StartTime string `json:"startTime"`
	// Version of the infra
	Version string `json:"version"`
	// User who created the infra
	CreatedBy *UserDetails `json:"createdBy"`
	// User who has updated the infra
	UpdatedBy *UserDetails `json:"updatedBy"`
	// Last Heartbeat status sent by the infra
	LastHeartbeat *string `json:"lastHeartbeat"`
	// hostname of the infra
	Hostname *string `json:"hostname"`
}

func (LinuxInfra) IsResourceDetails() {}
func (LinuxInfra) IsAudit()           {}

// Defines filter options for infras
type LinuxInfraFilterInput struct {
	// Name of the infra
	Name *string `json:"name"`
	// ID of the infra
	InfraID *string `json:"infraID"`
	// ID of the infra
	Description *string `json:"description"`
	// Status of infra
	IsActive *bool `json:"isActive"`
	// Tags of an infra
	Tags []*string `json:"tags"`
}

type ListChaosHubRequest struct {
	// Array of ChaosHub IDs for which details will be fetched
	ChaosHubIDs []string `json:"chaosHubIDs"`
	// Details for fetching filtered data
	Filter *ChaosHubFilterInput `json:"filter"`
}

// Defines the details for a workflow
type ListCloudFoundryInfraRequest struct {
	// Array of infra IDs for which details will be fetched
	InfraIDs []string `json:"infraIDs"`
	// Environment ID
	EnvironmentIDs []string `json:"environmentIDs"`
	// Details for fetching paginated data
	Pagination *Pagination `json:"pagination"`
	// Details for fetching filtered data
	Filter *CloudFoundryInfraFilterInput `json:"filter"`
}

// Defines the details for a infras with total infras count
type ListCloudFoundryInfraResponse struct {
	// Total number of infras
	TotalNoOfInfras int `json:"totalNoOfInfras"`
	// Details related to the infras
	Infras []*CloudFoundryInfra `json:"infras"`
}

type ListConditionRequest struct {
	Pagination *Pagination           `json:"pagination"`
	Filter     *ConditionFilterInput `json:"filter"`
}

type ListConditionResponse struct {
	TotalConditions int                  `json:"totalConditions"`
	Conditions      []*ConditionResponse `json:"conditions"`
}

type ListExperimentsWithActiveInfrasFilterInput struct {
	// Name of the workflow
	ExperimentName *string `json:"experimentName"`
	// Name of the infra in which the experiment is scheduled
	InfraName *string `json:"infraName"`
	// ID of the infrastructure in which the experiment is scheduled
	InfraIDs []*string `json:"infraIDs"`
	// Type of the experiment i.e. CRON, NON_CRON or Gameday
	ExperimentType *ScenarioType `json:"experimentType"`
	// Date range for filtering purpose
	DateRange *DateRange `json:"dateRange"`
	// Type of infrastructure
	InfraTypes []*InfrastructureType `json:"infraTypes"`
	// Tags based search
	Tags []*string `json:"tags"`
}

// Defines the details for workflow runs
type ListExperimentsWithActiveInfrasRequest struct {
	// Details for fetching paginated data
	Pagination *Pagination `json:"pagination"`
	// Details for fetching sorted data
	Sort *WorkflowRunSortInput `json:"sort"`
	// Details for fetching filtered data
	Filter *ListExperimentsWithActiveInfrasFilterInput `json:"filter"`
}

// Defines the details for a workflow with total workflow count
type ListExperimentsWithActiveInfrasResponse struct {
	// Total number of workflows
	TotalNoOfExperiments int `json:"totalNoOfExperiments"`
	// Details related to the workflows
	Experiments []*Workflow `json:"experiments"`
}

type ListGameDayRequest struct {
	GameDayIDs []string `json:"gameDayIDs"`
	// Details for fetching paginated data
	Pagination *Pagination `json:"pagination"`
	// Details for fetching filtered data
	Filter *GamedayFilterInput `json:"filter"`
	// Details for fetching sorted data
	Sort *GamedaySortInput `json:"sort"`
}

type ListGameDayRunsRequest struct {
	GameDayRunIDs []string `json:"gameDayRunIDs"`
	GameDayIDs    []string `json:"gameDayIDs"`
}

type ListGamedayResponse struct {
	// Total number of workflows
	TotalNoOfGamedays int `json:"totalNoOfGamedays"`
	// Details related to the workflows
	Gamedays []*GameDayResponse `json:"gamedays"`
}

// Defines the details for a workflow
type ListInfraRequest struct {
	// Array of infra IDs for which details will be fetched
	InfraIDs []string `json:"infraIDs"`
	// Environment ID
	EnvironmentIDs []string `json:"environmentIDs"`
	// Connector ID
	K8sConnectorIDs []string `json:"k8sConnectorIDs"`
	// Details for fetching paginated data
	Pagination *Pagination `json:"pagination"`
	// Details for fetching filtered data
	Filter *InfraFilterInput `json:"filter"`
}

// Defines the details for a infras with total infras count
type ListInfraResponse struct {
	// Total number of infras
	TotalNoOfInfras int `json:"totalNoOfInfras"`
	// Details related to the infras
	Infras []*KubernetesInfra `json:"infras"`
}

// Defines the details for a workflow
type ListLinuxInfraRequest struct {
	// Array of infra IDs for which details will be fetched
	InfraIDs []string `json:"infraIDs"`
	// Environment ID
	EnvironmentIDs []string `json:"environmentIDs"`
	// Details for fetching paginated data
	Pagination *Pagination `json:"pagination"`
	// Details for fetching filtered data
	Filter *LinuxInfraFilterInput `json:"filter"`
}

// Defines the details for a infras with total infras count
type ListLinuxInfraResponse struct {
	// Total number of infras
	TotalNoOfInfras int `json:"totalNoOfInfras"`
	// Details related to the infras
	Infras []*LinuxInfra `json:"infras"`
}

// Defines the details for a workflow
type ListProbeRequest struct {
	// Type of infrastructure associated with the probe
	InfrastructureType *InfrastructureType `json:"infrastructureType"`
	// ID of probes or array of probe IDs
	ProbeIDs []string `json:"probeIDs"`
	// Flag to either show probe executions or not
	ShowExecutions *bool `json:"showExecutions"`
	// Details for fetching paginated data
	Pagination *Pagination `json:"pagination"`
	// Details for fetching sorted data
	Sort *ProbeSortInput `json:"sort"`
	// Details for fetching filtered data
	Filter *ProbeFilterInput `json:"filter"`
}

// Defines the details for a list probe query with total probe count
type ListProbeResponse struct {
	// Total number of probes
	TotalNoOfProbes int `json:"totalNoOfProbes"`
	// Details related to the probes
	Probes []*Probe `json:"probes"`
}

type ListRuleRequest struct {
	Pagination *Pagination      `json:"pagination"`
	Filter     *RuleFilterInput `json:"filter"`
}

// Defines the details for a workflow
type ListWindowsInfraRequest struct {
	// Array of infra IDs for which details will be fetched
	InfraIDs []string `json:"infraIDs"`
	// Environment ID
	EnvironmentIDs []string `json:"environmentIDs"`
	// Details for fetching paginated data
	Pagination *Pagination `json:"pagination"`
	// Details for fetching filtered data
	Filter *WindowsInfraFilterInput `json:"filter"`
}

// Defines the details for a infras with total infras count
type ListWindowsInfraResponse struct {
	// Total number of infras
	TotalNoOfInfras int `json:"totalNoOfInfras"`
	// Details related to the infras
	Infras []*WindowsInfra `json:"infras"`
}

// Defines the details for a workflow
type ListWorkflowRequest struct {
	// Array of workflow IDs for which details will be fetched
	WorkflowIDs []*string `json:"workflowIDs"`
	// Array of identities for which details will be fetched
	Identities []*string `json:"identities"`
	// Details for fetching paginated data
	Pagination *Pagination `json:"pagination"`
	// Details for fetching sorted data
	Sort *WorkflowSortInput `json:"sort"`
	// Details for fetching filtered data
	Filter *WorkflowFilterInput `json:"filter"`
}

// Defines the details for a workflow with total workflow count
type ListWorkflowResponse struct {
	// Total number of workflows
	TotalNoOfWorkflows int `json:"totalNoOfWorkflows"`
	// Details related to the workflows
	Workflows []*Workflow `json:"workflows"`
}

// Defines the details for workflow runs
type ListWorkflowRunRequest struct {
	// Array of workflow run IDs for which details will be fetched
	WorkflowRunIDs []*string `json:"workflowRunIDs"`
	// Array of notify IDs for which details will be fetched
	NotifyIDs []*string `json:"notifyIDs"`
	// Array of workflow IDs for which details will be fetched
	WorkflowIDs []*string `json:"workflowIDs"`
	// Details for fetching paginated data
	Pagination *Pagination `json:"pagination"`
	// Details for fetching sorted data
	Sort *WorkflowRunSortInput `json:"sort"`
	// Details for fetching filtered data
	Filter *WorkflowRunFilterInput `json:"filter"`
}

// Defines the details of a workflow to sent as response
type ListWorkflowRunResponse struct {
	// Total number of workflow runs
	TotalNoOfWorkflowRuns int `json:"totalNoOfWorkflowRuns"`
	// Defines details of workflow runs
	WorkflowRuns []*WorkflowRun `json:"workflowRuns"`
}

type MachineSpec struct {
	InfraSpec *InfraSpec `json:"infraSpec"`
}

type MachineSpecInput struct {
	InfraSpec *InfraSpecInput `json:"infraSpec"`
}

// Defines the details of the maintainer
type Maintainer struct {
	// Name of the maintainer
	Name string `json:"name"`
	// Email of the maintainer
	Email string `json:"email"`
}

type Metadata struct {
	Name        string      `json:"name"`
	Version     string      `json:"version"`
	Annotations *Annotation `json:"annotations"`
}

// Defines the methods of the probe properties
type Method struct {
	// A GET request
	Get *Get `json:"get"`
	// A POST request
	Post *Post `json:"post"`
}

// Defines the input for methods of the probe properties
type MethodRequest struct {
	// A GET request
	Get *GETRequest `json:"get"`
	// A POST request
	Post *POSTRequest `json:"post"`
}

// Raw metrics details of the dynatrace probe
type Metrics struct {
	// Query to get Dynatrace metrics
	MetricsSelector string `json:"metricsSelector"`
	// Entity Selector of the metrics
	EntitySelector string `json:"entitySelector"`
}

// Defines the input for Raw metrics details of the dynatrace probe
type MetricsInput struct {
	// Query to get Dynatrace metrics
	MetricsSelector string `json:"metricsSelector"`
	// Entity Selector of the metrics
	EntitySelector string `json:"entitySelector"`
}

type NewInfraEventRequest struct {
	EventName   string `json:"eventName"`
	Description string `json:"description"`
	InfraID     string `json:"infraID"`
	AccessKey   string `json:"accessKey"`
}

type NonCronExperiment struct {
	ExperimentID           *string `json:"experimentID"`
	ExperimentTemplateName string  `json:"experimentTemplateName"`
	ChaosInfraID           string  `json:"chaosInfraID"`
	HubID                  string  `json:"hubID"`
}

type ObjectData struct {
	// Labels present in the resource
	Labels []string `json:"labels"`
	// Name of the resource
	Name string `json:"name"`
}

// Details of POST request
type Post struct {
	// Content Type of the request
	ContentType *string `json:"contentType"`
	// Body of the request
	Body *string `json:"body"`
	// Body Path of the HTTP body required for the http post request
	BodyPath *string `json:"bodyPath"`
	// Criteria of the response
	Criteria string `json:"criteria"`
	// Response Code of the response
	ResponseCode *string `json:"responseCode"`
	// Response Body of the response
	ResponseBody *string `json:"responseBody"`
}

// Details for input of the POST request
type POSTRequest struct {
	// Content Type of the request
	ContentType *string `json:"contentType"`
	// Body of the request
	Body *string `json:"body"`
	// Body Path of the request for Body
	BodyPath *string `json:"bodyPath"`
	// Criteria of the response
	Criteria string `json:"criteria"`
	// Response Code of the response
	ResponseCode *string `json:"responseCode"`
	// Response Body of the response
	ResponseBody *string `json:"responseBody"`
}

// Auth contains the Prometheus authentication details
type PROMAuthorization struct {
	// Flag to hold the authentication type
	AuthType *string `json:"authType"`
	// Flag to hold the basic auth credentials in base64 format
	Credentials *string `json:"credentials"`
	// Flag to hold the basic auth credentials file path
	CredentialsFile *string `json:"credentialsFile"`
}

// Defines the input for Prometheus authentication details
type PROMAuthorizationInput struct {
	// Flag to hold the authentication type
	AuthType *string `json:"authType"`
	// Flag to hold the basic auth credentials in base64 format
	Credentials *string `json:"credentials"`
	// Flag to hold the basic auth credentials file path
	CredentialsFile *string `json:"credentialsFile"`
}

// Defines the PROM probe properties
type PROMProbe struct {
	// Timeout of the Probe
	ProbeTimeout string `json:"probeTimeout"`
	// Interval of the Probe
	Interval string `json:"interval"`
	// Retry interval of the Probe
	Retry *int `json:"retry"`
	// Attempt contains the total attempt count for the probe
	Attempt int `json:"attempt"`
	// Polling interval of the Probe
	ProbePollingInterval *string `json:"probePollingInterval"`
	// Initial delay interval of the Probe in seconds
	InitialDelay *string `json:"initialDelay"`
	// Is stop on failure enabled in the Probe
	StopOnFailure *bool `json:"stopOnFailure"`
	// Endpoint of the Probe
	Endpoint string `json:"endpoint"`
	// Query of the Probe
	Query *string `json:"query"`
	// Query path of the Probe
	QueryPath *string `json:"queryPath"`
	// Comparator of the Probe
	Comparator *Comparator `json:"comparator"`
	// Verbosity for the probe logging
	Verbosity *string `json:"verbosity"`
	// Auth contains the authentication details for the prometheus probe
	Auth *PROMAuthorization `json:"auth"`
	// TLSConfig contains the tls configuration for the prometheus probe
	TLSConfig *TLSConfig `json:"tlsConfig"`
}

func (PROMProbe) IsCommonProbeProperties() {}

// Defines the input for PROM probe properties
type PROMProbeRequest struct {
	// Timeout of the Probe
	ProbeTimeout string `json:"probeTimeout"`
	// Interval of the Probe
	Interval string `json:"interval"`
	// Retry interval of the Probe
	Retry *int `json:"retry"`
	// Attempt contains the total attempt count for the probe
	Attempt int `json:"attempt"`
	// Polling interval of the Probe
	ProbePollingInterval *string `json:"probePollingInterval"`
	// Initial delay interval of the Probe in seconds
	InitialDelay *string `json:"initialDelay"`
	// Is stop on failure enabled in the Probe
	StopOnFailure *bool `json:"stopOnFailure"`
	// Endpoint of the Probe
	Endpoint string `json:"endpoint"`
	// Query of the Probe
	Query *string `json:"query"`
	// Query path of the Probe
	QueryPath *string `json:"queryPath"`
	// Comparator of the Probe
	Comparator *ComparatorInput `json:"comparator"`
	// Verbosity for the probe logging
	Verbosity *string `json:"verbosity"`
	// Auth contains the authentication details for the prometheus probe
	Auth *PROMAuthorizationInput `json:"auth"`
	// TLSConfig contains the tls configuration for the prometheus probe
	TLSConfig *TLSConfigInput `json:"tlsConfig"`
}

type PackageInformation struct {
	PackageName string         `json:"packageName"`
	Experiments []*Experiments `json:"experiments"`
}

// Defines data required to fetch paginated data
type Pagination struct {
	// Page number for which data will be fetched
	Page int `json:"page"`
	// Number of data to be fetched
	Limit int `json:"limit"`
}

// Defines the details for fetching the pod logs
type PodLogRequest struct {
	// Unique request
	RequestID string `json:"requestID"`
	// ID of the cluster
	InfraID string `json:"infraID"`
	// ID of a workflow run
	WorkflowRunID string `json:"workflowRunID"`
	// Name of the pod for which logs are required
	PodName string `json:"podName"`
	// Namespace where the pod is running
	PodNamespace string `json:"podNamespace"`
	// Type of the pod: chaosEngine or not pod
	PodType string `json:"podType"`
	// Name of the experiment pod fetched from execution data
	ExpPod *string `json:"expPod"`
	// Name of the runner pod fetched from execution data
	RunnerPod *string `json:"runnerPod"`
	// Namespace where the experiment is executing
	ChaosNamespace *string `json:"chaosNamespace"`
}

// Defines the response received for querying querying the pod logs
type PodLogResponse struct {
	// ID of the workflow run which is to be queried
	WorkflowRunID string `json:"workflowRunID"`
	// Name of the pod for which logs are queried
	PodName string `json:"podName"`
	// Type of the pod: chaosengine
	PodType string `json:"podType"`
	// Logs for the pod
	Log string `json:"log"`
}

type PredefinedWorkflowList struct {
	WorkflowName     string `json:"workflowName"`
	WorkflowCsv      string `json:"workflowCSV"`
	WorkflowManifest string `json:"workflowManifest"`
	Error            string `json:"error"`
}

// Defines the details of the Probe entity
type Probe struct {
	// ID of the probe
	ProbeID string `json:"probeID"`
	// Harness identifiers
	Identifiers *Identifiers `json:"identifiers"`
	// Name of the Probe
	Name string `json:"name"`
	// Description of the Probe
	Description *string `json:"description"`
	// Tags of the Probe
	Tags []string `json:"tags"`
	// Type of the Probe [From list of ProbeType enum]
	Type ProbeType `json:"type"`
	// Infrastructure type of the Probe
	InfrastructureType InfrastructureType `json:"infrastructureType"`
	// Kubernetes HTTP Properties of the specific type of the Probe
	KubernetesHTTPProperties *KubernetesHTTPProbe `json:"kubernetesHTTPProperties"`
	// Linux HTTP Properties of the specific type of the Probe
	LinuxHTTPProperties *LinuxHTTPProbe `json:"linuxHTTPProperties"`
	// Windows HTTP Properties of the specific type of the Probe
	WindowsHTTPProperties *WindowsHTTPProbe `json:"windowsHTTPProperties"`
	// Kubernetes CMD Properties of the specific type of the Probe
	KubernetesCMDProperties *KubernetesCMDProbe `json:"kubernetesCMDProperties"`
	// Linux CMD Properties of the specific type of the Probe
	LinuxCMDProperties *LinuxCMDProbe `json:"linuxCMDProperties"`
	// Kubernetes Datadog Properties of the specific type of the Probe
	KubernetesDatadogProperties *KubernetesDatadogProbe `json:"kubernetesDatadogProperties"`
	// Linux Datadog Properties of the specific type of the Probe
	LinuxDatadogProperties *LinuxDatadogProbe `json:"linuxDatadogProperties"`
	// K8S Properties of the specific type of the Probe
	K8sProperties *K8SProbe `json:"k8sProperties"`
	// Kubernetes Dynatrace Properties of the specific type of the Probe
	KubernetesDynatraceProperties *KubernetesDynatraceProbe `json:"kubernetesDynatraceProperties"`
	// Linux Dynatrace Properties of the specific type of the Probe
	LinuxDynatraceProperties *LinuxDynatraceProbe `json:"linuxDynatraceProperties"`
	// PROM Properties of the specific type of the Probe
	PromProperties *PROMProbe `json:"promProperties"`
	// SLO Properties of the specific type of the Probe
	SloProperties *SLOProbe `json:"sloProperties"`
	// APM Properties of the specific type of the Probe
	ApmProperties *APMProbe `json:"apmProperties"`
	// All execution histories of the probe
	RecentExecutions []*ProbeRecentExecutions `json:"recentExecutions"`
	// Referenced by how many faults
	ReferencedBy *int `json:"referencedBy"`
	// Is probe deleted or not
	IsRemoved bool `json:"isRemoved"`
	// Is probe enabled or disabled
	IsEnabled *bool `json:"isEnabled"`
	// Timestamp at which the Probe was last updated
	UpdatedAt string `json:"updatedAt"`
	// Timestamp at which the Probe was created
	CreatedAt string `json:"createdAt"`
	// User who has updated the Probe
	UpdatedBy *UserDetails `json:"updatedBy"`
	// User who has created the Probe
	CreatedBy *UserDetails `json:"createdBy"`
}

func (Probe) IsResourceDetails() {}
func (Probe) IsAudit()           {}

// Defines the input for Probe filter
type ProbeFilterInput struct {
	// Name of the Probe
	Name *string `json:"name"`
	// Date range for filtering purpose
	DateRange *DateRange `json:"dateRange"`
	// Type of the Probe [From list of ProbeType enum]
	Type []*ProbeType `json:"type"`
	// Is the probe enabled or disabled
	IsEnabled *bool `json:"isEnabled"`
	// Tags based search
	Tags []*string `json:"tags"`
}

type ProbeMap struct {
	FaultName string    `json:"faultName"`
	ProbeName []*string `json:"probeName"`
}

// Defines the Recent Executions of global probe in ListProbe API with different fault and execution history each time
type ProbeRecentExecutions struct {
	// Fault name
	FaultName string `json:"faultName"`
	// Fault Status
	Status *Status `json:"status"`
	// Fault executed by which experiment
	ExecutedByExperiment *ExecutedByExperiment `json:"executedByExperiment"`
}

// Defines the details required for updating a Chaos Probe
type ProbeRequest struct {
	// ID of the Probe
	ProbeID string `json:"probeID"`
	// Name of the Probe
	Name string `json:"name"`
	// Description of the Probe
	Description *string `json:"description"`
	// Tags of the Probe
	Tags []string `json:"tags"`
	// Type of the Probe [From list of ProbeType enum]
	Type ProbeType `json:"type"`
	// Is probe enabled or disabled
	IsEnabled *bool `json:"isEnabled"`
	// Is bulk enable for probe true or false
	IsBulkUpdateTrue *bool `json:"isBulkUpdateTrue"`
	// Infrastructure type of the Probe
	InfrastructureType InfrastructureType `json:"infrastructureType"`
	// HTTP Properties of the specific type of the Probe
	KubernetesHTTPProperties *KubernetesHTTPProbeRequest `json:"kubernetesHTTPProperties"`
	// HTTP Properties of the specific type of the Probe
	LinuxHTTPProperties *LinuxHTTPProbeRequest `json:"linuxHTTPProperties"`
	// HTTP Properties for windows of the specific type of the Probe
	WindowsHTTPProperties *WindowsHTTPProbeRequest `json:"windowsHTTPProperties"`
	// CMD Properties of the specific type of the Probe
	KubernetesCMDProperties *KubernetesCMDProbeRequest `json:"kubernetesCMDProperties"`
	// CMD Properties of the specific type of the Probe
	LinuxCMDProperties *LinuxCMDProbeRequest `json:"linuxCMDProperties"`
	// Datadog Properties of the specific type of the Probe
	KubernetesDatadogProperties *KubernetesDatadogProbeRequest `json:"kubernetesDatadogProperties"`
	// Datadog Properties of the specific type of the Probe
	LinuxDatadogProperties *LinuxDatadogProbeRequest `json:"linuxDatadogProperties"`
	// K8S Properties of the specific type of the Probe
	K8sProperties *K8SProbeRequest `json:"k8sProperties"`
	// Kubernetes Dynatrace Properties of the specific type of the Probe
	KubernetesDynatraceProperties *KubernetesDynatraceProbeRequest `json:"kubernetesDynatraceProperties"`
	// Linux Dynatrace Properties of the specific type of the Probe
	LinuxDynatraceProperties *LinuxDynatraceProbeRequest `json:"linuxDynatraceProperties"`
	// PROM Properties of the specific type of the Probe
	PromProperties *PROMProbeRequest `json:"promProperties"`
	// SLO Properties of the specific type of the Probe
	SloProperties *SLOProbeRequest `json:"sloProperties"`
}

// Defines sorting options for probes
type ProbeSortInput struct {
	// Field in which sorting will be done
	Field ProbeSortingField `json:"field"`
	// Bool value indicating whether the sorting will be done in ascending order
	Ascending *bool `json:"ascending"`
}

type Provider struct {
	Name string `json:"name"`
}

type PushProbeToChaosHubInput struct {
	// HubID for the selected ChaosHub
	ID string `json:"id"`
	// ProbeID to fetch the probe YAML
	ProbeID string `json:"probeID"`
	// Probe mode for execution
	Mode Mode `json:"mode"`
	// Probe Name
	ProbeName string `json:"probeName"`
	// Probe Description
	Description string `json:"description"`
	// Tags for the scenario
	Tags []string `json:"tags"`
}

type PushWorkflowToChaosHubInput struct {
	// HubID for the selected ChaosHub
	ID string `json:"id"`
	// Workflow Manifest to be pushed
	Manifest *string `json:"manifest"`
	// WorkflowID to fetch the manifest
	WorkflowID *string `json:"workflowID"`
	// Scenario Name
	ScenarioName string `json:"scenarioName"`
	// Scenario Description
	Description string `json:"description"`
	// Tags for the scenario
	Tags []string `json:"tags"`
	// Experiment info
	Experiments []*ExperimentInfoInput `json:"experiments"`
}

type QnARequest struct {
	QuestionType *QuestionType `json:"questionType"`
	OptionsMcq   []string      `json:"optionsMCQ"`
	Question     string        `json:"question"`
	Answer       string        `json:"answer"`
}

type QnAs struct {
	QuestionType QuestionType `json:"questionType"`
	OptionsMcq   []string     `json:"optionsMCQ"`
	Question     string       `json:"question"`
	Answer       string       `json:"answer"`
}

type ReRunChaosWorkflowResponse struct {
	NotifyID string `json:"notifyID"`
}

// Defines the Recent Executions of experiment referenced by the Probe
type RecentExecutions struct {
	// Fault name
	FaultName string `json:"faultName"`
	// Probe mode
	Mode Mode `json:"mode"`
	// Execution History
	ExecutionHistory []*ExecutionHistory `json:"executionHistory"`
}

type RecentWorkflowRun struct {
	// ID of the workflow run which is to be queried
	WorkflowRunID string `json:"workflowRunID"`
	// Notify ID for workflow run
	NotifyID *string `json:"notifyID"`
	// Phase of the workflow run
	Phase string `json:"phase"`
	// Resiliency score of the workflow
	ResiliencyScore *float64 `json:"resiliencyScore"`
	// Timestamp when the workflow was last updated
	UpdatedAt string `json:"updatedAt"`
	// Timestamp when the workflow was created
	CreatedAt string `json:"createdAt"`
	// User who created the workflow run
	CreatedBy *UserDetails `json:"createdBy"`
	// User who updated the workflow run
	UpdatedBy *UserDetails `json:"updatedBy"`
	// runSequence is the sequence number of experiment run
	RunSequence int `json:"runSequence"`
}

func (RecentWorkflowRun) IsAudit() {}

type Recurrence struct {
	Type RecurrenceType  `json:"type"`
	Spec *RecurrenceSpec `json:"spec"`
}

type RecurrenceInput struct {
	Type RecurrenceType       `json:"type"`
	Spec *RecurrenceSpecInput `json:"spec"`
}

type RecurrenceSpec struct {
	Until *int `json:"until"`
	Value *int `json:"value"`
}

type RecurrenceSpecInput struct {
	Until *int `json:"until"`
	Value *int `json:"value"`
}

// Defines the details for the new infra being connected
type RegisterCloudFoundryInfraRequest struct {
	// Name of the infra
	Name string `json:"name"`
	// Environment ID for the infra
	EnvironmentID string `json:"environmentID"`
	// Description of the infra
	Description *string `json:"description"`
	// Tags of the infra
	Tags []string `json:"tags"`
	// User inputs for the infra
	UserInputs *RegisterCloudFoundryInfraUserInputs `json:"userInputs"`
}

// Response received for registering a new infra
type RegisterCloudFoundryInfraResponse struct {
	// Unique ID for the newly registered infra
	InfraID string `json:"infraID"`
	// Infra name as sent in request
	Name string `json:"name"`
	// Infra access key
	AccessKey string `json:"accessKey"`
	// Infra server URL
	ServerURL string `json:"serverURL"`
	// Infra Version
	Version string `json:"version"`
	// Infra app manifest
	Manifest string `json:"manifest"`
}

// Defines the inputs for configuring cloud foundry infra
type RegisterCloudFoundryInfraUserInputs struct {
	// The polling interval for fetching the new tasks
	TaskPollInterval *string `json:"taskPollInterval"`
	// The interval at which abort is polled for a running task
	AbortPollInterval *string `json:"abortPollInterval"`
	// Specifies the maximum number of retries in case of a failure while sending a task status or result
	UpdateRetries *int `json:"updateRetries"`
	// Specifies the interval between the subsequent attempts to send a fault status or result, in case of a failure
	UpdateRetryInterval *string `json:"updateRetryInterval"`
	// The interval at which chaos infrastructure updates liveness heartbeat
	LivenessUpdateInterval *string `json:"livenessUpdateInterval"`
	// TLS certificate to be used with the infrastructure
	CustomTLSCertificate *string `json:"customTlsCertificate"`
	// URL for the http proxy
	HTTPProxy *string `json:"httpProxy"`
	// URL for the https proxy
	HTTPSProxy *string `json:"httpsProxy"`
	// Comma separated URLs that won't be proxied
	NoProxy *string `json:"noProxy"`
	// The timeout duration for the infrastructure http client
	HTTPClientTimeout *string `json:"httpClientTimeout"`
	// The origins to be exempted from CORS
	AllowedOrigins *string `json:"allowedOrigins"`
	// The port at which CFCO should be started
	HTTPPort *string `json:"httpPort"`
	// The interval at which app instances liveness is checked
	LivenessCheckInterval *string `json:"livenessCheckInterval"`
	// The maximum delay in the liveness update of app instances after which they get removed
	MaxHeartbeatDelay *string `json:"maxHeartbeatDelay"`
}

// Defines the details for the new infra being connected
type RegisterInfraRequest struct {
	// Name of the infra
	Name string `json:"name"`
	// Environment ID for the infra
	EnvironmentID string `json:"environmentID"`
	// Description of the infra
	Description *string `json:"description"`
	// Infra Platform Name eg. GKE,AWS, Others
	PlatformName string `json:"platformName"`
	// Namespace where the infra is being installed
	InfraNamespace *string `json:"infraNamespace"`
	// Name of service account used by infra
	ServiceAccount *string `json:"serviceAccount"`
	// Scope of the infra : ns or infra
	InfraScope InfraScope `json:"infraScope"`
	// Bool value indicating whether infra ns used already exists on infra or not
	InfraNsExists *bool `json:"infraNsExists"`
	// Bool value indicating whether service account used already exists on infra or not
	InfraSaExists *bool `json:"infraSaExists"`
	// InstallationType connector/manifest
	InstallationType InstallationType `json:"installationType"`
	// K8sConnectorID
	K8sConnectorID *string `json:"k8sConnectorID"`
	// Bool value indicating whether infra will skip ssl checks or not
	SkipSsl *bool `json:"skipSsl"`
	// Node selectors used by infra
	NodeSelector *string `json:"nodeSelector"`
	// Node tolerations used by infra
	Tolerations []*Toleration `json:"tolerations"`
	// Tune secret for infra
	IsSecretEnabled *bool `json:"isSecretEnabled"`
	// Tags of the infra
	Tags []string `json:"tags"`
	// set the user for security context in pod
	RunAsUser *int `json:"runAsUser"`
	// set the user group for security context in pod
	RunAsGroup *int `json:"runAsGroup"`
	// Value containing the scc-yaml used in openShift clusters
	SccYaml *string `json:"sccYaml"`
	// Boolean value indicating if chaos infrastructure has auto upgrade enabled or not
	IsAutoUpgradeEnabled bool `json:"isAutoUpgradeEnabled"`
}

// Response received for registering a new infra
type RegisterInfraResponse struct {
	// Token used to verify and retrieve the infra manifest
	Token string `json:"token"`
	// Unique ID for the newly registered infra
	InfraID string `json:"infraID"`
	// Infra name as sent in request
	Name string `json:"name"`
	// Infra Manifest
	Manifest string `json:"manifest"`
	// taskID sent for the brownfield deployment task
	TaskID *string `json:"taskID"`
}

// Defines the details for the new infra being connected
type RegisterLinuxInfraRequest struct {
	// Name of the infra
	Name string `json:"name"`
	// Environment ID for the infra
	EnvironmentID string `json:"environmentID"`
	// Description of the infra
	Description *string `json:"description"`
	// Tags of the infra
	Tags []string `json:"tags"`
}

// Response received for registering a new infra
type RegisterLinuxInfraResponse struct {
	// Unique ID for the newly registered infra
	InfraID string `json:"infraID"`
	// Infra name as sent in request
	Name string `json:"name"`
	// Infra access key
	AccessKey string `json:"accessKey"`
	// Linux Infra server URL
	ServerURL string `json:"serverURL"`
	// Linux Infra Version
	Version string `json:"version"`
	// Linux Infra AccountID
	AccountID string `json:"accountID"`
}

// Defines the details for the new infra being connected
type RegisterWindowsInfraRequest struct {
	// Name of the infra
	Name string `json:"name"`
	// Environment ID for the infra
	EnvironmentID string `json:"environmentID"`
	// Description of the infra
	Description *string `json:"description"`
	// Tags of the infra
	Tags []string `json:"tags"`
}

// Response received for registering a new infra
type RegisterWindowsInfraResponse struct {
	// Unique ID for the newly registered infra
	InfraID string `json:"infraID"`
	// Infra name as sent in request
	Name string `json:"name"`
	// Infra access key
	AccessKey string `json:"accessKey"`
	// Windows Infra server URL
	ServerURL string `json:"serverURL"`
	// Windows Infra Version
	Version string `json:"version"`
	// Windows Infra AccountID
	AccountID string `json:"accountID"`
}

type ResilienceScoreCategory struct {
	// Lower bound of the range(inclusive)
	ID int `json:"id"`
	// total experiments with avg resilience score between lower bound and upper bound(exclusive)
	Count int `json:"count"`
}

type Rule struct {
	Name         string        `json:"name"`
	Tags         []string      `json:"tags"`
	Description  *string       `json:"description"`
	IsEnabled    bool          `json:"isEnabled"`
	RuleID       string        `json:"ruleId"`
	UserGroupIds []string      `json:"userGroupIds"`
	TimeWindows  []*TimeWindow `json:"timeWindows"`
	Conditions   []*Condition  `json:"conditions"`
}

func (Rule) IsResourceDetails() {}

type RuleDetails struct {
	RuleID       *string             `json:"ruleId"`
	RuleName     *string             `json:"ruleName"`
	Message      *string             `json:"message"`
	Description  *string             `json:"description"`
	UserGroupIds []*string           `json:"userGroupIds"`
	TimeWindow   *TimeWindow         `json:"timeWindow"`
	Conditions   []*ConditionDetails `json:"conditions"`
}

type RuleFilterInput struct {
	Name        *string   `json:"name"`
	Description *string   `json:"description"`
	Tags        []*string `json:"tags"`
}

type RuleInput struct {
	Description  *string            `json:"description"`
	IsEnabled    bool               `json:"isEnabled"`
	Name         string             `json:"name"`
	RuleID       string             `json:"ruleId"`
	UserGroupIds []string           `json:"userGroupIds"`
	TimeWindows  []*TimeWindowInput `json:"timeWindows"`
	Tags         []*string          `json:"tags"`
	ConditionIds []string           `json:"conditionIds"`
}

type RuleResponse struct {
	CreatedAt   int          `json:"createdAt"`
	CreatedBy   *UserDetails `json:"createdBy"`
	UpdatedAt   int          `json:"updatedAt"`
	UpdatedBy   *UserDetails `json:"updatedBy"`
	Rule        *Rule        `json:"rule"`
	Identifiers *Identifiers `json:"identifiers"`
}

func (RuleResponse) IsAuditV2() {}

type RunChaosExperimentPipelineResponse struct {
	NotifyID string  `json:"notifyID"`
	Error    *string `json:"error"`
}

type RunChaosExperimentResponse struct {
	NotifyID string `json:"notifyID"`
}

type RunPipelineExperimentMetaData struct {
	Name     string `json:"name"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
}

// Run time inputs
type RunTimeInputs struct {
	// Experiment variables
	Experiment []*VariableMinimum `json:"experiment"`
	// Tasks
	Tasks []*TaskEntry `json:"tasks"`
}

// Defines the SLO probe properties
type SLOProbe struct {
	// Timeout of the Probe
	ProbeTimeout string `json:"probeTimeout"`
	// Interval of the Probe
	Interval string `json:"interval"`
	// Retry interval of the Probe
	Retry *int `json:"retry"`
	// Attempt contains the total attempt count for the probe
	Attempt int `json:"attempt"`
	// Polling interval of the Probe
	ProbePollingInterval *string `json:"probePollingInterval"`
	// Initial delay interval of the Probe in seconds
	InitialDelay *string `json:"initialDelay"`
	// Verbosity contains flag to set the verbosity of probe
	EvaluationTimeout string `json:"evaluationTimeout"`
	// Is stop on failure enabled in the Probe
	StopOnFailure *bool `json:"stopOnFailure"`
	// PlatformEndpoint for the monitoring service endpoint
	PlatformEndpoint string `json:"platformEndpoint"`
	// SLOIdentifier for fetching the details of the SLO
	SloIdentifier string `json:"sloIdentifier"`
	// EvaluationWindow is the time period for which the metrics will be evaluated
	EvaluationWindow *EvaluationWindow `json:"evaluationWindow"`
	// SLOSourceMetadata consists of required metadata details to fetch metric data
	SloSourceMetadata *SLOSourceMetadata `json:"sloSourceMetadata"`
	// Comparator check for the correctness of the probe output
	Comparator *Comparator `json:"comparator"`
	// Verbosity for the probe logging
	Verbosity *string `json:"verbosity"`
	// If Insecure HTTP verification should  be skipped
	InsecureSkipVerify *bool `json:"insecureSkipVerify"`
}

func (SLOProbe) IsCommonProbeProperties() {}

// Defines the input for SLO probes
type SLOProbeRequest struct {
	// Timeout of the Probe
	ProbeTimeout string `json:"probeTimeout"`
	// Interval of the Probe
	Interval string `json:"interval"`
	// Retry interval of the Probe
	Retry *int `json:"retry"`
	// Attempt contains the total attempt count for the probe
	Attempt int `json:"attempt"`
	// Polling interval of the Probe
	ProbePollingInterval *string `json:"probePollingInterval"`
	// Initial delay interval of the Probe in seconds
	InitialDelay *string `json:"initialDelay"`
	// Is stop on failure enabled in the Probe
	StopOnFailure *bool `json:"stopOnFailure"`
	// PlatformEndpoint for the monitoring service endpoint
	PlatformEndpoint string `json:"platformEndpoint"`
	// SLOIdentifier for fetching the details of the SLO
	SloIdentifier string `json:"sloIdentifier"`
	// EvaluationWindow is the time period for which the metrics will be evaluated
	EvaluationWindow *EvaluationWindowInput `json:"evaluationWindow"`
	// SLOSourceMetadata consists of required metadata details to fetch metric data
	SloSourceMetadata *SLOSourceMetadataInput `json:"sloSourceMetadata"`
	// Comparator check for the correctness of the probe output
	Comparator *ComparatorInput `json:"comparator"`
	// EvaluationTimeout is the timeout window in which the SLO metrics
	EvaluationTimeout string `json:"evaluationTimeout"`
	// Verbosity for the probe logging
	Verbosity *string `json:"verbosity"`
	// If Insecure HTTP verification should  be skipped
	InsecureSkipVerify *bool `json:"insecureSkipVerify"`
}

// SLOSourceMetadata consists of required metadata details to fetch metric data
type SLOSourceMetadata struct {
	// APITokenSecret for authenticating with the platform service
	APITokenSecret string `json:"apiTokenSecret"`
	// Scope required for fetching details
	Scope *Identifiers `json:"scope"`
}

// Defines the input for SLOSourceMetadata
type SLOSourceMetadataInput struct {
	// APITokenSecret for authenticating with the platform service
	APITokenSecret string `json:"apiTokenSecret"`
	// Scope required for fetching details
	Scope *IdentifiersRequest `json:"scope"`
}

type SSHKey struct {
	PublicKey  string `json:"publicKey"`
	PrivateKey string `json:"privateKey"`
}

type ScopedIdentifiers struct {
	// Harness AccountID
	AccountIdentifier string `json:"accountIdentifier"`
	// Harness OrgID
	OrgIdentifier *string `json:"orgIdentifier"`
	// Harness projectID
	ProjectIdentifier *string `json:"projectIdentifier"`
}

type ScopedIdentifiersRequest struct {
	// Harness AccountID
	AccountIdentifier string `json:"accountIdentifier"`
	// Harness OrgID
	OrgIdentifier *string `json:"orgIdentifier"`
	// Harness projectID
	ProjectIdentifier *string `json:"projectIdentifier"`
}

// SecretManager configures the options for TLS certificates
type SecretManager struct {
	// Flag to hold the secret identifier
	Identifier string `json:"identifier"`
}

type SecurityGovernance struct {
	Name                       *string                     `json:"name"`
	Type                       *string                     `json:"type"`
	StartedAt                  *int                        `json:"startedAt"`
	FinishedAt                 *int                        `json:"finishedAt"`
	Message                    *string                     `json:"message"`
	Phase                      *SecurityGovernancePhase    `json:"phase"`
	SecurityGovernanceNodeData *SecurityGovernanceNodeData `json:"securityGovernanceNodeData"`
}

type SecurityGovernanceNodeData struct {
	PassedRules  []*RuleDetails `json:"passedRules"`
	FailedRules  []*RuleDetails `json:"failedRules"`
	SkippedRules []*RuleDetails `json:"skippedRules"`
}

type Service struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

type ServiceInput struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

type Spec struct {
	DisplayName         string        `json:"displayName"`
	CategoryDescription string        `json:"categoryDescription"`
	Keywords            []string      `json:"keywords"`
	Maturity            string        `json:"maturity"`
	Maintainers         []*Maintainer `json:"maintainers"`
	MinKubeVersion      string        `json:"minKubeVersion"`
	Provider            *Provider     `json:"provider"`
	Links               []*Link       `json:"links"`
	Faults              []*FaultList  `json:"faults"`
	Experiments         []string      `json:"experiments"`
	ChaosExpCRDLink     string        `json:"chaosExpCRDLink"`
	Platforms           []string      `json:"platforms"`
	ChaosType           *string       `json:"chaosType"`
}

// Defines the APM splunk observability metrics properties
type SplunkObservabilityMetrics struct {
	// query of the metrics for splunk observability probe
	Query string `json:"query"`
	// duration in minutes for splunk observability probe
	DurationInMin int `json:"durationInMin"`
}

type StandardResponse struct {
	Message       string `json:"message"`
	CorrelationID string `json:"correlationId"`
	Response      string `json:"response"`
}

// Status defines whether a probe is pass or fail
type Status struct {
	// Verdict defines the verdict of the probe, range: Passed, Failed, N/A
	Verdict ProbeVerdict `json:"verdict"`
	// Description defines the description of probe status
	Description *string `json:"description"`
}

type StopGameDayRunExperimentRequest struct {
	GameDayRunID string `json:"gameDayRunID"`
	ExperimentID string `json:"experimentID"`
}

type StopGameDayRunRequest struct {
	GameDayRunID string `json:"gameDayRunID"`
}

type SummaryRequest struct {
	Notes      string               `json:"notes"`
	Qna        []*QnARequest        `json:"qna"`
	ActionItem []*ActionItemRequest `json:"actionItem"`
}

// Synthetics test parameters
type SyntheticsTest struct {
	// Type of the test; supports 'api' and 'browser' only
	TestType DatadogSyntheticsTestType `json:"testType"`
	// Public id of the test
	PublicID string `json:"publicId"`
}

// Defines the input for Synthetics test parameters
type SyntheticsTestRequest struct {
	// Type of the test; supports 'api' and 'browser' only
	TestType DatadogSyntheticsTestType `json:"testType"`
	// Public id of the test
	PublicID string `json:"publicId"`
}

// TLSConfig configures the options for TLS connections
type TLSConfig struct {
	// Flag to hold the ca file path
	CaFile *string `json:"caFile"`
	// Flag to hold the client cert file path
	CertFile *string `json:"certFile"`
	// Flag to hold the client key file path
	KeyFile *string `json:"keyFile"`
	// Flag to skip the tls certificates checks
	InsecureSkipVerify *bool `json:"insecureSkipVerify"`
}

// Defines the input for TLSConfig options for TLS connections
type TLSConfigInput struct {
	// Flag to hold the ca file path
	CaFile *string `json:"caFile"`
	// Flag to hold the client cert file path
	CertFile *string `json:"certFile"`
	// Flag to hold the client key file path
	KeyFile *string `json:"keyFile"`
	// Flag to skip the tls certificates checks
	InsecureSkipVerify *bool `json:"insecureSkipVerify"`
}

// Task entry
type TaskEntry struct {
	// Task name
	Key string `json:"key"`
	// Variable name and value
	Values []*VariableMinimum `json:"values"`
}

type TimeWindow struct {
	Duration   *string     `json:"duration"`
	EndTime    *int        `json:"endTime"`
	StartTime  int         `json:"startTime"`
	TimeZone   string      `json:"timeZone"`
	Recurrence *Recurrence `json:"recurrence"`
}

type TimeWindowInput struct {
	Duration   *string          `json:"duration"`
	EndTime    *int             `json:"endTime"`
	StartTime  int              `json:"startTime"`
	TimeZone   string           `json:"timeZone"`
	Recurrence *RecurrenceInput `json:"recurrence"`
}

type Toleration struct {
	TolerationSeconds *int    `json:"tolerationSeconds"`
	Key               *string `json:"key"`
	Operator          *string `json:"operator"`
	Effect            *string `json:"effect"`
	Value             *string `json:"value"`
}

// Defines input for updating cron experiment state
type UpdateCronExperimentStateRequest struct {
	// ID of the experiment
	ExperimentIDs []string `json:"experimentIDs"`
	// Action indicating whether to enable, disable or update the cron experiment
	Action UpdateCronExperimentAction `json:"action"`
	// Flag to check if single run status is enabled or not
	IsSingleRunCronEnabled *bool `json:"isSingleRunCronEnabled"`
	// Cron syntax to be updated
	CronSyntax *string `json:"cronSyntax"`
}

// Defines the response from UpdateCronExperimentState API
type UpdateCronExperimentStateResponse struct {
	// List of successful experiment ID
	SuccessIDs []string `json:"successIDs"`
	// List of failed experiment ID
	FailedIDs []string `json:"failedIDs"`
}

type UpdateExperimentNotesRequest struct {
	GameDayRunID string  `json:"gameDayRunID"`
	ExperimentID *string `json:"experimentID"`
	Notes        string  `json:"notes"`
}

type UpdateGameDayExperimentsRequest struct {
	GameDayID   string               `json:"gameDayID"`
	Experiments []*NonCronExperiment `json:"experiments"`
}

type UpdateGameDayRequest struct {
	GameDayID   string               `json:"gameDayID"`
	Name        *string              `json:"name"`
	Objective   *string              `json:"objective"`
	Description *string              `json:"description"`
	Summary     *SummaryRequest      `json:"summary"`
	IsRemoved   *bool                `json:"isRemoved"`
	Experiments []*NonCronExperiment `json:"experiments"`
}

type UpdateGameDayRunRequest struct {
	GameDayRunID string  `json:"gameDayRunID"`
	Completed    *bool   `json:"completed"`
	Notes        *string `json:"notes"`
}

type UpdateInfraRequest struct {
	// ID of the infrastructure to be updated
	InfraID string `json:"infraID"`
	// Name of the infra
	Name *string `json:"name"`
	// Environment ID for the infra
	EnvironmentID *string `json:"environmentID"`
	// Description of the infra
	Description *string `json:"description"`
	// Tags of the infra
	Tags []string `json:"tags"`
}

type UpdateSummaryRequest struct {
	ID             string          `json:"ID"`
	SummaryRequest *SummaryRequest `json:"summaryRequest"`
	Type           Type            `json:"type"`
}

type Upgrade struct {
	Status               UpgradeStatus `json:"status"`
	IsAutoUpgradeEnabled bool          `json:"isAutoUpgradeEnabled"`
}

type UserDetails struct {
	UserID   string `json:"userID"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// Variable name and value
type VariableMinimum struct {
	// Name of the variable
	Name string `json:"name"`
	// Value of the variable
	Value interface{} `json:"value"`
}

// Defines the details of the weightages of each chaos experiment in the workflow
type Weightages struct {
	// Name of the experiment
	ExperimentName string `json:"experimentName"`
	// Weightage of the experiment
	Weightage int `json:"weightage"`
}

// Defines the details of the weightages of each chaos experiment in the workflow
type WeightagesInput struct {
	// Name of the experiment
	ExperimentName string `json:"experimentName"`
	// Weightage of the experiment
	Weightage int `json:"weightage"`
}

// Defines the Windows HTTP probe properties
type WindowsHTTPProbe struct {
	// Timeout of the Probe
	ProbeTimeout string `json:"probeTimeout"`
	// Interval of the Probe
	Interval string `json:"interval"`
	// Attempt contains the total attempt count for the probe
	Attempt int `json:"attempt"`
	// Polling interval of the Probe
	ProbePollingInterval *string `json:"probePollingInterval"`
	// Initial delay interval of the Probe in seconds
	InitialDelay *string `json:"initialDelay"`
	// Is stop on failure enabled in the Probe
	StopOnFailure *bool `json:"stopOnFailure"`
	// URL of the Probe
	URL string `json:"url"`
	// HTTP method of the Probe
	Method *Method `json:"method"`
	// Auth contains the authentication details for the HTTP probe
	Auth *HTTPAuthorization `json:"auth"`
	// TLSConfig contains the tls configuration for the HTTP probe
	TLSConfig *TLSConfig `json:"tlsConfig"`
	// Headers contains the request headers
	Headers []*Headers `json:"headers"`
}

// Defines the input for windows HTTP probe properties
type WindowsHTTPProbeRequest struct {
	// Timeout of the Probe
	ProbeTimeout string `json:"probeTimeout"`
	// Interval of the Probe
	Interval string `json:"interval"`
	// Attempt contains the total attempt count for the probe
	Attempt int `json:"attempt"`
	// Polling interval of the Probe
	ProbePollingInterval *string `json:"probePollingInterval"`
	// Initial delay interval of the Probe in seconds
	InitialDelay *string `json:"initialDelay"`
	// Is stop on failure enabled in the Probe
	StopOnFailure *bool `json:"stopOnFailure"`
	// URL of the Probe
	URL string `json:"url"`
	// HTTP method of the Probe
	Method *MethodRequest `json:"method"`
	// Auth contains the authentication details for the HTTP probe
	Auth *HTTPAuthorizationInput `json:"auth"`
	// TLSConfig contains the tls configuration for the prometheus probe
	TLSConfig *TLSConfigInput `json:"tlsConfig"`
	// Headers contains the request headers
	Headers []*HeadersRequest `json:"headers"`
}

// Defines the details for a infra
type WindowsInfra struct {
	// ID of the infra
	InfraID string `json:"infraID"`
	// Name of the infra
	Name string `json:"name"`
	// Description of the infra
	Description *string `json:"description"`
	// Tags of the infra
	Tags []string `json:"tags"`
	// Environment ID for the infra
	EnvironmentID string `json:"environmentID"`
	// Boolean value indicating if chaos infrastructure is active or not
	IsActive bool `json:"isActive"`
	// Boolean value indicating if chaos infrastructure is confirmed or not
	IsInfraConfirmed bool `json:"isInfraConfirmed"`
	// Timestamp when the infra was last updated
	UpdatedAt string `json:"updatedAt"`
	// Timestamp when the infra was created
	CreatedAt string `json:"createdAt"`
	// Number of schedules created in the infra
	NoOfSchedules *int `json:"noOfSchedules"`
	// Number of workflows run in the infra
	NoOfWorkflows *int `json:"noOfWorkflows"`
	// Timestamp of the last workflow run in the infra
	LastWorkflowTimestamp *string `json:"lastWorkflowTimestamp"`
	// Timestamp when the infra got connected
	StartTime string `json:"startTime"`
	// Version of the infra
	Version string `json:"version"`
	// User who created the infra
	CreatedBy *UserDetails `json:"createdBy"`
	// User who has updated the infra
	UpdatedBy *UserDetails `json:"updatedBy"`
	// Last Heartbeat status sent by the infra
	LastHeartbeat *string `json:"lastHeartbeat"`
	// hostname of the infra
	Hostname *string `json:"hostname"`
}

func (WindowsInfra) IsResourceDetails() {}
func (WindowsInfra) IsAudit()           {}

// Defines filter options for infras
type WindowsInfraFilterInput struct {
	// Name of the infra
	Name *string `json:"name"`
	// ID of the infra
	InfraID *string `json:"infraID"`
	// ID of the infra
	Description *string `json:"description"`
	// Status of infra
	IsActive *bool `json:"isActive"`
	// Tags of an infra
	Tags []*string `json:"tags"`
}

// Defines the details for a workflow
type Workflow struct {
	// Harness identifiers
	Identifiers *Identifiers `json:"identifiers"`
	// Identity of the experiment
	Identity *string `json:"identity"`
	// ID of the workflow
	WorkflowID string `json:"workflowID"`
	// Type of the workflow
	WorkflowType *string `json:"workflowType"`
	// Manifest of the workflow
	WorkflowManifest string `json:"workflowManifest"`
	// Cron syntax of the workflow schedule
	CronSyntax string `json:"cronSyntax"`
	// If cron is enabled or disabled
	IsCronEnabled *bool `json:"isCronEnabled"`
	// Flag to check if single run status is enabled or not
	IsSingleRunCronEnabled *bool `json:"isSingleRunCronEnabled"`
	// Name of the workflow
	Name string `json:"name"`
	// Description of the workflow
	Description string `json:"description"`
	// Array containing weightage and name of each chaos experiment in the workflow
	Weightages []*Weightages `json:"weightages"`
	// Bool value indicating whether the workflow is a custom workflow or not
	IsCustomWorkflow bool `json:"isCustomWorkflow"`
	// Timestamp when the workflow was last updated
	UpdatedAt string `json:"updatedAt"`
	// Timestamp when the workflow was created
	CreatedAt string `json:"createdAt"`
	// Target infra in which the workflow will run
	Infra *Infrastructure `json:"infra"`
	// Bool value indicating if the workflow has removed
	IsRemoved bool `json:"isRemoved"`
	// Tags of the workflow
	Tags []string `json:"tags"`
	// User who created the workflow
	CreatedBy *UserDetails `json:"createdBy"`
	// Array of object containing details of recent workflow runs
	RecentWorkflowRunDetails []*RecentWorkflowRun `json:"recentWorkflowRunDetails"`
	// Array containing service identifier and environment identifier
	// for SRM change source events
	EventsMetadata []*EventMetadata `json:"eventsMetadata"`
	// Details of the user who updated the workflow
	UpdatedBy *UserDetails `json:"updatedBy"`
	// Timestamp when the experiment was last executed
	LastExecutedAt string `json:"lastExecutedAt"`
}

func (Workflow) IsResourceDetails() {}
func (Workflow) IsAudit()           {}

// Defines filter options for workflows
type WorkflowFilterInput struct {
	// Name of the workflow
	WorkflowName *string `json:"workflowName"`
	// Name of the infra in which the workflow is running
	InfraName *string `json:"infraName"`
	// IDs of the agent in which the workflow is running
	InfraIDs []*string `json:"infraIDs"`
	// Bool value indicating if Chaos Infrastructure is active
	InfraActive *bool `json:"infraActive"`
	// Scenario type of the workflow i.e. CRON or NON_CRON
	ScenarioType *ScenarioType `json:"scenarioType"`
	// Status of the latest workflow run
	Status *string `json:"status"`
	// Date range for filtering purpose
	DateRange *DateRange `json:"dateRange"`
	// Type of infras
	InfraTypes []*InfrastructureType `json:"infraTypes"`
	// Tags based search
	Tags []*string `json:"tags"`
	// Bool value to filter data based on cron enabled state
	IsCronEnabled *bool `json:"isCronEnabled"`
}

// Defines the details of a workflow run
type WorkflowRun struct {
	// Harness identifiers
	Identifiers *Identifiers `json:"identifiers"`
	// ID of the workflow run which is to be queried
	WorkflowRunID string `json:"workflowRunID"`
	// Type of the workflow
	WorkflowType *string `json:"workflowType"`
	// ID of the workflow
	WorkflowID string `json:"workflowID"`
	// Array containing weightage and name of each chaos experiment in the workflow
	Weightages []*Weightages `json:"weightages"`
	// Timestamp at which workflow run was last updated
	UpdatedAt string `json:"updatedAt"`
	// Timestamp at which workflow run was created
	CreatedAt string `json:"createdAt"`
	// Cron syntax of the workflow schedule
	CronSyntax string `json:"cronSyntax"`
	// Target infra in which the workflow will run
	Infra *Infrastructure `json:"infra"`
	// Name of the workflow
	WorkflowName string `json:"workflowName"`
	// Description of the workflow
	WorkflowDescription *string `json:"workflowDescription"`
	// Tag of the workflow
	WorkflowTags []*string `json:"workflowTags"`
	// Manifest of the workflow run
	WorkflowManifest string `json:"workflowManifest"`
	// If cron is enabled or disabled
	IsCronEnabled *bool `json:"isCronEnabled"`
	// Flag to check if single run status is enabled or not
	IsSingleRunCronEnabled *bool `json:"isSingleRunCronEnabled"`
	// Probe object containing reference of probeIDs
	Probe []*ProbeMap `json:"probe"`
	// Phase of the workflow run
	Phase WorkflowRunStatus `json:"phase"`
	// Resiliency score of the workflow
	ResiliencyScore *float64 `json:"resiliencyScore"`
	// Number of experiments passed
	ExperimentsPassed *int `json:"experimentsPassed"`
	// Number of experiments failed
	ExperimentsFailed *int `json:"experimentsFailed"`
	// Number of experiments awaited
	ExperimentsAwaited *int `json:"experimentsAwaited"`
	// Number of experiments stopped
	ExperimentsStopped *int `json:"experimentsStopped"`
	// Number of experiments which are not available
	ExperimentsNa *int `json:"experimentsNa"`
	// Total number of experiments
	TotalExperiments *int `json:"totalExperiments"`
	// Stores all the workflow run details related to the nodes of DAG graph and chaos results of the experiments
	ExecutionData string `json:"executionData"`
	// Bool value indicating if the workflow run has removed
	IsRemoved *bool `json:"isRemoved"`
	// User who has updated the workflow
	UpdatedBy *UserDetails `json:"updatedBy"`
	// User who has created the experiment run
	CreatedBy *UserDetails `json:"createdBy"`
	// Notify ID of the experiment run
	NotifyID *string `json:"notifyID"`
	// Error Response is the reason why experiment failed to run
	ErrorResponse *string `json:"errorResponse"`
	// Security Governance details of the workflow run
	SecurityGovernance *SecurityGovernance `json:"securityGovernance"`
	// runSequence is the sequence number of experiment run
	RunSequence int `json:"runSequence"`
	// experimentType is the type of experiment run
	ExperimentType string `json:"experimentType"`
	// kind of the experiment manifest
	ManifestVersion string `json:"manifestVersion"`
}

func (WorkflowRun) IsAudit() {}

// Defines input type for workflow run filter
type WorkflowRunFilterInput struct {
	// Sequence of the run
	RunSequence *int `json:"runSequence"`
	// Name of the workflow
	WorkflowName *string `json:"workflowName"`
	// Name of the infra infra
	InfraIDs []*string `json:"infraIDs"`
	// Type of the workflow
	WorkflowType *ScenarioType `json:"workflowType"`
	// Status of the workflow run
	WorkflowStatus *WorkflowRunStatus `json:"workflowStatus"`
	// Date range for filtering purpose
	DateRange *DateRange `json:"dateRange"`
	// ID of experiment run
	WorkflowRunID *string `json:"workflowRunID"`
	// Array of workflow run status
	WorkflowRunStatus []*string `json:"workflowRunStatus"`
	// Type of infras
	InfraTypes []*InfrastructureType `json:"infraTypes"`
}

// Defines the details for a workflow run
type WorkflowRunRequest struct {
	// ID of the workflow
	WorkflowID string `json:"workflowID"`
	// ID of the workflow run which is to be queried
	WorkflowRunID string `json:"workflowRunID"`
	// Name of the workflow
	WorkflowName string `json:"workflowName"`
	// Stores all the workflow run details related to the nodes of DAG graph and chaos results of the experiments
	ExecutionData string `json:"executionData"`
	// ID of the infra infra in which the workflow is running
	InfraID *InfraIdentity `json:"infraID"`
	// ID of the revision which consists manifest details
	RevisionID string `json:"revisionID"`
	// Notify ID is used to retrun re-run validation of an experiment
	NotifyID *string `json:"notifyID"`
	// Bool value indicating if the workflow run has completed
	Completed bool `json:"completed"`
	// Bool value indicating if the workflow run has removed
	IsRemoved *bool `json:"isRemoved"`
	// User who has updated the workflow
	UpdatedBy string `json:"updatedBy"`
}

// Defines sorting options for workflow runs
type WorkflowRunSortInput struct {
	// Field in which sorting will be done
	Field WorkflowSortingField `json:"field"`
	// Bool value indicating whether the sorting will be done in ascending order
	Ascending *bool `json:"ascending"`
}

// Defines sorting options for workflow
type WorkflowSortInput struct {
	// Field in which sorting will be done
	Field WorkflowSortingField `json:"field"`
	// Bool value indicating whether the sorting will be done in ascending order
	Ascending *bool `json:"ascending"`
}

type Workload struct {
	Label            *string  `json:"label"`
	Namespace        string   `json:"namespace"`
	Kind             *string  `json:"kind"`
	Services         []string `json:"services"`
	ApplicationMapID *string  `json:"applicationMapId"`
	Env              []*Env   `json:"env"`
}

type WorkloadInput struct {
	Label            *string     `json:"label"`
	Namespace        string      `json:"namespace"`
	Kind             *string     `json:"kind"`
	Services         []string    `json:"services"`
	ApplicationMapID *string     `json:"applicationMapId"`
	Env              []*EnvInput `json:"env"`
}

type ChaosHubAuthType string

const (
	ChaosHubAuthTypeSSH           ChaosHubAuthType = "Ssh"
	ChaosHubAuthTypeUsernameToken ChaosHubAuthType = "UsernameToken"
)

var AllChaosHubAuthType = []ChaosHubAuthType{
	ChaosHubAuthTypeSSH,
	ChaosHubAuthTypeUsernameToken,
}

func (e ChaosHubAuthType) IsValid() bool {
	switch e {
	case ChaosHubAuthTypeSSH, ChaosHubAuthTypeUsernameToken:
		return true
	}
	return false
}

func (e ChaosHubAuthType) String() string {
	return string(e)
}

func (e *ChaosHubAuthType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ChaosHubAuthType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ChaosHubAuthType", str)
	}
	return nil
}

func (e ChaosHubAuthType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ClusterType string

const (
	ClusterTypeKubernetes ClusterType = "KUBERNETES"
	ClusterTypeOpenshift  ClusterType = "OPENSHIFT"
	ClusterTypeHelm       ClusterType = "HELM"
)

var AllClusterType = []ClusterType{
	ClusterTypeKubernetes,
	ClusterTypeOpenshift,
	ClusterTypeHelm,
}

func (e ClusterType) IsValid() bool {
	switch e {
	case ClusterTypeKubernetes, ClusterTypeOpenshift, ClusterTypeHelm:
		return true
	}
	return false
}

func (e ClusterType) String() string {
	return string(e)
}

func (e *ClusterType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ClusterType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ClusterType", str)
	}
	return nil
}

func (e ClusterType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ConnectorScope string

const (
	ConnectorScopeAccount      ConnectorScope = "ACCOUNT"
	ConnectorScopeProject      ConnectorScope = "PROJECT"
	ConnectorScopeOrganisation ConnectorScope = "ORGANISATION"
)

var AllConnectorScope = []ConnectorScope{
	ConnectorScopeAccount,
	ConnectorScopeProject,
	ConnectorScopeOrganisation,
}

func (e ConnectorScope) IsValid() bool {
	switch e {
	case ConnectorScopeAccount, ConnectorScopeProject, ConnectorScopeOrganisation:
		return true
	}
	return false
}

func (e ConnectorScope) String() string {
	return string(e)
}

func (e *ConnectorScope) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ConnectorScope(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ConnectorScope", str)
	}
	return nil
}

func (e ConnectorScope) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type DatadogSyntheticsTestType string

const (
	DatadogSyntheticsTestTypeAPI     DatadogSyntheticsTestType = "api"
	DatadogSyntheticsTestTypeBrowser DatadogSyntheticsTestType = "browser"
)

var AllDatadogSyntheticsTestType = []DatadogSyntheticsTestType{
	DatadogSyntheticsTestTypeAPI,
	DatadogSyntheticsTestTypeBrowser,
}

func (e DatadogSyntheticsTestType) IsValid() bool {
	switch e {
	case DatadogSyntheticsTestTypeAPI, DatadogSyntheticsTestTypeBrowser:
		return true
	}
	return false
}

func (e DatadogSyntheticsTestType) String() string {
	return string(e)
}

func (e *DatadogSyntheticsTestType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DatadogSyntheticsTestType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid DatadogSyntheticsTestType", str)
	}
	return nil
}

func (e DatadogSyntheticsTestType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ExecutionPlane string

const (
	ExecutionPlaneLinux        ExecutionPlane = "LINUX"
	ExecutionPlaneKubernetes   ExecutionPlane = "KUBERNETES"
	ExecutionPlaneKubernetesv2 ExecutionPlane = "KUBERNETESV2"
	ExecutionPlaneWindows      ExecutionPlane = "WINDOWS"
	ExecutionPlaneCloudfoundry ExecutionPlane = "CLOUDFOUNDRY"
)

var AllExecutionPlane = []ExecutionPlane{
	ExecutionPlaneLinux,
	ExecutionPlaneKubernetes,
	ExecutionPlaneKubernetesv2,
	ExecutionPlaneWindows,
	ExecutionPlaneCloudfoundry,
}

func (e ExecutionPlane) IsValid() bool {
	switch e {
	case ExecutionPlaneLinux, ExecutionPlaneKubernetes, ExecutionPlaneKubernetesv2, ExecutionPlaneWindows, ExecutionPlaneCloudfoundry:
		return true
	}
	return false
}

func (e ExecutionPlane) String() string {
	return string(e)
}

func (e *ExecutionPlane) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ExecutionPlane(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ExecutionPlane", str)
	}
	return nil
}

func (e ExecutionPlane) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type FaultType string

const (
	FaultTypeFaultGroup FaultType = "FAULT_GROUP"
	FaultTypeFault      FaultType = "FAULT"
)

var AllFaultType = []FaultType{
	FaultTypeFaultGroup,
	FaultTypeFault,
}

func (e FaultType) IsValid() bool {
	switch e {
	case FaultTypeFaultGroup, FaultTypeFault:
		return true
	}
	return false
}

func (e FaultType) String() string {
	return string(e)
}

func (e *FaultType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = FaultType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid FaultType", str)
	}
	return nil
}

func (e FaultType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type FileType string

const (
	FileTypeExperiment FileType = "EXPERIMENT"
	FileTypeEngine     FileType = "ENGINE"
	FileTypeFault      FileType = "FAULT"
	FileTypeCsv        FileType = "CSV"
)

var AllFileType = []FileType{
	FileTypeExperiment,
	FileTypeEngine,
	FileTypeFault,
	FileTypeCsv,
}

func (e FileType) IsValid() bool {
	switch e {
	case FileTypeExperiment, FileTypeEngine, FileTypeFault, FileTypeCsv:
		return true
	}
	return false
}

func (e FileType) String() string {
	return string(e)
}

func (e *FileType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = FileType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid FileType", str)
	}
	return nil
}

func (e FileType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type GamedaySortingField string

const (
	GamedaySortingFieldName GamedaySortingField = "NAME"
	GamedaySortingFieldTime GamedaySortingField = "TIME"
)

var AllGamedaySortingField = []GamedaySortingField{
	GamedaySortingFieldName,
	GamedaySortingFieldTime,
}

func (e GamedaySortingField) IsValid() bool {
	switch e {
	case GamedaySortingFieldName, GamedaySortingFieldTime:
		return true
	}
	return false
}

func (e GamedaySortingField) String() string {
	return string(e)
}

func (e *GamedaySortingField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = GamedaySortingField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid GamedaySortingField", str)
	}
	return nil
}

func (e GamedaySortingField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type InfraScope string

const (
	InfraScopeNamespace InfraScope = "namespace"
	InfraScopeCluster   InfraScope = "cluster"
)

var AllInfraScope = []InfraScope{
	InfraScopeNamespace,
	InfraScopeCluster,
}

func (e InfraScope) IsValid() bool {
	switch e {
	case InfraScopeNamespace, InfraScopeCluster:
		return true
	}
	return false
}

func (e InfraScope) String() string {
	return string(e)
}

func (e *InfraScope) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = InfraScope(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid INFRA_SCOPE", str)
	}
	return nil
}

func (e InfraScope) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Defines the different types of Image Pull Policy
type ImagePullPolicy string

const (
	ImagePullPolicyIfNotPresent ImagePullPolicy = "IfNotPresent"
	ImagePullPolicyAlways       ImagePullPolicy = "Always"
	ImagePullPolicyNever        ImagePullPolicy = "Never"
)

var AllImagePullPolicy = []ImagePullPolicy{
	ImagePullPolicyIfNotPresent,
	ImagePullPolicyAlways,
	ImagePullPolicyNever,
}

func (e ImagePullPolicy) IsValid() bool {
	switch e {
	case ImagePullPolicyIfNotPresent, ImagePullPolicyAlways, ImagePullPolicyNever:
		return true
	}
	return false
}

func (e ImagePullPolicy) String() string {
	return string(e)
}

func (e *ImagePullPolicy) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ImagePullPolicy(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ImagePullPolicy", str)
	}
	return nil
}

func (e ImagePullPolicy) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type InfrastructureType string

const (
	InfrastructureTypeKubernetes   InfrastructureType = "Kubernetes"
	InfrastructureTypeKubernetesV2 InfrastructureType = "KubernetesV2"
	InfrastructureTypeWindows      InfrastructureType = "Windows"
	InfrastructureTypeLinux        InfrastructureType = "Linux"
	InfrastructureTypeCloudFoundry InfrastructureType = "CloudFoundry"
	InfrastructureTypeContainer    InfrastructureType = "Container"
)

var AllInfrastructureType = []InfrastructureType{
	InfrastructureTypeKubernetes,
	InfrastructureTypeKubernetesV2,
	InfrastructureTypeWindows,
	InfrastructureTypeLinux,
	InfrastructureTypeCloudFoundry,
	InfrastructureTypeContainer,
}

func (e InfrastructureType) IsValid() bool {
	switch e {
	case InfrastructureTypeKubernetes, InfrastructureTypeKubernetesV2, InfrastructureTypeWindows, InfrastructureTypeLinux, InfrastructureTypeCloudFoundry, InfrastructureTypeContainer:
		return true
	}
	return false
}

func (e InfrastructureType) String() string {
	return string(e)
}

func (e *InfrastructureType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = InfrastructureType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid InfrastructureType", str)
	}
	return nil
}

func (e InfrastructureType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// InstallationType defines the installation method used by the user
type InstallationType string

const (
	InstallationTypeConnector   InstallationType = "CONNECTOR"
	InstallationTypeManifest    InstallationType = "MANIFEST"
	InstallationTypeConnectorv2 InstallationType = "CONNECTORV2"
)

var AllInstallationType = []InstallationType{
	InstallationTypeConnector,
	InstallationTypeManifest,
	InstallationTypeConnectorv2,
}

func (e InstallationType) IsValid() bool {
	switch e {
	case InstallationTypeConnector, InstallationTypeManifest, InstallationTypeConnectorv2:
		return true
	}
	return false
}

func (e InstallationType) String() string {
	return string(e)
}

func (e *InstallationType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = InstallationType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid InstallationType", str)
	}
	return nil
}

func (e InstallationType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Defines the different modes of Probes
type Mode string

const (
	ModeSot        Mode = "SOT"
	ModeEot        Mode = "EOT"
	ModeEdge       Mode = "Edge"
	ModeContinuous Mode = "Continuous"
	ModeOnChaos    Mode = "OnChaos"
)

var AllMode = []Mode{
	ModeSot,
	ModeEot,
	ModeEdge,
	ModeContinuous,
	ModeOnChaos,
}

func (e Mode) IsValid() bool {
	switch e {
	case ModeSot, ModeEot, ModeEdge, ModeContinuous, ModeOnChaos:
		return true
	}
	return false
}

func (e Mode) String() string {
	return string(e)
}

func (e *Mode) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Mode(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Mode", str)
	}
	return nil
}

func (e Mode) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Operator string

const (
	OperatorEqualTo    Operator = "EQUAL_TO"
	OperatorNotEqualTo Operator = "NOT_EQUAL_TO"
)

var AllOperator = []Operator{
	OperatorEqualTo,
	OperatorNotEqualTo,
}

func (e Operator) IsValid() bool {
	switch e {
	case OperatorEqualTo, OperatorNotEqualTo:
		return true
	}
	return false
}

func (e Operator) String() string {
	return string(e)
}

func (e *Operator) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Operator(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Operator", str)
	}
	return nil
}

func (e Operator) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Probe Sorting Field
type ProbeSortingField string

const (
	ProbeSortingFieldTime    ProbeSortingField = "TIME"
	ProbeSortingFieldEnabled ProbeSortingField = "ENABLED"
)

var AllProbeSortingField = []ProbeSortingField{
	ProbeSortingFieldTime,
	ProbeSortingFieldEnabled,
}

func (e ProbeSortingField) IsValid() bool {
	switch e {
	case ProbeSortingFieldTime, ProbeSortingFieldEnabled:
		return true
	}
	return false
}

func (e ProbeSortingField) String() string {
	return string(e)
}

func (e *ProbeSortingField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ProbeSortingField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ProbeSortingField", str)
	}
	return nil
}

func (e ProbeSortingField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Defines the different statuses of Probes
type ProbeStatus string

const (
	ProbeStatusRunning   ProbeStatus = "Running"
	ProbeStatusCompleted ProbeStatus = "Completed"
	ProbeStatusStopped   ProbeStatus = "Stopped"
	ProbeStatusError     ProbeStatus = "Error"
	ProbeStatusQueued    ProbeStatus = "Queued"
	ProbeStatusNa        ProbeStatus = "NA"
)

var AllProbeStatus = []ProbeStatus{
	ProbeStatusRunning,
	ProbeStatusCompleted,
	ProbeStatusStopped,
	ProbeStatusError,
	ProbeStatusQueued,
	ProbeStatusNa,
}

func (e ProbeStatus) IsValid() bool {
	switch e {
	case ProbeStatusRunning, ProbeStatusCompleted, ProbeStatusStopped, ProbeStatusError, ProbeStatusQueued, ProbeStatusNa:
		return true
	}
	return false
}

func (e ProbeStatus) String() string {
	return string(e)
}

func (e *ProbeStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ProbeStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ProbeStatus", str)
	}
	return nil
}

func (e ProbeStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Defines the different types of Probes
type ProbeType string

const (
	ProbeTypeHTTPProbe      ProbeType = "httpProbe"
	ProbeTypeCmdProbe       ProbeType = "cmdProbe"
	ProbeTypePromProbe      ProbeType = "promProbe"
	ProbeTypeK8sProbe       ProbeType = "k8sProbe"
	ProbeTypeSloProbe       ProbeType = "sloProbe"
	ProbeTypeDatadogProbe   ProbeType = "datadogProbe"
	ProbeTypeDynatraceProbe ProbeType = "dynatraceProbe"
	ProbeTypeApmProbe       ProbeType = "apmProbe"
)

var AllProbeType = []ProbeType{
	ProbeTypeHTTPProbe,
	ProbeTypeCmdProbe,
	ProbeTypePromProbe,
	ProbeTypeK8sProbe,
	ProbeTypeSloProbe,
	ProbeTypeDatadogProbe,
	ProbeTypeDynatraceProbe,
	ProbeTypeApmProbe,
}

func (e ProbeType) IsValid() bool {
	switch e {
	case ProbeTypeHTTPProbe, ProbeTypeCmdProbe, ProbeTypePromProbe, ProbeTypeK8sProbe, ProbeTypeSloProbe, ProbeTypeDatadogProbe, ProbeTypeDynatraceProbe, ProbeTypeApmProbe:
		return true
	}
	return false
}

func (e ProbeType) String() string {
	return string(e)
}

func (e *ProbeType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ProbeType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ProbeType", str)
	}
	return nil
}

func (e ProbeType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Defines the older different statuses of Probes
type ProbeVerdict string

const (
	ProbeVerdictPassed  ProbeVerdict = "Passed"
	ProbeVerdictFailed  ProbeVerdict = "Failed"
	ProbeVerdictNa      ProbeVerdict = "NA"
	ProbeVerdictAwaited ProbeVerdict = "Awaited"
)

var AllProbeVerdict = []ProbeVerdict{
	ProbeVerdictPassed,
	ProbeVerdictFailed,
	ProbeVerdictNa,
	ProbeVerdictAwaited,
}

func (e ProbeVerdict) IsValid() bool {
	switch e {
	case ProbeVerdictPassed, ProbeVerdictFailed, ProbeVerdictNa, ProbeVerdictAwaited:
		return true
	}
	return false
}

func (e ProbeVerdict) String() string {
	return string(e)
}

func (e *ProbeVerdict) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ProbeVerdict(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ProbeVerdict", str)
	}
	return nil
}

func (e ProbeVerdict) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type QuestionType string

const (
	QuestionTypeMcq    QuestionType = "MCQ"
	QuestionTypeNonMcq QuestionType = "Non_MCQ"
)

var AllQuestionType = []QuestionType{
	QuestionTypeMcq,
	QuestionTypeNonMcq,
}

func (e QuestionType) IsValid() bool {
	switch e {
	case QuestionTypeMcq, QuestionTypeNonMcq:
		return true
	}
	return false
}

func (e QuestionType) String() string {
	return string(e)
}

func (e *QuestionType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = QuestionType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid QuestionType", str)
	}
	return nil
}

func (e QuestionType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type RecurrenceType string

const (
	RecurrenceTypeYearly  RecurrenceType = "Yearly"
	RecurrenceTypeMonthly RecurrenceType = "Monthly"
	RecurrenceTypeDaily   RecurrenceType = "Daily"
	RecurrenceTypeWeekly  RecurrenceType = "Weekly"
	RecurrenceTypeNone    RecurrenceType = "None"
)

var AllRecurrenceType = []RecurrenceType{
	RecurrenceTypeYearly,
	RecurrenceTypeMonthly,
	RecurrenceTypeDaily,
	RecurrenceTypeWeekly,
	RecurrenceTypeNone,
}

func (e RecurrenceType) IsValid() bool {
	switch e {
	case RecurrenceTypeYearly, RecurrenceTypeMonthly, RecurrenceTypeDaily, RecurrenceTypeWeekly, RecurrenceTypeNone:
		return true
	}
	return false
}

func (e RecurrenceType) String() string {
	return string(e)
}

func (e *RecurrenceType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = RecurrenceType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid RecurrenceType", str)
	}
	return nil
}

func (e RecurrenceType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ResourceType string

const (
	ResourceTypeGameday    ResourceType = "GAMEDAY"
	ResourceTypeExperiment ResourceType = "EXPERIMENT"
	ResourceTypeProbe      ResourceType = "PROBE"
)

var AllResourceType = []ResourceType{
	ResourceTypeGameday,
	ResourceTypeExperiment,
	ResourceTypeProbe,
}

func (e ResourceType) IsValid() bool {
	switch e {
	case ResourceTypeGameday, ResourceTypeExperiment, ResourceTypeProbe:
		return true
	}
	return false
}

func (e ResourceType) String() string {
	return string(e)
}

func (e *ResourceType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ResourceType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ResourceType", str)
	}
	return nil
}

func (e ResourceType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ScenarioType string

const (
	ScenarioTypeCron    ScenarioType = "CRON"
	ScenarioTypeNonCron ScenarioType = "NON_CRON"
	ScenarioTypeGameday ScenarioType = "GAMEDAY"
	ScenarioTypeAll     ScenarioType = "ALL"
)

var AllScenarioType = []ScenarioType{
	ScenarioTypeCron,
	ScenarioTypeNonCron,
	ScenarioTypeGameday,
	ScenarioTypeAll,
}

func (e ScenarioType) IsValid() bool {
	switch e {
	case ScenarioTypeCron, ScenarioTypeNonCron, ScenarioTypeGameday, ScenarioTypeAll:
		return true
	}
	return false
}

func (e ScenarioType) String() string {
	return string(e)
}

func (e *ScenarioType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ScenarioType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ScenarioType", str)
	}
	return nil
}

func (e ScenarioType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SecurityGovernancePhase string

const (
	SecurityGovernancePhasePassed SecurityGovernancePhase = "Passed"
	SecurityGovernancePhaseFailed SecurityGovernancePhase = "Failed"
)

var AllSecurityGovernancePhase = []SecurityGovernancePhase{
	SecurityGovernancePhasePassed,
	SecurityGovernancePhaseFailed,
}

func (e SecurityGovernancePhase) IsValid() bool {
	switch e {
	case SecurityGovernancePhasePassed, SecurityGovernancePhaseFailed:
		return true
	}
	return false
}

func (e SecurityGovernancePhase) String() string {
	return string(e)
}

func (e *SecurityGovernancePhase) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SecurityGovernancePhase(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SecurityGovernancePhase", str)
	}
	return nil
}

func (e SecurityGovernancePhase) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Task string

const (
	TaskAddQuestion      Task = "ADD_QUESTION"
	TaskAddActionItem    Task = "ADD_ACTION_ITEM"
	TaskAddNotes         Task = "ADD_NOTES"
	TaskUpdateActionItem Task = "UPDATE_ACTION_ITEM"
	TaskUpdateAnswer     Task = "UPDATE_ANSWER"
)

var AllTask = []Task{
	TaskAddQuestion,
	TaskAddActionItem,
	TaskAddNotes,
	TaskUpdateActionItem,
	TaskUpdateAnswer,
}

func (e Task) IsValid() bool {
	switch e {
	case TaskAddQuestion, TaskAddActionItem, TaskAddNotes, TaskUpdateActionItem, TaskUpdateAnswer:
		return true
	}
	return false
}

func (e Task) String() string {
	return string(e)
}

func (e *Task) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Task(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Task", str)
	}
	return nil
}

func (e Task) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Type string

const (
	TypeGameday    Type = "GAMEDAY"
	TypeGamedayRun Type = "GAMEDAY_RUN"
)

var AllType = []Type{
	TypeGameday,
	TypeGamedayRun,
}

func (e Type) IsValid() bool {
	switch e {
	case TypeGameday, TypeGamedayRun:
		return true
	}
	return false
}

func (e Type) String() string {
	return string(e)
}

func (e *Type) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Type(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Type", str)
	}
	return nil
}

func (e Type) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type UpdateCronExperimentAction string

const (
	UpdateCronExperimentActionEnable  UpdateCronExperimentAction = "Enable"
	UpdateCronExperimentActionDisable UpdateCronExperimentAction = "Disable"
	UpdateCronExperimentActionUpdate  UpdateCronExperimentAction = "Update"
)

var AllUpdateCronExperimentAction = []UpdateCronExperimentAction{
	UpdateCronExperimentActionEnable,
	UpdateCronExperimentActionDisable,
	UpdateCronExperimentActionUpdate,
}

func (e UpdateCronExperimentAction) IsValid() bool {
	switch e {
	case UpdateCronExperimentActionEnable, UpdateCronExperimentActionDisable, UpdateCronExperimentActionUpdate:
		return true
	}
	return false
}

func (e UpdateCronExperimentAction) String() string {
	return string(e)
}

func (e *UpdateCronExperimentAction) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UpdateCronExperimentAction(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UpdateCronExperimentAction", str)
	}
	return nil
}

func (e UpdateCronExperimentAction) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// UpdateStatus represents if infra needs to be updated
type UpdateStatus string

const (
	UpdateStatusAvailable   UpdateStatus = "AVAILABLE"
	UpdateStatusMandatory   UpdateStatus = "MANDATORY"
	UpdateStatusNotRequired UpdateStatus = "NOT_REQUIRED"
)

var AllUpdateStatus = []UpdateStatus{
	UpdateStatusAvailable,
	UpdateStatusMandatory,
	UpdateStatusNotRequired,
}

func (e UpdateStatus) IsValid() bool {
	switch e {
	case UpdateStatusAvailable, UpdateStatusMandatory, UpdateStatusNotRequired:
		return true
	}
	return false
}

func (e UpdateStatus) String() string {
	return string(e)
}

func (e *UpdateStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UpdateStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UpdateStatus", str)
	}
	return nil
}

func (e UpdateStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// UpgradeStatus represents the state of the infra upgraded
type UpgradeStatus string

const (
	UpgradeStatusUpgradingInfra    UpgradeStatus = "UPGRADING_INFRA"
	UpgradeStatusUpgradeSkipped    UpgradeStatus = "UPGRADE_SKIPPED"
	UpgradeStatusUpgradeSuccessful UpgradeStatus = "UPGRADE_SUCCESSFUL"
	UpgradeStatusUpgradeFailed     UpgradeStatus = "UPGRADE_FAILED"
	UpgradeStatusDetectingUpgrader UpgradeStatus = "DETECTING_UPGRADER"
	UpgradeStatusUpgraderDisabled  UpgradeStatus = "UPGRADER_DISABLED"
)

var AllUpgradeStatus = []UpgradeStatus{
	UpgradeStatusUpgradingInfra,
	UpgradeStatusUpgradeSkipped,
	UpgradeStatusUpgradeSuccessful,
	UpgradeStatusUpgradeFailed,
	UpgradeStatusDetectingUpgrader,
	UpgradeStatusUpgraderDisabled,
}

func (e UpgradeStatus) IsValid() bool {
	switch e {
	case UpgradeStatusUpgradingInfra, UpgradeStatusUpgradeSkipped, UpgradeStatusUpgradeSuccessful, UpgradeStatusUpgradeFailed, UpgradeStatusDetectingUpgrader, UpgradeStatusUpgraderDisabled:
		return true
	}
	return false
}

func (e UpgradeStatus) String() string {
	return string(e)
}

func (e *UpgradeStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UpgradeStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UpgradeStatus", str)
	}
	return nil
}

func (e UpgradeStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type WorkflowRunStatus string

const (
	WorkflowRunStatusAll                       WorkflowRunStatus = "All"
	WorkflowRunStatusRunning                   WorkflowRunStatus = "Running"
	WorkflowRunStatusCompleted                 WorkflowRunStatus = "Completed"
	WorkflowRunStatusCompletedWithError        WorkflowRunStatus = "Completed_With_Error"
	WorkflowRunStatusCompletedWithProbeFailure WorkflowRunStatus = "Completed_With_Probe_Failure"
	WorkflowRunStatusStopped                   WorkflowRunStatus = "Stopped"
	WorkflowRunStatusSkipped                   WorkflowRunStatus = "Skipped"
	WorkflowRunStatusError                     WorkflowRunStatus = "Error"
	WorkflowRunStatusTimeout                   WorkflowRunStatus = "Timeout"
	WorkflowRunStatusNa                        WorkflowRunStatus = "NA"
	WorkflowRunStatusQueued                    WorkflowRunStatus = "Queued"
	WorkflowRunStatusBlocked                   WorkflowRunStatus = "Blocked"
)

var AllWorkflowRunStatus = []WorkflowRunStatus{
	WorkflowRunStatusAll,
	WorkflowRunStatusRunning,
	WorkflowRunStatusCompleted,
	WorkflowRunStatusCompletedWithError,
	WorkflowRunStatusCompletedWithProbeFailure,
	WorkflowRunStatusStopped,
	WorkflowRunStatusSkipped,
	WorkflowRunStatusError,
	WorkflowRunStatusTimeout,
	WorkflowRunStatusNa,
	WorkflowRunStatusQueued,
	WorkflowRunStatusBlocked,
}

func (e WorkflowRunStatus) IsValid() bool {
	switch e {
	case WorkflowRunStatusAll, WorkflowRunStatusRunning, WorkflowRunStatusCompleted, WorkflowRunStatusCompletedWithError, WorkflowRunStatusCompletedWithProbeFailure, WorkflowRunStatusStopped, WorkflowRunStatusSkipped, WorkflowRunStatusError, WorkflowRunStatusTimeout, WorkflowRunStatusNa, WorkflowRunStatusQueued, WorkflowRunStatusBlocked:
		return true
	}
	return false
}

func (e WorkflowRunStatus) String() string {
	return string(e)
}

func (e *WorkflowRunStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = WorkflowRunStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid WorkflowRunStatus", str)
	}
	return nil
}

func (e WorkflowRunStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type WorkflowSortingField string

const (
	WorkflowSortingFieldLastExecuted WorkflowSortingField = "LAST_EXECUTED"
	WorkflowSortingFieldLastModified WorkflowSortingField = "LAST_MODIFIED"
	WorkflowSortingFieldName         WorkflowSortingField = "NAME"
)

var AllWorkflowSortingField = []WorkflowSortingField{
	WorkflowSortingFieldLastExecuted,
	WorkflowSortingFieldLastModified,
	WorkflowSortingFieldName,
}

func (e WorkflowSortingField) IsValid() bool {
	switch e {
	case WorkflowSortingFieldLastExecuted, WorkflowSortingFieldLastModified, WorkflowSortingFieldName:
		return true
	}
	return false
}

func (e WorkflowSortingField) String() string {
	return string(e)
}

func (e *WorkflowSortingField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = WorkflowSortingField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid WorkflowSortingField", str)
	}
	return nil
}

func (e WorkflowSortingField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type WorkflowType string

const (
	WorkflowTypeAll             WorkflowType = "All"
	WorkflowTypeGamedayWorkflow WorkflowType = "GamedayWorkflow"
	WorkflowTypeWorkflow        WorkflowType = "Workflow"
	WorkflowTypeCronWorkflow    WorkflowType = "CronWorkflow"
	WorkflowTypeChaosEngine     WorkflowType = "ChaosEngine"
	WorkflowTypeChaosSchedule   WorkflowType = "ChaosSchedule"
)

var AllWorkflowType = []WorkflowType{
	WorkflowTypeAll,
	WorkflowTypeGamedayWorkflow,
	WorkflowTypeWorkflow,
	WorkflowTypeCronWorkflow,
	WorkflowTypeChaosEngine,
	WorkflowTypeChaosSchedule,
}

func (e WorkflowType) IsValid() bool {
	switch e {
	case WorkflowTypeAll, WorkflowTypeGamedayWorkflow, WorkflowTypeWorkflow, WorkflowTypeCronWorkflow, WorkflowTypeChaosEngine, WorkflowTypeChaosSchedule:
		return true
	}
	return false
}

func (e WorkflowType) String() string {
	return string(e)
}

func (e *WorkflowType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = WorkflowType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid WorkflowType", str)
	}
	return nil
}

func (e WorkflowType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
