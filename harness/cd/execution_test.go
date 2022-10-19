package cd

import (
	"github.com/harness/harness-go-sdk/harness/cd/graphql"
	"github.com/stretchr/testify/require"
	"testing"
	time2 "time"
)

func TestExportExecutions(t *testing.T) {
	c := getClient()

	time1 := time2.Date(2022, 10, 7, 0, 0, 0, 0, time2.UTC)

	input := &graphql.ExportExecutionsInput{
		Filters: []*graphql.ExecutionFilter{
			{
				Pipeline: &graphql.IdFilter{
					Operator: graphql.IdOperatorTypes.Equals,
					Values:   []string{"dMybjgkpSOGeul6mJmsx5w"},
				},
			},
			{
				Tag: &graphql.DeploymentTagFilter{
					EntityType: graphql.DeploymentTagTypes.Service,
					Tags: []graphql.DeploymentTag{
						{
							Name:  "tag",
							Value: "tag",
						},
					},
				},
			},
			{
				StartTime: &graphql.TimeFilter{
					Operator:    graphql.TimeOperatorTypes.After,
					ValueMillis: time1.UnixMilli(),
				},
			},
		},
	}

	res, err := c.ExecutionClient.ExportExecutions(input)
	require.NoError(t, err)
	require.NotNil(t, res.DownloadLink)
	require.Exactly(t, res.ErrorMessage, "")
}

func TestGetWorkflowExecutionById(t *testing.T) {
	c := getClient()
	exec, err := c.ExecutionClient.GetWorkflowExecutionById("9nBH0M7dRyy_R9c-gf39sw")
	require.NoError(t, err)
	require.NotNil(t, exec)
}

func TestGetPipelineExecutionById(t *testing.T) {
	c := getClient()
	exec, err := c.ExecutionClient.GetPipelineExecutionById("mF3QM41vR5iqExZZMH3SjQ")
	require.NoError(t, err)
	require.NotNil(t, exec)
}

func TestStartWorkflowExecution(t *testing.T) {
	c := getClient()

	input := &graphql.StartExecutionInput{
		ApplicationId: "J4PP3exRS1C0XuH-BBYkNA",
		ExecutionType: graphql.ExecutionTypes.Workflow,
		EntityId:      "fxMFZE3ZQICV2BVfWTccpA",
		VariableInputs: []*graphql.VariableInput{
			{
				Name: "Environment",
				VariableValue: &graphql.VariableValue{
					Type:  graphql.VariableValueTypes.Name,
					Value: "dev",
				},
			},
			{
				Name: "Service",
				VariableValue: &graphql.VariableValue{
					Type:  graphql.VariableValueTypes.Name,
					Value: "nginx",
				},
			},
			{
				Name: "InfraDefinition",
				VariableValue: &graphql.VariableValue{
					Type:  graphql.VariableValueTypes.Name,
					Value: "k8s-dev",
				},
			},
		},
		ServiceInputs: []*graphql.ServiceInput{
			{
				Name: "nginx",
				ArtifactValueInput: &graphql.ArtifactValueInput{
					BuildNumber: &graphql.BuildNumberInput{
						ArtifactSourceName: "library_nginx",
						BuildNumber:        "latest",
					},
					ValueType: graphql.ArtifactInputTypes.BuildNumber,
				},
			},
		},
	}

	exec, err := c.ExecutionClient.StartExecution(input)
	require.NoError(t, err)
	require.NotNil(t, exec)
	require.NotNil(t, exec.WorkflowExecution)
	require.NotEmpty(t, exec.WorkflowExecution.Id)
	require.Nil(t, exec.PipelineExecution)
}

func TestStartPipelineExecution(t *testing.T) {
	c := getClient()

	input := &graphql.StartExecutionInput{
		ApplicationId: "J4PP3exRS1C0XuH-BBYkNA",
		ExecutionType: graphql.ExecutionTypes.Pipeline,
		EntityId:      "mrbA8gUfTc6luH6GUj0yag",
		VariableInputs: []*graphql.VariableInput{
			{
				Name: "Service",
				VariableValue: &graphql.VariableValue{
					Type:  graphql.VariableValueTypes.Name,
					Value: "nginx",
				},
			},
		},
		ServiceInputs: []*graphql.ServiceInput{
			{
				Name: "nginx",
				ArtifactValueInput: &graphql.ArtifactValueInput{
					BuildNumber: &graphql.BuildNumberInput{
						ArtifactSourceName: "library_nginx",
						BuildNumber:        "latest",
					},
					ValueType: graphql.ArtifactInputTypes.BuildNumber,
				},
			},
		},
	}

	exec, err := c.ExecutionClient.StartExecution(input)
	require.NoError(t, err)
	require.NotNil(t, exec)
	require.NotNil(t, exec.PipelineExecution)
	require.NotEmpty(t, exec.PipelineExecution.Id)
	require.Nil(t, exec.WorkflowExecution)
}
